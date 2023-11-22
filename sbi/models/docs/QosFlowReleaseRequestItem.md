# QosFlowReleaseRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**QosRules** | Pointer to **string** |  | [optional] 
**QosFlowDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewQosFlowReleaseRequestItem

`func NewQosFlowReleaseRequestItem(qfi int32, ) *QosFlowReleaseRequestItem`

NewQosFlowReleaseRequestItem instantiates a new QosFlowReleaseRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowReleaseRequestItemWithDefaults

`func NewQosFlowReleaseRequestItemWithDefaults() *QosFlowReleaseRequestItem`

NewQosFlowReleaseRequestItemWithDefaults instantiates a new QosFlowReleaseRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowReleaseRequestItem) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowReleaseRequestItem) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowReleaseRequestItem) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetQosRules

`func (o *QosFlowReleaseRequestItem) GetQosRules() string`

GetQosRules returns the QosRules field if non-nil, zero value otherwise.

### GetQosRulesOk

`func (o *QosFlowReleaseRequestItem) GetQosRulesOk() (*string, bool)`

GetQosRulesOk returns a tuple with the QosRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRules

`func (o *QosFlowReleaseRequestItem) SetQosRules(v string)`

SetQosRules sets QosRules field to given value.

### HasQosRules

`func (o *QosFlowReleaseRequestItem) HasQosRules() bool`

HasQosRules returns a boolean if a field has been set.

### GetQosFlowDescription

`func (o *QosFlowReleaseRequestItem) GetQosFlowDescription() string`

GetQosFlowDescription returns the QosFlowDescription field if non-nil, zero value otherwise.

### GetQosFlowDescriptionOk

`func (o *QosFlowReleaseRequestItem) GetQosFlowDescriptionOk() (*string, bool)`

GetQosFlowDescriptionOk returns a tuple with the QosFlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowDescription

`func (o *QosFlowReleaseRequestItem) SetQosFlowDescription(v string)`

SetQosFlowDescription sets QosFlowDescription field to given value.

### HasQosFlowDescription

`func (o *QosFlowReleaseRequestItem) HasQosFlowDescription() bool`

HasQosFlowDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


