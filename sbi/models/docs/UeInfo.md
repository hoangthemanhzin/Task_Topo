# UeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TadsInfo** | Pointer to [**UeContextInfo**](UeContextInfo.md) |  | [optional] 
**UserState** | Pointer to [**Model5GsUserState**](Model5GsUserState.md) |  | [optional] 
**Var5gSrvccInfo** | Pointer to [**Model5GSrvccInfo**](Model5GSrvccInfo.md) |  | [optional] 

## Methods

### NewUeInfo

`func NewUeInfo() *UeInfo`

NewUeInfo instantiates a new UeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeInfoWithDefaults

`func NewUeInfoWithDefaults() *UeInfo`

NewUeInfoWithDefaults instantiates a new UeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTadsInfo

`func (o *UeInfo) GetTadsInfo() UeContextInfo`

GetTadsInfo returns the TadsInfo field if non-nil, zero value otherwise.

### GetTadsInfoOk

`func (o *UeInfo) GetTadsInfoOk() (*UeContextInfo, bool)`

GetTadsInfoOk returns a tuple with the TadsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTadsInfo

`func (o *UeInfo) SetTadsInfo(v UeContextInfo)`

SetTadsInfo sets TadsInfo field to given value.

### HasTadsInfo

`func (o *UeInfo) HasTadsInfo() bool`

HasTadsInfo returns a boolean if a field has been set.

### GetUserState

`func (o *UeInfo) GetUserState() Model5GsUserState`

GetUserState returns the UserState field if non-nil, zero value otherwise.

### GetUserStateOk

`func (o *UeInfo) GetUserStateOk() (*Model5GsUserState, bool)`

GetUserStateOk returns a tuple with the UserState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserState

`func (o *UeInfo) SetUserState(v Model5GsUserState)`

SetUserState sets UserState field to given value.

### HasUserState

`func (o *UeInfo) HasUserState() bool`

HasUserState returns a boolean if a field has been set.

### GetVar5gSrvccInfo

`func (o *UeInfo) GetVar5gSrvccInfo() Model5GSrvccInfo`

GetVar5gSrvccInfo returns the Var5gSrvccInfo field if non-nil, zero value otherwise.

### GetVar5gSrvccInfoOk

`func (o *UeInfo) GetVar5gSrvccInfoOk() (*Model5GSrvccInfo, bool)`

GetVar5gSrvccInfoOk returns a tuple with the Var5gSrvccInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gSrvccInfo

`func (o *UeInfo) SetVar5gSrvccInfo(v Model5GSrvccInfo)`

SetVar5gSrvccInfo sets Var5gSrvccInfo field to given value.

### HasVar5gSrvccInfo

`func (o *UeInfo) HasVar5gSrvccInfo() bool`

HasVar5gSrvccInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


