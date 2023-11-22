# AmfEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**SubsChangeNotifyCorrelationId** | Pointer to **string** |  | [optional] 
**ReportList** | Pointer to [**[]AmfEventReport**](AmfEventReport.md) |  | [optional] 
**EventSubsSyncInfo** | Pointer to [**AmfEventSubsSyncInfo**](AmfEventSubsSyncInfo.md) |  | [optional] 

## Methods

### NewAmfEventNotification

`func NewAmfEventNotification() *AmfEventNotification`

NewAmfEventNotification instantiates a new AmfEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventNotificationWithDefaults

`func NewAmfEventNotificationWithDefaults() *AmfEventNotification`

NewAmfEventNotificationWithDefaults instantiates a new AmfEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyCorrelationId

`func (o *AmfEventNotification) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *AmfEventNotification) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *AmfEventNotification) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *AmfEventNotification) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetSubsChangeNotifyCorrelationId

`func (o *AmfEventNotification) GetSubsChangeNotifyCorrelationId() string`

GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field if non-nil, zero value otherwise.

### GetSubsChangeNotifyCorrelationIdOk

`func (o *AmfEventNotification) GetSubsChangeNotifyCorrelationIdOk() (*string, bool)`

GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyCorrelationId

`func (o *AmfEventNotification) SetSubsChangeNotifyCorrelationId(v string)`

SetSubsChangeNotifyCorrelationId sets SubsChangeNotifyCorrelationId field to given value.

### HasSubsChangeNotifyCorrelationId

`func (o *AmfEventNotification) HasSubsChangeNotifyCorrelationId() bool`

HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.

### GetReportList

`func (o *AmfEventNotification) GetReportList() []AmfEventReport`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *AmfEventNotification) GetReportListOk() (*[]AmfEventReport, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *AmfEventNotification) SetReportList(v []AmfEventReport)`

SetReportList sets ReportList field to given value.

### HasReportList

`func (o *AmfEventNotification) HasReportList() bool`

HasReportList returns a boolean if a field has been set.

### GetEventSubsSyncInfo

`func (o *AmfEventNotification) GetEventSubsSyncInfo() AmfEventSubsSyncInfo`

GetEventSubsSyncInfo returns the EventSubsSyncInfo field if non-nil, zero value otherwise.

### GetEventSubsSyncInfoOk

`func (o *AmfEventNotification) GetEventSubsSyncInfoOk() (*AmfEventSubsSyncInfo, bool)`

GetEventSubsSyncInfoOk returns a tuple with the EventSubsSyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubsSyncInfo

`func (o *AmfEventNotification) SetEventSubsSyncInfo(v AmfEventSubsSyncInfo)`

SetEventSubsSyncInfo sets EventSubsSyncInfo field to given value.

### HasEventSubsSyncInfo

`func (o *AmfEventNotification) HasEventSubsSyncInfo() bool`

HasEventSubsSyncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


