package registry

import (
	"bytes"
	"encoding/json"
	"etrib5gc/common"
	"etrib5gc/mesh/models"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	LB_RANDOM uint8 = iota
	LB_ROUND_ROBIN
	LB_LEAST_REQUEST
	K8S_POD_ID string = "K8S_POD_ID"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// represents a matched destination for a traffic routing
type MatchedGroup interface {
	Select() (*Endpoint, error)
}

type Registry struct {
	sman     *ServiceManager      //manage all subscribed services
	services []models.ServiceName //subscribed services
	server   *http.Server         //sbi server
	cli      *http.Client         //http client to the service controller
	wg       sync.WaitGroup
	quit     chan bool
	startsub chan bool //trigger subscribe
	config   Config
	id       string //a generated uuid or a Pod identity
	podId    string //pod identity (from os environment variable)
	hostname string //host name or pod name (in K8s)
	sbiIp    net.IP
	sbiPort  int
	labels   map[string]string //labels attached to this agent
}

func NewRegistry(cfg *Config, sbiIp net.IP, sbiPort int, labels map[string]string, services []models.ServiceName) (reg Registry) {
	_initLog()
	reg = Registry{
		config:   *cfg,
		services: services,
		cli:      &http.Client{}, //TODO: add security
		quit:     make(chan bool),
		startsub: make(chan bool),
		sbiIp:    sbiIp,
		sbiPort:  sbiPort,
		labels:   labels,
		podId:    os.Getenv(K8S_POD_ID),
		hostname: os.Getenv("HOSTNAME"),
		sman:     newServiceManager(),
	}
	if len(reg.hostname) == 0 {
		reg.hostname, _ = os.Hostname()
	}
	if !reg.isK8s() { //not in K8s
		reg.id = uuid.New().String()
	} else {
		reg.id = reg.podId //use podId as agent ID
	}

	if reg.config.Agent == nil {
		reg.config.Agent = DefaultAgentAddress()
	}

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

	router.POST("/endpoint", reg.onEndpointUpdate)
	router.GET("/ping", reg.onPing)
	addr := fmt.Sprintf("%s:%d", reg.config.Agent.Ip.String(), reg.config.Agent.Port)

	reg.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return
}

func (reg *Registry) isK8s() bool {
	return len(reg.podId) > 0
}

func (reg *Registry) Start() (err error) {
	if reg.config.Controller == nil {
		if reg.config.Controller, err = DefaultControllerAddress(); err != nil {
			return
		}
	}

	errch := make(chan error, 1)
	reg.wg.Add(3)
	go func() {
		defer reg.wg.Done()
		errch <- reg.server.ListenAndServe()
	}()
	t := time.NewTimer(100 * time.Millisecond)
	select {
	case <-t.C:
	case err = <-errch:
	}

	if err == nil {
		log.Infof("Agent listening to %s", reg.server.Addr)
		go reg.registerLoop()
		go reg.subscribeLoop()
	}
	return
}

func (reg *Registry) Terminate() {
	reg.server.Close()
	reg.wg.Wait()
	log.Info("Registry closed")
}

func (reg *Registry) registerLoop() {
	defer reg.wg.Done()
	defer log.Debug("quit registration loop")
	for {
		//do registering here
		if err := reg.register(); err == nil {
			close(reg.startsub) //trigger subscribe procedure
			return
		} else {
			log.Errorf(err.Error())
		}
		t := time.NewTimer(5 * time.Second)
		select {
		case <-t.C:
			continue
		case <-reg.quit:
			return
		}
	}
}

func (reg *Registry) subscribeLoop() {
	defer reg.wg.Done()
	defer log.Debug("quit subsribe loop")
	select {
	case <-reg.quit:
		return
	case <-reg.startsub: //registration has completed
		//start subscribe
	}

	for {
		//do subscribe here
		if err := reg.subscribe(); err == nil {
			//exit loop once we have a subscription
			return
		} else {
			log.Errorf(err.Error())
		}
		t := time.NewTimer(5 * time.Second)
		select {
		case <-t.C:
			continue
		case <-reg.quit:
			return
		}
	}
}

// register to the service controller
// should be removed when deploying on Kubernetes as the registration is done by the kublet
func (reg *Registry) register() (err error) {
	localip := common.GetLocalIP()
	msgbody, _ := json.Marshal(&models.RegistrationRequest{
		Id:        reg.id,
		Name:      reg.hostname,
		K8s:       reg.isK8s(),
		Ip:        localip.String(),
		SbiPort:   reg.sbiPort,
		AgentPort: reg.config.Agent.Port,
		Labels:    reg.labels,
	})
	body := bytes.NewBuffer(msgbody)
	url := fmt.Sprintf("http://%s/register", reg.config.Controller.String())
	log.Infof("Send registration: %s", url)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	var rsp *http.Response
	if rsp, err = reg.cli.Do(req); err == nil {
		var rspbody []byte
		defer rsp.Body.Close()
		if rspbody, err = ioutil.ReadAll(rsp.Body); err == nil {
			if rsp.StatusCode == 200 {
				var rspmsg models.RegistrationResponse
				if err = json.Unmarshal(rspbody, &rspmsg); err == nil {
					//reg.id = rspmsg.Id
					//log.Infof("Server send id: %s", rspmsg.Id)
				}
			} else {
				err = fmt.Errorf("Registration failed [%d]: %s (%s)", rsp.StatusCode, rsp.Status, string(rspbody))
			}
		}
	}

	return
}

func (reg *Registry) subscribe() (err error) {
	log.Infof("Subscribe services: %v", reg.services)

	msgbody, _ := json.Marshal(&models.SubscribeRequest{
		Id:       reg.id,
		Services: reg.services,
	})
	body := bytes.NewBuffer(msgbody)
	url := fmt.Sprintf("http://%s/subscribe", reg.config.Controller.String())
	req, _ := http.NewRequest(http.MethodGet, url, body)
	var rsp *http.Response
	if rsp, err = reg.cli.Do(req); err != nil {
		return
	}
	var rspmsg models.SubscribeResponse
	var dat []byte

	if dat, err = ioutil.ReadAll(rsp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(dat, &rspmsg); err != nil {
		return
	}
	reg.sman.initialize(&rspmsg)
	//process the subscribe response message
	//	reg.epman.init(rspmsg.Endpoints)
	//update subscribed services
	//	reg.sman.init(rspmsg.Services)

	return
}
func (reg *Registry) Id() string {
	return reg.id
}
func (reg *Registry) onEndpointUpdate(ctx *gin.Context) {
	log.Tracef("Receive Enpoint Update")
	var err error
	var dat []byte
	var msg models.EndpointUpdates
	//parse message
	if dat, err = ioutil.ReadAll(ctx.Request.Body); err == nil {
		if err = json.Unmarshal(dat, &msg); err == nil {
			reg.sman.update(&msg)
			ctx.JSON(http.StatusOK, map[string]interface{}{"Ok": true})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
}

func (reg *Registry) onPing(ctx *gin.Context) {
	log.Trace("Receive ping")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (reg *Registry) Search(id models.ServiceName, match models.RouteMatch) (MatchedGroup, error) {
	return reg.sman.search(id, match)
}

// retrieve an endpoint with its IP:Port, if not exist, create a new one
func (reg *Registry) GetEndpoint(addr string) (ep *Endpoint) {
	if ep = reg.sman.findEndpointByAddr(addr); ep != nil {
		return
	}
	ep = reg.sman.createEndpointWithAddr(addr)
	return
}
