# MmContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AccessType**](AccessType.md) |  | 
**NasSecurityMode** | Pointer to [**NasSecurityMode**](NasSecurityMode.md) |  | [optional] 
**EpsNasSecurityMode** | Pointer to [**EpsNasSecurityMode**](EpsNasSecurityMode.md) |  | [optional] 
**NasDownlinkCount** | Pointer to **int32** |  | [optional] 
**NasUplinkCount** | Pointer to **int32** |  | [optional] 
**UeSecurityCapability** | Pointer to **string** |  | [optional] 
**S1UeNetworkCapability** | Pointer to **string** |  | [optional] 
**AllowedNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NssaiMappingList** | Pointer to [**[]NssaiMapping**](NssaiMapping.md) |  | [optional] 
**AllowedHomeNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NsInstanceList** | Pointer to **[]string** |  | [optional] 
**ExpectedUEbehavior** | Pointer to [**ExpectedUeBehavior**](ExpectedUeBehavior.md) |  | [optional] 
**UeDifferentiationInfo** | Pointer to [**UeDifferentiationInfo**](UeDifferentiationInfo.md) |  | [optional] 
**PlmnAssiUeRadioCapId** | Pointer to **string** |  | [optional] 
**ManAssiUeRadioCapId** | Pointer to **string** |  | [optional] 
**UcmfDicEntryId** | Pointer to **string** |  | [optional] 
**N3IwfId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**WagfId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**TngfId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**AnN2ApId** | Pointer to **int32** |  | [optional] 
**NssaaStatusList** | Pointer to [**[]NssaaStatus**](NssaaStatus.md) |  | [optional] 
**PendingNssaiMappingList** | Pointer to [**[]NssaiMapping**](NssaiMapping.md) |  | [optional] 

## Methods

### NewMmContext

`func NewMmContext(accessType AccessType, ) *MmContext`

NewMmContext instantiates a new MmContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMmContextWithDefaults

`func NewMmContextWithDefaults() *MmContext`

NewMmContextWithDefaults instantiates a new MmContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *MmContext) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MmContext) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MmContext) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetNasSecurityMode

`func (o *MmContext) GetNasSecurityMode() NasSecurityMode`

GetNasSecurityMode returns the NasSecurityMode field if non-nil, zero value otherwise.

### GetNasSecurityModeOk

`func (o *MmContext) GetNasSecurityModeOk() (*NasSecurityMode, bool)`

GetNasSecurityModeOk returns a tuple with the NasSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasSecurityMode

`func (o *MmContext) SetNasSecurityMode(v NasSecurityMode)`

SetNasSecurityMode sets NasSecurityMode field to given value.

### HasNasSecurityMode

`func (o *MmContext) HasNasSecurityMode() bool`

HasNasSecurityMode returns a boolean if a field has been set.

### GetEpsNasSecurityMode

`func (o *MmContext) GetEpsNasSecurityMode() EpsNasSecurityMode`

GetEpsNasSecurityMode returns the EpsNasSecurityMode field if non-nil, zero value otherwise.

### GetEpsNasSecurityModeOk

`func (o *MmContext) GetEpsNasSecurityModeOk() (*EpsNasSecurityMode, bool)`

GetEpsNasSecurityModeOk returns a tuple with the EpsNasSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsNasSecurityMode

`func (o *MmContext) SetEpsNasSecurityMode(v EpsNasSecurityMode)`

SetEpsNasSecurityMode sets EpsNasSecurityMode field to given value.

### HasEpsNasSecurityMode

`func (o *MmContext) HasEpsNasSecurityMode() bool`

HasEpsNasSecurityMode returns a boolean if a field has been set.

### GetNasDownlinkCount

`func (o *MmContext) GetNasDownlinkCount() int32`

GetNasDownlinkCount returns the NasDownlinkCount field if non-nil, zero value otherwise.

### GetNasDownlinkCountOk

`func (o *MmContext) GetNasDownlinkCountOk() (*int32, bool)`

GetNasDownlinkCountOk returns a tuple with the NasDownlinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasDownlinkCount

`func (o *MmContext) SetNasDownlinkCount(v int32)`

SetNasDownlinkCount sets NasDownlinkCount field to given value.

### HasNasDownlinkCount

`func (o *MmContext) HasNasDownlinkCount() bool`

HasNasDownlinkCount returns a boolean if a field has been set.

### GetNasUplinkCount

