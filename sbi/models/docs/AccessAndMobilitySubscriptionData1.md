# AccessAndMobilitySubscriptionData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**InternalGroupIds** | Pointer to **[]string** |  | [optional] 
**SharedVnGroupDataIds** | Pointer to **map[string]string** |  | [optional] 
**SubscribedUeAmbr** | Pointer to [**AmbrRm**](AmbrRm.md) |  | [optional] 
**Nssai** | Pointer to [**Nssai1**](Nssai1.md) |  | [optional] 
**RatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**ForbiddenAreas** | Pointer to [**[]Area1**](Area1.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**ServiceAreaRestriction1**](ServiceAreaRestriction1.md) |  | [optional] 
**CoreNetworkTypeRestrictions** | Pointer to [**[]CoreNetworkType**](CoreNetworkType.md) |  | [optional] 
**RfspIndex** | Pointer to **int32** |  | [optional] 
**SubsRegTimer** | Pointer to **int32** |  | [optional] 
**UeUsageType** | Pointer to **int32** |  | [optional] 
**MpsPriority** | Pointer to **bool** |  | [optional] 
**McsPriority** | Pointer to **bool** |  | [optional] 
**ActiveTime** | Pointer to **int32** |  | [optional] 
**SorInfo** | Pointer to [**SorInfo1**](SorInfo1.md) |  | [optional] 
**SorInfoExpectInd** | Pointer to **bool** |  | [optional] 
**SorafRetrieval** | Pointer to **bool** |  | [optional] [default to false]
**SorUpdateIndicatorList** | Pointer to [**[]SorUpdateIndicator**](SorUpdateIndicator.md) |  | [optional] 
**UpuInfo** | Pointer to [**UpuInfo1**](UpuInfo1.md) |  | [optional] 
**MicoAllowed** | Pointer to **bool** |  | [optional] 
**SharedAmDataIds** | Pointer to **[]string** |  | [optional] 
**OdbPacketServices** | Pointer to [**OdbPacketServices**](OdbPacketServices.md) |  | [optional] 
**SubscribedDnnList** | Pointer to [**[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner**](AccessAndMobilitySubscriptionDataSubscribedDnnListInner.md) |  | [optional] 
**ServiceGapTime** | Pointer to **int32** |  | [optional] 
**MdtUserConsent** | Pointer to [**MdtUserConsent**](MdtUserConsent.md) |  | [optional] 
**MdtConfiguration** | Pointer to [**MdtConfiguration1**](MdtConfiguration1.md) |  | [optional] 
**TraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 
**CagData** | Pointer to [**CagData1**](CagData1.md) |  | [optional] 
**StnSr** | Pointer to **string** |  | [optional] 
**CMsisdn** | Pointer to **string** |  | [optional] 
**NbIoTUePriority** | Pointer to **int32** |  | [optional] 
**NssaiInclusionAllowed** | Pointer to **bool** |  | [optional] [default to false]
**RgWirelineCharacteristics** | Pointer to **string** |  | [optional] 
**EcRestrictionDataWb** | Pointer to [**EcRestrictionDataWb**](EcRestrictionDataWb.md) |  | [optional] 
**EcRestrictionDataNb** | Pointer to **bool** |  | [optional] [default to false]
**ExpectedUeBehaviourList** | Pointer to [**ExpectedUeBehaviourData1**](ExpectedUeBehaviourData1.md) |  | [optional] 
**PrimaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**SecondaryRatRestrictions** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**EdrxParametersList** | Pointer to [**[]EdrxParameters1**](EdrxParameters1.md) |  | [optional] 
**PtwParametersList** | Pointer to [**[]PtwParameters1**](PtwParameters1.md) |  | [optional] 
**IabOperationAllowed** | Pointer to **bool** |  | [optional] [default to false]
**WirelineForbiddenAreas** | Pointer to [**[]WirelineArea1**](WirelineArea1.md) |  | [optional] 
**WirelineServiceAreaRestriction** | Pointer to [**WirelineServiceAreaRestriction1**](WirelineServiceAreaRestriction1.md) |  | [optional] 

## Methods

### NewAccessAndMobilitySubscriptionData1

`func NewAccessAndMobilitySubscriptionData1() *AccessAndMobilitySubscriptionData1`

NewAccessAndMobilitySubscriptionData1 instantiates a new AccessAndMobilitySubscriptionData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessAndMobilitySubscriptionData1WithDefaults

`func NewAccessAndMobilitySubscriptionData1WithDefaults() *AccessAndMobilitySubscriptionData1`

NewAccessAndMobilitySubscriptionData1WithDefaults instantiates a new AccessAndMobilitySubscriptionData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData1) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AccessAndMobilitySubscriptionData1) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData1) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AccessAndMobilitySubscriptionData1) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetGpsis

