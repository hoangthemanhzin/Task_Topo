# EventsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEventsNotification

`func NewEventsNotification(evSubsUri string, evNotifs []AfEventNotification, ) *EventsNotification`

NewEventsNotification instantiates a new EventsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsNotificationWithDefaults

`func NewEventsNotificationWithDefaults() *EventsNotification`

NewEventsNotificationWithDefaults instantiates a new EventsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *EventsNotification) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *EventsNotification) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *EventsNotification) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *EventsNotification) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAddAccessInfo

`func (o *EventsNotification) GetAddAccessInfo() AdditionalAccessInfo`

GetAddAccessInfo returns the AddAccessInfo field if non-nil, zero value otherwise.

### GetAddAccessInfoOk

`func (o *EventsNotification) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetAddAccessInfoOk returns a tuple with the AddAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessInfo

`func (o *EventsNotification) SetAddAccessInfo(v AdditionalAccessInfo)`

SetAddAccessInfo sets AddAccessInfo field to given value.

### HasAddAccessInfo

`func (o *EventsNotification) HasAddAccessInfo() bool`

HasAddAccessInfo returns a boolean if a field has been set.

### GetRelAccessInfo

`func (o *EventsNotification) GetRelAccessInfo() AdditionalAccessInfo`

GetRelAccessInfo returns the RelAccessInfo field if non-nil, zero value otherwise.

### GetRelAccessInfoOk

`func (o *EventsNotification) GetRelAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetRelAccessInfoOk returns a tuple with the RelAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelAccessInfo

`func (o *EventsNotification) SetRelAccessInfo(v AdditionalAccessInfo)`

SetRelAccessInfo sets RelAccessInfo field to given value.

### HasRelAccessInfo

`func (o *EventsNotification) HasRelAccessInfo() bool`

HasRelAccessInfo returns a boolean if a field has been set.

### GetAnChargAddr

`func (o *EventsNotification) GetAnChargAddr() AccNetChargingAddress`

GetAnChargAddr returns the AnChargAddr field if non-nil, zero value otherwise.

### GetAnChargAddrOk

`func (o *EventsNotification) GetAnChargAddrOk() (*AccNetChargingAddress, bool)`

GetAnChargAddrOk returns a tuple with the AnChargAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnChargAddr

`func (o *EventsNotification) SetAnChargAddr(v AccNetChargingAddress)`

SetAnChargAddr sets AnChargAddr field to given value.

### HasAnChargAddr

`func (o *EventsNotification) HasAnChargAddr() bool`

HasAnChargAddr returns a boolean if a field has been set.

### GetAnChargIds

`func (o *EventsNotification) GetAnChargIds() []AccessNetChargingIdentifier`

GetAnChargIds returns the AnChargIds field if non-nil, zero value otherwise.

### GetAnChargIdsOk

`func (o *EventsNotification) GetAnChargIdsOk() (*[]AccessNetChargingIdentifier, bool)`

GetAnChargIdsOk returns a tuple with the AnChargIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnChargIds

`func (o *EventsNotification) SetAnChargIds(v []AccessNetChargingIdentifier)`

SetAnChargIds sets AnChargIds field to given value.

### HasAnChargIds

`func (o *EventsNotification) HasAnChargIds() bool`

HasAnChargIds returns a boolean if a field has been set.

### GetAnGwAddr

`func (o *EventsNotification) GetAnGwAddr() AnGwAddress`

GetAnGwAddr returns the AnGwAddr field if non-nil, zero value otherwise.

### GetAnGwAddrOk

`func (o *EventsNotification) GetAnGwAddrOk() (*AnGwAddress, bool)`

GetAnGwAddrOk returns a tuple with the AnGwAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnGwAddr

`func (o *EventsNotification) SetAnGwAddr(v AnGwAddress)`

SetAnGwAddr sets AnGwAddr field to given value.

### HasAnGwAddr

`func (o *EventsNotification) HasAnGwAddr() bool`

HasAnGwAddr returns a boolean if a field has been set.

### GetEvSubsUri

`func (o *EventsNotification) GetEvSubsUri() string`

GetEvSubsUri returns the EvSubsUri field if non-nil, zero value otherwise.

### GetEvSubsUriOk

`func (o *EventsNotification) GetEvSubsUriOk() (*string, bool)`

