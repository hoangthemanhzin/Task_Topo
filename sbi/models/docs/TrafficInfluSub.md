# TrafficInfluSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnns** | Pointer to **[]string** | Each element identifies a DNN. | [optional] 
**Snssais** | Pointer to [**[]Snssai**](Snssai.md) | Each element identifies a slice. | [optional] 
**InternalGroupIds** | Pointer to **[]string** | Each element identifies a group of users. | [optional] 
**Supis** | Pointer to **[]string** | Each element identifies the user. | [optional] 
**NotificationUri** | **string** |  | 
**Expiry** | Pointer to **time.Time** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewTrafficInfluSub

`func NewTrafficInfluSub(notificationUri string, ) *TrafficInfluSub`

NewTrafficInfluSub instantiates a new TrafficInfluSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInfluSubWithDefaults

`func NewTrafficInfluSubWithDefaults() *TrafficInfluSub`

NewTrafficInfluSubWithDefaults instantiates a new TrafficInfluSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnns

`func (o *TrafficInfluSub) GetDnns() []string`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *TrafficInfluSub) GetDnnsOk() (*[]string, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *TrafficInfluSub) SetDnns(v []string)`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *TrafficInfluSub) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetSnssais

`func (o *TrafficInfluSub) GetSnssais() []Snssai`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *TrafficInfluSub) GetSnssaisOk() (*[]Snssai, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *TrafficInfluSub) SetSnssais(v []Snssai)`

SetSnssais sets Snssais field to given value.

### HasSnssais

`func (o *TrafficInfluSub) HasSnssais() bool`

HasSnssais returns a boolean if a field has been set.

### GetInternalGroupIds

`func (o *TrafficInfluSub) GetInternalGroupIds() []string`

GetInternalGroupIds returns the InternalGroupIds field if non-nil, zero value otherwise.

### GetInternalGroupIdsOk

`func (o *TrafficInfluSub) GetInternalGroupIdsOk() (*[]string, bool)`

GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIds

`func (o *TrafficInfluSub) SetInternalGroupIds(v []string)`

SetInternalGroupIds sets InternalGroupIds field to given value.

### HasInternalGroupIds

`func (o *TrafficInfluSub) HasInternalGroupIds() bool`

HasInternalGroupIds returns a boolean if a field has been set.

### GetSupis

`func (o *TrafficInfluSub) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *TrafficInfluSub) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *TrafficInfluSub) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *TrafficInfluSub) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetNotificationUri

`func (o *TrafficInfluSub) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *TrafficInfluSub) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *TrafficInfluSub) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetExpiry

`func (o *TrafficInfluSub) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TrafficInfluSub) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TrafficInfluSub) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *TrafficInfluSub) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *TrafficInfluSub) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *TrafficInfluSub) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *TrafficInfluSub) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *TrafficInfluSub) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


