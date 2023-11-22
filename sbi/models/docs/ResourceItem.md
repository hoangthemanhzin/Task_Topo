# ResourceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonResourceUri** | **string** |  | 
**Items** | **[]string** |  | 

## Methods

### NewResourceItem

`func NewResourceItem(monResourceUri string, items []string, ) *ResourceItem`

NewResourceItem instantiates a new ResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceItemWithDefaults

`func NewResourceItemWithDefaults() *ResourceItem`

NewResourceItemWithDefaults instantiates a new ResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonResourceUri

`func (o *ResourceItem) GetMonResourceUri() string`

GetMonResourceUri returns the MonResourceUri field if non-nil, zero value otherwise.

### GetMonResourceUriOk

`func (o *ResourceItem) GetMonResourceUriOk() (*string, bool)`

GetMonResourceUriOk returns a tuple with the MonResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonResourceUri

`func (o *ResourceItem) SetMonResourceUri(v string)`

SetMonResourceUri sets MonResourceUri field to given value.


### GetItems

`func (o *ResourceItem) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceItem) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceItem) SetItems(v []string)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


