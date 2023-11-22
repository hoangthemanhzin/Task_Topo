# ErrorReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ProblemDetails**](ProblemDetails.md) |  | [optional] 
**RuleReports** | Pointer to [**[]RuleReport**](RuleReport.md) | Used to report the PCC rule failure. | [optional] 
**SessRuleReports** | Pointer to [**[]SessionRuleReport**](SessionRuleReport.md) | Used to report the session rule failure. | [optional] 
**PolDecFailureReports** | Pointer to [**[]PolicyDecisionFailureCode**](PolicyDecisionFailureCode.md) | Used to report failure of the policy decision and/or condition data. | [optional] 
**AltQosParamId** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorReport

`func NewErrorReport() *ErrorReport`

NewErrorReport instantiates a new ErrorReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorReportWithDefaults

`func NewErrorReportWithDefaults() *ErrorReport`

NewErrorReportWithDefaults instantiates a new ErrorReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorReport) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorReport) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorReport) SetError(v ProblemDetails)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorReport) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRuleReports

`func (o *ErrorReport) GetRuleReports() []RuleReport`

GetRuleReports returns the RuleReports field if non-nil, zero value otherwise.

### GetRuleReportsOk

`func (o *ErrorReport) GetRuleReportsOk() (*[]RuleReport, bool)`

GetRuleReportsOk returns a tuple with the RuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleReports

`func (o *ErrorReport) SetRuleReports(v []RuleReport)`

SetRuleReports sets RuleReports field to given value.

### HasRuleReports

`func (o *ErrorReport) HasRuleReports() bool`

HasRuleReports returns a boolean if a field has been set.

### GetSessRuleReports

`func (o *ErrorReport) GetSessRuleReports() []SessionRuleReport`

GetSessRuleReports returns the SessRuleReports field if non-nil, zero value otherwise.

### GetSessRuleReportsOk

`func (o *ErrorReport) GetSessRuleReportsOk() (*[]SessionRuleReport, bool)`

GetSessRuleReportsOk returns a tuple with the SessRuleReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRuleReports

`func (o *ErrorReport) SetSessRuleReports(v []SessionRuleReport)`

SetSessRuleReports sets SessRuleReports field to given value.

### HasSessRuleReports

`func (o *ErrorReport) HasSessRuleReports() bool`

HasSessRuleReports returns a boolean if a field has been set.

### GetPolDecFailureReports

`func (o *ErrorReport) GetPolDecFailureReports() []PolicyDecisionFailureCode`

GetPolDecFailureReports returns the PolDecFailureReports field if non-nil, zero value otherwise.

### GetPolDecFailureReportsOk

`func (o *ErrorReport) GetPolDecFailureReportsOk() (*[]PolicyDecisionFailureCode, bool)`

GetPolDecFailureReportsOk returns a tuple with the PolDecFailureReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDecFailureReports

`func (o *ErrorReport) SetPolDecFailureReports(v []PolicyDecisionFailureCode)`

SetPolDecFailureReports sets PolDecFailureReports field to given value.

### HasPolDecFailureReports

`func (o *ErrorReport) HasPolDecFailureReports() bool`

HasPolDecFailureReports returns a boolean if a field has been set.

### GetAltQosParamId

`func (o *ErrorReport) GetAltQosParamId() string`

GetAltQosParamId returns the AltQosParamId field if non-nil, zero value otherwise.

### GetAltQosParamIdOk

`func (o *ErrorReport) GetAltQosParamIdOk() (*string, bool)`

GetAltQosParamIdOk returns a tuple with the AltQosParamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosParamId

`func (o *ErrorReport) SetAltQosParamId(v string)`

SetAltQosParamId sets AltQosParamId field to given value.

### HasAltQosParamId

`func (o *ErrorReport) HasAltQosParamId() bool`

HasAltQosParamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


