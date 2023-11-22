# AfRoutingRequirementRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppReloc** | Pointer to **bool** |  | [optional] 
**RouteToLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) |  | [optional] 
**SpVal** | Pointer to [**SpatialValidityRm**](SpatialValidityRm.md) |  | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**UpPathChgSub** | Pointer to [**UpPathChgEvent**](UpPathChgEvent.md) |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewAfRoutingRequirementRm

`func NewAfRoutingRequirementRm() *AfRoutingRequirementRm`

NewAfRoutingRequirementRm instantiates a new AfRoutingRequirementRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfRoutingRequirementRmWithDefaults

`func NewAfRoutingRequirementRmWithDefaults() *AfRoutingRequirementRm`

NewAfRoutingRequirementRmWithDefaults instantiates a new AfRoutingRequirementRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppReloc

`func (o *AfRoutingRequirementRm) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *AfRoutingRequirementRm) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *AfRoutingRequirementRm) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *AfRoutingRequirementRm) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRouteToLocs

`func (o *AfRoutingRequirementRm) GetRouteToLocs() []RouteToLocation`

GetRouteToLocs returns the RouteToLocs field if non-nil, zero value otherwise.

### GetRouteToLocsOk

`func (o *AfRoutingRequirementRm) GetRouteToLocsOk() (*[]RouteToLocation, bool)`

GetRouteToLocsOk returns a tuple with the RouteToLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToLocs

`func (o *AfRoutingRequirementRm) SetRouteToLocs(v []RouteToLocation)`

SetRouteToLocs sets RouteToLocs field to given value.

### HasRouteToLocs

`func (o *AfRoutingRequirementRm) HasRouteToLocs() bool`

HasRouteToLocs returns a boolean if a field has been set.

### SetRouteToLocsNil

`func (o *AfRoutingRequirementRm) SetRouteToLocsNil(b bool)`

 SetRouteToLocsNil sets the value for RouteToLocs to be an explicit nil

### UnsetRouteToLocs
`func (o *AfRoutingRequirementRm) UnsetRouteToLocs()`

UnsetRouteToLocs ensures that no value is present for RouteToLocs, not even an explicit nil
### GetSpVal

`func (o *AfRoutingRequirementRm) GetSpVal() SpatialValidityRm`

GetSpVal returns the SpVal field if non-nil, zero value otherwise.

### GetSpValOk

`func (o *AfRoutingRequirementRm) GetSpValOk() (*SpatialValidityRm, bool)`

GetSpValOk returns a tuple with the SpVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVal

`func (o *AfRoutingRequirementRm) SetSpVal(v SpatialValidityRm)`

SetSpVal sets SpVal field to given value.

### HasSpVal

`func (o *AfRoutingRequirementRm) HasSpVal() bool`

HasSpVal returns a boolean if a field has been set.

### SetSpValNil

`func (o *AfRoutingRequirementRm) SetSpValNil(b bool)`

 SetSpValNil sets the value for SpVal to be an explicit nil

### UnsetSpVal
`func (o *AfRoutingRequirementRm) UnsetSpVal()`

UnsetSpVal ensures that no value is present for SpVal, not even an explicit nil
### GetTempVals

`func (o *AfRoutingRequirementRm) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *AfRoutingRequirementRm) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *AfRoutingRequirementRm) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *AfRoutingRequirementRm) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.

### SetTempValsNil

`func (o *AfRoutingRequirementRm) SetTempValsNil(b bool)`

 SetTempValsNil sets the value for TempVals to be an explicit nil

### UnsetTempVals
`func (o *AfRoutingRequirementRm) UnsetTempVals()`

UnsetTempVals ensures that no value is present for TempVals, not even an explicit nil
### GetUpPathChgSub

`func (o *AfRoutingRequirementRm) GetUpPathChgSub() UpPathChgEvent`

GetUpPathChgSub returns the UpPathChgSub field if non-nil, zero value otherwise.

### GetUpPathChgSubOk

`func (o *AfRoutingRequirementRm) GetUpPathChgSubOk() (*UpPathChgEvent, bool)`

GetUpPathChgSubOk returns a tuple with the UpPathChgSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgSub

`func (o *AfRoutingRequirementRm) SetUpPathChgSub(v UpPathChgEvent)`

SetUpPathChgSub sets UpPathChgSub field to given value.

### HasUpPathChgSub

`func (o *AfRoutingRequirementRm) HasUpPathChgSub() bool`

HasUpPathChgSub returns a boolean if a field has been set.

### SetUpPathChgSubNil

`func (o *AfRoutingRequirementRm) SetUpPathChgSubNil(b bool)`

 SetUpPathChgSubNil sets the value for UpPathChgSub to be an explicit nil

### UnsetUpPathChgSub
`func (o *AfRoutingRequirementRm) UnsetUpPathChgSub()`

UnsetUpPathChgSub ensures that no value is present for UpPathChgSub, not even an explicit nil
### GetAddrPreserInd

`func (o *AfRoutingRequirementRm) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *AfRoutingRequirementRm) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *AfRoutingRequirementRm) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *AfRoutingRequirementRm) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### SetAddrPreserIndNil

`func (o *AfRoutingRequirementRm) SetAddrPreserIndNil(b bool)`

 SetAddrPreserIndNil sets the value for AddrPreserInd to be an explicit nil

### UnsetAddrPreserInd
`func (o *AfRoutingRequirementRm) UnsetAddrPreserInd()`

UnsetAddrPreserInd ensures that no value is present for AddrPreserInd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


