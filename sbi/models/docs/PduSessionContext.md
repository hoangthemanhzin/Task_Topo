# PduSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** |  | 
**SmContextRef** | **string** |  | 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | **string** |  | 
**SelectedDnn** | Pointer to **string** |  | [optional] 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AllocatedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**HsmfId** | Pointer to **string** |  | [optional] 
**HsmfSetId** | Pointer to **string** |  | [optional] 
**HsmfServiceSetId** | Pointer to **string** |  | [optional] 
**SmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**VsmfId** | Pointer to **string** |  | [optional] 
**VsmfSetId** | Pointer to **string** |  | [optional] 
**VsmfServiceSetId** | Pointer to **string** |  | [optional] 
**VsmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**IsmfId** | Pointer to **string** |  | [optional] 
**IsmfSetId** | Pointer to **string** |  | [optional] 
**IsmfServiceSetId** | Pointer to **string** |  | [optional] 
**IsmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**NsInstance** | Pointer to **string** |  | [optional] 
**SmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**MaPduSession** | Pointer to **bool** |  | [optional] [default to false]
**CnAssistedRanPara** | Pointer to [**CnAssistedRanPara**](CnAssistedRanPara.md) |  | [optional] 

## Methods

### NewPduSessionContext

`func NewPduSessionContext(pduSessionId int32, smContextRef string, sNssai Snssai, dnn string, accessType AccessType, ) *PduSessionContext`

NewPduSessionContext instantiates a new PduSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionContextWithDefaults

`func NewPduSessionContextWithDefaults() *PduSessionContext`

NewPduSessionContextWithDefaults instantiates a new PduSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *PduSessionContext) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduSessionContext) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduSessionContext) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetSmContextRef

`func (o *PduSessionContext) GetSmContextRef() string`

GetSmContextRef returns the SmContextRef field if non-nil, zero value otherwise.

### GetSmContextRefOk

`func (o *PduSessionContext) GetSmContextRefOk() (*string, bool)`

GetSmContextRefOk returns a tuple with the SmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextRef

`func (o *PduSessionContext) SetSmContextRef(v string)`

SetSmContextRef sets SmContextRef field to given value.


### GetSNssai

`func (o *PduSessionContext) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *PduSessionContext) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *PduSessionContext) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetDnn

`func (o *PduSessionContext) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionContext) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionContext) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSelectedDnn

`func (o *PduSessionContext) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *PduSessionContext) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *PduSessionContext) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *PduSessionContext) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetAccessType

`func (o *PduSessionContext) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PduSessionContext) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PduSessionContext) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetAdditionalAccessType

`func (o *PduSessionContext) GetAdditionalAccessType() AccessType`

GetAdditionalAccessType returns the AdditionalAccessType field if non-nil, zero value otherwise.

### GetAdditionalAccessTypeOk

`func (o *PduSessionContext) GetAdditionalAccessTypeOk() (*AccessType, bool)`

GetAdditionalAccessTypeOk returns a tuple with the AdditionalAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAccessType

`func (o *PduSessionContext) SetAdditionalAccessType(v AccessType)`

SetAdditionalAccessType sets AdditionalAccessType field to given value.

### HasAdditionalAccessType

`func (o *PduSessionContext) HasAdditionalAccessType() bool`

HasAdditionalAccessType returns a boolean if a field has been set.

### GetAllocatedEbiList

`func (o *PduSessionContext) GetAllocatedEbiList() []EbiArpMapping`

GetAllocatedEbiList returns the AllocatedEbiList field if non-nil, zero value otherwise.

### GetAllocatedEbiListOk

`func (o *PduSessionContext) GetAllocatedEbiListOk() (*[]EbiArpMapping, bool)`

GetAllocatedEbiListOk returns a tuple with the AllocatedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedEbiList

`func (o *PduSessionContext) SetAllocatedEbiList(v []EbiArpMapping)`

SetAllocatedEbiList sets AllocatedEbiList field to given value.

### HasAllocatedEbiList

`func (o *PduSessionContext) HasAllocatedEbiList() bool`

HasAllocatedEbiList returns a boolean if a field has been set.

### GetHsmfId

`func (o *PduSessionContext) GetHsmfId() string`

GetHsmfId returns the HsmfId field if non-nil, zero value otherwise.

### GetHsmfIdOk

`func (o *PduSessionContext) GetHsmfIdOk() (*string, bool)`

GetHsmfIdOk returns a tuple with the HsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfId

`func (o *PduSessionContext) SetHsmfId(v string)`

SetHsmfId sets HsmfId field to given value.

### HasHsmfId

`func (o *PduSessionContext) HasHsmfId() bool`

HasHsmfId returns a boolean if a field has been set.

### GetHsmfSetId

`func (o *PduSessionContext) GetHsmfSetId() string`

GetHsmfSetId returns the HsmfSetId field if non-nil, zero value otherwise.

### GetHsmfSetIdOk

`func (o *PduSessionContext) GetHsmfSetIdOk() (*string, bool)`

GetHsmfSetIdOk returns a tuple with the HsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfSetId

