package pfcpmsg

import "etrib5gc/pfcp/pfcptypes"

type Message struct {
	Header Header
	Body   interface{}
}

func (message *Message) IsRequest() (IsRequest bool) {
	switch message.Header.MessageType {
	case PFCP_HEARTBEAT_REQUEST:
		IsRequest = true
	case PFCP_PFD_MANAGEMENT_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_SETUP_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_UPDATE_REQUEST:
		IsRequest = true
	case PFCP_ASSOCIATION_RELEASE_REQUEST:
		IsRequest = true
	case PFCP_NODE_REPORT_REQUEST:
		IsRequest = true
	case PFCP_SESSION_SET_DELETION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_ESTABLISHMENT_REQUEST:
		IsRequest = true
	case PFCP_SESSION_MODIFICATION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_DELETION_REQUEST:
		IsRequest = true
	case PFCP_SESSION_REPORT_REQUEST:
		IsRequest = true
	default:
		IsRequest = false
	}

	return
}

func (message *Message) IsResponse() (IsResponse bool) {
	IsResponse = false
	switch message.Header.MessageType {
	case PFCP_HEARTBEAT_RESPONSE:
		IsResponse = true
	case PFCP_PFD_MANAGEMENT_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_SETUP_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_UPDATE_RESPONSE:
		IsResponse = true
	case PFCP_ASSOCIATION_RELEASE_RESPONSE:
		IsResponse = true
	case PFCP_NODE_REPORT_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_SET_DELETION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_ESTABLISHMENT_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_MODIFICATION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_DELETION_RESPONSE:
		IsResponse = true
	case PFCP_SESSION_REPORT_RESPONSE:
		IsResponse = true
	default:
		IsResponse = false
	}

	return
}

type HeartbeatRequest struct {
	RecoveryTimeStamp *pfcptypes.RecoveryTimeStamp `tlv:"96"`
}

type HeartbeatResponse struct {
	RecoveryTimeStamp *pfcptypes.RecoveryTimeStamp `tlv:"96"`
}

type PFCPPFDManagementRequest struct {
	ApplicationIDsPFDs []ApplicationIDsPFDs `tlv:"58"`
}

type ApplicationIDsPFDs struct {
	ApplicationID pfcptypes.ApplicationID `tlv:"24"`
	PFD           *PFD                    `tlv:"59"`
}

type PFD struct {
	PFDContents []pfcptypes.PFDContents `tlv:"61"`
}

type PFCPPFDManagementResponse struct {
	Cause       *pfcptypes.Cause       `tlv:"19"`
	OffendingIE *pfcptypes.OffendingIE `tlv:"40"`
}

type PFCPAssociationSetupRequest struct {
	NodeID                         *pfcptypes.NodeID                         `tlv:"60"`
	RecoveryTimeStamp              *pfcptypes.RecoveryTimeStamp              `tlv:"96"`
	UPFunctionFeatures             *pfcptypes.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *pfcptypes.CPFunctionFeatures             `tlv:"89"`
	UserPlaneIPResourceInformation *pfcptypes.UserPlaneIPResourceInformation `tlv:"116"`
}

type PFCPAssociationSetupResponse struct {
	NodeID                         *pfcptypes.NodeID                         `tlv:"60"`
	Cause                          *pfcptypes.Cause                          `tlv:"19"`
	RecoveryTimeStamp              *pfcptypes.RecoveryTimeStamp              `tlv:"96"`
	UPFunctionFeatures             *pfcptypes.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *pfcptypes.CPFunctionFeatures             `tlv:"89"`
	UserPlaneIPResourceInformation *pfcptypes.UserPlaneIPResourceInformation `tlv:"116"`
}

type PFCPAssociationUpdateRequest struct {
	NodeID                         *pfcptypes.NodeID                         `tlv:"60"`
	UPFunctionFeatures             *pfcptypes.UPFunctionFeatures             `tlv:"43"`
	CPFunctionFeatures             *pfcptypes.CPFunctionFeatures             `tlv:"89"`
	PFCPAssociationReleaseRequest  *PFCPAssociationReleaseRequest            `tlv:"111"`
	GracefulReleasePeriod          *pfcptypes.GracefulReleasePeriod          `tlv:"112"`
	UserPlaneIPResourceInformation *pfcptypes.UserPlaneIPResourceInformation `tlv:"116"`
}

