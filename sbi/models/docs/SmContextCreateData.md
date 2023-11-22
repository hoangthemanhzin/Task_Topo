# SmContextCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** |  | [optional] 
**UnauthenticatedSupi** | Pointer to **bool** |  | [optional] [default to false]
**Pei** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**PduSessionId** | Pointer to **int32** |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**SelectedDnn** | Pointer to **string** |  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**HplmnSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ServingNfId** | **string** |  | 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**ServingNetwork** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**RequestType** | Pointer to [**RequestType**](RequestType.md) |  | [optional] 
**N1SmMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**AnType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PresenceInLadn** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**SmContextStatusUri** | **string** |  | 
**HSmfUri** | Pointer to **string** |  | [optional] 
**HSmfId** | Pointer to **string** |  | [optional] 
**SmfUri** | Pointer to **string** |  | [optional] 
**SmfId** | Pointer to **string** |  | [optional] 
**AdditionalHsmfUri** | Pointer to **[]string** |  | [optional] 
**AdditionalHsmfId** | Pointer to **[]string** |  | [optional] 
**AdditionalSmfUri** | Pointer to **[]string** |  | [optional] 
**AdditionalSmfId** | Pointer to **[]string** |  | [optional] 
**OldPduSessionId** | Pointer to **int32** |  | [optional] 
**PduSessionsActivateList** | Pointer to **[]int32** |  | [optional] 
**UeEpsPdnConnection** | Pointer to **string** |  | [optional] 
**HoState** | Pointer to [**HoState**](HoState.md) |  | [optional] 
**PcfId** | Pointer to **string** |  | [optional] 
**PcfGroupId** | Pointer to **string** |  | [optional] 
**PcfSetId** | Pointer to **string** |  | [optional] 
**NrfUri** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**SelMode** | Pointer to [**DnnSelectionMode**](DnnSelectionMode.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**TraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 
**UdmGroupId** | Pointer to **string** |  | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**EpsInterworkingInd** | Pointer to [**EpsInterworkingIndication**](EpsInterworkingIndication.md) |  | [optional] 
**IndirectForwardingFlag** | Pointer to **bool** |  | [optional] 
**DirectForwardingFlag** | Pointer to **bool** |  | [optional] 
**TargetId** | Pointer to [**NgRanTargetId**](NgRanTargetId.md) |  | [optional] 
**EpsBearerCtxStatus** | Pointer to **string** |  | [optional] 
**CpCiotEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CpOnlyInd** | Pointer to **bool** |  | [optional] [default to false]
**InvokeNef** | Pointer to **bool** |  | [optional] [default to false]
**MaRequestInd** | Pointer to **bool** |  | [optional] [default to false]
**MaNwUpgradeInd** | Pointer to **bool** |  | [optional] [default to false]
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**N2SmInfoExt1** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoTypeExt1** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**SmContextRef** | Pointer to **string** |  | [optional] 
**SmContextSmfId** | Pointer to **string** |  | [optional] 
**SmContextSmfSetId** | Pointer to **string** |  | [optional] 
**SmContextSmfServiceSetId** | Pointer to **string** |  | [optional] 
**SmContextSmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**ExtendedNasSmTimerInd** | Pointer to **bool** |  | [optional] [default to false]
**DlDataWaitingInd** | Pointer to **bool** |  | [optional] [default to false]
**DdnFailureSubs** | Pointer to [**DdnFailureSubs**](DdnFailureSubs.md) |  | [optional] 
**SmfTransferInd** | Pointer to **bool** |  | [optional] [default to false]
**OldSmfId** | Pointer to **string** |  | [optional] 
**OldSmContextRef** | Pointer to **string** |  | [optional] 
**WAgfInfo** | Pointer to [**WAgfInfo**](WAgfInfo.md) |  | [optional] 
**TngfInfo** | Pointer to [**TngfInfo**](TngfInfo.md) |  | [optional] 
**TwifInfo** | Pointer to [**TwifInfo**](TwifInfo.md) |  | [optional] 
**RanUnchangedInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmContextCreateData

`func NewSmContextCreateData(servingNfId string, servingNetwork PlmnIdNid, anType AccessType, smContextStatusUri string, ) *SmContextCreateData`

NewSmContextCreateData instantiates a new SmContextCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreateDataWithDefaults

`func NewSmContextCreateDataWithDefaults() *SmContextCreateData`

NewSmContextCreateDataWithDefaults instantiates a new SmContextCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *SmContextCreateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SmContextCreateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SmContextCreateData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *SmContextCreateData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUnauthenticatedSupi

`func (o *SmContextCreateData) GetUnauthenticatedSupi() bool`

GetUnauthenticatedSupi returns the UnauthenticatedSupi field if non-nil, zero value otherwise.

### GetUnauthenticatedSupiOk

`func (o *SmContextCreateData) GetUnauthenticatedSupiOk() (*bool, bool)`

GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSupi

`func (o *SmContextCreateData) SetUnauthenticatedSupi(v bool)`

SetUnauthenticatedSupi sets UnauthenticatedSupi field to given value.

### HasUnauthenticatedSupi

`func (o *SmContextCreateData) HasUnauthenticatedSupi() bool`

HasUnauthenticatedSupi returns a boolean if a field has been set.

### GetPei

`func (o *SmContextCreateData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *SmContextCreateData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *SmContextCreateData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *SmContextCreateData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetGpsi

`func (o *SmContextCreateData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmContextCreateData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmContextCreateData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmContextCreateData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmContextCreateData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContextCreateData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContextCreateData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *SmContextCreateData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetDnn

`func (o *SmContextCreateData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmContextCreateData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmContextCreateData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmContextCreateData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSelectedDnn

`func (o *SmContextCreateData) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *SmContextCreateData) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *SmContextCreateData) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *SmContextCreateData) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *SmContextCreateData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SmContextCreateData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SmContextCreateData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *SmContextCreateData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetHplmnSnssai

`func (o *SmContextCreateData) GetHplmnSnssai() Snssai`

GetHplmnSnssai returns the HplmnSnssai field if non-nil, zero value otherwise.

### GetHplmnSnssaiOk

`func (o *SmContextCreateData) GetHplmnSnssaiOk() (*Snssai, bool)`

GetHplmnSnssaiOk returns a tuple with the HplmnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHplmnSnssai

`func (o *SmContextCreateData) SetHplmnSnssai(v Snssai)`

SetHplmnSnssai sets HplmnSnssai field to given value.

### HasHplmnSnssai

`func (o *SmContextCreateData) HasHplmnSnssai() bool`

HasHplmnSnssai returns a boolean if a field has been set.

### GetServingNfId

`func (o *SmContextCreateData) GetServingNfId() string`

GetServingNfId returns the ServingNfId field if non-nil, zero value otherwise.

### GetServingNfIdOk

`func (o *SmContextCreateData) GetServingNfIdOk() (*string, bool)`

GetServingNfIdOk returns a tuple with the ServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfId

`func (o *SmContextCreateData) SetServingNfId(v string)`

SetServingNfId sets ServingNfId field to given value.


### GetGuami

`func (o *SmContextCreateData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *SmContextCreateData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *SmContextCreateData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *SmContextCreateData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServiceName

`func (o *SmContextCreateData) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SmContextCreateData) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SmContextCreateData) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SmContextCreateData) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmContextCreateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmContextCreateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmContextCreateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.


### GetRequestType

`func (o *SmContextCreateData) GetRequestType() RequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SmContextCreateData) GetRequestTypeOk() (*RequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SmContextCreateData) SetRequestType(v RequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *SmContextCreateData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetN1SmMsg

`func (o *SmContextCreateData) GetN1SmMsg() RefToBinaryData`

GetN1SmMsg returns the N1SmMsg field if non-nil, zero value otherwise.

### GetN1SmMsgOk

`func (o *SmContextCreateData) GetN1SmMsgOk() (*RefToBinaryData, bool)`

GetN1SmMsgOk returns a tuple with the N1SmMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmMsg

`func (o *SmContextCreateData) SetN1SmMsg(v RefToBinaryData)`

SetN1SmMsg sets N1SmMsg field to given value.

### HasN1SmMsg

`func (o *SmContextCreateData) HasN1SmMsg() bool`

HasN1SmMsg returns a boolean if a field has been set.

### GetAnType

`func (o *SmContextCreateData) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *SmContextCreateData) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *SmContextCreateData) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetAdditionalAnType

