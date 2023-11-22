package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/ausf/config"
	"etrib5gc/nfs/ausf/context"

	"etrib5gc/nfs/ausf/sbi/producer"
)

type AUSF struct {
	producer *producer.Producer   //handling Sbi requests received at the server
	context  *context.AusfContext // AUSF context
	cfg      *config.AusfConfig   // loaded AUSF config
}

func New(cfg *config.AusfConfig) (nf *AUSF, err error) {
	_initLog()

	nf = &AUSF{
		cfg: cfg,
	}

	// initialize AUSF context
	nf.context = context.New(cfg)

	//create sbi producer
	nf.producer = producer.New(nf.context)

	return
}

func (nf *AUSF) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdrServiceName(nf.context.PlmnId()),
		common.UdmServiceName(nf.context.PlmnId()),
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *AUSF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *AUSF) Start() (err error) {
	return
}

func (nf *AUSF) Terminate() {
}
