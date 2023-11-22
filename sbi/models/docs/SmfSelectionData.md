# SmfSelectionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsuppDnn** | Pointer to **bool** |  | [optional] 
**Candidates** | Pointer to [**map[string]CandidateForReplacement**](CandidateForReplacement.md) |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**MappingSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 

## Methods

### NewSmfSelectionData

`func NewSmfSelectionData() *SmfSelectionData`

NewSmfSelectionData instantiates a new SmfSelectionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfSelectionDataWithDefaults

`func NewSmfSelectionDataWithDefaults() *SmfSelectionData`

NewSmfSelectionDataWithDefaults instantiates a new SmfSelectionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsuppDnn

`func (o *SmfSelectionData) GetUnsuppDnn() bool`

GetUnsuppDnn returns the UnsuppDnn field if non-nil, zero value otherwise.

### GetUnsuppDnnOk

`func (o *SmfSelectionData) GetUnsuppDnnOk() (*bool, bool)`

GetUnsuppDnnOk returns a tuple with the UnsuppDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsuppDnn

`func (o *SmfSelectionData) SetUnsuppDnn(v bool)`

SetUnsuppDnn sets UnsuppDnn field to given value.

### HasUnsuppDnn

`func (o *SmfSelectionData) HasUnsuppDnn() bool`

HasUnsuppDnn returns a boolean if a field has been set.

### GetCandidates

`func (o *SmfSelectionData) GetCandidates() map[string]CandidateForReplacement`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *SmfSelectionData) GetCandidatesOk() (*map[string]CandidateForReplacement, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *SmfSelectionData) SetCandidates(v map[string]CandidateForReplacement)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *SmfSelectionData) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### SetCandidatesNil

`func (o *SmfSelectionData) SetCandidatesNil(b bool)`

 SetCandidatesNil sets the value for Candidates to be an explicit nil

### UnsetCandidates
`func (o *SmfSelectionData) UnsetCandidates()`

UnsetCandidates ensures that no value is present for Candidates, not even an explicit nil
### GetSnssai

`func (o *SmfSelectionData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmfSelectionData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmfSelectionData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *SmfSelectionData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetMappingSnssai

`func (o *SmfSelectionData) GetMappingSnssai() Snssai`

GetMappingSnssai returns the MappingSnssai field if non-nil, zero value otherwise.

### GetMappingSnssaiOk

`func (o *SmfSelectionData) GetMappingSnssaiOk() (*Snssai, bool)`

GetMappingSnssaiOk returns a tuple with the MappingSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingSnssai

`func (o *SmfSelectionData) SetMappingSnssai(v Snssai)`

SetMappingSnssai sets MappingSnssai field to given value.

### HasMappingSnssai

`func (o *SmfSelectionData) HasMappingSnssai() bool`

HasMappingSnssai returns a boolean if a field has been set.

### GetDnn

`func (o *SmfSelectionData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmfSelectionData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmfSelectionData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmfSelectionData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


