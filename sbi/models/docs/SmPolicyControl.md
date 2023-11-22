# SmPolicyControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**SmPolicyContextData**](SmPolicyContextData.md) |  | 
**Policy** | [**SmPolicyDecision**](SmPolicyDecision.md) |  | 

## Methods

### NewSmPolicyControl

`func NewSmPolicyControl(context SmPolicyContextData, policy SmPolicyDecision, ) *SmPolicyControl`

NewSmPolicyControl instantiates a new SmPolicyControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyControlWithDefaults

`func NewSmPolicyControlWithDefaults() *SmPolicyControl`

NewSmPolicyControlWithDefaults instantiates a new SmPolicyControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SmPolicyControl) GetContext() SmPolicyContextData`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SmPolicyControl) GetContextOk() (*SmPolicyContextData, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SmPolicyControl) SetContext(v SmPolicyContextData)`

SetContext sets Context field to given value.


### GetPolicy

`func (o *SmPolicyControl) GetPolicy() SmPolicyDecision`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SmPolicyControl) GetPolicyOk() (*SmPolicyDecision, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SmPolicyControl) SetPolicy(v SmPolicyDecision)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


