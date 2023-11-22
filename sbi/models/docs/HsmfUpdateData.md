# HsmfUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIndication** | [**RequestIndication**](RequestIndication.md) |  | 
**Pei** | Pointer to **string** |  | [optional] 
**VcnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**IcnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**AdditionalCnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PauseCharging** | Pointer to **bool** |  | [optional] 
**Pti** | Pointer to **int32** |  | [optional] 
**N1SmInfoFromUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UnknownN1SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**QosFlowsRelNotifyList** | Pointer to [**[]QosFlowItem**](QosFlowItem.md) |  | [optional] 
**QosFlowsNotifyList** | Pointer to [**[]QosFlowNotifyItem**](QosFlowNotifyItem.md) |  | [optional] 
**NotifyList** | Pointer to [**[]PduSessionNotifyItem**](PduSessionNotifyItem.md) |  | [optional] 
**EpsBearerId** | Pointer to **[]int32** |  | [optional] 
**HoPreparationIndication** | Pointer to **bool** |  | [optional] 
**RevokeEbiList** | Pointer to **[]int32** |  | [optional] 
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**NgApCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**Var5gMmCauseValue** | Pointer to **int32** |  | [optional] 
**AlwaysOnRequested** | Pointer to **bool** |  | [optional] [default to false]
**EpsInterworkingInd** | Pointer to [**EpsInterworkingIndication**](EpsInterworkingIndication.md) |  | [optional] 
**SecondaryRatUsageReport** | Pointer to [**[]SecondaryRatUsageReport**](SecondaryRatUsageReport.md) |  | [optional] 
**SecondaryRatUsageInfo** | Pointer to [**[]SecondaryRatUsageInfo**](SecondaryRatUsageInfo.md) |  | [optional] 
**AnTypeCanBeChanged** | Pointer to **bool** |  | [optional] [default to false]
**MaReleaseInd** | Pointer to [**MaReleaseIndication**](MaReleaseIndication.md) |  | [optional] 
**MaNwUpgradeInd** | Pointer to **bool** |  | [optional] [default to false]
**MaRequestInd** | Pointer to **bool** |  | [optional] [default to false]
**UnavailableAccessInd** | Pointer to [**UnavailableAccessIndication**](UnavailableAccessIndication.md) |  | [optional] 
**PsaInfo** | Pointer to [**[]PsaInformation**](PsaInformation.md) |  | [optional] 
**UlclBpInfo** | Pointer to [**UlclBpInformation**](UlclBpInformation.md) |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**PresenceInLadn** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 
**VsmfPduSessionUri** | Pointer to **string** |  | [optional] 
**VsmfId** | Pointer to **string** |  | [optional] 
**VSmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**IsmfPduSessionUri** | Pointer to **string** |  | [optional] 
**IsmfId** | Pointer to **string** |  | [optional] 
**ISmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**DlServingPlmnRateCtl** | Pointer to **int32** |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**RoamingChargingProfile** | Pointer to [**RoamingChargingProfile**](RoamingChargingProfile.md) |  | [optional] 
**MoExpDataCounter** | Pointer to [**MoExpDataCounter**](MoExpDataCounter.md) |  | [optional] 
**VplmnQos** | Pointer to [**VplmnQos**](VplmnQos.md) |  | [optional] 
**SecurityResult** | Pointer to [**SecurityResult**](SecurityResult.md) |  | [optional] 
**UpSecurityInfo** | Pointer to [**UpSecurityInfo**](UpSecurityInfo.md) |  | [optional] 
**AmfNfId** | Pointer to **string** |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**MaxIntegrityProtectedDataRateUl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**MaxIntegrityProtectedDataRateDl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 

## Methods

### NewHsmfUpdateData

`func NewHsmfUpdateData(requestIndication RequestIndication, ) *HsmfUpdateData`

NewHsmfUpdateData instantiates a new HsmfUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmfUpdateDataWithDefaults

`func NewHsmfUpdateDataWithDefaults() *HsmfUpdateData`

NewHsmfUpdateDataWithDefaults instantiates a new HsmfUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIndication

`func (o *HsmfUpdateData) GetRequestIndication() RequestIndication`

GetRequestIndication returns the RequestIndication field if non-nil, zero value otherwise.

### GetRequestIndicationOk

`func (o *HsmfUpdateData) GetRequestIndicationOk() (*RequestIndication, bool)`

GetRequestIndicationOk returns a tuple with the RequestIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIndication

`func (o *HsmfUpdateData) SetRequestIndication(v RequestIndication)`

SetRequestIndication sets RequestIndication field to given value.


### GetPei

