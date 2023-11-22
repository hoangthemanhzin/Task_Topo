# AreaScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EutraCellIdList** | Pointer to **[]string** |  | [optional] 
**NrCellIdList** | Pointer to **[]string** |  | [optional] 
**TacList** | Pointer to **[]string** |  | [optional] 
**TacInfoPerPlmn** | Pointer to [**map[string]TacInfo**](TacInfo.md) |  | [optional] 

## Methods

### NewAreaScope

`func NewAreaScope() *AreaScope`

NewAreaScope instantiates a new AreaScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaScopeWithDefaults

`func NewAreaScopeWithDefaults() *AreaScope`

NewAreaScopeWithDefaults instantiates a new AreaScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEutraCellIdList

`func (o *AreaScope) GetEutraCellIdList() []string`

GetEutraCellIdList returns the EutraCellIdList field if non-nil, zero value otherwise.

### GetEutraCellIdListOk

`func (o *AreaScope) GetEutraCellIdListOk() (*[]string, bool)`

GetEutraCellIdListOk returns a tuple with the EutraCellIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEutraCellIdList

`func (o *AreaScope) SetEutraCellIdList(v []string)`

SetEutraCellIdList sets EutraCellIdList field to given value.

### HasEutraCellIdList

`func (o *AreaScope) HasEutraCellIdList() bool`

HasEutraCellIdList returns a boolean if a field has been set.

### GetNrCellIdList

`func (o *AreaScope) GetNrCellIdList() []string`

GetNrCellIdList returns the NrCellIdList field if non-nil, zero value otherwise.

### GetNrCellIdListOk

`func (o *AreaScope) GetNrCellIdListOk() (*[]string, bool)`

GetNrCellIdListOk returns a tuple with the NrCellIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellIdList

`func (o *AreaScope) SetNrCellIdList(v []string)`

SetNrCellIdList sets NrCellIdList field to given value.

### HasNrCellIdList

`func (o *AreaScope) HasNrCellIdList() bool`

HasNrCellIdList returns a boolean if a field has been set.

### GetTacList

`func (o *AreaScope) GetTacList() []string`

GetTacList returns the TacList field if non-nil, zero value otherwise.

### GetTacListOk

`func (o *AreaScope) GetTacListOk() (*[]string, bool)`

GetTacListOk returns a tuple with the TacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacList

`func (o *AreaScope) SetTacList(v []string)`

SetTacList sets TacList field to given value.

### HasTacList

`func (o *AreaScope) HasTacList() bool`

HasTacList returns a boolean if a field has been set.

### GetTacInfoPerPlmn

`func (o *AreaScope) GetTacInfoPerPlmn() map[string]TacInfo`

GetTacInfoPerPlmn returns the TacInfoPerPlmn field if non-nil, zero value otherwise.

### GetTacInfoPerPlmnOk

`func (o *AreaScope) GetTacInfoPerPlmnOk() (*map[string]TacInfo, bool)`

GetTacInfoPerPlmnOk returns a tuple with the TacInfoPerPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacInfoPerPlmn

`func (o *AreaScope) SetTacInfoPerPlmn(v map[string]TacInfo)`

SetTacInfoPerPlmn sets TacInfoPerPlmn field to given value.

### HasTacInfoPerPlmn

`func (o *AreaScope) HasTacInfoPerPlmn() bool`

HasTacInfoPerPlmn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


