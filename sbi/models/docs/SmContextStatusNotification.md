# SmContextStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusInfo** | [**StatusInfo**](StatusInfo.md) |  | 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**DdnFailureStatus** | Pointer to **bool** |  | [optional] [default to false]
**NotifyCorrelationIdsForddnFailure** | Pointer to **[]string** |  | [optional] 
**NewIntermediateSmfId** | Pointer to **string** |  | [optional] 
**NewSmfId** | Pointer to **string** |  | [optional] 
**NewSmfSetId** | Pointer to **string** |  | [optional] 
**OldSmfId** | Pointer to **string** |  | [optional] 
**OldSmContextRef** | Pointer to **string** |  | [optional] 
**AltAnchorSmfUri** | Pointer to **string** |  | [optional] 
**AltAnchorSmfId** | Pointer to **string** |  | [optional] 

## Methods

### NewSmContextStatusNotification

`func NewSmContextStatusNotification(statusInfo StatusInfo, ) *SmContextStatusNotification`

NewSmContextStatusNotification instantiates a new SmContextStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextStatusNotificationWithDefaults

`func NewSmContextStatusNotificationWithDefaults() *SmContextStatusNotification`

NewSmContextStatusNotificationWithDefaults instantiates a new SmContextStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusInfo

`func (o *SmContextStatusNotification) GetStatusInfo() StatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *SmContextStatusNotification) GetStatusInfoOk() (*StatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *SmContextStatusNotification) SetStatusInfo(v StatusInfo)`

SetStatusInfo sets StatusInfo field to given value.


### GetSmallDataRateStatus

`func (o *SmContextStatusNotification) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextStatusNotification) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextStatusNotification) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextStatusNotification) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextStatusNotification) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextStatusNotification) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextStatusNotification) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextStatusNotification) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetDdnFailureStatus

`func (o *SmContextStatusNotification) GetDdnFailureStatus() bool`

GetDdnFailureStatus returns the DdnFailureStatus field if non-nil, zero value otherwise.

### GetDdnFailureStatusOk

`func (o *SmContextStatusNotification) GetDdnFailureStatusOk() (*bool, bool)`

GetDdnFailureStatusOk returns a tuple with the DdnFailureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnFailureStatus

`func (o *SmContextStatusNotification) SetDdnFailureStatus(v bool)`

SetDdnFailureStatus sets DdnFailureStatus field to given value.

### HasDdnFailureStatus

`func (o *SmContextStatusNotification) HasDdnFailureStatus() bool`

HasDdnFailureStatus returns a boolean if a field has been set.

### GetNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailure() []string`

GetNotifyCorrelationIdsForddnFailure returns the NotifyCorrelationIdsForddnFailure field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdsForddnFailureOk

`func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailureOk() (*[]string, bool)`

GetNotifyCorrelationIdsForddnFailureOk returns a tuple with the NotifyCorrelationIdsForddnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) SetNotifyCorrelationIdsForddnFailure(v []string)`

SetNotifyCorrelationIdsForddnFailure sets NotifyCorrelationIdsForddnFailure field to given value.

### HasNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) HasNotifyCorrelationIdsForddnFailure() bool`

HasNotifyCorrelationIdsForddnFailure returns a boolean if a field has been set.

### GetNewIntermediateSmfId

`func (o *SmContextStatusNotification) GetNewIntermediateSmfId() string`

GetNewIntermediateSmfId returns the NewIntermediateSmfId field if non-nil, zero value otherwise.

### GetNewIntermediateSmfIdOk

`func (o *SmContextStatusNotification) GetNewIntermediateSmfIdOk() (*string, bool)`

GetNewIntermediateSmfIdOk returns a tuple with the NewIntermediateSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIntermediateSmfId

`func (o *SmContextStatusNotification) SetNewIntermediateSmfId(v string)`

SetNewIntermediateSmfId sets NewIntermediateSmfId field to given value.

### HasNewIntermediateSmfId

`func (o *SmContextStatusNotification) HasNewIntermediateSmfId() bool`

HasNewIntermediateSmfId returns a boolean if a field has been set.

### GetNewSmfId

`func (o *SmContextStatusNotification) GetNewSmfId() string`

