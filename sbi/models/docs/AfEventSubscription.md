# AfEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AfEvent**](AfEvent.md) |  | 
**NotifMethod** | Pointer to [**AfNotifMethod**](AfNotifMethod.md) |  | [optional] 
**RepPeriod** | Pointer to **int32** |  | [optional] 
**WaitTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewAfEventSubscription

`func NewAfEventSubscription(event AfEvent, ) *AfEventSubscription`

NewAfEventSubscription instantiates a new AfEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventSubscriptionWithDefaults

`func NewAfEventSubscriptionWithDefaults() *AfEventSubscription`

NewAfEventSubscriptionWithDefaults instantiates a new AfEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AfEventSubscription) GetEvent() AfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AfEventSubscription) GetEventOk() (*AfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AfEventSubscription) SetEvent(v AfEvent)`

SetEvent sets Event field to given value.


### GetNotifMethod

`func (o *AfEventSubscription) GetNotifMethod() AfNotifMethod`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *AfEventSubscription) GetNotifMethodOk() (*AfNotifMethod, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *AfEventSubscription) SetNotifMethod(v AfNotifMethod)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *AfEventSubscription) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetRepPeriod

`func (o *AfEventSubscription) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *AfEventSubscription) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *AfEventSubscription) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *AfEventSubscription) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetWaitTime

`func (o *AfEventSubscription) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *AfEventSubscription) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *AfEventSubscription) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *AfEventSubscription) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


