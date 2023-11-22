# ReleaseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**NgApCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**Var5gMmCauseValue** | Pointer to **int32** |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**SecondaryRatUsageReport** | Pointer to [**[]SecondaryRatUsageReport**](SecondaryRatUsageReport.md) |  | [optional] 
**SecondaryRatUsageInfo** | Pointer to [**[]SecondaryRatUsageInfo**](SecondaryRatUsageInfo.md) |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 

## Methods

### NewReleaseData

`func NewReleaseData() *ReleaseData`

NewReleaseData instantiates a new ReleaseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseDataWithDefaults

`func NewReleaseDataWithDefaults() *ReleaseData`

NewReleaseDataWithDefaults instantiates a new ReleaseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *ReleaseData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ReleaseData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ReleaseData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ReleaseData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetNgApCause

`func (o *ReleaseData) GetNgApCause() NgApCause`

GetNgApCause returns the NgApCause field if non-nil, zero value otherwise.

### GetNgApCauseOk

`func (o *ReleaseData) GetNgApCauseOk() (*NgApCause, bool)`

GetNgApCauseOk returns a tuple with the NgApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgApCause

`func (o *ReleaseData) SetNgApCause(v NgApCause)`

SetNgApCause sets NgApCause field to given value.

### HasNgApCause

`func (o *ReleaseData) HasNgApCause() bool`

HasNgApCause returns a boolean if a field has been set.

### GetVar5gMmCauseValue

`func (o *ReleaseData) GetVar5gMmCauseValue() int32`

GetVar5gMmCauseValue returns the Var5gMmCauseValue field if non-nil, zero value otherwise.

### GetVar5gMmCauseValueOk

`func (o *ReleaseData) GetVar5gMmCauseValueOk() (*int32, bool)`

GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCauseValue

`func (o *ReleaseData) SetVar5gMmCauseValue(v int32)`

SetVar5gMmCauseValue sets Var5gMmCauseValue field to given value.

### HasVar5gMmCauseValue

`func (o *ReleaseData) HasVar5gMmCauseValue() bool`

HasVar5gMmCauseValue returns a boolean if a field has been set.

### GetUeLocation

`func (o *ReleaseData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *ReleaseData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *ReleaseData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *ReleaseData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *ReleaseData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *ReleaseData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *ReleaseData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *ReleaseData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *ReleaseData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *ReleaseData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *ReleaseData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *ReleaseData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetSecondaryRatUsageReport

`func (o *ReleaseData) GetSecondaryRatUsageReport() []SecondaryRatUsageReport`

GetSecondaryRatUsageReport returns the SecondaryRatUsageReport field if non-nil, zero value otherwise.

### GetSecondaryRatUsageReportOk

`func (o *ReleaseData) GetSecondaryRatUsageReportOk() (*[]SecondaryRatUsageReport, bool)`

GetSecondaryRatUsageReportOk returns a tuple with the SecondaryRatUsageReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageReport

`func (o *ReleaseData) SetSecondaryRatUsageReport(v []SecondaryRatUsageReport)`

SetSecondaryRatUsageReport sets SecondaryRatUsageReport field to given value.

### HasSecondaryRatUsageReport

`func (o *ReleaseData) HasSecondaryRatUsageReport() bool`

HasSecondaryRatUsageReport returns a boolean if a field has been set.

### GetSecondaryRatUsageInfo

`func (o *ReleaseData) GetSecondaryRatUsageInfo() []SecondaryRatUsageInfo`

GetSecondaryRatUsageInfo returns the SecondaryRatUsageInfo field if non-nil, zero value otherwise.

### GetSecondaryRatUsageInfoOk

`func (o *ReleaseData) GetSecondaryRatUsageInfoOk() (*[]SecondaryRatUsageInfo, bool)`

GetSecondaryRatUsageInfoOk returns a tuple with the SecondaryRatUsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageInfo

`func (o *ReleaseData) SetSecondaryRatUsageInfo(v []SecondaryRatUsageInfo)`

SetSecondaryRatUsageInfo sets SecondaryRatUsageInfo field to given value.

### HasSecondaryRatUsageInfo

`func (o *ReleaseData) HasSecondaryRatUsageInfo() bool`

HasSecondaryRatUsageInfo returns a boolean if a field has been set.

### GetN4Info

`func (o *ReleaseData) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *ReleaseData) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *ReleaseData) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *ReleaseData) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *ReleaseData) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *ReleaseData) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *ReleaseData) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *ReleaseData) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *ReleaseData) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *ReleaseData) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *ReleaseData) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *ReleaseData) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


