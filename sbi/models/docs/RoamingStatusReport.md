# RoamingStatusReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roaming** | **bool** |  | 
**NewServingPlmn** | [**PlmnId**](PlmnId.md) |  | 

## Methods

### NewRoamingStatusReport

`func NewRoamingStatusReport(roaming bool, newServingPlmn PlmnId, ) *RoamingStatusReport`

NewRoamingStatusReport instantiates a new RoamingStatusReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingStatusReportWithDefaults

`func NewRoamingStatusReportWithDefaults() *RoamingStatusReport`

NewRoamingStatusReportWithDefaults instantiates a new RoamingStatusReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoaming

`func (o *RoamingStatusReport) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *RoamingStatusReport) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *RoamingStatusReport) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.


### GetNewServingPlmn

`func (o *RoamingStatusReport) GetNewServingPlmn() PlmnId`

GetNewServingPlmn returns the NewServingPlmn field if non-nil, zero value otherwise.

### GetNewServingPlmnOk

`func (o *RoamingStatusReport) GetNewServingPlmnOk() (*PlmnId, bool)`

GetNewServingPlmnOk returns a tuple with the NewServingPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewServingPlmn

`func (o *RoamingStatusReport) SetNewServingPlmn(v PlmnId)`

SetNewServingPlmn sets NewServingPlmn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


