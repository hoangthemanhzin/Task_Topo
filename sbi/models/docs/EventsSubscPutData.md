# EventsSubscPutData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]AfEventSubscription**](AfEventSubscription.md) |  | 
**NotifUri** | Pointer to **string** |  | [optional] 
**ReqQosMonParams** | Pointer to [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) |  | [optional] 
**QosMon** | Pointer to [**QosMonitoringInformation**](QosMonitoringInformation.md) |  | [optional] 
**ReqAnis** | Pointer to [**[]RequiredAccessInfo**](RequiredAccessInfo.md) |  | [optional] 
**UsgThres** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**NotifCorreId** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AddAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**RelAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**AnChargAddr** | Pointer to [**AccNetChargingAddress**](AccNetChargingAddress.md) |  | [optional] 
**AnChargIds** | Pointer to [**[]AccessNetChargingIdentifier**](AccessNetChargingIdentifier.md) |  | [optional] 
**AnGwAddr** | Pointer to [**AnGwAddress**](AnGwAddress.md) |  | [optional] 
**EvSubsUri** | **string** |  | 
**EvNotifs** | [**[]AfEventNotification**](AfEventNotification.md) |  | 
**FailedResourcAllocReports** | Pointer to [**[]ResourcesAllocationInfo**](ResourcesAllocationInfo.md) |  | [optional] 
**SuccResourcAllocReports** | Pointer to [**[]ResourcesAllocationInfo**](ResourcesAllocationInfo.md) |  | [optional] 
**NoNetLocSupp** | Pointer to [**NetLocAccessSupport**](NetLocAccessSupport.md) |  | [optional] 
**OutOfCredReports** | Pointer to [**[]OutOfCreditInformation**](OutOfCreditInformation.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**QncReports** | Pointer to [**[]QosNotificationControlInfo**](QosNotificationControlInfo.md) |  | [optional] 
**QosMonReports** | Pointer to [**[]QosMonitoringReport**](QosMonitoringReport.md) |  | [optional] 
**RanNasRelCauses** | Pointer to [**[]RanNasRelCause**](RanNasRelCause.md) | Contains the RAN and/or NAS release cause. | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**UeLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeLocTime** | Pointer to **time.Time** |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**UsgRep** | Pointer to [**AccumulatedUsage**](AccumulatedUsage.md) |  | [optional] 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 

## Methods

### NewEventsSubscPutData

`func NewEventsSubscPutData(events []AfEventSubscription, evSubsUri string, evNotifs []AfEventNotification, ) *EventsSubscPutData`

NewEventsSubscPutData instantiates a new EventsSubscPutData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSubscPutDataWithDefaults

`func NewEventsSubscPutDataWithDefaults() *EventsSubscPutData`

NewEventsSubscPutDataWithDefaults instantiates a new EventsSubscPutData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsSubscPutData) GetEvents() []AfEventSubscription`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSubscPutData) GetEventsOk() (*[]AfEventSubscription, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSubscPutData) SetEvents(v []AfEventSubscription)`

SetEvents sets Events field to given value.


### GetNotifUri

`func (o *EventsSubscPutData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *EventsSubscPutData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *EventsSubscPutData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *EventsSubscPutData) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetReqQosMonParams

`func (o *EventsSubscPutData) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *EventsSubscPutData) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *EventsSubscPutData) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.

### HasReqQosMonParams

`func (o *EventsSubscPutData) HasReqQosMonParams() bool`

HasReqQosMonParams returns a boolean if a field has been set.

### GetQosMon

`func (o *EventsSubscPutData) GetQosMon() QosMonitoringInformation`

GetQosMon returns the QosMon field if non-nil, zero value otherwise.

### GetQosMonOk

`func (o *EventsSubscPutData) GetQosMonOk() (*QosMonitoringInformation, bool)`

GetQosMonOk returns a tuple with the QosMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMon

`func (o *EventsSubscPutData) SetQosMon(v QosMonitoringInformation)`

SetQosMon sets QosMon field to given value.

### HasQosMon

`func (o *EventsSubscPutData) HasQosMon() bool`

HasQosMon returns a boolean if a field has been set.

### GetReqAnis

`func (o *EventsSubscPutData) GetReqAnis() []RequiredAccessInfo`

GetReqAnis returns the ReqAnis field if non-nil, zero value otherwise.

### GetReqAnisOk

`func (o *EventsSubscPutData) GetReqAnisOk() (*[]RequiredAccessInfo, bool)`

GetReqAnisOk returns a tuple with the ReqAnis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAnis

`func (o *EventsSubscPutData) SetReqAnis(v []RequiredAccessInfo)`

SetReqAnis sets ReqAnis field to given value.

### HasReqAnis

`func (o *EventsSubscPutData) HasReqAnis() bool`