`func (o *AccessAndMobilitySubscriptionData1) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *AccessAndMobilitySubscriptionData1) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *AccessAndMobilitySubscriptionData1) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *AccessAndMobilitySubscriptionData1) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData1) GetInternalGroupIds() []string`

GetInternalGroupIds returns the InternalGroupIds field if non-nil, zero value otherwise.

### GetInternalGroupIdsOk

`func (o *AccessAndMobilitySubscriptionData1) GetInternalGroupIdsOk() (*[]string, bool)`

GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData1) SetInternalGroupIds(v []string)`

SetInternalGroupIds sets InternalGroupIds field to given value.

### HasInternalGroupIds

`func (o *AccessAndMobilitySubscriptionData1) HasInternalGroupIds() bool`

HasInternalGroupIds returns a boolean if a field has been set.

### GetSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData1) GetSharedVnGroupDataIds() map[string]string`

GetSharedVnGroupDataIds returns the SharedVnGroupDataIds field if non-nil, zero value otherwise.

### GetSharedVnGroupDataIdsOk

`func (o *AccessAndMobilitySubscriptionData1) GetSharedVnGroupDataIdsOk() (*map[string]string, bool)`

GetSharedVnGroupDataIdsOk returns a tuple with the SharedVnGroupDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData1) SetSharedVnGroupDataIds(v map[string]string)`

SetSharedVnGroupDataIds sets SharedVnGroupDataIds field to given value.

### HasSharedVnGroupDataIds

`func (o *AccessAndMobilitySubscriptionData1) HasSharedVnGroupDataIds() bool`

HasSharedVnGroupDataIds returns a boolean if a field has been set.

### GetSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData1) GetSubscribedUeAmbr() AmbrRm`

GetSubscribedUeAmbr returns the SubscribedUeAmbr field if non-nil, zero value otherwise.

### GetSubscribedUeAmbrOk

`func (o *AccessAndMobilitySubscriptionData1) GetSubscribedUeAmbrOk() (*AmbrRm, bool)`

GetSubscribedUeAmbrOk returns a tuple with the SubscribedUeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData1) SetSubscribedUeAmbr(v AmbrRm)`

SetSubscribedUeAmbr sets SubscribedUeAmbr field to given value.

### HasSubscribedUeAmbr

`func (o *AccessAndMobilitySubscriptionData1) HasSubscribedUeAmbr() bool`

HasSubscribedUeAmbr returns a boolean if a field has been set.

### GetNssai

`func (o *AccessAndMobilitySubscriptionData1) GetNssai() Nssai1`

GetNssai returns the Nssai field if non-nil, zero value otherwise.

### GetNssaiOk

`func (o *AccessAndMobilitySubscriptionData1) GetNssaiOk() (*Nssai1, bool)`

GetNssaiOk returns a tuple with the Nssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssai

`func (o *AccessAndMobilitySubscriptionData1) SetNssai(v Nssai1)`

SetNssai sets Nssai field to given value.

### HasNssai

`func (o *AccessAndMobilitySubscriptionData1) HasNssai() bool`

HasNssai returns a boolean if a field has been set.

### SetNssaiNil

`func (o *AccessAndMobilitySubscriptionData1) SetNssaiNil(b bool)`

 SetNssaiNil sets the value for Nssai to be an explicit nil

### UnsetNssai
`func (o *AccessAndMobilitySubscriptionData1) UnsetNssai()`

UnsetNssai ensures that no value is present for Nssai, not even an explicit nil
### GetRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) GetRatRestrictions() []RatType`

GetRatRestrictions returns the RatRestrictions field if non-nil, zero value otherwise.

### GetRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData1) GetRatRestrictionsOk() (*[]RatType, bool)`

GetRatRestrictionsOk returns a tuple with the RatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) SetRatRestrictions(v []RatType)`

SetRatRestrictions sets RatRestrictions field to given value.

### HasRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) HasRatRestrictions() bool`

HasRatRestrictions returns a boolean if a field has been set.

### GetForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) GetForbiddenAreas() []Area1`

GetForbiddenAreas returns the ForbiddenAreas field if non-nil, zero value otherwise.

### GetForbiddenAreasOk

`func (o *AccessAndMobilitySubscriptionData1) GetForbiddenAreasOk() (*[]Area1, bool)`

GetForbiddenAreasOk returns a tuple with the ForbiddenAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) SetForbiddenAreas(v []Area1)`

SetForbiddenAreas sets ForbiddenAreas field to given value.

### HasForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) HasForbiddenAreas() bool`

