# SmfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmfInstanceId** | **string** |  | 
**SmfSetId** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**PduSessionId** | **int32** |  | 
**SingleNssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | Pointer to **string** |  | [optional] 
**EmergencyServices** | Pointer to **bool** |  | [optional] 
**PcscfRestorationCallbackUri** | Pointer to **string** |  | [optional] 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**PgwFqdn** | Pointer to **string** |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]
**DeregCallbackUri** | Pointer to **string** |  | [optional] 
**RegistrationReason** | Pointer to [**RegistrationReason**](RegistrationReason.md) |  | [optional] 
**RegistrationTime** | Pointer to **time.Time** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 

## Methods

### NewSmfRegistration

`func NewSmfRegistration(smfInstanceId string, pduSessionId int32, singleNssai Snssai, plmnId PlmnId, ) *SmfRegistration`

NewSmfRegistration instantiates a new SmfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfRegistrationWithDefaults

`func NewSmfRegistrationWithDefaults() *SmfRegistration`

NewSmfRegistrationWithDefaults instantiates a new SmfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmfInstanceId

`func (o *SmfRegistration) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *SmfRegistration) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *SmfRegistration) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.


### GetSmfSetId

`func (o *SmfRegistration) GetSmfSetId() string`

GetSmfSetId returns the SmfSetId field if non-nil, zero value otherwise.

### GetSmfSetIdOk

`func (o *SmfRegistration) GetSmfSetIdOk() (*string, bool)`

GetSmfSetIdOk returns a tuple with the SmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSetId

`func (o *SmfRegistration) SetSmfSetId(v string)`

SetSmfSetId sets SmfSetId field to given value.

### HasSmfSetId

`func (o *SmfRegistration) HasSmfSetId() bool`

HasSmfSetId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmfRegistration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmfRegistration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmfRegistration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmfRegistration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmfRegistration) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmfRegistration) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmfRegistration) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetSingleNssai

`func (o *SmfRegistration) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *SmfRegistration) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *SmfRegistration) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.


### GetDnn

`func (o *SmfRegistration) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmfRegistration) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmfRegistration) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmfRegistration) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetEmergencyServices

`func (o *SmfRegistration) GetEmergencyServices() bool`

GetEmergencyServices returns the EmergencyServices field if non-nil, zero value otherwise.

### GetEmergencyServicesOk

`func (o *SmfRegistration) GetEmergencyServicesOk() (*bool, bool)`

GetEmergencyServicesOk returns a tuple with the EmergencyServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyServices

`func (o *SmfRegistration) SetEmergencyServices(v bool)`

SetEmergencyServices sets EmergencyServices field to given value.

### HasEmergencyServices

`func (o *SmfRegistration) HasEmergencyServices() bool`

HasEmergencyServices returns a boolean if a field has been set.

### GetPcscfRestorationCallbackUri

`func (o *SmfRegistration) GetPcscfRestorationCallbackUri() string`

GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field if non-nil, zero value otherwise.

### GetPcscfRestorationCallbackUriOk

`func (o *SmfRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool)`

GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestorationCallbackUri

`func (o *SmfRegistration) SetPcscfRestorationCallbackUri(v string)`

SetPcscfRestorationCallbackUri sets PcscfRestorationCallbackUri field to given value.

### HasPcscfRestorationCallbackUri

`func (o *SmfRegistration) HasPcscfRestorationCallbackUri() bool`

HasPcscfRestorationCallbackUri returns a boolean if a field has been set.

### GetPlmnId

`func (o *SmfRegistration) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SmfRegistration) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SmfRegistration) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetPgwFqdn

`func (o *SmfRegistration) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *SmfRegistration) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *SmfRegistration) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *SmfRegistration) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetEpdgInd

`func (o *SmfRegistration) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *SmfRegistration) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *SmfRegistration) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *SmfRegistration) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.

### GetDeregCallbackUri

`func (o *SmfRegistration) GetDeregCallbackUri() string`

GetDeregCallbackUri returns the DeregCallbackUri field if non-nil, zero value otherwise.

### GetDeregCallbackUriOk

`func (o *SmfRegistration) GetDeregCallbackUriOk() (*string, bool)`

GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregCallbackUri

`func (o *SmfRegistration) SetDeregCallbackUri(v string)`

SetDeregCallbackUri sets DeregCallbackUri field to given value.

### HasDeregCallbackUri

`func (o *SmfRegistration) HasDeregCallbackUri() bool`

HasDeregCallbackUri returns a boolean if a field has been set.

### GetRegistrationReason

`func (o *SmfRegistration) GetRegistrationReason() RegistrationReason`

GetRegistrationReason returns the RegistrationReason field if non-nil, zero value otherwise.

### GetRegistrationReasonOk

`func (o *SmfRegistration) GetRegistrationReasonOk() (*RegistrationReason, bool)`

GetRegistrationReasonOk returns a tuple with the RegistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationReason

`func (o *SmfRegistration) SetRegistrationReason(v RegistrationReason)`

SetRegistrationReason sets RegistrationReason field to given value.

### HasRegistrationReason

`func (o *SmfRegistration) HasRegistrationReason() bool`

HasRegistrationReason returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *SmfRegistration) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *SmfRegistration) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *SmfRegistration) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *SmfRegistration) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetContextInfo

`func (o *SmfRegistration) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SmfRegistration) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SmfRegistration) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SmfRegistration) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


