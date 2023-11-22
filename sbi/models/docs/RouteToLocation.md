# RouteToLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnai** | **string** |  | 
**RouteInfo** | Pointer to [**RouteInformation**](RouteInformation.md) |  | [optional] 
**RouteProfId** | Pointer to **string** |  | [optional] 

## Methods

### NewRouteToLocation

`func NewRouteToLocation(dnai string, ) *RouteToLocation`

NewRouteToLocation instantiates a new RouteToLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteToLocationWithDefaults

`func NewRouteToLocationWithDefaults() *RouteToLocation`

NewRouteToLocationWithDefaults instantiates a new RouteToLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnai

`func (o *RouteToLocation) GetDnai() string`

GetDnai returns the Dnai field if non-nil, zero value otherwise.

### GetDnaiOk

`func (o *RouteToLocation) GetDnaiOk() (*string, bool)`

GetDnaiOk returns a tuple with the Dnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnai

`func (o *RouteToLocation) SetDnai(v string)`

SetDnai sets Dnai field to given value.


### GetRouteInfo

`func (o *RouteToLocation) GetRouteInfo() RouteInformation`

GetRouteInfo returns the RouteInfo field if non-nil, zero value otherwise.

### GetRouteInfoOk

`func (o *RouteToLocation) GetRouteInfoOk() (*RouteInformation, bool)`

GetRouteInfoOk returns a tuple with the RouteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteInfo

`func (o *RouteToLocation) SetRouteInfo(v RouteInformation)`

SetRouteInfo sets RouteInfo field to given value.

### HasRouteInfo

`func (o *RouteToLocation) HasRouteInfo() bool`

HasRouteInfo returns a boolean if a field has been set.

### SetRouteInfoNil

`func (o *RouteToLocation) SetRouteInfoNil(b bool)`

 SetRouteInfoNil sets the value for RouteInfo to be an explicit nil

### UnsetRouteInfo
`func (o *RouteToLocation) UnsetRouteInfo()`

UnsetRouteInfo ensures that no value is present for RouteInfo, not even an explicit nil
### GetRouteProfId

`func (o *RouteToLocation) GetRouteProfId() string`

GetRouteProfId returns the RouteProfId field if non-nil, zero value otherwise.

### GetRouteProfIdOk

`func (o *RouteToLocation) GetRouteProfIdOk() (*string, bool)`

GetRouteProfIdOk returns a tuple with the RouteProfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteProfId

`func (o *RouteToLocation) SetRouteProfId(v string)`

SetRouteProfId sets RouteProfId field to given value.

### HasRouteProfId

`func (o *RouteToLocation) HasRouteProfId() bool`

HasRouteProfId returns a boolean if a field has been set.

### SetRouteProfIdNil

`func (o *RouteToLocation) SetRouteProfIdNil(b bool)`

 SetRouteProfIdNil sets the value for RouteProfId to be an explicit nil

### UnsetRouteProfId
`func (o *RouteToLocation) UnsetRouteProfId()`

UnsetRouteProfId ensures that no value is present for RouteProfId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


