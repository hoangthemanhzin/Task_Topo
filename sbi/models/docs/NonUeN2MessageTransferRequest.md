# NonUeN2MessageTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**N2InformationTransferReqData**](N2InformationTransferReqData.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewNonUeN2MessageTransferRequest

`func NewNonUeN2MessageTransferRequest() *NonUeN2MessageTransferRequest`

NewNonUeN2MessageTransferRequest instantiates a new NonUeN2MessageTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonUeN2MessageTransferRequestWithDefaults

`func NewNonUeN2MessageTransferRequestWithDefaults() *NonUeN2MessageTransferRequest`

NewNonUeN2MessageTransferRequestWithDefaults instantiates a new NonUeN2MessageTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *NonUeN2MessageTransferRequest) GetJsonData() N2InformationTransferReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *NonUeN2MessageTransferRequest) GetJsonDataOk() (*N2InformationTransferReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *NonUeN2MessageTransferRequest) SetJsonData(v N2InformationTransferReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *NonUeN2MessageTransferRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *NonUeN2MessageTransferRequest) GetBinaryDataN2Information() []byte`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *NonUeN2MessageTransferRequest) GetBinaryDataN2InformationOk() (*[]byte, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *NonUeN2MessageTransferRequest) SetBinaryDataN2Information(v []byte)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *NonUeN2MessageTransferRequest) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


