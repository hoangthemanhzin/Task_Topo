# QosMonitoringReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 
**UlDelays** | Pointer to **[]int32** |  | [optional] 
**DlDelays** | Pointer to **[]int32** |  | [optional] 
**RtDelays** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewQosMonitoringReport

`func NewQosMonitoringReport() *QosMonitoringReport`

NewQosMonitoringReport instantiates a new QosMonitoringReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosMonitoringReportWithDefaults

`func NewQosMonitoringReportWithDefaults() *QosMonitoringReport`

NewQosMonitoringReportWithDefaults instantiates a new QosMonitoringReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlows

`func (o *QosMonitoringReport) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *QosMonitoringReport) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *QosMonitoringReport) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *QosMonitoringReport) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetUlDelays

`func (o *QosMonitoringReport) GetUlDelays() []int32`

GetUlDelays returns the UlDelays field if non-nil, zero value otherwise.

### GetUlDelaysOk

`func (o *QosMonitoringReport) GetUlDelaysOk() (*[]int32, bool)`

GetUlDelaysOk returns a tuple with the UlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlDelays

`func (o *QosMonitoringReport) SetUlDelays(v []int32)`

SetUlDelays sets UlDelays field to given value.

### HasUlDelays

`func (o *QosMonitoringReport) HasUlDelays() bool`

HasUlDelays returns a boolean if a field has been set.

### GetDlDelays

`func (o *QosMonitoringReport) GetDlDelays() []int32`

GetDlDelays returns the DlDelays field if non-nil, zero value otherwise.

### GetDlDelaysOk

`func (o *QosMonitoringReport) GetDlDelaysOk() (*[]int32, bool)`

GetDlDelaysOk returns a tuple with the DlDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDelays

`func (o *QosMonitoringReport) SetDlDelays(v []int32)`

SetDlDelays sets DlDelays field to given value.

### HasDlDelays

`func (o *QosMonitoringReport) HasDlDelays() bool`

HasDlDelays returns a boolean if a field has been set.

### GetRtDelays

`func (o *QosMonitoringReport) GetRtDelays() []int32`

GetRtDelays returns the RtDelays field if non-nil, zero value otherwise.

### GetRtDelaysOk

`func (o *QosMonitoringReport) GetRtDelaysOk() (*[]int32, bool)`

GetRtDelaysOk returns a tuple with the RtDelays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtDelays

`func (o *QosMonitoringReport) SetRtDelays(v []int32)`

SetRtDelays sets RtDelays field to given value.

### HasRtDelays

`func (o *QosMonitoringReport) HasRtDelays() bool`

HasRtDelays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


