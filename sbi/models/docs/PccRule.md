# PccRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfos** | Pointer to [**[]FlowInformation**](FlowInformation.md) | An array of IP flow packet filter information. | [optional] 
**AppId** | Pointer to **string** | A reference to the application detection filter configured at the UPF. | [optional] 
**AppDescriptor** | Pointer to **string** |  | [optional] 
**ContVer** | Pointer to **int32** | Represents the content version of some content. | [optional] 
**PccRuleId** | **string** | Univocally identifies the PCC rule within a PDU session. | 
**Precedence** | Pointer to **int32** |  | [optional] 
**AfSigProtocol** | Pointer to [**AfSigProtocol**](AfSigProtocol.md) |  | [optional] 
**AppReloc** | Pointer to **bool** | Indication of application relocation possibility. | [optional] 
**RefQosData** | Pointer to **[]string** | A reference to the QosData policy decision type. It is the qosId described in subclause 5.6.2.8. | [optional] 
**RefAltQosParams** | Pointer to **[]string** | A Reference to the QosData policy decision type for the Alternative QoS parameter sets of the service data flow. | [optional] 
**RefTcData** | Pointer to **[]string** | A reference to the TrafficControlData policy decision type. It is the tcId described in subclause 5.6.2.10. | [optional] 
**RefChgData** | Pointer to **[]string** | A reference to the ChargingData policy decision type. It is the chgId described in subclause 5.6.2.11. | [optional] 
**RefChgN3gData** | Pointer to **[]string** | A reference to the ChargingData policy decision type only applicable to Non-3GPP access if \&quot;ATSSS\&quot; feature is supported. It is the chgId described in subclause 5.6.2.11. | [optional] 
**RefUmData** | Pointer to **[]string** | A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12. | [optional] 
**RefUmN3gData** | Pointer to **[]string** | A reference to UsageMonitoringData policy decision type only applicable to Non-3GPP access if \&quot;ATSSS\&quot; feature is supported. It is the umId described in subclause 5.6.2.12. | [optional] 
**RefCondData** | Pointer to **string** | A reference to the condition data. It is the condId described in subclause 5.6.2.9. | [optional] 
**RefQosMon** | Pointer to **[]string** | A reference to the QosMonitoringData policy decision type. It is the qmId described in subclause 5.6.2.40. | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 
**TscaiInputDl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**TscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**DdNotifCtrl** | Pointer to [**DownlinkDataNotificationControl**](DownlinkDataNotificationControl.md) |  | [optional] 
**DdNotifCtrl2** | Pointer to [**DownlinkDataNotificationControlRm**](DownlinkDataNotificationControlRm.md) |  | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 

## Methods

### NewPccRule

`func NewPccRule(pccRuleId string, ) *PccRule`

NewPccRule instantiates a new PccRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccRuleWithDefaults

`func NewPccRuleWithDefaults() *PccRule`

NewPccRuleWithDefaults instantiates a new PccRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfos

`func (o *PccRule) GetFlowInfos() []FlowInformation`

GetFlowInfos returns the FlowInfos field if non-nil, zero value otherwise.

### GetFlowInfosOk

`func (o *PccRule) GetFlowInfosOk() (*[]FlowInformation, bool)`

GetFlowInfosOk returns a tuple with the FlowInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfos

`func (o *PccRule) SetFlowInfos(v []FlowInformation)`

SetFlowInfos sets FlowInfos field to given value.

### HasFlowInfos

`func (o *PccRule) HasFlowInfos() bool`

HasFlowInfos returns a boolean if a field has been set.

### GetAppId

