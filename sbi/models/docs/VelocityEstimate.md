# VelocityEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HSpeed** | **float32** |  | 
**Bearing** | **int32** |  | 
**VSpeed** | **float32** |  | 
**VDirection** | [**VerticalDirection**](VerticalDirection.md) |  | 
**HUncertainty** | **float32** |  | 
**VUncertainty** | **float32** |  | 

## Methods

### NewVelocityEstimate

`func NewVelocityEstimate(hSpeed float32, bearing int32, vSpeed float32, vDirection VerticalDirection, hUncertainty float32, vUncertainty float32, ) *VelocityEstimate`

NewVelocityEstimate instantiates a new VelocityEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityEstimateWithDefaults

`func NewVelocityEstimateWithDefaults() *VelocityEstimate`

NewVelocityEstimateWithDefaults instantiates a new VelocityEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHSpeed

`func (o *VelocityEstimate) GetHSpeed() float32`

GetHSpeed returns the HSpeed field if non-nil, zero value otherwise.

### GetHSpeedOk

`func (o *VelocityEstimate) GetHSpeedOk() (*float32, bool)`

GetHSpeedOk returns a tuple with the HSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSpeed

`func (o *VelocityEstimate) SetHSpeed(v float32)`

SetHSpeed sets HSpeed field to given value.


### GetBearing

`func (o *VelocityEstimate) GetBearing() int32`

GetBearing returns the Bearing field if non-nil, zero value otherwise.

### GetBearingOk

`func (o *VelocityEstimate) GetBearingOk() (*int32, bool)`

GetBearingOk returns a tuple with the Bearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearing

`func (o *VelocityEstimate) SetBearing(v int32)`

SetBearing sets Bearing field to given value.


### GetVSpeed

`func (o *VelocityEstimate) GetVSpeed() float32`

GetVSpeed returns the VSpeed field if non-nil, zero value otherwise.

### GetVSpeedOk

`func (o *VelocityEstimate) GetVSpeedOk() (*float32, bool)`

GetVSpeedOk returns a tuple with the VSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSpeed

`func (o *VelocityEstimate) SetVSpeed(v float32)`

SetVSpeed sets VSpeed field to given value.


### GetVDirection

`func (o *VelocityEstimate) GetVDirection() VerticalDirection`

GetVDirection returns the VDirection field if non-nil, zero value otherwise.

### GetVDirectionOk

`func (o *VelocityEstimate) GetVDirectionOk() (*VerticalDirection, bool)`

GetVDirectionOk returns a tuple with the VDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDirection

`func (o *VelocityEstimate) SetVDirection(v VerticalDirection)`

SetVDirection sets VDirection field to given value.


### GetHUncertainty

`func (o *VelocityEstimate) GetHUncertainty() float32`

GetHUncertainty returns the HUncertainty field if non-nil, zero value otherwise.

### GetHUncertaintyOk

`func (o *VelocityEstimate) GetHUncertaintyOk() (*float32, bool)`

GetHUncertaintyOk returns a tuple with the HUncertainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHUncertainty

`func (o *VelocityEstimate) SetHUncertainty(v float32)`

SetHUncertainty sets HUncertainty field to given value.


### GetVUncertainty

`func (o *VelocityEstimate) GetVUncertainty() float32`

GetVUncertainty returns the VUncertainty field if non-nil, zero value otherwise.

### GetVUncertaintyOk

`func (o *VelocityEstimate) GetVUncertaintyOk() (*float32, bool)`

GetVUncertaintyOk returns a tuple with the VUncertainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVUncertainty

`func (o *VelocityEstimate) SetVUncertainty(v float32)`

SetVUncertainty sets VUncertainty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


