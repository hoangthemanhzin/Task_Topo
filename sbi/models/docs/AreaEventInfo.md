# AreaEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaDefinition** | [**[]ReportingArea**](ReportingArea.md) |  | 
**OccurrenceInfo** | Pointer to [**OccurrenceInfo**](OccurrenceInfo.md) |  | [optional] 
**MinimumInterval** | Pointer to **int32** |  | [optional] 
**MaximumInterval** | Pointer to **int32** |  | [optional] 
**SamplingInterval** | Pointer to **int32** |  | [optional] 
**ReportingDuration** | Pointer to **int32** |  | [optional] 
**ReportingLocationReq** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewAreaEventInfo

`func NewAreaEventInfo(areaDefinition []ReportingArea, ) *AreaEventInfo`

NewAreaEventInfo instantiates a new AreaEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaEventInfoWithDefaults

`func NewAreaEventInfoWithDefaults() *AreaEventInfo`

NewAreaEventInfoWithDefaults instantiates a new AreaEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaDefinition

`func (o *AreaEventInfo) GetAreaDefinition() []ReportingArea`

GetAreaDefinition returns the AreaDefinition field if non-nil, zero value otherwise.

### GetAreaDefinitionOk

`func (o *AreaEventInfo) GetAreaDefinitionOk() (*[]ReportingArea, bool)`

GetAreaDefinitionOk returns a tuple with the AreaDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaDefinition

`func (o *AreaEventInfo) SetAreaDefinition(v []ReportingArea)`

SetAreaDefinition sets AreaDefinition field to given value.


### GetOccurrenceInfo

`func (o *AreaEventInfo) GetOccurrenceInfo() OccurrenceInfo`

GetOccurrenceInfo returns the OccurrenceInfo field if non-nil, zero value otherwise.

### GetOccurrenceInfoOk

`func (o *AreaEventInfo) GetOccurrenceInfoOk() (*OccurrenceInfo, bool)`

GetOccurrenceInfoOk returns a tuple with the OccurrenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceInfo

`func (o *AreaEventInfo) SetOccurrenceInfo(v OccurrenceInfo)`

SetOccurrenceInfo sets OccurrenceInfo field to given value.

### HasOccurrenceInfo

`func (o *AreaEventInfo) HasOccurrenceInfo() bool`

HasOccurrenceInfo returns a boolean if a field has been set.

### GetMinimumInterval

`func (o *AreaEventInfo) GetMinimumInterval() int32`

GetMinimumInterval returns the MinimumInterval field if non-nil, zero value otherwise.

### GetMinimumIntervalOk

`func (o *AreaEventInfo) GetMinimumIntervalOk() (*int32, bool)`

GetMinimumIntervalOk returns a tuple with the MinimumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumInterval

`func (o *AreaEventInfo) SetMinimumInterval(v int32)`

SetMinimumInterval sets MinimumInterval field to given value.

### HasMinimumInterval

`func (o *AreaEventInfo) HasMinimumInterval() bool`

HasMinimumInterval returns a boolean if a field has been set.

### GetMaximumInterval

`func (o *AreaEventInfo) GetMaximumInterval() int32`

GetMaximumInterval returns the MaximumInterval field if non-nil, zero value otherwise.

### GetMaximumIntervalOk

`func (o *AreaEventInfo) GetMaximumIntervalOk() (*int32, bool)`

GetMaximumIntervalOk returns a tuple with the MaximumInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumInterval

`func (o *AreaEventInfo) SetMaximumInterval(v int32)`

SetMaximumInterval sets MaximumInterval field to given value.

### HasMaximumInterval

`func (o *AreaEventInfo) HasMaximumInterval() bool`

HasMaximumInterval returns a boolean if a field has been set.

### GetSamplingInterval

`func (o *AreaEventInfo) GetSamplingInterval() int32`

GetSamplingInterval returns the SamplingInterval field if non-nil, zero value otherwise.

### GetSamplingIntervalOk

`func (o *AreaEventInfo) GetSamplingIntervalOk() (*int32, bool)`

GetSamplingIntervalOk returns a tuple with the SamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingInterval

`func (o *AreaEventInfo) SetSamplingInterval(v int32)`

SetSamplingInterval sets SamplingInterval field to given value.

### HasSamplingInterval

`func (o *AreaEventInfo) HasSamplingInterval() bool`

HasSamplingInterval returns a boolean if a field has been set.

### GetReportingDuration

`func (o *AreaEventInfo) GetReportingDuration() int32`

GetReportingDuration returns the ReportingDuration field if non-nil, zero value otherwise.

### GetReportingDurationOk

`func (o *AreaEventInfo) GetReportingDurationOk() (*int32, bool)`

GetReportingDurationOk returns a tuple with the ReportingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingDuration

`func (o *AreaEventInfo) SetReportingDuration(v int32)`

SetReportingDuration sets ReportingDuration field to given value.

### HasReportingDuration

`func (o *AreaEventInfo) HasReportingDuration() bool`

HasReportingDuration returns a boolean if a field has been set.

### GetReportingLocationReq

`func (o *AreaEventInfo) GetReportingLocationReq() bool`

GetReportingLocationReq returns the ReportingLocationReq field if non-nil, zero value otherwise.

### GetReportingLocationReqOk

`func (o *AreaEventInfo) GetReportingLocationReqOk() (*bool, bool)`

GetReportingLocationReqOk returns a tuple with the ReportingLocationReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLocationReq

`func (o *AreaEventInfo) SetReportingLocationReq(v bool)`

SetReportingLocationReq sets ReportingLocationReq field to given value.

### HasReportingLocationReq

`func (o *AreaEventInfo) HasReportingLocationReq() bool`

HasReportingLocationReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


