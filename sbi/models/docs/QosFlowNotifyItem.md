# QosFlowNotifyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**NotificationCause** | [**NotificationCause**](NotificationCause.md) |  | 
**CurrentQosProfileIndex** | Pointer to **int32** |  | [optional] 
**NullQoSProfileIndex** | Pointer to **bool** |  | [optional] 

## Methods

### NewQosFlowNotifyItem

`func NewQosFlowNotifyItem(qfi int32, notificationCause NotificationCause, ) *QosFlowNotifyItem`

NewQosFlowNotifyItem instantiates a new QosFlowNotifyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowNotifyItemWithDefaults

`func NewQosFlowNotifyItemWithDefaults() *QosFlowNotifyItem`

NewQosFlowNotifyItemWithDefaults instantiates a new QosFlowNotifyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowNotifyItem) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowNotifyItem) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowNotifyItem) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetNotificationCause

`func (o *QosFlowNotifyItem) GetNotificationCause() NotificationCause`

GetNotificationCause returns the NotificationCause field if non-nil, zero value otherwise.

### GetNotificationCauseOk

`func (o *QosFlowNotifyItem) GetNotificationCauseOk() (*NotificationCause, bool)`

GetNotificationCauseOk returns a tuple with the NotificationCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCause

`func (o *QosFlowNotifyItem) SetNotificationCause(v NotificationCause)`

SetNotificationCause sets NotificationCause field to given value.


### GetCurrentQosProfileIndex

`func (o *QosFlowNotifyItem) GetCurrentQosProfileIndex() int32`

GetCurrentQosProfileIndex returns the CurrentQosProfileIndex field if non-nil, zero value otherwise.

### GetCurrentQosProfileIndexOk

`func (o *QosFlowNotifyItem) GetCurrentQosProfileIndexOk() (*int32, bool)`

GetCurrentQosProfileIndexOk returns a tuple with the CurrentQosProfileIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentQosProfileIndex

`func (o *QosFlowNotifyItem) SetCurrentQosProfileIndex(v int32)`

SetCurrentQosProfileIndex sets CurrentQosProfileIndex field to given value.

### HasCurrentQosProfileIndex

`func (o *QosFlowNotifyItem) HasCurrentQosProfileIndex() bool`

HasCurrentQosProfileIndex returns a boolean if a field has been set.

### GetNullQoSProfileIndex

`func (o *QosFlowNotifyItem) GetNullQoSProfileIndex() bool`

GetNullQoSProfileIndex returns the NullQoSProfileIndex field if non-nil, zero value otherwise.

### GetNullQoSProfileIndexOk

`func (o *QosFlowNotifyItem) GetNullQoSProfileIndexOk() (*bool, bool)`

GetNullQoSProfileIndexOk returns a tuple with the NullQoSProfileIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullQoSProfileIndex

`func (o *QosFlowNotifyItem) SetNullQoSProfileIndex(v bool)`

SetNullQoSProfileIndex sets NullQoSProfileIndex field to given value.

### HasNullQoSProfileIndex

`func (o *QosFlowNotifyItem) HasNullQoSProfileIndex() bool`

HasNullQoSProfileIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


