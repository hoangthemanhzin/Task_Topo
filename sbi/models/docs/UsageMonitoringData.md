# UsageMonitoringData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UmId** | **string** | Univocally identifies the usage monitoring policy data within a PDU session. | 
**VolumeThreshold** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**VolumeThresholdUplink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**VolumeThresholdDownlink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**TimeThreshold** | Pointer to **int32** |  | [optional] 
**MonitoringTime** | Pointer to **time.Time** |  | [optional] 
**NextVolThreshold** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**NextVolThresholdUplink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**NextVolThresholdDownlink** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**NextTimeThreshold** | Pointer to **int32** |  | [optional] 
**InactivityTime** | Pointer to **int32** |  | [optional] 
**ExUsagePccRuleIds** | Pointer to **[]string** | Contains the PCC rule identifier(s) which corresponding service data flow(s) shall be excluded from PDU Session usage monitoring. It is only included in the UsageMonitoringData instance for session level usage monitoring. | [optional] 

## Methods

### NewUsageMonitoringData

`func NewUsageMonitoringData(umId string, ) *UsageMonitoringData`

NewUsageMonitoringData instantiates a new UsageMonitoringData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMonitoringDataWithDefaults

`func NewUsageMonitoringDataWithDefaults() *UsageMonitoringData`

NewUsageMonitoringDataWithDefaults instantiates a new UsageMonitoringData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUmId

`func (o *UsageMonitoringData) GetUmId() string`

GetUmId returns the UmId field if non-nil, zero value otherwise.

### GetUmIdOk

`func (o *UsageMonitoringData) GetUmIdOk() (*string, bool)`

GetUmIdOk returns a tuple with the UmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmId

`func (o *UsageMonitoringData) SetUmId(v string)`

SetUmId sets UmId field to given value.


### GetVolumeThreshold

`func (o *UsageMonitoringData) GetVolumeThreshold() int64`

GetVolumeThreshold returns the VolumeThreshold field if non-nil, zero value otherwise.

### GetVolumeThresholdOk

`func (o *UsageMonitoringData) GetVolumeThresholdOk() (*int64, bool)`

GetVolumeThresholdOk returns a tuple with the VolumeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeThreshold

`func (o *UsageMonitoringData) SetVolumeThreshold(v int64)`

SetVolumeThreshold sets VolumeThreshold field to given value.

### HasVolumeThreshold

`func (o *UsageMonitoringData) HasVolumeThreshold() bool`

HasVolumeThreshold returns a boolean if a field has been set.

### SetVolumeThresholdNil

`func (o *UsageMonitoringData) SetVolumeThresholdNil(b bool)`

 SetVolumeThresholdNil sets the value for VolumeThreshold to be an explicit nil

### UnsetVolumeThreshold
`func (o *UsageMonitoringData) UnsetVolumeThreshold()`

UnsetVolumeThreshold ensures that no value is present for VolumeThreshold, not even an explicit nil
### GetVolumeThresholdUplink

`func (o *UsageMonitoringData) GetVolumeThresholdUplink() int64`

GetVolumeThresholdUplink returns the VolumeThresholdUplink field if non-nil, zero value otherwise.

### GetVolumeThresholdUplinkOk

`func (o *UsageMonitoringData) GetVolumeThresholdUplinkOk() (*int64, bool)`

GetVolumeThresholdUplinkOk returns a tuple with the VolumeThresholdUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeThresholdUplink

`func (o *UsageMonitoringData) SetVolumeThresholdUplink(v int64)`

SetVolumeThresholdUplink sets VolumeThresholdUplink field to given value.

### HasVolumeThresholdUplink

`func (o *UsageMonitoringData) HasVolumeThresholdUplink() bool`

HasVolumeThresholdUplink returns a boolean if a field has been set.

### SetVolumeThresholdUplinkNil

`func (o *UsageMonitoringData) SetVolumeThresholdUplinkNil(b bool)`

 SetVolumeThresholdUplinkNil sets the value for VolumeThresholdUplink to be an explicit nil

### UnsetVolumeThresholdUplink
`func (o *UsageMonitoringData) UnsetVolumeThresholdUplink()`

UnsetVolumeThresholdUplink ensures that no value is present for VolumeThresholdUplink, not even an explicit nil
### GetVolumeThresholdDownlink

`func (o *UsageMonitoringData) GetVolumeThresholdDownlink() int64`

GetVolumeThresholdDownlink returns the VolumeThresholdDownlink field if non-nil, zero value otherwise.

### GetVolumeThresholdDownlinkOk

`func (o *UsageMonitoringData) GetVolumeThresholdDownlinkOk() (*int64, bool)`

GetVolumeThresholdDownlinkOk returns a tuple with the VolumeThresholdDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeThresholdDownlink

`func (o *UsageMonitoringData) SetVolumeThresholdDownlink(v int64)`

SetVolumeThresholdDownlink sets VolumeThresholdDownlink field to given value.

### HasVolumeThresholdDownlink

