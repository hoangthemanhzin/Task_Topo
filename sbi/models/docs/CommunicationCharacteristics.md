# CommunicationCharacteristics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PpSubsRegTimer** | Pointer to [**PpSubsRegTimer**](PpSubsRegTimer.md) |  | [optional] 
**PpActiveTime** | Pointer to [**PpActiveTime**](PpActiveTime.md) |  | [optional] 
**PpDlPacketCount** | Pointer to **int32** |  | [optional] 
**PpDlPacketCountExt** | Pointer to [**PpDlPacketCountExt**](PpDlPacketCountExt.md) |  | [optional] 
**PpMaximumResponseTime** | Pointer to [**PpMaximumResponseTime**](PpMaximumResponseTime.md) |  | [optional] 
**PpMaximumLatency** | Pointer to [**PpMaximumLatency**](PpMaximumLatency.md) |  | [optional] 

## Methods

### NewCommunicationCharacteristics

`func NewCommunicationCharacteristics() *CommunicationCharacteristics`

NewCommunicationCharacteristics instantiates a new CommunicationCharacteristics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationCharacteristicsWithDefaults

`func NewCommunicationCharacteristicsWithDefaults() *CommunicationCharacteristics`

NewCommunicationCharacteristicsWithDefaults instantiates a new CommunicationCharacteristics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpSubsRegTimer

`func (o *CommunicationCharacteristics) GetPpSubsRegTimer() PpSubsRegTimer`

GetPpSubsRegTimer returns the PpSubsRegTimer field if non-nil, zero value otherwise.

### GetPpSubsRegTimerOk

`func (o *CommunicationCharacteristics) GetPpSubsRegTimerOk() (*PpSubsRegTimer, bool)`

GetPpSubsRegTimerOk returns a tuple with the PpSubsRegTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpSubsRegTimer

`func (o *CommunicationCharacteristics) SetPpSubsRegTimer(v PpSubsRegTimer)`

SetPpSubsRegTimer sets PpSubsRegTimer field to given value.

### HasPpSubsRegTimer

`func (o *CommunicationCharacteristics) HasPpSubsRegTimer() bool`

HasPpSubsRegTimer returns a boolean if a field has been set.

### SetPpSubsRegTimerNil

`func (o *CommunicationCharacteristics) SetPpSubsRegTimerNil(b bool)`

 SetPpSubsRegTimerNil sets the value for PpSubsRegTimer to be an explicit nil

### UnsetPpSubsRegTimer
`func (o *CommunicationCharacteristics) UnsetPpSubsRegTimer()`

UnsetPpSubsRegTimer ensures that no value is present for PpSubsRegTimer, not even an explicit nil
### GetPpActiveTime

`func (o *CommunicationCharacteristics) GetPpActiveTime() PpActiveTime`

GetPpActiveTime returns the PpActiveTime field if non-nil, zero value otherwise.

### GetPpActiveTimeOk

`func (o *CommunicationCharacteristics) GetPpActiveTimeOk() (*PpActiveTime, bool)`

GetPpActiveTimeOk returns a tuple with the PpActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpActiveTime

`func (o *CommunicationCharacteristics) SetPpActiveTime(v PpActiveTime)`

SetPpActiveTime sets PpActiveTime field to given value.

### HasPpActiveTime

`func (o *CommunicationCharacteristics) HasPpActiveTime() bool`

HasPpActiveTime returns a boolean if a field has been set.

### SetPpActiveTimeNil

`func (o *CommunicationCharacteristics) SetPpActiveTimeNil(b bool)`

 SetPpActiveTimeNil sets the value for PpActiveTime to be an explicit nil

### UnsetPpActiveTime
`func (o *CommunicationCharacteristics) UnsetPpActiveTime()`

UnsetPpActiveTime ensures that no value is present for PpActiveTime, not even an explicit nil
### GetPpDlPacketCount

`func (o *CommunicationCharacteristics) GetPpDlPacketCount() int32`

GetPpDlPacketCount returns the PpDlPacketCount field if non-nil, zero value otherwise.

### GetPpDlPacketCountOk

`func (o *CommunicationCharacteristics) GetPpDlPacketCountOk() (*int32, bool)`

GetPpDlPacketCountOk returns a tuple with the PpDlPacketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpDlPacketCount

`func (o *CommunicationCharacteristics) SetPpDlPacketCount(v int32)`

SetPpDlPacketCount sets PpDlPacketCount field to given value.

### HasPpDlPacketCount

`func (o *CommunicationCharacteristics) HasPpDlPacketCount() bool`

HasPpDlPacketCount returns a boolean if a field has been set.

