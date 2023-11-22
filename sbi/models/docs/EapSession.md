# EapSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EapPayload** | **string** | contains an EAP packet | 
**KSeaf** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]LinksValueSchema**](LinksValueSchema.md) |  | [optional] 
**AuthResult** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewEapSession

`func NewEapSession(eapPayload string, ) *EapSession`

NewEapSession instantiates a new EapSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapSessionWithDefaults

`func NewEapSessionWithDefaults() *EapSession`

NewEapSessionWithDefaults instantiates a new EapSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEapPayload

`func (o *EapSession) GetEapPayload() string`

GetEapPayload returns the EapPayload field if non-nil, zero value otherwise.

### GetEapPayloadOk

`func (o *EapSession) GetEapPayloadOk() (*string, bool)`

GetEapPayloadOk returns a tuple with the EapPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapPayload

`func (o *EapSession) SetEapPayload(v string)`

SetEapPayload sets EapPayload field to given value.


### SetEapPayloadNil

`func (o *EapSession) SetEapPayloadNil(b bool)`

 SetEapPayloadNil sets the value for EapPayload to be an explicit nil

### UnsetEapPayload
`func (o *EapSession) UnsetEapPayload()`

UnsetEapPayload ensures that no value is present for EapPayload, not even an explicit nil
### GetKSeaf

`func (o *EapSession) GetKSeaf() string`

GetKSeaf returns the KSeaf field if non-nil, zero value otherwise.

### GetKSeafOk

`func (o *EapSession) GetKSeafOk() (*string, bool)`

GetKSeafOk returns a tuple with the KSeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKSeaf

`func (o *EapSession) SetKSeaf(v string)`

SetKSeaf sets KSeaf field to given value.

### HasKSeaf

`func (o *EapSession) HasKSeaf() bool`

HasKSeaf returns a boolean if a field has been set.

### GetLinks

`func (o *EapSession) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EapSession) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EapSession) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EapSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthResult

`func (o *EapSession) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *EapSession) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *EapSession) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *EapSession) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetSupi

`func (o *EapSession) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EapSession) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EapSession) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EapSession) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EapSession) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EapSession) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EapSession) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EapSession) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


