package ue

import "etrib5gc/common"

// Event type for Sbi events
const (
	SBI_INIT_UE_STATUS uint8 = iota
	SBI_NAS_DL
	SBI_UE_SET_REQ
	SBI_UE_MOD_REQ
	SBI_UE_REL_CMD
	SBI_PDU_SET_REQ
	SBI_PDU_MOD_REQ
	SBI_PDU_REL_CMD
)

// Event type for Ngap Events
const (
	NGAP_INIT_UE uint8 = iota
	NGAP_UL_NAS
	NGAP_NAS_NON_DELIVERY
	NGAP_UE_SET_RSP
	NGAP_UE_SET_FAIL
	NGAP_UE_RRC_REP
	NGAP_UE_MOD_RSP
	NGAP_UE_MOD_FAIL
	NGAP_PDU_SET_RSP
	NGAP_PDU_REL_RSP
	NGAP_PDU_MOD_RSP
	NGAP_PDU_NOT
	NGAP_PDU_MOD_IND
)

func (uectx *UeContext) HandleNgap(ev *common.EventData) (err error) {
	if err = uectx.sendEvent(NgapEvent, ev); err != nil {
		uectx.Errorf("Handle event failed: %s", err.Error())
	}
	return
}
