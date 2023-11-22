# VsmfUpdatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosFlowsAddModList** | Pointer to [**[]QosFlowItem**](QosFlowItem.md) |  | [optional] 
**QosFlowsRelList** | Pointer to [**[]QosFlowItem**](QosFlowItem.md) |  | [optional] 
**QosFlowsFailedtoAddModList** | Pointer to [**[]QosFlowItem**](QosFlowItem.md) |  | [optional] 
**QosFlowsFailedtoRelList** | Pointer to [**[]QosFlowItem**](QosFlowItem.md) |  | [optional] 
**N1SmInfoFromUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UnknownN1SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** |  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**AssignedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**FailedToAssignEbiList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 
**ReleasedEbiList** | Pointer to **[]int32** |  | [optional] 
**SecondaryRatUsageReport** | Pointer to [**[]SecondaryRatUsageReport**](SecondaryRatUsageReport.md) |  | [optional] 
**SecondaryRatUsageInfo** | Pointer to [**[]SecondaryRatUsageInfo**](SecondaryRatUsageInfo.md) |  | [optional] 
**N4Info** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt1** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 
**N4InfoExt2** | Pointer to [**N4Information**](N4Information.md) |  | [optional] 

## Methods

### NewVsmfUpdatedData

`func NewVsmfUpdatedData() *VsmfUpdatedData`

NewVsmfUpdatedData instantiates a new VsmfUpdatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVsmfUpdatedDataWithDefaults

`func NewVsmfUpdatedDataWithDefaults() *VsmfUpdatedData`

NewVsmfUpdatedDataWithDefaults instantiates a new VsmfUpdatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosFlowsAddModList

`func (o *VsmfUpdatedData) GetQosFlowsAddModList() []QosFlowItem`

GetQosFlowsAddModList returns the QosFlowsAddModList field if non-nil, zero value otherwise.

### GetQosFlowsAddModListOk

`func (o *VsmfUpdatedData) GetQosFlowsAddModListOk() (*[]QosFlowItem, bool)`

GetQosFlowsAddModListOk returns a tuple with the QosFlowsAddModList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsAddModList

`func (o *VsmfUpdatedData) SetQosFlowsAddModList(v []QosFlowItem)`

SetQosFlowsAddModList sets QosFlowsAddModList field to given value.

### HasQosFlowsAddModList

`func (o *VsmfUpdatedData) HasQosFlowsAddModList() bool`

HasQosFlowsAddModList returns a boolean if a field has been set.

### GetQosFlowsRelList

`func (o *VsmfUpdatedData) GetQosFlowsRelList() []QosFlowItem`

GetQosFlowsRelList returns the QosFlowsRelList field if non-nil, zero value otherwise.

### GetQosFlowsRelListOk

`func (o *VsmfUpdatedData) GetQosFlowsRelListOk() (*[]QosFlowItem, bool)`

GetQosFlowsRelListOk returns a tuple with the QosFlowsRelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsRelList

`func (o *VsmfUpdatedData) SetQosFlowsRelList(v []QosFlowItem)`

SetQosFlowsRelList sets QosFlowsRelList field to given value.

### HasQosFlowsRelList

`func (o *VsmfUpdatedData) HasQosFlowsRelList() bool`

HasQosFlowsRelList returns a boolean if a field has been set.

### GetQosFlowsFailedtoAddModList

`func (o *VsmfUpdatedData) GetQosFlowsFailedtoAddModList() []QosFlowItem`

GetQosFlowsFailedtoAddModList returns the QosFlowsFailedtoAddModList field if non-nil, zero value otherwise.

### GetQosFlowsFailedtoAddModListOk

`func (o *VsmfUpdatedData) GetQosFlowsFailedtoAddModListOk() (*[]QosFlowItem, bool)`

GetQosFlowsFailedtoAddModListOk returns a tuple with the QosFlowsFailedtoAddModList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsFailedtoAddModList

`func (o *VsmfUpdatedData) SetQosFlowsFailedtoAddModList(v []QosFlowItem)`

SetQosFlowsFailedtoAddModList sets QosFlowsFailedtoAddModList field to given value.

### HasQosFlowsFailedtoAddModList

`func (o *VsmfUpdatedData) HasQosFlowsFailedtoAddModList() bool`

HasQosFlowsFailedtoAddModList returns a boolean if a field has been set.

### GetQosFlowsFailedtoRelList

`func (o *VsmfUpdatedData) GetQosFlowsFailedtoRelList() []QosFlowItem`

GetQosFlowsFailedtoRelList returns the QosFlowsFailedtoRelList field if non-nil, zero value otherwise.

