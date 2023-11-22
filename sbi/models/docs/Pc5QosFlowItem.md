# Pc5QosFlowItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pqi** | **int32** |  | 
**Pc5FlowBitRates** | Pointer to [**Pc5FlowBitRates**](Pc5FlowBitRates.md) |  | [optional] 
**Range** | Pointer to **int32** |  | [optional] 

## Methods

### NewPc5QosFlowItem

`func NewPc5QosFlowItem(pqi int32, ) *Pc5QosFlowItem`

NewPc5QosFlowItem instantiates a new Pc5QosFlowItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPc5QosFlowItemWithDefaults

`func NewPc5QosFlowItemWithDefaults() *Pc5QosFlowItem`

NewPc5QosFlowItemWithDefaults instantiates a new Pc5QosFlowItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPqi

`func (o *Pc5QosFlowItem) GetPqi() int32`

GetPqi returns the Pqi field if non-nil, zero value otherwise.

### GetPqiOk

`func (o *Pc5QosFlowItem) GetPqiOk() (*int32, bool)`

GetPqiOk returns a tuple with the Pqi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPqi

`func (o *Pc5QosFlowItem) SetPqi(v int32)`

SetPqi sets Pqi field to given value.


### GetPc5FlowBitRates

`func (o *Pc5QosFlowItem) GetPc5FlowBitRates() Pc5FlowBitRates`

GetPc5FlowBitRates returns the Pc5FlowBitRates field if non-nil, zero value otherwise.

### GetPc5FlowBitRatesOk

`func (o *Pc5QosFlowItem) GetPc5FlowBitRatesOk() (*Pc5FlowBitRates, bool)`

GetPc5FlowBitRatesOk returns a tuple with the Pc5FlowBitRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPc5FlowBitRates

`func (o *Pc5QosFlowItem) SetPc5FlowBitRates(v Pc5FlowBitRates)`

SetPc5FlowBitRates sets Pc5FlowBitRates field to given value.

### HasPc5FlowBitRates

`func (o *Pc5QosFlowItem) HasPc5FlowBitRates() bool`

HasPc5FlowBitRates returns a boolean if a field has been set.

### GetRange

`func (o *Pc5QosFlowItem) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Pc5QosFlowItem) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Pc5QosFlowItem) SetRange(v int32)`

SetRange sets Range field to given value.

### HasRange

`func (o *Pc5QosFlowItem) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


