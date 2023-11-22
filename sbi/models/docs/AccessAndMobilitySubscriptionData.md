# AccessAndMobilitySubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**InternalGroupIds** | Pointer to **[]string** |  | [optional] 
**SharedVnGroupDataIds** | Pointer to **map[string]string** |  | [optional] 
**SubscribedUeAmbr** | Pointer to [**AmbrRm**](AmbrRm.md) |  | [optional] 
**Nssai** | Pointer to [**Nssai**](Nssai.md) |  | [optional] 
**RatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**ForbiddenAreas** | Pointer to [**[]Area**](Area.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**CoreNetworkTypeRestrictions** | Pointer to [**[]CoreNetworkType**](CoreNetworkType.md) |  | [optional] 
**RfspIndex** | Pointer to **int32** |  | [optional] 
**SubsRegTimer** | Pointer to **int32** |  | [optional] 
**UeUsageType** | Pointer to **int32** |  | [optional] 
**MpsPriority** | Pointer to **bool** |  | [optional] 
**McsPriority** | Pointer to **bool** |  | [optional] 
**ActiveTime** | Pointer to **int32** |  | [optional] 
**SorInfo** | Pointer to [**SorInfo**](SorInfo.md) |  | [optional] 
**SorInfoExpectInd** | Pointer to **bool** |  | [optional] 
**SorafRetrieval** | Pointer to **bool** |  | [optional] [default to false]
**SorUpdateIndicatorList** | Pointer to [**[]SorUpdateIndicator**](SorUpdateIndicator.md) |  | [optional] 
**UpuInfo** | Pointer to [**UpuInfo**](UpuInfo.md) |  | [optional] 
**MicoAllowed** | Pointer to **bool** |  | [optional] 
**SharedAmDataIds** | Pointer to **[]string** |  | [optional] 
**OdbPacketServices** | Pointer to [**OdbPacketServices**](OdbPacketServices.md) |  | [optional] 
**SubscribedDnnList** | Pointer to [**[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner**](AccessAndMobilitySubscriptionDataSubscribedDnnListInner.md) |  | [optional] 
**ServiceGapTime** | Pointer to **int32** |  | [optional] 
**MdtUserConsent** | Pointer to [**MdtUserConsent**](MdtUserConsent.md) |  | [optional] 
**MdtConfiguration** | Pointer to [**MdtConfiguration**](MdtConfiguration.md) |  | [optional] 
**TraceData** | Pointer to [**TraceData1**](TraceData1.md) |  | [optional] 
**CagData** | Pointer to [**CagData**](CagData.md) |  | [optional] 
**StnSr** | Pointer to **string** |  | [optional] 
**CMsisdn** | Pointer to **string** |  | [optional] 
**NbIoTUePriority** | Pointer to **int32** |  | [optional] 
**NssaiInclusionAllowed** | Pointer to **bool** |  | [optional] [default to false]
**RgWirelineCharacteristics** | Pointer to **string** |  | [optional] 
**EcRestrictionDataWb** | Pointer to [**EcRestrictionDataWb**](EcRestrictionDataWb.md) |  | [optional] 
**EcRestrictionDataNb** | Pointer to **bool** |  | [optional] [default to false]
**ExpectedUeBehaviourList** | Pointer to [**ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) |  | [optional] 
**PrimaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**SecondaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**EdrxParametersList** | Pointer to [**[]EdrxParameters**](EdrxParameters.md) |  | [optional] 
**PtwParametersList** | Pointer to [**[]PtwParameters**](PtwParameters.md) |  | [optional] 
**IabOperationAllowed** | Pointer to **bool** |  | [optional] [default to false]
**WirelineForbiddenAreas** | Pointer to [**[]WirelineArea**](WirelineArea.md) |  | [optional] 
**WirelineServiceAreaRestriction** | Pointer to [**WirelineServiceAreaRestriction**](WirelineServiceAreaRestriction.md) |  | [optional] 

## Methods

### NewAccessAndMobilitySubscriptionData

`func NewAccessAndMobilitySubscriptionData() *AccessAndMobilitySubscriptionData`

NewAccessAndMobilitySubscriptionData instantiates a new AccessAndMobilitySubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessAndMobilitySubscriptionDataWithDefaults

`func NewAccessAndMobilitySubscriptionDataWithDefaults() *AccessAndMobilitySubscriptionData`

NewAccessAndMobilitySubscriptionDataWithDefaults instantiates a new AccessAndMobilitySubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AccessAndMobilitySubscriptionData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetGpsis

`func (o *AccessAndMobilitySubscriptionData) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *AccessAndMobilitySubscriptionData) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *AccessAndMobilitySubscriptionData) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *AccessAndMobilitySubscriptionData) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData) GetInternalGroupIds() []string`

GetInternalGroupIds returns the InternalGroupIds field if non-nil, zero value otherwise.

### GetInternalGroupIdsOk

`func (o *AccessAndMobilitySubscriptionData) GetInternalGroupIdsOk() (*[]string, bool)`

GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData) SetInternalGroupIds(v []string)`

SetInternalGroupIds sets InternalGroupIds field to given value.

### HasInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData) HasInternalGroupIds() bool`

HasInternalGroupIds returns a boolean if a field has been set.

### GetSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData) GetSharedVnGroupDataIds() map[string]string`

GetSharedVnGroupDataIds returns the SharedVnGroupDataIds field if non-nil, zero value otherwise.

### GetSharedVnGroupDataIdsOk

`func (o *AccessAndMobilitySubscriptionData) GetSharedVnGroupDataIdsOk() (*map[string]string, bool)`

GetSharedVnGroupDataIdsOk returns a tuple with the SharedVnGroupDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData) SetSharedVnGroupDataIds(v map[string]string)`

SetSharedVnGroupDataIds sets SharedVnGroupDataIds field to given value.

### HasSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData) HasSharedVnGroupDataIds() bool`

