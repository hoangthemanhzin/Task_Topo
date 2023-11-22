# QosFlowAddModifyRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**Ebi** | Pointer to **int32** |  | [optional] 
**QosRules** | Pointer to **string** |  | [optional] 
**QosFlowDescription** | Pointer to **string** |  | [optional] 
**QosFlowProfile** | Pointer to [**QosFlowProfile**](QosFlowProfile.md) |  | [optional] 
**AssociatedAnType** | Pointer to [**QosFlowAccessType**](QosFlowAccessType.md) |  | [optional] 

## Methods

### NewQosFlowAddModifyRequestItem

`func NewQosFlowAddModifyRequestItem(qfi int32, ) *QosFlowAddModifyRequestItem`

NewQosFlowAddModifyRequestItem instantiates a new QosFlowAddModifyRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowAddModifyRequestItemWithDefaults

`func NewQosFlowAddModifyRequestItemWithDefaults() *QosFlowAddModifyRequestItem`

NewQosFlowAddModifyRequestItemWithDefaults instantiates a new QosFlowAddModifyRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowAddModifyRequestItem) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowAddModifyRequestItem) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowAddModifyRequestItem) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetEbi

`func (o *QosFlowAddModifyRequestItem) GetEbi() int32`

GetEbi returns the Ebi field if non-nil, zero value otherwise.

### GetEbiOk

`func (o *QosFlowAddModifyRequestItem) GetEbiOk() (*int32, bool)`

GetEbiOk returns a tuple with the Ebi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbi

`func (o *QosFlowAddModifyRequestItem) SetEbi(v int32)`

SetEbi sets Ebi field to given value.

### HasEbi

`func (o *QosFlowAddModifyRequestItem) HasEbi() bool`

HasEbi returns a boolean if a field has been set.

### GetQosRules

`func (o *QosFlowAddModifyRequestItem) GetQosRules() string`

GetQosRules returns the QosRules field if non-nil, zero value otherwise.

### GetQosRulesOk

`func (o *QosFlowAddModifyRequestItem) GetQosRulesOk() (*string, bool)`

GetQosRulesOk returns a tuple with the QosRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRules

`func (o *QosFlowAddModifyRequestItem) SetQosRules(v string)`

SetQosRules sets QosRules field to given value.

### HasQosRules

`func (o *QosFlowAddModifyRequestItem) HasQosRules() bool`

HasQosRules returns a boolean if a field has been set.

### GetQosFlowDescription

`func (o *QosFlowAddModifyRequestItem) GetQosFlowDescription() string`

GetQosFlowDescription returns the QosFlowDescription field if non-nil, zero value otherwise.

### GetQosFlowDescriptionOk

`func (o *QosFlowAddModifyRequestItem) GetQosFlowDescriptionOk() (*string, bool)`

GetQosFlowDescriptionOk returns a tuple with the QosFlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowDescription

`func (o *QosFlowAddModifyRequestItem) SetQosFlowDescription(v string)`

SetQosFlowDescription sets QosFlowDescription field to given value.

### HasQosFlowDescription

`func (o *QosFlowAddModifyRequestItem) HasQosFlowDescription() bool`

HasQosFlowDescription returns a boolean if a field has been set.

### GetQosFlowProfile

`func (o *QosFlowAddModifyRequestItem) GetQosFlowProfile() QosFlowProfile`

GetQosFlowProfile returns the QosFlowProfile field if non-nil, zero value otherwise.

### GetQosFlowProfileOk

`func (o *QosFlowAddModifyRequestItem) GetQosFlowProfileOk() (*QosFlowProfile, bool)`

GetQosFlowProfileOk returns a tuple with the QosFlowProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowProfile

`func (o *QosFlowAddModifyRequestItem) SetQosFlowProfile(v QosFlowProfile)`

SetQosFlowProfile sets QosFlowProfile field to given value.

### HasQosFlowProfile

`func (o *QosFlowAddModifyRequestItem) HasQosFlowProfile() bool`

HasQosFlowProfile returns a boolean if a field has been set.

### GetAssociatedAnType

`func (o *QosFlowAddModifyRequestItem) GetAssociatedAnType() QosFlowAccessType`

GetAssociatedAnType returns the AssociatedAnType field if non-nil, zero value otherwise.

### GetAssociatedAnTypeOk

`func (o *QosFlowAddModifyRequestItem) GetAssociatedAnTypeOk() (*QosFlowAccessType, bool)`

GetAssociatedAnTypeOk returns a tuple with the AssociatedAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAnType

`func (o *QosFlowAddModifyRequestItem) SetAssociatedAnType(v QosFlowAccessType)`

SetAssociatedAnType sets AssociatedAnType field to given value.

### HasAssociatedAnType

`func (o *QosFlowAddModifyRequestItem) HasAssociatedAnType() bool`

HasAssociatedAnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


