package service

import (
	"etrib5gc/common"
	"etrib5gc/mesh/httpdp"
	"etrib5gc/sbi"
	"fmt"
	"net"
	"sync"

	"etrib5gc/nfs/upf/config"
	"etrib5gc/nfs/upf/context"

	"etrib5gc/nfs/upf/sbi/producer"
)

const (
	SBI_PORT int = 7888
)

type SbiServer interface {
	Start() error
	Terminate()
	Register([]sbi.SbiService) error
}

type UPF struct {
	producer  *producer.Producer  //handling Sbi requests received at the server
	ctx       *context.UpfContext // UPF context
	cfg       *config.UpfConfig   // loaded UPF config
	upmfcli   sbi.ConsumerClient
	upmfaddr  *config.UpmfAddress
	sbiserver SbiServer
	quit      chan bool
	wg        sync.WaitGroup
}

func New(cfg *config.UpfConfig) (nf *UPF, err error) {
	_initLog()
	nf = &UPF{
		cfg:      cfg,
		quit:     make(chan bool),
		upmfaddr: cfg.Upmf,
	}
	if nf.upmfaddr == nil {
		if nf.upmfaddr, err = config.DefaultUpmfAddress(); err != nil {
			err = common.WrapError("Get default UPMF address failed", err)
			return
		}
	}
	log.Infof("Upmf address: %s", nf.upmfaddr.String())
	nf.upmfcli = httpdp.NewClientWithAddr(nf.upmfaddr.String())

	//create context
	nf.ctx = context.New(cfg, nf.upmfcli)
	//create sbi producer
	nf.producer = producer.New(nf.ctx)

	//create Sbi Server
	ipstr := cfg.Ip
	if len(ipstr) == 0 {
		ipstr = "0.0.0.0"
	}
	var ip net.IP
	if ip = net.ParseIP(ipstr); ip == nil {
		err = fmt.Errorf("ParseIP failed: %s", ipstr)
		return
	}
	port := SBI_PORT
	if cfg.SbiPort > 0 {
		port = cfg.SbiPort
	}

	nf.sbiserver = httpdp.NewHttpServer(&httpdp.ServerConfig{
		Ip:   ip,
		Port: port,
	})
	//register routes to the server
	err = nf.sbiserver.Register(nf.producer.SbiServices())

	return
}

func (nf *UPF) Start() (err error) {
	if err = nf.sbiserver.Start(); err != nil {
		return
	}
	nf.wg.Add(1)
	go nf.registerLoop()
	return
}

func (nf *UPF) Terminate() {
	if nf.sbiserver != nil {
		nf.sbiserver.Terminate()
	}
	close(nf.quit)
	nf.wg.Wait()
}
