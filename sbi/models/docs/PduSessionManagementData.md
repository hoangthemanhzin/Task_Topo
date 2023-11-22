# PduSessionManagementData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionStatus** | Pointer to [**PduSessionStatus**](PduSessionStatus.md) |  | [optional] 
**PduSessionStatusTs** | Pointer to **time.Time** |  | [optional] 
**Dnai** | Pointer to **string** |  | [optional] 
**DnaiTs** | Pointer to **time.Time** |  | [optional] 
**N6TrafficRoutingInfo** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) |  | [optional] 
**N6TrafficRoutingInfoTs** | Pointer to **time.Time** |  | [optional] 
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Prefix** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) | UE IPv6 prefix. | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PduSessType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 
**IpAddrTs** | Pointer to **time.Time** |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**PduSessionId** | Pointer to **int32** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewPduSessionManagementData

`func NewPduSessionManagementData() *PduSessionManagementData`

NewPduSessionManagementData instantiates a new PduSessionManagementData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionManagementDataWithDefaults

`func NewPduSessionManagementDataWithDefaults() *PduSessionManagementData`

NewPduSessionManagementDataWithDefaults instantiates a new PduSessionManagementData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionStatus

`func (o *PduSessionManagementData) GetPduSessionStatus() PduSessionStatus`

GetPduSessionStatus returns the PduSessionStatus field if non-nil, zero value otherwise.

### GetPduSessionStatusOk

`func (o *PduSessionManagementData) GetPduSessionStatusOk() (*PduSessionStatus, bool)`

GetPduSessionStatusOk returns a tuple with the PduSessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionStatus

`func (o *PduSessionManagementData) SetPduSessionStatus(v PduSessionStatus)`

SetPduSessionStatus sets PduSessionStatus field to given value.

### HasPduSessionStatus

`func (o *PduSessionManagementData) HasPduSessionStatus() bool`

HasPduSessionStatus returns a boolean if a field has been set.

### GetPduSessionStatusTs

`func (o *PduSessionManagementData) GetPduSessionStatusTs() time.Time`

GetPduSessionStatusTs returns the PduSessionStatusTs field if non-nil, zero value otherwise.

### GetPduSessionStatusTsOk

`func (o *PduSessionManagementData) GetPduSessionStatusTsOk() (*time.Time, bool)`

GetPduSessionStatusTsOk returns a tuple with the PduSessionStatusTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionStatusTs

`func (o *PduSessionManagementData) SetPduSessionStatusTs(v time.Time)`

SetPduSessionStatusTs sets PduSessionStatusTs field to given value.

### HasPduSessionStatusTs

`func (o *PduSessionManagementData) HasPduSessionStatusTs() bool`

HasPduSessionStatusTs returns a boolean if a field has been set.

### GetDnai

`func (o *PduSessionManagementData) GetDnai() string`

GetDnai returns the Dnai field if non-nil, zero value otherwise.

### GetDnaiOk

`func (o *PduSessionManagementData) GetDnaiOk() (*string, bool)`

GetDnaiOk returns a tuple with the Dnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnai

`func (o *PduSessionManagementData) SetDnai(v string)`

SetDnai sets Dnai field to given value.

### HasDnai

`func (o *PduSessionManagementData) HasDnai() bool`

HasDnai returns a boolean if a field has been set.

### GetDnaiTs

`func (o *PduSessionManagementData) GetDnaiTs() time.Time`

GetDnaiTs returns the DnaiTs field if non-nil, zero value otherwise.

### GetDnaiTsOk

`func (o *PduSessionManagementData) GetDnaiTsOk() (*time.Time, bool)`

GetDnaiTsOk returns a tuple with the DnaiTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiTs

`func (o *PduSessionManagementData) SetDnaiTs(v time.Time)`

SetDnaiTs sets DnaiTs field to given value.

### HasDnaiTs

`func (o *PduSessionManagementData) HasDnaiTs() bool`

HasDnaiTs returns a boolean if a field has been set.

### GetN6TrafficRoutingInfo

`func (o *PduSessionManagementData) GetN6TrafficRoutingInfo() []RouteToLocation`

GetN6TrafficRoutingInfo returns the N6TrafficRoutingInfo field if non-nil, zero value otherwise.