`func (o *SmContextCreateData) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *SmContextCreateData) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *SmContextCreateData) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *SmContextCreateData) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.

### GetRatType

`func (o *SmContextCreateData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SmContextCreateData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SmContextCreateData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SmContextCreateData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetPresenceInLadn

`func (o *SmContextCreateData) GetPresenceInLadn() PresenceState`

GetPresenceInLadn returns the PresenceInLadn field if non-nil, zero value otherwise.

### GetPresenceInLadnOk

`func (o *SmContextCreateData) GetPresenceInLadnOk() (*PresenceState, bool)`

GetPresenceInLadnOk returns a tuple with the PresenceInLadn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInLadn

`func (o *SmContextCreateData) SetPresenceInLadn(v PresenceState)`

SetPresenceInLadn sets PresenceInLadn field to given value.

### HasPresenceInLadn

`func (o *SmContextCreateData) HasPresenceInLadn() bool`

HasPresenceInLadn returns a boolean if a field has been set.

### GetUeLocation

`func (o *SmContextCreateData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SmContextCreateData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SmContextCreateData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SmContextCreateData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmContextCreateData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmContextCreateData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmContextCreateData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmContextCreateData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *SmContextCreateData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *SmContextCreateData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *SmContextCreateData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *SmContextCreateData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetSmContextStatusUri

`func (o *SmContextCreateData) GetSmContextStatusUri() string`

GetSmContextStatusUri returns the SmContextStatusUri field if non-nil, zero value otherwise.

### GetSmContextStatusUriOk

`func (o *SmContextCreateData) GetSmContextStatusUriOk() (*string, bool)`

GetSmContextStatusUriOk returns a tuple with the SmContextStatusUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextStatusUri

`func (o *SmContextCreateData) SetSmContextStatusUri(v string)`

SetSmContextStatusUri sets SmContextStatusUri field to given value.


### GetHSmfUri

`func (o *SmContextCreateData) GetHSmfUri() string`

GetHSmfUri returns the HSmfUri field if non-nil, zero value otherwise.

### GetHSmfUriOk

`func (o *SmContextCreateData) GetHSmfUriOk() (*string, bool)`

GetHSmfUriOk returns a tuple with the HSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfUri

`func (o *SmContextCreateData) SetHSmfUri(v string)`

SetHSmfUri sets HSmfUri field to given value.

### HasHSmfUri

`func (o *SmContextCreateData) HasHSmfUri() bool`

HasHSmfUri returns a boolean if a field has been set.

### GetHSmfId

`func (o *SmContextCreateData) GetHSmfId() string`

GetHSmfId returns the HSmfId field if non-nil, zero value otherwise.

### GetHSmfIdOk

`func (o *SmContextCreateData) GetHSmfIdOk() (*string, bool)`

GetHSmfIdOk returns a tuple with the HSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfId

`func (o *SmContextCreateData) SetHSmfId(v string)`

SetHSmfId sets HSmfId field to given value.

### HasHSmfId

`func (o *SmContextCreateData) HasHSmfId() bool`

HasHSmfId returns a boolean if a field has been set.

### GetSmfUri

`func (o *SmContextCreateData) GetSmfUri() string`

GetSmfUri returns the SmfUri field if non-nil, zero value otherwise.

### GetSmfUriOk

`func (o *SmContextCreateData) GetSmfUriOk() (*string, bool)`

GetSmfUriOk returns a tuple with the SmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUri

`func (o *SmContextCreateData) SetSmfUri(v string)`

SetSmfUri sets SmfUri field to given value.

### HasSmfUri

`func (o *SmContextCreateData) HasSmfUri() bool`

HasSmfUri returns a boolean if a field has been set.

### GetSmfId

`func (o *SmContextCreateData) GetSmfId() string`

GetSmfId returns the SmfId field if non-nil, zero value otherwise.

### GetSmfIdOk

`func (o *SmContextCreateData) GetSmfIdOk() (*string, bool)`

GetSmfIdOk returns a tuple with the SmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfId

`func (o *SmContextCreateData) SetSmfId(v string)`

SetSmfId sets SmfId field to given value.

### HasSmfId

`func (o *SmContextCreateData) HasSmfId() bool`

HasSmfId returns a boolean if a field has been set.

### GetAdditionalHsmfUri

`func (o *SmContextCreateData) GetAdditionalHsmfUri() []string`

GetAdditionalHsmfUri returns the AdditionalHsmfUri field if non-nil, zero value otherwise.

### GetAdditionalHsmfUriOk

`func (o *SmContextCreateData) GetAdditionalHsmfUriOk() (*[]string, bool)`

GetAdditionalHsmfUriOk returns a tuple with the AdditionalHsmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHsmfUri

`func (o *SmContextCreateData) SetAdditionalHsmfUri(v []string)`

SetAdditionalHsmfUri sets AdditionalHsmfUri field to given value.

### HasAdditionalHsmfUri

`func (o *SmContextCreateData) HasAdditionalHsmfUri() bool`

HasAdditionalHsmfUri returns a boolean if a field has been set.

### GetAdditionalHsmfId

`func (o *SmContextCreateData) GetAdditionalHsmfId() []string`

GetAdditionalHsmfId returns the AdditionalHsmfId field if non-nil, zero value otherwise.

### GetAdditionalHsmfIdOk

`func (o *SmContextCreateData) GetAdditionalHsmfIdOk() (*[]string, bool)`

GetAdditionalHsmfIdOk returns a tuple with the AdditionalHsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHsmfId

`func (o *SmContextCreateData) SetAdditionalHsmfId(v []string)`

SetAdditionalHsmfId sets AdditionalHsmfId field to given value.

### HasAdditionalHsmfId

`func (o *SmContextCreateData) HasAdditionalHsmfId() bool`

HasAdditionalHsmfId returns a boolean if a field has been set.

### GetAdditionalSmfUri

`func (o *SmContextCreateData) GetAdditionalSmfUri() []string`

GetAdditionalSmfUri returns the AdditionalSmfUri field if non-nil, zero value otherwise.

### GetAdditionalSmfUriOk

`func (o *SmContextCreateData) GetAdditionalSmfUriOk() (*[]string, bool)`

GetAdditionalSmfUriOk returns a tuple with the AdditionalSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSmfUri

`func (o *SmContextCreateData) SetAdditionalSmfUri(v []string)`

SetAdditionalSmfUri sets AdditionalSmfUri field to given value.

### HasAdditionalSmfUri

`func (o *SmContextCreateData) HasAdditionalSmfUri() bool`

HasAdditionalSmfUri returns a boolean if a field has been set.

### GetAdditionalSmfId

`func (o *SmContextCreateData) GetAdditionalSmfId() []string`

GetAdditionalSmfId returns the AdditionalSmfId field if non-nil, zero value otherwise.

### GetAdditionalSmfIdOk

`func (o *SmContextCreateData) GetAdditionalSmfIdOk() (*[]string, bool)`

GetAdditionalSmfIdOk returns a tuple with the AdditionalSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSmfId

`func (o *SmContextCreateData) SetAdditionalSmfId(v []string)`

SetAdditionalSmfId sets AdditionalSmfId field to given value.

### HasAdditionalSmfId

`func (o *SmContextCreateData) HasAdditionalSmfId() bool`

HasAdditionalSmfId returns a boolean if a field has been set.

### GetOldPduSessionId

`func (o *SmContextCreateData) GetOldPduSessionId() int32`

GetOldPduSessionId returns the OldPduSessionId field if non-nil, zero value otherwise.

### GetOldPduSessionIdOk

`func (o *SmContextCreateData) GetOldPduSessionIdOk() (*int32, bool)`

GetOldPduSessionIdOk returns a tuple with the OldPduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionId

`func (o *SmContextCreateData) SetOldPduSessionId(v int32)`

SetOldPduSessionId sets OldPduSessionId field to given value.

### HasOldPduSessionId

`func (o *SmContextCreateData) HasOldPduSessionId() bool`

HasOldPduSessionId returns a boolean if a field has been set.

### GetPduSessionsActivateList

`func (o *SmContextCreateData) GetPduSessionsActivateList() []int32`

GetPduSessionsActivateList returns the PduSessionsActivateList field if non-nil, zero value otherwise.

### GetPduSessionsActivateListOk

`func (o *SmContextCreateData) GetPduSessionsActivateListOk() (*[]int32, bool)`

GetPduSessionsActivateListOk returns a tuple with the PduSessionsActivateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionsActivateList

`func (o *SmContextCreateData) SetPduSessionsActivateList(v []int32)`

SetPduSessionsActivateList sets PduSessionsActivateList field to given value.

### HasPduSessionsActivateList

`func (o *SmContextCreateData) HasPduSessionsActivateList() bool`

HasPduSessionsActivateList returns a boolean if a field has been set.

### GetUeEpsPdnConnection

`func (o *SmContextCreateData) GetUeEpsPdnConnection() string`

GetUeEpsPdnConnection returns the UeEpsPdnConnection field if non-nil, zero value otherwise.

### GetUeEpsPdnConnectionOk

`func (o *SmContextCreateData) GetUeEpsPdnConnectionOk() (*string, bool)`

GetUeEpsPdnConnectionOk returns a tuple with the UeEpsPdnConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeEpsPdnConnection

`func (o *SmContextCreateData) SetUeEpsPdnConnection(v string)`

SetUeEpsPdnConnection sets UeEpsPdnConnection field to given value.

### HasUeEpsPdnConnection

`func (o *SmContextCreateData) HasUeEpsPdnConnection() bool`

HasUeEpsPdnConnection returns a boolean if a field has been set.

### GetHoState

`func (o *SmContextCreateData) GetHoState() HoState`

GetHoState returns the HoState field if non-nil, zero value otherwise.

### GetHoStateOk

`func (o *SmContextCreateData) GetHoStateOk() (*HoState, bool)`

GetHoStateOk returns a tuple with the HoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoState

`func (o *SmContextCreateData) SetHoState(v HoState)`

SetHoState sets HoState field to given value.

### HasHoState

`func (o *SmContextCreateData) HasHoState() bool`

HasHoState returns a boolean if a field has been set.

### GetPcfId

`func (o *SmContextCreateData) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *SmContextCreateData) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *SmContextCreateData) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *SmContextCreateData) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfGroupId

`func (o *SmContextCreateData) GetPcfGroupId() string`

GetPcfGroupId returns the PcfGroupId field if non-nil, zero value otherwise.

### GetPcfGroupIdOk

`func (o *SmContextCreateData) GetPcfGroupIdOk() (*string, bool)`

GetPcfGroupIdOk returns a tuple with the PcfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfGroupId

`func (o *SmContextCreateData) SetPcfGroupId(v string)`

SetPcfGroupId sets PcfGroupId field to given value.

### HasPcfGroupId

`func (o *SmContextCreateData) HasPcfGroupId() bool`

HasPcfGroupId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *SmContextCreateData) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *SmContextCreateData) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *SmContextCreateData) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *SmContextCreateData) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetNrfUri

`func (o *SmContextCreateData) GetNrfUri() string`

GetNrfUri returns the NrfUri field if non-nil, zero value otherwise.

### GetNrfUriOk

`func (o *SmContextCreateData) GetNrfUriOk() (*string, bool)`

GetNrfUriOk returns a tuple with the NrfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfUri

`func (o *SmContextCreateData) SetNrfUri(v string)`

SetNrfUri sets NrfUri field to given value.

### HasNrfUri

`func (o *SmContextCreateData) HasNrfUri() bool`

HasNrfUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSelMode

