# MediaComponentRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**AfRoutReq** | Pointer to [**AfRoutingRequirementRm**](AfRoutingRequirementRm.md) |  | [optional] 
**QosReference** | Pointer to **string** |  | [optional] 
**AltSerReqs** | Pointer to **[]string** |  | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 
**ContVer** | Pointer to **int32** | Represents the content version of some content. | [optional] 
**Codecs** | Pointer to **[]string** |  | [optional] 
**DesMaxLatency** | Pointer to **float32** |  | [optional] 
**DesMaxLoss** | Pointer to **float32** |  | [optional] 
**FlusId** | Pointer to **string** |  | [optional] 
**FStatus** | Pointer to [**FlowStatus**](FlowStatus.md) |  | [optional] 
**MarBwDl** | Pointer to **string** |  | [optional] 
**MarBwUl** | Pointer to **string** |  | [optional] 
**MaxPacketLossRateDl** | Pointer to **int32** |  | [optional] 
**MaxPacketLossRateUl** | Pointer to **int32** |  | [optional] 
**MaxSuppBwDl** | Pointer to **string** |  | [optional] 
**MaxSuppBwUl** | Pointer to **string** |  | [optional] 
**MedCompN** | **int32** |  | 
**MedSubComps** | Pointer to [**map[string]MediaSubComponentRm**](MediaSubComponentRm.md) |  | [optional] 
**MedType** | Pointer to [**MediaType**](MediaType.md) |  | [optional] 
**MinDesBwDl** | Pointer to **string** |  | [optional] 
**MinDesBwUl** | Pointer to **string** |  | [optional] 
**MirBwDl** | Pointer to **string** |  | [optional] 
**MirBwUl** | Pointer to **string** |  | [optional] 
**PreemptCap** | Pointer to [**PreemptionCapabilityRm**](PreemptionCapabilityRm.md) |  | [optional] 
**PreemptVuln** | Pointer to [**PreemptionVulnerabilityRm**](PreemptionVulnerabilityRm.md) |  | [optional] 
**PrioSharingInd** | Pointer to [**PrioritySharingIndicator**](PrioritySharingIndicator.md) |  | [optional] 
**ResPrio** | Pointer to [**ReservPriority**](ReservPriority.md) |  | [optional] 
**RrBw** | Pointer to **string** |  | [optional] 
**RsBw** | Pointer to **string** |  | [optional] 
**SharingKeyDl** | Pointer to **int32** |  | [optional] 
**SharingKeyUl** | Pointer to **int32** |  | [optional] 
**TsnQos** | Pointer to [**TsnQosContainerRm**](TsnQosContainerRm.md) |  | [optional] 
**TscaiInputDl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 

## Methods

### NewMediaComponentRm

`func NewMediaComponentRm(medCompN int32, ) *MediaComponentRm`

NewMediaComponentRm instantiates a new MediaComponentRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaComponentRmWithDefaults

`func NewMediaComponentRmWithDefaults() *MediaComponentRm`

NewMediaComponentRmWithDefaults instantiates a new MediaComponentRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfAppId

`func (o *MediaComponentRm) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *MediaComponentRm) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *MediaComponentRm) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *MediaComponentRm) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetAfRoutReq

`func (o *MediaComponentRm) GetAfRoutReq() AfRoutingRequirementRm`

GetAfRoutReq returns the AfRoutReq field if non-nil, zero value otherwise.

### GetAfRoutReqOk

`func (o *MediaComponentRm) GetAfRoutReqOk() (*AfRoutingRequirementRm, bool)`

GetAfRoutReqOk returns a tuple with the AfRoutReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfRoutReq

`func (o *MediaComponentRm) SetAfRoutReq(v AfRoutingRequirementRm)`

SetAfRoutReq sets AfRoutReq field to given value.

### HasAfRoutReq

`func (o *MediaComponentRm) HasAfRoutReq() bool`

HasAfRoutReq returns a boolean if a field has been set.

### SetAfRoutReqNil

`func (o *MediaComponentRm) SetAfRoutReqNil(b bool)`

 SetAfRoutReqNil sets the value for AfRoutReq to be an explicit nil

### UnsetAfRoutReq
`func (o *MediaComponentRm) UnsetAfRoutReq()`

UnsetAfRoutReq ensures that no value is present for AfRoutReq, not even an explicit nil
### GetQosReference

`func (o *MediaComponentRm) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *MediaComponentRm) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *MediaComponentRm) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *MediaComponentRm) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### SetQosReferenceNil