### GetQosFlowsFailedtoRelListOk

`func (o *VsmfUpdatedData) GetQosFlowsFailedtoRelListOk() (*[]QosFlowItem, bool)`

GetQosFlowsFailedtoRelListOk returns a tuple with the QosFlowsFailedtoRelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowsFailedtoRelList

`func (o *VsmfUpdatedData) SetQosFlowsFailedtoRelList(v []QosFlowItem)`

SetQosFlowsFailedtoRelList sets QosFlowsFailedtoRelList field to given value.

### HasQosFlowsFailedtoRelList

`func (o *VsmfUpdatedData) HasQosFlowsFailedtoRelList() bool`

HasQosFlowsFailedtoRelList returns a boolean if a field has been set.

### GetN1SmInfoFromUe

`func (o *VsmfUpdatedData) GetN1SmInfoFromUe() RefToBinaryData`

GetN1SmInfoFromUe returns the N1SmInfoFromUe field if non-nil, zero value otherwise.

### GetN1SmInfoFromUeOk

`func (o *VsmfUpdatedData) GetN1SmInfoFromUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoFromUeOk returns a tuple with the N1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoFromUe

`func (o *VsmfUpdatedData) SetN1SmInfoFromUe(v RefToBinaryData)`

SetN1SmInfoFromUe sets N1SmInfoFromUe field to given value.

### HasN1SmInfoFromUe

`func (o *VsmfUpdatedData) HasN1SmInfoFromUe() bool`

HasN1SmInfoFromUe returns a boolean if a field has been set.

### GetUnknownN1SmInfo

`func (o *VsmfUpdatedData) GetUnknownN1SmInfo() RefToBinaryData`

GetUnknownN1SmInfo returns the UnknownN1SmInfo field if non-nil, zero value otherwise.

### GetUnknownN1SmInfoOk

`func (o *VsmfUpdatedData) GetUnknownN1SmInfoOk() (*RefToBinaryData, bool)`

GetUnknownN1SmInfoOk returns a tuple with the UnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownN1SmInfo

`func (o *VsmfUpdatedData) SetUnknownN1SmInfo(v RefToBinaryData)`

SetUnknownN1SmInfo sets UnknownN1SmInfo field to given value.

### HasUnknownN1SmInfo

`func (o *VsmfUpdatedData) HasUnknownN1SmInfo() bool`

HasUnknownN1SmInfo returns a boolean if a field has been set.

### GetUeLocation

`func (o *VsmfUpdatedData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *VsmfUpdatedData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *VsmfUpdatedData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *VsmfUpdatedData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *VsmfUpdatedData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *VsmfUpdatedData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *VsmfUpdatedData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *VsmfUpdatedData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *VsmfUpdatedData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *VsmfUpdatedData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *VsmfUpdatedData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *VsmfUpdatedData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetAssignedEbiList

`func (o *VsmfUpdatedData) GetAssignedEbiList() []EbiArpMapping`

GetAssignedEbiList returns the AssignedEbiList field if non-nil, zero value otherwise.

### GetAssignedEbiListOk

`func (o *VsmfUpdatedData) GetAssignedEbiListOk() (*[]EbiArpMapping, bool)`

GetAssignedEbiListOk returns a tuple with the AssignedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEbiList

`func (o *VsmfUpdatedData) SetAssignedEbiList(v []EbiArpMapping)`

SetAssignedEbiList sets AssignedEbiList field to given value.

### HasAssignedEbiList

`func (o *VsmfUpdatedData) HasAssignedEbiList() bool`

HasAssignedEbiList returns a boolean if a field has been set.

### GetFailedToAssignEbiList

`func (o *VsmfUpdatedData) GetFailedToAssignEbiList() []Arp`

GetFailedToAssignEbiList returns the FailedToAssignEbiList field if non-nil, zero value otherwise.

### GetFailedToAssignEbiListOk

`func (o *VsmfUpdatedData) GetFailedToAssignEbiListOk() (*[]Arp, bool)`

GetFailedToAssignEbiListOk returns a tuple with the FailedToAssignEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToAssignEbiList

`func (o *VsmfUpdatedData) SetFailedToAssignEbiList(v []Arp)`

SetFailedToAssignEbiList sets FailedToAssignEbiList field to given value.

### HasFailedToAssignEbiList

`func (o *VsmfUpdatedData) HasFailedToAssignEbiList() bool`

HasFailedToAssignEbiList returns a boolean if a field has been set.

### GetReleasedEbiList

`func (o *VsmfUpdatedData) GetReleasedEbiList() []int32`

GetReleasedEbiList returns the ReleasedEbiList field if non-nil, zero value otherwise.

### GetReleasedEbiListOk

`func (o *VsmfUpdatedData) GetReleasedEbiListOk() (*[]int32, bool)`

GetReleasedEbiListOk returns a tuple with the ReleasedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedEbiList

`func (o *VsmfUpdatedData) SetReleasedEbiList(v []int32)`

SetReleasedEbiList sets ReleasedEbiList field to given value.

### HasReleasedEbiList

`func (o *VsmfUpdatedData) HasReleasedEbiList() bool`

HasReleasedEbiList returns a boolean if a field has been set.

### GetSecondaryRatUsageReport

`func (o *VsmfUpdatedData) GetSecondaryRatUsageReport() []SecondaryRatUsageReport`

GetSecondaryRatUsageReport returns the SecondaryRatUsageReport field if non-nil, zero value otherwise.

### GetSecondaryRatUsageReportOk

`func (o *VsmfUpdatedData) GetSecondaryRatUsageReportOk() (*[]SecondaryRatUsageReport, bool)`

GetSecondaryRatUsageReportOk returns a tuple with the SecondaryRatUsageReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageReport

`func (o *VsmfUpdatedData) SetSecondaryRatUsageReport(v []SecondaryRatUsageReport)`

SetSecondaryRatUsageReport sets SecondaryRatUsageReport field to given value.

### HasSecondaryRatUsageReport

`func (o *VsmfUpdatedData) HasSecondaryRatUsageReport() bool`

HasSecondaryRatUsageReport returns a boolean if a field has been set.

### GetSecondaryRatUsageInfo

`func (o *VsmfUpdatedData) GetSecondaryRatUsageInfo() []SecondaryRatUsageInfo`

GetSecondaryRatUsageInfo returns the SecondaryRatUsageInfo field if non-nil, zero value otherwise.

### GetSecondaryRatUsageInfoOk

`func (o *VsmfUpdatedData) GetSecondaryRatUsageInfoOk() (*[]SecondaryRatUsageInfo, bool)`

GetSecondaryRatUsageInfoOk returns a tuple with the SecondaryRatUsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageInfo

`func (o *VsmfUpdatedData) SetSecondaryRatUsageInfo(v []SecondaryRatUsageInfo)`

SetSecondaryRatUsageInfo sets SecondaryRatUsageInfo field to given value.

### HasSecondaryRatUsageInfo

`func (o *VsmfUpdatedData) HasSecondaryRatUsageInfo() bool`

HasSecondaryRatUsageInfo returns a boolean if a field has been set.

### GetN4Info

`func (o *VsmfUpdatedData) GetN4Info() N4Information`

GetN4Info returns the N4Info field if non-nil, zero value otherwise.

### GetN4InfoOk

`func (o *VsmfUpdatedData) GetN4InfoOk() (*N4Information, bool)`

GetN4InfoOk returns a tuple with the N4Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4Info

`func (o *VsmfUpdatedData) SetN4Info(v N4Information)`

SetN4Info sets N4Info field to given value.

### HasN4Info

`func (o *VsmfUpdatedData) HasN4Info() bool`

HasN4Info returns a boolean if a field has been set.

### GetN4InfoExt1

`func (o *VsmfUpdatedData) GetN4InfoExt1() N4Information`

GetN4InfoExt1 returns the N4InfoExt1 field if non-nil, zero value otherwise.

### GetN4InfoExt1Ok

`func (o *VsmfUpdatedData) GetN4InfoExt1Ok() (*N4Information, bool)`

GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt1

`func (o *VsmfUpdatedData) SetN4InfoExt1(v N4Information)`

SetN4InfoExt1 sets N4InfoExt1 field to given value.

### HasN4InfoExt1

`func (o *VsmfUpdatedData) HasN4InfoExt1() bool`

HasN4InfoExt1 returns a boolean if a field has been set.

### GetN4InfoExt2

`func (o *VsmfUpdatedData) GetN4InfoExt2() N4Information`

GetN4InfoExt2 returns the N4InfoExt2 field if non-nil, zero value otherwise.

### GetN4InfoExt2Ok

`func (o *VsmfUpdatedData) GetN4InfoExt2Ok() (*N4Information, bool)`

GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4InfoExt2

`func (o *VsmfUpdatedData) SetN4InfoExt2(v N4Information)`

SetN4InfoExt2 sets N4InfoExt2 field to given value.

### HasN4InfoExt2

`func (o *VsmfUpdatedData) HasN4InfoExt2() bool`

HasN4InfoExt2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


