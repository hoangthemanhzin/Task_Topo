# VnGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionTypes** | Pointer to [**PduSessionTypes1**](PduSessionTypes1.md) |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppDescriptors** | Pointer to [**[]AppDescriptor**](AppDescriptor.md) |  | [optional] 

## Methods

### NewVnGroupData

`func NewVnGroupData() *VnGroupData`

NewVnGroupData instantiates a new VnGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnGroupDataWithDefaults

`func NewVnGroupDataWithDefaults() *VnGroupData`

NewVnGroupDataWithDefaults instantiates a new VnGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionTypes

`func (o *VnGroupData) GetPduSessionTypes() PduSessionTypes1`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *VnGroupData) GetPduSessionTypesOk() (*PduSessionTypes1, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *VnGroupData) SetPduSessionTypes(v PduSessionTypes1)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *VnGroupData) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetDnn

`func (o *VnGroupData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *VnGroupData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *VnGroupData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *VnGroupData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSingleNssai

`func (o *VnGroupData) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *VnGroupData) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *VnGroupData) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *VnGroupData) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetAppDescriptors

`func (o *VnGroupData) GetAppDescriptors() []AppDescriptor`

GetAppDescriptors returns the AppDescriptors field if non-nil, zero value otherwise.

### GetAppDescriptorsOk

`func (o *VnGroupData) GetAppDescriptorsOk() (*[]AppDescriptor, bool)`

GetAppDescriptorsOk returns a tuple with the AppDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptors

`func (o *VnGroupData) SetAppDescriptors(v []AppDescriptor)`

SetAppDescriptors sets AppDescriptors field to given value.

### HasAppDescriptors

`func (o *VnGroupData) HasAppDescriptors() bool`

HasAppDescriptors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


