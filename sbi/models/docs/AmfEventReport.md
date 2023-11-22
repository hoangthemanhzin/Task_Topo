# AmfEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AmfEventType**](AmfEventType.md) |  | 
**State** | [**AmfEventState**](AmfEventState.md) |  | 
**TimeStamp** | **time.Time** |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**AnyUe** | Pointer to **bool** |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**AreaList** | Pointer to [**[]AmfEventArea**](AmfEventArea.md) |  | [optional] 
**RefId** | Pointer to **int32** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**AdditionalLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**AccessTypeList** | Pointer to [**[]AccessType**](AccessType.md) |  | [optional] 
**RmInfoList** | Pointer to [**[]RmInfo**](RmInfo.md) |  | [optional] 
**CmInfoList** | Pointer to [**[]CmInfo**](CmInfo.md) |  | [optional] 
**Reachability** | Pointer to [**UeReachability**](UeReachability.md) |  | [optional] 
**CommFailure** | Pointer to [**CommunicationFailure**](CommunicationFailure.md) |  | [optional] 
**LossOfConnectReason** | Pointer to [**LossOfConnectivityReason**](LossOfConnectivityReason.md) |  | [optional] 
**NumberOfUes** | Pointer to **int32** |  | [optional] 
**Var5gsUserStateList** | Pointer to [**[]Model5GsUserStateInfo**](Model5GsUserStateInfo.md) |  | [optional] 
**TypeCode** | Pointer to **string** |  | [optional] 
**RegistrationNumber** | Pointer to **int32** |  | [optional] 
**MaxAvailabilityTime** | Pointer to **time.Time** |  | [optional] 
**UeIdExt** | Pointer to [**[]UEIdExt**](UEIdExt.md) |  | [optional] 
**IdleStatusIndication** | Pointer to [**IdleStatusIndication**](IdleStatusIndication.md) |  | [optional] 

## Methods

### NewAmfEventReport

`func NewAmfEventReport(type_ AmfEventType, state AmfEventState, timeStamp time.Time, ) *AmfEventReport`

NewAmfEventReport instantiates a new AmfEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfEventReportWithDefaults

`func NewAmfEventReportWithDefaults() *AmfEventReport`

NewAmfEventReportWithDefaults instantiates a new AmfEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AmfEventReport) GetType() AmfEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmfEventReport) GetTypeOk() (*AmfEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmfEventReport) SetType(v AmfEventType)`

SetType sets Type field to given value.


### GetState

`func (o *AmfEventReport) GetState() AmfEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AmfEventReport) GetStateOk() (*AmfEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AmfEventReport) SetState(v AmfEventState)`

SetState sets State field to given value.


### GetTimeStamp

`func (o *AmfEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AmfEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AmfEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSubscriptionId

`func (o *AmfEventReport) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AmfEventReport) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AmfEventReport) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AmfEventReport) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetAnyUe

`func (o *AmfEventReport) GetAnyUe() bool`

GetAnyUe returns the AnyUe field if non-nil, zero value otherwise.

### GetAnyUeOk

`func (o *AmfEventReport) GetAnyUeOk() (*bool, bool)`

GetAnyUeOk returns a tuple with the AnyUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUe

`func (o *AmfEventReport) SetAnyUe(v bool)`

SetAnyUe sets AnyUe field to given value.

### HasAnyUe

`func (o *AmfEventReport) HasAnyUe() bool`

HasAnyUe returns a boolean if a field has been set.

### GetSupi

`func (o *AmfEventReport) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AmfEventReport) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AmfEventReport) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AmfEventReport) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetAreaList

`func (o *AmfEventReport) GetAreaList() []AmfEventArea`

GetAreaList returns the AreaList field if non-nil, zero value otherwise.

### GetAreaListOk

`func (o *AmfEventReport) GetAreaListOk() (*[]AmfEventArea, bool)`

GetAreaListOk returns a tuple with the AreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaList

`func (o *AmfEventReport) SetAreaList(v []AmfEventArea)`

SetAreaList sets AreaList field to given value.

### HasAreaList

`func (o *AmfEventReport) HasAreaList() bool`

HasAreaList returns a boolean if a field has been set.

### GetRefId

`func (o *AmfEventReport) GetRefId() int32`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AmfEventReport) GetRefIdOk() (*int32, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AmfEventReport) SetRefId(v int32)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AmfEventReport) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetGpsi

