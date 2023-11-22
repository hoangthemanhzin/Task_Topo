# QosFlowTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QfiList** | **[]int32** |  | 
**TunnelInfo** | [**TunnelInfo**](TunnelInfo.md) |  | 

## Methods

### NewQosFlowTunnel

`func NewQosFlowTunnel(qfiList []int32, tunnelInfo TunnelInfo, ) *QosFlowTunnel`

NewQosFlowTunnel instantiates a new QosFlowTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowTunnelWithDefaults

`func NewQosFlowTunnelWithDefaults() *QosFlowTunnel`

NewQosFlowTunnelWithDefaults instantiates a new QosFlowTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQfiList

`func (o *QosFlowTunnel) GetQfiList() []int32`

GetQfiList returns the QfiList field if non-nil, zero value otherwise.

### GetQfiListOk

`func (o *QosFlowTunnel) GetQfiListOk() (*[]int32, bool)`

GetQfiListOk returns a tuple with the QfiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQfiList

`func (o *QosFlowTunnel) SetQfiList(v []int32)`

SetQfiList sets QfiList field to given value.


### GetTunnelInfo

`func (o *QosFlowTunnel) GetTunnelInfo() TunnelInfo`

GetTunnelInfo returns the TunnelInfo field if non-nil, zero value otherwise.

### GetTunnelInfoOk

`func (o *QosFlowTunnel) GetTunnelInfoOk() (*TunnelInfo, bool)`

GetTunnelInfoOk returns a tuple with the TunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInfo

`func (o *QosFlowTunnel) SetTunnelInfo(v TunnelInfo)`

SetTunnelInfo sets TunnelInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


