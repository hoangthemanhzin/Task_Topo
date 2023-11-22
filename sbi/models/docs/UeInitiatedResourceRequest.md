# UeInitiatedResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PccRuleId** | Pointer to **string** |  | [optional] 
**RuleOp** | [**RuleOperation**](RuleOperation.md) |  | 
**Precedence** | Pointer to **int32** |  | [optional] 
**PackFiltInfo** | [**[]PacketFilterInfo**](PacketFilterInfo.md) |  | 
**ReqQos** | Pointer to [**RequestedQos**](RequestedQos.md) |  | [optional] 

## Methods

### NewUeInitiatedResourceRequest

`func NewUeInitiatedResourceRequest(ruleOp RuleOperation, packFiltInfo []PacketFilterInfo, ) *UeInitiatedResourceRequest`

NewUeInitiatedResourceRequest instantiates a new UeInitiatedResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeInitiatedResourceRequestWithDefaults

`func NewUeInitiatedResourceRequestWithDefaults() *UeInitiatedResourceRequest`

NewUeInitiatedResourceRequestWithDefaults instantiates a new UeInitiatedResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPccRuleId

`func (o *UeInitiatedResourceRequest) GetPccRuleId() string`

GetPccRuleId returns the PccRuleId field if non-nil, zero value otherwise.

### GetPccRuleIdOk

`func (o *UeInitiatedResourceRequest) GetPccRuleIdOk() (*string, bool)`

GetPccRuleIdOk returns a tuple with the PccRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRuleId

`func (o *UeInitiatedResourceRequest) SetPccRuleId(v string)`

SetPccRuleId sets PccRuleId field to given value.

### HasPccRuleId

`func (o *UeInitiatedResourceRequest) HasPccRuleId() bool`

HasPccRuleId returns a boolean if a field has been set.

### GetRuleOp

`func (o *UeInitiatedResourceRequest) GetRuleOp() RuleOperation`

GetRuleOp returns the RuleOp field if non-nil, zero value otherwise.

### GetRuleOpOk

`func (o *UeInitiatedResourceRequest) GetRuleOpOk() (*RuleOperation, bool)`

GetRuleOpOk returns a tuple with the RuleOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOp

`func (o *UeInitiatedResourceRequest) SetRuleOp(v RuleOperation)`

SetRuleOp sets RuleOp field to given value.


### GetPrecedence

`func (o *UeInitiatedResourceRequest) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *UeInitiatedResourceRequest) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *UeInitiatedResourceRequest) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *UeInitiatedResourceRequest) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetPackFiltInfo

`func (o *UeInitiatedResourceRequest) GetPackFiltInfo() []PacketFilterInfo`

GetPackFiltInfo returns the PackFiltInfo field if non-nil, zero value otherwise.

### GetPackFiltInfoOk

`func (o *UeInitiatedResourceRequest) GetPackFiltInfoOk() (*[]PacketFilterInfo, bool)`

GetPackFiltInfoOk returns a tuple with the PackFiltInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackFiltInfo

`func (o *UeInitiatedResourceRequest) SetPackFiltInfo(v []PacketFilterInfo)`

SetPackFiltInfo sets PackFiltInfo field to given value.


### GetReqQos

`func (o *UeInitiatedResourceRequest) GetReqQos() RequestedQos`

GetReqQos returns the ReqQos field if non-nil, zero value otherwise.

### GetReqQosOk

`func (o *UeInitiatedResourceRequest) GetReqQosOk() (*RequestedQos, bool)`

GetReqQosOk returns a tuple with the ReqQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQos

`func (o *UeInitiatedResourceRequest) SetReqQos(v RequestedQos)`

SetReqQos sets ReqQos field to given value.

### HasReqQos

`func (o *UeInitiatedResourceRequest) HasReqQos() bool`

HasReqQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


