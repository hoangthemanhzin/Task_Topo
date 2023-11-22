# SmContextUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pei** | Pointer to **string** |  | [optional] 
**ServingNfId** | Pointer to **string** |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**AnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AnTypeToReactivate** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PresenceInLadn** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**HoState** | Pointer to [**HoState**](HoState.md) |  | [optional] 
**ToBeSwitched** | Pointer to **bool** |  | [optional] [default to false]
**FailedToBeSwitched** | Pointer to **bool** |  | [optional] 
**N1SmMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**TargetId** | Pointer to [**NgRanTargetId**](NgRanTargetId.md) |  | [optional] 
**TargetServingNfId** | Pointer to **string** |  | [optional] 
**SmContextStatusUri** | Pointer to **string** |  | [optional] 
**DataForwarding** | Pointer to **bool** |  | [optional] [default to false]
**N9ForwardingTunnel** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**N9DlForwardingTnlList** | Pointer to [**[]IndirectDataForwardingTunnelInfo**](IndirectDataForwardingTunnelInfo.md) |  | [optional] 
**N9UlForwardingTnlList** | Pointer to [**[]IndirectDataForwardingTunnelInfo**](IndirectDataForwardingTunnelInfo.md) |  | [optional] 
**EpsBearerSetup** | Pointer to **[]string** |  | [optional] 
**RevokeEbiList** | Pointer to **[]int32** |  | [optional] 
**Release** | Pointer to **bool** |  | [optional] [default to false]
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**NgApCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**Var5gMmCauseValue** | Pointer to **int32** |  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**TraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 
**EpsInterworkingInd** | Pointer to [**EpsInterworkingIndication**](EpsInterworkingIndication.md) |  | [optional] 
**AnTypeCanBeChanged** | Pointer to **bool** |  | [optional] [default to false]
**N2SmInfoExt1** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoTypeExt1** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**MaReleaseInd** | Pointer to [**MaReleaseIndication**](MaReleaseIndication.md) |  | [optional] 
**MaNwUpgradeInd** | Pointer to **bool** |  | [optional] [default to false]
**MaRequestInd** | Pointer to **bool** |  | [optional] [default to false]
**ExemptionInd** | Pointer to [**ExemptionInd**](ExemptionInd.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**MoExpDataCounter** | Pointer to [**MoExpDataCounter**](MoExpDataCounter.md) |  | [optional] 
**ExtendedNasSmTimerInd** | Pointer to **bool** |  | [optional] 
**ForwardingFTeid** | Pointer to **string** |  | [optional] 
**ForwardingBearerContexts** | Pointer to **[]string** |  | [optional] 
**DdnFailureSubs** | Pointer to [**DdnFailureSubs**](DdnFailureSubs.md) |  | [optional] 
**SkipN2PduSessionResRelInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSmContextUpdateData

`func NewSmContextUpdateData() *SmContextUpdateData`

NewSmContextUpdateData instantiates a new SmContextUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextUpdateDataWithDefaults

`func NewSmContextUpdateDataWithDefaults() *SmContextUpdateData`

NewSmContextUpdateDataWithDefaults instantiates a new SmContextUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPei

`func (o *SmContextUpdateData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *SmContextUpdateData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *SmContextUpdateData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *SmContextUpdateData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetServingNfId

`func (o *SmContextUpdateData) GetServingNfId() string`

GetServingNfId returns the ServingNfId field if non-nil, zero value otherwise.

### GetServingNfIdOk

`func (o *SmContextUpdateData) GetServingNfIdOk() (*string, bool)`

GetServingNfIdOk returns a tuple with the ServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfId

`func (o *SmContextUpdateData) SetServingNfId(v string)`

SetServingNfId sets ServingNfId field to given value.

### HasServingNfId

`func (o *SmContextUpdateData) HasServingNfId() bool`

HasServingNfId returns a boolean if a field has been set.

### GetGuami

`func (o *SmContextUpdateData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *SmContextUpdateData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *SmContextUpdateData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *SmContextUpdateData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmContextUpdateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmContextUpdateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmContextUpdateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *SmContextUpdateData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetBackupAmfInfo

`func (o *SmContextUpdateData) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *SmContextUpdateData) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *SmContextUpdateData) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *SmContextUpdateData) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### SetBackupAmfInfoNil

`func (o *SmContextUpdateData) SetBackupAmfInfoNil(b bool)`

 SetBackupAmfInfoNil sets the value for BackupAmfInfo to be an explicit nil

### UnsetBackupAmfInfo
`func (o *SmContextUpdateData) UnsetBackupAmfInfo()`

UnsetBackupAmfInfo ensures that no value is present for BackupAmfInfo, not even an explicit nil
### GetAnType

`func (o *SmContextUpdateData) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *SmContextUpdateData) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *SmContextUpdateData) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.