`func (o *MmContext) GetNasUplinkCount() int32`

GetNasUplinkCount returns the NasUplinkCount field if non-nil, zero value otherwise.

### GetNasUplinkCountOk

`func (o *MmContext) GetNasUplinkCountOk() (*int32, bool)`

GetNasUplinkCountOk returns a tuple with the NasUplinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasUplinkCount

`func (o *MmContext) SetNasUplinkCount(v int32)`

SetNasUplinkCount sets NasUplinkCount field to given value.

### HasNasUplinkCount

`func (o *MmContext) HasNasUplinkCount() bool`

HasNasUplinkCount returns a boolean if a field has been set.

### GetUeSecurityCapability

`func (o *MmContext) GetUeSecurityCapability() string`

GetUeSecurityCapability returns the UeSecurityCapability field if non-nil, zero value otherwise.

### GetUeSecurityCapabilityOk

`func (o *MmContext) GetUeSecurityCapabilityOk() (*string, bool)`

GetUeSecurityCapabilityOk returns a tuple with the UeSecurityCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSecurityCapability

`func (o *MmContext) SetUeSecurityCapability(v string)`

SetUeSecurityCapability sets UeSecurityCapability field to given value.

### HasUeSecurityCapability

`func (o *MmContext) HasUeSecurityCapability() bool`

HasUeSecurityCapability returns a boolean if a field has been set.

### GetS1UeNetworkCapability

`func (o *MmContext) GetS1UeNetworkCapability() string`

GetS1UeNetworkCapability returns the S1UeNetworkCapability field if non-nil, zero value otherwise.

### GetS1UeNetworkCapabilityOk

`func (o *MmContext) GetS1UeNetworkCapabilityOk() (*string, bool)`

GetS1UeNetworkCapabilityOk returns a tuple with the S1UeNetworkCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS1UeNetworkCapability

`func (o *MmContext) SetS1UeNetworkCapability(v string)`

SetS1UeNetworkCapability sets S1UeNetworkCapability field to given value.

### HasS1UeNetworkCapability

`func (o *MmContext) HasS1UeNetworkCapability() bool`

HasS1UeNetworkCapability returns a boolean if a field has been set.

### GetAllowedNssai

`func (o *MmContext) GetAllowedNssai() []Snssai`

GetAllowedNssai returns the AllowedNssai field if non-nil, zero value otherwise.

### GetAllowedNssaiOk

`func (o *MmContext) GetAllowedNssaiOk() (*[]Snssai, bool)`

GetAllowedNssaiOk returns a tuple with the AllowedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssai

`func (o *MmContext) SetAllowedNssai(v []Snssai)`

SetAllowedNssai sets AllowedNssai field to given value.

### HasAllowedNssai

`func (o *MmContext) HasAllowedNssai() bool`

HasAllowedNssai returns a boolean if a field has been set.

### GetNssaiMappingList

`func (o *MmContext) GetNssaiMappingList() []NssaiMapping`

GetNssaiMappingList returns the NssaiMappingList field if non-nil, zero value otherwise.

### GetNssaiMappingListOk

`func (o *MmContext) GetNssaiMappingListOk() (*[]NssaiMapping, bool)`

GetNssaiMappingListOk returns a tuple with the NssaiMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaiMappingList

`func (o *MmContext) SetNssaiMappingList(v []NssaiMapping)`

SetNssaiMappingList sets NssaiMappingList field to given value.

### HasNssaiMappingList

`func (o *MmContext) HasNssaiMappingList() bool`

HasNssaiMappingList returns a boolean if a field has been set.

### GetAllowedHomeNssai

`func (o *MmContext) GetAllowedHomeNssai() []Snssai`

GetAllowedHomeNssai returns the AllowedHomeNssai field if non-nil, zero value otherwise.

### GetAllowedHomeNssaiOk

`func (o *MmContext) GetAllowedHomeNssaiOk() (*[]Snssai, bool)`

GetAllowedHomeNssaiOk returns a tuple with the AllowedHomeNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHomeNssai

`func (o *MmContext) SetAllowedHomeNssai(v []Snssai)`

SetAllowedHomeNssai sets AllowedHomeNssai field to given value.

### HasAllowedHomeNssai

`func (o *MmContext) HasAllowedHomeNssai() bool`

HasAllowedHomeNssai returns a boolean if a field has been set.

### GetNsInstanceList

`func (o *MmContext) GetNsInstanceList() []string`

