# QosFlowUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qfi** | **int32** |  | 
**StartTimeStamp** | **time.Time** |  | 
**EndTimeStamp** | **time.Time** |  | 
**DownlinkVolume** | **int64** |  | 
**UplinkVolume** | **int64** |  | 

## Methods

### NewQosFlowUsageReport

`func NewQosFlowUsageReport(qfi int32, startTimeStamp time.Time, endTimeStamp time.Time, downlinkVolume int64, uplinkVolume int64, ) *QosFlowUsageReport`

NewQosFlowUsageReport instantiates a new QosFlowUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowUsageReportWithDefaults

`func NewQosFlowUsageReportWithDefaults() *QosFlowUsageReport`

NewQosFlowUsageReportWithDefaults instantiates a new QosFlowUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfi

`func (o *QosFlowUsageReport) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *QosFlowUsageReport) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *QosFlowUsageReport) SetQfi(v int32)`

SetQfi sets Qfi field to given value.


### GetStartTimeStamp

`func (o *QosFlowUsageReport) GetStartTimeStamp() time.Time`

GetStartTimeStamp returns the StartTimeStamp field if non-nil, zero value otherwise.

### GetStartTimeStampOk

`func (o *QosFlowUsageReport) GetStartTimeStampOk() (*time.Time, bool)`

GetStartTimeStampOk returns a tuple with the StartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeStamp

`func (o *QosFlowUsageReport) SetStartTimeStamp(v time.Time)`

SetStartTimeStamp sets StartTimeStamp field to given value.


### GetEndTimeStamp

`func (o *QosFlowUsageReport) GetEndTimeStamp() time.Time`

GetEndTimeStamp returns the EndTimeStamp field if non-nil, zero value otherwise.

### GetEndTimeStampOk

`func (o *QosFlowUsageReport) GetEndTimeStampOk() (*time.Time, bool)`

GetEndTimeStampOk returns a tuple with the EndTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeStamp

`func (o *QosFlowUsageReport) SetEndTimeStamp(v time.Time)`

SetEndTimeStamp sets EndTimeStamp field to given value.


### GetDownlinkVolume

`func (o *QosFlowUsageReport) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *QosFlowUsageReport) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *QosFlowUsageReport) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.


### GetUplinkVolume

`func (o *QosFlowUsageReport) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *QosFlowUsageReport) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *QosFlowUsageReport) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


