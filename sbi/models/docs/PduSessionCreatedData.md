# PduSessionCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionType** | [**PduSessionType**](PduSessionType.md) |  | 
**SscMode** | **string** |  | 
**HcnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**CnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**AdditionalCnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**SessionAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**QosFlowsSetupList** | Pointer to [**[]QosFlowSetupItem**](QosFlowSetupItem.md) |  | [optional] 
**HSmfInstanceId** | Pointer to **string** |  | [optional] 
**SmfInstanceId** | Pointer to **string** |  | [optional] 
**PduSessionId** | Pointer to **int32** |  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**EnablePauseCharging** | Pointer to **bool** |  | [optional] [default to false]
**UeIpv4Address** | Pointer to **string** |  | [optional] 
**UeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**N1SmInfoToUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**EpsPdnCnxInfo** | Pointer to [**EpsPdnCnxInfo**](EpsPdnCnxInfo.md) |  | [optional] 
**EpsBearerInfo** | Pointer to [**[]EpsBearerInfo**](EpsBearerInfo.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**MaxIntegrityProtectedDataRate** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**MaxIntegrityProtectedDataRateDl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**AlwaysOnGranted** | Pointer to **bool** |  | [optional] [default to false]
**Gpsi** | Pointer to **string** |  | [optional] 
**UpSecurity** | Pointer to [**UpSecurity**](UpSecurity.md) |  | [optional] 
**RoamingChargingProfile** | Pointer to [**RoamingChargingProfile**](RoamingChargingProfile.md) |  | [optional] 
**HSmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**SmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**Ipv6MultiHomingInd** | Pointer to **bool** |  | [optional] [default to false]
**MaAcceptedInd** | Pointer to **bool** |  | [optional] [default to false]
**HomeProvidedChargingId** | Pointer to **string** |  | [optional] 
**NefExtBufSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**SmallDataRateControlEnabled** | Pointer to **bool** |  | [optional] [default to false]
**UeIpv6InterfaceId** | Pointer to **string** |  | [optional] 
**Ipv6Index** | Pointer to **int32** |  | [optional] 
**DnAaaAddress** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**RedundantPduSessionInfo** | Pointer to [**RedundantPduSessionInformation**](RedundantPduSessionInformation.md) |  | [optional] 

## Methods

### NewPduSessionCreatedData

`func NewPduSessionCreatedData(pduSessionType PduSessionType, sscMode string, ) *PduSessionCreatedData`

NewPduSessionCreatedData instantiates a new PduSessionCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionCreatedDataWithDefaults

`func NewPduSessionCreatedDataWithDefaults() *PduSessionCreatedData`

NewPduSessionCreatedDataWithDefaults instantiates a new PduSessionCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionType

`func (o *PduSessionCreatedData) GetPduSessionType() PduSessionType`

GetPduSessionType returns the PduSessionType field if non-nil, zero value otherwise.

### GetPduSessionTypeOk

`func (o *PduSessionCreatedData) GetPduSessionTypeOk() (*PduSessionType, bool)`

GetPduSessionTypeOk returns a tuple with the PduSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionType

`func (o *PduSessionCreatedData) SetPduSessionType(v PduSessionType)`

SetPduSessionType sets PduSessionType field to given value.


### GetSscMode

`func (o *PduSessionCreatedData) GetSscMode() string`

GetSscMode returns the SscMode field if non-nil, zero value otherwise.

### GetSscModeOk

`func (o *PduSessionCreatedData) GetSscModeOk() (*string, bool)`

GetSscModeOk returns a tuple with the SscMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSscMode

`func (o *PduSessionCreatedData) SetSscMode(v string)`

SetSscMode sets SscMode field to given value.


### GetHcnTunnelInfo

`func (o *PduSessionCreatedData) GetHcnTunnelInfo() TunnelInfo`

GetHcnTunnelInfo returns the HcnTunnelInfo field if non-nil, zero value otherwise.

### GetHcnTunnelInfoOk

`func (o *PduSessionCreatedData) GetHcnTunnelInfoOk() (*TunnelInfo, bool)`

GetHcnTunnelInfoOk returns a tuple with the HcnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcnTunnelInfo

`func (o *PduSessionCreatedData) SetHcnTunnelInfo(v TunnelInfo)`

SetHcnTunnelInfo sets HcnTunnelInfo field to given value.

### HasHcnTunnelInfo

`func (o *PduSessionCreatedData) HasHcnTunnelInfo() bool`

HasHcnTunnelInfo returns a boolean if a field has been set.

### GetCnTunnelInfo

`func (o *PduSessionCreatedData) GetCnTunnelInfo() TunnelInfo`

GetCnTunnelInfo returns the CnTunnelInfo field if non-nil, zero value otherwise.

### GetCnTunnelInfoOk

`func (o *PduSessionCreatedData) GetCnTunnelInfoOk() (*TunnelInfo, bool)`

GetCnTunnelInfoOk returns a tuple with the CnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnTunnelInfo

`func (o *PduSessionCreatedData) SetCnTunnelInfo(v TunnelInfo)`

SetCnTunnelInfo sets CnTunnelInfo field to given value.

### HasCnTunnelInfo

`func (o *PduSessionCreatedData) HasCnTunnelInfo() bool`

HasCnTunnelInfo returns a boolean if a field has been set.

### GetAdditionalCnTunnelInfo

`func (o *PduSessionCreatedData) GetAdditionalCnTunnelInfo() TunnelInfo`

GetAdditionalCnTunnelInfo returns the AdditionalCnTunnelInfo field if non-nil, zero value otherwise.

### GetAdditionalCnTunnelInfoOk

`func (o *PduSessionCreatedData) GetAdditionalCnTunnelInfoOk() (*TunnelInfo, bool)`

GetAdditionalCnTunnelInfoOk returns a tuple with the AdditionalCnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCnTunnelInfo

`func (o *PduSessionCreatedData) SetAdditionalCnTunnelInfo(v TunnelInfo)`

SetAdditionalCnTunnelInfo sets AdditionalCnTunnelInfo field to given value.

### HasAdditionalCnTunnelInfo

`func (o *PduSessionCreatedData) HasAdditionalCnTunnelInfo() bool`

HasAdditionalCnTunnelInfo returns a boolean if a field has been set.

### GetSessionAmbr

`func (o *PduSessionCreatedData) GetSessionAmbr() Ambr`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *PduSessionCreatedData) GetSessionAmbrOk() (*Ambr, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *PduSessionCreatedData) SetSessionAmbr(v Ambr)`

SetSessionAmbr sets SessionAmbr field to given value.

### HasSessionAmbr

`func (o *PduSessionCreatedData) HasSessionAmbr() bool`

HasSessionAmbr returns a boolean if a field has been set.

### GetQosFlowsSetupList

`func (o *PduSessionCreatedData) GetQosFlowsSetupList() []QosFlowSetupItem`

GetQosFlowsSetupList returns the QosFlowsSetupList field if non-nil, zero value otherwise.

### GetQosFlowsSetupListOk

`func (o *PduSessionCreatedData) GetQosFlowsSetupListOk() (*[]QosFlowSetupItem, bool)`

GetQosFlowsSetupListOk returns a tuple with the QosFlowsSetupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsSetupList

`func (o *PduSessionCreatedData) SetQosFlowsSetupList(v []QosFlowSetupItem)`

SetQosFlowsSetupList sets QosFlowsSetupList field to given value.

### HasQosFlowsSetupList

`func (o *PduSessionCreatedData) HasQosFlowsSetupList() bool`

HasQosFlowsSetupList returns a boolean if a field has been set.

### GetHSmfInstanceId

`func (o *PduSessionCreatedData) GetHSmfInstanceId() string`

GetHSmfInstanceId returns the HSmfInstanceId field if non-nil, zero value otherwise.

### GetHSmfInstanceIdOk

`func (o *PduSessionCreatedData) GetHSmfInstanceIdOk() (*string, bool)`

GetHSmfInstanceIdOk returns a tuple with the HSmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfInstanceId

`func (o *PduSessionCreatedData) SetHSmfInstanceId(v string)`

SetHSmfInstanceId sets HSmfInstanceId field to given value.

### HasHSmfInstanceId

`func (o *PduSessionCreatedData) HasHSmfInstanceId() bool`

HasHSmfInstanceId returns a boolean if a field has been set.

### GetSmfInstanceId

`func (o *PduSessionCreatedData) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *PduSessionCreatedData) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *PduSessionCreatedData) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.

