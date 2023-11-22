# DatalinkReportingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DddTrafficDes** | Pointer to [**[]DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**Slice** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**DddStatusList** | Pointer to [**[]DlDataDeliveryStatus**](DlDataDeliveryStatus.md) |  | [optional] 

## Methods

### NewDatalinkReportingConfiguration

`func NewDatalinkReportingConfiguration() *DatalinkReportingConfiguration`

NewDatalinkReportingConfiguration instantiates a new DatalinkReportingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatalinkReportingConfigurationWithDefaults

`func NewDatalinkReportingConfigurationWithDefaults() *DatalinkReportingConfiguration`

NewDatalinkReportingConfigurationWithDefaults instantiates a new DatalinkReportingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDddTrafficDes

`func (o *DatalinkReportingConfiguration) GetDddTrafficDes() []DddTrafficDescriptor`

GetDddTrafficDes returns the DddTrafficDes field if non-nil, zero value otherwise.

### GetDddTrafficDesOk

`func (o *DatalinkReportingConfiguration) GetDddTrafficDesOk() (*[]DddTrafficDescriptor, bool)`

GetDddTrafficDesOk returns a tuple with the DddTrafficDes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTrafficDes

`func (o *DatalinkReportingConfiguration) SetDddTrafficDes(v []DddTrafficDescriptor)`

SetDddTrafficDes sets DddTrafficDes field to given value.

### HasDddTrafficDes

`func (o *DatalinkReportingConfiguration) HasDddTrafficDes() bool`

HasDddTrafficDes returns a boolean if a field has been set.

### GetDnn

`func (o *DatalinkReportingConfiguration) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DatalinkReportingConfiguration) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DatalinkReportingConfiguration) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *DatalinkReportingConfiguration) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSlice

`func (o *DatalinkReportingConfiguration) GetSlice() Snssai`

GetSlice returns the Slice field if non-nil, zero value otherwise.

### GetSliceOk

`func (o *DatalinkReportingConfiguration) GetSliceOk() (*Snssai, bool)`

GetSliceOk returns a tuple with the Slice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlice

`func (o *DatalinkReportingConfiguration) SetSlice(v Snssai)`

SetSlice sets Slice field to given value.

### HasSlice

`func (o *DatalinkReportingConfiguration) HasSlice() bool`

HasSlice returns a boolean if a field has been set.

### GetDddStatusList

`func (o *DatalinkReportingConfiguration) GetDddStatusList() []DlDataDeliveryStatus`

GetDddStatusList returns the DddStatusList field if non-nil, zero value otherwise.

### GetDddStatusListOk

`func (o *DatalinkReportingConfiguration) GetDddStatusListOk() (*[]DlDataDeliveryStatus, bool)`

GetDddStatusListOk returns a tuple with the DddStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddStatusList

`func (o *DatalinkReportingConfiguration) SetDddStatusList(v []DlDataDeliveryStatus)`

SetDddStatusList sets DddStatusList field to given value.

### HasDddStatusList

`func (o *DatalinkReportingConfiguration) HasDddStatusList() bool`

HasDddStatusList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


