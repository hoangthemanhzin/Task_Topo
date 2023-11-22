# CancelRelocateUEContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**UeContextCancelRelocateData**](UeContextCancelRelocateData.md) |  | [optional] 
**BinaryDataGtpcMessage** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewCancelRelocateUEContextRequest

`func NewCancelRelocateUEContextRequest() *CancelRelocateUEContextRequest`

NewCancelRelocateUEContextRequest instantiates a new CancelRelocateUEContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelRelocateUEContextRequestWithDefaults

`func NewCancelRelocateUEContextRequestWithDefaults() *CancelRelocateUEContextRequest`

NewCancelRelocateUEContextRequestWithDefaults instantiates a new CancelRelocateUEContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *CancelRelocateUEContextRequest) GetJsonData() UeContextCancelRelocateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *CancelRelocateUEContextRequest) GetJsonDataOk() (*UeContextCancelRelocateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *CancelRelocateUEContextRequest) SetJsonData(v UeContextCancelRelocateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *CancelRelocateUEContextRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataGtpcMessage

`func (o *CancelRelocateUEContextRequest) GetBinaryDataGtpcMessage() []byte`

GetBinaryDataGtpcMessage returns the BinaryDataGtpcMessage field if non-nil, zero value otherwise.

### GetBinaryDataGtpcMessageOk

`func (o *CancelRelocateUEContextRequest) GetBinaryDataGtpcMessageOk() (*[]byte, bool)`

GetBinaryDataGtpcMessageOk returns a tuple with the BinaryDataGtpcMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataGtpcMessage

`func (o *CancelRelocateUEContextRequest) SetBinaryDataGtpcMessage(v []byte)`

SetBinaryDataGtpcMessage sets BinaryDataGtpcMessage field to given value.

### HasBinaryDataGtpcMessage

`func (o *CancelRelocateUEContextRequest) HasBinaryDataGtpcMessage() bool`

HasBinaryDataGtpcMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


