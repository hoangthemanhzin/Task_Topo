# PfdChangeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**RemovalFlag** | Pointer to **bool** |  | [optional] [default to false]
**PartialFlag** | Pointer to **bool** |  | [optional] [default to false]
**Pfds** | Pointer to [**[]PfdContent**](PfdContent.md) |  | [optional] 

## Methods

### NewPfdChangeNotification

`func NewPfdChangeNotification(applicationId string, ) *PfdChangeNotification`

NewPfdChangeNotification instantiates a new PfdChangeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdChangeNotificationWithDefaults

`func NewPfdChangeNotificationWithDefaults() *PfdChangeNotification`

NewPfdChangeNotificationWithDefaults instantiates a new PfdChangeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PfdChangeNotification) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PfdChangeNotification) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PfdChangeNotification) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetRemovalFlag

`func (o *PfdChangeNotification) GetRemovalFlag() bool`

GetRemovalFlag returns the RemovalFlag field if non-nil, zero value otherwise.

### GetRemovalFlagOk

`func (o *PfdChangeNotification) GetRemovalFlagOk() (*bool, bool)`

GetRemovalFlagOk returns a tuple with the RemovalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalFlag

`func (o *PfdChangeNotification) SetRemovalFlag(v bool)`

SetRemovalFlag sets RemovalFlag field to given value.

### HasRemovalFlag

`func (o *PfdChangeNotification) HasRemovalFlag() bool`

HasRemovalFlag returns a boolean if a field has been set.

### GetPartialFlag

`func (o *PfdChangeNotification) GetPartialFlag() bool`

GetPartialFlag returns the PartialFlag field if non-nil, zero value otherwise.

### GetPartialFlagOk

`func (o *PfdChangeNotification) GetPartialFlagOk() (*bool, bool)`

GetPartialFlagOk returns a tuple with the PartialFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFlag

`func (o *PfdChangeNotification) SetPartialFlag(v bool)`

SetPartialFlag sets PartialFlag field to given value.

### HasPartialFlag

`func (o *PfdChangeNotification) HasPartialFlag() bool`

HasPartialFlag returns a boolean if a field has been set.

### GetPfds

`func (o *PfdChangeNotification) GetPfds() []PfdContent`

GetPfds returns the Pfds field if non-nil, zero value otherwise.

### GetPfdsOk

`func (o *PfdChangeNotification) GetPfdsOk() (*[]PfdContent, bool)`

GetPfdsOk returns a tuple with the Pfds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfds

`func (o *PfdChangeNotification) SetPfds(v []PfdContent)`

SetPfds sets Pfds field to given value.

### HasPfds

`func (o *PfdChangeNotification) HasPfds() bool`

HasPfds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


