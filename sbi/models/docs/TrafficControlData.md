# TrafficControlData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcId** | **string** | Univocally identifies the traffic control policy data within a PDU session. | 
**FlowStatus** | Pointer to [**FlowStatus**](FlowStatus.md) |  | [optional] 
**RedirectInfo** | Pointer to [**RedirectInformation**](RedirectInformation.md) |  | [optional] 
**AddRedirectInfo** | Pointer to [**[]RedirectInformation**](RedirectInformation.md) |  | [optional] 
**MuteNotif** | Pointer to **bool** | Indicates whether applicat&#39;on&#39;s start or stop notification is to be muted. | [optional] 
**TrafficSteeringPolIdDl** | Pointer to **string** | Reference to a pre-configured traffic steering policy for downlink traffic at the SMF. | [optional] 
**TrafficSteeringPolIdUl** | Pointer to **string** | Reference to a pre-configured traffic steering policy for uplink traffic at the SMF. | [optional] 
**RouteToLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) | A list of location which the traffic shall be routed to for the AF request | [optional] 
**TraffCorreInd** | Pointer to **bool** |  | [optional] 
**UpPathChgEvent** | Pointer to [**UpPathChgEvent**](UpPathChgEvent.md) |  | [optional] 
**SteerFun** | Pointer to [**SteeringFunctionality**](SteeringFunctionality.md) |  | [optional] 
**SteerModeDl** | Pointer to [**SteeringMode**](SteeringMode.md) |  | [optional] 
**SteerModeUl** | Pointer to [**SteeringMode**](SteeringMode.md) |  | [optional] 
**MulAccCtrl** | Pointer to [**MulticastAccessControl**](MulticastAccessControl.md) |  | [optional] 

## Methods

### NewTrafficControlData

`func NewTrafficControlData(tcId string, ) *TrafficControlData`

NewTrafficControlData instantiates a new TrafficControlData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficControlDataWithDefaults

`func NewTrafficControlDataWithDefaults() *TrafficControlData`

NewTrafficControlDataWithDefaults instantiates a new TrafficControlData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcId

`func (o *TrafficControlData) GetTcId() string`

GetTcId returns the TcId field if non-nil, zero value otherwise.

### GetTcIdOk

`func (o *TrafficControlData) GetTcIdOk() (*string, bool)`

GetTcIdOk returns a tuple with the TcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcId

`func (o *TrafficControlData) SetTcId(v string)`

SetTcId sets TcId field to given value.


### GetFlowStatus

`func (o *TrafficControlData) GetFlowStatus() FlowStatus`

GetFlowStatus returns the FlowStatus field if non-nil, zero value otherwise.

### GetFlowStatusOk

`func (o *TrafficControlData) GetFlowStatusOk() (*FlowStatus, bool)`

GetFlowStatusOk returns a tuple with the FlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowStatus

`func (o *TrafficControlData) SetFlowStatus(v FlowStatus)`

SetFlowStatus sets FlowStatus field to given value.

### HasFlowStatus

`func (o *TrafficControlData) HasFlowStatus() bool`

HasFlowStatus returns a boolean if a field has been set.

### GetRedirectInfo

`func (o *TrafficControlData) GetRedirectInfo() RedirectInformation`

GetRedirectInfo returns the RedirectInfo field if non-nil, zero value otherwise.

### GetRedirectInfoOk

`func (o *TrafficControlData) GetRedirectInfoOk() (*RedirectInformation, bool)`

GetRedirectInfoOk returns a tuple with the RedirectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectInfo

`func (o *TrafficControlData) SetRedirectInfo(v RedirectInformation)`

SetRedirectInfo sets RedirectInfo field to given value.

### HasRedirectInfo

`func (o *TrafficControlData) HasRedirectInfo() bool`

HasRedirectInfo returns a boolean if a field has been set.

### GetAddRedirectInfo

`func (o *TrafficControlData) GetAddRedirectInfo() []RedirectInformation`

GetAddRedirectInfo returns the AddRedirectInfo field if non-nil, zero value otherwise.

### GetAddRedirectInfoOk

`func (o *TrafficControlData) GetAddRedirectInfoOk() (*[]RedirectInformation, bool)`

GetAddRedirectInfoOk returns a tuple with the AddRedirectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRedirectInfo

`func (o *TrafficControlData) SetAddRedirectInfo(v []RedirectInformation)`

SetAddRedirectInfo sets AddRedirectInfo field to given value.

### HasAddRedirectInfo

`func (o *TrafficControlData) HasAddRedirectInfo() bool`

