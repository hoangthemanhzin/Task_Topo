# TrafficInfluData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpPathChgNotifCorreId** | Pointer to **string** | Contains the Notification Correlation Id allocated by the NEF for the UP path change notification. | [optional] 
**AppReloInd** | Pointer to **bool** | Identifies whether an application can be relocated once a location of the application has been selected. | [optional] 
**AfAppId** | Pointer to **string** | Identifies an application. | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**EthTrafficFilters** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet filters. Either \&quot;trafficFilters\&quot; or \&quot;ethTrafficFilters\&quot; shall be included if applicable. | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**InterGroupId** | Pointer to **string** |  | [optional] 
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
**SubscribedEvents** | Pointer to [**[]SubscribedEvent**](SubscribedEvent.md) |  | [optional] 
**DnaiChgType** | Pointer to [**DnaiChangeType**](DnaiChangeType.md) |  | [optional] 
**AfAckInd** | Pointer to **bool** |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ResUri** | Pointer to **string** |  | [optional] 

## Methods

### NewTrafficInfluData

`func NewTrafficInfluData() *TrafficInfluData`

NewTrafficInfluData instantiates a new TrafficInfluData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInfluDataWithDefaults

`func NewTrafficInfluDataWithDefaults() *TrafficInfluData`

NewTrafficInfluDataWithDefaults instantiates a new TrafficInfluData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpPathChgNotifCorreId

`func (o *TrafficInfluData) GetUpPathChgNotifCorreId() string`

GetUpPathChgNotifCorreId returns the UpPathChgNotifCorreId field if non-nil, zero value otherwise.

### GetUpPathChgNotifCorreIdOk

`func (o *TrafficInfluData) GetUpPathChgNotifCorreIdOk() (*string, bool)`

GetUpPathChgNotifCorreIdOk returns a tuple with the UpPathChgNotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifCorreId

`func (o *TrafficInfluData) SetUpPathChgNotifCorreId(v string)`

SetUpPathChgNotifCorreId sets UpPathChgNotifCorreId field to given value.

### HasUpPathChgNotifCorreId

`func (o *TrafficInfluData) HasUpPathChgNotifCorreId() bool`

HasUpPathChgNotifCorreId returns a boolean if a field has been set.

### GetAppReloInd

`func (o *TrafficInfluData) GetAppReloInd() bool`

GetAppReloInd returns the AppReloInd field if non-nil, zero value otherwise.

### GetAppReloIndOk

`func (o *TrafficInfluData) GetAppReloIndOk() (*bool, bool)`

GetAppReloIndOk returns a tuple with the AppReloInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloInd

`func (o *TrafficInfluData) SetAppReloInd(v bool)`

SetAppReloInd sets AppReloInd field to given value.

### HasAppReloInd

`func (o *TrafficInfluData) HasAppReloInd() bool`

HasAppReloInd returns a boolean if a field has been set.

### GetAfAppId

