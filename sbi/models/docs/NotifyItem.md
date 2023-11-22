# NotifyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** |  | 
**Changes** | [**[]ChangeItem**](ChangeItem.md) |  | 

## Methods

### NewNotifyItem

`func NewNotifyItem(resourceId string, changes []ChangeItem, ) *NotifyItem`

NewNotifyItem instantiates a new NotifyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyItemWithDefaults

`func NewNotifyItemWithDefaults() *NotifyItem`

NewNotifyItemWithDefaults instantiates a new NotifyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *NotifyItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *NotifyItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *NotifyItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetChanges

`func (o *NotifyItem) GetChanges() []ChangeItem`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *NotifyItem) GetChangesOk() (*[]ChangeItem, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *NotifyItem) SetChanges(v []ChangeItem)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


