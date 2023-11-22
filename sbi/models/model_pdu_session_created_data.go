/*
Nsmf_PDUSession

SMF PDU Session Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

import (
	"time"
)

type PduSessionCreatedData struct {
	PduSessionType PduSessionType `json:"pduSessionType"`

	SscMode string `json:"sscMode"`

	HcnTunnelInfo TunnelInfo `json:"hcnTunnelInfo,omitempty"`

	CnTunnelInfo TunnelInfo `json:"cnTunnelInfo,omitempty"`

	AdditionalCnTunnelInfo TunnelInfo `json:"additionalCnTunnelInfo,omitempty"`

	SessionAmbr Ambr `json:"sessionAmbr,omitempty"`

	QosFlowsSetupList []QosFlowSetupItem `json:"qosFlowsSetupList,omitempty"`

	HSmfInstanceId string `json:"hSmfInstanceId,omitempty"`

	SmfInstanceId string `json:"smfInstanceId,omitempty"`

	PduSessionId int32 `json:"pduSessionId,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	EnablePauseCharging bool `json:"enablePauseCharging,omitempty"`

	UeIpv4Address string `json:"ueIpv4Address,omitempty"`

	UeIpv6Prefix Ipv6Prefix `json:"ueIpv6Prefix,omitempty"`

	N1SmInfoToUe RefToBinaryData `json:"n1SmInfoToUe,omitempty"`

	EpsPdnCnxInfo EpsPdnCnxInfo `json:"epsPdnCnxInfo,omitempty"`

	EpsBearerInfo []EpsBearerInfo `json:"epsBearerInfo,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	MaxIntegrityProtectedDataRate MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRate,omitempty"`

	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`

	AlwaysOnGranted bool `json:"alwaysOnGranted,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	UpSecurity UpSecurity `json:"upSecurity,omitempty"`

	RoamingChargingProfile RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`

	HSmfServiceInstanceId string `json:"hSmfServiceInstanceId,omitempty"`

	SmfServiceInstanceId string `json:"smfServiceInstanceId,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	DnaiList []string `json:"dnaiList,omitempty"`

	Ipv6MultiHomingInd bool `json:"ipv6MultiHomingInd,omitempty"`

	MaAcceptedInd bool `json:"maAcceptedInd,omitempty"`

	HomeProvidedChargingId string `json:"homeProvidedChargingId,omitempty"`

	NefExtBufSupportInd bool `json:"nefExtBufSupportInd,omitempty"`

	SmallDataRateControlEnabled bool `json:"smallDataRateControlEnabled,omitempty"`

	UeIpv6InterfaceId string `json:"ueIpv6InterfaceId,omitempty"`

	Ipv6Index int32 `json:"ipv6Index,omitempty"`

	DnAaaAddress IpAddress `json:"dnAaaAddress,omitempty"`

	RedundantPduSessionInfo RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
}
