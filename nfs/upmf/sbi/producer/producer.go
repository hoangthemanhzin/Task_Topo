package producer

import (
	"etrib5gc/logctx"
	"etrib5gc/nfs/upmf/context"
	"etrib5gc/nfs/upmf/topo"
	"etrib5gc/sbi"
	"etrib5gc/sbi/upmf/upmf2activate"
	"etrib5gc/sbi/upmf/upmf2fe"
	"etrib5gc/sbi/upmf/upmf2smf"
	"etrib5gc/sbi/upmf/upmf2upf"
)

type Producer struct {
	logctx.LogWriter
	ctx *context.UpmfContext
	//topo context.UpfTopo
	topology *topo.Topo
}

func New(ctx *context.UpmfContext, topology *topo.Topo) *Producer {
	return &Producer{
		LogWriter: logctx.WithFields(logctx.Fields{
			"mod": "producer",
		}),
		ctx:      ctx,
		topology: topology,
	}
}
func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		upmf2smf.Service(prod),
		upmf2upf.Service(prod),
		upmf2fe.Service(prod),
		upmf2activate.Service(prod),
	}
}