`func (o *PccRule) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PccRule) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PccRule) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *PccRule) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppDescriptor

`func (o *PccRule) GetAppDescriptor() string`

GetAppDescriptor returns the AppDescriptor field if non-nil, zero value otherwise.

### GetAppDescriptorOk

`func (o *PccRule) GetAppDescriptorOk() (*string, bool)`

GetAppDescriptorOk returns a tuple with the AppDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptor

`func (o *PccRule) SetAppDescriptor(v string)`

SetAppDescriptor sets AppDescriptor field to given value.

### HasAppDescriptor

`func (o *PccRule) HasAppDescriptor() bool`

HasAppDescriptor returns a boolean if a field has been set.

### GetContVer

`func (o *PccRule) GetContVer() int32`

GetContVer returns the ContVer field if non-nil, zero value otherwise.

### GetContVerOk

`func (o *PccRule) GetContVerOk() (*int32, bool)`

GetContVerOk returns a tuple with the ContVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContVer

`func (o *PccRule) SetContVer(v int32)`

SetContVer sets ContVer field to given value.

### HasContVer

`func (o *PccRule) HasContVer() bool`

HasContVer returns a boolean if a field has been set.

### GetPccRuleId

`func (o *PccRule) GetPccRuleId() string`

GetPccRuleId returns the PccRuleId field if non-nil, zero value otherwise.

### GetPccRuleIdOk

`func (o *PccRule) GetPccRuleIdOk() (*string, bool)`

GetPccRuleIdOk returns a tuple with the PccRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRuleId

`func (o *PccRule) SetPccRuleId(v string)`

SetPccRuleId sets PccRuleId field to given value.


### GetPrecedence

`func (o *PccRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *PccRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *PccRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *PccRule) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetAfSigProtocol

`func (o *PccRule) GetAfSigProtocol() AfSigProtocol`

GetAfSigProtocol returns the AfSigProtocol field if non-nil, zero value otherwise.

### GetAfSigProtocolOk

`func (o *PccRule) GetAfSigProtocolOk() (*AfSigProtocol, bool)`

GetAfSigProtocolOk returns a tuple with the AfSigProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfSigProtocol

`func (o *PccRule) SetAfSigProtocol(v AfSigProtocol)`

SetAfSigProtocol sets AfSigProtocol field to given value.

### HasAfSigProtocol

`func (o *PccRule) HasAfSigProtocol() bool`

HasAfSigProtocol returns a boolean if a field has been set.

### GetAppReloc

