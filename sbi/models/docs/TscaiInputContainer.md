# TscaiInputContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periodicity** | Pointer to **int32** |  | [optional] 
**BurstArrivalTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTscaiInputContainer

`func NewTscaiInputContainer() *TscaiInputContainer`

NewTscaiInputContainer instantiates a new TscaiInputContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscaiInputContainerWithDefaults

`func NewTscaiInputContainerWithDefaults() *TscaiInputContainer`

NewTscaiInputContainerWithDefaults instantiates a new TscaiInputContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodicity

`func (o *TscaiInputContainer) GetPeriodicity() int32`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *TscaiInputContainer) GetPeriodicityOk() (*int32, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *TscaiInputContainer) SetPeriodicity(v int32)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *TscaiInputContainer) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### GetBurstArrivalTime

`func (o *TscaiInputContainer) GetBurstArrivalTime() time.Time`

GetBurstArrivalTime returns the BurstArrivalTime field if non-nil, zero value otherwise.

### GetBurstArrivalTimeOk

`func (o *TscaiInputContainer) GetBurstArrivalTimeOk() (*time.Time, bool)`

GetBurstArrivalTimeOk returns a tuple with the BurstArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstArrivalTime

`func (o *TscaiInputContainer) SetBurstArrivalTime(v time.Time)`

SetBurstArrivalTime sets BurstArrivalTime field to given value.

### HasBurstArrivalTime

`func (o *TscaiInputContainer) HasBurstArrivalTime() bool`

HasBurstArrivalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