`func (o *MediaComponentRm) SetQosReferenceNil(b bool)`

 SetQosReferenceNil sets the value for QosReference to be an explicit nil

### UnsetQosReference
`func (o *MediaComponentRm) UnsetQosReference()`

UnsetQosReference ensures that no value is present for QosReference, not even an explicit nil
### GetAltSerReqs

`func (o *MediaComponentRm) GetAltSerReqs() []string`

GetAltSerReqs returns the AltSerReqs field if non-nil, zero value otherwise.

### GetAltSerReqsOk

`func (o *MediaComponentRm) GetAltSerReqsOk() (*[]string, bool)`

GetAltSerReqsOk returns a tuple with the AltSerReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSerReqs

`func (o *MediaComponentRm) SetAltSerReqs(v []string)`

SetAltSerReqs sets AltSerReqs field to given value.

### HasAltSerReqs

`func (o *MediaComponentRm) HasAltSerReqs() bool`

HasAltSerReqs returns a boolean if a field has been set.

### SetAltSerReqsNil

`func (o *MediaComponentRm) SetAltSerReqsNil(b bool)`

 SetAltSerReqsNil sets the value for AltSerReqs to be an explicit nil

### UnsetAltSerReqs
`func (o *MediaComponentRm) UnsetAltSerReqs()`

UnsetAltSerReqs ensures that no value is present for AltSerReqs, not even an explicit nil
### GetDisUeNotif

`func (o *MediaComponentRm) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *MediaComponentRm) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *MediaComponentRm) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *MediaComponentRm) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetContVer

`func (o *MediaComponentRm) GetContVer() int32`

GetContVer returns the ContVer field if non-nil, zero value otherwise.

### GetContVerOk

`func (o *MediaComponentRm) GetContVerOk() (*int32, bool)`

GetContVerOk returns a tuple with the ContVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVer

`func (o *MediaComponentRm) SetContVer(v int32)`

SetContVer sets ContVer field to given value.

### HasContVer

`func (o *MediaComponentRm) HasContVer() bool`

HasContVer returns a boolean if a field has been set.

### GetCodecs

`func (o *MediaComponentRm) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *MediaComponentRm) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *MediaComponentRm) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *MediaComponentRm) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### GetDesMaxLatency

`func (o *MediaComponentRm) GetDesMaxLatency() float32`

GetDesMaxLatency returns the DesMaxLatency field if non-nil, zero value otherwise.

### GetDesMaxLatencyOk

`func (o *MediaComponentRm) GetDesMaxLatencyOk() (*float32, bool)`

GetDesMaxLatencyOk returns a tuple with the DesMaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesMaxLatency

`func (o *MediaComponentRm) SetDesMaxLatency(v float32)`

SetDesMaxLatency sets DesMaxLatency field to given value.

### HasDesMaxLatency

`func (o *MediaComponentRm) HasDesMaxLatency() bool`

HasDesMaxLatency returns a boolean if a field has been set.

### SetDesMaxLatencyNil

`func (o *MediaComponentRm) SetDesMaxLatencyNil(b bool)`

 SetDesMaxLatencyNil sets the value for DesMaxLatency to be an explicit nil

### UnsetDesMaxLatency
`func (o *MediaComponentRm) UnsetDesMaxLatency()`

UnsetDesMaxLatency ensures that no value is present for DesMaxLatency, not even an explicit nil
### GetDesMaxLoss

`func (o *MediaComponentRm) GetDesMaxLoss() float32`

GetDesMaxLoss returns the DesMaxLoss field if non-nil, zero value otherwise.

### GetDesMaxLossOk

`func (o *MediaComponentRm) GetDesMaxLossOk() (*float32, bool)`

GetDesMaxLossOk returns a tuple with the DesMaxLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesMaxLoss

`func (o *MediaComponentRm) SetDesMaxLoss(v float32)`

SetDesMaxLoss sets DesMaxLoss field to given value.

### HasDesMaxLoss

`func (o *MediaComponentRm) HasDesMaxLoss() bool`

HasDesMaxLoss returns a boolean if a field has been set.

### SetDesMaxLossNil

`func (o *MediaComponentRm) SetDesMaxLossNil(b bool)`

 SetDesMaxLossNil sets the value for DesMaxLoss to be an explicit nil

