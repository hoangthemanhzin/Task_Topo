# UeContextRelocateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeContext** | [**UeContext**](UeContext.md) |  | 
**TargetId** | [**NgRanTargetId**](NgRanTargetId.md) |  | 
**SourceToTargetData** | [**N2InfoContent**](N2InfoContent.md) |  | 
**ForwardRelocationRequest** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**PduSessionList** | Pointer to [**[]N2SmInformation**](N2SmInformation.md) |  | [optional] 
**UeRadioCapability** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**NgapCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewUeContextRelocateData

`func NewUeContextRelocateData(ueContext UeContext, targetId NgRanTargetId, sourceToTargetData N2InfoContent, forwardRelocationRequest RefToBinaryData, ) *UeContextRelocateData`

NewUeContextRelocateData instantiates a new UeContextRelocateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextRelocateDataWithDefaults

`func NewUeContextRelocateDataWithDefaults() *UeContextRelocateData`

NewUeContextRelocateDataWithDefaults instantiates a new UeContextRelocateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeContext

`func (o *UeContextRelocateData) GetUeContext() UeContext`

GetUeContext returns the UeContext field if non-nil, zero value otherwise.

### GetUeContextOk

`func (o *UeContextRelocateData) GetUeContextOk() (*UeContext, bool)`

GetUeContextOk returns a tuple with the UeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContext

`func (o *UeContextRelocateData) SetUeContext(v UeContext)`

SetUeContext sets UeContext field to given value.


### GetTargetId

`func (o *UeContextRelocateData) GetTargetId() NgRanTargetId`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *UeContextRelocateData) GetTargetIdOk() (*NgRanTargetId, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *UeContextRelocateData) SetTargetId(v NgRanTargetId)`

SetTargetId sets TargetId field to given value.


### GetSourceToTargetData

`func (o *UeContextRelocateData) GetSourceToTargetData() N2InfoContent`

GetSourceToTargetData returns the SourceToTargetData field if non-nil, zero value otherwise.

### GetSourceToTargetDataOk

`func (o *UeContextRelocateData) GetSourceToTargetDataOk() (*N2InfoContent, bool)`

GetSourceToTargetDataOk returns a tuple with the SourceToTargetData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceToTargetData

`func (o *UeContextRelocateData) SetSourceToTargetData(v N2InfoContent)`

SetSourceToTargetData sets SourceToTargetData field to given value.


### GetForwardRelocationRequest

`func (o *UeContextRelocateData) GetForwardRelocationRequest() RefToBinaryData`

GetForwardRelocationRequest returns the ForwardRelocationRequest field if non-nil, zero value otherwise.

### GetForwardRelocationRequestOk

`func (o *UeContextRelocateData) GetForwardRelocationRequestOk() (*RefToBinaryData, bool)`

GetForwardRelocationRequestOk returns a tuple with the ForwardRelocationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardRelocationRequest

`func (o *UeContextRelocateData) SetForwardRelocationRequest(v RefToBinaryData)`

SetForwardRelocationRequest sets ForwardRelocationRequest field to given value.


### GetPduSessionList

`func (o *UeContextRelocateData) GetPduSessionList() []N2SmInformation`

GetPduSessionList returns the PduSessionList field if non-nil, zero value otherwise.

### GetPduSessionListOk

`func (o *UeContextRelocateData) GetPduSessionListOk() (*[]N2SmInformation, bool)`

GetPduSessionListOk returns a tuple with the PduSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionList

`func (o *UeContextRelocateData) SetPduSessionList(v []N2SmInformation)`

SetPduSessionList sets PduSessionList field to given value.

### HasPduSessionList

`func (o *UeContextRelocateData) HasPduSessionList() bool`

HasPduSessionList returns a boolean if a field has been set.

### GetUeRadioCapability

`func (o *UeContextRelocateData) GetUeRadioCapability() N2InfoContent`

GetUeRadioCapability returns the UeRadioCapability field if non-nil, zero value otherwise.

### GetUeRadioCapabilityOk

`func (o *UeContextRelocateData) GetUeRadioCapabilityOk() (*N2InfoContent, bool)`

GetUeRadioCapabilityOk returns a tuple with the UeRadioCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability

`func (o *UeContextRelocateData) SetUeRadioCapability(v N2InfoContent)`

SetUeRadioCapability sets UeRadioCapability field to given value.

### HasUeRadioCapability

`func (o *UeContextRelocateData) HasUeRadioCapability() bool`

HasUeRadioCapability returns a boolean if a field has been set.

### GetNgapCause

`func (o *UeContextRelocateData) GetNgapCause() NgApCause`

GetNgapCause returns the NgapCause field if non-nil, zero value otherwise.

### GetNgapCauseOk

`func (o *UeContextRelocateData) GetNgapCauseOk() (*NgApCause, bool)`

GetNgapCauseOk returns a tuple with the NgapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapCause

`func (o *UeContextRelocateData) SetNgapCause(v NgApCause)`

SetNgapCause sets NgapCause field to given value.

### HasNgapCause

`func (o *UeContextRelocateData) HasNgapCause() bool`

HasNgapCause returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextRelocateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextRelocateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextRelocateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextRelocateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


