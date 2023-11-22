# NotifiedPosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEvent** | [**LocationEvent**](LocationEvent.md) |  | 
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**LocationEstimate** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** |  | [optional] 
**VelocityEstimate** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**PositioningDataList** | Pointer to [**[]PositioningMethodAndUsage**](PositioningMethodAndUsage.md) |  | [optional] 
**GnssPositioningDataList** | Pointer to [**[]GnssPositioningMethodAndUsage**](GnssPositioningMethodAndUsage.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**ServingNode** | Pointer to **string** |  | [optional] 
**TargetMmeName** | Pointer to **string** |  | [optional] 
**TargetMmeRealm** | Pointer to **string** |  | [optional] 
**UtranSrvccInd** | Pointer to **bool** |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**BarometricPressure** | Pointer to **int32** |  | [optional] 
**Altitude** | Pointer to **float64** |  | [optional] 
**HgmlcCallBackURI** | Pointer to **string** |  | [optional] 
**LdrReference** | Pointer to **string** |  | [optional] 
**ServingLMFIdentification** | Pointer to **string** |  | [optional] 
**TerminationCause** | Pointer to [**TerminationCause**](TerminationCause.md) |  | [optional] 

## Methods

### NewNotifiedPosInfo

`func NewNotifiedPosInfo(locationEvent LocationEvent, ) *NotifiedPosInfo`

NewNotifiedPosInfo instantiates a new NotifiedPosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifiedPosInfoWithDefaults

`func NewNotifiedPosInfoWithDefaults() *NotifiedPosInfo`

NewNotifiedPosInfoWithDefaults instantiates a new NotifiedPosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEvent

`func (o *NotifiedPosInfo) GetLocationEvent() LocationEvent`

GetLocationEvent returns the LocationEvent field if non-nil, zero value otherwise.

### GetLocationEventOk

`func (o *NotifiedPosInfo) GetLocationEventOk() (*LocationEvent, bool)`

GetLocationEventOk returns a tuple with the LocationEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvent

`func (o *NotifiedPosInfo) SetLocationEvent(v LocationEvent)`

SetLocationEvent sets LocationEvent field to given value.


### GetSupi

`func (o *NotifiedPosInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *NotifiedPosInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *NotifiedPosInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *NotifiedPosInfo) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *NotifiedPosInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *NotifiedPosInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *NotifiedPosInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *NotifiedPosInfo) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPei

`func (o *NotifiedPosInfo) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *NotifiedPosInfo) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *NotifiedPosInfo) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *NotifiedPosInfo) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetLocationEstimate

`func (o *NotifiedPosInfo) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *NotifiedPosInfo) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *NotifiedPosInfo) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *NotifiedPosInfo) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *NotifiedPosInfo) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *NotifiedPosInfo) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *NotifiedPosInfo) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *NotifiedPosInfo) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetVelocityEstimate

`func (o *NotifiedPosInfo) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *NotifiedPosInfo) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *NotifiedPosInfo) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *NotifiedPosInfo) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

### GetPositioningDataList

`func (o *NotifiedPosInfo) GetPositioningDataList() []PositioningMethodAndUsage`

GetPositioningDataList returns the PositioningDataList field if non-nil, zero value otherwise.

### GetPositioningDataListOk

`func (o *NotifiedPosInfo) GetPositioningDataListOk() (*[]PositioningMethodAndUsage, bool)`

GetPositioningDataListOk returns a tuple with the PositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningDataList

`func (o *NotifiedPosInfo) SetPositioningDataList(v []PositioningMethodAndUsage)`

SetPositioningDataList sets PositioningDataList field to given value.

### HasPositioningDataList

`func (o *NotifiedPosInfo) HasPositioningDataList() bool`

HasPositioningDataList returns a boolean if a field has been set.

### GetGnssPositioningDataList

`func (o *NotifiedPosInfo) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage`

GetGnssPositioningDataList returns the GnssPositioningDataList field if non-nil, zero value otherwise.

### GetGnssPositioningDataListOk

`func (o *NotifiedPosInfo) GetGnssPositioningDataListOk() (*[]GnssPositioningMethodAndUsage, bool)`

GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssPositioningDataList

`func (o *NotifiedPosInfo) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage)`

SetGnssPositioningDataList sets GnssPositioningDataList field to given value.

### HasGnssPositioningDataList

`func (o *NotifiedPosInfo) HasGnssPositioningDataList() bool`

HasGnssPositioningDataList returns a boolean if a field has been set.

### GetEcgi

