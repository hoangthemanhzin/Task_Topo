# UpdatedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | **string** | Identifies a fragment (subset of resource data) of a given resource. | 
**Value** | **interface{}** |  | 

## Methods

### NewUpdatedItem

`func NewUpdatedItem(item string, value interface{}, ) *UpdatedItem`

NewUpdatedItem instantiates a new UpdatedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedItemWithDefaults

`func NewUpdatedItemWithDefaults() *UpdatedItem`

NewUpdatedItemWithDefaults instantiates a new UpdatedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *UpdatedItem) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *UpdatedItem) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *UpdatedItem) SetItem(v string)`

SetItem sets Item field to given value.


### GetValue

`func (o *UpdatedItem) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatedItem) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatedItem) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *UpdatedItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdatedItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


