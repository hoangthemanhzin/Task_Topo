# N1MessageContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N1MessageClass** | [**N1MessageClass**](N1MessageClass.md) |  | 
**N1MessageContent** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**NfId** | Pointer to **string** |  | [optional] 
**ServiceInstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewN1MessageContainer

`func NewN1MessageContainer(n1MessageClass N1MessageClass, n1MessageContent RefToBinaryData, ) *N1MessageContainer`

NewN1MessageContainer instantiates a new N1MessageContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1MessageContainerWithDefaults

`func NewN1MessageContainerWithDefaults() *N1MessageContainer`

NewN1MessageContainerWithDefaults instantiates a new N1MessageContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN1MessageClass

`func (o *N1MessageContainer) GetN1MessageClass() N1MessageClass`

GetN1MessageClass returns the N1MessageClass field if non-nil, zero value otherwise.

### GetN1MessageClassOk

`func (o *N1MessageContainer) GetN1MessageClassOk() (*N1MessageClass, bool)`

GetN1MessageClassOk returns a tuple with the N1MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageClass

`func (o *N1MessageContainer) SetN1MessageClass(v N1MessageClass)`

SetN1MessageClass sets N1MessageClass field to given value.


### GetN1MessageContent

`func (o *N1MessageContainer) GetN1MessageContent() RefToBinaryData`

GetN1MessageContent returns the N1MessageContent field if non-nil, zero value otherwise.

### GetN1MessageContentOk

`func (o *N1MessageContainer) GetN1MessageContentOk() (*RefToBinaryData, bool)`

GetN1MessageContentOk returns a tuple with the N1MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageContent

`func (o *N1MessageContainer) SetN1MessageContent(v RefToBinaryData)`

SetN1MessageContent sets N1MessageContent field to given value.


### GetNfId

`func (o *N1MessageContainer) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *N1MessageContainer) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *N1MessageContainer) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *N1MessageContainer) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetServiceInstanceId

`func (o *N1MessageContainer) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *N1MessageContainer) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *N1MessageContainer) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *N1MessageContainer) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