`func (o *PduSessionContext) SetHsmfSetId(v string)`

SetHsmfSetId sets HsmfSetId field to given value.

### HasHsmfSetId

`func (o *PduSessionContext) HasHsmfSetId() bool`

HasHsmfSetId returns a boolean if a field has been set.

### GetHsmfServiceSetId

`func (o *PduSessionContext) GetHsmfServiceSetId() string`

GetHsmfServiceSetId returns the HsmfServiceSetId field if non-nil, zero value otherwise.

### GetHsmfServiceSetIdOk

`func (o *PduSessionContext) GetHsmfServiceSetIdOk() (*string, bool)`

GetHsmfServiceSetIdOk returns a tuple with the HsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfServiceSetId

`func (o *PduSessionContext) SetHsmfServiceSetId(v string)`

SetHsmfServiceSetId sets HsmfServiceSetId field to given value.

### HasHsmfServiceSetId

`func (o *PduSessionContext) HasHsmfServiceSetId() bool`

HasHsmfServiceSetId returns a boolean if a field has been set.

### GetSmfBinding

`func (o *PduSessionContext) GetSmfBinding() SbiBindingLevel`

GetSmfBinding returns the SmfBinding field if non-nil, zero value otherwise.

### GetSmfBindingOk

`func (o *PduSessionContext) GetSmfBindingOk() (*SbiBindingLevel, bool)`

GetSmfBindingOk returns a tuple with the SmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfBinding

`func (o *PduSessionContext) SetSmfBinding(v SbiBindingLevel)`

SetSmfBinding sets SmfBinding field to given value.

### HasSmfBinding

`func (o *PduSessionContext) HasSmfBinding() bool`

HasSmfBinding returns a boolean if a field has been set.

### GetVsmfId

`func (o *PduSessionContext) GetVsmfId() string`

GetVsmfId returns the VsmfId field if non-nil, zero value otherwise.

### GetVsmfIdOk

`func (o *PduSessionContext) GetVsmfIdOk() (*string, bool)`

GetVsmfIdOk returns a tuple with the VsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfId

`func (o *PduSessionContext) SetVsmfId(v string)`

SetVsmfId sets VsmfId field to given value.

### HasVsmfId

`func (o *PduSessionContext) HasVsmfId() bool`

HasVsmfId returns a boolean if a field has been set.

### GetVsmfSetId

`func (o *PduSessionContext) GetVsmfSetId() string`

GetVsmfSetId returns the VsmfSetId field if non-nil, zero value otherwise.

### GetVsmfSetIdOk

`func (o *PduSessionContext) GetVsmfSetIdOk() (*string, bool)`

GetVsmfSetIdOk returns a tuple with the VsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfSetId

`func (o *PduSessionContext) SetVsmfSetId(v string)`

SetVsmfSetId sets VsmfSetId field to given value.

### HasVsmfSetId

`func (o *PduSessionContext) HasVsmfSetId() bool`

HasVsmfSetId returns a boolean if a field has been set.

### GetVsmfServiceSetId

`func (o *PduSessionContext) GetVsmfServiceSetId() string`

GetVsmfServiceSetId returns the VsmfServiceSetId field if non-nil, zero value otherwise.

### GetVsmfServiceSetIdOk

`func (o *PduSessionContext) GetVsmfServiceSetIdOk() (*string, bool)`

GetVsmfServiceSetIdOk returns a tuple with the VsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfServiceSetId

`func (o *PduSessionContext) SetVsmfServiceSetId(v string)`

SetVsmfServiceSetId sets VsmfServiceSetId field to given value.

### HasVsmfServiceSetId

`func (o *PduSessionContext) HasVsmfServiceSetId() bool`

HasVsmfServiceSetId returns a boolean if a field has been set.

### GetVsmfBinding

`func (o *PduSessionContext) GetVsmfBinding() SbiBindingLevel`

GetVsmfBinding returns the VsmfBinding field if non-nil, zero value otherwise.

### GetVsmfBindingOk

`func (o *PduSessionContext) GetVsmfBindingOk() (*SbiBindingLevel, bool)`

GetVsmfBindingOk returns a tuple with the VsmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfBinding

`func (o *PduSessionContext) SetVsmfBinding(v SbiBindingLevel)`

SetVsmfBinding sets VsmfBinding field to given value.

### HasVsmfBinding

`func (o *PduSessionContext) HasVsmfBinding() bool`

HasVsmfBinding returns a boolean if a field has been set.

### GetIsmfId

`func (o *PduSessionContext) GetIsmfId() string`

GetIsmfId returns the IsmfId field if non-nil, zero value otherwise.

### GetIsmfIdOk

`func (o *PduSessionContext) GetIsmfIdOk() (*string, bool)`

GetIsmfIdOk returns a tuple with the IsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfId

`func (o *PduSessionContext) SetIsmfId(v string)`

SetIsmfId sets IsmfId field to given value.

### HasIsmfId

`func (o *PduSessionContext) HasIsmfId() bool`

