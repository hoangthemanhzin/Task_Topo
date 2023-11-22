# NotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** |  | 
**NotifItems** | [**[]UpdatedItem**](UpdatedItem.md) |  | 

## Methods

### NewNotificationItem

`func NewNotificationItem(resourceId string, notifItems []UpdatedItem, ) *NotificationItem`

NewNotificationItem instantiates a new NotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationItemWithDefaults

`func NewNotificationItemWithDefaults() *NotificationItem`

NewNotificationItemWithDefaults instantiates a new NotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *NotificationItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *NotificationItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *NotificationItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetNotifItems

`func (o *NotificationItem) GetNotifItems() []UpdatedItem`

GetNotifItems returns the NotifItems field if non-nil, zero value otherwise.

### GetNotifItemsOk

`func (o *NotificationItem) GetNotifItemsOk() (*[]UpdatedItem, bool)`

GetNotifItemsOk returns a tuple with the NotifItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifItems

`func (o *NotificationItem) SetNotifItems(v []UpdatedItem)`

SetNotifItems sets NotifItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