type PFCPAssociationUpdateResponse struct {
	NodeID             *pfcptypes.NodeID             `tlv:"60"`
	Cause              *pfcptypes.Cause              `tlv:"19"`
	UPFunctionFeatures *pfcptypes.UPFunctionFeatures `tlv:"43"`
	CPFunctionFeatures *pfcptypes.CPFunctionFeatures `tlv:"89"`
}

type PFCPAssociationReleaseRequest struct {
	NodeID *pfcptypes.NodeID `tlv:"60"`
}

type PFCPAssociationReleaseResponse struct {
	NodeID *pfcptypes.NodeID `tlv:"60"`
	Cause  *pfcptypes.Cause  `tlv:"19"`
}

type PFCPNodeReportRequest struct {
	NodeID                     *pfcptypes.NodeID                     `tlv:"60"`
	NodeReportType             *pfcptypes.NodeReportType             `tlv:"101"`
	UserPlanePathFailureReport *pfcptypes.UserPlanePathFailureReport `tlv:"102"`
}

type UserPlanePathFailure struct {
	RemoteGTPUPeer *pfcptypes.RemoteGTPUPeer `tlv:"103"`
}

type PFCPNodeReportResponse struct {
	NodeID      *pfcptypes.NodeID      `tlv:"60"`
	Cause       *pfcptypes.Cause       `tlv:"19"`
	OffendingIE *pfcptypes.OffendingIE `tlv:"40"`
}

type PFCPSessionSetDeletionRequest struct {
	NodeID     *pfcptypes.NodeID `tlv:"60"`
	SGWCFQCSID *pfcptypes.FQCSID `tlv:"65"`
	PGWCFQCSID *pfcptypes.FQCSID `tlv:"65"`
	SGWUFQCSID *pfcptypes.FQCSID `tlv:"65"`
	PGWUFQCSID *pfcptypes.FQCSID `tlv:"65"`
	TWANFQCSID *pfcptypes.FQCSID `tlv:"65"`
	EPDGFQCSID *pfcptypes.FQCSID `tlv:"65"`
	MMEFQCSID  *pfcptypes.FQCSID `tlv:"65"`
}

type PFCPSessionSetDeletionResponse struct {
	NodeID      *pfcptypes.NodeID      `tlv:"60"`
	Cause       *pfcptypes.Cause       `tlv:"19"`
	OffendingIE *pfcptypes.OffendingIE `tlv:"40"`
}