GetNsInstanceList returns the NsInstanceList field if non-nil, zero value otherwise.

### GetNsInstanceListOk

`func (o *MmContext) GetNsInstanceListOk() (*[]string, bool)`

GetNsInstanceListOk returns a tuple with the NsInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsInstanceList

`func (o *MmContext) SetNsInstanceList(v []string)`

SetNsInstanceList sets NsInstanceList field to given value.

### HasNsInstanceList

`func (o *MmContext) HasNsInstanceList() bool`

HasNsInstanceList returns a boolean if a field has been set.

### GetExpectedUEbehavior

`func (o *MmContext) GetExpectedUEbehavior() ExpectedUeBehavior`

GetExpectedUEbehavior returns the ExpectedUEbehavior field if non-nil, zero value otherwise.

### GetExpectedUEbehaviorOk

`func (o *MmContext) GetExpectedUEbehaviorOk() (*ExpectedUeBehavior, bool)`

GetExpectedUEbehaviorOk returns a tuple with the ExpectedUEbehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUEbehavior

`func (o *MmContext) SetExpectedUEbehavior(v ExpectedUeBehavior)`

SetExpectedUEbehavior sets ExpectedUEbehavior field to given value.

### HasExpectedUEbehavior

`func (o *MmContext) HasExpectedUEbehavior() bool`

HasExpectedUEbehavior returns a boolean if a field has been set.

### GetUeDifferentiationInfo

`func (o *MmContext) GetUeDifferentiationInfo() UeDifferentiationInfo`

GetUeDifferentiationInfo returns the UeDifferentiationInfo field if non-nil, zero value otherwise.

### GetUeDifferentiationInfoOk

`func (o *MmContext) GetUeDifferentiationInfoOk() (*UeDifferentiationInfo, bool)`

GetUeDifferentiationInfoOk returns a tuple with the UeDifferentiationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeDifferentiationInfo

`func (o *MmContext) SetUeDifferentiationInfo(v UeDifferentiationInfo)`

SetUeDifferentiationInfo sets UeDifferentiationInfo field to given value.

### HasUeDifferentiationInfo

`func (o *MmContext) HasUeDifferentiationInfo() bool`

HasUeDifferentiationInfo returns a boolean if a field has been set.

### GetPlmnAssiUeRadioCapId

`func (o *MmContext) GetPlmnAssiUeRadioCapId() string`

GetPlmnAssiUeRadioCapId returns the PlmnAssiUeRadioCapId field if non-nil, zero value otherwise.

### GetPlmnAssiUeRadioCapIdOk

`func (o *MmContext) GetPlmnAssiUeRadioCapIdOk() (*string, bool)`

GetPlmnAssiUeRadioCapIdOk returns a tuple with the PlmnAssiUeRadioCapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnAssiUeRadioCapId

`func (o *MmContext) SetPlmnAssiUeRadioCapId(v string)`

SetPlmnAssiUeRadioCapId sets PlmnAssiUeRadioCapId field to given value.

### HasPlmnAssiUeRadioCapId

`func (o *MmContext) HasPlmnAssiUeRadioCapId() bool`

HasPlmnAssiUeRadioCapId returns a boolean if a field has been set.

### GetManAssiUeRadioCapId

`func (o *MmContext) GetManAssiUeRadioCapId() string`

GetManAssiUeRadioCapId returns the ManAssiUeRadioCapId field if non-nil, zero value otherwise.

### GetManAssiUeRadioCapIdOk

`func (o *MmContext) GetManAssiUeRadioCapIdOk() (*string, bool)`

GetManAssiUeRadioCapIdOk returns a tuple with the ManAssiUeRadioCapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManAssiUeRadioCapId

`func (o *MmContext) SetManAssiUeRadioCapId(v string)`

SetManAssiUeRadioCapId sets ManAssiUeRadioCapId field to given value.

### HasManAssiUeRadioCapId

`func (o *MmContext) HasManAssiUeRadioCapId() bool`

HasManAssiUeRadioCapId returns a boolean if a field has been set.

### GetUcmfDicEntryId

`func (o *MmContext) GetUcmfDicEntryId() string`

GetUcmfDicEntryId returns the UcmfDicEntryId field if non-nil, zero value otherwise.

### GetUcmfDicEntryIdOk

`func (o *MmContext) GetUcmfDicEntryIdOk() (*string, bool)`

GetUcmfDicEntryIdOk returns a tuple with the UcmfDicEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmfDicEntryId

