# EventsSubscReqDataRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]AfEventSubscription**](AfEventSubscription.md) |  | 
**NotifUri** | Pointer to **string** |  | [optional] 
**ReqQosMonParams** | Pointer to [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) |  | [optional] 
**QosMon** | Pointer to [**QosMonitoringInformationRm**](QosMonitoringInformationRm.md) |  | [optional] 
**ReqAnis** | Pointer to [**[]RequiredAccessInfo**](RequiredAccessInfo.md) |  | [optional] 
**UsgThres** | Pointer to [**UsageThresholdRm**](UsageThresholdRm.md) |  | [optional] 
**NotifCorreId** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsSubscReqDataRm

`func NewEventsSubscReqDataRm(events []AfEventSubscription, ) *EventsSubscReqDataRm`

NewEventsSubscReqDataRm instantiates a new EventsSubscReqDataRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSubscReqDataRmWithDefaults

`func NewEventsSubscReqDataRmWithDefaults() *EventsSubscReqDataRm`

NewEventsSubscReqDataRmWithDefaults instantiates a new EventsSubscReqDataRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsSubscReqDataRm) GetEvents() []AfEventSubscription`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSubscReqDataRm) GetEventsOk() (*[]AfEventSubscription, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSubscReqDataRm) SetEvents(v []AfEventSubscription)`

SetEvents sets Events field to given value.


### GetNotifUri

`func (o *EventsSubscReqDataRm) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *EventsSubscReqDataRm) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *EventsSubscReqDataRm) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *EventsSubscReqDataRm) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetReqQosMonParams

`func (o *EventsSubscReqDataRm) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *EventsSubscReqDataRm) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *EventsSubscReqDataRm) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.

### HasReqQosMonParams

`func (o *EventsSubscReqDataRm) HasReqQosMonParams() bool`

HasReqQosMonParams returns a boolean if a field has been set.

### GetQosMon

`func (o *EventsSubscReqDataRm) GetQosMon() QosMonitoringInformationRm`

GetQosMon returns the QosMon field if non-nil, zero value otherwise.

### GetQosMonOk

`func (o *EventsSubscReqDataRm) GetQosMonOk() (*QosMonitoringInformationRm, bool)`

GetQosMonOk returns a tuple with the QosMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMon

`func (o *EventsSubscReqDataRm) SetQosMon(v QosMonitoringInformationRm)`

SetQosMon sets QosMon field to given value.

### HasQosMon

`func (o *EventsSubscReqDataRm) HasQosMon() bool`

HasQosMon returns a boolean if a field has been set.

### SetQosMonNil

`func (o *EventsSubscReqDataRm) SetQosMonNil(b bool)`

 SetQosMonNil sets the value for QosMon to be an explicit nil

### UnsetQosMon
`func (o *EventsSubscReqDataRm) UnsetQosMon()`

UnsetQosMon ensures that no value is present for QosMon, not even an explicit nil
### GetReqAnis

`func (o *EventsSubscReqDataRm) GetReqAnis() []RequiredAccessInfo`

GetReqAnis returns the ReqAnis field if non-nil, zero value otherwise.

### GetReqAnisOk

`func (o *EventsSubscReqDataRm) GetReqAnisOk() (*[]RequiredAccessInfo, bool)`

GetReqAnisOk returns a tuple with the ReqAnis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAnis

`func (o *EventsSubscReqDataRm) SetReqAnis(v []RequiredAccessInfo)`

SetReqAnis sets ReqAnis field to given value.

### HasReqAnis

`func (o *EventsSubscReqDataRm) HasReqAnis() bool`

HasReqAnis returns a boolean if a field has been set.

### GetUsgThres

`func (o *EventsSubscReqDataRm) GetUsgThres() UsageThresholdRm`

GetUsgThres returns the UsgThres field if non-nil, zero value otherwise.

### GetUsgThresOk

`func (o *EventsSubscReqDataRm) GetUsgThresOk() (*UsageThresholdRm, bool)`

GetUsgThresOk returns a tuple with the UsgThres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgThres

`func (o *EventsSubscReqDataRm) SetUsgThres(v UsageThresholdRm)`

SetUsgThres sets UsgThres field to given value.

### HasUsgThres

`func (o *EventsSubscReqDataRm) HasUsgThres() bool`

HasUsgThres returns a boolean if a field has been set.

### SetUsgThresNil

`func (o *EventsSubscReqDataRm) SetUsgThresNil(b bool)`

 SetUsgThresNil sets the value for UsgThres to be an explicit nil

### UnsetUsgThres
`func (o *EventsSubscReqDataRm) UnsetUsgThres()`

UnsetUsgThres ensures that no value is present for UsgThres, not even an explicit nil
### GetNotifCorreId

`func (o *EventsSubscReqDataRm) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *EventsSubscReqDataRm) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *EventsSubscReqDataRm) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.

### HasNotifCorreId

`func (o *EventsSubscReqDataRm) HasNotifCorreId() bool`

HasNotifCorreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


