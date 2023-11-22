# N2InformationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N2NotifySubscriptionId** | **string** |  | 
**N2InfoContainer** | Pointer to [**N2InfoContainer**](N2InfoContainer.md) |  | [optional] 
**ToReleaseSessionList** | Pointer to **[]int32** |  | [optional] 
**LcsCorrelationId** | Pointer to **string** |  | [optional] 
**NotifyReason** | Pointer to [**N2InfoNotifyReason**](N2InfoNotifyReason.md) |  | [optional] 
**SmfChangeInfoList** | Pointer to [**[]SmfChangeInfo**](SmfChangeInfo.md) |  | [optional] 
**RanNodeId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**InitialAmfName** | Pointer to **string** |  | [optional] 
**AnN2IPv4Addr** | Pointer to **string** |  | [optional] 
**AnN2IPv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**NotifySourceNgRan** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewN2InformationNotification

`func NewN2InformationNotification(n2NotifySubscriptionId string, ) *N2InformationNotification`

NewN2InformationNotification instantiates a new N2InformationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InformationNotificationWithDefaults

`func NewN2InformationNotificationWithDefaults() *N2InformationNotification`

NewN2InformationNotificationWithDefaults instantiates a new N2InformationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN2NotifySubscriptionId

`func (o *N2InformationNotification) GetN2NotifySubscriptionId() string`

GetN2NotifySubscriptionId returns the N2NotifySubscriptionId field if non-nil, zero value otherwise.

### GetN2NotifySubscriptionIdOk

`func (o *N2InformationNotification) GetN2NotifySubscriptionIdOk() (*string, bool)`

GetN2NotifySubscriptionIdOk returns a tuple with the N2NotifySubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2NotifySubscriptionId

`func (o *N2InformationNotification) SetN2NotifySubscriptionId(v string)`

SetN2NotifySubscriptionId sets N2NotifySubscriptionId field to given value.


### GetN2InfoContainer

`func (o *N2InformationNotification) GetN2InfoContainer() N2InfoContainer`

GetN2InfoContainer returns the N2InfoContainer field if non-nil, zero value otherwise.

### GetN2InfoContainerOk

`func (o *N2InformationNotification) GetN2InfoContainerOk() (*N2InfoContainer, bool)`

GetN2InfoContainerOk returns a tuple with the N2InfoContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InfoContainer

`func (o *N2InformationNotification) SetN2InfoContainer(v N2InfoContainer)`

SetN2InfoContainer sets N2InfoContainer field to given value.

### HasN2InfoContainer

`func (o *N2InformationNotification) HasN2InfoContainer() bool`

HasN2InfoContainer returns a boolean if a field has been set.

### GetToReleaseSessionList

`func (o *N2InformationNotification) GetToReleaseSessionList() []int32`

GetToReleaseSessionList returns the ToReleaseSessionList field if non-nil, zero value otherwise.

### GetToReleaseSessionListOk

`func (o *N2InformationNotification) GetToReleaseSessionListOk() (*[]int32, bool)`

GetToReleaseSessionListOk returns a tuple with the ToReleaseSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToReleaseSessionList

`func (o *N2InformationNotification) SetToReleaseSessionList(v []int32)`

SetToReleaseSessionList sets ToReleaseSessionList field to given value.

### HasToReleaseSessionList

`func (o *N2InformationNotification) HasToReleaseSessionList() bool`

HasToReleaseSessionList returns a boolean if a field has been set.

### GetLcsCorrelationId

`func (o *N2InformationNotification) GetLcsCorrelationId() string`

GetLcsCorrelationId returns the LcsCorrelationId field if non-nil, zero value otherwise.

### GetLcsCorrelationIdOk

`func (o *N2InformationNotification) GetLcsCorrelationIdOk() (*string, bool)`

GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsCorrelationId

`func (o *N2InformationNotification) SetLcsCorrelationId(v string)`

SetLcsCorrelationId sets LcsCorrelationId field to given value.

### HasLcsCorrelationId

`func (o *N2InformationNotification) HasLcsCorrelationId() bool`

HasLcsCorrelationId returns a boolean if a field has been set.

### GetNotifyReason

`func (o *N2InformationNotification) GetNotifyReason() N2InfoNotifyReason`

GetNotifyReason returns the NotifyReason field if non-nil, zero value otherwise.

### GetNotifyReasonOk

`func (o *N2InformationNotification) GetNotifyReasonOk() (*N2InfoNotifyReason, bool)`

GetNotifyReasonOk returns a tuple with the NotifyReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReason

`func (o *N2InformationNotification) SetNotifyReason(v N2InfoNotifyReason)`

SetNotifyReason sets NotifyReason field to given value.

### HasNotifyReason

`func (o *N2InformationNotification) HasNotifyReason() bool`

HasNotifyReason returns a boolean if a field has been set.