`func (o *NotifiedPosInfo) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *NotifiedPosInfo) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *NotifiedPosInfo) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *NotifiedPosInfo) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *NotifiedPosInfo) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *NotifiedPosInfo) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *NotifiedPosInfo) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *NotifiedPosInfo) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetServingNode

`func (o *NotifiedPosInfo) GetServingNode() string`

GetServingNode returns the ServingNode field if non-nil, zero value otherwise.

### GetServingNodeOk

`func (o *NotifiedPosInfo) GetServingNodeOk() (*string, bool)`

GetServingNodeOk returns a tuple with the ServingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNode

`func (o *NotifiedPosInfo) SetServingNode(v string)`

SetServingNode sets ServingNode field to given value.

### HasServingNode

`func (o *NotifiedPosInfo) HasServingNode() bool`

HasServingNode returns a boolean if a field has been set.

### GetTargetMmeName

`func (o *NotifiedPosInfo) GetTargetMmeName() string`

GetTargetMmeName returns the TargetMmeName field if non-nil, zero value otherwise.

### GetTargetMmeNameOk

`func (o *NotifiedPosInfo) GetTargetMmeNameOk() (*string, bool)`

GetTargetMmeNameOk returns a tuple with the TargetMmeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMmeName

`func (o *NotifiedPosInfo) SetTargetMmeName(v string)`

SetTargetMmeName sets TargetMmeName field to given value.

### HasTargetMmeName

`func (o *NotifiedPosInfo) HasTargetMmeName() bool`

HasTargetMmeName returns a boolean if a field has been set.

### GetTargetMmeRealm

`func (o *NotifiedPosInfo) GetTargetMmeRealm() string`

GetTargetMmeRealm returns the TargetMmeRealm field if non-nil, zero value otherwise.

### GetTargetMmeRealmOk

`func (o *NotifiedPosInfo) GetTargetMmeRealmOk() (*string, bool)`

GetTargetMmeRealmOk returns a tuple with the TargetMmeRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMmeRealm

`func (o *NotifiedPosInfo) SetTargetMmeRealm(v string)`

SetTargetMmeRealm sets TargetMmeRealm field to given value.

### HasTargetMmeRealm

`func (o *NotifiedPosInfo) HasTargetMmeRealm() bool`

HasTargetMmeRealm returns a boolean if a field has been set.

### GetUtranSrvccInd

`func (o *NotifiedPosInfo) GetUtranSrvccInd() bool`

GetUtranSrvccInd returns the UtranSrvccInd field if non-nil, zero value otherwise.

### GetUtranSrvccIndOk

`func (o *NotifiedPosInfo) GetUtranSrvccIndOk() (*bool, bool)`

GetUtranSrvccIndOk returns a tuple with the UtranSrvccInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtranSrvccInd

`func (o *NotifiedPosInfo) SetUtranSrvccInd(v bool)`

SetUtranSrvccInd sets UtranSrvccInd field to given value.

### HasUtranSrvccInd

`func (o *NotifiedPosInfo) HasUtranSrvccInd() bool`

HasUtranSrvccInd returns a boolean if a field has been set.

### GetCivicAddress

`func (o *NotifiedPosInfo) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *NotifiedPosInfo) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *NotifiedPosInfo) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *NotifiedPosInfo) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetBarometricPressure

`func (o *NotifiedPosInfo) GetBarometricPressure() int32`

GetBarometricPressure returns the BarometricPressure field if non-nil, zero value otherwise.

### GetBarometricPressureOk

`func (o *NotifiedPosInfo) GetBarometricPressureOk() (*int32, bool)`

GetBarometricPressureOk returns a tuple with the BarometricPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarometricPressure

`func (o *NotifiedPosInfo) SetBarometricPressure(v int32)`

SetBarometricPressure sets BarometricPressure field to given value.

### HasBarometricPressure

`func (o *NotifiedPosInfo) HasBarometricPressure() bool`

HasBarometricPressure returns a boolean if a field has been set.

### GetAltitude

`func (o *NotifiedPosInfo) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *NotifiedPosInfo) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *NotifiedPosInfo) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *NotifiedPosInfo) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetHgmlcCallBackURI

`func (o *NotifiedPosInfo) GetHgmlcCallBackURI() string`

GetHgmlcCallBackURI returns the HgmlcCallBackURI field if non-nil, zero value otherwise.

### GetHgmlcCallBackURIOk

`func (o *NotifiedPosInfo) GetHgmlcCallBackURIOk() (*string, bool)`

GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackURI

`func (o *NotifiedPosInfo) SetHgmlcCallBackURI(v string)`

SetHgmlcCallBackURI sets HgmlcCallBackURI field to given value.

### HasHgmlcCallBackURI

`func (o *NotifiedPosInfo) HasHgmlcCallBackURI() bool`

HasHgmlcCallBackURI returns a boolean if a field has been set.

### GetLdrReference

`func (o *NotifiedPosInfo) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *NotifiedPosInfo) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *NotifiedPosInfo) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.

### HasLdrReference

`func (o *NotifiedPosInfo) HasLdrReference() bool`

HasLdrReference returns a boolean if a field has been set.

### GetServingLMFIdentification

`func (o *NotifiedPosInfo) GetServingLMFIdentification() string`

GetServingLMFIdentification returns the ServingLMFIdentification field if non-nil, zero value otherwise.

### GetServingLMFIdentificationOk

`func (o *NotifiedPosInfo) GetServingLMFIdentificationOk() (*string, bool)`

GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingLMFIdentification

`func (o *NotifiedPosInfo) SetServingLMFIdentification(v string)`

SetServingLMFIdentification sets ServingLMFIdentification field to given value.

### HasServingLMFIdentification

`func (o *NotifiedPosInfo) HasServingLMFIdentification() bool`

HasServingLMFIdentification returns a boolean if a field has been set.

### GetTerminationCause

`func (o *NotifiedPosInfo) GetTerminationCause() TerminationCause`

GetTerminationCause returns the TerminationCause field if non-nil, zero value otherwise.

### GetTerminationCauseOk

`func (o *NotifiedPosInfo) GetTerminationCauseOk() (*TerminationCause, bool)`

GetTerminationCauseOk returns a tuple with the TerminationCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationCause

`func (o *NotifiedPosInfo) SetTerminationCause(v TerminationCause)`

SetTerminationCause sets TerminationCause field to given value.

### HasTerminationCause

`func (o *NotifiedPosInfo) HasTerminationCause() bool`

HasTerminationCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


