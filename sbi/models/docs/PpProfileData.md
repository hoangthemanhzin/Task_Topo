# PpProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMtcProviders** | Pointer to [**map[string][]AllowedMtcProviderInfo**](array.md) | A map (list of key-value pairs where PpDataType serves as key) of AllowedMtcProviderInfo lists. In addition to defined PpDataType, the key value \&quot;ALL\&quot; may be used to identify a map entry which contains a list of AllowedMtcProviderInfo that are allowed to provision all types of the PP data for the user using UDM ParameterProvision service. | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewPpProfileData

`func NewPpProfileData() *PpProfileData`

NewPpProfileData instantiates a new PpProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpProfileDataWithDefaults

`func NewPpProfileDataWithDefaults() *PpProfileData`

NewPpProfileDataWithDefaults instantiates a new PpProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMtcProviders

`func (o *PpProfileData) GetAllowedMtcProviders() map[string][]AllowedMtcProviderInfo`

GetAllowedMtcProviders returns the AllowedMtcProviders field if non-nil, zero value otherwise.

### GetAllowedMtcProvidersOk

`func (o *PpProfileData) GetAllowedMtcProvidersOk() (*map[string][]AllowedMtcProviderInfo, bool)`

GetAllowedMtcProvidersOk returns a tuple with the AllowedMtcProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProviders

`func (o *PpProfileData) SetAllowedMtcProviders(v map[string][]AllowedMtcProviderInfo)`

SetAllowedMtcProviders sets AllowedMtcProviders field to given value.

### HasAllowedMtcProviders

`func (o *PpProfileData) HasAllowedMtcProviders() bool`

HasAllowedMtcProviders returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PpProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PpProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PpProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PpProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