### UnsetDesMaxLoss
`func (o *MediaComponentRm) UnsetDesMaxLoss()`

UnsetDesMaxLoss ensures that no value is present for DesMaxLoss, not even an explicit nil
### GetFlusId

`func (o *MediaComponentRm) GetFlusId() string`

GetFlusId returns the FlusId field if non-nil, zero value otherwise.

### GetFlusIdOk

`func (o *MediaComponentRm) GetFlusIdOk() (*string, bool)`

GetFlusIdOk returns a tuple with the FlusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlusId

`func (o *MediaComponentRm) SetFlusId(v string)`

SetFlusId sets FlusId field to given value.

### HasFlusId

`func (o *MediaComponentRm) HasFlusId() bool`

HasFlusId returns a boolean if a field has been set.

### SetFlusIdNil

`func (o *MediaComponentRm) SetFlusIdNil(b bool)`

 SetFlusIdNil sets the value for FlusId to be an explicit nil

### UnsetFlusId
`func (o *MediaComponentRm) UnsetFlusId()`

UnsetFlusId ensures that no value is present for FlusId, not even an explicit nil
### GetFStatus

`func (o *MediaComponentRm) GetFStatus() FlowStatus`

GetFStatus returns the FStatus field if non-nil, zero value otherwise.

### GetFStatusOk

`func (o *MediaComponentRm) GetFStatusOk() (*FlowStatus, bool)`

GetFStatusOk returns a tuple with the FStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFStatus

`func (o *MediaComponentRm) SetFStatus(v FlowStatus)`

SetFStatus sets FStatus field to given value.

### HasFStatus

`func (o *MediaComponentRm) HasFStatus() bool`

HasFStatus returns a boolean if a field has been set.

### GetMarBwDl

`func (o *MediaComponentRm) GetMarBwDl() string`

GetMarBwDl returns the MarBwDl field if non-nil, zero value otherwise.

### GetMarBwDlOk

`func (o *MediaComponentRm) GetMarBwDlOk() (*string, bool)`

GetMarBwDlOk returns a tuple with the MarBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDl

`func (o *MediaComponentRm) SetMarBwDl(v string)`

SetMarBwDl sets MarBwDl field to given value.

### HasMarBwDl

`func (o *MediaComponentRm) HasMarBwDl() bool`

HasMarBwDl returns a boolean if a field has been set.

### SetMarBwDlNil

`func (o *MediaComponentRm) SetMarBwDlNil(b bool)`

 SetMarBwDlNil sets the value for MarBwDl to be an explicit nil

### UnsetMarBwDl
`func (o *MediaComponentRm) UnsetMarBwDl()`

UnsetMarBwDl ensures that no value is present for MarBwDl, not even an explicit nil
### GetMarBwUl

`func (o *MediaComponentRm) GetMarBwUl() string`

GetMarBwUl returns the MarBwUl field if non-nil, zero value otherwise.

### GetMarBwUlOk

`func (o *MediaComponentRm) GetMarBwUlOk() (*string, bool)`

GetMarBwUlOk returns a tuple with the MarBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUl

`func (o *MediaComponentRm) SetMarBwUl(v string)`

SetMarBwUl sets MarBwUl field to given value.

### HasMarBwUl

`func (o *MediaComponentRm) HasMarBwUl() bool`

HasMarBwUl returns a boolean if a field has been set.

### SetMarBwUlNil

`func (o *MediaComponentRm) SetMarBwUlNil(b bool)`

 SetMarBwUlNil sets the value for MarBwUl to be an explicit nil

### UnsetMarBwUl
`func (o *MediaComponentRm) UnsetMarBwUl()`

UnsetMarBwUl ensures that no value is present for MarBwUl, not even an explicit nil
### GetMaxPacketLossRateDl

