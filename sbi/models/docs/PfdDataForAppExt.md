# PfdDataForAppExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**Pfds** | [**[]PfdContent**](PfdContent.md) |  | 
**CachingTime** | Pointer to **time.Time** |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewPfdDataForAppExt

`func NewPfdDataForAppExt(applicationId string, pfds []PfdContent, ) *PfdDataForAppExt`

NewPfdDataForAppExt instantiates a new PfdDataForAppExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdDataForAppExtWithDefaults

`func NewPfdDataForAppExtWithDefaults() *PfdDataForAppExt`

NewPfdDataForAppExtWithDefaults instantiates a new PfdDataForAppExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PfdDataForAppExt) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PfdDataForAppExt) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PfdDataForAppExt) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetPfds

`func (o *PfdDataForAppExt) GetPfds() []PfdContent`

GetPfds returns the Pfds field if non-nil, zero value otherwise.

### GetPfdsOk

`func (o *PfdDataForAppExt) GetPfdsOk() (*[]PfdContent, bool)`

GetPfdsOk returns a tuple with the Pfds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfds

`func (o *PfdDataForAppExt) SetPfds(v []PfdContent)`

SetPfds sets Pfds field to given value.


### GetCachingTime

`func (o *PfdDataForAppExt) GetCachingTime() time.Time`

GetCachingTime returns the CachingTime field if non-nil, zero value otherwise.

### GetCachingTimeOk

`func (o *PfdDataForAppExt) GetCachingTimeOk() (*time.Time, bool)`

GetCachingTimeOk returns a tuple with the CachingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingTime

`func (o *PfdDataForAppExt) SetCachingTime(v time.Time)`

SetCachingTime sets CachingTime field to given value.

### HasCachingTime

`func (o *PfdDataForAppExt) HasCachingTime() bool`

HasCachingTime returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PfdDataForAppExt) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PfdDataForAppExt) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PfdDataForAppExt) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PfdDataForAppExt) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


