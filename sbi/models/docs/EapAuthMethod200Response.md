# EapAuthMethod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EapPayload** | **string** | contains an EAP packet | 
**Links** | [**map[string]LinksValueSchema**](LinksValueSchema.md) | URI : /{eapSessionUri} | 

## Methods

### NewEapAuthMethod200Response

`func NewEapAuthMethod200Response(eapPayload string, links map[string]LinksValueSchema, ) *EapAuthMethod200Response`

NewEapAuthMethod200Response instantiates a new EapAuthMethod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapAuthMethod200ResponseWithDefaults

`func NewEapAuthMethod200ResponseWithDefaults() *EapAuthMethod200Response`

NewEapAuthMethod200ResponseWithDefaults instantiates a new EapAuthMethod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEapPayload

`func (o *EapAuthMethod200Response) GetEapPayload() string`

GetEapPayload returns the EapPayload field if non-nil, zero value otherwise.

### GetEapPayloadOk

`func (o *EapAuthMethod200Response) GetEapPayloadOk() (*string, bool)`

GetEapPayloadOk returns a tuple with the EapPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapPayload

`func (o *EapAuthMethod200Response) SetEapPayload(v string)`

SetEapPayload sets EapPayload field to given value.


### SetEapPayloadNil

`func (o *EapAuthMethod200Response) SetEapPayloadNil(b bool)`

 SetEapPayloadNil sets the value for EapPayload to be an explicit nil

### UnsetEapPayload
`func (o *EapAuthMethod200Response) UnsetEapPayload()`

UnsetEapPayload ensures that no value is present for EapPayload, not even an explicit nil
### GetLinks

`func (o *EapAuthMethod200Response) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EapAuthMethod200Response) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EapAuthMethod200Response) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


