# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerType** | [**TriggerType**](TriggerType.md) |  | 
**TriggerCategory** | [**TriggerCategory**](TriggerCategory.md) |  | 
**TimeLimit** | Pointer to **int32** |  | [optional] 
**VolumeLimit** | Pointer to **int32** |  | [optional] 
**VolumeLimit64** | Pointer to **int32** |  | [optional] 
**EventLimit** | Pointer to **int32** |  | [optional] 
**MaxNumberOfccc** | Pointer to **int32** |  | [optional] 
**TariffTimeChange** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTrigger

`func NewTrigger(triggerType TriggerType, triggerCategory TriggerCategory, ) *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerType

`func (o *Trigger) GetTriggerType() TriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *Trigger) GetTriggerTypeOk() (*TriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *Trigger) SetTriggerType(v TriggerType)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerCategory

`func (o *Trigger) GetTriggerCategory() TriggerCategory`

GetTriggerCategory returns the TriggerCategory field if non-nil, zero value otherwise.

### GetTriggerCategoryOk

`func (o *Trigger) GetTriggerCategoryOk() (*TriggerCategory, bool)`

GetTriggerCategoryOk returns a tuple with the TriggerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCategory

`func (o *Trigger) SetTriggerCategory(v TriggerCategory)`

SetTriggerCategory sets TriggerCategory field to given value.


### GetTimeLimit

`func (o *Trigger) GetTimeLimit() int32`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *Trigger) GetTimeLimitOk() (*int32, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *Trigger) SetTimeLimit(v int32)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *Trigger) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetVolumeLimit

`func (o *Trigger) GetVolumeLimit() int32`

GetVolumeLimit returns the VolumeLimit field if non-nil, zero value otherwise.

### GetVolumeLimitOk

`func (o *Trigger) GetVolumeLimitOk() (*int32, bool)`

GetVolumeLimitOk returns a tuple with the VolumeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLimit

`func (o *Trigger) SetVolumeLimit(v int32)`

SetVolumeLimit sets VolumeLimit field to given value.

### HasVolumeLimit

`func (o *Trigger) HasVolumeLimit() bool`

HasVolumeLimit returns a boolean if a field has been set.

### GetVolumeLimit64

`func (o *Trigger) GetVolumeLimit64() int32`

GetVolumeLimit64 returns the VolumeLimit64 field if non-nil, zero value otherwise.

### GetVolumeLimit64Ok

`func (o *Trigger) GetVolumeLimit64Ok() (*int32, bool)`

GetVolumeLimit64Ok returns a tuple with the VolumeLimit64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLimit64

`func (o *Trigger) SetVolumeLimit64(v int32)`

SetVolumeLimit64 sets VolumeLimit64 field to given value.

### HasVolumeLimit64

`func (o *Trigger) HasVolumeLimit64() bool`

HasVolumeLimit64 returns a boolean if a field has been set.

### GetEventLimit

`func (o *Trigger) GetEventLimit() int32`

GetEventLimit returns the EventLimit field if non-nil, zero value otherwise.

### GetEventLimitOk

`func (o *Trigger) GetEventLimitOk() (*int32, bool)`

GetEventLimitOk returns a tuple with the EventLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLimit

`func (o *Trigger) SetEventLimit(v int32)`

SetEventLimit sets EventLimit field to given value.

### HasEventLimit

`func (o *Trigger) HasEventLimit() bool`

HasEventLimit returns a boolean if a field has been set.

### GetMaxNumberOfccc

`func (o *Trigger) GetMaxNumberOfccc() int32`

GetMaxNumberOfccc returns the MaxNumberOfccc field if non-nil, zero value otherwise.

### GetMaxNumberOfcccOk

`func (o *Trigger) GetMaxNumberOfcccOk() (*int32, bool)`

GetMaxNumberOfcccOk returns a tuple with the MaxNumberOfccc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfccc

`func (o *Trigger) SetMaxNumberOfccc(v int32)`

SetMaxNumberOfccc sets MaxNumberOfccc field to given value.

### HasMaxNumberOfccc

`func (o *Trigger) HasMaxNumberOfccc() bool`

HasMaxNumberOfccc returns a boolean if a field has been set.

### GetTariffTimeChange

`func (o *Trigger) GetTariffTimeChange() time.Time`

GetTariffTimeChange returns the TariffTimeChange field if non-nil, zero value otherwise.

### GetTariffTimeChangeOk

`func (o *Trigger) GetTariffTimeChangeOk() (*time.Time, bool)`

GetTariffTimeChangeOk returns a tuple with the TariffTimeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffTimeChange

`func (o *Trigger) SetTariffTimeChange(v time.Time)`

SetTariffTimeChange sets TariffTimeChange field to given value.

### HasTariffTimeChange

`func (o *Trigger) HasTariffTimeChange() bool`

HasTariffTimeChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


