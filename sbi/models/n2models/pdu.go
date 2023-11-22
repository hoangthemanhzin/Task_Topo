package n2models

import "etrib5gc/sbi/models"

//FROM PRAN TO AMF
type PduSessResSetRsp struct {
	SuccessList []UlPduSessionResourceInfo
	FailedList  []UlPduSessionResourceInfo
}

//FROM PRAN TO AMF
type PduSessResRelRsp struct {
	Loc  *models.UserLocation
	List []UlPduSessionResourceInfo
}

//FROM PRAN TO AMF
type PduSessResModRsp struct {
	Loc         *models.UserLocation
	SuccessList []UlPduSessionResourceInfo
	FailedList  []UlPduSessionResourceInfo
}

//FROM PRAN TO AMF
type PduSessResNot struct {
	Loc          *models.UserLocation
	NotifyList   []UlPduSessionResourceInfo
	ReleasedList []UlPduSessionResourceInfo
}

//FROM PRAN TO AMF
type PduSessResModInd struct {
	ModifyList []UlPduSessionResourceInfo
}

//FROM AMF TO PRAN
// Pdu session resource setup request
type PduSessResSetReq struct {
	NasPdu      []byte
	SessionList []DlPduSessionResourceInfo
	UeAmbr      *UeAmbr //NOTE: from Subscribed AmPolicy
}

//FROM AMF TO PRAN
// Pdu session resource release command
type PduSessResRelCmd struct {
	NasPdu      []byte
	SessionList []DlPduSessionResourceInfo
}

//FROM AMF TO PRAN
// Pdu session resource modify request
type PduSessResModReq struct {
	NasPdu      []byte
	SessionList []DlPduSessionResourceInfo
}

//FROM AMF TO PRAN
// Pdu session resource modify confirm
type PduSessResModCfm struct {
	NasPdu        []byte
	ConfirmedList []DlPduSessionResourceInfo
	FailedList    []DlPduSessionResourceInfo
}
