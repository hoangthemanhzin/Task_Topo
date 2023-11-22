# SmPolicyNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | Pointer to **string** |  | [optional] 
**SmPolicyDecision** | Pointer to [**SmPolicyDecision**](SmPolicyDecision.md) |  | [optional] 

## Methods

### NewSmPolicyNotification

`func NewSmPolicyNotification() *SmPolicyNotification`

NewSmPolicyNotification instantiates a new SmPolicyNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyNotificationWithDefaults

`func NewSmPolicyNotificationWithDefaults() *SmPolicyNotification`

NewSmPolicyNotificationWithDefaults instantiates a new SmPolicyNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *SmPolicyNotification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *SmPolicyNotification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *SmPolicyNotification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *SmPolicyNotification) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.

### GetSmPolicyDecision

`func (o *SmPolicyNotification) GetSmPolicyDecision() SmPolicyDecision`

GetSmPolicyDecision returns the SmPolicyDecision field if non-nil, zero value otherwise.

### GetSmPolicyDecisionOk

`func (o *SmPolicyNotification) GetSmPolicyDecisionOk() (*SmPolicyDecision, bool)`

GetSmPolicyDecisionOk returns a tuple with the SmPolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyDecision

`func (o *SmPolicyNotification) SetSmPolicyDecision(v SmPolicyDecision)`

SetSmPolicyDecision sets SmPolicyDecision field to given value.

### HasSmPolicyDecision

`func (o *SmPolicyNotification) HasSmPolicyDecision() bool`

HasSmPolicyDecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