type PFCPSessionEstablishmentRequest struct {
	NodeID                   *pfcptypes.NodeID                   `tlv:"60"`
	CPFSEID                  *pfcptypes.FSEID                    `tlv:"57"`
	CreatePDR                []*CreatePDR                        `tlv:"1"`
	CreateFAR                []*CreateFAR                        `tlv:"3"`
	CreateURR                []*CreateURR                        `tlv:"6"`
	CreateQER                []*CreateQER                        `tlv:"7"`
	CreateBAR                []*CreateBAR                        `tlv:"85"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint              `tlv:"127"`
	PDNType                  *pfcptypes.PDNType                  `tlv:"113"`
	SGWCFQCSID               *pfcptypes.FQCSID                   `tlv:"65"`
	MMEFQCSID                *pfcptypes.FQCSID                   `tlv:"65"`
	PGWCFQCSID               *pfcptypes.FQCSID                   `tlv:"65"`
	EPDGFQCSID               *pfcptypes.FQCSID                   `tlv:"65"`
	TWANFQCSID               *pfcptypes.FQCSID                   `tlv:"65"`
	UserPlaneInactivityTimer *pfcptypes.UserPlaneInactivityTimer `tlv:"117"`
	UserID                   *pfcptypes.UserID                   `tlv:"141"`
	TraceInformation         *pfcptypes.TraceInformation         `tlv:"152"`
}

type CreatePDR struct {
	PDRID                   *pfcptypes.PacketDetectionRuleID   `tlv:"56"`
	Precedence              *pfcptypes.Precedence              `tlv:"29"`
	PDI                     *PDI                               `tlv:"2"`
	OuterHeaderRemoval      *pfcptypes.OuterHeaderRemoval      `tlv:"95"`
	FARID                   *pfcptypes.FARID                   `tlv:"108"`
	URRID                   []*pfcptypes.URRID                 `tlv:"81"`
	QERID                   []*pfcptypes.QERID                 `tlv:"109"`
	ActivatePredefinedRules *pfcptypes.ActivatePredefinedRules `tlv:"106"`
}

type PDI struct {
	SourceInterface               *pfcptypes.SourceInterface               `tlv:"20"`
	LocalFTEID                    *pfcptypes.FTEID                         `tlv:"21"`
	NetworkInstance               *pfcptypes.NetworkInstance               `tlv:"22"`
	UEIPAddress                   *pfcptypes.UEIPAddress                   `tlv:"93"`
	TrafficEndpointID             *pfcptypes.TrafficEndpointID             `tlv:"131"`
	SDFFilter                     *pfcptypes.SDFFilter                     `tlv:"23"`
	ApplicationID                 *pfcptypes.ApplicationID                 `tlv:"24"`
	EthernetPDUSessionInformation *pfcptypes.EthernetPDUSessionInformation `tlv:"142"`
	EthernetPacketFilter          *EthernetPacketFilter                    `tlv:"132"`
	QFI                           []*pfcptypes.QFI                         `tlv:"124"`
	FramedRoute                   *pfcptypes.FramedRoute                   `tlv:"153"`
	FramedRouting                 *pfcptypes.FramedRouting                 `tlv:"154"`
	FramedIPv6Route               *pfcptypes.FramedIPv6Route               `tlv:"155"`
}

type EthernetPacketFilter struct {
	EthernetFilterID         *pfcptypes.EthernetFilterID         `tlv:"138"`
	EthernetFilterProperties *pfcptypes.EthernetFilterProperties `tlv:"139"`
	MACAddress               *pfcptypes.MACAddress               `tlv:"133"`
	Ethertype                *pfcptypes.Ethertype                `tlv:"136"`
	CTAG                     *pfcptypes.CTAG                     `tlv:"134"`
	STAG                     *pfcptypes.STAG                     `tlv:"135"`
	SDFFilter                *pfcptypes.SDFFilter                `tlv:"23"`
}

type CreateFAR struct {
	FARID                 *pfcptypes.FARID                 `tlv:"108"`
	ApplyAction           *pfcptypes.ApplyAction           `tlv:"44"`
	ForwardingParameters  *ForwardingParametersIEInFAR     `tlv:"4"`
	DuplicatingParameters *pfcptypes.DuplicatingParameters `tlv:"5"`
	BARID                 *pfcptypes.BARID                 `tlv:"88"`
}

type ForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcptypes.DestinationInterface  `tlv:"42"`
	NetworkInstance         *pfcptypes.NetworkInstance       `tlv:"22"`
	RedirectInformation     *pfcptypes.RedirectInformation   `tlv:"38"`
	OuterHeaderCreation     *pfcptypes.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking   *pfcptypes.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy        *pfcptypes.ForwardingPolicy      `tlv:"41"`
	HeaderEnrichment        *pfcptypes.HeaderEnrichment      `tlv:"98"`
	LinkedTrafficEndpointID *pfcptypes.TrafficEndpointID     `tlv:"131"`
	Proxying                *pfcptypes.Proxying              `tlv:"137"`
}

type DuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcptypes.DestinationInterface  `tlv:"42"`
	OuterHeaderCreation   *pfcptypes.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking *pfcptypes.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy      *pfcptypes.ForwardingPolicy      `tlv:"41"`
}

type CreateURR struct {
	URRID                     *pfcptypes.URRID                     `tlv:"81"`
	MeasurementMethod         *pfcptypes.MeasurementMethod         `tlv:"62"`
	ReportingTriggers         *pfcptypes.ReportingTriggers         `tlv:"37"`
	MeasurementPeriod         *pfcptypes.MeasurementPeriod         `tlv:"64"`
	VolumeThreshold           *pfcptypes.VolumeThreshold           `tlv:"31"`
	VolumeQuota               *pfcptypes.VolumeQuota               `tlv:"73"`
	TimeThreshold             *pfcptypes.TimeThreshold             `tlv:"32"`
	TimeQuota                 *pfcptypes.TimeQuota                 `tlv:"74"`
	QuotaHoldingTime          *pfcptypes.QuotaHoldingTime          `tlv:"71"`
	DroppedDLTrafficThreshold *pfcptypes.DroppedDLTrafficThreshold `tlv:"72"`
	MonitoringTime            *pfcptypes.MonitoringTime            `tlv:"33"`
	EventInformation          *EventInformation                    `tlv:"148"`
	SubsequentVolumeThreshold *pfcptypes.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *pfcptypes.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *pfcptypes.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *pfcptypes.SubsequentTimeQuota       `tlv:"122"`
	InactivityDetectionTime   *pfcptypes.InactivityDetectionTime   `tlv:"36"`
	LinkedURRID               *pfcptypes.LinkedURRID               `tlv:"82"`
	MeasurementInformation    *pfcptypes.MeasurementInformation    `tlv:"100"`
	TimeQuotaMechanism        *pfcptypes.TimeQuotaMechanism        `tlv:"115"`
	AggregatedURRs            []*AggregatedURRs                    `tlv:"118"`
	FARIDForQuotaAction       *pfcptypes.FARID                     `tlv:"108"`
	EthernetInactivityTimer   *pfcptypes.EthernetInactivityTimer   `tlv:"146"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime            `tlv:"147"`
	QuotaValidityTime         *pfcptypes.QuotaValidityTime         `tlv:"181"`
}

