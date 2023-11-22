package models

import "github.com/free5gc/ngap/ngapType"

type RadioCapabilityInfoIndication struct {
	RadioCap             string
	RadioCap4PagingNr    string
	RadioCap4PagingEutra string
}

type UEContextReleaseComplete struct {
	RanNodes []ngapType.RecommendedRANNodeItem
	Cells    []ngapType.RecommendedCellItem
	Sessions []int64
}

// TODO: should be a json-like struct
type HandoverRequire struct {
	HandoverType *ngapType.HandoverType
	TargetId     *ngapType.TargetID
	Sessions     *ngapType.PDUSessionResourceListHORqd
	Container    *ngapType.SourceToTargetTransparentContainer
}

type UeLocInfo struct {
	//To be defined
}

const (
	NAS_UPLINK_NAS_TRANSPORT uint8 = iota
	NAS_INITIAL_UE_MESSAGE
	NAS_NON_DELIVERY_INDICATION
)

type SbiNasDownlink struct {
	Code   uint8 //PDU content type
	Pdu    []byte
	UeCuId int64 //at CU
	//CuId       string //CU (aka PRAN) identity
	//may add more attributes
}

type NgapUplink struct {
	Code uint8
}
type NgapInitialContextSetupRequest struct {
	CoreNgapId int64 //at CU
	NasPdu     []byte
}
type NgapUeContextReleaseCommand struct {
	CoreNgapId int64 //at CU
}
type NgapPduSessionResourceSetupRequest struct {
	CoreNgapId int64 //at CU
}
type NgapPduSessionResourceModifyRequest struct {
	CoreNgapId int64 //at CU
}

type NgapPduSessionResourceReleaseCommand struct {
	CoreNgapId int64 //at CU
}
type NgapRerouteNasRequestCommand struct {
	CoreNgapId int64 //at CU
}
