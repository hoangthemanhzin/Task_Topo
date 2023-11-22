# DnnConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionTypes** | [**PduSessionTypes**](PduSessionTypes.md) |  | 
**SscModes** | [**SscModes**](SscModes.md) |  | 
**IwkEpsInd** | Pointer to **bool** |  | [optional] 
**Var5gQosProfile** | Pointer to [**SubscribedDefaultQos**](SubscribedDefaultQos.md) |  | [optional] 
**SessionAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**Var3gppChargingCharacteristics** | Pointer to **string** |  | [optional] 
**StaticIpAddress** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**UpSecurity** | Pointer to [**UpSecurity**](UpSecurity.md) |  | [optional] 
**PduSessionContinuityInd** | Pointer to [**PduSessionContinuityInd**](PduSessionContinuityInd.md) |  | [optional] 
**NiddNefId** | Pointer to **string** | Identity of the NEF | [optional] 
**NiddInfo** | Pointer to [**NiddInformation**](NiddInformation.md) |  | [optional] 
**RedundantSessionAllowed** | Pointer to **bool** |  | [optional] 
**AcsInfo** | Pointer to [**AcsInfo**](AcsInfo.md) |  | [optional] 
**Ipv4FrameRouteList** | Pointer to [**[]FrameRouteInfo**](FrameRouteInfo.md) |  | [optional] 
**Ipv6FrameRouteList** | Pointer to [**[]FrameRouteInfo**](FrameRouteInfo.md) |  | [optional] 
**AtsssAllowed** | Pointer to **bool** |  | [optional] [default to false]
**SecondaryAuth** | Pointer to **bool** |  | [optional] 
**DnAaaIpAddressAllocation** | Pointer to **bool** |  | [optional] 
**DnAaaAddress** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**IptvAccCtrlInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewDnnConfiguration

`func NewDnnConfiguration(pduSessionTypes PduSessionTypes, sscModes SscModes, ) *DnnConfiguration`

NewDnnConfiguration instantiates a new DnnConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnConfigurationWithDefaults

`func NewDnnConfigurationWithDefaults() *DnnConfiguration`

NewDnnConfigurationWithDefaults instantiates a new DnnConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionTypes

`func (o *DnnConfiguration) GetPduSessionTypes() PduSessionTypes`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *DnnConfiguration) GetPduSessionTypesOk() (*PduSessionTypes, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *DnnConfiguration) SetPduSessionTypes(v PduSessionTypes)`

SetPduSessionTypes sets PduSessionTypes field to given value.


### GetSscModes

`func (o *DnnConfiguration) GetSscModes() SscModes`

GetSscModes returns the SscModes field if non-nil, zero value otherwise.

### GetSscModesOk

`func (o *DnnConfiguration) GetSscModesOk() (*SscModes, bool)`

GetSscModesOk returns a tuple with the SscModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSscModes

`func (o *DnnConfiguration) SetSscModes(v SscModes)`

SetSscModes sets SscModes field to given value.


### GetIwkEpsInd

`func (o *DnnConfiguration) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *DnnConfiguration) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *DnnConfiguration) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *DnnConfiguration) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetVar5gQosProfile

`func (o *DnnConfiguration) GetVar5gQosProfile() SubscribedDefaultQos`

GetVar5gQosProfile returns the Var5gQosProfile field if non-nil, zero value otherwise.

### GetVar5gQosProfileOk

`func (o *DnnConfiguration) GetVar5gQosProfileOk() (*SubscribedDefaultQos, bool)`

GetVar5gQosProfileOk returns a tuple with the Var5gQosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gQosProfile

`func (o *DnnConfiguration) SetVar5gQosProfile(v SubscribedDefaultQos)`

SetVar5gQosProfile sets Var5gQosProfile field to given value.

### HasVar5gQosProfile

`func (o *DnnConfiguration) HasVar5gQosProfile() bool`

HasVar5gQosProfile returns a boolean if a field has been set.

### GetSessionAmbr

