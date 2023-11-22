# SteeringInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId1**](PlmnId1.md) |  | 
**AccessTechList** | Pointer to [**[]AccessTech**](AccessTech.md) |  | [optional] 

## Methods

### NewSteeringInfo

`func NewSteeringInfo(plmnId PlmnId1, ) *SteeringInfo`

NewSteeringInfo instantiates a new SteeringInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteeringInfoWithDefaults

`func NewSteeringInfoWithDefaults() *SteeringInfo`

NewSteeringInfoWithDefaults instantiates a new SteeringInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *SteeringInfo) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SteeringInfo) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SteeringInfo) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.


### GetAccessTechList

`func (o *SteeringInfo) GetAccessTechList() []AccessTech`

GetAccessTechList returns the AccessTechList field if non-nil, zero value otherwise.

### GetAccessTechListOk

`func (o *SteeringInfo) GetAccessTechListOk() (*[]AccessTech, bool)`

GetAccessTechListOk returns a tuple with the AccessTechList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTechList

`func (o *SteeringInfo) SetAccessTechList(v []AccessTech)`

SetAccessTechList sets AccessTechList field to given value.

### HasAccessTechList

`func (o *SteeringInfo) HasAccessTechList() bool`

HasAccessTechList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


