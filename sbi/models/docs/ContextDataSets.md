# ContextDataSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amf3Gpp** | Pointer to [**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md) |  | [optional] 
**AmfNon3Gpp** | Pointer to [**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md) |  | [optional] 
**SdmSubscriptions** | Pointer to [**[]SdmSubscription**](SdmSubscription.md) |  | [optional] 
**EeSubscriptions** | Pointer to [**[]EeSubscription**](EeSubscription.md) |  | [optional] 
**Smsf3GppAccess** | Pointer to [**SmsfRegistration**](SmsfRegistration.md) |  | [optional] 
**SmsfNon3GppAccess** | Pointer to [**SmsfRegistration**](SmsfRegistration.md) |  | [optional] 
**SubscriptionDataSubscriptions** | Pointer to [**[]SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md) |  | [optional] 
**SmfRegistrations** | Pointer to [**[]SmfRegistration**](SmfRegistration.md) |  | [optional] 
**IpSmGw** | Pointer to [**IpSmGwRegistration**](IpSmGwRegistration.md) |  | [optional] 

## Methods

### NewContextDataSets

`func NewContextDataSets() *ContextDataSets`

NewContextDataSets instantiates a new ContextDataSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextDataSetsWithDefaults

`func NewContextDataSetsWithDefaults() *ContextDataSets`

NewContextDataSetsWithDefaults instantiates a new ContextDataSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmf3Gpp

`func (o *ContextDataSets) GetAmf3Gpp() Amf3GppAccessRegistration`

GetAmf3Gpp returns the Amf3Gpp field if non-nil, zero value otherwise.

### GetAmf3GppOk

`func (o *ContextDataSets) GetAmf3GppOk() (*Amf3GppAccessRegistration, bool)`

GetAmf3GppOk returns a tuple with the Amf3Gpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmf3Gpp

`func (o *ContextDataSets) SetAmf3Gpp(v Amf3GppAccessRegistration)`

SetAmf3Gpp sets Amf3Gpp field to given value.

### HasAmf3Gpp

`func (o *ContextDataSets) HasAmf3Gpp() bool`

HasAmf3Gpp returns a boolean if a field has been set.

### GetAmfNon3Gpp

`func (o *ContextDataSets) GetAmfNon3Gpp() AmfNon3GppAccessRegistration`

GetAmfNon3Gpp returns the AmfNon3Gpp field if non-nil, zero value otherwise.

### GetAmfNon3GppOk

`func (o *ContextDataSets) GetAmfNon3GppOk() (*AmfNon3GppAccessRegistration, bool)`

GetAmfNon3GppOk returns a tuple with the AmfNon3Gpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfNon3Gpp

`func (o *ContextDataSets) SetAmfNon3Gpp(v AmfNon3GppAccessRegistration)`

SetAmfNon3Gpp sets AmfNon3Gpp field to given value.

### HasAmfNon3Gpp

`func (o *ContextDataSets) HasAmfNon3Gpp() bool`

HasAmfNon3Gpp returns a boolean if a field has been set.

### GetSdmSubscriptions

`func (o *ContextDataSets) GetSdmSubscriptions() []SdmSubscription`

GetSdmSubscriptions returns the SdmSubscriptions field if non-nil, zero value otherwise.

### GetSdmSubscriptionsOk

`func (o *ContextDataSets) GetSdmSubscriptionsOk() (*[]SdmSubscription, bool)`

GetSdmSubscriptionsOk returns a tuple with the SdmSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmSubscriptions

`func (o *ContextDataSets) SetSdmSubscriptions(v []SdmSubscription)`

SetSdmSubscriptions sets SdmSubscriptions field to given value.

### HasSdmSubscriptions

`func (o *ContextDataSets) HasSdmSubscriptions() bool`

HasSdmSubscriptions returns a boolean if a field has been set.

### GetEeSubscriptions

`func (o *ContextDataSets) GetEeSubscriptions() []EeSubscription`

GetEeSubscriptions returns the EeSubscriptions field if non-nil, zero value otherwise.

### GetEeSubscriptionsOk

`func (o *ContextDataSets) GetEeSubscriptionsOk() (*[]EeSubscription, bool)`

GetEeSubscriptionsOk returns a tuple with the EeSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeSubscriptions

`func (o *ContextDataSets) SetEeSubscriptions(v []EeSubscription)`