HasSharedVnGroupDataIds returns a boolean if a field has been set.

### GetSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData) GetSubscribedUeAmbr() AmbrRm`

GetSubscribedUeAmbr returns the SubscribedUeAmbr field if non-nil, zero value otherwise.

### GetSubscribedUeAmbrOk

`func (o *AccessAndMobilitySubscriptionData) GetSubscribedUeAmbrOk() (*AmbrRm, bool)`

GetSubscribedUeAmbrOk returns a tuple with the SubscribedUeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData) SetSubscribedUeAmbr(v AmbrRm)`

SetSubscribedUeAmbr sets SubscribedUeAmbr field to given value.

### HasSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData) HasSubscribedUeAmbr() bool`

HasSubscribedUeAmbr returns a boolean if a field has been set.

### GetNssai

`func (o *AccessAndMobilitySubscriptionData) GetNssai() Nssai`

GetNssai returns the Nssai field if non-nil, zero value otherwise.

### GetNssaiOk

`func (o *AccessAndMobilitySubscriptionData) GetNssaiOk() (*Nssai, bool)`

GetNssaiOk returns a tuple with the Nssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssai

`func (o *AccessAndMobilitySubscriptionData) SetNssai(v Nssai)`

SetNssai sets Nssai field to given value.

### HasNssai

`func (o *AccessAndMobilitySubscriptionData) HasNssai() bool`

HasNssai returns a boolean if a field has been set.

### SetNssaiNil

`func (o *AccessAndMobilitySubscriptionData) SetNssaiNil(b bool)`

 SetNssaiNil sets the value for Nssai to be an explicit nil

### UnsetNssai
`func (o *AccessAndMobilitySubscriptionData) UnsetNssai()`

UnsetNssai ensures that no value is present for Nssai, not even an explicit nil
### GetRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) GetRatRestrictions() []RatType`

GetRatRestrictions returns the RatRestrictions field if non-nil, zero value otherwise.

### GetRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData) GetRatRestrictionsOk() (*[]RatType, bool)`

GetRatRestrictionsOk returns a tuple with the RatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) SetRatRestrictions(v []RatType)`

SetRatRestrictions sets RatRestrictions field to given value.

### HasRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) HasRatRestrictions() bool`

HasRatRestrictions returns a boolean if a field has been set.

### GetForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) GetForbiddenAreas() []Area`

GetForbiddenAreas returns the ForbiddenAreas field if non-nil, zero value otherwise.

### GetForbiddenAreasOk

`func (o *AccessAndMobilitySubscriptionData) GetForbiddenAreasOk() (*[]Area, bool)`

