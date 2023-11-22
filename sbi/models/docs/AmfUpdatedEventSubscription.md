# AmfUpdatedEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**AmfEventSubscription**](AmfEventSubscription.md) |  | 
**ReportList** | Pointer to [**[]AmfEventReport**](AmfEventReport.md) |  | [optional] 

## Methods

### NewAmfUpdatedEventSubscription

`func NewAmfUpdatedEventSubscription(subscription AmfEventSubscription, ) *AmfUpdatedEventSubscription`

NewAmfUpdatedEventSubscription instantiates a new AmfUpdatedEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfUpdatedEventSubscriptionWithDefaults

`func NewAmfUpdatedEventSubscriptionWithDefaults() *AmfUpdatedEventSubscription`

NewAmfUpdatedEventSubscriptionWithDefaults instantiates a new AmfUpdatedEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *AmfUpdatedEventSubscription) GetSubscription() AmfEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AmfUpdatedEventSubscription) GetSubscriptionOk() (*AmfEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AmfUpdatedEventSubscription) SetSubscription(v AmfEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetReportList

`func (o *AmfUpdatedEventSubscription) GetReportList() []AmfEventReport`

GetReportList returns the ReportList field if non-nil, zero value otherwise.

### GetReportListOk

`func (o *AmfUpdatedEventSubscription) GetReportListOk() (*[]AmfEventReport, bool)`

GetReportListOk returns a tuple with the ReportList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportList

`func (o *AmfUpdatedEventSubscription) SetReportList(v []AmfEventReport)`

SetReportList sets ReportList field to given value.

### HasReportList

`func (o *AmfUpdatedEventSubscription) HasReportList() bool`

HasReportList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


