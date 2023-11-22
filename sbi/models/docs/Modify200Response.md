# Modify200Response

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
**Report** | [**[]ReportItem**](ReportItem.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**UniqueSubscription** | Pointer to **bool** |  | [optional] 

## Methods

### NewModify200Response

`func NewModify200Response(nfInstanceId string, callbackReference string, monitoredResourceUris []string, report []ReportItem, ) *Modify200Response`

NewModify200Response instantiates a new Modify200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModify200ResponseWithDefaults

`func NewModify200ResponseWithDefaults() *Modify200Response`

NewModify200ResponseWithDefaults instantiates a new Modify200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *Modify200Response) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *Modify200Response) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *Modify200Response) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetImplicitUnsubscribe

`func (o *Modify200Response) GetImplicitUnsubscribe() bool`

GetImplicitUnsubscribe returns the ImplicitUnsubscribe field if non-nil, zero value otherwise.

### GetImplicitUnsubscribeOk

`func (o *Modify200Response) GetImplicitUnsubscribeOk() (*bool, bool)`

GetImplicitUnsubscribeOk returns a tuple with the ImplicitUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitUnsubscribe

`func (o *Modify200Response) SetImplicitUnsubscribe(v bool)`

SetImplicitUnsubscribe sets ImplicitUnsubscribe field to given value.

### HasImplicitUnsubscribe

`func (o *Modify200Response) HasImplicitUnsubscribe() bool`

HasImplicitUnsubscribe returns a boolean if a field has been set.

### GetExpires

`func (o *Modify200Response) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Modify200Response) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Modify200Response) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Modify200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetCallbackReference

`func (o *Modify200Response) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *Modify200Response) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *Modify200Response) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetAmfServiceName

`func (o *Modify200Response) GetAmfServiceName() ServiceName`

GetAmfServiceName returns the AmfServiceName field if non-nil, zero value otherwise.

### GetAmfServiceNameOk

`func (o *Modify200Response) GetAmfServiceNameOk() (*ServiceName, bool)`

GetAmfServiceNameOk returns a tuple with the AmfServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceName

`func (o *Modify200Response) SetAmfServiceName(v ServiceName)`

SetAmfServiceName sets AmfServiceName field to given value.

### HasAmfServiceName

`func (o *Modify200Response) HasAmfServiceName() bool`

HasAmfServiceName returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *Modify200Response) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *Modify200Response) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *Modify200Response) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetSingleNssai

`func (o *Modify200Response) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *Modify200Response) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *Modify200Response) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *Modify200Response) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetDnn

`func (o *Modify200Response) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Modify200Response) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Modify200Response) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *Modify200Response) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Modify200Response) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Modify200Response) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Modify200Response) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Modify200Response) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPlmnId

`func (o *Modify200Response) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Modify200Response) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Modify200Response) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *Modify200Response) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetImmediateReport

`func (o *Modify200Response) GetImmediateReport() bool`

GetImmediateReport returns the ImmediateReport field if non-nil, zero value otherwise.

### GetImmediateReportOk

`func (o *Modify200Response) GetImmediateReportOk() (*bool, bool)`

GetImmediateReportOk returns a tuple with the ImmediateReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReport

`func (o *Modify200Response) SetImmediateReport(v bool)`

SetImmediateReport sets ImmediateReport field to given value.

### HasImmediateReport

`func (o *Modify200Response) HasImmediateReport() bool`

HasImmediateReport returns a boolean if a field has been set.

### GetReport

`func (o *Modify200Response) GetReport() []ReportItem`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *Modify200Response) GetReportOk() (*[]ReportItem, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *Modify200Response) SetReport(v []ReportItem)`

SetReport sets Report field to given value.


### GetSupportedFeatures

`func (o *Modify200Response) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Modify200Response) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Modify200Response) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Modify200Response) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetContextInfo

`func (o *Modify200Response) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *Modify200Response) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *Modify200Response) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *Modify200Response) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *Modify200Response) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *Modify200Response) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *Modify200Response) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *Modify200Response) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


