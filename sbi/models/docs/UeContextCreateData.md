# UeContextCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeContext** | [**UeContext**](UeContext.md) |  | 
**TargetId** | [**NgRanTargetId**](NgRanTargetId.md) |  | 
**SourceToTargetData** | [**N2InfoContent**](N2InfoContent.md) |  | 
**PduSessionList** | [**[]N2SmInformation**](N2SmInformation.md) |  | 
**N2NotifyUri** | Pointer to **string** |  | [optional] 
**UeRadioCapability** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**NgapCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 

## Methods

### NewUeContextCreateData

`func NewUeContextCreateData(ueContext UeContext, targetId NgRanTargetId, sourceToTargetData N2InfoContent, pduSessionList []N2SmInformation, ) *UeContextCreateData`

NewUeContextCreateData instantiates a new UeContextCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextCreateDataWithDefaults

`func NewUeContextCreateDataWithDefaults() *UeContextCreateData`

NewUeContextCreateDataWithDefaults instantiates a new UeContextCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeContext

`func (o *UeContextCreateData) GetUeContext() UeContext`

GetUeContext returns the UeContext field if non-nil, zero value otherwise.

### GetUeContextOk

`func (o *UeContextCreateData) GetUeContextOk() (*UeContext, bool)`

GetUeContextOk returns a tuple with the UeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContext

`func (o *UeContextCreateData) SetUeContext(v UeContext)`

SetUeContext sets UeContext field to given value.


### GetTargetId

`func (o *UeContextCreateData) GetTargetId() NgRanTargetId`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *UeContextCreateData) GetTargetIdOk() (*NgRanTargetId, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *UeContextCreateData) SetTargetId(v NgRanTargetId)`

SetTargetId sets TargetId field to given value.


### GetSourceToTargetData

`func (o *UeContextCreateData) GetSourceToTargetData() N2InfoContent`

GetSourceToTargetData returns the SourceToTargetData field if non-nil, zero value otherwise.

### GetSourceToTargetDataOk

`func (o *UeContextCreateData) GetSourceToTargetDataOk() (*N2InfoContent, bool)`

GetSourceToTargetDataOk returns a tuple with the SourceToTargetData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceToTargetData

`func (o *UeContextCreateData) SetSourceToTargetData(v N2InfoContent)`

SetSourceToTargetData sets SourceToTargetData field to given value.


### GetPduSessionList

`func (o *UeContextCreateData) GetPduSessionList() []N2SmInformation`

GetPduSessionList returns the PduSessionList field if non-nil, zero value otherwise.

### GetPduSessionListOk

`func (o *UeContextCreateData) GetPduSessionListOk() (*[]N2SmInformation, bool)`

GetPduSessionListOk returns a tuple with the PduSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionList

`func (o *UeContextCreateData) SetPduSessionList(v []N2SmInformation)`

SetPduSessionList sets PduSessionList field to given value.


### GetN2NotifyUri

`func (o *UeContextCreateData) GetN2NotifyUri() string`

GetN2NotifyUri returns the N2NotifyUri field if non-nil, zero value otherwise.

### GetN2NotifyUriOk

`func (o *UeContextCreateData) GetN2NotifyUriOk() (*string, bool)`

GetN2NotifyUriOk returns a tuple with the N2NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2NotifyUri

`func (o *UeContextCreateData) SetN2NotifyUri(v string)`

SetN2NotifyUri sets N2NotifyUri field to given value.

### HasN2NotifyUri

`func (o *UeContextCreateData) HasN2NotifyUri() bool`

HasN2NotifyUri returns a boolean if a field has been set.

### GetUeRadioCapability

`func (o *UeContextCreateData) GetUeRadioCapability() N2InfoContent`

GetUeRadioCapability returns the UeRadioCapability field if non-nil, zero value otherwise.

### GetUeRadioCapabilityOk

`func (o *UeContextCreateData) GetUeRadioCapabilityOk() (*N2InfoContent, bool)`

GetUeRadioCapabilityOk returns a tuple with the UeRadioCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability

`func (o *UeContextCreateData) SetUeRadioCapability(v N2InfoContent)`

SetUeRadioCapability sets UeRadioCapability field to given value.

### HasUeRadioCapability

`func (o *UeContextCreateData) HasUeRadioCapability() bool`

HasUeRadioCapability returns a boolean if a field has been set.

### GetNgapCause

`func (o *UeContextCreateData) GetNgapCause() NgApCause`

GetNgapCause returns the NgapCause field if non-nil, zero value otherwise.

### GetNgapCauseOk

`func (o *UeContextCreateData) GetNgapCauseOk() (*NgApCause, bool)`

GetNgapCauseOk returns a tuple with the NgapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapCause

`func (o *UeContextCreateData) SetNgapCause(v NgApCause)`

SetNgapCause sets NgapCause field to given value.

### HasNgapCause

`func (o *UeContextCreateData) HasNgapCause() bool`

HasNgapCause returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServingNetwork

`func (o *UeContextCreateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *UeContextCreateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *UeContextCreateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *UeContextCreateData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


