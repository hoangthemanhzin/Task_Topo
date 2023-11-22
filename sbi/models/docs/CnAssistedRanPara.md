# CnAssistedRanPara

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StationaryIndication** | Pointer to [**StationaryIndication**](StationaryIndication.md) |  | [optional] 
**CommunicationDurationTime** | Pointer to **int32** |  | [optional] 
**PeriodicTime** | Pointer to **int32** |  | [optional] 
**ScheduledCommunicationTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**ScheduledCommunicationType** | Pointer to [**ScheduledCommunicationType**](ScheduledCommunicationType.md) |  | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfile**](TrafficProfile.md) |  | [optional] 
**BatteryIndication** | Pointer to [**BatteryIndication**](BatteryIndication.md) |  | [optional] 

## Methods

### NewCnAssistedRanPara

`func NewCnAssistedRanPara() *CnAssistedRanPara`

NewCnAssistedRanPara instantiates a new CnAssistedRanPara object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCnAssistedRanParaWithDefaults

`func NewCnAssistedRanParaWithDefaults() *CnAssistedRanPara`

NewCnAssistedRanParaWithDefaults instantiates a new CnAssistedRanPara object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStationaryIndication

`func (o *CnAssistedRanPara) GetStationaryIndication() StationaryIndication`

GetStationaryIndication returns the StationaryIndication field if non-nil, zero value otherwise.

### GetStationaryIndicationOk

`func (o *CnAssistedRanPara) GetStationaryIndicationOk() (*StationaryIndication, bool)`

GetStationaryIndicationOk returns a tuple with the StationaryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryIndication

`func (o *CnAssistedRanPara) SetStationaryIndication(v StationaryIndication)`

SetStationaryIndication sets StationaryIndication field to given value.

### HasStationaryIndication

`func (o *CnAssistedRanPara) HasStationaryIndication() bool`

HasStationaryIndication returns a boolean if a field has been set.

### GetCommunicationDurationTime

`func (o *CnAssistedRanPara) GetCommunicationDurationTime() int32`

GetCommunicationDurationTime returns the CommunicationDurationTime field if non-nil, zero value otherwise.

### GetCommunicationDurationTimeOk

`func (o *CnAssistedRanPara) GetCommunicationDurationTimeOk() (*int32, bool)`

GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDurationTime

`func (o *CnAssistedRanPara) SetCommunicationDurationTime(v int32)`

SetCommunicationDurationTime sets CommunicationDurationTime field to given value.

### HasCommunicationDurationTime

`func (o *CnAssistedRanPara) HasCommunicationDurationTime() bool`

HasCommunicationDurationTime returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *CnAssistedRanPara) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *CnAssistedRanPara) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *CnAssistedRanPara) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *CnAssistedRanPara) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### GetScheduledCommunicationTime

`func (o *CnAssistedRanPara) GetScheduledCommunicationTime() ScheduledCommunicationTime`

GetScheduledCommunicationTime returns the ScheduledCommunicationTime field if non-nil, zero value otherwise.

### GetScheduledCommunicationTimeOk

`func (o *CnAssistedRanPara) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime, bool)`

GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationTime

`func (o *CnAssistedRanPara) SetScheduledCommunicationTime(v ScheduledCommunicationTime)`

SetScheduledCommunicationTime sets ScheduledCommunicationTime field to given value.

### HasScheduledCommunicationTime

`func (o *CnAssistedRanPara) HasScheduledCommunicationTime() bool`

HasScheduledCommunicationTime returns a boolean if a field has been set.

### GetScheduledCommunicationType

`func (o *CnAssistedRanPara) GetScheduledCommunicationType() ScheduledCommunicationType`

GetScheduledCommunicationType returns the ScheduledCommunicationType field if non-nil, zero value otherwise.

### GetScheduledCommunicationTypeOk

`func (o *CnAssistedRanPara) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool)`

GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationType

`func (o *CnAssistedRanPara) SetScheduledCommunicationType(v ScheduledCommunicationType)`

SetScheduledCommunicationType sets ScheduledCommunicationType field to given value.

### HasScheduledCommunicationType

`func (o *CnAssistedRanPara) HasScheduledCommunicationType() bool`

HasScheduledCommunicationType returns a boolean if a field has been set.

### GetTrafficProfile

`func (o *CnAssistedRanPara) GetTrafficProfile() TrafficProfile`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *CnAssistedRanPara) GetTrafficProfileOk() (*TrafficProfile, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *CnAssistedRanPara) SetTrafficProfile(v TrafficProfile)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *CnAssistedRanPara) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetBatteryIndication

`func (o *CnAssistedRanPara) GetBatteryIndication() BatteryIndication`

GetBatteryIndication returns the BatteryIndication field if non-nil, zero value otherwise.

### GetBatteryIndicationOk

`func (o *CnAssistedRanPara) GetBatteryIndicationOk() (*BatteryIndication, bool)`

GetBatteryIndicationOk returns a tuple with the BatteryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryIndication

`func (o *CnAssistedRanPara) SetBatteryIndication(v BatteryIndication)`

SetBatteryIndication sets BatteryIndication field to given value.

### HasBatteryIndication

`func (o *CnAssistedRanPara) HasBatteryIndication() bool`

HasBatteryIndication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


