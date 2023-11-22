# SmContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** |  | 
**Dnn** | **string** |  | 
**SelectedDnn** | Pointer to **string** |  | [optional] 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**HplmnSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**PduSessionType** | [**PduSessionType**](PduSessionType.md) |  | 
**Gpsi** | Pointer to **string** |  | [optional] 
**HSmfUri** | Pointer to **string** |  | [optional] 
**SmfUri** | Pointer to **string** |  | [optional] 
**PduSessionRef** | Pointer to **string** |  | [optional] 
**PcfId** | Pointer to **string** |  | [optional] 
**PcfGroupId** | Pointer to **string** |  | [optional] 
**PcfSetId** | Pointer to **string** |  | [optional] 
**SelMode** | Pointer to [**DnnSelectionMode**](DnnSelectionMode.md) |  | [optional] 
**UdmGroupId** | Pointer to **string** |  | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**SessionAmbr** | [**Ambr**](Ambr.md) |  | 
**QosFlowsList** | [**[]QosFlowSetupItem**](QosFlowSetupItem.md) |  | 
**HSmfInstanceId** | Pointer to **string** |  | [optional] 
**SmfInstanceId** | Pointer to **string** |  | [optional] 
**PduSessionSmfSetId** | Pointer to **string** |  | [optional] 
**PduSessionSmfServiceSetId** | Pointer to **string** |  | [optional] 
**PduSessionSmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**EnablePauseCharging** | Pointer to **bool** |  | [optional] [default to false]
**UeIpv4Address** | Pointer to **string** |  | [optional] 
**UeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**EpsPdnCnxInfo** | Pointer to [**EpsPdnCnxInfo**](EpsPdnCnxInfo.md) |  | [optional] 
**EpsBearerInfo** | Pointer to [**[]EpsBearerInfo**](EpsBearerInfo.md) |  | [optional] 
**MaxIntegrityProtectedDataRate** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**MaxIntegrityProtectedDataRateDl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**AlwaysOnGranted** | Pointer to **bool** |  | [optional] [default to false]
**UpSecurity** | Pointer to [**UpSecurity**](UpSecurity.md) |  | [optional] 
**HSmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**SmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 
**ForwardingInd** | Pointer to **bool** |  | [optional] [default to false]
**PsaTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**ChargingId** | Pointer to **string** |  | [optional] 
**ChargingInfo** | Pointer to [**ChargingInformation**](ChargingInformation.md) |  | [optional] 
**RoamingChargingProfile** | Pointer to [**RoamingChargingProfile**](RoamingChargingProfile.md) |  | [optional] 
**NefExtBufSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**Ipv6Index** | Pointer to **int32** |  | [optional] 
**DnAaaAddress** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**RedundantPduSessionInfo** | Pointer to [**RedundantPduSessionInformation**](RedundantPduSessionInformation.md) |  | [optional] 
**RanTunnelInfo** | Pointer to [**QosFlowTunnel**](QosFlowTunnel.md) |  | [optional] 
**AddRanTunnelInfo** | Pointer to [**[]QosFlowTunnel**](QosFlowTunnel.md) |  | [optional] 
**RedRanTunnelInfo** | Pointer to [**QosFlowTunnel**](QosFlowTunnel.md) |  | [optional] 
**AddRedRanTunnelInfo** | Pointer to [**[]QosFlowTunnel**](QosFlowTunnel.md) |  | [optional] 
**DlsetSupportInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmContext

`func NewSmContext(pduSessionId int32, dnn string, sNssai Snssai, pduSessionType PduSessionType, sessionAmbr Ambr, qosFlowsList []QosFlowSetupItem, ) *SmContext`

NewSmContext instantiates a new SmContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextWithDefaults

`func NewSmContextWithDefaults() *SmContext`

NewSmContextWithDefaults instantiates a new SmContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *SmContext) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContext) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContext) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetDnn

`func (o *SmContext) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmContext) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmContext) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSelectedDnn

`func (o *SmContext) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *SmContext) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *SmContext) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *SmContext) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *SmContext) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SmContext) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SmContext) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetHplmnSnssai

`func (o *SmContext) GetHplmnSnssai() Snssai`

GetHplmnSnssai returns the HplmnSnssai field if non-nil, zero value otherwise.

