# N1N2MessageTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**N1N2MessageTransferReqData**](N1N2MessageTransferReqData.md) |  | [optional] 
**BinaryDataN1Message** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 
**BinaryMtData** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewN1N2MessageTransferRequest

`func NewN1N2MessageTransferRequest() *N1N2MessageTransferRequest`

NewN1N2MessageTransferRequest instantiates a new N1N2MessageTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1N2MessageTransferRequestWithDefaults

`func NewN1N2MessageTransferRequestWithDefaults() *N1N2MessageTransferRequest`

NewN1N2MessageTransferRequestWithDefaults instantiates a new N1N2MessageTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *N1N2MessageTransferRequest) GetJsonData() N1N2MessageTransferReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *N1N2MessageTransferRequest) GetJsonDataOk() (*N1N2MessageTransferReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *N1N2MessageTransferRequest) SetJsonData(v N1N2MessageTransferReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *N1N2MessageTransferRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1Message

`func (o *N1N2MessageTransferRequest) GetBinaryDataN1Message() []byte`

GetBinaryDataN1Message returns the BinaryDataN1Message field if non-nil, zero value otherwise.

### GetBinaryDataN1MessageOk

`func (o *N1N2MessageTransferRequest) GetBinaryDataN1MessageOk() (*[]byte, bool)`

GetBinaryDataN1MessageOk returns a tuple with the BinaryDataN1Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1Message

`func (o *N1N2MessageTransferRequest) SetBinaryDataN1Message(v []byte)`

SetBinaryDataN1Message sets BinaryDataN1Message field to given value.

### HasBinaryDataN1Message

`func (o *N1N2MessageTransferRequest) HasBinaryDataN1Message() bool`

HasBinaryDataN1Message returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *N1N2MessageTransferRequest) GetBinaryDataN2Information() []byte`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *N1N2MessageTransferRequest) GetBinaryDataN2InformationOk() (*[]byte, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *N1N2MessageTransferRequest) SetBinaryDataN2Information(v []byte)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *N1N2MessageTransferRequest) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.

### GetBinaryMtData

`func (o *N1N2MessageTransferRequest) GetBinaryMtData() []byte`

GetBinaryMtData returns the BinaryMtData field if non-nil, zero value otherwise.

### GetBinaryMtDataOk

`func (o *N1N2MessageTransferRequest) GetBinaryMtDataOk() (*[]byte, bool)`

GetBinaryMtDataOk returns a tuple with the BinaryMtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMtData

`func (o *N1N2MessageTransferRequest) SetBinaryMtData(v []byte)`

SetBinaryMtData sets BinaryMtData field to given value.

### HasBinaryMtData

`func (o *N1N2MessageTransferRequest) HasBinaryMtData() bool`

HasBinaryMtData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


