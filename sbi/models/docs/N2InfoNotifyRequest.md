# N2InfoNotifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**N2InformationNotification**](N2InformationNotification.md) |  | [optional] 
**BinaryDataN1Message** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewN2InfoNotifyRequest

`func NewN2InfoNotifyRequest() *N2InfoNotifyRequest`

NewN2InfoNotifyRequest instantiates a new N2InfoNotifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InfoNotifyRequestWithDefaults

`func NewN2InfoNotifyRequestWithDefaults() *N2InfoNotifyRequest`

NewN2InfoNotifyRequestWithDefaults instantiates a new N2InfoNotifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *N2InfoNotifyRequest) GetJsonData() N2InformationNotification`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *N2InfoNotifyRequest) GetJsonDataOk() (*N2InformationNotification, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *N2InfoNotifyRequest) SetJsonData(v N2InformationNotification)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *N2InfoNotifyRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1Message

`func (o *N2InfoNotifyRequest) GetBinaryDataN1Message() []byte`

GetBinaryDataN1Message returns the BinaryDataN1Message field if non-nil, zero value otherwise.

### GetBinaryDataN1MessageOk

`func (o *N2InfoNotifyRequest) GetBinaryDataN1MessageOk() (*[]byte, bool)`

GetBinaryDataN1MessageOk returns a tuple with the BinaryDataN1Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1Message

`func (o *N2InfoNotifyRequest) SetBinaryDataN1Message(v []byte)`

SetBinaryDataN1Message sets BinaryDataN1Message field to given value.

### HasBinaryDataN1Message

`func (o *N2InfoNotifyRequest) HasBinaryDataN1Message() bool`

HasBinaryDataN1Message returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *N2InfoNotifyRequest) GetBinaryDataN2Information() []byte`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *N2InfoNotifyRequest) GetBinaryDataN2InformationOk() (*[]byte, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *N2InfoNotifyRequest) SetBinaryDataN2Information(v []byte)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *N2InfoNotifyRequest) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


