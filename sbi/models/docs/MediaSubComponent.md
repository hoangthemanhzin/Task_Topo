# MediaSubComponent

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
**TosTrCl** | Pointer to **string** | 2-octet string, where each octet is encoded in hexadecimal representation. The first octet contains the IPv4 Type-of-Service or the IPv6 Traffic-Class field and the second octet contains the ToS/Traffic Class mask field. | [optional] 
**FlowUsage** | Pointer to [**FlowUsage**](FlowUsage.md) |  | [optional] 

## Methods

### NewMediaSubComponent

`func NewMediaSubComponent(fNum int32, ) *MediaSubComponent`

NewMediaSubComponent instantiates a new MediaSubComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaSubComponentWithDefaults

`func NewMediaSubComponentWithDefaults() *MediaSubComponent`

NewMediaSubComponentWithDefaults instantiates a new MediaSubComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfSigProtocol

`func (o *MediaSubComponent) GetAfSigProtocol() AfSigProtocol`

GetAfSigProtocol returns the AfSigProtocol field if non-nil, zero value otherwise.

### GetAfSigProtocolOk

`func (o *MediaSubComponent) GetAfSigProtocolOk() (*AfSigProtocol, bool)`

GetAfSigProtocolOk returns a tuple with the AfSigProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfSigProtocol

`func (o *MediaSubComponent) SetAfSigProtocol(v AfSigProtocol)`

SetAfSigProtocol sets AfSigProtocol field to given value.

### HasAfSigProtocol

`func (o *MediaSubComponent) HasAfSigProtocol() bool`

HasAfSigProtocol returns a boolean if a field has been set.

### GetEthfDescs

`func (o *MediaSubComponent) GetEthfDescs() []EthFlowDescription`

GetEthfDescs returns the EthfDescs field if non-nil, zero value otherwise.

### GetEthfDescsOk

`func (o *MediaSubComponent) GetEthfDescsOk() (*[]EthFlowDescription, bool)`

GetEthfDescsOk returns a tuple with the EthfDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthfDescs

`func (o *MediaSubComponent) SetEthfDescs(v []EthFlowDescription)`

SetEthfDescs sets EthfDescs field to given value.

### HasEthfDescs

`func (o *MediaSubComponent) HasEthfDescs() bool`

HasEthfDescs returns a boolean if a field has been set.

### GetFNum

`func (o *MediaSubComponent) GetFNum() int32`

GetFNum returns the FNum field if non-nil, zero value otherwise.

### GetFNumOk

`func (o *MediaSubComponent) GetFNumOk() (*int32, bool)`

GetFNumOk returns a tuple with the FNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFNum

`func (o *MediaSubComponent) SetFNum(v int32)`

SetFNum sets FNum field to given value.


### GetFDescs

`func (o *MediaSubComponent) GetFDescs() []string`

GetFDescs returns the FDescs field if non-nil, zero value otherwise.

### GetFDescsOk

`func (o *MediaSubComponent) GetFDescsOk() (*[]string, bool)`

GetFDescsOk returns a tuple with the FDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDescs

`func (o *MediaSubComponent) SetFDescs(v []string)`

SetFDescs sets FDescs field to given value.

### HasFDescs

`func (o *MediaSubComponent) HasFDescs() bool`

HasFDescs returns a boolean if a field has been set.

### GetFStatus

`func (o *MediaSubComponent) GetFStatus() FlowStatus`

GetFStatus returns the FStatus field if non-nil, zero value otherwise.

### GetFStatusOk

`func (o *MediaSubComponent) GetFStatusOk() (*FlowStatus, bool)`

GetFStatusOk returns a tuple with the FStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFStatus

`func (o *MediaSubComponent) SetFStatus(v FlowStatus)`

SetFStatus sets FStatus field to given value.

### HasFStatus

`func (o *MediaSubComponent) HasFStatus() bool`

HasFStatus returns a boolean if a field has been set.

### GetMarBwDl

`func (o *MediaSubComponent) GetMarBwDl() string`

GetMarBwDl returns the MarBwDl field if non-nil, zero value otherwise.

### GetMarBwDlOk

`func (o *MediaSubComponent) GetMarBwDlOk() (*string, bool)`

GetMarBwDlOk returns a tuple with the MarBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDl

`func (o *MediaSubComponent) SetMarBwDl(v string)`

SetMarBwDl sets MarBwDl field to given value.

### HasMarBwDl

`func (o *MediaSubComponent) HasMarBwDl() bool`

HasMarBwDl returns a boolean if a field has been set.

### GetMarBwUl

`func (o *MediaSubComponent) GetMarBwUl() string`

GetMarBwUl returns the MarBwUl field if non-nil, zero value otherwise.

### GetMarBwUlOk

`func (o *MediaSubComponent) GetMarBwUlOk() (*string, bool)`

GetMarBwUlOk returns a tuple with the MarBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUl

`func (o *MediaSubComponent) SetMarBwUl(v string)`

SetMarBwUl sets MarBwUl field to given value.

### HasMarBwUl

`func (o *MediaSubComponent) HasMarBwUl() bool`

HasMarBwUl returns a boolean if a field has been set.

### GetTosTrCl

`func (o *MediaSubComponent) GetTosTrCl() string`

GetTosTrCl returns the TosTrCl field if non-nil, zero value otherwise.

### GetTosTrClOk

`func (o *MediaSubComponent) GetTosTrClOk() (*string, bool)`

GetTosTrClOk returns a tuple with the TosTrCl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosTrCl

`func (o *MediaSubComponent) SetTosTrCl(v string)`

SetTosTrCl sets TosTrCl field to given value.

### HasTosTrCl

`func (o *MediaSubComponent) HasTosTrCl() bool`

HasTosTrCl returns a boolean if a field has been set.

### GetFlowUsage

`func (o *MediaSubComponent) GetFlowUsage() FlowUsage`

GetFlowUsage returns the FlowUsage field if non-nil, zero value otherwise.

### GetFlowUsageOk

`func (o *MediaSubComponent) GetFlowUsageOk() (*FlowUsage, bool)`

GetFlowUsageOk returns a tuple with the FlowUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUsage

`func (o *MediaSubComponent) SetFlowUsage(v FlowUsage)`

SetFlowUsage sets FlowUsage field to given value.

### HasFlowUsage

`func (o *MediaSubComponent) HasFlowUsage() bool`

HasFlowUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


