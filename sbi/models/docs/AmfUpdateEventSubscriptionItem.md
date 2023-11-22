# AmfUpdateEventSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | **string** |  | 
**Value** | Pointer to [**AmfEvent**](AmfEvent.md) |  | [optional] 

## Methods

### NewAmfUpdateEventSubscriptionItem

`func NewAmfUpdateEventSubscriptionItem(op string, path string, ) *AmfUpdateEventSubscriptionItem`

NewAmfUpdateEventSubscriptionItem instantiates a new AmfUpdateEventSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfUpdateEventSubscriptionItemWithDefaults

`func NewAmfUpdateEventSubscriptionItemWithDefaults() *AmfUpdateEventSubscriptionItem`

NewAmfUpdateEventSubscriptionItemWithDefaults instantiates a new AmfUpdateEventSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AmfUpdateEventSubscriptionItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AmfUpdateEventSubscriptionItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AmfUpdateEventSubscriptionItem) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *AmfUpdateEventSubscriptionItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AmfUpdateEventSubscriptionItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AmfUpdateEventSubscriptionItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AmfUpdateEventSubscriptionItem) GetValue() AmfEvent`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AmfUpdateEventSubscriptionItem) GetValueOk() (*AmfEvent, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AmfUpdateEventSubscriptionItem) SetValue(v AmfEvent)`

SetValue sets Value field to given value.

### HasValue

`func (o *AmfUpdateEventSubscriptionItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


