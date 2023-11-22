# ApplicationDataSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** |  | 
**DataFilters** | Pointer to [**[]DataFilter**](DataFilter.md) |  | [optional] 
**Expiry** | Pointer to **time.Time** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationDataSubs

`func NewApplicationDataSubs(notificationUri string, ) *ApplicationDataSubs`

NewApplicationDataSubs instantiates a new ApplicationDataSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataSubsWithDefaults

`func NewApplicationDataSubsWithDefaults() *ApplicationDataSubs`

NewApplicationDataSubsWithDefaults instantiates a new ApplicationDataSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *ApplicationDataSubs) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *ApplicationDataSubs) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *ApplicationDataSubs) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetDataFilters

`func (o *ApplicationDataSubs) GetDataFilters() []DataFilter`

GetDataFilters returns the DataFilters field if non-nil, zero value otherwise.

### GetDataFiltersOk

`func (o *ApplicationDataSubs) GetDataFiltersOk() (*[]DataFilter, bool)`

GetDataFiltersOk returns a tuple with the DataFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFilters

`func (o *ApplicationDataSubs) SetDataFilters(v []DataFilter)`

SetDataFilters sets DataFilters field to given value.

### HasDataFilters

`func (o *ApplicationDataSubs) HasDataFilters() bool`

HasDataFilters returns a boolean if a field has been set.

### GetExpiry

`func (o *ApplicationDataSubs) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ApplicationDataSubs) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ApplicationDataSubs) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ApplicationDataSubs) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ApplicationDataSubs) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ApplicationDataSubs) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ApplicationDataSubs) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ApplicationDataSubs) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


