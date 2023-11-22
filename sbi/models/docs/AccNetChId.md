# AccNetChId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccNetChaIdValue** | **int32** |  | 
**RefPccRuleIds** | Pointer to **[]string** | Contains the identifier of the PCC rule(s) associated to the provided Access Network Charging Identifier. | [optional] 
**SessionChScope** | Pointer to **bool** | When it is included and set to true, indicates the Access Network Charging Identifier applies to the whole PDU Session | [optional] 

## Methods

### NewAccNetChId

`func NewAccNetChId(accNetChaIdValue int32, ) *AccNetChId`

NewAccNetChId instantiates a new AccNetChId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccNetChIdWithDefaults

`func NewAccNetChIdWithDefaults() *AccNetChId`

NewAccNetChIdWithDefaults instantiates a new AccNetChId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccNetChaIdValue

`func (o *AccNetChId) GetAccNetChaIdValue() int32`

GetAccNetChaIdValue returns the AccNetChaIdValue field if non-nil, zero value otherwise.

### GetAccNetChaIdValueOk

`func (o *AccNetChId) GetAccNetChaIdValueOk() (*int32, bool)`

GetAccNetChaIdValueOk returns a tuple with the AccNetChaIdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccNetChaIdValue

`func (o *AccNetChId) SetAccNetChaIdValue(v int32)`

SetAccNetChaIdValue sets AccNetChaIdValue field to given value.


### GetRefPccRuleIds

`func (o *AccNetChId) GetRefPccRuleIds() []string`

GetRefPccRuleIds returns the RefPccRuleIds field if non-nil, zero value otherwise.

### GetRefPccRuleIdsOk

`func (o *AccNetChId) GetRefPccRuleIdsOk() (*[]string, bool)`

GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPccRuleIds

`func (o *AccNetChId) SetRefPccRuleIds(v []string)`

SetRefPccRuleIds sets RefPccRuleIds field to given value.

### HasRefPccRuleIds

`func (o *AccNetChId) HasRefPccRuleIds() bool`

HasRefPccRuleIds returns a boolean if a field has been set.

### GetSessionChScope

`func (o *AccNetChId) GetSessionChScope() bool`

GetSessionChScope returns the SessionChScope field if non-nil, zero value otherwise.

### GetSessionChScopeOk

`func (o *AccNetChId) GetSessionChScopeOk() (*bool, bool)`

GetSessionChScopeOk returns a tuple with the SessionChScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionChScope

`func (o *AccNetChId) SetSessionChScope(v bool)`

SetSessionChScope sets SessionChScope field to given value.

### HasSessionChScope

`func (o *AccNetChId) HasSessionChScope() bool`

HasSessionChScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


