# ChangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**ChangeType**](ChangeType.md) |  | 
**Path** | **string** |  | 
**From** | Pointer to **string** |  | [optional] 
**OrigValue** | Pointer to **interface{}** |  | [optional] 
**NewValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChangeItem

`func NewChangeItem(op ChangeType, path string, ) *ChangeItem`

NewChangeItem instantiates a new ChangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeItemWithDefaults

`func NewChangeItemWithDefaults() *ChangeItem`

NewChangeItemWithDefaults instantiates a new ChangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ChangeItem) GetOp() ChangeType`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ChangeItem) GetOpOk() (*ChangeType, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ChangeItem) SetOp(v ChangeType)`

SetOp sets Op field to given value.


### GetPath

`func (o *ChangeItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChangeItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChangeItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetFrom

`func (o *ChangeItem) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChangeItem) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChangeItem) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ChangeItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetOrigValue

`func (o *ChangeItem) GetOrigValue() interface{}`

GetOrigValue returns the OrigValue field if non-nil, zero value otherwise.

### GetOrigValueOk

`func (o *ChangeItem) GetOrigValueOk() (*interface{}, bool)`

GetOrigValueOk returns a tuple with the OrigValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigValue

`func (o *ChangeItem) SetOrigValue(v interface{})`

SetOrigValue sets OrigValue field to given value.

### HasOrigValue

`func (o *ChangeItem) HasOrigValue() bool`

HasOrigValue returns a boolean if a field has been set.

### SetOrigValueNil

`func (o *ChangeItem) SetOrigValueNil(b bool)`

 SetOrigValueNil sets the value for OrigValue to be an explicit nil

### UnsetOrigValue
`func (o *ChangeItem) UnsetOrigValue()`

UnsetOrigValue ensures that no value is present for OrigValue, not even an explicit nil
### GetNewValue

`func (o *ChangeItem) GetNewValue() interface{}`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *ChangeItem) GetNewValueOk() (*interface{}, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *ChangeItem) SetNewValue(v interface{})`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *ChangeItem) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *ChangeItem) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *ChangeItem) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


