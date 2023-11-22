# NasSecurityMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrityAlgorithm** | [**IntegrityAlgorithm**](IntegrityAlgorithm.md) |  | 
**CipheringAlgorithm** | [**CipheringAlgorithm**](CipheringAlgorithm.md) |  | 

## Methods

### NewNasSecurityMode

`func NewNasSecurityMode(integrityAlgorithm IntegrityAlgorithm, cipheringAlgorithm CipheringAlgorithm, ) *NasSecurityMode`

NewNasSecurityMode instantiates a new NasSecurityMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasSecurityModeWithDefaults

`func NewNasSecurityModeWithDefaults() *NasSecurityMode`

NewNasSecurityModeWithDefaults instantiates a new NasSecurityMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrityAlgorithm

`func (o *NasSecurityMode) GetIntegrityAlgorithm() IntegrityAlgorithm`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *NasSecurityMode) GetIntegrityAlgorithmOk() (*IntegrityAlgorithm, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *NasSecurityMode) SetIntegrityAlgorithm(v IntegrityAlgorithm)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.


### GetCipheringAlgorithm

`func (o *NasSecurityMode) GetCipheringAlgorithm() CipheringAlgorithm`

GetCipheringAlgorithm returns the CipheringAlgorithm field if non-nil, zero value otherwise.

### GetCipheringAlgorithmOk

`func (o *NasSecurityMode) GetCipheringAlgorithmOk() (*CipheringAlgorithm, bool)`

GetCipheringAlgorithmOk returns a tuple with the CipheringAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipheringAlgorithm

`func (o *NasSecurityMode) SetCipheringAlgorithm(v CipheringAlgorithm)`

SetCipheringAlgorithm sets CipheringAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


