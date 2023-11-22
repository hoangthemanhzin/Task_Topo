# PWSResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NgapMessageType** | **int32** |  | 
**SerialNumber** | **int32** |  | 
**MessageIdentifier** | **int32** |  | 
**UnknownTaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 

## Methods

### NewPWSResponseData

`func NewPWSResponseData(ngapMessageType int32, serialNumber int32, messageIdentifier int32, ) *PWSResponseData`

NewPWSResponseData instantiates a new PWSResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPWSResponseDataWithDefaults

`func NewPWSResponseDataWithDefaults() *PWSResponseData`

NewPWSResponseDataWithDefaults instantiates a new PWSResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNgapMessageType

`func (o *PWSResponseData) GetNgapMessageType() int32`

GetNgapMessageType returns the NgapMessageType field if non-nil, zero value otherwise.

### GetNgapMessageTypeOk

`func (o *PWSResponseData) GetNgapMessageTypeOk() (*int32, bool)`

GetNgapMessageTypeOk returns a tuple with the NgapMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapMessageType

`func (o *PWSResponseData) SetNgapMessageType(v int32)`

SetNgapMessageType sets NgapMessageType field to given value.


### GetSerialNumber

`func (o *PWSResponseData) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PWSResponseData) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PWSResponseData) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.


### GetMessageIdentifier

`func (o *PWSResponseData) GetMessageIdentifier() int32`

GetMessageIdentifier returns the MessageIdentifier field if non-nil, zero value otherwise.

### GetMessageIdentifierOk

`func (o *PWSResponseData) GetMessageIdentifierOk() (*int32, bool)`

GetMessageIdentifierOk returns a tuple with the MessageIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIdentifier

`func (o *PWSResponseData) SetMessageIdentifier(v int32)`

SetMessageIdentifier sets MessageIdentifier field to given value.


### GetUnknownTaiList

`func (o *PWSResponseData) GetUnknownTaiList() []Tai`

GetUnknownTaiList returns the UnknownTaiList field if non-nil, zero value otherwise.

### GetUnknownTaiListOk

`func (o *PWSResponseData) GetUnknownTaiListOk() (*[]Tai, bool)`

GetUnknownTaiListOk returns a tuple with the UnknownTaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownTaiList

`func (o *PWSResponseData) SetUnknownTaiList(v []Tai)`

SetUnknownTaiList sets UnknownTaiList field to given value.

### HasUnknownTaiList

`func (o *PWSResponseData) HasUnknownTaiList() bool`

HasUnknownTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


