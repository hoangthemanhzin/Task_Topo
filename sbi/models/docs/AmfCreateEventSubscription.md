# AmfCreateEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**AmfEventSubscription**](AmfEventSubscription.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 

## Methods

### NewAmfCreateEventSubscription

`func NewAmfCreateEventSubscription(subscription AmfEventSubscription, ) *AmfCreateEventSubscription`

NewAmfCreateEventSubscription instantiates a new AmfCreateEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfCreateEventSubscriptionWithDefaults

`func NewAmfCreateEventSubscriptionWithDefaults() *AmfCreateEventSubscription`

NewAmfCreateEventSubscriptionWithDefaults instantiates a new AmfCreateEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *AmfCreateEventSubscription) GetSubscription() AmfEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AmfCreateEventSubscription) GetSubscriptionOk() (*AmfEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AmfCreateEventSubscription) SetSubscription(v AmfEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSupportedFeatures

`func (o *AmfCreateEventSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AmfCreateEventSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AmfCreateEventSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AmfCreateEventSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *AmfCreateEventSubscription) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *AmfCreateEventSubscription) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *AmfCreateEventSubscription) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *AmfCreateEventSubscription) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


