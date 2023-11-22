package producer

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/pran/nas"

	"etrib5gc/nfs/pran/ue"
)

type CuContext interface {
	FindByCuNgapId(int64) *ue.UeContext
}

type Producer struct {
	ctx CuContext
}

func New(c CuContext) *Producer {
	_initLog()
	return &Producer{
		ctx: c,
	}
}

func (prod *Producer) SbiServices() []sbi.SbiService {
	return []sbi.SbiService{
		nas.Service(prod),
	}
}