HasAddRedirectInfo returns a boolean if a field has been set.

### GetMuteNotif

`func (o *TrafficControlData) GetMuteNotif() bool`

GetMuteNotif returns the MuteNotif field if non-nil, zero value otherwise.

### GetMuteNotifOk

`func (o *TrafficControlData) GetMuteNotifOk() (*bool, bool)`

GetMuteNotifOk returns a tuple with the MuteNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteNotif

`func (o *TrafficControlData) SetMuteNotif(v bool)`

SetMuteNotif sets MuteNotif field to given value.

### HasMuteNotif

`func (o *TrafficControlData) HasMuteNotif() bool`

HasMuteNotif returns a boolean if a field has been set.

### GetTrafficSteeringPolIdDl

`func (o *TrafficControlData) GetTrafficSteeringPolIdDl() string`

GetTrafficSteeringPolIdDl returns the TrafficSteeringPolIdDl field if non-nil, zero value otherwise.

### GetTrafficSteeringPolIdDlOk

`func (o *TrafficControlData) GetTrafficSteeringPolIdDlOk() (*string, bool)`

GetTrafficSteeringPolIdDlOk returns a tuple with the TrafficSteeringPolIdDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSteeringPolIdDl

`func (o *TrafficControlData) SetTrafficSteeringPolIdDl(v string)`

SetTrafficSteeringPolIdDl sets TrafficSteeringPolIdDl field to given value.

### HasTrafficSteeringPolIdDl

`func (o *TrafficControlData) HasTrafficSteeringPolIdDl() bool`

HasTrafficSteeringPolIdDl returns a boolean if a field has been set.

### SetTrafficSteeringPolIdDlNil

`func (o *TrafficControlData) SetTrafficSteeringPolIdDlNil(b bool)`

 SetTrafficSteeringPolIdDlNil sets the value for TrafficSteeringPolIdDl to be an explicit nil

### UnsetTrafficSteeringPolIdDl
`func (o *TrafficControlData) UnsetTrafficSteeringPolIdDl()`

UnsetTrafficSteeringPolIdDl ensures that no value is present for TrafficSteeringPolIdDl, not even an explicit nil
### GetTrafficSteeringPolIdUl

`func (o *TrafficControlData) GetTrafficSteeringPolIdUl() string`

GetTrafficSteeringPolIdUl returns the TrafficSteeringPolIdUl field if non-nil, zero value otherwise.

### GetTrafficSteeringPolIdUlOk

`func (o *TrafficControlData) GetTrafficSteeringPolIdUlOk() (*string, bool)`

GetTrafficSteeringPolIdUlOk returns a tuple with the TrafficSteeringPolIdUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSteeringPolIdUl

`func (o *TrafficControlData) SetTrafficSteeringPolIdUl(v string)`

SetTrafficSteeringPolIdUl sets TrafficSteeringPolIdUl field to given value.

### HasTrafficSteeringPolIdUl

`func (o *TrafficControlData) HasTrafficSteeringPolIdUl() bool`

HasTrafficSteeringPolIdUl returns a boolean if a field has been set.

### SetTrafficSteeringPolIdUlNil

`func (o *TrafficControlData) SetTrafficSteeringPolIdUlNil(b bool)`

 SetTrafficSteeringPolIdUlNil sets the value for TrafficSteeringPolIdUl to be an explicit nil

### UnsetTrafficSteeringPolIdUl
`func (o *TrafficControlData) UnsetTrafficSteeringPolIdUl()`

UnsetTrafficSteeringPolIdUl ensures that no value is present for TrafficSteeringPolIdUl, not even an explicit nil
### GetRouteToLocs

`func (o *TrafficControlData) GetRouteToLocs() []RouteToLocation`

GetRouteToLocs returns the RouteToLocs field if non-nil, zero value otherwise.

### GetRouteToLocsOk

`func (o *TrafficControlData) GetRouteToLocsOk() (*[]RouteToLocation, bool)`

GetRouteToLocsOk returns a tuple with the RouteToLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToLocs

`func (o *TrafficControlData) SetRouteToLocs(v []RouteToLocation)`

SetRouteToLocs sets RouteToLocs field to given value.

### HasRouteToLocs

`func (o *TrafficControlData) HasRouteToLocs() bool`

HasRouteToLocs returns a boolean if a field has been set.

### GetTraffCorreInd

`func (o *TrafficControlData) GetTraffCorreInd() bool`

GetTraffCorreInd returns the TraffCorreInd field if non-nil, zero value otherwise.