### GetHplmnSnssaiOk

`func (o *SmContext) GetHplmnSnssaiOk() (*Snssai, bool)`

GetHplmnSnssaiOk returns a tuple with the HplmnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHplmnSnssai

`func (o *SmContext) SetHplmnSnssai(v Snssai)`

SetHplmnSnssai sets HplmnSnssai field to given value.

### HasHplmnSnssai

`func (o *SmContext) HasHplmnSnssai() bool`

HasHplmnSnssai returns a boolean if a field has been set.

### GetPduSessionType

`func (o *SmContext) GetPduSessionType() PduSessionType`

GetPduSessionType returns the PduSessionType field if non-nil, zero value otherwise.

### GetPduSessionTypeOk

`func (o *SmContext) GetPduSessionTypeOk() (*PduSessionType, bool)`

GetPduSessionTypeOk returns a tuple with the PduSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionType

`func (o *SmContext) SetPduSessionType(v PduSessionType)`

SetPduSessionType sets PduSessionType field to given value.


### GetGpsi

`func (o *SmContext) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmContext) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmContext) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmContext) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetHSmfUri

`func (o *SmContext) GetHSmfUri() string`

GetHSmfUri returns the HSmfUri field if non-nil, zero value otherwise.

### GetHSmfUriOk

`func (o *SmContext) GetHSmfUriOk() (*string, bool)`

GetHSmfUriOk returns a tuple with the HSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfUri

`func (o *SmContext) SetHSmfUri(v string)`

SetHSmfUri sets HSmfUri field to given value.

### HasHSmfUri

`func (o *SmContext) HasHSmfUri() bool`

HasHSmfUri returns a boolean if a field has been set.

### GetSmfUri

`func (o *SmContext) GetSmfUri() string`

GetSmfUri returns the SmfUri field if non-nil, zero value otherwise.

### GetSmfUriOk

`func (o *SmContext) GetSmfUriOk() (*string, bool)`

GetSmfUriOk returns a tuple with the SmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUri

`func (o *SmContext) SetSmfUri(v string)`

SetSmfUri sets SmfUri field to given value.

### HasSmfUri

`func (o *SmContext) HasSmfUri() bool`

HasSmfUri returns a boolean if a field has been set.

### GetPduSessionRef

`func (o *SmContext) GetPduSessionRef() string`

GetPduSessionRef returns the PduSessionRef field if non-nil, zero value otherwise.

### GetPduSessionRefOk

`func (o *SmContext) GetPduSessionRefOk() (*string, bool)`

GetPduSessionRefOk returns a tuple with the PduSessionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionRef

`func (o *SmContext) SetPduSessionRef(v string)`

SetPduSessionRef sets PduSessionRef field to given value.

### HasPduSessionRef

`func (o *SmContext) HasPduSessionRef() bool`

HasPduSessionRef returns a boolean if a field has been set.

### GetPcfId

`func (o *SmContext) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *SmContext) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *SmContext) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *SmContext) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfGroupId

`func (o *SmContext) GetPcfGroupId() string`

GetPcfGroupId returns the PcfGroupId field if non-nil, zero value otherwise.

### GetPcfGroupIdOk

`func (o *SmContext) GetPcfGroupIdOk() (*string, bool)`

GetPcfGroupIdOk returns a tuple with the PcfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfGroupId

`func (o *SmContext) SetPcfGroupId(v string)`

SetPcfGroupId sets PcfGroupId field to given value.

### HasPcfGroupId

`func (o *SmContext) HasPcfGroupId() bool`

HasPcfGroupId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *SmContext) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *SmContext) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *SmContext) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *SmContext) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetSelMode

`func (o *SmContext) GetSelMode() DnnSelectionMode`

GetSelMode returns the SelMode field if non-nil, zero value otherwise.

### GetSelModeOk

`func (o *SmContext) GetSelModeOk() (*DnnSelectionMode, bool)`

GetSelModeOk returns a tuple with the SelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelMode

`func (o *SmContext) SetSelMode(v DnnSelectionMode)`

SetSelMode sets SelMode field to given value.

### HasSelMode

`func (o *SmContext) HasSelMode() bool`

HasSelMode returns a boolean if a field has been set.

### GetUdmGroupId

