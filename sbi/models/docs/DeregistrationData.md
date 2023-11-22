# DeregistrationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeregReason** | [**DeregistrationReason**](DeregistrationReason.md) |  | 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**PduSessionId** | Pointer to **int32** |  | [optional] 
**NewSmfInstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewDeregistrationData

`func NewDeregistrationData(deregReason DeregistrationReason, ) *DeregistrationData`

NewDeregistrationData instantiates a new DeregistrationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregistrationDataWithDefaults

`func NewDeregistrationDataWithDefaults() *DeregistrationData`

NewDeregistrationDataWithDefaults instantiates a new DeregistrationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeregReason

`func (o *DeregistrationData) GetDeregReason() DeregistrationReason`

GetDeregReason returns the DeregReason field if non-nil, zero value otherwise.

### GetDeregReasonOk

`func (o *DeregistrationData) GetDeregReasonOk() (*DeregistrationReason, bool)`

GetDeregReasonOk returns a tuple with the DeregReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregReason

`func (o *DeregistrationData) SetDeregReason(v DeregistrationReason)`

SetDeregReason sets DeregReason field to given value.


### GetAccessType

`func (o *DeregistrationData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *DeregistrationData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *DeregistrationData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *DeregistrationData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPduSessionId

`func (o *DeregistrationData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *DeregistrationData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *DeregistrationData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *DeregistrationData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetNewSmfInstanceId

`func (o *DeregistrationData) GetNewSmfInstanceId() string`

GetNewSmfInstanceId returns the NewSmfInstanceId field if non-nil, zero value otherwise.

### GetNewSmfInstanceIdOk

`func (o *DeregistrationData) GetNewSmfInstanceIdOk() (*string, bool)`

GetNewSmfInstanceIdOk returns a tuple with the NewSmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfInstanceId

`func (o *DeregistrationData) SetNewSmfInstanceId(v string)`

SetNewSmfInstanceId sets NewSmfInstanceId field to given value.

### HasNewSmfInstanceId

`func (o *DeregistrationData) HasNewSmfInstanceId() bool`

HasNewSmfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