### GetSmfChangeInfoList

`func (o *N2InformationNotification) GetSmfChangeInfoList() []SmfChangeInfo`

GetSmfChangeInfoList returns the SmfChangeInfoList field if non-nil, zero value otherwise.

### GetSmfChangeInfoListOk

`func (o *N2InformationNotification) GetSmfChangeInfoListOk() (*[]SmfChangeInfo, bool)`

GetSmfChangeInfoListOk returns a tuple with the SmfChangeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfChangeInfoList

`func (o *N2InformationNotification) SetSmfChangeInfoList(v []SmfChangeInfo)`

SetSmfChangeInfoList sets SmfChangeInfoList field to given value.

### HasSmfChangeInfoList

`func (o *N2InformationNotification) HasSmfChangeInfoList() bool`

HasSmfChangeInfoList returns a boolean if a field has been set.

### GetRanNodeId

`func (o *N2InformationNotification) GetRanNodeId() GlobalRanNodeId`

GetRanNodeId returns the RanNodeId field if non-nil, zero value otherwise.

### GetRanNodeIdOk

`func (o *N2InformationNotification) GetRanNodeIdOk() (*GlobalRanNodeId, bool)`

GetRanNodeIdOk returns a tuple with the RanNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNodeId

`func (o *N2InformationNotification) SetRanNodeId(v GlobalRanNodeId)`

SetRanNodeId sets RanNodeId field to given value.

### HasRanNodeId

`func (o *N2InformationNotification) HasRanNodeId() bool`

HasRanNodeId returns a boolean if a field has been set.

### GetInitialAmfName

`func (o *N2InformationNotification) GetInitialAmfName() string`

GetInitialAmfName returns the InitialAmfName field if non-nil, zero value otherwise.

### GetInitialAmfNameOk

`func (o *N2InformationNotification) GetInitialAmfNameOk() (*string, bool)`

GetInitialAmfNameOk returns a tuple with the InitialAmfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmfName

`func (o *N2InformationNotification) SetInitialAmfName(v string)`

SetInitialAmfName sets InitialAmfName field to given value.

### HasInitialAmfName

`func (o *N2InformationNotification) HasInitialAmfName() bool`

HasInitialAmfName returns a boolean if a field has been set.

### GetAnN2IPv4Addr

`func (o *N2InformationNotification) GetAnN2IPv4Addr() string`

GetAnN2IPv4Addr returns the AnN2IPv4Addr field if non-nil, zero value otherwise.

### GetAnN2IPv4AddrOk

`func (o *N2InformationNotification) GetAnN2IPv4AddrOk() (*string, bool)`

GetAnN2IPv4AddrOk returns a tuple with the AnN2IPv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2IPv4Addr

`func (o *N2InformationNotification) SetAnN2IPv4Addr(v string)`

SetAnN2IPv4Addr sets AnN2IPv4Addr field to given value.

### HasAnN2IPv4Addr

`func (o *N2InformationNotification) HasAnN2IPv4Addr() bool`

HasAnN2IPv4Addr returns a boolean if a field has been set.

### GetAnN2IPv6Addr

`func (o *N2InformationNotification) GetAnN2IPv6Addr() Ipv6Addr`

GetAnN2IPv6Addr returns the AnN2IPv6Addr field if non-nil, zero value otherwise.

### GetAnN2IPv6AddrOk

`func (o *N2InformationNotification) GetAnN2IPv6AddrOk() (*Ipv6Addr, bool)`

GetAnN2IPv6AddrOk returns a tuple with the AnN2IPv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2IPv6Addr

`func (o *N2InformationNotification) SetAnN2IPv6Addr(v Ipv6Addr)`

SetAnN2IPv6Addr sets AnN2IPv6Addr field to given value.

### HasAnN2IPv6Addr

`func (o *N2InformationNotification) HasAnN2IPv6Addr() bool`

HasAnN2IPv6Addr returns a boolean if a field has been set.

### GetGuami

`func (o *N2InformationNotification) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *N2InformationNotification) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *N2InformationNotification) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *N2InformationNotification) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetNotifySourceNgRan

`func (o *N2InformationNotification) GetNotifySourceNgRan() bool`

GetNotifySourceNgRan returns the NotifySourceNgRan field if non-nil, zero value otherwise.

### GetNotifySourceNgRanOk

`func (o *N2InformationNotification) GetNotifySourceNgRanOk() (*bool, bool)`

GetNotifySourceNgRanOk returns a tuple with the NotifySourceNgRan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySourceNgRan

`func (o *N2InformationNotification) SetNotifySourceNgRan(v bool)`

SetNotifySourceNgRan sets NotifySourceNgRan field to given value.

### HasNotifySourceNgRan

`func (o *N2InformationNotification) HasNotifySourceNgRan() bool`

HasNotifySourceNgRan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