`func (o *HsmfUpdateData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *HsmfUpdateData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *HsmfUpdateData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *HsmfUpdateData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetVcnTunnelInfo

`func (o *HsmfUpdateData) GetVcnTunnelInfo() TunnelInfo`

GetVcnTunnelInfo returns the VcnTunnelInfo field if non-nil, zero value otherwise.

### GetVcnTunnelInfoOk

`func (o *HsmfUpdateData) GetVcnTunnelInfoOk() (*TunnelInfo, bool)`

GetVcnTunnelInfoOk returns a tuple with the VcnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcnTunnelInfo

`func (o *HsmfUpdateData) SetVcnTunnelInfo(v TunnelInfo)`

SetVcnTunnelInfo sets VcnTunnelInfo field to given value.

### HasVcnTunnelInfo

`func (o *HsmfUpdateData) HasVcnTunnelInfo() bool`

HasVcnTunnelInfo returns a boolean if a field has been set.

### GetIcnTunnelInfo

`func (o *HsmfUpdateData) GetIcnTunnelInfo() TunnelInfo`

GetIcnTunnelInfo returns the IcnTunnelInfo field if non-nil, zero value otherwise.

### GetIcnTunnelInfoOk

`func (o *HsmfUpdateData) GetIcnTunnelInfoOk() (*TunnelInfo, bool)`

GetIcnTunnelInfoOk returns a tuple with the IcnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcnTunnelInfo

`func (o *HsmfUpdateData) SetIcnTunnelInfo(v TunnelInfo)`

SetIcnTunnelInfo sets IcnTunnelInfo field to given value.

### HasIcnTunnelInfo

`func (o *HsmfUpdateData) HasIcnTunnelInfo() bool`

HasIcnTunnelInfo returns a boolean if a field has been set.

### GetAdditionalCnTunnelInfo

`func (o *HsmfUpdateData) GetAdditionalCnTunnelInfo() TunnelInfo`

GetAdditionalCnTunnelInfo returns the AdditionalCnTunnelInfo field if non-nil, zero value otherwise.

### GetAdditionalCnTunnelInfoOk

`func (o *HsmfUpdateData) GetAdditionalCnTunnelInfoOk() (*TunnelInfo, bool)`

GetAdditionalCnTunnelInfoOk returns a tuple with the AdditionalCnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCnTunnelInfo

`func (o *HsmfUpdateData) SetAdditionalCnTunnelInfo(v TunnelInfo)`

SetAdditionalCnTunnelInfo sets AdditionalCnTunnelInfo field to given value.

### HasAdditionalCnTunnelInfo

`func (o *HsmfUpdateData) HasAdditionalCnTunnelInfo() bool`

HasAdditionalCnTunnelInfo returns a boolean if a field has been set.

### GetServingNetwork

`func (o *HsmfUpdateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *HsmfUpdateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *HsmfUpdateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *HsmfUpdateData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetAnType

`func (o *HsmfUpdateData) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *HsmfUpdateData) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *HsmfUpdateData) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.

### HasAnType

`func (o *HsmfUpdateData) HasAnType() bool`

HasAnType returns a boolean if a field has been set.

### GetAdditionalAnType

`func (o *HsmfUpdateData) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *HsmfUpdateData) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *HsmfUpdateData) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *HsmfUpdateData) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.

### GetRatType

