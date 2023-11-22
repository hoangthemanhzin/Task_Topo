# DdnFailureSubInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyCorrelationId** | **string** |  | 
**DddTrafficDescriptorList** | Pointer to [**[]DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 

## Methods

### NewDdnFailureSubInfo

`func NewDdnFailureSubInfo(notifyCorrelationId string, ) *DdnFailureSubInfo`

NewDdnFailureSubInfo instantiates a new DdnFailureSubInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnFailureSubInfoWithDefaults

`func NewDdnFailureSubInfoWithDefaults() *DdnFailureSubInfo`

NewDdnFailureSubInfoWithDefaults instantiates a new DdnFailureSubInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyCorrelationId

`func (o *DdnFailureSubInfo) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *DdnFailureSubInfo) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *DdnFailureSubInfo) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.


### GetDddTrafficDescriptorList

`func (o *DdnFailureSubInfo) GetDddTrafficDescriptorList() []DddTrafficDescriptor`

GetDddTrafficDescriptorList returns the DddTrafficDescriptorList field if non-nil, zero value otherwise.

### GetDddTrafficDescriptorListOk

`func (o *DdnFailureSubInfo) GetDddTrafficDescriptorListOk() (*[]DddTrafficDescriptor, bool)`

GetDddTrafficDescriptorListOk returns a tuple with the DddTrafficDescriptorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTrafficDescriptorList

`func (o *DdnFailureSubInfo) SetDddTrafficDescriptorList(v []DddTrafficDescriptor)`

SetDddTrafficDescriptorList sets DddTrafficDescriptorList field to given value.

### HasDddTrafficDescriptorList

`func (o *DdnFailureSubInfo) HasDddTrafficDescriptorList() bool`

HasDddTrafficDescriptorList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


