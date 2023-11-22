# EventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**SmfEvent**](SmfEvent.md) |  | 
**DnaiChgType** | Pointer to [**DnaiChangeType**](DnaiChangeType.md) |  | [optional] 
**DddTraDescriptors** | Pointer to [**[]DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 
**DddStati** | Pointer to [**[]DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEventSubscription

`func NewEventSubscription(event SmfEvent, ) *EventSubscription`

NewEventSubscription instantiates a new EventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionWithDefaults

`func NewEventSubscriptionWithDefaults() *EventSubscription`

NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventSubscription) GetEvent() SmfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventSubscription) GetEventOk() (*SmfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventSubscription) SetEvent(v SmfEvent)`

SetEvent sets Event field to given value.


### GetDnaiChgType

`func (o *EventSubscription) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *EventSubscription) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *EventSubscription) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.

### HasDnaiChgType

`func (o *EventSubscription) HasDnaiChgType() bool`

HasDnaiChgType returns a boolean if a field has been set.

### GetDddTraDescriptors

`func (o *EventSubscription) GetDddTraDescriptors() []DddTrafficDescriptor`

GetDddTraDescriptors returns the DddTraDescriptors field if non-nil, zero value otherwise.

### GetDddTraDescriptorsOk

`func (o *EventSubscription) GetDddTraDescriptorsOk() (*[]DddTrafficDescriptor, bool)`

GetDddTraDescriptorsOk returns a tuple with the DddTraDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTraDescriptors

`func (o *EventSubscription) SetDddTraDescriptors(v []DddTrafficDescriptor)`

SetDddTraDescriptors sets DddTraDescriptors field to given value.

### HasDddTraDescriptors

`func (o *EventSubscription) HasDddTraDescriptors() bool`

HasDddTraDescriptors returns a boolean if a field has been set.

### GetDddStati

`func (o *EventSubscription) GetDddStati() []DlDataDeliveryStatus`

GetDddStati returns the DddStati field if non-nil, zero value otherwise.

### GetDddStatiOk

`func (o *EventSubscription) GetDddStatiOk() (*[]DlDataDeliveryStatus, bool)`

GetDddStatiOk returns a tuple with the DddStati field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddStati

`func (o *EventSubscription) SetDddStati(v []DlDataDeliveryStatus)`

SetDddStati sets DddStati field to given value.

### HasDddStati

`func (o *EventSubscription) HasDddStati() bool`

HasDddStati returns a boolean if a field has been set.

### GetAppIds

`func (o *EventSubscription) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventSubscription) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventSubscription) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventSubscription) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