### HasAnType

`func (o *SmContextUpdateData) HasAnType() bool`

HasAnType returns a boolean if a field has been set.

### GetAdditionalAnType

`func (o *SmContextUpdateData) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *SmContextUpdateData) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *SmContextUpdateData) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *SmContextUpdateData) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.

### GetAnTypeToReactivate

`func (o *SmContextUpdateData) GetAnTypeToReactivate() AccessType`

GetAnTypeToReactivate returns the AnTypeToReactivate field if non-nil, zero value otherwise.

### GetAnTypeToReactivateOk

`func (o *SmContextUpdateData) GetAnTypeToReactivateOk() (*AccessType, bool)`

GetAnTypeToReactivateOk returns a tuple with the AnTypeToReactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnTypeToReactivate

`func (o *SmContextUpdateData) SetAnTypeToReactivate(v AccessType)`

SetAnTypeToReactivate sets AnTypeToReactivate field to given value.

### HasAnTypeToReactivate

`func (o *SmContextUpdateData) HasAnTypeToReactivate() bool`

HasAnTypeToReactivate returns a boolean if a field has been set.

### GetRatType

`func (o *SmContextUpdateData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SmContextUpdateData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SmContextUpdateData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SmContextUpdateData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetPresenceInLadn

`func (o *SmContextUpdateData) GetPresenceInLadn() PresenceState`

GetPresenceInLadn returns the PresenceInLadn field if non-nil, zero value otherwise.

### GetPresenceInLadnOk

`func (o *SmContextUpdateData) GetPresenceInLadnOk() (*PresenceState, bool)`

GetPresenceInLadnOk returns a tuple with the PresenceInLadn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInLadn

`func (o *SmContextUpdateData) SetPresenceInLadn(v PresenceState)`

SetPresenceInLadn sets PresenceInLadn field to given value.

### HasPresenceInLadn

`func (o *SmContextUpdateData) HasPresenceInLadn() bool`

HasPresenceInLadn returns a boolean if a field has been set.

### GetUeLocation

`func (o *SmContextUpdateData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SmContextUpdateData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SmContextUpdateData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SmContextUpdateData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmContextUpdateData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmContextUpdateData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmContextUpdateData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmContextUpdateData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *SmContextUpdateData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *SmContextUpdateData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *SmContextUpdateData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *SmContextUpdateData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetUpCnxState

`func (o *SmContextUpdateData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *SmContextUpdateData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *SmContextUpdateData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *SmContextUpdateData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetHoState

`func (o *SmContextUpdateData) GetHoState() HoState`

GetHoState returns the HoState field if non-nil, zero value otherwise.

### GetHoStateOk

`func (o *SmContextUpdateData) GetHoStateOk() (*HoState, bool)`

GetHoStateOk returns a tuple with the HoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoState

`func (o *SmContextUpdateData) SetHoState(v HoState)`

SetHoState sets HoState field to given value.

### HasHoState

`func (o *SmContextUpdateData) HasHoState() bool`

HasHoState returns a boolean if a field has been set.

### GetToBeSwitched

`func (o *SmContextUpdateData) GetToBeSwitched() bool`

GetToBeSwitched returns the ToBeSwitched field if non-nil, zero value otherwise.

### GetToBeSwitchedOk

`func (o *SmContextUpdateData) GetToBeSwitchedOk() (*bool, bool)`

GetToBeSwitchedOk returns a tuple with the ToBeSwitched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeSwitched

`func (o *SmContextUpdateData) SetToBeSwitched(v bool)`

SetToBeSwitched sets ToBeSwitched field to given value.

### HasToBeSwitched

`func (o *SmContextUpdateData) HasToBeSwitched() bool`

HasToBeSwitched returns a boolean if a field has been set.

### GetFailedToBeSwitched

`func (o *SmContextUpdateData) GetFailedToBeSwitched() bool`

GetFailedToBeSwitched returns the FailedToBeSwitched field if non-nil, zero value otherwise.

### GetFailedToBeSwitchedOk

`func (o *SmContextUpdateData) GetFailedToBeSwitchedOk() (*bool, bool)`

GetFailedToBeSwitchedOk returns a tuple with the FailedToBeSwitched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToBeSwitched

`func (o *SmContextUpdateData) SetFailedToBeSwitched(v bool)`

SetFailedToBeSwitched sets FailedToBeSwitched field to given value.

### HasFailedToBeSwitched

`func (o *SmContextUpdateData) HasFailedToBeSwitched() bool`

HasFailedToBeSwitched returns a boolean if a field has been set.

### GetN1SmMsg

`func (o *SmContextUpdateData) GetN1SmMsg() RefToBinaryData`

GetN1SmMsg returns the N1SmMsg field if non-nil, zero value otherwise.

### GetN1SmMsgOk

`func (o *SmContextUpdateData) GetN1SmMsgOk() (*RefToBinaryData, bool)`

GetN1SmMsgOk returns a tuple with the N1SmMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmMsg

`func (o *SmContextUpdateData) SetN1SmMsg(v RefToBinaryData)`

SetN1SmMsg sets N1SmMsg field to given value.

### HasN1SmMsg

`func (o *SmContextUpdateData) HasN1SmMsg() bool`

HasN1SmMsg returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextUpdateData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextUpdateData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextUpdateData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextUpdateData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextUpdateData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextUpdateData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextUpdateData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextUpdateData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetTargetId

`func (o *SmContextUpdateData) GetTargetId() NgRanTargetId`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SmContextUpdateData) GetTargetIdOk() (*NgRanTargetId, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SmContextUpdateData) SetTargetId(v NgRanTargetId)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SmContextUpdateData) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetServingNfId

