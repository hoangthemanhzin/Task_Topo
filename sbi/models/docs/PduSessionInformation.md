# PduSessionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | **string** |  | 
**UeIpv4** | Pointer to **string** |  | [optional] 
**UeIpv6** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**UeMac** | Pointer to **string** |  | [optional] 

## Methods

### NewPduSessionInformation

`func NewPduSessionInformation(snssai Snssai, dnn string, ) *PduSessionInformation`

NewPduSessionInformation instantiates a new PduSessionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionInformationWithDefaults

`func NewPduSessionInformationWithDefaults() *PduSessionInformation`

NewPduSessionInformationWithDefaults instantiates a new PduSessionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *PduSessionInformation) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PduSessionInformation) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PduSessionInformation) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnn

`func (o *PduSessionInformation) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionInformation) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionInformation) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetUeIpv4

`func (o *PduSessionInformation) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *PduSessionInformation) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *PduSessionInformation) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *PduSessionInformation) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *PduSessionInformation) GetUeIpv6() Ipv6Prefix`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *PduSessionInformation) GetUeIpv6Ok() (*Ipv6Prefix, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *PduSessionInformation) SetUeIpv6(v Ipv6Prefix)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *PduSessionInformation) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetIpDomain

`func (o *PduSessionInformation) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *PduSessionInformation) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *PduSessionInformation) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *PduSessionInformation) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUeMac

`func (o *PduSessionInformation) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *PduSessionInformation) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *PduSessionInformation) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *PduSessionInformation) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


