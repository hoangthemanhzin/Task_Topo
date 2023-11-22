# TrafficInfluDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpPathChgNotifCorreId** | Pointer to **string** | Contains the Notification Correlation Id allocated by the NEF for the UP path change notification. | [optional] 
**AppReloInd** | Pointer to **bool** | Identifies whether an application can be relocated once a location of the application has been selected. | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**EthTrafficFilters** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet filters. Either \&quot;trafficFilters\&quot; or \&quot;ethTrafficFilters\&quot; shall be included if applicable. | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**InternalGroupId** | Pointer to **string** |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**TrafficFilters** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Identifies IP packet filters. Either \&quot;trafficFilters\&quot; or \&quot;ethTrafficFilters\&quot; shall be included if applicable. | [optional] 
**TrafficRoutes** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) | Identifies the N6 traffic routing requirement. | [optional] 
**TraffCorreInd** | Pointer to **bool** |  | [optional] 
**ValidStartTime** | Pointer to **time.Time** |  | [optional] 
**ValidEndTime** | Pointer to **time.Time** |  | [optional] 
**TempValidities** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) | Identifies the temporal validities for the N6 traffic routing requirement. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo1**](NetworkAreaInfo1.md) |  | [optional] 
**UpPathChgNotifUri** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **[]string** |  | [optional] 
**AfAckInd** | Pointer to **bool** |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrafficInfluDataPatch

`func NewTrafficInfluDataPatch() *TrafficInfluDataPatch`

NewTrafficInfluDataPatch instantiates a new TrafficInfluDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInfluDataPatchWithDefaults

`func NewTrafficInfluDataPatchWithDefaults() *TrafficInfluDataPatch`

NewTrafficInfluDataPatchWithDefaults instantiates a new TrafficInfluDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreId() string`

GetUpPathChgNotifCorreId returns the UpPathChgNotifCorreId field if non-nil, zero value otherwise.

### GetUpPathChgNotifCorreIdOk

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreIdOk() (*string, bool)`

GetUpPathChgNotifCorreIdOk returns a tuple with the UpPathChgNotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) SetUpPathChgNotifCorreId(v string)`

SetUpPathChgNotifCorreId sets UpPathChgNotifCorreId field to given value.

### HasUpPathChgNotifCorreId

`func (o *TrafficInfluDataPatch) HasUpPathChgNotifCorreId() bool`

HasUpPathChgNotifCorreId returns a boolean if a field has been set.

### GetAppReloInd

`func (o *TrafficInfluDataPatch) GetAppReloInd() bool`

GetAppReloInd returns the AppReloInd field if non-nil, zero value otherwise.

### GetAppReloIndOk

`func (o *TrafficInfluDataPatch) GetAppReloIndOk() (*bool, bool)`

GetAppReloIndOk returns a tuple with the AppReloInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloInd

`func (o *TrafficInfluDataPatch) SetAppReloInd(v bool)`

SetAppReloInd sets AppReloInd field to given value.

### HasAppReloInd

`func (o *TrafficInfluDataPatch) HasAppReloInd() bool`

HasAppReloInd returns a boolean if a field has been set.

### GetDnn

`func (o *TrafficInfluDataPatch) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TrafficInfluDataPatch) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TrafficInfluDataPatch) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TrafficInfluDataPatch) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetEthTrafficFilters

`func (o *TrafficInfluDataPatch) GetEthTrafficFilters() []EthFlowDescription`

GetEthTrafficFilters returns the EthTrafficFilters field if non-nil, zero value otherwise.

### GetEthTrafficFiltersOk

`func (o *TrafficInfluDataPatch) GetEthTrafficFiltersOk() (*[]EthFlowDescription, bool)`

GetEthTrafficFiltersOk returns a tuple with the EthTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilters

`func (o *TrafficInfluDataPatch) SetEthTrafficFilters(v []EthFlowDescription)`

SetEthTrafficFilters sets EthTrafficFilters field to given value.

