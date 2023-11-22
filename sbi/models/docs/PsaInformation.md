# PsaInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PsaInd** | Pointer to [**PsaIndication**](PsaIndication.md) |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**UeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**PsaUpfId** | Pointer to **string** |  | [optional] 

## Methods

### NewPsaInformation

`func NewPsaInformation() *PsaInformation`

NewPsaInformation instantiates a new PsaInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsaInformationWithDefaults

`func NewPsaInformationWithDefaults() *PsaInformation`

NewPsaInformationWithDefaults instantiates a new PsaInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsaInd

`func (o *PsaInformation) GetPsaInd() PsaIndication`

GetPsaInd returns the PsaInd field if non-nil, zero value otherwise.

### GetPsaIndOk

`func (o *PsaInformation) GetPsaIndOk() (*PsaIndication, bool)`

GetPsaIndOk returns a tuple with the PsaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsaInd

`func (o *PsaInformation) SetPsaInd(v PsaIndication)`

SetPsaInd sets PsaInd field to given value.

### HasPsaInd

`func (o *PsaInformation) HasPsaInd() bool`

HasPsaInd returns a boolean if a field has been set.

### GetDnaiList

`func (o *PsaInformation) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *PsaInformation) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *PsaInformation) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *PsaInformation) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetUeIpv6Prefix

`func (o *PsaInformation) GetUeIpv6Prefix() Ipv6Prefix`

GetUeIpv6Prefix returns the UeIpv6Prefix field if non-nil, zero value otherwise.

### GetUeIpv6PrefixOk

`func (o *PsaInformation) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Prefix

`func (o *PsaInformation) SetUeIpv6Prefix(v Ipv6Prefix)`

SetUeIpv6Prefix sets UeIpv6Prefix field to given value.

### HasUeIpv6Prefix

`func (o *PsaInformation) HasUeIpv6Prefix() bool`

HasUeIpv6Prefix returns a boolean if a field has been set.

### GetPsaUpfId

`func (o *PsaInformation) GetPsaUpfId() string`

GetPsaUpfId returns the PsaUpfId field if non-nil, zero value otherwise.

### GetPsaUpfIdOk

`func (o *PsaInformation) GetPsaUpfIdOk() (*string, bool)`

GetPsaUpfIdOk returns a tuple with the PsaUpfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsaUpfId

`func (o *PsaInformation) SetPsaUpfId(v string)`

SetPsaUpfId sets PsaUpfId field to given value.

### HasPsaUpfId

`func (o *PsaInformation) HasPsaUpfId() bool`

HasPsaUpfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


