# UpPathChgEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** |  | 
**NotifCorreId** | **string** | It is used to set the value of Notification Correlation ID in the notification sent by the SMF. | 
**DnaiChgType** | [**DnaiChangeType**](DnaiChangeType.md) |  | 
**AfAckInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpPathChgEvent

`func NewUpPathChgEvent(notificationUri string, notifCorreId string, dnaiChgType DnaiChangeType, ) *UpPathChgEvent`

NewUpPathChgEvent instantiates a new UpPathChgEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpPathChgEventWithDefaults

`func NewUpPathChgEventWithDefaults() *UpPathChgEvent`

NewUpPathChgEventWithDefaults instantiates a new UpPathChgEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *UpPathChgEvent) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *UpPathChgEvent) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *UpPathChgEvent) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetNotifCorreId

`func (o *UpPathChgEvent) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *UpPathChgEvent) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *UpPathChgEvent) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.


### GetDnaiChgType

`func (o *UpPathChgEvent) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *UpPathChgEvent) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *UpPathChgEvent) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.


### GetAfAckInd

`func (o *UpPathChgEvent) GetAfAckInd() bool`

GetAfAckInd returns the AfAckInd field if non-nil, zero value otherwise.

### GetAfAckIndOk

`func (o *UpPathChgEvent) GetAfAckIndOk() (*bool, bool)`

GetAfAckIndOk returns a tuple with the AfAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckInd

`func (o *UpPathChgEvent) SetAfAckInd(v bool)`

SetAfAckInd sets AfAckInd field to given value.

### HasAfAckInd

`func (o *UpPathChgEvent) HasAfAckInd() bool`

HasAfAckInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


