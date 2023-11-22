# VsmfUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIndication** | [**RequestIndication**](RequestIndication.md) |  | 
**SessionAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**QosFlowsAddModRequestList** | Pointer to [**[]QosFlowAddModifyRequestItem**](QosFlowAddModifyRequestItem.md) |  | [optional] 
**QosFlowsRelRequestList** | Pointer to [**[]QosFlowReleaseRequestItem**](QosFlowReleaseRequestItem.md) |  | [optional] 
**EpsBearerInfo** | Pointer to [**[]EpsBearerInfo**](EpsBearerInfo.md) |  | [optional] 
**AssignEbiList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 
**RevokeEbiList** | Pointer to **[]int32** |  | [optional] 
**ModifiedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**Pti** | Pointer to **int32** |  | [optional] 
**N1SmInfoToUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**AlwaysOnGranted** | Pointer to **bool** |  | [optional] [default to false]
**HsmfPduSessionUri** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**N1smCause** | Pointer to **string** |  | [optional] 
**BackOffTimer** | Pointer to **int32** |  | [optional] 
**MaReleaseInd** | Pointer to [**MaReleaseIndication**](MaReleaseIndication.md) |  | [optional] 
**MaAcceptedInd** | Pointer to **bool** |  | [optional] [default to false]
**AdditionalCnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**SmallDataRateControlEnabled** | Pointer to **bool** |  | [optional] 
**QosMonitoringInfo** | Pointer to [**QosMonitoringInfo**](QosMonitoringInfo.md) |  | [optional] 
**EpsPdnCnxInfo** | Pointer to [**EpsPdnCnxInfo**](EpsPdnCnxInfo.md) |  | [optional] 

## Methods

### NewVsmfUpdateData

`func NewVsmfUpdateData(requestIndication RequestIndication, ) *VsmfUpdateData`

NewVsmfUpdateData instantiates a new VsmfUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVsmfUpdateDataWithDefaults

`func NewVsmfUpdateDataWithDefaults() *VsmfUpdateData`

NewVsmfUpdateDataWithDefaults instantiates a new VsmfUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIndication

`func (o *VsmfUpdateData) GetRequestIndication() RequestIndication`

GetRequestIndication returns the RequestIndication field if non-nil, zero value otherwise.

### GetRequestIndicationOk

`func (o *VsmfUpdateData) GetRequestIndicationOk() (*RequestIndication, bool)`

GetRequestIndicationOk returns a tuple with the RequestIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIndication

`func (o *VsmfUpdateData) SetRequestIndication(v RequestIndication)`

SetRequestIndication sets RequestIndication field to given value.


### GetSessionAmbr

