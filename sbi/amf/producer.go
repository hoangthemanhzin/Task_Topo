package amf

import (
	"etrib5gc/sbi/amf/comm"
	"etrib5gc/sbi/amf/ee"
	"etrib5gc/sbi/amf/loc"
	"etrib5gc/sbi/amf/mt"
	"etrib5gc/sbi/amf/ran"
)

type Producer interface {
	comm.Producer
	ee.Producer
	loc.Producer
	mt.Producer
	ran.Producer
}
type DProducer interface {
	ran.DProducer
}
