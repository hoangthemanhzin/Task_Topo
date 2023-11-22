# UEContextTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**UeContextTransferReqData**](UeContextTransferReqData.md) |  | [optional] 
**BinaryDataN1Message** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewUEContextTransferRequest

`func NewUEContextTransferRequest() *UEContextTransferRequest`

NewUEContextTransferRequest instantiates a new UEContextTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUEContextTransferRequestWithDefaults

`func NewUEContextTransferRequestWithDefaults() *UEContextTransferRequest`

NewUEContextTransferRequestWithDefaults instantiates a new UEContextTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *UEContextTransferRequest) GetJsonData() UeContextTransferReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *UEContextTransferRequest) GetJsonDataOk() (*UeContextTransferReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *UEContextTransferRequest) SetJsonData(v UeContextTransferReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *UEContextTransferRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1Message

`func (o *UEContextTransferRequest) GetBinaryDataN1Message() []byte`

GetBinaryDataN1Message returns the BinaryDataN1Message field if non-nil, zero value otherwise.

### GetBinaryDataN1MessageOk

`func (o *UEContextTransferRequest) GetBinaryDataN1MessageOk() (*[]byte, bool)`

GetBinaryDataN1MessageOk returns a tuple with the BinaryDataN1Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1Message

`func (o *UEContextTransferRequest) SetBinaryDataN1Message(v []byte)`

SetBinaryDataN1Message sets BinaryDataN1Message field to given value.

### HasBinaryDataN1Message

`func (o *UEContextTransferRequest) HasBinaryDataN1Message() bool`

HasBinaryDataN1Message returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