HasForbiddenAreas returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) GetServiceAreaRestriction() ServiceAreaRestriction1`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *AccessAndMobilitySubscriptionData1) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction1, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) SetServiceAreaRestriction(v ServiceAreaRestriction1)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData1) GetCoreNetworkTypeRestrictions() []CoreNetworkType`

GetCoreNetworkTypeRestrictions returns the CoreNetworkTypeRestrictions field if non-nil, zero value otherwise.

### GetCoreNetworkTypeRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData1) GetCoreNetworkTypeRestrictionsOk() (*[]CoreNetworkType, bool)`

GetCoreNetworkTypeRestrictionsOk returns a tuple with the CoreNetworkTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData1) SetCoreNetworkTypeRestrictions(v []CoreNetworkType)`

SetCoreNetworkTypeRestrictions sets CoreNetworkTypeRestrictions field to given value.

### HasCoreNetworkTypeRestrictions

`func (o *AccessAndMobilitySubscriptionData1) HasCoreNetworkTypeRestrictions() bool`

HasCoreNetworkTypeRestrictions returns a boolean if a field has been set.

### GetRfspIndex

`func (o *AccessAndMobilitySubscriptionData1) GetRfspIndex() int32`

GetRfspIndex returns the RfspIndex field if non-nil, zero value otherwise.

### GetRfspIndexOk

`func (o *AccessAndMobilitySubscriptionData1) GetRfspIndexOk() (*int32, bool)`

GetRfspIndexOk returns a tuple with the RfspIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfspIndex

`func (o *AccessAndMobilitySubscriptionData1) SetRfspIndex(v int32)`

SetRfspIndex sets RfspIndex field to given value.

### HasRfspIndex

`func (o *AccessAndMobilitySubscriptionData1) HasRfspIndex() bool`

HasRfspIndex returns a boolean if a field has been set.

### SetRfspIndexNil

`func (o *AccessAndMobilitySubscriptionData1) SetRfspIndexNil(b bool)`

 SetRfspIndexNil sets the value for RfspIndex to be an explicit nil

### UnsetRfspIndex
`func (o *AccessAndMobilitySubscriptionData1) UnsetRfspIndex()`

UnsetRfspIndex ensures that no value is present for RfspIndex, not even an explicit nil
### GetSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData1) GetSubsRegTimer() int32`

GetSubsRegTimer returns the SubsRegTimer field if non-nil, zero value otherwise.

### GetSubsRegTimerOk

`func (o *AccessAndMobilitySubscriptionData1) GetSubsRegTimerOk() (*int32, bool)`

GetSubsRegTimerOk returns a tuple with the SubsRegTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData1) SetSubsRegTimer(v int32)`

SetSubsRegTimer sets SubsRegTimer field to given value.

### HasSubsRegTimer

`func (o *AccessAndMobilitySubscriptionData1) HasSubsRegTimer() bool`

HasSubsRegTimer returns a boolean if a field has been set.

### SetSubsRegTimerNil

`func (o *AccessAndMobilitySubscriptionData1) SetSubsRegTimerNil(b bool)`

 SetSubsRegTimerNil sets the value for SubsRegTimer to be an explicit nil

### UnsetSubsRegTimer
`func (o *AccessAndMobilitySubscriptionData1) UnsetSubsRegTimer()`

UnsetSubsRegTimer ensures that no value is present for SubsRegTimer, not even an explicit nil
### GetUeUsageType

`func (o *AccessAndMobilitySubscriptionData1) GetUeUsageType() int32`

GetUeUsageType returns the UeUsageType field if non-nil, zero value otherwise.

### GetUeUsageTypeOk

`func (o *AccessAndMobilitySubscriptionData1) GetUeUsageTypeOk() (*int32, bool)`

GetUeUsageTypeOk returns a tuple with the UeUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeUsageType

`func (o *AccessAndMobilitySubscriptionData1) SetUeUsageType(v int32)`

SetUeUsageType sets UeUsageType field to given value.

### HasUeUsageType

`func (o *AccessAndMobilitySubscriptionData1) HasUeUsageType() bool`

HasUeUsageType returns a boolean if a field has been set.

### GetMpsPriority

`func (o *AccessAndMobilitySubscriptionData1) GetMpsPriority() bool`

GetMpsPriority returns the MpsPriority field if non-nil, zero value otherwise.

### GetMpsPriorityOk

`func (o *AccessAndMobilitySubscriptionData1) GetMpsPriorityOk() (*bool, bool)`

GetMpsPriorityOk returns a tuple with the MpsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsPriority

`func (o *AccessAndMobilitySubscriptionData1) SetMpsPriority(v bool)`

SetMpsPriority sets MpsPriority field to given value.

### HasMpsPriority

`func (o *AccessAndMobilitySubscriptionData1) HasMpsPriority() bool`

HasMpsPriority returns a boolean if a field has been set.

### GetMcsPriority

`func (o *AccessAndMobilitySubscriptionData1) GetMcsPriority() bool`

GetMcsPriority returns the McsPriority field if non-nil, zero value otherwise.

### GetMcsPriorityOk

`func (o *AccessAndMobilitySubscriptionData1) GetMcsPriorityOk() (*bool, bool)`

GetMcsPriorityOk returns a tuple with the McsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsPriority

`func (o *AccessAndMobilitySubscriptionData1) SetMcsPriority(v bool)`

SetMcsPriority sets McsPriority field to given value.

### HasMcsPriority

`func (o *AccessAndMobilitySubscriptionData1) HasMcsPriority() bool`

HasMcsPriority returns a boolean if a field has been set.

### GetActiveTime

`func (o *AccessAndMobilitySubscriptionData1) GetActiveTime() int32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *AccessAndMobilitySubscriptionData1) GetActiveTimeOk() (*int32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *AccessAndMobilitySubscriptionData1) SetActiveTime(v int32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *AccessAndMobilitySubscriptionData1) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### SetActiveTimeNil

