# UeContextCreateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**NgapCause** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 
**TargetToSourceFailureData** | Pointer to [**N2InfoContent**](N2InfoContent.md) |  | [optional] 

## Methods

### NewUeContextCreateError

`func NewUeContextCreateError(error_ ProblemDetails, ) *UeContextCreateError`

NewUeContextCreateError instantiates a new UeContextCreateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextCreateErrorWithDefaults

`func NewUeContextCreateErrorWithDefaults() *UeContextCreateError`

NewUeContextCreateErrorWithDefaults instantiates a new UeContextCreateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UeContextCreateError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UeContextCreateError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UeContextCreateError) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetNgapCause

`func (o *UeContextCreateError) GetNgapCause() NgApCause`

GetNgapCause returns the NgapCause field if non-nil, zero value otherwise.

### GetNgapCauseOk

`func (o *UeContextCreateError) GetNgapCauseOk() (*NgApCause, bool)`

GetNgapCauseOk returns a tuple with the NgapCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgapCause

`func (o *UeContextCreateError) SetNgapCause(v NgApCause)`

SetNgapCause sets NgapCause field to given value.

### HasNgapCause

`func (o *UeContextCreateError) HasNgapCause() bool`

HasNgapCause returns a boolean if a field has been set.

### GetTargetToSourceFailureData

`func (o *UeContextCreateError) GetTargetToSourceFailureData() N2InfoContent`

GetTargetToSourceFailureData returns the TargetToSourceFailureData field if non-nil, zero value otherwise.

### GetTargetToSourceFailureDataOk

`func (o *UeContextCreateError) GetTargetToSourceFailureDataOk() (*N2InfoContent, bool)`

GetTargetToSourceFailureDataOk returns a tuple with the TargetToSourceFailureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetToSourceFailureData

`func (o *UeContextCreateError) SetTargetToSourceFailureData(v N2InfoContent)`

SetTargetToSourceFailureData sets TargetToSourceFailureData field to given value.

### HasTargetToSourceFailureData

`func (o *UeContextCreateError) HasTargetToSourceFailureData() bool`

HasTargetToSourceFailureData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