### GetN6TrafficRoutingInfoOk

`func (o *PduSessionManagementData) GetN6TrafficRoutingInfoOk() (*[]RouteToLocation, bool)`

GetN6TrafficRoutingInfoOk returns a tuple with the N6TrafficRoutingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6TrafficRoutingInfo

`func (o *PduSessionManagementData) SetN6TrafficRoutingInfo(v []RouteToLocation)`

SetN6TrafficRoutingInfo sets N6TrafficRoutingInfo field to given value.

### HasN6TrafficRoutingInfo

`func (o *PduSessionManagementData) HasN6TrafficRoutingInfo() bool`

HasN6TrafficRoutingInfo returns a boolean if a field has been set.

### GetN6TrafficRoutingInfoTs

`func (o *PduSessionManagementData) GetN6TrafficRoutingInfoTs() time.Time`

GetN6TrafficRoutingInfoTs returns the N6TrafficRoutingInfoTs field if non-nil, zero value otherwise.

### GetN6TrafficRoutingInfoTsOk

`func (o *PduSessionManagementData) GetN6TrafficRoutingInfoTsOk() (*time.Time, bool)`

GetN6TrafficRoutingInfoTsOk returns a tuple with the N6TrafficRoutingInfoTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6TrafficRoutingInfoTs

`func (o *PduSessionManagementData) SetN6TrafficRoutingInfoTs(v time.Time)`

SetN6TrafficRoutingInfoTs sets N6TrafficRoutingInfoTs field to given value.

### HasN6TrafficRoutingInfoTs

`func (o *PduSessionManagementData) HasN6TrafficRoutingInfoTs() bool`

HasN6TrafficRoutingInfoTs returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *PduSessionManagementData) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PduSessionManagementData) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PduSessionManagementData) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PduSessionManagementData) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefix

`func (o *PduSessionManagementData) GetIpv6Prefix() []Ipv6Prefix`

GetIpv6Prefix returns the Ipv6Prefix field if non-nil, zero value otherwise.

### GetIpv6PrefixOk

`func (o *PduSessionManagementData) GetIpv6PrefixOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefix

`func (o *PduSessionManagementData) SetIpv6Prefix(v []Ipv6Prefix)`

SetIpv6Prefix sets Ipv6Prefix field to given value.

### HasIpv6Prefix

`func (o *PduSessionManagementData) HasIpv6Prefix() bool`

HasIpv6Prefix returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *PduSessionManagementData) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *PduSessionManagementData) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *PduSessionManagementData) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *PduSessionManagementData) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetPduSessType

`func (o *PduSessionManagementData) GetPduSessType() PduSessionType`

GetPduSessType returns the PduSessType field if non-nil, zero value otherwise.

### GetPduSessTypeOk

`func (o *PduSessionManagementData) GetPduSessTypeOk() (*PduSessionType, bool)`

GetPduSessTypeOk returns a tuple with the PduSessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessType

`func (o *PduSessionManagementData) SetPduSessType(v PduSessionType)`

SetPduSessType sets PduSessType field to given value.

### HasPduSessType

`func (o *PduSessionManagementData) HasPduSessType() bool`

HasPduSessType returns a boolean if a field has been set.

### GetIpAddrTs

`func (o *PduSessionManagementData) GetIpAddrTs() time.Time`

GetIpAddrTs returns the IpAddrTs field if non-nil, zero value otherwise.

### GetIpAddrTsOk

`func (o *PduSessionManagementData) GetIpAddrTsOk() (*time.Time, bool)`

GetIpAddrTsOk returns a tuple with the IpAddrTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddrTs

`func (o *PduSessionManagementData) SetIpAddrTs(v time.Time)`

SetIpAddrTs sets IpAddrTs field to given value.

### HasIpAddrTs

`func (o *PduSessionManagementData) HasIpAddrTs() bool`

HasIpAddrTs returns a boolean if a field has been set.

### GetDnn

`func (o *PduSessionManagementData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionManagementData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionManagementData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PduSessionManagementData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetPduSessionId

`func (o *PduSessionManagementData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduSessionManagementData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduSessionManagementData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *PduSessionManagementData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PduSessionManagementData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PduSessionManagementData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PduSessionManagementData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PduSessionManagementData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


