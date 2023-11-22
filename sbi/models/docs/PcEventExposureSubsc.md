# PcEventExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubs** | [**[]PcEvent**](PcEvent.md) |  | 
**EventsRepInfo** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**FilterDnns** | Pointer to **[]string** |  | [optional] 
**FilterSnssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**FilterServices** | Pointer to [**[]ServiceIdentification**](ServiceIdentification.md) |  | [optional] 
**NotifUri** | **string** |  | 
**NotifId** | **string** |  | 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewPcEventExposureSubsc

`func NewPcEventExposureSubsc(eventSubs []PcEvent, notifUri string, notifId string, ) *PcEventExposureSubsc`

NewPcEventExposureSubsc instantiates a new PcEventExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcEventExposureSubscWithDefaults

`func NewPcEventExposureSubscWithDefaults() *PcEventExposureSubsc`

NewPcEventExposureSubscWithDefaults instantiates a new PcEventExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubs

`func (o *PcEventExposureSubsc) GetEventSubs() []PcEvent`

GetEventSubs returns the EventSubs field if non-nil, zero value otherwise.

### GetEventSubsOk

`func (o *PcEventExposureSubsc) GetEventSubsOk() (*[]PcEvent, bool)`

GetEventSubsOk returns a tuple with the EventSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubs

`func (o *PcEventExposureSubsc) SetEventSubs(v []PcEvent)`

SetEventSubs sets EventSubs field to given value.


### GetEventsRepInfo

`func (o *PcEventExposureSubsc) GetEventsRepInfo() ReportingInformation`

GetEventsRepInfo returns the EventsRepInfo field if non-nil, zero value otherwise.

### GetEventsRepInfoOk

`func (o *PcEventExposureSubsc) GetEventsRepInfoOk() (*ReportingInformation, bool)`

GetEventsRepInfoOk returns a tuple with the EventsRepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRepInfo

`func (o *PcEventExposureSubsc) SetEventsRepInfo(v ReportingInformation)`

SetEventsRepInfo sets EventsRepInfo field to given value.

### HasEventsRepInfo

`func (o *PcEventExposureSubsc) HasEventsRepInfo() bool`

HasEventsRepInfo returns a boolean if a field has been set.

### GetGroupId

`func (o *PcEventExposureSubsc) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PcEventExposureSubsc) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PcEventExposureSubsc) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PcEventExposureSubsc) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetFilterDnns

`func (o *PcEventExposureSubsc) GetFilterDnns() []string`

GetFilterDnns returns the FilterDnns field if non-nil, zero value otherwise.

### GetFilterDnnsOk

`func (o *PcEventExposureSubsc) GetFilterDnnsOk() (*[]string, bool)`

GetFilterDnnsOk returns a tuple with the FilterDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDnns

`func (o *PcEventExposureSubsc) SetFilterDnns(v []string)`

SetFilterDnns sets FilterDnns field to given value.

### HasFilterDnns

`func (o *PcEventExposureSubsc) HasFilterDnns() bool`

HasFilterDnns returns a boolean if a field has been set.

### GetFilterSnssais

`func (o *PcEventExposureSubsc) GetFilterSnssais() []Snssai`

GetFilterSnssais returns the FilterSnssais field if non-nil, zero value otherwise.

### GetFilterSnssaisOk

`func (o *PcEventExposureSubsc) GetFilterSnssaisOk() (*[]Snssai, bool)`

GetFilterSnssaisOk returns a tuple with the FilterSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSnssais

`func (o *PcEventExposureSubsc) SetFilterSnssais(v []Snssai)`

SetFilterSnssais sets FilterSnssais field to given value.

### HasFilterSnssais

`func (o *PcEventExposureSubsc) HasFilterSnssais() bool`

HasFilterSnssais returns a boolean if a field has been set.

### GetFilterServices

`func (o *PcEventExposureSubsc) GetFilterServices() []ServiceIdentification`

GetFilterServices returns the FilterServices field if non-nil, zero value otherwise.

### GetFilterServicesOk

`func (o *PcEventExposureSubsc) GetFilterServicesOk() (*[]ServiceIdentification, bool)`

GetFilterServicesOk returns a tuple with the FilterServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterServices

`func (o *PcEventExposureSubsc) SetFilterServices(v []ServiceIdentification)`

SetFilterServices sets FilterServices field to given value.

### HasFilterServices

`func (o *PcEventExposureSubsc) HasFilterServices() bool`

HasFilterServices returns a boolean if a field has been set.

### GetNotifUri

`func (o *PcEventExposureSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *PcEventExposureSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *PcEventExposureSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetNotifId

`func (o *PcEventExposureSubsc) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *PcEventExposureSubsc) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *PcEventExposureSubsc) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetSuppFeat

`func (o *PcEventExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PcEventExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PcEventExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PcEventExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