`func (o *MmContext) SetUcmfDicEntryId(v string)`

SetUcmfDicEntryId sets UcmfDicEntryId field to given value.

### HasUcmfDicEntryId

`func (o *MmContext) HasUcmfDicEntryId() bool`

HasUcmfDicEntryId returns a boolean if a field has been set.

### GetN3IwfId

`func (o *MmContext) GetN3IwfId() GlobalRanNodeId`

GetN3IwfId returns the N3IwfId field if non-nil, zero value otherwise.

### GetN3IwfIdOk

`func (o *MmContext) GetN3IwfIdOk() (*GlobalRanNodeId, bool)`

GetN3IwfIdOk returns a tuple with the N3IwfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3IwfId

`func (o *MmContext) SetN3IwfId(v GlobalRanNodeId)`

SetN3IwfId sets N3IwfId field to given value.

### HasN3IwfId

`func (o *MmContext) HasN3IwfId() bool`

HasN3IwfId returns a boolean if a field has been set.

### GetWagfId

`func (o *MmContext) GetWagfId() GlobalRanNodeId`

GetWagfId returns the WagfId field if non-nil, zero value otherwise.

### GetWagfIdOk

`func (o *MmContext) GetWagfIdOk() (*GlobalRanNodeId, bool)`

GetWagfIdOk returns a tuple with the WagfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagfId

`func (o *MmContext) SetWagfId(v GlobalRanNodeId)`

SetWagfId sets WagfId field to given value.

### HasWagfId

`func (o *MmContext) HasWagfId() bool`

HasWagfId returns a boolean if a field has been set.

### GetTngfId

`func (o *MmContext) GetTngfId() GlobalRanNodeId`

GetTngfId returns the TngfId field if non-nil, zero value otherwise.

### GetTngfIdOk

`func (o *MmContext) GetTngfIdOk() (*GlobalRanNodeId, bool)`

GetTngfIdOk returns a tuple with the TngfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfId

`func (o *MmContext) SetTngfId(v GlobalRanNodeId)`

SetTngfId sets TngfId field to given value.

### HasTngfId

`func (o *MmContext) HasTngfId() bool`

HasTngfId returns a boolean if a field has been set.

### GetAnN2ApId

`func (o *MmContext) GetAnN2ApId() int32`

GetAnN2ApId returns the AnN2ApId field if non-nil, zero value otherwise.

### GetAnN2ApIdOk

`func (o *MmContext) GetAnN2ApIdOk() (*int32, bool)`

GetAnN2ApIdOk returns a tuple with the AnN2ApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2ApId

`func (o *MmContext) SetAnN2ApId(v int32)`

SetAnN2ApId sets AnN2ApId field to given value.

### HasAnN2ApId

`func (o *MmContext) HasAnN2ApId() bool`

HasAnN2ApId returns a boolean if a field has been set.

### GetNssaaStatusList

`func (o *MmContext) GetNssaaStatusList() []NssaaStatus`

GetNssaaStatusList returns the NssaaStatusList field if non-nil, zero value otherwise.

### GetNssaaStatusListOk

`func (o *MmContext) GetNssaaStatusListOk() (*[]NssaaStatus, bool)`

GetNssaaStatusListOk returns a tuple with the NssaaStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaaStatusList

`func (o *MmContext) SetNssaaStatusList(v []NssaaStatus)`

SetNssaaStatusList sets NssaaStatusList field to given value.

### HasNssaaStatusList

`func (o *MmContext) HasNssaaStatusList() bool`

HasNssaaStatusList returns a boolean if a field has been set.

### GetPendingNssaiMappingList

`func (o *MmContext) GetPendingNssaiMappingList() []NssaiMapping`

GetPendingNssaiMappingList returns the PendingNssaiMappingList field if non-nil, zero value otherwise.

### GetPendingNssaiMappingListOk

`func (o *MmContext) GetPendingNssaiMappingListOk() (*[]NssaiMapping, bool)`

GetPendingNssaiMappingListOk returns a tuple with the PendingNssaiMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingNssaiMappingList

`func (o *MmContext) SetPendingNssaiMappingList(v []NssaiMapping)`

SetPendingNssaiMappingList sets PendingNssaiMappingList field to given value.

### HasPendingNssaiMappingList

`func (o *MmContext) HasPendingNssaiMappingList() bool`

HasPendingNssaiMappingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


