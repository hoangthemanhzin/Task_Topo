package producer

import (
	//	"fmt"

	"etrib5gc/nfs/amf/context"
	"etrib5gc/sbi"
	amfcomm "etrib5gc/sbi/amf/comm"
	amfee "etrib5gc/sbi/amf/ee"
	amfloc "etrib5gc/sbi/amf/loc"
	amfmt "etrib5gc/sbi/amf/mt"
	amfran "etrib5gc/sbi/amf/ran"
)

/*
type Gmm interface {
	HandleN2Msg(int64, *n2models.N2Msg) error
	InitUeContext(*models.Callback, *n2models.InitUeContextRequest) error
}
*/
type Producer struct {
	ctx *context.AmfContext
	//	gmm Gmm
}

func New(ctx *context.AmfContext) *Producer {
	_initLog()
	return &Producer{
		ctx: ctx,
		//		gmm: gmm,
	}
}

func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		amfcomm.Service(prod),
		amfran.Service(prod),
		amfee.Service(prod),
		amfloc.Service(prod),
		amfmt.Service(prod),
	}
}
