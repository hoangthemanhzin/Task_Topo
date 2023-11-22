# QosData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosId** | **string** | Univocally identifies the QoS control policy data within a PDU session. | 
**Var5qi** | Pointer to **int32** |  | [optional] 
**MaxbrUl** | Pointer to **string** |  | [optional] 
**MaxbrDl** | Pointer to **string** |  | [optional] 
**GbrUl** | Pointer to **string** |  | [optional] 
**GbrDl** | Pointer to **string** |  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**Qnc** | Pointer to **bool** | Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow. | [optional] 
**PriorityLevel** | Pointer to **int32** |  | [optional] 
**AverWindow** | Pointer to **int32** |  | [optional] [default to 2000]
**MaxDataBurstVol** | Pointer to **int32** |  | [optional] 
**ReflectiveQos** | Pointer to **bool** | Indicates whether the QoS information is reflective for the corresponding service data flow. | [optional] 
**SharingKeyDl** | Pointer to **string** | Indicates, by containing the same value, what PCC rules may share resource in downlink direction. | [optional] 
**SharingKeyUl** | Pointer to **string** | Indicates, by containing the same value, what PCC rules may share resource in uplink direction. | [optional] 
**MaxPacketLossRateDl** | Pointer to **int32** |  | [optional] 
**MaxPacketLossRateUl** | Pointer to **int32** |  | [optional] 
**DefQosFlowIndication** | Pointer to **bool** | Indicates that the dynamic PCC rule shall always have its binding with the QoS Flow associated with the default QoS rule | [optional] 
**ExtMaxDataBurstVol** | Pointer to **int32** |  | [optional] 
**PacketDelayBudget** | Pointer to **int32** |  | [optional] 
**PacketErrorRate** | Pointer to **string** |  | [optional] 

## Methods

### NewQosData

`func NewQosData(qosId string, ) *QosData`

NewQosData instantiates a new QosData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosDataWithDefaults

`func NewQosDataWithDefaults() *QosData`

NewQosDataWithDefaults instantiates a new QosData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosId

`func (o *QosData) GetQosId() string`

GetQosId returns the QosId field if non-nil, zero value otherwise.

### GetQosIdOk

`func (o *QosData) GetQosIdOk() (*string, bool)`

GetQosIdOk returns a tuple with the QosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosId

`func (o *QosData) SetQosId(v string)`

SetQosId sets QosId field to given value.


### GetVar5qi

`func (o *QosData) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *QosData) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *QosData) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *QosData) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetMaxbrUl

`func (o *QosData) GetMaxbrUl() string`

GetMaxbrUl returns the MaxbrUl field if non-nil, zero value otherwise.

### GetMaxbrUlOk

`func (o *QosData) GetMaxbrUlOk() (*string, bool)`

GetMaxbrUlOk returns a tuple with the MaxbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrUl

`func (o *QosData) SetMaxbrUl(v string)`

SetMaxbrUl sets MaxbrUl field to given value.

### HasMaxbrUl

`func (o *QosData) HasMaxbrUl() bool`

HasMaxbrUl returns a boolean if a field has been set.

### SetMaxbrUlNil

`func (o *QosData) SetMaxbrUlNil(b bool)`

 SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil

### UnsetMaxbrUl
`func (o *QosData) UnsetMaxbrUl()`

UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
### GetMaxbrDl

`func (o *QosData) GetMaxbrDl() string`

GetMaxbrDl returns the MaxbrDl field if non-nil, zero value otherwise.

### GetMaxbrDlOk

`func (o *QosData) GetMaxbrDlOk() (*string, bool)`

GetMaxbrDlOk returns a tuple with the MaxbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrDl

`func (o *QosData) SetMaxbrDl(v string)`

SetMaxbrDl sets MaxbrDl field to given value.

### HasMaxbrDl

`func (o *QosData) HasMaxbrDl() bool`

HasMaxbrDl returns a boolean if a field has been set.

### SetMaxbrDlNil

`func (o *QosData) SetMaxbrDlNil(b bool)`

 SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil

### UnsetMaxbrDl
`func (o *QosData) UnsetMaxbrDl()`

UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
### GetGbrUl

`func (o *QosData) GetGbrUl() string`

GetGbrUl returns the GbrUl field if non-nil, zero value otherwise.

### GetGbrUlOk

`func (o *QosData) GetGbrUlOk() (*string, bool)`

GetGbrUlOk returns a tuple with the GbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrUl

`func (o *QosData) SetGbrUl(v string)`

SetGbrUl sets GbrUl field to given value.

### HasGbrUl

`func (o *QosData) HasGbrUl() bool`

HasGbrUl returns a boolean if a field has been set.

### SetGbrUlNil

`func (o *QosData) SetGbrUlNil(b bool)`

 SetGbrUlNil sets the value for GbrUl to be an explicit nil

### UnsetGbrUl
`func (o *QosData) UnsetGbrUl()`

UnsetGbrUl ensures that no value is present for GbrUl, not even an explicit nil
### GetGbrDl

`func (o *QosData) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *QosData) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *QosData) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *QosData) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### SetGbrDlNil

