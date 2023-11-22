# NiddAuthUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationData** | [**AuthorizationData**](AuthorizationData.md) |  | 
**InvalidityInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewNiddAuthUpdateInfo

`func NewNiddAuthUpdateInfo(authorizationData AuthorizationData, ) *NiddAuthUpdateInfo`

NewNiddAuthUpdateInfo instantiates a new NiddAuthUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddAuthUpdateInfoWithDefaults

`func NewNiddAuthUpdateInfoWithDefaults() *NiddAuthUpdateInfo`

NewNiddAuthUpdateInfoWithDefaults instantiates a new NiddAuthUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationData

`func (o *NiddAuthUpdateInfo) GetAuthorizationData() AuthorizationData`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *NiddAuthUpdateInfo) GetAuthorizationDataOk() (*AuthorizationData, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *NiddAuthUpdateInfo) SetAuthorizationData(v AuthorizationData)`

SetAuthorizationData sets AuthorizationData field to given value.


### GetInvalidityInd

`func (o *NiddAuthUpdateInfo) GetInvalidityInd() bool`

GetInvalidityInd returns the InvalidityInd field if non-nil, zero value otherwise.

### GetInvalidityIndOk

`func (o *NiddAuthUpdateInfo) GetInvalidityIndOk() (*bool, bool)`

GetInvalidityIndOk returns a tuple with the InvalidityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidityInd

`func (o *NiddAuthUpdateInfo) SetInvalidityInd(v bool)`

SetInvalidityInd sets InvalidityInd field to given value.

### HasInvalidityInd

`func (o *NiddAuthUpdateInfo) HasInvalidityInd() bool`

HasInvalidityInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


