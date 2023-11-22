# AuthenticationInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**AuthType**](AuthType.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**AuthenticationVector** | Pointer to [**AuthenticationVector**](AuthenticationVector.md) |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationInfoResult

`func NewAuthenticationInfoResult(authType AuthType, ) *AuthenticationInfoResult`

NewAuthenticationInfoResult instantiates a new AuthenticationInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoResultWithDefaults

`func NewAuthenticationInfoResultWithDefaults() *AuthenticationInfoResult`

NewAuthenticationInfoResultWithDefaults instantiates a new AuthenticationInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthenticationInfoResult) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthenticationInfoResult) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthenticationInfoResult) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetSupportedFeatures

`func (o *AuthenticationInfoResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfoResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfoResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAuthenticationVector

`func (o *AuthenticationInfoResult) GetAuthenticationVector() AuthenticationVector`

GetAuthenticationVector returns the AuthenticationVector field if non-nil, zero value otherwise.

### GetAuthenticationVectorOk

`func (o *AuthenticationInfoResult) GetAuthenticationVectorOk() (*AuthenticationVector, bool)`

GetAuthenticationVectorOk returns a tuple with the AuthenticationVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationVector

`func (o *AuthenticationInfoResult) SetAuthenticationVector(v AuthenticationVector)`

SetAuthenticationVector sets AuthenticationVector field to given value.

### HasAuthenticationVector

`func (o *AuthenticationInfoResult) HasAuthenticationVector() bool`

HasAuthenticationVector returns a boolean if a field has been set.

### GetSupi

`func (o *AuthenticationInfoResult) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AuthenticationInfoResult) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AuthenticationInfoResult) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AuthenticationInfoResult) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


