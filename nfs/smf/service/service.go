package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/pfcp"
	"etrib5gc/pfcp/pfcptypes"
	"etrib5gc/sbi"

	"etrib5gc/nfs/smf/config"
	"etrib5gc/nfs/smf/context"
	"etrib5gc/nfs/smf/upman"

	"etrib5gc/nfs/smf/sbi/producer"
)

type SMF struct {
	producer  *producer.Producer  //handling Sbi requests received at the server
	ctx       *context.SmfContext // SMF context
	cfg       *config.SmfConfig   // loaded SMF config
	upmanager *upman.UpManager    //for creating tunnel
	pfcp      pfcp.PfcpSender     //for sending PFCP message
}

func New(cfg *config.SmfConfig) (nf *SMF, err error) {
	_initLog()
	nf = &SMF{
		cfg: cfg,
	}

	nodeid := pfcptypes.NodeID{
		NodeIdType: pfcptypes.NodeIdTypeIpv4Address,
		IP:         common.GetLocalIP(),
	}
	nf.pfcp = pfcp.NewSbiPfcp(&nodeid)

	if nf.upmanager, err = upman.NewUpManager(nf.pfcp); err != nil {
		log.Errorf(err.Error())
		return
	}
	nf.ctx = context.New(cfg)
	//create sbi producer
	nf.producer = producer.New(nf.ctx, nf.upmanager)

	return
}

func (nf *SMF) Services() []sbi.SbiService {
	return nf.producer.SbiServices()
}

func (nf *SMF) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdmServiceName(nf.ctx.PlmnId()),
		common.PcfServiceName(nf.ctx.PlmnId()),
		common.UpmfServiceName(nf.ctx.PlmnId()),
	}

	for _, amfid := range nf.cfg.AmfList {
		names = append(names, common.AmfServiceName(nf.ctx.PlmnId(), amfid))
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}
	return
}

func (nf *SMF) Start() (err error) {
	/*
		nf.upmanager.Start()
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
	*/
	return
}
func (nf *SMF) Terminate() {
	/*
		nf.ctx.Clean()
		nf.upmanager.Stop()
	*/
}
