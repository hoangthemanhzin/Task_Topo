# Amf3GppAccessRegistrationModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guami** | [**Guami**](Guami.md) |  | 
**PurgeFlag** | Pointer to **bool** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**ImsVoPs** | Pointer to [**ImsVoPs**](ImsVoPs.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**EpsInterworkingInfo** | Pointer to [**EpsInterworkingInfo**](EpsInterworkingInfo.md) |  | [optional] 
**UeSrvccCapability** | Pointer to **bool** |  | [optional] 

## Methods

### NewAmf3GppAccessRegistrationModification

`func NewAmf3GppAccessRegistrationModification(guami Guami, ) *Amf3GppAccessRegistrationModification`

NewAmf3GppAccessRegistrationModification instantiates a new Amf3GppAccessRegistrationModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmf3GppAccessRegistrationModificationWithDefaults

`func NewAmf3GppAccessRegistrationModificationWithDefaults() *Amf3GppAccessRegistrationModification`

NewAmf3GppAccessRegistrationModificationWithDefaults instantiates a new Amf3GppAccessRegistrationModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuami

`func (o *Amf3GppAccessRegistrationModification) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *Amf3GppAccessRegistrationModification) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *Amf3GppAccessRegistrationModification) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetPurgeFlag

`func (o *Amf3GppAccessRegistrationModification) GetPurgeFlag() bool`

GetPurgeFlag returns the PurgeFlag field if non-nil, zero value otherwise.

### GetPurgeFlagOk

`func (o *Amf3GppAccessRegistrationModification) GetPurgeFlagOk() (*bool, bool)`

GetPurgeFlagOk returns a tuple with the PurgeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeFlag

`func (o *Amf3GppAccessRegistrationModification) SetPurgeFlag(v bool)`

SetPurgeFlag sets PurgeFlag field to given value.

### HasPurgeFlag

`func (o *Amf3GppAccessRegistrationModification) HasPurgeFlag() bool`

HasPurgeFlag returns a boolean if a field has been set.

### GetPei

`func (o *Amf3GppAccessRegistrationModification) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *Amf3GppAccessRegistrationModification) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *Amf3GppAccessRegistrationModification) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *Amf3GppAccessRegistrationModification) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetImsVoPs

`func (o *Amf3GppAccessRegistrationModification) GetImsVoPs() ImsVoPs`

GetImsVoPs returns the ImsVoPs field if non-nil, zero value otherwise.

### GetImsVoPsOk

`func (o *Amf3GppAccessRegistrationModification) GetImsVoPsOk() (*ImsVoPs, bool)`

GetImsVoPsOk returns a tuple with the ImsVoPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsVoPs

`func (o *Amf3GppAccessRegistrationModification) SetImsVoPs(v ImsVoPs)`

SetImsVoPs sets ImsVoPs field to given value.

### HasImsVoPs

`func (o *Amf3GppAccessRegistrationModification) HasImsVoPs() bool`

HasImsVoPs returns a boolean if a field has been set.

### GetBackupAmfInfo

`func (o *Amf3GppAccessRegistrationModification) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *Amf3GppAccessRegistrationModification) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *Amf3GppAccessRegistrationModification) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *Amf3GppAccessRegistrationModification) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetEpsInterworkingInfo

`func (o *Amf3GppAccessRegistrationModification) GetEpsInterworkingInfo() EpsInterworkingInfo`

GetEpsInterworkingInfo returns the EpsInterworkingInfo field if non-nil, zero value otherwise.

### GetEpsInterworkingInfoOk

`func (o *Amf3GppAccessRegistrationModification) GetEpsInterworkingInfoOk() (*EpsInterworkingInfo, bool)`

GetEpsInterworkingInfoOk returns a tuple with the EpsInterworkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInfo

`func (o *Amf3GppAccessRegistrationModification) SetEpsInterworkingInfo(v EpsInterworkingInfo)`

SetEpsInterworkingInfo sets EpsInterworkingInfo field to given value.

### HasEpsInterworkingInfo

`func (o *Amf3GppAccessRegistrationModification) HasEpsInterworkingInfo() bool`

HasEpsInterworkingInfo returns a boolean if a field has been set.

### GetUeSrvccCapability

`func (o *Amf3GppAccessRegistrationModification) GetUeSrvccCapability() bool`

GetUeSrvccCapability returns the UeSrvccCapability field if non-nil, zero value otherwise.

### GetUeSrvccCapabilityOk

`func (o *Amf3GppAccessRegistrationModification) GetUeSrvccCapabilityOk() (*bool, bool)`

GetUeSrvccCapabilityOk returns a tuple with the UeSrvccCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSrvccCapability

`func (o *Amf3GppAccessRegistrationModification) SetUeSrvccCapability(v bool)`

SetUeSrvccCapability sets UeSrvccCapability field to given value.

### HasUeSrvccCapability

`func (o *Amf3GppAccessRegistrationModification) HasUeSrvccCapability() bool`

HasUeSrvccCapability returns a boolean if a field has been set.

### SetUeSrvccCapabilityNil

`func (o *Amf3GppAccessRegistrationModification) SetUeSrvccCapabilityNil(b bool)`

 SetUeSrvccCapabilityNil sets the value for UeSrvccCapability to be an explicit nil

### UnsetUeSrvccCapability
`func (o *Amf3GppAccessRegistrationModification) UnsetUeSrvccCapability()`

UnsetUeSrvccCapability ensures that no value is present for UeSrvccCapability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


