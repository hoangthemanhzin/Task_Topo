# PduSessionTsnBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsnBridgeInfo** | [**TsnBridgeInfo**](TsnBridgeInfo.md) |  | 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 

## Methods

### NewPduSessionTsnBridge

`func NewPduSessionTsnBridge(tsnBridgeInfo TsnBridgeInfo, ) *PduSessionTsnBridge`

NewPduSessionTsnBridge instantiates a new PduSessionTsnBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionTsnBridgeWithDefaults

`func NewPduSessionTsnBridgeWithDefaults() *PduSessionTsnBridge`

NewPduSessionTsnBridgeWithDefaults instantiates a new PduSessionTsnBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsnBridgeInfo

`func (o *PduSessionTsnBridge) GetTsnBridgeInfo() TsnBridgeInfo`

GetTsnBridgeInfo returns the TsnBridgeInfo field if non-nil, zero value otherwise.

### GetTsnBridgeInfoOk

`func (o *PduSessionTsnBridge) GetTsnBridgeInfoOk() (*TsnBridgeInfo, bool)`

GetTsnBridgeInfoOk returns a tuple with the TsnBridgeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeInfo

`func (o *PduSessionTsnBridge) SetTsnBridgeInfo(v TsnBridgeInfo)`

SetTsnBridgeInfo sets TsnBridgeInfo field to given value.


### GetTsnBridgeManCont

`func (o *PduSessionTsnBridge) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *PduSessionTsnBridge) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *PduSessionTsnBridge) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *PduSessionTsnBridge) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *PduSessionTsnBridge) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *PduSessionTsnBridge) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *PduSessionTsnBridge) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *PduSessionTsnBridge) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *PduSessionTsnBridge) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *PduSessionTsnBridge) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *PduSessionTsnBridge) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *PduSessionTsnBridge) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


