# ConfirmationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResStar** | **string** |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewConfirmationData

`func NewConfirmationData(resStar string, ) *ConfirmationData`

NewConfirmationData instantiates a new ConfirmationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmationDataWithDefaults

`func NewConfirmationDataWithDefaults() *ConfirmationData`

NewConfirmationDataWithDefaults instantiates a new ConfirmationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResStar

`func (o *ConfirmationData) GetResStar() string`

GetResStar returns the ResStar field if non-nil, zero value otherwise.

### GetResStarOk

`func (o *ConfirmationData) GetResStarOk() (*string, bool)`

GetResStarOk returns a tuple with the ResStar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResStar

`func (o *ConfirmationData) SetResStar(v string)`

SetResStar sets ResStar field to given value.


### SetResStarNil

`func (o *ConfirmationData) SetResStarNil(b bool)`

 SetResStarNil sets the value for ResStar to be an explicit nil

### UnsetResStar
`func (o *ConfirmationData) UnsetResStar()`

UnsetResStar ensures that no value is present for ResStar, not even an explicit nil
### GetSupportedFeatures

`func (o *ConfirmationData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ConfirmationData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ConfirmationData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ConfirmationData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


