# SmContextCreateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ExtProblemDetails**](ExtProblemDetails.md) |  | 
**N1SmMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSmContextCreateError

`func NewSmContextCreateError(error_ ExtProblemDetails, ) *SmContextCreateError`

NewSmContextCreateError instantiates a new SmContextCreateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreateErrorWithDefaults

`func NewSmContextCreateErrorWithDefaults() *SmContextCreateError`

NewSmContextCreateErrorWithDefaults instantiates a new SmContextCreateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SmContextCreateError) GetError() ExtProblemDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SmContextCreateError) GetErrorOk() (*ExtProblemDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SmContextCreateError) SetError(v ExtProblemDetails)`

SetError sets Error field to given value.


### GetN1SmMsg

`func (o *SmContextCreateError) GetN1SmMsg() RefToBinaryData`

GetN1SmMsg returns the N1SmMsg field if non-nil, zero value otherwise.

### GetN1SmMsgOk

`func (o *SmContextCreateError) GetN1SmMsgOk() (*RefToBinaryData, bool)`

GetN1SmMsgOk returns a tuple with the N1SmMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmMsg

`func (o *SmContextCreateError) SetN1SmMsg(v RefToBinaryData)`

SetN1SmMsg sets N1SmMsg field to given value.

### HasN1SmMsg

`func (o *SmContextCreateError) HasN1SmMsg() bool`

HasN1SmMsg returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextCreateError) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextCreateError) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextCreateError) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextCreateError) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextCreateError) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextCreateError) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextCreateError) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextCreateError) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *SmContextCreateError) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *SmContextCreateError) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *SmContextCreateError) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *SmContextCreateError) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