`func (o *SmContext) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *SmContext) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *SmContext) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *SmContext) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *SmContext) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *SmContext) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *SmContext) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *SmContext) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetSessionAmbr

`func (o *SmContext) GetSessionAmbr() Ambr`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *SmContext) GetSessionAmbrOk() (*Ambr, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *SmContext) SetSessionAmbr(v Ambr)`

SetSessionAmbr sets SessionAmbr field to given value.


### GetQosFlowsList

`func (o *SmContext) GetQosFlowsList() []QosFlowSetupItem`

GetQosFlowsList returns the QosFlowsList field if non-nil, zero value otherwise.

### GetQosFlowsListOk

`func (o *SmContext) GetQosFlowsListOk() (*[]QosFlowSetupItem, bool)`

GetQosFlowsListOk returns a tuple with the QosFlowsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsList

`func (o *SmContext) SetQosFlowsList(v []QosFlowSetupItem)`

SetQosFlowsList sets QosFlowsList field to given value.


### GetHSmfInstanceId

`func (o *SmContext) GetHSmfInstanceId() string`

GetHSmfInstanceId returns the HSmfInstanceId field if non-nil, zero value otherwise.

### GetHSmfInstanceIdOk

`func (o *SmContext) GetHSmfInstanceIdOk() (*string, bool)`

GetHSmfInstanceIdOk returns a tuple with the HSmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfInstanceId

`func (o *SmContext) SetHSmfInstanceId(v string)`

SetHSmfInstanceId sets HSmfInstanceId field to given value.

### HasHSmfInstanceId

`func (o *SmContext) HasHSmfInstanceId() bool`

HasHSmfInstanceId returns a boolean if a field has been set.

### GetSmfInstanceId

`func (o *SmContext) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *SmContext) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *SmContext) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.

### HasSmfInstanceId

`func (o *SmContext) HasSmfInstanceId() bool`

HasSmfInstanceId returns a boolean if a field has been set.

### GetPduSessionSmfSetId

`func (o *SmContext) GetPduSessionSmfSetId() string`

GetPduSessionSmfSetId returns the PduSessionSmfSetId field if non-nil, zero value otherwise.

### GetPduSessionSmfSetIdOk

`func (o *SmContext) GetPduSessionSmfSetIdOk() (*string, bool)`

GetPduSessionSmfSetIdOk returns a tuple with the PduSessionSmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionSmfSetId

`func (o *SmContext) SetPduSessionSmfSetId(v string)`

SetPduSessionSmfSetId sets PduSessionSmfSetId field to given value.

### HasPduSessionSmfSetId

`func (o *SmContext) HasPduSessionSmfSetId() bool`

HasPduSessionSmfSetId returns a boolean if a field has been set.

### GetPduSessionSmfServiceSetId

`func (o *SmContext) GetPduSessionSmfServiceSetId() string`

GetPduSessionSmfServiceSetId returns the PduSessionSmfServiceSetId field if non-nil, zero value otherwise.

### GetPduSessionSmfServiceSetIdOk

`func (o *SmContext) GetPduSessionSmfServiceSetIdOk() (*string, bool)`

GetPduSessionSmfServiceSetIdOk returns a tuple with the PduSessionSmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionSmfServiceSetId

`func (o *SmContext) SetPduSessionSmfServiceSetId(v string)`

SetPduSessionSmfServiceSetId sets PduSessionSmfServiceSetId field to given value.

### HasPduSessionSmfServiceSetId

`func (o *SmContext) HasPduSessionSmfServiceSetId() bool`

HasPduSessionSmfServiceSetId returns a boolean if a field has been set.

### GetPduSessionSmfBinding

`func (o *SmContext) GetPduSessionSmfBinding() SbiBindingLevel`

GetPduSessionSmfBinding returns the PduSessionSmfBinding field if non-nil, zero value otherwise.

### GetPduSessionSmfBindingOk

`func (o *SmContext) GetPduSessionSmfBindingOk() (*SbiBindingLevel, bool)`

GetPduSessionSmfBindingOk returns a tuple with the PduSessionSmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionSmfBinding

`func (o *SmContext) SetPduSessionSmfBinding(v SbiBindingLevel)`

SetPduSessionSmfBinding sets PduSessionSmfBinding field to given value.

### HasPduSessionSmfBinding

`func (o *SmContext) HasPduSessionSmfBinding() bool`

HasPduSessionSmfBinding returns a boolean if a field has been set.

### GetEnablePauseCharging

`func (o *SmContext) GetEnablePauseCharging() bool`

GetEnablePauseCharging returns the EnablePauseCharging field if non-nil, zero value otherwise.

### GetEnablePauseChargingOk

`func (o *SmContext) GetEnablePauseChargingOk() (*bool, bool)`

GetEnablePauseChargingOk returns a tuple with the EnablePauseCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePauseCharging

`func (o *SmContext) SetEnablePauseCharging(v bool)`

SetEnablePauseCharging sets EnablePauseCharging field to given value.

### HasEnablePauseCharging

`func (o *SmContext) HasEnablePauseCharging() bool`

HasEnablePauseCharging returns a boolean if a field has been set.

### GetUeIpv4Address

`func (o *SmContext) GetUeIpv4Address() string`

GetUeIpv4Address returns the UeIpv4Address field if non-nil, zero value otherwise.

### GetUeIpv4AddressOk

`func (o *SmContext) GetUeIpv4AddressOk() (*string, bool)`

GetUeIpv4AddressOk returns a tuple with the UeIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Address

`func (o *SmContext) SetUeIpv4Address(v string)`

SetUeIpv4Address sets UeIpv4Address field to given value.

### HasUeIpv4Address

`func (o *SmContext) HasUeIpv4Address() bool`

HasUeIpv4Address returns a boolean if a field has been set.

### GetUeIpv6Prefix

`func (o *SmContext) GetUeIpv6Prefix() Ipv6Prefix`

GetUeIpv6Prefix returns the UeIpv6Prefix field if non-nil, zero value otherwise.

### GetUeIpv6PrefixOk

`func (o *SmContext) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Prefix

