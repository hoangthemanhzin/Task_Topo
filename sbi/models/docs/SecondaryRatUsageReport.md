# SecondaryRatUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondaryRatType** | [**RatType**](RatType.md) |  | 
**QosFlowsUsageData** | [**[]QosFlowUsageReport**](QosFlowUsageReport.md) |  | 

## Methods

### NewSecondaryRatUsageReport

`func NewSecondaryRatUsageReport(secondaryRatType RatType, qosFlowsUsageData []QosFlowUsageReport, ) *SecondaryRatUsageReport`

NewSecondaryRatUsageReport instantiates a new SecondaryRatUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryRatUsageReportWithDefaults

`func NewSecondaryRatUsageReportWithDefaults() *SecondaryRatUsageReport`

NewSecondaryRatUsageReportWithDefaults instantiates a new SecondaryRatUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondaryRatType

`func (o *SecondaryRatUsageReport) GetSecondaryRatType() RatType`

GetSecondaryRatType returns the SecondaryRatType field if non-nil, zero value otherwise.

### GetSecondaryRatTypeOk

`func (o *SecondaryRatUsageReport) GetSecondaryRatTypeOk() (*RatType, bool)`

GetSecondaryRatTypeOk returns a tuple with the SecondaryRatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatType

`func (o *SecondaryRatUsageReport) SetSecondaryRatType(v RatType)`

SetSecondaryRatType sets SecondaryRatType field to given value.


### GetQosFlowsUsageData

`func (o *SecondaryRatUsageReport) GetQosFlowsUsageData() []QosFlowUsageReport`

GetQosFlowsUsageData returns the QosFlowsUsageData field if non-nil, zero value otherwise.

### GetQosFlowsUsageDataOk

`func (o *SecondaryRatUsageReport) GetQosFlowsUsageDataOk() (*[]QosFlowUsageReport, bool)`

GetQosFlowsUsageDataOk returns a tuple with the QosFlowsUsageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsUsageData

`func (o *SecondaryRatUsageReport) SetQosFlowsUsageData(v []QosFlowUsageReport)`

SetQosFlowsUsageData sets QosFlowsUsageData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


