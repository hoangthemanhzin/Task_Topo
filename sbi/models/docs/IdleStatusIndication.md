# IdleStatusIndication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**ActiveTime** | Pointer to **int32** |  | [optional] 
**SubsRegTimer** | Pointer to **int32** |  | [optional] 
**EdrxCycleLength** | Pointer to **int32** |  | [optional] 
**SuggestedNumOfDlPackets** | Pointer to **int32** |  | [optional] 

## Methods

### NewIdleStatusIndication

`func NewIdleStatusIndication() *IdleStatusIndication`

NewIdleStatusIndication instantiates a new IdleStatusIndication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdleStatusIndicationWithDefaults

`func NewIdleStatusIndicationWithDefaults() *IdleStatusIndication`

NewIdleStatusIndicationWithDefaults instantiates a new IdleStatusIndication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStamp

`func (o *IdleStatusIndication) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *IdleStatusIndication) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *IdleStatusIndication) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *IdleStatusIndication) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetActiveTime

`func (o *IdleStatusIndication) GetActiveTime() int32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *IdleStatusIndication) GetActiveTimeOk() (*int32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *IdleStatusIndication) SetActiveTime(v int32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *IdleStatusIndication) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### GetSubsRegTimer

`func (o *IdleStatusIndication) GetSubsRegTimer() int32`

GetSubsRegTimer returns the SubsRegTimer field if non-nil, zero value otherwise.

### GetSubsRegTimerOk

`func (o *IdleStatusIndication) GetSubsRegTimerOk() (*int32, bool)`

GetSubsRegTimerOk returns a tuple with the SubsRegTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsRegTimer

`func (o *IdleStatusIndication) SetSubsRegTimer(v int32)`

SetSubsRegTimer sets SubsRegTimer field to given value.

### HasSubsRegTimer

`func (o *IdleStatusIndication) HasSubsRegTimer() bool`

HasSubsRegTimer returns a boolean if a field has been set.

### GetEdrxCycleLength

`func (o *IdleStatusIndication) GetEdrxCycleLength() int32`

GetEdrxCycleLength returns the EdrxCycleLength field if non-nil, zero value otherwise.

### GetEdrxCycleLengthOk

`func (o *IdleStatusIndication) GetEdrxCycleLengthOk() (*int32, bool)`

GetEdrxCycleLengthOk returns a tuple with the EdrxCycleLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdrxCycleLength

`func (o *IdleStatusIndication) SetEdrxCycleLength(v int32)`

SetEdrxCycleLength sets EdrxCycleLength field to given value.

### HasEdrxCycleLength

`func (o *IdleStatusIndication) HasEdrxCycleLength() bool`

HasEdrxCycleLength returns a boolean if a field has been set.

### GetSuggestedNumOfDlPackets

`func (o *IdleStatusIndication) GetSuggestedNumOfDlPackets() int32`

GetSuggestedNumOfDlPackets returns the SuggestedNumOfDlPackets field if non-nil, zero value otherwise.

### GetSuggestedNumOfDlPacketsOk

`func (o *IdleStatusIndication) GetSuggestedNumOfDlPacketsOk() (*int32, bool)`

GetSuggestedNumOfDlPacketsOk returns a tuple with the SuggestedNumOfDlPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedNumOfDlPackets

`func (o *IdleStatusIndication) SetSuggestedNumOfDlPackets(v int32)`

SetSuggestedNumOfDlPackets sets SuggestedNumOfDlPackets field to given value.

### HasSuggestedNumOfDlPackets

`func (o *IdleStatusIndication) HasSuggestedNumOfDlPackets() bool`

HasSuggestedNumOfDlPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


