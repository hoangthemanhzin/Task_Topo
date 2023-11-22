# AssignEbiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ProblemDetails**](ProblemDetails.md) |  | 
**FailureDetails** | [**AssignEbiFailed**](AssignEbiFailed.md) |  | 

## Methods

### NewAssignEbiError

`func NewAssignEbiError(error_ ProblemDetails, failureDetails AssignEbiFailed, ) *AssignEbiError`

NewAssignEbiError instantiates a new AssignEbiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignEbiErrorWithDefaults

`func NewAssignEbiErrorWithDefaults() *AssignEbiError`

NewAssignEbiErrorWithDefaults instantiates a new AssignEbiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AssignEbiError) GetError() ProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AssignEbiError) GetErrorOk() (*ProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AssignEbiError) SetError(v ProblemDetails)`

SetError sets Error field to given value.


### GetFailureDetails

`func (o *AssignEbiError) GetFailureDetails() AssignEbiFailed`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *AssignEbiError) GetFailureDetailsOk() (*AssignEbiFailed, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *AssignEbiError) SetFailureDetails(v AssignEbiFailed)`

SetFailureDetails sets FailureDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