SetEeSubscriptions sets EeSubscriptions field to given value.

### HasEeSubscriptions

`func (o *ContextDataSets) HasEeSubscriptions() bool`

HasEeSubscriptions returns a boolean if a field has been set.

### GetSmsf3GppAccess

`func (o *ContextDataSets) GetSmsf3GppAccess() SmsfRegistration`

GetSmsf3GppAccess returns the Smsf3GppAccess field if non-nil, zero value otherwise.

### GetSmsf3GppAccessOk

`func (o *ContextDataSets) GetSmsf3GppAccessOk() (*SmsfRegistration, bool)`

GetSmsf3GppAccessOk returns a tuple with the Smsf3GppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsf3GppAccess

`func (o *ContextDataSets) SetSmsf3GppAccess(v SmsfRegistration)`

SetSmsf3GppAccess sets Smsf3GppAccess field to given value.

### HasSmsf3GppAccess

`func (o *ContextDataSets) HasSmsf3GppAccess() bool`

HasSmsf3GppAccess returns a boolean if a field has been set.

### GetSmsfNon3GppAccess

`func (o *ContextDataSets) GetSmsfNon3GppAccess() SmsfRegistration`

GetSmsfNon3GppAccess returns the SmsfNon3GppAccess field if non-nil, zero value otherwise.

### GetSmsfNon3GppAccessOk

`func (o *ContextDataSets) GetSmsfNon3GppAccessOk() (*SmsfRegistration, bool)`

GetSmsfNon3GppAccessOk returns a tuple with the SmsfNon3GppAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfNon3GppAccess

`func (o *ContextDataSets) SetSmsfNon3GppAccess(v SmsfRegistration)`

SetSmsfNon3GppAccess sets SmsfNon3GppAccess field to given value.

### HasSmsfNon3GppAccess

`func (o *ContextDataSets) HasSmsfNon3GppAccess() bool`

HasSmsfNon3GppAccess returns a boolean if a field has been set.

### GetSubscriptionDataSubscriptions

`func (o *ContextDataSets) GetSubscriptionDataSubscriptions() []SubscriptionDataSubscriptions`

GetSubscriptionDataSubscriptions returns the SubscriptionDataSubscriptions field if non-nil, zero value otherwise.

### GetSubscriptionDataSubscriptionsOk

`func (o *ContextDataSets) GetSubscriptionDataSubscriptionsOk() (*[]SubscriptionDataSubscriptions, bool)`

GetSubscriptionDataSubscriptionsOk returns a tuple with the SubscriptionDataSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDataSubscriptions

`func (o *ContextDataSets) SetSubscriptionDataSubscriptions(v []SubscriptionDataSubscriptions)`

SetSubscriptionDataSubscriptions sets SubscriptionDataSubscriptions field to given value.

### HasSubscriptionDataSubscriptions

`func (o *ContextDataSets) HasSubscriptionDataSubscriptions() bool`

HasSubscriptionDataSubscriptions returns a boolean if a field has been set.

### GetSmfRegistrations

`func (o *ContextDataSets) GetSmfRegistrations() []SmfRegistration`

GetSmfRegistrations returns the SmfRegistrations field if non-nil, zero value otherwise.

### GetSmfRegistrationsOk

`func (o *ContextDataSets) GetSmfRegistrationsOk() (*[]SmfRegistration, bool)`

GetSmfRegistrationsOk returns a tuple with the SmfRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfRegistrations

`func (o *ContextDataSets) SetSmfRegistrations(v []SmfRegistration)`

SetSmfRegistrations sets SmfRegistrations field to given value.

### HasSmfRegistrations

`func (o *ContextDataSets) HasSmfRegistrations() bool`

HasSmfRegistrations returns a boolean if a field has been set.

### GetIpSmGw

`func (o *ContextDataSets) GetIpSmGw() IpSmGwRegistration`

GetIpSmGw returns the IpSmGw field if non-nil, zero value otherwise.

### GetIpSmGwOk

`func (o *ContextDataSets) GetIpSmGwOk() (*IpSmGwRegistration, bool)`

GetIpSmGwOk returns a tuple with the IpSmGw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGw

`func (o *ContextDataSets) SetIpSmGw(v IpSmGwRegistration)`

SetIpSmGw sets IpSmGw field to given value.

### HasIpSmGw

`func (o *ContextDataSets) HasIpSmGw() bool`

HasIpSmGw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


