# PpData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationCharacteristics** | Pointer to [**CommunicationCharacteristics**](CommunicationCharacteristics.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ExpectedUeBehaviourParameters** | Pointer to [**ExpectedUeBehaviour**](ExpectedUeBehaviour.md) |  | [optional] 
**EcRestriction** | Pointer to [**EcRestriction**](EcRestriction.md) |  | [optional] 
**AcsInfo** | Pointer to [**AcsInfoRm**](AcsInfoRm.md) |  | [optional] 
**StnSr** | Pointer to **string** |  | [optional] 
**LcsPrivacy** | Pointer to [**LcsPrivacy**](LcsPrivacy.md) |  | [optional] 
**SorInfo** | Pointer to [**SorInfo**](SorInfo.md) |  | [optional] 

## Methods

### NewPpData

`func NewPpData() *PpData`

NewPpData instantiates a new PpData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpDataWithDefaults

`func NewPpDataWithDefaults() *PpData`

NewPpDataWithDefaults instantiates a new PpData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationCharacteristics

`func (o *PpData) GetCommunicationCharacteristics() CommunicationCharacteristics`

GetCommunicationCharacteristics returns the CommunicationCharacteristics field if non-nil, zero value otherwise.

### GetCommunicationCharacteristicsOk

`func (o *PpData) GetCommunicationCharacteristicsOk() (*CommunicationCharacteristics, bool)`

GetCommunicationCharacteristicsOk returns a tuple with the CommunicationCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationCharacteristics

`func (o *PpData) SetCommunicationCharacteristics(v CommunicationCharacteristics)`

SetCommunicationCharacteristics sets CommunicationCharacteristics field to given value.

### HasCommunicationCharacteristics

`func (o *PpData) HasCommunicationCharacteristics() bool`

HasCommunicationCharacteristics returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PpData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PpData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PpData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PpData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetExpectedUeBehaviourParameters

`func (o *PpData) GetExpectedUeBehaviourParameters() ExpectedUeBehaviour`

GetExpectedUeBehaviourParameters returns the ExpectedUeBehaviourParameters field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourParametersOk

`func (o *PpData) GetExpectedUeBehaviourParametersOk() (*ExpectedUeBehaviour, bool)`

GetExpectedUeBehaviourParametersOk returns a tuple with the ExpectedUeBehaviourParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourParameters

`func (o *PpData) SetExpectedUeBehaviourParameters(v ExpectedUeBehaviour)`

SetExpectedUeBehaviourParameters sets ExpectedUeBehaviourParameters field to given value.

### HasExpectedUeBehaviourParameters

`func (o *PpData) HasExpectedUeBehaviourParameters() bool`

HasExpectedUeBehaviourParameters returns a boolean if a field has been set.

### GetEcRestriction

`func (o *PpData) GetEcRestriction() EcRestriction`

GetEcRestriction returns the EcRestriction field if non-nil, zero value otherwise.

### GetEcRestrictionOk

`func (o *PpData) GetEcRestrictionOk() (*EcRestriction, bool)`

GetEcRestrictionOk returns a tuple with the EcRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestriction

`func (o *PpData) SetEcRestriction(v EcRestriction)`

SetEcRestriction sets EcRestriction field to given value.

### HasEcRestriction

`func (o *PpData) HasEcRestriction() bool`

HasEcRestriction returns a boolean if a field has been set.

### GetAcsInfo

`func (o *PpData) GetAcsInfo() AcsInfoRm`

GetAcsInfo returns the AcsInfo field if non-nil, zero value otherwise.

### GetAcsInfoOk

`func (o *PpData) GetAcsInfoOk() (*AcsInfoRm, bool)`

GetAcsInfoOk returns a tuple with the AcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfo

`func (o *PpData) SetAcsInfo(v AcsInfoRm)`

SetAcsInfo sets AcsInfo field to given value.

### HasAcsInfo

`func (o *PpData) HasAcsInfo() bool`

HasAcsInfo returns a boolean if a field has been set.

### GetStnSr

`func (o *PpData) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *PpData) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *PpData) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *PpData) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### SetStnSrNil

`func (o *PpData) SetStnSrNil(b bool)`

 SetStnSrNil sets the value for StnSr to be an explicit nil

### UnsetStnSr
`func (o *PpData) UnsetStnSr()`

UnsetStnSr ensures that no value is present for StnSr, not even an explicit nil
### GetLcsPrivacy

`func (o *PpData) GetLcsPrivacy() LcsPrivacy`

GetLcsPrivacy returns the LcsPrivacy field if non-nil, zero value otherwise.

### GetLcsPrivacyOk

`func (o *PpData) GetLcsPrivacyOk() (*LcsPrivacy, bool)`

GetLcsPrivacyOk returns a tuple with the LcsPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacy

`func (o *PpData) SetLcsPrivacy(v LcsPrivacy)`

SetLcsPrivacy sets LcsPrivacy field to given value.

### HasLcsPrivacy

`func (o *PpData) HasLcsPrivacy() bool`

HasLcsPrivacy returns a boolean if a field has been set.

### GetSorInfo

`func (o *PpData) GetSorInfo() SorInfo`

GetSorInfo returns the SorInfo field if non-nil, zero value otherwise.

### GetSorInfoOk

`func (o *PpData) GetSorInfoOk() (*SorInfo, bool)`

GetSorInfoOk returns a tuple with the SorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfo

`func (o *PpData) SetSorInfo(v SorInfo)`

SetSorInfo sets SorInfo field to given value.

### HasSorInfo

`func (o *PpData) HasSorInfo() bool`

HasSorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


