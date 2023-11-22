# AuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationData** | [**[]UserIdentifier**](UserIdentifier.md) |  | 
**ValidityTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAuthorizationData

`func NewAuthorizationData(authorizationData []UserIdentifier, ) *AuthorizationData`

NewAuthorizationData instantiates a new AuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDataWithDefaults

`func NewAuthorizationDataWithDefaults() *AuthorizationData`

NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationData

`func (o *AuthorizationData) GetAuthorizationData() []UserIdentifier`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *AuthorizationData) GetAuthorizationDataOk() (*[]UserIdentifier, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *AuthorizationData) SetAuthorizationData(v []UserIdentifier)`

SetAuthorizationData sets AuthorizationData field to given value.


### GetValidityTime

`func (o *AuthorizationData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AuthorizationData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AuthorizationData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *AuthorizationData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


