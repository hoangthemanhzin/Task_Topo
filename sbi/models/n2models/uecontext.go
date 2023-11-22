package n2models

import "etrib5gc/sbi/models"

// Initial Context Setup Request
type InitCtxSetupReq struct {
	PduList  []DlPduSessionResourceInfo
	NasPdu   []byte
	SecKey   []byte
	UeSecCap UeSecCap
	// Old AMF: new amf should get old amf's amf name
	OldAmf       string
	Guami        Guami
	AllowedNssai []models.AllowedSnssai
	UeAmbr       UeAmbr
	UeRadCap     string
	Rfsp         *int64
	//RrcInfo *rrcInactiveAssistanceInfo

	// rrcInactiveTransitionReportRequest: configured by amf
	// This IE is used to request the NG-RAN node to report or stop reporting to the 5GC
	// when the UE enters or leaves RRC_INACTIVE state. (TS 38.413 9.3.1.91)

	// accessType indicate amfUe send this msg for which accessType
	// emergencyFallbackIndicator: configured by amf (TS 23.501 5.16.4.11)
	// coreNetworkAssistanceInfo TS 23.501 5.4.6, 5.4.6.2

	// Mobility Restriction List TS 23.501 5.3.4
	// TS 23.501 5.3.4.1.1: For a given UE, the core network determines the Mobility restrictions
	// based on UE subscription information.
	// TS 38.413 9.3.1.85: This IE defines roaming or access restrictions for subsequent mobility action for
	// which the NR-RAN provides information about the target of the mobility action towards
	// the UE, e.g., handover, or for SCG selection during dual connectivity operation or for
	// assigning proper RNAs. If the NG-RAN receives the Mobility Restriction List IE, it shall
	// overwrite previously received mobility restriction information.

}

//from PRAN to AMF messages
// Initial Context Setup Response
type InitCtxSetupRsp struct {
	SuccessList []UlPduSessionResourceInfo
	FailedList  []UlPduSessionResourceInfo
	Diag        *CritDiag
}

// Initial Context Setup Failure
type InitCtxSetupFailure struct {
	Cause Cause
}

// Ue Context Release Request
// RAN --> AMF
type UeCtxRelReq struct {
	Cause       Cause
	SuccessList []int64
}

// Ue Context Release Command
// AMF --> RAN
type UeCtxRelCmd struct {
	Cause Cause
	PP    *PagingPriority
}

//Ue Context Release Complete
// RAN --> AMF

type UeCtxRelCmpl struct {
	Loc         *models.UserLocation
	RecRanNodes []RecRanNode
	RecCells    []RecCell
	Sessions    []int64
	//TODO: add more attribute
}

// Ue Context Modification Request
// AMF --> RAN
type UeCtxModReq struct {
	PP                 PagingPriority
	SecKey             []byte
	Rfsp               *int64
	UeAmbr             *UeAmbr
	RrcInactTranRepReq *uint64
	CoreAssist         *CoreNetAssistInfo
	Emerg              *EmergFlblkInd
	OldAmfNgapId       *int64
	//TODO: add more attributes
}

// Ue Context Modification Response
// RAN --> AMF
type UeCtxModRsp struct {
	Loc      *models.UserLocation
	RrcState uint16
	//Diag *CritDiag
}

// Ue Context Modification Failure
// RAN --> AMF
type UeCtxModFail struct {
	Cause Cause
	Diag  *CritDiag
}

//RAN --> AMF

type RrcInactTranRep struct {
	Loc      *models.UserLocation
	RrcState uint16
}
