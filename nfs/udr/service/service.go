package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/udr/config"
	"etrib5gc/nfs/udr/context"

	"etrib5gc/nfs/udr/sbi/producer"
)

type UDR struct {
	producer *producer.Producer  //handling Sbi requests received at the server
	ctx      *context.UdrContext // UDR context
	cfg      *config.UdrConfig   // loaded UDR config
}

func New(cfg *config.UdrConfig) (nf *UDR, err error) {
	_initLog()
	nf = &UDR{
		cfg: cfg,
	}

	// initialize UDR context
	nf.ctx = context.New(cfg)

	//create sbi producer
	nf.producer = producer.New(nf.ctx)
	return
}

func (nf *UDR) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *UDR) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdmServiceName(nf.ctx.PlmnId()),
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *UDR) Start() (err error) {
	return
}

func (nf *UDR) Terminate() {
}
