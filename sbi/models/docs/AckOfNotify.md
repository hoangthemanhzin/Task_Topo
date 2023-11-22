# AckOfNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**AckResult** | [**AfResultInfo**](AfResultInfo.md) |  | 
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 

## Methods

### NewAckOfNotify

`func NewAckOfNotify(notifId string, ackResult AfResultInfo, ) *AckOfNotify`

NewAckOfNotify instantiates a new AckOfNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAckOfNotifyWithDefaults

`func NewAckOfNotifyWithDefaults() *AckOfNotify`

NewAckOfNotifyWithDefaults instantiates a new AckOfNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *AckOfNotify) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *AckOfNotify) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *AckOfNotify) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetAckResult

`func (o *AckOfNotify) GetAckResult() AfResultInfo`

GetAckResult returns the AckResult field if non-nil, zero value otherwise.

### GetAckResultOk

`func (o *AckOfNotify) GetAckResultOk() (*AfResultInfo, bool)`

GetAckResultOk returns a tuple with the AckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckResult

`func (o *AckOfNotify) SetAckResult(v AfResultInfo)`

SetAckResult sets AckResult field to given value.


### GetSupi

`func (o *AckOfNotify) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AckOfNotify) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AckOfNotify) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AckOfNotify) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *AckOfNotify) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AckOfNotify) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AckOfNotify) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AckOfNotify) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


