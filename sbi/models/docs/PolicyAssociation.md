# PolicyAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**PolicyAssociationRequest**](PolicyAssociationRequest.md) |  | [optional] 
**UePolicy** | Pointer to **string** |  | [optional] 
**N2Pc5Pol** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the PCF subscribes. Only values \&quot;LOC_CH\&quot; and \&quot;PRA_CH\&quot; are permitted. | [optional] 
**Pras** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**SuppFeat** | **string** |  | 

## Methods

### NewPolicyAssociation

`func NewPolicyAssociation(suppFeat string, ) *PolicyAssociation`

NewPolicyAssociation instantiates a new PolicyAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationWithDefaults

`func NewPolicyAssociationWithDefaults() *PolicyAssociation`

NewPolicyAssociationWithDefaults instantiates a new PolicyAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *PolicyAssociation) GetRequest() PolicyAssociationRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PolicyAssociation) GetRequestOk() (*PolicyAssociationRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PolicyAssociation) SetRequest(v PolicyAssociationRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PolicyAssociation) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetUePolicy

`func (o *PolicyAssociation) GetUePolicy() string`

GetUePolicy returns the UePolicy field if non-nil, zero value otherwise.

### GetUePolicyOk

`func (o *PolicyAssociation) GetUePolicyOk() (*string, bool)`

GetUePolicyOk returns a tuple with the UePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicy

`func (o *PolicyAssociation) SetUePolicy(v string)`

SetUePolicy sets UePolicy field to given value.

### HasUePolicy

`func (o *PolicyAssociation) HasUePolicy() bool`

HasUePolicy returns a boolean if a field has been set.

### GetN2Pc5Pol

`func (o *PolicyAssociation) GetN2Pc5Pol() N2InfoContent`

GetN2Pc5Pol returns the N2Pc5Pol field if non-nil, zero value otherwise.

### GetN2Pc5PolOk

`func (o *PolicyAssociation) GetN2Pc5PolOk() (*N2InfoContent, bool)`

GetN2Pc5PolOk returns a tuple with the N2Pc5Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2Pc5Pol

`func (o *PolicyAssociation) SetN2Pc5Pol(v N2InfoContent)`

SetN2Pc5Pol sets N2Pc5Pol field to given value.

### HasN2Pc5Pol

`func (o *PolicyAssociation) HasN2Pc5Pol() bool`

HasN2Pc5Pol returns a boolean if a field has been set.

### GetTriggers

`func (o *PolicyAssociation) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyAssociation) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyAssociation) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyAssociation) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetPras

`func (o *PolicyAssociation) GetPras() map[string]PresenceInfo`

GetPras returns the Pras field if non-nil, zero value otherwise.

### GetPrasOk

`func (o *PolicyAssociation) GetPrasOk() (*map[string]PresenceInfo, bool)`

GetPrasOk returns a tuple with the Pras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPras

`func (o *PolicyAssociation) SetPras(v map[string]PresenceInfo)`

SetPras sets Pras field to given value.

### HasPras

`func (o *PolicyAssociation) HasPras() bool`

HasPras returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PolicyAssociation) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PolicyAssociation) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PolicyAssociation) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