GetForbiddenAreasOk returns a tuple with the ForbiddenAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) SetForbiddenAreas(v []Area)`

SetForbiddenAreas sets ForbiddenAreas field to given value.

### HasForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) HasForbiddenAreas() bool`

HasForbiddenAreas returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) GetServiceAreaRestriction() ServiceAreaRestriction`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *AccessAndMobilitySubscriptionData) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) SetServiceAreaRestriction(v ServiceAreaRestriction)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData) GetCoreNetworkTypeRestrictions() []CoreNetworkType`

GetCoreNetworkTypeRestrictions returns the CoreNetworkTypeRestrictions field if non-nil, zero value otherwise.

### GetCoreNetworkTypeRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData) GetCoreNetworkTypeRestrictionsOk() (*[]CoreNetworkType, bool)`

GetCoreNetworkTypeRestrictionsOk returns a tuple with the CoreNetworkTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData) SetCoreNetworkTypeRestrictions(v []CoreNetworkType)`

SetCoreNetworkTypeRestrictions sets CoreNetworkTypeRestrictions field to given value.

### HasCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData) HasCoreNetworkTypeRestrictions() bool`

HasCoreNetworkTypeRestrictions returns a boolean if a field has been set.

### GetRfspIndex

`func (o *AccessAndMobilitySubscriptionData) GetRfspIndex() int32`

GetRfspIndex returns the RfspIndex field if non-nil, zero value otherwise.

### GetRfspIndexOk

`func (o *AccessAndMobilitySubscriptionData) GetRfspIndexOk() (*int32, bool)`

GetRfspIndexOk returns a tuple with the RfspIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfspIndex

`func (o *AccessAndMobilitySubscriptionData) SetRfspIndex(v int32)`

SetRfspIndex sets RfspIndex field to given value.

### HasRfspIndex

`func (o *AccessAndMobilitySubscriptionData) HasRfspIndex() bool`

HasRfspIndex returns a boolean if a field has been set.

### SetRfspIndexNil

`func (o *AccessAndMobilitySubscriptionData) SetRfspIndexNil(b bool)`

 SetRfspIndexNil sets the value for RfspIndex to be an explicit nil

### UnsetRfspIndex
`func (o *AccessAndMobilitySubscriptionData) UnsetRfspIndex()`

UnsetRfspIndex ensures that no value is present for RfspIndex, not even an explicit nil
### GetSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData) GetSubsRegTimer() int32`

GetSubsRegTimer returns the SubsRegTimer field if non-nil, zero value otherwise.

### GetSubsRegTimerOk

`func (o *AccessAndMobilitySubscriptionData) GetSubsRegTimerOk() (*int32, bool)`

GetSubsRegTimerOk returns a tuple with the SubsRegTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData) SetSubsRegTimer(v int32)`

SetSubsRegTimer sets SubsRegTimer field to given value.

### HasSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData) HasSubsRegTimer() bool`

HasSubsRegTimer returns a boolean if a field has been set.

### SetSubsRegTimerNil

`func (o *AccessAndMobilitySubscriptionData) SetSubsRegTimerNil(b bool)`

 SetSubsRegTimerNil sets the value for SubsRegTimer to be an explicit nil

### UnsetSubsRegTimer
`func (o *AccessAndMobilitySubscriptionData) UnsetSubsRegTimer()`

UnsetSubsRegTimer ensures that no value is present for SubsRegTimer, not even an explicit nil
### GetUeUsageType

`func (o *AccessAndMobilitySubscriptionData) GetUeUsageType() int32`

GetUeUsageType returns the UeUsageType field if non-nil, zero value otherwise.

### GetUeUsageTypeOk

`func (o *AccessAndMobilitySubscriptionData) GetUeUsageTypeOk() (*int32, bool)`

GetUeUsageTypeOk returns a tuple with the UeUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeUsageType

`func (o *AccessAndMobilitySubscriptionData) SetUeUsageType(v int32)`

SetUeUsageType sets UeUsageType field to given value.

### HasUeUsageType

`func (o *AccessAndMobilitySubscriptionData) HasUeUsageType() bool`

HasUeUsageType returns a boolean if a field has been set.

### GetMpsPriority

`func (o *AccessAndMobilitySubscriptionData) GetMpsPriority() bool`

GetMpsPriority returns the MpsPriority field if non-nil, zero value otherwise.

### GetMpsPriorityOk

