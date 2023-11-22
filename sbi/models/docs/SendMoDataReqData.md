# SendMoDataReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoData** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**MoExpDataCounter** | Pointer to [**MoExpDataCounter**](MoExpDataCounter.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 

## Methods

### NewSendMoDataReqData

`func NewSendMoDataReqData(moData RefToBinaryData, ) *SendMoDataReqData`

NewSendMoDataReqData instantiates a new SendMoDataReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMoDataReqDataWithDefaults

`func NewSendMoDataReqDataWithDefaults() *SendMoDataReqData`

NewSendMoDataReqDataWithDefaults instantiates a new SendMoDataReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoData

`func (o *SendMoDataReqData) GetMoData() RefToBinaryData`

GetMoData returns the MoData field if non-nil, zero value otherwise.

### GetMoDataOk

`func (o *SendMoDataReqData) GetMoDataOk() (*RefToBinaryData, bool)`

GetMoDataOk returns a tuple with the MoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoData

`func (o *SendMoDataReqData) SetMoData(v RefToBinaryData)`

SetMoData sets MoData field to given value.


### GetMoExpDataCounter

`func (o *SendMoDataReqData) GetMoExpDataCounter() MoExpDataCounter`

GetMoExpDataCounter returns the MoExpDataCounter field if non-nil, zero value otherwise.

### GetMoExpDataCounterOk

`func (o *SendMoDataReqData) GetMoExpDataCounterOk() (*MoExpDataCounter, bool)`

GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoExpDataCounter

`func (o *SendMoDataReqData) SetMoExpDataCounter(v MoExpDataCounter)`

SetMoExpDataCounter sets MoExpDataCounter field to given value.

### HasMoExpDataCounter

`func (o *SendMoDataReqData) HasMoExpDataCounter() bool`

HasMoExpDataCounter returns a boolean if a field has been set.

### GetUeLocation

`func (o *SendMoDataReqData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SendMoDataReqData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SendMoDataReqData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SendMoDataReqData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


