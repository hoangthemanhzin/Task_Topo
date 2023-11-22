# UeId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** |  | 
**GpsiList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUeId

`func NewUeId(supi string, ) *UeId`

NewUeId instantiates a new UeId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeIdWithDefaults

`func NewUeIdWithDefaults() *UeId`

NewUeIdWithDefaults instantiates a new UeId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeId) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeId) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeId) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsiList

`func (o *UeId) GetGpsiList() []string`

GetGpsiList returns the GpsiList field if non-nil, zero value otherwise.

### GetGpsiListOk

`func (o *UeId) GetGpsiListOk() (*[]string, bool)`

GetGpsiListOk returns a tuple with the GpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiList

`func (o *UeId) SetGpsiList(v []string)`

SetGpsiList sets GpsiList field to given value.

### HasGpsiList

`func (o *UeId) HasGpsiList() bool`

HasGpsiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


