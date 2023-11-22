# SmPolicySnssaiDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**SmPolicyDnnData** | Pointer to [**map[string]SmPolicyDnnDataPatch**](SmPolicyDnnDataPatch.md) |  | [optional] 

## Methods

### NewSmPolicySnssaiDataPatch

`func NewSmPolicySnssaiDataPatch(snssai Snssai, ) *SmPolicySnssaiDataPatch`

NewSmPolicySnssaiDataPatch instantiates a new SmPolicySnssaiDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicySnssaiDataPatchWithDefaults

`func NewSmPolicySnssaiDataPatchWithDefaults() *SmPolicySnssaiDataPatch`

NewSmPolicySnssaiDataPatchWithDefaults instantiates a new SmPolicySnssaiDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *SmPolicySnssaiDataPatch) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmPolicySnssaiDataPatch) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmPolicySnssaiDataPatch) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSmPolicyDnnData

`func (o *SmPolicySnssaiDataPatch) GetSmPolicyDnnData() map[string]SmPolicyDnnDataPatch`

GetSmPolicyDnnData returns the SmPolicyDnnData field if non-nil, zero value otherwise.

### GetSmPolicyDnnDataOk

`func (o *SmPolicySnssaiDataPatch) GetSmPolicyDnnDataOk() (*map[string]SmPolicyDnnDataPatch, bool)`

GetSmPolicyDnnDataOk returns a tuple with the SmPolicyDnnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyDnnData

`func (o *SmPolicySnssaiDataPatch) SetSmPolicyDnnData(v map[string]SmPolicyDnnDataPatch)`

SetSmPolicyDnnData sets SmPolicyDnnData field to given value.

### HasSmPolicyDnnData

`func (o *SmPolicySnssaiDataPatch) HasSmPolicyDnnData() bool`

HasSmPolicyDnnData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


