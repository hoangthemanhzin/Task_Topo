package n2models

import "etrib5gc/sbi/models"

type AuthUeCtx struct {
	Supi     string
	Kamf     []byte
	Rand     []byte
	AuthType models.AuthType
	PlmnId   models.PlmnId
	Eap      string
	Success  bool
}

type DlPduSessionResourceInfo struct {
	Id       int64 //pdu session id
	NasPdu   []byte
	Snssai   models.Snssai
	Transfer []byte
}

type UlPduSessionResourceInfo struct {
	Id       int64 //pdu session id
	Transfer []byte
}

// recommended ran node
type RecRanNode struct {
	//TODO
}

// recommended cell
type RecCell struct {
	//TODO
}

// core network assistance information
type CoreNetAssistInfo struct {
	//TODO
}

// Emergency Fallback Indicator
type EmergFlblkInd struct {
	//TODO
}

type MobilityRestrictionList struct {
	//TODO
}

type Guami struct {
	PlmnId models.PlmnId
	AmfId  string
}

type UeId struct {
	CuId string
	Id   int64
}

type Cause struct {
	Present uint8
	Value   uint8
}

type CritDiag struct {
	//TODO
}

type PagingPriority struct {
	//TODO
}

type UeRadCap4Paging struct {
	//TODO
}

type UeAmbr struct {
	Ul int64
	Dl int64
}

type RrcInactiveAssistanceInfo struct {
	//TODO
}

// ue security cabability
type UeSecCap struct {
	Nr    *SecCap
	Eutra *SecCap
}

type SecCap struct {
	Enc [2]byte
	Int [2]byte
}
