package service

import (
	"etrib5gc/common"
	meshmodels "etrib5gc/mesh/models"
	"etrib5gc/sbi"

	config "etrib5gc/nfs/pran/config"
	"etrib5gc/nfs/pran/context"

	"etrib5gc/nfs/pran/ngap"

	"etrib5gc/nfs/pran/sbi/producer"
)

const (
	NGAP_PORT int = 38412
)

type PRAN struct {
	producer *producer.Producer //handling Sbi requests received at the server
	ngap     *ngap.Server       //ngap server handling Ran connections and ngap messages
	context  *context.CuContext // PRAN context
	cfg      *config.PRanConfig // loaded PRAN config
}

func New(cfg *config.PRanConfig) (nf *PRAN, err error) {
	_initLog()
	nf = &PRAN{
		cfg: cfg,
	}

	// initialize PRAN context
	if nf.context, err = context.NewCuContext(cfg); err != nil {
		return
	}

	//create NGAP server (including a NAS handler)
	ngapcfg := cfg.Ngap
	if ngapcfg == nil {
		ngapcfg = &config.NgapConfig{}
	}

	var iplist []string
	var port int
	ipdict := make(map[string]bool)
	if ngapcfg.IpList != nil {
		for _, ip := range *ngapcfg.IpList {
			ipdict[ip] = true
		}
	}
	for ip, _ := range ipdict {
		iplist = append(iplist, ip)
	}
	if len(iplist) == 0 {
		iplist = []string{"0.0.0.0"}
	}
	port = NGAP_PORT
	if ngapcfg.Port != nil {
		port = *ngapcfg.Port
	}
	nf.ngap = ngap.NewServer(nf.context, iplist, port)

	//create sbi producer
	nf.producer = producer.New(nf.context)

	return
}

func (nf *PRAN) Services() []sbi.SbiService {
	if nf.producer == nil {
		panic("sbi producer has not been created")
	}
	return nf.producer.SbiServices()
}

func (nf *PRAN) SubscribedServices() (services []meshmodels.ServiceName) {
	names := []string{
		common.UdmServiceName(nf.context.PlmnId()),
		common.DamfServiceName(nf.context.PlmnId(), nf.context.Id()),
		common.AusfServiceName(nf.context.PlmnId()),
	}
	for _, amfid := range nf.cfg.AmfList {
		names = append(names, common.AmfServiceName(nf.context.PlmnId(), amfid))
	}
	for _, name := range names {
		services = append(services, meshmodels.ServiceName(name))
	}

	return
}

func (nf *PRAN) Start() (err error) {
	// start ngap server
	err = nf.ngap.Run()
	return
}

func (nf *PRAN) Terminate() {
	nf.context.Clean()
	//stop ngap server
	nf.ngap.Stop()
}
