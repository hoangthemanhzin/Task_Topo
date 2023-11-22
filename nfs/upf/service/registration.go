package service

import (
	"etrib5gc/common"
	"etrib5gc/sbi/models/n42"
	"etrib5gc/sbi/upmf/upmf2upf"
	"time"
)

func (nf *UPF) registerLoop() {
	defer nf.wg.Done()
	defer log.Info("quit registration loop")
	for {
		//do registering here
		if err := nf.register(); err == nil {
			return
		} else {
			log.Errorf("Registration error: %+v", err)
		}
		t := time.NewTimer(5 * time.Second)
		select {
		case <-t.C:
			continue
		case <-nf.quit:
			return
		}
	}
}

func (nf *UPF) register() (err error) {
	req := n42.RegistrationRequest{
		UpfId:   nf.cfg.Id,
		Ip:      common.GetLocalIP().String(),
		Slices:  nf.cfg.Slices,
		SbiPort: nf.cfg.SbiPort,
		Time:    time.Now().UnixNano(),
		Infs:    make(map[string]n42.NetInfConfig),
	}
	addInfs(req, nf)
	var rsp *n42.RegistrationResponse
	if rsp, err = upmf2upf.Register(nf.upmfcli, req); err == nil {
		log.Infof("Registration success: %v", *rsp)
	}
	return
}

// add Infor Infs RegistrationRequest :
func addInfs(req n42.RegistrationRequest, nf *UPF) {
	n42DnnInfoConfig := n42.DnnInfoConfig{
		Cidr: "",
	}
	for _, value := range nf.cfg.DnnList {
		n42DnnInfoConfig.Cidr = value.Cidr
		if value.Dnn != "" {
			req.Infs[value.Dnn] = n42.NetInfConfig{
				Addr:    value.Addr,
				DnnInfo: &n42DnnInfoConfig,
			}
		}
	}
	for _, ifValue := range nf.cfg.IfList {
		if ifValue.Name != "" {
			req.Infs[ifValue.Name] = n42.NetInfConfig{
				Addr: ifValue.Ip,
			}
		}
	}
}
