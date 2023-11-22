# PdnConnectivityStatReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PdnConnStat** | [**PdnConnectivityStatus**](PdnConnectivityStatus.md) |  | 
**Dnn** | Pointer to **string** |  | [optional] 
**PduSeId** | Pointer to **int32** |  | [optional] 
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PduSessType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewPdnConnectivityStatReport

`func NewPdnConnectivityStatReport(pdnConnStat PdnConnectivityStatus, ) *PdnConnectivityStatReport`

NewPdnConnectivityStatReport instantiates a new PdnConnectivityStatReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdnConnectivityStatReportWithDefaults

`func NewPdnConnectivityStatReportWithDefaults() *PdnConnectivityStatReport`

NewPdnConnectivityStatReportWithDefaults instantiates a new PdnConnectivityStatReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPdnConnStat

`func (o *PdnConnectivityStatReport) GetPdnConnStat() PdnConnectivityStatus`

GetPdnConnStat returns the PdnConnStat field if non-nil, zero value otherwise.

### GetPdnConnStatOk

`func (o *PdnConnectivityStatReport) GetPdnConnStatOk() (*PdnConnectivityStatus, bool)`

GetPdnConnStatOk returns a tuple with the PdnConnStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnConnStat

`func (o *PdnConnectivityStatReport) SetPdnConnStat(v PdnConnectivityStatus)`

SetPdnConnStat sets PdnConnStat field to given value.


### GetDnn

`func (o *PdnConnectivityStatReport) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PdnConnectivityStatReport) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PdnConnectivityStatReport) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PdnConnectivityStatReport) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetPduSeId

`func (o *PdnConnectivityStatReport) GetPduSeId() int32`

GetPduSeId returns the PduSeId field if non-nil, zero value otherwise.

### GetPduSeIdOk

`func (o *PdnConnectivityStatReport) GetPduSeIdOk() (*int32, bool)`

GetPduSeIdOk returns a tuple with the PduSeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSeId

`func (o *PdnConnectivityStatReport) SetPduSeId(v int32)`

SetPduSeId sets PduSeId field to given value.

### HasPduSeId

`func (o *PdnConnectivityStatReport) HasPduSeId() bool`

HasPduSeId returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *PdnConnectivityStatReport) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *PdnConnectivityStatReport) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *PdnConnectivityStatReport) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *PdnConnectivityStatReport) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *PdnConnectivityStatReport) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *PdnConnectivityStatReport) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *PdnConnectivityStatReport) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *PdnConnectivityStatReport) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *PdnConnectivityStatReport) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *PdnConnectivityStatReport) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *PdnConnectivityStatReport) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *PdnConnectivityStatReport) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetPduSessType

`func (o *PdnConnectivityStatReport) GetPduSessType() PduSessionType`

GetPduSessType returns the PduSessType field if non-nil, zero value otherwise.

### GetPduSessTypeOk

`func (o *PdnConnectivityStatReport) GetPduSessTypeOk() (*PduSessionType, bool)`

GetPduSessTypeOk returns a tuple with the PduSessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessType

`func (o *PdnConnectivityStatReport) SetPduSessType(v PduSessionType)`

SetPduSessType sets PduSessType field to given value.

### HasPduSessType

`func (o *PdnConnectivityStatReport) HasPduSessType() bool`

HasPduSessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


