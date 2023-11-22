# PositioningMethodAndUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**PositioningMethod**](PositioningMethod.md) |  | 
**Mode** | [**PositioningMode**](PositioningMode.md) |  | 
**Usage** | [**Usage**](Usage.md) |  | 
**MethodCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewPositioningMethodAndUsage

`func NewPositioningMethodAndUsage(method PositioningMethod, mode PositioningMode, usage Usage, ) *PositioningMethodAndUsage`

NewPositioningMethodAndUsage instantiates a new PositioningMethodAndUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositioningMethodAndUsageWithDefaults

`func NewPositioningMethodAndUsageWithDefaults() *PositioningMethodAndUsage`

NewPositioningMethodAndUsageWithDefaults instantiates a new PositioningMethodAndUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *PositioningMethodAndUsage) GetMethod() PositioningMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PositioningMethodAndUsage) GetMethodOk() (*PositioningMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PositioningMethodAndUsage) SetMethod(v PositioningMethod)`

SetMethod sets Method field to given value.


### GetMode

`func (o *PositioningMethodAndUsage) GetMode() PositioningMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PositioningMethodAndUsage) GetModeOk() (*PositioningMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PositioningMethodAndUsage) SetMode(v PositioningMode)`

SetMode sets Mode field to given value.


### GetUsage

`func (o *PositioningMethodAndUsage) GetUsage() Usage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PositioningMethodAndUsage) GetUsageOk() (*Usage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PositioningMethodAndUsage) SetUsage(v Usage)`

SetUsage sets Usage field to given value.


### GetMethodCode

`func (o *PositioningMethodAndUsage) GetMethodCode() int32`

GetMethodCode returns the MethodCode field if non-nil, zero value otherwise.

### GetMethodCodeOk

`func (o *PositioningMethodAndUsage) GetMethodCodeOk() (*int32, bool)`

GetMethodCodeOk returns a tuple with the MethodCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodCode

`func (o *PositioningMethodAndUsage) SetMethodCode(v int32)`

SetMethodCode sets MethodCode field to given value.

### HasMethodCode

`func (o *PositioningMethodAndUsage) HasMethodCode() bool`

HasMethodCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


