# AfResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfStatus** | [**AfResultStatus**](AfResultStatus.md) |  | 
**TrafficRoute** | Pointer to [**RouteToLocation**](RouteToLocation.md) |  | [optional] 

## Methods

### NewAfResultInfo

`func NewAfResultInfo(afStatus AfResultStatus, ) *AfResultInfo`

NewAfResultInfo instantiates a new AfResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfResultInfoWithDefaults

`func NewAfResultInfoWithDefaults() *AfResultInfo`

NewAfResultInfoWithDefaults instantiates a new AfResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfStatus

`func (o *AfResultInfo) GetAfStatus() AfResultStatus`

GetAfStatus returns the AfStatus field if non-nil, zero value otherwise.

### GetAfStatusOk

`func (o *AfResultInfo) GetAfStatusOk() (*AfResultStatus, bool)`

GetAfStatusOk returns a tuple with the AfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfStatus

`func (o *AfResultInfo) SetAfStatus(v AfResultStatus)`

SetAfStatus sets AfStatus field to given value.


### GetTrafficRoute

`func (o *AfResultInfo) GetTrafficRoute() RouteToLocation`

GetTrafficRoute returns the TrafficRoute field if non-nil, zero value otherwise.

### GetTrafficRouteOk

`func (o *AfResultInfo) GetTrafficRouteOk() (*RouteToLocation, bool)`

GetTrafficRouteOk returns a tuple with the TrafficRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRoute

`func (o *AfResultInfo) SetTrafficRoute(v RouteToLocation)`

SetTrafficRoute sets TrafficRoute field to given value.

### HasTrafficRoute

`func (o *AfResultInfo) HasTrafficRoute() bool`

HasTrafficRoute returns a boolean if a field has been set.

### SetTrafficRouteNil

`func (o *AfResultInfo) SetTrafficRouteNil(b bool)`

 SetTrafficRouteNil sets the value for TrafficRoute to be an explicit nil

### UnsetTrafficRoute
`func (o *AfResultInfo) UnsetTrafficRoute()`

UnsetTrafficRoute ensures that no value is present for TrafficRoute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