`func (o *AccessAndMobilitySubscriptionData1) SetActiveTimeNil(b bool)`

 SetActiveTimeNil sets the value for ActiveTime to be an explicit nil

### UnsetActiveTime
`func (o *AccessAndMobilitySubscriptionData1) UnsetActiveTime()`

UnsetActiveTime ensures that no value is present for ActiveTime, not even an explicit nil
### GetSorInfo

`func (o *AccessAndMobilitySubscriptionData1) GetSorInfo() SorInfo1`

GetSorInfo returns the SorInfo field if non-nil, zero value otherwise.

### GetSorInfoOk

`func (o *AccessAndMobilitySubscriptionData1) GetSorInfoOk() (*SorInfo1, bool)`

GetSorInfoOk returns a tuple with the SorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfo

`func (o *AccessAndMobilitySubscriptionData1) SetSorInfo(v SorInfo1)`

SetSorInfo sets SorInfo field to given value.

### HasSorInfo

`func (o *AccessAndMobilitySubscriptionData1) HasSorInfo() bool`

HasSorInfo returns a boolean if a field has been set.

### GetSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData1) GetSorInfoExpectInd() bool`

GetSorInfoExpectInd returns the SorInfoExpectInd field if non-nil, zero value otherwise.

### GetSorInfoExpectIndOk

`func (o *AccessAndMobilitySubscriptionData1) GetSorInfoExpectIndOk() (*bool, bool)`

GetSorInfoExpectIndOk returns a tuple with the SorInfoExpectInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData1) SetSorInfoExpectInd(v bool)`

SetSorInfoExpectInd sets SorInfoExpectInd field to given value.

### HasSorInfoExpectInd

`func (o *AccessAndMobilitySubscriptionData1) HasSorInfoExpectInd() bool`

HasSorInfoExpectInd returns a boolean if a field has been set.

### GetSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData1) GetSorafRetrieval() bool`

GetSorafRetrieval returns the SorafRetrieval field if non-nil, zero value otherwise.

### GetSorafRetrievalOk

`func (o *AccessAndMobilitySubscriptionData1) GetSorafRetrievalOk() (*bool, bool)`

GetSorafRetrievalOk returns a tuple with the SorafRetrieval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData1) SetSorafRetrieval(v bool)`

SetSorafRetrieval sets SorafRetrieval field to given value.

### HasSorafRetrieval

`func (o *AccessAndMobilitySubscriptionData1) HasSorafRetrieval() bool`

HasSorafRetrieval returns a boolean if a field has been set.

### GetSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData1) GetSorUpdateIndicatorList() []SorUpdateIndicator`

GetSorUpdateIndicatorList returns the SorUpdateIndicatorList field if non-nil, zero value otherwise.

### GetSorUpdateIndicatorListOk

`func (o *AccessAndMobilitySubscriptionData1) GetSorUpdateIndicatorListOk() (*[]SorUpdateIndicator, bool)`

