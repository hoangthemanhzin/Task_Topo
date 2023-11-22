# MediaSubComponentRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfSigProtocol** | Pointer to [**AfSigProtocol**](AfSigProtocol.md) |  | [optional] 
**EthfDescs** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**FNum** | **int32** |  | 
**FDescs** | Pointer to **[]string** |  | [optional] 
**FStatus** | Pointer to [**FlowStatus**](FlowStatus.md) |  | [optional] 
**MarBwDl** | Pointer to **string** |  | [optional] 
**MarBwUl** | Pointer to **string** |  | [optional] 
**TosTrCl** | Pointer to **string** | this data type is defined in the same way as the TosTrafficClass data type, but with the OpenAPI nullable property set to true | [optional] 
**FlowUsage** | Pointer to [**FlowUsage**](FlowUsage.md) |  | [optional] 

## Methods

### NewMediaSubComponentRm

`func NewMediaSubComponentRm(fNum int32, ) *MediaSubComponentRm`

NewMediaSubComponentRm instantiates a new MediaSubComponentRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaSubComponentRmWithDefaults

`func NewMediaSubComponentRmWithDefaults() *MediaSubComponentRm`

NewMediaSubComponentRmWithDefaults instantiates a new MediaSubComponentRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfSigProtocol

`func (o *MediaSubComponentRm) GetAfSigProtocol() AfSigProtocol`

GetAfSigProtocol returns the AfSigProtocol field if non-nil, zero value otherwise.

### GetAfSigProtocolOk

`func (o *MediaSubComponentRm) GetAfSigProtocolOk() (*AfSigProtocol, bool)`

GetAfSigProtocolOk returns a tuple with the AfSigProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfSigProtocol

`func (o *MediaSubComponentRm) SetAfSigProtocol(v AfSigProtocol)`

SetAfSigProtocol sets AfSigProtocol field to given value.

### HasAfSigProtocol

`func (o *MediaSubComponentRm) HasAfSigProtocol() bool`

HasAfSigProtocol returns a boolean if a field has been set.

### GetEthfDescs

`func (o *MediaSubComponentRm) GetEthfDescs() []EthFlowDescription`

GetEthfDescs returns the EthfDescs field if non-nil, zero value otherwise.

### GetEthfDescsOk

`func (o *MediaSubComponentRm) GetEthfDescsOk() (*[]EthFlowDescription, bool)`

GetEthfDescsOk returns a tuple with the EthfDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthfDescs

`func (o *MediaSubComponentRm) SetEthfDescs(v []EthFlowDescription)`

SetEthfDescs sets EthfDescs field to given value.

### HasEthfDescs

`func (o *MediaSubComponentRm) HasEthfDescs() bool`

HasEthfDescs returns a boolean if a field has been set.

### SetEthfDescsNil

`func (o *MediaSubComponentRm) SetEthfDescsNil(b bool)`

 SetEthfDescsNil sets the value for EthfDescs to be an explicit nil

### UnsetEthfDescs
`func (o *MediaSubComponentRm) UnsetEthfDescs()`

UnsetEthfDescs ensures that no value is present for EthfDescs, not even an explicit nil
### GetFNum

`func (o *MediaSubComponentRm) GetFNum() int32`

GetFNum returns the FNum field if non-nil, zero value otherwise.

### GetFNumOk

`func (o *MediaSubComponentRm) GetFNumOk() (*int32, bool)`

GetFNumOk returns a tuple with the FNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFNum

`func (o *MediaSubComponentRm) SetFNum(v int32)`

SetFNum sets FNum field to given value.


### GetFDescs

`func (o *MediaSubComponentRm) GetFDescs() []string`

GetFDescs returns the FDescs field if non-nil, zero value otherwise.

### GetFDescsOk

`func (o *MediaSubComponentRm) GetFDescsOk() (*[]string, bool)`

GetFDescsOk returns a tuple with the FDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDescs

`func (o *MediaSubComponentRm) SetFDescs(v []string)`

SetFDescs sets FDescs field to given value.

### HasFDescs

`func (o *MediaSubComponentRm) HasFDescs() bool`

HasFDescs returns a boolean if a field has been set.

### SetFDescsNil

`func (o *MediaSubComponentRm) SetFDescsNil(b bool)`

 SetFDescsNil sets the value for FDescs to be an explicit nil

### UnsetFDescs
`func (o *MediaSubComponentRm) UnsetFDescs()`

