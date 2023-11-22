# AmfNon3GppAccessRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**PurgeFlag** | Pointer to **bool** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**ImsVoPs** | [**ImsVoPs**](ImsVoPs.md) |  | 
**DeregCallbackUri** | **string** |  | 
**AmfServiceNameDereg** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**PcscfRestorationCallbackUri** | Pointer to **string** |  | [optional] 
**AmfServiceNamePcscfRest** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**Guami** | [**Guami**](Guami.md) |  | 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**RatType** | [**RatType**](RatType.md) |  | 
**UrrpIndicator** | Pointer to **bool** |  | [optional] 
**AmfEeSubscriptionId** | Pointer to **string** |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** |  | [optional] 
**VgmlcAddress** | Pointer to [**VgmlcAddress**](VgmlcAddress.md) |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**NoEeSubscriptionInd** | Pointer to **bool** |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 

## Methods

### NewAmfNon3GppAccessRegistration

`func NewAmfNon3GppAccessRegistration(amfInstanceId string, imsVoPs ImsVoPs, deregCallbackUri string, guami Guami, ratType RatType, ) *AmfNon3GppAccessRegistration`

NewAmfNon3GppAccessRegistration instantiates a new AmfNon3GppAccessRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfNon3GppAccessRegistrationWithDefaults

`func NewAmfNon3GppAccessRegistrationWithDefaults() *AmfNon3GppAccessRegistration`

NewAmfNon3GppAccessRegistrationWithDefaults instantiates a new AmfNon3GppAccessRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *AmfNon3GppAccessRegistration) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *AmfNon3GppAccessRegistration) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *AmfNon3GppAccessRegistration) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AmfNon3GppAccessRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AmfNon3GppAccessRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPurgeFlag

`func (o *AmfNon3GppAccessRegistration) GetPurgeFlag() bool`

GetPurgeFlag returns the PurgeFlag field if non-nil, zero value otherwise.

### GetPurgeFlagOk

`func (o *AmfNon3GppAccessRegistration) GetPurgeFlagOk() (*bool, bool)`

GetPurgeFlagOk returns a tuple with the PurgeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeFlag

`func (o *AmfNon3GppAccessRegistration) SetPurgeFlag(v bool)`

SetPurgeFlag sets PurgeFlag field to given value.

### HasPurgeFlag

`func (o *AmfNon3GppAccessRegistration) HasPurgeFlag() bool`

HasPurgeFlag returns a boolean if a field has been set.

### GetPei

`func (o *AmfNon3GppAccessRegistration) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AmfNon3GppAccessRegistration) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AmfNon3GppAccessRegistration) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AmfNon3GppAccessRegistration) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetImsVoPs

`func (o *AmfNon3GppAccessRegistration) GetImsVoPs() ImsVoPs`

GetImsVoPs returns the ImsVoPs field if non-nil, zero value otherwise.

### GetImsVoPsOk

`func (o *AmfNon3GppAccessRegistration) GetImsVoPsOk() (*ImsVoPs, bool)`

GetImsVoPsOk returns a tuple with the ImsVoPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVoPs

`func (o *AmfNon3GppAccessRegistration) SetImsVoPs(v ImsVoPs)`

SetImsVoPs sets ImsVoPs field to given value.


### GetDeregCallbackUri

`func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *AmfNon3GppAccessRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.


### GetAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDereg() ServiceName`

GetAmfServiceNameDereg returns the AmfServiceNameDereg field if non-nil, zero value otherwise.

### GetAmfServiceNameDeregOk

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDeregOk() (*ServiceName, bool)`

GetAmfServiceNameDeregOk returns a tuple with the AmfServiceNameDereg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) SetAmfServiceNameDereg(v ServiceName)`

SetAmfServiceNameDereg sets AmfServiceNameDereg field to given value.

### HasAmfServiceNameDereg

`func (o *AmfNon3GppAccessRegistration) HasAmfServiceNameDereg() bool`

HasAmfServiceNameDereg returns a boolean if a field has been set.

### GetPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUri() string`

GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field if non-nil, zero value otherwise.

### GetPcscfRestorationCallbackUriOk

`func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool)`

GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) SetPcscfRestorationCallbackUri(v string)`

SetPcscfRestorationCallbackUri sets PcscfRestorationCallbackUri field to given value.

### HasPcscfRestorationCallbackUri

`func (o *AmfNon3GppAccessRegistration) HasPcscfRestorationCallbackUri() bool`

HasPcscfRestorationCallbackUri returns a boolean if a field has been set.

### GetAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRest() ServiceName`

