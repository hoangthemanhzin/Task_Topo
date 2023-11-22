# ExpectedUeBehaviourData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StationaryIndication** | Pointer to [**StationaryIndication**](StationaryIndication.md) |  | [optional] 
**CommunicationDurationTime** | Pointer to **int32** |  | [optional] 
**PeriodicTime** | Pointer to **int32** |  | [optional] 
**ScheduledCommunicationTime** | Pointer to [**ScheduledCommunicationTime1**](ScheduledCommunicationTime1.md) |  | [optional] 
**ScheduledCommunicationType** | Pointer to [**ScheduledCommunicationType**](ScheduledCommunicationType.md) |  | [optional] 
**ExpectedUmts** | Pointer to [**[]LocationArea1**](LocationArea1.md) | Identifies the UE&#39;s expected geographical movement. The attribute is only applicable in 5G. | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfile**](TrafficProfile.md) |  | [optional] 
**BatteryIndication** | Pointer to [**BatteryIndication**](BatteryIndication.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewExpectedUeBehaviourData1

`func NewExpectedUeBehaviourData1() *ExpectedUeBehaviourData1`

NewExpectedUeBehaviourData1 instantiates a new ExpectedUeBehaviourData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedUeBehaviourData1WithDefaults

`func NewExpectedUeBehaviourData1WithDefaults() *ExpectedUeBehaviourData1`

NewExpectedUeBehaviourData1WithDefaults instantiates a new ExpectedUeBehaviourData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStationaryIndication

`func (o *ExpectedUeBehaviourData1) GetStationaryIndication() StationaryIndication`

GetStationaryIndication returns the StationaryIndication field if non-nil, zero value otherwise.

### GetStationaryIndicationOk

`func (o *ExpectedUeBehaviourData1) GetStationaryIndicationOk() (*StationaryIndication, bool)`

GetStationaryIndicationOk returns a tuple with the StationaryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryIndication

`func (o *ExpectedUeBehaviourData1) SetStationaryIndication(v StationaryIndication)`

SetStationaryIndication sets StationaryIndication field to given value.

### HasStationaryIndication

`func (o *ExpectedUeBehaviourData1) HasStationaryIndication() bool`

HasStationaryIndication returns a boolean if a field has been set.

### GetCommunicationDurationTime

`func (o *ExpectedUeBehaviourData1) GetCommunicationDurationTime() int32`

GetCommunicationDurationTime returns the CommunicationDurationTime field if non-nil, zero value otherwise.

### GetCommunicationDurationTimeOk

`func (o *ExpectedUeBehaviourData1) GetCommunicationDurationTimeOk() (*int32, bool)`

GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDurationTime

`func (o *ExpectedUeBehaviourData1) SetCommunicationDurationTime(v int32)`

SetCommunicationDurationTime sets CommunicationDurationTime field to given value.

### HasCommunicationDurationTime

`func (o *ExpectedUeBehaviourData1) HasCommunicationDurationTime() bool`

HasCommunicationDurationTime returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *ExpectedUeBehaviourData1) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *ExpectedUeBehaviourData1) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *ExpectedUeBehaviourData1) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *ExpectedUeBehaviourData1) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### GetScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData1) GetScheduledCommunicationTime() ScheduledCommunicationTime1`

GetScheduledCommunicationTime returns the ScheduledCommunicationTime field if non-nil, zero value otherwise.

### GetScheduledCommunicationTimeOk

`func (o *ExpectedUeBehaviourData1) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime1, bool)`

GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData1) SetScheduledCommunicationTime(v ScheduledCommunicationTime1)`

SetScheduledCommunicationTime sets ScheduledCommunicationTime field to given value.

### HasScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData1) HasScheduledCommunicationTime() bool`

HasScheduledCommunicationTime returns a boolean if a field has been set.

### GetScheduledCommunicationType

`func (o *ExpectedUeBehaviourData1) GetScheduledCommunicationType() ScheduledCommunicationType`

GetScheduledCommunicationType returns the ScheduledCommunicationType field if non-nil, zero value otherwise.

### GetScheduledCommunicationTypeOk

`func (o *ExpectedUeBehaviourData1) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool)`

GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationType

`func (o *ExpectedUeBehaviourData1) SetScheduledCommunicationType(v ScheduledCommunicationType)`

SetScheduledCommunicationType sets ScheduledCommunicationType field to given value.

### HasScheduledCommunicationType

`func (o *ExpectedUeBehaviourData1) HasScheduledCommunicationType() bool`

HasScheduledCommunicationType returns a boolean if a field has been set.

### GetExpectedUmts

`func (o *ExpectedUeBehaviourData1) GetExpectedUmts() []LocationArea1`

GetExpectedUmts returns the ExpectedUmts field if non-nil, zero value otherwise.

### GetExpectedUmtsOk

`func (o *ExpectedUeBehaviourData1) GetExpectedUmtsOk() (*[]LocationArea1, bool)`

GetExpectedUmtsOk returns a tuple with the ExpectedUmts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUmts

`func (o *ExpectedUeBehaviourData1) SetExpectedUmts(v []LocationArea1)`

SetExpectedUmts sets ExpectedUmts field to given value.

### HasExpectedUmts

`func (o *ExpectedUeBehaviourData1) HasExpectedUmts() bool`

HasExpectedUmts returns a boolean if a field has been set.

### GetTrafficProfile

`func (o *ExpectedUeBehaviourData1) GetTrafficProfile() TrafficProfile`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *ExpectedUeBehaviourData1) GetTrafficProfileOk() (*TrafficProfile, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *ExpectedUeBehaviourData1) SetTrafficProfile(v TrafficProfile)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *ExpectedUeBehaviourData1) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetBatteryIndication

`func (o *ExpectedUeBehaviourData1) GetBatteryIndication() BatteryIndication`

GetBatteryIndication returns the BatteryIndication field if non-nil, zero value otherwise.

### GetBatteryIndicationOk

`func (o *ExpectedUeBehaviourData1) GetBatteryIndicationOk() (*BatteryIndication, bool)`

GetBatteryIndicationOk returns a tuple with the BatteryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryIndication

`func (o *ExpectedUeBehaviourData1) SetBatteryIndication(v BatteryIndication)`

SetBatteryIndication sets BatteryIndication field to given value.

### HasBatteryIndication

`func (o *ExpectedUeBehaviourData1) HasBatteryIndication() bool`

HasBatteryIndication returns a boolean if a field has been set.

### GetValidityTime

`func (o *ExpectedUeBehaviourData1) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *ExpectedUeBehaviourData1) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *ExpectedUeBehaviourData1) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *ExpectedUeBehaviourData1) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


