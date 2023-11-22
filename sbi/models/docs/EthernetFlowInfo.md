# EthernetFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EthFlows** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**FlowNumber** | **int32** |  | 

## Methods

### NewEthernetFlowInfo

`func NewEthernetFlowInfo(flowNumber int32, ) *EthernetFlowInfo`

NewEthernetFlowInfo instantiates a new EthernetFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetFlowInfoWithDefaults

`func NewEthernetFlowInfoWithDefaults() *EthernetFlowInfo`

NewEthernetFlowInfoWithDefaults instantiates a new EthernetFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEthFlows

`func (o *EthernetFlowInfo) GetEthFlows() []EthFlowDescription`

GetEthFlows returns the EthFlows field if non-nil, zero value otherwise.

### GetEthFlowsOk

`func (o *EthernetFlowInfo) GetEthFlowsOk() (*[]EthFlowDescription, bool)`

GetEthFlowsOk returns a tuple with the EthFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlows

`func (o *EthernetFlowInfo) SetEthFlows(v []EthFlowDescription)`

SetEthFlows sets EthFlows field to given value.

### HasEthFlows

`func (o *EthernetFlowInfo) HasEthFlows() bool`

HasEthFlows returns a boolean if a field has been set.

### GetFlowNumber

`func (o *EthernetFlowInfo) GetFlowNumber() int32`

GetFlowNumber returns the FlowNumber field if non-nil, zero value otherwise.

### GetFlowNumberOk

`func (o *EthernetFlowInfo) GetFlowNumberOk() (*int32, bool)`

GetFlowNumberOk returns a tuple with the FlowNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowNumber

`func (o *EthernetFlowInfo) SetFlowNumber(v int32)`

SetFlowNumber sets FlowNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


