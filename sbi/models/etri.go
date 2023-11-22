package models

type PduSessionInfo struct {
	Id       int32
	Snssai   Snssai
	N1SmMsg  []byte
	N2SmInfo []byte
}

type PDUSessionResourceInfo struct {
	Id       int64
	Transfer string
}
type PDUSessionResourceSetupResponse struct {
	OkList     []PDUSessionResourceInfo
	FailedList []PDUSessionResourceInfo
}
type PDUSessionResourceReleaseResponse []PDUSessionResourceInfo

type PDUSessionResourceModifyResponse struct {
	OkList     []PDUSessionResourceInfo
	FailedList []PDUSessionResourceInfo
}

type PDUSessionResourceNotify struct {
	NotifyList   []PDUSessionResourceInfo
	ReleasedList []PDUSessionResourceInfo
}

type PDUSessionResourceModifyIndication []PDUSessionResourceInfo

type InitialContextSetupFailure []PDUSessionResourceInfo

type InitialContextSetupResponse struct {
	OkList     []PDUSessionResourceInfo
	FailedList []PDUSessionResourceInfo
}

type UEContextReleaseRequest []int64
