# SequenceNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqnScheme** | Pointer to [**SqnScheme**](SqnScheme.md) |  | [optional] 
**Sqn** | Pointer to **string** |  | [optional] 
**LastIndexes** | Pointer to **map[string]int32** |  | [optional] 
**IndLength** | Pointer to **int32** |  | [optional] 
**DifSign** | Pointer to [**Sign**](Sign.md) |  | [optional] 

## Methods

### NewSequenceNumber

`func NewSequenceNumber() *SequenceNumber`

NewSequenceNumber instantiates a new SequenceNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSequenceNumberWithDefaults

`func NewSequenceNumberWithDefaults() *SequenceNumber`

NewSequenceNumberWithDefaults instantiates a new SequenceNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqnScheme

`func (o *SequenceNumber) GetSqnScheme() SqnScheme`

GetSqnScheme returns the SqnScheme field if non-nil, zero value otherwise.

### GetSqnSchemeOk

`func (o *SequenceNumber) GetSqnSchemeOk() (*SqnScheme, bool)`

GetSqnSchemeOk returns a tuple with the SqnScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqnScheme

`func (o *SequenceNumber) SetSqnScheme(v SqnScheme)`

SetSqnScheme sets SqnScheme field to given value.

### HasSqnScheme

`func (o *SequenceNumber) HasSqnScheme() bool`

HasSqnScheme returns a boolean if a field has been set.

### GetSqn

`func (o *SequenceNumber) GetSqn() string`

GetSqn returns the Sqn field if non-nil, zero value otherwise.

### GetSqnOk

`func (o *SequenceNumber) GetSqnOk() (*string, bool)`

GetSqnOk returns a tuple with the Sqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqn

`func (o *SequenceNumber) SetSqn(v string)`

SetSqn sets Sqn field to given value.

### HasSqn

`func (o *SequenceNumber) HasSqn() bool`

HasSqn returns a boolean if a field has been set.

### GetLastIndexes

`func (o *SequenceNumber) GetLastIndexes() map[string]int32`

GetLastIndexes returns the LastIndexes field if non-nil, zero value otherwise.

### GetLastIndexesOk

`func (o *SequenceNumber) GetLastIndexesOk() (*map[string]int32, bool)`

GetLastIndexesOk returns a tuple with the LastIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexes

`func (o *SequenceNumber) SetLastIndexes(v map[string]int32)`

SetLastIndexes sets LastIndexes field to given value.

### HasLastIndexes

`func (o *SequenceNumber) HasLastIndexes() bool`

HasLastIndexes returns a boolean if a field has been set.

### GetIndLength

`func (o *SequenceNumber) GetIndLength() int32`

GetIndLength returns the IndLength field if non-nil, zero value otherwise.

### GetIndLengthOk

`func (o *SequenceNumber) GetIndLengthOk() (*int32, bool)`

GetIndLengthOk returns a tuple with the IndLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndLength

`func (o *SequenceNumber) SetIndLength(v int32)`

SetIndLength sets IndLength field to given value.

### HasIndLength

`func (o *SequenceNumber) HasIndLength() bool`

HasIndLength returns a boolean if a field has been set.

### GetDifSign

`func (o *SequenceNumber) GetDifSign() Sign`

GetDifSign returns the DifSign field if non-nil, zero value otherwise.

### GetDifSignOk

`func (o *SequenceNumber) GetDifSignOk() (*Sign, bool)`

GetDifSignOk returns a tuple with the DifSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifSign

`func (o *SequenceNumber) SetDifSign(v Sign)`

SetDifSign sets DifSign field to given value.

### HasDifSign

`func (o *SequenceNumber) HasDifSign() bool`

HasDifSign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