`func (o *DnnConfiguration) GetSessionAmbr() Ambr`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *DnnConfiguration) GetSessionAmbrOk() (*Ambr, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *DnnConfiguration) SetSessionAmbr(v Ambr)`

SetSessionAmbr sets SessionAmbr field to given value.

### HasSessionAmbr

`func (o *DnnConfiguration) HasSessionAmbr() bool`

HasSessionAmbr returns a boolean if a field has been set.

### GetVar3gppChargingCharacteristics

`func (o *DnnConfiguration) GetVar3gppChargingCharacteristics() string`

GetVar3gppChargingCharacteristics returns the Var3gppChargingCharacteristics field if non-nil, zero value otherwise.

### GetVar3gppChargingCharacteristicsOk

`func (o *DnnConfiguration) GetVar3gppChargingCharacteristicsOk() (*string, bool)`

GetVar3gppChargingCharacteristicsOk returns a tuple with the Var3gppChargingCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppChargingCharacteristics

`func (o *DnnConfiguration) SetVar3gppChargingCharacteristics(v string)`

SetVar3gppChargingCharacteristics sets Var3gppChargingCharacteristics field to given value.

### HasVar3gppChargingCharacteristics

`func (o *DnnConfiguration) HasVar3gppChargingCharacteristics() bool`

HasVar3gppChargingCharacteristics returns a boolean if a field has been set.

### GetStaticIpAddress

`func (o *DnnConfiguration) GetStaticIpAddress() []IpAddress`

GetStaticIpAddress returns the StaticIpAddress field if non-nil, zero value otherwise.

### GetStaticIpAddressOk

`func (o *DnnConfiguration) GetStaticIpAddressOk() (*[]IpAddress, bool)`

GetStaticIpAddressOk returns a tuple with the StaticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpAddress

`func (o *DnnConfiguration) SetStaticIpAddress(v []IpAddress)`

SetStaticIpAddress sets StaticIpAddress field to given value.

### HasStaticIpAddress

`func (o *DnnConfiguration) HasStaticIpAddress() bool`

HasStaticIpAddress returns a boolean if a field has been set.

### GetUpSecurity

`func (o *DnnConfiguration) GetUpSecurity() UpSecurity`

GetUpSecurity returns the UpSecurity field if non-nil, zero value otherwise.

### GetUpSecurityOk

`func (o *DnnConfiguration) GetUpSecurityOk() (*UpSecurity, bool)`

GetUpSecurityOk returns a tuple with the UpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurity

`func (o *DnnConfiguration) SetUpSecurity(v UpSecurity)`

SetUpSecurity sets UpSecurity field to given value.

### HasUpSecurity

`func (o *DnnConfiguration) HasUpSecurity() bool`

HasUpSecurity returns a boolean if a field has been set.

### GetPduSessionContinuityInd

`func (o *DnnConfiguration) GetPduSessionContinuityInd() PduSessionContinuityInd`

GetPduSessionContinuityInd returns the PduSessionContinuityInd field if non-nil, zero value otherwise.

### GetPduSessionContinuityIndOk

`func (o *DnnConfiguration) GetPduSessionContinuityIndOk() (*PduSessionContinuityInd, bool)`

GetPduSessionContinuityIndOk returns a tuple with the PduSessionContinuityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionContinuityInd

`func (o *DnnConfiguration) SetPduSessionContinuityInd(v PduSessionContinuityInd)`

SetPduSessionContinuityInd sets PduSessionContinuityInd field to given value.

### HasPduSessionContinuityInd

`func (o *DnnConfiguration) HasPduSessionContinuityInd() bool`

HasPduSessionContinuityInd returns a boolean if a field has been set.

### GetNiddNefId

`func (o *DnnConfiguration) GetNiddNefId() string`

GetNiddNefId returns the NiddNefId field if non-nil, zero value otherwise.

### GetNiddNefIdOk

`func (o *DnnConfiguration) GetNiddNefIdOk() (*string, bool)`

GetNiddNefIdOk returns a tuple with the NiddNefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddNefId

`func (o *DnnConfiguration) SetNiddNefId(v string)`

SetNiddNefId sets NiddNefId field to given value.

### HasNiddNefId

`func (o *DnnConfiguration) HasNiddNefId() bool`

HasNiddNefId returns a boolean if a field has been set.

### GetNiddInfo

`func (o *DnnConfiguration) GetNiddInfo() NiddInformation`

GetNiddInfo returns the NiddInfo field if non-nil, zero value otherwise.

### GetNiddInfoOk

`func (o *DnnConfiguration) GetNiddInfoOk() (*NiddInformation, bool)`

GetNiddInfoOk returns a tuple with the NiddInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddInfo

`func (o *DnnConfiguration) SetNiddInfo(v NiddInformation)`

SetNiddInfo sets NiddInfo field to given value.

### HasNiddInfo

`func (o *DnnConfiguration) HasNiddInfo() bool`

HasNiddInfo returns a boolean if a field has been set.

### GetRedundantSessionAllowed

`func (o *DnnConfiguration) GetRedundantSessionAllowed() bool`

GetRedundantSessionAllowed returns the RedundantSessionAllowed field if non-nil, zero value otherwise.

### GetRedundantSessionAllowedOk

`func (o *DnnConfiguration) GetRedundantSessionAllowedOk() (*bool, bool)`

GetRedundantSessionAllowedOk returns a tuple with the RedundantSessionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantSessionAllowed

`func (o *DnnConfiguration) SetRedundantSessionAllowed(v bool)`

SetRedundantSessionAllowed sets RedundantSessionAllowed field to given value.

### HasRedundantSessionAllowed

`func (o *DnnConfiguration) HasRedundantSessionAllowed() bool`

HasRedundantSessionAllowed returns a boolean if a field has been set.

### GetAcsInfo

`func (o *DnnConfiguration) GetAcsInfo() AcsInfo`

GetAcsInfo returns the AcsInfo field if non-nil, zero value otherwise.

### GetAcsInfoOk

`func (o *DnnConfiguration) GetAcsInfoOk() (*AcsInfo, bool)`

GetAcsInfoOk returns a tuple with the AcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfo

`func (o *DnnConfiguration) SetAcsInfo(v AcsInfo)`

SetAcsInfo sets AcsInfo field to given value.

### HasAcsInfo

`func (o *DnnConfiguration) HasAcsInfo() bool`

HasAcsInfo returns a boolean if a field has been set.

### GetIpv4FrameRouteList

`func (o *DnnConfiguration) GetIpv4FrameRouteList() []FrameRouteInfo`

GetIpv4FrameRouteList returns the Ipv4FrameRouteList field if non-nil, zero value otherwise.

### GetIpv4FrameRouteListOk

`func (o *DnnConfiguration) GetIpv4FrameRouteListOk() (*[]FrameRouteInfo, bool)`

GetIpv4FrameRouteListOk returns a tuple with the Ipv4FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4FrameRouteList

`func (o *DnnConfiguration) SetIpv4FrameRouteList(v []FrameRouteInfo)`

SetIpv4FrameRouteList sets Ipv4FrameRouteList field to given value.

### HasIpv4FrameRouteList

`func (o *DnnConfiguration) HasIpv4FrameRouteList() bool`

HasIpv4FrameRouteList returns a boolean if a field has been set.

### GetIpv6FrameRouteList

`func (o *DnnConfiguration) GetIpv6FrameRouteList() []FrameRouteInfo`

GetIpv6FrameRouteList returns the Ipv6FrameRouteList field if non-nil, zero value otherwise.

### GetIpv6FrameRouteListOk

`func (o *DnnConfiguration) GetIpv6FrameRouteListOk() (*[]FrameRouteInfo, bool)`

GetIpv6FrameRouteListOk returns a tuple with the Ipv6FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6FrameRouteList

`func (o *DnnConfiguration) SetIpv6FrameRouteList(v []FrameRouteInfo)`

SetIpv6FrameRouteList sets Ipv6FrameRouteList field to given value.

### HasIpv6FrameRouteList

`func (o *DnnConfiguration) HasIpv6FrameRouteList() bool`

HasIpv6FrameRouteList returns a boolean if a field has been set.

### GetAtsssAllowed

`func (o *DnnConfiguration) GetAtsssAllowed() bool`

GetAtsssAllowed returns the AtsssAllowed field if non-nil, zero value otherwise.

### GetAtsssAllowedOk

`func (o *DnnConfiguration) GetAtsssAllowedOk() (*bool, bool)`

GetAtsssAllowedOk returns a tuple with the AtsssAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssAllowed

`func (o *DnnConfiguration) SetAtsssAllowed(v bool)`

SetAtsssAllowed sets AtsssAllowed field to given value.

### HasAtsssAllowed

`func (o *DnnConfiguration) HasAtsssAllowed() bool`

HasAtsssAllowed returns a boolean if a field has been set.

### GetSecondaryAuth

`func (o *DnnConfiguration) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *DnnConfiguration) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *DnnConfiguration) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *DnnConfiguration) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### GetDnAaaIpAddressAllocation

