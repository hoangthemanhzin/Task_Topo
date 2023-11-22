# RedirectInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectEnabled** | Pointer to **bool** | Indicates the redirect is enable. | [optional] 
**RedirectAddressType** | Pointer to [**RedirectAddressType**](RedirectAddressType.md) |  | [optional] 
**RedirectServerAddress** | Pointer to **string** | Indicates the address of the redirect server. If \&quot;redirectAddressType\&quot; attribute indicates the IPV4_ADDR, the encoding is the same as the Ipv4Addr data type defined in 3GPP TS 29.571.If \&quot;redirectAddressType\&quot; attribute indicates the IPV6_ADDR, the encoding is the same as the Ipv6Addr data type defined in 3GPP TS 29.571.If \&quot;redirectAddressType\&quot; attribute indicates the URL or SIP_URI, the encoding is the same as the Uri data type defined in 3GPP TS 29.571. | [optional] 

## Methods

### NewRedirectInformation

`func NewRedirectInformation() *RedirectInformation`

NewRedirectInformation instantiates a new RedirectInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectInformationWithDefaults

`func NewRedirectInformationWithDefaults() *RedirectInformation`

NewRedirectInformationWithDefaults instantiates a new RedirectInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectEnabled

`func (o *RedirectInformation) GetRedirectEnabled() bool`

GetRedirectEnabled returns the RedirectEnabled field if non-nil, zero value otherwise.

### GetRedirectEnabledOk

`func (o *RedirectInformation) GetRedirectEnabledOk() (*bool, bool)`

GetRedirectEnabledOk returns a tuple with the RedirectEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectEnabled

`func (o *RedirectInformation) SetRedirectEnabled(v bool)`

SetRedirectEnabled sets RedirectEnabled field to given value.

### HasRedirectEnabled

`func (o *RedirectInformation) HasRedirectEnabled() bool`

HasRedirectEnabled returns a boolean if a field has been set.

### GetRedirectAddressType

`func (o *RedirectInformation) GetRedirectAddressType() RedirectAddressType`

GetRedirectAddressType returns the RedirectAddressType field if non-nil, zero value otherwise.

### GetRedirectAddressTypeOk

`func (o *RedirectInformation) GetRedirectAddressTypeOk() (*RedirectAddressType, bool)`

GetRedirectAddressTypeOk returns a tuple with the RedirectAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAddressType

`func (o *RedirectInformation) SetRedirectAddressType(v RedirectAddressType)`

SetRedirectAddressType sets RedirectAddressType field to given value.

### HasRedirectAddressType

`func (o *RedirectInformation) HasRedirectAddressType() bool`

HasRedirectAddressType returns a boolean if a field has been set.

### GetRedirectServerAddress

`func (o *RedirectInformation) GetRedirectServerAddress() string`

GetRedirectServerAddress returns the RedirectServerAddress field if non-nil, zero value otherwise.

### GetRedirectServerAddressOk

`func (o *RedirectInformation) GetRedirectServerAddressOk() (*string, bool)`

GetRedirectServerAddressOk returns a tuple with the RedirectServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectServerAddress

`func (o *RedirectInformation) SetRedirectServerAddress(v string)`

SetRedirectServerAddress sets RedirectServerAddress field to given value.

### HasRedirectServerAddress

`func (o *RedirectInformation) HasRedirectServerAddress() bool`

HasRedirectServerAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