`func (o *HsmfUpdateData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *HsmfUpdateData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *HsmfUpdateData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *HsmfUpdateData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetUeLocation

`func (o *HsmfUpdateData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *HsmfUpdateData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *HsmfUpdateData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *HsmfUpdateData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *HsmfUpdateData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *HsmfUpdateData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *HsmfUpdateData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *HsmfUpdateData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *HsmfUpdateData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *HsmfUpdateData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *HsmfUpdateData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *HsmfUpdateData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetPauseCharging

`func (o *HsmfUpdateData) GetPauseCharging() bool`

GetPauseCharging returns the PauseCharging field if non-nil, zero value otherwise.

### GetPauseChargingOk

`func (o *HsmfUpdateData) GetPauseChargingOk() (*bool, bool)`

GetPauseChargingOk returns a tuple with the PauseCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseCharging

`func (o *HsmfUpdateData) SetPauseCharging(v bool)`

SetPauseCharging sets PauseCharging field to given value.

### HasPauseCharging

`func (o *HsmfUpdateData) HasPauseCharging() bool`

HasPauseCharging returns a boolean if a field has been set.

### GetPti

`func (o *HsmfUpdateData) GetPti() int32`

GetPti returns the Pti field if non-nil, zero value otherwise.

### GetPtiOk

`func (o *HsmfUpdateData) GetPtiOk() (*int32, bool)`

GetPtiOk returns a tuple with the Pti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPti

`func (o *HsmfUpdateData) SetPti(v int32)`

SetPti sets Pti field to given value.

### HasPti

`func (o *HsmfUpdateData) HasPti() bool`

HasPti returns a boolean if a field has been set.

### GetN1SmInfoFromUe

`func (o *HsmfUpdateData) GetN1SmInfoFromUe() RefToBinaryData`

GetN1SmInfoFromUe returns the N1SmInfoFromUe field if non-nil, zero value otherwise.

### GetN1SmInfoFromUeOk

`func (o *HsmfUpdateData) GetN1SmInfoFromUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoFromUeOk returns a tuple with the N1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoFromUe

`func (o *HsmfUpdateData) SetN1SmInfoFromUe(v RefToBinaryData)`

SetN1SmInfoFromUe sets N1SmInfoFromUe field to given value.

### HasN1SmInfoFromUe

`func (o *HsmfUpdateData) HasN1SmInfoFromUe() bool`

HasN1SmInfoFromUe returns a boolean if a field has been set.

### GetUnknownN1SmInfo

`func (o *HsmfUpdateData) GetUnknownN1SmInfo() RefToBinaryData`

GetUnknownN1SmInfo returns the UnknownN1SmInfo field if non-nil, zero value otherwise.

### GetUnknownN1SmInfoOk

`func (o *HsmfUpdateData) GetUnknownN1SmInfoOk() (*RefToBinaryData, bool)`

GetUnknownN1SmInfoOk returns a tuple with the UnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownN1SmInfo

`func (o *HsmfUpdateData) SetUnknownN1SmInfo(v RefToBinaryData)`

SetUnknownN1SmInfo sets UnknownN1SmInfo field to given value.

### HasUnknownN1SmInfo

`func (o *HsmfUpdateData) HasUnknownN1SmInfo() bool`

HasUnknownN1SmInfo returns a boolean if a field has been set.

### GetQosFlowsRelNotifyList

`func (o *HsmfUpdateData) GetQosFlowsRelNotifyList() []QosFlowItem`

GetQosFlowsRelNotifyList returns the QosFlowsRelNotifyList field if non-nil, zero value otherwise.

### GetQosFlowsRelNotifyListOk

`func (o *HsmfUpdateData) GetQosFlowsRelNotifyListOk() (*[]QosFlowItem, bool)`

GetQosFlowsRelNotifyListOk returns a tuple with the QosFlowsRelNotifyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsRelNotifyList

`func (o *HsmfUpdateData) SetQosFlowsRelNotifyList(v []QosFlowItem)`

SetQosFlowsRelNotifyList sets QosFlowsRelNotifyList field to given value.

### HasQosFlowsRelNotifyList

`func (o *HsmfUpdateData) HasQosFlowsRelNotifyList() bool`

HasQosFlowsRelNotifyList returns a boolean if a field has been set.

### GetQosFlowsNotifyList

`func (o *HsmfUpdateData) GetQosFlowsNotifyList() []QosFlowNotifyItem`

GetQosFlowsNotifyList returns the QosFlowsNotifyList field if non-nil, zero value otherwise.

### GetQosFlowsNotifyListOk

`func (o *HsmfUpdateData) GetQosFlowsNotifyListOk() (*[]QosFlowNotifyItem, bool)`

GetQosFlowsNotifyListOk returns a tuple with the QosFlowsNotifyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsNotifyList

`func (o *HsmfUpdateData) SetQosFlowsNotifyList(v []QosFlowNotifyItem)`

SetQosFlowsNotifyList sets QosFlowsNotifyList field to given value.

### HasQosFlowsNotifyList

`func (o *HsmfUpdateData) HasQosFlowsNotifyList() bool`

HasQosFlowsNotifyList returns a boolean if a field has been set.

### GetNotifyList

`func (o *HsmfUpdateData) GetNotifyList() []PduSessionNotifyItem`

GetNotifyList returns the NotifyList field if non-nil, zero value otherwise.

### GetNotifyListOk

`func (o *HsmfUpdateData) GetNotifyListOk() (*[]PduSessionNotifyItem, bool)`

GetNotifyListOk returns a tuple with the NotifyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyList

`func (o *HsmfUpdateData) SetNotifyList(v []PduSessionNotifyItem)`

SetNotifyList sets NotifyList field to given value.

### HasNotifyList

`func (o *HsmfUpdateData) HasNotifyList() bool`

HasNotifyList returns a boolean if a field has been set.

### GetEpsBearerId

`func (o *HsmfUpdateData) GetEpsBearerId() []int32`

GetEpsBearerId returns the EpsBearerId field if non-nil, zero value otherwise.

### GetEpsBearerIdOk

`func (o *HsmfUpdateData) GetEpsBearerIdOk() (*[]int32, bool)`

GetEpsBearerIdOk returns a tuple with the EpsBearerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerId

`func (o *HsmfUpdateData) SetEpsBearerId(v []int32)`

SetEpsBearerId sets EpsBearerId field to given value.

### HasEpsBearerId

`func (o *HsmfUpdateData) HasEpsBearerId() bool`

HasEpsBearerId returns a boolean if a field has been set.

### GetHoPreparationIndication

`func (o *HsmfUpdateData) GetHoPreparationIndication() bool`

GetHoPreparationIndication returns the HoPreparationIndication field if non-nil, zero value otherwise.

### GetHoPreparationIndicationOk

`func (o *HsmfUpdateData) GetHoPreparationIndicationOk() (*bool, bool)`

GetHoPreparationIndicationOk returns a tuple with the HoPreparationIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoPreparationIndication

`func (o *HsmfUpdateData) SetHoPreparationIndication(v bool)`

SetHoPreparationIndication sets HoPreparationIndication field to given value.

### HasHoPreparationIndication

`func (o *HsmfUpdateData) HasHoPreparationIndication() bool`

HasHoPreparationIndication returns a boolean if a field has been set.

### GetRevokeEbiList

`func (o *HsmfUpdateData) GetRevokeEbiList() []int32`

GetRevokeEbiList returns the RevokeEbiList field if non-nil, zero value otherwise.

### GetRevokeEbiListOk

`func (o *HsmfUpdateData) GetRevokeEbiListOk() (*[]int32, bool)`

GetRevokeEbiListOk returns a tuple with the RevokeEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeEbiList

`func (o *HsmfUpdateData) SetRevokeEbiList(v []int32)`

SetRevokeEbiList sets RevokeEbiList field to given value.

### HasRevokeEbiList

`func (o *HsmfUpdateData) HasRevokeEbiList() bool`

HasRevokeEbiList returns a boolean if a field has been set.

### GetCause

`func (o *HsmfUpdateData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HsmfUpdateData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HsmfUpdateData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HsmfUpdateData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetNgApCause

`func (o *HsmfUpdateData) GetNgApCause() NgApCause`

GetNgApCause returns the NgApCause field if non-nil, zero value otherwise.

### GetNgApCauseOk

`func (o *HsmfUpdateData) GetNgApCauseOk() (*NgApCause, bool)`

GetNgApCauseOk returns a tuple with the NgApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgApCause

`func (o *HsmfUpdateData) SetNgApCause(v NgApCause)`

SetNgApCause sets NgApCause field to given value.

### HasNgApCause

`func (o *HsmfUpdateData) HasNgApCause() bool`

HasNgApCause returns a boolean if a field has been set.

### GetVar5gMmCauseValue

`func (o *HsmfUpdateData) GetVar5gMmCauseValue() int32`

GetVar5gMmCauseValue returns the Var5gMmCauseValue field if non-nil, zero value otherwise.

### GetVar5gMmCauseValueOk

`func (o *HsmfUpdateData) GetVar5gMmCauseValueOk() (*int32, bool)`

GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCauseValue

`func (o *HsmfUpdateData) SetVar5gMmCauseValue(v int32)`

SetVar5gMmCauseValue sets Var5gMmCauseValue field to given value.

### HasVar5gMmCauseValue

`func (o *HsmfUpdateData) HasVar5gMmCauseValue() bool`

HasVar5gMmCauseValue returns a boolean if a field has been set.

### GetAlwaysOnRequested

`func (o *HsmfUpdateData) GetAlwaysOnRequested() bool`

GetAlwaysOnRequested returns the AlwaysOnRequested field if non-nil, zero value otherwise.

### GetAlwaysOnRequestedOk

`func (o *HsmfUpdateData) GetAlwaysOnRequestedOk() (*bool, bool)`

GetAlwaysOnRequestedOk returns a tuple with the AlwaysOnRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnRequested

`func (o *HsmfUpdateData) SetAlwaysOnRequested(v bool)`

SetAlwaysOnRequested sets AlwaysOnRequested field to given value.

### HasAlwaysOnRequested

`func (o *HsmfUpdateData) HasAlwaysOnRequested() bool`

HasAlwaysOnRequested returns a boolean if a field has been set.

### GetEpsInterworkingInd

`func (o *HsmfUpdateData) GetEpsInterworkingInd() EpsInterworkingIndication`

GetEpsInterworkingInd returns the EpsInterworkingInd field if non-nil, zero value otherwise.

### GetEpsInterworkingIndOk

`func (o *HsmfUpdateData) GetEpsInterworkingIndOk() (*EpsInterworkingIndication, bool)`

GetEpsInterworkingIndOk returns a tuple with the EpsInterworkingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInd

`func (o *HsmfUpdateData) SetEpsInterworkingInd(v EpsInterworkingIndication)`

SetEpsInterworkingInd sets EpsInterworkingInd field to given value.

### HasEpsInterworkingInd

`func (o *HsmfUpdateData) HasEpsInterworkingInd() bool`

HasEpsInterworkingInd returns a boolean if a field has been set.

### GetSecondaryRatUsageReport

`func (o *HsmfUpdateData) GetSecondaryRatUsageReport() []SecondaryRatUsageReport`

GetSecondaryRatUsageReport returns the SecondaryRatUsageReport field if non-nil, zero value otherwise.

### GetSecondaryRatUsageReportOk

`func (o *HsmfUpdateData) GetSecondaryRatUsageReportOk() (*[]SecondaryRatUsageReport, bool)`

GetSecondaryRatUsageReportOk returns a tuple with the SecondaryRatUsageReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageReport

`func (o *HsmfUpdateData) SetSecondaryRatUsageReport(v []SecondaryRatUsageReport)`

SetSecondaryRatUsageReport sets SecondaryRatUsageReport field to given value.

### HasSecondaryRatUsageReport

`func (o *HsmfUpdateData) HasSecondaryRatUsageReport() bool`

HasSecondaryRatUsageReport returns a boolean if a field has been set.

### GetSecondaryRatUsageInfo

`func (o *HsmfUpdateData) GetSecondaryRatUsageInfo() []SecondaryRatUsageInfo`

GetSecondaryRatUsageInfo returns the SecondaryRatUsageInfo field if non-nil, zero value otherwise.

### GetSecondaryRatUsageInfoOk

`func (o *HsmfUpdateData) GetSecondaryRatUsageInfoOk() (*[]SecondaryRatUsageInfo, bool)`

GetSecondaryRatUsageInfoOk returns a tuple with the SecondaryRatUsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageInfo

`func (o *HsmfUpdateData) SetSecondaryRatUsageInfo(v []SecondaryRatUsageInfo)`

SetSecondaryRatUsageInfo sets SecondaryRatUsageInfo field to given value.

### HasSecondaryRatUsageInfo

`func (o *HsmfUpdateData) HasSecondaryRatUsageInfo() bool`

HasSecondaryRatUsageInfo returns a boolean if a field has been set.

### GetAnTypeCanBeChanged

`func (o *HsmfUpdateData) GetAnTypeCanBeChanged() bool`

GetAnTypeCanBeChanged returns the AnTypeCanBeChanged field if non-nil, zero value otherwise.

### GetAnTypeCanBeChangedOk

`func (o *HsmfUpdateData) GetAnTypeCanBeChangedOk() (*bool, bool)`

GetAnTypeCanBeChangedOk returns a tuple with the AnTypeCanBeChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnTypeCanBeChanged

`func (o *HsmfUpdateData) SetAnTypeCanBeChanged(v bool)`

SetAnTypeCanBeChanged sets AnTypeCanBeChanged field to given value.

### HasAnTypeCanBeChanged

`func (o *HsmfUpdateData) HasAnTypeCanBeChanged() bool`

HasAnTypeCanBeChanged returns a boolean if a field has been set.

### GetMaReleaseInd

`func (o *HsmfUpdateData) GetMaReleaseInd() MaReleaseIndication`

GetMaReleaseInd returns the MaReleaseInd field if non-nil, zero value otherwise.

### GetMaReleaseIndOk

`func (o *HsmfUpdateData) GetMaReleaseIndOk() (*MaReleaseIndication, bool)`

GetMaReleaseIndOk returns a tuple with the MaReleaseInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaReleaseInd

`func (o *HsmfUpdateData) SetMaReleaseInd(v MaReleaseIndication)`

SetMaReleaseInd sets MaReleaseInd field to given value.

### HasMaReleaseInd

`func (o *HsmfUpdateData) HasMaReleaseInd() bool`

HasMaReleaseInd returns a boolean if a field has been set.

### GetMaNwUpgradeInd

`func (o *HsmfUpdateData) GetMaNwUpgradeInd() bool`

GetMaNwUpgradeInd returns the MaNwUpgradeInd field if non-nil, zero value otherwise.

### GetMaNwUpgradeIndOk

`func (o *HsmfUpdateData) GetMaNwUpgradeIndOk() (*bool, bool)`

GetMaNwUpgradeIndOk returns a tuple with the MaNwUpgradeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaNwUpgradeInd

`func (o *HsmfUpdateData) SetMaNwUpgradeInd(v bool)`

SetMaNwUpgradeInd sets MaNwUpgradeInd field to given value.

### HasMaNwUpgradeInd

`func (o *HsmfUpdateData) HasMaNwUpgradeInd() bool`

HasMaNwUpgradeInd returns a boolean if a field has been set.

### GetMaRequestInd

`func (o *HsmfUpdateData) GetMaRequestInd() bool`

GetMaRequestInd returns the MaRequestInd field if non-nil, zero value otherwise.

### GetMaRequestIndOk

`func (o *HsmfUpdateData) GetMaRequestIndOk() (*bool, bool)`

GetMaRequestIndOk returns a tuple with the MaRequestInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaRequestInd

`func (o *HsmfUpdateData) SetMaRequestInd(v bool)`

SetMaRequestInd sets MaRequestInd field to given value.

### HasMaRequestInd

`func (o *HsmfUpdateData) HasMaRequestInd() bool`

HasMaRequestInd returns a boolean if a field has been set.

### GetUnavailableAccessInd

`func (o *HsmfUpdateData) GetUnavailableAccessInd() UnavailableAccessIndication`

GetUnavailableAccessInd returns the UnavailableAccessInd field if non-nil, zero value otherwise.

### GetUnavailableAccessIndOk

`func (o *HsmfUpdateData) GetUnavailableAccessIndOk() (*UnavailableAccessIndication, bool)`

GetUnavailableAccessIndOk returns a tuple with the UnavailableAccessInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableAccessInd

`func (o *HsmfUpdateData) SetUnavailableAccessInd(v UnavailableAccessIndication)`

SetUnavailableAccessInd sets UnavailableAccessInd field to given value.

### HasUnavailableAccessInd

`func (o *HsmfUpdateData) HasUnavailableAccessInd() bool`

HasUnavailableAccessInd returns a boolean if a field has been set.

### GetPsaInfo

`func (o *HsmfUpdateData) GetPsaInfo() []PsaInformation`

GetPsaInfo returns the PsaInfo field if non-nil, zero value otherwise.

### GetPsaInfoOk

`func (o *HsmfUpdateData) GetPsaInfoOk() (*[]PsaInformation, bool)`

GetPsaInfoOk returns a tuple with the PsaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsaInfo

`func (o *HsmfUpdateData) SetPsaInfo(v []PsaInformation)`

SetPsaInfo sets PsaInfo field to given value.

### HasPsaInfo

`func (o *HsmfUpdateData) HasPsaInfo() bool`

HasPsaInfo returns a boolean if a field has been set.

### GetUlclBpInfo

`func (o *HsmfUpdateData) GetUlclBpInfo() UlclBpInformation`

GetUlclBpInfo returns the UlclBpInfo field if non-nil, zero value otherwise.

### GetUlclBpInfoOk

`func (o *HsmfUpdateData) GetUlclBpInfoOk() (*UlclBpInformation, bool)`

GetUlclBpInfoOk returns a tuple with the UlclBpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlclBpInfo

`func (o *HsmfUpdateData) SetUlclBpInfo(v UlclBpInformation)`

SetUlclBpInfo sets UlclBpInfo field to given value.

### HasUlclBpInfo

`func (o *HsmfUpdateData) HasUlclBpInfo() bool`

HasUlclBpInfo returns a boolean if a field has been set.

### GetN4Info

`func (o *HsmfUpdateData) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *HsmfUpdateData) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *HsmfUpdateData) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *HsmfUpdateData) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *HsmfUpdateData) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *HsmfUpdateData) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *HsmfUpdateData) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *HsmfUpdateData) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *HsmfUpdateData) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *HsmfUpdateData) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *HsmfUpdateData) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *HsmfUpdateData) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.

### GetPresenceInLadn

`func (o *HsmfUpdateData) GetPresenceInLadn() PresenceState`

GetPresenceInLadn returns the PresenceInLadn field if non-nil, zero value otherwise.

### GetPresenceInLadnOk

`func (o *HsmfUpdateData) GetPresenceInLadnOk() (*PresenceState, bool)`

GetPresenceInLadnOk returns a tuple with the PresenceInLadn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInLadn

`func (o *HsmfUpdateData) SetPresenceInLadn(v PresenceState)`

SetPresenceInLadn sets PresenceInLadn field to given value.

### HasPresenceInLadn

`func (o *HsmfUpdateData) HasPresenceInLadn() bool`

HasPresenceInLadn returns a boolean if a field has been set.

### GetVsmfPduSessionUri

`func (o *HsmfUpdateData) GetVsmfPduSessionUri() string`

GetVsmfPduSessionUri returns the VsmfPduSessionUri field if non-nil, zero value otherwise.

### GetVsmfPduSessionUriOk

`func (o *HsmfUpdateData) GetVsmfPduSessionUriOk() (*string, bool)`

GetVsmfPduSessionUriOk returns a tuple with the VsmfPduSessionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfPduSessionUri

`func (o *HsmfUpdateData) SetVsmfPduSessionUri(v string)`

SetVsmfPduSessionUri sets VsmfPduSessionUri field to given value.

### HasVsmfPduSessionUri

`func (o *HsmfUpdateData) HasVsmfPduSessionUri() bool`

HasVsmfPduSessionUri returns a boolean if a field has been set.

### GetVsmfId

`func (o *HsmfUpdateData) GetVsmfId() string`

GetVsmfId returns the VsmfId field if non-nil, zero value otherwise.

### GetVsmfIdOk

`func (o *HsmfUpdateData) GetVsmfIdOk() (*string, bool)`

GetVsmfIdOk returns a tuple with the VsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfId

`func (o *HsmfUpdateData) SetVsmfId(v string)`

SetVsmfId sets VsmfId field to given value.

### HasVsmfId

`func (o *HsmfUpdateData) HasVsmfId() bool`

HasVsmfId returns a boolean if a field has been set.

### GetVSmfServiceInstanceId

`func (o *HsmfUpdateData) GetVSmfServiceInstanceId() string`

GetVSmfServiceInstanceId returns the VSmfServiceInstanceId field if non-nil, zero value otherwise.

### GetVSmfServiceInstanceIdOk

`func (o *HsmfUpdateData) GetVSmfServiceInstanceIdOk() (*string, bool)`

GetVSmfServiceInstanceIdOk returns a tuple with the VSmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSmfServiceInstanceId

`func (o *HsmfUpdateData) SetVSmfServiceInstanceId(v string)`

SetVSmfServiceInstanceId sets VSmfServiceInstanceId field to given value.

### HasVSmfServiceInstanceId

`func (o *HsmfUpdateData) HasVSmfServiceInstanceId() bool`

HasVSmfServiceInstanceId returns a boolean if a field has been set.

### GetIsmfPduSessionUri

`func (o *HsmfUpdateData) GetIsmfPduSessionUri() string`

GetIsmfPduSessionUri returns the IsmfPduSessionUri field if non-nil, zero value otherwise.

### GetIsmfPduSessionUriOk

`func (o *HsmfUpdateData) GetIsmfPduSessionUriOk() (*string, bool)`

GetIsmfPduSessionUriOk returns a tuple with the IsmfPduSessionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfPduSessionUri

`func (o *HsmfUpdateData) SetIsmfPduSessionUri(v string)`

SetIsmfPduSessionUri sets IsmfPduSessionUri field to given value.

### HasIsmfPduSessionUri

`func (o *HsmfUpdateData) HasIsmfPduSessionUri() bool`

HasIsmfPduSessionUri returns a boolean if a field has been set.

### GetIsmfId

`func (o *HsmfUpdateData) GetIsmfId() string`

GetIsmfId returns the IsmfId field if non-nil, zero value otherwise.

### GetIsmfIdOk

`func (o *HsmfUpdateData) GetIsmfIdOk() (*string, bool)`

GetIsmfIdOk returns a tuple with the IsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfId

`func (o *HsmfUpdateData) SetIsmfId(v string)`

SetIsmfId sets IsmfId field to given value.

### HasIsmfId

`func (o *HsmfUpdateData) HasIsmfId() bool`

HasIsmfId returns a boolean if a field has been set.

### GetISmfServiceInstanceId

`func (o *HsmfUpdateData) GetISmfServiceInstanceId() string`

GetISmfServiceInstanceId returns the ISmfServiceInstanceId field if non-nil, zero value otherwise.

### GetISmfServiceInstanceIdOk

`func (o *HsmfUpdateData) GetISmfServiceInstanceIdOk() (*string, bool)`

GetISmfServiceInstanceIdOk returns a tuple with the ISmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISmfServiceInstanceId

`func (o *HsmfUpdateData) SetISmfServiceInstanceId(v string)`

SetISmfServiceInstanceId sets ISmfServiceInstanceId field to given value.

### HasISmfServiceInstanceId

`func (o *HsmfUpdateData) HasISmfServiceInstanceId() bool`

HasISmfServiceInstanceId returns a boolean if a field has been set.

### GetDlServingPlmnRateCtl

`func (o *HsmfUpdateData) GetDlServingPlmnRateCtl() int32`

GetDlServingPlmnRateCtl returns the DlServingPlmnRateCtl field if non-nil, zero value otherwise.

### GetDlServingPlmnRateCtlOk

`func (o *HsmfUpdateData) GetDlServingPlmnRateCtlOk() (*int32, bool)`

GetDlServingPlmnRateCtlOk returns a tuple with the DlServingPlmnRateCtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlServingPlmnRateCtl

`func (o *HsmfUpdateData) SetDlServingPlmnRateCtl(v int32)`

SetDlServingPlmnRateCtl sets DlServingPlmnRateCtl field to given value.

### HasDlServingPlmnRateCtl

`func (o *HsmfUpdateData) HasDlServingPlmnRateCtl() bool`

HasDlServingPlmnRateCtl returns a boolean if a field has been set.

### SetDlServingPlmnRateCtlNil

`func (o *HsmfUpdateData) SetDlServingPlmnRateCtlNil(b bool)`

 SetDlServingPlmnRateCtlNil sets the value for DlServingPlmnRateCtl to be an explicit nil

### UnsetDlServingPlmnRateCtl
`func (o *HsmfUpdateData) UnsetDlServingPlmnRateCtl()`

UnsetDlServingPlmnRateCtl ensures that no value is present for DlServingPlmnRateCtl, not even an explicit nil
### GetDnaiList

`func (o *HsmfUpdateData) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *HsmfUpdateData) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *HsmfUpdateData) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *HsmfUpdateData) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *HsmfUpdateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *HsmfUpdateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *HsmfUpdateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *HsmfUpdateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetRoamingChargingProfile

`func (o *HsmfUpdateData) GetRoamingChargingProfile() RoamingChargingProfile`

GetRoamingChargingProfile returns the RoamingChargingProfile field if non-nil, zero value otherwise.

### GetRoamingChargingProfileOk

`func (o *HsmfUpdateData) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool)`

GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingChargingProfile

`func (o *HsmfUpdateData) SetRoamingChargingProfile(v RoamingChargingProfile)`

SetRoamingChargingProfile sets RoamingChargingProfile field to given value.

### HasRoamingChargingProfile

`func (o *HsmfUpdateData) HasRoamingChargingProfile() bool`

HasRoamingChargingProfile returns a boolean if a field has been set.

### GetMoExpDataCounter

`func (o *HsmfUpdateData) GetMoExpDataCounter() MoExpDataCounter`

GetMoExpDataCounter returns the MoExpDataCounter field if non-nil, zero value otherwise.

### GetMoExpDataCounterOk

`func (o *HsmfUpdateData) GetMoExpDataCounterOk() (*MoExpDataCounter, bool)`

GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoExpDataCounter

`func (o *HsmfUpdateData) SetMoExpDataCounter(v MoExpDataCounter)`

SetMoExpDataCounter sets MoExpDataCounter field to given value.

### HasMoExpDataCounter