GetSorUpdateIndicatorListOk returns a tuple with the SorUpdateIndicatorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData1) SetSorUpdateIndicatorList(v []SorUpdateIndicator)`

SetSorUpdateIndicatorList sets SorUpdateIndicatorList field to given value.

### HasSorUpdateIndicatorList

`func (o *AccessAndMobilitySubscriptionData1) HasSorUpdateIndicatorList() bool`

HasSorUpdateIndicatorList returns a boolean if a field has been set.

### GetUpuInfo

`func (o *AccessAndMobilitySubscriptionData1) GetUpuInfo() UpuInfo1`

GetUpuInfo returns the UpuInfo field if non-nil, zero value otherwise.

### GetUpuInfoOk

`func (o *AccessAndMobilitySubscriptionData1) GetUpuInfoOk() (*UpuInfo1, bool)`

GetUpuInfoOk returns a tuple with the UpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuInfo

`func (o *AccessAndMobilitySubscriptionData1) SetUpuInfo(v UpuInfo1)`

SetUpuInfo sets UpuInfo field to given value.

### HasUpuInfo

`func (o *AccessAndMobilitySubscriptionData1) HasUpuInfo() bool`

HasUpuInfo returns a boolean if a field has been set.

### GetMicoAllowed

`func (o *AccessAndMobilitySubscriptionData1) GetMicoAllowed() bool`

GetMicoAllowed returns the MicoAllowed field if non-nil, zero value otherwise.

### GetMicoAllowedOk

`func (o *AccessAndMobilitySubscriptionData1) GetMicoAllowedOk() (*bool, bool)`

GetMicoAllowedOk returns a tuple with the MicoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicoAllowed

`func (o *AccessAndMobilitySubscriptionData1) SetMicoAllowed(v bool)`

SetMicoAllowed sets MicoAllowed field to given value.

### HasMicoAllowed

`func (o *AccessAndMobilitySubscriptionData1) HasMicoAllowed() bool`

HasMicoAllowed returns a boolean if a field has been set.

### GetSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData1) GetSharedAmDataIds() []string`

GetSharedAmDataIds returns the SharedAmDataIds field if non-nil, zero value otherwise.

### GetSharedAmDataIdsOk

`func (o *AccessAndMobilitySubscriptionData1) GetSharedAmDataIdsOk() (*[]string, bool)`

GetSharedAmDataIdsOk returns a tuple with the SharedAmDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData1) SetSharedAmDataIds(v []string)`

SetSharedAmDataIds sets SharedAmDataIds field to given value.

### HasSharedAmDataIds

`func (o *AccessAndMobilitySubscriptionData1) HasSharedAmDataIds() bool`

HasSharedAmDataIds returns a boolean if a field has been set.

### GetOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData1) GetOdbPacketServices() OdbPacketServices`

GetOdbPacketServices returns the OdbPacketServices field if non-nil, zero value otherwise.

### GetOdbPacketServicesOk

`func (o *AccessAndMobilitySubscriptionData1) GetOdbPacketServicesOk() (*OdbPacketServices, bool)`

GetOdbPacketServicesOk returns a tuple with the OdbPacketServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData1) SetOdbPacketServices(v OdbPacketServices)`

SetOdbPacketServices sets OdbPacketServices field to given value.

### HasOdbPacketServices

`func (o *AccessAndMobilitySubscriptionData1) HasOdbPacketServices() bool`

HasOdbPacketServices returns a boolean if a field has been set.

### GetSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData1) GetSubscribedDnnList() []AccessAndMobilitySubscriptionDataSubscribedDnnListInner`

GetSubscribedDnnList returns the SubscribedDnnList field if non-nil, zero value otherwise.

### GetSubscribedDnnListOk

`func (o *AccessAndMobilitySubscriptionData1) GetSubscribedDnnListOk() (*[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool)`

GetSubscribedDnnListOk returns a tuple with the SubscribedDnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData1) SetSubscribedDnnList(v []AccessAndMobilitySubscriptionDataSubscribedDnnListInner)`

SetSubscribedDnnList sets SubscribedDnnList field to given value.

### HasSubscribedDnnList

`func (o *AccessAndMobilitySubscriptionData1) HasSubscribedDnnList() bool`

HasSubscribedDnnList returns a boolean if a field has been set.

### GetServiceGapTime

`func (o *AccessAndMobilitySubscriptionData1) GetServiceGapTime() int32`

GetServiceGapTime returns the ServiceGapTime field if non-nil, zero value otherwise.

### GetServiceGapTimeOk

`func (o *AccessAndMobilitySubscriptionData1) GetServiceGapTimeOk() (*int32, bool)`

GetServiceGapTimeOk returns a tuple with the ServiceGapTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGapTime

`func (o *AccessAndMobilitySubscriptionData1) SetServiceGapTime(v int32)`

SetServiceGapTime sets ServiceGapTime field to given value.

### HasServiceGapTime

`func (o *AccessAndMobilitySubscriptionData1) HasServiceGapTime() bool`

HasServiceGapTime returns a boolean if a field has been set.

### GetMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData1) GetMdtUserConsent() MdtUserConsent`

