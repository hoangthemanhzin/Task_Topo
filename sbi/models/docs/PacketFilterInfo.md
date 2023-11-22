# PacketFilterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackFiltId** | Pointer to **string** | An identifier of packet filter. | [optional] 
**PackFiltCont** | Pointer to **string** | Defines a packet filter for an IP flow. | [optional] 
**TosTrafficClass** | Pointer to **string** | Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field. | [optional] 
**Spi** | Pointer to **string** | The security parameter index of the IPSec packet. | [optional] 
**FlowLabel** | Pointer to **string** | The Ipv6 flow label header field. | [optional] 
**FlowDirection** | Pointer to [**FlowDirection**](FlowDirection.md) |  | [optional] 

## Methods

### NewPacketFilterInfo

`func NewPacketFilterInfo() *PacketFilterInfo`

NewPacketFilterInfo instantiates a new PacketFilterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketFilterInfoWithDefaults

`func NewPacketFilterInfoWithDefaults() *PacketFilterInfo`

NewPacketFilterInfoWithDefaults instantiates a new PacketFilterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackFiltId

`func (o *PacketFilterInfo) GetPackFiltId() string`

GetPackFiltId returns the PackFiltId field if non-nil, zero value otherwise.

### GetPackFiltIdOk

`func (o *PacketFilterInfo) GetPackFiltIdOk() (*string, bool)`

GetPackFiltIdOk returns a tuple with the PackFiltId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackFiltId

`func (o *PacketFilterInfo) SetPackFiltId(v string)`

SetPackFiltId sets PackFiltId field to given value.

### HasPackFiltId

`func (o *PacketFilterInfo) HasPackFiltId() bool`

HasPackFiltId returns a boolean if a field has been set.

### GetPackFiltCont

`func (o *PacketFilterInfo) GetPackFiltCont() string`

GetPackFiltCont returns the PackFiltCont field if non-nil, zero value otherwise.

### GetPackFiltContOk

`func (o *PacketFilterInfo) GetPackFiltContOk() (*string, bool)`

GetPackFiltContOk returns a tuple with the PackFiltCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackFiltCont

`func (o *PacketFilterInfo) SetPackFiltCont(v string)`

SetPackFiltCont sets PackFiltCont field to given value.

### HasPackFiltCont

`func (o *PacketFilterInfo) HasPackFiltCont() bool`

HasPackFiltCont returns a boolean if a field has been set.

### GetTosTrafficClass

`func (o *PacketFilterInfo) GetTosTrafficClass() string`

GetTosTrafficClass returns the TosTrafficClass field if non-nil, zero value otherwise.

### GetTosTrafficClassOk

`func (o *PacketFilterInfo) GetTosTrafficClassOk() (*string, bool)`

GetTosTrafficClassOk returns a tuple with the TosTrafficClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosTrafficClass

`func (o *PacketFilterInfo) SetTosTrafficClass(v string)`

SetTosTrafficClass sets TosTrafficClass field to given value.

### HasTosTrafficClass

`func (o *PacketFilterInfo) HasTosTrafficClass() bool`

HasTosTrafficClass returns a boolean if a field has been set.

### GetSpi

`func (o *PacketFilterInfo) GetSpi() string`

GetSpi returns the Spi field if non-nil, zero value otherwise.

### GetSpiOk

`func (o *PacketFilterInfo) GetSpiOk() (*string, bool)`

GetSpiOk returns a tuple with the Spi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpi

`func (o *PacketFilterInfo) SetSpi(v string)`

SetSpi sets Spi field to given value.

### HasSpi

`func (o *PacketFilterInfo) HasSpi() bool`

HasSpi returns a boolean if a field has been set.

### GetFlowLabel

`func (o *PacketFilterInfo) GetFlowLabel() string`

GetFlowLabel returns the FlowLabel field if non-nil, zero value otherwise.

### GetFlowLabelOk

`func (o *PacketFilterInfo) GetFlowLabelOk() (*string, bool)`

GetFlowLabelOk returns a tuple with the FlowLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowLabel

`func (o *PacketFilterInfo) SetFlowLabel(v string)`

SetFlowLabel sets FlowLabel field to given value.

### HasFlowLabel

`func (o *PacketFilterInfo) HasFlowLabel() bool`

HasFlowLabel returns a boolean if a field has been set.

### GetFlowDirection

`func (o *PacketFilterInfo) GetFlowDirection() FlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *PacketFilterInfo) GetFlowDirectionOk() (*FlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *PacketFilterInfo) SetFlowDirection(v FlowDirection)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *PacketFilterInfo) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


