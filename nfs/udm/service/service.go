package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/udm/config"
	"etrib5gc/nfs/udm/context"

	"etrib5gc/nfs/udm/sbi/producer"
)

type UDM struct {
	producer *producer.Producer  //handling Sbi requests received at the server
	context  *context.UdmContext // UDM context
	cfg      *config.UdmConfig   // loaded UDM config
}

func New(cfg *config.UdmConfig) (nf *UDM, err error) {
	_initLog()
	nf = &UDM{
		cfg: cfg,
	}

	//create context
	nf.context = context.New(cfg)
	//create sbi producer
	nf.producer = producer.New(nf.context)

	return
}

func (nf *UDM) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *UDM) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdrServiceName(nf.context.PlmnId()),
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *UDM) Start() (err error) {
	return
}

func (nf *UDM) Terminate() {
}
