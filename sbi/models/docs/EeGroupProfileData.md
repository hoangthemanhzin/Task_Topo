# EeGroupProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedEventTypes** | Pointer to [**[]EventType**](EventType.md) |  | [optional] 
**AllowedMtcProvider** | Pointer to [**map[string][]MtcProvider**](array.md) | A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \&quot;ALL\&quot; may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types. | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewEeGroupProfileData

`func NewEeGroupProfileData() *EeGroupProfileData`

NewEeGroupProfileData instantiates a new EeGroupProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeGroupProfileDataWithDefaults

`func NewEeGroupProfileDataWithDefaults() *EeGroupProfileData`

NewEeGroupProfileDataWithDefaults instantiates a new EeGroupProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedEventTypes

`func (o *EeGroupProfileData) GetRestrictedEventTypes() []EventType`

GetRestrictedEventTypes returns the RestrictedEventTypes field if non-nil, zero value otherwise.

### GetRestrictedEventTypesOk

`func (o *EeGroupProfileData) GetRestrictedEventTypesOk() (*[]EventType, bool)`

GetRestrictedEventTypesOk returns a tuple with the RestrictedEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedEventTypes

`func (o *EeGroupProfileData) SetRestrictedEventTypes(v []EventType)`

SetRestrictedEventTypes sets RestrictedEventTypes field to given value.

### HasRestrictedEventTypes

`func (o *EeGroupProfileData) HasRestrictedEventTypes() bool`

HasRestrictedEventTypes returns a boolean if a field has been set.

### GetAllowedMtcProvider

`func (o *EeGroupProfileData) GetAllowedMtcProvider() map[string][]MtcProvider`

GetAllowedMtcProvider returns the AllowedMtcProvider field if non-nil, zero value otherwise.

### GetAllowedMtcProviderOk

`func (o *EeGroupProfileData) GetAllowedMtcProviderOk() (*map[string][]MtcProvider, bool)`

GetAllowedMtcProviderOk returns a tuple with the AllowedMtcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProvider

`func (o *EeGroupProfileData) SetAllowedMtcProvider(v map[string][]MtcProvider)`

SetAllowedMtcProvider sets AllowedMtcProvider field to given value.

### HasAllowedMtcProvider

`func (o *EeGroupProfileData) HasAllowedMtcProvider() bool`

HasAllowedMtcProvider returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeGroupProfileData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeGroupProfileData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeGroupProfileData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeGroupProfileData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