type AggregatedURRs struct {
	AggregatedURRID *pfcptypes.AggregatedURRID `tlv:"120"`
	Multiplier      *pfcptypes.Multiplier      `tlv:"119"`
}

type AdditionalMonitoringTime struct {
	MonitoringTime            *pfcptypes.MonitoringTime            `tlv:"33"`
	SubsequentVolumeThreshold *pfcptypes.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *pfcptypes.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *pfcptypes.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *pfcptypes.SubsequentTimeQuota       `tlv:"122"`
}

type EventInformation struct {
	EventID        *pfcptypes.EventID        `tlv:"150"`
	EventThreshold *pfcptypes.EventThreshold `tlv:"151"`
}

type CreateQER struct {
	QERID              *pfcptypes.QERID              `tlv:"109"`
	QERCorrelationID   *pfcptypes.QERCorrelationID   `tlv:"28"`
	GateStatus         *pfcptypes.GateStatus         `tlv:"25"`
	MaximumBitrate     *pfcptypes.MBR                `tlv:"26"`
	GuaranteedBitrate  *pfcptypes.GBR                `tlv:"27"`
	PacketRate         *pfcptypes.PacketRate         `tlv:"94"`
	DLFlowLevelMarking *pfcptypes.DLFlowLevelMarking `tlv:"97"`
	QoSFlowIdentifier  *pfcptypes.QFI                `tlv:"124"`
	ReflectiveQoS      *pfcptypes.RQI                `tlv:"123"`
}

type CreateBAR struct {
	BARID                          *pfcptypes.BARID                          `tlv:"88"`
	DownlinkDataNotificationDelay  *pfcptypes.DownlinkDataNotificationDelay  `tlv:"46"`
	SuggestedBufferingPacketsCount *pfcptypes.SuggestedBufferingPacketsCount `tlv:"140"`
}

type CreateTrafficEndpoint struct {
	TrafficEndpointID             *pfcptypes.TrafficEndpointID             `tlv:"131"`
	LocalFTEID                    *pfcptypes.FTEID                         `tlv:"21"`
	NetworkInstance               *pfcptypes.NetworkInstance               `tlv:"22"`
	UEIPAddress                   *pfcptypes.UEIPAddress                   `tlv:"93"`
	EthernetPDUSessionInformation *pfcptypes.EthernetPDUSessionInformation `tlv:"142"`
	FramedRoute                   *pfcptypes.FramedRoute                   `tlv:"153"`
	FramedRouting                 *pfcptypes.FramedRouting                 `tlv:"154"`
	FramedIPv6Route               *pfcptypes.FramedIPv6Route               `tlv:"155"`
}

type PFCPSessionEstablishmentResponse struct {
	NodeID                     *pfcptypes.NodeID           `tlv:"60"`
	Cause                      *pfcptypes.Cause            `tlv:"19"`
	OffendingIE                *pfcptypes.OffendingIE      `tlv:"40"`
	UPFSEID                    *pfcptypes.FSEID            `tlv:"57"`
	CreatedPDR                 *CreatedPDR                 `tlv:"8"`
	LoadControlInformation     *LoadControlInformation     `tlv:"51"`
	OverloadControlInformation *OverloadControlInformation `tlv:"54"`
	SGWUFQCSID                 *pfcptypes.FQCSID           `tlv:"65"`
	PGWUFQCSID                 *pfcptypes.FQCSID           `tlv:"65"`
	FailedRuleID               *pfcptypes.FailedRuleID     `tlv:"114"`
	CreatedTrafficEndpoint     *CreatedTrafficEndpoint     `tlv:"128"`
}

type CreatedPDR struct {
	PDRID      *pfcptypes.PacketDetectionRuleID `tlv:"56"`
	LocalFTEID *pfcptypes.FTEID                 `tlv:"21"`
}

type LoadControlInformation struct {
	LoadControlSequenceNumber *pfcptypes.SequenceNumber `tlv:"52"`
	LoadMetric                *pfcptypes.Metric         `tlv:"53"`
}

