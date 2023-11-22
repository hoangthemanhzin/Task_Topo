# ServiceParameterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | Identifies an application. | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**InterGroupId** | Pointer to **string** |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**UeIpv4** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**UeIpv6** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**UeMac** | Pointer to **string** |  | [optional] 
**AnyUeInd** | Pointer to **bool** |  | [optional] 
**ParamOverPc5** | Pointer to **string** |  | [optional] 
**ParamOverUu** | Pointer to **string** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 
**ResUri** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceParameterData

`func NewServiceParameterData() *ServiceParameterData`

NewServiceParameterData instantiates a new ServiceParameterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceParameterDataWithDefaults

`func NewServiceParameterDataWithDefaults() *ServiceParameterData`

NewServiceParameterDataWithDefaults instantiates a new ServiceParameterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ServiceParameterData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceParameterData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceParameterData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceParameterData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnn

`func (o *ServiceParameterData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ServiceParameterData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ServiceParameterData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ServiceParameterData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *ServiceParameterData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ServiceParameterData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ServiceParameterData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ServiceParameterData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetInterGroupId

`func (o *ServiceParameterData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *ServiceParameterData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *ServiceParameterData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *ServiceParameterData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *ServiceParameterData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ServiceParameterData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ServiceParameterData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ServiceParameterData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeIpv4

`func (o *ServiceParameterData) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *ServiceParameterData) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *ServiceParameterData) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *ServiceParameterData) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *ServiceParameterData) GetUeIpv6() string`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *ServiceParameterData) GetUeIpv6Ok() (*string, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *ServiceParameterData) SetUeIpv6(v string)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *ServiceParameterData) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetUeMac

`func (o *ServiceParameterData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *ServiceParameterData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *ServiceParameterData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *ServiceParameterData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *ServiceParameterData) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *ServiceParameterData) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *ServiceParameterData) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *ServiceParameterData) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetParamOverPc5

`func (o *ServiceParameterData) GetParamOverPc5() string`

GetParamOverPc5 returns the ParamOverPc5 field if non-nil, zero value otherwise.

### GetParamOverPc5Ok

`func (o *ServiceParameterData) GetParamOverPc5Ok() (*string, bool)`

GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverPc5

`func (o *ServiceParameterData) SetParamOverPc5(v string)`

SetParamOverPc5 sets ParamOverPc5 field to given value.

### HasParamOverPc5

`func (o *ServiceParameterData) HasParamOverPc5() bool`

HasParamOverPc5 returns a boolean if a field has been set.

### GetParamOverUu

`func (o *ServiceParameterData) GetParamOverUu() string`

GetParamOverUu returns the ParamOverUu field if non-nil, zero value otherwise.

### GetParamOverUuOk

`func (o *ServiceParameterData) GetParamOverUuOk() (*string, bool)`

GetParamOverUuOk returns a tuple with the ParamOverUu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverUu

`func (o *ServiceParameterData) SetParamOverUu(v string)`

SetParamOverUu sets ParamOverUu field to given value.

### HasParamOverUu

`func (o *ServiceParameterData) HasParamOverUu() bool`

HasParamOverUu returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ServiceParameterData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ServiceParameterData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ServiceParameterData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ServiceParameterData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResUri

`func (o *ServiceParameterData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *ServiceParameterData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *ServiceParameterData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *ServiceParameterData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