GetMdtUserConsent returns the MdtUserConsent field if non-nil, zero value otherwise.

### GetMdtUserConsentOk

`func (o *AccessAndMobilitySubscriptionData1) GetMdtUserConsentOk() (*MdtUserConsent, bool)`

GetMdtUserConsentOk returns a tuple with the MdtUserConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData1) SetMdtUserConsent(v MdtUserConsent)`

SetMdtUserConsent sets MdtUserConsent field to given value.

### HasMdtUserConsent

`func (o *AccessAndMobilitySubscriptionData1) HasMdtUserConsent() bool`

HasMdtUserConsent returns a boolean if a field has been set.

### GetMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData1) GetMdtConfiguration() MdtConfiguration1`

GetMdtConfiguration returns the MdtConfiguration field if non-nil, zero value otherwise.

### GetMdtConfigurationOk

`func (o *AccessAndMobilitySubscriptionData1) GetMdtConfigurationOk() (*MdtConfiguration1, bool)`

GetMdtConfigurationOk returns a tuple with the MdtConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData1) SetMdtConfiguration(v MdtConfiguration1)`

SetMdtConfiguration sets MdtConfiguration field to given value.

### HasMdtConfiguration

`func (o *AccessAndMobilitySubscriptionData1) HasMdtConfiguration() bool`

HasMdtConfiguration returns a boolean if a field has been set.

### GetTraceData

`func (o *AccessAndMobilitySubscriptionData1) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *AccessAndMobilitySubscriptionData1) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *AccessAndMobilitySubscriptionData1) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *AccessAndMobilitySubscriptionData1) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *AccessAndMobilitySubscriptionData1) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *AccessAndMobilitySubscriptionData1) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetCagData

`func (o *AccessAndMobilitySubscriptionData1) GetCagData() CagData1`

GetCagData returns the CagData field if non-nil, zero value otherwise.

### GetCagDataOk

`func (o *AccessAndMobilitySubscriptionData1) GetCagDataOk() (*CagData1, bool)`

GetCagDataOk returns a tuple with the CagData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagData

`func (o *AccessAndMobilitySubscriptionData1) SetCagData(v CagData1)`

SetCagData sets CagData field to given value.

### HasCagData

`func (o *AccessAndMobilitySubscriptionData1) HasCagData() bool`

HasCagData returns a boolean if a field has been set.

### GetStnSr

`func (o *AccessAndMobilitySubscriptionData1) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *AccessAndMobilitySubscriptionData1) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *AccessAndMobilitySubscriptionData1) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *AccessAndMobilitySubscriptionData1) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### GetCMsisdn

`func (o *AccessAndMobilitySubscriptionData1) GetCMsisdn() string`

GetCMsisdn returns the CMsisdn field if non-nil, zero value otherwise.

### GetCMsisdnOk

`func (o *AccessAndMobilitySubscriptionData1) GetCMsisdnOk() (*string, bool)`

GetCMsisdnOk returns a tuple with the CMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCMsisdn

`func (o *AccessAndMobilitySubscriptionData1) SetCMsisdn(v string)`

SetCMsisdn sets CMsisdn field to given value.

### HasCMsisdn

`func (o *AccessAndMobilitySubscriptionData1) HasCMsisdn() bool`

HasCMsisdn returns a boolean if a field has been set.

### GetNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData1) GetNbIoTUePriority() int32`

GetNbIoTUePriority returns the NbIoTUePriority field if non-nil, zero value otherwise.

### GetNbIoTUePriorityOk

`func (o *AccessAndMobilitySubscriptionData1) GetNbIoTUePriorityOk() (*int32, bool)`

GetNbIoTUePriorityOk returns a tuple with the NbIoTUePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData1) SetNbIoTUePriority(v int32)`

SetNbIoTUePriority sets NbIoTUePriority field to given value.

### HasNbIoTUePriority

`func (o *AccessAndMobilitySubscriptionData1) HasNbIoTUePriority() bool`

HasNbIoTUePriority returns a boolean if a field has been set.

### GetNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData1) GetNssaiInclusionAllowed() bool`

GetNssaiInclusionAllowed returns the NssaiInclusionAllowed field if non-nil, zero value otherwise.

### GetNssaiInclusionAllowedOk

`func (o *AccessAndMobilitySubscriptionData1) GetNssaiInclusionAllowedOk() (*bool, bool)`