`func (o *VsmfUpdateData) GetSessionAmbr() Ambr`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *VsmfUpdateData) GetSessionAmbrOk() (*Ambr, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *VsmfUpdateData) SetSessionAmbr(v Ambr)`

SetSessionAmbr sets SessionAmbr field to given value.

### HasSessionAmbr

`func (o *VsmfUpdateData) HasSessionAmbr() bool`

HasSessionAmbr returns a boolean if a field has been set.

### GetQosFlowsAddModRequestList

`func (o *VsmfUpdateData) GetQosFlowsAddModRequestList() []QosFlowAddModifyRequestItem`

GetQosFlowsAddModRequestList returns the QosFlowsAddModRequestList field if non-nil, zero value otherwise.

### GetQosFlowsAddModRequestListOk

`func (o *VsmfUpdateData) GetQosFlowsAddModRequestListOk() (*[]QosFlowAddModifyRequestItem, bool)`

GetQosFlowsAddModRequestListOk returns a tuple with the QosFlowsAddModRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsAddModRequestList

`func (o *VsmfUpdateData) SetQosFlowsAddModRequestList(v []QosFlowAddModifyRequestItem)`

SetQosFlowsAddModRequestList sets QosFlowsAddModRequestList field to given value.

### HasQosFlowsAddModRequestList

`func (o *VsmfUpdateData) HasQosFlowsAddModRequestList() bool`

HasQosFlowsAddModRequestList returns a boolean if a field has been set.

### GetQosFlowsRelRequestList

`func (o *VsmfUpdateData) GetQosFlowsRelRequestList() []QosFlowReleaseRequestItem`

GetQosFlowsRelRequestList returns the QosFlowsRelRequestList field if non-nil, zero value otherwise.

### GetQosFlowsRelRequestListOk

`func (o *VsmfUpdateData) GetQosFlowsRelRequestListOk() (*[]QosFlowReleaseRequestItem, bool)`

GetQosFlowsRelRequestListOk returns a tuple with the QosFlowsRelRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsRelRequestList

`func (o *VsmfUpdateData) SetQosFlowsRelRequestList(v []QosFlowReleaseRequestItem)`

SetQosFlowsRelRequestList sets QosFlowsRelRequestList field to given value.

### HasQosFlowsRelRequestList

`func (o *VsmfUpdateData) HasQosFlowsRelRequestList() bool`

HasQosFlowsRelRequestList returns a boolean if a field has been set.

### GetEpsBearerInfo

`func (o *VsmfUpdateData) GetEpsBearerInfo() []EpsBearerInfo`

GetEpsBearerInfo returns the EpsBearerInfo field if non-nil, zero value otherwise.

### GetEpsBearerInfoOk

`func (o *VsmfUpdateData) GetEpsBearerInfoOk() (*[]EpsBearerInfo, bool)`

GetEpsBearerInfoOk returns a tuple with the EpsBearerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerInfo

`func (o *VsmfUpdateData) SetEpsBearerInfo(v []EpsBearerInfo)`

SetEpsBearerInfo sets EpsBearerInfo field to given value.

### HasEpsBearerInfo

`func (o *VsmfUpdateData) HasEpsBearerInfo() bool`

HasEpsBearerInfo returns a boolean if a field has been set.

### GetAssignEbiList

`func (o *VsmfUpdateData) GetAssignEbiList() []Arp`

GetAssignEbiList returns the AssignEbiList field if non-nil, zero value otherwise.

### GetAssignEbiListOk

`func (o *VsmfUpdateData) GetAssignEbiListOk() (*[]Arp, bool)`

GetAssignEbiListOk returns a tuple with the AssignEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignEbiList

`func (o *VsmfUpdateData) SetAssignEbiList(v []Arp)`

SetAssignEbiList sets AssignEbiList field to given value.

### HasAssignEbiList

`func (o *VsmfUpdateData) HasAssignEbiList() bool`

HasAssignEbiList returns a boolean if a field has been set.

### GetRevokeEbiList

`func (o *VsmfUpdateData) GetRevokeEbiList() []int32`

GetRevokeEbiList returns the RevokeEbiList field if non-nil, zero value otherwise.

### GetRevokeEbiListOk

`func (o *VsmfUpdateData) GetRevokeEbiListOk() (*[]int32, bool)`

GetRevokeEbiListOk returns a tuple with the RevokeEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeEbiList

`func (o *VsmfUpdateData) SetRevokeEbiList(v []int32)`

SetRevokeEbiList sets RevokeEbiList field to given value.

### HasRevokeEbiList

`func (o *VsmfUpdateData) HasRevokeEbiList() bool`

HasRevokeEbiList returns a boolean if a field has been set.

### GetModifiedEbiList

`func (o *VsmfUpdateData) GetModifiedEbiList() []EbiArpMapping`

GetModifiedEbiList returns the ModifiedEbiList field if non-nil, zero value otherwise.

### GetModifiedEbiListOk

`func (o *VsmfUpdateData) GetModifiedEbiListOk() (*[]EbiArpMapping, bool)`

GetModifiedEbiListOk returns a tuple with the ModifiedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedEbiList

`func (o *VsmfUpdateData) SetModifiedEbiList(v []EbiArpMapping)`

SetModifiedEbiList sets ModifiedEbiList field to given value.

### HasModifiedEbiList

`func (o *VsmfUpdateData) HasModifiedEbiList() bool`

HasModifiedEbiList returns a boolean if a field has been set.

### GetPti

`func (o *VsmfUpdateData) GetPti() int32`

GetPti returns the Pti field if non-nil, zero value otherwise.

### GetPtiOk

`func (o *VsmfUpdateData) GetPtiOk() (*int32, bool)`

GetPtiOk returns a tuple with the Pti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPti

`func (o *VsmfUpdateData) SetPti(v int32)`

SetPti sets Pti field to given value.

### HasPti

`func (o *VsmfUpdateData) HasPti() bool`

HasPti returns a boolean if a field has been set.

### GetN1SmInfoToUe

`func (o *VsmfUpdateData) GetN1SmInfoToUe() RefToBinaryData`

GetN1SmInfoToUe returns the N1SmInfoToUe field if non-nil, zero value otherwise.

### GetN1SmInfoToUeOk

`func (o *VsmfUpdateData) GetN1SmInfoToUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoToUe

`func (o *VsmfUpdateData) SetN1SmInfoToUe(v RefToBinaryData)`

SetN1SmInfoToUe sets N1SmInfoToUe field to given value.

### HasN1SmInfoToUe

`func (o *VsmfUpdateData) HasN1SmInfoToUe() bool`

HasN1SmInfoToUe returns a boolean if a field has been set.

### GetAlwaysOnGranted

`func (o *VsmfUpdateData) GetAlwaysOnGranted() bool`

GetAlwaysOnGranted returns the AlwaysOnGranted field if non-nil, zero value otherwise.

### GetAlwaysOnGrantedOk

`func (o *VsmfUpdateData) GetAlwaysOnGrantedOk() (*bool, bool)`

GetAlwaysOnGrantedOk returns a tuple with the AlwaysOnGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnGranted

`func (o *VsmfUpdateData) SetAlwaysOnGranted(v bool)`

SetAlwaysOnGranted sets AlwaysOnGranted field to given value.

### HasAlwaysOnGranted

`func (o *VsmfUpdateData) HasAlwaysOnGranted() bool`

HasAlwaysOnGranted returns a boolean if a field has been set.

### GetHsmfPduSessionUri

`func (o *VsmfUpdateData) GetHsmfPduSessionUri() string`

GetHsmfPduSessionUri returns the HsmfPduSessionUri field if non-nil, zero value otherwise.

### GetHsmfPduSessionUriOk

`func (o *VsmfUpdateData) GetHsmfPduSessionUriOk() (*string, bool)`

GetHsmfPduSessionUriOk returns a tuple with the HsmfPduSessionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfPduSessionUri

`func (o *VsmfUpdateData) SetHsmfPduSessionUri(v string)`

SetHsmfPduSessionUri sets HsmfPduSessionUri field to given value.

### HasHsmfPduSessionUri

`func (o *VsmfUpdateData) HasHsmfPduSessionUri() bool`

HasHsmfPduSessionUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *VsmfUpdateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *VsmfUpdateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *VsmfUpdateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *VsmfUpdateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetCause

`func (o *VsmfUpdateData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *VsmfUpdateData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *VsmfUpdateData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *VsmfUpdateData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetN1smCause

`func (o *VsmfUpdateData) GetN1smCause() string`

GetN1smCause returns the N1smCause field if non-nil, zero value otherwise.

### GetN1smCauseOk

`func (o *VsmfUpdateData) GetN1smCauseOk() (*string, bool)`

GetN1smCauseOk returns a tuple with the N1smCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1smCause

`func (o *VsmfUpdateData) SetN1smCause(v string)`

SetN1smCause sets N1smCause field to given value.

### HasN1smCause

`func (o *VsmfUpdateData) HasN1smCause() bool`

HasN1smCause returns a boolean if a field has been set.

### GetBackOffTimer

`func (o *VsmfUpdateData) GetBackOffTimer() int32`

GetBackOffTimer returns the BackOffTimer field if non-nil, zero value otherwise.

### GetBackOffTimerOk

`func (o *VsmfUpdateData) GetBackOffTimerOk() (*int32, bool)`

GetBackOffTimerOk returns a tuple with the BackOffTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOffTimer

`func (o *VsmfUpdateData) SetBackOffTimer(v int32)`

SetBackOffTimer sets BackOffTimer field to given value.

### HasBackOffTimer

`func (o *VsmfUpdateData) HasBackOffTimer() bool`

HasBackOffTimer returns a boolean if a field has been set.

### GetMaReleaseInd

`func (o *VsmfUpdateData) GetMaReleaseInd() MaReleaseIndication`

GetMaReleaseInd returns the MaReleaseInd field if non-nil, zero value otherwise.

### GetMaReleaseIndOk

`func (o *VsmfUpdateData) GetMaReleaseIndOk() (*MaReleaseIndication, bool)`

GetMaReleaseIndOk returns a tuple with the MaReleaseInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaReleaseInd

`func (o *VsmfUpdateData) SetMaReleaseInd(v MaReleaseIndication)`

SetMaReleaseInd sets MaReleaseInd field to given value.

### HasMaReleaseInd

`func (o *VsmfUpdateData) HasMaReleaseInd() bool`

HasMaReleaseInd returns a boolean if a field has been set.

### GetMaAcceptedInd

`func (o *VsmfUpdateData) GetMaAcceptedInd() bool`

GetMaAcceptedInd returns the MaAcceptedInd field if non-nil, zero value otherwise.

### GetMaAcceptedIndOk

`func (o *VsmfUpdateData) GetMaAcceptedIndOk() (*bool, bool)`

GetMaAcceptedIndOk returns a tuple with the MaAcceptedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaAcceptedInd

`func (o *VsmfUpdateData) SetMaAcceptedInd(v bool)`

SetMaAcceptedInd sets MaAcceptedInd field to given value.

### HasMaAcceptedInd

`func (o *VsmfUpdateData) HasMaAcceptedInd() bool`

HasMaAcceptedInd returns a boolean if a field has been set.

### GetAdditionalCnTunnelInfo

`func (o *VsmfUpdateData) GetAdditionalCnTunnelInfo() TunnelInfo`

GetAdditionalCnTunnelInfo returns the AdditionalCnTunnelInfo field if non-nil, zero value otherwise.

### GetAdditionalCnTunnelInfoOk

`func (o *VsmfUpdateData) GetAdditionalCnTunnelInfoOk() (*TunnelInfo, bool)`

GetAdditionalCnTunnelInfoOk returns a tuple with the AdditionalCnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCnTunnelInfo

`func (o *VsmfUpdateData) SetAdditionalCnTunnelInfo(v TunnelInfo)`

SetAdditionalCnTunnelInfo sets AdditionalCnTunnelInfo field to given value.

### HasAdditionalCnTunnelInfo

`func (o *VsmfUpdateData) HasAdditionalCnTunnelInfo() bool`

HasAdditionalCnTunnelInfo returns a boolean if a field has been set.

### GetDnaiList

`func (o *VsmfUpdateData) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *VsmfUpdateData) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *VsmfUpdateData) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *VsmfUpdateData) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetN4Info

