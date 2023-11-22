# AfEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AfEvent**](AfEvent.md) |  | 
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 

## Methods

### NewAfEventNotification

`func NewAfEventNotification(event AfEvent, ) *AfEventNotification`

NewAfEventNotification instantiates a new AfEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventNotificationWithDefaults

`func NewAfEventNotificationWithDefaults() *AfEventNotification`

NewAfEventNotificationWithDefaults instantiates a new AfEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AfEventNotification) GetEvent() AfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AfEventNotification) GetEventOk() (*AfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AfEventNotification) SetEvent(v AfEvent)`

SetEvent sets Event field to given value.


### GetFlows

`func (o *AfEventNotification) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *AfEventNotification) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *AfEventNotification) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *AfEventNotification) HasFlows() bool`

HasFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


