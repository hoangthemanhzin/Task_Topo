# EeProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedEventTypes** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**AllowedMtcProvider** | Pointer to [**map[string][]MtcProvider**](array.md) | A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \&quot;ALL\&quot; may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types. | [optional] 

## Methods

### NewEeProfileData

`func NewEeProfileData() *EeProfileData`

NewEeProfileData instantiates a new EeProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeProfileDataWithDefaults

`func NewEeProfileDataWithDefaults() *EeProfileData`

NewEeProfileDataWithDefaults instantiates a new EeProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedEventTypes

`func (o *EeProfileData) GetRestrictedEventTypes() []EventType`

GetRestrictedEventTypes returns the RestrictedEventTypes field if non-nil, zero value otherwise.

### GetRestrictedEventTypesOk

`func (o *EeProfileData) GetRestrictedEventTypesOk() (*[]EventType, bool)`

GetRestrictedEventTypesOk returns a tuple with the RestrictedEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedEventTypes

`func (o *EeProfileData) SetRestrictedEventTypes(v []EventType)`

SetRestrictedEventTypes sets RestrictedEventTypes field to given value.

### HasRestrictedEventTypes

`func (o *EeProfileData) HasRestrictedEventTypes() bool`

HasRestrictedEventTypes returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAllowedMtcProvider

`func (o *EeProfileData) GetAllowedMtcProvider() map[string][]MtcProvider`

GetAllowedMtcProvider returns the AllowedMtcProvider field if non-nil, zero value otherwise.

### GetAllowedMtcProviderOk

`func (o *EeProfileData) GetAllowedMtcProviderOk() (*map[string][]MtcProvider, bool)`

GetAllowedMtcProviderOk returns a tuple with the AllowedMtcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProvider

`func (o *EeProfileData) SetAllowedMtcProvider(v map[string][]MtcProvider)`

SetAllowedMtcProvider sets AllowedMtcProvider field to given value.

### HasAllowedMtcProvider

`func (o *EeProfileData) HasAllowedMtcProvider() bool`

HasAllowedMtcProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