`func (o *SmContext) SetUeIpv6Prefix(v Ipv6Prefix)`

SetUeIpv6Prefix sets UeIpv6Prefix field to given value.

### HasUeIpv6Prefix

`func (o *SmContext) HasUeIpv6Prefix() bool`

HasUeIpv6Prefix returns a boolean if a field has been set.

### GetEpsPdnCnxInfo

`func (o *SmContext) GetEpsPdnCnxInfo() EpsPdnCnxInfo`

GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field if non-nil, zero value otherwise.

### GetEpsPdnCnxInfoOk

`func (o *SmContext) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool)`

GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsPdnCnxInfo

`func (o *SmContext) SetEpsPdnCnxInfo(v EpsPdnCnxInfo)`

SetEpsPdnCnxInfo sets EpsPdnCnxInfo field to given value.

### HasEpsPdnCnxInfo

`func (o *SmContext) HasEpsPdnCnxInfo() bool`

HasEpsPdnCnxInfo returns a boolean if a field has been set.

### GetEpsBearerInfo

`func (o *SmContext) GetEpsBearerInfo() []EpsBearerInfo`

GetEpsBearerInfo returns the EpsBearerInfo field if non-nil, zero value otherwise.

### GetEpsBearerInfoOk

`func (o *SmContext) GetEpsBearerInfoOk() (*[]EpsBearerInfo, bool)`

GetEpsBearerInfoOk returns a tuple with the EpsBearerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerInfo

`func (o *SmContext) SetEpsBearerInfo(v []EpsBearerInfo)`

SetEpsBearerInfo sets EpsBearerInfo field to given value.

### HasEpsBearerInfo

`func (o *SmContext) HasEpsBearerInfo() bool`

HasEpsBearerInfo returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRate

`func (o *SmContext) GetMaxIntegrityProtectedDataRate() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRate returns the MaxIntegrityProtectedDataRate field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateOk

`func (o *SmContext) GetMaxIntegrityProtectedDataRateOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateOk returns a tuple with the MaxIntegrityProtectedDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRate

`func (o *SmContext) SetMaxIntegrityProtectedDataRate(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRate sets MaxIntegrityProtectedDataRate field to given value.

### HasMaxIntegrityProtectedDataRate

`func (o *SmContext) HasMaxIntegrityProtectedDataRate() bool`

HasMaxIntegrityProtectedDataRate returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateDl

`func (o *SmContext) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateDlOk

`func (o *SmContext) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateDl

`func (o *SmContext) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateDl sets MaxIntegrityProtectedDataRateDl field to given value.

### HasMaxIntegrityProtectedDataRateDl

`func (o *SmContext) HasMaxIntegrityProtectedDataRateDl() bool`

HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.

### GetAlwaysOnGranted

`func (o *SmContext) GetAlwaysOnGranted() bool`

GetAlwaysOnGranted returns the AlwaysOnGranted field if non-nil, zero value otherwise.

### GetAlwaysOnGrantedOk

`func (o *SmContext) GetAlwaysOnGrantedOk() (*bool, bool)`

GetAlwaysOnGrantedOk returns a tuple with the AlwaysOnGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnGranted

`func (o *SmContext) SetAlwaysOnGranted(v bool)`

SetAlwaysOnGranted sets AlwaysOnGranted field to given value.

### HasAlwaysOnGranted

`func (o *SmContext) HasAlwaysOnGranted() bool`

HasAlwaysOnGranted returns a boolean if a field has been set.

### GetUpSecurity

`func (o *SmContext) GetUpSecurity() UpSecurity`

GetUpSecurity returns the UpSecurity field if non-nil, zero value otherwise.

### GetUpSecurityOk

`func (o *SmContext) GetUpSecurityOk() (*UpSecurity, bool)`

GetUpSecurityOk returns a tuple with the UpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurity

`func (o *SmContext) SetUpSecurity(v UpSecurity)`

SetUpSecurity sets UpSecurity field to given value.

### HasUpSecurity

`func (o *SmContext) HasUpSecurity() bool`

HasUpSecurity returns a boolean if a field has been set.

### GetHSmfServiceInstanceId

`func (o *SmContext) GetHSmfServiceInstanceId() string`

GetHSmfServiceInstanceId returns the HSmfServiceInstanceId field if non-nil, zero value otherwise.

### GetHSmfServiceInstanceIdOk

`func (o *SmContext) GetHSmfServiceInstanceIdOk() (*string, bool)`

GetHSmfServiceInstanceIdOk returns a tuple with the HSmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfServiceInstanceId

`func (o *SmContext) SetHSmfServiceInstanceId(v string)`

SetHSmfServiceInstanceId sets HSmfServiceInstanceId field to given value.

### HasHSmfServiceInstanceId

`func (o *SmContext) HasHSmfServiceInstanceId() bool`

HasHSmfServiceInstanceId returns a boolean if a field has been set.

### GetSmfServiceInstanceId

`func (o *SmContext) GetSmfServiceInstanceId() string`

GetSmfServiceInstanceId returns the SmfServiceInstanceId field if non-nil, zero value otherwise.

### GetSmfServiceInstanceIdOk

`func (o *SmContext) GetSmfServiceInstanceIdOk() (*string, bool)`

GetSmfServiceInstanceIdOk returns a tuple with the SmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServiceInstanceId

`func (o *SmContext) SetSmfServiceInstanceId(v string)`

SetSmfServiceInstanceId sets SmfServiceInstanceId field to given value.

### HasSmfServiceInstanceId

`func (o *SmContext) HasSmfServiceInstanceId() bool`

HasSmfServiceInstanceId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *SmContext) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *SmContext) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *SmContext) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *SmContext) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetForwardingInd

`func (o *SmContext) GetForwardingInd() bool`

GetForwardingInd returns the ForwardingInd field if non-nil, zero value otherwise.

### GetForwardingIndOk

`func (o *SmContext) GetForwardingIndOk() (*bool, bool)`

GetForwardingIndOk returns a tuple with the ForwardingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingInd

`func (o *SmContext) SetForwardingInd(v bool)`

SetForwardingInd sets ForwardingInd field to given value.

### HasForwardingInd

`func (o *SmContext) HasForwardingInd() bool`

HasForwardingInd returns a boolean if a field has been set.

### GetPsaTunnelInfo

`func (o *SmContext) GetPsaTunnelInfo() TunnelInfo`

GetPsaTunnelInfo returns the PsaTunnelInfo field if non-nil, zero value otherwise.

### GetPsaTunnelInfoOk

`func (o *SmContext) GetPsaTunnelInfoOk() (*TunnelInfo, bool)`

GetPsaTunnelInfoOk returns a tuple with the PsaTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsaTunnelInfo

