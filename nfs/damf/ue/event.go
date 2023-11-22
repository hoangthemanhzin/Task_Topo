package ue

import "etrib5gc/common"

const (
	NAS_INIT_UE uint8 = iota
	NAS_UL_TRANSPORT
	NAS_NON_DELIVERY
)

func (ue *UeContext) HandleSbi(ev *common.EventData) (err error) {
	err = ue.sendEvent(SbiNasEvent, ev)
	return
}
