# MonitoringReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **int32** |  | 
**EventType** | [**EventType**](EventType.md) |  | 
**Report** | Pointer to [**Report**](Report.md) |  | [optional] 
**ReachabilityForSmsReport** | Pointer to [**ReachabilityForSmsReport**](ReachabilityForSmsReport.md) |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**TimeStamp** | **time.Time** |  | 

## Methods

### NewMonitoringReport

`func NewMonitoringReport(referenceId int32, eventType EventType, timeStamp time.Time, ) *MonitoringReport`

NewMonitoringReport instantiates a new MonitoringReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringReportWithDefaults

`func NewMonitoringReportWithDefaults() *MonitoringReport`

NewMonitoringReportWithDefaults instantiates a new MonitoringReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *MonitoringReport) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MonitoringReport) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MonitoringReport) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetEventType

`func (o *MonitoringReport) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitoringReport) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitoringReport) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetReport

`func (o *MonitoringReport) GetReport() Report`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *MonitoringReport) GetReportOk() (*Report, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *MonitoringReport) SetReport(v Report)`

SetReport sets Report field to given value.

### HasReport

`func (o *MonitoringReport) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetReachabilityForSmsReport

`func (o *MonitoringReport) GetReachabilityForSmsReport() ReachabilityForSmsReport`

GetReachabilityForSmsReport returns the ReachabilityForSmsReport field if non-nil, zero value otherwise.

### GetReachabilityForSmsReportOk

`func (o *MonitoringReport) GetReachabilityForSmsReportOk() (*ReachabilityForSmsReport, bool)`

GetReachabilityForSmsReportOk returns a tuple with the ReachabilityForSmsReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityForSmsReport

`func (o *MonitoringReport) SetReachabilityForSmsReport(v ReachabilityForSmsReport)`

SetReachabilityForSmsReport sets ReachabilityForSmsReport field to given value.

### HasReachabilityForSmsReport

`func (o *MonitoringReport) HasReachabilityForSmsReport() bool`

HasReachabilityForSmsReport returns a boolean if a field has been set.

### GetGpsi

`func (o *MonitoringReport) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *MonitoringReport) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *MonitoringReport) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *MonitoringReport) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetTimeStamp

`func (o *MonitoringReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *MonitoringReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *MonitoringReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


