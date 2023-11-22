package context

import (
	"etrib5gc/nfs/upf/config"
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

type UpfContext struct {
	//cfg    *config.UpfConfig
	plmnid  models.PlmnId
	upmfcli sbi.ConsumerClient //for sending request to UPMF
}

func New(cfg *config.UpfConfig, upmfcli sbi.ConsumerClient) *UpfContext {
	_initLog()
	ret := &UpfContext{
		//	cfg:    cfg,
		plmnid:  models.PlmnId(cfg.PlmnId),
		upmfcli: upmfcli,
	}
	return ret
}

func (ctx *UpfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}

func (ctx *UpfContext) UpmfCli() sbi.ConsumerClient {
	return ctx.upmfcli
}
