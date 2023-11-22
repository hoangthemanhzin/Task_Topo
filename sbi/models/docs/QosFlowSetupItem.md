# QosFlowSetupItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**QosRules** | **string** |  | 
**Ebi** | Pointer to **int32** |  | [optional] 
**QosFlowDescription** | Pointer to **string** |  | [optional] 
**QosFlowProfile** | Pointer to [**QosFlowProfile**](QosFlowProfile.md) |  | [optional] 
**AssociatedAnType** | Pointer to [**QosFlowAccessType**](QosFlowAccessType.md) |  | [optional] 
**DefaultQosRuleInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewQosFlowSetupItem

`func NewQosFlowSetupItem(qfi int32, qosRules string, ) *QosFlowSetupItem`

NewQosFlowSetupItem instantiates a new QosFlowSetupItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowSetupItemWithDefaults

`func NewQosFlowSetupItemWithDefaults() *QosFlowSetupItem`

NewQosFlowSetupItemWithDefaults instantiates a new QosFlowSetupItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowSetupItem) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowSetupItem) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowSetupItem) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetQosRules

`func (o *QosFlowSetupItem) GetQosRules() string`

GetQosRules returns the QosRules field if non-nil, zero value otherwise.

### GetQosRulesOk

`func (o *QosFlowSetupItem) GetQosRulesOk() (*string, bool)`

GetQosRulesOk returns a tuple with the QosRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRules

`func (o *QosFlowSetupItem) SetQosRules(v string)`

SetQosRules sets QosRules field to given value.


### GetEbi

`func (o *QosFlowSetupItem) GetEbi() int32`

GetEbi returns the Ebi field if non-nil, zero value otherwise.

### GetEbiOk

`func (o *QosFlowSetupItem) GetEbiOk() (*int32, bool)`

GetEbiOk returns a tuple with the Ebi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbi

`func (o *QosFlowSetupItem) SetEbi(v int32)`

SetEbi sets Ebi field to given value.

### HasEbi

`func (o *QosFlowSetupItem) HasEbi() bool`

HasEbi returns a boolean if a field has been set.

### GetQosFlowDescription

`func (o *QosFlowSetupItem) GetQosFlowDescription() string`

GetQosFlowDescription returns the QosFlowDescription field if non-nil, zero value otherwise.

### GetQosFlowDescriptionOk

`func (o *QosFlowSetupItem) GetQosFlowDescriptionOk() (*string, bool)`

GetQosFlowDescriptionOk returns a tuple with the QosFlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowDescription

`func (o *QosFlowSetupItem) SetQosFlowDescription(v string)`

SetQosFlowDescription sets QosFlowDescription field to given value.

### HasQosFlowDescription

`func (o *QosFlowSetupItem) HasQosFlowDescription() bool`

HasQosFlowDescription returns a boolean if a field has been set.

### GetQosFlowProfile

`func (o *QosFlowSetupItem) GetQosFlowProfile() QosFlowProfile`

GetQosFlowProfile returns the QosFlowProfile field if non-nil, zero value otherwise.

### GetQosFlowProfileOk

`func (o *QosFlowSetupItem) GetQosFlowProfileOk() (*QosFlowProfile, bool)`

GetQosFlowProfileOk returns a tuple with the QosFlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowProfile

`func (o *QosFlowSetupItem) SetQosFlowProfile(v QosFlowProfile)`

SetQosFlowProfile sets QosFlowProfile field to given value.

### HasQosFlowProfile

`func (o *QosFlowSetupItem) HasQosFlowProfile() bool`

HasQosFlowProfile returns a boolean if a field has been set.

### GetAssociatedAnType

`func (o *QosFlowSetupItem) GetAssociatedAnType() QosFlowAccessType`

GetAssociatedAnType returns the AssociatedAnType field if non-nil, zero value otherwise.

### GetAssociatedAnTypeOk

`func (o *QosFlowSetupItem) GetAssociatedAnTypeOk() (*QosFlowAccessType, bool)`

GetAssociatedAnTypeOk returns a tuple with the AssociatedAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAnType

`func (o *QosFlowSetupItem) SetAssociatedAnType(v QosFlowAccessType)`

SetAssociatedAnType sets AssociatedAnType field to given value.

### HasAssociatedAnType

`func (o *QosFlowSetupItem) HasAssociatedAnType() bool`

HasAssociatedAnType returns a boolean if a field has been set.

### GetDefaultQosRuleInd

`func (o *QosFlowSetupItem) GetDefaultQosRuleInd() bool`

GetDefaultQosRuleInd returns the DefaultQosRuleInd field if non-nil, zero value otherwise.

### GetDefaultQosRuleIndOk

`func (o *QosFlowSetupItem) GetDefaultQosRuleIndOk() (*bool, bool)`

GetDefaultQosRuleIndOk returns a tuple with the DefaultQosRuleInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQosRuleInd

`func (o *QosFlowSetupItem) SetDefaultQosRuleInd(v bool)`

SetDefaultQosRuleInd sets DefaultQosRuleInd field to given value.

### HasDefaultQosRuleInd

`func (o *QosFlowSetupItem) HasDefaultQosRuleInd() bool`

HasDefaultQosRuleInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


