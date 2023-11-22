# ProvideLocInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLoc** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**AdditionalLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**GeoInfo** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**LocationAge** | Pointer to **int32** |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 

## Methods

### NewProvideLocInfo

`func NewProvideLocInfo() *ProvideLocInfo`

NewProvideLocInfo instantiates a new ProvideLocInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvideLocInfoWithDefaults

`func NewProvideLocInfoWithDefaults() *ProvideLocInfo`

NewProvideLocInfoWithDefaults instantiates a new ProvideLocInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLoc

`func (o *ProvideLocInfo) GetCurrentLoc() bool`

GetCurrentLoc returns the CurrentLoc field if non-nil, zero value otherwise.

### GetCurrentLocOk

`func (o *ProvideLocInfo) GetCurrentLocOk() (*bool, bool)`

GetCurrentLocOk returns a tuple with the CurrentLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLoc

`func (o *ProvideLocInfo) SetCurrentLoc(v bool)`

SetCurrentLoc sets CurrentLoc field to given value.

### HasCurrentLoc

`func (o *ProvideLocInfo) HasCurrentLoc() bool`

HasCurrentLoc returns a boolean if a field has been set.

### GetLocation

`func (o *ProvideLocInfo) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProvideLocInfo) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProvideLocInfo) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProvideLocInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAdditionalLocation

`func (o *ProvideLocInfo) GetAdditionalLocation() UserLocation`

GetAdditionalLocation returns the AdditionalLocation field if non-nil, zero value otherwise.

### GetAdditionalLocationOk

`func (o *ProvideLocInfo) GetAdditionalLocationOk() (*UserLocation, bool)`

GetAdditionalLocationOk returns a tuple with the AdditionalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLocation

`func (o *ProvideLocInfo) SetAdditionalLocation(v UserLocation)`

SetAdditionalLocation sets AdditionalLocation field to given value.

### HasAdditionalLocation

`func (o *ProvideLocInfo) HasAdditionalLocation() bool`

HasAdditionalLocation returns a boolean if a field has been set.

### GetGeoInfo

`func (o *ProvideLocInfo) GetGeoInfo() GeographicArea`

GetGeoInfo returns the GeoInfo field if non-nil, zero value otherwise.

### GetGeoInfoOk

`func (o *ProvideLocInfo) GetGeoInfoOk() (*GeographicArea, bool)`

GetGeoInfoOk returns a tuple with the GeoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoInfo

`func (o *ProvideLocInfo) SetGeoInfo(v GeographicArea)`

SetGeoInfo sets GeoInfo field to given value.

### HasGeoInfo

`func (o *ProvideLocInfo) HasGeoInfo() bool`

HasGeoInfo returns a boolean if a field has been set.

### GetLocationAge

`func (o *ProvideLocInfo) GetLocationAge() int32`

GetLocationAge returns the LocationAge field if non-nil, zero value otherwise.

### GetLocationAgeOk

`func (o *ProvideLocInfo) GetLocationAgeOk() (*int32, bool)`

GetLocationAgeOk returns a tuple with the LocationAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAge

`func (o *ProvideLocInfo) SetLocationAge(v int32)`

SetLocationAge sets LocationAge field to given value.

### HasLocationAge

`func (o *ProvideLocInfo) HasLocationAge() bool`

HasLocationAge returns a boolean if a field has been set.

### GetRatType

`func (o *ProvideLocInfo) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *ProvideLocInfo) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *ProvideLocInfo) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *ProvideLocInfo) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetTimezone

`func (o *ProvideLocInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProvideLocInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProvideLocInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProvideLocInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ProvideLocInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProvideLocInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProvideLocInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProvideLocInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *ProvideLocInfo) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *ProvideLocInfo) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *ProvideLocInfo) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *ProvideLocInfo) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


