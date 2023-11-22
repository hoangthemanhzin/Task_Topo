# AmPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PraInfos** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**SubscCats** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAmPolicyData

`func NewAmPolicyData() *AmPolicyData`

NewAmPolicyData instantiates a new AmPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmPolicyDataWithDefaults

`func NewAmPolicyDataWithDefaults() *AmPolicyData`

NewAmPolicyDataWithDefaults instantiates a new AmPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPraInfos

`func (o *AmPolicyData) GetPraInfos() map[string]PresenceInfo`

GetPraInfos returns the PraInfos field if non-nil, zero value otherwise.

### GetPraInfosOk

`func (o *AmPolicyData) GetPraInfosOk() (*map[string]PresenceInfo, bool)`

GetPraInfosOk returns a tuple with the PraInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInfos

`func (o *AmPolicyData) SetPraInfos(v map[string]PresenceInfo)`

SetPraInfos sets PraInfos field to given value.

### HasPraInfos

`func (o *AmPolicyData) HasPraInfos() bool`

HasPraInfos returns a boolean if a field has been set.

### GetSubscCats

`func (o *AmPolicyData) GetSubscCats() []string`

GetSubscCats returns the SubscCats field if non-nil, zero value otherwise.

### GetSubscCatsOk

`func (o *AmPolicyData) GetSubscCatsOk() (*[]string, bool)`

GetSubscCatsOk returns a tuple with the SubscCats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscCats

`func (o *AmPolicyData) SetSubscCats(v []string)`

SetSubscCats sets SubscCats field to given value.

### HasSubscCats

`func (o *AmPolicyData) HasSubscCats() bool`

HasSubscCats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


