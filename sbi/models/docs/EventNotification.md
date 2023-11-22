# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**SmfEvent**](SmfEvent.md) |  | 
**TimeStamp** | **time.Time** |  | 
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**SourceDnai** | Pointer to **string** |  | [optional] 
**TargetDnai** | Pointer to **string** |  | [optional] 
**DnaiChgType** | Pointer to [**DnaiChangeType**](DnaiChangeType.md) |  | [optional] 
**SourceUeIpv4Addr** | Pointer to **string** |  | [optional] 
**SourceUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**TargetUeIpv4Addr** | Pointer to **string** |  | [optional] 
**TargetUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**SourceTraRouting** | Pointer to [**RouteToLocation**](RouteToLocation.md) |  | [optional] 
**TargetTraRouting** | Pointer to [**RouteToLocation**](RouteToLocation.md) |  | [optional] 
**UeMac** | Pointer to **string** |  | [optional] 
**AdIpv4Addr** | Pointer to **string** |  | [optional] 
**AdIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**ReIpv4Addr** | Pointer to **string** |  | [optional] 
**ReIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AccType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**PduSeId** | Pointer to **int32** |  | [optional] 
**DddStatus** | Pointer to [**DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 
**DddTraDescriptor** | Pointer to [**DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 
**MaxWaitTime** | Pointer to **time.Time** |  | [optional] 
**CommFailure** | Pointer to [**CommunicationFailure**](CommunicationFailure.md) |  | [optional] 
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PduSessType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 
**Qfi** | Pointer to **int32** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**EthfDescs** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**FDescs** | Pointer to **[]string** |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UlDelays** | Pointer to **[]int32** |  | [optional] 
**DlDelays** | Pointer to **[]int32** |  | [optional] 
**RtDelays** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification(event SmfEvent, timeStamp time.Time, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventNotification) GetEvent() SmfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventNotification) GetEventOk() (*SmfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventNotification) SetEvent(v SmfEvent)`

SetEvent sets Event field to given value.


### GetTimeStamp

`func (o *EventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *EventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *EventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSupi

`func (o *EventNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EventNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EventNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EventNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *EventNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EventNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EventNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EventNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSourceDnai

`func (o *EventNotification) GetSourceDnai() string`

GetSourceDnai returns the SourceDnai field if non-nil, zero value otherwise.

### GetSourceDnaiOk

`func (o *EventNotification) GetSourceDnaiOk() (*string, bool)`

GetSourceDnaiOk returns a tuple with the SourceDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDnai

`func (o *EventNotification) SetSourceDnai(v string)`

SetSourceDnai sets SourceDnai field to given value.

### HasSourceDnai

`func (o *EventNotification) HasSourceDnai() bool`

HasSourceDnai returns a boolean if a field has been set.

### GetTargetDnai

`func (o *EventNotification) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *EventNotification) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *EventNotification) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *EventNotification) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetDnaiChgType

`func (o *EventNotification) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *EventNotification) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *EventNotification) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.

### HasDnaiChgType

`func (o *EventNotification) HasDnaiChgType() bool`

HasDnaiChgType returns a boolean if a field has been set.

### GetSourceUeIpv4Addr

`func (o *EventNotification) GetSourceUeIpv4Addr() string`

GetSourceUeIpv4Addr returns the SourceUeIpv4Addr field if non-nil, zero value otherwise.

### GetSourceUeIpv4AddrOk

`func (o *EventNotification) GetSourceUeIpv4AddrOk() (*string, bool)`

GetSourceUeIpv4AddrOk returns a tuple with the SourceUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv4Addr

`func (o *EventNotification) SetSourceUeIpv4Addr(v string)`

SetSourceUeIpv4Addr sets SourceUeIpv4Addr field to given value.

### HasSourceUeIpv4Addr

`func (o *EventNotification) HasSourceUeIpv4Addr() bool`

HasSourceUeIpv4Addr returns a boolean if a field has been set.

### GetSourceUeIpv6Prefix

`func (o *EventNotification) GetSourceUeIpv6Prefix() Ipv6Prefix`

GetSourceUeIpv6Prefix returns the SourceUeIpv6Prefix field if non-nil, zero value otherwise.

### GetSourceUeIpv6PrefixOk

`func (o *EventNotification) GetSourceUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSourceUeIpv6PrefixOk returns a tuple with the SourceUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv6Prefix

