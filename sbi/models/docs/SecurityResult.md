# SecurityResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrityProtectionResult** | Pointer to [**ProtectionResult**](ProtectionResult.md) |  | [optional] 
**ConfidentialityProtectionResult** | Pointer to [**ProtectionResult**](ProtectionResult.md) |  | [optional] 

## Methods

### NewSecurityResult

`func NewSecurityResult() *SecurityResult`

NewSecurityResult instantiates a new SecurityResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityResultWithDefaults

`func NewSecurityResultWithDefaults() *SecurityResult`

NewSecurityResultWithDefaults instantiates a new SecurityResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrityProtectionResult

`func (o *SecurityResult) GetIntegrityProtectionResult() ProtectionResult`

GetIntegrityProtectionResult returns the IntegrityProtectionResult field if non-nil, zero value otherwise.

### GetIntegrityProtectionResultOk

`func (o *SecurityResult) GetIntegrityProtectionResultOk() (*ProtectionResult, bool)`

GetIntegrityProtectionResultOk returns a tuple with the IntegrityProtectionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityProtectionResult

`func (o *SecurityResult) SetIntegrityProtectionResult(v ProtectionResult)`

SetIntegrityProtectionResult sets IntegrityProtectionResult field to given value.

### HasIntegrityProtectionResult

`func (o *SecurityResult) HasIntegrityProtectionResult() bool`

HasIntegrityProtectionResult returns a boolean if a field has been set.

### GetConfidentialityProtectionResult

`func (o *SecurityResult) GetConfidentialityProtectionResult() ProtectionResult`

GetConfidentialityProtectionResult returns the ConfidentialityProtectionResult field if non-nil, zero value otherwise.

### GetConfidentialityProtectionResultOk

`func (o *SecurityResult) GetConfidentialityProtectionResultOk() (*ProtectionResult, bool)`

GetConfidentialityProtectionResultOk returns a tuple with the ConfidentialityProtectionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityProtectionResult

`func (o *SecurityResult) SetConfidentialityProtectionResult(v ProtectionResult)`

SetConfidentialityProtectionResult sets ConfidentialityProtectionResult field to given value.

### HasConfidentialityProtectionResult

`func (o *SecurityResult) HasConfidentialityProtectionResult() bool`

HasConfidentialityProtectionResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