`func (o *AccessAndMobilitySubscriptionData) GetMpsPriorityOk() (*bool, bool)`

GetMpsPriorityOk returns a tuple with the MpsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsPriority

`func (o *AccessAndMobilitySubscriptionData) SetMpsPriority(v bool)`

SetMpsPriority sets MpsPriority field to given value.

### HasMpsPriority

`func (o *AccessAndMobilitySubscriptionData) HasMpsPriority() bool`

HasMpsPriority returns a boolean if a field has been set.

### GetMcsPriority

`func (o *AccessAndMobilitySubscriptionData) GetMcsPriority() bool`

GetMcsPriority returns the McsPriority field if non-nil, zero value otherwise.

### GetMcsPriorityOk

`func (o *AccessAndMobilitySubscriptionData) GetMcsPriorityOk() (*bool, bool)`

GetMcsPriorityOk returns a tuple with the McsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsPriority

`func (o *AccessAndMobilitySubscriptionData) SetMcsPriority(v bool)`

SetMcsPriority sets McsPriority field to given value.

### HasMcsPriority

`func (o *AccessAndMobilitySubscriptionData) HasMcsPriority() bool`

HasMcsPriority returns a boolean if a field has been set.

### GetActiveTime

`func (o *AccessAndMobilitySubscriptionData) GetActiveTime() int32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *AccessAndMobilitySubscriptionData) GetActiveTimeOk() (*int32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *AccessAndMobilitySubscriptionData) SetActiveTime(v int32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *AccessAndMobilitySubscriptionData) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### SetActiveTimeNil

`func (o *AccessAndMobilitySubscriptionData) SetActiveTimeNil(b bool)`

 SetActiveTimeNil sets the value for ActiveTime to be an explicit nil

### UnsetActiveTime
`func (o *AccessAndMobilitySubscriptionData) UnsetActiveTime()`

UnsetActiveTime ensures that no value is present for ActiveTime, not even an explicit nil
### GetSorInfo

`func (o *AccessAndMobilitySubscriptionData) GetSorInfo() SorInfo`

GetSorInfo returns the SorInfo field if non-nil, zero value otherwise.

### GetSorInfoOk

`func (o *AccessAndMobilitySubscriptionData) GetSorInfoOk() (*SorInfo, bool)`

GetSorInfoOk returns a tuple with the SorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfo

`func (o *AccessAndMobilitySubscriptionData) SetSorInfo(v SorInfo)`

SetSorInfo sets SorInfo field to given value.

### HasSorInfo

`func (o *AccessAndMobilitySubscriptionData) HasSorInfo() bool`

HasSorInfo returns a boolean if a field has been set.

### GetSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData) GetSorInfoExpectInd() bool`

GetSorInfoExpectInd returns the SorInfoExpectInd field if non-nil, zero value otherwise.

### GetSorInfoExpectIndOk

`func (o *AccessAndMobilitySubscriptionData) GetSorInfoExpectIndOk() (*bool, bool)`

GetSorInfoExpectIndOk returns a tuple with the SorInfoExpectInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData) SetSorInfoExpectInd(v bool)`

SetSorInfoExpectInd sets SorInfoExpectInd field to given value.

### HasSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData) HasSorInfoExpectInd() bool`

HasSorInfoExpectInd returns a boolean if a field has been set.

### GetSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData) GetSorafRetrieval() bool`

GetSorafRetrieval returns the SorafRetrieval field if non-nil, zero value otherwise.

### GetSorafRetrievalOk

`func (o *AccessAndMobilitySubscriptionData) GetSorafRetrievalOk() (*bool, bool)`

GetSorafRetrievalOk returns a tuple with the SorafRetrieval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData) SetSorafRetrieval(v bool)`

SetSorafRetrieval sets SorafRetrieval field to given value.

### HasSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData) HasSorafRetrieval() bool`

HasSorafRetrieval returns a boolean if a field has been set.

### GetSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData) GetSorUpdateIndicatorList() []SorUpdateIndicator`

GetSorUpdateIndicatorList returns the SorUpdateIndicatorList field if non-nil, zero value otherwise.

### GetSorUpdateIndicatorListOk

`func (o *AccessAndMobilitySubscriptionData) GetSorUpdateIndicatorListOk() (*[]SorUpdateIndicator, bool)`

