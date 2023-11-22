# AppSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AscReqData** | Pointer to [**AppSessionContextReqData**](AppSessionContextReqData.md) |  | [optional] 
**AscRespData** | Pointer to [**AppSessionContextRespData**](AppSessionContextRespData.md) |  | [optional] 
**EvsNotif** | Pointer to [**EventsNotification**](EventsNotification.md) |  | [optional] 

## Methods

### NewAppSessionContext

`func NewAppSessionContext() *AppSessionContext`

NewAppSessionContext instantiates a new AppSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSessionContextWithDefaults

`func NewAppSessionContextWithDefaults() *AppSessionContext`

NewAppSessionContextWithDefaults instantiates a new AppSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAscReqData

`func (o *AppSessionContext) GetAscReqData() AppSessionContextReqData`

GetAscReqData returns the AscReqData field if non-nil, zero value otherwise.

### GetAscReqDataOk

`func (o *AppSessionContext) GetAscReqDataOk() (*AppSessionContextReqData, bool)`

GetAscReqDataOk returns a tuple with the AscReqData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAscReqData

`func (o *AppSessionContext) SetAscReqData(v AppSessionContextReqData)`

SetAscReqData sets AscReqData field to given value.

### HasAscReqData

`func (o *AppSessionContext) HasAscReqData() bool`

HasAscReqData returns a boolean if a field has been set.

### GetAscRespData

`func (o *AppSessionContext) GetAscRespData() AppSessionContextRespData`

GetAscRespData returns the AscRespData field if non-nil, zero value otherwise.

### GetAscRespDataOk

`func (o *AppSessionContext) GetAscRespDataOk() (*AppSessionContextRespData, bool)`

GetAscRespDataOk returns a tuple with the AscRespData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAscRespData

`func (o *AppSessionContext) SetAscRespData(v AppSessionContextRespData)`

SetAscRespData sets AscRespData field to given value.

### HasAscRespData

`func (o *AppSessionContext) HasAscRespData() bool`

HasAscRespData returns a boolean if a field has been set.

### GetEvsNotif

`func (o *AppSessionContext) GetEvsNotif() EventsNotification`

GetEvsNotif returns the EvsNotif field if non-nil, zero value otherwise.

### GetEvsNotifOk

`func (o *AppSessionContext) GetEvsNotifOk() (*EventsNotification, bool)`

GetEvsNotifOk returns a tuple with the EvsNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvsNotif

`func (o *AppSessionContext) SetEvsNotif(v EventsNotification)`

SetEvsNotif sets EvsNotif field to given value.

### HasEvsNotif

`func (o *AppSessionContext) HasEvsNotif() bool`

HasEvsNotif returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


