# SharedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDataId** | **string** |  | 
**SharedAmData** | Pointer to [**AccessAndMobilitySubscriptionData1**](AccessAndMobilitySubscriptionData1.md) |  | [optional] 
**SharedSmsSubsData** | Pointer to [**SmsSubscriptionData1**](SmsSubscriptionData1.md) |  | [optional] 
**SharedSmsMngSubsData** | Pointer to [**SmsManagementSubscriptionData1**](SmsManagementSubscriptionData1.md) |  | [optional] 
**SharedDnnConfigurations** | Pointer to [**map[string]DnnConfiguration1**](DnnConfiguration1.md) |  | [optional] 
**SharedTraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 
**SharedSnssaiInfos** | Pointer to [**map[string]SnssaiInfo**](SnssaiInfo.md) |  | [optional] 
**SharedVnGroupDatas** | Pointer to [**map[string]VnGroupData**](VnGroupData.md) |  | [optional] 

## Methods

### NewSharedData

`func NewSharedData(sharedDataId string, ) *SharedData`

NewSharedData instantiates a new SharedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDataWithDefaults

`func NewSharedDataWithDefaults() *SharedData`

NewSharedDataWithDefaults instantiates a new SharedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedDataId

`func (o *SharedData) GetSharedDataId() string`

GetSharedDataId returns the SharedDataId field if non-nil, zero value otherwise.

### GetSharedDataIdOk

`func (o *SharedData) GetSharedDataIdOk() (*string, bool)`

GetSharedDataIdOk returns a tuple with the SharedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataId

`func (o *SharedData) SetSharedDataId(v string)`

SetSharedDataId sets SharedDataId field to given value.


### GetSharedAmData

`func (o *SharedData) GetSharedAmData() AccessAndMobilitySubscriptionData1`

GetSharedAmData returns the SharedAmData field if non-nil, zero value otherwise.

### GetSharedAmDataOk

`func (o *SharedData) GetSharedAmDataOk() (*AccessAndMobilitySubscriptionData1, bool)`

GetSharedAmDataOk returns a tuple with the SharedAmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAmData

`func (o *SharedData) SetSharedAmData(v AccessAndMobilitySubscriptionData1)`

SetSharedAmData sets SharedAmData field to given value.

### HasSharedAmData

`func (o *SharedData) HasSharedAmData() bool`

HasSharedAmData returns a boolean if a field has been set.

### GetSharedSmsSubsData

`func (o *SharedData) GetSharedSmsSubsData() SmsSubscriptionData1`

GetSharedSmsSubsData returns the SharedSmsSubsData field if non-nil, zero value otherwise.

### GetSharedSmsSubsDataOk

`func (o *SharedData) GetSharedSmsSubsDataOk() (*SmsSubscriptionData1, bool)`

GetSharedSmsSubsDataOk returns a tuple with the SharedSmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsSubsData

`func (o *SharedData) SetSharedSmsSubsData(v SmsSubscriptionData1)`

SetSharedSmsSubsData sets SharedSmsSubsData field to given value.

### HasSharedSmsSubsData

`func (o *SharedData) HasSharedSmsSubsData() bool`

HasSharedSmsSubsData returns a boolean if a field has been set.

### GetSharedSmsMngSubsData

`func (o *SharedData) GetSharedSmsMngSubsData() SmsManagementSubscriptionData1`

GetSharedSmsMngSubsData returns the SharedSmsMngSubsData field if non-nil, zero value otherwise.

### GetSharedSmsMngSubsDataOk

`func (o *SharedData) GetSharedSmsMngSubsDataOk() (*SmsManagementSubscriptionData1, bool)`

GetSharedSmsMngSubsDataOk returns a tuple with the SharedSmsMngSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsMngSubsData

`func (o *SharedData) SetSharedSmsMngSubsData(v SmsManagementSubscriptionData1)`

SetSharedSmsMngSubsData sets SharedSmsMngSubsData field to given value.

### HasSharedSmsMngSubsData

`func (o *SharedData) HasSharedSmsMngSubsData() bool`

HasSharedSmsMngSubsData returns a boolean if a field has been set.

### GetSharedDnnConfigurations

`func (o *SharedData) GetSharedDnnConfigurations() map[string]DnnConfiguration1`

GetSharedDnnConfigurations returns the SharedDnnConfigurations field if non-nil, zero value otherwise.

### GetSharedDnnConfigurationsOk

`func (o *SharedData) GetSharedDnnConfigurationsOk() (*map[string]DnnConfiguration1, bool)`

GetSharedDnnConfigurationsOk returns a tuple with the SharedDnnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDnnConfigurations

`func (o *SharedData) SetSharedDnnConfigurations(v map[string]DnnConfiguration1)`

SetSharedDnnConfigurations sets SharedDnnConfigurations field to given value.

### HasSharedDnnConfigurations

`func (o *SharedData) HasSharedDnnConfigurations() bool`

HasSharedDnnConfigurations returns a boolean if a field has been set.

### GetSharedTraceData

`func (o *SharedData) GetSharedTraceData() TraceData`

GetSharedTraceData returns the SharedTraceData field if non-nil, zero value otherwise.

### GetSharedTraceDataOk

`func (o *SharedData) GetSharedTraceDataOk() (*TraceData, bool)`

GetSharedTraceDataOk returns a tuple with the SharedTraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTraceData

`func (o *SharedData) SetSharedTraceData(v TraceData)`

SetSharedTraceData sets SharedTraceData field to given value.

### HasSharedTraceData

`func (o *SharedData) HasSharedTraceData() bool`

HasSharedTraceData returns a boolean if a field has been set.

### SetSharedTraceDataNil

`func (o *SharedData) SetSharedTraceDataNil(b bool)`

 SetSharedTraceDataNil sets the value for SharedTraceData to be an explicit nil

### UnsetSharedTraceData
`func (o *SharedData) UnsetSharedTraceData()`

UnsetSharedTraceData ensures that no value is present for SharedTraceData, not even an explicit nil
### GetSharedSnssaiInfos

`func (o *SharedData) GetSharedSnssaiInfos() map[string]SnssaiInfo`

GetSharedSnssaiInfos returns the SharedSnssaiInfos field if non-nil, zero value otherwise.

### GetSharedSnssaiInfosOk

`func (o *SharedData) GetSharedSnssaiInfosOk() (*map[string]SnssaiInfo, bool)`

GetSharedSnssaiInfosOk returns a tuple with the SharedSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSnssaiInfos

`func (o *SharedData) SetSharedSnssaiInfos(v map[string]SnssaiInfo)`

SetSharedSnssaiInfos sets SharedSnssaiInfos field to given value.

### HasSharedSnssaiInfos

`func (o *SharedData) HasSharedSnssaiInfos() bool`

HasSharedSnssaiInfos returns a boolean if a field has been set.

### GetSharedVnGroupDatas

`func (o *SharedData) GetSharedVnGroupDatas() map[string]VnGroupData`

GetSharedVnGroupDatas returns the SharedVnGroupDatas field if non-nil, zero value otherwise.

### GetSharedVnGroupDatasOk

`func (o *SharedData) GetSharedVnGroupDatasOk() (*map[string]VnGroupData, bool)`

GetSharedVnGroupDatasOk returns a tuple with the SharedVnGroupDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVnGroupDatas

`func (o *SharedData) SetSharedVnGroupDatas(v map[string]VnGroupData)`

SetSharedVnGroupDatas sets SharedVnGroupDatas field to given value.

### HasSharedVnGroupDatas

`func (o *SharedData) HasSharedVnGroupDatas() bool`

HasSharedVnGroupDatas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


