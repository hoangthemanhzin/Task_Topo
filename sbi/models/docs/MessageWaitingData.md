# MessageWaitingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MwdList** | Pointer to [**[]SmscData**](SmscData.md) |  | [optional] 

## Methods

### NewMessageWaitingData

`func NewMessageWaitingData() *MessageWaitingData`

NewMessageWaitingData instantiates a new MessageWaitingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWaitingDataWithDefaults

`func NewMessageWaitingDataWithDefaults() *MessageWaitingData`

NewMessageWaitingDataWithDefaults instantiates a new MessageWaitingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMwdList

`func (o *MessageWaitingData) GetMwdList() []SmscData`

GetMwdList returns the MwdList field if non-nil, zero value otherwise.

### GetMwdListOk

`func (o *MessageWaitingData) GetMwdListOk() (*[]SmscData, bool)`

GetMwdListOk returns a tuple with the MwdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMwdList

`func (o *MessageWaitingData) SetMwdList(v []SmscData)`

SetMwdList sets MwdList field to given value.

### HasMwdList

`func (o *MessageWaitingData) HasMwdList() bool`

HasMwdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


