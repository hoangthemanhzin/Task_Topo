# ExposureDataChangeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | Pointer to **string** |  | [optional] 
**AccessAndMobilityData** | Pointer to [**AccessAndMobilityData**](AccessAndMobilityData.md) |  | [optional] 
**PduSessionManagementData** | Pointer to [**[]PduSessionManagementData**](PduSessionManagementData.md) |  | [optional] 
**DelResources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExposureDataChangeNotification

`func NewExposureDataChangeNotification() *ExposureDataChangeNotification`

NewExposureDataChangeNotification instantiates a new ExposureDataChangeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExposureDataChangeNotificationWithDefaults

`func NewExposureDataChangeNotificationWithDefaults() *ExposureDataChangeNotification`

NewExposureDataChangeNotificationWithDefaults instantiates a new ExposureDataChangeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *ExposureDataChangeNotification) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ExposureDataChangeNotification) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ExposureDataChangeNotification) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ExposureDataChangeNotification) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAccessAndMobilityData

`func (o *ExposureDataChangeNotification) GetAccessAndMobilityData() AccessAndMobilityData`

GetAccessAndMobilityData returns the AccessAndMobilityData field if non-nil, zero value otherwise.

### GetAccessAndMobilityDataOk

`func (o *ExposureDataChangeNotification) GetAccessAndMobilityDataOk() (*AccessAndMobilityData, bool)`

GetAccessAndMobilityDataOk returns a tuple with the AccessAndMobilityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndMobilityData

`func (o *ExposureDataChangeNotification) SetAccessAndMobilityData(v AccessAndMobilityData)`

SetAccessAndMobilityData sets AccessAndMobilityData field to given value.

### HasAccessAndMobilityData

`func (o *ExposureDataChangeNotification) HasAccessAndMobilityData() bool`

HasAccessAndMobilityData returns a boolean if a field has been set.

### GetPduSessionManagementData

`func (o *ExposureDataChangeNotification) GetPduSessionManagementData() []PduSessionManagementData`

GetPduSessionManagementData returns the PduSessionManagementData field if non-nil, zero value otherwise.

### GetPduSessionManagementDataOk

`func (o *ExposureDataChangeNotification) GetPduSessionManagementDataOk() (*[]PduSessionManagementData, bool)`

GetPduSessionManagementDataOk returns a tuple with the PduSessionManagementData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionManagementData

`func (o *ExposureDataChangeNotification) SetPduSessionManagementData(v []PduSessionManagementData)`

SetPduSessionManagementData sets PduSessionManagementData field to given value.

### HasPduSessionManagementData

`func (o *ExposureDataChangeNotification) HasPduSessionManagementData() bool`

HasPduSessionManagementData returns a boolean if a field has been set.

### GetDelResources

`func (o *ExposureDataChangeNotification) GetDelResources() []string`

GetDelResources returns the DelResources field if non-nil, zero value otherwise.

### GetDelResourcesOk

`func (o *ExposureDataChangeNotification) GetDelResourcesOk() (*[]string, bool)`

GetDelResourcesOk returns a tuple with the DelResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelResources

`func (o *ExposureDataChangeNotification) SetDelResources(v []string)`

SetDelResources sets DelResources field to given value.

### HasDelResources

`func (o *ExposureDataChangeNotification) HasDelResources() bool`

HasDelResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