`func (o *AmfEventReport) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AmfEventReport) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AmfEventReport) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AmfEventReport) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPei

`func (o *AmfEventReport) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AmfEventReport) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AmfEventReport) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AmfEventReport) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetLocation

`func (o *AmfEventReport) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AmfEventReport) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AmfEventReport) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AmfEventReport) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAdditionalLocation

`func (o *AmfEventReport) GetAdditionalLocation() UserLocation`

GetAdditionalLocation returns the AdditionalLocation field if non-nil, zero value otherwise.

### GetAdditionalLocationOk

`func (o *AmfEventReport) GetAdditionalLocationOk() (*UserLocation, bool)`

GetAdditionalLocationOk returns a tuple with the AdditionalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLocation

`func (o *AmfEventReport) SetAdditionalLocation(v UserLocation)`

SetAdditionalLocation sets AdditionalLocation field to given value.

### HasAdditionalLocation

`func (o *AmfEventReport) HasAdditionalLocation() bool`

HasAdditionalLocation returns a boolean if a field has been set.

### GetTimezone

`func (o *AmfEventReport) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AmfEventReport) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AmfEventReport) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AmfEventReport) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAccessTypeList

`func (o *AmfEventReport) GetAccessTypeList() []AccessType`

GetAccessTypeList returns the AccessTypeList field if non-nil, zero value otherwise.

### GetAccessTypeListOk

`func (o *AmfEventReport) GetAccessTypeListOk() (*[]AccessType, bool)`

GetAccessTypeListOk returns a tuple with the AccessTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypeList

`func (o *AmfEventReport) SetAccessTypeList(v []AccessType)`

SetAccessTypeList sets AccessTypeList field to given value.

### HasAccessTypeList

`func (o *AmfEventReport) HasAccessTypeList() bool`

HasAccessTypeList returns a boolean if a field has been set.

### GetRmInfoList

`func (o *AmfEventReport) GetRmInfoList() []RmInfo`

GetRmInfoList returns the RmInfoList field if non-nil, zero value otherwise.

### GetRmInfoListOk

`func (o *AmfEventReport) GetRmInfoListOk() (*[]RmInfo, bool)`

GetRmInfoListOk returns a tuple with the RmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmInfoList

`func (o *AmfEventReport) SetRmInfoList(v []RmInfo)`

SetRmInfoList sets RmInfoList field to given value.

### HasRmInfoList

`func (o *AmfEventReport) HasRmInfoList() bool`

HasRmInfoList returns a boolean if a field has been set.

### GetCmInfoList

`func (o *AmfEventReport) GetCmInfoList() []CmInfo`

GetCmInfoList returns the CmInfoList field if non-nil, zero value otherwise.

### GetCmInfoListOk

`func (o *AmfEventReport) GetCmInfoListOk() (*[]CmInfo, bool)`

GetCmInfoListOk returns a tuple with the CmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmInfoList

`func (o *AmfEventReport) SetCmInfoList(v []CmInfo)`

SetCmInfoList sets CmInfoList field to given value.

### HasCmInfoList

`func (o *AmfEventReport) HasCmInfoList() bool`

HasCmInfoList returns a boolean if a field has been set.

### GetReachability

`func (o *AmfEventReport) GetReachability() UeReachability`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *AmfEventReport) GetReachabilityOk() (*UeReachability, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *AmfEventReport) SetReachability(v UeReachability)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *AmfEventReport) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetCommFailure

`func (o *AmfEventReport) GetCommFailure() CommunicationFailure`

GetCommFailure returns the CommFailure field if non-nil, zero value otherwise.

### GetCommFailureOk

`func (o *AmfEventReport) GetCommFailureOk() (*CommunicationFailure, bool)`

GetCommFailureOk returns a tuple with the CommFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommFailure

`func (o *AmfEventReport) SetCommFailure(v CommunicationFailure)`

SetCommFailure sets CommFailure field to given value.

### HasCommFailure

`func (o *AmfEventReport) HasCommFailure() bool`

HasCommFailure returns a boolean if a field has been set.

### GetLossOfConnectReason

`func (o *AmfEventReport) GetLossOfConnectReason() LossOfConnectivityReason`

GetLossOfConnectReason returns the LossOfConnectReason field if non-nil, zero value otherwise.

### GetLossOfConnectReasonOk

`func (o *AmfEventReport) GetLossOfConnectReasonOk() (*LossOfConnectivityReason, bool)`

