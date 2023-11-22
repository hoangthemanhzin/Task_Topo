# AmfEventSubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** |  | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**RefIdList** | **[]int32** |  | 
**OldSubId** | Pointer to **string** |  | [optional] 

## Methods

### NewAmfEventSubscriptionInfo

`func NewAmfEventSubscriptionInfo(subId string, refIdList []int32, ) *AmfEventSubscriptionInfo`

NewAmfEventSubscriptionInfo instantiates a new AmfEventSubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventSubscriptionInfoWithDefaults

`func NewAmfEventSubscriptionInfoWithDefaults() *AmfEventSubscriptionInfo`

NewAmfEventSubscriptionInfoWithDefaults instantiates a new AmfEventSubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *AmfEventSubscriptionInfo) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *AmfEventSubscriptionInfo) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *AmfEventSubscriptionInfo) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetNotifyCorrelationId

`func (o *AmfEventSubscriptionInfo) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *AmfEventSubscriptionInfo) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *AmfEventSubscriptionInfo) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *AmfEventSubscriptionInfo) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetRefIdList

`func (o *AmfEventSubscriptionInfo) GetRefIdList() []int32`

GetRefIdList returns the RefIdList field if non-nil, zero value otherwise.

### GetRefIdListOk

`func (o *AmfEventSubscriptionInfo) GetRefIdListOk() (*[]int32, bool)`

GetRefIdListOk returns a tuple with the RefIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefIdList

`func (o *AmfEventSubscriptionInfo) SetRefIdList(v []int32)`

SetRefIdList sets RefIdList field to given value.


### GetOldSubId

`func (o *AmfEventSubscriptionInfo) GetOldSubId() string`

GetOldSubId returns the OldSubId field if non-nil, zero value otherwise.

### GetOldSubIdOk

`func (o *AmfEventSubscriptionInfo) GetOldSubIdOk() (*string, bool)`

GetOldSubIdOk returns a tuple with the OldSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSubId

`func (o *AmfEventSubscriptionInfo) SetOldSubId(v string)`

SetOldSubId sets OldSubId field to given value.

### HasOldSubId

`func (o *AmfEventSubscriptionInfo) HasOldSubId() bool`

HasOldSubId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


