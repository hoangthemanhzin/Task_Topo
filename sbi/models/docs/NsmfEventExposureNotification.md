# NsmfEventExposureNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** | Notification correlation ID | 
**EventNotifs** | [**[]EventNotification**](EventNotification.md) | Notifications about Individual Events | 
**AckUri** | Pointer to **string** |  | [optional] 

## Methods

### NewNsmfEventExposureNotification

`func NewNsmfEventExposureNotification(notifId string, eventNotifs []EventNotification, ) *NsmfEventExposureNotification`

NewNsmfEventExposureNotification instantiates a new NsmfEventExposureNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsmfEventExposureNotificationWithDefaults

`func NewNsmfEventExposureNotificationWithDefaults() *NsmfEventExposureNotification`

NewNsmfEventExposureNotificationWithDefaults instantiates a new NsmfEventExposureNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *NsmfEventExposureNotification) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *NsmfEventExposureNotification) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *NsmfEventExposureNotification) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *NsmfEventExposureNotification) GetEventNotifs() []EventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *NsmfEventExposureNotification) GetEventNotifsOk() (*[]EventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *NsmfEventExposureNotification) SetEventNotifs(v []EventNotification)`

SetEventNotifs sets EventNotifs field to given value.


### GetAckUri

`func (o *NsmfEventExposureNotification) GetAckUri() string`

GetAckUri returns the AckUri field if non-nil, zero value otherwise.

### GetAckUriOk

`func (o *NsmfEventExposureNotification) GetAckUriOk() (*string, bool)`

GetAckUriOk returns a tuple with the AckUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUri

`func (o *NsmfEventExposureNotification) SetAckUri(v string)`

SetAckUri sets AckUri field to given value.

### HasAckUri

`func (o *NsmfEventExposureNotification) HasAckUri() bool`

HasAckUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


