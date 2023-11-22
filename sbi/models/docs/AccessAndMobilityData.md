# AccessAndMobilityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**LocationTs** | Pointer to **time.Time** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**TimeZoneTs** | Pointer to **time.Time** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RegStates** | Pointer to [**[]RmInfo**](RmInfo.md) |  | [optional] 
**RegStatesTs** | Pointer to **time.Time** |  | [optional] 
**ConnStates** | Pointer to [**[]CmInfo**](CmInfo.md) |  | [optional] 
**ConnStatesTs** | Pointer to **time.Time** |  | [optional] 
**ReachabilityStatus** | Pointer to [**UeReachability**](UeReachability.md) |  | [optional] 
**ReachabilityStatusTs** | Pointer to **time.Time** |  | [optional] 
**SmsOverNasStatus** | Pointer to [**SmsSupport**](SmsSupport.md) |  | [optional] 
**SmsOverNasStatusTs** | Pointer to **time.Time** |  | [optional] 
**RoamingStatus** | Pointer to **bool** | True  The serving PLMN of the UE is different from the HPLMN of the UE; False  The serving PLMN of the UE is the HPLMN of the UE. | [optional] 
**RoamingStatusTs** | Pointer to **time.Time** |  | [optional] 
**CurrentPlmn** | Pointer to [**PlmnId1**](PlmnId1.md) |  | [optional] 
**CurrentPlmnTs** | Pointer to **time.Time** |  | [optional] 
**RatType** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**RatTypesTs** | Pointer to **time.Time** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessAndMobilityData

`func NewAccessAndMobilityData() *AccessAndMobilityData`

NewAccessAndMobilityData instantiates a new AccessAndMobilityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessAndMobilityDataWithDefaults

`func NewAccessAndMobilityDataWithDefaults() *AccessAndMobilityData`

NewAccessAndMobilityDataWithDefaults instantiates a new AccessAndMobilityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *AccessAndMobilityData) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AccessAndMobilityData) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AccessAndMobilityData) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AccessAndMobilityData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationTs

`func (o *AccessAndMobilityData) GetLocationTs() time.Time`

GetLocationTs returns the LocationTs field if non-nil, zero value otherwise.

### GetLocationTsOk

`func (o *AccessAndMobilityData) GetLocationTsOk() (*time.Time, bool)`

GetLocationTsOk returns a tuple with the LocationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTs

`func (o *AccessAndMobilityData) SetLocationTs(v time.Time)`

SetLocationTs sets LocationTs field to given value.

### HasLocationTs

`func (o *AccessAndMobilityData) HasLocationTs() bool`

HasLocationTs returns a boolean if a field has been set.

### GetTimeZone

