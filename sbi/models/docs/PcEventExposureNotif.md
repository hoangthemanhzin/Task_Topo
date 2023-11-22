# PcEventExposureNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**EventNotifs** | [**[]PcEventNotification**](PcEventNotification.md) |  | 

## Methods

### NewPcEventExposureNotif

`func NewPcEventExposureNotif(notifId string, eventNotifs []PcEventNotification, ) *PcEventExposureNotif`

NewPcEventExposureNotif instantiates a new PcEventExposureNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcEventExposureNotifWithDefaults

`func NewPcEventExposureNotifWithDefaults() *PcEventExposureNotif`

NewPcEventExposureNotifWithDefaults instantiates a new PcEventExposureNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *PcEventExposureNotif) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *PcEventExposureNotif) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *PcEventExposureNotif) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *PcEventExposureNotif) GetEventNotifs() []PcEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *PcEventExposureNotif) GetEventNotifsOk() (*[]PcEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *PcEventExposureNotif) SetEventNotifs(v []PcEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


