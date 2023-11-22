# PartialSuccessReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCause** | [**FailureCause**](FailureCause.md) |  | 
**RuleReports** | Pointer to [**[]RuleReport**](RuleReport.md) | Information about the PCC rules provisioned by the PCF not successfully installed/activated. | [optional] 
**SessRuleReports** | Pointer to [**[]SessionRuleReport**](SessionRuleReport.md) | Information about the session rules provisioned by the PCF not successfully installed. | [optional] 
**UeCampingRep** | Pointer to [**UeCampingRep**](UeCampingRep.md) |  | [optional] 
**PolicyDecFailureReports** | Pointer to [**[]PolicyDecisionFailureCode**](PolicyDecisionFailureCode.md) | Contains the type(s) of failed policy decision and/or condition data. | [optional] 

## Methods

### NewPartialSuccessReport

`func NewPartialSuccessReport(failureCause FailureCause, ) *PartialSuccessReport`

NewPartialSuccessReport instantiates a new PartialSuccessReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialSuccessReportWithDefaults

`func NewPartialSuccessReportWithDefaults() *PartialSuccessReport`

NewPartialSuccessReportWithDefaults instantiates a new PartialSuccessReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCause

`func (o *PartialSuccessReport) GetFailureCause() FailureCause`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *PartialSuccessReport) GetFailureCauseOk() (*FailureCause, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *PartialSuccessReport) SetFailureCause(v FailureCause)`

SetFailureCause sets FailureCause field to given value.


### GetRuleReports

`func (o *PartialSuccessReport) GetRuleReports() []RuleReport`

GetRuleReports returns the RuleReports field if non-nil, zero value otherwise.

### GetRuleReportsOk

`func (o *PartialSuccessReport) GetRuleReportsOk() (*[]RuleReport, bool)`

GetRuleReportsOk returns a tuple with the RuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleReports

`func (o *PartialSuccessReport) SetRuleReports(v []RuleReport)`

SetRuleReports sets RuleReports field to given value.

### HasRuleReports

`func (o *PartialSuccessReport) HasRuleReports() bool`

HasRuleReports returns a boolean if a field has been set.

### GetSessRuleReports

`func (o *PartialSuccessReport) GetSessRuleReports() []SessionRuleReport`

GetSessRuleReports returns the SessRuleReports field if non-nil, zero value otherwise.

### GetSessRuleReportsOk

`func (o *PartialSuccessReport) GetSessRuleReportsOk() (*[]SessionRuleReport, bool)`

GetSessRuleReportsOk returns a tuple with the SessRuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRuleReports

`func (o *PartialSuccessReport) SetSessRuleReports(v []SessionRuleReport)`

SetSessRuleReports sets SessRuleReports field to given value.

### HasSessRuleReports

`func (o *PartialSuccessReport) HasSessRuleReports() bool`

HasSessRuleReports returns a boolean if a field has been set.

### GetUeCampingRep

`func (o *PartialSuccessReport) GetUeCampingRep() UeCampingRep`

GetUeCampingRep returns the UeCampingRep field if non-nil, zero value otherwise.

### GetUeCampingRepOk

`func (o *PartialSuccessReport) GetUeCampingRepOk() (*UeCampingRep, bool)`

GetUeCampingRepOk returns a tuple with the UeCampingRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCampingRep

`func (o *PartialSuccessReport) SetUeCampingRep(v UeCampingRep)`

SetUeCampingRep sets UeCampingRep field to given value.

### HasUeCampingRep

`func (o *PartialSuccessReport) HasUeCampingRep() bool`

HasUeCampingRep returns a boolean if a field has been set.

### GetPolicyDecFailureReports

`func (o *PartialSuccessReport) GetPolicyDecFailureReports() []PolicyDecisionFailureCode`

GetPolicyDecFailureReports returns the PolicyDecFailureReports field if non-nil, zero value otherwise.

### GetPolicyDecFailureReportsOk

`func (o *PartialSuccessReport) GetPolicyDecFailureReportsOk() (*[]PolicyDecisionFailureCode, bool)`

GetPolicyDecFailureReportsOk returns a tuple with the PolicyDecFailureReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecFailureReports

`func (o *PartialSuccessReport) SetPolicyDecFailureReports(v []PolicyDecisionFailureCode)`

SetPolicyDecFailureReports sets PolicyDecFailureReports field to given value.

### HasPolicyDecFailureReports

`func (o *PartialSuccessReport) HasPolicyDecFailureReports() bool`

HasPolicyDecFailureReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


