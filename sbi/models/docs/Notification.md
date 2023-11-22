# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BdtRefId** | **string** | string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154. | 
**CandPolicies** | Pointer to [**[]TransferPolicy**](TransferPolicy.md) | Contains a list of the candidate transfer policies from which the AF may select a new transfer policy due to a network performance is below the criteria set by the operator. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**TimeWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification(bdtRefId string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBdtRefId

`func (o *Notification) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *Notification) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *Notification) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.


### GetCandPolicies

`func (o *Notification) GetCandPolicies() []TransferPolicy`

GetCandPolicies returns the CandPolicies field if non-nil, zero value otherwise.

### GetCandPoliciesOk

`func (o *Notification) GetCandPoliciesOk() (*[]TransferPolicy, bool)`

GetCandPoliciesOk returns a tuple with the CandPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandPolicies

`func (o *Notification) SetCandPolicies(v []TransferPolicy)`

SetCandPolicies sets CandPolicies field to given value.

### HasCandPolicies

`func (o *Notification) HasCandPolicies() bool`

HasCandPolicies returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *Notification) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *Notification) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *Notification) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *Notification) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetTimeWindow

`func (o *Notification) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *Notification) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *Notification) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *Notification) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