type OverloadControlInformation struct {
	OverloadControlSequenceNumber   *pfcptypes.SequenceNumber `tlv:"52"`
	OverloadReductionMetric         *pfcptypes.Metric         `tlv:"53"`
	PeriodOfValidity                *pfcptypes.Timer          `tlv:"55"`
	OverloadControlInformationFlags *pfcptypes.OCIFlags       `tlv:"110"`
}

type CreatedTrafficEndpoint struct {
	TrafficEndpointID *pfcptypes.TrafficEndpointID `tlv:"131"`
	LocalFTEID        *pfcptypes.FTEID             `tlv:"21"`
}

type PFCPSessionModificationRequest struct {
	CPFSEID                  *pfcptypes.FSEID                         `tlv:"57"`
	RemovePDR                []*RemovePDR                             `tlv:"15"`
	RemoveFAR                []*RemoveFAR                             `tlv:"16"`
	RemoveURR                []*RemoveURR                             `tlv:"17"`
	RemoveQER                []*RemoveQER                             `tlv:"18"`
	RemoveBAR                []*RemoveBAR                             `tlv:"87"`
	RemoveTrafficEndpoint    *RemoveTrafficEndpoint                   `tlv:"130"`
	CreatePDR                []*CreatePDR                             `tlv:"1"`
	CreateFAR                []*CreateFAR                             `tlv:"3"`
	CreateURR                []*CreateURR                             `tlv:"6"`
	CreateQER                []*CreateQER                             `tlv:"7"`
	CreateBAR                []*CreateBAR                             `tlv:"85"`
	CreateTrafficEndpoint    *CreateTrafficEndpoint                   `tlv:"127"`
	UpdatePDR                []*UpdatePDR                             `tlv:"9"`
	UpdateFAR                []*UpdateFAR                             `tlv:"10"`
	UpdateURR                []*UpdateURR                             `tlv:"13"`
	UpdateQER                []*UpdateQER                             `tlv:"14"`
	UpdateBAR                *UpdateBARPFCPSessionModificationRequest `tlv:"86"`
	UpdateTrafficEndpoint    *UpdateTrafficEndpoint                   `tlv:"129"`
	PFCPSMReqFlags           *pfcptypes.PFCPSMReqFlags                `tlv:"49"`
	QueryURR                 *QueryURR                                `tlv:"77"`
	PGWCFQCSID               *pfcptypes.FQCSID                        `tlv:"65"`
	SGWCFQCSID               *pfcptypes.FQCSID                        `tlv:"65"`
	MMEFQCSID                *pfcptypes.FQCSID                        `tlv:"65"`
	EPDGFQCSID               *pfcptypes.FQCSID                        `tlv:"65"`
	TWANFQCSID               *pfcptypes.FQCSID                        `tlv:"65"`
	UserPlaneInactivityTimer *pfcptypes.UserPlaneInactivityTimer      `tlv:"117"`
	QueryURRReference        *pfcptypes.QueryURRReference             `tlv:"125"`
	TraceInformation         *pfcptypes.TraceInformation              `tlv:"152"`
}

type UpdatePDR struct {
	PDRID                     *pfcptypes.PacketDetectionRuleID     `tlv:"56"`
	OuterHeaderRemoval        *pfcptypes.OuterHeaderRemoval        `tlv:"95"`
	Precedence                *pfcptypes.Precedence                `tlv:"29"`
	PDI                       *PDI                                 `tlv:"2"`
	FARID                     *pfcptypes.FARID                     `tlv:"108"`
	URRID                     []*pfcptypes.URRID                   `tlv:"81"`
	QERID                     []*pfcptypes.QERID                   `tlv:"109"`
	ActivatePredefinedRules   *pfcptypes.ActivatePredefinedRules   `tlv:"106"`
	DeactivatePredefinedRules *pfcptypes.DeactivatePredefinedRules `tlv:"107"`
}

type UpdateFAR struct {
	FARID                       *pfcptypes.FARID                       `tlv:"108"`
	ApplyAction                 *pfcptypes.ApplyAction                 `tlv:"44"`
	UpdateForwardingParameters  *UpdateForwardingParametersIEInFAR     `tlv:"11"`
	UpdateDuplicatingParameters *pfcptypes.UpdateDuplicatingParameters `tlv:"105"`
	BARID                       *pfcptypes.BARID                       `tlv:"88"`
}

