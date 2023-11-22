package producer

import (
	//	"fmt"

	"etrib5gc/nfs/damf/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/amf/comm"
	"etrib5gc/sbi/amf/ee"
	"etrib5gc/sbi/amf/ran"
)

const (
	T3502_DURATION = 100 //miliseconds
)

type Producer struct {
	ctx *context.DamfContext
}

func New(ctx *context.DamfContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
	}
}

func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		comm.Service(prod),
		ee.Service(prod),
		ran.DamfService(prod),
	}
}
