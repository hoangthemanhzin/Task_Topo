# IpFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpFlows** | Pointer to **[]string** |  | [optional] 
**FlowNumber** | **int32** |  | 

## Methods

### NewIpFlowInfo

`func NewIpFlowInfo(flowNumber int32, ) *IpFlowInfo`

NewIpFlowInfo instantiates a new IpFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpFlowInfoWithDefaults

`func NewIpFlowInfoWithDefaults() *IpFlowInfo`

NewIpFlowInfoWithDefaults instantiates a new IpFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFlows

`func (o *IpFlowInfo) GetIpFlows() []string`

GetIpFlows returns the IpFlows field if non-nil, zero value otherwise.

### GetIpFlowsOk

`func (o *IpFlowInfo) GetIpFlowsOk() (*[]string, bool)`

GetIpFlowsOk returns a tuple with the IpFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFlows

`func (o *IpFlowInfo) SetIpFlows(v []string)`

SetIpFlows sets IpFlows field to given value.

### HasIpFlows

`func (o *IpFlowInfo) HasIpFlows() bool`

HasIpFlows returns a boolean if a field has been set.

### GetFlowNumber

`func (o *IpFlowInfo) GetFlowNumber() int32`

GetFlowNumber returns the FlowNumber field if non-nil, zero value otherwise.

### GetFlowNumberOk

`func (o *IpFlowInfo) GetFlowNumberOk() (*int32, bool)`

GetFlowNumberOk returns a tuple with the FlowNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowNumber

`func (o *IpFlowInfo) SetFlowNumber(v int32)`

SetFlowNumber sets FlowNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


