# TunnelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**GtpTeid** | **string** |  | 
**AnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewTunnelInfo

`func NewTunnelInfo(gtpTeid string, ) *TunnelInfo`

NewTunnelInfo instantiates a new TunnelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelInfoWithDefaults

`func NewTunnelInfoWithDefaults() *TunnelInfo`

NewTunnelInfoWithDefaults instantiates a new TunnelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *TunnelInfo) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *TunnelInfo) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *TunnelInfo) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *TunnelInfo) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *TunnelInfo) GetIpv6Addr() Ipv6Addr`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *TunnelInfo) GetIpv6AddrOk() (*Ipv6Addr, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *TunnelInfo) SetIpv6Addr(v Ipv6Addr)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *TunnelInfo) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetGtpTeid

`func (o *TunnelInfo) GetGtpTeid() string`

GetGtpTeid returns the GtpTeid field if non-nil, zero value otherwise.

### GetGtpTeidOk

`func (o *TunnelInfo) GetGtpTeidOk() (*string, bool)`

GetGtpTeidOk returns a tuple with the GtpTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtpTeid

`func (o *TunnelInfo) SetGtpTeid(v string)`

SetGtpTeid sets GtpTeid field to given value.


### GetAnType

`func (o *TunnelInfo) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *TunnelInfo) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *TunnelInfo) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.

### HasAnType

`func (o *TunnelInfo) HasAnType() bool`

HasAnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


