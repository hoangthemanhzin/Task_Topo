# AccumulatedUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**TotalVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DownlinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**UplinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewAccumulatedUsage

`func NewAccumulatedUsage() *AccumulatedUsage`

NewAccumulatedUsage instantiates a new AccumulatedUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccumulatedUsageWithDefaults

`func NewAccumulatedUsageWithDefaults() *AccumulatedUsage`

NewAccumulatedUsageWithDefaults instantiates a new AccumulatedUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *AccumulatedUsage) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AccumulatedUsage) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AccumulatedUsage) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AccumulatedUsage) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTotalVolume

`func (o *AccumulatedUsage) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *AccumulatedUsage) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *AccumulatedUsage) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *AccumulatedUsage) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *AccumulatedUsage) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *AccumulatedUsage) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *AccumulatedUsage) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *AccumulatedUsage) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.

### GetUplinkVolume

`func (o *AccumulatedUsage) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *AccumulatedUsage) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *AccumulatedUsage) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *AccumulatedUsage) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


