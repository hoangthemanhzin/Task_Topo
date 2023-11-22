# ExpectedUeBehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpMoveTrajectory** | [**[]UserLocation**](UserLocation.md) |  | 
**ValidityTime** | **time.Time** |  | 

## Methods

### NewExpectedUeBehavior

`func NewExpectedUeBehavior(expMoveTrajectory []UserLocation, validityTime time.Time, ) *ExpectedUeBehavior`

NewExpectedUeBehavior instantiates a new ExpectedUeBehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedUeBehaviorWithDefaults

`func NewExpectedUeBehaviorWithDefaults() *ExpectedUeBehavior`

NewExpectedUeBehaviorWithDefaults instantiates a new ExpectedUeBehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpMoveTrajectory

`func (o *ExpectedUeBehavior) GetExpMoveTrajectory() []UserLocation`

GetExpMoveTrajectory returns the ExpMoveTrajectory field if non-nil, zero value otherwise.

### GetExpMoveTrajectoryOk

`func (o *ExpectedUeBehavior) GetExpMoveTrajectoryOk() (*[]UserLocation, bool)`

GetExpMoveTrajectoryOk returns a tuple with the ExpMoveTrajectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMoveTrajectory

`func (o *ExpectedUeBehavior) SetExpMoveTrajectory(v []UserLocation)`

SetExpMoveTrajectory sets ExpMoveTrajectory field to given value.


### GetValidityTime

`func (o *ExpectedUeBehavior) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *ExpectedUeBehavior) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *ExpectedUeBehavior) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


