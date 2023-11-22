# UserIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** |  | 
**Gpsi** | Pointer to **string** |  | [optional] 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUserIdentifier

`func NewUserIdentifier(supi string, ) *UserIdentifier`

NewUserIdentifier instantiates a new UserIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentifierWithDefaults

`func NewUserIdentifierWithDefaults() *UserIdentifier`

NewUserIdentifierWithDefaults instantiates a new UserIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UserIdentifier) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UserIdentifier) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UserIdentifier) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsi

`func (o *UserIdentifier) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UserIdentifier) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UserIdentifier) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UserIdentifier) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetValidityTime

`func (o *UserIdentifier) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *UserIdentifier) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *UserIdentifier) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *UserIdentifier) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


