# UeContextCancelRelocateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** |  | [optional] 
**RelocationCancelRequest** | [**RefToBinaryData**](RefToBinaryData.md) |  | 

## Methods

### NewUeContextCancelRelocateData

`func NewUeContextCancelRelocateData(relocationCancelRequest RefToBinaryData, ) *UeContextCancelRelocateData`

NewUeContextCancelRelocateData instantiates a new UeContextCancelRelocateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextCancelRelocateDataWithDefaults

`func NewUeContextCancelRelocateDataWithDefaults() *UeContextCancelRelocateData`

NewUeContextCancelRelocateDataWithDefaults instantiates a new UeContextCancelRelocateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeContextCancelRelocateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeContextCancelRelocateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeContextCancelRelocateData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeContextCancelRelocateData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetRelocationCancelRequest

`func (o *UeContextCancelRelocateData) GetRelocationCancelRequest() RefToBinaryData`

GetRelocationCancelRequest returns the RelocationCancelRequest field if non-nil, zero value otherwise.

### GetRelocationCancelRequestOk

`func (o *UeContextCancelRelocateData) GetRelocationCancelRequestOk() (*RefToBinaryData, bool)`

GetRelocationCancelRequestOk returns a tuple with the RelocationCancelRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocationCancelRequest

`func (o *UeContextCancelRelocateData) SetRelocationCancelRequest(v RefToBinaryData)`

SetRelocationCancelRequest sets RelocationCancelRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


