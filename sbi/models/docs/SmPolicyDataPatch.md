# SmPolicyDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UmData** | Pointer to [**map[string]UsageMonData**](UsageMonData.md) |  | [optional] 
**SmPolicySnssaiData** | Pointer to [**map[string]SmPolicySnssaiDataPatch**](SmPolicySnssaiDataPatch.md) |  | [optional] 

## Methods

### NewSmPolicyDataPatch

`func NewSmPolicyDataPatch() *SmPolicyDataPatch`

NewSmPolicyDataPatch instantiates a new SmPolicyDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDataPatchWithDefaults

`func NewSmPolicyDataPatchWithDefaults() *SmPolicyDataPatch`

NewSmPolicyDataPatchWithDefaults instantiates a new SmPolicyDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUmData

`func (o *SmPolicyDataPatch) GetUmData() map[string]UsageMonData`

GetUmData returns the UmData field if non-nil, zero value otherwise.

### GetUmDataOk

`func (o *SmPolicyDataPatch) GetUmDataOk() (*map[string]UsageMonData, bool)`

GetUmDataOk returns a tuple with the UmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmData

`func (o *SmPolicyDataPatch) SetUmData(v map[string]UsageMonData)`

SetUmData sets UmData field to given value.

### HasUmData

`func (o *SmPolicyDataPatch) HasUmData() bool`

HasUmData returns a boolean if a field has been set.

### SetUmDataNil

`func (o *SmPolicyDataPatch) SetUmDataNil(b bool)`

 SetUmDataNil sets the value for UmData to be an explicit nil

### UnsetUmData
`func (o *SmPolicyDataPatch) UnsetUmData()`

UnsetUmData ensures that no value is present for UmData, not even an explicit nil
### GetSmPolicySnssaiData

`func (o *SmPolicyDataPatch) GetSmPolicySnssaiData() map[string]SmPolicySnssaiDataPatch`

GetSmPolicySnssaiData returns the SmPolicySnssaiData field if non-nil, zero value otherwise.

### GetSmPolicySnssaiDataOk

`func (o *SmPolicyDataPatch) GetSmPolicySnssaiDataOk() (*map[string]SmPolicySnssaiDataPatch, bool)`

GetSmPolicySnssaiDataOk returns a tuple with the SmPolicySnssaiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicySnssaiData

`func (o *SmPolicyDataPatch) SetSmPolicySnssaiData(v map[string]SmPolicySnssaiDataPatch)`

SetSmPolicySnssaiData sets SmPolicySnssaiData field to given value.

### HasSmPolicySnssaiData

`func (o *SmPolicyDataPatch) HasSmPolicySnssaiData() bool`

HasSmPolicySnssaiData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


