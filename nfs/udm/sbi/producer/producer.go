package producer

import (
	//	"fmt"

	"etrib5gc/nfs/udm/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/udm/ee"
	"etrib5gc/sbi/udm/mt"
	"etrib5gc/sbi/udm/niddau"
	"etrib5gc/sbi/udm/pp"
	"etrib5gc/sbi/udm/sdm"
	"etrib5gc/sbi/udm/ueau"
	"etrib5gc/sbi/udm/uecm"
)

type Producer struct {
	ctx *context.UdmContext
}

func New(ctx *context.UdmContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
	}
}
func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		ee.Service(prod),
		pp.Service(prod),
		mt.Service(prod),
		sdm.Service(prod),
		ueau.Service(prod),
		uecm.Service(prod),
		niddau.Service(prod),
	}
}
