# MediaComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**AfRoutReq** | Pointer to [**AfRoutingRequirement**](AfRoutingRequirement.md) |  | [optional] 
**QosReference** | Pointer to **string** |  | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 
**AltSerReqs** | Pointer to **[]string** |  | [optional] 
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
**MedSubComps** | Pointer to [**map[string]MediaSubComponent**](MediaSubComponent.md) |  | [optional] 
**MedType** | Pointer to [**MediaType**](MediaType.md) |  | [optional] 
**MinDesBwDl** | Pointer to **string** |  | [optional] 
**MinDesBwUl** | Pointer to **string** |  | [optional] 
**MirBwDl** | Pointer to **string** |  | [optional] 
**MirBwUl** | Pointer to **string** |  | [optional] 
**PreemptCap** | Pointer to [**PreemptionCapability**](PreemptionCapability.md) |  | [optional] 
**PreemptVuln** | Pointer to [**PreemptionVulnerability**](PreemptionVulnerability.md) |  | [optional] 
**PrioSharingInd** | Pointer to [**PrioritySharingIndicator**](PrioritySharingIndicator.md) |  | [optional] 
**ResPrio** | Pointer to [**ReservPriority**](ReservPriority.md) |  | [optional] 
**RrBw** | Pointer to **string** |  | [optional] 
**RsBw** | Pointer to **string** |  | [optional] 
**SharingKeyDl** | Pointer to **int32** |  | [optional] 
**SharingKeyUl** | Pointer to **int32** |  | [optional] 
**TsnQos** | Pointer to [**TsnQosContainer**](TsnQosContainer.md) |  | [optional] 
**TscaiInputDl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 

## Methods

### NewMediaComponent

`func NewMediaComponent(medCompN int32, ) *MediaComponent`

NewMediaComponent instantiates a new MediaComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaComponentWithDefaults

`func NewMediaComponentWithDefaults() *MediaComponent`

NewMediaComponentWithDefaults instantiates a new MediaComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfAppId

`func (o *MediaComponent) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *MediaComponent) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *MediaComponent) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *MediaComponent) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetAfRoutReq

`func (o *MediaComponent) GetAfRoutReq() AfRoutingRequirement`

GetAfRoutReq returns the AfRoutReq field if non-nil, zero value otherwise.

### GetAfRoutReqOk

`func (o *MediaComponent) GetAfRoutReqOk() (*AfRoutingRequirement, bool)`

GetAfRoutReqOk returns a tuple with the AfRoutReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfRoutReq

`func (o *MediaComponent) SetAfRoutReq(v AfRoutingRequirement)`

SetAfRoutReq sets AfRoutReq field to given value.

### HasAfRoutReq

`func (o *MediaComponent) HasAfRoutReq() bool`

HasAfRoutReq returns a boolean if a field has been set.

### GetQosReference

`func (o *MediaComponent) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *MediaComponent) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *MediaComponent) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *MediaComponent) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetDisUeNotif

`func (o *MediaComponent) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *MediaComponent) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *MediaComponent) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *MediaComponent) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetAltSerReqs

`func (o *MediaComponent) GetAltSerReqs() []string`

GetAltSerReqs returns the AltSerReqs field if non-nil, zero value otherwise.

### GetAltSerReqsOk

`func (o *MediaComponent) GetAltSerReqsOk() (*[]string, bool)`

GetAltSerReqsOk returns a tuple with the AltSerReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSerReqs

`func (o *MediaComponent) SetAltSerReqs(v []string)`

SetAltSerReqs sets AltSerReqs field to given value.

### HasAltSerReqs

`func (o *MediaComponent) HasAltSerReqs() bool`

HasAltSerReqs returns a boolean if a field has been set.

### GetContVer

`func (o *MediaComponent) GetContVer() int32`

GetContVer returns the ContVer field if non-nil, zero value otherwise.

### GetContVerOk

`func (o *MediaComponent) GetContVerOk() (*int32, bool)`

GetContVerOk returns a tuple with the ContVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVer

`func (o *MediaComponent) SetContVer(v int32)`

SetContVer sets ContVer field to given value.

### HasContVer

`func (o *MediaComponent) HasContVer() bool`

HasContVer returns a boolean if a field has been set.

### GetCodecs

`func (o *MediaComponent) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *MediaComponent) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *MediaComponent) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *MediaComponent) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### GetDesMaxLatency

