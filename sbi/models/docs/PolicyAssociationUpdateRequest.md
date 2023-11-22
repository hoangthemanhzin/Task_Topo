# PolicyAssociationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationUri** | Pointer to **string** |  | [optional] 
**AltNotifIpv4Addrs** | Pointer to **[]string** | Alternate or backup IPv4 Address(es) where to send Notifications. | [optional] 
**AltNotifIpv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) | Alternate or backup IPv6 Address(es) where to send Notifications. | [optional] 
**AltNotifFqdns** | Pointer to **[]string** | Alternate or backup FQDN(s) where to send Notifications. | [optional] 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the NF service consumer observes. | [optional] 
**PraStatuses** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Map of PRA status information. | [optional] 
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UePolDelResult** | Pointer to **string** |  | [optional] 
**UePolTransFailNotif** | Pointer to [**UePolicyTransferFailureNotification**](UePolicyTransferFailureNotification.md) |  | [optional] 
**UePolReq** | Pointer to **string** |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServingNfId** | Pointer to **string** |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ConnectState** | Pointer to [**CmState**](CmState.md) |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyAssociationUpdateRequest

`func NewPolicyAssociationUpdateRequest() *PolicyAssociationUpdateRequest`

NewPolicyAssociationUpdateRequest instantiates a new PolicyAssociationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationUpdateRequestWithDefaults

`func NewPolicyAssociationUpdateRequestWithDefaults() *PolicyAssociationUpdateRequest`

NewPolicyAssociationUpdateRequestWithDefaults instantiates a new PolicyAssociationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationUri

`func (o *PolicyAssociationUpdateRequest) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *PolicyAssociationUpdateRequest) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *PolicyAssociationUpdateRequest) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.

### HasNotificationUri

`func (o *PolicyAssociationUpdateRequest) HasNotificationUri() bool`

HasNotificationUri returns a boolean if a field has been set.

### GetAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4Addrs() []string`

GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv4AddrsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4AddrsOk() (*[]string, bool)`

GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv4Addrs(v []string)`

SetAltNotifIpv4Addrs sets AltNotifIpv4Addrs field to given value.

### HasAltNotifIpv4Addrs

`func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv4Addrs() bool`

HasAltNotifIpv4Addrs returns a boolean if a field has been set.

### GetAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6Addrs() []Ipv6Addr`

GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv6AddrsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr)`

SetAltNotifIpv6Addrs sets AltNotifIpv6Addrs field to given value.

### HasAltNotifIpv6Addrs

`func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv6Addrs() bool`

HasAltNotifIpv6Addrs returns a boolean if a field has been set.

### GetAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdns() []string`

GetAltNotifFqdns returns the AltNotifFqdns field if non-nil, zero value otherwise.

### GetAltNotifFqdnsOk

`func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdnsOk() (*[]string, bool)`

GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) SetAltNotifFqdns(v []string)`

SetAltNotifFqdns sets AltNotifFqdns field to given value.

### HasAltNotifFqdns

`func (o *PolicyAssociationUpdateRequest) HasAltNotifFqdns() bool`

HasAltNotifFqdns returns a boolean if a field has been set.

### GetTriggers

`func (o *PolicyAssociationUpdateRequest) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyAssociationUpdateRequest) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyAssociationUpdateRequest) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyAssociationUpdateRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetPraStatuses

`func (o *PolicyAssociationUpdateRequest) GetPraStatuses() map[string]PresenceInfo`

GetPraStatuses returns the PraStatuses field if non-nil, zero value otherwise.

### GetPraStatusesOk

`func (o *PolicyAssociationUpdateRequest) GetPraStatusesOk() (*map[string]PresenceInfo, bool)`

GetPraStatusesOk returns a tuple with the PraStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraStatuses

`func (o *PolicyAssociationUpdateRequest) SetPraStatuses(v map[string]PresenceInfo)`

SetPraStatuses sets PraStatuses field to given value.

### HasPraStatuses

`func (o *PolicyAssociationUpdateRequest) HasPraStatuses() bool`

HasPraStatuses returns a boolean if a field has been set.

### GetUserLoc

