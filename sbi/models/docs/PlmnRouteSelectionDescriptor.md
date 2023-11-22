# PlmnRouteSelectionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServingPlmn** | [**PlmnId1**](PlmnId1.md) |  | 
**SnssaiRouteSelDescs** | Pointer to [**[]SnssaiRouteSelectionDescriptor**](SnssaiRouteSelectionDescriptor.md) |  | [optional] 

## Methods

### NewPlmnRouteSelectionDescriptor

`func NewPlmnRouteSelectionDescriptor(servingPlmn PlmnId1, ) *PlmnRouteSelectionDescriptor`

NewPlmnRouteSelectionDescriptor instantiates a new PlmnRouteSelectionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnRouteSelectionDescriptorWithDefaults

`func NewPlmnRouteSelectionDescriptorWithDefaults() *PlmnRouteSelectionDescriptor`

NewPlmnRouteSelectionDescriptorWithDefaults instantiates a new PlmnRouteSelectionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServingPlmn

`func (o *PlmnRouteSelectionDescriptor) GetServingPlmn() PlmnId1`

GetServingPlmn returns the ServingPlmn field if non-nil, zero value otherwise.

### GetServingPlmnOk

`func (o *PlmnRouteSelectionDescriptor) GetServingPlmnOk() (*PlmnId1, bool)`

GetServingPlmnOk returns a tuple with the ServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPlmn

`func (o *PlmnRouteSelectionDescriptor) SetServingPlmn(v PlmnId1)`

SetServingPlmn sets ServingPlmn field to given value.


### GetSnssaiRouteSelDescs

`func (o *PlmnRouteSelectionDescriptor) GetSnssaiRouteSelDescs() []SnssaiRouteSelectionDescriptor`

GetSnssaiRouteSelDescs returns the SnssaiRouteSelDescs field if non-nil, zero value otherwise.

### GetSnssaiRouteSelDescsOk

`func (o *PlmnRouteSelectionDescriptor) GetSnssaiRouteSelDescsOk() (*[]SnssaiRouteSelectionDescriptor, bool)`

GetSnssaiRouteSelDescsOk returns a tuple with the SnssaiRouteSelDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiRouteSelDescs

`func (o *PlmnRouteSelectionDescriptor) SetSnssaiRouteSelDescs(v []SnssaiRouteSelectionDescriptor)`

SetSnssaiRouteSelDescs sets SnssaiRouteSelDescs field to given value.

### HasSnssaiRouteSelDescs

`func (o *PlmnRouteSelectionDescriptor) HasSnssaiRouteSelDescs() bool`

HasSnssaiRouteSelDescs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


