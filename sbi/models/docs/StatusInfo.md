# StatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceStatus** | [**ResourceStatus**](ResourceStatus.md) |  | 
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**CnAssistedRanPara** | Pointer to [**CnAssistedRanPara**](CnAssistedRanPara.md) |  | [optional] 
**AnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewStatusInfo

`func NewStatusInfo(resourceStatus ResourceStatus, ) *StatusInfo`

NewStatusInfo instantiates a new StatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusInfoWithDefaults

`func NewStatusInfoWithDefaults() *StatusInfo`

NewStatusInfoWithDefaults instantiates a new StatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceStatus

`func (o *StatusInfo) GetResourceStatus() ResourceStatus`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *StatusInfo) GetResourceStatusOk() (*ResourceStatus, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *StatusInfo) SetResourceStatus(v ResourceStatus)`

SetResourceStatus sets ResourceStatus field to given value.


### GetCause

`func (o *StatusInfo) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *StatusInfo) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *StatusInfo) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *StatusInfo) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCnAssistedRanPara

`func (o *StatusInfo) GetCnAssistedRanPara() CnAssistedRanPara`

GetCnAssistedRanPara returns the CnAssistedRanPara field if non-nil, zero value otherwise.

### GetCnAssistedRanParaOk

`func (o *StatusInfo) GetCnAssistedRanParaOk() (*CnAssistedRanPara, bool)`

GetCnAssistedRanParaOk returns a tuple with the CnAssistedRanPara field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnAssistedRanPara

`func (o *StatusInfo) SetCnAssistedRanPara(v CnAssistedRanPara)`

SetCnAssistedRanPara sets CnAssistedRanPara field to given value.

### HasCnAssistedRanPara

`func (o *StatusInfo) HasCnAssistedRanPara() bool`

HasCnAssistedRanPara returns a boolean if a field has been set.

### GetAnType

`func (o *StatusInfo) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *StatusInfo) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *StatusInfo) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.

### HasAnType

`func (o *StatusInfo) HasAnType() bool`

HasAnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


