# UsageMonDataScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUsageMonDataScope

`func NewUsageMonDataScope(snssai Snssai, ) *UsageMonDataScope`

NewUsageMonDataScope instantiates a new UsageMonDataScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMonDataScopeWithDefaults

`func NewUsageMonDataScopeWithDefaults() *UsageMonDataScope`

NewUsageMonDataScopeWithDefaults instantiates a new UsageMonDataScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *UsageMonDataScope) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *UsageMonDataScope) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *UsageMonDataScope) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnn

`func (o *UsageMonDataScope) GetDnn() []string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *UsageMonDataScope) GetDnnOk() (*[]string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *UsageMonDataScope) SetDnn(v []string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *UsageMonDataScope) HasDnn() bool`

HasDnn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


