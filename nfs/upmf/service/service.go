package service

import (
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/upmf/config"
	"etrib5gc/nfs/upmf/context"

	"etrib5gc/nfs/upmf/topo"

	"etrib5gc/nfs/upmf/sbi/producer"
)

type UPMF struct {
	producer *producer.Producer   //handling Sbi requests received at the server
	ctx      *context.UpmfContext // UPMF context
	cfg      *config.UpmfConfig   // loaded UPMF config
	topology *topo.Topo
}

func New(cfg *config.UpmfConfig) (nf *UPMF, err error) {
	_initLog()
	nf = &UPMF{
		cfg: cfg,
	}
	//create context
	nf.ctx = context.New(cfg)
	nf.topology = topo.NewTopo(&cfg.Topo)
	//create sbi producer
	nf.producer = producer.New(nf.ctx, nf.topology)
	return
}

func (nf *UPMF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *UPMF) SubscribedServices() (services []meshmodels.ServiceName) {
	return
}
func (nf *UPMF) Start() (err error) {
	nf.topology.Start()
	return
}

func (nf *UPMF) Terminate() {
	nf.topology.Stop()
}