`func (o *EventNotification) SetSourceUeIpv6Prefix(v Ipv6Prefix)`

SetSourceUeIpv6Prefix sets SourceUeIpv6Prefix field to given value.

### HasSourceUeIpv6Prefix

`func (o *EventNotification) HasSourceUeIpv6Prefix() bool`

HasSourceUeIpv6Prefix returns a boolean if a field has been set.

### GetTargetUeIpv4Addr

`func (o *EventNotification) GetTargetUeIpv4Addr() string`

GetTargetUeIpv4Addr returns the TargetUeIpv4Addr field if non-nil, zero value otherwise.

### GetTargetUeIpv4AddrOk

`func (o *EventNotification) GetTargetUeIpv4AddrOk() (*string, bool)`

GetTargetUeIpv4AddrOk returns a tuple with the TargetUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUeIpv4Addr

`func (o *EventNotification) SetTargetUeIpv4Addr(v string)`

SetTargetUeIpv4Addr sets TargetUeIpv4Addr field to given value.

### HasTargetUeIpv4Addr

`func (o *EventNotification) HasTargetUeIpv4Addr() bool`

HasTargetUeIpv4Addr returns a boolean if a field has been set.

### GetTargetUeIpv6Prefix

`func (o *EventNotification) GetTargetUeIpv6Prefix() Ipv6Prefix`

GetTargetUeIpv6Prefix returns the TargetUeIpv6Prefix field if non-nil, zero value otherwise.

### GetTargetUeIpv6PrefixOk

`func (o *EventNotification) GetTargetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetTargetUeIpv6PrefixOk returns a tuple with the TargetUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUeIpv6Prefix

`func (o *EventNotification) SetTargetUeIpv6Prefix(v Ipv6Prefix)`

SetTargetUeIpv6Prefix sets TargetUeIpv6Prefix field to given value.

### HasTargetUeIpv6Prefix

`func (o *EventNotification) HasTargetUeIpv6Prefix() bool`

HasTargetUeIpv6Prefix returns a boolean if a field has been set.

### GetSourceTraRouting

`func (o *EventNotification) GetSourceTraRouting() RouteToLocation`

GetSourceTraRouting returns the SourceTraRouting field if non-nil, zero value otherwise.

### GetSourceTraRoutingOk

`func (o *EventNotification) GetSourceTraRoutingOk() (*RouteToLocation, bool)`

GetSourceTraRoutingOk returns a tuple with the SourceTraRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTraRouting

`func (o *EventNotification) SetSourceTraRouting(v RouteToLocation)`

SetSourceTraRouting sets SourceTraRouting field to given value.

### HasSourceTraRouting

`func (o *EventNotification) HasSourceTraRouting() bool`

HasSourceTraRouting returns a boolean if a field has been set.

### SetSourceTraRoutingNil

`func (o *EventNotification) SetSourceTraRoutingNil(b bool)`

 SetSourceTraRoutingNil sets the value for SourceTraRouting to be an explicit nil

### UnsetSourceTraRouting
`func (o *EventNotification) UnsetSourceTraRouting()`

UnsetSourceTraRouting ensures that no value is present for SourceTraRouting, not even an explicit nil
### GetTargetTraRouting

`func (o *EventNotification) GetTargetTraRouting() RouteToLocation`

GetTargetTraRouting returns the TargetTraRouting field if non-nil, zero value otherwise.

### GetTargetTraRoutingOk

`func (o *EventNotification) GetTargetTraRoutingOk() (*RouteToLocation, bool)`

GetTargetTraRoutingOk returns a tuple with the TargetTraRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTraRouting

`func (o *EventNotification) SetTargetTraRouting(v RouteToLocation)`

SetTargetTraRouting sets TargetTraRouting field to given value.

### HasTargetTraRouting

`func (o *EventNotification) HasTargetTraRouting() bool`

HasTargetTraRouting returns a boolean if a field has been set.

### SetTargetTraRoutingNil

`func (o *EventNotification) SetTargetTraRoutingNil(b bool)`

 SetTargetTraRoutingNil sets the value for TargetTraRouting to be an explicit nil

### UnsetTargetTraRouting
`func (o *EventNotification) UnsetTargetTraRouting()`

UnsetTargetTraRouting ensures that no value is present for TargetTraRouting, not even an explicit nil
### GetUeMac

`func (o *EventNotification) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *EventNotification) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *EventNotification) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *EventNotification) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetAdIpv4Addr

