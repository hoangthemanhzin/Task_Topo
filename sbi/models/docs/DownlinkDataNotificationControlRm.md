# DownlinkDataNotificationControlRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifCtrlInds** | Pointer to [**[]NotificationControlIndication**](NotificationControlIndication.md) |  | [optional] 
**TypesOfNotif** | Pointer to [**[]DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 

## Methods

### NewDownlinkDataNotificationControlRm

`func NewDownlinkDataNotificationControlRm() *DownlinkDataNotificationControlRm`

NewDownlinkDataNotificationControlRm instantiates a new DownlinkDataNotificationControlRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownlinkDataNotificationControlRmWithDefaults

`func NewDownlinkDataNotificationControlRmWithDefaults() *DownlinkDataNotificationControlRm`

NewDownlinkDataNotificationControlRmWithDefaults instantiates a new DownlinkDataNotificationControlRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifCtrlInds

`func (o *DownlinkDataNotificationControlRm) GetNotifCtrlInds() []NotificationControlIndication`

GetNotifCtrlInds returns the NotifCtrlInds field if non-nil, zero value otherwise.

### GetNotifCtrlIndsOk

`func (o *DownlinkDataNotificationControlRm) GetNotifCtrlIndsOk() (*[]NotificationControlIndication, bool)`

GetNotifCtrlIndsOk returns a tuple with the NotifCtrlInds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCtrlInds

`func (o *DownlinkDataNotificationControlRm) SetNotifCtrlInds(v []NotificationControlIndication)`

SetNotifCtrlInds sets NotifCtrlInds field to given value.

### HasNotifCtrlInds

`func (o *DownlinkDataNotificationControlRm) HasNotifCtrlInds() bool`

HasNotifCtrlInds returns a boolean if a field has been set.

### SetNotifCtrlIndsNil

`func (o *DownlinkDataNotificationControlRm) SetNotifCtrlIndsNil(b bool)`

 SetNotifCtrlIndsNil sets the value for NotifCtrlInds to be an explicit nil

### UnsetNotifCtrlInds
`func (o *DownlinkDataNotificationControlRm) UnsetNotifCtrlInds()`

UnsetNotifCtrlInds ensures that no value is present for NotifCtrlInds, not even an explicit nil
### GetTypesOfNotif

`func (o *DownlinkDataNotificationControlRm) GetTypesOfNotif() []DlDataDeliveryStatus`

GetTypesOfNotif returns the TypesOfNotif field if non-nil, zero value otherwise.

### GetTypesOfNotifOk

`func (o *DownlinkDataNotificationControlRm) GetTypesOfNotifOk() (*[]DlDataDeliveryStatus, bool)`

GetTypesOfNotifOk returns a tuple with the TypesOfNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypesOfNotif

`func (o *DownlinkDataNotificationControlRm) SetTypesOfNotif(v []DlDataDeliveryStatus)`

SetTypesOfNotif sets TypesOfNotif field to given value.

### HasTypesOfNotif

`func (o *DownlinkDataNotificationControlRm) HasTypesOfNotif() bool`

HasTypesOfNotif returns a boolean if a field has been set.

### SetTypesOfNotifNil

`func (o *DownlinkDataNotificationControlRm) SetTypesOfNotifNil(b bool)`

 SetTypesOfNotifNil sets the value for TypesOfNotif to be an explicit nil

### UnsetTypesOfNotif
`func (o *DownlinkDataNotificationControlRm) UnsetTypesOfNotif()`

UnsetTypesOfNotif ensures that no value is present for TypesOfNotif, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


