# DownlinkDataNotificationControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifCtrlInds** | Pointer to [**[]NotificationControlIndication**](NotificationControlIndication.md) |  | [optional] 
**TypesOfNotif** | Pointer to [**[]DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 

## Methods

### NewDownlinkDataNotificationControl

`func NewDownlinkDataNotificationControl() *DownlinkDataNotificationControl`

NewDownlinkDataNotificationControl instantiates a new DownlinkDataNotificationControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownlinkDataNotificationControlWithDefaults

`func NewDownlinkDataNotificationControlWithDefaults() *DownlinkDataNotificationControl`

NewDownlinkDataNotificationControlWithDefaults instantiates a new DownlinkDataNotificationControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifCtrlInds

`func (o *DownlinkDataNotificationControl) GetNotifCtrlInds() []NotificationControlIndication`

GetNotifCtrlInds returns the NotifCtrlInds field if non-nil, zero value otherwise.

### GetNotifCtrlIndsOk

`func (o *DownlinkDataNotificationControl) GetNotifCtrlIndsOk() (*[]NotificationControlIndication, bool)`

GetNotifCtrlIndsOk returns a tuple with the NotifCtrlInds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCtrlInds

`func (o *DownlinkDataNotificationControl) SetNotifCtrlInds(v []NotificationControlIndication)`

SetNotifCtrlInds sets NotifCtrlInds field to given value.

### HasNotifCtrlInds

`func (o *DownlinkDataNotificationControl) HasNotifCtrlInds() bool`

HasNotifCtrlInds returns a boolean if a field has been set.

### GetTypesOfNotif

`func (o *DownlinkDataNotificationControl) GetTypesOfNotif() []DlDataDeliveryStatus`

GetTypesOfNotif returns the TypesOfNotif field if non-nil, zero value otherwise.

### GetTypesOfNotifOk

`func (o *DownlinkDataNotificationControl) GetTypesOfNotifOk() (*[]DlDataDeliveryStatus, bool)`

GetTypesOfNotifOk returns a tuple with the TypesOfNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypesOfNotif

`func (o *DownlinkDataNotificationControl) SetTypesOfNotif(v []DlDataDeliveryStatus)`

SetTypesOfNotif sets TypesOfNotif field to given value.

### HasTypesOfNotif

`func (o *DownlinkDataNotificationControl) HasTypesOfNotif() bool`

HasTypesOfNotif returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