HasIsmfId returns a boolean if a field has been set.

### GetIsmfSetId

`func (o *PduSessionContext) GetIsmfSetId() string`

GetIsmfSetId returns the IsmfSetId field if non-nil, zero value otherwise.

### GetIsmfSetIdOk

`func (o *PduSessionContext) GetIsmfSetIdOk() (*string, bool)`

GetIsmfSetIdOk returns a tuple with the IsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfSetId

`func (o *PduSessionContext) SetIsmfSetId(v string)`

SetIsmfSetId sets IsmfSetId field to given value.

### HasIsmfSetId

`func (o *PduSessionContext) HasIsmfSetId() bool`

HasIsmfSetId returns a boolean if a field has been set.

### GetIsmfServiceSetId

`func (o *PduSessionContext) GetIsmfServiceSetId() string`

GetIsmfServiceSetId returns the IsmfServiceSetId field if non-nil, zero value otherwise.

### GetIsmfServiceSetIdOk

`func (o *PduSessionContext) GetIsmfServiceSetIdOk() (*string, bool)`

GetIsmfServiceSetIdOk returns a tuple with the IsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfServiceSetId

`func (o *PduSessionContext) SetIsmfServiceSetId(v string)`

SetIsmfServiceSetId sets IsmfServiceSetId field to given value.

### HasIsmfServiceSetId

`func (o *PduSessionContext) HasIsmfServiceSetId() bool`

HasIsmfServiceSetId returns a boolean if a field has been set.

### GetIsmfBinding

`func (o *PduSessionContext) GetIsmfBinding() SbiBindingLevel`

GetIsmfBinding returns the IsmfBinding field if non-nil, zero value otherwise.

### GetIsmfBindingOk

`func (o *PduSessionContext) GetIsmfBindingOk() (*SbiBindingLevel, bool)`

GetIsmfBindingOk returns a tuple with the IsmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfBinding

`func (o *PduSessionContext) SetIsmfBinding(v SbiBindingLevel)`

SetIsmfBinding sets IsmfBinding field to given value.

### HasIsmfBinding

`func (o *PduSessionContext) HasIsmfBinding() bool`

HasIsmfBinding returns a boolean if a field has been set.

### GetNsInstance

`func (o *PduSessionContext) GetNsInstance() string`

GetNsInstance returns the NsInstance field if non-nil, zero value otherwise.

### GetNsInstanceOk

`func (o *PduSessionContext) GetNsInstanceOk() (*string, bool)`

GetNsInstanceOk returns a tuple with the NsInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsInstance

`func (o *PduSessionContext) SetNsInstance(v string)`

SetNsInstance sets NsInstance field to given value.

### HasNsInstance

`func (o *PduSessionContext) HasNsInstance() bool`

HasNsInstance returns a boolean if a field has been set.

### GetSmfServiceInstanceId

`func (o *PduSessionContext) GetSmfServiceInstanceId() string`

GetSmfServiceInstanceId returns the SmfServiceInstanceId field if non-nil, zero value otherwise.

### GetSmfServiceInstanceIdOk

`func (o *PduSessionContext) GetSmfServiceInstanceIdOk() (*string, bool)`

GetSmfServiceInstanceIdOk returns a tuple with the SmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServiceInstanceId

`func (o *PduSessionContext) SetSmfServiceInstanceId(v string)`

SetSmfServiceInstanceId sets SmfServiceInstanceId field to given value.

### HasSmfServiceInstanceId

`func (o *PduSessionContext) HasSmfServiceInstanceId() bool`

HasSmfServiceInstanceId returns a boolean if a field has been set.

### GetMaPduSession

`func (o *PduSessionContext) GetMaPduSession() bool`

GetMaPduSession returns the MaPduSession field if non-nil, zero value otherwise.

### GetMaPduSessionOk

`func (o *PduSessionContext) GetMaPduSessionOk() (*bool, bool)`

GetMaPduSessionOk returns a tuple with the MaPduSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaPduSession

`func (o *PduSessionContext) SetMaPduSession(v bool)`

SetMaPduSession sets MaPduSession field to given value.

### HasMaPduSession

`func (o *PduSessionContext) HasMaPduSession() bool`

HasMaPduSession returns a boolean if a field has been set.

### GetCnAssistedRanPara

`func (o *PduSessionContext) GetCnAssistedRanPara() CnAssistedRanPara`

GetCnAssistedRanPara returns the CnAssistedRanPara field if non-nil, zero value otherwise.

### GetCnAssistedRanParaOk

`func (o *PduSessionContext) GetCnAssistedRanParaOk() (*CnAssistedRanPara, bool)`

GetCnAssistedRanParaOk returns a tuple with the CnAssistedRanPara field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnAssistedRanPara

`func (o *PduSessionContext) SetCnAssistedRanPara(v CnAssistedRanPara)`

SetCnAssistedRanPara sets CnAssistedRanPara field to given value.

### HasCnAssistedRanPara

`func (o *PduSessionContext) HasCnAssistedRanPara() bool`

HasCnAssistedRanPara returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


