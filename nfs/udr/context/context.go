package context

import (
	"etrib5gc/nfs/udr/config"
	"etrib5gc/sbi/models"
)

type UdrContext struct {
	cfg    *config.UdrConfig
	uelist UdrUeList
	plmnid models.PlmnId
}

func New(cfg *config.UdrConfig) *UdrContext {
	ret := &UdrContext{
		cfg: cfg,
	}
	return ret
}

func (ctx *UdrContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}