`func (o *AccessAndMobilityData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AccessAndMobilityData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AccessAndMobilityData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AccessAndMobilityData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTimeZoneTs

`func (o *AccessAndMobilityData) GetTimeZoneTs() time.Time`

GetTimeZoneTs returns the TimeZoneTs field if non-nil, zero value otherwise.

### GetTimeZoneTsOk

`func (o *AccessAndMobilityData) GetTimeZoneTsOk() (*time.Time, bool)`

GetTimeZoneTsOk returns a tuple with the TimeZoneTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneTs

`func (o *AccessAndMobilityData) SetTimeZoneTs(v time.Time)`

SetTimeZoneTs sets TimeZoneTs field to given value.

### HasTimeZoneTs

`func (o *AccessAndMobilityData) HasTimeZoneTs() bool`

HasTimeZoneTs returns a boolean if a field has been set.

### GetAccessType

`func (o *AccessAndMobilityData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AccessAndMobilityData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AccessAndMobilityData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AccessAndMobilityData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRegStates

`func (o *AccessAndMobilityData) GetRegStates() []RmInfo`

GetRegStates returns the RegStates field if non-nil, zero value otherwise.

### GetRegStatesOk

`func (o *AccessAndMobilityData) GetRegStatesOk() (*[]RmInfo, bool)`

GetRegStatesOk returns a tuple with the RegStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegStates

`func (o *AccessAndMobilityData) SetRegStates(v []RmInfo)`

SetRegStates sets RegStates field to given value.

### HasRegStates

`func (o *AccessAndMobilityData) HasRegStates() bool`

HasRegStates returns a boolean if a field has been set.

### GetRegStatesTs

`func (o *AccessAndMobilityData) GetRegStatesTs() time.Time`

GetRegStatesTs returns the RegStatesTs field if non-nil, zero value otherwise.

### GetRegStatesTsOk

`func (o *AccessAndMobilityData) GetRegStatesTsOk() (*time.Time, bool)`

GetRegStatesTsOk returns a tuple with the RegStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegStatesTs

`func (o *AccessAndMobilityData) SetRegStatesTs(v time.Time)`

SetRegStatesTs sets RegStatesTs field to given value.

### HasRegStatesTs

`func (o *AccessAndMobilityData) HasRegStatesTs() bool`

HasRegStatesTs returns a boolean if a field has been set.

### GetConnStates

`func (o *AccessAndMobilityData) GetConnStates() []CmInfo`

GetConnStates returns the ConnStates field if non-nil, zero value otherwise.

### GetConnStatesOk

`func (o *AccessAndMobilityData) GetConnStatesOk() (*[]CmInfo, bool)`

GetConnStatesOk returns a tuple with the ConnStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnStates

`func (o *AccessAndMobilityData) SetConnStates(v []CmInfo)`

SetConnStates sets ConnStates field to given value.

### HasConnStates

`func (o *AccessAndMobilityData) HasConnStates() bool`

HasConnStates returns a boolean if a field has been set.

### GetConnStatesTs

`func (o *AccessAndMobilityData) GetConnStatesTs() time.Time`

GetConnStatesTs returns the ConnStatesTs field if non-nil, zero value otherwise.

### GetConnStatesTsOk

`func (o *AccessAndMobilityData) GetConnStatesTsOk() (*time.Time, bool)`

GetConnStatesTsOk returns a tuple with the ConnStatesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnStatesTs

`func (o *AccessAndMobilityData) SetConnStatesTs(v time.Time)`

SetConnStatesTs sets ConnStatesTs field to given value.

### HasConnStatesTs

`func (o *AccessAndMobilityData) HasConnStatesTs() bool`

HasConnStatesTs returns a boolean if a field has been set.

### GetReachabilityStatus

`func (o *AccessAndMobilityData) GetReachabilityStatus() UeReachability`

GetReachabilityStatus returns the ReachabilityStatus field if non-nil, zero value otherwise.

### GetReachabilityStatusOk

`func (o *AccessAndMobilityData) GetReachabilityStatusOk() (*UeReachability, bool)`

GetReachabilityStatusOk returns a tuple with the ReachabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityStatus

`func (o *AccessAndMobilityData) SetReachabilityStatus(v UeReachability)`

SetReachabilityStatus sets ReachabilityStatus field to given value.

### HasReachabilityStatus

`func (o *AccessAndMobilityData) HasReachabilityStatus() bool`

HasReachabilityStatus returns a boolean if a field has been set.

### GetReachabilityStatusTs

`func (o *AccessAndMobilityData) GetReachabilityStatusTs() time.Time`

GetReachabilityStatusTs returns the ReachabilityStatusTs field if non-nil, zero value otherwise.

### GetReachabilityStatusTsOk

`func (o *AccessAndMobilityData) GetReachabilityStatusTsOk() (*time.Time, bool)`

GetReachabilityStatusTsOk returns a tuple with the ReachabilityStatusTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityStatusTs

`func (o *AccessAndMobilityData) SetReachabilityStatusTs(v time.Time)`

SetReachabilityStatusTs sets ReachabilityStatusTs field to given value.

### HasReachabilityStatusTs

`func (o *AccessAndMobilityData) HasReachabilityStatusTs() bool`

HasReachabilityStatusTs returns a boolean if a field has been set.

### GetSmsOverNasStatus

`func (o *AccessAndMobilityData) GetSmsOverNasStatus() SmsSupport`

GetSmsOverNasStatus returns the SmsOverNasStatus field if non-nil, zero value otherwise.

### GetSmsOverNasStatusOk

`func (o *AccessAndMobilityData) GetSmsOverNasStatusOk() (*SmsSupport, bool)`

GetSmsOverNasStatusOk returns a tuple with the SmsOverNasStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsOverNasStatus

`func (o *AccessAndMobilityData) SetSmsOverNasStatus(v SmsSupport)`

SetSmsOverNasStatus sets SmsOverNasStatus field to given value.

### HasSmsOverNasStatus

`func (o *AccessAndMobilityData) HasSmsOverNasStatus() bool`

HasSmsOverNasStatus returns a boolean if a field has been set.

### GetSmsOverNasStatusTs

`func (o *AccessAndMobilityData) GetSmsOverNasStatusTs() time.Time`

GetSmsOverNasStatusTs returns the SmsOverNasStatusTs field if non-nil, zero value otherwise.

### GetSmsOverNasStatusTsOk

`func (o *AccessAndMobilityData) GetSmsOverNasStatusTsOk() (*time.Time, bool)`

GetSmsOverNasStatusTsOk returns a tuple with the SmsOverNasStatusTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsOverNasStatusTs

`func (o *AccessAndMobilityData) SetSmsOverNasStatusTs(v time.Time)`

SetSmsOverNasStatusTs sets SmsOverNasStatusTs field to given value.

### HasSmsOverNasStatusTs

`func (o *AccessAndMobilityData) HasSmsOverNasStatusTs() bool`

HasSmsOverNasStatusTs returns a boolean if a field has been set.

### GetRoamingStatus

`func (o *AccessAndMobilityData) GetRoamingStatus() bool`

GetRoamingStatus returns the RoamingStatus field if non-nil, zero value otherwise.

### GetRoamingStatusOk

`func (o *AccessAndMobilityData) GetRoamingStatusOk() (*bool, bool)`

GetRoamingStatusOk returns a tuple with the RoamingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingStatus

`func (o *AccessAndMobilityData) SetRoamingStatus(v bool)`

SetRoamingStatus sets RoamingStatus field to given value.

### HasRoamingStatus

`func (o *AccessAndMobilityData) HasRoamingStatus() bool`

HasRoamingStatus returns a boolean if a field has been set.

### GetRoamingStatusTs

`func (o *AccessAndMobilityData) GetRoamingStatusTs() time.Time`

GetRoamingStatusTs returns the RoamingStatusTs field if non-nil, zero value otherwise.

### GetRoamingStatusTsOk

`func (o *AccessAndMobilityData) GetRoamingStatusTsOk() (*time.Time, bool)`

GetRoamingStatusTsOk returns a tuple with the RoamingStatusTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingStatusTs

`func (o *AccessAndMobilityData) SetRoamingStatusTs(v time.Time)`

SetRoamingStatusTs sets RoamingStatusTs field to given value.

### HasRoamingStatusTs

`func (o *AccessAndMobilityData) HasRoamingStatusTs() bool`

HasRoamingStatusTs returns a boolean if a field has been set.

### GetCurrentPlmn

`func (o *AccessAndMobilityData) GetCurrentPlmn() PlmnId1`

GetCurrentPlmn returns the CurrentPlmn field if non-nil, zero value otherwise.

### GetCurrentPlmnOk

`func (o *AccessAndMobilityData) GetCurrentPlmnOk() (*PlmnId1, bool)`

GetCurrentPlmnOk returns a tuple with the CurrentPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlmn

`func (o *AccessAndMobilityData) SetCurrentPlmn(v PlmnId1)`

SetCurrentPlmn sets CurrentPlmn field to given value.

### HasCurrentPlmn

`func (o *AccessAndMobilityData) HasCurrentPlmn() bool`

HasCurrentPlmn returns a boolean if a field has been set.

### GetCurrentPlmnTs

`func (o *AccessAndMobilityData) GetCurrentPlmnTs() time.Time`

GetCurrentPlmnTs returns the CurrentPlmnTs field if non-nil, zero value otherwise.

### GetCurrentPlmnTsOk

`func (o *AccessAndMobilityData) GetCurrentPlmnTsOk() (*time.Time, bool)`

GetCurrentPlmnTsOk returns a tuple with the CurrentPlmnTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlmnTs

`func (o *AccessAndMobilityData) SetCurrentPlmnTs(v time.Time)`

SetCurrentPlmnTs sets CurrentPlmnTs field to given value.

### HasCurrentPlmnTs

`func (o *AccessAndMobilityData) HasCurrentPlmnTs() bool`

HasCurrentPlmnTs returns a boolean if a field has been set.

### GetRatType

`func (o *AccessAndMobilityData) GetRatType() []RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *AccessAndMobilityData) GetRatTypeOk() (*[]RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *AccessAndMobilityData) SetRatType(v []RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *AccessAndMobilityData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetRatTypesTs

`func (o *AccessAndMobilityData) GetRatTypesTs() time.Time`

GetRatTypesTs returns the RatTypesTs field if non-nil, zero value otherwise.

### GetRatTypesTsOk

`func (o *AccessAndMobilityData) GetRatTypesTsOk() (*time.Time, bool)`

GetRatTypesTsOk returns a tuple with the RatTypesTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypesTs

`func (o *AccessAndMobilityData) SetRatTypesTs(v time.Time)`

SetRatTypesTs sets RatTypesTs field to given value.

### HasRatTypesTs

`func (o *AccessAndMobilityData) HasRatTypesTs() bool`

HasRatTypesTs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AccessAndMobilityData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AccessAndMobilityData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AccessAndMobilityData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AccessAndMobilityData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