`func (o *SmContextCreateData) GetSelMode() DnnSelectionMode`

GetSelMode returns the SelMode field if non-nil, zero value otherwise.

### GetSelModeOk

`func (o *SmContextCreateData) GetSelModeOk() (*DnnSelectionMode, bool)`

GetSelModeOk returns a tuple with the SelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelMode

`func (o *SmContextCreateData) SetSelMode(v DnnSelectionMode)`

SetSelMode sets SelMode field to given value.

### HasSelMode

`func (o *SmContextCreateData) HasSelMode() bool`

HasSelMode returns a boolean if a field has been set.

### GetBackupAmfInfo

`func (o *SmContextCreateData) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *SmContextCreateData) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *SmContextCreateData) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *SmContextCreateData) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetTraceData

`func (o *SmContextCreateData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmContextCreateData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmContextCreateData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmContextCreateData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmContextCreateData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmContextCreateData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetUdmGroupId

`func (o *SmContextCreateData) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *SmContextCreateData) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *SmContextCreateData) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *SmContextCreateData) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *SmContextCreateData) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *SmContextCreateData) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *SmContextCreateData) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *SmContextCreateData) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetEpsInterworkingInd

`func (o *SmContextCreateData) GetEpsInterworkingInd() EpsInterworkingIndication`

GetEpsInterworkingInd returns the EpsInterworkingInd field if non-nil, zero value otherwise.

### GetEpsInterworkingIndOk

`func (o *SmContextCreateData) GetEpsInterworkingIndOk() (*EpsInterworkingIndication, bool)`

GetEpsInterworkingIndOk returns a tuple with the EpsInterworkingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInd

`func (o *SmContextCreateData) SetEpsInterworkingInd(v EpsInterworkingIndication)`

SetEpsInterworkingInd sets EpsInterworkingInd field to given value.

### HasEpsInterworkingInd

`func (o *SmContextCreateData) HasEpsInterworkingInd() bool`

HasEpsInterworkingInd returns a boolean if a field has been set.

### GetIndirectForwardingFlag

`func (o *SmContextCreateData) GetIndirectForwardingFlag() bool`

GetIndirectForwardingFlag returns the IndirectForwardingFlag field if non-nil, zero value otherwise.

### GetIndirectForwardingFlagOk

`func (o *SmContextCreateData) GetIndirectForwardingFlagOk() (*bool, bool)`

GetIndirectForwardingFlagOk returns a tuple with the IndirectForwardingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectForwardingFlag

`func (o *SmContextCreateData) SetIndirectForwardingFlag(v bool)`

SetIndirectForwardingFlag sets IndirectForwardingFlag field to given value.

### HasIndirectForwardingFlag

`func (o *SmContextCreateData) HasIndirectForwardingFlag() bool`

HasIndirectForwardingFlag returns a boolean if a field has been set.

### GetDirectForwardingFlag

`func (o *SmContextCreateData) GetDirectForwardingFlag() bool`

GetDirectForwardingFlag returns the DirectForwardingFlag field if non-nil, zero value otherwise.

### GetDirectForwardingFlagOk

`func (o *SmContextCreateData) GetDirectForwardingFlagOk() (*bool, bool)`

GetDirectForwardingFlagOk returns a tuple with the DirectForwardingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectForwardingFlag

`func (o *SmContextCreateData) SetDirectForwardingFlag(v bool)`

SetDirectForwardingFlag sets DirectForwardingFlag field to given value.

### HasDirectForwardingFlag

`func (o *SmContextCreateData) HasDirectForwardingFlag() bool`

HasDirectForwardingFlag returns a boolean if a field has been set.

### GetTargetId

`func (o *SmContextCreateData) GetTargetId() NgRanTargetId`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SmContextCreateData) GetTargetIdOk() (*NgRanTargetId, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SmContextCreateData) SetTargetId(v NgRanTargetId)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SmContextCreateData) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetEpsBearerCtxStatus

`func (o *SmContextCreateData) GetEpsBearerCtxStatus() string`

GetEpsBearerCtxStatus returns the EpsBearerCtxStatus field if non-nil, zero value otherwise.

### GetEpsBearerCtxStatusOk

`func (o *SmContextCreateData) GetEpsBearerCtxStatusOk() (*string, bool)`

GetEpsBearerCtxStatusOk returns a tuple with the EpsBearerCtxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerCtxStatus

`func (o *SmContextCreateData) SetEpsBearerCtxStatus(v string)`

SetEpsBearerCtxStatus sets EpsBearerCtxStatus field to given value.

### HasEpsBearerCtxStatus

`func (o *SmContextCreateData) HasEpsBearerCtxStatus() bool`

HasEpsBearerCtxStatus returns a boolean if a field has been set.

### GetCpCiotEnabled

`func (o *SmContextCreateData) GetCpCiotEnabled() bool`

GetCpCiotEnabled returns the CpCiotEnabled field if non-nil, zero value otherwise.

### GetCpCiotEnabledOk

`func (o *SmContextCreateData) GetCpCiotEnabledOk() (*bool, bool)`

GetCpCiotEnabledOk returns a tuple with the CpCiotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpCiotEnabled

`func (o *SmContextCreateData) SetCpCiotEnabled(v bool)`

SetCpCiotEnabled sets CpCiotEnabled field to given value.

### HasCpCiotEnabled

`func (o *SmContextCreateData) HasCpCiotEnabled() bool`

HasCpCiotEnabled returns a boolean if a field has been set.

### GetCpOnlyInd

`func (o *SmContextCreateData) GetCpOnlyInd() bool`

GetCpOnlyInd returns the CpOnlyInd field if non-nil, zero value otherwise.

### GetCpOnlyIndOk

`func (o *SmContextCreateData) GetCpOnlyIndOk() (*bool, bool)`

GetCpOnlyIndOk returns a tuple with the CpOnlyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpOnlyInd

`func (o *SmContextCreateData) SetCpOnlyInd(v bool)`

SetCpOnlyInd sets CpOnlyInd field to given value.

### HasCpOnlyInd

`func (o *SmContextCreateData) HasCpOnlyInd() bool`

HasCpOnlyInd returns a boolean if a field has been set.

### GetInvokeNef

`func (o *SmContextCreateData) GetInvokeNef() bool`

GetInvokeNef returns the InvokeNef field if non-nil, zero value otherwise.

### GetInvokeNefOk

`func (o *SmContextCreateData) GetInvokeNefOk() (*bool, bool)`

GetInvokeNefOk returns a tuple with the InvokeNef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeNef

`func (o *SmContextCreateData) SetInvokeNef(v bool)`

SetInvokeNef sets InvokeNef field to given value.

### HasInvokeNef

`func (o *SmContextCreateData) HasInvokeNef() bool`

HasInvokeNef returns a boolean if a field has been set.

### GetMaRequestInd

`func (o *SmContextCreateData) GetMaRequestInd() bool`

GetMaRequestInd returns the MaRequestInd field if non-nil, zero value otherwise.

### GetMaRequestIndOk

`func (o *SmContextCreateData) GetMaRequestIndOk() (*bool, bool)`

GetMaRequestIndOk returns a tuple with the MaRequestInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaRequestInd

`func (o *SmContextCreateData) SetMaRequestInd(v bool)`

SetMaRequestInd sets MaRequestInd field to given value.

### HasMaRequestInd

`func (o *SmContextCreateData) HasMaRequestInd() bool`

HasMaRequestInd returns a boolean if a field has been set.

### GetMaNwUpgradeInd

`func (o *SmContextCreateData) GetMaNwUpgradeInd() bool`

GetMaNwUpgradeInd returns the MaNwUpgradeInd field if non-nil, zero value otherwise.

### GetMaNwUpgradeIndOk

`func (o *SmContextCreateData) GetMaNwUpgradeIndOk() (*bool, bool)`

GetMaNwUpgradeIndOk returns a tuple with the MaNwUpgradeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaNwUpgradeInd

`func (o *SmContextCreateData) SetMaNwUpgradeInd(v bool)`

SetMaNwUpgradeInd sets MaNwUpgradeInd field to given value.

### HasMaNwUpgradeInd

`func (o *SmContextCreateData) HasMaNwUpgradeInd() bool`

HasMaNwUpgradeInd returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextCreateData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextCreateData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextCreateData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextCreateData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextCreateData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextCreateData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextCreateData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextCreateData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetN2SmInfoExt1

`func (o *SmContextCreateData) GetN2SmInfoExt1() RefToBinaryData`

GetN2SmInfoExt1 returns the N2SmInfoExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoExt1Ok

`func (o *SmContextCreateData) GetN2SmInfoExt1Ok() (*RefToBinaryData, bool)`

GetN2SmInfoExt1Ok returns a tuple with the N2SmInfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoExt1

`func (o *SmContextCreateData) SetN2SmInfoExt1(v RefToBinaryData)`

SetN2SmInfoExt1 sets N2SmInfoExt1 field to given value.

### HasN2SmInfoExt1

`func (o *SmContextCreateData) HasN2SmInfoExt1() bool`

HasN2SmInfoExt1 returns a boolean if a field has been set.

### GetN2SmInfoTypeExt1

`func (o *SmContextCreateData) GetN2SmInfoTypeExt1() N2SmInfoType`

GetN2SmInfoTypeExt1 returns the N2SmInfoTypeExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoTypeExt1Ok

`func (o *SmContextCreateData) GetN2SmInfoTypeExt1Ok() (*N2SmInfoType, bool)`

GetN2SmInfoTypeExt1Ok returns a tuple with the N2SmInfoTypeExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoTypeExt1

`func (o *SmContextCreateData) SetN2SmInfoTypeExt1(v N2SmInfoType)`

SetN2SmInfoTypeExt1 sets N2SmInfoTypeExt1 field to given value.

### HasN2SmInfoTypeExt1

`func (o *SmContextCreateData) HasN2SmInfoTypeExt1() bool`

HasN2SmInfoTypeExt1 returns a boolean if a field has been set.

### GetSmContextRef

`func (o *SmContextCreateData) GetSmContextRef() string`

GetSmContextRef returns the SmContextRef field if non-nil, zero value otherwise.

### GetSmContextRefOk

`func (o *SmContextCreateData) GetSmContextRefOk() (*string, bool)`

GetSmContextRefOk returns a tuple with the SmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextRef

`func (o *SmContextCreateData) SetSmContextRef(v string)`

SetSmContextRef sets SmContextRef field to given value.

### HasSmContextRef

`func (o *SmContextCreateData) HasSmContextRef() bool`

HasSmContextRef returns a boolean if a field has been set.

### GetSmContextSmfId

`func (o *SmContextCreateData) GetSmContextSmfId() string`

GetSmContextSmfId returns the SmContextSmfId field if non-nil, zero value otherwise.

### GetSmContextSmfIdOk

`func (o *SmContextCreateData) GetSmContextSmfIdOk() (*string, bool)`

GetSmContextSmfIdOk returns a tuple with the SmContextSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfId

`func (o *SmContextCreateData) SetSmContextSmfId(v string)`

SetSmContextSmfId sets SmContextSmfId field to given value.

### HasSmContextSmfId

`func (o *SmContextCreateData) HasSmContextSmfId() bool`

HasSmContextSmfId returns a boolean if a field has been set.

### GetSmContextSmfSetId

`func (o *SmContextCreateData) GetSmContextSmfSetId() string`

GetSmContextSmfSetId returns the SmContextSmfSetId field if non-nil, zero value otherwise.

### GetSmContextSmfSetIdOk

`func (o *SmContextCreateData) GetSmContextSmfSetIdOk() (*string, bool)`

GetSmContextSmfSetIdOk returns a tuple with the SmContextSmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfSetId

`func (o *SmContextCreateData) SetSmContextSmfSetId(v string)`

SetSmContextSmfSetId sets SmContextSmfSetId field to given value.

### HasSmContextSmfSetId

`func (o *SmContextCreateData) HasSmContextSmfSetId() bool`

HasSmContextSmfSetId returns a boolean if a field has been set.

### GetSmContextSmfServiceSetId

`func (o *SmContextCreateData) GetSmContextSmfServiceSetId() string`

GetSmContextSmfServiceSetId returns the SmContextSmfServiceSetId field if non-nil, zero value otherwise.

### GetSmContextSmfServiceSetIdOk

`func (o *SmContextCreateData) GetSmContextSmfServiceSetIdOk() (*string, bool)`

GetSmContextSmfServiceSetIdOk returns a tuple with the SmContextSmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfServiceSetId

`func (o *SmContextCreateData) SetSmContextSmfServiceSetId(v string)`

SetSmContextSmfServiceSetId sets SmContextSmfServiceSetId field to given value.

### HasSmContextSmfServiceSetId

`func (o *SmContextCreateData) HasSmContextSmfServiceSetId() bool`

HasSmContextSmfServiceSetId returns a boolean if a field has been set.

### GetSmContextSmfBinding

`func (o *SmContextCreateData) GetSmContextSmfBinding() SbiBindingLevel`

GetSmContextSmfBinding returns the SmContextSmfBinding field if non-nil, zero value otherwise.

### GetSmContextSmfBindingOk

`func (o *SmContextCreateData) GetSmContextSmfBindingOk() (*SbiBindingLevel, bool)`

GetSmContextSmfBindingOk returns a tuple with the SmContextSmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfBinding

`func (o *SmContextCreateData) SetSmContextSmfBinding(v SbiBindingLevel)`

SetSmContextSmfBinding sets SmContextSmfBinding field to given value.

### HasSmContextSmfBinding

`func (o *SmContextCreateData) HasSmContextSmfBinding() bool`

HasSmContextSmfBinding returns a boolean if a field has been set.

### GetUpCnxState

`func (o *SmContextCreateData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *SmContextCreateData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *SmContextCreateData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *SmContextCreateData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *SmContextCreateData) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextCreateData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextCreateData) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextCreateData) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextCreateData) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextCreateData) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextCreateData) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextCreateData) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetExtendedNasSmTimerInd

`func (o *SmContextCreateData) GetExtendedNasSmTimerInd() bool`

GetExtendedNasSmTimerInd returns the ExtendedNasSmTimerInd field if non-nil, zero value otherwise.

### GetExtendedNasSmTimerIndOk

`func (o *SmContextCreateData) GetExtendedNasSmTimerIndOk() (*bool, bool)`

GetExtendedNasSmTimerIndOk returns a tuple with the ExtendedNasSmTimerInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedNasSmTimerInd

`func (o *SmContextCreateData) SetExtendedNasSmTimerInd(v bool)`

SetExtendedNasSmTimerInd sets ExtendedNasSmTimerInd field to given value.

### HasExtendedNasSmTimerInd

`func (o *SmContextCreateData) HasExtendedNasSmTimerInd() bool`

HasExtendedNasSmTimerInd returns a boolean if a field has been set.

### GetDlDataWaitingInd

`func (o *SmContextCreateData) GetDlDataWaitingInd() bool`

GetDlDataWaitingInd returns the DlDataWaitingInd field if non-nil, zero value otherwise.

### GetDlDataWaitingIndOk

`func (o *SmContextCreateData) GetDlDataWaitingIndOk() (*bool, bool)`

GetDlDataWaitingIndOk returns a tuple with the DlDataWaitingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDataWaitingInd

`func (o *SmContextCreateData) SetDlDataWaitingInd(v bool)`

SetDlDataWaitingInd sets DlDataWaitingInd field to given value.

### HasDlDataWaitingInd

`func (o *SmContextCreateData) HasDlDataWaitingInd() bool`

HasDlDataWaitingInd returns a boolean if a field has been set.

### GetDdnFailureSubs

`func (o *SmContextCreateData) GetDdnFailureSubs() DdnFailureSubs`

GetDdnFailureSubs returns the DdnFailureSubs field if non-nil, zero value otherwise.

### GetDdnFailureSubsOk

`func (o *SmContextCreateData) GetDdnFailureSubsOk() (*DdnFailureSubs, bool)`

GetDdnFailureSubsOk returns a tuple with the DdnFailureSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnFailureSubs

`func (o *SmContextCreateData) SetDdnFailureSubs(v DdnFailureSubs)`

SetDdnFailureSubs sets DdnFailureSubs field to given value.

### HasDdnFailureSubs

`func (o *SmContextCreateData) HasDdnFailureSubs() bool`

HasDdnFailureSubs returns a boolean if a field has been set.

### GetSmfTransferInd

`func (o *SmContextCreateData) GetSmfTransferInd() bool`

GetSmfTransferInd returns the SmfTransferInd field if non-nil, zero value otherwise.

