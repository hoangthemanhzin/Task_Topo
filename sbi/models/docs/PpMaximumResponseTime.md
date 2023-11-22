# PpMaximumResponseTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumResponseTime** | **int32** |  | 
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** |  | 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 
**MtcProviderInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewPpMaximumResponseTime

`func NewPpMaximumResponseTime(maximumResponseTime int32, afInstanceId string, referenceId int32, ) *PpMaximumResponseTime`

NewPpMaximumResponseTime instantiates a new PpMaximumResponseTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpMaximumResponseTimeWithDefaults

`func NewPpMaximumResponseTimeWithDefaults() *PpMaximumResponseTime`

NewPpMaximumResponseTimeWithDefaults instantiates a new PpMaximumResponseTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumResponseTime

`func (o *PpMaximumResponseTime) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *PpMaximumResponseTime) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *PpMaximumResponseTime) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.


### GetAfInstanceId

`func (o *PpMaximumResponseTime) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *PpMaximumResponseTime) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *PpMaximumResponseTime) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *PpMaximumResponseTime) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PpMaximumResponseTime) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PpMaximumResponseTime) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetValidityTime

`func (o *PpMaximumResponseTime) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *PpMaximumResponseTime) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *PpMaximumResponseTime) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *PpMaximumResponseTime) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *PpMaximumResponseTime) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *PpMaximumResponseTime) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *PpMaximumResponseTime) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *PpMaximumResponseTime) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