GetNssaiInclusionAllowedOk returns a tuple with the NssaiInclusionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData1) SetNssaiInclusionAllowed(v bool)`

SetNssaiInclusionAllowed sets NssaiInclusionAllowed field to given value.

### HasNssaiInclusionAllowed

`func (o *AccessAndMobilitySubscriptionData1) HasNssaiInclusionAllowed() bool`

HasNssaiInclusionAllowed returns a boolean if a field has been set.

### GetRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData1) GetRgWirelineCharacteristics() string`

GetRgWirelineCharacteristics returns the RgWirelineCharacteristics field if non-nil, zero value otherwise.

### GetRgWirelineCharacteristicsOk

`func (o *AccessAndMobilitySubscriptionData1) GetRgWirelineCharacteristicsOk() (*string, bool)`

GetRgWirelineCharacteristicsOk returns a tuple with the RgWirelineCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData1) SetRgWirelineCharacteristics(v string)`

SetRgWirelineCharacteristics sets RgWirelineCharacteristics field to given value.

### HasRgWirelineCharacteristics

`func (o *AccessAndMobilitySubscriptionData1) HasRgWirelineCharacteristics() bool`

HasRgWirelineCharacteristics returns a boolean if a field has been set.

### GetEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData1) GetEcRestrictionDataWb() EcRestrictionDataWb`

GetEcRestrictionDataWb returns the EcRestrictionDataWb field if non-nil, zero value otherwise.

### GetEcRestrictionDataWbOk

`func (o *AccessAndMobilitySubscriptionData1) GetEcRestrictionDataWbOk() (*EcRestrictionDataWb, bool)`

GetEcRestrictionDataWbOk returns a tuple with the EcRestrictionDataWb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData1) SetEcRestrictionDataWb(v EcRestrictionDataWb)`

SetEcRestrictionDataWb sets EcRestrictionDataWb field to given value.

### HasEcRestrictionDataWb

`func (o *AccessAndMobilitySubscriptionData1) HasEcRestrictionDataWb() bool`

HasEcRestrictionDataWb returns a boolean if a field has been set.

### GetEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData1) GetEcRestrictionDataNb() bool`

GetEcRestrictionDataNb returns the EcRestrictionDataNb field if non-nil, zero value otherwise.

### GetEcRestrictionDataNbOk

`func (o *AccessAndMobilitySubscriptionData1) GetEcRestrictionDataNbOk() (*bool, bool)`

GetEcRestrictionDataNbOk returns a tuple with the EcRestrictionDataNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData1) SetEcRestrictionDataNb(v bool)`

SetEcRestrictionDataNb sets EcRestrictionDataNb field to given value.

### HasEcRestrictionDataNb

`func (o *AccessAndMobilitySubscriptionData1) HasEcRestrictionDataNb() bool`

HasEcRestrictionDataNb returns a boolean if a field has been set.

### GetExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData1) GetExpectedUeBehaviourList() ExpectedUeBehaviourData1`

GetExpectedUeBehaviourList returns the ExpectedUeBehaviourList field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourListOk

`func (o *AccessAndMobilitySubscriptionData1) GetExpectedUeBehaviourListOk() (*ExpectedUeBehaviourData1, bool)`

GetExpectedUeBehaviourListOk returns a tuple with the ExpectedUeBehaviourList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData1) SetExpectedUeBehaviourList(v ExpectedUeBehaviourData1)`

SetExpectedUeBehaviourList sets ExpectedUeBehaviourList field to given value.

### HasExpectedUeBehaviourList

`func (o *AccessAndMobilitySubscriptionData1) HasExpectedUeBehaviourList() bool`

HasExpectedUeBehaviourList returns a boolean if a field has been set.

### GetPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) GetPrimaryRatRestrictions() []RatType`

GetPrimaryRatRestrictions returns the PrimaryRatRestrictions field if non-nil, zero value otherwise.

### GetPrimaryRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData1) GetPrimaryRatRestrictionsOk() (*[]RatType, bool)`

GetPrimaryRatRestrictionsOk returns a tuple with the PrimaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) SetPrimaryRatRestrictions(v []RatType)`

SetPrimaryRatRestrictions sets PrimaryRatRestrictions field to given value.

### HasPrimaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) HasPrimaryRatRestrictions() bool`

HasPrimaryRatRestrictions returns a boolean if a field has been set.

### GetSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) GetSecondaryRatRestrictions() []RatType`

GetSecondaryRatRestrictions returns the SecondaryRatRestrictions field if non-nil, zero value otherwise.