GetEvSubsUriOk returns a tuple with the EvSubsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsUri

`func (o *EventsNotification) SetEvSubsUri(v string)`

SetEvSubsUri sets EvSubsUri field to given value.


### GetEvNotifs

`func (o *EventsNotification) GetEvNotifs() []AfEventNotification`

GetEvNotifs returns the EvNotifs field if non-nil, zero value otherwise.

### GetEvNotifsOk

`func (o *EventsNotification) GetEvNotifsOk() (*[]AfEventNotification, bool)`

GetEvNotifsOk returns a tuple with the EvNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvNotifs

`func (o *EventsNotification) SetEvNotifs(v []AfEventNotification)`

SetEvNotifs sets EvNotifs field to given value.


### GetFailedResourcAllocReports

`func (o *EventsNotification) GetFailedResourcAllocReports() []ResourcesAllocationInfo`

GetFailedResourcAllocReports returns the FailedResourcAllocReports field if non-nil, zero value otherwise.

### GetFailedResourcAllocReportsOk

`func (o *EventsNotification) GetFailedResourcAllocReportsOk() (*[]ResourcesAllocationInfo, bool)`

GetFailedResourcAllocReportsOk returns a tuple with the FailedResourcAllocReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedResourcAllocReports

`func (o *EventsNotification) SetFailedResourcAllocReports(v []ResourcesAllocationInfo)`

SetFailedResourcAllocReports sets FailedResourcAllocReports field to given value.

### HasFailedResourcAllocReports

`func (o *EventsNotification) HasFailedResourcAllocReports() bool`

HasFailedResourcAllocReports returns a boolean if a field has been set.

### GetSuccResourcAllocReports

`func (o *EventsNotification) GetSuccResourcAllocReports() []ResourcesAllocationInfo`

GetSuccResourcAllocReports returns the SuccResourcAllocReports field if non-nil, zero value otherwise.

### GetSuccResourcAllocReportsOk

`func (o *EventsNotification) GetSuccResourcAllocReportsOk() (*[]ResourcesAllocationInfo, bool)`

GetSuccResourcAllocReportsOk returns a tuple with the SuccResourcAllocReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccResourcAllocReports

`func (o *EventsNotification) SetSuccResourcAllocReports(v []ResourcesAllocationInfo)`

SetSuccResourcAllocReports sets SuccResourcAllocReports field to given value.

### HasSuccResourcAllocReports

`func (o *EventsNotification) HasSuccResourcAllocReports() bool`

HasSuccResourcAllocReports returns a boolean if a field has been set.

### GetNoNetLocSupp

`func (o *EventsNotification) GetNoNetLocSupp() NetLocAccessSupport`

GetNoNetLocSupp returns the NoNetLocSupp field if non-nil, zero value otherwise.

### GetNoNetLocSuppOk

`func (o *EventsNotification) GetNoNetLocSuppOk() (*NetLocAccessSupport, bool)`

GetNoNetLocSuppOk returns a tuple with the NoNetLocSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNetLocSupp

`func (o *EventsNotification) SetNoNetLocSupp(v NetLocAccessSupport)`

SetNoNetLocSupp sets NoNetLocSupp field to given value.

### HasNoNetLocSupp

`func (o *EventsNotification) HasNoNetLocSupp() bool`

HasNoNetLocSupp returns a boolean if a field has been set.

### GetOutOfCredReports

`func (o *EventsNotification) GetOutOfCredReports() []OutOfCreditInformation`

GetOutOfCredReports returns the OutOfCredReports field if non-nil, zero value otherwise.

### GetOutOfCredReportsOk

`func (o *EventsNotification) GetOutOfCredReportsOk() (*[]OutOfCreditInformation, bool)`

GetOutOfCredReportsOk returns a tuple with the OutOfCredReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfCredReports

`func (o *EventsNotification) SetOutOfCredReports(v []OutOfCreditInformation)`

SetOutOfCredReports sets OutOfCredReports field to given value.

### HasOutOfCredReports

`func (o *EventsNotification) HasOutOfCredReports() bool`

HasOutOfCredReports returns a boolean if a field has been set.

### GetPlmnId