### GetTraffCorreIndOk

`func (o *TrafficControlData) GetTraffCorreIndOk() (*bool, bool)`

GetTraffCorreIndOk returns a tuple with the TraffCorreInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffCorreInd

`func (o *TrafficControlData) SetTraffCorreInd(v bool)`

SetTraffCorreInd sets TraffCorreInd field to given value.

### HasTraffCorreInd

`func (o *TrafficControlData) HasTraffCorreInd() bool`

HasTraffCorreInd returns a boolean if a field has been set.

### GetUpPathChgEvent

`func (o *TrafficControlData) GetUpPathChgEvent() UpPathChgEvent`

GetUpPathChgEvent returns the UpPathChgEvent field if non-nil, zero value otherwise.

### GetUpPathChgEventOk

`func (o *TrafficControlData) GetUpPathChgEventOk() (*UpPathChgEvent, bool)`

GetUpPathChgEventOk returns a tuple with the UpPathChgEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgEvent

`func (o *TrafficControlData) SetUpPathChgEvent(v UpPathChgEvent)`

SetUpPathChgEvent sets UpPathChgEvent field to given value.

### HasUpPathChgEvent

`func (o *TrafficControlData) HasUpPathChgEvent() bool`

HasUpPathChgEvent returns a boolean if a field has been set.

### SetUpPathChgEventNil

`func (o *TrafficControlData) SetUpPathChgEventNil(b bool)`

 SetUpPathChgEventNil sets the value for UpPathChgEvent to be an explicit nil

### UnsetUpPathChgEvent
`func (o *TrafficControlData) UnsetUpPathChgEvent()`

UnsetUpPathChgEvent ensures that no value is present for UpPathChgEvent, not even an explicit nil
### GetSteerFun

`func (o *TrafficControlData) GetSteerFun() SteeringFunctionality`

GetSteerFun returns the SteerFun field if non-nil, zero value otherwise.

### GetSteerFunOk

`func (o *TrafficControlData) GetSteerFunOk() (*SteeringFunctionality, bool)`

GetSteerFunOk returns a tuple with the SteerFun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteerFun

`func (o *TrafficControlData) SetSteerFun(v SteeringFunctionality)`

SetSteerFun sets SteerFun field to given value.

### HasSteerFun

`func (o *TrafficControlData) HasSteerFun() bool`

HasSteerFun returns a boolean if a field has been set.

### GetSteerModeDl

`func (o *TrafficControlData) GetSteerModeDl() SteeringMode`

GetSteerModeDl returns the SteerModeDl field if non-nil, zero value otherwise.

### GetSteerModeDlOk

`func (o *TrafficControlData) GetSteerModeDlOk() (*SteeringMode, bool)`

GetSteerModeDlOk returns a tuple with the SteerModeDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteerModeDl

`func (o *TrafficControlData) SetSteerModeDl(v SteeringMode)`

SetSteerModeDl sets SteerModeDl field to given value.

### HasSteerModeDl

`func (o *TrafficControlData) HasSteerModeDl() bool`

HasSteerModeDl returns a boolean if a field has been set.

### GetSteerModeUl

`func (o *TrafficControlData) GetSteerModeUl() SteeringMode`

GetSteerModeUl returns the SteerModeUl field if non-nil, zero value otherwise.

### GetSteerModeUlOk

`func (o *TrafficControlData) GetSteerModeUlOk() (*SteeringMode, bool)`

GetSteerModeUlOk returns a tuple with the SteerModeUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteerModeUl

`func (o *TrafficControlData) SetSteerModeUl(v SteeringMode)`

SetSteerModeUl sets SteerModeUl field to given value.

### HasSteerModeUl

`func (o *TrafficControlData) HasSteerModeUl() bool`

HasSteerModeUl returns a boolean if a field has been set.

### GetMulAccCtrl

`func (o *TrafficControlData) GetMulAccCtrl() MulticastAccessControl`

GetMulAccCtrl returns the MulAccCtrl field if non-nil, zero value otherwise.

### GetMulAccCtrlOk

`func (o *TrafficControlData) GetMulAccCtrlOk() (*MulticastAccessControl, bool)`

GetMulAccCtrlOk returns a tuple with the MulAccCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulAccCtrl

`func (o *TrafficControlData) SetMulAccCtrl(v MulticastAccessControl)`

SetMulAccCtrl sets MulAccCtrl field to given value.

### HasMulAccCtrl

`func (o *TrafficControlData) HasMulAccCtrl() bool`

HasMulAccCtrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


