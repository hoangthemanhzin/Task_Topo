# FlowInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowDescription** | Pointer to **string** | Defines a packet filter for an IP flow. | [optional] 
**EthFlowDescription** | Pointer to [**EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**PackFiltId** | Pointer to **string** | An identifier of packet filter. | [optional] 
**PacketFilterUsage** | Pointer to **bool** | The packet shall be sent to the UE. | [optional] 
**TosTrafficClass** | Pointer to **string** | Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field. | [optional] 
**Spi** | Pointer to **string** | the security parameter index of the IPSec packet. | [optional] 
**FlowLabel** | Pointer to **string** | the Ipv6 flow label header field. | [optional] 
**FlowDirection** | Pointer to [**FlowDirectionRm**](FlowDirectionRm.md) |  | [optional] 

## Methods

### NewFlowInformation

`func NewFlowInformation() *FlowInformation`

NewFlowInformation instantiates a new FlowInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInformationWithDefaults

`func NewFlowInformationWithDefaults() *FlowInformation`

NewFlowInformationWithDefaults instantiates a new FlowInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowDescription

`func (o *FlowInformation) GetFlowDescription() string`

GetFlowDescription returns the FlowDescription field if non-nil, zero value otherwise.

### GetFlowDescriptionOk

`func (o *FlowInformation) GetFlowDescriptionOk() (*string, bool)`

GetFlowDescriptionOk returns a tuple with the FlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescription

`func (o *FlowInformation) SetFlowDescription(v string)`

SetFlowDescription sets FlowDescription field to given value.

### HasFlowDescription

`func (o *FlowInformation) HasFlowDescription() bool`

HasFlowDescription returns a boolean if a field has been set.

### GetEthFlowDescription

`func (o *FlowInformation) GetEthFlowDescription() EthFlowDescription`

GetEthFlowDescription returns the EthFlowDescription field if non-nil, zero value otherwise.

### GetEthFlowDescriptionOk

`func (o *FlowInformation) GetEthFlowDescriptionOk() (*EthFlowDescription, bool)`

GetEthFlowDescriptionOk returns a tuple with the EthFlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowDescription

`func (o *FlowInformation) SetEthFlowDescription(v EthFlowDescription)`

SetEthFlowDescription sets EthFlowDescription field to given value.

### HasEthFlowDescription

`func (o *FlowInformation) HasEthFlowDescription() bool`

HasEthFlowDescription returns a boolean if a field has been set.

### GetPackFiltId

`func (o *FlowInformation) GetPackFiltId() string`

GetPackFiltId returns the PackFiltId field if non-nil, zero value otherwise.

### GetPackFiltIdOk

`func (o *FlowInformation) GetPackFiltIdOk() (*string, bool)`

GetPackFiltIdOk returns a tuple with the PackFiltId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackFiltId

`func (o *FlowInformation) SetPackFiltId(v string)`

SetPackFiltId sets PackFiltId field to given value.

### HasPackFiltId

`func (o *FlowInformation) HasPackFiltId() bool`

HasPackFiltId returns a boolean if a field has been set.

### GetPacketFilterUsage

`func (o *FlowInformation) GetPacketFilterUsage() bool`

GetPacketFilterUsage returns the PacketFilterUsage field if non-nil, zero value otherwise.

### GetPacketFilterUsageOk

`func (o *FlowInformation) GetPacketFilterUsageOk() (*bool, bool)`

GetPacketFilterUsageOk returns a tuple with the PacketFilterUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketFilterUsage

`func (o *FlowInformation) SetPacketFilterUsage(v bool)`

SetPacketFilterUsage sets PacketFilterUsage field to given value.

### HasPacketFilterUsage

`func (o *FlowInformation) HasPacketFilterUsage() bool`

HasPacketFilterUsage returns a boolean if a field has been set.

### GetTosTrafficClass

`func (o *FlowInformation) GetTosTrafficClass() string`

GetTosTrafficClass returns the TosTrafficClass field if non-nil, zero value otherwise.

### GetTosTrafficClassOk

`func (o *FlowInformation) GetTosTrafficClassOk() (*string, bool)`

GetTosTrafficClassOk returns a tuple with the TosTrafficClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosTrafficClass

`func (o *FlowInformation) SetTosTrafficClass(v string)`

SetTosTrafficClass sets TosTrafficClass field to given value.

### HasTosTrafficClass

`func (o *FlowInformation) HasTosTrafficClass() bool`

HasTosTrafficClass returns a boolean if a field has been set.

### SetTosTrafficClassNil

`func (o *FlowInformation) SetTosTrafficClassNil(b bool)`

 SetTosTrafficClassNil sets the value for TosTrafficClass to be an explicit nil

### UnsetTosTrafficClass
`func (o *FlowInformation) UnsetTosTrafficClass()`

UnsetTosTrafficClass ensures that no value is present for TosTrafficClass, not even an explicit nil
### GetSpi

`func (o *FlowInformation) GetSpi() string`

GetSpi returns the Spi field if non-nil, zero value otherwise.

### GetSpiOk

`func (o *FlowInformation) GetSpiOk() (*string, bool)`

GetSpiOk returns a tuple with the Spi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpi

`func (o *FlowInformation) SetSpi(v string)`

SetSpi sets Spi field to given value.

### HasSpi

`func (o *FlowInformation) HasSpi() bool`

HasSpi returns a boolean if a field has been set.

### SetSpiNil

`func (o *FlowInformation) SetSpiNil(b bool)`

 SetSpiNil sets the value for Spi to be an explicit nil

### UnsetSpi
`func (o *FlowInformation) UnsetSpi()`

UnsetSpi ensures that no value is present for Spi, not even an explicit nil
### GetFlowLabel

`func (o *FlowInformation) GetFlowLabel() string`

GetFlowLabel returns the FlowLabel field if non-nil, zero value otherwise.

### GetFlowLabelOk

`func (o *FlowInformation) GetFlowLabelOk() (*string, bool)`

GetFlowLabelOk returns a tuple with the FlowLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowLabel

`func (o *FlowInformation) SetFlowLabel(v string)`

SetFlowLabel sets FlowLabel field to given value.

### HasFlowLabel

`func (o *FlowInformation) HasFlowLabel() bool`

HasFlowLabel returns a boolean if a field has been set.

### SetFlowLabelNil

`func (o *FlowInformation) SetFlowLabelNil(b bool)`

 SetFlowLabelNil sets the value for FlowLabel to be an explicit nil

### UnsetFlowLabel
`func (o *FlowInformation) UnsetFlowLabel()`

UnsetFlowLabel ensures that no value is present for FlowLabel, not even an explicit nil
### GetFlowDirection

`func (o *FlowInformation) GetFlowDirection() FlowDirectionRm`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *FlowInformation) GetFlowDirectionOk() (*FlowDirectionRm, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *FlowInformation) SetFlowDirection(v FlowDirectionRm)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *FlowInformation) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


