# ExpectedUeBehaviour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** |  | 
**StationaryIndication** | Pointer to [**StationaryIndicationRm**](StationaryIndicationRm.md) |  | [optional] 
**CommunicationDurationTime** | Pointer to **int32** |  | [optional] 
**ScheduledCommunicationType** | Pointer to [**ScheduledCommunicationTypeRm**](ScheduledCommunicationTypeRm.md) |  | [optional] 
**PeriodicTime** | Pointer to **int32** |  | [optional] 
**ScheduledCommunicationTime** | Pointer to [**ScheduledCommunicationTimeRm**](ScheduledCommunicationTimeRm.md) |  | [optional] 
**ExpectedUmts** | Pointer to [**[]LocationArea**](LocationArea.md) | Identifies the UE&#39;s expected geographical movement. The attribute is only applicable in 5G. | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfileRm**](TrafficProfileRm.md) |  | [optional] 
**BatteryIndication** | Pointer to [**BatteryIndicationRm**](BatteryIndicationRm.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 
**MtcProviderInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewExpectedUeBehaviour

`func NewExpectedUeBehaviour(afInstanceId string, referenceId int32, ) *ExpectedUeBehaviour`

NewExpectedUeBehaviour instantiates a new ExpectedUeBehaviour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedUeBehaviourWithDefaults

`func NewExpectedUeBehaviourWithDefaults() *ExpectedUeBehaviour`

NewExpectedUeBehaviourWithDefaults instantiates a new ExpectedUeBehaviour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *ExpectedUeBehaviour) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *ExpectedUeBehaviour) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *ExpectedUeBehaviour) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *ExpectedUeBehaviour) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ExpectedUeBehaviour) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ExpectedUeBehaviour) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetStationaryIndication

`func (o *ExpectedUeBehaviour) GetStationaryIndication() StationaryIndicationRm`

GetStationaryIndication returns the StationaryIndication field if non-nil, zero value otherwise.

### GetStationaryIndicationOk

`func (o *ExpectedUeBehaviour) GetStationaryIndicationOk() (*StationaryIndicationRm, bool)`

GetStationaryIndicationOk returns a tuple with the StationaryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryIndication

`func (o *ExpectedUeBehaviour) SetStationaryIndication(v StationaryIndicationRm)`

SetStationaryIndication sets StationaryIndication field to given value.

### HasStationaryIndication

`func (o *ExpectedUeBehaviour) HasStationaryIndication() bool`

HasStationaryIndication returns a boolean if a field has been set.

### GetCommunicationDurationTime

`func (o *ExpectedUeBehaviour) GetCommunicationDurationTime() int32`

GetCommunicationDurationTime returns the CommunicationDurationTime field if non-nil, zero value otherwise.

### GetCommunicationDurationTimeOk

`func (o *ExpectedUeBehaviour) GetCommunicationDurationTimeOk() (*int32, bool)`

GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDurationTime

`func (o *ExpectedUeBehaviour) SetCommunicationDurationTime(v int32)`

SetCommunicationDurationTime sets CommunicationDurationTime field to given value.

### HasCommunicationDurationTime

`func (o *ExpectedUeBehaviour) HasCommunicationDurationTime() bool`

HasCommunicationDurationTime returns a boolean if a field has been set.

### SetCommunicationDurationTimeNil

`func (o *ExpectedUeBehaviour) SetCommunicationDurationTimeNil(b bool)`

 SetCommunicationDurationTimeNil sets the value for CommunicationDurationTime to be an explicit nil

### UnsetCommunicationDurationTime
`func (o *ExpectedUeBehaviour) UnsetCommunicationDurationTime()`

UnsetCommunicationDurationTime ensures that no value is present for CommunicationDurationTime, not even an explicit nil
### GetScheduledCommunicationType

`func (o *ExpectedUeBehaviour) GetScheduledCommunicationType() ScheduledCommunicationTypeRm`

GetScheduledCommunicationType returns the ScheduledCommunicationType field if non-nil, zero value otherwise.

### GetScheduledCommunicationTypeOk

`func (o *ExpectedUeBehaviour) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationTypeRm, bool)`

GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationType

`func (o *ExpectedUeBehaviour) SetScheduledCommunicationType(v ScheduledCommunicationTypeRm)`

