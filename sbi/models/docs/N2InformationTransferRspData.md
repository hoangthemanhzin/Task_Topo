# N2InformationTransferRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**N2InformationTransferResult**](N2InformationTransferResult.md) |  | 
**PwsRspData** | Pointer to [**PWSResponseData**](PWSResponseData.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewN2InformationTransferRspData

`func NewN2InformationTransferRspData(result N2InformationTransferResult, ) *N2InformationTransferRspData`

NewN2InformationTransferRspData instantiates a new N2InformationTransferRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InformationTransferRspDataWithDefaults

`func NewN2InformationTransferRspDataWithDefaults() *N2InformationTransferRspData`

NewN2InformationTransferRspDataWithDefaults instantiates a new N2InformationTransferRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *N2InformationTransferRspData) GetResult() N2InformationTransferResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *N2InformationTransferRspData) GetResultOk() (*N2InformationTransferResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *N2InformationTransferRspData) SetResult(v N2InformationTransferResult)`

SetResult sets Result field to given value.


### GetPwsRspData

`func (o *N2InformationTransferRspData) GetPwsRspData() PWSResponseData`

GetPwsRspData returns the PwsRspData field if non-nil, zero value otherwise.

### GetPwsRspDataOk

`func (o *N2InformationTransferRspData) GetPwsRspDataOk() (*PWSResponseData, bool)`

GetPwsRspDataOk returns a tuple with the PwsRspData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwsRspData

`func (o *N2InformationTransferRspData) SetPwsRspData(v PWSResponseData)`

SetPwsRspData sets PwsRspData field to given value.

### HasPwsRspData

`func (o *N2InformationTransferRspData) HasPwsRspData() bool`

HasPwsRspData returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *N2InformationTransferRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *N2InformationTransferRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *N2InformationTransferRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *N2InformationTransferRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


