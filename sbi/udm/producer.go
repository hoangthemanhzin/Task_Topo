package udm

import (
	"etrib5gc/sbi/udm/ee"
	"etrib5gc/sbi/udm/mt"
	"etrib5gc/sbi/udm/niddau"
	"etrib5gc/sbi/udm/pp"
	"etrib5gc/sbi/udm/sdm"
	"etrib5gc/sbi/udm/ueau"
	"etrib5gc/sbi/udm/uecm"
)

type Producer interface {
	ee.Producer
	mt.Producer
	sdm.Producer
	pp.Producer
	ueau.Producer
	uecm.Producer
	niddau.Producer
}
