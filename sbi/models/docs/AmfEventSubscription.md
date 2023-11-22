# AmfEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventList** | [**[]AmfEvent**](AmfEvent.md) |  | 
**EventNotifyUri** | **string** |  | 
**NotifyCorrelationId** | **string** |  | 
**NfId** | **string** |  | 
**SubsChangeNotifyUri** | Pointer to **string** |  | [optional] 
**SubsChangeNotifyCorrelationId** | Pointer to **string** |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**AnyUE** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**AmfEventMode**](AmfEventMode.md) |  | [optional] 
**SourceNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 

## Methods

### NewAmfEventSubscription

`func NewAmfEventSubscription(eventList []AmfEvent, eventNotifyUri string, notifyCorrelationId string, nfId string, ) *AmfEventSubscription`

NewAmfEventSubscription instantiates a new AmfEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventSubscriptionWithDefaults

`func NewAmfEventSubscriptionWithDefaults() *AmfEventSubscription`

NewAmfEventSubscriptionWithDefaults instantiates a new AmfEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventList

`func (o *AmfEventSubscription) GetEventList() []AmfEvent`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *AmfEventSubscription) GetEventListOk() (*[]AmfEvent, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *AmfEventSubscription) SetEventList(v []AmfEvent)`

SetEventList sets EventList field to given value.


### GetEventNotifyUri

`func (o *AmfEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *AmfEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *AmfEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetNotifyCorrelationId

`func (o *AmfEventSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *AmfEventSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *AmfEventSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.


### GetNfId

`func (o *AmfEventSubscription) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *AmfEventSubscription) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *AmfEventSubscription) SetNfId(v string)`

SetNfId sets NfId field to given value.


### GetSubsChangeNotifyUri

`func (o *AmfEventSubscription) GetSubsChangeNotifyUri() string`

GetSubsChangeNotifyUri returns the SubsChangeNotifyUri field if non-nil, zero value otherwise.

### GetSubsChangeNotifyUriOk

`func (o *AmfEventSubscription) GetSubsChangeNotifyUriOk() (*string, bool)`

GetSubsChangeNotifyUriOk returns a tuple with the SubsChangeNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyUri

`func (o *AmfEventSubscription) SetSubsChangeNotifyUri(v string)`

SetSubsChangeNotifyUri sets SubsChangeNotifyUri field to given value.

### HasSubsChangeNotifyUri

`func (o *AmfEventSubscription) HasSubsChangeNotifyUri() bool`

HasSubsChangeNotifyUri returns a boolean if a field has been set.

### GetSubsChangeNotifyCorrelationId

`func (o *AmfEventSubscription) GetSubsChangeNotifyCorrelationId() string`

GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field if non-nil, zero value otherwise.

### GetSubsChangeNotifyCorrelationIdOk

`func (o *AmfEventSubscription) GetSubsChangeNotifyCorrelationIdOk() (*string, bool)`

GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyCorrelationId

`func (o *AmfEventSubscription) SetSubsChangeNotifyCorrelationId(v string)`

SetSubsChangeNotifyCorrelationId sets SubsChangeNotifyCorrelationId field to given value.

### HasSubsChangeNotifyCorrelationId

`func (o *AmfEventSubscription) HasSubsChangeNotifyCorrelationId() bool`

HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.

### GetSupi

`func (o *AmfEventSubscription) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AmfEventSubscription) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AmfEventSubscription) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AmfEventSubscription) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGroupId

`func (o *AmfEventSubscription) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AmfEventSubscription) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AmfEventSubscription) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AmfEventSubscription) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGpsi

`func (o *AmfEventSubscription) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AmfEventSubscription) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AmfEventSubscription) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AmfEventSubscription) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPei

`func (o *AmfEventSubscription) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AmfEventSubscription) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AmfEventSubscription) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AmfEventSubscription) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetAnyUE

`func (o *AmfEventSubscription) GetAnyUE() bool`

GetAnyUE returns the AnyUE field if non-nil, zero value otherwise.

### GetAnyUEOk

`func (o *AmfEventSubscription) GetAnyUEOk() (*bool, bool)`

GetAnyUEOk returns a tuple with the AnyUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUE

`func (o *AmfEventSubscription) SetAnyUE(v bool)`

SetAnyUE sets AnyUE field to given value.

### HasAnyUE

`func (o *AmfEventSubscription) HasAnyUE() bool`

HasAnyUE returns a boolean if a field has been set.

### GetOptions

`func (o *AmfEventSubscription) GetOptions() AmfEventMode`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AmfEventSubscription) GetOptionsOk() (*AmfEventMode, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AmfEventSubscription) SetOptions(v AmfEventMode)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AmfEventSubscription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSourceNfType

`func (o *AmfEventSubscription) GetSourceNfType() NFType`

GetSourceNfType returns the SourceNfType field if non-nil, zero value otherwise.

### GetSourceNfTypeOk

`func (o *AmfEventSubscription) GetSourceNfTypeOk() (*NFType, bool)`

GetSourceNfTypeOk returns a tuple with the SourceNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNfType

`func (o *AmfEventSubscription) SetSourceNfType(v NFType)`

SetSourceNfType sets SourceNfType field to given value.

### HasSourceNfType

`func (o *AmfEventSubscription) HasSourceNfType() bool`

HasSourceNfType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


