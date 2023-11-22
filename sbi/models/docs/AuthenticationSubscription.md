# AuthenticationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | [**AuthMethod**](AuthMethod.md) |  | 
**EncPermanentKey** | Pointer to **string** |  | [optional] 
**ProtectionParameterId** | Pointer to **string** |  | [optional] 
**SequenceNumber** | Pointer to [**SequenceNumber**](SequenceNumber.md) |  | [optional] 
**AuthenticationManagementField** | Pointer to **string** |  | [optional] 
**AlgorithmId** | Pointer to **string** |  | [optional] 
**EncOpcKey** | Pointer to **string** |  | [optional] 
**EncTopcKey** | Pointer to **string** |  | [optional] 
**VectorGenerationInHss** | Pointer to **bool** |  | [optional] [default to false]
**N5gcAuthMethod** | Pointer to [**AuthMethod**](AuthMethod.md) |  | [optional] 
**RgAuthenticationInd** | Pointer to **bool** |  | [optional] [default to false]
**Supi** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationSubscription

`func NewAuthenticationSubscription(authenticationMethod AuthMethod, ) *AuthenticationSubscription`

NewAuthenticationSubscription instantiates a new AuthenticationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSubscriptionWithDefaults

`func NewAuthenticationSubscriptionWithDefaults() *AuthenticationSubscription`

NewAuthenticationSubscriptionWithDefaults instantiates a new AuthenticationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethod

`func (o *AuthenticationSubscription) GetAuthenticationMethod() AuthMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AuthenticationSubscription) GetAuthenticationMethodOk() (*AuthMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AuthenticationSubscription) SetAuthenticationMethod(v AuthMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetEncPermanentKey

`func (o *AuthenticationSubscription) GetEncPermanentKey() string`

GetEncPermanentKey returns the EncPermanentKey field if non-nil, zero value otherwise.

### GetEncPermanentKeyOk

`func (o *AuthenticationSubscription) GetEncPermanentKeyOk() (*string, bool)`

GetEncPermanentKeyOk returns a tuple with the EncPermanentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncPermanentKey

`func (o *AuthenticationSubscription) SetEncPermanentKey(v string)`

SetEncPermanentKey sets EncPermanentKey field to given value.

### HasEncPermanentKey

`func (o *AuthenticationSubscription) HasEncPermanentKey() bool`

HasEncPermanentKey returns a boolean if a field has been set.

### GetProtectionParameterId

`func (o *AuthenticationSubscription) GetProtectionParameterId() string`

GetProtectionParameterId returns the ProtectionParameterId field if non-nil, zero value otherwise.

### GetProtectionParameterIdOk

`func (o *AuthenticationSubscription) GetProtectionParameterIdOk() (*string, bool)`

GetProtectionParameterIdOk returns a tuple with the ProtectionParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionParameterId

`func (o *AuthenticationSubscription) SetProtectionParameterId(v string)`

SetProtectionParameterId sets ProtectionParameterId field to given value.

### HasProtectionParameterId

`func (o *AuthenticationSubscription) HasProtectionParameterId() bool`

HasProtectionParameterId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *AuthenticationSubscription) GetSequenceNumber() SequenceNumber`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *AuthenticationSubscription) GetSequenceNumberOk() (*SequenceNumber, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *AuthenticationSubscription) SetSequenceNumber(v SequenceNumber)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *AuthenticationSubscription) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetAuthenticationManagementField

`func (o *AuthenticationSubscription) GetAuthenticationManagementField() string`

GetAuthenticationManagementField returns the AuthenticationManagementField field if non-nil, zero value otherwise.

### GetAuthenticationManagementFieldOk

`func (o *AuthenticationSubscription) GetAuthenticationManagementFieldOk() (*string, bool)`

GetAuthenticationManagementFieldOk returns a tuple with the AuthenticationManagementField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationManagementField

`func (o *AuthenticationSubscription) SetAuthenticationManagementField(v string)`

SetAuthenticationManagementField sets AuthenticationManagementField field to given value.

### HasAuthenticationManagementField

`func (o *AuthenticationSubscription) HasAuthenticationManagementField() bool`

HasAuthenticationManagementField returns a boolean if a field has been set.

### GetAlgorithmId

`func (o *AuthenticationSubscription) GetAlgorithmId() string`

GetAlgorithmId returns the AlgorithmId field if non-nil, zero value otherwise.

### GetAlgorithmIdOk

`func (o *AuthenticationSubscription) GetAlgorithmIdOk() (*string, bool)`