GetAmfServiceNamePcscfRest returns the AmfServiceNamePcscfRest field if non-nil, zero value otherwise.

### GetAmfServiceNamePcscfRestOk

`func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRestOk() (*ServiceName, bool)`

GetAmfServiceNamePcscfRestOk returns a tuple with the AmfServiceNamePcscfRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) SetAmfServiceNamePcscfRest(v ServiceName)`

SetAmfServiceNamePcscfRest sets AmfServiceNamePcscfRest field to given value.

### HasAmfServiceNamePcscfRest

`func (o *AmfNon3GppAccessRegistration) HasAmfServiceNamePcscfRest() bool`

HasAmfServiceNamePcscfRest returns a boolean if a field has been set.

### GetGuami

`func (o *AmfNon3GppAccessRegistration) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *AmfNon3GppAccessRegistration) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *AmfNon3GppAccessRegistration) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *AmfNon3GppAccessRegistration) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetRatType

`func (o *AmfNon3GppAccessRegistration) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *AmfNon3GppAccessRegistration) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *AmfNon3GppAccessRegistration) SetRatType(v RatType)`

SetRatType sets RatType field to given value.


### GetUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) GetUrrpIndicator() bool`

GetUrrpIndicator returns the UrrpIndicator field if non-nil, zero value otherwise.

### GetUrrpIndicatorOk

`func (o *AmfNon3GppAccessRegistration) GetUrrpIndicatorOk() (*bool, bool)`

GetUrrpIndicatorOk returns a tuple with the UrrpIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) SetUrrpIndicator(v bool)`

SetUrrpIndicator sets UrrpIndicator field to given value.

### HasUrrpIndicator

`func (o *AmfNon3GppAccessRegistration) HasUrrpIndicator() bool`

HasUrrpIndicator returns a boolean if a field has been set.

### GetAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionId() string`

GetAmfEeSubscriptionId returns the AmfEeSubscriptionId field if non-nil, zero value otherwise.

### GetAmfEeSubscriptionIdOk

`func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionIdOk() (*string, bool)`

GetAmfEeSubscriptionIdOk returns a tuple with the AmfEeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) SetAmfEeSubscriptionId(v string)`

SetAmfEeSubscriptionId sets AmfEeSubscriptionId field to given value.

### HasAmfEeSubscriptionId

`func (o *AmfNon3GppAccessRegistration) HasAmfEeSubscriptionId() bool`

HasAmfEeSubscriptionId returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *AmfNon3GppAccessRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *AmfNon3GppAccessRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *AmfNon3GppAccessRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *AmfNon3GppAccessRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) GetVgmlcAddress() VgmlcAddress`

GetVgmlcAddress returns the VgmlcAddress field if non-nil, zero value otherwise.

### GetVgmlcAddressOk

`func (o *AmfNon3GppAccessRegistration) GetVgmlcAddressOk() (*VgmlcAddress, bool)`

GetVgmlcAddressOk returns a tuple with the VgmlcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) SetVgmlcAddress(v VgmlcAddress)`

SetVgmlcAddress sets VgmlcAddress field to given value.

### HasVgmlcAddress

`func (o *AmfNon3GppAccessRegistration) HasVgmlcAddress() bool`

HasVgmlcAddress returns a boolean if a field has been set.

### GetContextInfo

`func (o *AmfNon3GppAccessRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *AmfNon3GppAccessRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *AmfNon3GppAccessRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *AmfNon3GppAccessRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionInd() bool`

GetNoEeSubscriptionInd returns the NoEeSubscriptionInd field if non-nil, zero value otherwise.

### GetNoEeSubscriptionIndOk

`func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionIndOk() (*bool, bool)`

GetNoEeSubscriptionIndOk returns a tuple with the NoEeSubscriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) SetNoEeSubscriptionInd(v bool)`

SetNoEeSubscriptionInd sets NoEeSubscriptionInd field to given value.

### HasNoEeSubscriptionInd

`func (o *AmfNon3GppAccessRegistration) HasNoEeSubscriptionInd() bool`

HasNoEeSubscriptionInd returns a boolean if a field has been set.

### GetSupi

`func (o *AmfNon3GppAccessRegistration) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AmfNon3GppAccessRegistration) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AmfNon3GppAccessRegistration) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AmfNon3GppAccessRegistration) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


