# SdmSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** |  | 
**ImplicitUnsubscribe** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**CallbackReference** | **string** |  | 
**AmfServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ImmediateReport** | Pointer to **bool** |  | [optional] [default to false]
**Report** | Pointer to [**SubscriptionDataSets**](SubscriptionDataSets.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**UniqueSubscription** | Pointer to **bool** |  | [optional] 

## Methods

### NewSdmSubscription

`func NewSdmSubscription(nfInstanceId string, callbackReference string, monitoredResourceUris []string, ) *SdmSubscription`

NewSdmSubscription instantiates a new SdmSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdmSubscriptionWithDefaults

`func NewSdmSubscriptionWithDefaults() *SdmSubscription`

NewSdmSubscriptionWithDefaults instantiates a new SdmSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SdmSubscription) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SdmSubscription) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SdmSubscription) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetImplicitUnsubscribe

`func (o *SdmSubscription) GetImplicitUnsubscribe() bool`

GetImplicitUnsubscribe returns the ImplicitUnsubscribe field if non-nil, zero value otherwise.

### GetImplicitUnsubscribeOk

`func (o *SdmSubscription) GetImplicitUnsubscribeOk() (*bool, bool)`

GetImplicitUnsubscribeOk returns a tuple with the ImplicitUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitUnsubscribe

`func (o *SdmSubscription) SetImplicitUnsubscribe(v bool)`

SetImplicitUnsubscribe sets ImplicitUnsubscribe field to given value.

### HasImplicitUnsubscribe

`func (o *SdmSubscription) HasImplicitUnsubscribe() bool`

HasImplicitUnsubscribe returns a boolean if a field has been set.

### GetExpires

`func (o *SdmSubscription) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SdmSubscription) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SdmSubscription) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SdmSubscription) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetCallbackReference

`func (o *SdmSubscription) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *SdmSubscription) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *SdmSubscription) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetAmfServiceName

`func (o *SdmSubscription) GetAmfServiceName() ServiceName`

GetAmfServiceName returns the AmfServiceName field if non-nil, zero value otherwise.

### GetAmfServiceNameOk

`func (o *SdmSubscription) GetAmfServiceNameOk() (*ServiceName, bool)`

GetAmfServiceNameOk returns a tuple with the AmfServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceName

`func (o *SdmSubscription) SetAmfServiceName(v ServiceName)`

SetAmfServiceName sets AmfServiceName field to given value.

### HasAmfServiceName

`func (o *SdmSubscription) HasAmfServiceName() bool`

HasAmfServiceName returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *SdmSubscription) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SdmSubscription) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SdmSubscription) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetSingleNssai

`func (o *SdmSubscription) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *SdmSubscription) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *SdmSubscription) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *SdmSubscription) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetDnn

`func (o *SdmSubscription) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SdmSubscription) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SdmSubscription) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SdmSubscription) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SdmSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SdmSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SdmSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SdmSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPlmnId

`func (o *SdmSubscription) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SdmSubscription) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SdmSubscription) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *SdmSubscription) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetImmediateReport

`func (o *SdmSubscription) GetImmediateReport() bool`

GetImmediateReport returns the ImmediateReport field if non-nil, zero value otherwise.

### GetImmediateReportOk

`func (o *SdmSubscription) GetImmediateReportOk() (*bool, bool)`

GetImmediateReportOk returns a tuple with the ImmediateReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReport

`func (o *SdmSubscription) SetImmediateReport(v bool)`

SetImmediateReport sets ImmediateReport field to given value.

### HasImmediateReport

`func (o *SdmSubscription) HasImmediateReport() bool`

HasImmediateReport returns a boolean if a field has been set.

### GetReport

`func (o *SdmSubscription) GetReport() SubscriptionDataSets`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *SdmSubscription) GetReportOk() (*SubscriptionDataSets, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *SdmSubscription) SetReport(v SubscriptionDataSets)`

SetReport sets Report field to given value.

### HasReport

`func (o *SdmSubscription) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SdmSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SdmSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SdmSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SdmSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetContextInfo

`func (o *SdmSubscription) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SdmSubscription) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SdmSubscription) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SdmSubscription) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *SdmSubscription) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *SdmSubscription) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *SdmSubscription) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *SdmSubscription) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