`func (o *HsmfUpdateData) HasMoExpDataCounter() bool`

HasMoExpDataCounter returns a boolean if a field has been set.

### GetVplmnQos

`func (o *HsmfUpdateData) GetVplmnQos() VplmnQos`

GetVplmnQos returns the VplmnQos field if non-nil, zero value otherwise.

### GetVplmnQosOk

`func (o *HsmfUpdateData) GetVplmnQosOk() (*VplmnQos, bool)`

GetVplmnQosOk returns a tuple with the VplmnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVplmnQos

`func (o *HsmfUpdateData) SetVplmnQos(v VplmnQos)`

SetVplmnQos sets VplmnQos field to given value.

### HasVplmnQos

`func (o *HsmfUpdateData) HasVplmnQos() bool`

HasVplmnQos returns a boolean if a field has been set.

### GetSecurityResult

`func (o *HsmfUpdateData) GetSecurityResult() SecurityResult`

GetSecurityResult returns the SecurityResult field if non-nil, zero value otherwise.

### GetSecurityResultOk

`func (o *HsmfUpdateData) GetSecurityResultOk() (*SecurityResult, bool)`

GetSecurityResultOk returns a tuple with the SecurityResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityResult

`func (o *HsmfUpdateData) SetSecurityResult(v SecurityResult)`

