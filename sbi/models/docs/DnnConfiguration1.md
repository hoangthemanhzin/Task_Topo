# DnnConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionTypes** | [**PduSessionTypes1**](PduSessionTypes1.md) |  | 
**SscModes** | [**SscModes1**](SscModes1.md) |  | 
**IwkEpsInd** | Pointer to **bool** |  | [optional] 
**Var5gQosProfile** | Pointer to [**SubscribedDefaultQos1**](SubscribedDefaultQos1.md) |  | [optional] 
**SessionAmbr** | Pointer to [**Ambr1**](Ambr1.md) |  | [optional] 
**Var3gppChargingCharacteristics** | Pointer to **string** |  | [optional] 
**StaticIpAddress** | Pointer to [**[]IpAddress1**](IpAddress1.md) |  | [optional] 
**UpSecurity** | Pointer to [**UpSecurity1**](UpSecurity1.md) |  | [optional] 
**PduSessionContinuityInd** | Pointer to [**PduSessionContinuityInd**](PduSessionContinuityInd.md) |  | [optional] 
**NiddNefId** | Pointer to **string** | Identity of the NEF | [optional] 
**NiddInfo** | Pointer to [**NiddInformation1**](NiddInformation1.md) |  | [optional] 
**RedundantSessionAllowed** | Pointer to **bool** |  | [optional] 
**AcsInfo** | Pointer to [**AcsInfo1**](AcsInfo1.md) |  | [optional] 
**Ipv4FrameRouteList** | Pointer to [**[]FrameRouteInfo1**](FrameRouteInfo1.md) |  | [optional] 
**Ipv6FrameRouteList** | Pointer to [**[]FrameRouteInfo1**](FrameRouteInfo1.md) |  | [optional] 
**AtsssAllowed** | Pointer to **bool** |  | [optional] [default to false]
**SecondaryAuth** | Pointer to **bool** |  | [optional] 
**DnAaaIpAddressAllocation** | Pointer to **bool** |  | [optional] 
**DnAaaAddress** | Pointer to [**IpAddress1**](IpAddress1.md) |  | [optional] 
**IptvAccCtrlInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewDnnConfiguration1

`func NewDnnConfiguration1(pduSessionTypes PduSessionTypes1, sscModes SscModes1, ) *DnnConfiguration1`

NewDnnConfiguration1 instantiates a new DnnConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnConfiguration1WithDefaults

`func NewDnnConfiguration1WithDefaults() *DnnConfiguration1`

NewDnnConfiguration1WithDefaults instantiates a new DnnConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionTypes

`func (o *DnnConfiguration1) GetPduSessionTypes() PduSessionTypes1`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *DnnConfiguration1) GetPduSessionTypesOk() (*PduSessionTypes1, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *DnnConfiguration1) SetPduSessionTypes(v PduSessionTypes1)`

SetPduSessionTypes sets PduSessionTypes field to given value.


### GetSscModes

`func (o *DnnConfiguration1) GetSscModes() SscModes1`

GetSscModes returns the SscModes field if non-nil, zero value otherwise.

### GetSscModesOk

`func (o *DnnConfiguration1) GetSscModesOk() (*SscModes1, bool)`

GetSscModesOk returns a tuple with the SscModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSscModes

`func (o *DnnConfiguration1) SetSscModes(v SscModes1)`

SetSscModes sets SscModes field to given value.


### GetIwkEpsInd

`func (o *DnnConfiguration1) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *DnnConfiguration1) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *DnnConfiguration1) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *DnnConfiguration1) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetVar5gQosProfile

`func (o *DnnConfiguration1) GetVar5gQosProfile() SubscribedDefaultQos1`

GetVar5gQosProfile returns the Var5gQosProfile field if non-nil, zero value otherwise.

### GetVar5gQosProfileOk

`func (o *DnnConfiguration1) GetVar5gQosProfileOk() (*SubscribedDefaultQos1, bool)`