### GetSmfTransferIndOk

`func (o *SmContextCreateData) GetSmfTransferIndOk() (*bool, bool)`

GetSmfTransferIndOk returns a tuple with the SmfTransferInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfTransferInd

`func (o *SmContextCreateData) SetSmfTransferInd(v bool)`

SetSmfTransferInd sets SmfTransferInd field to given value.

### HasSmfTransferInd

`func (o *SmContextCreateData) HasSmfTransferInd() bool`

HasSmfTransferInd returns a boolean if a field has been set.

### GetOldSmfId

`func (o *SmContextCreateData) GetOldSmfId() string`

GetOldSmfId returns the OldSmfId field if non-nil, zero value otherwise.

### GetOldSmfIdOk

`func (o *SmContextCreateData) GetOldSmfIdOk() (*string, bool)`

GetOldSmfIdOk returns a tuple with the OldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmfId

`func (o *SmContextCreateData) SetOldSmfId(v string)`

SetOldSmfId sets OldSmfId field to given value.

### HasOldSmfId

`func (o *SmContextCreateData) HasOldSmfId() bool`

HasOldSmfId returns a boolean if a field has been set.

### GetOldSmContextRef

`func (o *SmContextCreateData) GetOldSmContextRef() string`

GetOldSmContextRef returns the OldSmContextRef field if non-nil, zero value otherwise.

### GetOldSmContextRefOk

`func (o *SmContextCreateData) GetOldSmContextRefOk() (*string, bool)`

GetOldSmContextRefOk returns a tuple with the OldSmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmContextRef

`func (o *SmContextCreateData) SetOldSmContextRef(v string)`

SetOldSmContextRef sets OldSmContextRef field to given value.

### HasOldSmContextRef

`func (o *SmContextCreateData) HasOldSmContextRef() bool`

HasOldSmContextRef returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *SmContextCreateData) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *SmContextCreateData) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *SmContextCreateData) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *SmContextCreateData) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### GetTngfInfo

`func (o *SmContextCreateData) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *SmContextCreateData) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *SmContextCreateData) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *SmContextCreateData) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### GetTwifInfo

`func (o *SmContextCreateData) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *SmContextCreateData) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *SmContextCreateData) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *SmContextCreateData) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### GetRanUnchangedInd

`func (o *SmContextCreateData) GetRanUnchangedInd() bool`

GetRanUnchangedInd returns the RanUnchangedInd field if non-nil, zero value otherwise.

### GetRanUnchangedIndOk

`func (o *SmContextCreateData) GetRanUnchangedIndOk() (*bool, bool)`

GetRanUnchangedIndOk returns a tuple with the RanUnchangedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUnchangedInd

`func (o *SmContextCreateData) SetRanUnchangedInd(v bool)`

SetRanUnchangedInd sets RanUnchangedInd field to given value.

### HasRanUnchangedInd

`func (o *SmContextCreateData) HasRanUnchangedInd() bool`

HasRanUnchangedInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


