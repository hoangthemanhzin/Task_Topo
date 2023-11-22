package context

import (
	"etrib5gc/nfs/upmf/config"
	"etrib5gc/sbi/models"
)

type UpmfContext struct {
	cfg    *config.UpmfConfig
	plmnid models.PlmnId
}

func New(cfg *config.UpmfConfig) *UpmfContext {
	//_initLog()
	ret := &UpmfContext{
		cfg:    cfg,
		plmnid: models.PlmnId(cfg.PlmnId),
	}
	return ret
}

func (ctx *UpmfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}