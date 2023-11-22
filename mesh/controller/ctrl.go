package controller

import (
	"encoding/json"
	"etrib5gc/mesh/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	EP_LEFT uint8 = iota
	EP_JOIN
)

type Controller struct {
	server *http.Server
	hb     int
	//db          *registryDb
	evch  chan epEvent
	epman *EndpointManager
	sman  *ServiceManager
	wg    sync.WaitGroup
	quit  chan bool
}

type epEvent struct {
	evtype uint8
	dat    interface{}
}

func New(cfg *Config) (ctrl *Controller, err error) {
	_initLog()
	ctrl = &Controller{
		evch: make(chan epEvent, 256),
		quit: make(chan bool),
		hb:   cfg.Heartbeat,
	}
	ctrl.epman = newEndpointManager(ctrl)
	ctrl.sman = newServiceManager(ctrl.epman, cfg.Services)
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host",
			"Token", "X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	router.POST("/register", ctrl.onRegister)
	router.GET("/subscribe", ctrl.onSubscribe)
	addr := fmt.Sprintf("0.0.0.0:%d", CONTROLLER_PORT)
	if cfg.Addr != nil {
		addr = fmt.Sprintf("%s:%d", cfg.Addr.Ip.String(), cfg.Addr.Port)
	}

	ctrl.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return
}

func (c *Controller) Start() (err error) {
	c.wg.Add(1)
	errch := make(chan error, 1)
	go func() {
		defer c.wg.Done()
		errch <- c.server.ListenAndServe()
	}()
	t := time.NewTimer(100 * time.Millisecond)
	select {
	case <-t.C:
	case err = <-errch:
		return
	}

	log.Infof("Controller listen to %s", c.server.Addr)
	go c.epEventLoop()
	return
}

func (c *Controller) Terminate() {
	log.Info("Controller is terminating")
	c.epman.close()
	c.server.Close()
	close(c.quit)
	c.wg.Wait()

	//	c.db.clean()
	close(c.evch)
}
func (c *Controller) sendEpEvent(evtype uint8, ep *Endpoint) {
	c.evch <- epEvent{
		evtype: evtype,
		dat:    ep,
	}
}
func (c *Controller) onRegister(ctx *gin.Context) {
	log.Infof("on registrer")
	var err error
	var dat []byte
	var msg models.RegistrationRequest
	//parse message
	if dat, err = ioutil.ReadAll(ctx.Request.Body); err == nil {
		if err = json.Unmarshal(dat, &msg); err == nil {
			//register the  endpoint
			var ep *Endpoint
			if ep, err = c.epman.registerEndpoint(&msg); err == nil {
				ctx.JSON(http.StatusOK, models.RegistrationResponse{
					Id: ep.id,
				})
				return
			}
		}
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
}

func (c *Controller) onSubscribe(ctx *gin.Context) {
	log.Infof("on subscribe")
	var err error
	var dat []byte
	var msg models.SubscribeRequest
	//parse message
	if dat, err = ioutil.ReadAll(ctx.Request.Body); err == nil {
		if err = json.Unmarshal(dat, &msg); err == nil {
			//add subscibed services for endpoint
			if ep := c.endpointSubscribe(&msg); ep != nil {
				ctx.JSON(http.StatusOK, c.buildSubscribeData(ep))
				return
			} else {
				err = fmt.Errorf("Can't add subscriber")
			}
		}
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)

}

func (c *Controller) epEventLoop() {
	for ev := range c.evch {
		switch ev.evtype {
		case EP_LEFT: // an Endpoint was removed
			if ep, ok := ev.dat.(*Endpoint); ok {
				c.onEpLeft(ep)
			}
		case EP_JOIN:
			//an endpoint joined
			if ep, ok := ev.dat.(*Endpoint); ok {
				c.onEpJoin(ep)
			}
		}
	}
}
func (c *Controller) onEpLeft(ep *Endpoint) {
	log.Warnf("endpoint %s is removed:%v", ep.id, ep.labels)
	//remove endpoint as a subscriber
	ep.unsubscribeServices()
	//notify of endpoint leaving
	subs := c.sman.getSubscribers(ep)
	for _, sub := range subs {
		sub.notifyEpLeft(ep)
	}
}

func (c *Controller) onEpJoin(ep *Endpoint) {
	log.Infof("new endpoint %s joined: %v", ep.id, ep.labels)
	//notify of endpoint joining
	subs := c.sman.getSubscribers(ep)
	for _, sub := range subs {
		sub.notifyEpJoin(ep)
	}
}

// subscribe endpoint
func (c *Controller) endpointSubscribe(msg *models.SubscribeRequest) (ep *Endpoint) {
	//find the endpoint with its identity
	if ep = c.epman.findById(msg.Id); ep == nil {
		log.Errorf("endpoint %s not found", msg.Id)
		return
	}
	//subscribe services for the endpoint
	c.sman.addSubscriber(ep, msg.Services)
	return
}

// compose subscription data for the subsribed endpoint
func (c *Controller) buildSubscribeData(ep *Endpoint) (dat models.SubscribeResponse) {
	dat.Endpoints = make(map[string]models.Endpoint)
	dat.Services = make(map[models.ServiceName]models.Service)
	//get relevant endpoints
	endpoints := c.epman.getEndpoints(ep)
	//add relevant endpoints
	for _, item := range endpoints {
		dat.Endpoints[item.id] = models.Endpoint{
			Id:      item.id,
			Labels:  item.labels,
			Ip:      item.ip.String(),
			SbiPort: item.sbiPort,
		}
	}

	//add service definitions
	for sid, s := range ep.services {
		dat.Services[sid] = models.Service{
			Id:        sid,
			Selectors: s.selectors,
			Routes:    s.routes,
			Groups:    s.groups,
		}
	}
	//TODO: add routes to service
	return
}
