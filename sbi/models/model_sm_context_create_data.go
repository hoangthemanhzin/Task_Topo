package models

type SmContextCreateData struct {
	Supi string `json:"supi"`

	UnauthenticatedSupi *bool `json:"unauthenticatedSupi,omitempty"`

	Pei string `json:"pei,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	PduSessionId int32 `json:"pduSessionId"`

	Dnn string `json:"dnn,omitempty"`

	RanNets []string `json:"rannets,omitempty"`

	SelectedDnn string `json:"selectedDnn,omitempty"`

	SNssai Snssai `json:"sNssai"`

	HplmnSnssai *Snssai `json:"hplmnSnssai,omitempty"`

	ServingNfId string `json:"servingNfId"`

	Guami *Guami `json:"guami,omitempty"`

	ServiceName ServiceName `json:"serviceName,omitempty"` //string

	ServingNetwork PlmnIdNid `json:"servingNetwork"`

	RequestType RequestType `json:"requestType,omitempty"` //string

	N1SmMsg RefToBinaryData `json:"n1SmMsg,omitempty"` //string

	AnType AccessType `json:"anType"` //string

	AdditionalAnType AccessType `json:"additionalAnType,omitempty"` //string

	RatType RatType `json:"ratType,omitempty"` //string

	PresenceInLadn PresenceState `json:"presenceInLadn,omitempty"`

	UeLocation *UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`

	SmContextStatusUri string `json:"smContextStatusUri"`

	HSmfUri string `json:"hSmfUri,omitempty"`

	HSmfId string `json:"hSmfId,omitempty"`

	SmfUri string `json:"smfUri,omitempty"`

	SmfId string `json:"smfId,omitempty"`

	AdditionalHsmfUri []string `json:"additionalHsmfUri,omitempty"`

	AdditionalHsmfId []string `json:"additionalHsmfId,omitempty"`

	AdditionalSmfUri []string `json:"additionalSmfUri,omitempty"`

	AdditionalSmfId []string `json:"additionalSmfId,omitempty"`

	OldPduSessionId *int32 `json:"oldPduSessionId,omitempty"`

	PduSessionsActivateList []int32 `json:"pduSessionsActivateList,omitempty"`

	UeEpsPdnConnection string `json:"ueEpsPdnConnection,omitempty"`

	HoState HoState `json:"hoState,omitempty"` //string

	PcfId string `json:"pcfId,omitempty"`

	PcfGroupId string `json:"pcfGroupId,omitempty"`

	PcfSetId string `json:"pcfSetId,omitempty"`

	NrfUri string `json:"nrfUri,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	SelMode DnnSelectionMode `json:"selMode,omitempty"` //string

	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	UdmGroupId string `json:"udmGroupId,omitempty"`

	RoutingIndicator string `json:"routingIndicator,omitempty"`

	EpsInterworkingInd *EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`

	IndirectForwardingFlag *bool `json:"indirectForwardingFlag,omitempty"`

	DirectForwardingFlag *bool `json:"directForwardingFlag,omitempty"`

	TargetId *NgRanTargetId `json:"targetId,omitempty"`

	EpsBearerCtxStatus string `json:"epsBearerCtxStatus,omitempty"`

	CpCiotEnabled *bool `json:"cpCiotEnabled,omitempty"`

	CpOnlyInd *bool `json:"cpOnlyInd,omitempty"`

	InvokeNef *bool `json:"invokeNef,omitempty"`

	MaRequestInd *bool `json:"maRequestInd,omitempty"`

	MaNwUpgradeInd *bool `json:"maNwUpgradeInd,omitempty"`

	N2SmInfo RefToBinaryData `json:"n2SmInfo,omitempty"` //string

	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"` //string

	N2SmInfoExt1 RefToBinaryData `json:"n2SmInfoExt1,omitempty"` //string

	N2SmInfoTypeExt1 N2SmInfoType `json:"n2SmInfoTypeExt1,omitempty"` //string

	SmContextRef string `json:"smContextRef,omitempty"`

	SmContextSmfId string `json:"smContextSmfId,omitempty"`

	SmContextSmfSetId string `json:"smContextSmfSetId,omitempty"`

	SmContextSmfServiceSetId string `json:"smContextSmfServiceSetId,omitempty"`

	SmContextSmfBinding SbiBindingLevel `json:"smContextSmfBinding,omitempty"`

	UpCnxState UpCnxState `json:"upCnxState,omitempty"` //string

	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`

	ApnRateStatus *ApnRateStatus `json:"apnRateStatus,omitempty"`

	ExtendedNasSmTimerInd *bool `json:"extendedNasSmTimerInd,omitempty"`

	DlDataWaitingInd *bool `json:"dlDataWaitingInd,omitempty"`

	DdnFailureSubs *DdnFailureSubs `json:"ddnFailureSubs,omitempty"`

	SmfTransferInd *bool `json:"smfTransferInd,omitempty"`

	OldSmfId string `json:"oldSmfId,omitempty"`

	OldSmContextRef string `json:"oldSmContextRef,omitempty"`

	WAgfInfo *WAgfInfo `json:"wAgfInfo,omitempty"`

	TngfInfo *TngfInfo `json:"tngfInfo,omitempty"`

	TwifInfo *TwifInfo `json:"twifInfo,omitempty"`

	RanUnchangedInd *bool `json:"ranUnchangedInd,omitempty"`
}
