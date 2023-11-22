# SdmSubsModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **time.Time** |  | [optional] 
**MonitoredResourceUris** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSdmSubsModification

`func NewSdmSubsModification() *SdmSubsModification`

NewSdmSubsModification instantiates a new SdmSubsModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdmSubsModificationWithDefaults

`func NewSdmSubsModificationWithDefaults() *SdmSubsModification`

NewSdmSubsModificationWithDefaults instantiates a new SdmSubsModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *SdmSubsModification) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SdmSubsModification) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SdmSubsModification) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SdmSubsModification) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *SdmSubsModification) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SdmSubsModification) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SdmSubsModification) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.

### HasMonitoredResourceUris

`func (o *SdmSubsModification) HasMonitoredResourceUris() bool`

HasMonitoredResourceUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