### HasSmfInstanceId

`func (o *PduSessionCreatedData) HasSmfInstanceId() bool`

HasSmfInstanceId returns a boolean if a field has been set.

### GetPduSessionId

`func (o *PduSessionCreatedData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduSessionCreatedData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduSessionCreatedData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *PduSessionCreatedData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetSNssai

`func (o *PduSessionCreatedData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *PduSessionCreatedData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *PduSessionCreatedData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *PduSessionCreatedData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetEnablePauseCharging

`func (o *PduSessionCreatedData) GetEnablePauseCharging() bool`

GetEnablePauseCharging returns the EnablePauseCharging field if non-nil, zero value otherwise.

### GetEnablePauseChargingOk

`func (o *PduSessionCreatedData) GetEnablePauseChargingOk() (*bool, bool)`

GetEnablePauseChargingOk returns a tuple with the EnablePauseCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePauseCharging

`func (o *PduSessionCreatedData) SetEnablePauseCharging(v bool)`

SetEnablePauseCharging sets EnablePauseCharging field to given value.

### HasEnablePauseCharging

`func (o *PduSessionCreatedData) HasEnablePauseCharging() bool`

HasEnablePauseCharging returns a boolean if a field has been set.

### GetUeIpv4Address

`func (o *PduSessionCreatedData) GetUeIpv4Address() string`

GetUeIpv4Address returns the UeIpv4Address field if non-nil, zero value otherwise.

### GetUeIpv4AddressOk

`func (o *PduSessionCreatedData) GetUeIpv4AddressOk() (*string, bool)`

GetUeIpv4AddressOk returns a tuple with the UeIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Address

`func (o *PduSessionCreatedData) SetUeIpv4Address(v string)`

SetUeIpv4Address sets UeIpv4Address field to given value.

### HasUeIpv4Address

`func (o *PduSessionCreatedData) HasUeIpv4Address() bool`

HasUeIpv4Address returns a boolean if a field has been set.

### GetUeIpv6Prefix

`func (o *PduSessionCreatedData) GetUeIpv6Prefix() Ipv6Prefix`

GetUeIpv6Prefix returns the UeIpv6Prefix field if non-nil, zero value otherwise.

### GetUeIpv6PrefixOk

`func (o *PduSessionCreatedData) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Prefix

`func (o *PduSessionCreatedData) SetUeIpv6Prefix(v Ipv6Prefix)`

SetUeIpv6Prefix sets UeIpv6Prefix field to given value.

### HasUeIpv6Prefix

`func (o *PduSessionCreatedData) HasUeIpv6Prefix() bool`

HasUeIpv6Prefix returns a boolean if a field has been set.

### GetN1SmInfoToUe

`func (o *PduSessionCreatedData) GetN1SmInfoToUe() RefToBinaryData`

GetN1SmInfoToUe returns the N1SmInfoToUe field if non-nil, zero value otherwise.

### GetN1SmInfoToUeOk

`func (o *PduSessionCreatedData) GetN1SmInfoToUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoToUe

`func (o *PduSessionCreatedData) SetN1SmInfoToUe(v RefToBinaryData)`

SetN1SmInfoToUe sets N1SmInfoToUe field to given value.

### HasN1SmInfoToUe

`func (o *PduSessionCreatedData) HasN1SmInfoToUe() bool`

HasN1SmInfoToUe returns a boolean if a field has been set.

### GetEpsPdnCnxInfo

`func (o *PduSessionCreatedData) GetEpsPdnCnxInfo() EpsPdnCnxInfo`

GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field if non-nil, zero value otherwise.

### GetEpsPdnCnxInfoOk

`func (o *PduSessionCreatedData) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool)`

GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsPdnCnxInfo

`func (o *PduSessionCreatedData) SetEpsPdnCnxInfo(v EpsPdnCnxInfo)`

SetEpsPdnCnxInfo sets EpsPdnCnxInfo field to given value.

### HasEpsPdnCnxInfo

`func (o *PduSessionCreatedData) HasEpsPdnCnxInfo() bool`

HasEpsPdnCnxInfo returns a boolean if a field has been set.

### GetEpsBearerInfo

`func (o *PduSessionCreatedData) GetEpsBearerInfo() []EpsBearerInfo`

GetEpsBearerInfo returns the EpsBearerInfo field if non-nil, zero value otherwise.

### GetEpsBearerInfoOk

`func (o *PduSessionCreatedData) GetEpsBearerInfoOk() (*[]EpsBearerInfo, bool)`

GetEpsBearerInfoOk returns a tuple with the EpsBearerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerInfo

`func (o *PduSessionCreatedData) SetEpsBearerInfo(v []EpsBearerInfo)`

SetEpsBearerInfo sets EpsBearerInfo field to given value.

### HasEpsBearerInfo

`func (o *PduSessionCreatedData) HasEpsBearerInfo() bool`

HasEpsBearerInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PduSessionCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PduSessionCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PduSessionCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PduSessionCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRate

`func (o *PduSessionCreatedData) GetMaxIntegrityProtectedDataRate() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRate returns the MaxIntegrityProtectedDataRate field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateOk

`func (o *PduSessionCreatedData) GetMaxIntegrityProtectedDataRateOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateOk returns a tuple with the MaxIntegrityProtectedDataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRate

`func (o *PduSessionCreatedData) SetMaxIntegrityProtectedDataRate(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRate sets MaxIntegrityProtectedDataRate field to given value.

### HasMaxIntegrityProtectedDataRate

`func (o *PduSessionCreatedData) HasMaxIntegrityProtectedDataRate() bool`

HasMaxIntegrityProtectedDataRate returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreatedData) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateDlOk

`func (o *PduSessionCreatedData) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreatedData) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateDl sets MaxIntegrityProtectedDataRateDl field to given value.

### HasMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreatedData) HasMaxIntegrityProtectedDataRateDl() bool`

HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.

### GetAlwaysOnGranted

`func (o *PduSessionCreatedData) GetAlwaysOnGranted() bool`

GetAlwaysOnGranted returns the AlwaysOnGranted field if non-nil, zero value otherwise.

### GetAlwaysOnGrantedOk

`func (o *PduSessionCreatedData) GetAlwaysOnGrantedOk() (*bool, bool)`

GetAlwaysOnGrantedOk returns a tuple with the AlwaysOnGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnGranted

`func (o *PduSessionCreatedData) SetAlwaysOnGranted(v bool)`

SetAlwaysOnGranted sets AlwaysOnGranted field to given value.

### HasAlwaysOnGranted

`func (o *PduSessionCreatedData) HasAlwaysOnGranted() bool`

HasAlwaysOnGranted returns a boolean if a field has been set.

### GetGpsi

`func (o *PduSessionCreatedData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PduSessionCreatedData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PduSessionCreatedData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PduSessionCreatedData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetUpSecurity

`func (o *PduSessionCreatedData) GetUpSecurity() UpSecurity`

GetUpSecurity returns the UpSecurity field if non-nil, zero value otherwise.

### GetUpSecurityOk

`func (o *PduSessionCreatedData) GetUpSecurityOk() (*UpSecurity, bool)`

GetUpSecurityOk returns a tuple with the UpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurity

`func (o *PduSessionCreatedData) SetUpSecurity(v UpSecurity)`

SetUpSecurity sets UpSecurity field to given value.

### HasUpSecurity

`func (o *PduSessionCreatedData) HasUpSecurity() bool`

HasUpSecurity returns a boolean if a field has been set.

### GetRoamingChargingProfile

`func (o *PduSessionCreatedData) GetRoamingChargingProfile() RoamingChargingProfile`

GetRoamingChargingProfile returns the RoamingChargingProfile field if non-nil, zero value otherwise.

### GetRoamingChargingProfileOk

`func (o *PduSessionCreatedData) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool)`

GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingChargingProfile

`func (o *PduSessionCreatedData) SetRoamingChargingProfile(v RoamingChargingProfile)`

SetRoamingChargingProfile sets RoamingChargingProfile field to given value.

### HasRoamingChargingProfile

`func (o *PduSessionCreatedData) HasRoamingChargingProfile() bool`

HasRoamingChargingProfile returns a boolean if a field has been set.

### GetHSmfServiceInstanceId

`func (o *PduSessionCreatedData) GetHSmfServiceInstanceId() string`

GetHSmfServiceInstanceId returns the HSmfServiceInstanceId field if non-nil, zero value otherwise.

### GetHSmfServiceInstanceIdOk

`func (o *PduSessionCreatedData) GetHSmfServiceInstanceIdOk() (*string, bool)`

GetHSmfServiceInstanceIdOk returns a tuple with the HSmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfServiceInstanceId

`func (o *PduSessionCreatedData) SetHSmfServiceInstanceId(v string)`

SetHSmfServiceInstanceId sets HSmfServiceInstanceId field to given value.

### HasHSmfServiceInstanceId

`func (o *PduSessionCreatedData) HasHSmfServiceInstanceId() bool`

HasHSmfServiceInstanceId returns a boolean if a field has been set.

### GetSmfServiceInstanceId

`func (o *PduSessionCreatedData) GetSmfServiceInstanceId() string`

GetSmfServiceInstanceId returns the SmfServiceInstanceId field if non-nil, zero value otherwise.

### GetSmfServiceInstanceIdOk

`func (o *PduSessionCreatedData) GetSmfServiceInstanceIdOk() (*string, bool)`

GetSmfServiceInstanceIdOk returns a tuple with the SmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServiceInstanceId

`func (o *PduSessionCreatedData) SetSmfServiceInstanceId(v string)`

SetSmfServiceInstanceId sets SmfServiceInstanceId field to given value.

### HasSmfServiceInstanceId

`func (o *PduSessionCreatedData) HasSmfServiceInstanceId() bool`

HasSmfServiceInstanceId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *PduSessionCreatedData) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *PduSessionCreatedData) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *PduSessionCreatedData) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *PduSessionCreatedData) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetDnaiList

`func (o *PduSessionCreatedData) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *PduSessionCreatedData) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *PduSessionCreatedData) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *PduSessionCreatedData) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetIpv6MultiHomingInd

