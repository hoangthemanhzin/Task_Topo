# HsmfUpdatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N1SmInfoToUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**RoamingChargingProfile** | Pointer to [**RoamingChargingProfile**](RoamingChargingProfile.md) |  | [optional] 
**UpSecurity** | Pointer to [**UpSecurity**](UpSecurity.md) |  | [optional] 
**MaxIntegrityProtectedDataRateUl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**MaxIntegrityProtectedDataRateDl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**Ipv6MultiHomingInd** | Pointer to **bool** |  | [optional] [default to false]
**QosFlowsSetupList** | Pointer to [**[]QosFlowSetupItem**](QosFlowSetupItem.md) |  | [optional] 
**SessionAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**EpsPdnCnxInfo** | Pointer to [**EpsPdnCnxInfo**](EpsPdnCnxInfo.md) |  | [optional] 
**EpsBearerInfo** | Pointer to [**[]EpsBearerInfo**](EpsBearerInfo.md) |  | [optional] 
**Pti** | Pointer to **int32** |  | [optional] 

## Methods

### NewHsmfUpdatedData

`func NewHsmfUpdatedData() *HsmfUpdatedData`

NewHsmfUpdatedData instantiates a new HsmfUpdatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmfUpdatedDataWithDefaults

`func NewHsmfUpdatedDataWithDefaults() *HsmfUpdatedData`

NewHsmfUpdatedDataWithDefaults instantiates a new HsmfUpdatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN1SmInfoToUe

`func (o *HsmfUpdatedData) GetN1SmInfoToUe() RefToBinaryData`

GetN1SmInfoToUe returns the N1SmInfoToUe field if non-nil, zero value otherwise.

### GetN1SmInfoToUeOk

`func (o *HsmfUpdatedData) GetN1SmInfoToUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoToUe

`func (o *HsmfUpdatedData) SetN1SmInfoToUe(v RefToBinaryData)`

SetN1SmInfoToUe sets N1SmInfoToUe field to given value.

### HasN1SmInfoToUe

`func (o *HsmfUpdatedData) HasN1SmInfoToUe() bool`

HasN1SmInfoToUe returns a boolean if a field has been set.

### GetN4Info

`func (o *HsmfUpdatedData) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *HsmfUpdatedData) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *HsmfUpdatedData) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *HsmfUpdatedData) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *HsmfUpdatedData) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *HsmfUpdatedData) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *HsmfUpdatedData) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *HsmfUpdatedData) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *HsmfUpdatedData) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *HsmfUpdatedData) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *HsmfUpdatedData) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *HsmfUpdatedData) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.

### GetDnaiList

`func (o *HsmfUpdatedData) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *HsmfUpdatedData) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *HsmfUpdatedData) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *HsmfUpdatedData) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *HsmfUpdatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *HsmfUpdatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *HsmfUpdatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *HsmfUpdatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetRoamingChargingProfile

`func (o *HsmfUpdatedData) GetRoamingChargingProfile() RoamingChargingProfile`

GetRoamingChargingProfile returns the RoamingChargingProfile field if non-nil, zero value otherwise.

### GetRoamingChargingProfileOk

`func (o *HsmfUpdatedData) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool)`

GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingChargingProfile

`func (o *HsmfUpdatedData) SetRoamingChargingProfile(v RoamingChargingProfile)`

SetRoamingChargingProfile sets RoamingChargingProfile field to given value.

### HasRoamingChargingProfile

`func (o *HsmfUpdatedData) HasRoamingChargingProfile() bool`

HasRoamingChargingProfile returns a boolean if a field has been set.

### GetUpSecurity

`func (o *HsmfUpdatedData) GetUpSecurity() UpSecurity`

GetUpSecurity returns the UpSecurity field if non-nil, zero value otherwise.

### GetUpSecurityOk

`func (o *HsmfUpdatedData) GetUpSecurityOk() (*UpSecurity, bool)`

GetUpSecurityOk returns a tuple with the UpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurity

`func (o *HsmfUpdatedData) SetUpSecurity(v UpSecurity)`

SetUpSecurity sets UpSecurity field to given value.

### HasUpSecurity

`func (o *HsmfUpdatedData) HasUpSecurity() bool`

HasUpSecurity returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateUl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateUl returns the MaxIntegrityProtectedDataRateUl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateUlOk

`func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateUlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateUlOk returns a tuple with the MaxIntegrityProtectedDataRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdatedData) SetMaxIntegrityProtectedDataRateUl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateUl sets MaxIntegrityProtectedDataRateUl field to given value.

### HasMaxIntegrityProtectedDataRateUl

`func (o *HsmfUpdatedData) HasMaxIntegrityProtectedDataRateUl() bool`

HasMaxIntegrityProtectedDataRateUl returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateDlOk

