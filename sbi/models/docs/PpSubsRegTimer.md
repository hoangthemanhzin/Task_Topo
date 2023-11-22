# PpSubsRegTimer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubsRegTimer** | **int32** |  | 
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** |  | 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 
**MtcProviderInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewPpSubsRegTimer

`func NewPpSubsRegTimer(subsRegTimer int32, afInstanceId string, referenceId int32, ) *PpSubsRegTimer`

NewPpSubsRegTimer instantiates a new PpSubsRegTimer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpSubsRegTimerWithDefaults

`func NewPpSubsRegTimerWithDefaults() *PpSubsRegTimer`

NewPpSubsRegTimerWithDefaults instantiates a new PpSubsRegTimer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubsRegTimer

`func (o *PpSubsRegTimer) GetSubsRegTimer() int32`

GetSubsRegTimer returns the SubsRegTimer field if non-nil, zero value otherwise.

### GetSubsRegTimerOk

`func (o *PpSubsRegTimer) GetSubsRegTimerOk() (*int32, bool)`

GetSubsRegTimerOk returns a tuple with the SubsRegTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsRegTimer

`func (o *PpSubsRegTimer) SetSubsRegTimer(v int32)`

SetSubsRegTimer sets SubsRegTimer field to given value.


### GetAfInstanceId

`func (o *PpSubsRegTimer) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *PpSubsRegTimer) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *PpSubsRegTimer) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *PpSubsRegTimer) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PpSubsRegTimer) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PpSubsRegTimer) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetValidityTime

`func (o *PpSubsRegTimer) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *PpSubsRegTimer) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *PpSubsRegTimer) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *PpSubsRegTimer) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *PpSubsRegTimer) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *PpSubsRegTimer) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *PpSubsRegTimer) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *PpSubsRegTimer) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