`func (o *TrafficInfluData) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *TrafficInfluData) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *TrafficInfluData) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *TrafficInfluData) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetDnn

`func (o *TrafficInfluData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TrafficInfluData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TrafficInfluData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TrafficInfluData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetEthTrafficFilters

`func (o *TrafficInfluData) GetEthTrafficFilters() []EthFlowDescription`

GetEthTrafficFilters returns the EthTrafficFilters field if non-nil, zero value otherwise.

### GetEthTrafficFiltersOk

`func (o *TrafficInfluData) GetEthTrafficFiltersOk() (*[]EthFlowDescription, bool)`

GetEthTrafficFiltersOk returns a tuple with the EthTrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilters

`func (o *TrafficInfluData) SetEthTrafficFilters(v []EthFlowDescription)`

SetEthTrafficFilters sets EthTrafficFilters field to given value.

### HasEthTrafficFilters

`func (o *TrafficInfluData) HasEthTrafficFilters() bool`

HasEthTrafficFilters returns a boolean if a field has been set.

### GetSnssai

`func (o *TrafficInfluData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TrafficInfluData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TrafficInfluData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TrafficInfluData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetInterGroupId

`func (o *TrafficInfluData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *TrafficInfluData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *TrafficInfluData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *TrafficInfluData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *TrafficInfluData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *TrafficInfluData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *TrafficInfluData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *TrafficInfluData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *TrafficInfluData) GetTrafficFilters() []FlowInfo`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *TrafficInfluData) GetTrafficFiltersOk() (*[]FlowInfo, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *TrafficInfluData) SetTrafficFilters(v []FlowInfo)`

SetTrafficFilters sets TrafficFilters field to given value.

### HasTrafficFilters

`func (o *TrafficInfluData) HasTrafficFilters() bool`

HasTrafficFilters returns a boolean if a field has been set.

### GetTrafficRoutes

`func (o *TrafficInfluData) GetTrafficRoutes() []RouteToLocation`

GetTrafficRoutes returns the TrafficRoutes field if non-nil, zero value otherwise.

### GetTrafficRoutesOk

`func (o *TrafficInfluData) GetTrafficRoutesOk() (*[]RouteToLocation, bool)`

GetTrafficRoutesOk returns a tuple with the TrafficRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRoutes

`func (o *TrafficInfluData) SetTrafficRoutes(v []RouteToLocation)`

SetTrafficRoutes sets TrafficRoutes field to given value.

### HasTrafficRoutes

`func (o *TrafficInfluData) HasTrafficRoutes() bool`

HasTrafficRoutes returns a boolean if a field has been set.

### GetTraffCorreInd

`func (o *TrafficInfluData) GetTraffCorreInd() bool`

GetTraffCorreInd returns the TraffCorreInd field if non-nil, zero value otherwise.

### GetTraffCorreIndOk

`func (o *TrafficInfluData) GetTraffCorreIndOk() (*bool, bool)`

GetTraffCorreIndOk returns a tuple with the TraffCorreInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffCorreInd

`func (o *TrafficInfluData) SetTraffCorreInd(v bool)`

SetTraffCorreInd sets TraffCorreInd field to given value.

### HasTraffCorreInd

`func (o *TrafficInfluData) HasTraffCorreInd() bool`

HasTraffCorreInd returns a boolean if a field has been set.

### GetValidStartTime

`func (o *TrafficInfluData) GetValidStartTime() time.Time`

GetValidStartTime returns the ValidStartTime field if non-nil, zero value otherwise.

### GetValidStartTimeOk

`func (o *TrafficInfluData) GetValidStartTimeOk() (*time.Time, bool)`

GetValidStartTimeOk returns a tuple with the ValidStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidStartTime

`func (o *TrafficInfluData) SetValidStartTime(v time.Time)`

SetValidStartTime sets ValidStartTime field to given value.

### HasValidStartTime

`func (o *TrafficInfluData) HasValidStartTime() bool`

HasValidStartTime returns a boolean if a field has been set.

### GetValidEndTime

`func (o *TrafficInfluData) GetValidEndTime() time.Time`

GetValidEndTime returns the ValidEndTime field if non-nil, zero value otherwise.

### GetValidEndTimeOk

`func (o *TrafficInfluData) GetValidEndTimeOk() (*time.Time, bool)`

GetValidEndTimeOk returns a tuple with the ValidEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidEndTime

`func (o *TrafficInfluData) SetValidEndTime(v time.Time)`

SetValidEndTime sets ValidEndTime field to given value.

### HasValidEndTime

`func (o *TrafficInfluData) HasValidEndTime() bool`

HasValidEndTime returns a boolean if a field has been set.

### GetTempValidities

`func (o *TrafficInfluData) GetTempValidities() []TemporalValidity`

GetTempValidities returns the TempValidities field if non-nil, zero value otherwise.

### GetTempValiditiesOk

`func (o *TrafficInfluData) GetTempValiditiesOk() (*[]TemporalValidity, bool)`

GetTempValiditiesOk returns a tuple with the TempValidities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempValidities

`func (o *TrafficInfluData) SetTempValidities(v []TemporalValidity)`

SetTempValidities sets TempValidities field to given value.

### HasTempValidities

`func (o *TrafficInfluData) HasTempValidities() bool`

HasTempValidities returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *TrafficInfluData) GetNwAreaInfo() NetworkAreaInfo1`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *TrafficInfluData) GetNwAreaInfoOk() (*NetworkAreaInfo1, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *TrafficInfluData) SetNwAreaInfo(v NetworkAreaInfo1)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *TrafficInfluData) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetUpPathChgNotifUri

`func (o *TrafficInfluData) GetUpPathChgNotifUri() string`

GetUpPathChgNotifUri returns the UpPathChgNotifUri field if non-nil, zero value otherwise.

### GetUpPathChgNotifUriOk

`func (o *TrafficInfluData) GetUpPathChgNotifUriOk() (*string, bool)`

GetUpPathChgNotifUriOk returns a tuple with the UpPathChgNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgNotifUri

`func (o *TrafficInfluData) SetUpPathChgNotifUri(v string)`

SetUpPathChgNotifUri sets UpPathChgNotifUri field to given value.

### HasUpPathChgNotifUri

`func (o *TrafficInfluData) HasUpPathChgNotifUri() bool`

HasUpPathChgNotifUri returns a boolean if a field has been set.

### GetHeaders

`func (o *TrafficInfluData) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *TrafficInfluData) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *TrafficInfluData) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *TrafficInfluData) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSubscribedEvents

`func (o *TrafficInfluData) GetSubscribedEvents() []SubscribedEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *TrafficInfluData) GetSubscribedEventsOk() (*[]SubscribedEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *TrafficInfluData) SetSubscribedEvents(v []SubscribedEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *TrafficInfluData) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetDnaiChgType

`func (o *TrafficInfluData) GetDnaiChgType() DnaiChangeType`

GetDnaiChgType returns the DnaiChgType field if non-nil, zero value otherwise.

### GetDnaiChgTypeOk

`func (o *TrafficInfluData) GetDnaiChgTypeOk() (*DnaiChangeType, bool)`

GetDnaiChgTypeOk returns a tuple with the DnaiChgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiChgType

`func (o *TrafficInfluData) SetDnaiChgType(v DnaiChangeType)`

SetDnaiChgType sets DnaiChgType field to given value.

### HasDnaiChgType

`func (o *TrafficInfluData) HasDnaiChgType() bool`

HasDnaiChgType returns a boolean if a field has been set.

### GetAfAckInd

`func (o *TrafficInfluData) GetAfAckInd() bool`

GetAfAckInd returns the AfAckInd field if non-nil, zero value otherwise.

### GetAfAckIndOk

`func (o *TrafficInfluData) GetAfAckIndOk() (*bool, bool)`

GetAfAckIndOk returns a tuple with the AfAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAckInd

`func (o *TrafficInfluData) SetAfAckInd(v bool)`

SetAfAckInd sets AfAckInd field to given value.

### HasAfAckInd

`func (o *TrafficInfluData) HasAfAckInd() bool`

HasAfAckInd returns a boolean if a field has been set.

### GetAddrPreserInd

`func (o *TrafficInfluData) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *TrafficInfluData) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *TrafficInfluData) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *TrafficInfluData) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *TrafficInfluData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *TrafficInfluData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *TrafficInfluData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *TrafficInfluData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetResUri

`func (o *TrafficInfluData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *TrafficInfluData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *TrafficInfluData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *TrafficInfluData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


