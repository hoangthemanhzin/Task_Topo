# AmfEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AmfEventType**](AmfEventType.md) |  | 
**ImmediateFlag** | Pointer to **bool** |  | [optional] [default to false]
**AreaList** | Pointer to [**[]AmfEventArea**](AmfEventArea.md) |  | [optional] 
**LocationFilterList** | Pointer to [**[]LocationFilter**](LocationFilter.md) |  | [optional] 
**RefId** | Pointer to **int32** |  | [optional] 
**TrafficDescriptorList** | Pointer to [**[]TrafficDescriptor**](TrafficDescriptor.md) |  | [optional] 
**ReportUeReachable** | Pointer to **bool** |  | [optional] [default to false]
**ReachabilityFilter** | Pointer to [**ReachabilityFilter**](ReachabilityFilter.md) |  | [optional] 
**MaxReports** | Pointer to **int32** |  | [optional] 
**MaxResponseTime** | Pointer to **int32** |  | [optional] 
**IdleStatusInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAmfEvent

`func NewAmfEvent(type_ AmfEventType, ) *AmfEvent`

NewAmfEvent instantiates a new AmfEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventWithDefaults

`func NewAmfEventWithDefaults() *AmfEvent`

NewAmfEventWithDefaults instantiates a new AmfEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AmfEvent) GetType() AmfEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmfEvent) GetTypeOk() (*AmfEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmfEvent) SetType(v AmfEventType)`

SetType sets Type field to given value.


### GetImmediateFlag

`func (o *AmfEvent) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *AmfEvent) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *AmfEvent) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *AmfEvent) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.

### GetAreaList

`func (o *AmfEvent) GetAreaList() []AmfEventArea`

GetAreaList returns the AreaList field if non-nil, zero value otherwise.

### GetAreaListOk

`func (o *AmfEvent) GetAreaListOk() (*[]AmfEventArea, bool)`

GetAreaListOk returns a tuple with the AreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaList

`func (o *AmfEvent) SetAreaList(v []AmfEventArea)`

SetAreaList sets AreaList field to given value.

### HasAreaList

`func (o *AmfEvent) HasAreaList() bool`

HasAreaList returns a boolean if a field has been set.

### GetLocationFilterList

`func (o *AmfEvent) GetLocationFilterList() []LocationFilter`

GetLocationFilterList returns the LocationFilterList field if non-nil, zero value otherwise.

### GetLocationFilterListOk

`func (o *AmfEvent) GetLocationFilterListOk() (*[]LocationFilter, bool)`

GetLocationFilterListOk returns a tuple with the LocationFilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFilterList

`func (o *AmfEvent) SetLocationFilterList(v []LocationFilter)`

SetLocationFilterList sets LocationFilterList field to given value.

### HasLocationFilterList

`func (o *AmfEvent) HasLocationFilterList() bool`

HasLocationFilterList returns a boolean if a field has been set.

### GetRefId

`func (o *AmfEvent) GetRefId() int32`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AmfEvent) GetRefIdOk() (*int32, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AmfEvent) SetRefId(v int32)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AmfEvent) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTrafficDescriptorList

`func (o *AmfEvent) GetTrafficDescriptorList() []TrafficDescriptor`

GetTrafficDescriptorList returns the TrafficDescriptorList field if non-nil, zero value otherwise.

### GetTrafficDescriptorListOk

`func (o *AmfEvent) GetTrafficDescriptorListOk() (*[]TrafficDescriptor, bool)`

GetTrafficDescriptorListOk returns a tuple with the TrafficDescriptorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDescriptorList

`func (o *AmfEvent) SetTrafficDescriptorList(v []TrafficDescriptor)`

SetTrafficDescriptorList sets TrafficDescriptorList field to given value.

### HasTrafficDescriptorList

`func (o *AmfEvent) HasTrafficDescriptorList() bool`

HasTrafficDescriptorList returns a boolean if a field has been set.

### GetReportUeReachable

`func (o *AmfEvent) GetReportUeReachable() bool`

GetReportUeReachable returns the ReportUeReachable field if non-nil, zero value otherwise.

### GetReportUeReachableOk

`func (o *AmfEvent) GetReportUeReachableOk() (*bool, bool)`

GetReportUeReachableOk returns a tuple with the ReportUeReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUeReachable

`func (o *AmfEvent) SetReportUeReachable(v bool)`

SetReportUeReachable sets ReportUeReachable field to given value.

### HasReportUeReachable

`func (o *AmfEvent) HasReportUeReachable() bool`

HasReportUeReachable returns a boolean if a field has been set.

### GetReachabilityFilter

`func (o *AmfEvent) GetReachabilityFilter() ReachabilityFilter`

GetReachabilityFilter returns the ReachabilityFilter field if non-nil, zero value otherwise.

### GetReachabilityFilterOk

`func (o *AmfEvent) GetReachabilityFilterOk() (*ReachabilityFilter, bool)`

GetReachabilityFilterOk returns a tuple with the ReachabilityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityFilter

`func (o *AmfEvent) SetReachabilityFilter(v ReachabilityFilter)`

SetReachabilityFilter sets ReachabilityFilter field to given value.

### HasReachabilityFilter

`func (o *AmfEvent) HasReachabilityFilter() bool`

HasReachabilityFilter returns a boolean if a field has been set.

### GetMaxReports

`func (o *AmfEvent) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *AmfEvent) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *AmfEvent) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *AmfEvent) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *AmfEvent) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *AmfEvent) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *AmfEvent) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *AmfEvent) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetIdleStatusInd

`func (o *AmfEvent) GetIdleStatusInd() bool`

GetIdleStatusInd returns the IdleStatusInd field if non-nil, zero value otherwise.

### GetIdleStatusIndOk

`func (o *AmfEvent) GetIdleStatusIndOk() (*bool, bool)`

GetIdleStatusIndOk returns a tuple with the IdleStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusInd

`func (o *AmfEvent) SetIdleStatusInd(v bool)`

SetIdleStatusInd sets IdleStatusInd field to given value.

### HasIdleStatusInd

`func (o *AmfEvent) HasIdleStatusInd() bool`

HasIdleStatusInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


