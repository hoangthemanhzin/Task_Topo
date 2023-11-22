# RuleReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PccRuleIds** | **[]string** | Contains the identifier of the affected PCC rule(s). | 
**RuleStatus** | [**RuleStatus**](RuleStatus.md) |  | 
**ContVers** | Pointer to **[]int32** | Indicates the version of a PCC rule. | [optional] 
**FailureCode** | Pointer to [**FailureCode**](FailureCode.md) |  | [optional] 
**FinUnitAct** | Pointer to [**FinalUnitAction**](FinalUnitAction.md) |  | [optional] 
**RanNasRelCauses** | Pointer to [**[]RanNasRelCause**](RanNasRelCause.md) | indicates the RAN or NAS release cause code information. | [optional] 

## Methods

### NewRuleReport

`func NewRuleReport(pccRuleIds []string, ruleStatus RuleStatus, ) *RuleReport`

NewRuleReport instantiates a new RuleReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleReportWithDefaults

`func NewRuleReportWithDefaults() *RuleReport`

NewRuleReportWithDefaults instantiates a new RuleReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPccRuleIds

`func (o *RuleReport) GetPccRuleIds() []string`

GetPccRuleIds returns the PccRuleIds field if non-nil, zero value otherwise.

### GetPccRuleIdsOk

`func (o *RuleReport) GetPccRuleIdsOk() (*[]string, bool)`

GetPccRuleIdsOk returns a tuple with the PccRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRuleIds

`func (o *RuleReport) SetPccRuleIds(v []string)`

SetPccRuleIds sets PccRuleIds field to given value.


### GetRuleStatus

`func (o *RuleReport) GetRuleStatus() RuleStatus`

GetRuleStatus returns the RuleStatus field if non-nil, zero value otherwise.

### GetRuleStatusOk

`func (o *RuleReport) GetRuleStatusOk() (*RuleStatus, bool)`

GetRuleStatusOk returns a tuple with the RuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatus

`func (o *RuleReport) SetRuleStatus(v RuleStatus)`

SetRuleStatus sets RuleStatus field to given value.


### GetContVers

`func (o *RuleReport) GetContVers() []int32`

GetContVers returns the ContVers field if non-nil, zero value otherwise.

### GetContVersOk

`func (o *RuleReport) GetContVersOk() (*[]int32, bool)`

GetContVersOk returns a tuple with the ContVers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVers

`func (o *RuleReport) SetContVers(v []int32)`

SetContVers sets ContVers field to given value.

### HasContVers

`func (o *RuleReport) HasContVers() bool`

HasContVers returns a boolean if a field has been set.

### GetFailureCode

`func (o *RuleReport) GetFailureCode() FailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *RuleReport) GetFailureCodeOk() (*FailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *RuleReport) SetFailureCode(v FailureCode)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *RuleReport) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFinUnitAct

`func (o *RuleReport) GetFinUnitAct() FinalUnitAction`

GetFinUnitAct returns the FinUnitAct field if non-nil, zero value otherwise.

### GetFinUnitActOk

`func (o *RuleReport) GetFinUnitActOk() (*FinalUnitAction, bool)`

GetFinUnitActOk returns a tuple with the FinUnitAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinUnitAct

`func (o *RuleReport) SetFinUnitAct(v FinalUnitAction)`

SetFinUnitAct sets FinUnitAct field to given value.

### HasFinUnitAct

`func (o *RuleReport) HasFinUnitAct() bool`

HasFinUnitAct returns a boolean if a field has been set.

### GetRanNasRelCauses

`func (o *RuleReport) GetRanNasRelCauses() []RanNasRelCause`

GetRanNasRelCauses returns the RanNasRelCauses field if non-nil, zero value otherwise.

### GetRanNasRelCausesOk

`func (o *RuleReport) GetRanNasRelCausesOk() (*[]RanNasRelCause, bool)`

GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNasRelCauses

`func (o *RuleReport) SetRanNasRelCauses(v []RanNasRelCause)`

SetRanNasRelCauses sets RanNasRelCauses field to given value.

### HasRanNasRelCauses

`func (o *RuleReport) HasRanNasRelCauses() bool`

HasRanNasRelCauses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


