package smf

import (
	"etrib5gc/sbi/smf/ee"
	"etrib5gc/sbi/smf/nidd"
	"etrib5gc/sbi/smf/pdu"
)

type Producer interface {
	pdu.Producer
	ee.Producer
	nidd.Producer
}
