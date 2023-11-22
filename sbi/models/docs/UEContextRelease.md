# UEContextRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** |  | [optional] 
**UnauthenticatedSupi** | Pointer to **bool** |  | [optional] [default to false]
**NgapCause** | [**NgApCause**](NgApCause.md) |  | 

## Methods

### NewUEContextRelease

`func NewUEContextRelease(ngapCause NgApCause, ) *UEContextRelease`

NewUEContextRelease instantiates a new UEContextRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUEContextReleaseWithDefaults

`func NewUEContextReleaseWithDefaults() *UEContextRelease`

NewUEContextReleaseWithDefaults instantiates a new UEContextRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UEContextRelease) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UEContextRelease) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UEContextRelease) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UEContextRelease) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUnauthenticatedSupi

`func (o *UEContextRelease) GetUnauthenticatedSupi() bool`

GetUnauthenticatedSupi returns the UnauthenticatedSupi field if non-nil, zero value otherwise.

### GetUnauthenticatedSupiOk

`func (o *UEContextRelease) GetUnauthenticatedSupiOk() (*bool, bool)`

GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSupi

`func (o *UEContextRelease) SetUnauthenticatedSupi(v bool)`

SetUnauthenticatedSupi sets UnauthenticatedSupi field to given value.

### HasUnauthenticatedSupi

`func (o *UEContextRelease) HasUnauthenticatedSupi() bool`

HasUnauthenticatedSupi returns a boolean if a field has been set.

### GetNgapCause

`func (o *UEContextRelease) GetNgapCause() NgApCause`

GetNgapCause returns the NgapCause field if non-nil, zero value otherwise.

### GetNgapCauseOk

`func (o *UEContextRelease) GetNgapCauseOk() (*NgApCause, bool)`

GetNgapCauseOk returns a tuple with the NgapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapCause

`func (o *UEContextRelease) SetNgapCause(v NgApCause)`

SetNgapCause sets NgapCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


