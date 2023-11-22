# EcRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** |  | 
**PlmnEcInfos** | Pointer to [**[]PlmnEcInfo**](PlmnEcInfo.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewEcRestriction

`func NewEcRestriction(afInstanceId string, referenceId int32, ) *EcRestriction`

NewEcRestriction instantiates a new EcRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcRestrictionWithDefaults

`func NewEcRestrictionWithDefaults() *EcRestriction`

NewEcRestrictionWithDefaults instantiates a new EcRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *EcRestriction) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *EcRestriction) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *EcRestriction) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *EcRestriction) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *EcRestriction) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *EcRestriction) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetPlmnEcInfos

`func (o *EcRestriction) GetPlmnEcInfos() []PlmnEcInfo`

GetPlmnEcInfos returns the PlmnEcInfos field if non-nil, zero value otherwise.

### GetPlmnEcInfosOk

`func (o *EcRestriction) GetPlmnEcInfosOk() (*[]PlmnEcInfo, bool)`

GetPlmnEcInfosOk returns a tuple with the PlmnEcInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnEcInfos

`func (o *EcRestriction) SetPlmnEcInfos(v []PlmnEcInfo)`

SetPlmnEcInfos sets PlmnEcInfos field to given value.

### HasPlmnEcInfos

`func (o *EcRestriction) HasPlmnEcInfos() bool`

HasPlmnEcInfos returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *EcRestriction) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *EcRestriction) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *EcRestriction) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *EcRestriction) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


