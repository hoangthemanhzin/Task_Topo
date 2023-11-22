package models

/*

const (
	//update actions
	UPDATE_SM_CONTEXT_REL_DUP uint8 = iota
	UPDATE_SM_CONTEXT_N1N2
	UPDATE_SM_CONTEXT_UPCNXSTATE
	UPDATE_SM_CONTEXT_HOSTATE
	UPDATE_SM_CONTEXT_AN_CHANGE
)
*/
type SmContextUpdateData struct {
	//	Procedure	uint8 //to clearly differentiate update actions

	Pei string `json:"pei,omitempty"`

	ServingNfId string `json:"servingNfId,omitempty"`

	Guami Guami `json:"guami,omitempty"`

	ServingNetwork PlmnIdNid `json:"servingNetwork,omitempty"`

	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	AnType AccessType `json:"anType,omitempty"`

	AdditionalAnType AccessType `json:"additionalAnType,omitempty"`

	AnTypeToReactivate AccessType `json:"anTypeToReactivate,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	PresenceInLadn PresenceState `json:"presenceInLadn,omitempty"`

	UeLocation UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation UserLocation `json:"addUeLocation,omitempty"`

	UpCnxState UpCnxState `json:"upCnxState,omitempty"`

	HoState HoState `json:"hoState,omitempty"`

	ToBeSwitched bool `json:"toBeSwitched,omitempty"`

	FailedToBeSwitched bool `json:"failedToBeSwitched,omitempty"`

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`

	TargetId *NgRanTargetId `json:"targetId,omitempty"`

	TargetServingNfId string `json:"targetServingNfId,omitempty"`

	SmContextStatusUri string `json:"smContextStatusUri,omitempty"`

	DataForwarding bool `json:"dataForwarding,omitempty"`

	N9ForwardingTunnel TunnelInfo `json:"n9ForwardingTunnel,omitempty"`

	N9DlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`

	N9UlForwardingTnlList []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`

	EpsBearerSetup []string `json:"epsBearerSetup,omitempty"`

	RevokeEbiList []int32 `json:"revokeEbiList,omitempty"`

	Release bool `json:"release,omitempty"`

	Cause Cause `json:"cause,omitempty"`

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	Var5gMmCauseValue int32 `json:"5gMmCauseValue,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	EpsInterworkingInd EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`

	AnTypeCanBeChanged bool `json:"anTypeCanBeChanged,omitempty"`

	N2SmInfoExt1 RefToBinaryData `json:"n2SmInfoExt1,omitempty"`

	N2SmInfoTypeExt1 N2SmInfoType `json:"n2SmInfoTypeExt1,omitempty"`

	MaReleaseInd MaReleaseIndication `json:"maReleaseInd,omitempty"`

	MaNwUpgradeInd bool `json:"maNwUpgradeInd,omitempty"`

	MaRequestInd bool `json:"maRequestInd,omitempty"`

	ExemptionInd ExemptionInd `json:"exemptionInd,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	MoExpDataCounter MoExpDataCounter `json:"moExpDataCounter,omitempty"`

	ExtendedNasSmTimerInd bool `json:"extendedNasSmTimerInd,omitempty"`

	ForwardingFTeid string `json:"forwardingFTeid,omitempty"`

	ForwardingBearerContexts []string `json:"forwardingBearerContexts,omitempty"`

	DdnFailureSubs DdnFailureSubs `json:"ddnFailureSubs,omitempty"`

	SkipN2PduSessionResRelInd bool `json:"skipN2PduSessionResRelInd,omitempty"`
}