type UpdateForwardingParametersIEInFAR struct {
	DestinationInterface    *pfcptypes.DestinationInterface  `tlv:"42"`
	NetworkInstance         *pfcptypes.NetworkInstance       `tlv:"22"`
	RedirectInformation     *pfcptypes.RedirectInformation   `tlv:"38"`
	OuterHeaderCreation     *pfcptypes.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking   *pfcptypes.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy        *pfcptypes.ForwardingPolicy      `tlv:"41"`
	HeaderEnrichment        *pfcptypes.HeaderEnrichment      `tlv:"98"`
	PFCPSMReqFlags          *pfcptypes.PFCPSMReqFlags        `tlv:"49"`
	LinkedTrafficEndpointID *pfcptypes.TrafficEndpointID     `tlv:"131"`
}

type UpdateDuplicatingParametersIEInFAR struct {
	DestinationInterface  *pfcptypes.DestinationInterface  `tlv:"42"`
	OuterHeaderCreation   *pfcptypes.OuterHeaderCreation   `tlv:"84"`
	TransportLevelMarking *pfcptypes.TransportLevelMarking `tlv:"30"`
	ForwardingPolicy      *pfcptypes.ForwardingPolicy      `tlv:"41"`
}

type UpdateURR struct {
	URRID                     *pfcptypes.URRID                     `tlv:"81"`
	MeasurementMethod         *pfcptypes.MeasurementMethod         `tlv:"62"`
	ReportingTriggers         *pfcptypes.ReportingTriggers         `tlv:"37"`
	MeasurementPeriod         *pfcptypes.MeasurementPeriod         `tlv:"64"`
	VolumeThreshold           *pfcptypes.VolumeThreshold           `tlv:"31"`
	VolumeQuota               *pfcptypes.VolumeQuota               `tlv:"73"`
	TimeThreshold             *pfcptypes.TimeThreshold             `tlv:"32"`
	TimeQuota                 *pfcptypes.TimeQuota                 `tlv:"74"`
	QuotaHoldingTime          *pfcptypes.QuotaHoldingTime          `tlv:"71"`
	DroppedDLTrafficThreshold *pfcptypes.DroppedDLTrafficThreshold `tlv:"72"`
	MonitoringTime            *pfcptypes.MonitoringTime            `tlv:"33"`
	EventInformation          *EventInformation                    `tlv:"148"`
	SubsequentVolumeThreshold *pfcptypes.SubsequentVolumeThreshold `tlv:"34"`
	SubsequentTimeThreshold   *pfcptypes.SubsequentTimeThreshold   `tlv:"35"`
	SubsequentVolumeQuota     *pfcptypes.SubsequentVolumeQuota     `tlv:"121"`
	SubsequentTimeQuota       *pfcptypes.SubsequentTimeQuota       `tlv:"122"`
	InactivityDetectionTime   *pfcptypes.InactivityDetectionTime   `tlv:"36"`
	LinkedURRID               *pfcptypes.LinkedURRID               `tlv:"82"`
	MeasurementInformation    *pfcptypes.MeasurementInformation    `tlv:"100"`
	TimeQuotaMechanism        *pfcptypes.TimeQuotaMechanism        `tlv:"115"`
	AggregatedURRs            *AggregatedURRs                      `tlv:"118"`
	FARIDForQuotaAction       *pfcptypes.FARID                     `tlv:"108"`
	EthernetInactivityTimer   *pfcptypes.EthernetInactivityTimer   `tlv:"146"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime            `tlv:"147"`
	QuotaValidityTime         *pfcptypes.QuotaValidityTime         `tlv:"181"`
}

type UpdateQER struct {
	QERID              *pfcptypes.QERID              `tlv:"109"`
	QERCorrelationID   *pfcptypes.QERCorrelationID   `tlv:"28"`
	GateStatus         *pfcptypes.GateStatus         `tlv:"25"`
	MaximumBitrate     *pfcptypes.MBR                `tlv:"26"`
	GuaranteedBitrate  *pfcptypes.GBR                `tlv:"27"`
	PacketRate         *pfcptypes.PacketRate         `tlv:"94"`
	DLFlowLevelMarking *pfcptypes.DLFlowLevelMarking `tlv:"97"`
	QoSFlowIdentifier  *pfcptypes.QFI                `tlv:"124"`
	ReflectiveQoS      *pfcptypes.RQI                `tlv:"123"`
}

type RemovePDR struct {
	PDRID *pfcptypes.PacketDetectionRuleID `tlv:"56"`
}

type RemoveFAR struct {
	FARID *pfcptypes.FARID `tlv:"108"`
}

type RemoveURR struct {
	URRID *pfcptypes.URRID `tlv:"81"`
}

type RemoveQER struct {
	QERID *pfcptypes.QERID `tlv:"109"`
}

type QueryURR struct {
	URRID *pfcptypes.URRID `tlv:"81"`
}

type UpdateBARPFCPSessionModificationRequest struct {
	BARID                          *pfcptypes.BARID                          `tlv:"88"`
	DownlinkDataNotificationDelay  *pfcptypes.DownlinkDataNotificationDelay  `tlv:"46"`
	SuggestedBufferingPacketsCount *pfcptypes.SuggestedBufferingPacketsCount `tlv:"140"`
}

type RemoveBAR struct {
	BARID *pfcptypes.BARID `tlv:"88"`
}

type UpdateTrafficEndpoint struct {
	TrafficEndpointID *pfcptypes.TrafficEndpointID `tlv:"131"`
	LocalFTEID        *pfcptypes.FTEID             `tlv:"21"`
	NetworkInstance   *pfcptypes.NetworkInstance   `tlv:"22"`
	UEIPAddress       *pfcptypes.UEIPAddress       `tlv:"93"`
	FramedRoute       *pfcptypes.FramedRoute       `tlv:"153"`
	FramedRouting     *pfcptypes.FramedRouting     `tlv:"154"`
	FramedIPv6Route   *pfcptypes.FramedIPv6Route   `tlv:"155"`
}

type RemoveTrafficEndpoint struct {
	TrafficEndpointID *pfcptypes.TrafficEndpointID `tlv:"131"`
}

type PFCPSessionModificationResponse struct {
	Cause                             *pfcptypes.Cause                              `tlv:"19"`
	OffendingIE                       *pfcptypes.OffendingIE                        `tlv:"40"`
	CreatedPDR                        *CreatedPDR                                   `tlv:"8"`
	LoadControlInformation            *LoadControlInformation                       `tlv:"51"`
	OverloadControlInformation        *OverloadControlInformation                   `tlv:"54"`
	UsageReport                       []*UsageReportPFCPSessionModificationResponse `tlv:"78"`
	FailedRuleID                      *pfcptypes.FailedRuleID                       `tlv:"114"`
	AdditionalUsageReportsInformation *pfcptypes.AdditionalUsageReportsInformation  `tlv:"126"`
	CreatedUpdatedTrafficEndpoint     *CreatedTrafficEndpoint                       `tlv:"128"`
}

type UsageReportPFCPSessionModificationResponse struct {
	URRID                      *pfcptypes.URRID               `tlv:"81"`
	URSEQN                     *pfcptypes.URSEQN              `tlv:"104"`
	UsageReportTrigger         *pfcptypes.UsageReportTrigger  `tlv:"63"`
	StartTime                  *pfcptypes.StartTime           `tlv:"75"`
	EndTime                    *pfcptypes.EndTime             `tlv:"76"`
	VolumeMeasurement          *pfcptypes.VolumeMeasurement   `tlv:"66"`
	DurationMeasurement        *pfcptypes.DurationMeasurement `tlv:"67"`
	TimeOfFirstPacket          *pfcptypes.TimeOfFirstPacket   `tlv:"69"`
	TimeOfLastPacket           *pfcptypes.TimeOfLastPacket    `tlv:"70"`
	UsageInformation           *pfcptypes.UsageInformation    `tlv:"90"`
	QueryURRReference          *pfcptypes.QueryURRReference   `tlv:"125"`
	EthernetTrafficInformation *EthernetTrafficInformation    `tlv:"143"`
}

type PFCPSessionDeletionRequest struct{}

type PFCPSessionDeletionResponse struct {
	Cause                      *pfcptypes.Cause                          `tlv:"19"`
	OffendingIE                *pfcptypes.OffendingIE                    `tlv:"40"`
	LoadControlInformation     *LoadControlInformation                   `tlv:"51"`
	OverloadControlInformation *OverloadControlInformation               `tlv:"54"`
	UsageReport                []*UsageReportPFCPSessionDeletionResponse `tlv:"79"`
}

type UsageReportPFCPSessionDeletionResponse struct {
	URRID                      *pfcptypes.URRID               `tlv:"81"`
	URSEQN                     *pfcptypes.URSEQN              `tlv:"104"`
	UsageReportTrigger         *pfcptypes.UsageReportTrigger  `tlv:"63"`
	StartTime                  *pfcptypes.StartTime           `tlv:"75"`
	EndTime                    *pfcptypes.EndTime             `tlv:"76"`
	VolumeMeasurement          *pfcptypes.VolumeMeasurement   `tlv:"66"`
	DurationMeasurement        *pfcptypes.DurationMeasurement `tlv:"67"`
	TimeOfFirstPacket          *pfcptypes.TimeOfFirstPacket   `tlv:"69"`
	TimeOfLastPacket           *pfcptypes.TimeOfLastPacket    `tlv:"70"`
	UsageInformation           *pfcptypes.UsageInformation    `tlv:"90"`
	EthernetTrafficInformation *EthernetTrafficInformation    `tlv:"143"`
}

type PFCPSessionReportRequest struct {
	ReportType                        *pfcptypes.ReportType                        `tlv:"39"`
	DownlinkDataReport                *DownlinkDataReport                          `tlv:"83"`
	UsageReport                       []*UsageReportPFCPSessionReportRequest       `tlv:"80"`
	ErrorIndicationReport             *ErrorIndicationReport                       `tlv:"99"`
	LoadControlInformation            *LoadControlInformation                      `tlv:"51"`
	OverloadControlInformation        *OverloadControlInformation                  `tlv:"54"`
	AdditionalUsageReportsInformation *pfcptypes.AdditionalUsageReportsInformation `tlv:"126"`
}

type DownlinkDataReport struct {
	PDRID                          *pfcptypes.PacketDetectionRuleID          `tlv:"56"`
	DownlinkDataServiceInformation *pfcptypes.DownlinkDataServiceInformation `tlv:"45"`
}

type UsageReportPFCPSessionReportRequest struct {
	URRID                           *pfcptypes.URRID                 `tlv:"81"`
	URSEQN                          *pfcptypes.URSEQN                `tlv:"104"`
	UsageReportTrigger              *pfcptypes.UsageReportTrigger    `tlv:"63"`
	StartTime                       *pfcptypes.StartTime             `tlv:"75"`
	EndTime                         *pfcptypes.EndTime               `tlv:"76"`
	VolumeMeasurement               *pfcptypes.VolumeMeasurement     `tlv:"66"`
	DurationMeasurement             *pfcptypes.DurationMeasurement   `tlv:"67"`
	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68"`
	UEIPAddress                     *pfcptypes.UEIPAddress           `tlv:"93"`
	NetworkInstance                 *pfcptypes.NetworkInstance       `tlv:"22"`
	TimeOfFirstPacket               *pfcptypes.TimeOfFirstPacket     `tlv:"69"`
	TimeOfLastPacket                *pfcptypes.TimeOfLastPacket      `tlv:"70"`
	UsageInformation                *pfcptypes.UsageInformation      `tlv:"90"`
	QueryURRReference               *pfcptypes.QueryURRReference     `tlv:"125"`
	EventReporting                  *EventReporting                  `tlv:"149"`
	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143"`
}