`func (o *QosData) SetGbrDlNil(b bool)`

 SetGbrDlNil sets the value for GbrDl to be an explicit nil

### UnsetGbrDl
`func (o *QosData) UnsetGbrDl()`

UnsetGbrDl ensures that no value is present for GbrDl, not even an explicit nil
### GetArp

`func (o *QosData) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *QosData) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *QosData) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *QosData) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetQnc

`func (o *QosData) GetQnc() bool`

GetQnc returns the Qnc field if non-nil, zero value otherwise.

### GetQncOk

`func (o *QosData) GetQncOk() (*bool, bool)`

GetQncOk returns a tuple with the Qnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQnc

`func (o *QosData) SetQnc(v bool)`

SetQnc sets Qnc field to given value.

### HasQnc

`func (o *QosData) HasQnc() bool`

HasQnc returns a boolean if a field has been set.

### GetPriorityLevel

`func (o *QosData) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *QosData) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *QosData) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *QosData) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### SetPriorityLevelNil

`func (o *QosData) SetPriorityLevelNil(b bool)`

 SetPriorityLevelNil sets the value for PriorityLevel to be an explicit nil

### UnsetPriorityLevel
`func (o *QosData) UnsetPriorityLevel()`

UnsetPriorityLevel ensures that no value is present for PriorityLevel, not even an explicit nil
### GetAverWindow

`func (o *QosData) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *QosData) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *QosData) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *QosData) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### SetAverWindowNil

`func (o *QosData) SetAverWindowNil(b bool)`

 SetAverWindowNil sets the value for AverWindow to be an explicit nil

### UnsetAverWindow
`func (o *QosData) UnsetAverWindow()`

UnsetAverWindow ensures that no value is present for AverWindow, not even an explicit nil
### GetMaxDataBurstVol

`func (o *QosData) GetMaxDataBurstVol() int32`

GetMaxDataBurstVol returns the MaxDataBurstVol field if non-nil, zero value otherwise.

### GetMaxDataBurstVolOk

`func (o *QosData) GetMaxDataBurstVolOk() (*int32, bool)`

GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataBurstVol

`func (o *QosData) SetMaxDataBurstVol(v int32)`

SetMaxDataBurstVol sets MaxDataBurstVol field to given value.

### HasMaxDataBurstVol

`func (o *QosData) HasMaxDataBurstVol() bool`

HasMaxDataBurstVol returns a boolean if a field has been set.

### SetMaxDataBurstVolNil

`func (o *QosData) SetMaxDataBurstVolNil(b bool)`

 SetMaxDataBurstVolNil sets the value for MaxDataBurstVol to be an explicit nil

### UnsetMaxDataBurstVol
`func (o *QosData) UnsetMaxDataBurstVol()`

UnsetMaxDataBurstVol ensures that no value is present for MaxDataBurstVol, not even an explicit nil
### GetReflectiveQos

`func (o *QosData) GetReflectiveQos() bool`

GetReflectiveQos returns the ReflectiveQos field if non-nil, zero value otherwise.

### GetReflectiveQosOk

`func (o *QosData) GetReflectiveQosOk() (*bool, bool)`

GetReflectiveQosOk returns a tuple with the ReflectiveQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectiveQos

`func (o *QosData) SetReflectiveQos(v bool)`

SetReflectiveQos sets ReflectiveQos field to given value.

### HasReflectiveQos

`func (o *QosData) HasReflectiveQos() bool`

HasReflectiveQos returns a boolean if a field has been set.

### GetSharingKeyDl

`func (o *QosData) GetSharingKeyDl() string`

GetSharingKeyDl returns the SharingKeyDl field if non-nil, zero value otherwise.

### GetSharingKeyDlOk

`func (o *QosData) GetSharingKeyDlOk() (*string, bool)`

GetSharingKeyDlOk returns a tuple with the SharingKeyDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyDl

`func (o *QosData) SetSharingKeyDl(v string)`

SetSharingKeyDl sets SharingKeyDl field to given value.

### HasSharingKeyDl

`func (o *QosData) HasSharingKeyDl() bool`

HasSharingKeyDl returns a boolean if a field has been set.

### GetSharingKeyUl

`func (o *QosData) GetSharingKeyUl() string`

GetSharingKeyUl returns the SharingKeyUl field if non-nil, zero value otherwise.

### GetSharingKeyUlOk

`func (o *QosData) GetSharingKeyUlOk() (*string, bool)`

GetSharingKeyUlOk returns a tuple with the SharingKeyUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyUl

`func (o *QosData) SetSharingKeyUl(v string)`

SetSharingKeyUl sets SharingKeyUl field to given value.

### HasSharingKeyUl

`func (o *QosData) HasSharingKeyUl() bool`

HasSharingKeyUl returns a boolean if a field has been set.

### GetMaxPacketLossRateDl

`func (o *QosData) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *QosData) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *QosData) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *QosData) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.

