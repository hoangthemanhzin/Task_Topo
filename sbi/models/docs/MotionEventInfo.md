# MotionEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinearDistance** | **int32** |  | 
**OccurrenceInfo** | Pointer to [**OccurrenceInfo**](OccurrenceInfo.md) |  | [optional] 
**MinimumInterval** | Pointer to **int32** |  | [optional] 
**MaximumInterval** | Pointer to **int32** |  | [optional] 
**SamplingInterval** | Pointer to **int32** |  | [optional] 
**ReportingDuration** | Pointer to **int32** |  | [optional] 
**ReportingLocationReq** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewMotionEventInfo

`func NewMotionEventInfo(linearDistance int32, ) *MotionEventInfo`

NewMotionEventInfo instantiates a new MotionEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotionEventInfoWithDefaults

`func NewMotionEventInfoWithDefaults() *MotionEventInfo`

NewMotionEventInfoWithDefaults instantiates a new MotionEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinearDistance

`func (o *MotionEventInfo) GetLinearDistance() int32`

GetLinearDistance returns the LinearDistance field if non-nil, zero value otherwise.

### GetLinearDistanceOk

`func (o *MotionEventInfo) GetLinearDistanceOk() (*int32, bool)`

GetLinearDistanceOk returns a tuple with the LinearDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearDistance

`func (o *MotionEventInfo) SetLinearDistance(v int32)`

SetLinearDistance sets LinearDistance field to given value.


### GetOccurrenceInfo

`func (o *MotionEventInfo) GetOccurrenceInfo() OccurrenceInfo`

GetOccurrenceInfo returns the OccurrenceInfo field if non-nil, zero value otherwise.

### GetOccurrenceInfoOk

`func (o *MotionEventInfo) GetOccurrenceInfoOk() (*OccurrenceInfo, bool)`

GetOccurrenceInfoOk returns a tuple with the OccurrenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceInfo

`func (o *MotionEventInfo) SetOccurrenceInfo(v OccurrenceInfo)`

SetOccurrenceInfo sets OccurrenceInfo field to given value.

### HasOccurrenceInfo

`func (o *MotionEventInfo) HasOccurrenceInfo() bool`

HasOccurrenceInfo returns a boolean if a field has been set.

### GetMinimumInterval

`func (o *MotionEventInfo) GetMinimumInterval() int32`

GetMinimumInterval returns the MinimumInterval field if non-nil, zero value otherwise.

### GetMinimumIntervalOk

`func (o *MotionEventInfo) GetMinimumIntervalOk() (*int32, bool)`

GetMinimumIntervalOk returns a tuple with the MinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInterval

`func (o *MotionEventInfo) SetMinimumInterval(v int32)`

SetMinimumInterval sets MinimumInterval field to given value.

### HasMinimumInterval

`func (o *MotionEventInfo) HasMinimumInterval() bool`

HasMinimumInterval returns a boolean if a field has been set.

### GetMaximumInterval

`func (o *MotionEventInfo) GetMaximumInterval() int32`

GetMaximumInterval returns the MaximumInterval field if non-nil, zero value otherwise.

### GetMaximumIntervalOk

`func (o *MotionEventInfo) GetMaximumIntervalOk() (*int32, bool)`

GetMaximumIntervalOk returns a tuple with the MaximumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInterval

`func (o *MotionEventInfo) SetMaximumInterval(v int32)`

SetMaximumInterval sets MaximumInterval field to given value.

### HasMaximumInterval

`func (o *MotionEventInfo) HasMaximumInterval() bool`

HasMaximumInterval returns a boolean if a field has been set.

### GetSamplingInterval

`func (o *MotionEventInfo) GetSamplingInterval() int32`

GetSamplingInterval returns the SamplingInterval field if non-nil, zero value otherwise.

### GetSamplingIntervalOk

`func (o *MotionEventInfo) GetSamplingIntervalOk() (*int32, bool)`

GetSamplingIntervalOk returns a tuple with the SamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingInterval

`func (o *MotionEventInfo) SetSamplingInterval(v int32)`

SetSamplingInterval sets SamplingInterval field to given value.

### HasSamplingInterval

`func (o *MotionEventInfo) HasSamplingInterval() bool`

HasSamplingInterval returns a boolean if a field has been set.

### GetReportingDuration

`func (o *MotionEventInfo) GetReportingDuration() int32`

GetReportingDuration returns the ReportingDuration field if non-nil, zero value otherwise.

### GetReportingDurationOk

`func (o *MotionEventInfo) GetReportingDurationOk() (*int32, bool)`

GetReportingDurationOk returns a tuple with the ReportingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingDuration

`func (o *MotionEventInfo) SetReportingDuration(v int32)`

SetReportingDuration sets ReportingDuration field to given value.

### HasReportingDuration

`func (o *MotionEventInfo) HasReportingDuration() bool`

HasReportingDuration returns a boolean if a field has been set.

### GetReportingLocationReq

`func (o *MotionEventInfo) GetReportingLocationReq() bool`

GetReportingLocationReq returns the ReportingLocationReq field if non-nil, zero value otherwise.

### GetReportingLocationReqOk

`func (o *MotionEventInfo) GetReportingLocationReqOk() (*bool, bool)`

GetReportingLocationReqOk returns a tuple with the ReportingLocationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLocationReq

`func (o *MotionEventInfo) SetReportingLocationReq(v bool)`

SetReportingLocationReq sets ReportingLocationReq field to given value.

### HasReportingLocationReq

`func (o *MotionEventInfo) HasReportingLocationReq() bool`

HasReportingLocationReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


