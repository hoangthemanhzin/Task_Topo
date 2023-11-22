# Nssai1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**DefaultSingleNssais** | [**[]Snssai**](Snssai.md) |  | 
**SingleNssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**ProvisioningTime** | Pointer to **time.Time** |  | [optional] 
**AdditionalSnssaiData** | Pointer to [**map[string]AdditionalSnssaiData**](AdditionalSnssaiData.md) |  | [optional] 

## Methods

### NewNssai1

`func NewNssai1(defaultSingleNssais []Snssai, ) *Nssai1`

NewNssai1 instantiates a new Nssai1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssai1WithDefaults

`func NewNssai1WithDefaults() *Nssai1`

NewNssai1WithDefaults instantiates a new Nssai1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *Nssai1) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Nssai1) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Nssai1) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Nssai1) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetDefaultSingleNssais

`func (o *Nssai1) GetDefaultSingleNssais() []Snssai`

GetDefaultSingleNssais returns the DefaultSingleNssais field if non-nil, zero value otherwise.

### GetDefaultSingleNssaisOk

`func (o *Nssai1) GetDefaultSingleNssaisOk() (*[]Snssai, bool)`

GetDefaultSingleNssaisOk returns a tuple with the DefaultSingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleNssais

`func (o *Nssai1) SetDefaultSingleNssais(v []Snssai)`

SetDefaultSingleNssais sets DefaultSingleNssais field to given value.


### GetSingleNssais

`func (o *Nssai1) GetSingleNssais() []Snssai`

GetSingleNssais returns the SingleNssais field if non-nil, zero value otherwise.

### GetSingleNssaisOk

`func (o *Nssai1) GetSingleNssaisOk() (*[]Snssai, bool)`

GetSingleNssaisOk returns a tuple with the SingleNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssais

`func (o *Nssai1) SetSingleNssais(v []Snssai)`

SetSingleNssais sets SingleNssais field to given value.

### HasSingleNssais

`func (o *Nssai1) HasSingleNssais() bool`

HasSingleNssais returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *Nssai1) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *Nssai1) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *Nssai1) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *Nssai1) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.

### GetAdditionalSnssaiData

`func (o *Nssai1) GetAdditionalSnssaiData() map[string]AdditionalSnssaiData`

GetAdditionalSnssaiData returns the AdditionalSnssaiData field if non-nil, zero value otherwise.

### GetAdditionalSnssaiDataOk

`func (o *Nssai1) GetAdditionalSnssaiDataOk() (*map[string]AdditionalSnssaiData, bool)`

GetAdditionalSnssaiDataOk returns a tuple with the AdditionalSnssaiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSnssaiData

`func (o *Nssai1) SetAdditionalSnssaiData(v map[string]AdditionalSnssaiData)`

SetAdditionalSnssaiData sets AdditionalSnssaiData field to given value.

### HasAdditionalSnssaiData

`func (o *Nssai1) HasAdditionalSnssaiData() bool`

HasAdditionalSnssaiData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


