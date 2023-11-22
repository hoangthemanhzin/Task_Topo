# VolumeTimedReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeStamp** | **time.Time** |  | 
**EndTimeStamp** | **time.Time** |  | 
**DownlinkVolume** | **int64** |  | 
**UplinkVolume** | **int64** |  | 

## Methods

### NewVolumeTimedReport

`func NewVolumeTimedReport(startTimeStamp time.Time, endTimeStamp time.Time, downlinkVolume int64, uplinkVolume int64, ) *VolumeTimedReport`

NewVolumeTimedReport instantiates a new VolumeTimedReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTimedReportWithDefaults

`func NewVolumeTimedReportWithDefaults() *VolumeTimedReport`

NewVolumeTimedReportWithDefaults instantiates a new VolumeTimedReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeStamp

`func (o *VolumeTimedReport) GetStartTimeStamp() time.Time`

GetStartTimeStamp returns the StartTimeStamp field if non-nil, zero value otherwise.

### GetStartTimeStampOk

`func (o *VolumeTimedReport) GetStartTimeStampOk() (*time.Time, bool)`

GetStartTimeStampOk returns a tuple with the StartTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeStamp

`func (o *VolumeTimedReport) SetStartTimeStamp(v time.Time)`

SetStartTimeStamp sets StartTimeStamp field to given value.


### GetEndTimeStamp

`func (o *VolumeTimedReport) GetEndTimeStamp() time.Time`

GetEndTimeStamp returns the EndTimeStamp field if non-nil, zero value otherwise.

### GetEndTimeStampOk

`func (o *VolumeTimedReport) GetEndTimeStampOk() (*time.Time, bool)`

GetEndTimeStampOk returns a tuple with the EndTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeStamp

`func (o *VolumeTimedReport) SetEndTimeStamp(v time.Time)`

SetEndTimeStamp sets EndTimeStamp field to given value.


### GetDownlinkVolume

`func (o *VolumeTimedReport) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *VolumeTimedReport) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *VolumeTimedReport) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.


### GetUplinkVolume

`func (o *VolumeTimedReport) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *VolumeTimedReport) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *VolumeTimedReport) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


