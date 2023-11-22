package ausf

import (
	"etrib5gc/sbi/ausf/sor"
	"etrib5gc/sbi/ausf/uea"
	"etrib5gc/sbi/ausf/upu"
)

type Producer interface {
	sor.Producer
	uea.Producer
	upu.Producer
}
