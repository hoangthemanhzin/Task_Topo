# VsmfUpdateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**Pti** | Pointer to **int32** |  | [optional] 
**N1smCause** | Pointer to **string** |  | [optional] 
**N1SmInfoFromUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UnknownN1SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**FailedToAssignEbiList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 
**NgApCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**Var5gMmCauseValue** | Pointer to **int32** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 

## Methods

### NewVsmfUpdateError

`func NewVsmfUpdateError(error_ ProblemDetails, ) *VsmfUpdateError`

NewVsmfUpdateError instantiates a new VsmfUpdateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVsmfUpdateErrorWithDefaults

`func NewVsmfUpdateErrorWithDefaults() *VsmfUpdateError`

NewVsmfUpdateErrorWithDefaults instantiates a new VsmfUpdateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *VsmfUpdateError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VsmfUpdateError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VsmfUpdateError) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetPti

`func (o *VsmfUpdateError) GetPti() int32`

GetPti returns the Pti field if non-nil, zero value otherwise.

### GetPtiOk

`func (o *VsmfUpdateError) GetPtiOk() (*int32, bool)`

GetPtiOk returns a tuple with the Pti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPti

`func (o *VsmfUpdateError) SetPti(v int32)`

SetPti sets Pti field to given value.

### HasPti

`func (o *VsmfUpdateError) HasPti() bool`

HasPti returns a boolean if a field has been set.

### GetN1smCause

`func (o *VsmfUpdateError) GetN1smCause() string`

GetN1smCause returns the N1smCause field if non-nil, zero value otherwise.

### GetN1smCauseOk

`func (o *VsmfUpdateError) GetN1smCauseOk() (*string, bool)`

GetN1smCauseOk returns a tuple with the N1smCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1smCause

`func (o *VsmfUpdateError) SetN1smCause(v string)`

SetN1smCause sets N1smCause field to given value.

### HasN1smCause

`func (o *VsmfUpdateError) HasN1smCause() bool`

HasN1smCause returns a boolean if a field has been set.

### GetN1SmInfoFromUe

`func (o *VsmfUpdateError) GetN1SmInfoFromUe() RefToBinaryData`

GetN1SmInfoFromUe returns the N1SmInfoFromUe field if non-nil, zero value otherwise.

### GetN1SmInfoFromUeOk

`func (o *VsmfUpdateError) GetN1SmInfoFromUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoFromUeOk returns a tuple with the N1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoFromUe

`func (o *VsmfUpdateError) SetN1SmInfoFromUe(v RefToBinaryData)`

SetN1SmInfoFromUe sets N1SmInfoFromUe field to given value.

### HasN1SmInfoFromUe

`func (o *VsmfUpdateError) HasN1SmInfoFromUe() bool`

HasN1SmInfoFromUe returns a boolean if a field has been set.

### GetUnknownN1SmInfo

`func (o *VsmfUpdateError) GetUnknownN1SmInfo() RefToBinaryData`

GetUnknownN1SmInfo returns the UnknownN1SmInfo field if non-nil, zero value otherwise.

### GetUnknownN1SmInfoOk

`func (o *VsmfUpdateError) GetUnknownN1SmInfoOk() (*RefToBinaryData, bool)`

GetUnknownN1SmInfoOk returns a tuple with the UnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownN1SmInfo

`func (o *VsmfUpdateError) SetUnknownN1SmInfo(v RefToBinaryData)`

SetUnknownN1SmInfo sets UnknownN1SmInfo field to given value.

### HasUnknownN1SmInfo

`func (o *VsmfUpdateError) HasUnknownN1SmInfo() bool`

HasUnknownN1SmInfo returns a boolean if a field has been set.

### GetFailedToAssignEbiList

`func (o *VsmfUpdateError) GetFailedToAssignEbiList() []Arp`

GetFailedToAssignEbiList returns the FailedToAssignEbiList field if non-nil, zero value otherwise.

### GetFailedToAssignEbiListOk

`func (o *VsmfUpdateError) GetFailedToAssignEbiListOk() (*[]Arp, bool)`

GetFailedToAssignEbiListOk returns a tuple with the FailedToAssignEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToAssignEbiList

`func (o *VsmfUpdateError) SetFailedToAssignEbiList(v []Arp)`

SetFailedToAssignEbiList sets FailedToAssignEbiList field to given value.

### HasFailedToAssignEbiList

`func (o *VsmfUpdateError) HasFailedToAssignEbiList() bool`

HasFailedToAssignEbiList returns a boolean if a field has been set.

### GetNgApCause

`func (o *VsmfUpdateError) GetNgApCause() NgApCause`

GetNgApCause returns the NgApCause field if non-nil, zero value otherwise.

### GetNgApCauseOk

`func (o *VsmfUpdateError) GetNgApCauseOk() (*NgApCause, bool)`

GetNgApCauseOk returns a tuple with the NgApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgApCause

`func (o *VsmfUpdateError) SetNgApCause(v NgApCause)`

SetNgApCause sets NgApCause field to given value.

### HasNgApCause

`func (o *VsmfUpdateError) HasNgApCause() bool`

HasNgApCause returns a boolean if a field has been set.

### GetVar5gMmCauseValue

`func (o *VsmfUpdateError) GetVar5gMmCauseValue() int32`

GetVar5gMmCauseValue returns the Var5gMmCauseValue field if non-nil, zero value otherwise.

### GetVar5gMmCauseValueOk

`func (o *VsmfUpdateError) GetVar5gMmCauseValueOk() (*int32, bool)`

GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCauseValue

`func (o *VsmfUpdateError) SetVar5gMmCauseValue(v int32)`

SetVar5gMmCauseValue sets Var5gMmCauseValue field to given value.

### HasVar5gMmCauseValue

`func (o *VsmfUpdateError) HasVar5gMmCauseValue() bool`

HasVar5gMmCauseValue returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *VsmfUpdateError) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *VsmfUpdateError) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *VsmfUpdateError) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *VsmfUpdateError) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetN4Info

`func (o *VsmfUpdateError) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *VsmfUpdateError) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *VsmfUpdateError) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *VsmfUpdateError) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *VsmfUpdateError) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *VsmfUpdateError) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *VsmfUpdateError) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *VsmfUpdateError) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *VsmfUpdateError) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *VsmfUpdateError) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *VsmfUpdateError) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *VsmfUpdateError) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