### HasEthTrafficFilters

`func (o *TrafficInfluDataPatch) HasEthTrafficFilters() bool`

HasEthTrafficFilters returns a boolean if a field has been set.

### GetSnssai

`func (o *TrafficInfluDataPatch) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TrafficInfluDataPatch) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TrafficInfluDataPatch) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TrafficInfluDataPatch) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetInternalGroupId

`func (o *TrafficInfluDataPatch) GetInternalGroupId() string`

GetInternalGroupId returns the InternalGroupId field if non-nil, zero value otherwise.

### GetInternalGroupIdOk

`func (o *TrafficInfluDataPatch) GetInternalGroupIdOk() (*string, bool)`

GetInternalGroupIdOk returns a tuple with the InternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupId

`func (o *TrafficInfluDataPatch) SetInternalGroupId(v string)`

SetInternalGroupId sets InternalGroupId field to given value.

### HasInternalGroupId

`func (o *TrafficInfluDataPatch) HasInternalGroupId() bool`

HasInternalGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *TrafficInfluDataPatch) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *TrafficInfluDataPatch) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *TrafficInfluDataPatch) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *TrafficInfluDataPatch) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *TrafficInfluDataPatch) GetTrafficFilters() []FlowInfo`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *TrafficInfluDataPatch) GetTrafficFiltersOk() (*[]FlowInfo, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *TrafficInfluDataPatch) SetTrafficFilters(v []FlowInfo)`

SetTrafficFilters sets TrafficFilters field to given value.

### HasTrafficFilters

`func (o *TrafficInfluDataPatch) HasTrafficFilters() bool`

HasTrafficFilters returns a boolean if a field has been set.

### GetTrafficRoutes

`func (o *TrafficInfluDataPatch) GetTrafficRoutes() []RouteToLocation`

GetTrafficRoutes returns the TrafficRoutes field if non-nil, zero value otherwise.

### GetTrafficRoutesOk

`func (o *TrafficInfluDataPatch) GetTrafficRoutesOk() (*[]RouteToLocation, bool)`

GetTrafficRoutesOk returns a tuple with the TrafficRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRoutes

`func (o *TrafficInfluDataPatch) SetTrafficRoutes(v []RouteToLocation)`

SetTrafficRoutes sets TrafficRoutes field to given value.

### HasTrafficRoutes

`func (o *TrafficInfluDataPatch) HasTrafficRoutes() bool`

HasTrafficRoutes returns a boolean if a field has been set.

### GetTraffCorreInd

`func (o *TrafficInfluDataPatch) GetTraffCorreInd() bool`

GetTraffCorreInd returns the TraffCorreInd field if non-nil, zero value otherwise.

### GetTraffCorreIndOk

`func (o *TrafficInfluDataPatch) GetTraffCorreIndOk() (*bool, bool)`

GetTraffCorreIndOk returns a tuple with the TraffCorreInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffCorreInd

`func (o *TrafficInfluDataPatch) SetTraffCorreInd(v bool)`

SetTraffCorreInd sets TraffCorreInd field to given value.

### HasTraffCorreInd

`func (o *TrafficInfluDataPatch) HasTraffCorreInd() bool`

HasTraffCorreInd returns a boolean if a field has been set.

### GetValidStartTime

`func (o *TrafficInfluDataPatch) GetValidStartTime() time.Time`

GetValidStartTime returns the ValidStartTime field if non-nil, zero value otherwise.

### GetValidStartTimeOk

`func (o *TrafficInfluDataPatch) GetValidStartTimeOk() (*time.Time, bool)`

GetValidStartTimeOk returns a tuple with the ValidStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidStartTime

`func (o *TrafficInfluDataPatch) SetValidStartTime(v time.Time)`

SetValidStartTime sets ValidStartTime field to given value.

### HasValidStartTime

`func (o *TrafficInfluDataPatch) HasValidStartTime() bool`

HasValidStartTime returns a boolean if a field has been set.

### GetValidEndTime

`func (o *TrafficInfluDataPatch) GetValidEndTime() time.Time`

GetValidEndTime returns the ValidEndTime field if non-nil, zero value otherwise.

### GetValidEndTimeOk

`func (o *TrafficInfluDataPatch) GetValidEndTimeOk() (*time.Time, bool)`

GetValidEndTimeOk returns a tuple with the ValidEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidEndTime

`func (o *TrafficInfluDataPatch) SetValidEndTime(v time.Time)`

SetValidEndTime sets ValidEndTime field to given value.

### HasValidEndTime

`func (o *TrafficInfluDataPatch) HasValidEndTime() bool`

HasValidEndTime returns a boolean if a field has been set.

### GetTempValidities

`func (o *TrafficInfluDataPatch) GetTempValidities() []TemporalValidity`

GetTempValidities returns the TempValidities field if non-nil, zero value otherwise.

### GetTempValiditiesOk

`func (o *TrafficInfluDataPatch) GetTempValiditiesOk() (*[]TemporalValidity, bool)`

GetTempValiditiesOk returns a tuple with the TempValidities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidities

`func (o *TrafficInfluDataPatch) SetTempValidities(v []TemporalValidity)`

SetTempValidities sets TempValidities field to given value.

### HasTempValidities

`func (o *TrafficInfluDataPatch) HasTempValidities() bool`

HasTempValidities returns a boolean if a field has been set.

### SetTempValiditiesNil

`func (o *TrafficInfluDataPatch) SetTempValiditiesNil(b bool)`

 SetTempValiditiesNil sets the value for TempValidities to be an explicit nil

### UnsetTempValidities
`func (o *TrafficInfluDataPatch) UnsetTempValidities()`

UnsetTempValidities ensures that no value is present for TempValidities, not even an explicit nil
### GetNwAreaInfo

`func (o *TrafficInfluDataPatch) GetNwAreaInfo() NetworkAreaInfo1`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *TrafficInfluDataPatch) GetNwAreaInfoOk() (*NetworkAreaInfo1, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *TrafficInfluDataPatch) SetNwAreaInfo(v NetworkAreaInfo1)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *TrafficInfluDataPatch) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifUri() string`