`func (o *EventNotification) GetAdIpv4Addr() string`

GetAdIpv4Addr returns the AdIpv4Addr field if non-nil, zero value otherwise.

### GetAdIpv4AddrOk

`func (o *EventNotification) GetAdIpv4AddrOk() (*string, bool)`

GetAdIpv4AddrOk returns a tuple with the AdIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdIpv4Addr

`func (o *EventNotification) SetAdIpv4Addr(v string)`

SetAdIpv4Addr sets AdIpv4Addr field to given value.

### HasAdIpv4Addr

`func (o *EventNotification) HasAdIpv4Addr() bool`

HasAdIpv4Addr returns a boolean if a field has been set.

### GetAdIpv6Prefix

`func (o *EventNotification) GetAdIpv6Prefix() Ipv6Prefix`

GetAdIpv6Prefix returns the AdIpv6Prefix field if non-nil, zero value otherwise.

### GetAdIpv6PrefixOk

`func (o *EventNotification) GetAdIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetAdIpv6PrefixOk returns a tuple with the AdIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdIpv6Prefix

`func (o *EventNotification) SetAdIpv6Prefix(v Ipv6Prefix)`

SetAdIpv6Prefix sets AdIpv6Prefix field to given value.

### HasAdIpv6Prefix

`func (o *EventNotification) HasAdIpv6Prefix() bool`

HasAdIpv6Prefix returns a boolean if a field has been set.

### GetReIpv4Addr

`func (o *EventNotification) GetReIpv4Addr() string`

GetReIpv4Addr returns the ReIpv4Addr field if non-nil, zero value otherwise.

### GetReIpv4AddrOk

`func (o *EventNotification) GetReIpv4AddrOk() (*string, bool)`

GetReIpv4AddrOk returns a tuple with the ReIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReIpv4Addr

`func (o *EventNotification) SetReIpv4Addr(v string)`

SetReIpv4Addr sets ReIpv4Addr field to given value.

### HasReIpv4Addr

`func (o *EventNotification) HasReIpv4Addr() bool`

HasReIpv4Addr returns a boolean if a field has been set.

### GetReIpv6Prefix

`func (o *EventNotification) GetReIpv6Prefix() Ipv6Prefix`

GetReIpv6Prefix returns the ReIpv6Prefix field if non-nil, zero value otherwise.

### GetReIpv6PrefixOk

`func (o *EventNotification) GetReIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetReIpv6PrefixOk returns a tuple with the ReIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReIpv6Prefix

`func (o *EventNotification) SetReIpv6Prefix(v Ipv6Prefix)`

SetReIpv6Prefix sets ReIpv6Prefix field to given value.

### HasReIpv6Prefix

`func (o *EventNotification) HasReIpv6Prefix() bool`

HasReIpv6Prefix returns a boolean if a field has been set.

### GetPlmnId

