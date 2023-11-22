# UEAuthenticationCtx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**AuthType**](AuthType.md) |  | 
**Var5gAuthData** | [**UEAuthenticationCtx5gAuthData**](UEAuthenticationCtx5gAuthData.md) |  | 
**Links** | [**map[string]LinksValueSchema**](LinksValueSchema.md) |  | 
**ServingNetworkName** | Pointer to **string** |  | [optional] 

## Methods

### NewUEAuthenticationCtx

`func NewUEAuthenticationCtx(authType AuthType, var5gAuthData UEAuthenticationCtx5gAuthData, links map[string]LinksValueSchema, ) *UEAuthenticationCtx`

NewUEAuthenticationCtx instantiates a new UEAuthenticationCtx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUEAuthenticationCtxWithDefaults

`func NewUEAuthenticationCtxWithDefaults() *UEAuthenticationCtx`

NewUEAuthenticationCtxWithDefaults instantiates a new UEAuthenticationCtx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *UEAuthenticationCtx) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UEAuthenticationCtx) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UEAuthenticationCtx) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetVar5gAuthData

`func (o *UEAuthenticationCtx) GetVar5gAuthData() UEAuthenticationCtx5gAuthData`

GetVar5gAuthData returns the Var5gAuthData field if non-nil, zero value otherwise.

### GetVar5gAuthDataOk

`func (o *UEAuthenticationCtx) GetVar5gAuthDataOk() (*UEAuthenticationCtx5gAuthData, bool)`

GetVar5gAuthDataOk returns a tuple with the Var5gAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gAuthData

`func (o *UEAuthenticationCtx) SetVar5gAuthData(v UEAuthenticationCtx5gAuthData)`

SetVar5gAuthData sets Var5gAuthData field to given value.


### GetLinks

`func (o *UEAuthenticationCtx) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UEAuthenticationCtx) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UEAuthenticationCtx) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.


### GetServingNetworkName

`func (o *UEAuthenticationCtx) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *UEAuthenticationCtx) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *UEAuthenticationCtx) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.

### HasServingNetworkName

`func (o *UEAuthenticationCtx) HasServingNetworkName() bool`

HasServingNetworkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


