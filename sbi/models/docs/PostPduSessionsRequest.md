# PostPduSessionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**PduSessionCreateData**](PduSessionCreateData.md) |  | [optional] 
**BinaryDataN1SmInfoFromUe** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 
**BinaryDataUnknownN1SmInfo** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewPostPduSessionsRequest

`func NewPostPduSessionsRequest() *PostPduSessionsRequest`

NewPostPduSessionsRequest instantiates a new PostPduSessionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPduSessionsRequestWithDefaults

`func NewPostPduSessionsRequestWithDefaults() *PostPduSessionsRequest`

NewPostPduSessionsRequestWithDefaults instantiates a new PostPduSessionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *PostPduSessionsRequest) GetJsonData() PduSessionCreateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *PostPduSessionsRequest) GetJsonDataOk() (*PduSessionCreateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *PostPduSessionsRequest) SetJsonData(v PduSessionCreateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *PostPduSessionsRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmInfoFromUe

`func (o *PostPduSessionsRequest) GetBinaryDataN1SmInfoFromUe() []byte`

GetBinaryDataN1SmInfoFromUe returns the BinaryDataN1SmInfoFromUe field if non-nil, zero value otherwise.

### GetBinaryDataN1SmInfoFromUeOk

`func (o *PostPduSessionsRequest) GetBinaryDataN1SmInfoFromUeOk() (*[]byte, bool)`

GetBinaryDataN1SmInfoFromUeOk returns a tuple with the BinaryDataN1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmInfoFromUe

`func (o *PostPduSessionsRequest) SetBinaryDataN1SmInfoFromUe(v []byte)`

SetBinaryDataN1SmInfoFromUe sets BinaryDataN1SmInfoFromUe field to given value.

### HasBinaryDataN1SmInfoFromUe

`func (o *PostPduSessionsRequest) HasBinaryDataN1SmInfoFromUe() bool`

HasBinaryDataN1SmInfoFromUe returns a boolean if a field has been set.

### GetBinaryDataUnknownN1SmInfo

`func (o *PostPduSessionsRequest) GetBinaryDataUnknownN1SmInfo() []byte`

GetBinaryDataUnknownN1SmInfo returns the BinaryDataUnknownN1SmInfo field if non-nil, zero value otherwise.

### GetBinaryDataUnknownN1SmInfoOk

`func (o *PostPduSessionsRequest) GetBinaryDataUnknownN1SmInfoOk() (*[]byte, bool)`

GetBinaryDataUnknownN1SmInfoOk returns a tuple with the BinaryDataUnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUnknownN1SmInfo

`func (o *PostPduSessionsRequest) SetBinaryDataUnknownN1SmInfo(v []byte)`

SetBinaryDataUnknownN1SmInfo sets BinaryDataUnknownN1SmInfo field to given value.

### HasBinaryDataUnknownN1SmInfo

`func (o *PostPduSessionsRequest) HasBinaryDataUnknownN1SmInfo() bool`

HasBinaryDataUnknownN1SmInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


