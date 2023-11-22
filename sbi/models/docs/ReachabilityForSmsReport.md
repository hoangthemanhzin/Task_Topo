# ReachabilityForSmsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsfAccessType** | [**AccessType**](AccessType.md) |  | 
**MaxAvailabilityTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReachabilityForSmsReport

`func NewReachabilityForSmsReport(smsfAccessType AccessType, ) *ReachabilityForSmsReport`

NewReachabilityForSmsReport instantiates a new ReachabilityForSmsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityForSmsReportWithDefaults

`func NewReachabilityForSmsReportWithDefaults() *ReachabilityForSmsReport`

NewReachabilityForSmsReportWithDefaults instantiates a new ReachabilityForSmsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsfAccessType

`func (o *ReachabilityForSmsReport) GetSmsfAccessType() AccessType`

GetSmsfAccessType returns the SmsfAccessType field if non-nil, zero value otherwise.

### GetSmsfAccessTypeOk

`func (o *ReachabilityForSmsReport) GetSmsfAccessTypeOk() (*AccessType, bool)`

GetSmsfAccessTypeOk returns a tuple with the SmsfAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfAccessType

`func (o *ReachabilityForSmsReport) SetSmsfAccessType(v AccessType)`

SetSmsfAccessType sets SmsfAccessType field to given value.


### GetMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) GetMaxAvailabilityTime() time.Time`

GetMaxAvailabilityTime returns the MaxAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxAvailabilityTimeOk

`func (o *ReachabilityForSmsReport) GetMaxAvailabilityTimeOk() (*time.Time, bool)`

GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) SetMaxAvailabilityTime(v time.Time)`

SetMaxAvailabilityTime sets MaxAvailabilityTime field to given value.

### HasMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) HasMaxAvailabilityTime() bool`

HasMaxAvailabilityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