type ApplicationDetectionInformation struct {
	ApplicationID         *pfcptypes.ApplicationID         `tlv:"24"`
	ApplicationInstanceID *pfcptypes.ApplicationInstanceID `tlv:"91"`
	FlowInformation       *pfcptypes.FlowInformation       `tlv:"92"`
}

type EventReporting struct {
	EventID *pfcptypes.EventID `tlv:"150"`
}

type EthernetTrafficInformation struct {
	MACAddressesDetected *pfcptypes.MACAddressesDetected `tlv:"144"`
	MACAddressesRemoved  *pfcptypes.MACAddressesRemoved  `tlv:"145"`
}

type ErrorIndicationReport struct {
	RemoteFTEID *pfcptypes.FTEID `tlv:"21"`
}

type PFCPSessionReportResponse struct {
	Cause        *pfcptypes.Cause                              `tlv:"19"`
	OffendingIE  *pfcptypes.OffendingIE                        `tlv:"40"`
	UpdateBAR    *pfcptypes.UpdateBARPFCPSessionReportResponse `tlv:"12"`
	SxSRRspFlags *pfcptypes.PFCPSRRspFlags                     `tlv:"50"`
}

type UpdateBARIEInPFCPSessionReportResponse struct {
	BARID                           *pfcptypes.BARID                           `tlv:"88"`
	DownlinkDataNotificationDelay   *pfcptypes.DownlinkDataNotificationDelay   `tlv:"46"`
	DLBufferingDuration             *pfcptypes.DLBufferingDuration             `tlv:"47"`
	DLBufferingSuggestedPacketCount *pfcptypes.DLBufferingSuggestedPacketCount `tlv:"48"`
	SuggestedBufferingPacketsCount  *pfcptypes.SuggestedBufferingPacketsCount  `tlv:"140"`
}