### SetPpDlPacketCountNil

`func (o *CommunicationCharacteristics) SetPpDlPacketCountNil(b bool)`

 SetPpDlPacketCountNil sets the value for PpDlPacketCount to be an explicit nil

### UnsetPpDlPacketCount
`func (o *CommunicationCharacteristics) UnsetPpDlPacketCount()`

UnsetPpDlPacketCount ensures that no value is present for PpDlPacketCount, not even an explicit nil
### GetPpDlPacketCountExt

`func (o *CommunicationCharacteristics) GetPpDlPacketCountExt() PpDlPacketCountExt`

GetPpDlPacketCountExt returns the PpDlPacketCountExt field if non-nil, zero value otherwise.

### GetPpDlPacketCountExtOk

`func (o *CommunicationCharacteristics) GetPpDlPacketCountExtOk() (*PpDlPacketCountExt, bool)`

GetPpDlPacketCountExtOk returns a tuple with the PpDlPacketCountExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpDlPacketCountExt

`func (o *CommunicationCharacteristics) SetPpDlPacketCountExt(v PpDlPacketCountExt)`

SetPpDlPacketCountExt sets PpDlPacketCountExt field to given value.

### HasPpDlPacketCountExt

`func (o *CommunicationCharacteristics) HasPpDlPacketCountExt() bool`

HasPpDlPacketCountExt returns a boolean if a field has been set.

### SetPpDlPacketCountExtNil

`func (o *CommunicationCharacteristics) SetPpDlPacketCountExtNil(b bool)`

 SetPpDlPacketCountExtNil sets the value for PpDlPacketCountExt to be an explicit nil

### UnsetPpDlPacketCountExt
`func (o *CommunicationCharacteristics) UnsetPpDlPacketCountExt()`

UnsetPpDlPacketCountExt ensures that no value is present for PpDlPacketCountExt, not even an explicit nil
### GetPpMaximumResponseTime

`func (o *CommunicationCharacteristics) GetPpMaximumResponseTime() PpMaximumResponseTime`

GetPpMaximumResponseTime returns the PpMaximumResponseTime field if non-nil, zero value otherwise.

### GetPpMaximumResponseTimeOk

`func (o *CommunicationCharacteristics) GetPpMaximumResponseTimeOk() (*PpMaximumResponseTime, bool)`

GetPpMaximumResponseTimeOk returns a tuple with the PpMaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpMaximumResponseTime

`func (o *CommunicationCharacteristics) SetPpMaximumResponseTime(v PpMaximumResponseTime)`

SetPpMaximumResponseTime sets PpMaximumResponseTime field to given value.

### HasPpMaximumResponseTime

`func (o *CommunicationCharacteristics) HasPpMaximumResponseTime() bool`

HasPpMaximumResponseTime returns a boolean if a field has been set.

### SetPpMaximumResponseTimeNil

`func (o *CommunicationCharacteristics) SetPpMaximumResponseTimeNil(b bool)`

 SetPpMaximumResponseTimeNil sets the value for PpMaximumResponseTime to be an explicit nil

### UnsetPpMaximumResponseTime
`func (o *CommunicationCharacteristics) UnsetPpMaximumResponseTime()`

UnsetPpMaximumResponseTime ensures that no value is present for PpMaximumResponseTime, not even an explicit nil
### GetPpMaximumLatency

`func (o *CommunicationCharacteristics) GetPpMaximumLatency() PpMaximumLatency`

GetPpMaximumLatency returns the PpMaximumLatency field if non-nil, zero value otherwise.

### GetPpMaximumLatencyOk

`func (o *CommunicationCharacteristics) GetPpMaximumLatencyOk() (*PpMaximumLatency, bool)`

GetPpMaximumLatencyOk returns a tuple with the PpMaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpMaximumLatency

`func (o *CommunicationCharacteristics) SetPpMaximumLatency(v PpMaximumLatency)`

SetPpMaximumLatency sets PpMaximumLatency field to given value.

### HasPpMaximumLatency

`func (o *CommunicationCharacteristics) HasPpMaximumLatency() bool`

HasPpMaximumLatency returns a boolean if a field has been set.

### SetPpMaximumLatencyNil

`func (o *CommunicationCharacteristics) SetPpMaximumLatencyNil(b bool)`

 SetPpMaximumLatencyNil sets the value for PpMaximumLatency to be an explicit nil

### UnsetPpMaximumLatency
`func (o *CommunicationCharacteristics) UnsetPpMaximumLatency()`

UnsetPpMaximumLatency ensures that no value is present for PpMaximumLatency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