GetLossOfConnectReasonOk returns a tuple with the LossOfConnectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossOfConnectReason

`func (o *AmfEventReport) SetLossOfConnectReason(v LossOfConnectivityReason)`

SetLossOfConnectReason sets LossOfConnectReason field to given value.

### HasLossOfConnectReason

`func (o *AmfEventReport) HasLossOfConnectReason() bool`

HasLossOfConnectReason returns a boolean if a field has been set.

### GetNumberOfUes

`func (o *AmfEventReport) GetNumberOfUes() int32`

GetNumberOfUes returns the NumberOfUes field if non-nil, zero value otherwise.

### GetNumberOfUesOk

`func (o *AmfEventReport) GetNumberOfUesOk() (*int32, bool)`

GetNumberOfUesOk returns a tuple with the NumberOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUes

`func (o *AmfEventReport) SetNumberOfUes(v int32)`

SetNumberOfUes sets NumberOfUes field to given value.

### HasNumberOfUes

`func (o *AmfEventReport) HasNumberOfUes() bool`

HasNumberOfUes returns a boolean if a field has been set.

### GetVar5gsUserStateList

`func (o *AmfEventReport) GetVar5gsUserStateList() []Model5GsUserStateInfo`

GetVar5gsUserStateList returns the Var5gsUserStateList field if non-nil, zero value otherwise.

### GetVar5gsUserStateListOk

`func (o *AmfEventReport) GetVar5gsUserStateListOk() (*[]Model5GsUserStateInfo, bool)`

GetVar5gsUserStateListOk returns a tuple with the Var5gsUserStateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gsUserStateList

`func (o *AmfEventReport) SetVar5gsUserStateList(v []Model5GsUserStateInfo)`

SetVar5gsUserStateList sets Var5gsUserStateList field to given value.

### HasVar5gsUserStateList

`func (o *AmfEventReport) HasVar5gsUserStateList() bool`

HasVar5gsUserStateList returns a boolean if a field has been set.

### GetTypeCode

`func (o *AmfEventReport) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *AmfEventReport) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *AmfEventReport) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *AmfEventReport) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *AmfEventReport) GetRegistrationNumber() int32`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *AmfEventReport) GetRegistrationNumberOk() (*int32, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *AmfEventReport) SetRegistrationNumber(v int32)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *AmfEventReport) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetMaxAvailabilityTime

`func (o *AmfEventReport) GetMaxAvailabilityTime() time.Time`

GetMaxAvailabilityTime returns the MaxAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxAvailabilityTimeOk

`func (o *AmfEventReport) GetMaxAvailabilityTimeOk() (*time.Time, bool)`

GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailabilityTime

`func (o *AmfEventReport) SetMaxAvailabilityTime(v time.Time)`

SetMaxAvailabilityTime sets MaxAvailabilityTime field to given value.

### HasMaxAvailabilityTime

`func (o *AmfEventReport) HasMaxAvailabilityTime() bool`

HasMaxAvailabilityTime returns a boolean if a field has been set.

### GetUeIdExt

`func (o *AmfEventReport) GetUeIdExt() []UEIdExt`

GetUeIdExt returns the UeIdExt field if non-nil, zero value otherwise.

### GetUeIdExtOk

`func (o *AmfEventReport) GetUeIdExtOk() (*[]UEIdExt, bool)`

GetUeIdExtOk returns a tuple with the UeIdExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIdExt

`func (o *AmfEventReport) SetUeIdExt(v []UEIdExt)`

SetUeIdExt sets UeIdExt field to given value.

### HasUeIdExt

`func (o *AmfEventReport) HasUeIdExt() bool`

HasUeIdExt returns a boolean if a field has been set.

### GetIdleStatusIndication

`func (o *AmfEventReport) GetIdleStatusIndication() IdleStatusIndication`

GetIdleStatusIndication returns the IdleStatusIndication field if non-nil, zero value otherwise.

### GetIdleStatusIndicationOk

`func (o *AmfEventReport) GetIdleStatusIndicationOk() (*IdleStatusIndication, bool)`

GetIdleStatusIndicationOk returns a tuple with the IdleStatusIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusIndication

`func (o *AmfEventReport) SetIdleStatusIndication(v IdleStatusIndication)`

SetIdleStatusIndication sets IdleStatusIndication field to given value.

### HasIdleStatusIndication

`func (o *AmfEventReport) HasIdleStatusIndication() bool`

HasIdleStatusIndication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


