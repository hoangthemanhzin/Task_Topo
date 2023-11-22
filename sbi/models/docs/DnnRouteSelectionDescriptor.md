# DnnRouteSelectionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**SscModes** | Pointer to [**[]SscMode**](SscMode.md) |  | [optional] 
**PduSessTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**AtsssInfo** | Pointer to **bool** | Indicates whether MA PDU session establishment is allowed for this DNN. When set to value true MA PDU session establishment is allowed for this DNN. | [optional] [default to false]

## Methods

### NewDnnRouteSelectionDescriptor

`func NewDnnRouteSelectionDescriptor(dnn string, ) *DnnRouteSelectionDescriptor`

NewDnnRouteSelectionDescriptor instantiates a new DnnRouteSelectionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnRouteSelectionDescriptorWithDefaults

`func NewDnnRouteSelectionDescriptorWithDefaults() *DnnRouteSelectionDescriptor`

NewDnnRouteSelectionDescriptorWithDefaults instantiates a new DnnRouteSelectionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnRouteSelectionDescriptor) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnRouteSelectionDescriptor) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnRouteSelectionDescriptor) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSscModes

`func (o *DnnRouteSelectionDescriptor) GetSscModes() []SscMode`

GetSscModes returns the SscModes field if non-nil, zero value otherwise.

### GetSscModesOk

`func (o *DnnRouteSelectionDescriptor) GetSscModesOk() (*[]SscMode, bool)`

GetSscModesOk returns a tuple with the SscModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSscModes

`func (o *DnnRouteSelectionDescriptor) SetSscModes(v []SscMode)`

SetSscModes sets SscModes field to given value.

### HasSscModes

`func (o *DnnRouteSelectionDescriptor) HasSscModes() bool`

HasSscModes returns a boolean if a field has been set.

### GetPduSessTypes

`func (o *DnnRouteSelectionDescriptor) GetPduSessTypes() []PduSessionType`

GetPduSessTypes returns the PduSessTypes field if non-nil, zero value otherwise.

### GetPduSessTypesOk

`func (o *DnnRouteSelectionDescriptor) GetPduSessTypesOk() (*[]PduSessionType, bool)`

GetPduSessTypesOk returns a tuple with the PduSessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessTypes

`func (o *DnnRouteSelectionDescriptor) SetPduSessTypes(v []PduSessionType)`

SetPduSessTypes sets PduSessTypes field to given value.

### HasPduSessTypes

`func (o *DnnRouteSelectionDescriptor) HasPduSessTypes() bool`

HasPduSessTypes returns a boolean if a field has been set.

### GetAtsssInfo

`func (o *DnnRouteSelectionDescriptor) GetAtsssInfo() bool`

GetAtsssInfo returns the AtsssInfo field if non-nil, zero value otherwise.

### GetAtsssInfoOk

`func (o *DnnRouteSelectionDescriptor) GetAtsssInfoOk() (*bool, bool)`

GetAtsssInfoOk returns a tuple with the AtsssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssInfo

`func (o *DnnRouteSelectionDescriptor) SetAtsssInfo(v bool)`

SetAtsssInfo sets AtsssInfo field to given value.

### HasAtsssInfo

`func (o *DnnRouteSelectionDescriptor) HasAtsssInfo() bool`

HasAtsssInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


