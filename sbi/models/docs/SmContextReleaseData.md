# SmContextReleaseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**NgApCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**Var5gMmCauseValue** | Pointer to **int32** |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**VsmfReleaseOnly** | Pointer to **bool** |  | [optional] [default to false]
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**IsmfReleaseOnly** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSmContextReleaseData

`func NewSmContextReleaseData() *SmContextReleaseData`

NewSmContextReleaseData instantiates a new SmContextReleaseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextReleaseDataWithDefaults

`func NewSmContextReleaseDataWithDefaults() *SmContextReleaseData`

NewSmContextReleaseDataWithDefaults instantiates a new SmContextReleaseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *SmContextReleaseData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SmContextReleaseData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SmContextReleaseData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *SmContextReleaseData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetNgApCause

`func (o *SmContextReleaseData) GetNgApCause() NgApCause`

GetNgApCause returns the NgApCause field if non-nil, zero value otherwise.

### GetNgApCauseOk

`func (o *SmContextReleaseData) GetNgApCauseOk() (*NgApCause, bool)`

GetNgApCauseOk returns a tuple with the NgApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgApCause

`func (o *SmContextReleaseData) SetNgApCause(v NgApCause)`

SetNgApCause sets NgApCause field to given value.

### HasNgApCause

`func (o *SmContextReleaseData) HasNgApCause() bool`

HasNgApCause returns a boolean if a field has been set.

### GetVar5gMmCauseValue

`func (o *SmContextReleaseData) GetVar5gMmCauseValue() int32`

GetVar5gMmCauseValue returns the Var5gMmCauseValue field if non-nil, zero value otherwise.

### GetVar5gMmCauseValueOk

`func (o *SmContextReleaseData) GetVar5gMmCauseValueOk() (*int32, bool)`

GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCauseValue

`func (o *SmContextReleaseData) SetVar5gMmCauseValue(v int32)`

SetVar5gMmCauseValue sets Var5gMmCauseValue field to given value.

### HasVar5gMmCauseValue

`func (o *SmContextReleaseData) HasVar5gMmCauseValue() bool`

HasVar5gMmCauseValue returns a boolean if a field has been set.

### GetUeLocation

`func (o *SmContextReleaseData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SmContextReleaseData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SmContextReleaseData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SmContextReleaseData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmContextReleaseData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmContextReleaseData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmContextReleaseData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmContextReleaseData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *SmContextReleaseData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *SmContextReleaseData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *SmContextReleaseData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *SmContextReleaseData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetVsmfReleaseOnly

`func (o *SmContextReleaseData) GetVsmfReleaseOnly() bool`

GetVsmfReleaseOnly returns the VsmfReleaseOnly field if non-nil, zero value otherwise.

### GetVsmfReleaseOnlyOk

`func (o *SmContextReleaseData) GetVsmfReleaseOnlyOk() (*bool, bool)`

GetVsmfReleaseOnlyOk returns a tuple with the VsmfReleaseOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfReleaseOnly

`func (o *SmContextReleaseData) SetVsmfReleaseOnly(v bool)`

SetVsmfReleaseOnly sets VsmfReleaseOnly field to given value.

### HasVsmfReleaseOnly

`func (o *SmContextReleaseData) HasVsmfReleaseOnly() bool`

HasVsmfReleaseOnly returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextReleaseData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextReleaseData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextReleaseData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextReleaseData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextReleaseData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextReleaseData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextReleaseData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextReleaseData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetIsmfReleaseOnly

`func (o *SmContextReleaseData) GetIsmfReleaseOnly() bool`

GetIsmfReleaseOnly returns the IsmfReleaseOnly field if non-nil, zero value otherwise.

### GetIsmfReleaseOnlyOk

`func (o *SmContextReleaseData) GetIsmfReleaseOnlyOk() (*bool, bool)`

GetIsmfReleaseOnlyOk returns a tuple with the IsmfReleaseOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfReleaseOnly

`func (o *SmContextReleaseData) SetIsmfReleaseOnly(v bool)`

SetIsmfReleaseOnly sets IsmfReleaseOnly field to given value.

### HasIsmfReleaseOnly

`func (o *SmContextReleaseData) HasIsmfReleaseOnly() bool`

HasIsmfReleaseOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


