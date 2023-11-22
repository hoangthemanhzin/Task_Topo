# AuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | **string** |  | 
**MtcProviderInformation** | **string** |  | 
**AuthUpdateCallbackUri** | **string** |  | 
**AfId** | Pointer to **string** |  | [optional] 
**NefId** | Pointer to **string** | Identity of the NEF | [optional] 

## Methods

### NewAuthorizationInfo

`func NewAuthorizationInfo(snssai Snssai, dnn string, mtcProviderInformation string, authUpdateCallbackUri string, ) *AuthorizationInfo`

NewAuthorizationInfo instantiates a new AuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationInfoWithDefaults

`func NewAuthorizationInfoWithDefaults() *AuthorizationInfo`

NewAuthorizationInfoWithDefaults instantiates a new AuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *AuthorizationInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AuthorizationInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AuthorizationInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnn

`func (o *AuthorizationInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AuthorizationInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AuthorizationInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetMtcProviderInformation

`func (o *AuthorizationInfo) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *AuthorizationInfo) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *AuthorizationInfo) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.


### GetAuthUpdateCallbackUri

`func (o *AuthorizationInfo) GetAuthUpdateCallbackUri() string`

GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field if non-nil, zero value otherwise.

### GetAuthUpdateCallbackUriOk

`func (o *AuthorizationInfo) GetAuthUpdateCallbackUriOk() (*string, bool)`

GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUpdateCallbackUri

`func (o *AuthorizationInfo) SetAuthUpdateCallbackUri(v string)`

SetAuthUpdateCallbackUri sets AuthUpdateCallbackUri field to given value.


### GetAfId

`func (o *AuthorizationInfo) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *AuthorizationInfo) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *AuthorizationInfo) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *AuthorizationInfo) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetNefId

`func (o *AuthorizationInfo) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *AuthorizationInfo) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *AuthorizationInfo) SetNefId(v string)`

SetNefId sets NefId field to given value.

### HasNefId

`func (o *AuthorizationInfo) HasNefId() bool`

HasNefId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


