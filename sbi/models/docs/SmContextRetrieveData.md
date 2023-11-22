# SmContextRetrieveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetMmeCap** | Pointer to [**MmeCapabilities**](MmeCapabilities.md) |  | [optional] 
**SmContextType** | Pointer to [**SmContextType**](SmContextType.md) |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**NotToTransferEbiList** | Pointer to **[]int32** |  | [optional] 
**RanUnchangedInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSmContextRetrieveData

`func NewSmContextRetrieveData() *SmContextRetrieveData`

NewSmContextRetrieveData instantiates a new SmContextRetrieveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextRetrieveDataWithDefaults

`func NewSmContextRetrieveDataWithDefaults() *SmContextRetrieveData`

NewSmContextRetrieveDataWithDefaults instantiates a new SmContextRetrieveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetMmeCap

`func (o *SmContextRetrieveData) GetTargetMmeCap() MmeCapabilities`

GetTargetMmeCap returns the TargetMmeCap field if non-nil, zero value otherwise.

### GetTargetMmeCapOk

`func (o *SmContextRetrieveData) GetTargetMmeCapOk() (*MmeCapabilities, bool)`

GetTargetMmeCapOk returns a tuple with the TargetMmeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMmeCap

`func (o *SmContextRetrieveData) SetTargetMmeCap(v MmeCapabilities)`

SetTargetMmeCap sets TargetMmeCap field to given value.

### HasTargetMmeCap

`func (o *SmContextRetrieveData) HasTargetMmeCap() bool`

HasTargetMmeCap returns a boolean if a field has been set.

### GetSmContextType

`func (o *SmContextRetrieveData) GetSmContextType() SmContextType`

GetSmContextType returns the SmContextType field if non-nil, zero value otherwise.

### GetSmContextTypeOk

`func (o *SmContextRetrieveData) GetSmContextTypeOk() (*SmContextType, bool)`

GetSmContextTypeOk returns a tuple with the SmContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextType

`func (o *SmContextRetrieveData) SetSmContextType(v SmContextType)`

SetSmContextType sets SmContextType field to given value.

### HasSmContextType

`func (o *SmContextRetrieveData) HasSmContextType() bool`

HasSmContextType returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmContextRetrieveData) GetServingNetwork() PlmnId`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmContextRetrieveData) GetServingNetworkOk() (*PlmnId, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmContextRetrieveData) SetServingNetwork(v PlmnId)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *SmContextRetrieveData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetNotToTransferEbiList

`func (o *SmContextRetrieveData) GetNotToTransferEbiList() []int32`

GetNotToTransferEbiList returns the NotToTransferEbiList field if non-nil, zero value otherwise.

### GetNotToTransferEbiListOk

`func (o *SmContextRetrieveData) GetNotToTransferEbiListOk() (*[]int32, bool)`

GetNotToTransferEbiListOk returns a tuple with the NotToTransferEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotToTransferEbiList

`func (o *SmContextRetrieveData) SetNotToTransferEbiList(v []int32)`

SetNotToTransferEbiList sets NotToTransferEbiList field to given value.

### HasNotToTransferEbiList

`func (o *SmContextRetrieveData) HasNotToTransferEbiList() bool`

HasNotToTransferEbiList returns a boolean if a field has been set.

### GetRanUnchangedInd

`func (o *SmContextRetrieveData) GetRanUnchangedInd() bool`

GetRanUnchangedInd returns the RanUnchangedInd field if non-nil, zero value otherwise.

### GetRanUnchangedIndOk

`func (o *SmContextRetrieveData) GetRanUnchangedIndOk() (*bool, bool)`

GetRanUnchangedIndOk returns a tuple with the RanUnchangedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUnchangedInd

`func (o *SmContextRetrieveData) SetRanUnchangedInd(v bool)`

SetRanUnchangedInd sets RanUnchangedInd field to given value.

### HasRanUnchangedInd

`func (o *SmContextRetrieveData) HasRanUnchangedInd() bool`

HasRanUnchangedInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