SetSecurityResult sets SecurityResult field to given value.

### HasSecurityResult

`func (o *HsmfUpdateData) HasSecurityResult() bool`

HasSecurityResult returns a boolean if a field has been set.

### GetUpSecurityInfo

`func (o *HsmfUpdateData) GetUpSecurityInfo() UpSecurityInfo`

GetUpSecurityInfo returns the UpSecurityInfo field if non-nil, zero value otherwise.

### GetUpSecurityInfoOk

`func (o *HsmfUpdateData) GetUpSecurityInfoOk() (*UpSecurityInfo, bool)`

GetUpSecurityInfoOk returns a tuple with the UpSecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurityInfo

`func (o *HsmfUpdateData) SetUpSecurityInfo(v UpSecurityInfo)`

SetUpSecurityInfo sets UpSecurityInfo field to given value.

### HasUpSecurityInfo

`func (o *HsmfUpdateData) HasUpSecurityInfo() bool`

HasUpSecurityInfo returns a boolean if a field has been set.

### GetAmfNfId

`func (o *HsmfUpdateData) GetAmfNfId() string`

GetAmfNfId returns the AmfNfId field if non-nil, zero value otherwise.

### GetAmfNfIdOk

`func (o *HsmfUpdateData) GetAmfNfIdOk() (*string, bool)`

