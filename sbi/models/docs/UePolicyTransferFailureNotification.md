# UePolicyTransferFailureNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | [**N1N2MessageTransferCause**](N1N2MessageTransferCause.md) |  | 
**Ptis** | **[]int32** |  | 

## Methods

### NewUePolicyTransferFailureNotification

`func NewUePolicyTransferFailureNotification(cause N1N2MessageTransferCause, ptis []int32, ) *UePolicyTransferFailureNotification`

NewUePolicyTransferFailureNotification instantiates a new UePolicyTransferFailureNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUePolicyTransferFailureNotificationWithDefaults

`func NewUePolicyTransferFailureNotificationWithDefaults() *UePolicyTransferFailureNotification`

NewUePolicyTransferFailureNotificationWithDefaults instantiates a new UePolicyTransferFailureNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *UePolicyTransferFailureNotification) GetCause() N1N2MessageTransferCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *UePolicyTransferFailureNotification) GetCauseOk() (*N1N2MessageTransferCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *UePolicyTransferFailureNotification) SetCause(v N1N2MessageTransferCause)`

SetCause sets Cause field to given value.


### GetPtis

`func (o *UePolicyTransferFailureNotification) GetPtis() []int32`

GetPtis returns the Ptis field if non-nil, zero value otherwise.

### GetPtisOk

`func (o *UePolicyTransferFailureNotification) GetPtisOk() (*[]int32, bool)`

GetPtisOk returns a tuple with the Ptis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtis

`func (o *UePolicyTransferFailureNotification) SetPtis(v []int32)`

SetPtis sets Ptis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


