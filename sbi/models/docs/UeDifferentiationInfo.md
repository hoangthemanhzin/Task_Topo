# UeDifferentiationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodicComInd** | Pointer to [**PeriodicCommunicationIndicator**](PeriodicCommunicationIndicator.md) |  | [optional] 
**PeriodicTime** | Pointer to **int32** |  | [optional] 
**ScheduledComTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**StationaryInd** | Pointer to [**StationaryIndication**](StationaryIndication.md) |  | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfile**](TrafficProfile.md) |  | [optional] 
**BatteryInd** | Pointer to [**BatteryIndication**](BatteryIndication.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUeDifferentiationInfo

`func NewUeDifferentiationInfo() *UeDifferentiationInfo`

NewUeDifferentiationInfo instantiates a new UeDifferentiationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeDifferentiationInfoWithDefaults

`func NewUeDifferentiationInfoWithDefaults() *UeDifferentiationInfo`

NewUeDifferentiationInfoWithDefaults instantiates a new UeDifferentiationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodicComInd

`func (o *UeDifferentiationInfo) GetPeriodicComInd() PeriodicCommunicationIndicator`

GetPeriodicComInd returns the PeriodicComInd field if non-nil, zero value otherwise.

### GetPeriodicComIndOk

`func (o *UeDifferentiationInfo) GetPeriodicComIndOk() (*PeriodicCommunicationIndicator, bool)`

GetPeriodicComIndOk returns a tuple with the PeriodicComInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicComInd

`func (o *UeDifferentiationInfo) SetPeriodicComInd(v PeriodicCommunicationIndicator)`

SetPeriodicComInd sets PeriodicComInd field to given value.

### HasPeriodicComInd

`func (o *UeDifferentiationInfo) HasPeriodicComInd() bool`

HasPeriodicComInd returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *UeDifferentiationInfo) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *UeDifferentiationInfo) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *UeDifferentiationInfo) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *UeDifferentiationInfo) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### GetScheduledComTime

`func (o *UeDifferentiationInfo) GetScheduledComTime() ScheduledCommunicationTime`

GetScheduledComTime returns the ScheduledComTime field if non-nil, zero value otherwise.

### GetScheduledComTimeOk

`func (o *UeDifferentiationInfo) GetScheduledComTimeOk() (*ScheduledCommunicationTime, bool)`

GetScheduledComTimeOk returns a tuple with the ScheduledComTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledComTime

`func (o *UeDifferentiationInfo) SetScheduledComTime(v ScheduledCommunicationTime)`

SetScheduledComTime sets ScheduledComTime field to given value.

### HasScheduledComTime

`func (o *UeDifferentiationInfo) HasScheduledComTime() bool`

HasScheduledComTime returns a boolean if a field has been set.

### GetStationaryInd

`func (o *UeDifferentiationInfo) GetStationaryInd() StationaryIndication`

GetStationaryInd returns the StationaryInd field if non-nil, zero value otherwise.

### GetStationaryIndOk

`func (o *UeDifferentiationInfo) GetStationaryIndOk() (*StationaryIndication, bool)`

GetStationaryIndOk returns a tuple with the StationaryInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryInd

`func (o *UeDifferentiationInfo) SetStationaryInd(v StationaryIndication)`

SetStationaryInd sets StationaryInd field to given value.

### HasStationaryInd

`func (o *UeDifferentiationInfo) HasStationaryInd() bool`

HasStationaryInd returns a boolean if a field has been set.

### GetTrafficProfile

`func (o *UeDifferentiationInfo) GetTrafficProfile() TrafficProfile`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *UeDifferentiationInfo) GetTrafficProfileOk() (*TrafficProfile, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *UeDifferentiationInfo) SetTrafficProfile(v TrafficProfile)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *UeDifferentiationInfo) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetBatteryInd

`func (o *UeDifferentiationInfo) GetBatteryInd() BatteryIndication`

GetBatteryInd returns the BatteryInd field if non-nil, zero value otherwise.

### GetBatteryIndOk

`func (o *UeDifferentiationInfo) GetBatteryIndOk() (*BatteryIndication, bool)`

GetBatteryIndOk returns a tuple with the BatteryInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryInd

`func (o *UeDifferentiationInfo) SetBatteryInd(v BatteryIndication)`

SetBatteryInd sets BatteryInd field to given value.

### HasBatteryInd

`func (o *UeDifferentiationInfo) HasBatteryInd() bool`

HasBatteryInd returns a boolean if a field has been set.

### GetValidityTime

`func (o *UeDifferentiationInfo) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *UeDifferentiationInfo) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *UeDifferentiationInfo) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *UeDifferentiationInfo) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


