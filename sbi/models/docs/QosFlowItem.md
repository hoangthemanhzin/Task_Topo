# QosFlowItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**CurrentQosProfileIndex** | Pointer to **int32** |  | [optional] 
**NullQoSProfileIndex** | Pointer to **bool** |  | [optional] 

## Methods

### NewQosFlowItem

`func NewQosFlowItem(qfi int32, ) *QosFlowItem`

NewQosFlowItem instantiates a new QosFlowItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowItemWithDefaults

`func NewQosFlowItemWithDefaults() *QosFlowItem`

NewQosFlowItemWithDefaults instantiates a new QosFlowItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowItem) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowItem) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowItem) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetCause

`func (o *QosFlowItem) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *QosFlowItem) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *QosFlowItem) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *QosFlowItem) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCurrentQosProfileIndex

`func (o *QosFlowItem) GetCurrentQosProfileIndex() int32`

GetCurrentQosProfileIndex returns the CurrentQosProfileIndex field if non-nil, zero value otherwise.

### GetCurrentQosProfileIndexOk

`func (o *QosFlowItem) GetCurrentQosProfileIndexOk() (*int32, bool)`

GetCurrentQosProfileIndexOk returns a tuple with the CurrentQosProfileIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentQosProfileIndex

`func (o *QosFlowItem) SetCurrentQosProfileIndex(v int32)`

SetCurrentQosProfileIndex sets CurrentQosProfileIndex field to given value.

### HasCurrentQosProfileIndex

`func (o *QosFlowItem) HasCurrentQosProfileIndex() bool`

HasCurrentQosProfileIndex returns a boolean if a field has been set.

### GetNullQoSProfileIndex

`func (o *QosFlowItem) GetNullQoSProfileIndex() bool`

GetNullQoSProfileIndex returns the NullQoSProfileIndex field if non-nil, zero value otherwise.

### GetNullQoSProfileIndexOk

`func (o *QosFlowItem) GetNullQoSProfileIndexOk() (*bool, bool)`

GetNullQoSProfileIndexOk returns a tuple with the NullQoSProfileIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullQoSProfileIndex

`func (o *QosFlowItem) SetNullQoSProfileIndex(v bool)`

SetNullQoSProfileIndex sets NullQoSProfileIndex field to given value.

### HasNullQoSProfileIndex

`func (o *QosFlowItem) HasNullQoSProfileIndex() bool`

HasNullQoSProfileIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