### SetMaxPacketLossRateDlNil

`func (o *QosData) SetMaxPacketLossRateDlNil(b bool)`

 SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil

### UnsetMaxPacketLossRateDl
`func (o *QosData) UnsetMaxPacketLossRateDl()`

UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
### GetMaxPacketLossRateUl

`func (o *QosData) GetMaxPacketLossRateUl() int32`

GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateUlOk

`func (o *QosData) GetMaxPacketLossRateUlOk() (*int32, bool)`

GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateUl

`func (o *QosData) SetMaxPacketLossRateUl(v int32)`

SetMaxPacketLossRateUl sets MaxPacketLossRateUl field to given value.

### HasMaxPacketLossRateUl

`func (o *QosData) HasMaxPacketLossRateUl() bool`

HasMaxPacketLossRateUl returns a boolean if a field has been set.

### SetMaxPacketLossRateUlNil

`func (o *QosData) SetMaxPacketLossRateUlNil(b bool)`

 SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil

### UnsetMaxPacketLossRateUl
`func (o *QosData) UnsetMaxPacketLossRateUl()`

UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
### GetDefQosFlowIndication

`func (o *QosData) GetDefQosFlowIndication() bool`

GetDefQosFlowIndication returns the DefQosFlowIndication field if non-nil, zero value otherwise.

### GetDefQosFlowIndicationOk

`func (o *QosData) GetDefQosFlowIndicationOk() (*bool, bool)`

GetDefQosFlowIndicationOk returns a tuple with the DefQosFlowIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefQosFlowIndication

`func (o *QosData) SetDefQosFlowIndication(v bool)`

SetDefQosFlowIndication sets DefQosFlowIndication field to given value.

### HasDefQosFlowIndication

`func (o *QosData) HasDefQosFlowIndication() bool`

HasDefQosFlowIndication returns a boolean if a field has been set.

### GetExtMaxDataBurstVol

`func (o *QosData) GetExtMaxDataBurstVol() int32`

GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field if non-nil, zero value otherwise.

### GetExtMaxDataBurstVolOk

`func (o *QosData) GetExtMaxDataBurstVolOk() (*int32, bool)`

GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMaxDataBurstVol

`func (o *QosData) SetExtMaxDataBurstVol(v int32)`

SetExtMaxDataBurstVol sets ExtMaxDataBurstVol field to given value.

### HasExtMaxDataBurstVol

`func (o *QosData) HasExtMaxDataBurstVol() bool`

HasExtMaxDataBurstVol returns a boolean if a field has been set.

### SetExtMaxDataBurstVolNil

`func (o *QosData) SetExtMaxDataBurstVolNil(b bool)`

 SetExtMaxDataBurstVolNil sets the value for ExtMaxDataBurstVol to be an explicit nil

### UnsetExtMaxDataBurstVol
`func (o *QosData) UnsetExtMaxDataBurstVol()`

UnsetExtMaxDataBurstVol ensures that no value is present for ExtMaxDataBurstVol, not even an explicit nil
### GetPacketDelayBudget

`func (o *QosData) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *QosData) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *QosData) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.

### HasPacketDelayBudget

`func (o *QosData) HasPacketDelayBudget() bool`

HasPacketDelayBudget returns a boolean if a field has been set.

### GetPacketErrorRate

`func (o *QosData) GetPacketErrorRate() string`

GetPacketErrorRate returns the PacketErrorRate field if non-nil, zero value otherwise.

### GetPacketErrorRateOk

`func (o *QosData) GetPacketErrorRateOk() (*string, bool)`

GetPacketErrorRateOk returns a tuple with the PacketErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketErrorRate

`func (o *QosData) SetPacketErrorRate(v string)`

SetPacketErrorRate sets PacketErrorRate field to given value.

### HasPacketErrorRate

`func (o *QosData) HasPacketErrorRate() bool`

HasPacketErrorRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


