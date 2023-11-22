# UsageThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**TotalVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DownlinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**UplinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewUsageThreshold

`func NewUsageThreshold() *UsageThreshold`

NewUsageThreshold instantiates a new UsageThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageThresholdWithDefaults

`func NewUsageThresholdWithDefaults() *UsageThreshold`

NewUsageThresholdWithDefaults instantiates a new UsageThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *UsageThreshold) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UsageThreshold) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UsageThreshold) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UsageThreshold) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTotalVolume

`func (o *UsageThreshold) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *UsageThreshold) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *UsageThreshold) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *UsageThreshold) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *UsageThreshold) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *UsageThreshold) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *UsageThreshold) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *UsageThreshold) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.

### GetUplinkVolume

`func (o *UsageThreshold) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *UsageThreshold) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *UsageThreshold) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *UsageThreshold) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


