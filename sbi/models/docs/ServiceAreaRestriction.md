# ServiceAreaRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictionType** | Pointer to [**RestrictionType**](RestrictionType.md) |  | [optional] 
**Areas** | Pointer to [**[]Area**](Area.md) |  | [optional] 
**MaxNumOfTAs** | Pointer to **int32** |  | [optional] 
**MaxNumOfTAsForNotAllowedAreas** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceAreaRestriction

`func NewServiceAreaRestriction() *ServiceAreaRestriction`

NewServiceAreaRestriction instantiates a new ServiceAreaRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAreaRestrictionWithDefaults

`func NewServiceAreaRestrictionWithDefaults() *ServiceAreaRestriction`

NewServiceAreaRestrictionWithDefaults instantiates a new ServiceAreaRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictionType

`func (o *ServiceAreaRestriction) GetRestrictionType() RestrictionType`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *ServiceAreaRestriction) GetRestrictionTypeOk() (*RestrictionType, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *ServiceAreaRestriction) SetRestrictionType(v RestrictionType)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *ServiceAreaRestriction) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### GetAreas

`func (o *ServiceAreaRestriction) GetAreas() []Area`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *ServiceAreaRestriction) GetAreasOk() (*[]Area, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *ServiceAreaRestriction) SetAreas(v []Area)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *ServiceAreaRestriction) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetMaxNumOfTAs

`func (o *ServiceAreaRestriction) GetMaxNumOfTAs() int32`

GetMaxNumOfTAs returns the MaxNumOfTAs field if non-nil, zero value otherwise.

### GetMaxNumOfTAsOk

`func (o *ServiceAreaRestriction) GetMaxNumOfTAsOk() (*int32, bool)`

GetMaxNumOfTAsOk returns a tuple with the MaxNumOfTAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTAs

`func (o *ServiceAreaRestriction) SetMaxNumOfTAs(v int32)`

SetMaxNumOfTAs sets MaxNumOfTAs field to given value.

### HasMaxNumOfTAs

`func (o *ServiceAreaRestriction) HasMaxNumOfTAs() bool`

HasMaxNumOfTAs returns a boolean if a field has been set.

### GetMaxNumOfTAsForNotAllowedAreas

`func (o *ServiceAreaRestriction) GetMaxNumOfTAsForNotAllowedAreas() int32`

GetMaxNumOfTAsForNotAllowedAreas returns the MaxNumOfTAsForNotAllowedAreas field if non-nil, zero value otherwise.

### GetMaxNumOfTAsForNotAllowedAreasOk

`func (o *ServiceAreaRestriction) GetMaxNumOfTAsForNotAllowedAreasOk() (*int32, bool)`

GetMaxNumOfTAsForNotAllowedAreasOk returns a tuple with the MaxNumOfTAsForNotAllowedAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfTAsForNotAllowedAreas

`func (o *ServiceAreaRestriction) SetMaxNumOfTAsForNotAllowedAreas(v int32)`

SetMaxNumOfTAsForNotAllowedAreas sets MaxNumOfTAsForNotAllowedAreas field to given value.

### HasMaxNumOfTAsForNotAllowedAreas

`func (o *ServiceAreaRestriction) HasMaxNumOfTAsForNotAllowedAreas() bool`

HasMaxNumOfTAsForNotAllowedAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


