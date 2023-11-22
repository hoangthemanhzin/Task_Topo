# AdditionalAccessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AccessType**](AccessType.md) |  | 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 

## Methods

### NewAdditionalAccessInfo

`func NewAdditionalAccessInfo(accessType AccessType, ) *AdditionalAccessInfo`

NewAdditionalAccessInfo instantiates a new AdditionalAccessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalAccessInfoWithDefaults

`func NewAdditionalAccessInfoWithDefaults() *AdditionalAccessInfo`

NewAdditionalAccessInfoWithDefaults instantiates a new AdditionalAccessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *AdditionalAccessInfo) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AdditionalAccessInfo) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AdditionalAccessInfo) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetRatType

`func (o *AdditionalAccessInfo) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *AdditionalAccessInfo) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *AdditionalAccessInfo) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *AdditionalAccessInfo) HasRatType() bool`

HasRatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


