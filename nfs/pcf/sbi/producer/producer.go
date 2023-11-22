package producer

import (
	//	"fmt"

	"etrib5gc/nfs/pcf/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/amf/ee"
	"etrib5gc/sbi/pcf/ampc"
	"etrib5gc/sbi/pcf/btdpc"
	"etrib5gc/sbi/pcf/pa"
	"etrib5gc/sbi/pcf/smpc"
	"etrib5gc/sbi/pcf/uepc"
)

type Producer struct {
	ctx *context.PcfContext
}

func New(ctx *context.PcfContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
	}
}
func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		ee.Service(prod),
		pa.Service(prod),
		ampc.Service(prod),
		smpc.Service(prod),
		uepc.Service(prod),
		btdpc.Service(prod),
	}
}
