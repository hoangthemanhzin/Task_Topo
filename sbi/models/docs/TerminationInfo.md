# TerminationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermCause** | [**TerminationCause**](TerminationCause.md) |  | 
**ResUri** | **string** |  | 

## Methods

### NewTerminationInfo

`func NewTerminationInfo(termCause TerminationCause, resUri string, ) *TerminationInfo`

NewTerminationInfo instantiates a new TerminationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminationInfoWithDefaults

`func NewTerminationInfoWithDefaults() *TerminationInfo`

NewTerminationInfoWithDefaults instantiates a new TerminationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermCause

`func (o *TerminationInfo) GetTermCause() TerminationCause`

GetTermCause returns the TermCause field if non-nil, zero value otherwise.

### GetTermCauseOk

`func (o *TerminationInfo) GetTermCauseOk() (*TerminationCause, bool)`

GetTermCauseOk returns a tuple with the TermCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermCause

`func (o *TerminationInfo) SetTermCause(v TerminationCause)`

SetTermCause sets TermCause field to given value.


### GetResUri

`func (o *TerminationInfo) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *TerminationInfo) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *TerminationInfo) SetResUri(v string)`

SetResUri sets ResUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


