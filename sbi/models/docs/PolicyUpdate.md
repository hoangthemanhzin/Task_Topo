# PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** |  | 
**UePolicy** | Pointer to **string** |  | [optional] 
**N2Pc5Pol** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the PCF subscribes. Only values \&quot;LOC_CH\&quot; and \&quot;PRA_CH\&quot; are permitted. | [optional] 
**Pras** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Map of PRA information. | [optional] 

## Methods

### NewPolicyUpdate

`func NewPolicyUpdate(resourceUri string, ) *PolicyUpdate`

NewPolicyUpdate instantiates a new PolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateWithDefaults

`func NewPolicyUpdateWithDefaults() *PolicyUpdate`

NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *PolicyUpdate) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *PolicyUpdate) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *PolicyUpdate) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetUePolicy

`func (o *PolicyUpdate) GetUePolicy() string`

GetUePolicy returns the UePolicy field if non-nil, zero value otherwise.

### GetUePolicyOk

`func (o *PolicyUpdate) GetUePolicyOk() (*string, bool)`

GetUePolicyOk returns a tuple with the UePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicy

`func (o *PolicyUpdate) SetUePolicy(v string)`

SetUePolicy sets UePolicy field to given value.

### HasUePolicy

`func (o *PolicyUpdate) HasUePolicy() bool`

HasUePolicy returns a boolean if a field has been set.

### GetN2Pc5Pol

`func (o *PolicyUpdate) GetN2Pc5Pol() N2InfoContent`

GetN2Pc5Pol returns the N2Pc5Pol field if non-nil, zero value otherwise.

### GetN2Pc5PolOk

`func (o *PolicyUpdate) GetN2Pc5PolOk() (*N2InfoContent, bool)`

GetN2Pc5PolOk returns a tuple with the N2Pc5Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2Pc5Pol

`func (o *PolicyUpdate) SetN2Pc5Pol(v N2InfoContent)`

SetN2Pc5Pol sets N2Pc5Pol field to given value.

### HasN2Pc5Pol

`func (o *PolicyUpdate) HasN2Pc5Pol() bool`

HasN2Pc5Pol returns a boolean if a field has been set.

### GetTriggers

`func (o *PolicyUpdate) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyUpdate) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyUpdate) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyUpdate) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *PolicyUpdate) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *PolicyUpdate) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetPras

`func (o *PolicyUpdate) GetPras() map[string]PresenceInfo`

GetPras returns the Pras field if non-nil, zero value otherwise.

### GetPrasOk

`func (o *PolicyUpdate) GetPrasOk() (*map[string]PresenceInfo, bool)`

GetPrasOk returns a tuple with the Pras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPras

`func (o *PolicyUpdate) SetPras(v map[string]PresenceInfo)`

SetPras sets Pras field to given value.

### HasPras

`func (o *PolicyUpdate) HasPras() bool`

HasPras returns a boolean if a field has been set.

### SetPrasNil

`func (o *PolicyUpdate) SetPrasNil(b bool)`

 SetPrasNil sets the value for Pras to be an explicit nil

### UnsetPras
`func (o *PolicyUpdate) UnsetPras()`

UnsetPras ensures that no value is present for Pras, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