`func (o *PccRule) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *PccRule) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *PccRule) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *PccRule) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRefQosData

`func (o *PccRule) GetRefQosData() []string`

GetRefQosData returns the RefQosData field if non-nil, zero value otherwise.

### GetRefQosDataOk

`func (o *PccRule) GetRefQosDataOk() (*[]string, bool)`

GetRefQosDataOk returns a tuple with the RefQosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefQosData

`func (o *PccRule) SetRefQosData(v []string)`

SetRefQosData sets RefQosData field to given value.

### HasRefQosData

`func (o *PccRule) HasRefQosData() bool`

HasRefQosData returns a boolean if a field has been set.

### GetRefAltQosParams

`func (o *PccRule) GetRefAltQosParams() []string`

GetRefAltQosParams returns the RefAltQosParams field if non-nil, zero value otherwise.

### GetRefAltQosParamsOk

`func (o *PccRule) GetRefAltQosParamsOk() (*[]string, bool)`

GetRefAltQosParamsOk returns a tuple with the RefAltQosParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefAltQosParams

`func (o *PccRule) SetRefAltQosParams(v []string)`

SetRefAltQosParams sets RefAltQosParams field to given value.

### HasRefAltQosParams

`func (o *PccRule) HasRefAltQosParams() bool`

HasRefAltQosParams returns a boolean if a field has been set.

### GetRefTcData

`func (o *PccRule) GetRefTcData() []string`

GetRefTcData returns the RefTcData field if non-nil, zero value otherwise.

### GetRefTcDataOk

`func (o *PccRule) GetRefTcDataOk() (*[]string, bool)`

GetRefTcDataOk returns a tuple with the RefTcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTcData

`func (o *PccRule) SetRefTcData(v []string)`

SetRefTcData sets RefTcData field to given value.

### HasRefTcData

`func (o *PccRule) HasRefTcData() bool`

HasRefTcData returns a boolean if a field has been set.

### GetRefChgData

`func (o *PccRule) GetRefChgData() []string`

GetRefChgData returns the RefChgData field if non-nil, zero value otherwise.

### GetRefChgDataOk

`func (o *PccRule) GetRefChgDataOk() (*[]string, bool)`

GetRefChgDataOk returns a tuple with the RefChgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefChgData

`func (o *PccRule) SetRefChgData(v []string)`

SetRefChgData sets RefChgData field to given value.

### HasRefChgData

`func (o *PccRule) HasRefChgData() bool`

HasRefChgData returns a boolean if a field has been set.

### SetRefChgDataNil

`func (o *PccRule) SetRefChgDataNil(b bool)`

 SetRefChgDataNil sets the value for RefChgData to be an explicit nil

### UnsetRefChgData
`func (o *PccRule) UnsetRefChgData()`

UnsetRefChgData ensures that no value is present for RefChgData, not even an explicit nil
### GetRefChgN3gData

`func (o *PccRule) GetRefChgN3gData() []string`

GetRefChgN3gData returns the RefChgN3gData field if non-nil, zero value otherwise.

### GetRefChgN3gDataOk

`func (o *PccRule) GetRefChgN3gDataOk() (*[]string, bool)`

GetRefChgN3gDataOk returns a tuple with the RefChgN3gData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefChgN3gData

`func (o *PccRule) SetRefChgN3gData(v []string)`

SetRefChgN3gData sets RefChgN3gData field to given value.

### HasRefChgN3gData

`func (o *PccRule) HasRefChgN3gData() bool`

HasRefChgN3gData returns a boolean if a field has been set.

### SetRefChgN3gDataNil

`func (o *PccRule) SetRefChgN3gDataNil(b bool)`

 SetRefChgN3gDataNil sets the value for RefChgN3gData to be an explicit nil

### UnsetRefChgN3gData
`func (o *PccRule) UnsetRefChgN3gData()`

UnsetRefChgN3gData ensures that no value is present for RefChgN3gData, not even an explicit nil
### GetRefUmData

`func (o *PccRule) GetRefUmData() []string`

GetRefUmData returns the RefUmData field if non-nil, zero value otherwise.

### GetRefUmDataOk

`func (o *PccRule) GetRefUmDataOk() (*[]string, bool)`

GetRefUmDataOk returns a tuple with the RefUmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmData

`func (o *PccRule) SetRefUmData(v []string)`

SetRefUmData sets RefUmData field to given value.

### HasRefUmData

`func (o *PccRule) HasRefUmData() bool`

HasRefUmData returns a boolean if a field has been set.

### SetRefUmDataNil

`func (o *PccRule) SetRefUmDataNil(b bool)`

 SetRefUmDataNil sets the value for RefUmData to be an explicit nil

### UnsetRefUmData
`func (o *PccRule) UnsetRefUmData()`

UnsetRefUmData ensures that no value is present for RefUmData, not even an explicit nil
### GetRefUmN3gData

`func (o *PccRule) GetRefUmN3gData() []string`

GetRefUmN3gData returns the RefUmN3gData field if non-nil, zero value otherwise.

### GetRefUmN3gDataOk

`func (o *PccRule) GetRefUmN3gDataOk() (*[]string, bool)`

GetRefUmN3gDataOk returns a tuple with the RefUmN3gData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmN3gData

`func (o *PccRule) SetRefUmN3gData(v []string)`

SetRefUmN3gData sets RefUmN3gData field to given value.

### HasRefUmN3gData

`func (o *PccRule) HasRefUmN3gData() bool`

HasRefUmN3gData returns a boolean if a field has been set.

### SetRefUmN3gDataNil

`func (o *PccRule) SetRefUmN3gDataNil(b bool)`

 SetRefUmN3gDataNil sets the value for RefUmN3gData to be an explicit nil

### UnsetRefUmN3gData
`func (o *PccRule) UnsetRefUmN3gData()`

UnsetRefUmN3gData ensures that no value is present for RefUmN3gData, not even an explicit nil
### GetRefCondData

`func (o *PccRule) GetRefCondData() string`

GetRefCondData returns the RefCondData field if non-nil, zero value otherwise.

### GetRefCondDataOk

`func (o *PccRule) GetRefCondDataOk() (*string, bool)`

GetRefCondDataOk returns a tuple with the RefCondData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCondData

`func (o *PccRule) SetRefCondData(v string)`

SetRefCondData sets RefCondData field to given value.

### HasRefCondData

`func (o *PccRule) HasRefCondData() bool`

HasRefCondData returns a boolean if a field has been set.

### SetRefCondDataNil

`func (o *PccRule) SetRefCondDataNil(b bool)`

 SetRefCondDataNil sets the value for RefCondData to be an explicit nil

### UnsetRefCondData
`func (o *PccRule) UnsetRefCondData()`

UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil
### GetRefQosMon

`func (o *PccRule) GetRefQosMon() []string`

GetRefQosMon returns the RefQosMon field if non-nil, zero value otherwise.

### GetRefQosMonOk

`func (o *PccRule) GetRefQosMonOk() (*[]string, bool)`

GetRefQosMonOk returns a tuple with the RefQosMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefQosMon

`func (o *PccRule) SetRefQosMon(v []string)`

SetRefQosMon sets RefQosMon field to given value.

### HasRefQosMon

`func (o *PccRule) HasRefQosMon() bool`

HasRefQosMon returns a boolean if a field has been set.

### SetRefQosMonNil

`func (o *PccRule) SetRefQosMonNil(b bool)`

 SetRefQosMonNil sets the value for RefQosMon to be an explicit nil

### UnsetRefQosMon
`func (o *PccRule) UnsetRefQosMon()`

UnsetRefQosMon ensures that no value is present for RefQosMon, not even an explicit nil
### GetAddrPreserInd

`func (o *PccRule) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *PccRule) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *PccRule) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *PccRule) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### SetAddrPreserIndNil

