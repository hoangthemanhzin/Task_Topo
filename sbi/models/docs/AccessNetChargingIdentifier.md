# AccessNetChargingIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccNetChaIdValue** | **int32** |  | 
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 

## Methods

### NewAccessNetChargingIdentifier

`func NewAccessNetChargingIdentifier(accNetChaIdValue int32, ) *AccessNetChargingIdentifier`

NewAccessNetChargingIdentifier instantiates a new AccessNetChargingIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessNetChargingIdentifierWithDefaults

`func NewAccessNetChargingIdentifierWithDefaults() *AccessNetChargingIdentifier`

NewAccessNetChargingIdentifierWithDefaults instantiates a new AccessNetChargingIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccNetChaIdValue

`func (o *AccessNetChargingIdentifier) GetAccNetChaIdValue() int32`

GetAccNetChaIdValue returns the AccNetChaIdValue field if non-nil, zero value otherwise.

### GetAccNetChaIdValueOk

`func (o *AccessNetChargingIdentifier) GetAccNetChaIdValueOk() (*int32, bool)`

GetAccNetChaIdValueOk returns a tuple with the AccNetChaIdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccNetChaIdValue

`func (o *AccessNetChargingIdentifier) SetAccNetChaIdValue(v int32)`

SetAccNetChaIdValue sets AccNetChaIdValue field to given value.


### GetFlows

`func (o *AccessNetChargingIdentifier) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *AccessNetChargingIdentifier) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *AccessNetChargingIdentifier) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *AccessNetChargingIdentifier) HasFlows() bool`

HasFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


