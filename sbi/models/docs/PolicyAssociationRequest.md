# PolicyAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | **string** |  | 
**AltNotifIpv4Addrs** | Pointer to **[]string** | Alternate or backup IPv4 Address(es) where to send Notifications. | [optional] 
**AltNotifIpv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) | Alternate or backup IPv6 Address(es) where to send Notifications. | [optional] 
**AltNotifFqdns** | Pointer to **[]string** | Alternate or backup FQDN(s) where to send Notifications. | [optional] 
**Supi** | **string** |  | 
**Gpsi** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**Pei** | Pointer to **string** |  | [optional] 
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**ServingPlmn** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**HPcfId** | Pointer to **string** |  | [optional] 
**UePolReq** | Pointer to **string** |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**ServingNfId** | Pointer to **string** |  | [optional] 
**Pc5Capab** | Pointer to [**Pc5Capability**](Pc5Capability.md) |  | [optional] 
**SuppFeat** | **string** |  | 

## Methods

### NewPolicyAssociationRequest

`func NewPolicyAssociationRequest(notificationUri string, supi string, suppFeat string, ) *PolicyAssociationRequest`

NewPolicyAssociationRequest instantiates a new PolicyAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationRequestWithDefaults

`func NewPolicyAssociationRequestWithDefaults() *PolicyAssociationRequest`

NewPolicyAssociationRequestWithDefaults instantiates a new PolicyAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *PolicyAssociationRequest) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *PolicyAssociationRequest) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *PolicyAssociationRequest) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) GetAltNotifIpv4Addrs() []string`

GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv4AddrsOk

`func (o *PolicyAssociationRequest) GetAltNotifIpv4AddrsOk() (*[]string, bool)`

GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) SetAltNotifIpv4Addrs(v []string)`

SetAltNotifIpv4Addrs sets AltNotifIpv4Addrs field to given value.

### HasAltNotifIpv4Addrs

`func (o *PolicyAssociationRequest) HasAltNotifIpv4Addrs() bool`

HasAltNotifIpv4Addrs returns a boolean if a field has been set.

### GetAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) GetAltNotifIpv6Addrs() []Ipv6Addr`

GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv6AddrsOk

`func (o *PolicyAssociationRequest) GetAltNotifIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr)`

SetAltNotifIpv6Addrs sets AltNotifIpv6Addrs field to given value.

### HasAltNotifIpv6Addrs

`func (o *PolicyAssociationRequest) HasAltNotifIpv6Addrs() bool`

HasAltNotifIpv6Addrs returns a boolean if a field has been set.

### GetAltNotifFqdns

`func (o *PolicyAssociationRequest) GetAltNotifFqdns() []string`

GetAltNotifFqdns returns the AltNotifFqdns field if non-nil, zero value otherwise.

### GetAltNotifFqdnsOk

`func (o *PolicyAssociationRequest) GetAltNotifFqdnsOk() (*[]string, bool)`

GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifFqdns

`func (o *PolicyAssociationRequest) SetAltNotifFqdns(v []string)`

SetAltNotifFqdns sets AltNotifFqdns field to given value.

### HasAltNotifFqdns

`func (o *PolicyAssociationRequest) HasAltNotifFqdns() bool`

HasAltNotifFqdns returns a boolean if a field has been set.

### GetSupi

`func (o *PolicyAssociationRequest) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PolicyAssociationRequest) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PolicyAssociationRequest) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsi

`func (o *PolicyAssociationRequest) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PolicyAssociationRequest) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PolicyAssociationRequest) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PolicyAssociationRequest) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetAccessType

`func (o *PolicyAssociationRequest) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PolicyAssociationRequest) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PolicyAssociationRequest) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PolicyAssociationRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPei

`func (o *PolicyAssociationRequest) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *PolicyAssociationRequest) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *PolicyAssociationRequest) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *PolicyAssociationRequest) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetUserLoc

`func (o *PolicyAssociationRequest) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *PolicyAssociationRequest) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *PolicyAssociationRequest) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *PolicyAssociationRequest) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetTimeZone

`func (o *PolicyAssociationRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PolicyAssociationRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PolicyAssociationRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PolicyAssociationRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetServingPlmn

`func (o *PolicyAssociationRequest) GetServingPlmn() PlmnIdNid`

GetServingPlmn returns the ServingPlmn field if non-nil, zero value otherwise.

### GetServingPlmnOk

`func (o *PolicyAssociationRequest) GetServingPlmnOk() (*PlmnIdNid, bool)`

GetServingPlmnOk returns a tuple with the ServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPlmn

`func (o *PolicyAssociationRequest) SetServingPlmn(v PlmnIdNid)`

SetServingPlmn sets ServingPlmn field to given value.

### HasServingPlmn

`func (o *PolicyAssociationRequest) HasServingPlmn() bool`

HasServingPlmn returns a boolean if a field has been set.

### GetRatType

`func (o *PolicyAssociationRequest) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *PolicyAssociationRequest) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *PolicyAssociationRequest) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *PolicyAssociationRequest) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetGroupIds

`func (o *PolicyAssociationRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PolicyAssociationRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PolicyAssociationRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PolicyAssociationRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetHPcfId

`func (o *PolicyAssociationRequest) GetHPcfId() string`

GetHPcfId returns the HPcfId field if non-nil, zero value otherwise.

### GetHPcfIdOk

`func (o *PolicyAssociationRequest) GetHPcfIdOk() (*string, bool)`

GetHPcfIdOk returns a tuple with the HPcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHPcfId

`func (o *PolicyAssociationRequest) SetHPcfId(v string)`

SetHPcfId sets HPcfId field to given value.

### HasHPcfId

`func (o *PolicyAssociationRequest) HasHPcfId() bool`

HasHPcfId returns a boolean if a field has been set.

### GetUePolReq

`func (o *PolicyAssociationRequest) GetUePolReq() string`

GetUePolReq returns the UePolReq field if non-nil, zero value otherwise.

### GetUePolReqOk

`func (o *PolicyAssociationRequest) GetUePolReqOk() (*string, bool)`

GetUePolReqOk returns a tuple with the UePolReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolReq

`func (o *PolicyAssociationRequest) SetUePolReq(v string)`

SetUePolReq sets UePolReq field to given value.

### HasUePolReq

`func (o *PolicyAssociationRequest) HasUePolReq() bool`

HasUePolReq returns a boolean if a field has been set.

### GetGuami

`func (o *PolicyAssociationRequest) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *PolicyAssociationRequest) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *PolicyAssociationRequest) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *PolicyAssociationRequest) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServiceName

`func (o *PolicyAssociationRequest) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PolicyAssociationRequest) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PolicyAssociationRequest) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PolicyAssociationRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServingNfId

`func (o *PolicyAssociationRequest) GetServingNfId() string`

GetServingNfId returns the ServingNfId field if non-nil, zero value otherwise.

### GetServingNfIdOk

`func (o *PolicyAssociationRequest) GetServingNfIdOk() (*string, bool)`

GetServingNfIdOk returns a tuple with the ServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfId

`func (o *PolicyAssociationRequest) SetServingNfId(v string)`

SetServingNfId sets ServingNfId field to given value.

### HasServingNfId

`func (o *PolicyAssociationRequest) HasServingNfId() bool`

HasServingNfId returns a boolean if a field has been set.

### GetPc5Capab

`func (o *PolicyAssociationRequest) GetPc5Capab() Pc5Capability`

GetPc5Capab returns the Pc5Capab field if non-nil, zero value otherwise.

### GetPc5CapabOk

`func (o *PolicyAssociationRequest) GetPc5CapabOk() (*Pc5Capability, bool)`

GetPc5CapabOk returns a tuple with the Pc5Capab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPc5Capab

`func (o *PolicyAssociationRequest) SetPc5Capab(v Pc5Capability)`

SetPc5Capab sets Pc5Capab field to given value.

### HasPc5Capab

`func (o *PolicyAssociationRequest) HasPc5Capab() bool`

HasPc5Capab returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PolicyAssociationRequest) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PolicyAssociationRequest) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PolicyAssociationRequest) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


