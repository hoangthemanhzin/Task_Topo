# TimePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**Periodicity**](Periodicity.md) |  | 
**MaxNumPeriod** | Pointer to **int32** |  | [optional] 

## Methods

### NewTimePeriod

`func NewTimePeriod(period Periodicity, ) *TimePeriod`

NewTimePeriod instantiates a new TimePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePeriodWithDefaults

`func NewTimePeriodWithDefaults() *TimePeriod`

NewTimePeriodWithDefaults instantiates a new TimePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *TimePeriod) GetPeriod() Periodicity`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TimePeriod) GetPeriodOk() (*Periodicity, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TimePeriod) SetPeriod(v Periodicity)`

SetPeriod sets Period field to given value.


### GetMaxNumPeriod

`func (o *TimePeriod) GetMaxNumPeriod() int32`

GetMaxNumPeriod returns the MaxNumPeriod field if non-nil, zero value otherwise.

### GetMaxNumPeriodOk

`func (o *TimePeriod) GetMaxNumPeriodOk() (*int32, bool)`

GetMaxNumPeriodOk returns a tuple with the MaxNumPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumPeriod

`func (o *TimePeriod) SetMaxNumPeriod(v int32)`

SetMaxNumPeriod sets MaxNumPeriod field to given value.

### HasMaxNumPeriod

`func (o *TimePeriod) HasMaxNumPeriod() bool`

HasMaxNumPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