GetAmfNfIdOk returns a tuple with the AmfNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfNfId

`func (o *HsmfUpdateData) SetAmfNfId(v string)`

SetAmfNfId sets AmfNfId field to given value.

### HasAmfNfId

`func (o *HsmfUpdateData) HasAmfNfId() bool`

HasAmfNfId returns a boolean if a field has been set.

### GetGuami

`func (o *HsmfUpdateData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *HsmfUpdateData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *HsmfUpdateData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *HsmfUpdateData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdateData) GetMaxIntegrityProtectedDataRateUl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateUl returns the MaxIntegrityProtectedDataRateUl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateUlOk

`func (o *HsmfUpdateData) GetMaxIntegrityProtectedDataRateUlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateUlOk returns a tuple with the MaxIntegrityProtectedDataRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdateData) SetMaxIntegrityProtectedDataRateUl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateUl sets MaxIntegrityProtectedDataRateUl field to given value.

### HasMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdateData) HasMaxIntegrityProtectedDataRateUl() bool`

HasMaxIntegrityProtectedDataRateUl returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdateData) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateDlOk

`func (o *HsmfUpdateData) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdateData) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateDl sets MaxIntegrityProtectedDataRateDl field to given value.

### HasMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdateData) HasMaxIntegrityProtectedDataRateDl() bool`

HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.

### GetUpCnxState

`func (o *HsmfUpdateData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *HsmfUpdateData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *HsmfUpdateData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *HsmfUpdateData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


