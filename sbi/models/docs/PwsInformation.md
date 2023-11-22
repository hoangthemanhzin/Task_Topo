# PwsInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageIdentifier** | **int32** |  | 
**SerialNumber** | **int32** |  | 
**PwsContainer** | [**N2InfoContent**](N2InfoContent.md) |  | 
**BcEmptyAreaList** | Pointer to [**[]GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**SendRanResponse** | Pointer to **bool** |  | [optional] [default to false]
**OmcId** | Pointer to **string** |  | [optional] 

## Methods

### NewPwsInformation

`func NewPwsInformation(messageIdentifier int32, serialNumber int32, pwsContainer N2InfoContent, ) *PwsInformation`

NewPwsInformation instantiates a new PwsInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwsInformationWithDefaults

`func NewPwsInformationWithDefaults() *PwsInformation`

NewPwsInformationWithDefaults instantiates a new PwsInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageIdentifier

`func (o *PwsInformation) GetMessageIdentifier() int32`

GetMessageIdentifier returns the MessageIdentifier field if non-nil, zero value otherwise.

### GetMessageIdentifierOk

`func (o *PwsInformation) GetMessageIdentifierOk() (*int32, bool)`

GetMessageIdentifierOk returns a tuple with the MessageIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIdentifier

`func (o *PwsInformation) SetMessageIdentifier(v int32)`

SetMessageIdentifier sets MessageIdentifier field to given value.


### GetSerialNumber

`func (o *PwsInformation) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PwsInformation) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PwsInformation) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.


### GetPwsContainer

`func (o *PwsInformation) GetPwsContainer() N2InfoContent`

GetPwsContainer returns the PwsContainer field if non-nil, zero value otherwise.

### GetPwsContainerOk

`func (o *PwsInformation) GetPwsContainerOk() (*N2InfoContent, bool)`

GetPwsContainerOk returns a tuple with the PwsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwsContainer

`func (o *PwsInformation) SetPwsContainer(v N2InfoContent)`

SetPwsContainer sets PwsContainer field to given value.


### GetBcEmptyAreaList

`func (o *PwsInformation) GetBcEmptyAreaList() []GlobalRanNodeId`

GetBcEmptyAreaList returns the BcEmptyAreaList field if non-nil, zero value otherwise.

### GetBcEmptyAreaListOk

`func (o *PwsInformation) GetBcEmptyAreaListOk() (*[]GlobalRanNodeId, bool)`

GetBcEmptyAreaListOk returns a tuple with the BcEmptyAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcEmptyAreaList

`func (o *PwsInformation) SetBcEmptyAreaList(v []GlobalRanNodeId)`

SetBcEmptyAreaList sets BcEmptyAreaList field to given value.

### HasBcEmptyAreaList

`func (o *PwsInformation) HasBcEmptyAreaList() bool`

HasBcEmptyAreaList returns a boolean if a field has been set.

### GetSendRanResponse

`func (o *PwsInformation) GetSendRanResponse() bool`

GetSendRanResponse returns the SendRanResponse field if non-nil, zero value otherwise.

### GetSendRanResponseOk

`func (o *PwsInformation) GetSendRanResponseOk() (*bool, bool)`

GetSendRanResponseOk returns a tuple with the SendRanResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRanResponse

`func (o *PwsInformation) SetSendRanResponse(v bool)`

SetSendRanResponse sets SendRanResponse field to given value.

### HasSendRanResponse

`func (o *PwsInformation) HasSendRanResponse() bool`

HasSendRanResponse returns a boolean if a field has been set.

### GetOmcId

`func (o *PwsInformation) GetOmcId() string`

GetOmcId returns the OmcId field if non-nil, zero value otherwise.

### GetOmcIdOk

`func (o *PwsInformation) GetOmcIdOk() (*string, bool)`

GetOmcIdOk returns a tuple with the OmcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmcId

`func (o *PwsInformation) SetOmcId(v string)`

SetOmcId sets OmcId field to given value.

### HasOmcId

`func (o *PwsInformation) HasOmcId() bool`

HasOmcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