GetVar5gQosProfileOk returns a tuple with the Var5gQosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gQosProfile

`func (o *DnnConfiguration1) SetVar5gQosProfile(v SubscribedDefaultQos1)`

SetVar5gQosProfile sets Var5gQosProfile field to given value.

### HasVar5gQosProfile

`func (o *DnnConfiguration1) HasVar5gQosProfile() bool`

HasVar5gQosProfile returns a boolean if a field has been set.

### GetSessionAmbr

`func (o *DnnConfiguration1) GetSessionAmbr() Ambr1`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *DnnConfiguration1) GetSessionAmbrOk() (*Ambr1, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *DnnConfiguration1) SetSessionAmbr(v Ambr1)`

SetSessionAmbr sets SessionAmbr field to given value.

### HasSessionAmbr

`func (o *DnnConfiguration1) HasSessionAmbr() bool`

HasSessionAmbr returns a boolean if a field has been set.

### GetVar3gppChargingCharacteristics

`func (o *DnnConfiguration1) GetVar3gppChargingCharacteristics() string`

GetVar3gppChargingCharacteristics returns the Var3gppChargingCharacteristics field if non-nil, zero value otherwise.

### GetVar3gppChargingCharacteristicsOk

`func (o *DnnConfiguration1) GetVar3gppChargingCharacteristicsOk() (*string, bool)`

GetVar3gppChargingCharacteristicsOk returns a tuple with the Var3gppChargingCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppChargingCharacteristics

`func (o *DnnConfiguration1) SetVar3gppChargingCharacteristics(v string)`

SetVar3gppChargingCharacteristics sets Var3gppChargingCharacteristics field to given value.

### HasVar3gppChargingCharacteristics

`func (o *DnnConfiguration1) HasVar3gppChargingCharacteristics() bool`

HasVar3gppChargingCharacteristics returns a boolean if a field has been set.

### GetStaticIpAddress

`func (o *DnnConfiguration1) GetStaticIpAddress() []IpAddress1`

GetStaticIpAddress returns the StaticIpAddress field if non-nil, zero value otherwise.

### GetStaticIpAddressOk

`func (o *DnnConfiguration1) GetStaticIpAddressOk() (*[]IpAddress1, bool)`

GetStaticIpAddressOk returns a tuple with the StaticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpAddress

`func (o *DnnConfiguration1) SetStaticIpAddress(v []IpAddress1)`

SetStaticIpAddress sets StaticIpAddress field to given value.

### HasStaticIpAddress

`func (o *DnnConfiguration1) HasStaticIpAddress() bool`

HasStaticIpAddress returns a boolean if a field has been set.

### GetUpSecurity

`func (o *DnnConfiguration1) GetUpSecurity() UpSecurity1`

GetUpSecurity returns the UpSecurity field if non-nil, zero value otherwise.

### GetUpSecurityOk

`func (o *DnnConfiguration1) GetUpSecurityOk() (*UpSecurity1, bool)`

GetUpSecurityOk returns a tuple with the UpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurity

`func (o *DnnConfiguration1) SetUpSecurity(v UpSecurity1)`

SetUpSecurity sets UpSecurity field to given value.

### HasUpSecurity

`func (o *DnnConfiguration1) HasUpSecurity() bool`

HasUpSecurity returns a boolean if a field has been set.

### GetPduSessionContinuityInd

`func (o *DnnConfiguration1) GetPduSessionContinuityInd() PduSessionContinuityInd`

GetPduSessionContinuityInd returns the PduSessionContinuityInd field if non-nil, zero value otherwise.

### GetPduSessionContinuityIndOk

`func (o *DnnConfiguration1) GetPduSessionContinuityIndOk() (*PduSessionContinuityInd, bool)`

GetPduSessionContinuityIndOk returns a tuple with the PduSessionContinuityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionContinuityInd

`func (o *DnnConfiguration1) SetPduSessionContinuityInd(v PduSessionContinuityInd)`

SetPduSessionContinuityInd sets PduSessionContinuityInd field to given value.

### HasPduSessionContinuityInd

`func (o *DnnConfiguration1) HasPduSessionContinuityInd() bool`

HasPduSessionContinuityInd returns a boolean if a field has been set.

### GetNiddNefId

`func (o *DnnConfiguration1) GetNiddNefId() string`

GetNiddNefId returns the NiddNefId field if non-nil, zero value otherwise.

### GetNiddNefIdOk

`func (o *DnnConfiguration1) GetNiddNefIdOk() (*string, bool)`

GetNiddNefIdOk returns a tuple with the NiddNefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddNefId

`func (o *DnnConfiguration1) SetNiddNefId(v string)`

SetNiddNefId sets NiddNefId field to given value.

### HasNiddNefId

`func (o *DnnConfiguration1) HasNiddNefId() bool`

HasNiddNefId returns a boolean if a field has been set.

### GetNiddInfo

`func (o *DnnConfiguration1) GetNiddInfo() NiddInformation1`

GetNiddInfo returns the NiddInfo field if non-nil, zero value otherwise.

### GetNiddInfoOk

`func (o *DnnConfiguration1) GetNiddInfoOk() (*NiddInformation1, bool)`

GetNiddInfoOk returns a tuple with the NiddInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddInfo

`func (o *DnnConfiguration1) SetNiddInfo(v NiddInformation1)`

SetNiddInfo sets NiddInfo field to given value.

### HasNiddInfo

`func (o *DnnConfiguration1) HasNiddInfo() bool`

HasNiddInfo returns a boolean if a field has been set.

### GetRedundantSessionAllowed

`func (o *DnnConfiguration1) GetRedundantSessionAllowed() bool`

GetRedundantSessionAllowed returns the RedundantSessionAllowed field if non-nil, zero value otherwise.

### GetRedundantSessionAllowedOk

`func (o *DnnConfiguration1) GetRedundantSessionAllowedOk() (*bool, bool)`

GetRedundantSessionAllowedOk returns a tuple with the RedundantSessionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantSessionAllowed

`func (o *DnnConfiguration1) SetRedundantSessionAllowed(v bool)`

SetRedundantSessionAllowed sets RedundantSessionAllowed field to given value.

### HasRedundantSessionAllowed

`func (o *DnnConfiguration1) HasRedundantSessionAllowed() bool`

HasRedundantSessionAllowed returns a boolean if a field has been set.

### GetAcsInfo

`func (o *DnnConfiguration1) GetAcsInfo() AcsInfo1`

GetAcsInfo returns the AcsInfo field if non-nil, zero value otherwise.

### GetAcsInfoOk

`func (o *DnnConfiguration1) GetAcsInfoOk() (*AcsInfo1, bool)`

GetAcsInfoOk returns a tuple with the AcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfo

`func (o *DnnConfiguration1) SetAcsInfo(v AcsInfo1)`

SetAcsInfo sets AcsInfo field to given value.

### HasAcsInfo

`func (o *DnnConfiguration1) HasAcsInfo() bool`

HasAcsInfo returns a boolean if a field has been set.

### GetIpv4FrameRouteList

`func (o *DnnConfiguration1) GetIpv4FrameRouteList() []FrameRouteInfo1`

GetIpv4FrameRouteList returns the Ipv4FrameRouteList field if non-nil, zero value otherwise.

### GetIpv4FrameRouteListOk

`func (o *DnnConfiguration1) GetIpv4FrameRouteListOk() (*[]FrameRouteInfo1, bool)`

GetIpv4FrameRouteListOk returns a tuple with the Ipv4FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4FrameRouteList

`func (o *DnnConfiguration1) SetIpv4FrameRouteList(v []FrameRouteInfo1)`

SetIpv4FrameRouteList sets Ipv4FrameRouteList field to given value.

### HasIpv4FrameRouteList

`func (o *DnnConfiguration1) HasIpv4FrameRouteList() bool`

HasIpv4FrameRouteList returns a boolean if a field has been set.

### GetIpv6FrameRouteList

`func (o *DnnConfiguration1) GetIpv6FrameRouteList() []FrameRouteInfo1`

GetIpv6FrameRouteList returns the Ipv6FrameRouteList field if non-nil, zero value otherwise.

### GetIpv6FrameRouteListOk

`func (o *DnnConfiguration1) GetIpv6FrameRouteListOk() (*[]FrameRouteInfo1, bool)`

GetIpv6FrameRouteListOk returns a tuple with the Ipv6FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6FrameRouteList

`func (o *DnnConfiguration1) SetIpv6FrameRouteList(v []FrameRouteInfo1)`

SetIpv6FrameRouteList sets Ipv6FrameRouteList field to given value.

### HasIpv6FrameRouteList

`func (o *DnnConfiguration1) HasIpv6FrameRouteList() bool`

HasIpv6FrameRouteList returns a boolean if a field has been set.

### GetAtsssAllowed

`func (o *DnnConfiguration1) GetAtsssAllowed() bool`

GetAtsssAllowed returns the AtsssAllowed field if non-nil, zero value otherwise.

### GetAtsssAllowedOk

`func (o *DnnConfiguration1) GetAtsssAllowedOk() (*bool, bool)`

GetAtsssAllowedOk returns a tuple with the AtsssAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssAllowed

`func (o *DnnConfiguration1) SetAtsssAllowed(v bool)`

SetAtsssAllowed sets AtsssAllowed field to given value.

### HasAtsssAllowed

`func (o *DnnConfiguration1) HasAtsssAllowed() bool`

HasAtsssAllowed returns a boolean if a field has been set.

### GetSecondaryAuth

`func (o *DnnConfiguration1) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *DnnConfiguration1) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *DnnConfiguration1) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *DnnConfiguration1) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### GetDnAaaIpAddressAllocation

