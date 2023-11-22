# CommunicationFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NasReleaseCode** | Pointer to **string** |  | [optional] 
**RanReleaseCode** | Pointer to [**NgApCause**](NgApCause.md) |  | [optional] 

## Methods

### NewCommunicationFailure

`func NewCommunicationFailure() *CommunicationFailure`

NewCommunicationFailure instantiates a new CommunicationFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationFailureWithDefaults

`func NewCommunicationFailureWithDefaults() *CommunicationFailure`

NewCommunicationFailureWithDefaults instantiates a new CommunicationFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNasReleaseCode

`func (o *CommunicationFailure) GetNasReleaseCode() string`

GetNasReleaseCode returns the NasReleaseCode field if non-nil, zero value otherwise.

### GetNasReleaseCodeOk

`func (o *CommunicationFailure) GetNasReleaseCodeOk() (*string, bool)`

GetNasReleaseCodeOk returns a tuple with the NasReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasReleaseCode

`func (o *CommunicationFailure) SetNasReleaseCode(v string)`

SetNasReleaseCode sets NasReleaseCode field to given value.

### HasNasReleaseCode

`func (o *CommunicationFailure) HasNasReleaseCode() bool`

HasNasReleaseCode returns a boolean if a field has been set.

### GetRanReleaseCode

`func (o *CommunicationFailure) GetRanReleaseCode() NgApCause`

GetRanReleaseCode returns the RanReleaseCode field if non-nil, zero value otherwise.

### GetRanReleaseCodeOk

`func (o *CommunicationFailure) GetRanReleaseCodeOk() (*NgApCause, bool)`

GetRanReleaseCodeOk returns a tuple with the RanReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanReleaseCode

`func (o *CommunicationFailure) SetRanReleaseCode(v NgApCause)`

SetRanReleaseCode sets RanReleaseCode field to given value.

### HasRanReleaseCode

`func (o *CommunicationFailure) HasRanReleaseCode() bool`

HasRanReleaseCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


