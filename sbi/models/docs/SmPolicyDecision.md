# SmPolicyDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessRules** | Pointer to [**map[string]SessionRule**](SessionRule.md) | A map of Sessionrules with the content being the SessionRule as described in subclause 5.6.2.7. | [optional] 
**PccRules** | Pointer to [**map[string]PccRule**](PccRule.md) | A map of PCC rules with the content being the PCCRule as described in subclause 5.6.2.6. | [optional] 
**PcscfRestIndication** | Pointer to **bool** | If it is included and set to true, it indicates the P-CSCF Restoration is requested. | [optional] 
**QosDecs** | Pointer to [**map[string]QosData**](QosData.md) | Map of QoS data policy decisions. | [optional] 
**ChgDecs** | Pointer to [**map[string]ChargingData**](ChargingData.md) | Map of Charging data policy decisions. | [optional] 
**ChargingInfo** | Pointer to [**ChargingInformation**](ChargingInformation.md) |  | [optional] 
**TraffContDecs** | Pointer to [**map[string]TrafficControlData**](TrafficControlData.md) | Map of Traffic Control data policy decisions. | [optional] 
**UmDecs** | Pointer to [**map[string]UsageMonitoringData**](UsageMonitoringData.md) | Map of Usage Monitoring data policy decisions. | [optional] 
**QosChars** | Pointer to [**map[string]QosCharacteristics**](QosCharacteristics.md) | Map of QoS characteristics for non standard 5QIs. This map uses the 5QI values as keys. | [optional] 
**QosMonDecs** | Pointer to [**map[string]QosMonitoringData**](QosMonitoringData.md) | Map of QoS Monitoring data policy decisions. | [optional] 
**ReflectiveQoSTimer** | Pointer to **int32** |  | [optional] 
**Conds** | Pointer to [**map[string]ConditionData**](ConditionData.md) | A map of condition data with the content being as described in subclause 5.6.2.9. | [optional] 
**RevalidationTime** | Pointer to **time.Time** |  | [optional] 
**Offline** | Pointer to **bool** | Indicates the offline charging is applicable to the PDU session when it is included and set to true. | [optional] 
**Online** | Pointer to **bool** | Indicates the online charging is applicable to the PDU session when it is included and set to true. | [optional] 
**PolicyCtrlReqTriggers** | Pointer to [**[]PolicyControlRequestTrigger**](PolicyControlRequestTrigger.md) | Defines the policy control request triggers subscribed by the PCF. | [optional] 
**LastReqRuleData** | Pointer to [**[]RequestedRuleData**](RequestedRuleData.md) | Defines the last list of rule control data requested by the PCF. | [optional] 
**LastReqUsageData** | Pointer to [**RequestedUsageData**](RequestedUsageData.md) |  | [optional] 
**PraInfos** | Pointer to [**map[string]PresenceInfoRm**](PresenceInfoRm.md) | Map of PRA information. | [optional] 
**Ipv4Index** | Pointer to **int32** |  | [optional] 
**Ipv6Index** | Pointer to **int32** |  | [optional] 
**QosFlowUsage** | Pointer to [**QosFlowUsage**](QosFlowUsage.md) |  | [optional] 
**RelCause** | Pointer to [**SmPolicyAssociationReleaseCause**](SmPolicyAssociationReleaseCause.md) |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**RedSessIndication** | Pointer to **bool** | Indicates whether the PDU session is a redundant PDU session. If absent it means the PDU session is not a redundant PDU session. | [optional] 

## Methods

### NewSmPolicyDecision

`func NewSmPolicyDecision() *SmPolicyDecision`

NewSmPolicyDecision instantiates a new SmPolicyDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDecisionWithDefaults

`func NewSmPolicyDecisionWithDefaults() *SmPolicyDecision`

NewSmPolicyDecisionWithDefaults instantiates a new SmPolicyDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessRules

`func (o *SmPolicyDecision) GetSessRules() map[string]SessionRule`

GetSessRules returns the SessRules field if non-nil, zero value otherwise.

### GetSessRulesOk

`func (o *SmPolicyDecision) GetSessRulesOk() (*map[string]SessionRule, bool)`

GetSessRulesOk returns a tuple with the SessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRules

`func (o *SmPolicyDecision) SetSessRules(v map[string]SessionRule)`

SetSessRules sets SessRules field to given value.

### HasSessRules

`func (o *SmPolicyDecision) HasSessRules() bool`

HasSessRules returns a boolean if a field has been set.

### GetPccRules

`func (o *SmPolicyDecision) GetPccRules() map[string]PccRule`

GetPccRules returns the PccRules field if non-nil, zero value otherwise.

