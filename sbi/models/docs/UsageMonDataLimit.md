# UsageMonDataLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitId** | **string** |  | 
**Scopes** | Pointer to [**map[string]UsageMonDataScope**](UsageMonDataScope.md) |  | [optional] 
**UmLevel** | Pointer to [**UsageMonLevel**](UsageMonLevel.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**UsageLimit** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**ResetPeriod** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewUsageMonDataLimit

`func NewUsageMonDataLimit(limitId string, ) *UsageMonDataLimit`

NewUsageMonDataLimit instantiates a new UsageMonDataLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMonDataLimitWithDefaults

`func NewUsageMonDataLimitWithDefaults() *UsageMonDataLimit`

NewUsageMonDataLimitWithDefaults instantiates a new UsageMonDataLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitId

`func (o *UsageMonDataLimit) GetLimitId() string`

GetLimitId returns the LimitId field if non-nil, zero value otherwise.

### GetLimitIdOk

`func (o *UsageMonDataLimit) GetLimitIdOk() (*string, bool)`

GetLimitIdOk returns a tuple with the LimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitId

`func (o *UsageMonDataLimit) SetLimitId(v string)`

SetLimitId sets LimitId field to given value.


### GetScopes

`func (o *UsageMonDataLimit) GetScopes() map[string]UsageMonDataScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UsageMonDataLimit) GetScopesOk() (*map[string]UsageMonDataScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UsageMonDataLimit) SetScopes(v map[string]UsageMonDataScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UsageMonDataLimit) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUmLevel

`func (o *UsageMonDataLimit) GetUmLevel() UsageMonLevel`

GetUmLevel returns the UmLevel field if non-nil, zero value otherwise.

### GetUmLevelOk

`func (o *UsageMonDataLimit) GetUmLevelOk() (*UsageMonLevel, bool)`

GetUmLevelOk returns a tuple with the UmLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmLevel

`func (o *UsageMonDataLimit) SetUmLevel(v UsageMonLevel)`

SetUmLevel sets UmLevel field to given value.

### HasUmLevel

`func (o *UsageMonDataLimit) HasUmLevel() bool`

HasUmLevel returns a boolean if a field has been set.

### GetStartDate

`func (o *UsageMonDataLimit) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageMonDataLimit) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageMonDataLimit) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageMonDataLimit) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageMonDataLimit) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageMonDataLimit) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageMonDataLimit) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageMonDataLimit) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *UsageMonDataLimit) GetUsageLimit() UsageThreshold`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UsageMonDataLimit) GetUsageLimitOk() (*UsageThreshold, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *UsageMonDataLimit) SetUsageLimit(v UsageThreshold)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *UsageMonDataLimit) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetResetPeriod

`func (o *UsageMonDataLimit) GetResetPeriod() TimePeriod`

GetResetPeriod returns the ResetPeriod field if non-nil, zero value otherwise.

### GetResetPeriodOk

`func (o *UsageMonDataLimit) GetResetPeriodOk() (*TimePeriod, bool)`

GetResetPeriodOk returns a tuple with the ResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPeriod

`func (o *UsageMonDataLimit) SetResetPeriod(v TimePeriod)`

SetResetPeriod sets ResetPeriod field to given value.

### HasResetPeriod

`func (o *UsageMonDataLimit) HasResetPeriod() bool`

HasResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


