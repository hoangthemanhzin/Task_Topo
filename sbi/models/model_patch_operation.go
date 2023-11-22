/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package models

type PatchOperation string

// List of PatchOperation
const (
	PATCHOPERATION_ADD     PatchOperation = "add"
	PATCHOPERATION_COPY    PatchOperation = "copy"
	PATCHOPERATION_MOVE    PatchOperation = "move"
	PATCHOPERATION_REMOVE  PatchOperation = "remove"
	PATCHOPERATION_REPLACE PatchOperation = "replace"
	PATCHOPERATION_TEST    PatchOperation = "test"
)