### GetSecondaryRatRestrictionsOk

`func (o *AccessAndMobilitySubscriptionData1) GetSecondaryRatRestrictionsOk() (*[]RatType, bool)`

GetSecondaryRatRestrictionsOk returns a tuple with the SecondaryRatRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) SetSecondaryRatRestrictions(v []RatType)`

SetSecondaryRatRestrictions sets SecondaryRatRestrictions field to given value.

### HasSecondaryRatRestrictions

`func (o *AccessAndMobilitySubscriptionData1) HasSecondaryRatRestrictions() bool`

HasSecondaryRatRestrictions returns a boolean if a field has been set.

### GetEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData1) GetEdrxParametersList() []EdrxParameters1`

GetEdrxParametersList returns the EdrxParametersList field if non-nil, zero value otherwise.

### GetEdrxParametersListOk

`func (o *AccessAndMobilitySubscriptionData1) GetEdrxParametersListOk() (*[]EdrxParameters1, bool)`

GetEdrxParametersListOk returns a tuple with the EdrxParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData1) SetEdrxParametersList(v []EdrxParameters1)`

SetEdrxParametersList sets EdrxParametersList field to given value.

### HasEdrxParametersList

`func (o *AccessAndMobilitySubscriptionData1) HasEdrxParametersList() bool`

HasEdrxParametersList returns a boolean if a field has been set.

### GetPtwParametersList

`func (o *AccessAndMobilitySubscriptionData1) GetPtwParametersList() []PtwParameters1`

GetPtwParametersList returns the PtwParametersList field if non-nil, zero value otherwise.

### GetPtwParametersListOk

`func (o *AccessAndMobilitySubscriptionData1) GetPtwParametersListOk() (*[]PtwParameters1, bool)`

GetPtwParametersListOk returns a tuple with the PtwParametersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtwParametersList

`func (o *AccessAndMobilitySubscriptionData1) SetPtwParametersList(v []PtwParameters1)`

SetPtwParametersList sets PtwParametersList field to given value.

### HasPtwParametersList

`func (o *AccessAndMobilitySubscriptionData1) HasPtwParametersList() bool`

HasPtwParametersList returns a boolean if a field has been set.

### GetIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData1) GetIabOperationAllowed() bool`

GetIabOperationAllowed returns the IabOperationAllowed field if non-nil, zero value otherwise.

### GetIabOperationAllowedOk

`func (o *AccessAndMobilitySubscriptionData1) GetIabOperationAllowedOk() (*bool, bool)`

GetIabOperationAllowedOk returns a tuple with the IabOperationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData1) SetIabOperationAllowed(v bool)`

SetIabOperationAllowed sets IabOperationAllowed field to given value.

### HasIabOperationAllowed

`func (o *AccessAndMobilitySubscriptionData1) HasIabOperationAllowed() bool`

HasIabOperationAllowed returns a boolean if a field has been set.

### GetWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) GetWirelineForbiddenAreas() []WirelineArea1`

GetWirelineForbiddenAreas returns the WirelineForbiddenAreas field if non-nil, zero value otherwise.

### GetWirelineForbiddenAreasOk

`func (o *AccessAndMobilitySubscriptionData1) GetWirelineForbiddenAreasOk() (*[]WirelineArea1, bool)`

GetWirelineForbiddenAreasOk returns a tuple with the WirelineForbiddenAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) SetWirelineForbiddenAreas(v []WirelineArea1)`

SetWirelineForbiddenAreas sets WirelineForbiddenAreas field to given value.

### HasWirelineForbiddenAreas

`func (o *AccessAndMobilitySubscriptionData1) HasWirelineForbiddenAreas() bool`

HasWirelineForbiddenAreas returns a boolean if a field has been set.

### GetWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) GetWirelineServiceAreaRestriction() WirelineServiceAreaRestriction1`

GetWirelineServiceAreaRestriction returns the WirelineServiceAreaRestriction field if non-nil, zero value otherwise.

### GetWirelineServiceAreaRestrictionOk

`func (o *AccessAndMobilitySubscriptionData1) GetWirelineServiceAreaRestrictionOk() (*WirelineServiceAreaRestriction1, bool)`

GetWirelineServiceAreaRestrictionOk returns a tuple with the WirelineServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) SetWirelineServiceAreaRestriction(v WirelineServiceAreaRestriction1)`

SetWirelineServiceAreaRestriction sets WirelineServiceAreaRestriction field to given value.

### HasWirelineServiceAreaRestriction

`func (o *AccessAndMobilitySubscriptionData1) HasWirelineServiceAreaRestriction() bool`

HasWirelineServiceAreaRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


