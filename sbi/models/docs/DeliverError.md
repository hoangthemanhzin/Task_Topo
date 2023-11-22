# DeliverError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Cause** | Pointer to **string** |  | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**AccessTokenError** | Pointer to [**AccessTokenErr**](AccessTokenErr.md) |  | [optional] 
**AccessTokenRequest** | Pointer to [**AccessTokenReq**](AccessTokenReq.md) |  | [optional] 
**NrfId** | Pointer to **string** |  | [optional] 
**MaxWaitingTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeliverError

`func NewDeliverError() *DeliverError`

NewDeliverError instantiates a new DeliverError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverErrorWithDefaults

`func NewDeliverErrorWithDefaults() *DeliverError`

NewDeliverErrorWithDefaults instantiates a new DeliverError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeliverError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeliverError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeliverError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeliverError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *DeliverError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliverError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliverError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliverError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *DeliverError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeliverError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeliverError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeliverError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *DeliverError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DeliverError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DeliverError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DeliverError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *DeliverError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *DeliverError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *DeliverError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *DeliverError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *DeliverError) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *DeliverError) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *DeliverError) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *DeliverError) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *DeliverError) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *DeliverError) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *DeliverError) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *DeliverError) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DeliverError) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DeliverError) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DeliverError) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DeliverError) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *DeliverError) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *DeliverError) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *DeliverError) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *DeliverError) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *DeliverError) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *DeliverError) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *DeliverError) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *DeliverError) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *DeliverError) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *DeliverError) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *DeliverError) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *DeliverError) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetMaxWaitingTime

`func (o *DeliverError) GetMaxWaitingTime() int32`

GetMaxWaitingTime returns the MaxWaitingTime field if non-nil, zero value otherwise.

### GetMaxWaitingTimeOk

`func (o *DeliverError) GetMaxWaitingTimeOk() (*int32, bool)`

GetMaxWaitingTimeOk returns a tuple with the MaxWaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitingTime

`func (o *DeliverError) SetMaxWaitingTime(v int32)`

SetMaxWaitingTime sets MaxWaitingTime field to given value.

### HasMaxWaitingTime

`func (o *DeliverError) HasMaxWaitingTime() bool`

HasMaxWaitingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


