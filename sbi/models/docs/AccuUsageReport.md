# AccuUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefUmIds** | **string** | An id referencing UsageMonitoringData objects associated with this usage report. | 
**VolUsage** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**VolUsageUplink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**VolUsageDownlink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**TimeUsage** | Pointer to **int32** |  | [optional] 
**NextVolUsage** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**NextVolUsageUplink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**NextVolUsageDownlink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**NextTimeUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccuUsageReport

`func NewAccuUsageReport(refUmIds string, ) *AccuUsageReport`

NewAccuUsageReport instantiates a new AccuUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccuUsageReportWithDefaults

`func NewAccuUsageReportWithDefaults() *AccuUsageReport`

NewAccuUsageReportWithDefaults instantiates a new AccuUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefUmIds

`func (o *AccuUsageReport) GetRefUmIds() string`

GetRefUmIds returns the RefUmIds field if non-nil, zero value otherwise.

### GetRefUmIdsOk

`func (o *AccuUsageReport) GetRefUmIdsOk() (*string, bool)`

GetRefUmIdsOk returns a tuple with the RefUmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmIds

`func (o *AccuUsageReport) SetRefUmIds(v string)`

SetRefUmIds sets RefUmIds field to given value.


### GetVolUsage

`func (o *AccuUsageReport) GetVolUsage() int64`

GetVolUsage returns the VolUsage field if non-nil, zero value otherwise.

### GetVolUsageOk

`func (o *AccuUsageReport) GetVolUsageOk() (*int64, bool)`

GetVolUsageOk returns a tuple with the VolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolUsage

`func (o *AccuUsageReport) SetVolUsage(v int64)`

SetVolUsage sets VolUsage field to given value.

### HasVolUsage

`func (o *AccuUsageReport) HasVolUsage() bool`

HasVolUsage returns a boolean if a field has been set.

### GetVolUsageUplink

`func (o *AccuUsageReport) GetVolUsageUplink() int64`

GetVolUsageUplink returns the VolUsageUplink field if non-nil, zero value otherwise.

### GetVolUsageUplinkOk

`func (o *AccuUsageReport) GetVolUsageUplinkOk() (*int64, bool)`

GetVolUsageUplinkOk returns a tuple with the VolUsageUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolUsageUplink

`func (o *AccuUsageReport) SetVolUsageUplink(v int64)`

SetVolUsageUplink sets VolUsageUplink field to given value.

### HasVolUsageUplink

`func (o *AccuUsageReport) HasVolUsageUplink() bool`

HasVolUsageUplink returns a boolean if a field has been set.

### GetVolUsageDownlink

`func (o *AccuUsageReport) GetVolUsageDownlink() int64`

GetVolUsageDownlink returns the VolUsageDownlink field if non-nil, zero value otherwise.

### GetVolUsageDownlinkOk

`func (o *AccuUsageReport) GetVolUsageDownlinkOk() (*int64, bool)`

GetVolUsageDownlinkOk returns a tuple with the VolUsageDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolUsageDownlink

`func (o *AccuUsageReport) SetVolUsageDownlink(v int64)`

SetVolUsageDownlink sets VolUsageDownlink field to given value.

### HasVolUsageDownlink

`func (o *AccuUsageReport) HasVolUsageDownlink() bool`

HasVolUsageDownlink returns a boolean if a field has been set.

### GetTimeUsage

`func (o *AccuUsageReport) GetTimeUsage() int32`

GetTimeUsage returns the TimeUsage field if non-nil, zero value otherwise.

### GetTimeUsageOk

`func (o *AccuUsageReport) GetTimeUsageOk() (*int32, bool)`

GetTimeUsageOk returns a tuple with the TimeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUsage

`func (o *AccuUsageReport) SetTimeUsage(v int32)`

SetTimeUsage sets TimeUsage field to given value.

### HasTimeUsage

`func (o *AccuUsageReport) HasTimeUsage() bool`

HasTimeUsage returns a boolean if a field has been set.

### GetNextVolUsage

`func (o *AccuUsageReport) GetNextVolUsage() int64`

GetNextVolUsage returns the NextVolUsage field if non-nil, zero value otherwise.

### GetNextVolUsageOk

`func (o *AccuUsageReport) GetNextVolUsageOk() (*int64, bool)`

GetNextVolUsageOk returns a tuple with the NextVolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolUsage

`func (o *AccuUsageReport) SetNextVolUsage(v int64)`

SetNextVolUsage sets NextVolUsage field to given value.

### HasNextVolUsage

`func (o *AccuUsageReport) HasNextVolUsage() bool`

HasNextVolUsage returns a boolean if a field has been set.

### GetNextVolUsageUplink

`func (o *AccuUsageReport) GetNextVolUsageUplink() int64`

GetNextVolUsageUplink returns the NextVolUsageUplink field if non-nil, zero value otherwise.

### GetNextVolUsageUplinkOk

`func (o *AccuUsageReport) GetNextVolUsageUplinkOk() (*int64, bool)`

GetNextVolUsageUplinkOk returns a tuple with the NextVolUsageUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolUsageUplink

`func (o *AccuUsageReport) SetNextVolUsageUplink(v int64)`

SetNextVolUsageUplink sets NextVolUsageUplink field to given value.

### HasNextVolUsageUplink

`func (o *AccuUsageReport) HasNextVolUsageUplink() bool`

HasNextVolUsageUplink returns a boolean if a field has been set.

### GetNextVolUsageDownlink

`func (o *AccuUsageReport) GetNextVolUsageDownlink() int64`

GetNextVolUsageDownlink returns the NextVolUsageDownlink field if non-nil, zero value otherwise.

### GetNextVolUsageDownlinkOk

`func (o *AccuUsageReport) GetNextVolUsageDownlinkOk() (*int64, bool)`

GetNextVolUsageDownlinkOk returns a tuple with the NextVolUsageDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolUsageDownlink

`func (o *AccuUsageReport) SetNextVolUsageDownlink(v int64)`

SetNextVolUsageDownlink sets NextVolUsageDownlink field to given value.

### HasNextVolUsageDownlink

`func (o *AccuUsageReport) HasNextVolUsageDownlink() bool`

HasNextVolUsageDownlink returns a boolean if a field has been set.

### GetNextTimeUsage

`func (o *AccuUsageReport) GetNextTimeUsage() int32`

GetNextTimeUsage returns the NextTimeUsage field if non-nil, zero value otherwise.

### GetNextTimeUsageOk

`func (o *AccuUsageReport) GetNextTimeUsageOk() (*int32, bool)`

GetNextTimeUsageOk returns a tuple with the NextTimeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTimeUsage

`func (o *AccuUsageReport) SetNextTimeUsage(v int32)`

SetNextTimeUsage sets NextTimeUsage field to given value.

### HasNextTimeUsage

`func (o *AccuUsageReport) HasNextTimeUsage() bool`

HasNextTimeUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


