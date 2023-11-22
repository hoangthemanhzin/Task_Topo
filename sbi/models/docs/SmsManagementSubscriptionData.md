# SmsManagementSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**MtSmsSubscribed** | Pointer to **bool** |  | [optional] 
**MtSmsBarringAll** | Pointer to **bool** |  | [optional] 
**MtSmsBarringRoaming** | Pointer to **bool** |  | [optional] 
**MoSmsSubscribed** | Pointer to **bool** |  | [optional] 
**MoSmsBarringAll** | Pointer to **bool** |  | [optional] 
**MoSmsBarringRoaming** | Pointer to **bool** |  | [optional] 
**SharedSmsMngDataIds** | Pointer to **[]string** |  | [optional] 
**TraceData** | Pointer to [**TraceData1**](TraceData1.md) |  | [optional] 

## Methods

### NewSmsManagementSubscriptionData

`func NewSmsManagementSubscriptionData() *SmsManagementSubscriptionData`

NewSmsManagementSubscriptionData instantiates a new SmsManagementSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsManagementSubscriptionDataWithDefaults

`func NewSmsManagementSubscriptionDataWithDefaults() *SmsManagementSubscriptionData`

NewSmsManagementSubscriptionDataWithDefaults instantiates a new SmsManagementSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *SmsManagementSubscriptionData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmsManagementSubscriptionData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmsManagementSubscriptionData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmsManagementSubscriptionData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) GetMtSmsSubscribed() bool`

GetMtSmsSubscribed returns the MtSmsSubscribed field if non-nil, zero value otherwise.

### GetMtSmsSubscribedOk

`func (o *SmsManagementSubscriptionData) GetMtSmsSubscribedOk() (*bool, bool)`

GetMtSmsSubscribedOk returns a tuple with the MtSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) SetMtSmsSubscribed(v bool)`

SetMtSmsSubscribed sets MtSmsSubscribed field to given value.

### HasMtSmsSubscribed

`func (o *SmsManagementSubscriptionData) HasMtSmsSubscribed() bool`

HasMtSmsSubscribed returns a boolean if a field has been set.

### GetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringAll() bool`

GetMtSmsBarringAll returns the MtSmsBarringAll field if non-nil, zero value otherwise.

### GetMtSmsBarringAllOk

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringAllOk() (*bool, bool)`

GetMtSmsBarringAllOk returns a tuple with the MtSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) SetMtSmsBarringAll(v bool)`

SetMtSmsBarringAll sets MtSmsBarringAll field to given value.

### HasMtSmsBarringAll

`func (o *SmsManagementSubscriptionData) HasMtSmsBarringAll() bool`

HasMtSmsBarringAll returns a boolean if a field has been set.

### GetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoaming() bool`

GetMtSmsBarringRoaming returns the MtSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMtSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoamingOk() (*bool, bool)`

GetMtSmsBarringRoamingOk returns a tuple with the MtSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) SetMtSmsBarringRoaming(v bool)`

SetMtSmsBarringRoaming sets MtSmsBarringRoaming field to given value.

### HasMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) HasMtSmsBarringRoaming() bool`

HasMtSmsBarringRoaming returns a boolean if a field has been set.

### GetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) GetMoSmsSubscribed() bool`

GetMoSmsSubscribed returns the MoSmsSubscribed field if non-nil, zero value otherwise.

### GetMoSmsSubscribedOk

`func (o *SmsManagementSubscriptionData) GetMoSmsSubscribedOk() (*bool, bool)`

GetMoSmsSubscribedOk returns a tuple with the MoSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) SetMoSmsSubscribed(v bool)`

SetMoSmsSubscribed sets MoSmsSubscribed field to given value.

### HasMoSmsSubscribed

`func (o *SmsManagementSubscriptionData) HasMoSmsSubscribed() bool`

HasMoSmsSubscribed returns a boolean if a field has been set.

### GetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringAll() bool`

GetMoSmsBarringAll returns the MoSmsBarringAll field if non-nil, zero value otherwise.

### GetMoSmsBarringAllOk

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringAllOk() (*bool, bool)`

GetMoSmsBarringAllOk returns a tuple with the MoSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) SetMoSmsBarringAll(v bool)`

SetMoSmsBarringAll sets MoSmsBarringAll field to given value.

### HasMoSmsBarringAll

`func (o *SmsManagementSubscriptionData) HasMoSmsBarringAll() bool`

HasMoSmsBarringAll returns a boolean if a field has been set.

### GetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoaming() bool`

GetMoSmsBarringRoaming returns the MoSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMoSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoamingOk() (*bool, bool)`

GetMoSmsBarringRoamingOk returns a tuple with the MoSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) SetMoSmsBarringRoaming(v bool)`

SetMoSmsBarringRoaming sets MoSmsBarringRoaming field to given value.

### HasMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData) HasMoSmsBarringRoaming() bool`

HasMoSmsBarringRoaming returns a boolean if a field has been set.

### GetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIds() []string`

GetSharedSmsMngDataIds returns the SharedSmsMngDataIds field if non-nil, zero value otherwise.

### GetSharedSmsMngDataIdsOk

`func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIdsOk() (*[]string, bool)`

GetSharedSmsMngDataIdsOk returns a tuple with the SharedSmsMngDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) SetSharedSmsMngDataIds(v []string)`

SetSharedSmsMngDataIds sets SharedSmsMngDataIds field to given value.

### HasSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData) HasSharedSmsMngDataIds() bool`

HasSharedSmsMngDataIds returns a boolean if a field has been set.

### GetTraceData

`func (o *SmsManagementSubscriptionData) GetTraceData() TraceData1`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmsManagementSubscriptionData) GetTraceDataOk() (*TraceData1, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmsManagementSubscriptionData) SetTraceData(v TraceData1)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmsManagementSubscriptionData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmsManagementSubscriptionData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmsManagementSubscriptionData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


