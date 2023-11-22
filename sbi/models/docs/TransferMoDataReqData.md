# TransferMoDataReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoData** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**MoExpDataCounter** | Pointer to [**MoExpDataCounter**](MoExpDataCounter.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 

## Methods

### NewTransferMoDataReqData

`func NewTransferMoDataReqData(moData RefToBinaryData, ) *TransferMoDataReqData`

NewTransferMoDataReqData instantiates a new TransferMoDataReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferMoDataReqDataWithDefaults

`func NewTransferMoDataReqDataWithDefaults() *TransferMoDataReqData`

NewTransferMoDataReqDataWithDefaults instantiates a new TransferMoDataReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoData

`func (o *TransferMoDataReqData) GetMoData() RefToBinaryData`

GetMoData returns the MoData field if non-nil, zero value otherwise.

### GetMoDataOk

`func (o *TransferMoDataReqData) GetMoDataOk() (*RefToBinaryData, bool)`

GetMoDataOk returns a tuple with the MoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoData

`func (o *TransferMoDataReqData) SetMoData(v RefToBinaryData)`

SetMoData sets MoData field to given value.


### GetMoExpDataCounter

`func (o *TransferMoDataReqData) GetMoExpDataCounter() MoExpDataCounter`

GetMoExpDataCounter returns the MoExpDataCounter field if non-nil, zero value otherwise.

### GetMoExpDataCounterOk

`func (o *TransferMoDataReqData) GetMoExpDataCounterOk() (*MoExpDataCounter, bool)`

GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoExpDataCounter

`func (o *TransferMoDataReqData) SetMoExpDataCounter(v MoExpDataCounter)`

SetMoExpDataCounter sets MoExpDataCounter field to given value.

### HasMoExpDataCounter

`func (o *TransferMoDataReqData) HasMoExpDataCounter() bool`

HasMoExpDataCounter returns a boolean if a field has been set.

### GetUeLocation

`func (o *TransferMoDataReqData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *TransferMoDataReqData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *TransferMoDataReqData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *TransferMoDataReqData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


