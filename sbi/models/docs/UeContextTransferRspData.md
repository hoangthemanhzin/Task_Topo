# UeContextTransferRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeContext** | [**UeContext**](UeContext.md) |  | 
**UeRadioCapability** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**UeNbiotRadioCapability** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewUeContextTransferRspData

`func NewUeContextTransferRspData(ueContext UeContext, ) *UeContextTransferRspData`

NewUeContextTransferRspData instantiates a new UeContextTransferRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextTransferRspDataWithDefaults

`func NewUeContextTransferRspDataWithDefaults() *UeContextTransferRspData`

NewUeContextTransferRspDataWithDefaults instantiates a new UeContextTransferRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeContext

`func (o *UeContextTransferRspData) GetUeContext() UeContext`

GetUeContext returns the UeContext field if non-nil, zero value otherwise.

### GetUeContextOk

`func (o *UeContextTransferRspData) GetUeContextOk() (*UeContext, bool)`

GetUeContextOk returns a tuple with the UeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContext

`func (o *UeContextTransferRspData) SetUeContext(v UeContext)`

SetUeContext sets UeContext field to given value.


### GetUeRadioCapability

`func (o *UeContextTransferRspData) GetUeRadioCapability() N2InfoContent`

GetUeRadioCapability returns the UeRadioCapability field if non-nil, zero value otherwise.

### GetUeRadioCapabilityOk

`func (o *UeContextTransferRspData) GetUeRadioCapabilityOk() (*N2InfoContent, bool)`

GetUeRadioCapabilityOk returns a tuple with the UeRadioCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability

`func (o *UeContextTransferRspData) SetUeRadioCapability(v N2InfoContent)`

SetUeRadioCapability sets UeRadioCapability field to given value.

### HasUeRadioCapability

`func (o *UeContextTransferRspData) HasUeRadioCapability() bool`

HasUeRadioCapability returns a boolean if a field has been set.

### GetUeNbiotRadioCapability

`func (o *UeContextTransferRspData) GetUeNbiotRadioCapability() N2InfoContent`

GetUeNbiotRadioCapability returns the UeNbiotRadioCapability field if non-nil, zero value otherwise.

### GetUeNbiotRadioCapabilityOk

`func (o *UeContextTransferRspData) GetUeNbiotRadioCapabilityOk() (*N2InfoContent, bool)`

GetUeNbiotRadioCapabilityOk returns a tuple with the UeNbiotRadioCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeNbiotRadioCapability

`func (o *UeContextTransferRspData) SetUeNbiotRadioCapability(v N2InfoContent)`

SetUeNbiotRadioCapability sets UeNbiotRadioCapability field to given value.

### HasUeNbiotRadioCapability

`func (o *UeContextTransferRspData) HasUeNbiotRadioCapability() bool`

HasUeNbiotRadioCapability returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextTransferRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextTransferRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextTransferRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextTransferRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


