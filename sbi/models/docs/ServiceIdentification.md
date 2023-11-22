# ServiceIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServEthFlows** | Pointer to [**[]EthernetFlowInfo**](EthernetFlowInfo.md) |  | [optional] 
**ServIpFlows** | Pointer to [**[]IpFlowInfo**](IpFlowInfo.md) |  | [optional] 
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 

## Methods

### NewServiceIdentification

`func NewServiceIdentification() *ServiceIdentification`

NewServiceIdentification instantiates a new ServiceIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIdentificationWithDefaults

`func NewServiceIdentificationWithDefaults() *ServiceIdentification`

NewServiceIdentificationWithDefaults instantiates a new ServiceIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServEthFlows

`func (o *ServiceIdentification) GetServEthFlows() []EthernetFlowInfo`

GetServEthFlows returns the ServEthFlows field if non-nil, zero value otherwise.

### GetServEthFlowsOk

`func (o *ServiceIdentification) GetServEthFlowsOk() (*[]EthernetFlowInfo, bool)`

GetServEthFlowsOk returns a tuple with the ServEthFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServEthFlows

`func (o *ServiceIdentification) SetServEthFlows(v []EthernetFlowInfo)`

SetServEthFlows sets ServEthFlows field to given value.

### HasServEthFlows

`func (o *ServiceIdentification) HasServEthFlows() bool`

HasServEthFlows returns a boolean if a field has been set.

### GetServIpFlows

`func (o *ServiceIdentification) GetServIpFlows() []IpFlowInfo`

GetServIpFlows returns the ServIpFlows field if non-nil, zero value otherwise.

### GetServIpFlowsOk

`func (o *ServiceIdentification) GetServIpFlowsOk() (*[]IpFlowInfo, bool)`

GetServIpFlowsOk returns a tuple with the ServIpFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServIpFlows

`func (o *ServiceIdentification) SetServIpFlows(v []IpFlowInfo)`

SetServIpFlows sets ServIpFlows field to given value.

### HasServIpFlows

`func (o *ServiceIdentification) HasServIpFlows() bool`

HasServIpFlows returns a boolean if a field has been set.

### GetAfAppId

`func (o *ServiceIdentification) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *ServiceIdentification) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *ServiceIdentification) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *ServiceIdentification) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