`func (o *DnnConfiguration) GetDnAaaIpAddressAllocation() bool`

GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field if non-nil, zero value otherwise.

### GetDnAaaIpAddressAllocationOk

`func (o *DnnConfiguration) GetDnAaaIpAddressAllocationOk() (*bool, bool)`

GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaIpAddressAllocation

`func (o *DnnConfiguration) SetDnAaaIpAddressAllocation(v bool)`

SetDnAaaIpAddressAllocation sets DnAaaIpAddressAllocation field to given value.

### HasDnAaaIpAddressAllocation

`func (o *DnnConfiguration) HasDnAaaIpAddressAllocation() bool`

HasDnAaaIpAddressAllocation returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *DnnConfiguration) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *DnnConfiguration) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *DnnConfiguration) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *DnnConfiguration) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### GetIptvAccCtrlInfo

`func (o *DnnConfiguration) GetIptvAccCtrlInfo() string`

GetIptvAccCtrlInfo returns the IptvAccCtrlInfo field if non-nil, zero value otherwise.

### GetIptvAccCtrlInfoOk

`func (o *DnnConfiguration) GetIptvAccCtrlInfoOk() (*string, bool)`

GetIptvAccCtrlInfoOk returns a tuple with the IptvAccCtrlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvAccCtrlInfo

`func (o *DnnConfiguration) SetIptvAccCtrlInfo(v string)`

SetIptvAccCtrlInfo sets IptvAccCtrlInfo field to given value.

### HasIptvAccCtrlInfo

`func (o *DnnConfiguration) HasIptvAccCtrlInfo() bool`

HasIptvAccCtrlInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


