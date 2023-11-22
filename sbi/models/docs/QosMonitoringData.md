# QosMonitoringData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QmId** | **string** | Univocally identifies the QoS monitoring policy data within a PDU session. | 
**ReqQosMonParams** | [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) | indicates the UL packet delay, DL packet delay and/or round trip packet delay between the UE and the UPF is to be monitored when the QoS Monitoring for URLLC is enabled for the service data flow. | 
**RepFreqs** | [**[]ReportingFrequency**](ReportingFrequency.md) |  | 
**RepThreshDl** | Pointer to **int32** | Indicates the period of time in units of miliiseconds for DL packet delay. | [optional] 
**RepThreshUl** | Pointer to **int32** | Indicates the period of time in units of miliiseconds for UL packet delay. | [optional] 
**RepThreshRp** | Pointer to **int32** | Indicates the period of time in units of miliiseconds for round trip packet delay. | [optional] 
**WaitTime** | Pointer to **int32** |  | [optional] 
**RepPeriod** | Pointer to **int32** |  | [optional] 
**NotifyUri** | Pointer to **string** |  | [optional] 
**NotifyCorreId** | Pointer to **string** |  | [optional] 

## Methods

### NewQosMonitoringData

`func NewQosMonitoringData(qmId string, reqQosMonParams []RequestedQosMonitoringParameter, repFreqs []ReportingFrequency, ) *QosMonitoringData`

NewQosMonitoringData instantiates a new QosMonitoringData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosMonitoringDataWithDefaults

`func NewQosMonitoringDataWithDefaults() *QosMonitoringData`

NewQosMonitoringDataWithDefaults instantiates a new QosMonitoringData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQmId

`func (o *QosMonitoringData) GetQmId() string`

GetQmId returns the QmId field if non-nil, zero value otherwise.

### GetQmIdOk

`func (o *QosMonitoringData) GetQmIdOk() (*string, bool)`

GetQmIdOk returns a tuple with the QmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQmId

`func (o *QosMonitoringData) SetQmId(v string)`

SetQmId sets QmId field to given value.


### GetReqQosMonParams

