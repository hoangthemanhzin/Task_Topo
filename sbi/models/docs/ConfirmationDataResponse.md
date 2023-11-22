# ConfirmationDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResult** | [**AuthResult**](AuthResult.md) |  | 
**Supi** | Pointer to **string** |  | [optional] 
**Kseaf** | Pointer to **string** |  | [optional] 

## Methods

### NewConfirmationDataResponse

`func NewConfirmationDataResponse(authResult AuthResult, ) *ConfirmationDataResponse`

NewConfirmationDataResponse instantiates a new ConfirmationDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmationDataResponseWithDefaults

`func NewConfirmationDataResponseWithDefaults() *ConfirmationDataResponse`

NewConfirmationDataResponseWithDefaults instantiates a new ConfirmationDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResult

`func (o *ConfirmationDataResponse) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *ConfirmationDataResponse) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *ConfirmationDataResponse) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.


### GetSupi

`func (o *ConfirmationDataResponse) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ConfirmationDataResponse) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ConfirmationDataResponse) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ConfirmationDataResponse) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetKseaf

`func (o *ConfirmationDataResponse) GetKseaf() string`

GetKseaf returns the Kseaf field if non-nil, zero value otherwise.

### GetKseafOk

`func (o *ConfirmationDataResponse) GetKseafOk() (*string, bool)`

GetKseafOk returns a tuple with the Kseaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKseaf

`func (o *ConfirmationDataResponse) SetKseaf(v string)`

SetKseaf sets Kseaf field to given value.

### HasKseaf

`func (o *ConfirmationDataResponse) HasKseaf() bool`

HasKseaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


