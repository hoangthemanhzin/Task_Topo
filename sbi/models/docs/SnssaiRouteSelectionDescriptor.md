# SnssaiRouteSelectionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**DnnRouteSelDescs** | Pointer to [**[]DnnRouteSelectionDescriptor**](DnnRouteSelectionDescriptor.md) |  | [optional] 

## Methods

### NewSnssaiRouteSelectionDescriptor

`func NewSnssaiRouteSelectionDescriptor(snssai Snssai, ) *SnssaiRouteSelectionDescriptor`

NewSnssaiRouteSelectionDescriptor instantiates a new SnssaiRouteSelectionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnssaiRouteSelectionDescriptorWithDefaults

`func NewSnssaiRouteSelectionDescriptorWithDefaults() *SnssaiRouteSelectionDescriptor`

NewSnssaiRouteSelectionDescriptorWithDefaults instantiates a new SnssaiRouteSelectionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *SnssaiRouteSelectionDescriptor) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SnssaiRouteSelectionDescriptor) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SnssaiRouteSelectionDescriptor) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnnRouteSelDescs

`func (o *SnssaiRouteSelectionDescriptor) GetDnnRouteSelDescs() []DnnRouteSelectionDescriptor`

GetDnnRouteSelDescs returns the DnnRouteSelDescs field if non-nil, zero value otherwise.

### GetDnnRouteSelDescsOk

`func (o *SnssaiRouteSelectionDescriptor) GetDnnRouteSelDescsOk() (*[]DnnRouteSelectionDescriptor, bool)`

GetDnnRouteSelDescsOk returns a tuple with the DnnRouteSelDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnRouteSelDescs

`func (o *SnssaiRouteSelectionDescriptor) SetDnnRouteSelDescs(v []DnnRouteSelectionDescriptor)`

SetDnnRouteSelDescs sets DnnRouteSelDescs field to given value.

### HasDnnRouteSelDescs

`func (o *SnssaiRouteSelectionDescriptor) HasDnnRouteSelDescs() bool`

HasDnnRouteSelDescs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


