# AuthEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** |  | 
**Success** | **bool** |  | 
**TimeStamp** | **time.Time** |  | 
**AuthType** | [**AuthType**](AuthType.md) |  | 
**ServingNetworkName** | **string** |  | 
**AuthRemovalInd** | Pointer to **bool** |  | [optional] [default to false]
**NfSetId** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthEvent

`func NewAuthEvent(nfInstanceId string, success bool, timeStamp time.Time, authType AuthType, servingNetworkName string, ) *AuthEvent`

NewAuthEvent instantiates a new AuthEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthEventWithDefaults

`func NewAuthEventWithDefaults() *AuthEvent`

NewAuthEventWithDefaults instantiates a new AuthEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *AuthEvent) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *AuthEvent) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *AuthEvent) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetSuccess

`func (o *AuthEvent) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AuthEvent) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AuthEvent) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTimeStamp

`func (o *AuthEvent) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AuthEvent) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AuthEvent) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetAuthType

`func (o *AuthEvent) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthEvent) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthEvent) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetServingNetworkName

`func (o *AuthEvent) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *AuthEvent) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *AuthEvent) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.


### GetAuthRemovalInd

`func (o *AuthEvent) GetAuthRemovalInd() bool`

GetAuthRemovalInd returns the AuthRemovalInd field if non-nil, zero value otherwise.

### GetAuthRemovalIndOk

`func (o *AuthEvent) GetAuthRemovalIndOk() (*bool, bool)`

GetAuthRemovalIndOk returns a tuple with the AuthRemovalInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRemovalInd

`func (o *AuthEvent) SetAuthRemovalInd(v bool)`

SetAuthRemovalInd sets AuthRemovalInd field to given value.

### HasAuthRemovalInd

`func (o *AuthEvent) HasAuthRemovalInd() bool`

HasAuthRemovalInd returns a boolean if a field has been set.

### GetNfSetId

`func (o *AuthEvent) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *AuthEvent) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *AuthEvent) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.

### HasNfSetId

`func (o *AuthEvent) HasNfSetId() bool`

HasNfSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


