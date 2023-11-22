# DeliverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**DeliverReqData**](DeliverReqData.md) |  | [optional] 
**BinaryMtData** | Pointer to [**[]byte**]([]byte.md) |  | [optional] 

## Methods

### NewDeliverRequest

`func NewDeliverRequest() *DeliverRequest`

NewDeliverRequest instantiates a new DeliverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverRequestWithDefaults

`func NewDeliverRequestWithDefaults() *DeliverRequest`

NewDeliverRequestWithDefaults instantiates a new DeliverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *DeliverRequest) GetJsonData() DeliverReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *DeliverRequest) GetJsonDataOk() (*DeliverReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *DeliverRequest) SetJsonData(v DeliverReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *DeliverRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryMtData

`func (o *DeliverRequest) GetBinaryMtData() []byte`

GetBinaryMtData returns the BinaryMtData field if non-nil, zero value otherwise.

### GetBinaryMtDataOk

`func (o *DeliverRequest) GetBinaryMtDataOk() (*[]byte, bool)`

GetBinaryMtDataOk returns a tuple with the BinaryMtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryMtData

`func (o *DeliverRequest) SetBinaryMtData(v []byte)`

SetBinaryMtData sets BinaryMtData field to given value.

### HasBinaryMtData

`func (o *DeliverRequest) HasBinaryMtData() bool`

HasBinaryMtData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


