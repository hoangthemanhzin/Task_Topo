# SmsfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsfInstanceId** | **string** |  | 
**SmsfSetId** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SmsfMAPAddress** | Pointer to **string** |  | [optional] 
**SmsfDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 

## Methods

### NewSmsfRegistration

`func NewSmsfRegistration(smsfInstanceId string, plmnId PlmnId, ) *SmsfRegistration`

NewSmsfRegistration instantiates a new SmsfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsfRegistrationWithDefaults

`func NewSmsfRegistrationWithDefaults() *SmsfRegistration`

NewSmsfRegistrationWithDefaults instantiates a new SmsfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsfInstanceId

`func (o *SmsfRegistration) GetSmsfInstanceId() string`

GetSmsfInstanceId returns the SmsfInstanceId field if non-nil, zero value otherwise.

### GetSmsfInstanceIdOk

`func (o *SmsfRegistration) GetSmsfInstanceIdOk() (*string, bool)`

GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInstanceId

`func (o *SmsfRegistration) SetSmsfInstanceId(v string)`

SetSmsfInstanceId sets SmsfInstanceId field to given value.


### GetSmsfSetId

`func (o *SmsfRegistration) GetSmsfSetId() string`

GetSmsfSetId returns the SmsfSetId field if non-nil, zero value otherwise.

### GetSmsfSetIdOk

`func (o *SmsfRegistration) GetSmsfSetIdOk() (*string, bool)`

GetSmsfSetIdOk returns a tuple with the SmsfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSetId

`func (o *SmsfRegistration) SetSmsfSetId(v string)`

SetSmsfSetId sets SmsfSetId field to given value.

### HasSmsfSetId

`func (o *SmsfRegistration) HasSmsfSetId() bool`

HasSmsfSetId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmsfRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmsfRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmsfRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmsfRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPlmnId

`func (o *SmsfRegistration) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SmsfRegistration) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SmsfRegistration) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSmsfMAPAddress

`func (o *SmsfRegistration) GetSmsfMAPAddress() string`

GetSmsfMAPAddress returns the SmsfMAPAddress field if non-nil, zero value otherwise.

### GetSmsfMAPAddressOk

`func (o *SmsfRegistration) GetSmsfMAPAddressOk() (*string, bool)`

GetSmsfMAPAddressOk returns a tuple with the SmsfMAPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfMAPAddress

`func (o *SmsfRegistration) SetSmsfMAPAddress(v string)`

SetSmsfMAPAddress sets SmsfMAPAddress field to given value.

### HasSmsfMAPAddress

`func (o *SmsfRegistration) HasSmsfMAPAddress() bool`

HasSmsfMAPAddress returns a boolean if a field has been set.

### GetSmsfDiameterAddress

`func (o *SmsfRegistration) GetSmsfDiameterAddress() NetworkNodeDiameterAddress`

GetSmsfDiameterAddress returns the SmsfDiameterAddress field if non-nil, zero value otherwise.

### GetSmsfDiameterAddressOk

`func (o *SmsfRegistration) GetSmsfDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetSmsfDiameterAddressOk returns a tuple with the SmsfDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfDiameterAddress

`func (o *SmsfRegistration) SetSmsfDiameterAddress(v NetworkNodeDiameterAddress)`

SetSmsfDiameterAddress sets SmsfDiameterAddress field to given value.

### HasSmsfDiameterAddress

`func (o *SmsfRegistration) HasSmsfDiameterAddress() bool`

HasSmsfDiameterAddress returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *SmsfRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *SmsfRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *SmsfRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *SmsfRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetContextInfo

`func (o *SmsfRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SmsfRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SmsfRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SmsfRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


