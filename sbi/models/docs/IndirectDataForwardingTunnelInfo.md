# IndirectDataForwardingTunnelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**GtpTeid** | **string** |  | 
**DrbId** | Pointer to **int32** |  | [optional] 
**AdditionalTnlNb** | Pointer to **int32** |  | [optional] 

## Methods

### NewIndirectDataForwardingTunnelInfo

`func NewIndirectDataForwardingTunnelInfo(gtpTeid string, ) *IndirectDataForwardingTunnelInfo`

NewIndirectDataForwardingTunnelInfo instantiates a new IndirectDataForwardingTunnelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirectDataForwardingTunnelInfoWithDefaults

`func NewIndirectDataForwardingTunnelInfoWithDefaults() *IndirectDataForwardingTunnelInfo`

NewIndirectDataForwardingTunnelInfoWithDefaults instantiates a new IndirectDataForwardingTunnelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *IndirectDataForwardingTunnelInfo) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *IndirectDataForwardingTunnelInfo) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *IndirectDataForwardingTunnelInfo) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *IndirectDataForwardingTunnelInfo) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *IndirectDataForwardingTunnelInfo) GetIpv6Addr() Ipv6Addr`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *IndirectDataForwardingTunnelInfo) GetIpv6AddrOk() (*Ipv6Addr, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *IndirectDataForwardingTunnelInfo) SetIpv6Addr(v Ipv6Addr)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *IndirectDataForwardingTunnelInfo) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetGtpTeid

`func (o *IndirectDataForwardingTunnelInfo) GetGtpTeid() string`

GetGtpTeid returns the GtpTeid field if non-nil, zero value otherwise.

### GetGtpTeidOk

`func (o *IndirectDataForwardingTunnelInfo) GetGtpTeidOk() (*string, bool)`

GetGtpTeidOk returns a tuple with the GtpTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtpTeid

`func (o *IndirectDataForwardingTunnelInfo) SetGtpTeid(v string)`

SetGtpTeid sets GtpTeid field to given value.


### GetDrbId

`func (o *IndirectDataForwardingTunnelInfo) GetDrbId() int32`

GetDrbId returns the DrbId field if non-nil, zero value otherwise.

### GetDrbIdOk

`func (o *IndirectDataForwardingTunnelInfo) GetDrbIdOk() (*int32, bool)`

GetDrbIdOk returns a tuple with the DrbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrbId

`func (o *IndirectDataForwardingTunnelInfo) SetDrbId(v int32)`

SetDrbId sets DrbId field to given value.

### HasDrbId

`func (o *IndirectDataForwardingTunnelInfo) HasDrbId() bool`

HasDrbId returns a boolean if a field has been set.

### GetAdditionalTnlNb

`func (o *IndirectDataForwardingTunnelInfo) GetAdditionalTnlNb() int32`

GetAdditionalTnlNb returns the AdditionalTnlNb field if non-nil, zero value otherwise.

### GetAdditionalTnlNbOk

`func (o *IndirectDataForwardingTunnelInfo) GetAdditionalTnlNbOk() (*int32, bool)`

GetAdditionalTnlNbOk returns a tuple with the AdditionalTnlNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTnlNb

`func (o *IndirectDataForwardingTunnelInfo) SetAdditionalTnlNb(v int32)`

SetAdditionalTnlNb sets AdditionalTnlNb field to given value.

### HasAdditionalTnlNb

`func (o *IndirectDataForwardingTunnelInfo) HasAdditionalTnlNb() bool`

HasAdditionalTnlNb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


