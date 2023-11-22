# ProvisionedDataSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmData** | Pointer to [**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] 
**SmfSelData** | Pointer to [**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md) |  | [optional] 
**SmsSubsData** | Pointer to [**SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] 
**SmData** | Pointer to [**[]SessionManagementSubscriptionData**](SessionManagementSubscriptionData.md) |  | [optional] 
**TraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 
**SmsMngData** | Pointer to [**SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] 
**LcsPrivacyData** | Pointer to [**LcsPrivacyData**](LcsPrivacyData.md) |  | [optional] 
**LcsMoData** | Pointer to [**LcsMoData**](LcsMoData.md) |  | [optional] 
**LcsBcaData** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 
**V2xData** | Pointer to [**V2xSubscriptionData**](V2xSubscriptionData.md) |  | [optional] 

## Methods

### NewProvisionedDataSets

`func NewProvisionedDataSets() *ProvisionedDataSets`

NewProvisionedDataSets instantiates a new ProvisionedDataSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedDataSetsWithDefaults

`func NewProvisionedDataSetsWithDefaults() *ProvisionedDataSets`

NewProvisionedDataSetsWithDefaults instantiates a new ProvisionedDataSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmData

`func (o *ProvisionedDataSets) GetAmData() AccessAndMobilitySubscriptionData`

GetAmData returns the AmData field if non-nil, zero value otherwise.

### GetAmDataOk

`func (o *ProvisionedDataSets) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool)`

GetAmDataOk returns a tuple with the AmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmData

`func (o *ProvisionedDataSets) SetAmData(v AccessAndMobilitySubscriptionData)`

SetAmData sets AmData field to given value.

### HasAmData

`func (o *ProvisionedDataSets) HasAmData() bool`

HasAmData returns a boolean if a field has been set.

### GetSmfSelData

`func (o *ProvisionedDataSets) GetSmfSelData() SmfSelectionSubscriptionData`

GetSmfSelData returns the SmfSelData field if non-nil, zero value otherwise.

### GetSmfSelDataOk

`func (o *ProvisionedDataSets) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool)`

GetSmfSelDataOk returns a tuple with the SmfSelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelData

`func (o *ProvisionedDataSets) SetSmfSelData(v SmfSelectionSubscriptionData)`

SetSmfSelData sets SmfSelData field to given value.

### HasSmfSelData

`func (o *ProvisionedDataSets) HasSmfSelData() bool`

HasSmfSelData returns a boolean if a field has been set.

### GetSmsSubsData

`func (o *ProvisionedDataSets) GetSmsSubsData() SmsSubscriptionData`

GetSmsSubsData returns the SmsSubsData field if non-nil, zero value otherwise.

### GetSmsSubsDataOk

`func (o *ProvisionedDataSets) GetSmsSubsDataOk() (*SmsSubscriptionData, bool)`

GetSmsSubsDataOk returns a tuple with the SmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSubsData

`func (o *ProvisionedDataSets) SetSmsSubsData(v SmsSubscriptionData)`

SetSmsSubsData sets SmsSubsData field to given value.

### HasSmsSubsData

`func (o *ProvisionedDataSets) HasSmsSubsData() bool`

HasSmsSubsData returns a boolean if a field has been set.

### GetSmData

`func (o *ProvisionedDataSets) GetSmData() []SessionManagementSubscriptionData`

GetSmData returns the SmData field if non-nil, zero value otherwise.

### GetSmDataOk

`func (o *ProvisionedDataSets) GetSmDataOk() (*[]SessionManagementSubscriptionData, bool)`

GetSmDataOk returns a tuple with the SmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmData

`func (o *ProvisionedDataSets) SetSmData(v []SessionManagementSubscriptionData)`

SetSmData sets SmData field to given value.

### HasSmData

`func (o *ProvisionedDataSets) HasSmData() bool`

HasSmData returns a boolean if a field has been set.

### GetTraceData

`func (o *ProvisionedDataSets) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *ProvisionedDataSets) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *ProvisionedDataSets) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *ProvisionedDataSets) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *ProvisionedDataSets) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *ProvisionedDataSets) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSmsMngData

`func (o *ProvisionedDataSets) GetSmsMngData() SmsManagementSubscriptionData`

GetSmsMngData returns the SmsMngData field if non-nil, zero value otherwise.

### GetSmsMngDataOk

`func (o *ProvisionedDataSets) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool)`

GetSmsMngDataOk returns a tuple with the SmsMngData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMngData

`func (o *ProvisionedDataSets) SetSmsMngData(v SmsManagementSubscriptionData)`

SetSmsMngData sets SmsMngData field to given value.

### HasSmsMngData

`func (o *ProvisionedDataSets) HasSmsMngData() bool`

HasSmsMngData returns a boolean if a field has been set.

### GetLcsPrivacyData

`func (o *ProvisionedDataSets) GetLcsPrivacyData() LcsPrivacyData`

GetLcsPrivacyData returns the LcsPrivacyData field if non-nil, zero value otherwise.

### GetLcsPrivacyDataOk

`func (o *ProvisionedDataSets) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool)`

GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacyData

`func (o *ProvisionedDataSets) SetLcsPrivacyData(v LcsPrivacyData)`

SetLcsPrivacyData sets LcsPrivacyData field to given value.

### HasLcsPrivacyData

`func (o *ProvisionedDataSets) HasLcsPrivacyData() bool`

HasLcsPrivacyData returns a boolean if a field has been set.

### GetLcsMoData

`func (o *ProvisionedDataSets) GetLcsMoData() LcsMoData`

GetLcsMoData returns the LcsMoData field if non-nil, zero value otherwise.

### GetLcsMoDataOk

`func (o *ProvisionedDataSets) GetLcsMoDataOk() (*LcsMoData, bool)`

GetLcsMoDataOk returns a tuple with the LcsMoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsMoData

`func (o *ProvisionedDataSets) SetLcsMoData(v LcsMoData)`

SetLcsMoData sets LcsMoData field to given value.

### HasLcsMoData

`func (o *ProvisionedDataSets) HasLcsMoData() bool`

HasLcsMoData returns a boolean if a field has been set.

### GetLcsBcaData

`func (o *ProvisionedDataSets) GetLcsBcaData() LcsBroadcastAssistanceTypesData`

GetLcsBcaData returns the LcsBcaData field if non-nil, zero value otherwise.

### GetLcsBcaDataOk

`func (o *ProvisionedDataSets) GetLcsBcaDataOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetLcsBcaDataOk returns a tuple with the LcsBcaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsBcaData

`func (o *ProvisionedDataSets) SetLcsBcaData(v LcsBroadcastAssistanceTypesData)`

SetLcsBcaData sets LcsBcaData field to given value.

### HasLcsBcaData

`func (o *ProvisionedDataSets) HasLcsBcaData() bool`

HasLcsBcaData returns a boolean if a field has been set.

### GetV2xData

`func (o *ProvisionedDataSets) GetV2xData() V2xSubscriptionData`

GetV2xData returns the V2xData field if non-nil, zero value otherwise.

### GetV2xDataOk

`func (o *ProvisionedDataSets) GetV2xDataOk() (*V2xSubscriptionData, bool)`

GetV2xDataOk returns a tuple with the V2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xData

`func (o *ProvisionedDataSets) SetV2xData(v V2xSubscriptionData)`

SetV2xData sets V2xData field to given value.

### HasV2xData

`func (o *ProvisionedDataSets) HasV2xData() bool`

HasV2xData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