`func (o *MediaComponent) GetDesMaxLatency() float32`

GetDesMaxLatency returns the DesMaxLatency field if non-nil, zero value otherwise.

### GetDesMaxLatencyOk

`func (o *MediaComponent) GetDesMaxLatencyOk() (*float32, bool)`

GetDesMaxLatencyOk returns a tuple with the DesMaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesMaxLatency

`func (o *MediaComponent) SetDesMaxLatency(v float32)`

SetDesMaxLatency sets DesMaxLatency field to given value.

### HasDesMaxLatency

`func (o *MediaComponent) HasDesMaxLatency() bool`

HasDesMaxLatency returns a boolean if a field has been set.

### GetDesMaxLoss

`func (o *MediaComponent) GetDesMaxLoss() float32`

GetDesMaxLoss returns the DesMaxLoss field if non-nil, zero value otherwise.

### GetDesMaxLossOk

`func (o *MediaComponent) GetDesMaxLossOk() (*float32, bool)`

GetDesMaxLossOk returns a tuple with the DesMaxLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesMaxLoss

`func (o *MediaComponent) SetDesMaxLoss(v float32)`

SetDesMaxLoss sets DesMaxLoss field to given value.

### HasDesMaxLoss

`func (o *MediaComponent) HasDesMaxLoss() bool`

HasDesMaxLoss returns a boolean if a field has been set.

### GetFlusId

`func (o *MediaComponent) GetFlusId() string`

GetFlusId returns the FlusId field if non-nil, zero value otherwise.

### GetFlusIdOk

`func (o *MediaComponent) GetFlusIdOk() (*string, bool)`

GetFlusIdOk returns a tuple with the FlusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlusId

`func (o *MediaComponent) SetFlusId(v string)`

SetFlusId sets FlusId field to given value.

### HasFlusId

`func (o *MediaComponent) HasFlusId() bool`

HasFlusId returns a boolean if a field has been set.

### GetFStatus

`func (o *MediaComponent) GetFStatus() FlowStatus`

GetFStatus returns the FStatus field if non-nil, zero value otherwise.

### GetFStatusOk

`func (o *MediaComponent) GetFStatusOk() (*FlowStatus, bool)`

GetFStatusOk returns a tuple with the FStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFStatus

`func (o *MediaComponent) SetFStatus(v FlowStatus)`

SetFStatus sets FStatus field to given value.

### HasFStatus

`func (o *MediaComponent) HasFStatus() bool`

HasFStatus returns a boolean if a field has been set.

### GetMarBwDl

`func (o *MediaComponent) GetMarBwDl() string`

GetMarBwDl returns the MarBwDl field if non-nil, zero value otherwise.

### GetMarBwDlOk

`func (o *MediaComponent) GetMarBwDlOk() (*string, bool)`

GetMarBwDlOk returns a tuple with the MarBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDl

`func (o *MediaComponent) SetMarBwDl(v string)`

SetMarBwDl sets MarBwDl field to given value.

### HasMarBwDl

`func (o *MediaComponent) HasMarBwDl() bool`

HasMarBwDl returns a boolean if a field has been set.

### GetMarBwUl

`func (o *MediaComponent) GetMarBwUl() string`

GetMarBwUl returns the MarBwUl field if non-nil, zero value otherwise.

### GetMarBwUlOk

`func (o *MediaComponent) GetMarBwUlOk() (*string, bool)`

GetMarBwUlOk returns a tuple with the MarBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUl

`func (o *MediaComponent) SetMarBwUl(v string)`

SetMarBwUl sets MarBwUl field to given value.

### HasMarBwUl

`func (o *MediaComponent) HasMarBwUl() bool`

HasMarBwUl returns a boolean if a field has been set.

### GetMaxPacketLossRateDl