`func (o *PduSessionCreatedData) GetIpv6MultiHomingInd() bool`

GetIpv6MultiHomingInd returns the Ipv6MultiHomingInd field if non-nil, zero value otherwise.

### GetIpv6MultiHomingIndOk

`func (o *PduSessionCreatedData) GetIpv6MultiHomingIndOk() (*bool, bool)`

GetIpv6MultiHomingIndOk returns a tuple with the Ipv6MultiHomingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MultiHomingInd

`func (o *PduSessionCreatedData) SetIpv6MultiHomingInd(v bool)`

SetIpv6MultiHomingInd sets Ipv6MultiHomingInd field to given value.

### HasIpv6MultiHomingInd

`func (o *PduSessionCreatedData) HasIpv6MultiHomingInd() bool`

HasIpv6MultiHomingInd returns a boolean if a field has been set.

### GetMaAcceptedInd

`func (o *PduSessionCreatedData) GetMaAcceptedInd() bool`

GetMaAcceptedInd returns the MaAcceptedInd field if non-nil, zero value otherwise.

### GetMaAcceptedIndOk

`func (o *PduSessionCreatedData) GetMaAcceptedIndOk() (*bool, bool)`

GetMaAcceptedIndOk returns a tuple with the MaAcceptedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaAcceptedInd

`func (o *PduSessionCreatedData) SetMaAcceptedInd(v bool)`

SetMaAcceptedInd sets MaAcceptedInd field to given value.

### HasMaAcceptedInd

`func (o *PduSessionCreatedData) HasMaAcceptedInd() bool`

HasMaAcceptedInd returns a boolean if a field has been set.

### GetHomeProvidedChargingId

`func (o *PduSessionCreatedData) GetHomeProvidedChargingId() string`

GetHomeProvidedChargingId returns the HomeProvidedChargingId field if non-nil, zero value otherwise.

### GetHomeProvidedChargingIdOk

`func (o *PduSessionCreatedData) GetHomeProvidedChargingIdOk() (*string, bool)`

GetHomeProvidedChargingIdOk returns a tuple with the HomeProvidedChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeProvidedChargingId

`func (o *PduSessionCreatedData) SetHomeProvidedChargingId(v string)`

SetHomeProvidedChargingId sets HomeProvidedChargingId field to given value.

### HasHomeProvidedChargingId

`func (o *PduSessionCreatedData) HasHomeProvidedChargingId() bool`

HasHomeProvidedChargingId returns a boolean if a field has been set.

### GetNefExtBufSupportInd

`func (o *PduSessionCreatedData) GetNefExtBufSupportInd() bool`

