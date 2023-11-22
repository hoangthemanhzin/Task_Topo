# PduSessionTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSessionType** | [**PduSessionType**](PduSessionType.md) |  | 
**AllowedSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewPduSessionTypes

`func NewPduSessionTypes(defaultSessionType PduSessionType, ) *PduSessionTypes`

NewPduSessionTypes instantiates a new PduSessionTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionTypesWithDefaults

`func NewPduSessionTypesWithDefaults() *PduSessionTypes`

NewPduSessionTypesWithDefaults instantiates a new PduSessionTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSessionType

`func (o *PduSessionTypes) GetDefaultSessionType() PduSessionType`

GetDefaultSessionType returns the DefaultSessionType field if non-nil, zero value otherwise.

### GetDefaultSessionTypeOk

`func (o *PduSessionTypes) GetDefaultSessionTypeOk() (*PduSessionType, bool)`

GetDefaultSessionTypeOk returns a tuple with the DefaultSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSessionType

`func (o *PduSessionTypes) SetDefaultSessionType(v PduSessionType)`

SetDefaultSessionType sets DefaultSessionType field to given value.


### GetAllowedSessionTypes

`func (o *PduSessionTypes) GetAllowedSessionTypes() []PduSessionType`

GetAllowedSessionTypes returns the AllowedSessionTypes field if non-nil, zero value otherwise.

### GetAllowedSessionTypesOk

`func (o *PduSessionTypes) GetAllowedSessionTypesOk() (*[]PduSessionType, bool)`

GetAllowedSessionTypesOk returns a tuple with the AllowedSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSessionTypes

`func (o *PduSessionTypes) SetAllowedSessionTypes(v []PduSessionType)`

SetAllowedSessionTypes sets AllowedSessionTypes field to given value.

### HasAllowedSessionTypes

`func (o *PduSessionTypes) HasAllowedSessionTypes() bool`

HasAllowedSessionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


