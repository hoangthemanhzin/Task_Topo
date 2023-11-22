# PgwInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**PgwFqdn** | **string** |  | 
**PlmnId** | Pointer to [**PlmnId1**](PlmnId1.md) |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPgwInfo1

`func NewPgwInfo1(dnn string, pgwFqdn string, ) *PgwInfo1`

NewPgwInfo1 instantiates a new PgwInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgwInfo1WithDefaults

`func NewPgwInfo1WithDefaults() *PgwInfo1`

NewPgwInfo1WithDefaults instantiates a new PgwInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PgwInfo1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PgwInfo1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PgwInfo1) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetPgwFqdn

`func (o *PgwInfo1) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *PgwInfo1) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *PgwInfo1) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.


### GetPlmnId

`func (o *PgwInfo1) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PgwInfo1) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PgwInfo1) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PgwInfo1) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetEpdgInd

`func (o *PgwInfo1) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *PgwInfo1) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *PgwInfo1) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *PgwInfo1) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


