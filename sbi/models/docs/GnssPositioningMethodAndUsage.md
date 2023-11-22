# GnssPositioningMethodAndUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**PositioningMode**](PositioningMode.md) |  | 
**Gnss** | [**GnssId**](GnssId.md) |  | 
**Usage** | [**Usage**](Usage.md) |  | 

## Methods

### NewGnssPositioningMethodAndUsage

`func NewGnssPositioningMethodAndUsage(mode PositioningMode, gnss GnssId, usage Usage, ) *GnssPositioningMethodAndUsage`

NewGnssPositioningMethodAndUsage instantiates a new GnssPositioningMethodAndUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGnssPositioningMethodAndUsageWithDefaults

`func NewGnssPositioningMethodAndUsageWithDefaults() *GnssPositioningMethodAndUsage`

NewGnssPositioningMethodAndUsageWithDefaults instantiates a new GnssPositioningMethodAndUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GnssPositioningMethodAndUsage) GetMode() PositioningMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GnssPositioningMethodAndUsage) GetModeOk() (*PositioningMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GnssPositioningMethodAndUsage) SetMode(v PositioningMode)`

SetMode sets Mode field to given value.


### GetGnss

`func (o *GnssPositioningMethodAndUsage) GetGnss() GnssId`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *GnssPositioningMethodAndUsage) GetGnssOk() (*GnssId, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *GnssPositioningMethodAndUsage) SetGnss(v GnssId)`

SetGnss sets Gnss field to given value.


### GetUsage

`func (o *GnssPositioningMethodAndUsage) GetUsage() Usage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GnssPositioningMethodAndUsage) GetUsageOk() (*Usage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GnssPositioningMethodAndUsage) SetUsage(v Usage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


