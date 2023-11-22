# N2InformationTransferError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**PwsErrorInfo** | Pointer to [**PWSErrorData**](PWSErrorData.md) |  | [optional] 

## Methods

### NewN2InformationTransferError

`func NewN2InformationTransferError(error_ ProblemDetails, ) *N2InformationTransferError`

NewN2InformationTransferError instantiates a new N2InformationTransferError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InformationTransferErrorWithDefaults

`func NewN2InformationTransferErrorWithDefaults() *N2InformationTransferError`

NewN2InformationTransferErrorWithDefaults instantiates a new N2InformationTransferError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *N2InformationTransferError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *N2InformationTransferError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *N2InformationTransferError) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetPwsErrorInfo

`func (o *N2InformationTransferError) GetPwsErrorInfo() PWSErrorData`

GetPwsErrorInfo returns the PwsErrorInfo field if non-nil, zero value otherwise.

### GetPwsErrorInfoOk

`func (o *N2InformationTransferError) GetPwsErrorInfoOk() (*PWSErrorData, bool)`

GetPwsErrorInfoOk returns a tuple with the PwsErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwsErrorInfo

`func (o *N2InformationTransferError) SetPwsErrorInfo(v PWSErrorData)`

SetPwsErrorInfo sets PwsErrorInfo field to given value.

### HasPwsErrorInfo

`func (o *N2InformationTransferError) HasPwsErrorInfo() bool`

HasPwsErrorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