`func (o *SmContext) SetPsaTunnelInfo(v TunnelInfo)`

SetPsaTunnelInfo sets PsaTunnelInfo field to given value.

### HasPsaTunnelInfo

`func (o *SmContext) HasPsaTunnelInfo() bool`

HasPsaTunnelInfo returns a boolean if a field has been set.

### GetChargingId

`func (o *SmContext) GetChargingId() string`

GetChargingId returns the ChargingId field if non-nil, zero value otherwise.

### GetChargingIdOk

`func (o *SmContext) GetChargingIdOk() (*string, bool)`

GetChargingIdOk returns a tuple with the ChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingId

`func (o *SmContext) SetChargingId(v string)`

SetChargingId sets ChargingId field to given value.

### HasChargingId

`func (o *SmContext) HasChargingId() bool`

HasChargingId returns a boolean if a field has been set.

### GetChargingInfo

`func (o *SmContext) GetChargingInfo() ChargingInformation`

GetChargingInfo returns the ChargingInfo field if non-nil, zero value otherwise.

### GetChargingInfoOk

`func (o *SmContext) GetChargingInfoOk() (*ChargingInformation, bool)`

GetChargingInfoOk returns a tuple with the ChargingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingInfo

`func (o *SmContext) SetChargingInfo(v ChargingInformation)`

SetChargingInfo sets ChargingInfo field to given value.

### HasChargingInfo

`func (o *SmContext) HasChargingInfo() bool`

HasChargingInfo returns a boolean if a field has been set.

### GetRoamingChargingProfile

`func (o *SmContext) GetRoamingChargingProfile() RoamingChargingProfile`

GetRoamingChargingProfile returns the RoamingChargingProfile field if non-nil, zero value otherwise.

### GetRoamingChargingProfileOk

`func (o *SmContext) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool)`

GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingChargingProfile

`func (o *SmContext) SetRoamingChargingProfile(v RoamingChargingProfile)`

SetRoamingChargingProfile sets RoamingChargingProfile field to given value.

### HasRoamingChargingProfile

`func (o *SmContext) HasRoamingChargingProfile() bool`

HasRoamingChargingProfile returns a boolean if a field has been set.

### GetNefExtBufSupportInd

`func (o *SmContext) GetNefExtBufSupportInd() bool`

GetNefExtBufSupportInd returns the NefExtBufSupportInd field if non-nil, zero value otherwise.

### GetNefExtBufSupportIndOk

`func (o *SmContext) GetNefExtBufSupportIndOk() (*bool, bool)`

GetNefExtBufSupportIndOk returns a tuple with the NefExtBufSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefExtBufSupportInd

`func (o *SmContext) SetNefExtBufSupportInd(v bool)`

SetNefExtBufSupportInd sets NefExtBufSupportInd field to given value.

### HasNefExtBufSupportInd

`func (o *SmContext) HasNefExtBufSupportInd() bool`

HasNefExtBufSupportInd returns a boolean if a field has been set.

### GetIpv6Index

`func (o *SmContext) GetIpv6Index() int32`

GetIpv6Index returns the Ipv6Index field if non-nil, zero value otherwise.

### GetIpv6IndexOk

`func (o *SmContext) GetIpv6IndexOk() (*int32, bool)`

GetIpv6IndexOk returns a tuple with the Ipv6Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Index

`func (o *SmContext) SetIpv6Index(v int32)`

SetIpv6Index sets Ipv6Index field to given value.

### HasIpv6Index

`func (o *SmContext) HasIpv6Index() bool`

HasIpv6Index returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *SmContext) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *SmContext) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *SmContext) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *SmContext) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### GetRedundantPduSessionInfo

`func (o *SmContext) GetRedundantPduSessionInfo() RedundantPduSessionInformation`

GetRedundantPduSessionInfo returns the RedundantPduSessionInfo field if non-nil, zero value otherwise.

### GetRedundantPduSessionInfoOk

`func (o *SmContext) GetRedundantPduSessionInfoOk() (*RedundantPduSessionInformation, bool)`

GetRedundantPduSessionInfoOk returns a tuple with the RedundantPduSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantPduSessionInfo

`func (o *SmContext) SetRedundantPduSessionInfo(v RedundantPduSessionInformation)`

