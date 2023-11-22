# MoExpDataCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counter** | **int32** |  | 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMoExpDataCounter

`func NewMoExpDataCounter(counter int32, ) *MoExpDataCounter`

NewMoExpDataCounter instantiates a new MoExpDataCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoExpDataCounterWithDefaults

`func NewMoExpDataCounterWithDefaults() *MoExpDataCounter`

NewMoExpDataCounterWithDefaults instantiates a new MoExpDataCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounter

`func (o *MoExpDataCounter) GetCounter() int32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MoExpDataCounter) GetCounterOk() (*int32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MoExpDataCounter) SetCounter(v int32)`

SetCounter sets Counter field to given value.


### GetTimeStamp

`func (o *MoExpDataCounter) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *MoExpDataCounter) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *MoExpDataCounter) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *MoExpDataCounter) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


