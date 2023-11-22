package producer

import (
	"etrib5gc/logctx"
	"etrib5gc/nfs/smf/context"
	"etrib5gc/nfs/smf/upman"
	"etrib5gc/sbi"
	"etrib5gc/sbi/smf/ee"
	"etrib5gc/sbi/smf/nidd"
	"etrib5gc/sbi/smf/pdu"
)

var _logfields logctx.Fields = logctx.Fields{
	"mod": "producer",
}

type Producer struct {
	logctx.LogWriter
	ctx       *context.SmfContext
	upmanager *upman.UpManager
}

func New(ctx *context.SmfContext, upmanager *upman.UpManager) *Producer {
	return &Producer{
		LogWriter: logctx.WithFields(_logfields),
		upmanager: upmanager,
		ctx:       ctx,
	}
}

func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		ee.Service(prod),
		pdu.Service(prod),
		nidd.Service(prod),
	}
}
