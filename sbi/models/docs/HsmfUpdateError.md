# HsmfUpdateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**Pti** | Pointer to **int32** |  | [optional] 
**N1smCause** | Pointer to **string** |  | [optional] 
**N1SmInfoToUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**BackOffTimer** | Pointer to **int32** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewHsmfUpdateError

`func NewHsmfUpdateError(error_ ProblemDetails, ) *HsmfUpdateError`

NewHsmfUpdateError instantiates a new HsmfUpdateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmfUpdateErrorWithDefaults

`func NewHsmfUpdateErrorWithDefaults() *HsmfUpdateError`

NewHsmfUpdateErrorWithDefaults instantiates a new HsmfUpdateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *HsmfUpdateError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HsmfUpdateError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HsmfUpdateError) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetPti

`func (o *HsmfUpdateError) GetPti() int32`

GetPti returns the Pti field if non-nil, zero value otherwise.

### GetPtiOk

`func (o *HsmfUpdateError) GetPtiOk() (*int32, bool)`

GetPtiOk returns a tuple with the Pti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPti

`func (o *HsmfUpdateError) SetPti(v int32)`

SetPti sets Pti field to given value.

### HasPti

`func (o *HsmfUpdateError) HasPti() bool`

HasPti returns a boolean if a field has been set.

### GetN1smCause

`func (o *HsmfUpdateError) GetN1smCause() string`

GetN1smCause returns the N1smCause field if non-nil, zero value otherwise.

### GetN1smCauseOk

`func (o *HsmfUpdateError) GetN1smCauseOk() (*string, bool)`

GetN1smCauseOk returns a tuple with the N1smCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1smCause

`func (o *HsmfUpdateError) SetN1smCause(v string)`

SetN1smCause sets N1smCause field to given value.

### HasN1smCause

`func (o *HsmfUpdateError) HasN1smCause() bool`

HasN1smCause returns a boolean if a field has been set.

### GetN1SmInfoToUe

`func (o *HsmfUpdateError) GetN1SmInfoToUe() RefToBinaryData`

GetN1SmInfoToUe returns the N1SmInfoToUe field if non-nil, zero value otherwise.

### GetN1SmInfoToUeOk

`func (o *HsmfUpdateError) GetN1SmInfoToUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoToUe

`func (o *HsmfUpdateError) SetN1SmInfoToUe(v RefToBinaryData)`

SetN1SmInfoToUe sets N1SmInfoToUe field to given value.

### HasN1SmInfoToUe

`func (o *HsmfUpdateError) HasN1SmInfoToUe() bool`

HasN1SmInfoToUe returns a boolean if a field has been set.

### GetBackOffTimer

`func (o *HsmfUpdateError) GetBackOffTimer() int32`

GetBackOffTimer returns the BackOffTimer field if non-nil, zero value otherwise.

### GetBackOffTimerOk

`func (o *HsmfUpdateError) GetBackOffTimerOk() (*int32, bool)`

GetBackOffTimerOk returns a tuple with the BackOffTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOffTimer

`func (o *HsmfUpdateError) SetBackOffTimer(v int32)`

SetBackOffTimer sets BackOffTimer field to given value.

### HasBackOffTimer

`func (o *HsmfUpdateError) HasBackOffTimer() bool`

HasBackOffTimer returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *HsmfUpdateError) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *HsmfUpdateError) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *HsmfUpdateError) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *HsmfUpdateError) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


