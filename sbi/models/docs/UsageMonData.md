# UsageMonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitId** | **string** |  | 
**Scopes** | Pointer to [**map[string]UsageMonDataScope**](UsageMonDataScope.md) |  | [optional] 
**UmLevel** | Pointer to [**UsageMonLevel**](UsageMonLevel.md) |  | [optional] 
**AllowedUsage** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**ResetTime** | Pointer to **time.Time** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewUsageMonData

`func NewUsageMonData(limitId string, ) *UsageMonData`

NewUsageMonData instantiates a new UsageMonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMonDataWithDefaults

`func NewUsageMonDataWithDefaults() *UsageMonData`

NewUsageMonDataWithDefaults instantiates a new UsageMonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitId

`func (o *UsageMonData) GetLimitId() string`

GetLimitId returns the LimitId field if non-nil, zero value otherwise.

### GetLimitIdOk

`func (o *UsageMonData) GetLimitIdOk() (*string, bool)`

GetLimitIdOk returns a tuple with the LimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitId

`func (o *UsageMonData) SetLimitId(v string)`

SetLimitId sets LimitId field to given value.


### GetScopes

`func (o *UsageMonData) GetScopes() map[string]UsageMonDataScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UsageMonData) GetScopesOk() (*map[string]UsageMonDataScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UsageMonData) SetScopes(v map[string]UsageMonDataScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UsageMonData) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUmLevel

`func (o *UsageMonData) GetUmLevel() UsageMonLevel`

GetUmLevel returns the UmLevel field if non-nil, zero value otherwise.

### GetUmLevelOk

`func (o *UsageMonData) GetUmLevelOk() (*UsageMonLevel, bool)`

GetUmLevelOk returns a tuple with the UmLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmLevel

`func (o *UsageMonData) SetUmLevel(v UsageMonLevel)`

SetUmLevel sets UmLevel field to given value.

### HasUmLevel

`func (o *UsageMonData) HasUmLevel() bool`

HasUmLevel returns a boolean if a field has been set.

### GetAllowedUsage

`func (o *UsageMonData) GetAllowedUsage() UsageThreshold`

GetAllowedUsage returns the AllowedUsage field if non-nil, zero value otherwise.

### GetAllowedUsageOk

`func (o *UsageMonData) GetAllowedUsageOk() (*UsageThreshold, bool)`

GetAllowedUsageOk returns a tuple with the AllowedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsage

`func (o *UsageMonData) SetAllowedUsage(v UsageThreshold)`

SetAllowedUsage sets AllowedUsage field to given value.

### HasAllowedUsage

`func (o *UsageMonData) HasAllowedUsage() bool`

HasAllowedUsage returns a boolean if a field has been set.

### GetResetTime

`func (o *UsageMonData) GetResetTime() time.Time`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *UsageMonData) GetResetTimeOk() (*time.Time, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *UsageMonData) SetResetTime(v time.Time)`

SetResetTime sets ResetTime field to given value.

### HasResetTime

`func (o *UsageMonData) HasResetTime() bool`

HasResetTime returns a boolean if a field has been set.

### GetSuppFeat

`func (o *UsageMonData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *UsageMonData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *UsageMonData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *UsageMonData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