### GetPccRulesOk

`func (o *SmPolicyDecision) GetPccRulesOk() (*map[string]PccRule, bool)`

GetPccRulesOk returns a tuple with the PccRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRules

`func (o *SmPolicyDecision) SetPccRules(v map[string]PccRule)`

SetPccRules sets PccRules field to given value.

### HasPccRules

`func (o *SmPolicyDecision) HasPccRules() bool`

HasPccRules returns a boolean if a field has been set.

### SetPccRulesNil

`func (o *SmPolicyDecision) SetPccRulesNil(b bool)`

 SetPccRulesNil sets the value for PccRules to be an explicit nil

### UnsetPccRules
`func (o *SmPolicyDecision) UnsetPccRules()`

UnsetPccRules ensures that no value is present for PccRules, not even an explicit nil
### GetPcscfRestIndication

`func (o *SmPolicyDecision) GetPcscfRestIndication() bool`

GetPcscfRestIndication returns the PcscfRestIndication field if non-nil, zero value otherwise.

### GetPcscfRestIndicationOk

`func (o *SmPolicyDecision) GetPcscfRestIndicationOk() (*bool, bool)`

GetPcscfRestIndicationOk returns a tuple with the PcscfRestIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfRestIndication

`func (o *SmPolicyDecision) SetPcscfRestIndication(v bool)`

SetPcscfRestIndication sets PcscfRestIndication field to given value.

### HasPcscfRestIndication

`func (o *SmPolicyDecision) HasPcscfRestIndication() bool`

HasPcscfRestIndication returns a boolean if a field has been set.

### GetQosDecs

`func (o *SmPolicyDecision) GetQosDecs() map[string]QosData`

GetQosDecs returns the QosDecs field if non-nil, zero value otherwise.

### GetQosDecsOk

`func (o *SmPolicyDecision) GetQosDecsOk() (*map[string]QosData, bool)`

GetQosDecsOk returns a tuple with the QosDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosDecs

`func (o *SmPolicyDecision) SetQosDecs(v map[string]QosData)`

SetQosDecs sets QosDecs field to given value.

### HasQosDecs

`func (o *SmPolicyDecision) HasQosDecs() bool`

HasQosDecs returns a boolean if a field has been set.

### GetChgDecs

`func (o *SmPolicyDecision) GetChgDecs() map[string]ChargingData`

GetChgDecs returns the ChgDecs field if non-nil, zero value otherwise.

### GetChgDecsOk

`func (o *SmPolicyDecision) GetChgDecsOk() (*map[string]ChargingData, bool)`

GetChgDecsOk returns a tuple with the ChgDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChgDecs

`func (o *SmPolicyDecision) SetChgDecs(v map[string]ChargingData)`

SetChgDecs sets ChgDecs field to given value.

### HasChgDecs

`func (o *SmPolicyDecision) HasChgDecs() bool`

HasChgDecs returns a boolean if a field has been set.

### SetChgDecsNil

`func (o *SmPolicyDecision) SetChgDecsNil(b bool)`

 SetChgDecsNil sets the value for ChgDecs to be an explicit nil

### UnsetChgDecs
`func (o *SmPolicyDecision) UnsetChgDecs()`

UnsetChgDecs ensures that no value is present for ChgDecs, not even an explicit nil
### GetChargingInfo

`func (o *SmPolicyDecision) GetChargingInfo() ChargingInformation`

GetChargingInfo returns the ChargingInfo field if non-nil, zero value otherwise.

### GetChargingInfoOk

`func (o *SmPolicyDecision) GetChargingInfoOk() (*ChargingInformation, bool)`

GetChargingInfoOk returns a tuple with the ChargingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingInfo

`func (o *SmPolicyDecision) SetChargingInfo(v ChargingInformation)`

SetChargingInfo sets ChargingInfo field to given value.

### HasChargingInfo

`func (o *SmPolicyDecision) HasChargingInfo() bool`

HasChargingInfo returns a boolean if a field has been set.

### GetTraffContDecs

`func (o *SmPolicyDecision) GetTraffContDecs() map[string]TrafficControlData`

GetTraffContDecs returns the TraffContDecs field if non-nil, zero value otherwise.

### GetTraffContDecsOk

`func (o *SmPolicyDecision) GetTraffContDecsOk() (*map[string]TrafficControlData, bool)`

GetTraffContDecsOk returns a tuple with the TraffContDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffContDecs

`func (o *SmPolicyDecision) SetTraffContDecs(v map[string]TrafficControlData)`

SetTraffContDecs sets TraffContDecs field to given value.

### HasTraffContDecs