GetUpPathChgNotifUri returns the UpPathChgNotifUri field if non-nil, zero value otherwise.

### GetUpPathChgNotifUriOk

`func (o *TrafficInfluDataPatch) GetUpPathChgNotifUriOk() (*string, bool)`

GetUpPathChgNotifUriOk returns a tuple with the UpPathChgNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) SetUpPathChgNotifUri(v string)`

SetUpPathChgNotifUri sets UpPathChgNotifUri field to given value.

### HasUpPathChgNotifUri

`func (o *TrafficInfluDataPatch) HasUpPathChgNotifUri() bool`

HasUpPathChgNotifUri returns a boolean if a field has been set.

### GetHeaders

`func (o *TrafficInfluDataPatch) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TrafficInfluDataPatch) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TrafficInfluDataPatch) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TrafficInfluDataPatch) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAfAckInd

`func (o *TrafficInfluDataPatch) GetAfAckInd() bool`

GetAfAckInd returns the AfAckInd field if non-nil, zero value otherwise.

### GetAfAckIndOk

`func (o *TrafficInfluDataPatch) GetAfAckIndOk() (*bool, bool)`

GetAfAckIndOk returns a tuple with the AfAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckInd

`func (o *TrafficInfluDataPatch) SetAfAckInd(v bool)`

SetAfAckInd sets AfAckInd field to given value.

### HasAfAckInd

`func (o *TrafficInfluDataPatch) HasAfAckInd() bool`

HasAfAckInd returns a boolean if a field has been set.

### GetAddrPreserInd

`func (o *TrafficInfluDataPatch) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *TrafficInfluDataPatch) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *TrafficInfluDataPatch) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *TrafficInfluDataPatch) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


