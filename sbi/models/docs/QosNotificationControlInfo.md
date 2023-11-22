# QosNotificationControlInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifType** | [**QosNotifType**](QosNotifType.md) |  | 
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 
**AltSerReq** | Pointer to **string** |  | [optional] 

## Methods

### NewQosNotificationControlInfo

`func NewQosNotificationControlInfo(notifType QosNotifType, ) *QosNotificationControlInfo`

NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosNotificationControlInfoWithDefaults

`func NewQosNotificationControlInfoWithDefaults() *QosNotificationControlInfo`

NewQosNotificationControlInfoWithDefaults instantiates a new QosNotificationControlInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifType

`func (o *QosNotificationControlInfo) GetNotifType() QosNotifType`

GetNotifType returns the NotifType field if non-nil, zero value otherwise.

### GetNotifTypeOk

`func (o *QosNotificationControlInfo) GetNotifTypeOk() (*QosNotifType, bool)`

GetNotifTypeOk returns a tuple with the NotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifType

`func (o *QosNotificationControlInfo) SetNotifType(v QosNotifType)`

SetNotifType sets NotifType field to given value.


### GetFlows

`func (o *QosNotificationControlInfo) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *QosNotificationControlInfo) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *QosNotificationControlInfo) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *QosNotificationControlInfo) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetAltSerReq

`func (o *QosNotificationControlInfo) GetAltSerReq() string`

GetAltSerReq returns the AltSerReq field if non-nil, zero value otherwise.

### GetAltSerReqOk

`func (o *QosNotificationControlInfo) GetAltSerReqOk() (*string, bool)`

GetAltSerReqOk returns a tuple with the AltSerReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSerReq

`func (o *QosNotificationControlInfo) SetAltSerReq(v string)`

SetAltSerReq sets AltSerReq field to given value.

### HasAltSerReq

`func (o *QosNotificationControlInfo) HasAltSerReq() bool`

HasAltSerReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