`func (o *MediaComponentRm) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *MediaComponentRm) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *MediaComponentRm) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *MediaComponentRm) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.

### SetMaxPacketLossRateDlNil

`func (o *MediaComponentRm) SetMaxPacketLossRateDlNil(b bool)`

 SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil

### UnsetMaxPacketLossRateDl
`func (o *MediaComponentRm) UnsetMaxPacketLossRateDl()`

UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
### GetMaxPacketLossRateUl

`func (o *MediaComponentRm) GetMaxPacketLossRateUl() int32`

GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateUlOk

`func (o *MediaComponentRm) GetMaxPacketLossRateUlOk() (*int32, bool)`

GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateUl

`func (o *MediaComponentRm) SetMaxPacketLossRateUl(v int32)`

SetMaxPacketLossRateUl sets MaxPacketLossRateUl field to given value.

### HasMaxPacketLossRateUl

`func (o *MediaComponentRm) HasMaxPacketLossRateUl() bool`

HasMaxPacketLossRateUl returns a boolean if a field has been set.

### SetMaxPacketLossRateUlNil

`func (o *MediaComponentRm) SetMaxPacketLossRateUlNil(b bool)`

 SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil

### UnsetMaxPacketLossRateUl
`func (o *MediaComponentRm) UnsetMaxPacketLossRateUl()`

UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
### GetMaxSuppBwDl

`func (o *MediaComponentRm) GetMaxSuppBwDl() string`

GetMaxSuppBwDl returns the MaxSuppBwDl field if non-nil, zero value otherwise.

### GetMaxSuppBwDlOk

`func (o *MediaComponentRm) GetMaxSuppBwDlOk() (*string, bool)`

GetMaxSuppBwDlOk returns a tuple with the MaxSuppBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSuppBwDl

`func (o *MediaComponentRm) SetMaxSuppBwDl(v string)`

SetMaxSuppBwDl sets MaxSuppBwDl field to given value.

### HasMaxSuppBwDl

`func (o *MediaComponentRm) HasMaxSuppBwDl() bool`

HasMaxSuppBwDl returns a boolean if a field has been set.

### SetMaxSuppBwDlNil

`func (o *MediaComponentRm) SetMaxSuppBwDlNil(b bool)`

 SetMaxSuppBwDlNil sets the value for MaxSuppBwDl to be an explicit nil

### UnsetMaxSuppBwDl
`func (o *MediaComponentRm) UnsetMaxSuppBwDl()`

UnsetMaxSuppBwDl ensures that no value is present for MaxSuppBwDl, not even an explicit nil
### GetMaxSuppBwUl

`func (o *MediaComponentRm) GetMaxSuppBwUl() string`

GetMaxSuppBwUl returns the MaxSuppBwUl field if non-nil, zero value otherwise.

### GetMaxSuppBwUlOk

`func (o *MediaComponentRm) GetMaxSuppBwUlOk() (*string, bool)`

GetMaxSuppBwUlOk returns a tuple with the MaxSuppBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSuppBwUl

`func (o *MediaComponentRm) SetMaxSuppBwUl(v string)`

SetMaxSuppBwUl sets MaxSuppBwUl field to given value.

### HasMaxSuppBwUl

`func (o *MediaComponentRm) HasMaxSuppBwUl() bool`

HasMaxSuppBwUl returns a boolean if a field has been set.

### SetMaxSuppBwUlNil

`func (o *MediaComponentRm) SetMaxSuppBwUlNil(b bool)`

 SetMaxSuppBwUlNil sets the value for MaxSuppBwUl to be an explicit nil

### UnsetMaxSuppBwUl
`func (o *MediaComponentRm) UnsetMaxSuppBwUl()`

UnsetMaxSuppBwUl ensures that no value is present for MaxSuppBwUl, not even an explicit nil
### GetMedCompN

`func (o *MediaComponentRm) GetMedCompN() int32`

GetMedCompN returns the MedCompN field if non-nil, zero value otherwise.

### GetMedCompNOk

`func (o *MediaComponentRm) GetMedCompNOk() (*int32, bool)`

GetMedCompNOk returns a tuple with the MedCompN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedCompN

`func (o *MediaComponentRm) SetMedCompN(v int32)`

SetMedCompN sets MedCompN field to given value.


### GetMedSubComps

`func (o *MediaComponentRm) GetMedSubComps() map[string]MediaSubComponentRm`

GetMedSubComps returns the MedSubComps field if non-nil, zero value otherwise.

### GetMedSubCompsOk

`func (o *MediaComponentRm) GetMedSubCompsOk() (*map[string]MediaSubComponentRm, bool)`

GetMedSubCompsOk returns a tuple with the MedSubComps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedSubComps

`func (o *MediaComponentRm) SetMedSubComps(v map[string]MediaSubComponentRm)`

SetMedSubComps sets MedSubComps field to given value.

### HasMedSubComps

`func (o *MediaComponentRm) HasMedSubComps() bool`

HasMedSubComps returns a boolean if a field has been set.

### GetMedType

`func (o *MediaComponentRm) GetMedType() MediaType`

GetMedType returns the MedType field if non-nil, zero value otherwise.

### GetMedTypeOk

`func (o *MediaComponentRm) GetMedTypeOk() (*MediaType, bool)`

GetMedTypeOk returns a tuple with the MedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedType

`func (o *MediaComponentRm) SetMedType(v MediaType)`

SetMedType sets MedType field to given value.

### HasMedType

`func (o *MediaComponentRm) HasMedType() bool`

HasMedType returns a boolean if a field has been set.

### GetMinDesBwDl

`func (o *MediaComponentRm) GetMinDesBwDl() string`

GetMinDesBwDl returns the MinDesBwDl field if non-nil, zero value otherwise.

### GetMinDesBwDlOk

`func (o *MediaComponentRm) GetMinDesBwDlOk() (*string, bool)`

GetMinDesBwDlOk returns a tuple with the MinDesBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwDl

`func (o *MediaComponentRm) SetMinDesBwDl(v string)`

SetMinDesBwDl sets MinDesBwDl field to given value.

### HasMinDesBwDl

`func (o *MediaComponentRm) HasMinDesBwDl() bool`

HasMinDesBwDl returns a boolean if a field has been set.

### SetMinDesBwDlNil

`func (o *MediaComponentRm) SetMinDesBwDlNil(b bool)`

 SetMinDesBwDlNil sets the value for MinDesBwDl to be an explicit nil

### UnsetMinDesBwDl
`func (o *MediaComponentRm) UnsetMinDesBwDl()`

UnsetMinDesBwDl ensures that no value is present for MinDesBwDl, not even an explicit nil
### GetMinDesBwUl

`func (o *MediaComponentRm) GetMinDesBwUl() string`

GetMinDesBwUl returns the MinDesBwUl field if non-nil, zero value otherwise.

### GetMinDesBwUlOk

`func (o *MediaComponentRm) GetMinDesBwUlOk() (*string, bool)`

GetMinDesBwUlOk returns a tuple with the MinDesBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwUl

`func (o *MediaComponentRm) SetMinDesBwUl(v string)`

SetMinDesBwUl sets MinDesBwUl field to given value.

### HasMinDesBwUl

`func (o *MediaComponentRm) HasMinDesBwUl() bool`

HasMinDesBwUl returns a boolean if a field has been set.

### SetMinDesBwUlNil

`func (o *MediaComponentRm) SetMinDesBwUlNil(b bool)`

 SetMinDesBwUlNil sets the value for MinDesBwUl to be an explicit nil

### UnsetMinDesBwUl
`func (o *MediaComponentRm) UnsetMinDesBwUl()`

UnsetMinDesBwUl ensures that no value is present for MinDesBwUl, not even an explicit nil
### GetMirBwDl

`func (o *MediaComponentRm) GetMirBwDl() string`

GetMirBwDl returns the MirBwDl field if non-nil, zero value otherwise.

### GetMirBwDlOk

`func (o *MediaComponentRm) GetMirBwDlOk() (*string, bool)`

GetMirBwDlOk returns a tuple with the MirBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwDl

`func (o *MediaComponentRm) SetMirBwDl(v string)`

SetMirBwDl sets MirBwDl field to given value.

### HasMirBwDl

`func (o *MediaComponentRm) HasMirBwDl() bool`

HasMirBwDl returns a boolean if a field has been set.

### SetMirBwDlNil

`func (o *MediaComponentRm) SetMirBwDlNil(b bool)`

 SetMirBwDlNil sets the value for MirBwDl to be an explicit nil

### UnsetMirBwDl
`func (o *MediaComponentRm) UnsetMirBwDl()`

UnsetMirBwDl ensures that no value is present for MirBwDl, not even an explicit nil
### GetMirBwUl

`func (o *MediaComponentRm) GetMirBwUl() string`

GetMirBwUl returns the MirBwUl field if non-nil, zero value otherwise.

### GetMirBwUlOk

`func (o *MediaComponentRm) GetMirBwUlOk() (*string, bool)`

GetMirBwUlOk returns a tuple with the MirBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwUl

`func (o *MediaComponentRm) SetMirBwUl(v string)`

SetMirBwUl sets MirBwUl field to given value.

### HasMirBwUl

`func (o *MediaComponentRm) HasMirBwUl() bool`

HasMirBwUl returns a boolean if a field has been set.

### SetMirBwUlNil

`func (o *MediaComponentRm) SetMirBwUlNil(b bool)`

 SetMirBwUlNil sets the value for MirBwUl to be an explicit nil

### UnsetMirBwUl
`func (o *MediaComponentRm) UnsetMirBwUl()`

UnsetMirBwUl ensures that no value is present for MirBwUl, not even an explicit nil
### GetPreemptCap

`func (o *MediaComponentRm) GetPreemptCap() PreemptionCapabilityRm`

GetPreemptCap returns the PreemptCap field if non-nil, zero value otherwise.

### GetPreemptCapOk

`func (o *MediaComponentRm) GetPreemptCapOk() (*PreemptionCapabilityRm, bool)`

GetPreemptCapOk returns a tuple with the PreemptCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptCap

`func (o *MediaComponentRm) SetPreemptCap(v PreemptionCapabilityRm)`

SetPreemptCap sets PreemptCap field to given value.

### HasPreemptCap

`func (o *MediaComponentRm) HasPreemptCap() bool`

HasPreemptCap returns a boolean if a field has been set.

### GetPreemptVuln

`func (o *MediaComponentRm) GetPreemptVuln() PreemptionVulnerabilityRm`

GetPreemptVuln returns the PreemptVuln field if non-nil, zero value otherwise.

### GetPreemptVulnOk

`func (o *MediaComponentRm) GetPreemptVulnOk() (*PreemptionVulnerabilityRm, bool)`

GetPreemptVulnOk returns a tuple with the PreemptVuln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptVuln

`func (o *MediaComponentRm) SetPreemptVuln(v PreemptionVulnerabilityRm)`

SetPreemptVuln sets PreemptVuln field to given value.

### HasPreemptVuln

`func (o *MediaComponentRm) HasPreemptVuln() bool`

HasPreemptVuln returns a boolean if a field has been set.

### GetPrioSharingInd

`func (o *MediaComponentRm) GetPrioSharingInd() PrioritySharingIndicator`

GetPrioSharingInd returns the PrioSharingInd field if non-nil, zero value otherwise.

### GetPrioSharingIndOk

`func (o *MediaComponentRm) GetPrioSharingIndOk() (*PrioritySharingIndicator, bool)`

GetPrioSharingIndOk returns a tuple with the PrioSharingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioSharingInd

`func (o *MediaComponentRm) SetPrioSharingInd(v PrioritySharingIndicator)`

SetPrioSharingInd sets PrioSharingInd field to given value.

### HasPrioSharingInd

`func (o *MediaComponentRm) HasPrioSharingInd() bool`

HasPrioSharingInd returns a boolean if a field has been set.

### GetResPrio

`func (o *MediaComponentRm) GetResPrio() ReservPriority`

GetResPrio returns the ResPrio field if non-nil, zero value otherwise.

### GetResPrioOk

`func (o *MediaComponentRm) GetResPrioOk() (*ReservPriority, bool)`

GetResPrioOk returns a tuple with the ResPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPrio

`func (o *MediaComponentRm) SetResPrio(v ReservPriority)`

SetResPrio sets ResPrio field to given value.

### HasResPrio

`func (o *MediaComponentRm) HasResPrio() bool`

HasResPrio returns a boolean if a field has been set.

### GetRrBw

`func (o *MediaComponentRm) GetRrBw() string`

GetRrBw returns the RrBw field if non-nil, zero value otherwise.

### GetRrBwOk

`func (o *MediaComponentRm) GetRrBwOk() (*string, bool)`

GetRrBwOk returns a tuple with the RrBw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrBw

`func (o *MediaComponentRm) SetRrBw(v string)`

SetRrBw sets RrBw field to given value.

### HasRrBw

`func (o *MediaComponentRm) HasRrBw() bool`

HasRrBw returns a boolean if a field has been set.

### SetRrBwNil

`func (o *MediaComponentRm) SetRrBwNil(b bool)`

 SetRrBwNil sets the value for RrBw to be an explicit nil

### UnsetRrBw
`func (o *MediaComponentRm) UnsetRrBw()`

UnsetRrBw ensures that no value is present for RrBw, not even an explicit nil
### GetRsBw

`func (o *MediaComponentRm) GetRsBw() string`

GetRsBw returns the RsBw field if non-nil, zero value otherwise.

### GetRsBwOk

`func (o *MediaComponentRm) GetRsBwOk() (*string, bool)`

GetRsBwOk returns a tuple with the RsBw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsBw

`func (o *MediaComponentRm) SetRsBw(v string)`

SetRsBw sets RsBw field to given value.

### HasRsBw

`func (o *MediaComponentRm) HasRsBw() bool`

HasRsBw returns a boolean if a field has been set.

### SetRsBwNil

`func (o *MediaComponentRm) SetRsBwNil(b bool)`

 SetRsBwNil sets the value for RsBw to be an explicit nil

### UnsetRsBw
`func (o *MediaComponentRm) UnsetRsBw()`

UnsetRsBw ensures that no value is present for RsBw, not even an explicit nil
### GetSharingKeyDl

`func (o *MediaComponentRm) GetSharingKeyDl() int32`

GetSharingKeyDl returns the SharingKeyDl field if non-nil, zero value otherwise.

### GetSharingKeyDlOk

`func (o *MediaComponentRm) GetSharingKeyDlOk() (*int32, bool)`

GetSharingKeyDlOk returns a tuple with the SharingKeyDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyDl

`func (o *MediaComponentRm) SetSharingKeyDl(v int32)`

SetSharingKeyDl sets SharingKeyDl field to given value.

### HasSharingKeyDl

`func (o *MediaComponentRm) HasSharingKeyDl() bool`

HasSharingKeyDl returns a boolean if a field has been set.

### SetSharingKeyDlNil

`func (o *MediaComponentRm) SetSharingKeyDlNil(b bool)`

 SetSharingKeyDlNil sets the value for SharingKeyDl to be an explicit nil

### UnsetSharingKeyDl
`func (o *MediaComponentRm) UnsetSharingKeyDl()`

UnsetSharingKeyDl ensures that no value is present for SharingKeyDl, not even an explicit nil
### GetSharingKeyUl

`func (o *MediaComponentRm) GetSharingKeyUl() int32`

GetSharingKeyUl returns the SharingKeyUl field if non-nil, zero value otherwise.

### GetSharingKeyUlOk

`func (o *MediaComponentRm) GetSharingKeyUlOk() (*int32, bool)`

GetSharingKeyUlOk returns a tuple with the SharingKeyUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyUl

`func (o *MediaComponentRm) SetSharingKeyUl(v int32)`

SetSharingKeyUl sets SharingKeyUl field to given value.

### HasSharingKeyUl

`func (o *MediaComponentRm) HasSharingKeyUl() bool`

HasSharingKeyUl returns a boolean if a field has been set.

### SetSharingKeyUlNil

`func (o *MediaComponentRm) SetSharingKeyUlNil(b bool)`

 SetSharingKeyUlNil sets the value for SharingKeyUl to be an explicit nil

### UnsetSharingKeyUl
`func (o *MediaComponentRm) UnsetSharingKeyUl()`

UnsetSharingKeyUl ensures that no value is present for SharingKeyUl, not even an explicit nil
### GetTsnQos

`func (o *MediaComponentRm) GetTsnQos() TsnQosContainerRm`

GetTsnQos returns the TsnQos field if non-nil, zero value otherwise.

### GetTsnQosOk

`func (o *MediaComponentRm) GetTsnQosOk() (*TsnQosContainerRm, bool)`

GetTsnQosOk returns a tuple with the TsnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnQos

`func (o *MediaComponentRm) SetTsnQos(v TsnQosContainerRm)`

SetTsnQos sets TsnQos field to given value.

### HasTsnQos

`func (o *MediaComponentRm) HasTsnQos() bool`

HasTsnQos returns a boolean if a field has been set.

### SetTsnQosNil

`func (o *MediaComponentRm) SetTsnQosNil(b bool)`

 SetTsnQosNil sets the value for TsnQos to be an explicit nil

### UnsetTsnQos
`func (o *MediaComponentRm) UnsetTsnQos()`

UnsetTsnQos ensures that no value is present for TsnQos, not even an explicit nil
### GetTscaiInputDl

`func (o *MediaComponentRm) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *MediaComponentRm) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *MediaComponentRm) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *MediaComponentRm) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *MediaComponentRm) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *MediaComponentRm) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *MediaComponentRm) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *MediaComponentRm) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *MediaComponentRm) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *MediaComponentRm) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *MediaComponentRm) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *MediaComponentRm) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


