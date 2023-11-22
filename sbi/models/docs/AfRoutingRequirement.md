# AfRoutingRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppReloc** | Pointer to **bool** |  | [optional] 
**RouteToLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) |  | [optional] 
**SpVal** | Pointer to [**SpatialValidity**](SpatialValidity.md) |  | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**UpPathChgSub** | Pointer to [**UpPathChgEvent**](UpPathChgEvent.md) |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewAfRoutingRequirement

`func NewAfRoutingRequirement() *AfRoutingRequirement`

NewAfRoutingRequirement instantiates a new AfRoutingRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfRoutingRequirementWithDefaults

`func NewAfRoutingRequirementWithDefaults() *AfRoutingRequirement`

NewAfRoutingRequirementWithDefaults instantiates a new AfRoutingRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppReloc

`func (o *AfRoutingRequirement) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *AfRoutingRequirement) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *AfRoutingRequirement) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *AfRoutingRequirement) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRouteToLocs

`func (o *AfRoutingRequirement) GetRouteToLocs() []RouteToLocation`

GetRouteToLocs returns the RouteToLocs field if non-nil, zero value otherwise.

### GetRouteToLocsOk

`func (o *AfRoutingRequirement) GetRouteToLocsOk() (*[]RouteToLocation, bool)`

GetRouteToLocsOk returns a tuple with the RouteToLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToLocs

`func (o *AfRoutingRequirement) SetRouteToLocs(v []RouteToLocation)`

SetRouteToLocs sets RouteToLocs field to given value.

### HasRouteToLocs

`func (o *AfRoutingRequirement) HasRouteToLocs() bool`

HasRouteToLocs returns a boolean if a field has been set.

### GetSpVal

`func (o *AfRoutingRequirement) GetSpVal() SpatialValidity`

GetSpVal returns the SpVal field if non-nil, zero value otherwise.

### GetSpValOk

`func (o *AfRoutingRequirement) GetSpValOk() (*SpatialValidity, bool)`

GetSpValOk returns a tuple with the SpVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVal

`func (o *AfRoutingRequirement) SetSpVal(v SpatialValidity)`

SetSpVal sets SpVal field to given value.

### HasSpVal

`func (o *AfRoutingRequirement) HasSpVal() bool`

HasSpVal returns a boolean if a field has been set.

### GetTempVals

`func (o *AfRoutingRequirement) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *AfRoutingRequirement) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *AfRoutingRequirement) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *AfRoutingRequirement) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.

### GetUpPathChgSub

`func (o *AfRoutingRequirement) GetUpPathChgSub() UpPathChgEvent`

GetUpPathChgSub returns the UpPathChgSub field if non-nil, zero value otherwise.

### GetUpPathChgSubOk

`func (o *AfRoutingRequirement) GetUpPathChgSubOk() (*UpPathChgEvent, bool)`

GetUpPathChgSubOk returns a tuple with the UpPathChgSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgSub

`func (o *AfRoutingRequirement) SetUpPathChgSub(v UpPathChgEvent)`

SetUpPathChgSub sets UpPathChgSub field to given value.

### HasUpPathChgSub

`func (o *AfRoutingRequirement) HasUpPathChgSub() bool`

HasUpPathChgSub returns a boolean if a field has been set.

### SetUpPathChgSubNil

`func (o *AfRoutingRequirement) SetUpPathChgSubNil(b bool)`

 SetUpPathChgSubNil sets the value for UpPathChgSub to be an explicit nil

### UnsetUpPathChgSub
`func (o *AfRoutingRequirement) UnsetUpPathChgSub()`

UnsetUpPathChgSub ensures that no value is present for UpPathChgSub, not even an explicit nil
### GetAddrPreserInd

`func (o *AfRoutingRequirement) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *AfRoutingRequirement) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *AfRoutingRequirement) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *AfRoutingRequirement) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


