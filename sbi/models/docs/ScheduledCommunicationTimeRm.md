# ScheduledCommunicationTimeRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | Pointer to **[]int32** | Identifies the day(s) of the week. If absent, it indicates every day of the week. | [optional] 
**TimeOfDayStart** | Pointer to **string** | String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC). | [optional] 
**TimeOfDayEnd** | Pointer to **string** | String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC). | [optional] 

## Methods

### NewScheduledCommunicationTimeRm

`func NewScheduledCommunicationTimeRm() *ScheduledCommunicationTimeRm`

NewScheduledCommunicationTimeRm instantiates a new ScheduledCommunicationTimeRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledCommunicationTimeRmWithDefaults

`func NewScheduledCommunicationTimeRmWithDefaults() *ScheduledCommunicationTimeRm`

NewScheduledCommunicationTimeRmWithDefaults instantiates a new ScheduledCommunicationTimeRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfWeek

`func (o *ScheduledCommunicationTimeRm) GetDaysOfWeek() []int32`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ScheduledCommunicationTimeRm) GetDaysOfWeekOk() (*[]int32, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ScheduledCommunicationTimeRm) SetDaysOfWeek(v []int32)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *ScheduledCommunicationTimeRm) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### GetTimeOfDayStart

`func (o *ScheduledCommunicationTimeRm) GetTimeOfDayStart() string`

GetTimeOfDayStart returns the TimeOfDayStart field if non-nil, zero value otherwise.

### GetTimeOfDayStartOk

`func (o *ScheduledCommunicationTimeRm) GetTimeOfDayStartOk() (*string, bool)`

GetTimeOfDayStartOk returns a tuple with the TimeOfDayStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayStart

`func (o *ScheduledCommunicationTimeRm) SetTimeOfDayStart(v string)`

SetTimeOfDayStart sets TimeOfDayStart field to given value.

### HasTimeOfDayStart

`func (o *ScheduledCommunicationTimeRm) HasTimeOfDayStart() bool`

HasTimeOfDayStart returns a boolean if a field has been set.

### GetTimeOfDayEnd

`func (o *ScheduledCommunicationTimeRm) GetTimeOfDayEnd() string`

GetTimeOfDayEnd returns the TimeOfDayEnd field if non-nil, zero value otherwise.

### GetTimeOfDayEndOk

`func (o *ScheduledCommunicationTimeRm) GetTimeOfDayEndOk() (*string, bool)`

GetTimeOfDayEndOk returns a tuple with the TimeOfDayEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayEnd

`func (o *ScheduledCommunicationTimeRm) SetTimeOfDayEnd(v string)`

SetTimeOfDayEnd sets TimeOfDayEnd field to given value.

### HasTimeOfDayEnd

`func (o *ScheduledCommunicationTimeRm) HasTimeOfDayEnd() bool`

HasTimeOfDayEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


