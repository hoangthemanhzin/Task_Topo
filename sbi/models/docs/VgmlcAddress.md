# VgmlcAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VgmlcAddressIpv4** | Pointer to **string** |  | [optional] 
**VgmlcAddressIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**VgmlcFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewVgmlcAddress

`func NewVgmlcAddress() *VgmlcAddress`

NewVgmlcAddress instantiates a new VgmlcAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVgmlcAddressWithDefaults

`func NewVgmlcAddressWithDefaults() *VgmlcAddress`

NewVgmlcAddressWithDefaults instantiates a new VgmlcAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVgmlcAddressIpv4

`func (o *VgmlcAddress) GetVgmlcAddressIpv4() string`

GetVgmlcAddressIpv4 returns the VgmlcAddressIpv4 field if non-nil, zero value otherwise.

### GetVgmlcAddressIpv4Ok

`func (o *VgmlcAddress) GetVgmlcAddressIpv4Ok() (*string, bool)`

GetVgmlcAddressIpv4Ok returns a tuple with the VgmlcAddressIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddressIpv4

`func (o *VgmlcAddress) SetVgmlcAddressIpv4(v string)`

SetVgmlcAddressIpv4 sets VgmlcAddressIpv4 field to given value.

### HasVgmlcAddressIpv4

`func (o *VgmlcAddress) HasVgmlcAddressIpv4() bool`

HasVgmlcAddressIpv4 returns a boolean if a field has been set.

### GetVgmlcAddressIpv6

`func (o *VgmlcAddress) GetVgmlcAddressIpv6() Ipv6Addr`

GetVgmlcAddressIpv6 returns the VgmlcAddressIpv6 field if non-nil, zero value otherwise.

### GetVgmlcAddressIpv6Ok

`func (o *VgmlcAddress) GetVgmlcAddressIpv6Ok() (*Ipv6Addr, bool)`

GetVgmlcAddressIpv6Ok returns a tuple with the VgmlcAddressIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddressIpv6

`func (o *VgmlcAddress) SetVgmlcAddressIpv6(v Ipv6Addr)`

SetVgmlcAddressIpv6 sets VgmlcAddressIpv6 field to given value.

### HasVgmlcAddressIpv6

`func (o *VgmlcAddress) HasVgmlcAddressIpv6() bool`

HasVgmlcAddressIpv6 returns a boolean if a field has been set.

### GetVgmlcFqdn

`func (o *VgmlcAddress) GetVgmlcFqdn() string`

GetVgmlcFqdn returns the VgmlcFqdn field if non-nil, zero value otherwise.

### GetVgmlcFqdnOk

`func (o *VgmlcAddress) GetVgmlcFqdnOk() (*string, bool)`

GetVgmlcFqdnOk returns a tuple with the VgmlcFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcFqdn

`func (o *VgmlcAddress) SetVgmlcFqdn(v string)`

SetVgmlcFqdn sets VgmlcFqdn field to given value.

### HasVgmlcFqdn

`func (o *VgmlcAddress) HasVgmlcFqdn() bool`

HasVgmlcFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


