# AllowedSnssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSnssai** | [**Snssai**](Snssai.md) |  | 
**NsiInformationList** | Pointer to [**[]NsiInformation**](NsiInformation.md) |  | [optional] 
**MappedHomeSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewAllowedSnssai

`func NewAllowedSnssai(allowedSnssai Snssai, ) *AllowedSnssai`

NewAllowedSnssai instantiates a new AllowedSnssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedSnssaiWithDefaults

`func NewAllowedSnssaiWithDefaults() *AllowedSnssai`

NewAllowedSnssaiWithDefaults instantiates a new AllowedSnssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSnssai

`func (o *AllowedSnssai) GetAllowedSnssai() Snssai`

GetAllowedSnssai returns the AllowedSnssai field if non-nil, zero value otherwise.

### GetAllowedSnssaiOk

`func (o *AllowedSnssai) GetAllowedSnssaiOk() (*Snssai, bool)`

GetAllowedSnssaiOk returns a tuple with the AllowedSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnssai

`func (o *AllowedSnssai) SetAllowedSnssai(v Snssai)`

SetAllowedSnssai sets AllowedSnssai field to given value.


### GetNsiInformationList

`func (o *AllowedSnssai) GetNsiInformationList() []NsiInformation`

GetNsiInformationList returns the NsiInformationList field if non-nil, zero value otherwise.

### GetNsiInformationListOk

`func (o *AllowedSnssai) GetNsiInformationListOk() (*[]NsiInformation, bool)`

GetNsiInformationListOk returns a tuple with the NsiInformationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiInformationList

`func (o *AllowedSnssai) SetNsiInformationList(v []NsiInformation)`

SetNsiInformationList sets NsiInformationList field to given value.

### HasNsiInformationList

`func (o *AllowedSnssai) HasNsiInformationList() bool`

HasNsiInformationList returns a boolean if a field has been set.

### GetMappedHomeSnssai

`func (o *AllowedSnssai) GetMappedHomeSnssai() Snssai`

GetMappedHomeSnssai returns the MappedHomeSnssai field if non-nil, zero value otherwise.

### GetMappedHomeSnssaiOk

`func (o *AllowedSnssai) GetMappedHomeSnssaiOk() (*Snssai, bool)`

GetMappedHomeSnssaiOk returns a tuple with the MappedHomeSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedHomeSnssai

`func (o *AllowedSnssai) SetMappedHomeSnssai(v Snssai)`

SetMappedHomeSnssai sets MappedHomeSnssai field to given value.

### HasMappedHomeSnssai

`func (o *AllowedSnssai) HasMappedHomeSnssai() bool`

HasMappedHomeSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


