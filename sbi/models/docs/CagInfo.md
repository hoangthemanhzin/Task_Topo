# CagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCagList** | **[]string** |  | 
**CagOnlyIndicator** | Pointer to **bool** |  | [optional] 

## Methods

### NewCagInfo

`func NewCagInfo(allowedCagList []string, ) *CagInfo`

NewCagInfo instantiates a new CagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCagInfoWithDefaults

`func NewCagInfoWithDefaults() *CagInfo`

NewCagInfoWithDefaults instantiates a new CagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedCagList

`func (o *CagInfo) GetAllowedCagList() []string`

GetAllowedCagList returns the AllowedCagList field if non-nil, zero value otherwise.

### GetAllowedCagListOk

`func (o *CagInfo) GetAllowedCagListOk() (*[]string, bool)`

GetAllowedCagListOk returns a tuple with the AllowedCagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCagList

`func (o *CagInfo) SetAllowedCagList(v []string)`

SetAllowedCagList sets AllowedCagList field to given value.


### GetCagOnlyIndicator

`func (o *CagInfo) GetCagOnlyIndicator() bool`

GetCagOnlyIndicator returns the CagOnlyIndicator field if non-nil, zero value otherwise.

### GetCagOnlyIndicatorOk

`func (o *CagInfo) GetCagOnlyIndicatorOk() (*bool, bool)`

GetCagOnlyIndicatorOk returns a tuple with the CagOnlyIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagOnlyIndicator

`func (o *CagInfo) SetCagOnlyIndicator(v bool)`

SetCagOnlyIndicator sets CagOnlyIndicator field to given value.

### HasCagOnlyIndicator

`func (o *CagInfo) HasCagOnlyIndicator() bool`

HasCagOnlyIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