`func (o *UsageMonitoringData) HasVolumeThresholdDownlink() bool`

HasVolumeThresholdDownlink returns a boolean if a field has been set.

### SetVolumeThresholdDownlinkNil

`func (o *UsageMonitoringData) SetVolumeThresholdDownlinkNil(b bool)`

 SetVolumeThresholdDownlinkNil sets the value for VolumeThresholdDownlink to be an explicit nil

### UnsetVolumeThresholdDownlink
`func (o *UsageMonitoringData) UnsetVolumeThresholdDownlink()`

UnsetVolumeThresholdDownlink ensures that no value is present for VolumeThresholdDownlink, not even an explicit nil
### GetTimeThreshold

`func (o *UsageMonitoringData) GetTimeThreshold() int32`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *UsageMonitoringData) GetTimeThresholdOk() (*int32, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *UsageMonitoringData) SetTimeThreshold(v int32)`

SetTimeThreshold sets TimeThreshold field to given value.

### HasTimeThreshold

`func (o *UsageMonitoringData) HasTimeThreshold() bool`

HasTimeThreshold returns a boolean if a field has been set.

### SetTimeThresholdNil

`func (o *UsageMonitoringData) SetTimeThresholdNil(b bool)`

 SetTimeThresholdNil sets the value for TimeThreshold to be an explicit nil

### UnsetTimeThreshold
`func (o *UsageMonitoringData) UnsetTimeThreshold()`

UnsetTimeThreshold ensures that no value is present for TimeThreshold, not even an explicit nil
### GetMonitoringTime

`func (o *UsageMonitoringData) GetMonitoringTime() time.Time`

GetMonitoringTime returns the MonitoringTime field if non-nil, zero value otherwise.

### GetMonitoringTimeOk

`func (o *UsageMonitoringData) GetMonitoringTimeOk() (*time.Time, bool)`

GetMonitoringTimeOk returns a tuple with the MonitoringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringTime

`func (o *UsageMonitoringData) SetMonitoringTime(v time.Time)`

SetMonitoringTime sets MonitoringTime field to given value.

### HasMonitoringTime

`func (o *UsageMonitoringData) HasMonitoringTime() bool`

HasMonitoringTime returns a boolean if a field has been set.

### SetMonitoringTimeNil

`func (o *UsageMonitoringData) SetMonitoringTimeNil(b bool)`

 SetMonitoringTimeNil sets the value for MonitoringTime to be an explicit nil

### UnsetMonitoringTime
`func (o *UsageMonitoringData) UnsetMonitoringTime()`

UnsetMonitoringTime ensures that no value is present for MonitoringTime, not even an explicit nil
### GetNextVolThreshold

`func (o *UsageMonitoringData) GetNextVolThreshold() int64`

GetNextVolThreshold returns the NextVolThreshold field if non-nil, zero value otherwise.

### GetNextVolThresholdOk

`func (o *UsageMonitoringData) GetNextVolThresholdOk() (*int64, bool)`

GetNextVolThresholdOk returns a tuple with the NextVolThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolThreshold

`func (o *UsageMonitoringData) SetNextVolThreshold(v int64)`

SetNextVolThreshold sets NextVolThreshold field to given value.

### HasNextVolThreshold

`func (o *UsageMonitoringData) HasNextVolThreshold() bool`

HasNextVolThreshold returns a boolean if a field has been set.

### SetNextVolThresholdNil

`func (o *UsageMonitoringData) SetNextVolThresholdNil(b bool)`

 SetNextVolThresholdNil sets the value for NextVolThreshold to be an explicit nil

### UnsetNextVolThreshold
`func (o *UsageMonitoringData) UnsetNextVolThreshold()`

UnsetNextVolThreshold ensures that no value is present for NextVolThreshold, not even an explicit nil
### GetNextVolThresholdUplink

`func (o *UsageMonitoringData) GetNextVolThresholdUplink() int64`

GetNextVolThresholdUplink returns the NextVolThresholdUplink field if non-nil, zero value otherwise.

### GetNextVolThresholdUplinkOk

`func (o *UsageMonitoringData) GetNextVolThresholdUplinkOk() (*int64, bool)`

GetNextVolThresholdUplinkOk returns a tuple with the NextVolThresholdUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolThresholdUplink

`func (o *UsageMonitoringData) SetNextVolThresholdUplink(v int64)`

SetNextVolThresholdUplink sets NextVolThresholdUplink field to given value.

### HasNextVolThresholdUplink

`func (o *UsageMonitoringData) HasNextVolThresholdUplink() bool`

HasNextVolThresholdUplink returns a boolean if a field has been set.

### SetNextVolThresholdUplinkNil

`func (o *UsageMonitoringData) SetNextVolThresholdUplinkNil(b bool)`

 SetNextVolThresholdUplinkNil sets the value for NextVolThresholdUplink to be an explicit nil

### UnsetNextVolThresholdUplink
`func (o *UsageMonitoringData) UnsetNextVolThresholdUplink()`

UnsetNextVolThresholdUplink ensures that no value is present for NextVolThresholdUplink, not even an explicit nil
### GetNextVolThresholdDownlink

`func (o *UsageMonitoringData) GetNextVolThresholdDownlink() int64`

GetNextVolThresholdDownlink returns the NextVolThresholdDownlink field if non-nil, zero value otherwise.

### GetNextVolThresholdDownlinkOk

`func (o *UsageMonitoringData) GetNextVolThresholdDownlinkOk() (*int64, bool)`

GetNextVolThresholdDownlinkOk returns a tuple with the NextVolThresholdDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVolThresholdDownlink

`func (o *UsageMonitoringData) SetNextVolThresholdDownlink(v int64)`

SetNextVolThresholdDownlink sets NextVolThresholdDownlink field to given value.

### HasNextVolThresholdDownlink

`func (o *UsageMonitoringData) HasNextVolThresholdDownlink() bool`

HasNextVolThresholdDownlink returns a boolean if a field has been set.

### SetNextVolThresholdDownlinkNil

`func (o *UsageMonitoringData) SetNextVolThresholdDownlinkNil(b bool)`

 SetNextVolThresholdDownlinkNil sets the value for NextVolThresholdDownlink to be an explicit nil

### UnsetNextVolThresholdDownlink
`func (o *UsageMonitoringData) UnsetNextVolThresholdDownlink()`

UnsetNextVolThresholdDownlink ensures that no value is present for NextVolThresholdDownlink, not even an explicit nil
### GetNextTimeThreshold

`func (o *UsageMonitoringData) GetNextTimeThreshold() int32`

GetNextTimeThreshold returns the NextTimeThreshold field if non-nil, zero value otherwise.

### GetNextTimeThresholdOk

`func (o *UsageMonitoringData) GetNextTimeThresholdOk() (*int32, bool)`

GetNextTimeThresholdOk returns a tuple with the NextTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTimeThreshold

`func (o *UsageMonitoringData) SetNextTimeThreshold(v int32)`

SetNextTimeThreshold sets NextTimeThreshold field to given value.

### HasNextTimeThreshold

`func (o *UsageMonitoringData) HasNextTimeThreshold() bool`

HasNextTimeThreshold returns a boolean if a field has been set.

### SetNextTimeThresholdNil

`func (o *UsageMonitoringData) SetNextTimeThresholdNil(b bool)`

 SetNextTimeThresholdNil sets the value for NextTimeThreshold to be an explicit nil

### UnsetNextTimeThreshold
`func (o *UsageMonitoringData) UnsetNextTimeThreshold()`

UnsetNextTimeThreshold ensures that no value is present for NextTimeThreshold, not even an explicit nil
### GetInactivityTime

`func (o *UsageMonitoringData) GetInactivityTime() int32`

GetInactivityTime returns the InactivityTime field if non-nil, zero value otherwise.

### GetInactivityTimeOk

`func (o *UsageMonitoringData) GetInactivityTimeOk() (*int32, bool)`

GetInactivityTimeOk returns a tuple with the InactivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityTime

`func (o *UsageMonitoringData) SetInactivityTime(v int32)`

SetInactivityTime sets InactivityTime field to given value.

### HasInactivityTime

`func (o *UsageMonitoringData) HasInactivityTime() bool`

HasInactivityTime returns a boolean if a field has been set.

### SetInactivityTimeNil

`func (o *UsageMonitoringData) SetInactivityTimeNil(b bool)`

 SetInactivityTimeNil sets the value for InactivityTime to be an explicit nil

### UnsetInactivityTime
`func (o *UsageMonitoringData) UnsetInactivityTime()`

UnsetInactivityTime ensures that no value is present for InactivityTime, not even an explicit nil
### GetExUsagePccRuleIds

`func (o *UsageMonitoringData) GetExUsagePccRuleIds() []string`

GetExUsagePccRuleIds returns the ExUsagePccRuleIds field if non-nil, zero value otherwise.

### GetExUsagePccRuleIdsOk

`func (o *UsageMonitoringData) GetExUsagePccRuleIdsOk() (*[]string, bool)`

GetExUsagePccRuleIdsOk returns a tuple with the ExUsagePccRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExUsagePccRuleIds

`func (o *UsageMonitoringData) SetExUsagePccRuleIds(v []string)`

SetExUsagePccRuleIds sets ExUsagePccRuleIds field to given value.

### HasExUsagePccRuleIds

`func (o *UsageMonitoringData) HasExUsagePccRuleIds() bool`

HasExUsagePccRuleIds returns a boolean if a field has been set.

### SetExUsagePccRuleIdsNil

`func (o *UsageMonitoringData) SetExUsagePccRuleIdsNil(b bool)`

 SetExUsagePccRuleIdsNil sets the value for ExUsagePccRuleIds to be an explicit nil

### UnsetExUsagePccRuleIds
`func (o *UsageMonitoringData) UnsetExUsagePccRuleIds()`

UnsetExUsagePccRuleIds ensures that no value is present for ExUsagePccRuleIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


