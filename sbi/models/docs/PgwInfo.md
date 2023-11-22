# PgwInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**PgwFqdn** | **string** |  | 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPgwInfo

`func NewPgwInfo(dnn string, pgwFqdn string, ) *PgwInfo`

NewPgwInfo instantiates a new PgwInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgwInfoWithDefaults

`func NewPgwInfoWithDefaults() *PgwInfo`

NewPgwInfoWithDefaults instantiates a new PgwInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PgwInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PgwInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PgwInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetPgwFqdn

`func (o *PgwInfo) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *PgwInfo) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *PgwInfo) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.


### GetPlmnId

`func (o *PgwInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PgwInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PgwInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PgwInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetEpdgInd

`func (o *PgwInfo) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *PgwInfo) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *PgwInfo) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *PgwInfo) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


