# Dynamic5Qi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | [**QosResourceType**](QosResourceType.md) |  | 
**PriorityLevel** | **int32** |  | 
**PacketDelayBudget** | **int32** |  | 
**PacketErrRate** | **string** |  | 
**AverWindow** | Pointer to **int32** |  | [optional] [default to 2000]
**MaxDataBurstVol** | Pointer to **int32** |  | [optional] 
**ExtMaxDataBurstVol** | Pointer to **int32** |  | [optional] 
**ExtPacketDelBudget** | Pointer to **int32** |  | [optional] 
**CnPacketDelayBudgetDl** | Pointer to **int32** |  | [optional] 
**CnPacketDelayBudgetUl** | Pointer to **int32** |  | [optional] 

## Methods

### NewDynamic5Qi

`func NewDynamic5Qi(resourceType QosResourceType, priorityLevel int32, packetDelayBudget int32, packetErrRate string, ) *Dynamic5Qi`

NewDynamic5Qi instantiates a new Dynamic5Qi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamic5QiWithDefaults

`func NewDynamic5QiWithDefaults() *Dynamic5Qi`

NewDynamic5QiWithDefaults instantiates a new Dynamic5Qi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *Dynamic5Qi) GetResourceType() QosResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Dynamic5Qi) GetResourceTypeOk() (*QosResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Dynamic5Qi) SetResourceType(v QosResourceType)`

SetResourceType sets ResourceType field to given value.


### GetPriorityLevel

`func (o *Dynamic5Qi) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *Dynamic5Qi) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *Dynamic5Qi) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.


### GetPacketDelayBudget

`func (o *Dynamic5Qi) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *Dynamic5Qi) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *Dynamic5Qi) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.


### GetPacketErrRate

`func (o *Dynamic5Qi) GetPacketErrRate() string`

GetPacketErrRate returns the PacketErrRate field if non-nil, zero value otherwise.

### GetPacketErrRateOk

`func (o *Dynamic5Qi) GetPacketErrRateOk() (*string, bool)`

GetPacketErrRateOk returns a tuple with the PacketErrRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketErrRate

`func (o *Dynamic5Qi) SetPacketErrRate(v string)`

SetPacketErrRate sets PacketErrRate field to given value.


### GetAverWindow

`func (o *Dynamic5Qi) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *Dynamic5Qi) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *Dynamic5Qi) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *Dynamic5Qi) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### GetMaxDataBurstVol

`func (o *Dynamic5Qi) GetMaxDataBurstVol() int32`

GetMaxDataBurstVol returns the MaxDataBurstVol field if non-nil, zero value otherwise.

### GetMaxDataBurstVolOk

`func (o *Dynamic5Qi) GetMaxDataBurstVolOk() (*int32, bool)`

GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataBurstVol

`func (o *Dynamic5Qi) SetMaxDataBurstVol(v int32)`

SetMaxDataBurstVol sets MaxDataBurstVol field to given value.

### HasMaxDataBurstVol

`func (o *Dynamic5Qi) HasMaxDataBurstVol() bool`

HasMaxDataBurstVol returns a boolean if a field has been set.

### GetExtMaxDataBurstVol

`func (o *Dynamic5Qi) GetExtMaxDataBurstVol() int32`

GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field if non-nil, zero value otherwise.

### GetExtMaxDataBurstVolOk

`func (o *Dynamic5Qi) GetExtMaxDataBurstVolOk() (*int32, bool)`

GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMaxDataBurstVol

`func (o *Dynamic5Qi) SetExtMaxDataBurstVol(v int32)`

SetExtMaxDataBurstVol sets ExtMaxDataBurstVol field to given value.

### HasExtMaxDataBurstVol

`func (o *Dynamic5Qi) HasExtMaxDataBurstVol() bool`

HasExtMaxDataBurstVol returns a boolean if a field has been set.

### GetExtPacketDelBudget

`func (o *Dynamic5Qi) GetExtPacketDelBudget() int32`

GetExtPacketDelBudget returns the ExtPacketDelBudget field if non-nil, zero value otherwise.

### GetExtPacketDelBudgetOk

`func (o *Dynamic5Qi) GetExtPacketDelBudgetOk() (*int32, bool)`

GetExtPacketDelBudgetOk returns a tuple with the ExtPacketDelBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtPacketDelBudget

`func (o *Dynamic5Qi) SetExtPacketDelBudget(v int32)`

SetExtPacketDelBudget sets ExtPacketDelBudget field to given value.

### HasExtPacketDelBudget

`func (o *Dynamic5Qi) HasExtPacketDelBudget() bool`

HasExtPacketDelBudget returns a boolean if a field has been set.

### GetCnPacketDelayBudgetDl

`func (o *Dynamic5Qi) GetCnPacketDelayBudgetDl() int32`

GetCnPacketDelayBudgetDl returns the CnPacketDelayBudgetDl field if non-nil, zero value otherwise.

### GetCnPacketDelayBudgetDlOk

`func (o *Dynamic5Qi) GetCnPacketDelayBudgetDlOk() (*int32, bool)`

GetCnPacketDelayBudgetDlOk returns a tuple with the CnPacketDelayBudgetDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnPacketDelayBudgetDl

`func (o *Dynamic5Qi) SetCnPacketDelayBudgetDl(v int32)`

SetCnPacketDelayBudgetDl sets CnPacketDelayBudgetDl field to given value.

### HasCnPacketDelayBudgetDl

`func (o *Dynamic5Qi) HasCnPacketDelayBudgetDl() bool`

HasCnPacketDelayBudgetDl returns a boolean if a field has been set.

### GetCnPacketDelayBudgetUl

`func (o *Dynamic5Qi) GetCnPacketDelayBudgetUl() int32`

GetCnPacketDelayBudgetUl returns the CnPacketDelayBudgetUl field if non-nil, zero value otherwise.

### GetCnPacketDelayBudgetUlOk

`func (o *Dynamic5Qi) GetCnPacketDelayBudgetUlOk() (*int32, bool)`

GetCnPacketDelayBudgetUlOk returns a tuple with the CnPacketDelayBudgetUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnPacketDelayBudgetUl

`func (o *Dynamic5Qi) SetCnPacketDelayBudgetUl(v int32)`

SetCnPacketDelayBudgetUl sets CnPacketDelayBudgetUl field to given value.

### HasCnPacketDelayBudgetUl

`func (o *Dynamic5Qi) HasCnPacketDelayBudgetUl() bool`

HasCnPacketDelayBudgetUl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


