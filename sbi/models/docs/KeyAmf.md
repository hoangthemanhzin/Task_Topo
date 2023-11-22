# KeyAmf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | [**KeyAmfType**](KeyAmfType.md) |  | 
**KeyVal** | **string** |  | 

## Methods

### NewKeyAmf

`func NewKeyAmf(keyType KeyAmfType, keyVal string, ) *KeyAmf`

NewKeyAmf instantiates a new KeyAmf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyAmfWithDefaults

`func NewKeyAmfWithDefaults() *KeyAmf`

NewKeyAmfWithDefaults instantiates a new KeyAmf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *KeyAmf) GetKeyType() KeyAmfType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *KeyAmf) GetKeyTypeOk() (*KeyAmfType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *KeyAmf) SetKeyType(v KeyAmfType)`

SetKeyType sets KeyType field to given value.


### GetKeyVal

`func (o *KeyAmf) GetKeyVal() string`

GetKeyVal returns the KeyVal field if non-nil, zero value otherwise.

### GetKeyValOk

`func (o *KeyAmf) GetKeyValOk() (*string, bool)`

GetKeyValOk returns a tuple with the KeyVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVal

`func (o *KeyAmf) SetKeyVal(v string)`

SetKeyVal sets KeyVal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


