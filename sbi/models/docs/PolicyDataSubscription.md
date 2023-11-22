# PolicyDataSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** |  | 
**NotifId** | Pointer to **string** |  | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**MonResItems** | Pointer to [**[]ResourceItem**](ResourceItem.md) |  | [optional] 
**Expiry** | Pointer to **time.Time** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyDataSubscription

`func NewPolicyDataSubscription(notificationUri string, monitoredResourceUris []string, ) *PolicyDataSubscription`

NewPolicyDataSubscription instantiates a new PolicyDataSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDataSubscriptionWithDefaults

`func NewPolicyDataSubscriptionWithDefaults() *PolicyDataSubscription`

NewPolicyDataSubscriptionWithDefaults instantiates a new PolicyDataSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *PolicyDataSubscription) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *PolicyDataSubscription) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *PolicyDataSubscription) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetNotifId

`func (o *PolicyDataSubscription) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *PolicyDataSubscription) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *PolicyDataSubscription) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.

### HasNotifId

`func (o *PolicyDataSubscription) HasNotifId() bool`

HasNotifId returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *PolicyDataSubscription) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *PolicyDataSubscription) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *PolicyDataSubscription) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetMonResItems

`func (o *PolicyDataSubscription) GetMonResItems() []ResourceItem`

GetMonResItems returns the MonResItems field if non-nil, zero value otherwise.

### GetMonResItemsOk

`func (o *PolicyDataSubscription) GetMonResItemsOk() (*[]ResourceItem, bool)`

GetMonResItemsOk returns a tuple with the MonResItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonResItems

`func (o *PolicyDataSubscription) SetMonResItems(v []ResourceItem)`

SetMonResItems sets MonResItems field to given value.

### HasMonResItems

`func (o *PolicyDataSubscription) HasMonResItems() bool`

HasMonResItems returns a boolean if a field has been set.

### GetExpiry

`func (o *PolicyDataSubscription) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PolicyDataSubscription) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PolicyDataSubscription) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PolicyDataSubscription) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PolicyDataSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PolicyDataSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PolicyDataSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PolicyDataSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


