# EventsSubscReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]AfEventSubscription**](AfEventSubscription.md) |  | 
**NotifUri** | Pointer to **string** |  | [optional] 
**ReqQosMonParams** | Pointer to [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) |  | [optional] 
**QosMon** | Pointer to [**QosMonitoringInformation**](QosMonitoringInformation.md) |  | [optional] 
**ReqAnis** | Pointer to [**[]RequiredAccessInfo**](RequiredAccessInfo.md) |  | [optional] 
**UsgThres** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**NotifCorreId** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsSubscReqData

`func NewEventsSubscReqData(events []AfEventSubscription, ) *EventsSubscReqData`

NewEventsSubscReqData instantiates a new EventsSubscReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSubscReqDataWithDefaults

`func NewEventsSubscReqDataWithDefaults() *EventsSubscReqData`

NewEventsSubscReqDataWithDefaults instantiates a new EventsSubscReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsSubscReqData) GetEvents() []AfEventSubscription`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSubscReqData) GetEventsOk() (*[]AfEventSubscription, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSubscReqData) SetEvents(v []AfEventSubscription)`

SetEvents sets Events field to given value.


### GetNotifUri

`func (o *EventsSubscReqData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *EventsSubscReqData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *EventsSubscReqData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *EventsSubscReqData) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetReqQosMonParams

`func (o *EventsSubscReqData) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *EventsSubscReqData) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *EventsSubscReqData) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.

### HasReqQosMonParams

`func (o *EventsSubscReqData) HasReqQosMonParams() bool`

HasReqQosMonParams returns a boolean if a field has been set.

### GetQosMon

`func (o *EventsSubscReqData) GetQosMon() QosMonitoringInformation`

GetQosMon returns the QosMon field if non-nil, zero value otherwise.

### GetQosMonOk

`func (o *EventsSubscReqData) GetQosMonOk() (*QosMonitoringInformation, bool)`

GetQosMonOk returns a tuple with the QosMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMon

`func (o *EventsSubscReqData) SetQosMon(v QosMonitoringInformation)`

SetQosMon sets QosMon field to given value.

### HasQosMon

`func (o *EventsSubscReqData) HasQosMon() bool`

HasQosMon returns a boolean if a field has been set.

### GetReqAnis

`func (o *EventsSubscReqData) GetReqAnis() []RequiredAccessInfo`

GetReqAnis returns the ReqAnis field if non-nil, zero value otherwise.

### GetReqAnisOk

`func (o *EventsSubscReqData) GetReqAnisOk() (*[]RequiredAccessInfo, bool)`

GetReqAnisOk returns a tuple with the ReqAnis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAnis

`func (o *EventsSubscReqData) SetReqAnis(v []RequiredAccessInfo)`

SetReqAnis sets ReqAnis field to given value.

### HasReqAnis

`func (o *EventsSubscReqData) HasReqAnis() bool`

HasReqAnis returns a boolean if a field has been set.

### GetUsgThres

`func (o *EventsSubscReqData) GetUsgThres() UsageThreshold`

GetUsgThres returns the UsgThres field if non-nil, zero value otherwise.

### GetUsgThresOk

`func (o *EventsSubscReqData) GetUsgThresOk() (*UsageThreshold, bool)`

GetUsgThresOk returns a tuple with the UsgThres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgThres

`func (o *EventsSubscReqData) SetUsgThres(v UsageThreshold)`

SetUsgThres sets UsgThres field to given value.

### HasUsgThres

`func (o *EventsSubscReqData) HasUsgThres() bool`

HasUsgThres returns a boolean if a field has been set.

### GetNotifCorreId

`func (o *EventsSubscReqData) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *EventsSubscReqData) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *EventsSubscReqData) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.

### HasNotifCorreId

`func (o *EventsSubscReqData) HasNotifCorreId() bool`

HasNotifCorreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


