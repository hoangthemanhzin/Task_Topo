# UeN1N2InfoSubscriptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N2InformationClass** | Pointer to [**N2InformationClass**](N2InformationClass.md) |  | [optional] 
**N2NotifyCallbackUri** | Pointer to **string** |  | [optional] 
**N1MessageClass** | Pointer to [**N1MessageClass**](N1MessageClass.md) |  | [optional] 
**N1NotifyCallbackUri** | Pointer to **string** |  | [optional] 
**NfId** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 

## Methods

### NewUeN1N2InfoSubscriptionCreateData

`func NewUeN1N2InfoSubscriptionCreateData() *UeN1N2InfoSubscriptionCreateData`

NewUeN1N2InfoSubscriptionCreateData instantiates a new UeN1N2InfoSubscriptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeN1N2InfoSubscriptionCreateDataWithDefaults

`func NewUeN1N2InfoSubscriptionCreateDataWithDefaults() *UeN1N2InfoSubscriptionCreateData`

NewUeN1N2InfoSubscriptionCreateDataWithDefaults instantiates a new UeN1N2InfoSubscriptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN2InformationClass

`func (o *UeN1N2InfoSubscriptionCreateData) GetN2InformationClass() N2InformationClass`

GetN2InformationClass returns the N2InformationClass field if non-nil, zero value otherwise.

### GetN2InformationClassOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetN2InformationClassOk() (*N2InformationClass, bool)`

GetN2InformationClassOk returns a tuple with the N2InformationClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InformationClass

`func (o *UeN1N2InfoSubscriptionCreateData) SetN2InformationClass(v N2InformationClass)`

SetN2InformationClass sets N2InformationClass field to given value.

### HasN2InformationClass

`func (o *UeN1N2InfoSubscriptionCreateData) HasN2InformationClass() bool`

HasN2InformationClass returns a boolean if a field has been set.

### GetN2NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) GetN2NotifyCallbackUri() string`

GetN2NotifyCallbackUri returns the N2NotifyCallbackUri field if non-nil, zero value otherwise.

### GetN2NotifyCallbackUriOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetN2NotifyCallbackUriOk() (*string, bool)`

GetN2NotifyCallbackUriOk returns a tuple with the N2NotifyCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) SetN2NotifyCallbackUri(v string)`

SetN2NotifyCallbackUri sets N2NotifyCallbackUri field to given value.

### HasN2NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) HasN2NotifyCallbackUri() bool`

HasN2NotifyCallbackUri returns a boolean if a field has been set.

### GetN1MessageClass

`func (o *UeN1N2InfoSubscriptionCreateData) GetN1MessageClass() N1MessageClass`

GetN1MessageClass returns the N1MessageClass field if non-nil, zero value otherwise.

### GetN1MessageClassOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetN1MessageClassOk() (*N1MessageClass, bool)`

GetN1MessageClassOk returns a tuple with the N1MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageClass

`func (o *UeN1N2InfoSubscriptionCreateData) SetN1MessageClass(v N1MessageClass)`

SetN1MessageClass sets N1MessageClass field to given value.

### HasN1MessageClass

`func (o *UeN1N2InfoSubscriptionCreateData) HasN1MessageClass() bool`

HasN1MessageClass returns a boolean if a field has been set.

### GetN1NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) GetN1NotifyCallbackUri() string`

GetN1NotifyCallbackUri returns the N1NotifyCallbackUri field if non-nil, zero value otherwise.

### GetN1NotifyCallbackUriOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetN1NotifyCallbackUriOk() (*string, bool)`

GetN1NotifyCallbackUriOk returns a tuple with the N1NotifyCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) SetN1NotifyCallbackUri(v string)`

SetN1NotifyCallbackUri sets N1NotifyCallbackUri field to given value.

### HasN1NotifyCallbackUri

`func (o *UeN1N2InfoSubscriptionCreateData) HasN1NotifyCallbackUri() bool`

HasN1NotifyCallbackUri returns a boolean if a field has been set.

### GetNfId

`func (o *UeN1N2InfoSubscriptionCreateData) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *UeN1N2InfoSubscriptionCreateData) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *UeN1N2InfoSubscriptionCreateData) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeN1N2InfoSubscriptionCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeN1N2InfoSubscriptionCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeN1N2InfoSubscriptionCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *UeN1N2InfoSubscriptionCreateData) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *UeN1N2InfoSubscriptionCreateData) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *UeN1N2InfoSubscriptionCreateData) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *UeN1N2InfoSubscriptionCreateData) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


