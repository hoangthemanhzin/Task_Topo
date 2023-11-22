# ResourcesAllocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**McResourcStatus** | Pointer to [**MediaComponentResourcesStatus**](MediaComponentResourcesStatus.md) |  | [optional] 
**Flows** | Pointer to [**[]Flows**](Flows.md) |  | [optional] 
**AltSerReq** | Pointer to **string** |  | [optional] 

## Methods

### NewResourcesAllocationInfo

`func NewResourcesAllocationInfo() *ResourcesAllocationInfo`

NewResourcesAllocationInfo instantiates a new ResourcesAllocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesAllocationInfoWithDefaults

`func NewResourcesAllocationInfoWithDefaults() *ResourcesAllocationInfo`

NewResourcesAllocationInfoWithDefaults instantiates a new ResourcesAllocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcResourcStatus

`func (o *ResourcesAllocationInfo) GetMcResourcStatus() MediaComponentResourcesStatus`

GetMcResourcStatus returns the McResourcStatus field if non-nil, zero value otherwise.

### GetMcResourcStatusOk

`func (o *ResourcesAllocationInfo) GetMcResourcStatusOk() (*MediaComponentResourcesStatus, bool)`

GetMcResourcStatusOk returns a tuple with the McResourcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcResourcStatus

`func (o *ResourcesAllocationInfo) SetMcResourcStatus(v MediaComponentResourcesStatus)`

SetMcResourcStatus sets McResourcStatus field to given value.

### HasMcResourcStatus

`func (o *ResourcesAllocationInfo) HasMcResourcStatus() bool`

HasMcResourcStatus returns a boolean if a field has been set.

### GetFlows

`func (o *ResourcesAllocationInfo) GetFlows() []Flows`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ResourcesAllocationInfo) GetFlowsOk() (*[]Flows, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ResourcesAllocationInfo) SetFlows(v []Flows)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *ResourcesAllocationInfo) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetAltSerReq

`func (o *ResourcesAllocationInfo) GetAltSerReq() string`

GetAltSerReq returns the AltSerReq field if non-nil, zero value otherwise.

### GetAltSerReqOk

`func (o *ResourcesAllocationInfo) GetAltSerReqOk() (*string, bool)`

GetAltSerReqOk returns a tuple with the AltSerReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSerReq

`func (o *ResourcesAllocationInfo) SetAltSerReq(v string)`

SetAltSerReq sets AltSerReq field to given value.

### HasAltSerReq

`func (o *ResourcesAllocationInfo) HasAltSerReq() bool`

HasAltSerReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


