# AppSessionContextRespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAuthInfo** | Pointer to [**ServAuthInfo**](ServAuthInfo.md) |  | [optional] 
**UeIds** | Pointer to [**[]UeIdentityInfo**](UeIdentityInfo.md) |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSessionContextRespData

`func NewAppSessionContextRespData() *AppSessionContextRespData`

NewAppSessionContextRespData instantiates a new AppSessionContextRespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSessionContextRespDataWithDefaults

`func NewAppSessionContextRespDataWithDefaults() *AppSessionContextRespData`

NewAppSessionContextRespDataWithDefaults instantiates a new AppSessionContextRespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAuthInfo

`func (o *AppSessionContextRespData) GetServAuthInfo() ServAuthInfo`

GetServAuthInfo returns the ServAuthInfo field if non-nil, zero value otherwise.

### GetServAuthInfoOk

`func (o *AppSessionContextRespData) GetServAuthInfoOk() (*ServAuthInfo, bool)`

GetServAuthInfoOk returns a tuple with the ServAuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAuthInfo

`func (o *AppSessionContextRespData) SetServAuthInfo(v ServAuthInfo)`

SetServAuthInfo sets ServAuthInfo field to given value.

### HasServAuthInfo

`func (o *AppSessionContextRespData) HasServAuthInfo() bool`

HasServAuthInfo returns a boolean if a field has been set.

### GetUeIds

`func (o *AppSessionContextRespData) GetUeIds() []UeIdentityInfo`

GetUeIds returns the UeIds field if non-nil, zero value otherwise.

### GetUeIdsOk

`func (o *AppSessionContextRespData) GetUeIdsOk() (*[]UeIdentityInfo, bool)`

GetUeIdsOk returns a tuple with the UeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIds

`func (o *AppSessionContextRespData) SetUeIds(v []UeIdentityInfo)`

SetUeIds sets UeIds field to given value.

### HasUeIds

`func (o *AppSessionContextRespData) HasUeIds() bool`

HasUeIds returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AppSessionContextRespData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AppSessionContextRespData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AppSessionContextRespData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AppSessionContextRespData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


