# RgAuthCtx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResult** | [**AuthResult**](AuthResult.md) |  | 
**Supi** | Pointer to **string** |  | [optional] 
**AuthInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRgAuthCtx

`func NewRgAuthCtx(authResult AuthResult, ) *RgAuthCtx`

NewRgAuthCtx instantiates a new RgAuthCtx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRgAuthCtxWithDefaults

`func NewRgAuthCtxWithDefaults() *RgAuthCtx`

NewRgAuthCtxWithDefaults instantiates a new RgAuthCtx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResult

`func (o *RgAuthCtx) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *RgAuthCtx) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *RgAuthCtx) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.


### GetSupi

`func (o *RgAuthCtx) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *RgAuthCtx) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *RgAuthCtx) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *RgAuthCtx) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetAuthInd

`func (o *RgAuthCtx) GetAuthInd() bool`

GetAuthInd returns the AuthInd field if non-nil, zero value otherwise.

### GetAuthIndOk

`func (o *RgAuthCtx) GetAuthIndOk() (*bool, bool)`

GetAuthIndOk returns a tuple with the AuthInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInd

`func (o *RgAuthCtx) SetAuthInd(v bool)`

SetAuthInd sets AuthInd field to given value.

### HasAuthInd

`func (o *RgAuthCtx) HasAuthInd() bool`

HasAuthInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


