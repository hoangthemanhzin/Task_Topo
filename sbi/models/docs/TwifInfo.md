# TwifInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4EndpointAddresses** | Pointer to **[]string** |  | [optional] 
**Ipv6EndpointAddresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**EndpointFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewTwifInfo

`func NewTwifInfo() *TwifInfo`

NewTwifInfo instantiates a new TwifInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwifInfoWithDefaults

`func NewTwifInfoWithDefaults() *TwifInfo`

NewTwifInfoWithDefaults instantiates a new TwifInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4EndpointAddresses

`func (o *TwifInfo) GetIpv4EndpointAddresses() []string`

GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv4EndpointAddressesOk

`func (o *TwifInfo) GetIpv4EndpointAddressesOk() (*[]string, bool)`

GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4EndpointAddresses

`func (o *TwifInfo) SetIpv4EndpointAddresses(v []string)`

SetIpv4EndpointAddresses sets Ipv4EndpointAddresses field to given value.

### HasIpv4EndpointAddresses

`func (o *TwifInfo) HasIpv4EndpointAddresses() bool`

HasIpv4EndpointAddresses returns a boolean if a field has been set.

### GetIpv6EndpointAddresses

`func (o *TwifInfo) GetIpv6EndpointAddresses() []Ipv6Addr`

GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv6EndpointAddressesOk

`func (o *TwifInfo) GetIpv6EndpointAddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndpointAddresses

`func (o *TwifInfo) SetIpv6EndpointAddresses(v []Ipv6Addr)`

SetIpv6EndpointAddresses sets Ipv6EndpointAddresses field to given value.

### HasIpv6EndpointAddresses

`func (o *TwifInfo) HasIpv6EndpointAddresses() bool`

HasIpv6EndpointAddresses returns a boolean if a field has been set.

### GetEndpointFqdn

`func (o *TwifInfo) GetEndpointFqdn() string`

GetEndpointFqdn returns the EndpointFqdn field if non-nil, zero value otherwise.

### GetEndpointFqdnOk

`func (o *TwifInfo) GetEndpointFqdnOk() (*string, bool)`

GetEndpointFqdnOk returns a tuple with the EndpointFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointFqdn

`func (o *TwifInfo) SetEndpointFqdn(v string)`

SetEndpointFqdn sets EndpointFqdn field to given value.

### HasEndpointFqdn

`func (o *TwifInfo) HasEndpointFqdn() bool`

HasEndpointFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


