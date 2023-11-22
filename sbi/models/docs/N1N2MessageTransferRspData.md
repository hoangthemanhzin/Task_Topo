# N1N2MessageTransferRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | [**N1N2MessageTransferCause**](N1N2MessageTransferCause.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewN1N2MessageTransferRspData

`func NewN1N2MessageTransferRspData(cause N1N2MessageTransferCause, ) *N1N2MessageTransferRspData`

NewN1N2MessageTransferRspData instantiates a new N1N2MessageTransferRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1N2MessageTransferRspDataWithDefaults

`func NewN1N2MessageTransferRspDataWithDefaults() *N1N2MessageTransferRspData`

NewN1N2MessageTransferRspDataWithDefaults instantiates a new N1N2MessageTransferRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *N1N2MessageTransferRspData) GetCause() N1N2MessageTransferCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *N1N2MessageTransferRspData) GetCauseOk() (*N1N2MessageTransferCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *N1N2MessageTransferRspData) SetCause(v N1N2MessageTransferCause)`

SetCause sets Cause field to given value.


### GetSupportedFeatures

`func (o *N1N2MessageTransferRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *N1N2MessageTransferRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *N1N2MessageTransferRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *N1N2MessageTransferRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


