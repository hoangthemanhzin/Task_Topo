package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	"etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/context"

	"etrib5gc/nfs/amf/sbi/producer"
)

type AMF struct {
	producer *producer.Producer  //handling Sbi requests received at the server
	ctx      *context.AmfContext // AMF context
	cfg      *config.AmfConfig   // loaded AMF config
}

func New(cfg *config.AmfConfig) (nf *AMF, err error) {
	_initLog()
	nf = &AMF{
		cfg: cfg,
	}

	// initialize AMF context
	if nf.ctx, err = context.NewAmfContext(cfg); err != nil {
		return
	}

	//create sbi producer
	nf.producer = producer.New(nf.ctx)

	return
}

func (nf *AMF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *AMF) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdmServiceName(nf.ctx.PlmnId()),
		common.PcfServiceName(nf.ctx.PlmnId()),
		common.AusfServiceName(nf.ctx.PlmnId()),
	}
	for _, r := range nf.cfg.RanList {
		names = append(names, common.RanServiceName(nf.ctx.PlmnId(), r))
	}

	for _, slice := range nf.cfg.Slices {
		names = append(names, common.SmfServiceName(nf.ctx.PlmnId(), &slice))
	}

	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *AMF) Start() (err error) {
	return
}

func (nf *AMF) Terminate() {
	nf.ctx.Clean()
}
