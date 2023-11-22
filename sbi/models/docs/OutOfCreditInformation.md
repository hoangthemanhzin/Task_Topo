# OutOfCreditInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinUnitAct** | [**FinalUnitAction**](FinalUnitAction.md) |  | 
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 

## Methods

### NewOutOfCreditInformation

`func NewOutOfCreditInformation(finUnitAct FinalUnitAction, ) *OutOfCreditInformation`

NewOutOfCreditInformation instantiates a new OutOfCreditInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfCreditInformationWithDefaults

`func NewOutOfCreditInformationWithDefaults() *OutOfCreditInformation`

NewOutOfCreditInformationWithDefaults instantiates a new OutOfCreditInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinUnitAct

`func (o *OutOfCreditInformation) GetFinUnitAct() FinalUnitAction`

GetFinUnitAct returns the FinUnitAct field if non-nil, zero value otherwise.

### GetFinUnitActOk

`func (o *OutOfCreditInformation) GetFinUnitActOk() (*FinalUnitAction, bool)`

GetFinUnitActOk returns a tuple with the FinUnitAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinUnitAct

`func (o *OutOfCreditInformation) SetFinUnitAct(v FinalUnitAction)`

SetFinUnitAct sets FinUnitAct field to given value.


### GetFlows

`func (o *OutOfCreditInformation) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *OutOfCreditInformation) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *OutOfCreditInformation) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *OutOfCreditInformation) HasFlows() bool`

HasFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


