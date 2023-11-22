package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/pcf/config"
	"etrib5gc/nfs/pcf/context"

	"etrib5gc/nfs/pcf/sbi/producer"
)

type PCF struct {
	producer *producer.Producer  //handling Sbi requests received at the server
	ctx      *context.PcfContext // PCF context
	cfg      *config.PcfConfig   // loaded PCF config
}

func New(cfg *config.PcfConfig) (nf *PCF, err error) {
	_initLog()
	nf = &PCF{
		cfg: cfg,
	}

	//create context
	nf.ctx = context.New(cfg)
	//create sbi producer
	nf.producer = producer.New(nf.ctx)

	return
}

func (nf *PCF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *PCF) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{

		common.UdmServiceName(nf.ctx.PlmnId()),
		common.UdrServiceName(nf.ctx.PlmnId()),
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}
func (nf *PCF) Start() (err error) {
	return
}

func (nf *PCF) Terminate() {
}
