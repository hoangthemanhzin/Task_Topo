# SmscData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmscMapAddress** | Pointer to **string** |  | [optional] 
**SmscDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress1**](NetworkNodeDiameterAddress1.md) |  | [optional] 

## Methods

### NewSmscData

`func NewSmscData() *SmscData`

NewSmscData instantiates a new SmscData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmscDataWithDefaults

`func NewSmscDataWithDefaults() *SmscData`

NewSmscDataWithDefaults instantiates a new SmscData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmscMapAddress

`func (o *SmscData) GetSmscMapAddress() string`

GetSmscMapAddress returns the SmscMapAddress field if non-nil, zero value otherwise.

### GetSmscMapAddressOk

`func (o *SmscData) GetSmscMapAddressOk() (*string, bool)`

GetSmscMapAddressOk returns a tuple with the SmscMapAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscMapAddress

`func (o *SmscData) SetSmscMapAddress(v string)`

SetSmscMapAddress sets SmscMapAddress field to given value.

### HasSmscMapAddress

`func (o *SmscData) HasSmscMapAddress() bool`

HasSmscMapAddress returns a boolean if a field has been set.

### GetSmscDiameterAddress

`func (o *SmscData) GetSmscDiameterAddress() NetworkNodeDiameterAddress1`

GetSmscDiameterAddress returns the SmscDiameterAddress field if non-nil, zero value otherwise.

### GetSmscDiameterAddressOk

`func (o *SmscData) GetSmscDiameterAddressOk() (*NetworkNodeDiameterAddress1, bool)`

GetSmscDiameterAddressOk returns a tuple with the SmscDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscDiameterAddress

`func (o *SmscData) SetSmscDiameterAddress(v NetworkNodeDiameterAddress1)`

SetSmscDiameterAddress sets SmscDiameterAddress field to given value.

### HasSmscDiameterAddress

`func (o *SmscData) HasSmscDiameterAddress() bool`

HasSmscDiameterAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


