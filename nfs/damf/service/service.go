package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/damf/config"
	"etrib5gc/nfs/damf/context"

	"etrib5gc/nfs/damf/sbi/producer"
)

type DAMF struct {
	producer *producer.Producer   //handling Sbi requests received at the server
	context  *context.DamfContext // DAMF context
	cfg      *config.DamfConfig   // loaded DAMF config
}

func New(cfg *config.DamfConfig) (nf *DAMF, err error) {
	_initLog()
	nf = &DAMF{
		cfg: cfg,
	}

	//create context
	nf.context = context.NewDamfContext(cfg)
	//create sbi producer
	nf.producer = producer.New(nf.context)

	return
}

func (nf *DAMF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *DAMF) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdmServiceName(nf.context.PlmnId()),
		common.AusfServiceName(nf.context.PlmnId()),
		common.RanServiceName(nf.context.PlmnId(), nf.context.Id()),
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *DAMF) Start() (err error) {
	return
}

func (nf *DAMF) Terminate() {
	//clear all UeContext
	nf.context.Clean()
}
