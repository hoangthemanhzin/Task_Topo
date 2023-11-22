# N2SmInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** |  | 
**N2InfoContent** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**HomePlmnSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**IwkSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SubjectToHo** | Pointer to **bool** |  | [optional] 

## Methods

### NewN2SmInformation

`func NewN2SmInformation(pduSessionId int32, ) *N2SmInformation`

NewN2SmInformation instantiates a new N2SmInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2SmInformationWithDefaults

`func NewN2SmInformationWithDefaults() *N2SmInformation`

NewN2SmInformationWithDefaults instantiates a new N2SmInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *N2SmInformation) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *N2SmInformation) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *N2SmInformation) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetN2InfoContent

`func (o *N2SmInformation) GetN2InfoContent() N2InfoContent`

GetN2InfoContent returns the N2InfoContent field if non-nil, zero value otherwise.

### GetN2InfoContentOk

`func (o *N2SmInformation) GetN2InfoContentOk() (*N2InfoContent, bool)`

GetN2InfoContentOk returns a tuple with the N2InfoContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InfoContent

`func (o *N2SmInformation) SetN2InfoContent(v N2InfoContent)`

SetN2InfoContent sets N2InfoContent field to given value.

### HasN2InfoContent

`func (o *N2SmInformation) HasN2InfoContent() bool`

HasN2InfoContent returns a boolean if a field has been set.

### GetSNssai

`func (o *N2SmInformation) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *N2SmInformation) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *N2SmInformation) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *N2SmInformation) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetHomePlmnSnssai

`func (o *N2SmInformation) GetHomePlmnSnssai() Snssai`

GetHomePlmnSnssai returns the HomePlmnSnssai field if non-nil, zero value otherwise.

### GetHomePlmnSnssaiOk

`func (o *N2SmInformation) GetHomePlmnSnssaiOk() (*Snssai, bool)`

GetHomePlmnSnssaiOk returns a tuple with the HomePlmnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePlmnSnssai

`func (o *N2SmInformation) SetHomePlmnSnssai(v Snssai)`

SetHomePlmnSnssai sets HomePlmnSnssai field to given value.

### HasHomePlmnSnssai

`func (o *N2SmInformation) HasHomePlmnSnssai() bool`

HasHomePlmnSnssai returns a boolean if a field has been set.

### GetIwkSnssai

`func (o *N2SmInformation) GetIwkSnssai() Snssai`

GetIwkSnssai returns the IwkSnssai field if non-nil, zero value otherwise.

### GetIwkSnssaiOk

`func (o *N2SmInformation) GetIwkSnssaiOk() (*Snssai, bool)`

GetIwkSnssaiOk returns a tuple with the IwkSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkSnssai

`func (o *N2SmInformation) SetIwkSnssai(v Snssai)`

SetIwkSnssai sets IwkSnssai field to given value.

### HasIwkSnssai

`func (o *N2SmInformation) HasIwkSnssai() bool`

HasIwkSnssai returns a boolean if a field has been set.

### GetSubjectToHo

`func (o *N2SmInformation) GetSubjectToHo() bool`

GetSubjectToHo returns the SubjectToHo field if non-nil, zero value otherwise.

### GetSubjectToHoOk

`func (o *N2SmInformation) GetSubjectToHoOk() (*bool, bool)`

GetSubjectToHoOk returns a tuple with the SubjectToHo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectToHo

`func (o *N2SmInformation) SetSubjectToHo(v bool)`

SetSubjectToHo sets SubjectToHo field to given value.

### HasSubjectToHo

`func (o *N2SmInformation) HasSubjectToHo() bool`

HasSubjectToHo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