HasReqAnis returns a boolean if a field has been set.

### GetUsgThres

`func (o *EventsSubscPutData) GetUsgThres() UsageThreshold`

GetUsgThres returns the UsgThres field if non-nil, zero value otherwise.

### GetUsgThresOk

`func (o *EventsSubscPutData) GetUsgThresOk() (*UsageThreshold, bool)`

GetUsgThresOk returns a tuple with the UsgThres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgThres

`func (o *EventsSubscPutData) SetUsgThres(v UsageThreshold)`

SetUsgThres sets UsgThres field to given value.

### HasUsgThres

`func (o *EventsSubscPutData) HasUsgThres() bool`

HasUsgThres returns a boolean if a field has been set.

### GetNotifCorreId

`func (o *EventsSubscPutData) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *EventsSubscPutData) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *EventsSubscPutData) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.

### HasNotifCorreId

`func (o *EventsSubscPutData) HasNotifCorreId() bool`

HasNotifCorreId returns a boolean if a field has been set.

### GetAccessType

`func (o *EventsSubscPutData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *EventsSubscPutData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *EventsSubscPutData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *EventsSubscPutData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAddAccessInfo

`func (o *EventsSubscPutData) GetAddAccessInfo() AdditionalAccessInfo`

GetAddAccessInfo returns the AddAccessInfo field if non-nil, zero value otherwise.

### GetAddAccessInfoOk

`func (o *EventsSubscPutData) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetAddAccessInfoOk returns a tuple with the AddAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessInfo

`func (o *EventsSubscPutData) SetAddAccessInfo(v AdditionalAccessInfo)`

SetAddAccessInfo sets AddAccessInfo field to given value.

### HasAddAccessInfo

`func (o *EventsSubscPutData) HasAddAccessInfo() bool`

HasAddAccessInfo returns a boolean if a field has been set.

### GetRelAccessInfo

`func (o *EventsSubscPutData) GetRelAccessInfo() AdditionalAccessInfo`

GetRelAccessInfo returns the RelAccessInfo field if non-nil, zero value otherwise.

### GetRelAccessInfoOk

`func (o *EventsSubscPutData) GetRelAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetRelAccessInfoOk returns a tuple with the RelAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelAccessInfo

`func (o *EventsSubscPutData) SetRelAccessInfo(v AdditionalAccessInfo)`

SetRelAccessInfo sets RelAccessInfo field to given value.

### HasRelAccessInfo

`func (o *EventsSubscPutData) HasRelAccessInfo() bool`

HasRelAccessInfo returns a boolean if a field has been set.

### GetAnChargAddr

`func (o *EventsSubscPutData) GetAnChargAddr() AccNetChargingAddress`

GetAnChargAddr returns the AnChargAddr field if non-nil, zero value otherwise.

### GetAnChargAddrOk

`func (o *EventsSubscPutData) GetAnChargAddrOk() (*AccNetChargingAddress, bool)`

GetAnChargAddrOk returns a tuple with the AnChargAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnChargAddr

`func (o *EventsSubscPutData) SetAnChargAddr(v AccNetChargingAddress)`

SetAnChargAddr sets AnChargAddr field to given value.

### HasAnChargAddr

`func (o *EventsSubscPutData) HasAnChargAddr() bool`

HasAnChargAddr returns a boolean if a field has been set.

### GetAnChargIds

`func (o *EventsSubscPutData) GetAnChargIds() []AccessNetChargingIdentifier`

GetAnChargIds returns the AnChargIds field if non-nil, zero value otherwise.

### GetAnChargIdsOk

`func (o *EventsSubscPutData) GetAnChargIdsOk() (*[]AccessNetChargingIdentifier, bool)`

GetAnChargIdsOk returns a tuple with the AnChargIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnChargIds

`func (o *EventsSubscPutData) SetAnChargIds(v []AccessNetChargingIdentifier)`

SetAnChargIds sets AnChargIds field to given value.

### HasAnChargIds

`func (o *EventsSubscPutData) HasAnChargIds() bool`

HasAnChargIds returns a boolean if a field has been set.

### GetAnGwAddr

`func (o *EventsSubscPutData) GetAnGwAddr() AnGwAddress`

GetAnGwAddr returns the AnGwAddr field if non-nil, zero value otherwise.

### GetAnGwAddrOk

`func (o *EventsSubscPutData) GetAnGwAddrOk() (*AnGwAddress, bool)`

GetAnGwAddrOk returns a tuple with the AnGwAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnGwAddr

`func (o *EventsSubscPutData) SetAnGwAddr(v AnGwAddress)`

SetAnGwAddr sets AnGwAddr field to given value.

### HasAnGwAddr

`func (o *EventsSubscPutData) HasAnGwAddr() bool`

HasAnGwAddr returns a boolean if a field has been set.

### GetEvSubsUri

`func (o *EventsSubscPutData) GetEvSubsUri() string`

GetEvSubsUri returns the EvSubsUri field if non-nil, zero value otherwise.

### GetEvSubsUriOk

`func (o *EventsSubscPutData) GetEvSubsUriOk() (*string, bool)`

GetEvSubsUriOk returns a tuple with the EvSubsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsUri

`func (o *EventsSubscPutData) SetEvSubsUri(v string)`

SetEvSubsUri sets EvSubsUri field to given value.


### GetEvNotifs

`func (o *EventsSubscPutData) GetEvNotifs() []AfEventNotification`

GetEvNotifs returns the EvNotifs field if non-nil, zero value otherwise.

### GetEvNotifsOk

`func (o *EventsSubscPutData) GetEvNotifsOk() (*[]AfEventNotification, bool)`

GetEvNotifsOk returns a tuple with the EvNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvNotifs

`func (o *EventsSubscPutData) SetEvNotifs(v []AfEventNotification)`

SetEvNotifs sets EvNotifs field to given value.


### GetFailedResourcAllocReports

`func (o *EventsSubscPutData) GetFailedResourcAllocReports() []ResourcesAllocationInfo`

GetFailedResourcAllocReports returns the FailedResourcAllocReports field if non-nil, zero value otherwise.

### GetFailedResourcAllocReportsOk

`func (o *EventsSubscPutData) GetFailedResourcAllocReportsOk() (*[]ResourcesAllocationInfo, bool)`

GetFailedResourcAllocReportsOk returns a tuple with the FailedResourcAllocReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedResourcAllocReports

`func (o *EventsSubscPutData) SetFailedResourcAllocReports(v []ResourcesAllocationInfo)`

SetFailedResourcAllocReports sets FailedResourcAllocReports field to given value.

### HasFailedResourcAllocReports

`func (o *EventsSubscPutData) HasFailedResourcAllocReports() bool`

HasFailedResourcAllocReports returns a boolean if a field has been set.

### GetSuccResourcAllocReports

`func (o *EventsSubscPutData) GetSuccResourcAllocReports() []ResourcesAllocationInfo`

GetSuccResourcAllocReports returns the SuccResourcAllocReports field if non-nil, zero value otherwise.

### GetSuccResourcAllocReportsOk

`func (o *EventsSubscPutData) GetSuccResourcAllocReportsOk() (*[]ResourcesAllocationInfo, bool)`

GetSuccResourcAllocReportsOk returns a tuple with the SuccResourcAllocReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccResourcAllocReports

`func (o *EventsSubscPutData) SetSuccResourcAllocReports(v []ResourcesAllocationInfo)`

SetSuccResourcAllocReports sets SuccResourcAllocReports field to given value.

### HasSuccResourcAllocReports

`func (o *EventsSubscPutData) HasSuccResourcAllocReports() bool`

HasSuccResourcAllocReports returns a boolean if a field has been set.

### GetNoNetLocSupp

`func (o *EventsSubscPutData) GetNoNetLocSupp() NetLocAccessSupport`

GetNoNetLocSupp returns the NoNetLocSupp field if non-nil, zero value otherwise.

### GetNoNetLocSuppOk

`func (o *EventsSubscPutData) GetNoNetLocSuppOk() (*NetLocAccessSupport, bool)`

GetNoNetLocSuppOk returns a tuple with the NoNetLocSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNetLocSupp

`func (o *EventsSubscPutData) SetNoNetLocSupp(v NetLocAccessSupport)`

SetNoNetLocSupp sets NoNetLocSupp field to given value.

### HasNoNetLocSupp

`func (o *EventsSubscPutData) HasNoNetLocSupp() bool`

HasNoNetLocSupp returns a boolean if a field has been set.

### GetOutOfCredReports

`func (o *EventsSubscPutData) GetOutOfCredReports() []OutOfCreditInformation`

GetOutOfCredReports returns the OutOfCredReports field if non-nil, zero value otherwise.

### GetOutOfCredReportsOk

`func (o *EventsSubscPutData) GetOutOfCredReportsOk() (*[]OutOfCreditInformation, bool)`

GetOutOfCredReportsOk returns a tuple with the OutOfCredReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfCredReports

`func (o *EventsSubscPutData) SetOutOfCredReports(v []OutOfCreditInformation)`

SetOutOfCredReports sets OutOfCredReports field to given value.

### HasOutOfCredReports

`func (o *EventsSubscPutData) HasOutOfCredReports() bool`

HasOutOfCredReports returns a boolean if a field has been set.

### GetPlmnId

`func (o *EventsSubscPutData) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *EventsSubscPutData) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *EventsSubscPutData) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *EventsSubscPutData) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetQncReports

`func (o *EventsSubscPutData) GetQncReports() []QosNotificationControlInfo`

GetQncReports returns the QncReports field if non-nil, zero value otherwise.

### GetQncReportsOk

`func (o *EventsSubscPutData) GetQncReportsOk() (*[]QosNotificationControlInfo, bool)`

GetQncReportsOk returns a tuple with the QncReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQncReports

`func (o *EventsSubscPutData) SetQncReports(v []QosNotificationControlInfo)`

SetQncReports sets QncReports field to given value.

### HasQncReports

`func (o *EventsSubscPutData) HasQncReports() bool`

HasQncReports returns a boolean if a field has been set.

### GetQosMonReports

`func (o *EventsSubscPutData) GetQosMonReports() []QosMonitoringReport`

GetQosMonReports returns the QosMonReports field if non-nil, zero value otherwise.

### GetQosMonReportsOk

`func (o *EventsSubscPutData) GetQosMonReportsOk() (*[]QosMonitoringReport, bool)`

GetQosMonReportsOk returns a tuple with the QosMonReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonReports

`func (o *EventsSubscPutData) SetQosMonReports(v []QosMonitoringReport)`

SetQosMonReports sets QosMonReports field to given value.

### HasQosMonReports

`func (o *EventsSubscPutData) HasQosMonReports() bool`

HasQosMonReports returns a boolean if a field has been set.

### GetRanNasRelCauses

`func (o *EventsSubscPutData) GetRanNasRelCauses() []RanNasRelCause`

GetRanNasRelCauses returns the RanNasRelCauses field if non-nil, zero value otherwise.

### GetRanNasRelCausesOk

`func (o *EventsSubscPutData) GetRanNasRelCausesOk() (*[]RanNasRelCause, bool)`

GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNasRelCauses

`func (o *EventsSubscPutData) SetRanNasRelCauses(v []RanNasRelCause)`

SetRanNasRelCauses sets RanNasRelCauses field to given value.

### HasRanNasRelCauses

`func (o *EventsSubscPutData) HasRanNasRelCauses() bool`

HasRanNasRelCauses returns a boolean if a field has been set.

### GetRatType

`func (o *EventsSubscPutData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *EventsSubscPutData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *EventsSubscPutData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *EventsSubscPutData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetUeLoc

`func (o *EventsSubscPutData) GetUeLoc() UserLocation`

GetUeLoc returns the UeLoc field if non-nil, zero value otherwise.

### GetUeLocOk

`func (o *EventsSubscPutData) GetUeLocOk() (*UserLocation, bool)`

GetUeLocOk returns a tuple with the UeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLoc

`func (o *EventsSubscPutData) SetUeLoc(v UserLocation)`

SetUeLoc sets UeLoc field to given value.

### HasUeLoc

`func (o *EventsSubscPutData) HasUeLoc() bool`

HasUeLoc returns a boolean if a field has been set.

### GetUeLocTime

`func (o *EventsSubscPutData) GetUeLocTime() time.Time`

GetUeLocTime returns the UeLocTime field if non-nil, zero value otherwise.

### GetUeLocTimeOk

`func (o *EventsSubscPutData) GetUeLocTimeOk() (*time.Time, bool)`

GetUeLocTimeOk returns a tuple with the UeLocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocTime

`func (o *EventsSubscPutData) SetUeLocTime(v time.Time)`

SetUeLocTime sets UeLocTime field to given value.

### HasUeLocTime

`func (o *EventsSubscPutData) HasUeLocTime() bool`

HasUeLocTime returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *EventsSubscPutData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *EventsSubscPutData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *EventsSubscPutData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *EventsSubscPutData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetUsgRep

`func (o *EventsSubscPutData) GetUsgRep() AccumulatedUsage`

GetUsgRep returns the UsgRep field if non-nil, zero value otherwise.

### GetUsgRepOk

`func (o *EventsSubscPutData) GetUsgRepOk() (*AccumulatedUsage, bool)`

GetUsgRepOk returns a tuple with the UsgRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgRep

`func (o *EventsSubscPutData) SetUsgRep(v AccumulatedUsage)`

SetUsgRep sets UsgRep field to given value.

### HasUsgRep

`func (o *EventsSubscPutData) HasUsgRep() bool`

HasUsgRep returns a boolean if a field has been set.

### GetTsnBridgeManCont

`func (o *EventsSubscPutData) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *EventsSubscPutData) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *EventsSubscPutData) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *EventsSubscPutData) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *EventsSubscPutData) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *EventsSubscPutData) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *EventsSubscPutData) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *EventsSubscPutData) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *EventsSubscPutData) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *EventsSubscPutData) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *EventsSubscPutData) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *EventsSubscPutData) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


