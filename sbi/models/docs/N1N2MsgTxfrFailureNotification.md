# N1N2MsgTxfrFailureNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | [**N1N2MessageTransferCause**](N1N2MessageTransferCause.md) |  | 
**N1n2MsgDataUri** | **string** |  | 

## Methods

### NewN1N2MsgTxfrFailureNotification

`func NewN1N2MsgTxfrFailureNotification(cause N1N2MessageTransferCause, n1n2MsgDataUri string, ) *N1N2MsgTxfrFailureNotification`

NewN1N2MsgTxfrFailureNotification instantiates a new N1N2MsgTxfrFailureNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1N2MsgTxfrFailureNotificationWithDefaults

`func NewN1N2MsgTxfrFailureNotificationWithDefaults() *N1N2MsgTxfrFailureNotification`

NewN1N2MsgTxfrFailureNotificationWithDefaults instantiates a new N1N2MsgTxfrFailureNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *N1N2MsgTxfrFailureNotification) GetCause() N1N2MessageTransferCause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *N1N2MsgTxfrFailureNotification) GetCauseOk() (*N1N2MessageTransferCause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *N1N2MsgTxfrFailureNotification) SetCause(v N1N2MessageTransferCause)`

SetCause sets Cause field to given value.


### GetN1n2MsgDataUri

`func (o *N1N2MsgTxfrFailureNotification) GetN1n2MsgDataUri() string`

GetN1n2MsgDataUri returns the N1n2MsgDataUri field if non-nil, zero value otherwise.

### GetN1n2MsgDataUriOk

`func (o *N1N2MsgTxfrFailureNotification) GetN1n2MsgDataUriOk() (*string, bool)`

GetN1n2MsgDataUriOk returns a tuple with the N1n2MsgDataUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1n2MsgDataUri

`func (o *N1N2MsgTxfrFailureNotification) SetN1n2MsgDataUri(v string)`

SetN1n2MsgDataUri sets N1n2MsgDataUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