`func (o *QosMonitoringData) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *QosMonitoringData) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *QosMonitoringData) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.


### GetRepFreqs

`func (o *QosMonitoringData) GetRepFreqs() []ReportingFrequency`

GetRepFreqs returns the RepFreqs field if non-nil, zero value otherwise.

### GetRepFreqsOk

`func (o *QosMonitoringData) GetRepFreqsOk() (*[]ReportingFrequency, bool)`

GetRepFreqsOk returns a tuple with the RepFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepFreqs

`func (o *QosMonitoringData) SetRepFreqs(v []ReportingFrequency)`

SetRepFreqs sets RepFreqs field to given value.


### GetRepThreshDl

`func (o *QosMonitoringData) GetRepThreshDl() int32`

GetRepThreshDl returns the RepThreshDl field if non-nil, zero value otherwise.

### GetRepThreshDlOk

`func (o *QosMonitoringData) GetRepThreshDlOk() (*int32, bool)`

GetRepThreshDlOk returns a tuple with the RepThreshDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshDl

`func (o *QosMonitoringData) SetRepThreshDl(v int32)`

SetRepThreshDl sets RepThreshDl field to given value.

### HasRepThreshDl

`func (o *QosMonitoringData) HasRepThreshDl() bool`

HasRepThreshDl returns a boolean if a field has been set.

### SetRepThreshDlNil

`func (o *QosMonitoringData) SetRepThreshDlNil(b bool)`

 SetRepThreshDlNil sets the value for RepThreshDl to be an explicit nil

### UnsetRepThreshDl
`func (o *QosMonitoringData) UnsetRepThreshDl()`

UnsetRepThreshDl ensures that no value is present for RepThreshDl, not even an explicit nil
### GetRepThreshUl

`func (o *QosMonitoringData) GetRepThreshUl() int32`

GetRepThreshUl returns the RepThreshUl field if non-nil, zero value otherwise.

### GetRepThreshUlOk

`func (o *QosMonitoringData) GetRepThreshUlOk() (*int32, bool)`

GetRepThreshUlOk returns a tuple with the RepThreshUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshUl

`func (o *QosMonitoringData) SetRepThreshUl(v int32)`

SetRepThreshUl sets RepThreshUl field to given value.

### HasRepThreshUl

`func (o *QosMonitoringData) HasRepThreshUl() bool`

HasRepThreshUl returns a boolean if a field has been set.

### SetRepThreshUlNil

`func (o *QosMonitoringData) SetRepThreshUlNil(b bool)`

 SetRepThreshUlNil sets the value for RepThreshUl to be an explicit nil

### UnsetRepThreshUl
`func (o *QosMonitoringData) UnsetRepThreshUl()`

UnsetRepThreshUl ensures that no value is present for RepThreshUl, not even an explicit nil
### GetRepThreshRp

`func (o *QosMonitoringData) GetRepThreshRp() int32`

GetRepThreshRp returns the RepThreshRp field if non-nil, zero value otherwise.

### GetRepThreshRpOk

`func (o *QosMonitoringData) GetRepThreshRpOk() (*int32, bool)`

GetRepThreshRpOk returns a tuple with the RepThreshRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshRp

`func (o *QosMonitoringData) SetRepThreshRp(v int32)`

SetRepThreshRp sets RepThreshRp field to given value.

### HasRepThreshRp

`func (o *QosMonitoringData) HasRepThreshRp() bool`

HasRepThreshRp returns a boolean if a field has been set.

### SetRepThreshRpNil

`func (o *QosMonitoringData) SetRepThreshRpNil(b bool)`

 SetRepThreshRpNil sets the value for RepThreshRp to be an explicit nil

### UnsetRepThreshRp
`func (o *QosMonitoringData) UnsetRepThreshRp()`

UnsetRepThreshRp ensures that no value is present for RepThreshRp, not even an explicit nil
### GetWaitTime

`func (o *QosMonitoringData) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *QosMonitoringData) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *QosMonitoringData) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *QosMonitoringData) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

### SetWaitTimeNil

`func (o *QosMonitoringData) SetWaitTimeNil(b bool)`

 SetWaitTimeNil sets the value for WaitTime to be an explicit nil

### UnsetWaitTime
`func (o *QosMonitoringData) UnsetWaitTime()`

UnsetWaitTime ensures that no value is present for WaitTime, not even an explicit nil
### GetRepPeriod

`func (o *QosMonitoringData) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *QosMonitoringData) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *QosMonitoringData) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *QosMonitoringData) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### SetRepPeriodNil

`func (o *QosMonitoringData) SetRepPeriodNil(b bool)`

 SetRepPeriodNil sets the value for RepPeriod to be an explicit nil

### UnsetRepPeriod
`func (o *QosMonitoringData) UnsetRepPeriod()`

UnsetRepPeriod ensures that no value is present for RepPeriod, not even an explicit nil
### GetNotifyUri

`func (o *QosMonitoringData) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *QosMonitoringData) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *QosMonitoringData) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.

### HasNotifyUri

`func (o *QosMonitoringData) HasNotifyUri() bool`

HasNotifyUri returns a boolean if a field has been set.

### GetNotifyCorreId

`func (o *QosMonitoringData) GetNotifyCorreId() string`

GetNotifyCorreId returns the NotifyCorreId field if non-nil, zero value otherwise.

### GetNotifyCorreIdOk

`func (o *QosMonitoringData) GetNotifyCorreIdOk() (*string, bool)`

GetNotifyCorreIdOk returns a tuple with the NotifyCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorreId

`func (o *QosMonitoringData) SetNotifyCorreId(v string)`

SetNotifyCorreId sets NotifyCorreId field to given value.

### HasNotifyCorreId

`func (o *QosMonitoringData) HasNotifyCorreId() bool`

HasNotifyCorreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


