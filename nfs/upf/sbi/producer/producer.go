package producer

import (
	"etrib5gc/logctx"
	"etrib5gc/nfs/upf/context"
	"etrib5gc/sbi"
	"etrib5gc/sbi/upf/upf2smf"
	"etrib5gc/sbi/upf/upf2upmf"
)

type Producer struct {
	logctx.LogWriter
	ctx *context.UpfContext
}

func New(ctx *context.UpfContext) *Producer {
	return &Producer{
		LogWriter: logctx.WithFields(logctx.Fields{
			"mod": "producer",
		}),
		ctx: ctx,
	}
}
func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		upf2smf.Service(prod),
		upf2upmf.Service(prod),
	}
}
