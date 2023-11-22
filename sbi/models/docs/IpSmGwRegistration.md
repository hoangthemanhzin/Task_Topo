# IpSmGwRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSmGwMapAddress** | Pointer to **string** |  | [optional] 
**IpSmGwDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**UnriIndicator** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIpSmGwRegistration

`func NewIpSmGwRegistration() *IpSmGwRegistration`

NewIpSmGwRegistration instantiates a new IpSmGwRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSmGwRegistrationWithDefaults

`func NewIpSmGwRegistrationWithDefaults() *IpSmGwRegistration`

NewIpSmGwRegistrationWithDefaults instantiates a new IpSmGwRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSmGwMapAddress

`func (o *IpSmGwRegistration) GetIpSmGwMapAddress() string`

GetIpSmGwMapAddress returns the IpSmGwMapAddress field if non-nil, zero value otherwise.

### GetIpSmGwMapAddressOk

`func (o *IpSmGwRegistration) GetIpSmGwMapAddressOk() (*string, bool)`

GetIpSmGwMapAddressOk returns a tuple with the IpSmGwMapAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwMapAddress

`func (o *IpSmGwRegistration) SetIpSmGwMapAddress(v string)`

SetIpSmGwMapAddress sets IpSmGwMapAddress field to given value.

### HasIpSmGwMapAddress

`func (o *IpSmGwRegistration) HasIpSmGwMapAddress() bool`

HasIpSmGwMapAddress returns a boolean if a field has been set.

### GetIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) GetIpSmGwDiameterAddress() NetworkNodeDiameterAddress`

GetIpSmGwDiameterAddress returns the IpSmGwDiameterAddress field if non-nil, zero value otherwise.

### GetIpSmGwDiameterAddressOk

`func (o *IpSmGwRegistration) GetIpSmGwDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetIpSmGwDiameterAddressOk returns a tuple with the IpSmGwDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) SetIpSmGwDiameterAddress(v NetworkNodeDiameterAddress)`

SetIpSmGwDiameterAddress sets IpSmGwDiameterAddress field to given value.

### HasIpSmGwDiameterAddress

`func (o *IpSmGwRegistration) HasIpSmGwDiameterAddress() bool`

HasIpSmGwDiameterAddress returns a boolean if a field has been set.

### GetUnriIndicator

`func (o *IpSmGwRegistration) GetUnriIndicator() bool`

GetUnriIndicator returns the UnriIndicator field if non-nil, zero value otherwise.

### GetUnriIndicatorOk

`func (o *IpSmGwRegistration) GetUnriIndicatorOk() (*bool, bool)`

GetUnriIndicatorOk returns a tuple with the UnriIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnriIndicator

`func (o *IpSmGwRegistration) SetUnriIndicator(v bool)`

SetUnriIndicator sets UnriIndicator field to given value.

### HasUnriIndicator

`func (o *IpSmGwRegistration) HasUnriIndicator() bool`

HasUnriIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


