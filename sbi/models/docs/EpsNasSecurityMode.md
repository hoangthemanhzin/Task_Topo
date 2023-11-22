# EpsNasSecurityMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrityAlgorithm** | [**EpsNasIntegrityAlgorithm**](EpsNasIntegrityAlgorithm.md) |  | 
**CipheringAlgorithm** | [**EpsNasCipheringAlgorithm**](EpsNasCipheringAlgorithm.md) |  | 

## Methods

### NewEpsNasSecurityMode

`func NewEpsNasSecurityMode(integrityAlgorithm EpsNasIntegrityAlgorithm, cipheringAlgorithm EpsNasCipheringAlgorithm, ) *EpsNasSecurityMode`

NewEpsNasSecurityMode instantiates a new EpsNasSecurityMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpsNasSecurityModeWithDefaults

`func NewEpsNasSecurityModeWithDefaults() *EpsNasSecurityMode`

NewEpsNasSecurityModeWithDefaults instantiates a new EpsNasSecurityMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrityAlgorithm

`func (o *EpsNasSecurityMode) GetIntegrityAlgorithm() EpsNasIntegrityAlgorithm`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *EpsNasSecurityMode) GetIntegrityAlgorithmOk() (*EpsNasIntegrityAlgorithm, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *EpsNasSecurityMode) SetIntegrityAlgorithm(v EpsNasIntegrityAlgorithm)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.


### GetCipheringAlgorithm

`func (o *EpsNasSecurityMode) GetCipheringAlgorithm() EpsNasCipheringAlgorithm`

GetCipheringAlgorithm returns the CipheringAlgorithm field if non-nil, zero value otherwise.

### GetCipheringAlgorithmOk

`func (o *EpsNasSecurityMode) GetCipheringAlgorithmOk() (*EpsNasCipheringAlgorithm, bool)`

GetCipheringAlgorithmOk returns a tuple with the CipheringAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipheringAlgorithm

`func (o *EpsNasSecurityMode) SetCipheringAlgorithm(v EpsNasCipheringAlgorithm)`

SetCipheringAlgorithm sets CipheringAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