SetScheduledCommunicationType sets ScheduledCommunicationType field to given value.

### HasScheduledCommunicationType

`func (o *ExpectedUeBehaviour) HasScheduledCommunicationType() bool`

HasScheduledCommunicationType returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *ExpectedUeBehaviour) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *ExpectedUeBehaviour) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *ExpectedUeBehaviour) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *ExpectedUeBehaviour) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### SetPeriodicTimeNil

`func (o *ExpectedUeBehaviour) SetPeriodicTimeNil(b bool)`

 SetPeriodicTimeNil sets the value for PeriodicTime to be an explicit nil

### UnsetPeriodicTime
`func (o *ExpectedUeBehaviour) UnsetPeriodicTime()`

UnsetPeriodicTime ensures that no value is present for PeriodicTime, not even an explicit nil
### GetScheduledCommunicationTime

`func (o *ExpectedUeBehaviour) GetScheduledCommunicationTime() ScheduledCommunicationTimeRm`

GetScheduledCommunicationTime returns the ScheduledCommunicationTime field if non-nil, zero value otherwise.

### GetScheduledCommunicationTimeOk

`func (o *ExpectedUeBehaviour) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTimeRm, bool)`

GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationTime

`func (o *ExpectedUeBehaviour) SetScheduledCommunicationTime(v ScheduledCommunicationTimeRm)`

SetScheduledCommunicationTime sets ScheduledCommunicationTime field to given value.

### HasScheduledCommunicationTime

`func (o *ExpectedUeBehaviour) HasScheduledCommunicationTime() bool`

HasScheduledCommunicationTime returns a boolean if a field has been set.

### GetExpectedUmts

`func (o *ExpectedUeBehaviour) GetExpectedUmts() []LocationArea`

GetExpectedUmts returns the ExpectedUmts field if non-nil, zero value otherwise.

### GetExpectedUmtsOk

`func (o *ExpectedUeBehaviour) GetExpectedUmtsOk() (*[]LocationArea, bool)`

GetExpectedUmtsOk returns a tuple with the ExpectedUmts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUmts

`func (o *ExpectedUeBehaviour) SetExpectedUmts(v []LocationArea)`

SetExpectedUmts sets ExpectedUmts field to given value.

### HasExpectedUmts

`func (o *ExpectedUeBehaviour) HasExpectedUmts() bool`

HasExpectedUmts returns a boolean if a field has been set.

### SetExpectedUmtsNil

`func (o *ExpectedUeBehaviour) SetExpectedUmtsNil(b bool)`

 SetExpectedUmtsNil sets the value for ExpectedUmts to be an explicit nil

### UnsetExpectedUmts
`func (o *ExpectedUeBehaviour) UnsetExpectedUmts()`

UnsetExpectedUmts ensures that no value is present for ExpectedUmts, not even an explicit nil
### GetTrafficProfile

`func (o *ExpectedUeBehaviour) GetTrafficProfile() TrafficProfileRm`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *ExpectedUeBehaviour) GetTrafficProfileOk() (*TrafficProfileRm, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *ExpectedUeBehaviour) SetTrafficProfile(v TrafficProfileRm)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *ExpectedUeBehaviour) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetBatteryIndication

`func (o *ExpectedUeBehaviour) GetBatteryIndication() BatteryIndicationRm`

GetBatteryIndication returns the BatteryIndication field if non-nil, zero value otherwise.

### GetBatteryIndicationOk

`func (o *ExpectedUeBehaviour) GetBatteryIndicationOk() (*BatteryIndicationRm, bool)`

GetBatteryIndicationOk returns a tuple with the BatteryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryIndication

`func (o *ExpectedUeBehaviour) SetBatteryIndication(v BatteryIndicationRm)`

SetBatteryIndication sets BatteryIndication field to given value.

### HasBatteryIndication

`func (o *ExpectedUeBehaviour) HasBatteryIndication() bool`

HasBatteryIndication returns a boolean if a field has been set.

### GetValidityTime

`func (o *ExpectedUeBehaviour) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *ExpectedUeBehaviour) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *ExpectedUeBehaviour) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *ExpectedUeBehaviour) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *ExpectedUeBehaviour) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *ExpectedUeBehaviour) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *ExpectedUeBehaviour) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *ExpectedUeBehaviour) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


