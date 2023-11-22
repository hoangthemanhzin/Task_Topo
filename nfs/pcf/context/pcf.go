package context

import (
	"etrib5gc/nfs/pcf/config"
	"etrib5gc/sbi/models"
)

type PcfContext struct {
	conf   *config.PcfConfig
	uelist PcfUeList
	plmnid models.PlmnId
}

func New(conf *config.PcfConfig) *PcfContext {
	_initLog()
	ret := &PcfContext{
		conf:   conf,
		plmnid: models.PlmnId(conf.PlmnId),
	}
	return ret
}

func (ctx *PcfContext) PlmnId() *models.PlmnId {
	return &ctx.plmnid
}