SetRedundantPduSessionInfo sets RedundantPduSessionInfo field to given value.

### HasRedundantPduSessionInfo

`func (o *SmContext) HasRedundantPduSessionInfo() bool`

HasRedundantPduSessionInfo returns a boolean if a field has been set.

### GetRanTunnelInfo

`func (o *SmContext) GetRanTunnelInfo() QosFlowTunnel`

GetRanTunnelInfo returns the RanTunnelInfo field if non-nil, zero value otherwise.

### GetRanTunnelInfoOk

`func (o *SmContext) GetRanTunnelInfoOk() (*QosFlowTunnel, bool)`

GetRanTunnelInfoOk returns a tuple with the RanTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanTunnelInfo

`func (o *SmContext) SetRanTunnelInfo(v QosFlowTunnel)`

SetRanTunnelInfo sets RanTunnelInfo field to given value.

### HasRanTunnelInfo

`func (o *SmContext) HasRanTunnelInfo() bool`

HasRanTunnelInfo returns a boolean if a field has been set.

### GetAddRanTunnelInfo

`func (o *SmContext) GetAddRanTunnelInfo() []QosFlowTunnel`

GetAddRanTunnelInfo returns the AddRanTunnelInfo field if non-nil, zero value otherwise.

### GetAddRanTunnelInfoOk

`func (o *SmContext) GetAddRanTunnelInfoOk() (*[]QosFlowTunnel, bool)`

GetAddRanTunnelInfoOk returns a tuple with the AddRanTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRanTunnelInfo

`func (o *SmContext) SetAddRanTunnelInfo(v []QosFlowTunnel)`

SetAddRanTunnelInfo sets AddRanTunnelInfo field to given value.

### HasAddRanTunnelInfo

`func (o *SmContext) HasAddRanTunnelInfo() bool`

HasAddRanTunnelInfo returns a boolean if a field has been set.

### GetRedRanTunnelInfo

`func (o *SmContext) GetRedRanTunnelInfo() QosFlowTunnel`

GetRedRanTunnelInfo returns the RedRanTunnelInfo field if non-nil, zero value otherwise.

### GetRedRanTunnelInfoOk

`func (o *SmContext) GetRedRanTunnelInfoOk() (*QosFlowTunnel, bool)`

GetRedRanTunnelInfoOk returns a tuple with the RedRanTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedRanTunnelInfo

`func (o *SmContext) SetRedRanTunnelInfo(v QosFlowTunnel)`

SetRedRanTunnelInfo sets RedRanTunnelInfo field to given value.

### HasRedRanTunnelInfo

`func (o *SmContext) HasRedRanTunnelInfo() bool`

HasRedRanTunnelInfo returns a boolean if a field has been set.

### GetAddRedRanTunnelInfo

`func (o *SmContext) GetAddRedRanTunnelInfo() []QosFlowTunnel`

GetAddRedRanTunnelInfo returns the AddRedRanTunnelInfo field if non-nil, zero value otherwise.

### GetAddRedRanTunnelInfoOk

`func (o *SmContext) GetAddRedRanTunnelInfoOk() (*[]QosFlowTunnel, bool)`

GetAddRedRanTunnelInfoOk returns a tuple with the AddRedRanTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRedRanTunnelInfo

`func (o *SmContext) SetAddRedRanTunnelInfo(v []QosFlowTunnel)`

SetAddRedRanTunnelInfo sets AddRedRanTunnelInfo field to given value.

### HasAddRedRanTunnelInfo

`func (o *SmContext) HasAddRedRanTunnelInfo() bool`

HasAddRedRanTunnelInfo returns a boolean if a field has been set.

### GetDlsetSupportInd

`func (o *SmContext) GetDlsetSupportInd() bool`

GetDlsetSupportInd returns the DlsetSupportInd field if non-nil, zero value otherwise.

### GetDlsetSupportIndOk

`func (o *SmContext) GetDlsetSupportIndOk() (*bool, bool)`

GetDlsetSupportIndOk returns a tuple with the DlsetSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlsetSupportInd

`func (o *SmContext) SetDlsetSupportInd(v bool)`

SetDlsetSupportInd sets DlsetSupportInd field to given value.

### HasDlsetSupportInd

`func (o *SmContext) HasDlsetSupportInd() bool`

HasDlsetSupportInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