`func (o *DnnConfiguration1) GetDnAaaIpAddressAllocation() bool`

GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field if non-nil, zero value otherwise.

### GetDnAaaIpAddressAllocationOk

`func (o *DnnConfiguration1) GetDnAaaIpAddressAllocationOk() (*bool, bool)`

GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaIpAddressAllocation

`func (o *DnnConfiguration1) SetDnAaaIpAddressAllocation(v bool)`

SetDnAaaIpAddressAllocation sets DnAaaIpAddressAllocation field to given value.

### HasDnAaaIpAddressAllocation

`func (o *DnnConfiguration1) HasDnAaaIpAddressAllocation() bool`

HasDnAaaIpAddressAllocation returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *DnnConfiguration1) GetDnAaaAddress() IpAddress1`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *DnnConfiguration1) GetDnAaaAddressOk() (*IpAddress1, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *DnnConfiguration1) SetDnAaaAddress(v IpAddress1)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *DnnConfiguration1) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### GetIptvAccCtrlInfo

`func (o *DnnConfiguration1) GetIptvAccCtrlInfo() string`

GetIptvAccCtrlInfo returns the IptvAccCtrlInfo field if non-nil, zero value otherwise.

### GetIptvAccCtrlInfoOk

`func (o *DnnConfiguration1) GetIptvAccCtrlInfoOk() (*string, bool)`

GetIptvAccCtrlInfoOk returns a tuple with the IptvAccCtrlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvAccCtrlInfo

`func (o *DnnConfiguration1) SetIptvAccCtrlInfo(v string)`

SetIptvAccCtrlInfo sets IptvAccCtrlInfo field to given value.

### HasIptvAccCtrlInfo

`func (o *DnnConfiguration1) HasIptvAccCtrlInfo() bool`

HasIptvAccCtrlInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


