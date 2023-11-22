# ProblemDetails2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code for this occurrence of the problem. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available. | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam1**](InvalidParam1.md) | Description of invalid parameters, for a request rejected due to invalid parameters. | [optional] 

## Methods

### NewProblemDetails2

`func NewProblemDetails2() *ProblemDetails2`

NewProblemDetails2 instantiates a new ProblemDetails2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDetails2WithDefaults

`func NewProblemDetails2WithDefaults() *ProblemDetails2`

NewProblemDetails2WithDefaults instantiates a new ProblemDetails2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProblemDetails2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProblemDetails2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProblemDetails2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProblemDetails2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ProblemDetails2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemDetails2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemDetails2) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProblemDetails2) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ProblemDetails2) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProblemDetails2) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProblemDetails2) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProblemDetails2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *ProblemDetails2) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemDetails2) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemDetails2) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ProblemDetails2) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *ProblemDetails2) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProblemDetails2) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProblemDetails2) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ProblemDetails2) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *ProblemDetails2) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ProblemDetails2) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ProblemDetails2) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ProblemDetails2) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *ProblemDetails2) GetInvalidParams() []InvalidParam1`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *ProblemDetails2) GetInvalidParamsOk() (*[]InvalidParam1, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *ProblemDetails2) SetInvalidParams(v []InvalidParam1)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *ProblemDetails2) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