GetSorUpdateIndicatorListOk returns a tuple with the SorUpdateIndicatorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData) SetSorUpdateIndicatorList(v []SorUpdateIndicator)`

SetSorUpdateIndicatorList sets SorUpdateIndicatorList field to given value.

### HasSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData) HasSorUpdateIndicatorList() bool`

HasSorUpdateIndicatorList returns a boolean if a field has been set.

### GetUpuInfo

`func (o *AccessAndMobilitySubscriptionData) GetUpuInfo() UpuInfo`

GetUpuInfo returns the UpuInfo field if non-nil, zero value otherwise.

### GetUpuInfoOk

`func (o *AccessAndMobilitySubscriptionData) GetUpuInfoOk() (*UpuInfo, bool)`

GetUpuInfoOk returns a tuple with the UpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuInfo

`func (o *AccessAndMobilitySubscriptionData) SetUpuInfo(v UpuInfo)`

SetUpuInfo sets UpuInfo field to given value.

### HasUpuInfo

`func (o *AccessAndMobilitySubscriptionData) HasUpuInfo() bool`

HasUpuInfo returns a boolean if a field has been set.

### GetMicoAllowed

`func (o *AccessAndMobilitySubscriptionData) GetMicoAllowed() bool`

GetMicoAllowed returns the MicoAllowed field if non-nil, zero value otherwise.

### GetMicoAllowedOk

`func (o *AccessAndMobilitySubscriptionData) GetMicoAllowedOk() (*bool, bool)`

GetMicoAllowedOk returns a tuple with the MicoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicoAllowed

`func (o *AccessAndMobilitySubscriptionData) SetMicoAllowed(v bool)`

SetMicoAllowed sets MicoAllowed field to given value.

### HasMicoAllowed

`func (o *AccessAndMobilitySubscriptionData) HasMicoAllowed() bool`

HasMicoAllowed returns a boolean if a field has been set.

### GetSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData) GetSharedAmDataIds() []string`

GetSharedAmDataIds returns the SharedAmDataIds field if non-nil, zero value otherwise.

### GetSharedAmDataIdsOk

`func (o *AccessAndMobilitySubscriptionData) GetSharedAmDataIdsOk() (*[]string, bool)`

GetSharedAmDataIdsOk returns a tuple with the SharedAmDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData) SetSharedAmDataIds(v []string)`

SetSharedAmDataIds sets SharedAmDataIds field to given value.

### HasSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData) HasSharedAmDataIds() bool`

HasSharedAmDataIds returns a boolean if a field has been set.

### GetOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData) GetOdbPacketServices() OdbPacketServices`

GetOdbPacketServices returns the OdbPacketServices field if non-nil, zero value otherwise.

### GetOdbPacketServicesOk

`func (o *AccessAndMobilitySubscriptionData) GetOdbPacketServicesOk() (*OdbPacketServices, bool)`

GetOdbPacketServicesOk returns a tuple with the OdbPacketServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData) SetOdbPacketServices(v OdbPacketServices)`

SetOdbPacketServices sets OdbPacketServices field to given value.

### HasOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData) HasOdbPacketServices() bool`

HasOdbPacketServices returns a boolean if a field has been set.

### GetSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData) GetSubscribedDnnList() []AccessAndMobilitySubscriptionDataSubscribedDnnListInner`

GetSubscribedDnnList returns the SubscribedDnnList field if non-nil, zero value otherwise.

### GetSubscribedDnnListOk

`func (o *AccessAndMobilitySubscriptionData) GetSubscribedDnnListOk() (*[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool)`

GetSubscribedDnnListOk returns a tuple with the SubscribedDnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData) SetSubscribedDnnList(v []AccessAndMobilitySubscriptionDataSubscribedDnnListInner)`

SetSubscribedDnnList sets SubscribedDnnList field to given value.

### HasSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData) HasSubscribedDnnList() bool`

HasSubscribedDnnList returns a boolean if a field has been set.

### GetServiceGapTime

`func (o *AccessAndMobilitySubscriptionData) GetServiceGapTime() int32`

GetServiceGapTime returns the ServiceGapTime field if non-nil, zero value otherwise.

### GetServiceGapTimeOk

`func (o *AccessAndMobilitySubscriptionData) GetServiceGapTimeOk() (*int32, bool)`

