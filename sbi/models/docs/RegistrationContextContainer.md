# RegistrationContextContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeContext** | [**UeContext**](UeContext.md) |  | 
**LocalTimeZone** | Pointer to **string** |  | [optional] 
**AnType** | [**AccessType**](AccessType.md) |  | 
**AnN2ApId** | **int32** |  | 
**RanNodeId** | [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | 
**InitialAmfName** | **string** |  | 
**UserLocation** | [**UserLocation**](UserLocation.md) |  | 
**RrcEstCause** | Pointer to **string** |  | [optional] 
**UeContextRequest** | Pointer to **bool** |  | [optional] [default to false]
**InitialAmfN2ApId** | Pointer to **int32** |  | [optional] 
**AnN2IPv4Addr** | Pointer to **string** |  | [optional] 
**AnN2IPv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AllowedNssai** | Pointer to [**AllowedNssai**](AllowedNssai.md) |  | [optional] 
**ConfiguredNssai** | Pointer to [**[]ConfiguredSnssai**](ConfiguredSnssai.md) |  | [optional] 
**RejectedNssaiInPlmn** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RejectedNssaiInTa** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**SelectedPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**IabNodeInd** | Pointer to **bool** |  | [optional] [default to false]
**CeModeBInd** | Pointer to [**CeModeBInd**](CeModeBInd.md) |  | [optional] 
**LteMInd** | Pointer to [**LteMInd**](LteMInd.md) |  | [optional] 
**AuthenticatedInd** | Pointer to **bool** |  | [optional] [default to false]
**NpnAccessInfo** | Pointer to [**NpnAccessInfo**](NpnAccessInfo.md) |  | [optional] 

## Methods

### NewRegistrationContextContainer

`func NewRegistrationContextContainer(ueContext UeContext, anType AccessType, anN2ApId int32, ranNodeId GlobalRanNodeId, initialAmfName string, userLocation UserLocation, ) *RegistrationContextContainer`

NewRegistrationContextContainer instantiates a new RegistrationContextContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationContextContainerWithDefaults

`func NewRegistrationContextContainerWithDefaults() *RegistrationContextContainer`

NewRegistrationContextContainerWithDefaults instantiates a new RegistrationContextContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeContext

`func (o *RegistrationContextContainer) GetUeContext() UeContext`

GetUeContext returns the UeContext field if non-nil, zero value otherwise.

### GetUeContextOk

`func (o *RegistrationContextContainer) GetUeContextOk() (*UeContext, bool)`

GetUeContextOk returns a tuple with the UeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContext

`func (o *RegistrationContextContainer) SetUeContext(v UeContext)`

SetUeContext sets UeContext field to given value.


### GetLocalTimeZone

`func (o *RegistrationContextContainer) GetLocalTimeZone() string`

GetLocalTimeZone returns the LocalTimeZone field if non-nil, zero value otherwise.

### GetLocalTimeZoneOk

`func (o *RegistrationContextContainer) GetLocalTimeZoneOk() (*string, bool)`

GetLocalTimeZoneOk returns a tuple with the LocalTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTimeZone

`func (o *RegistrationContextContainer) SetLocalTimeZone(v string)`

SetLocalTimeZone sets LocalTimeZone field to given value.

### HasLocalTimeZone

`func (o *RegistrationContextContainer) HasLocalTimeZone() bool`

HasLocalTimeZone returns a boolean if a field has been set.

### GetAnType

`func (o *RegistrationContextContainer) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *RegistrationContextContainer) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *RegistrationContextContainer) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetAnN2ApId

`func (o *RegistrationContextContainer) GetAnN2ApId() int32`

GetAnN2ApId returns the AnN2ApId field if non-nil, zero value otherwise.

### GetAnN2ApIdOk

`func (o *RegistrationContextContainer) GetAnN2ApIdOk() (*int32, bool)`

GetAnN2ApIdOk returns a tuple with the AnN2ApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2ApId

`func (o *RegistrationContextContainer) SetAnN2ApId(v int32)`

SetAnN2ApId sets AnN2ApId field to given value.


### GetRanNodeId

`func (o *RegistrationContextContainer) GetRanNodeId() GlobalRanNodeId`

GetRanNodeId returns the RanNodeId field if non-nil, zero value otherwise.

### GetRanNodeIdOk

`func (o *RegistrationContextContainer) GetRanNodeIdOk() (*GlobalRanNodeId, bool)`

GetRanNodeIdOk returns a tuple with the RanNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNodeId

`func (o *RegistrationContextContainer) SetRanNodeId(v GlobalRanNodeId)`

SetRanNodeId sets RanNodeId field to given value.


### GetInitialAmfName

`func (o *RegistrationContextContainer) GetInitialAmfName() string`

GetInitialAmfName returns the InitialAmfName field if non-nil, zero value otherwise.

### GetInitialAmfNameOk

`func (o *RegistrationContextContainer) GetInitialAmfNameOk() (*string, bool)`

GetInitialAmfNameOk returns a tuple with the InitialAmfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmfName

`func (o *RegistrationContextContainer) SetInitialAmfName(v string)`

SetInitialAmfName sets InitialAmfName field to given value.


### GetUserLocation

`func (o *RegistrationContextContainer) GetUserLocation() UserLocation`

GetUserLocation returns the UserLocation field if non-nil, zero value otherwise.

### GetUserLocationOk

`func (o *RegistrationContextContainer) GetUserLocationOk() (*UserLocation, bool)`

GetUserLocationOk returns a tuple with the UserLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocation

`func (o *RegistrationContextContainer) SetUserLocation(v UserLocation)`

SetUserLocation sets UserLocation field to given value.


### GetRrcEstCause

`func (o *RegistrationContextContainer) GetRrcEstCause() string`

GetRrcEstCause returns the RrcEstCause field if non-nil, zero value otherwise.

### GetRrcEstCauseOk

`func (o *RegistrationContextContainer) GetRrcEstCauseOk() (*string, bool)`

GetRrcEstCauseOk returns a tuple with the RrcEstCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrcEstCause

`func (o *RegistrationContextContainer) SetRrcEstCause(v string)`

SetRrcEstCause sets RrcEstCause field to given value.

### HasRrcEstCause

`func (o *RegistrationContextContainer) HasRrcEstCause() bool`

HasRrcEstCause returns a boolean if a field has been set.

### GetUeContextRequest

`func (o *RegistrationContextContainer) GetUeContextRequest() bool`

GetUeContextRequest returns the UeContextRequest field if non-nil, zero value otherwise.

### GetUeContextRequestOk

`func (o *RegistrationContextContainer) GetUeContextRequestOk() (*bool, bool)`

GetUeContextRequestOk returns a tuple with the UeContextRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeContextRequest

`func (o *RegistrationContextContainer) SetUeContextRequest(v bool)`

SetUeContextRequest sets UeContextRequest field to given value.

### HasUeContextRequest

`func (o *RegistrationContextContainer) HasUeContextRequest() bool`

HasUeContextRequest returns a boolean if a field has been set.

### GetInitialAmfN2ApId

`func (o *RegistrationContextContainer) GetInitialAmfN2ApId() int32`

GetInitialAmfN2ApId returns the InitialAmfN2ApId field if non-nil, zero value otherwise.

### GetInitialAmfN2ApIdOk

`func (o *RegistrationContextContainer) GetInitialAmfN2ApIdOk() (*int32, bool)`

GetInitialAmfN2ApIdOk returns a tuple with the InitialAmfN2ApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmfN2ApId

`func (o *RegistrationContextContainer) SetInitialAmfN2ApId(v int32)`

SetInitialAmfN2ApId sets InitialAmfN2ApId field to given value.

### HasInitialAmfN2ApId

`func (o *RegistrationContextContainer) HasInitialAmfN2ApId() bool`

HasInitialAmfN2ApId returns a boolean if a field has been set.

### GetAnN2IPv4Addr

`func (o *RegistrationContextContainer) GetAnN2IPv4Addr() string`

GetAnN2IPv4Addr returns the AnN2IPv4Addr field if non-nil, zero value otherwise.

### GetAnN2IPv4AddrOk

`func (o *RegistrationContextContainer) GetAnN2IPv4AddrOk() (*string, bool)`

GetAnN2IPv4AddrOk returns a tuple with the AnN2IPv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2IPv4Addr

`func (o *RegistrationContextContainer) SetAnN2IPv4Addr(v string)`

SetAnN2IPv4Addr sets AnN2IPv4Addr field to given value.

### HasAnN2IPv4Addr

`func (o *RegistrationContextContainer) HasAnN2IPv4Addr() bool`

HasAnN2IPv4Addr returns a boolean if a field has been set.

### GetAnN2IPv6Addr

`func (o *RegistrationContextContainer) GetAnN2IPv6Addr() Ipv6Addr`

GetAnN2IPv6Addr returns the AnN2IPv6Addr field if non-nil, zero value otherwise.

### GetAnN2IPv6AddrOk

`func (o *RegistrationContextContainer) GetAnN2IPv6AddrOk() (*Ipv6Addr, bool)`

GetAnN2IPv6AddrOk returns a tuple with the AnN2IPv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnN2IPv6Addr

`func (o *RegistrationContextContainer) SetAnN2IPv6Addr(v Ipv6Addr)`

SetAnN2IPv6Addr sets AnN2IPv6Addr field to given value.

### HasAnN2IPv6Addr

`func (o *RegistrationContextContainer) HasAnN2IPv6Addr() bool`

HasAnN2IPv6Addr returns a boolean if a field has been set.

### GetAllowedNssai

`func (o *RegistrationContextContainer) GetAllowedNssai() AllowedNssai`

GetAllowedNssai returns the AllowedNssai field if non-nil, zero value otherwise.

### GetAllowedNssaiOk

`func (o *RegistrationContextContainer) GetAllowedNssaiOk() (*AllowedNssai, bool)`

GetAllowedNssaiOk returns a tuple with the AllowedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssai

`func (o *RegistrationContextContainer) SetAllowedNssai(v AllowedNssai)`

SetAllowedNssai sets AllowedNssai field to given value.

### HasAllowedNssai

`func (o *RegistrationContextContainer) HasAllowedNssai() bool`

HasAllowedNssai returns a boolean if a field has been set.

### GetConfiguredNssai

`func (o *RegistrationContextContainer) GetConfiguredNssai() []ConfiguredSnssai`

GetConfiguredNssai returns the ConfiguredNssai field if non-nil, zero value otherwise.

### GetConfiguredNssaiOk

`func (o *RegistrationContextContainer) GetConfiguredNssaiOk() (*[]ConfiguredSnssai, bool)`

GetConfiguredNssaiOk returns a tuple with the ConfiguredNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredNssai

`func (o *RegistrationContextContainer) SetConfiguredNssai(v []ConfiguredSnssai)`

SetConfiguredNssai sets ConfiguredNssai field to given value.

### HasConfiguredNssai

`func (o *RegistrationContextContainer) HasConfiguredNssai() bool`

HasConfiguredNssai returns a boolean if a field has been set.

### GetRejectedNssaiInPlmn

`func (o *RegistrationContextContainer) GetRejectedNssaiInPlmn() []Snssai`

GetRejectedNssaiInPlmn returns the RejectedNssaiInPlmn field if non-nil, zero value otherwise.

### GetRejectedNssaiInPlmnOk

`func (o *RegistrationContextContainer) GetRejectedNssaiInPlmnOk() (*[]Snssai, bool)`

GetRejectedNssaiInPlmnOk returns a tuple with the RejectedNssaiInPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNssaiInPlmn

`func (o *RegistrationContextContainer) SetRejectedNssaiInPlmn(v []Snssai)`

SetRejectedNssaiInPlmn sets RejectedNssaiInPlmn field to given value.

### HasRejectedNssaiInPlmn

`func (o *RegistrationContextContainer) HasRejectedNssaiInPlmn() bool`

HasRejectedNssaiInPlmn returns a boolean if a field has been set.

### GetRejectedNssaiInTa

`func (o *RegistrationContextContainer) GetRejectedNssaiInTa() []Snssai`

GetRejectedNssaiInTa returns the RejectedNssaiInTa field if non-nil, zero value otherwise.

### GetRejectedNssaiInTaOk

`func (o *RegistrationContextContainer) GetRejectedNssaiInTaOk() (*[]Snssai, bool)`

GetRejectedNssaiInTaOk returns a tuple with the RejectedNssaiInTa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNssaiInTa

`func (o *RegistrationContextContainer) SetRejectedNssaiInTa(v []Snssai)`

SetRejectedNssaiInTa sets RejectedNssaiInTa field to given value.

### HasRejectedNssaiInTa

`func (o *RegistrationContextContainer) HasRejectedNssaiInTa() bool`

HasRejectedNssaiInTa returns a boolean if a field has been set.

### GetSelectedPlmnId

`func (o *RegistrationContextContainer) GetSelectedPlmnId() PlmnId`

GetSelectedPlmnId returns the SelectedPlmnId field if non-nil, zero value otherwise.

### GetSelectedPlmnIdOk

`func (o *RegistrationContextContainer) GetSelectedPlmnIdOk() (*PlmnId, bool)`

GetSelectedPlmnIdOk returns a tuple with the SelectedPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPlmnId

`func (o *RegistrationContextContainer) SetSelectedPlmnId(v PlmnId)`

SetSelectedPlmnId sets SelectedPlmnId field to given value.

### HasSelectedPlmnId

`func (o *RegistrationContextContainer) HasSelectedPlmnId() bool`

HasSelectedPlmnId returns a boolean if a field has been set.

### GetIabNodeInd

`func (o *RegistrationContextContainer) GetIabNodeInd() bool`

GetIabNodeInd returns the IabNodeInd field if non-nil, zero value otherwise.

### GetIabNodeIndOk

`func (o *RegistrationContextContainer) GetIabNodeIndOk() (*bool, bool)`

GetIabNodeIndOk returns a tuple with the IabNodeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIabNodeInd

`func (o *RegistrationContextContainer) SetIabNodeInd(v bool)`

SetIabNodeInd sets IabNodeInd field to given value.

### HasIabNodeInd

`func (o *RegistrationContextContainer) HasIabNodeInd() bool`

HasIabNodeInd returns a boolean if a field has been set.

### GetCeModeBInd

`func (o *RegistrationContextContainer) GetCeModeBInd() CeModeBInd`

GetCeModeBInd returns the CeModeBInd field if non-nil, zero value otherwise.

### GetCeModeBIndOk

`func (o *RegistrationContextContainer) GetCeModeBIndOk() (*CeModeBInd, bool)`

GetCeModeBIndOk returns a tuple with the CeModeBInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeModeBInd

`func (o *RegistrationContextContainer) SetCeModeBInd(v CeModeBInd)`

SetCeModeBInd sets CeModeBInd field to given value.

### HasCeModeBInd

`func (o *RegistrationContextContainer) HasCeModeBInd() bool`

HasCeModeBInd returns a boolean if a field has been set.

### GetLteMInd

`func (o *RegistrationContextContainer) GetLteMInd() LteMInd`

GetLteMInd returns the LteMInd field if non-nil, zero value otherwise.

### GetLteMIndOk

`func (o *RegistrationContextContainer) GetLteMIndOk() (*LteMInd, bool)`

GetLteMIndOk returns a tuple with the LteMInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteMInd

`func (o *RegistrationContextContainer) SetLteMInd(v LteMInd)`

SetLteMInd sets LteMInd field to given value.

### HasLteMInd

`func (o *RegistrationContextContainer) HasLteMInd() bool`

HasLteMInd returns a boolean if a field has been set.

### GetAuthenticatedInd

`func (o *RegistrationContextContainer) GetAuthenticatedInd() bool`

GetAuthenticatedInd returns the AuthenticatedInd field if non-nil, zero value otherwise.

### GetAuthenticatedIndOk

`func (o *RegistrationContextContainer) GetAuthenticatedIndOk() (*bool, bool)`

GetAuthenticatedIndOk returns a tuple with the AuthenticatedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedInd

`func (o *RegistrationContextContainer) SetAuthenticatedInd(v bool)`

SetAuthenticatedInd sets AuthenticatedInd field to given value.

### HasAuthenticatedInd

`func (o *RegistrationContextContainer) HasAuthenticatedInd() bool`

HasAuthenticatedInd returns a boolean if a field has been set.

### GetNpnAccessInfo

`func (o *RegistrationContextContainer) GetNpnAccessInfo() NpnAccessInfo`

GetNpnAccessInfo returns the NpnAccessInfo field if non-nil, zero value otherwise.

### GetNpnAccessInfoOk

`func (o *RegistrationContextContainer) GetNpnAccessInfoOk() (*NpnAccessInfo, bool)`

GetNpnAccessInfoOk returns a tuple with the NpnAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpnAccessInfo

`func (o *RegistrationContextContainer) SetNpnAccessInfo(v NpnAccessInfo)`

SetNpnAccessInfo sets NpnAccessInfo field to given value.

### HasNpnAccessInfo

`func (o *RegistrationContextContainer) HasNpnAccessInfo() bool`

HasNpnAccessInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


