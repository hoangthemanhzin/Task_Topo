# UePolicySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PraInfos** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**SubscCats** | Pointer to **[]string** |  | [optional] 
**UePolicySections** | Pointer to [**map[string]UePolicySection**](UePolicySection.md) |  | [optional] 
**Upsis** | Pointer to **[]string** |  | [optional] 
**AllowedRouteSelDescs** | Pointer to [**map[string]PlmnRouteSelectionDescriptor**](PlmnRouteSelectionDescriptor.md) |  | [optional] 
**AndspInd** | Pointer to **bool** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**OsIds** | Pointer to **[]string** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewUePolicySet

`func NewUePolicySet() *UePolicySet`

NewUePolicySet instantiates a new UePolicySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUePolicySetWithDefaults

`func NewUePolicySetWithDefaults() *UePolicySet`

NewUePolicySetWithDefaults instantiates a new UePolicySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPraInfos

`func (o *UePolicySet) GetPraInfos() map[string]PresenceInfo`

GetPraInfos returns the PraInfos field if non-nil, zero value otherwise.

### GetPraInfosOk

`func (o *UePolicySet) GetPraInfosOk() (*map[string]PresenceInfo, bool)`

GetPraInfosOk returns a tuple with the PraInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInfos

`func (o *UePolicySet) SetPraInfos(v map[string]PresenceInfo)`

SetPraInfos sets PraInfos field to given value.

### HasPraInfos

`func (o *UePolicySet) HasPraInfos() bool`

HasPraInfos returns a boolean if a field has been set.

### GetSubscCats

`func (o *UePolicySet) GetSubscCats() []string`

GetSubscCats returns the SubscCats field if non-nil, zero value otherwise.

### GetSubscCatsOk

`func (o *UePolicySet) GetSubscCatsOk() (*[]string, bool)`

GetSubscCatsOk returns a tuple with the SubscCats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscCats

`func (o *UePolicySet) SetSubscCats(v []string)`

SetSubscCats sets SubscCats field to given value.

### HasSubscCats

`func (o *UePolicySet) HasSubscCats() bool`

HasSubscCats returns a boolean if a field has been set.

### GetUePolicySections

`func (o *UePolicySet) GetUePolicySections() map[string]UePolicySection`

GetUePolicySections returns the UePolicySections field if non-nil, zero value otherwise.

### GetUePolicySectionsOk

`func (o *UePolicySet) GetUePolicySectionsOk() (*map[string]UePolicySection, bool)`

GetUePolicySectionsOk returns a tuple with the UePolicySections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicySections

`func (o *UePolicySet) SetUePolicySections(v map[string]UePolicySection)`

SetUePolicySections sets UePolicySections field to given value.

### HasUePolicySections

`func (o *UePolicySet) HasUePolicySections() bool`

HasUePolicySections returns a boolean if a field has been set.

### GetUpsis

`func (o *UePolicySet) GetUpsis() []string`

GetUpsis returns the Upsis field if non-nil, zero value otherwise.

### GetUpsisOk

`func (o *UePolicySet) GetUpsisOk() (*[]string, bool)`

GetUpsisOk returns a tuple with the Upsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsis

`func (o *UePolicySet) SetUpsis(v []string)`

SetUpsis sets Upsis field to given value.

### HasUpsis

`func (o *UePolicySet) HasUpsis() bool`

HasUpsis returns a boolean if a field has been set.

### GetAllowedRouteSelDescs

`func (o *UePolicySet) GetAllowedRouteSelDescs() map[string]PlmnRouteSelectionDescriptor`

GetAllowedRouteSelDescs returns the AllowedRouteSelDescs field if non-nil, zero value otherwise.

### GetAllowedRouteSelDescsOk

`func (o *UePolicySet) GetAllowedRouteSelDescsOk() (*map[string]PlmnRouteSelectionDescriptor, bool)`

GetAllowedRouteSelDescsOk returns a tuple with the AllowedRouteSelDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRouteSelDescs

`func (o *UePolicySet) SetAllowedRouteSelDescs(v map[string]PlmnRouteSelectionDescriptor)`

SetAllowedRouteSelDescs sets AllowedRouteSelDescs field to given value.

### HasAllowedRouteSelDescs

`func (o *UePolicySet) HasAllowedRouteSelDescs() bool`

HasAllowedRouteSelDescs returns a boolean if a field has been set.

### GetAndspInd

`func (o *UePolicySet) GetAndspInd() bool`

GetAndspInd returns the AndspInd field if non-nil, zero value otherwise.

### GetAndspIndOk

`func (o *UePolicySet) GetAndspIndOk() (*bool, bool)`

GetAndspIndOk returns a tuple with the AndspInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndspInd

`func (o *UePolicySet) SetAndspInd(v bool)`

SetAndspInd sets AndspInd field to given value.

### HasAndspInd

`func (o *UePolicySet) HasAndspInd() bool`

HasAndspInd returns a boolean if a field has been set.

### GetPei

`func (o *UePolicySet) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *UePolicySet) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *UePolicySet) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *UePolicySet) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetOsIds

`func (o *UePolicySet) GetOsIds() []string`

GetOsIds returns the OsIds field if non-nil, zero value otherwise.

### GetOsIdsOk

`func (o *UePolicySet) GetOsIdsOk() (*[]string, bool)`

GetOsIdsOk returns a tuple with the OsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsIds

`func (o *UePolicySet) SetOsIds(v []string)`

SetOsIds sets OsIds field to given value.

### HasOsIds

`func (o *UePolicySet) HasOsIds() bool`

HasOsIds returns a boolean if a field has been set.

### GetSuppFeat

`func (o *UePolicySet) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *UePolicySet) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *UePolicySet) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *UePolicySet) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