GetServiceGapTimeOk returns a tuple with the ServiceGapTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGapTime

`func (o *AccessAndMobilitySubscriptionData) SetServiceGapTime(v int32)`

SetServiceGapTime sets ServiceGapTime field to given value.

### HasServiceGapTime

`func (o *AccessAndMobilitySubscriptionData) HasServiceGapTime() bool`

HasServiceGapTime returns a boolean if a field has been set.

### GetMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData) GetMdtUserConsent() MdtUserConsent`

GetMdtUserConsent returns the MdtUserConsent field if non-nil, zero value otherwise.

### GetMdtUserConsentOk

`func (o *AccessAndMobilitySubscriptionData) GetMdtUserConsentOk() (*MdtUserConsent, bool)`

GetMdtUserConsentOk returns a tuple with the MdtUserConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData) SetMdtUserConsent(v MdtUserConsent)`

SetMdtUserConsent sets MdtUserConsent field to given value.

### HasMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData) HasMdtUserConsent() bool`

HasMdtUserConsent returns a boolean if a field has been set.

### GetMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData) GetMdtConfiguration() MdtConfiguration`

GetMdtConfiguration returns the MdtConfiguration field if non-nil, zero value otherwise.

### GetMdtConfigurationOk

`func (o *AccessAndMobilitySubscriptionData) GetMdtConfigurationOk() (*MdtConfiguration, bool)`

GetMdtConfigurationOk returns a tuple with the MdtConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData) SetMdtConfiguration(v MdtConfiguration)`

SetMdtConfiguration sets MdtConfiguration field to given value.

### HasMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData) HasMdtConfiguration() bool`

HasMdtConfiguration returns a boolean if a field has been set.

### GetTraceData

`func (o *AccessAndMobilitySubscriptionData) GetTraceData() TraceData1`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *AccessAndMobilitySubscriptionData) GetTraceDataOk() (*TraceData1, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *AccessAndMobilitySubscriptionData) SetTraceData(v TraceData1)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *AccessAndMobilitySubscriptionData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *AccessAndMobilitySubscriptionData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *AccessAndMobilitySubscriptionData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetCagData

`func (o *AccessAndMobilitySubscriptionData) GetCagData() CagData`

GetCagData returns the CagData field if non-nil, zero value otherwise.

### GetCagDataOk

`func (o *AccessAndMobilitySubscriptionData) GetCagDataOk() (*CagData, bool)`

GetCagDataOk returns a tuple with the CagData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagData

`func (o *AccessAndMobilitySubscriptionData) SetCagData(v CagData)`

SetCagData sets CagData field to given value.

### HasCagData

`func (o *AccessAndMobilitySubscriptionData) HasCagData() bool`

HasCagData returns a boolean if a field has been set.

### GetStnSr

`func (o *AccessAndMobilitySubscriptionData) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *AccessAndMobilitySubscriptionData) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *AccessAndMobilitySubscriptionData) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *AccessAndMobilitySubscriptionData) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### GetCMsisdn

`func (o *AccessAndMobilitySubscriptionData) GetCMsisdn() string`

GetCMsisdn returns the CMsisdn field if non-nil, zero value otherwise.

### GetCMsisdnOk

`func (o *AccessAndMobilitySubscriptionData) GetCMsisdnOk() (*string, bool)`

GetCMsisdnOk returns a tuple with the CMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCMsisdn

`func (o *AccessAndMobilitySubscriptionData) SetCMsisdn(v string)`

SetCMsisdn sets CMsisdn field to given value.

### HasCMsisdn

`func (o *AccessAndMobilitySubscriptionData) HasCMsisdn() bool`

HasCMsisdn returns a boolean if a field has been set.

### GetNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData) GetNbIoTUePriority() int32`

GetNbIoTUePriority returns the NbIoTUePriority field if non-nil, zero value otherwise.

### GetNbIoTUePriorityOk

`func (o *AccessAndMobilitySubscriptionData) GetNbIoTUePriorityOk() (*int32, bool)`

GetNbIoTUePriorityOk returns a tuple with the NbIoTUePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData) SetNbIoTUePriority(v int32)`

SetNbIoTUePriority sets NbIoTUePriority field to given value.

### HasNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData) HasNbIoTUePriority() bool`

