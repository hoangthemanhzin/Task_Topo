/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// PolicyDataChangeNotification - Contains changed policy data for which notification was requested.
type PolicyDataChangeNotification struct {
	AmPolicyData AmPolicyData `json:"amPolicyData,omitempty"`

	UePolicySet UePolicySet `json:"uePolicySet,omitempty"`

	PlmnUePolicySet UePolicySet `json:"plmnUePolicySet,omitempty"`

	SmPolicyData SmPolicyData `json:"smPolicyData,omitempty"`

	UsageMonData UsageMonData `json:"usageMonData,omitempty"`

	SponsorConnectivityData SponsorConnectivityData `json:"SponsorConnectivityData,omitempty"`

	BdtData BdtData `json:"bdtData,omitempty"`

	OpSpecData OperatorSpecificDataContainer `json:"opSpecData,omitempty"`

	OpSpecDataMap map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`

	UeId string `json:"ueId,omitempty"`

	SponsorId string `json:"sponsorId,omitempty"`

	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId,omitempty"`

	UsageMonId string `json:"usageMonId,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	DelResources []string `json:"delResources,omitempty"`

	NotifId string `json:"notifId,omitempty"`

	ReportedFragments []NotificationItem `json:"reportedFragments,omitempty"`
}
