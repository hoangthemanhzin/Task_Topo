# UeContextInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportVoPS** | Pointer to **bool** |  | [optional] 
**SupportVoPSn3gpp** | Pointer to **bool** |  | [optional] 
**LastActTime** | Pointer to **time.Time** |  | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewUeContextInfo

`func NewUeContextInfo() *UeContextInfo`

NewUeContextInfo instantiates a new UeContextInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextInfoWithDefaults

`func NewUeContextInfoWithDefaults() *UeContextInfo`

NewUeContextInfoWithDefaults instantiates a new UeContextInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportVoPS

`func (o *UeContextInfo) GetSupportVoPS() bool`

GetSupportVoPS returns the SupportVoPS field if non-nil, zero value otherwise.

### GetSupportVoPSOk

`func (o *UeContextInfo) GetSupportVoPSOk() (*bool, bool)`

GetSupportVoPSOk returns a tuple with the SupportVoPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVoPS

`func (o *UeContextInfo) SetSupportVoPS(v bool)`

SetSupportVoPS sets SupportVoPS field to given value.

### HasSupportVoPS

`func (o *UeContextInfo) HasSupportVoPS() bool`

HasSupportVoPS returns a boolean if a field has been set.

### GetSupportVoPSn3gpp

`func (o *UeContextInfo) GetSupportVoPSn3gpp() bool`

GetSupportVoPSn3gpp returns the SupportVoPSn3gpp field if non-nil, zero value otherwise.

### GetSupportVoPSn3gppOk

`func (o *UeContextInfo) GetSupportVoPSn3gppOk() (*bool, bool)`

GetSupportVoPSn3gppOk returns a tuple with the SupportVoPSn3gpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVoPSn3gpp

`func (o *UeContextInfo) SetSupportVoPSn3gpp(v bool)`

SetSupportVoPSn3gpp sets SupportVoPSn3gpp field to given value.

### HasSupportVoPSn3gpp

`func (o *UeContextInfo) HasSupportVoPSn3gpp() bool`

HasSupportVoPSn3gpp returns a boolean if a field has been set.

### GetLastActTime

`func (o *UeContextInfo) GetLastActTime() time.Time`

GetLastActTime returns the LastActTime field if non-nil, zero value otherwise.

### GetLastActTimeOk

`func (o *UeContextInfo) GetLastActTimeOk() (*time.Time, bool)`

GetLastActTimeOk returns a tuple with the LastActTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActTime

`func (o *UeContextInfo) SetLastActTime(v time.Time)`

SetLastActTime sets LastActTime field to given value.

### HasLastActTime

`func (o *UeContextInfo) HasLastActTime() bool`

HasLastActTime returns a boolean if a field has been set.

### GetAccessType

`func (o *UeContextInfo) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UeContextInfo) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UeContextInfo) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *UeContextInfo) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRatType

`func (o *UeContextInfo) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *UeContextInfo) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *UeContextInfo) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *UeContextInfo) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