UnsetFDescs ensures that no value is present for FDescs, not even an explicit nil
### GetFStatus

`func (o *MediaSubComponentRm) GetFStatus() FlowStatus`

GetFStatus returns the FStatus field if non-nil, zero value otherwise.

### GetFStatusOk

`func (o *MediaSubComponentRm) GetFStatusOk() (*FlowStatus, bool)`

GetFStatusOk returns a tuple with the FStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFStatus

`func (o *MediaSubComponentRm) SetFStatus(v FlowStatus)`

SetFStatus sets FStatus field to given value.

### HasFStatus

`func (o *MediaSubComponentRm) HasFStatus() bool`

HasFStatus returns a boolean if a field has been set.

### GetMarBwDl

`func (o *MediaSubComponentRm) GetMarBwDl() string`

GetMarBwDl returns the MarBwDl field if non-nil, zero value otherwise.

### GetMarBwDlOk

`func (o *MediaSubComponentRm) GetMarBwDlOk() (*string, bool)`

GetMarBwDlOk returns a tuple with the MarBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDl

`func (o *MediaSubComponentRm) SetMarBwDl(v string)`

SetMarBwDl sets MarBwDl field to given value.

### HasMarBwDl

`func (o *MediaSubComponentRm) HasMarBwDl() bool`

HasMarBwDl returns a boolean if a field has been set.

### SetMarBwDlNil

`func (o *MediaSubComponentRm) SetMarBwDlNil(b bool)`

 SetMarBwDlNil sets the value for MarBwDl to be an explicit nil

### UnsetMarBwDl
`func (o *MediaSubComponentRm) UnsetMarBwDl()`

UnsetMarBwDl ensures that no value is present for MarBwDl, not even an explicit nil
### GetMarBwUl

`func (o *MediaSubComponentRm) GetMarBwUl() string`

GetMarBwUl returns the MarBwUl field if non-nil, zero value otherwise.

### GetMarBwUlOk

`func (o *MediaSubComponentRm) GetMarBwUlOk() (*string, bool)`

GetMarBwUlOk returns a tuple with the MarBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUl

`func (o *MediaSubComponentRm) SetMarBwUl(v string)`

SetMarBwUl sets MarBwUl field to given value.

### HasMarBwUl

`func (o *MediaSubComponentRm) HasMarBwUl() bool`

HasMarBwUl returns a boolean if a field has been set.

### SetMarBwUlNil

`func (o *MediaSubComponentRm) SetMarBwUlNil(b bool)`

 SetMarBwUlNil sets the value for MarBwUl to be an explicit nil

### UnsetMarBwUl
`func (o *MediaSubComponentRm) UnsetMarBwUl()`

UnsetMarBwUl ensures that no value is present for MarBwUl, not even an explicit nil
### GetTosTrCl

`func (o *MediaSubComponentRm) GetTosTrCl() string`

GetTosTrCl returns the TosTrCl field if non-nil, zero value otherwise.

### GetTosTrClOk

`func (o *MediaSubComponentRm) GetTosTrClOk() (*string, bool)`

GetTosTrClOk returns a tuple with the TosTrCl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosTrCl

`func (o *MediaSubComponentRm) SetTosTrCl(v string)`

SetTosTrCl sets TosTrCl field to given value.

### HasTosTrCl

`func (o *MediaSubComponentRm) HasTosTrCl() bool`

HasTosTrCl returns a boolean if a field has been set.

### SetTosTrClNil

`func (o *MediaSubComponentRm) SetTosTrClNil(b bool)`

 SetTosTrClNil sets the value for TosTrCl to be an explicit nil

### UnsetTosTrCl
`func (o *MediaSubComponentRm) UnsetTosTrCl()`

UnsetTosTrCl ensures that no value is present for TosTrCl, not even an explicit nil
### GetFlowUsage

`func (o *MediaSubComponentRm) GetFlowUsage() FlowUsage`

GetFlowUsage returns the FlowUsage field if non-nil, zero value otherwise.

### GetFlowUsageOk

`func (o *MediaSubComponentRm) GetFlowUsageOk() (*FlowUsage, bool)`

GetFlowUsageOk returns a tuple with the FlowUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUsage

`func (o *MediaSubComponentRm) SetFlowUsage(v FlowUsage)`

SetFlowUsage sets FlowUsage field to given value.

### HasFlowUsage

`func (o *MediaSubComponentRm) HasFlowUsage() bool`

HasFlowUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