GetNefExtBufSupportInd returns the NefExtBufSupportInd field if non-nil, zero value otherwise.

### GetNefExtBufSupportIndOk

`func (o *PduSessionCreatedData) GetNefExtBufSupportIndOk() (*bool, bool)`

GetNefExtBufSupportIndOk returns a tuple with the NefExtBufSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefExtBufSupportInd

`func (o *PduSessionCreatedData) SetNefExtBufSupportInd(v bool)`

SetNefExtBufSupportInd sets NefExtBufSupportInd field to given value.

### HasNefExtBufSupportInd

`func (o *PduSessionCreatedData) HasNefExtBufSupportInd() bool`

HasNefExtBufSupportInd returns a boolean if a field has been set.

### GetSmallDataRateControlEnabled

`func (o *PduSessionCreatedData) GetSmallDataRateControlEnabled() bool`

GetSmallDataRateControlEnabled returns the SmallDataRateControlEnabled field if non-nil, zero value otherwise.

### GetSmallDataRateControlEnabledOk

`func (o *PduSessionCreatedData) GetSmallDataRateControlEnabledOk() (*bool, bool)`

GetSmallDataRateControlEnabledOk returns a tuple with the SmallDataRateControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateControlEnabled

`func (o *PduSessionCreatedData) SetSmallDataRateControlEnabled(v bool)`

SetSmallDataRateControlEnabled sets SmallDataRateControlEnabled field to given value.

### HasSmallDataRateControlEnabled

`func (o *PduSessionCreatedData) HasSmallDataRateControlEnabled() bool`

HasSmallDataRateControlEnabled returns a boolean if a field has been set.

### GetUeIpv6InterfaceId

`func (o *PduSessionCreatedData) GetUeIpv6InterfaceId() string`

GetUeIpv6InterfaceId returns the UeIpv6InterfaceId field if non-nil, zero value otherwise.

### GetUeIpv6InterfaceIdOk

`func (o *PduSessionCreatedData) GetUeIpv6InterfaceIdOk() (*string, bool)`

GetUeIpv6InterfaceIdOk returns a tuple with the UeIpv6InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6InterfaceId

`func (o *PduSessionCreatedData) SetUeIpv6InterfaceId(v string)`

SetUeIpv6InterfaceId sets UeIpv6InterfaceId field to given value.

### HasUeIpv6InterfaceId

`func (o *PduSessionCreatedData) HasUeIpv6InterfaceId() bool`

HasUeIpv6InterfaceId returns a boolean if a field has been set.

### GetIpv6Index

`func (o *PduSessionCreatedData) GetIpv6Index() int32`

GetIpv6Index returns the Ipv6Index field if non-nil, zero value otherwise.

### GetIpv6IndexOk

`func (o *PduSessionCreatedData) GetIpv6IndexOk() (*int32, bool)`

GetIpv6IndexOk returns a tuple with the Ipv6Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Index

`func (o *PduSessionCreatedData) SetIpv6Index(v int32)`

SetIpv6Index sets Ipv6Index field to given value.

### HasIpv6Index

`func (o *PduSessionCreatedData) HasIpv6Index() bool`

HasIpv6Index returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *PduSessionCreatedData) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *PduSessionCreatedData) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *PduSessionCreatedData) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *PduSessionCreatedData) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### GetRedundantPduSessionInfo

`func (o *PduSessionCreatedData) GetRedundantPduSessionInfo() RedundantPduSessionInformation`

GetRedundantPduSessionInfo returns the RedundantPduSessionInfo field if non-nil, zero value otherwise.

### GetRedundantPduSessionInfoOk

`func (o *PduSessionCreatedData) GetRedundantPduSessionInfoOk() (*RedundantPduSessionInformation, bool)`

GetRedundantPduSessionInfoOk returns a tuple with the RedundantPduSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantPduSessionInfo

`func (o *PduSessionCreatedData) SetRedundantPduSessionInfo(v RedundantPduSessionInformation)`

SetRedundantPduSessionInfo sets RedundantPduSessionInfo field to given value.

### HasRedundantPduSessionInfo

`func (o *PduSessionCreatedData) HasRedundantPduSessionInfo() bool`

HasRedundantPduSessionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


