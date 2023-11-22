package producer

import (
	"etrib5gc/nfs/ausf/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/ausf/sor"
	"etrib5gc/sbi/ausf/uea"
	"etrib5gc/sbi/ausf/upu"
)

type Producer struct {
	ctx *context.AusfContext
}

func New(ctx *context.AusfContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
	}
}

func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		sor.Service(prod),
		uea.Service(prod),
		upu.Service(prod),
	}
}