`func (o *SmContextUpdateData) GetTargetServingNfId() string`

GetTargetServingNfId returns the TargetServingNfId field if non-nil, zero value otherwise.

### GetTargetServingNfIdOk

`func (o *SmContextUpdateData) GetTargetServingNfIdOk() (*string, bool)`

GetTargetServingNfIdOk returns a tuple with the TargetServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServingNfId

`func (o *SmContextUpdateData) SetTargetServingNfId(v string)`

SetTargetServingNfId sets TargetServingNfId field to given value.

### HasTargetServingNfId

`func (o *SmContextUpdateData) HasTargetServingNfId() bool`

HasTargetServingNfId returns a boolean if a field has been set.

### GetSmContextStatusUri

`func (o *SmContextUpdateData) GetSmContextStatusUri() string`

GetSmContextStatusUri returns the SmContextStatusUri field if non-nil, zero value otherwise.

### GetSmContextStatusUriOk

`func (o *SmContextUpdateData) GetSmContextStatusUriOk() (*string, bool)`

GetSmContextStatusUriOk returns a tuple with the SmContextStatusUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextStatusUri

`func (o *SmContextUpdateData) SetSmContextStatusUri(v string)`

SetSmContextStatusUri sets SmContextStatusUri field to given value.

### HasSmContextStatusUri

`func (o *SmContextUpdateData) HasSmContextStatusUri() bool`

HasSmContextStatusUri returns a boolean if a field has been set.

### GetDataForwarding

