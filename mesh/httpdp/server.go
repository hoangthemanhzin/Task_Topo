package httpdp

import (
	"etrib5gc/sbi"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type HttpRoute struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type HttpRoutes []HttpRoute

type HttpService struct {
	Group  string
	Routes []HttpRoute
}

// httpServer
type httpServer struct {
	config *ServerConfig
	server *http.Server
	wg     sync.WaitGroup
}

type ServerConfig struct {
	Ip   net.IP
	Port int
}

func NewHttpServer(config *ServerConfig) *httpServer {
	_initLog()
	ret := &httpServer{
		config: config,
	}
	//router := gin.Default()
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

	//addr := fmt.Sprintf("%s:%d", config.Ip.String(), config.Port)
	addr := fmt.Sprintf("%s:%d", config.Ip.String(), config.Port)

	ret.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return ret
}

// create a http server, register services and their handlers
func (srv *httpServer) Register(services []sbi.SbiService) (err error) {
	router := srv.server.Handler.(*gin.Engine)
	for _, s := range services {
		addHttpRoutes(router, s.Group, MakeHttpRoutes(s.Routes, s.Handler))
	}

	return
}

func (s *httpServer) Start() (err error) {
	errch := make(chan error, 1)
	go func() {
		defer s.wg.Done()
		s.wg.Add(1)
		errch <- s.server.ListenAndServe()
		/*
			if s.config.Scheme == "http" {
				if err := s.server.ListenAndServe(); err != nil {
					//log.Errorf("Http server failed to listen", err)
				}
				return
			}

			if err :=s.server.ListenAndServeTLS(s.config.Tls.Pem, s.config.Tls.Key); err != nil {
				//log.Errorf("Http server failed to listen", err)
			}
		*/

	}()
	t := time.NewTimer(100 * time.Millisecond)
	select {
	case <-t.C:
	case err = <-errch:
	}
	if err == nil {
		log.Infof("Sbi server running at %s", s.server.Addr)
	}
	return
}

func (s *httpServer) Terminate() {
	s.server.Close()
	s.wg.Wait()
	log.Info("Sbi server closed")
}

func addHttpRoutes(engine *gin.Engine, groupname string, routes []HttpRoute) *gin.RouterGroup {
	group := engine.Group(groupname)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// IndexHandler is the index handler.
func HttpIndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from EtriB5GC!")
}
