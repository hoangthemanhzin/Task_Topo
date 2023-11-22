# Nssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**DefaultSingleNssais** | [**[]Snssai**](Snssai.md) |  | 
**SingleNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**ProvisioningTime** | Pointer to **time.Time** |  | [optional] 
**AdditionalSnssaiData** | Pointer to [**map[string]AdditionalSnssaiData**](AdditionalSnssaiData.md) |  | [optional] 

## Methods

### NewNssai

`func NewNssai(defaultSingleNssais []Snssai, ) *Nssai`

NewNssai instantiates a new Nssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssaiWithDefaults

`func NewNssaiWithDefaults() *Nssai`

NewNssaiWithDefaults instantiates a new Nssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *Nssai) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Nssai) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Nssai) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Nssai) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetDefaultSingleNssais

`func (o *Nssai) GetDefaultSingleNssais() []Snssai`

GetDefaultSingleNssais returns the DefaultSingleNssais field if non-nil, zero value otherwise.

### GetDefaultSingleNssaisOk

`func (o *Nssai) GetDefaultSingleNssaisOk() (*[]Snssai, bool)`

GetDefaultSingleNssaisOk returns a tuple with the DefaultSingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleNssais

`func (o *Nssai) SetDefaultSingleNssais(v []Snssai)`

SetDefaultSingleNssais sets DefaultSingleNssais field to given value.


### GetSingleNssais

`func (o *Nssai) GetSingleNssais() []Snssai`

GetSingleNssais returns the SingleNssais field if non-nil, zero value otherwise.

### GetSingleNssaisOk

`func (o *Nssai) GetSingleNssaisOk() (*[]Snssai, bool)`

GetSingleNssaisOk returns a tuple with the SingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssais

`func (o *Nssai) SetSingleNssais(v []Snssai)`

SetSingleNssais sets SingleNssais field to given value.

### HasSingleNssais

`func (o *Nssai) HasSingleNssais() bool`

HasSingleNssais returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *Nssai) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *Nssai) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *Nssai) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *Nssai) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.

### GetAdditionalSnssaiData

`func (o *Nssai) GetAdditionalSnssaiData() map[string]AdditionalSnssaiData`

GetAdditionalSnssaiData returns the AdditionalSnssaiData field if non-nil, zero value otherwise.

### GetAdditionalSnssaiDataOk

`func (o *Nssai) GetAdditionalSnssaiDataOk() (*map[string]AdditionalSnssaiData, bool)`

GetAdditionalSnssaiDataOk returns a tuple with the AdditionalSnssaiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSnssaiData

`func (o *Nssai) SetAdditionalSnssaiData(v map[string]AdditionalSnssaiData)`

SetAdditionalSnssaiData sets AdditionalSnssaiData field to given value.

### HasAdditionalSnssaiData

`func (o *Nssai) HasAdditionalSnssaiData() bool`

HasAdditionalSnssaiData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


