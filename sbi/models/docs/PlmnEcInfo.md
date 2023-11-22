# PlmnEcInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**EcRestrictionDataWb** | Pointer to [**EcRestrictionDataWb**](EcRestrictionDataWb.md) |  | [optional] 
**EcRestrictionDataNb** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPlmnEcInfo

`func NewPlmnEcInfo(plmnId PlmnId, ) *PlmnEcInfo`

NewPlmnEcInfo instantiates a new PlmnEcInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnEcInfoWithDefaults

`func NewPlmnEcInfoWithDefaults() *PlmnEcInfo`

NewPlmnEcInfoWithDefaults instantiates a new PlmnEcInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *PlmnEcInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PlmnEcInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PlmnEcInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetEcRestrictionDataWb

`func (o *PlmnEcInfo) GetEcRestrictionDataWb() EcRestrictionDataWb`

GetEcRestrictionDataWb returns the EcRestrictionDataWb field if non-nil, zero value otherwise.

### GetEcRestrictionDataWbOk

`func (o *PlmnEcInfo) GetEcRestrictionDataWbOk() (*EcRestrictionDataWb, bool)`

GetEcRestrictionDataWbOk returns a tuple with the EcRestrictionDataWb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataWb

`func (o *PlmnEcInfo) SetEcRestrictionDataWb(v EcRestrictionDataWb)`

SetEcRestrictionDataWb sets EcRestrictionDataWb field to given value.

### HasEcRestrictionDataWb

`func (o *PlmnEcInfo) HasEcRestrictionDataWb() bool`

HasEcRestrictionDataWb returns a boolean if a field has been set.

### GetEcRestrictionDataNb

`func (o *PlmnEcInfo) GetEcRestrictionDataNb() bool`

GetEcRestrictionDataNb returns the EcRestrictionDataNb field if non-nil, zero value otherwise.

### GetEcRestrictionDataNbOk

`func (o *PlmnEcInfo) GetEcRestrictionDataNbOk() (*bool, bool)`

GetEcRestrictionDataNbOk returns a tuple with the EcRestrictionDataNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataNb

`func (o *PlmnEcInfo) SetEcRestrictionDataNb(v bool)`

SetEcRestrictionDataNb sets EcRestrictionDataNb field to given value.

### HasEcRestrictionDataNb

`func (o *PlmnEcInfo) HasEcRestrictionDataNb() bool`

HasEcRestrictionDataNb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


