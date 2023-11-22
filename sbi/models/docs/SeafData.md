# SeafData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NgKsi** | [**NgKsi**](NgKsi.md) |  | 
**KeyAmf** | [**KeyAmf**](KeyAmf.md) |  | 
**Nh** | Pointer to **string** |  | [optional] 
**Ncc** | Pointer to **int32** |  | [optional] 
**KeyAmfChangeInd** | Pointer to **bool** |  | [optional] 
**KeyAmfHDerivationInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeafData

`func NewSeafData(ngKsi NgKsi, keyAmf KeyAmf, ) *SeafData`

NewSeafData instantiates a new SeafData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeafDataWithDefaults

`func NewSeafDataWithDefaults() *SeafData`

NewSeafDataWithDefaults instantiates a new SeafData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNgKsi

`func (o *SeafData) GetNgKsi() NgKsi`

GetNgKsi returns the NgKsi field if non-nil, zero value otherwise.

### GetNgKsiOk

`func (o *SeafData) GetNgKsiOk() (*NgKsi, bool)`

GetNgKsiOk returns a tuple with the NgKsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgKsi

`func (o *SeafData) SetNgKsi(v NgKsi)`

SetNgKsi sets NgKsi field to given value.


### GetKeyAmf

`func (o *SeafData) GetKeyAmf() KeyAmf`

GetKeyAmf returns the KeyAmf field if non-nil, zero value otherwise.

### GetKeyAmfOk

`func (o *SeafData) GetKeyAmfOk() (*KeyAmf, bool)`

GetKeyAmfOk returns a tuple with the KeyAmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAmf

`func (o *SeafData) SetKeyAmf(v KeyAmf)`

SetKeyAmf sets KeyAmf field to given value.


### GetNh

`func (o *SeafData) GetNh() string`

GetNh returns the Nh field if non-nil, zero value otherwise.

### GetNhOk

`func (o *SeafData) GetNhOk() (*string, bool)`

GetNhOk returns a tuple with the Nh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNh

`func (o *SeafData) SetNh(v string)`

SetNh sets Nh field to given value.

### HasNh

`func (o *SeafData) HasNh() bool`

HasNh returns a boolean if a field has been set.

### GetNcc

`func (o *SeafData) GetNcc() int32`

GetNcc returns the Ncc field if non-nil, zero value otherwise.

### GetNccOk

`func (o *SeafData) GetNccOk() (*int32, bool)`

GetNccOk returns a tuple with the Ncc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcc

`func (o *SeafData) SetNcc(v int32)`

SetNcc sets Ncc field to given value.

### HasNcc

`func (o *SeafData) HasNcc() bool`

HasNcc returns a boolean if a field has been set.

### GetKeyAmfChangeInd

`func (o *SeafData) GetKeyAmfChangeInd() bool`

GetKeyAmfChangeInd returns the KeyAmfChangeInd field if non-nil, zero value otherwise.

### GetKeyAmfChangeIndOk

`func (o *SeafData) GetKeyAmfChangeIndOk() (*bool, bool)`

GetKeyAmfChangeIndOk returns a tuple with the KeyAmfChangeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAmfChangeInd

`func (o *SeafData) SetKeyAmfChangeInd(v bool)`

SetKeyAmfChangeInd sets KeyAmfChangeInd field to given value.

### HasKeyAmfChangeInd

`func (o *SeafData) HasKeyAmfChangeInd() bool`

HasKeyAmfChangeInd returns a boolean if a field has been set.

### GetKeyAmfHDerivationInd

`func (o *SeafData) GetKeyAmfHDerivationInd() bool`

GetKeyAmfHDerivationInd returns the KeyAmfHDerivationInd field if non-nil, zero value otherwise.

### GetKeyAmfHDerivationIndOk

`func (o *SeafData) GetKeyAmfHDerivationIndOk() (*bool, bool)`

GetKeyAmfHDerivationIndOk returns a tuple with the KeyAmfHDerivationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAmfHDerivationInd

`func (o *SeafData) SetKeyAmfHDerivationInd(v bool)`

SetKeyAmfHDerivationInd sets KeyAmfHDerivationInd field to given value.

### HasKeyAmfHDerivationInd

`func (o *SeafData) HasKeyAmfHDerivationInd() bool`

HasKeyAmfHDerivationInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