`func (o *EventsNotification) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *EventsNotification) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *EventsNotification) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *EventsNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetQncReports

`func (o *EventsNotification) GetQncReports() []QosNotificationControlInfo`

GetQncReports returns the QncReports field if non-nil, zero value otherwise.

### GetQncReportsOk

`func (o *EventsNotification) GetQncReportsOk() (*[]QosNotificationControlInfo, bool)`

GetQncReportsOk returns a tuple with the QncReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQncReports

`func (o *EventsNotification) SetQncReports(v []QosNotificationControlInfo)`

SetQncReports sets QncReports field to given value.

### HasQncReports

`func (o *EventsNotification) HasQncReports() bool`

HasQncReports returns a boolean if a field has been set.

### GetQosMonReports

`func (o *EventsNotification) GetQosMonReports() []QosMonitoringReport`

GetQosMonReports returns the QosMonReports field if non-nil, zero value otherwise.

### GetQosMonReportsOk

`func (o *EventsNotification) GetQosMonReportsOk() (*[]QosMonitoringReport, bool)`

GetQosMonReportsOk returns a tuple with the QosMonReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonReports

`func (o *EventsNotification) SetQosMonReports(v []QosMonitoringReport)`

SetQosMonReports sets QosMonReports field to given value.

### HasQosMonReports

`func (o *EventsNotification) HasQosMonReports() bool`

HasQosMonReports returns a boolean if a field has been set.

### GetRanNasRelCauses

`func (o *EventsNotification) GetRanNasRelCauses() []RanNasRelCause`

GetRanNasRelCauses returns the RanNasRelCauses field if non-nil, zero value otherwise.

### GetRanNasRelCausesOk

`func (o *EventsNotification) GetRanNasRelCausesOk() (*[]RanNasRelCause, bool)`

GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNasRelCauses

`func (o *EventsNotification) SetRanNasRelCauses(v []RanNasRelCause)`

SetRanNasRelCauses sets RanNasRelCauses field to given value.

### HasRanNasRelCauses

`func (o *EventsNotification) HasRanNasRelCauses() bool`

HasRanNasRelCauses returns a boolean if a field has been set.

### GetRatType

`func (o *EventsNotification) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *EventsNotification) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *EventsNotification) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *EventsNotification) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetUeLoc

`func (o *EventsNotification) GetUeLoc() UserLocation`

GetUeLoc returns the UeLoc field if non-nil, zero value otherwise.

### GetUeLocOk

`func (o *EventsNotification) GetUeLocOk() (*UserLocation, bool)`

GetUeLocOk returns a tuple with the UeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLoc

`func (o *EventsNotification) SetUeLoc(v UserLocation)`

SetUeLoc sets UeLoc field to given value.

### HasUeLoc

`func (o *EventsNotification) HasUeLoc() bool`

HasUeLoc returns a boolean if a field has been set.

### GetUeLocTime

`func (o *EventsNotification) GetUeLocTime() time.Time`

GetUeLocTime returns the UeLocTime field if non-nil, zero value otherwise.

### GetUeLocTimeOk

`func (o *EventsNotification) GetUeLocTimeOk() (*time.Time, bool)`

GetUeLocTimeOk returns a tuple with the UeLocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocTime

`func (o *EventsNotification) SetUeLocTime(v time.Time)`

SetUeLocTime sets UeLocTime field to given value.

### HasUeLocTime

`func (o *EventsNotification) HasUeLocTime() bool`

HasUeLocTime returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *EventsNotification) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *EventsNotification) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *EventsNotification) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *EventsNotification) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetUsgRep

`func (o *EventsNotification) GetUsgRep() AccumulatedUsage`

GetUsgRep returns the UsgRep field if non-nil, zero value otherwise.

### GetUsgRepOk

`func (o *EventsNotification) GetUsgRepOk() (*AccumulatedUsage, bool)`

GetUsgRepOk returns a tuple with the UsgRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgRep

`func (o *EventsNotification) SetUsgRep(v AccumulatedUsage)`

SetUsgRep sets UsgRep field to given value.

### HasUsgRep

`func (o *EventsNotification) HasUsgRep() bool`

HasUsgRep returns a boolean if a field has been set.

### GetTsnBridgeManCont

`func (o *EventsNotification) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *EventsNotification) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *EventsNotification) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *EventsNotification) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *EventsNotification) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *EventsNotification) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *EventsNotification) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *EventsNotification) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *EventsNotification) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *EventsNotification) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *EventsNotification) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *EventsNotification) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