`func (o *EventNotification) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *EventNotification) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *EventNotification) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *EventNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetAccType

`func (o *EventNotification) GetAccType() AccessType`

GetAccType returns the AccType field if non-nil, zero value otherwise.

### GetAccTypeOk

`func (o *EventNotification) GetAccTypeOk() (*AccessType, bool)`

GetAccTypeOk returns a tuple with the AccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccType

`func (o *EventNotification) SetAccType(v AccessType)`

SetAccType sets AccType field to given value.

### HasAccType

`func (o *EventNotification) HasAccType() bool`

HasAccType returns a boolean if a field has been set.

### GetPduSeId

`func (o *EventNotification) GetPduSeId() int32`

GetPduSeId returns the PduSeId field if non-nil, zero value otherwise.

### GetPduSeIdOk

`func (o *EventNotification) GetPduSeIdOk() (*int32, bool)`

GetPduSeIdOk returns a tuple with the PduSeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSeId

`func (o *EventNotification) SetPduSeId(v int32)`

SetPduSeId sets PduSeId field to given value.

### HasPduSeId

`func (o *EventNotification) HasPduSeId() bool`

HasPduSeId returns a boolean if a field has been set.

### GetDddStatus

`func (o *EventNotification) GetDddStatus() DlDataDeliveryStatus`

GetDddStatus returns the DddStatus field if non-nil, zero value otherwise.

### GetDddStatusOk

`func (o *EventNotification) GetDddStatusOk() (*DlDataDeliveryStatus, bool)`

GetDddStatusOk returns a tuple with the DddStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddStatus

`func (o *EventNotification) SetDddStatus(v DlDataDeliveryStatus)`

SetDddStatus sets DddStatus field to given value.

### HasDddStatus

`func (o *EventNotification) HasDddStatus() bool`

HasDddStatus returns a boolean if a field has been set.

### GetDddTraDescriptor

`func (o *EventNotification) GetDddTraDescriptor() DddTrafficDescriptor`

GetDddTraDescriptor returns the DddTraDescriptor field if non-nil, zero value otherwise.

### GetDddTraDescriptorOk

`func (o *EventNotification) GetDddTraDescriptorOk() (*DddTrafficDescriptor, bool)`

GetDddTraDescriptorOk returns a tuple with the DddTraDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTraDescriptor

`func (o *EventNotification) SetDddTraDescriptor(v DddTrafficDescriptor)`

SetDddTraDescriptor sets DddTraDescriptor field to given value.

### HasDddTraDescriptor

`func (o *EventNotification) HasDddTraDescriptor() bool`

HasDddTraDescriptor returns a boolean if a field has been set.

### GetMaxWaitTime

`func (o *EventNotification) GetMaxWaitTime() time.Time`

GetMaxWaitTime returns the MaxWaitTime field if non-nil, zero value otherwise.

### GetMaxWaitTimeOk

`func (o *EventNotification) GetMaxWaitTimeOk() (*time.Time, bool)`

GetMaxWaitTimeOk returns a tuple with the MaxWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitTime

`func (o *EventNotification) SetMaxWaitTime(v time.Time)`

SetMaxWaitTime sets MaxWaitTime field to given value.

### HasMaxWaitTime

`func (o *EventNotification) HasMaxWaitTime() bool`

HasMaxWaitTime returns a boolean if a field has been set.

### GetCommFailure

`func (o *EventNotification) GetCommFailure() CommunicationFailure`

GetCommFailure returns the CommFailure field if non-nil, zero value otherwise.

### GetCommFailureOk

`func (o *EventNotification) GetCommFailureOk() (*CommunicationFailure, bool)`

GetCommFailureOk returns a tuple with the CommFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommFailure

`func (o *EventNotification) SetCommFailure(v CommunicationFailure)`

SetCommFailure sets CommFailure field to given value.

### HasCommFailure

`func (o *EventNotification) HasCommFailure() bool`

HasCommFailure returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *EventNotification) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *EventNotification) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *EventNotification) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *EventNotification) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *EventNotification) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *EventNotification) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *EventNotification) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *EventNotification) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *EventNotification) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *EventNotification) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *EventNotification) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *EventNotification) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetPduSessType

`func (o *EventNotification) GetPduSessType() PduSessionType`

GetPduSessType returns the PduSessType field if non-nil, zero value otherwise.

### GetPduSessTypeOk

`func (o *EventNotification) GetPduSessTypeOk() (*PduSessionType, bool)`

GetPduSessTypeOk returns a tuple with the PduSessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessType

`func (o *EventNotification) SetPduSessType(v PduSessionType)`

SetPduSessType sets PduSessType field to given value.

### HasPduSessType

`func (o *EventNotification) HasPduSessType() bool`

HasPduSessType returns a boolean if a field has been set.

### GetQfi

`func (o *EventNotification) GetQfi() int32`

GetQfi returns the Qfi field if non-nil, zero value otherwise.

### GetQfiOk

`func (o *EventNotification) GetQfiOk() (*int32, bool)`

GetQfiOk returns a tuple with the Qfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfi

`func (o *EventNotification) SetQfi(v int32)`

SetQfi sets Qfi field to given value.

### HasQfi

`func (o *EventNotification) HasQfi() bool`

HasQfi returns a boolean if a field has been set.

### GetAppId

`func (o *EventNotification) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EventNotification) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EventNotification) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EventNotification) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetEthfDescs

