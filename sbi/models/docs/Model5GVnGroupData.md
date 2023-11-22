# Model5GVnGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AppDescriptors** | Pointer to [**[]AppDescriptor**](AppDescriptor.md) |  | [optional] 
**SecondaryAuth** | Pointer to **bool** |  | [optional] 
**DnAaaAddress** | Pointer to [**IpAddress1**](IpAddress1.md) |  | [optional] 

## Methods

### NewModel5GVnGroupData

`func NewModel5GVnGroupData(dnn string, sNssai Snssai, ) *Model5GVnGroupData`

NewModel5GVnGroupData instantiates a new Model5GVnGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GVnGroupDataWithDefaults

`func NewModel5GVnGroupDataWithDefaults() *Model5GVnGroupData`

NewModel5GVnGroupDataWithDefaults instantiates a new Model5GVnGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *Model5GVnGroupData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Model5GVnGroupData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Model5GVnGroupData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSNssai

`func (o *Model5GVnGroupData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *Model5GVnGroupData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *Model5GVnGroupData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetPduSessionTypes

`func (o *Model5GVnGroupData) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *Model5GVnGroupData) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *Model5GVnGroupData) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *Model5GVnGroupData) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetAppDescriptors

`func (o *Model5GVnGroupData) GetAppDescriptors() []AppDescriptor`

GetAppDescriptors returns the AppDescriptors field if non-nil, zero value otherwise.

### GetAppDescriptorsOk

`func (o *Model5GVnGroupData) GetAppDescriptorsOk() (*[]AppDescriptor, bool)`

GetAppDescriptorsOk returns a tuple with the AppDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptors

`func (o *Model5GVnGroupData) SetAppDescriptors(v []AppDescriptor)`

SetAppDescriptors sets AppDescriptors field to given value.

### HasAppDescriptors

`func (o *Model5GVnGroupData) HasAppDescriptors() bool`

HasAppDescriptors returns a boolean if a field has been set.

### GetSecondaryAuth

`func (o *Model5GVnGroupData) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *Model5GVnGroupData) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *Model5GVnGroupData) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *Model5GVnGroupData) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *Model5GVnGroupData) GetDnAaaAddress() IpAddress1`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *Model5GVnGroupData) GetDnAaaAddressOk() (*IpAddress1, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *Model5GVnGroupData) SetDnAaaAddress(v IpAddress1)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *Model5GVnGroupData) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


