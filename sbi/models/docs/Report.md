# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPei** | **string** |  | 
**Roaming** | **bool** |  | 
**NewServingPlmn** | [**PlmnId**](PlmnId.md) |  | 
**NewCnType** | [**CnType**](CnType.md) |  | 
**OldCnType** | Pointer to [**CnType**](CnType.md) |  | [optional] 
**OldCmInfoList** | Pointer to [**[]CmInfo**](CmInfo.md) |  | [optional] 
**NewCmInfoList** | [**[]CmInfo**](CmInfo.md) |  | 
**LossOfConnectReason** | [**LossOfConnectivityReason**](LossOfConnectivityReason.md) |  | 
**Location** | [**UserLocation**](UserLocation.md) |  | 
**PdnConnStat** | [**PdnConnectivityStatus**](PdnConnectivityStatus.md) |  | 
**Dnn** | Pointer to **string** |  | [optional] 
**PduSeId** | Pointer to **int32** |  | [optional] 
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Prefixes** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PduSessType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewReport

`func NewReport(newPei string, roaming bool, newServingPlmn PlmnId, newCnType CnType, newCmInfoList []CmInfo, lossOfConnectReason LossOfConnectivityReason, location UserLocation, pdnConnStat PdnConnectivityStatus, ) *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPei

`func (o *Report) GetNewPei() string`

GetNewPei returns the NewPei field if non-nil, zero value otherwise.

### GetNewPeiOk

`func (o *Report) GetNewPeiOk() (*string, bool)`

GetNewPeiOk returns a tuple with the NewPei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPei

`func (o *Report) SetNewPei(v string)`

SetNewPei sets NewPei field to given value.


### GetRoaming

`func (o *Report) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *Report) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *Report) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.


### GetNewServingPlmn

`func (o *Report) GetNewServingPlmn() PlmnId`

GetNewServingPlmn returns the NewServingPlmn field if non-nil, zero value otherwise.

### GetNewServingPlmnOk

`func (o *Report) GetNewServingPlmnOk() (*PlmnId, bool)`

GetNewServingPlmnOk returns a tuple with the NewServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewServingPlmn

`func (o *Report) SetNewServingPlmn(v PlmnId)`

SetNewServingPlmn sets NewServingPlmn field to given value.


### GetNewCnType

`func (o *Report) GetNewCnType() CnType`

GetNewCnType returns the NewCnType field if non-nil, zero value otherwise.

### GetNewCnTypeOk

`func (o *Report) GetNewCnTypeOk() (*CnType, bool)`

GetNewCnTypeOk returns a tuple with the NewCnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCnType

`func (o *Report) SetNewCnType(v CnType)`

SetNewCnType sets NewCnType field to given value.


### GetOldCnType

`func (o *Report) GetOldCnType() CnType`

GetOldCnType returns the OldCnType field if non-nil, zero value otherwise.

### GetOldCnTypeOk

`func (o *Report) GetOldCnTypeOk() (*CnType, bool)`

GetOldCnTypeOk returns a tuple with the OldCnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCnType

`func (o *Report) SetOldCnType(v CnType)`

SetOldCnType sets OldCnType field to given value.

### HasOldCnType

`func (o *Report) HasOldCnType() bool`

HasOldCnType returns a boolean if a field has been set.

### GetOldCmInfoList

`func (o *Report) GetOldCmInfoList() []CmInfo`

GetOldCmInfoList returns the OldCmInfoList field if non-nil, zero value otherwise.

### GetOldCmInfoListOk

`func (o *Report) GetOldCmInfoListOk() (*[]CmInfo, bool)`

GetOldCmInfoListOk returns a tuple with the OldCmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCmInfoList

`func (o *Report) SetOldCmInfoList(v []CmInfo)`

SetOldCmInfoList sets OldCmInfoList field to given value.

### HasOldCmInfoList

`func (o *Report) HasOldCmInfoList() bool`

HasOldCmInfoList returns a boolean if a field has been set.

### GetNewCmInfoList

`func (o *Report) GetNewCmInfoList() []CmInfo`

GetNewCmInfoList returns the NewCmInfoList field if non-nil, zero value otherwise.

### GetNewCmInfoListOk

`func (o *Report) GetNewCmInfoListOk() (*[]CmInfo, bool)`

GetNewCmInfoListOk returns a tuple with the NewCmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCmInfoList

`func (o *Report) SetNewCmInfoList(v []CmInfo)`

SetNewCmInfoList sets NewCmInfoList field to given value.


### GetLossOfConnectReason

`func (o *Report) GetLossOfConnectReason() LossOfConnectivityReason`

GetLossOfConnectReason returns the LossOfConnectReason field if non-nil, zero value otherwise.

### GetLossOfConnectReasonOk

`func (o *Report) GetLossOfConnectReasonOk() (*LossOfConnectivityReason, bool)`

GetLossOfConnectReasonOk returns a tuple with the LossOfConnectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfConnectReason

`func (o *Report) SetLossOfConnectReason(v LossOfConnectivityReason)`

SetLossOfConnectReason sets LossOfConnectReason field to given value.


### GetLocation

`func (o *Report) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Report) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Report) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.


### GetPdnConnStat

`func (o *Report) GetPdnConnStat() PdnConnectivityStatus`

GetPdnConnStat returns the PdnConnStat field if non-nil, zero value otherwise.

### GetPdnConnStatOk

`func (o *Report) GetPdnConnStatOk() (*PdnConnectivityStatus, bool)`

GetPdnConnStatOk returns a tuple with the PdnConnStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnConnStat

`func (o *Report) SetPdnConnStat(v PdnConnectivityStatus)`

SetPdnConnStat sets PdnConnStat field to given value.


### GetDnn

`func (o *Report) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Report) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Report) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *Report) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetPduSeId

`func (o *Report) GetPduSeId() int32`

GetPduSeId returns the PduSeId field if non-nil, zero value otherwise.

### GetPduSeIdOk

`func (o *Report) GetPduSeIdOk() (*int32, bool)`

GetPduSeIdOk returns a tuple with the PduSeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSeId

`func (o *Report) SetPduSeId(v int32)`

SetPduSeId sets PduSeId field to given value.

### HasPduSeId

`func (o *Report) HasPduSeId() bool`

HasPduSeId returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *Report) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *Report) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *Report) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *Report) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Prefixes

`func (o *Report) GetIpv6Prefixes() []Ipv6Prefix`

GetIpv6Prefixes returns the Ipv6Prefixes field if non-nil, zero value otherwise.

### GetIpv6PrefixesOk

`func (o *Report) GetIpv6PrefixesOk() (*[]Ipv6Prefix, bool)`

GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefixes

`func (o *Report) SetIpv6Prefixes(v []Ipv6Prefix)`

SetIpv6Prefixes sets Ipv6Prefixes field to given value.

### HasIpv6Prefixes

`func (o *Report) HasIpv6Prefixes() bool`

HasIpv6Prefixes returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *Report) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *Report) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *Report) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *Report) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.

### GetPduSessType

`func (o *Report) GetPduSessType() PduSessionType`

GetPduSessType returns the PduSessType field if non-nil, zero value otherwise.

### GetPduSessTypeOk

`func (o *Report) GetPduSessTypeOk() (*PduSessionType, bool)`

GetPduSessTypeOk returns a tuple with the PduSessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessType

`func (o *Report) SetPduSessType(v PduSessionType)`

SetPduSessType sets PduSessType field to given value.

### HasPduSessType

`func (o *Report) HasPduSessType() bool`

HasPduSessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


