/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

// BdtPolicy - Represents an Individual BDT policy resource.
type BdtPolicy struct {
	BdtPolData BdtPolicyData `json:"bdtPolData,omitempty"`

	BdtReqData BdtReqData `json:"bdtReqData,omitempty"`
}