`func (o *PolicyAssociationUpdateRequest) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *PolicyAssociationUpdateRequest) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *PolicyAssociationUpdateRequest) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *PolicyAssociationUpdateRequest) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetUePolDelResult

`func (o *PolicyAssociationUpdateRequest) GetUePolDelResult() string`

GetUePolDelResult returns the UePolDelResult field if non-nil, zero value otherwise.

### GetUePolDelResultOk

`func (o *PolicyAssociationUpdateRequest) GetUePolDelResultOk() (*string, bool)`

GetUePolDelResultOk returns a tuple with the UePolDelResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolDelResult

`func (o *PolicyAssociationUpdateRequest) SetUePolDelResult(v string)`

SetUePolDelResult sets UePolDelResult field to given value.

### HasUePolDelResult

`func (o *PolicyAssociationUpdateRequest) HasUePolDelResult() bool`

HasUePolDelResult returns a boolean if a field has been set.

### GetUePolTransFailNotif

`func (o *PolicyAssociationUpdateRequest) GetUePolTransFailNotif() UePolicyTransferFailureNotification`

GetUePolTransFailNotif returns the UePolTransFailNotif field if non-nil, zero value otherwise.

### GetUePolTransFailNotifOk

`func (o *PolicyAssociationUpdateRequest) GetUePolTransFailNotifOk() (*UePolicyTransferFailureNotification, bool)`

GetUePolTransFailNotifOk returns a tuple with the UePolTransFailNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolTransFailNotif

`func (o *PolicyAssociationUpdateRequest) SetUePolTransFailNotif(v UePolicyTransferFailureNotification)`

SetUePolTransFailNotif sets UePolTransFailNotif field to given value.

### HasUePolTransFailNotif

`func (o *PolicyAssociationUpdateRequest) HasUePolTransFailNotif() bool`

HasUePolTransFailNotif returns a boolean if a field has been set.

### GetUePolReq

`func (o *PolicyAssociationUpdateRequest) GetUePolReq() string`

GetUePolReq returns the UePolReq field if non-nil, zero value otherwise.

### GetUePolReqOk

`func (o *PolicyAssociationUpdateRequest) GetUePolReqOk() (*string, bool)`

GetUePolReqOk returns a tuple with the UePolReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolReq

`func (o *PolicyAssociationUpdateRequest) SetUePolReq(v string)`

SetUePolReq sets UePolReq field to given value.

### HasUePolReq

`func (o *PolicyAssociationUpdateRequest) HasUePolReq() bool`

HasUePolReq returns a boolean if a field has been set.

### GetGuami

`func (o *PolicyAssociationUpdateRequest) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *PolicyAssociationUpdateRequest) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *PolicyAssociationUpdateRequest) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *PolicyAssociationUpdateRequest) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServingNfId

`func (o *PolicyAssociationUpdateRequest) GetServingNfId() string`

GetServingNfId returns the ServingNfId field if non-nil, zero value otherwise.

### GetServingNfIdOk

`func (o *PolicyAssociationUpdateRequest) GetServingNfIdOk() (*string, bool)`

GetServingNfIdOk returns a tuple with the ServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfId

`func (o *PolicyAssociationUpdateRequest) SetServingNfId(v string)`

SetServingNfId sets ServingNfId field to given value.

### HasServingNfId

`func (o *PolicyAssociationUpdateRequest) HasServingNfId() bool`

HasServingNfId returns a boolean if a field has been set.

### GetPlmnId

`func (o *PolicyAssociationUpdateRequest) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PolicyAssociationUpdateRequest) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PolicyAssociationUpdateRequest) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PolicyAssociationUpdateRequest) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetConnectState

`func (o *PolicyAssociationUpdateRequest) GetConnectState() CmState`

GetConnectState returns the ConnectState field if non-nil, zero value otherwise.

### GetConnectStateOk

`func (o *PolicyAssociationUpdateRequest) GetConnectStateOk() (*CmState, bool)`

GetConnectStateOk returns a tuple with the ConnectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectState

`func (o *PolicyAssociationUpdateRequest) SetConnectState(v CmState)`

SetConnectState sets ConnectState field to given value.

### HasConnectState

`func (o *PolicyAssociationUpdateRequest) HasConnectState() bool`

HasConnectState returns a boolean if a field has been set.

### GetGroupIds

`func (o *PolicyAssociationUpdateRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PolicyAssociationUpdateRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PolicyAssociationUpdateRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PolicyAssociationUpdateRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