`func (o *PccRule) SetAddrPreserIndNil(b bool)`

 SetAddrPreserIndNil sets the value for AddrPreserInd to be an explicit nil

### UnsetAddrPreserInd
`func (o *PccRule) UnsetAddrPreserInd()`

UnsetAddrPreserInd ensures that no value is present for AddrPreserInd, not even an explicit nil
### GetTscaiInputDl

`func (o *PccRule) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *PccRule) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *PccRule) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *PccRule) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *PccRule) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *PccRule) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *PccRule) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *PccRule) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *PccRule) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *PccRule) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *PccRule) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *PccRule) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil
### GetDdNotifCtrl

`func (o *PccRule) GetDdNotifCtrl() DownlinkDataNotificationControl`

GetDdNotifCtrl returns the DdNotifCtrl field if non-nil, zero value otherwise.

### GetDdNotifCtrlOk

`func (o *PccRule) GetDdNotifCtrlOk() (*DownlinkDataNotificationControl, bool)`

GetDdNotifCtrlOk returns a tuple with the DdNotifCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdNotifCtrl

`func (o *PccRule) SetDdNotifCtrl(v DownlinkDataNotificationControl)`

SetDdNotifCtrl sets DdNotifCtrl field to given value.

### HasDdNotifCtrl

`func (o *PccRule) HasDdNotifCtrl() bool`

HasDdNotifCtrl returns a boolean if a field has been set.

### GetDdNotifCtrl2

`func (o *PccRule) GetDdNotifCtrl2() DownlinkDataNotificationControlRm`

GetDdNotifCtrl2 returns the DdNotifCtrl2 field if non-nil, zero value otherwise.

### GetDdNotifCtrl2Ok

`func (o *PccRule) GetDdNotifCtrl2Ok() (*DownlinkDataNotificationControlRm, bool)`

GetDdNotifCtrl2Ok returns a tuple with the DdNotifCtrl2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdNotifCtrl2

`func (o *PccRule) SetDdNotifCtrl2(v DownlinkDataNotificationControlRm)`

SetDdNotifCtrl2 sets DdNotifCtrl2 field to given value.

### HasDdNotifCtrl2

`func (o *PccRule) HasDdNotifCtrl2() bool`

HasDdNotifCtrl2 returns a boolean if a field has been set.

### SetDdNotifCtrl2Nil

`func (o *PccRule) SetDdNotifCtrl2Nil(b bool)`

 SetDdNotifCtrl2Nil sets the value for DdNotifCtrl2 to be an explicit nil

### UnsetDdNotifCtrl2
`func (o *PccRule) UnsetDdNotifCtrl2()`

UnsetDdNotifCtrl2 ensures that no value is present for DdNotifCtrl2, not even an explicit nil
### GetDisUeNotif

`func (o *PccRule) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *PccRule) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *PccRule) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *PccRule) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### SetDisUeNotifNil

`func (o *PccRule) SetDisUeNotifNil(b bool)`

 SetDisUeNotifNil sets the value for DisUeNotif to be an explicit nil

### UnsetDisUeNotif
`func (o *PccRule) UnsetDisUeNotif()`

UnsetDisUeNotif ensures that no value is present for DisUeNotif, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


