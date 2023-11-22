# SmPolicyDeleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLocationInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**UserLocationInfoTime** | Pointer to **time.Time** |  | [optional] 
**RanNasRelCauses** | Pointer to [**[]RanNasRelCause**](RanNasRelCause.md) | Contains the RAN and/or NAS release cause. | [optional] 
**AccuUsageReports** | Pointer to [**[]AccuUsageReport**](AccuUsageReport.md) | Contains the usage report | [optional] 
**PduSessRelCause** | Pointer to [**PduSessionRelCause**](PduSessionRelCause.md) |  | [optional] 
**QosMonReports** | Pointer to [**[]QosMonitoringReport**](QosMonitoringReport.md) |  | [optional] 

## Methods

### NewSmPolicyDeleteData

`func NewSmPolicyDeleteData() *SmPolicyDeleteData`

NewSmPolicyDeleteData instantiates a new SmPolicyDeleteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDeleteDataWithDefaults

`func NewSmPolicyDeleteDataWithDefaults() *SmPolicyDeleteData`

NewSmPolicyDeleteDataWithDefaults instantiates a new SmPolicyDeleteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLocationInfo

`func (o *SmPolicyDeleteData) GetUserLocationInfo() UserLocation`

GetUserLocationInfo returns the UserLocationInfo field if non-nil, zero value otherwise.

### GetUserLocationInfoOk

`func (o *SmPolicyDeleteData) GetUserLocationInfoOk() (*UserLocation, bool)`

GetUserLocationInfoOk returns a tuple with the UserLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfo

`func (o *SmPolicyDeleteData) SetUserLocationInfo(v UserLocation)`

SetUserLocationInfo sets UserLocationInfo field to given value.

### HasUserLocationInfo

`func (o *SmPolicyDeleteData) HasUserLocationInfo() bool`

HasUserLocationInfo returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmPolicyDeleteData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmPolicyDeleteData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmPolicyDeleteData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmPolicyDeleteData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmPolicyDeleteData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmPolicyDeleteData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmPolicyDeleteData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *SmPolicyDeleteData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetUserLocationInfoTime

`func (o *SmPolicyDeleteData) GetUserLocationInfoTime() time.Time`

GetUserLocationInfoTime returns the UserLocationInfoTime field if non-nil, zero value otherwise.

### GetUserLocationInfoTimeOk

`func (o *SmPolicyDeleteData) GetUserLocationInfoTimeOk() (*time.Time, bool)`

GetUserLocationInfoTimeOk returns a tuple with the UserLocationInfoTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfoTime

`func (o *SmPolicyDeleteData) SetUserLocationInfoTime(v time.Time)`

SetUserLocationInfoTime sets UserLocationInfoTime field to given value.

### HasUserLocationInfoTime

`func (o *SmPolicyDeleteData) HasUserLocationInfoTime() bool`

HasUserLocationInfoTime returns a boolean if a field has been set.

### GetRanNasRelCauses

`func (o *SmPolicyDeleteData) GetRanNasRelCauses() []RanNasRelCause`

GetRanNasRelCauses returns the RanNasRelCauses field if non-nil, zero value otherwise.

### GetRanNasRelCausesOk

`func (o *SmPolicyDeleteData) GetRanNasRelCausesOk() (*[]RanNasRelCause, bool)`

GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNasRelCauses

`func (o *SmPolicyDeleteData) SetRanNasRelCauses(v []RanNasRelCause)`

SetRanNasRelCauses sets RanNasRelCauses field to given value.

### HasRanNasRelCauses

`func (o *SmPolicyDeleteData) HasRanNasRelCauses() bool`

HasRanNasRelCauses returns a boolean if a field has been set.

### GetAccuUsageReports

`func (o *SmPolicyDeleteData) GetAccuUsageReports() []AccuUsageReport`

GetAccuUsageReports returns the AccuUsageReports field if non-nil, zero value otherwise.

### GetAccuUsageReportsOk

`func (o *SmPolicyDeleteData) GetAccuUsageReportsOk() (*[]AccuUsageReport, bool)`

GetAccuUsageReportsOk returns a tuple with the AccuUsageReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuUsageReports

`func (o *SmPolicyDeleteData) SetAccuUsageReports(v []AccuUsageReport)`

SetAccuUsageReports sets AccuUsageReports field to given value.

### HasAccuUsageReports

`func (o *SmPolicyDeleteData) HasAccuUsageReports() bool`

HasAccuUsageReports returns a boolean if a field has been set.

### GetPduSessRelCause

`func (o *SmPolicyDeleteData) GetPduSessRelCause() PduSessionRelCause`

GetPduSessRelCause returns the PduSessRelCause field if non-nil, zero value otherwise.

### GetPduSessRelCauseOk

`func (o *SmPolicyDeleteData) GetPduSessRelCauseOk() (*PduSessionRelCause, bool)`

GetPduSessRelCauseOk returns a tuple with the PduSessRelCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessRelCause

`func (o *SmPolicyDeleteData) SetPduSessRelCause(v PduSessionRelCause)`

SetPduSessRelCause sets PduSessRelCause field to given value.

### HasPduSessRelCause

`func (o *SmPolicyDeleteData) HasPduSessRelCause() bool`

HasPduSessRelCause returns a boolean if a field has been set.

### GetQosMonReports

`func (o *SmPolicyDeleteData) GetQosMonReports() []QosMonitoringReport`

GetQosMonReports returns the QosMonReports field if non-nil, zero value otherwise.

### GetQosMonReportsOk

`func (o *SmPolicyDeleteData) GetQosMonReportsOk() (*[]QosMonitoringReport, bool)`

GetQosMonReportsOk returns a tuple with the QosMonReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonReports

`func (o *SmPolicyDeleteData) SetQosMonReports(v []QosMonitoringReport)`

SetQosMonReports sets QosMonReports field to given value.

### HasQosMonReports

`func (o *SmPolicyDeleteData) HasQosMonReports() bool`

HasQosMonReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


