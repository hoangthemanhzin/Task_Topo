package producer

import (
	//	"fmt"

	"etrib5gc/nfs/udr/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/udr/dr"
	"etrib5gc/sbi/udr/group"
)

type Producer struct {
	ctx *context.UdrContext
}

func New(ctx *context.UdrContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
	}
}
func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		group.Service(prod),
		dr.Service(prod),
	}
}