`func (o *SmContextUpdateData) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *SmContextUpdateData) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *SmContextUpdateData) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *SmContextUpdateData) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetN9ForwardingTunnel

`func (o *SmContextUpdateData) GetN9ForwardingTunnel() TunnelInfo`

GetN9ForwardingTunnel returns the N9ForwardingTunnel field if non-nil, zero value otherwise.

### GetN9ForwardingTunnelOk

`func (o *SmContextUpdateData) GetN9ForwardingTunnelOk() (*TunnelInfo, bool)`

GetN9ForwardingTunnelOk returns a tuple with the N9ForwardingTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN9ForwardingTunnel

`func (o *SmContextUpdateData) SetN9ForwardingTunnel(v TunnelInfo)`

SetN9ForwardingTunnel sets N9ForwardingTunnel field to given value.

### HasN9ForwardingTunnel

`func (o *SmContextUpdateData) HasN9ForwardingTunnel() bool`

HasN9ForwardingTunnel returns a boolean if a field has been set.

### GetN9DlForwardingTnlList

`func (o *SmContextUpdateData) GetN9DlForwardingTnlList() []IndirectDataForwardingTunnelInfo`

GetN9DlForwardingTnlList returns the N9DlForwardingTnlList field if non-nil, zero value otherwise.

### GetN9DlForwardingTnlListOk

`func (o *SmContextUpdateData) GetN9DlForwardingTnlListOk() (*[]IndirectDataForwardingTunnelInfo, bool)`

GetN9DlForwardingTnlListOk returns a tuple with the N9DlForwardingTnlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN9DlForwardingTnlList

`func (o *SmContextUpdateData) SetN9DlForwardingTnlList(v []IndirectDataForwardingTunnelInfo)`

SetN9DlForwardingTnlList sets N9DlForwardingTnlList field to given value.

### HasN9DlForwardingTnlList

`func (o *SmContextUpdateData) HasN9DlForwardingTnlList() bool`

HasN9DlForwardingTnlList returns a boolean if a field has been set.

### GetN9UlForwardingTnlList

`func (o *SmContextUpdateData) GetN9UlForwardingTnlList() []IndirectDataForwardingTunnelInfo`

GetN9UlForwardingTnlList returns the N9UlForwardingTnlList field if non-nil, zero value otherwise.

### GetN9UlForwardingTnlListOk

`func (o *SmContextUpdateData) GetN9UlForwardingTnlListOk() (*[]IndirectDataForwardingTunnelInfo, bool)`

GetN9UlForwardingTnlListOk returns a tuple with the N9UlForwardingTnlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN9UlForwardingTnlList

`func (o *SmContextUpdateData) SetN9UlForwardingTnlList(v []IndirectDataForwardingTunnelInfo)`

SetN9UlForwardingTnlList sets N9UlForwardingTnlList field to given value.

### HasN9UlForwardingTnlList

`func (o *SmContextUpdateData) HasN9UlForwardingTnlList() bool`

HasN9UlForwardingTnlList returns a boolean if a field has been set.

### GetEpsBearerSetup

`func (o *SmContextUpdateData) GetEpsBearerSetup() []string`

GetEpsBearerSetup returns the EpsBearerSetup field if non-nil, zero value otherwise.

### GetEpsBearerSetupOk

`func (o *SmContextUpdateData) GetEpsBearerSetupOk() (*[]string, bool)`

GetEpsBearerSetupOk returns a tuple with the EpsBearerSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerSetup

`func (o *SmContextUpdateData) SetEpsBearerSetup(v []string)`

SetEpsBearerSetup sets EpsBearerSetup field to given value.

### HasEpsBearerSetup

`func (o *SmContextUpdateData) HasEpsBearerSetup() bool`

HasEpsBearerSetup returns a boolean if a field has been set.

### GetRevokeEbiList

`func (o *SmContextUpdateData) GetRevokeEbiList() []int32`

GetRevokeEbiList returns the RevokeEbiList field if non-nil, zero value otherwise.

### GetRevokeEbiListOk

`func (o *SmContextUpdateData) GetRevokeEbiListOk() (*[]int32, bool)`

GetRevokeEbiListOk returns a tuple with the RevokeEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeEbiList

`func (o *SmContextUpdateData) SetRevokeEbiList(v []int32)`

SetRevokeEbiList sets RevokeEbiList field to given value.

### HasRevokeEbiList

`func (o *SmContextUpdateData) HasRevokeEbiList() bool`

HasRevokeEbiList returns a boolean if a field has been set.

### GetRelease

`func (o *SmContextUpdateData) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SmContextUpdateData) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SmContextUpdateData) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *SmContextUpdateData) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCause

`func (o *SmContextUpdateData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SmContextUpdateData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SmContextUpdateData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *SmContextUpdateData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetNgApCause

`func (o *SmContextUpdateData) GetNgApCause() NgApCause`

GetNgApCause returns the NgApCause field if non-nil, zero value otherwise.

### GetNgApCauseOk

`func (o *SmContextUpdateData) GetNgApCauseOk() (*NgApCause, bool)`

GetNgApCauseOk returns a tuple with the NgApCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgApCause

`func (o *SmContextUpdateData) SetNgApCause(v NgApCause)`

SetNgApCause sets NgApCause field to given value.

### HasNgApCause

`func (o *SmContextUpdateData) HasNgApCause() bool`

HasNgApCause returns a boolean if a field has been set.

### GetVar5gMmCauseValue

`func (o *SmContextUpdateData) GetVar5gMmCauseValue() int32`

GetVar5gMmCauseValue returns the Var5gMmCauseValue field if non-nil, zero value otherwise.

### GetVar5gMmCauseValueOk

`func (o *SmContextUpdateData) GetVar5gMmCauseValueOk() (*int32, bool)`

GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCauseValue

`func (o *SmContextUpdateData) SetVar5gMmCauseValue(v int32)`

SetVar5gMmCauseValue sets Var5gMmCauseValue field to given value.

### HasVar5gMmCauseValue

`func (o *SmContextUpdateData) HasVar5gMmCauseValue() bool`

HasVar5gMmCauseValue returns a boolean if a field has been set.

### GetSNssai

`func (o *SmContextUpdateData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SmContextUpdateData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SmContextUpdateData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *SmContextUpdateData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetTraceData

`func (o *SmContextUpdateData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmContextUpdateData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmContextUpdateData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmContextUpdateData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmContextUpdateData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmContextUpdateData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetEpsInterworkingInd

`func (o *SmContextUpdateData) GetEpsInterworkingInd() EpsInterworkingIndication`

GetEpsInterworkingInd returns the EpsInterworkingInd field if non-nil, zero value otherwise.

### GetEpsInterworkingIndOk

`func (o *SmContextUpdateData) GetEpsInterworkingIndOk() (*EpsInterworkingIndication, bool)`

GetEpsInterworkingIndOk returns a tuple with the EpsInterworkingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInd

`func (o *SmContextUpdateData) SetEpsInterworkingInd(v EpsInterworkingIndication)`

SetEpsInterworkingInd sets EpsInterworkingInd field to given value.

### HasEpsInterworkingInd

`func (o *SmContextUpdateData) HasEpsInterworkingInd() bool`

HasEpsInterworkingInd returns a boolean if a field has been set.

### GetAnTypeCanBeChanged

`func (o *SmContextUpdateData) GetAnTypeCanBeChanged() bool`

GetAnTypeCanBeChanged returns the AnTypeCanBeChanged field if non-nil, zero value otherwise.

### GetAnTypeCanBeChangedOk

`func (o *SmContextUpdateData) GetAnTypeCanBeChangedOk() (*bool, bool)`

GetAnTypeCanBeChangedOk returns a tuple with the AnTypeCanBeChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnTypeCanBeChanged

`func (o *SmContextUpdateData) SetAnTypeCanBeChanged(v bool)`

SetAnTypeCanBeChanged sets AnTypeCanBeChanged field to given value.

### HasAnTypeCanBeChanged

`func (o *SmContextUpdateData) HasAnTypeCanBeChanged() bool`

HasAnTypeCanBeChanged returns a boolean if a field has been set.

### GetN2SmInfoExt1

`func (o *SmContextUpdateData) GetN2SmInfoExt1() RefToBinaryData`

GetN2SmInfoExt1 returns the N2SmInfoExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoExt1Ok

`func (o *SmContextUpdateData) GetN2SmInfoExt1Ok() (*RefToBinaryData, bool)`

GetN2SmInfoExt1Ok returns a tuple with the N2SmInfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoExt1

`func (o *SmContextUpdateData) SetN2SmInfoExt1(v RefToBinaryData)`

SetN2SmInfoExt1 sets N2SmInfoExt1 field to given value.

### HasN2SmInfoExt1

`func (o *SmContextUpdateData) HasN2SmInfoExt1() bool`

HasN2SmInfoExt1 returns a boolean if a field has been set.

### GetN2SmInfoTypeExt1

`func (o *SmContextUpdateData) GetN2SmInfoTypeExt1() N2SmInfoType`

GetN2SmInfoTypeExt1 returns the N2SmInfoTypeExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoTypeExt1Ok

`func (o *SmContextUpdateData) GetN2SmInfoTypeExt1Ok() (*N2SmInfoType, bool)`

GetN2SmInfoTypeExt1Ok returns a tuple with the N2SmInfoTypeExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoTypeExt1

`func (o *SmContextUpdateData) SetN2SmInfoTypeExt1(v N2SmInfoType)`

SetN2SmInfoTypeExt1 sets N2SmInfoTypeExt1 field to given value.

### HasN2SmInfoTypeExt1

`func (o *SmContextUpdateData) HasN2SmInfoTypeExt1() bool`

HasN2SmInfoTypeExt1 returns a boolean if a field has been set.

### GetMaReleaseInd

`func (o *SmContextUpdateData) GetMaReleaseInd() MaReleaseIndication`

GetMaReleaseInd returns the MaReleaseInd field if non-nil, zero value otherwise.

### GetMaReleaseIndOk

`func (o *SmContextUpdateData) GetMaReleaseIndOk() (*MaReleaseIndication, bool)`

GetMaReleaseIndOk returns a tuple with the MaReleaseInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaReleaseInd

`func (o *SmContextUpdateData) SetMaReleaseInd(v MaReleaseIndication)`

SetMaReleaseInd sets MaReleaseInd field to given value.

### HasMaReleaseInd

`func (o *SmContextUpdateData) HasMaReleaseInd() bool`

HasMaReleaseInd returns a boolean if a field has been set.

### GetMaNwUpgradeInd

`func (o *SmContextUpdateData) GetMaNwUpgradeInd() bool`

GetMaNwUpgradeInd returns the MaNwUpgradeInd field if non-nil, zero value otherwise.

### GetMaNwUpgradeIndOk

`func (o *SmContextUpdateData) GetMaNwUpgradeIndOk() (*bool, bool)`

GetMaNwUpgradeIndOk returns a tuple with the MaNwUpgradeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaNwUpgradeInd

`func (o *SmContextUpdateData) SetMaNwUpgradeInd(v bool)`

SetMaNwUpgradeInd sets MaNwUpgradeInd field to given value.

### HasMaNwUpgradeInd

`func (o *SmContextUpdateData) HasMaNwUpgradeInd() bool`

HasMaNwUpgradeInd returns a boolean if a field has been set.

### GetMaRequestInd

`func (o *SmContextUpdateData) GetMaRequestInd() bool`

GetMaRequestInd returns the MaRequestInd field if non-nil, zero value otherwise.

### GetMaRequestIndOk

`func (o *SmContextUpdateData) GetMaRequestIndOk() (*bool, bool)`

GetMaRequestIndOk returns a tuple with the MaRequestInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaRequestInd

`func (o *SmContextUpdateData) SetMaRequestInd(v bool)`

SetMaRequestInd sets MaRequestInd field to given value.

### HasMaRequestInd

`func (o *SmContextUpdateData) HasMaRequestInd() bool`

HasMaRequestInd returns a boolean if a field has been set.

### GetExemptionInd

`func (o *SmContextUpdateData) GetExemptionInd() ExemptionInd`

GetExemptionInd returns the ExemptionInd field if non-nil, zero value otherwise.

### GetExemptionIndOk

`func (o *SmContextUpdateData) GetExemptionIndOk() (*ExemptionInd, bool)`

GetExemptionIndOk returns a tuple with the ExemptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptionInd

`func (o *SmContextUpdateData) SetExemptionInd(v ExemptionInd)`

SetExemptionInd sets ExemptionInd field to given value.

### HasExemptionInd

`func (o *SmContextUpdateData) HasExemptionInd() bool`

HasExemptionInd returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextUpdateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextUpdateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextUpdateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextUpdateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMoExpDataCounter

`func (o *SmContextUpdateData) GetMoExpDataCounter() MoExpDataCounter`

GetMoExpDataCounter returns the MoExpDataCounter field if non-nil, zero value otherwise.

### GetMoExpDataCounterOk

`func (o *SmContextUpdateData) GetMoExpDataCounterOk() (*MoExpDataCounter, bool)`

GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoExpDataCounter

`func (o *SmContextUpdateData) SetMoExpDataCounter(v MoExpDataCounter)`

SetMoExpDataCounter sets MoExpDataCounter field to given value.

### HasMoExpDataCounter

`func (o *SmContextUpdateData) HasMoExpDataCounter() bool`

HasMoExpDataCounter returns a boolean if a field has been set.

### GetExtendedNasSmTimerInd

`func (o *SmContextUpdateData) GetExtendedNasSmTimerInd() bool`

GetExtendedNasSmTimerInd returns the ExtendedNasSmTimerInd field if non-nil, zero value otherwise.

### GetExtendedNasSmTimerIndOk

`func (o *SmContextUpdateData) GetExtendedNasSmTimerIndOk() (*bool, bool)`

GetExtendedNasSmTimerIndOk returns a tuple with the ExtendedNasSmTimerInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedNasSmTimerInd

`func (o *SmContextUpdateData) SetExtendedNasSmTimerInd(v bool)`

SetExtendedNasSmTimerInd sets ExtendedNasSmTimerInd field to given value.

### HasExtendedNasSmTimerInd

`func (o *SmContextUpdateData) HasExtendedNasSmTimerInd() bool`

HasExtendedNasSmTimerInd returns a boolean if a field has been set.

### GetForwardingFTeid

`func (o *SmContextUpdateData) GetForwardingFTeid() string`

GetForwardingFTeid returns the ForwardingFTeid field if non-nil, zero value otherwise.

### GetForwardingFTeidOk

`func (o *SmContextUpdateData) GetForwardingFTeidOk() (*string, bool)`

GetForwardingFTeidOk returns a tuple with the ForwardingFTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingFTeid

`func (o *SmContextUpdateData) SetForwardingFTeid(v string)`

SetForwardingFTeid sets ForwardingFTeid field to given value.

### HasForwardingFTeid

`func (o *SmContextUpdateData) HasForwardingFTeid() bool`

HasForwardingFTeid returns a boolean if a field has been set.

### GetForwardingBearerContexts

`func (o *SmContextUpdateData) GetForwardingBearerContexts() []string`

GetForwardingBearerContexts returns the ForwardingBearerContexts field if non-nil, zero value otherwise.

### GetForwardingBearerContextsOk

`func (o *SmContextUpdateData) GetForwardingBearerContextsOk() (*[]string, bool)`

GetForwardingBearerContextsOk returns a tuple with the ForwardingBearerContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingBearerContexts

`func (o *SmContextUpdateData) SetForwardingBearerContexts(v []string)`

SetForwardingBearerContexts sets ForwardingBearerContexts field to given value.

### HasForwardingBearerContexts

`func (o *SmContextUpdateData) HasForwardingBearerContexts() bool`

HasForwardingBearerContexts returns a boolean if a field has been set.

### GetDdnFailureSubs

`func (o *SmContextUpdateData) GetDdnFailureSubs() DdnFailureSubs`

GetDdnFailureSubs returns the DdnFailureSubs field if non-nil, zero value otherwise.

### GetDdnFailureSubsOk

`func (o *SmContextUpdateData) GetDdnFailureSubsOk() (*DdnFailureSubs, bool)`

GetDdnFailureSubsOk returns a tuple with the DdnFailureSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnFailureSubs

`func (o *SmContextUpdateData) SetDdnFailureSubs(v DdnFailureSubs)`

SetDdnFailureSubs sets DdnFailureSubs field to given value.

### HasDdnFailureSubs

`func (o *SmContextUpdateData) HasDdnFailureSubs() bool`

HasDdnFailureSubs returns a boolean if a field has been set.

### GetSkipN2PduSessionResRelInd

`func (o *SmContextUpdateData) GetSkipN2PduSessionResRelInd() bool`

GetSkipN2PduSessionResRelInd returns the SkipN2PduSessionResRelInd field if non-nil, zero value otherwise.

### GetSkipN2PduSessionResRelIndOk

`func (o *SmContextUpdateData) GetSkipN2PduSessionResRelIndOk() (*bool, bool)`

GetSkipN2PduSessionResRelIndOk returns a tuple with the SkipN2PduSessionResRelInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipN2PduSessionResRelInd

`func (o *SmContextUpdateData) SetSkipN2PduSessionResRelInd(v bool)`

SetSkipN2PduSessionResRelInd sets SkipN2PduSessionResRelInd field to given value.

### HasSkipN2PduSessionResRelInd

`func (o *SmContextUpdateData) HasSkipN2PduSessionResRelInd() bool`

HasSkipN2PduSessionResRelInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