`func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdatedData) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateDl sets MaxIntegrityProtectedDataRateDl field to given value.

### HasMaxIntegrityProtectedDataRateDl

`func (o *HsmfUpdatedData) HasMaxIntegrityProtectedDataRateDl() bool`

HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.

### GetIpv6MultiHomingInd

`func (o *HsmfUpdatedData) GetIpv6MultiHomingInd() bool`

GetIpv6MultiHomingInd returns the Ipv6MultiHomingInd field if non-nil, zero value otherwise.

### GetIpv6MultiHomingIndOk

`func (o *HsmfUpdatedData) GetIpv6MultiHomingIndOk() (*bool, bool)`

GetIpv6MultiHomingIndOk returns a tuple with the Ipv6MultiHomingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MultiHomingInd

`func (o *HsmfUpdatedData) SetIpv6MultiHomingInd(v bool)`

SetIpv6MultiHomingInd sets Ipv6MultiHomingInd field to given value.

### HasIpv6MultiHomingInd

`func (o *HsmfUpdatedData) HasIpv6MultiHomingInd() bool`

HasIpv6MultiHomingInd returns a boolean if a field has been set.

### GetQosFlowsSetupList

`func (o *HsmfUpdatedData) GetQosFlowsSetupList() []QosFlowSetupItem`

GetQosFlowsSetupList returns the QosFlowsSetupList field if non-nil, zero value otherwise.

### GetQosFlowsSetupListOk

`func (o *HsmfUpdatedData) GetQosFlowsSetupListOk() (*[]QosFlowSetupItem, bool)`

GetQosFlowsSetupListOk returns a tuple with the QosFlowsSetupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsSetupList

`func (o *HsmfUpdatedData) SetQosFlowsSetupList(v []QosFlowSetupItem)`

SetQosFlowsSetupList sets QosFlowsSetupList field to given value.

### HasQosFlowsSetupList

`func (o *HsmfUpdatedData) HasQosFlowsSetupList() bool`

HasQosFlowsSetupList returns a boolean if a field has been set.

### GetSessionAmbr

`func (o *HsmfUpdatedData) GetSessionAmbr() Ambr`

GetSessionAmbr returns the SessionAmbr field if non-nil, zero value otherwise.

### GetSessionAmbrOk

`func (o *HsmfUpdatedData) GetSessionAmbrOk() (*Ambr, bool)`

GetSessionAmbrOk returns a tuple with the SessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAmbr

`func (o *HsmfUpdatedData) SetSessionAmbr(v Ambr)`

SetSessionAmbr sets SessionAmbr field to given value.

### HasSessionAmbr

`func (o *HsmfUpdatedData) HasSessionAmbr() bool`

HasSessionAmbr returns a boolean if a field has been set.

### GetEpsPdnCnxInfo

`func (o *HsmfUpdatedData) GetEpsPdnCnxInfo() EpsPdnCnxInfo`

GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field if non-nil, zero value otherwise.

### GetEpsPdnCnxInfoOk

`func (o *HsmfUpdatedData) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool)`

GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsPdnCnxInfo

`func (o *HsmfUpdatedData) SetEpsPdnCnxInfo(v EpsPdnCnxInfo)`

SetEpsPdnCnxInfo sets EpsPdnCnxInfo field to given value.

### HasEpsPdnCnxInfo

`func (o *HsmfUpdatedData) HasEpsPdnCnxInfo() bool`

HasEpsPdnCnxInfo returns a boolean if a field has been set.

### GetEpsBearerInfo

`func (o *HsmfUpdatedData) GetEpsBearerInfo() []EpsBearerInfo`

GetEpsBearerInfo returns the EpsBearerInfo field if non-nil, zero value otherwise.

### GetEpsBearerInfoOk

`func (o *HsmfUpdatedData) GetEpsBearerInfoOk() (*[]EpsBearerInfo, bool)`

GetEpsBearerInfoOk returns a tuple with the EpsBearerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerInfo

`func (o *HsmfUpdatedData) SetEpsBearerInfo(v []EpsBearerInfo)`

SetEpsBearerInfo sets EpsBearerInfo field to given value.

### HasEpsBearerInfo

`func (o *HsmfUpdatedData) HasEpsBearerInfo() bool`

HasEpsBearerInfo returns a boolean if a field has been set.

### GetPti

`func (o *HsmfUpdatedData) GetPti() int32`

GetPti returns the Pti field if non-nil, zero value otherwise.

### GetPtiOk

`func (o *HsmfUpdatedData) GetPtiOk() (*int32, bool)`

GetPtiOk returns a tuple with the Pti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPti

`func (o *HsmfUpdatedData) SetPti(v int32)`

SetPti sets Pti field to given value.

### HasPti

`func (o *HsmfUpdatedData) HasPti() bool`

HasPti returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


