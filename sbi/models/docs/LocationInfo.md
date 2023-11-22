# LocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**RegistrationLocationInfoList** | [**[]RegistrationLocationInfo**](RegistrationLocationInfo.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationInfo

`func NewLocationInfo(registrationLocationInfoList []RegistrationLocationInfo, ) *LocationInfo`

NewLocationInfo instantiates a new LocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoWithDefaults

`func NewLocationInfoWithDefaults() *LocationInfo`

NewLocationInfoWithDefaults instantiates a new LocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *LocationInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *LocationInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *LocationInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *LocationInfo) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *LocationInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocationInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocationInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *LocationInfo) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetRegistrationLocationInfoList

`func (o *LocationInfo) GetRegistrationLocationInfoList() []RegistrationLocationInfo`

GetRegistrationLocationInfoList returns the RegistrationLocationInfoList field if non-nil, zero value otherwise.

### GetRegistrationLocationInfoListOk

`func (o *LocationInfo) GetRegistrationLocationInfoListOk() (*[]RegistrationLocationInfo, bool)`

GetRegistrationLocationInfoListOk returns a tuple with the RegistrationLocationInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationLocationInfoList

`func (o *LocationInfo) SetRegistrationLocationInfoList(v []RegistrationLocationInfo)`

SetRegistrationLocationInfoList sets RegistrationLocationInfoList field to given value.


### GetSupportedFeatures

`func (o *LocationInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *LocationInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *LocationInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *LocationInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