GetAlgorithmIdOk returns a tuple with the AlgorithmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmId

`func (o *AuthenticationSubscription) SetAlgorithmId(v string)`

SetAlgorithmId sets AlgorithmId field to given value.

### HasAlgorithmId

`func (o *AuthenticationSubscription) HasAlgorithmId() bool`

HasAlgorithmId returns a boolean if a field has been set.

### GetEncOpcKey

`func (o *AuthenticationSubscription) GetEncOpcKey() string`

GetEncOpcKey returns the EncOpcKey field if non-nil, zero value otherwise.

### GetEncOpcKeyOk

`func (o *AuthenticationSubscription) GetEncOpcKeyOk() (*string, bool)`

GetEncOpcKeyOk returns a tuple with the EncOpcKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncOpcKey

`func (o *AuthenticationSubscription) SetEncOpcKey(v string)`

SetEncOpcKey sets EncOpcKey field to given value.

### HasEncOpcKey

`func (o *AuthenticationSubscription) HasEncOpcKey() bool`

HasEncOpcKey returns a boolean if a field has been set.

### GetEncTopcKey

`func (o *AuthenticationSubscription) GetEncTopcKey() string`

GetEncTopcKey returns the EncTopcKey field if non-nil, zero value otherwise.

### GetEncTopcKeyOk

`func (o *AuthenticationSubscription) GetEncTopcKeyOk() (*string, bool)`

GetEncTopcKeyOk returns a tuple with the EncTopcKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncTopcKey

`func (o *AuthenticationSubscription) SetEncTopcKey(v string)`

SetEncTopcKey sets EncTopcKey field to given value.

### HasEncTopcKey

`func (o *AuthenticationSubscription) HasEncTopcKey() bool`

HasEncTopcKey returns a boolean if a field has been set.

### GetVectorGenerationInHss

`func (o *AuthenticationSubscription) GetVectorGenerationInHss() bool`

GetVectorGenerationInHss returns the VectorGenerationInHss field if non-nil, zero value otherwise.

### GetVectorGenerationInHssOk

`func (o *AuthenticationSubscription) GetVectorGenerationInHssOk() (*bool, bool)`

GetVectorGenerationInHssOk returns a tuple with the VectorGenerationInHss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorGenerationInHss

`func (o *AuthenticationSubscription) SetVectorGenerationInHss(v bool)`

SetVectorGenerationInHss sets VectorGenerationInHss field to given value.

### HasVectorGenerationInHss

`func (o *AuthenticationSubscription) HasVectorGenerationInHss() bool`

HasVectorGenerationInHss returns a boolean if a field has been set.

### GetN5gcAuthMethod

`func (o *AuthenticationSubscription) GetN5gcAuthMethod() AuthMethod`

GetN5gcAuthMethod returns the N5gcAuthMethod field if non-nil, zero value otherwise.

### GetN5gcAuthMethodOk

`func (o *AuthenticationSubscription) GetN5gcAuthMethodOk() (*AuthMethod, bool)`

GetN5gcAuthMethodOk returns a tuple with the N5gcAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN5gcAuthMethod

`func (o *AuthenticationSubscription) SetN5gcAuthMethod(v AuthMethod)`

SetN5gcAuthMethod sets N5gcAuthMethod field to given value.

### HasN5gcAuthMethod

`func (o *AuthenticationSubscription) HasN5gcAuthMethod() bool`

HasN5gcAuthMethod returns a boolean if a field has been set.

### GetRgAuthenticationInd

`func (o *AuthenticationSubscription) GetRgAuthenticationInd() bool`

GetRgAuthenticationInd returns the RgAuthenticationInd field if non-nil, zero value otherwise.

### GetRgAuthenticationIndOk

`func (o *AuthenticationSubscription) GetRgAuthenticationIndOk() (*bool, bool)`

GetRgAuthenticationIndOk returns a tuple with the RgAuthenticationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgAuthenticationInd

`func (o *AuthenticationSubscription) SetRgAuthenticationInd(v bool)`

SetRgAuthenticationInd sets RgAuthenticationInd field to given value.

### HasRgAuthenticationInd

`func (o *AuthenticationSubscription) HasRgAuthenticationInd() bool`

HasRgAuthenticationInd returns a boolean if a field has been set.

### GetSupi

`func (o *AuthenticationSubscription) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AuthenticationSubscription) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AuthenticationSubscription) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AuthenticationSubscription) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


