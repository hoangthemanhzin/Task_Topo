# TransferMoDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**TransferMoDataReqData**](TransferMoDataReqData.md) |  | [optional] 
**BinaryMoData** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewTransferMoDataRequest

`func NewTransferMoDataRequest() *TransferMoDataRequest`

NewTransferMoDataRequest instantiates a new TransferMoDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferMoDataRequestWithDefaults

`func NewTransferMoDataRequestWithDefaults() *TransferMoDataRequest`

NewTransferMoDataRequestWithDefaults instantiates a new TransferMoDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *TransferMoDataRequest) GetJsonData() TransferMoDataReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *TransferMoDataRequest) GetJsonDataOk() (*TransferMoDataReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *TransferMoDataRequest) SetJsonData(v TransferMoDataReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *TransferMoDataRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryMoData

`func (o *TransferMoDataRequest) GetBinaryMoData() []byte`

GetBinaryMoData returns the BinaryMoData field if non-nil, zero value otherwise.

### GetBinaryMoDataOk

`func (o *TransferMoDataRequest) GetBinaryMoDataOk() (*[]byte, bool)`

GetBinaryMoDataOk returns a tuple with the BinaryMoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMoData

`func (o *TransferMoDataRequest) SetBinaryMoData(v []byte)`

SetBinaryMoData sets BinaryMoData field to given value.

### HasBinaryMoData

`func (o *TransferMoDataRequest) HasBinaryMoData() bool`

HasBinaryMoData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


