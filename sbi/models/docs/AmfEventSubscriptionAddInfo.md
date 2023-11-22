# AmfEventSubscriptionAddInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingInfo** | Pointer to **[]string** |  | [optional] 
**SubscribingNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**EventSyncInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewAmfEventSubscriptionAddInfo

`func NewAmfEventSubscriptionAddInfo() *AmfEventSubscriptionAddInfo`

NewAmfEventSubscriptionAddInfo instantiates a new AmfEventSubscriptionAddInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventSubscriptionAddInfoWithDefaults

`func NewAmfEventSubscriptionAddInfoWithDefaults() *AmfEventSubscriptionAddInfo`

NewAmfEventSubscriptionAddInfoWithDefaults instantiates a new AmfEventSubscriptionAddInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingInfo

`func (o *AmfEventSubscriptionAddInfo) GetBindingInfo() []string`

GetBindingInfo returns the BindingInfo field if non-nil, zero value otherwise.

### GetBindingInfoOk

`func (o *AmfEventSubscriptionAddInfo) GetBindingInfoOk() (*[]string, bool)`

GetBindingInfoOk returns a tuple with the BindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingInfo

`func (o *AmfEventSubscriptionAddInfo) SetBindingInfo(v []string)`

SetBindingInfo sets BindingInfo field to given value.

### HasBindingInfo

`func (o *AmfEventSubscriptionAddInfo) HasBindingInfo() bool`

HasBindingInfo returns a boolean if a field has been set.

### GetSubscribingNfType

`func (o *AmfEventSubscriptionAddInfo) GetSubscribingNfType() NFType`

GetSubscribingNfType returns the SubscribingNfType field if non-nil, zero value otherwise.

### GetSubscribingNfTypeOk

`func (o *AmfEventSubscriptionAddInfo) GetSubscribingNfTypeOk() (*NFType, bool)`

GetSubscribingNfTypeOk returns a tuple with the SubscribingNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingNfType

`func (o *AmfEventSubscriptionAddInfo) SetSubscribingNfType(v NFType)`

SetSubscribingNfType sets SubscribingNfType field to given value.

### HasSubscribingNfType

`func (o *AmfEventSubscriptionAddInfo) HasSubscribingNfType() bool`

HasSubscribingNfType returns a boolean if a field has been set.

### GetEventSyncInd

`func (o *AmfEventSubscriptionAddInfo) GetEventSyncInd() bool`

GetEventSyncInd returns the EventSyncInd field if non-nil, zero value otherwise.

### GetEventSyncIndOk

`func (o *AmfEventSubscriptionAddInfo) GetEventSyncIndOk() (*bool, bool)`

GetEventSyncIndOk returns a tuple with the EventSyncInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSyncInd

`func (o *AmfEventSubscriptionAddInfo) SetEventSyncInd(v bool)`

SetEventSyncInd sets EventSyncInd field to given value.

### HasEventSyncInd

`func (o *AmfEventSubscriptionAddInfo) HasEventSyncInd() bool`

HasEventSyncInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


