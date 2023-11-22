/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type MediaType string

// List of MediaType
const (
	MEDIATYPE_AUDIO       MediaType = "AUDIO"
	MEDIATYPE_VIDEO       MediaType = "VIDEO"
	MEDIATYPE_DATA        MediaType = "DATA"
	MEDIATYPE_APPLICATION MediaType = "APPLICATION"
	MEDIATYPE_CONTROL     MediaType = "CONTROL"
	MEDIATYPE_TEXT        MediaType = "TEXT"
	MEDIATYPE_MESSAGE     MediaType = "MESSAGE"
	MEDIATYPE_OTHER       MediaType = "OTHER"
)
