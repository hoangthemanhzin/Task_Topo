# SmsManagementSubscriptionData1

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
**TraceData** | Pointer to [**TraceData**](TraceData.md) |  | [optional] 

## Methods

### NewSmsManagementSubscriptionData1

`func NewSmsManagementSubscriptionData1() *SmsManagementSubscriptionData1`

NewSmsManagementSubscriptionData1 instantiates a new SmsManagementSubscriptionData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsManagementSubscriptionData1WithDefaults

`func NewSmsManagementSubscriptionData1WithDefaults() *SmsManagementSubscriptionData1`

NewSmsManagementSubscriptionData1WithDefaults instantiates a new SmsManagementSubscriptionData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *SmsManagementSubscriptionData1) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmsManagementSubscriptionData1) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmsManagementSubscriptionData1) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmsManagementSubscriptionData1) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData1) GetMtSmsSubscribed() bool`

GetMtSmsSubscribed returns the MtSmsSubscribed field if non-nil, zero value otherwise.

### GetMtSmsSubscribedOk

`func (o *SmsManagementSubscriptionData1) GetMtSmsSubscribedOk() (*bool, bool)`

GetMtSmsSubscribedOk returns a tuple with the MtSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsSubscribed

`func (o *SmsManagementSubscriptionData1) SetMtSmsSubscribed(v bool)`

SetMtSmsSubscribed sets MtSmsSubscribed field to given value.

### HasMtSmsSubscribed

`func (o *SmsManagementSubscriptionData1) HasMtSmsSubscribed() bool`

HasMtSmsSubscribed returns a boolean if a field has been set.

### GetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData1) GetMtSmsBarringAll() bool`

GetMtSmsBarringAll returns the MtSmsBarringAll field if non-nil, zero value otherwise.

### GetMtSmsBarringAllOk

`func (o *SmsManagementSubscriptionData1) GetMtSmsBarringAllOk() (*bool, bool)`

GetMtSmsBarringAllOk returns a tuple with the MtSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringAll

`func (o *SmsManagementSubscriptionData1) SetMtSmsBarringAll(v bool)`

SetMtSmsBarringAll sets MtSmsBarringAll field to given value.

### HasMtSmsBarringAll

`func (o *SmsManagementSubscriptionData1) HasMtSmsBarringAll() bool`

HasMtSmsBarringAll returns a boolean if a field has been set.

### GetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) GetMtSmsBarringRoaming() bool`

GetMtSmsBarringRoaming returns the MtSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMtSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData1) GetMtSmsBarringRoamingOk() (*bool, bool)`

GetMtSmsBarringRoamingOk returns a tuple with the MtSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) SetMtSmsBarringRoaming(v bool)`

SetMtSmsBarringRoaming sets MtSmsBarringRoaming field to given value.

### HasMtSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) HasMtSmsBarringRoaming() bool`

HasMtSmsBarringRoaming returns a boolean if a field has been set.

### GetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData1) GetMoSmsSubscribed() bool`

GetMoSmsSubscribed returns the MoSmsSubscribed field if non-nil, zero value otherwise.

### GetMoSmsSubscribedOk

`func (o *SmsManagementSubscriptionData1) GetMoSmsSubscribedOk() (*bool, bool)`

GetMoSmsSubscribedOk returns a tuple with the MoSmsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsSubscribed

`func (o *SmsManagementSubscriptionData1) SetMoSmsSubscribed(v bool)`

SetMoSmsSubscribed sets MoSmsSubscribed field to given value.

### HasMoSmsSubscribed

`func (o *SmsManagementSubscriptionData1) HasMoSmsSubscribed() bool`

HasMoSmsSubscribed returns a boolean if a field has been set.

### GetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData1) GetMoSmsBarringAll() bool`

GetMoSmsBarringAll returns the MoSmsBarringAll field if non-nil, zero value otherwise.

### GetMoSmsBarringAllOk

`func (o *SmsManagementSubscriptionData1) GetMoSmsBarringAllOk() (*bool, bool)`

GetMoSmsBarringAllOk returns a tuple with the MoSmsBarringAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringAll

`func (o *SmsManagementSubscriptionData1) SetMoSmsBarringAll(v bool)`

SetMoSmsBarringAll sets MoSmsBarringAll field to given value.

### HasMoSmsBarringAll

`func (o *SmsManagementSubscriptionData1) HasMoSmsBarringAll() bool`

HasMoSmsBarringAll returns a boolean if a field has been set.

### GetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) GetMoSmsBarringRoaming() bool`

GetMoSmsBarringRoaming returns the MoSmsBarringRoaming field if non-nil, zero value otherwise.

### GetMoSmsBarringRoamingOk

`func (o *SmsManagementSubscriptionData1) GetMoSmsBarringRoamingOk() (*bool, bool)`

GetMoSmsBarringRoamingOk returns a tuple with the MoSmsBarringRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) SetMoSmsBarringRoaming(v bool)`

SetMoSmsBarringRoaming sets MoSmsBarringRoaming field to given value.

### HasMoSmsBarringRoaming

`func (o *SmsManagementSubscriptionData1) HasMoSmsBarringRoaming() bool`

HasMoSmsBarringRoaming returns a boolean if a field has been set.

### GetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData1) GetSharedSmsMngDataIds() []string`

GetSharedSmsMngDataIds returns the SharedSmsMngDataIds field if non-nil, zero value otherwise.

### GetSharedSmsMngDataIdsOk

`func (o *SmsManagementSubscriptionData1) GetSharedSmsMngDataIdsOk() (*[]string, bool)`

GetSharedSmsMngDataIdsOk returns a tuple with the SharedSmsMngDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData1) SetSharedSmsMngDataIds(v []string)`

SetSharedSmsMngDataIds sets SharedSmsMngDataIds field to given value.

### HasSharedSmsMngDataIds

`func (o *SmsManagementSubscriptionData1) HasSharedSmsMngDataIds() bool`

HasSharedSmsMngDataIds returns a boolean if a field has been set.

### GetTraceData

`func (o *SmsManagementSubscriptionData1) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmsManagementSubscriptionData1) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmsManagementSubscriptionData1) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmsManagementSubscriptionData1) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmsManagementSubscriptionData1) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmsManagementSubscriptionData1) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