GetNewSmfId returns the NewSmfId field if non-nil, zero value otherwise.

### GetNewSmfIdOk

`func (o *SmContextStatusNotification) GetNewSmfIdOk() (*string, bool)`

GetNewSmfIdOk returns a tuple with the NewSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfId

`func (o *SmContextStatusNotification) SetNewSmfId(v string)`

SetNewSmfId sets NewSmfId field to given value.

### HasNewSmfId

`func (o *SmContextStatusNotification) HasNewSmfId() bool`

HasNewSmfId returns a boolean if a field has been set.

### GetNewSmfSetId

`func (o *SmContextStatusNotification) GetNewSmfSetId() string`

GetNewSmfSetId returns the NewSmfSetId field if non-nil, zero value otherwise.

### GetNewSmfSetIdOk

`func (o *SmContextStatusNotification) GetNewSmfSetIdOk() (*string, bool)`

GetNewSmfSetIdOk returns a tuple with the NewSmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfSetId

`func (o *SmContextStatusNotification) SetNewSmfSetId(v string)`

SetNewSmfSetId sets NewSmfSetId field to given value.

### HasNewSmfSetId

`func (o *SmContextStatusNotification) HasNewSmfSetId() bool`

HasNewSmfSetId returns a boolean if a field has been set.

### GetOldSmfId

`func (o *SmContextStatusNotification) GetOldSmfId() string`

GetOldSmfId returns the OldSmfId field if non-nil, zero value otherwise.

### GetOldSmfIdOk

`func (o *SmContextStatusNotification) GetOldSmfIdOk() (*string, bool)`

GetOldSmfIdOk returns a tuple with the OldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmfId

`func (o *SmContextStatusNotification) SetOldSmfId(v string)`

SetOldSmfId sets OldSmfId field to given value.

### HasOldSmfId

`func (o *SmContextStatusNotification) HasOldSmfId() bool`

HasOldSmfId returns a boolean if a field has been set.

### GetOldSmContextRef

`func (o *SmContextStatusNotification) GetOldSmContextRef() string`

GetOldSmContextRef returns the OldSmContextRef field if non-nil, zero value otherwise.

### GetOldSmContextRefOk

`func (o *SmContextStatusNotification) GetOldSmContextRefOk() (*string, bool)`

GetOldSmContextRefOk returns a tuple with the OldSmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmContextRef

`func (o *SmContextStatusNotification) SetOldSmContextRef(v string)`

SetOldSmContextRef sets OldSmContextRef field to given value.

### HasOldSmContextRef

`func (o *SmContextStatusNotification) HasOldSmContextRef() bool`

HasOldSmContextRef returns a boolean if a field has been set.

### GetAltAnchorSmfUri

`func (o *SmContextStatusNotification) GetAltAnchorSmfUri() string`

GetAltAnchorSmfUri returns the AltAnchorSmfUri field if non-nil, zero value otherwise.

### GetAltAnchorSmfUriOk

`func (o *SmContextStatusNotification) GetAltAnchorSmfUriOk() (*string, bool)`

GetAltAnchorSmfUriOk returns a tuple with the AltAnchorSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAnchorSmfUri

`func (o *SmContextStatusNotification) SetAltAnchorSmfUri(v string)`

SetAltAnchorSmfUri sets AltAnchorSmfUri field to given value.

### HasAltAnchorSmfUri

`func (o *SmContextStatusNotification) HasAltAnchorSmfUri() bool`

HasAltAnchorSmfUri returns a boolean if a field has been set.

### GetAltAnchorSmfId

`func (o *SmContextStatusNotification) GetAltAnchorSmfId() string`

GetAltAnchorSmfId returns the AltAnchorSmfId field if non-nil, zero value otherwise.

### GetAltAnchorSmfIdOk

`func (o *SmContextStatusNotification) GetAltAnchorSmfIdOk() (*string, bool)`

GetAltAnchorSmfIdOk returns a tuple with the AltAnchorSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAnchorSmfId

`func (o *SmContextStatusNotification) SetAltAnchorSmfId(v string)`

SetAltAnchorSmfId sets AltAnchorSmfId field to given value.

### HasAltAnchorSmfId

`func (o *SmContextStatusNotification) HasAltAnchorSmfId() bool`

HasAltAnchorSmfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


