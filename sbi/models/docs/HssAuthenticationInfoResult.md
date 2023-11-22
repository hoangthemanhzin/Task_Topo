# HssAuthenticationInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**HssAuthenticationVectors** | [**HssAuthenticationVectors**](HssAuthenticationVectors.md) |  | 

## Methods

### NewHssAuthenticationInfoResult

`func NewHssAuthenticationInfoResult(hssAuthenticationVectors HssAuthenticationVectors, ) *HssAuthenticationInfoResult`

NewHssAuthenticationInfoResult instantiates a new HssAuthenticationInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHssAuthenticationInfoResultWithDefaults

`func NewHssAuthenticationInfoResultWithDefaults() *HssAuthenticationInfoResult`

NewHssAuthenticationInfoResultWithDefaults instantiates a new HssAuthenticationInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *HssAuthenticationInfoResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *HssAuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *HssAuthenticationInfoResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *HssAuthenticationInfoResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetHssAuthenticationVectors

`func (o *HssAuthenticationInfoResult) GetHssAuthenticationVectors() HssAuthenticationVectors`

GetHssAuthenticationVectors returns the HssAuthenticationVectors field if non-nil, zero value otherwise.

### GetHssAuthenticationVectorsOk

`func (o *HssAuthenticationInfoResult) GetHssAuthenticationVectorsOk() (*HssAuthenticationVectors, bool)`

GetHssAuthenticationVectorsOk returns a tuple with the HssAuthenticationVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssAuthenticationVectors

`func (o *HssAuthenticationInfoResult) SetHssAuthenticationVectors(v HssAuthenticationVectors)`

SetHssAuthenticationVectors sets HssAuthenticationVectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


