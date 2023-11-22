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

type ApnRateStatus struct {
	RemainPacketsUl int32 `json:"remainPacketsUl,omitempty"`

	RemainPacketsDl int32 `json:"remainPacketsDl,omitempty"`

	ValidityTime time.Time `json:"validityTime,omitempty"`

	RemainExReportsUl int32 `json:"remainExReportsUl,omitempty"`

	RemainExReportsDl int32 `json:"remainExReportsDl,omitempty"`
}
