/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type N2InformationNotification struct {
	N2NotifySubscriptionId string `json:"n2NotifySubscriptionId"`

	N2InfoContainer N2InfoContainer `json:"n2InfoContainer,omitempty"`

	ToReleaseSessionList []int32 `json:"toReleaseSessionList,omitempty"`

	LcsCorrelationId string `json:"lcsCorrelationId,omitempty"`

	NotifyReason N2InfoNotifyReason `json:"notifyReason,omitempty"`

	SmfChangeInfoList []SmfChangeInfo `json:"smfChangeInfoList,omitempty"`

	RanNodeId GlobalRanNodeId `json:"ranNodeId,omitempty"`

	InitialAmfName string `json:"initialAmfName,omitempty"`

	AnN2IPv4Addr string `json:"anN2IPv4Addr,omitempty"`

	AnN2IPv6Addr Ipv6Addr `json:"anN2IPv6Addr,omitempty"`

	Guami Guami `json:"guami,omitempty"`

	NotifySourceNgRan bool `json:"notifySourceNgRan,omitempty"`
}