HasNbIoTUePriority returns a boolean if a field has been set.

### GetNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData) GetNssaiInclusionAllowed() bool`

GetNssaiInclusionAllowed returns the NssaiInclusionAllowed field if non-nil, zero value otherwise.

### GetNssaiInclusionAllowedOk

`func (o *AccessAndMobilitySubscriptionData) GetNssaiInclusionAllowedOk() (*bool, bool)`

GetNssaiInclusionAllowedOk returns a tuple with the NssaiInclusionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData) SetNssaiInclusionAllowed(v bool)`

SetNssaiInclusionAllowed sets NssaiInclusionAllowed field to given value.

### HasNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData) HasNssaiInclusionAllowed() bool`

HasNssaiInclusionAllowed returns a boolean if a field has been set.

### GetRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData) GetRgWirelineCharacteristics() string`

GetRgWirelineCharacteristics returns the RgWirelineCharacteristics field if non-nil, zero value otherwise.

### GetRgWirelineCharacteristicsOk

`func (o *AccessAndMobilitySubscriptionData) GetRgWirelineCharacteristicsOk() (*string, bool)`

GetRgWirelineCharacteristicsOk returns a tuple with the RgWirelineCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData) SetRgWirelineCharacteristics(v string)`

SetRgWirelineCharacteristics sets RgWirelineCharacteristics field to given value.

### HasRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData) HasRgWirelineCharacteristics() bool`

HasRgWirelineCharacteristics returns a boolean if a field has been set.

### GetEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData) GetEcRestrictionDataWb() EcRestrictionDataWb`

GetEcRestrictionDataWb returns the EcRestrictionDataWb field if non-nil, zero value otherwise.

### GetEcRestrictionDataWbOk

`func (o *AccessAndMobilitySubscriptionData) GetEcRestrictionDataWbOk() (*EcRestrictionDataWb, bool)`

GetEcRestrictionDataWbOk returns a tuple with the EcRestrictionDataWb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData) SetEcRestrictionDataWb(v EcRestrictionDataWb)`

SetEcRestrictionDataWb sets EcRestrictionDataWb field to given value.

### HasEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData) HasEcRestrictionDataWb() bool`

HasEcRestrictionDataWb returns a boolean if a field has been set.

### GetEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData) GetEcRestrictionDataNb() bool`

GetEcRestrictionDataNb returns the EcRestrictionDataNb field if non-nil, zero value otherwise.

### GetEcRestrictionDataNbOk

`func (o *AccessAndMobilitySubscriptionData) GetEcRestrictionDataNbOk() (*bool, bool)`

GetEcRestrictionDataNbOk returns a tuple with the EcRestrictionDataNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData) SetEcRestrictionDataNb(v bool)`

SetEcRestrictionDataNb sets EcRestrictionDataNb field to given value.

### HasEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData) HasEcRestrictionDataNb() bool`

HasEcRestrictionDataNb returns a boolean if a field has been set.

### GetExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData) GetExpectedUeBehaviourList() ExpectedUeBehaviourData`

GetExpectedUeBehaviourList returns the ExpectedUeBehaviourList field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourListOk

`func (o *AccessAndMobilitySubscriptionData) GetExpectedUeBehaviourListOk() (*ExpectedUeBehaviourData, bool)`

GetExpectedUeBehaviourListOk returns a tuple with the ExpectedUeBehaviourList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData) SetExpectedUeBehaviourList(v ExpectedUeBehaviourData)`

SetExpectedUeBehaviourList sets ExpectedUeBehaviourList field to given value.

### HasExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData) HasExpectedUeBehaviourList() bool`

HasExpectedUeBehaviourList returns a boolean if a field has been set.

### GetPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) GetPrimaryRatRestrictions() []RatType`

GetPrimaryRatRestrictions returns the PrimaryRatRestrictions field if non-nil, zero value otherwise.

### GetPrimaryRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData) GetPrimaryRatRestrictionsOk() (*[]RatType, bool)`

GetPrimaryRatRestrictionsOk returns a tuple with the PrimaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) SetPrimaryRatRestrictions(v []RatType)`

SetPrimaryRatRestrictions sets PrimaryRatRestrictions field to given value.

### HasPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) HasPrimaryRatRestrictions() bool`

HasPrimaryRatRestrictions returns a boolean if a field has been set.

### GetSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) GetSecondaryRatRestrictions() []RatType`

GetSecondaryRatRestrictions returns the SecondaryRatRestrictions field if non-nil, zero value otherwise.

### GetSecondaryRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData) GetSecondaryRatRestrictionsOk() (*[]RatType, bool)`

GetSecondaryRatRestrictionsOk returns a tuple with the SecondaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) SetSecondaryRatRestrictions(v []RatType)`

SetSecondaryRatRestrictions sets SecondaryRatRestrictions field to given value.

### HasSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData) HasSecondaryRatRestrictions() bool`

HasSecondaryRatRestrictions returns a boolean if a field has been set.

### GetEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData) GetEdrxParametersList() []EdrxParameters`

GetEdrxParametersList returns the EdrxParametersList field if non-nil, zero value otherwise.

### GetEdrxParametersListOk

`func (o *AccessAndMobilitySubscriptionData) GetEdrxParametersListOk() (*[]EdrxParameters, bool)`

GetEdrxParametersListOk returns a tuple with the EdrxParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData) SetEdrxParametersList(v []EdrxParameters)`

SetEdrxParametersList sets EdrxParametersList field to given value.

### HasEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData) HasEdrxParametersList() bool`

HasEdrxParametersList returns a boolean if a field has been set.

### GetPtwParametersList

`func (o *AccessAndMobilitySubscriptionData) GetPtwParametersList() []PtwParameters`

GetPtwParametersList returns the PtwParametersList field if non-nil, zero value otherwise.

### GetPtwParametersListOk

`func (o *AccessAndMobilitySubscriptionData) GetPtwParametersListOk() (*[]PtwParameters, bool)`

GetPtwParametersListOk returns a tuple with the PtwParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtwParametersList

`func (o *AccessAndMobilitySubscriptionData) SetPtwParametersList(v []PtwParameters)`

SetPtwParametersList sets PtwParametersList field to given value.

### HasPtwParametersList

`func (o *AccessAndMobilitySubscriptionData) HasPtwParametersList() bool`

HasPtwParametersList returns a boolean if a field has been set.

### GetIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData) GetIabOperationAllowed() bool`

GetIabOperationAllowed returns the IabOperationAllowed field if non-nil, zero value otherwise.

### GetIabOperationAllowedOk

`func (o *AccessAndMobilitySubscriptionData) GetIabOperationAllowedOk() (*bool, bool)`

GetIabOperationAllowedOk returns a tuple with the IabOperationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData) SetIabOperationAllowed(v bool)`

SetIabOperationAllowed sets IabOperationAllowed field to given value.

### HasIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData) HasIabOperationAllowed() bool`

HasIabOperationAllowed returns a boolean if a field has been set.

### GetWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) GetWirelineForbiddenAreas() []WirelineArea`

GetWirelineForbiddenAreas returns the WirelineForbiddenAreas field if non-nil, zero value otherwise.

### GetWirelineForbiddenAreasOk

`func (o *AccessAndMobilitySubscriptionData) GetWirelineForbiddenAreasOk() (*[]WirelineArea, bool)`

GetWirelineForbiddenAreasOk returns a tuple with the WirelineForbiddenAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) SetWirelineForbiddenAreas(v []WirelineArea)`

SetWirelineForbiddenAreas sets WirelineForbiddenAreas field to given value.

### HasWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData) HasWirelineForbiddenAreas() bool`

HasWirelineForbiddenAreas returns a boolean if a field has been set.

### GetWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) GetWirelineServiceAreaRestriction() WirelineServiceAreaRestriction`

GetWirelineServiceAreaRestriction returns the WirelineServiceAreaRestriction field if non-nil, zero value otherwise.

### GetWirelineServiceAreaRestrictionOk

`func (o *AccessAndMobilitySubscriptionData) GetWirelineServiceAreaRestrictionOk() (*WirelineServiceAreaRestriction, bool)`

GetWirelineServiceAreaRestrictionOk returns a tuple with the WirelineServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) SetWirelineServiceAreaRestriction(v WirelineServiceAreaRestriction)`

SetWirelineServiceAreaRestriction sets WirelineServiceAreaRestriction field to given value.

### HasWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData) HasWirelineServiceAreaRestriction() bool`

HasWirelineServiceAreaRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


