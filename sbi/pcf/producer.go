package pcf

import (
	"etrib5gc/sbi/pcf/ampc"
	"etrib5gc/sbi/pcf/btdpc"
	"etrib5gc/sbi/pcf/ee"
	"etrib5gc/sbi/pcf/pa"
	"etrib5gc/sbi/pcf/smpc"
	"etrib5gc/sbi/pcf/uepc"
)

type Producer interface {
	pa.Producer
	ee.Producer
	ampc.Producer
	smpc.Producer
	uepc.Producer
	btdpc.Producer
}
