package mesh

import (
	"etrib5gc/mesh/httpdp"
	"etrib5gc/mesh/models"
	"etrib5gc/mesh/registry"
	"etrib5gc/sbi"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var _mesh *b5gcMesh

type App interface {
	Services() []sbi.SbiService
	SubscribedServices() []models.ServiceName
}

type Running interface {
	Start() error
	Terminate()
}
type DpServer interface {
	Running
	Register([]sbi.SbiService) error
}

type b5gcMesh struct {
	cfg      MeshConfig
	app      App
	regcli   registry.Registry
	server   DpServer
	runnings []Running
}

func newMesh(config *MeshConfig, app App) (m *b5gcMesh, err error) {
	sbi := config.Sbi
	if sbi == nil {
		sbi = DefaultSbiAddress()
	}

	m = &b5gcMesh{
		cfg: *config,
		app: app,
		server: httpdp.NewHttpServer(&httpdp.ServerConfig{
			Ip:   sbi.Ip,
			Port: sbi.Port,
		}),
		regcli: registry.NewRegistry(&config.Registry, sbi.Ip, sbi.Port, config.Labels, app.SubscribedServices()),
	}

	m.cfg.Sbi = sbi

	if err = m.server.Register(app.Services()); err != nil {
		return
	}
	err = m.start()
	return
}

func (m *b5gcMesh) start() (err error) {
	//register services to execute
	services := []Running{}
	services = append(services, m.server)
	services = append(services, &m.regcli)
	//services = append(services, m.reporter)

	//execute service sequentially
	for _, service := range services {
		if err = service.Start(); err != nil {
			m.stop()
			return
		}
		m.runnings = append(m.runnings, service)
	}
	return
}

func (m *b5gcMesh) stop() (err error) {
	for _, service := range m.runnings {
		service.Terminate()
	}
	//	log.Info("mesh is terminated")
	return
}
