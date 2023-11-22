# PfdContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PfdId** | Pointer to **string** | Identifies a PDF of an application identifier. | [optional] 
**FlowDescriptions** | Pointer to **[]string** | Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. | [optional] 
**Urls** | Pointer to **[]string** | Indicates a URL or a regular expression which is used to match the significant parts of the URL. | [optional] 
**DomainNames** | Pointer to **[]string** | Indicates an FQDN or a regular expression as a domain name matching criteria. | [optional] 
**DnProtocol** | Pointer to [**DomainNameProtocol**](DomainNameProtocol.md) |  | [optional] 

## Methods

### NewPfdContent

`func NewPfdContent() *PfdContent`

NewPfdContent instantiates a new PfdContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdContentWithDefaults

`func NewPfdContentWithDefaults() *PfdContent`

NewPfdContentWithDefaults instantiates a new PfdContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPfdId

`func (o *PfdContent) GetPfdId() string`

GetPfdId returns the PfdId field if non-nil, zero value otherwise.

### GetPfdIdOk

`func (o *PfdContent) GetPfdIdOk() (*string, bool)`

GetPfdIdOk returns a tuple with the PfdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdId

`func (o *PfdContent) SetPfdId(v string)`

SetPfdId sets PfdId field to given value.

### HasPfdId

`func (o *PfdContent) HasPfdId() bool`

HasPfdId returns a boolean if a field has been set.

### GetFlowDescriptions

`func (o *PfdContent) GetFlowDescriptions() []string`

GetFlowDescriptions returns the FlowDescriptions field if non-nil, zero value otherwise.

### GetFlowDescriptionsOk

`func (o *PfdContent) GetFlowDescriptionsOk() (*[]string, bool)`

GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescriptions

`func (o *PfdContent) SetFlowDescriptions(v []string)`

SetFlowDescriptions sets FlowDescriptions field to given value.

### HasFlowDescriptions

`func (o *PfdContent) HasFlowDescriptions() bool`

HasFlowDescriptions returns a boolean if a field has been set.

### GetUrls

`func (o *PfdContent) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PfdContent) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PfdContent) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *PfdContent) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetDomainNames

`func (o *PfdContent) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PfdContent) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PfdContent) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *PfdContent) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetDnProtocol

`func (o *PfdContent) GetDnProtocol() DomainNameProtocol`

GetDnProtocol returns the DnProtocol field if non-nil, zero value otherwise.

### GetDnProtocolOk

`func (o *PfdContent) GetDnProtocolOk() (*DomainNameProtocol, bool)`

GetDnProtocolOk returns a tuple with the DnProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnProtocol

`func (o *PfdContent) SetDnProtocol(v DomainNameProtocol)`

SetDnProtocol sets DnProtocol field to given value.

### HasDnProtocol

`func (o *PfdContent) HasDnProtocol() bool`

HasDnProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