`func (o *EventNotification) GetEthfDescs() []EthFlowDescription`

GetEthfDescs returns the EthfDescs field if non-nil, zero value otherwise.

### GetEthfDescsOk

`func (o *EventNotification) GetEthfDescsOk() (*[]EthFlowDescription, bool)`

GetEthfDescsOk returns a tuple with the EthfDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthfDescs

`func (o *EventNotification) SetEthfDescs(v []EthFlowDescription)`

SetEthfDescs sets EthfDescs field to given value.

### HasEthfDescs

`func (o *EventNotification) HasEthfDescs() bool`

HasEthfDescs returns a boolean if a field has been set.

### GetFDescs

`func (o *EventNotification) GetFDescs() []string`

GetFDescs returns the FDescs field if non-nil, zero value otherwise.

### GetFDescsOk

`func (o *EventNotification) GetFDescsOk() (*[]string, bool)`

GetFDescsOk returns a tuple with the FDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDescs

`func (o *EventNotification) SetFDescs(v []string)`

SetFDescs sets FDescs field to given value.

### HasFDescs

`func (o *EventNotification) HasFDescs() bool`

HasFDescs returns a boolean if a field has been set.

### GetDnn

`func (o *EventNotification) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EventNotification) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EventNotification) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EventNotification) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *EventNotification) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EventNotification) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EventNotification) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EventNotification) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetUlDelays

`func (o *EventNotification) GetUlDelays() []int32`

GetUlDelays returns the UlDelays field if non-nil, zero value otherwise.

### GetUlDelaysOk

`func (o *EventNotification) GetUlDelaysOk() (*[]int32, bool)`

GetUlDelaysOk returns a tuple with the UlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlDelays

`func (o *EventNotification) SetUlDelays(v []int32)`

SetUlDelays sets UlDelays field to given value.

### HasUlDelays

`func (o *EventNotification) HasUlDelays() bool`

HasUlDelays returns a boolean if a field has been set.

### GetDlDelays

`func (o *EventNotification) GetDlDelays() []int32`

GetDlDelays returns the DlDelays field if non-nil, zero value otherwise.

### GetDlDelaysOk

`func (o *EventNotification) GetDlDelaysOk() (*[]int32, bool)`

GetDlDelaysOk returns a tuple with the DlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDelays

`func (o *EventNotification) SetDlDelays(v []int32)`

SetDlDelays sets DlDelays field to given value.

### HasDlDelays

`func (o *EventNotification) HasDlDelays() bool`

HasDlDelays returns a boolean if a field has been set.

### GetRtDelays

`func (o *EventNotification) GetRtDelays() []int32`

GetRtDelays returns the RtDelays field if non-nil, zero value otherwise.

### GetRtDelaysOk

`func (o *EventNotification) GetRtDelaysOk() (*[]int32, bool)`

GetRtDelaysOk returns a tuple with the RtDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtDelays

`func (o *EventNotification) SetRtDelays(v []int32)`

SetRtDelays sets RtDelays field to given value.

### HasRtDelays

`func (o *EventNotification) HasRtDelays() bool`

HasRtDelays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


