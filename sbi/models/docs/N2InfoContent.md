# N2InfoContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NgapMessageType** | Pointer to **int32** |  | [optional] 
**NgapIeType** | Pointer to [**NgapIeType**](NgapIeType.md) |  | [optional] 
**NgapData** | [**RefToBinaryData**](RefToBinaryData.md) |  | 

## Methods

### NewN2InfoContent

`func NewN2InfoContent(ngapData RefToBinaryData, ) *N2InfoContent`

NewN2InfoContent instantiates a new N2InfoContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InfoContentWithDefaults

`func NewN2InfoContentWithDefaults() *N2InfoContent`

NewN2InfoContentWithDefaults instantiates a new N2InfoContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNgapMessageType

`func (o *N2InfoContent) GetNgapMessageType() int32`

GetNgapMessageType returns the NgapMessageType field if non-nil, zero value otherwise.

### GetNgapMessageTypeOk

`func (o *N2InfoContent) GetNgapMessageTypeOk() (*int32, bool)`

GetNgapMessageTypeOk returns a tuple with the NgapMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapMessageType

`func (o *N2InfoContent) SetNgapMessageType(v int32)`

SetNgapMessageType sets NgapMessageType field to given value.

### HasNgapMessageType

`func (o *N2InfoContent) HasNgapMessageType() bool`

HasNgapMessageType returns a boolean if a field has been set.

### GetNgapIeType

`func (o *N2InfoContent) GetNgapIeType() NgapIeType`

GetNgapIeType returns the NgapIeType field if non-nil, zero value otherwise.

### GetNgapIeTypeOk

`func (o *N2InfoContent) GetNgapIeTypeOk() (*NgapIeType, bool)`

GetNgapIeTypeOk returns a tuple with the NgapIeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapIeType

`func (o *N2InfoContent) SetNgapIeType(v NgapIeType)`

SetNgapIeType sets NgapIeType field to given value.

### HasNgapIeType

`func (o *N2InfoContent) HasNgapIeType() bool`

HasNgapIeType returns a boolean if a field has been set.

### GetNgapData

`func (o *N2InfoContent) GetNgapData() RefToBinaryData`

GetNgapData returns the NgapData field if non-nil, zero value otherwise.

### GetNgapDataOk

`func (o *N2InfoContent) GetNgapDataOk() (*RefToBinaryData, bool)`

GetNgapDataOk returns a tuple with the NgapData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapData

`func (o *N2InfoContent) SetNgapData(v RefToBinaryData)`

SetNgapData sets NgapData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


