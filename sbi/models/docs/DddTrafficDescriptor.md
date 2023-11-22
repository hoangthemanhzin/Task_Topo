# DddTrafficDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addr** | Pointer to **string** |  | [optional] 
**Ipv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PortNumber** | Pointer to **int32** |  | [optional] 
**MacAddr** | Pointer to **string** |  | [optional] 

## Methods

### NewDddTrafficDescriptor

`func NewDddTrafficDescriptor() *DddTrafficDescriptor`

NewDddTrafficDescriptor instantiates a new DddTrafficDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDddTrafficDescriptorWithDefaults

`func NewDddTrafficDescriptorWithDefaults() *DddTrafficDescriptor`

NewDddTrafficDescriptorWithDefaults instantiates a new DddTrafficDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addr

`func (o *DddTrafficDescriptor) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *DddTrafficDescriptor) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *DddTrafficDescriptor) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *DddTrafficDescriptor) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *DddTrafficDescriptor) GetIpv6Addr() Ipv6Addr`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *DddTrafficDescriptor) GetIpv6AddrOk() (*Ipv6Addr, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *DddTrafficDescriptor) SetIpv6Addr(v Ipv6Addr)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *DddTrafficDescriptor) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetPortNumber

`func (o *DddTrafficDescriptor) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *DddTrafficDescriptor) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *DddTrafficDescriptor) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *DddTrafficDescriptor) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### GetMacAddr

`func (o *DddTrafficDescriptor) GetMacAddr() string`

GetMacAddr returns the MacAddr field if non-nil, zero value otherwise.

### GetMacAddrOk

`func (o *DddTrafficDescriptor) GetMacAddrOk() (*string, bool)`

GetMacAddrOk returns a tuple with the MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr

`func (o *DddTrafficDescriptor) SetMacAddr(v string)`

SetMacAddr sets MacAddr field to given value.

### HasMacAddr

`func (o *DddTrafficDescriptor) HasMacAddr() bool`

HasMacAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