`func (o *SmPolicyDecision) HasTraffContDecs() bool`

HasTraffContDecs returns a boolean if a field has been set.

### GetUmDecs

`func (o *SmPolicyDecision) GetUmDecs() map[string]UsageMonitoringData`

GetUmDecs returns the UmDecs field if non-nil, zero value otherwise.

### GetUmDecsOk

`func (o *SmPolicyDecision) GetUmDecsOk() (*map[string]UsageMonitoringData, bool)`

GetUmDecsOk returns a tuple with the UmDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmDecs

`func (o *SmPolicyDecision) SetUmDecs(v map[string]UsageMonitoringData)`

SetUmDecs sets UmDecs field to given value.

### HasUmDecs

`func (o *SmPolicyDecision) HasUmDecs() bool`

HasUmDecs returns a boolean if a field has been set.

### SetUmDecsNil

`func (o *SmPolicyDecision) SetUmDecsNil(b bool)`

 SetUmDecsNil sets the value for UmDecs to be an explicit nil

### UnsetUmDecs
`func (o *SmPolicyDecision) UnsetUmDecs()`

UnsetUmDecs ensures that no value is present for UmDecs, not even an explicit nil
### GetQosChars

`func (o *SmPolicyDecision) GetQosChars() map[string]QosCharacteristics`

GetQosChars returns the QosChars field if non-nil, zero value otherwise.

### GetQosCharsOk

`func (o *SmPolicyDecision) GetQosCharsOk() (*map[string]QosCharacteristics, bool)`

GetQosCharsOk returns a tuple with the QosChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosChars

`func (o *SmPolicyDecision) SetQosChars(v map[string]QosCharacteristics)`

SetQosChars sets QosChars field to given value.

### HasQosChars

`func (o *SmPolicyDecision) HasQosChars() bool`

HasQosChars returns a boolean if a field has been set.

### GetQosMonDecs

`func (o *SmPolicyDecision) GetQosMonDecs() map[string]QosMonitoringData`

GetQosMonDecs returns the QosMonDecs field if non-nil, zero value otherwise.

### GetQosMonDecsOk

`func (o *SmPolicyDecision) GetQosMonDecsOk() (*map[string]QosMonitoringData, bool)`

GetQosMonDecsOk returns a tuple with the QosMonDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonDecs

`func (o *SmPolicyDecision) SetQosMonDecs(v map[string]QosMonitoringData)`

SetQosMonDecs sets QosMonDecs field to given value.

### HasQosMonDecs

`func (o *SmPolicyDecision) HasQosMonDecs() bool`

HasQosMonDecs returns a boolean if a field has been set.

### SetQosMonDecsNil

`func (o *SmPolicyDecision) SetQosMonDecsNil(b bool)`

 SetQosMonDecsNil sets the value for QosMonDecs to be an explicit nil

### UnsetQosMonDecs
`func (o *SmPolicyDecision) UnsetQosMonDecs()`

UnsetQosMonDecs ensures that no value is present for QosMonDecs, not even an explicit nil
### GetReflectiveQoSTimer

`func (o *SmPolicyDecision) GetReflectiveQoSTimer() int32`

GetReflectiveQoSTimer returns the ReflectiveQoSTimer field if non-nil, zero value otherwise.

### GetReflectiveQoSTimerOk

`func (o *SmPolicyDecision) GetReflectiveQoSTimerOk() (*int32, bool)`

GetReflectiveQoSTimerOk returns a tuple with the ReflectiveQoSTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectiveQoSTimer

`func (o *SmPolicyDecision) SetReflectiveQoSTimer(v int32)`

SetReflectiveQoSTimer sets ReflectiveQoSTimer field to given value.

### HasReflectiveQoSTimer

`func (o *SmPolicyDecision) HasReflectiveQoSTimer() bool`

HasReflectiveQoSTimer returns a boolean if a field has been set.

### GetConds

`func (o *SmPolicyDecision) GetConds() map[string]ConditionData`

GetConds returns the Conds field if non-nil, zero value otherwise.

### GetCondsOk

`func (o *SmPolicyDecision) GetCondsOk() (*map[string]ConditionData, bool)`

GetCondsOk returns a tuple with the Conds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConds

`func (o *SmPolicyDecision) SetConds(v map[string]ConditionData)`

SetConds sets Conds field to given value.

### HasConds

`func (o *SmPolicyDecision) HasConds() bool`

HasConds returns a boolean if a field has been set.

### SetCondsNil

`func (o *SmPolicyDecision) SetCondsNil(b bool)`

 SetCondsNil sets the value for Conds to be an explicit nil

