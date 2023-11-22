# ExtendedProblemDetails

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
**AcceptableServInfo** | Pointer to [**AcceptableServiceInfo**](AcceptableServiceInfo.md) |  | [optional] 

## Methods

### NewExtendedProblemDetails

`func NewExtendedProblemDetails() *ExtendedProblemDetails`

NewExtendedProblemDetails instantiates a new ExtendedProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedProblemDetailsWithDefaults

`func NewExtendedProblemDetailsWithDefaults() *ExtendedProblemDetails`

NewExtendedProblemDetailsWithDefaults instantiates a new ExtendedProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExtendedProblemDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedProblemDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedProblemDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExtendedProblemDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ExtendedProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExtendedProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExtendedProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExtendedProblemDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ExtendedProblemDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtendedProblemDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtendedProblemDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtendedProblemDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *ExtendedProblemDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ExtendedProblemDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ExtendedProblemDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ExtendedProblemDetails) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *ExtendedProblemDetails) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ExtendedProblemDetails) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ExtendedProblemDetails) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ExtendedProblemDetails) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *ExtendedProblemDetails) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ExtendedProblemDetails) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ExtendedProblemDetails) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ExtendedProblemDetails) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *ExtendedProblemDetails) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *ExtendedProblemDetails) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *ExtendedProblemDetails) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *ExtendedProblemDetails) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ExtendedProblemDetails) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ExtendedProblemDetails) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ExtendedProblemDetails) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ExtendedProblemDetails) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *ExtendedProblemDetails) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *ExtendedProblemDetails) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *ExtendedProblemDetails) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *ExtendedProblemDetails) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *ExtendedProblemDetails) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *ExtendedProblemDetails) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *ExtendedProblemDetails) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *ExtendedProblemDetails) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *ExtendedProblemDetails) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *ExtendedProblemDetails) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *ExtendedProblemDetails) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *ExtendedProblemDetails) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetAcceptableServInfo

`func (o *ExtendedProblemDetails) GetAcceptableServInfo() AcceptableServiceInfo`

GetAcceptableServInfo returns the AcceptableServInfo field if non-nil, zero value otherwise.

### GetAcceptableServInfoOk

`func (o *ExtendedProblemDetails) GetAcceptableServInfoOk() (*AcceptableServiceInfo, bool)`

GetAcceptableServInfoOk returns a tuple with the AcceptableServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableServInfo

`func (o *ExtendedProblemDetails) SetAcceptableServInfo(v AcceptableServiceInfo)`

SetAcceptableServInfo sets AcceptableServInfo field to given value.

### HasAcceptableServInfo

`func (o *ExtendedProblemDetails) HasAcceptableServInfo() bool`

HasAcceptableServInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


