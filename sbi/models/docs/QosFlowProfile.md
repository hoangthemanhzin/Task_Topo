# QosFlowProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | **int32** |  | 
**NonDynamic5Qi** | Pointer to [**NonDynamic5Qi**](NonDynamic5Qi.md) |  | [optional] 
**Dynamic5Qi** | Pointer to [**Dynamic5Qi**](Dynamic5Qi.md) |  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**GbrQosFlowInfo** | Pointer to [**GbrQosFlowInformation**](GbrQosFlowInformation.md) |  | [optional] 
**Rqa** | Pointer to [**ReflectiveQoSAttribute**](ReflectiveQoSAttribute.md) |  | [optional] 
**AdditionalQosFlowInfo** | Pointer to [**AdditionalQosFlowInfo**](AdditionalQosFlowInfo.md) |  | [optional] 
**QosMonitoringReq** | Pointer to [**QosMonitoringReq**](QosMonitoringReq.md) |  | [optional] 
**QosRepPeriod** | Pointer to **int32** |  | [optional] 

## Methods

### NewQosFlowProfile

`func NewQosFlowProfile(var5qi int32, ) *QosFlowProfile`

NewQosFlowProfile instantiates a new QosFlowProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosFlowProfileWithDefaults

`func NewQosFlowProfileWithDefaults() *QosFlowProfile`

NewQosFlowProfileWithDefaults instantiates a new QosFlowProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *QosFlowProfile) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *QosFlowProfile) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *QosFlowProfile) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.


### GetNonDynamic5Qi

`func (o *QosFlowProfile) GetNonDynamic5Qi() NonDynamic5Qi`

GetNonDynamic5Qi returns the NonDynamic5Qi field if non-nil, zero value otherwise.

### GetNonDynamic5QiOk

`func (o *QosFlowProfile) GetNonDynamic5QiOk() (*NonDynamic5Qi, bool)`

GetNonDynamic5QiOk returns a tuple with the NonDynamic5Qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDynamic5Qi

`func (o *QosFlowProfile) SetNonDynamic5Qi(v NonDynamic5Qi)`

SetNonDynamic5Qi sets NonDynamic5Qi field to given value.

### HasNonDynamic5Qi

`func (o *QosFlowProfile) HasNonDynamic5Qi() bool`

HasNonDynamic5Qi returns a boolean if a field has been set.

### GetDynamic5Qi

`func (o *QosFlowProfile) GetDynamic5Qi() Dynamic5Qi`

GetDynamic5Qi returns the Dynamic5Qi field if non-nil, zero value otherwise.

### GetDynamic5QiOk

`func (o *QosFlowProfile) GetDynamic5QiOk() (*Dynamic5Qi, bool)`

GetDynamic5QiOk returns a tuple with the Dynamic5Qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5Qi

`func (o *QosFlowProfile) SetDynamic5Qi(v Dynamic5Qi)`

SetDynamic5Qi sets Dynamic5Qi field to given value.

### HasDynamic5Qi

`func (o *QosFlowProfile) HasDynamic5Qi() bool`

HasDynamic5Qi returns a boolean if a field has been set.

### GetArp

`func (o *QosFlowProfile) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *QosFlowProfile) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *QosFlowProfile) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *QosFlowProfile) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetGbrQosFlowInfo

`func (o *QosFlowProfile) GetGbrQosFlowInfo() GbrQosFlowInformation`

GetGbrQosFlowInfo returns the GbrQosFlowInfo field if non-nil, zero value otherwise.

### GetGbrQosFlowInfoOk

`func (o *QosFlowProfile) GetGbrQosFlowInfoOk() (*GbrQosFlowInformation, bool)`

GetGbrQosFlowInfoOk returns a tuple with the GbrQosFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrQosFlowInfo

`func (o *QosFlowProfile) SetGbrQosFlowInfo(v GbrQosFlowInformation)`

SetGbrQosFlowInfo sets GbrQosFlowInfo field to given value.

### HasGbrQosFlowInfo

`func (o *QosFlowProfile) HasGbrQosFlowInfo() bool`

HasGbrQosFlowInfo returns a boolean if a field has been set.

### GetRqa

`func (o *QosFlowProfile) GetRqa() ReflectiveQoSAttribute`

GetRqa returns the Rqa field if non-nil, zero value otherwise.

### GetRqaOk

`func (o *QosFlowProfile) GetRqaOk() (*ReflectiveQoSAttribute, bool)`

GetRqaOk returns a tuple with the Rqa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRqa

`func (o *QosFlowProfile) SetRqa(v ReflectiveQoSAttribute)`

SetRqa sets Rqa field to given value.

### HasRqa

`func (o *QosFlowProfile) HasRqa() bool`

HasRqa returns a boolean if a field has been set.

### GetAdditionalQosFlowInfo

`func (o *QosFlowProfile) GetAdditionalQosFlowInfo() AdditionalQosFlowInfo`

GetAdditionalQosFlowInfo returns the AdditionalQosFlowInfo field if non-nil, zero value otherwise.

### GetAdditionalQosFlowInfoOk

`func (o *QosFlowProfile) GetAdditionalQosFlowInfoOk() (*AdditionalQosFlowInfo, bool)`

GetAdditionalQosFlowInfoOk returns a tuple with the AdditionalQosFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQosFlowInfo

`func (o *QosFlowProfile) SetAdditionalQosFlowInfo(v AdditionalQosFlowInfo)`

SetAdditionalQosFlowInfo sets AdditionalQosFlowInfo field to given value.

### HasAdditionalQosFlowInfo

`func (o *QosFlowProfile) HasAdditionalQosFlowInfo() bool`

HasAdditionalQosFlowInfo returns a boolean if a field has been set.

### GetQosMonitoringReq

`func (o *QosFlowProfile) GetQosMonitoringReq() QosMonitoringReq`

GetQosMonitoringReq returns the QosMonitoringReq field if non-nil, zero value otherwise.

### GetQosMonitoringReqOk

`func (o *QosFlowProfile) GetQosMonitoringReqOk() (*QosMonitoringReq, bool)`

GetQosMonitoringReqOk returns a tuple with the QosMonitoringReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonitoringReq

`func (o *QosFlowProfile) SetQosMonitoringReq(v QosMonitoringReq)`

SetQosMonitoringReq sets QosMonitoringReq field to given value.

### HasQosMonitoringReq

`func (o *QosFlowProfile) HasQosMonitoringReq() bool`

HasQosMonitoringReq returns a boolean if a field has been set.

### GetQosRepPeriod

`func (o *QosFlowProfile) GetQosRepPeriod() int32`

GetQosRepPeriod returns the QosRepPeriod field if non-nil, zero value otherwise.

### GetQosRepPeriodOk

`func (o *QosFlowProfile) GetQosRepPeriodOk() (*int32, bool)`

GetQosRepPeriodOk returns a tuple with the QosRepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRepPeriod

`func (o *QosFlowProfile) SetQosRepPeriod(v int32)`

SetQosRepPeriod sets QosRepPeriod field to given value.

### HasQosRepPeriod

`func (o *QosFlowProfile) HasQosRepPeriod() bool`

HasQosRepPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