### UnsetConds
`func (o *SmPolicyDecision) UnsetConds()`

UnsetConds ensures that no value is present for Conds, not even an explicit nil
### GetRevalidationTime

`func (o *SmPolicyDecision) GetRevalidationTime() time.Time`

GetRevalidationTime returns the RevalidationTime field if non-nil, zero value otherwise.

### GetRevalidationTimeOk

`func (o *SmPolicyDecision) GetRevalidationTimeOk() (*time.Time, bool)`

GetRevalidationTimeOk returns a tuple with the RevalidationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevalidationTime

`func (o *SmPolicyDecision) SetRevalidationTime(v time.Time)`

SetRevalidationTime sets RevalidationTime field to given value.

### HasRevalidationTime

`func (o *SmPolicyDecision) HasRevalidationTime() bool`

HasRevalidationTime returns a boolean if a field has been set.

### GetOffline

`func (o *SmPolicyDecision) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *SmPolicyDecision) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *SmPolicyDecision) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *SmPolicyDecision) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOnline

`func (o *SmPolicyDecision) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SmPolicyDecision) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SmPolicyDecision) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SmPolicyDecision) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPolicyCtrlReqTriggers

`func (o *SmPolicyDecision) GetPolicyCtrlReqTriggers() []PolicyControlRequestTrigger`

GetPolicyCtrlReqTriggers returns the PolicyCtrlReqTriggers field if non-nil, zero value otherwise.

### GetPolicyCtrlReqTriggersOk

`func (o *SmPolicyDecision) GetPolicyCtrlReqTriggersOk() (*[]PolicyControlRequestTrigger, bool)`

GetPolicyCtrlReqTriggersOk returns a tuple with the PolicyCtrlReqTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCtrlReqTriggers

`func (o *SmPolicyDecision) SetPolicyCtrlReqTriggers(v []PolicyControlRequestTrigger)`

SetPolicyCtrlReqTriggers sets PolicyCtrlReqTriggers field to given value.

### HasPolicyCtrlReqTriggers

`func (o *SmPolicyDecision) HasPolicyCtrlReqTriggers() bool`

HasPolicyCtrlReqTriggers returns a boolean if a field has been set.

### SetPolicyCtrlReqTriggersNil

`func (o *SmPolicyDecision) SetPolicyCtrlReqTriggersNil(b bool)`

 SetPolicyCtrlReqTriggersNil sets the value for PolicyCtrlReqTriggers to be an explicit nil

### UnsetPolicyCtrlReqTriggers
`func (o *SmPolicyDecision) UnsetPolicyCtrlReqTriggers()`

UnsetPolicyCtrlReqTriggers ensures that no value is present for PolicyCtrlReqTriggers, not even an explicit nil
### GetLastReqRuleData

`func (o *SmPolicyDecision) GetLastReqRuleData() []RequestedRuleData`

GetLastReqRuleData returns the LastReqRuleData field if non-nil, zero value otherwise.

### GetLastReqRuleDataOk

`func (o *SmPolicyDecision) GetLastReqRuleDataOk() (*[]RequestedRuleData, bool)`

GetLastReqRuleDataOk returns a tuple with the LastReqRuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReqRuleData

`func (o *SmPolicyDecision) SetLastReqRuleData(v []RequestedRuleData)`

SetLastReqRuleData sets LastReqRuleData field to given value.

### HasLastReqRuleData

`func (o *SmPolicyDecision) HasLastReqRuleData() bool`

HasLastReqRuleData returns a boolean if a field has been set.

### GetLastReqUsageData

`func (o *SmPolicyDecision) GetLastReqUsageData() RequestedUsageData`

GetLastReqUsageData returns the LastReqUsageData field if non-nil, zero value otherwise.

### GetLastReqUsageDataOk

`func (o *SmPolicyDecision) GetLastReqUsageDataOk() (*RequestedUsageData, bool)`

GetLastReqUsageDataOk returns a tuple with the LastReqUsageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReqUsageData

`func (o *SmPolicyDecision) SetLastReqUsageData(v RequestedUsageData)`

SetLastReqUsageData sets LastReqUsageData field to given value.

### HasLastReqUsageData

`func (o *SmPolicyDecision) HasLastReqUsageData() bool`

HasLastReqUsageData returns a boolean if a field has been set.

### GetPraInfos

`func (o *SmPolicyDecision) GetPraInfos() map[string]PresenceInfoRm`

GetPraInfos returns the PraInfos field if non-nil, zero value otherwise.

### GetPraInfosOk

`func (o *SmPolicyDecision) GetPraInfosOk() (*map[string]PresenceInfoRm, bool)`

GetPraInfosOk returns a tuple with the PraInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInfos

`func (o *SmPolicyDecision) SetPraInfos(v map[string]PresenceInfoRm)`

SetPraInfos sets PraInfos field to given value.

### HasPraInfos

`func (o *SmPolicyDecision) HasPraInfos() bool`

HasPraInfos returns a boolean if a field has been set.

### SetPraInfosNil

`func (o *SmPolicyDecision) SetPraInfosNil(b bool)`

 SetPraInfosNil sets the value for PraInfos to be an explicit nil

### UnsetPraInfos
`func (o *SmPolicyDecision) UnsetPraInfos()`

UnsetPraInfos ensures that no value is present for PraInfos, not even an explicit nil
### GetIpv4Index

`func (o *SmPolicyDecision) GetIpv4Index() int32`

GetIpv4Index returns the Ipv4Index field if non-nil, zero value otherwise.

### GetIpv4IndexOk

`func (o *SmPolicyDecision) GetIpv4IndexOk() (*int32, bool)`

GetIpv4IndexOk returns a tuple with the Ipv4Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Index

`func (o *SmPolicyDecision) SetIpv4Index(v int32)`

SetIpv4Index sets Ipv4Index field to given value.

### HasIpv4Index

`func (o *SmPolicyDecision) HasIpv4Index() bool`

HasIpv4Index returns a boolean if a field has been set.

### GetIpv6Index

`func (o *SmPolicyDecision) GetIpv6Index() int32`

GetIpv6Index returns the Ipv6Index field if non-nil, zero value otherwise.

### GetIpv6IndexOk

`func (o *SmPolicyDecision) GetIpv6IndexOk() (*int32, bool)`

GetIpv6IndexOk returns a tuple with the Ipv6Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Index

`func (o *SmPolicyDecision) SetIpv6Index(v int32)`

SetIpv6Index sets Ipv6Index field to given value.

### HasIpv6Index

`func (o *SmPolicyDecision) HasIpv6Index() bool`

HasIpv6Index returns a boolean if a field has been set.

### GetQosFlowUsage

`func (o *SmPolicyDecision) GetQosFlowUsage() QosFlowUsage`

GetQosFlowUsage returns the QosFlowUsage field if non-nil, zero value otherwise.

### GetQosFlowUsageOk

`func (o *SmPolicyDecision) GetQosFlowUsageOk() (*QosFlowUsage, bool)`

GetQosFlowUsageOk returns a tuple with the QosFlowUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowUsage

`func (o *SmPolicyDecision) SetQosFlowUsage(v QosFlowUsage)`

SetQosFlowUsage sets QosFlowUsage field to given value.

### HasQosFlowUsage

`func (o *SmPolicyDecision) HasQosFlowUsage() bool`

HasQosFlowUsage returns a boolean if a field has been set.

### GetRelCause

`func (o *SmPolicyDecision) GetRelCause() SmPolicyAssociationReleaseCause`

GetRelCause returns the RelCause field if non-nil, zero value otherwise.

### GetRelCauseOk

`func (o *SmPolicyDecision) GetRelCauseOk() (*SmPolicyAssociationReleaseCause, bool)`

GetRelCauseOk returns a tuple with the RelCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelCause

`func (o *SmPolicyDecision) SetRelCause(v SmPolicyAssociationReleaseCause)`

SetRelCause sets RelCause field to given value.

### HasRelCause

`func (o *SmPolicyDecision) HasRelCause() bool`

HasRelCause returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SmPolicyDecision) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SmPolicyDecision) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SmPolicyDecision) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SmPolicyDecision) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetTsnBridgeManCont

`func (o *SmPolicyDecision) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *SmPolicyDecision) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *SmPolicyDecision) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *SmPolicyDecision) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *SmPolicyDecision) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *SmPolicyDecision) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *SmPolicyDecision) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *SmPolicyDecision) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *SmPolicyDecision) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *SmPolicyDecision) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *SmPolicyDecision) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *SmPolicyDecision) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.

### GetRedSessIndication

`func (o *SmPolicyDecision) GetRedSessIndication() bool`

GetRedSessIndication returns the RedSessIndication field if non-nil, zero value otherwise.

### GetRedSessIndicationOk

`func (o *SmPolicyDecision) GetRedSessIndicationOk() (*bool, bool)`

GetRedSessIndicationOk returns a tuple with the RedSessIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedSessIndication

`func (o *SmPolicyDecision) SetRedSessIndication(v bool)`

SetRedSessIndication sets RedSessIndication field to given value.

### HasRedSessIndication

`func (o *SmPolicyDecision) HasRedSessIndication() bool`

HasRedSessIndication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


