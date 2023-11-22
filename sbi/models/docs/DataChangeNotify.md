# DataChangeNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalCallbackReference** | Pointer to **[]string** |  | [optional] 
**UeId** | Pointer to **string** |  | [optional] 
**NotifyItems** | Pointer to [**[]NotifyItem**](NotifyItem.md) |  | [optional] 
**SdmSubscription** | Pointer to [**SdmSubscription1**](SdmSubscription1.md) |  | [optional] 
**AdditionalSdmSubscriptions** | Pointer to [**[]SdmSubscription1**](SdmSubscription1.md) |  | [optional] 
**SubscriptionDataSubscriptions** | Pointer to [**[]SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md) |  | [optional] 

## Methods

### NewDataChangeNotify

`func NewDataChangeNotify() *DataChangeNotify`

NewDataChangeNotify instantiates a new DataChangeNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataChangeNotifyWithDefaults

`func NewDataChangeNotifyWithDefaults() *DataChangeNotify`

NewDataChangeNotifyWithDefaults instantiates a new DataChangeNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalCallbackReference

`func (o *DataChangeNotify) GetOriginalCallbackReference() []string`

GetOriginalCallbackReference returns the OriginalCallbackReference field if non-nil, zero value otherwise.

### GetOriginalCallbackReferenceOk

`func (o *DataChangeNotify) GetOriginalCallbackReferenceOk() (*[]string, bool)`

GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCallbackReference

`func (o *DataChangeNotify) SetOriginalCallbackReference(v []string)`

SetOriginalCallbackReference sets OriginalCallbackReference field to given value.

### HasOriginalCallbackReference

`func (o *DataChangeNotify) HasOriginalCallbackReference() bool`

HasOriginalCallbackReference returns a boolean if a field has been set.

### GetUeId

`func (o *DataChangeNotify) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *DataChangeNotify) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *DataChangeNotify) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *DataChangeNotify) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetNotifyItems

`func (o *DataChangeNotify) GetNotifyItems() []NotifyItem`

GetNotifyItems returns the NotifyItems field if non-nil, zero value otherwise.

### GetNotifyItemsOk

`func (o *DataChangeNotify) GetNotifyItemsOk() (*[]NotifyItem, bool)`

GetNotifyItemsOk returns a tuple with the NotifyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyItems

`func (o *DataChangeNotify) SetNotifyItems(v []NotifyItem)`

SetNotifyItems sets NotifyItems field to given value.

### HasNotifyItems

`func (o *DataChangeNotify) HasNotifyItems() bool`

HasNotifyItems returns a boolean if a field has been set.

### GetSdmSubscription

`func (o *DataChangeNotify) GetSdmSubscription() SdmSubscription1`

GetSdmSubscription returns the SdmSubscription field if non-nil, zero value otherwise.

### GetSdmSubscriptionOk

`func (o *DataChangeNotify) GetSdmSubscriptionOk() (*SdmSubscription1, bool)`

GetSdmSubscriptionOk returns a tuple with the SdmSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmSubscription

`func (o *DataChangeNotify) SetSdmSubscription(v SdmSubscription1)`

SetSdmSubscription sets SdmSubscription field to given value.

### HasSdmSubscription

`func (o *DataChangeNotify) HasSdmSubscription() bool`

HasSdmSubscription returns a boolean if a field has been set.

### GetAdditionalSdmSubscriptions

`func (o *DataChangeNotify) GetAdditionalSdmSubscriptions() []SdmSubscription1`

GetAdditionalSdmSubscriptions returns the AdditionalSdmSubscriptions field if non-nil, zero value otherwise.

### GetAdditionalSdmSubscriptionsOk

`func (o *DataChangeNotify) GetAdditionalSdmSubscriptionsOk() (*[]SdmSubscription1, bool)`

GetAdditionalSdmSubscriptionsOk returns a tuple with the AdditionalSdmSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSdmSubscriptions

`func (o *DataChangeNotify) SetAdditionalSdmSubscriptions(v []SdmSubscription1)`

SetAdditionalSdmSubscriptions sets AdditionalSdmSubscriptions field to given value.

### HasAdditionalSdmSubscriptions

`func (o *DataChangeNotify) HasAdditionalSdmSubscriptions() bool`

HasAdditionalSdmSubscriptions returns a boolean if a field has been set.

### GetSubscriptionDataSubscriptions

`func (o *DataChangeNotify) GetSubscriptionDataSubscriptions() []SubscriptionDataSubscriptions`

GetSubscriptionDataSubscriptions returns the SubscriptionDataSubscriptions field if non-nil, zero value otherwise.

### GetSubscriptionDataSubscriptionsOk

`func (o *DataChangeNotify) GetSubscriptionDataSubscriptionsOk() (*[]SubscriptionDataSubscriptions, bool)`

GetSubscriptionDataSubscriptionsOk returns a tuple with the SubscriptionDataSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDataSubscriptions

`func (o *DataChangeNotify) SetSubscriptionDataSubscriptions(v []SubscriptionDataSubscriptions)`

SetSubscriptionDataSubscriptions sets SubscriptionDataSubscriptions field to given value.

### HasSubscriptionDataSubscriptions

`func (o *DataChangeNotify) HasSubscriptionDataSubscriptions() bool`

HasSubscriptionDataSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