`func (o *VsmfUpdateData) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *VsmfUpdateData) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *VsmfUpdateData) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *VsmfUpdateData) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *VsmfUpdateData) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *VsmfUpdateData) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *VsmfUpdateData) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *VsmfUpdateData) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *VsmfUpdateData) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *VsmfUpdateData) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *VsmfUpdateData) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *VsmfUpdateData) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.

### GetSmallDataRateControlEnabled

`func (o *VsmfUpdateData) GetSmallDataRateControlEnabled() bool`

GetSmallDataRateControlEnabled returns the SmallDataRateControlEnabled field if non-nil, zero value otherwise.

### GetSmallDataRateControlEnabledOk

`func (o *VsmfUpdateData) GetSmallDataRateControlEnabledOk() (*bool, bool)`

GetSmallDataRateControlEnabledOk returns a tuple with the SmallDataRateControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateControlEnabled

`func (o *VsmfUpdateData) SetSmallDataRateControlEnabled(v bool)`

SetSmallDataRateControlEnabled sets SmallDataRateControlEnabled field to given value.

### HasSmallDataRateControlEnabled

`func (o *VsmfUpdateData) HasSmallDataRateControlEnabled() bool`

HasSmallDataRateControlEnabled returns a boolean if a field has been set.

### GetQosMonitoringInfo

`func (o *VsmfUpdateData) GetQosMonitoringInfo() QosMonitoringInfo`

GetQosMonitoringInfo returns the QosMonitoringInfo field if non-nil, zero value otherwise.

### GetQosMonitoringInfoOk

`func (o *VsmfUpdateData) GetQosMonitoringInfoOk() (*QosMonitoringInfo, bool)`

GetQosMonitoringInfoOk returns a tuple with the QosMonitoringInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonitoringInfo

`func (o *VsmfUpdateData) SetQosMonitoringInfo(v QosMonitoringInfo)`

SetQosMonitoringInfo sets QosMonitoringInfo field to given value.

### HasQosMonitoringInfo

`func (o *VsmfUpdateData) HasQosMonitoringInfo() bool`

HasQosMonitoringInfo returns a boolean if a field has been set.

### GetEpsPdnCnxInfo

`func (o *VsmfUpdateData) GetEpsPdnCnxInfo() EpsPdnCnxInfo`

GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field if non-nil, zero value otherwise.

### GetEpsPdnCnxInfoOk

`func (o *VsmfUpdateData) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool)`

GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsPdnCnxInfo

`func (o *VsmfUpdateData) SetEpsPdnCnxInfo(v EpsPdnCnxInfo)`

SetEpsPdnCnxInfo sets EpsPdnCnxInfo field to given value.

### HasEpsPdnCnxInfo

`func (o *VsmfUpdateData) HasEpsPdnCnxInfo() bool`

HasEpsPdnCnxInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


