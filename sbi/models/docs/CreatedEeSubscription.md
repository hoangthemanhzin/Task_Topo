# CreatedEeSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 
**NumberOfUes** | Pointer to **int32** |  | [optional] 
**EventReports** | Pointer to [**[]MonitoringReport**](MonitoringReport.md) |  | [optional] 
**EpcStatusInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreatedEeSubscription

`func NewCreatedEeSubscription(eeSubscription EeSubscription, ) *CreatedEeSubscription`

NewCreatedEeSubscription instantiates a new CreatedEeSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedEeSubscriptionWithDefaults

`func NewCreatedEeSubscriptionWithDefaults() *CreatedEeSubscription`

NewCreatedEeSubscriptionWithDefaults instantiates a new CreatedEeSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEeSubscription

`func (o *CreatedEeSubscription) GetEeSubscription() EeSubscription`

GetEeSubscription returns the EeSubscription field if non-nil, zero value otherwise.

### GetEeSubscriptionOk

`func (o *CreatedEeSubscription) GetEeSubscriptionOk() (*EeSubscription, bool)`

GetEeSubscriptionOk returns a tuple with the EeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeSubscription

`func (o *CreatedEeSubscription) SetEeSubscription(v EeSubscription)`

SetEeSubscription sets EeSubscription field to given value.


### GetNumberOfUes

`func (o *CreatedEeSubscription) GetNumberOfUes() int32`

GetNumberOfUes returns the NumberOfUes field if non-nil, zero value otherwise.

### GetNumberOfUesOk

`func (o *CreatedEeSubscription) GetNumberOfUesOk() (*int32, bool)`

GetNumberOfUesOk returns a tuple with the NumberOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUes

`func (o *CreatedEeSubscription) SetNumberOfUes(v int32)`

SetNumberOfUes sets NumberOfUes field to given value.

### HasNumberOfUes

`func (o *CreatedEeSubscription) HasNumberOfUes() bool`

HasNumberOfUes returns a boolean if a field has been set.

### GetEventReports

`func (o *CreatedEeSubscription) GetEventReports() []MonitoringReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *CreatedEeSubscription) GetEventReportsOk() (*[]MonitoringReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *CreatedEeSubscription) SetEventReports(v []MonitoringReport)`

SetEventReports sets EventReports field to given value.

### HasEventReports

`func (o *CreatedEeSubscription) HasEventReports() bool`

HasEventReports returns a boolean if a field has been set.

### GetEpcStatusInd

`func (o *CreatedEeSubscription) GetEpcStatusInd() bool`

GetEpcStatusInd returns the EpcStatusInd field if non-nil, zero value otherwise.

### GetEpcStatusIndOk

`func (o *CreatedEeSubscription) GetEpcStatusIndOk() (*bool, bool)`

GetEpcStatusIndOk returns a tuple with the EpcStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpcStatusInd

`func (o *CreatedEeSubscription) SetEpcStatusInd(v bool)`

SetEpcStatusInd sets EpcStatusInd field to given value.

### HasEpcStatusInd

`func (o *CreatedEeSubscription) HasEpcStatusInd() bool`

HasEpcStatusInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


