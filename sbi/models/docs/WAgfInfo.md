# WAgfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4EndpointAddresses** | Pointer to **[]string** |  | [optional] 
**Ipv6EndpointAddresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**EndpointFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewWAgfInfo

`func NewWAgfInfo() *WAgfInfo`

NewWAgfInfo instantiates a new WAgfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAgfInfoWithDefaults

`func NewWAgfInfoWithDefaults() *WAgfInfo`

NewWAgfInfoWithDefaults instantiates a new WAgfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4EndpointAddresses

`func (o *WAgfInfo) GetIpv4EndpointAddresses() []string`

GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv4EndpointAddressesOk

`func (o *WAgfInfo) GetIpv4EndpointAddressesOk() (*[]string, bool)`

GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4EndpointAddresses

`func (o *WAgfInfo) SetIpv4EndpointAddresses(v []string)`

SetIpv4EndpointAddresses sets Ipv4EndpointAddresses field to given value.

### HasIpv4EndpointAddresses

`func (o *WAgfInfo) HasIpv4EndpointAddresses() bool`

HasIpv4EndpointAddresses returns a boolean if a field has been set.

### GetIpv6EndpointAddresses

`func (o *WAgfInfo) GetIpv6EndpointAddresses() []Ipv6Addr`

GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv6EndpointAddressesOk

`func (o *WAgfInfo) GetIpv6EndpointAddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndpointAddresses

`func (o *WAgfInfo) SetIpv6EndpointAddresses(v []Ipv6Addr)`

SetIpv6EndpointAddresses sets Ipv6EndpointAddresses field to given value.

### HasIpv6EndpointAddresses

`func (o *WAgfInfo) HasIpv6EndpointAddresses() bool`

HasIpv6EndpointAddresses returns a boolean if a field has been set.

### GetEndpointFqdn

`func (o *WAgfInfo) GetEndpointFqdn() string`

GetEndpointFqdn returns the EndpointFqdn field if non-nil, zero value otherwise.

### GetEndpointFqdnOk

`func (o *WAgfInfo) GetEndpointFqdnOk() (*string, bool)`

GetEndpointFqdnOk returns a tuple with the EndpointFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointFqdn

`func (o *WAgfInfo) SetEndpointFqdn(v string)`

SetEndpointFqdn sets EndpointFqdn field to given value.

### HasEndpointFqdn

`func (o *WAgfInfo) HasEndpointFqdn() bool`

HasEndpointFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