`func (o *MediaComponent) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *MediaComponent) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *MediaComponent) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *MediaComponent) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.

### SetMaxPacketLossRateDlNil

`func (o *MediaComponent) SetMaxPacketLossRateDlNil(b bool)`

 SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil

### UnsetMaxPacketLossRateDl
`func (o *MediaComponent) UnsetMaxPacketLossRateDl()`

UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
### GetMaxPacketLossRateUl

`func (o *MediaComponent) GetMaxPacketLossRateUl() int32`

GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateUlOk

`func (o *MediaComponent) GetMaxPacketLossRateUlOk() (*int32, bool)`

GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateUl

`func (o *MediaComponent) SetMaxPacketLossRateUl(v int32)`

SetMaxPacketLossRateUl sets MaxPacketLossRateUl field to given value.

### HasMaxPacketLossRateUl

`func (o *MediaComponent) HasMaxPacketLossRateUl() bool`

HasMaxPacketLossRateUl returns a boolean if a field has been set.

### SetMaxPacketLossRateUlNil

`func (o *MediaComponent) SetMaxPacketLossRateUlNil(b bool)`

 SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil

### UnsetMaxPacketLossRateUl
`func (o *MediaComponent) UnsetMaxPacketLossRateUl()`

UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
### GetMaxSuppBwDl

`func (o *MediaComponent) GetMaxSuppBwDl() string`

GetMaxSuppBwDl returns the MaxSuppBwDl field if non-nil, zero value otherwise.

### GetMaxSuppBwDlOk

`func (o *MediaComponent) GetMaxSuppBwDlOk() (*string, bool)`

GetMaxSuppBwDlOk returns a tuple with the MaxSuppBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSuppBwDl

`func (o *MediaComponent) SetMaxSuppBwDl(v string)`

SetMaxSuppBwDl sets MaxSuppBwDl field to given value.

### HasMaxSuppBwDl

`func (o *MediaComponent) HasMaxSuppBwDl() bool`

HasMaxSuppBwDl returns a boolean if a field has been set.

### GetMaxSuppBwUl

`func (o *MediaComponent) GetMaxSuppBwUl() string`

GetMaxSuppBwUl returns the MaxSuppBwUl field if non-nil, zero value otherwise.

### GetMaxSuppBwUlOk

`func (o *MediaComponent) GetMaxSuppBwUlOk() (*string, bool)`

GetMaxSuppBwUlOk returns a tuple with the MaxSuppBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSuppBwUl

`func (o *MediaComponent) SetMaxSuppBwUl(v string)`

SetMaxSuppBwUl sets MaxSuppBwUl field to given value.

### HasMaxSuppBwUl

`func (o *MediaComponent) HasMaxSuppBwUl() bool`

HasMaxSuppBwUl returns a boolean if a field has been set.

### GetMedCompN

`func (o *MediaComponent) GetMedCompN() int32`

GetMedCompN returns the MedCompN field if non-nil, zero value otherwise.

### GetMedCompNOk

`func (o *MediaComponent) GetMedCompNOk() (*int32, bool)`

GetMedCompNOk returns a tuple with the MedCompN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedCompN

`func (o *MediaComponent) SetMedCompN(v int32)`

SetMedCompN sets MedCompN field to given value.


### GetMedSubComps

`func (o *MediaComponent) GetMedSubComps() map[string]MediaSubComponent`

GetMedSubComps returns the MedSubComps field if non-nil, zero value otherwise.

### GetMedSubCompsOk

`func (o *MediaComponent) GetMedSubCompsOk() (*map[string]MediaSubComponent, bool)`

GetMedSubCompsOk returns a tuple with the MedSubComps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedSubComps

`func (o *MediaComponent) SetMedSubComps(v map[string]MediaSubComponent)`

SetMedSubComps sets MedSubComps field to given value.

### HasMedSubComps

`func (o *MediaComponent) HasMedSubComps() bool`

HasMedSubComps returns a boolean if a field has been set.

### GetMedType

`func (o *MediaComponent) GetMedType() MediaType`

GetMedType returns the MedType field if non-nil, zero value otherwise.

### GetMedTypeOk

`func (o *MediaComponent) GetMedTypeOk() (*MediaType, bool)`

GetMedTypeOk returns a tuple with the MedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedType

`func (o *MediaComponent) SetMedType(v MediaType)`

SetMedType sets MedType field to given value.

### HasMedType

`func (o *MediaComponent) HasMedType() bool`

HasMedType returns a boolean if a field has been set.

### GetMinDesBwDl

`func (o *MediaComponent) GetMinDesBwDl() string`

GetMinDesBwDl returns the MinDesBwDl field if non-nil, zero value otherwise.

### GetMinDesBwDlOk

`func (o *MediaComponent) GetMinDesBwDlOk() (*string, bool)`

GetMinDesBwDlOk returns a tuple with the MinDesBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwDl

`func (o *MediaComponent) SetMinDesBwDl(v string)`

SetMinDesBwDl sets MinDesBwDl field to given value.

### HasMinDesBwDl

`func (o *MediaComponent) HasMinDesBwDl() bool`

HasMinDesBwDl returns a boolean if a field has been set.

### GetMinDesBwUl

`func (o *MediaComponent) GetMinDesBwUl() string`

GetMinDesBwUl returns the MinDesBwUl field if non-nil, zero value otherwise.

### GetMinDesBwUlOk

`func (o *MediaComponent) GetMinDesBwUlOk() (*string, bool)`

GetMinDesBwUlOk returns a tuple with the MinDesBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwUl

`func (o *MediaComponent) SetMinDesBwUl(v string)`

SetMinDesBwUl sets MinDesBwUl field to given value.

### HasMinDesBwUl

`func (o *MediaComponent) HasMinDesBwUl() bool`

HasMinDesBwUl returns a boolean if a field has been set.

### GetMirBwDl

`func (o *MediaComponent) GetMirBwDl() string`

GetMirBwDl returns the MirBwDl field if non-nil, zero value otherwise.

### GetMirBwDlOk

`func (o *MediaComponent) GetMirBwDlOk() (*string, bool)`

GetMirBwDlOk returns a tuple with the MirBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwDl

`func (o *MediaComponent) SetMirBwDl(v string)`

SetMirBwDl sets MirBwDl field to given value.

### HasMirBwDl

`func (o *MediaComponent) HasMirBwDl() bool`

HasMirBwDl returns a boolean if a field has been set.

### GetMirBwUl

`func (o *MediaComponent) GetMirBwUl() string`

GetMirBwUl returns the MirBwUl field if non-nil, zero value otherwise.

### GetMirBwUlOk

`func (o *MediaComponent) GetMirBwUlOk() (*string, bool)`

GetMirBwUlOk returns a tuple with the MirBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwUl

`func (o *MediaComponent) SetMirBwUl(v string)`

SetMirBwUl sets MirBwUl field to given value.

### HasMirBwUl

`func (o *MediaComponent) HasMirBwUl() bool`

HasMirBwUl returns a boolean if a field has been set.

### GetPreemptCap

`func (o *MediaComponent) GetPreemptCap() PreemptionCapability`

GetPreemptCap returns the PreemptCap field if non-nil, zero value otherwise.

### GetPreemptCapOk

`func (o *MediaComponent) GetPreemptCapOk() (*PreemptionCapability, bool)`

GetPreemptCapOk returns a tuple with the PreemptCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptCap

`func (o *MediaComponent) SetPreemptCap(v PreemptionCapability)`

SetPreemptCap sets PreemptCap field to given value.

### HasPreemptCap

`func (o *MediaComponent) HasPreemptCap() bool`

HasPreemptCap returns a boolean if a field has been set.

### GetPreemptVuln

`func (o *MediaComponent) GetPreemptVuln() PreemptionVulnerability`

GetPreemptVuln returns the PreemptVuln field if non-nil, zero value otherwise.

### GetPreemptVulnOk

`func (o *MediaComponent) GetPreemptVulnOk() (*PreemptionVulnerability, bool)`

GetPreemptVulnOk returns a tuple with the PreemptVuln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptVuln

`func (o *MediaComponent) SetPreemptVuln(v PreemptionVulnerability)`

SetPreemptVuln sets PreemptVuln field to given value.

### HasPreemptVuln

`func (o *MediaComponent) HasPreemptVuln() bool`

HasPreemptVuln returns a boolean if a field has been set.

### GetPrioSharingInd

`func (o *MediaComponent) GetPrioSharingInd() PrioritySharingIndicator`

GetPrioSharingInd returns the PrioSharingInd field if non-nil, zero value otherwise.

### GetPrioSharingIndOk

`func (o *MediaComponent) GetPrioSharingIndOk() (*PrioritySharingIndicator, bool)`

GetPrioSharingIndOk returns a tuple with the PrioSharingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioSharingInd

`func (o *MediaComponent) SetPrioSharingInd(v PrioritySharingIndicator)`

SetPrioSharingInd sets PrioSharingInd field to given value.

### HasPrioSharingInd

`func (o *MediaComponent) HasPrioSharingInd() bool`

HasPrioSharingInd returns a boolean if a field has been set.

### GetResPrio

`func (o *MediaComponent) GetResPrio() ReservPriority`

GetResPrio returns the ResPrio field if non-nil, zero value otherwise.

### GetResPrioOk

`func (o *MediaComponent) GetResPrioOk() (*ReservPriority, bool)`

GetResPrioOk returns a tuple with the ResPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPrio

`func (o *MediaComponent) SetResPrio(v ReservPriority)`

SetResPrio sets ResPrio field to given value.

### HasResPrio

`func (o *MediaComponent) HasResPrio() bool`

HasResPrio returns a boolean if a field has been set.

### GetRrBw

`func (o *MediaComponent) GetRrBw() string`

GetRrBw returns the RrBw field if non-nil, zero value otherwise.

### GetRrBwOk

`func (o *MediaComponent) GetRrBwOk() (*string, bool)`

GetRrBwOk returns a tuple with the RrBw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrBw

`func (o *MediaComponent) SetRrBw(v string)`

SetRrBw sets RrBw field to given value.

### HasRrBw

`func (o *MediaComponent) HasRrBw() bool`

HasRrBw returns a boolean if a field has been set.

### GetRsBw

`func (o *MediaComponent) GetRsBw() string`

GetRsBw returns the RsBw field if non-nil, zero value otherwise.

### GetRsBwOk

`func (o *MediaComponent) GetRsBwOk() (*string, bool)`

GetRsBwOk returns a tuple with the RsBw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsBw

`func (o *MediaComponent) SetRsBw(v string)`

SetRsBw sets RsBw field to given value.

### HasRsBw

`func (o *MediaComponent) HasRsBw() bool`

HasRsBw returns a boolean if a field has been set.

### GetSharingKeyDl

`func (o *MediaComponent) GetSharingKeyDl() int32`

GetSharingKeyDl returns the SharingKeyDl field if non-nil, zero value otherwise.

### GetSharingKeyDlOk

`func (o *MediaComponent) GetSharingKeyDlOk() (*int32, bool)`

GetSharingKeyDlOk returns a tuple with the SharingKeyDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyDl

`func (o *MediaComponent) SetSharingKeyDl(v int32)`

SetSharingKeyDl sets SharingKeyDl field to given value.

### HasSharingKeyDl

`func (o *MediaComponent) HasSharingKeyDl() bool`

HasSharingKeyDl returns a boolean if a field has been set.

### GetSharingKeyUl

`func (o *MediaComponent) GetSharingKeyUl() int32`

GetSharingKeyUl returns the SharingKeyUl field if non-nil, zero value otherwise.

### GetSharingKeyUlOk

`func (o *MediaComponent) GetSharingKeyUlOk() (*int32, bool)`

GetSharingKeyUlOk returns a tuple with the SharingKeyUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyUl

`func (o *MediaComponent) SetSharingKeyUl(v int32)`

SetSharingKeyUl sets SharingKeyUl field to given value.

### HasSharingKeyUl

`func (o *MediaComponent) HasSharingKeyUl() bool`

HasSharingKeyUl returns a boolean if a field has been set.

### GetTsnQos

`func (o *MediaComponent) GetTsnQos() TsnQosContainer`

GetTsnQos returns the TsnQos field if non-nil, zero value otherwise.

### GetTsnQosOk

`func (o *MediaComponent) GetTsnQosOk() (*TsnQosContainer, bool)`

GetTsnQosOk returns a tuple with the TsnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnQos

`func (o *MediaComponent) SetTsnQos(v TsnQosContainer)`

SetTsnQos sets TsnQos field to given value.

### HasTsnQos

`func (o *MediaComponent) HasTsnQos() bool`

HasTsnQos returns a boolean if a field has been set.

### GetTscaiInputDl

`func (o *MediaComponent) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *MediaComponent) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *MediaComponent) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *MediaComponent) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *MediaComponent) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *MediaComponent) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *MediaComponent) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *MediaComponent) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *MediaComponent) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *MediaComponent) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *MediaComponent) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *MediaComponent) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


