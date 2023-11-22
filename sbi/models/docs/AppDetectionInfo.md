# AppDetectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | A reference to the application detection filter configured at the UPF | 
**InstanceId** | Pointer to **string** | Identifier sent by the SMF in order to allow correlation of application Start and Stop events to the specific service data flow description, if service data flow descriptions are deducible. | [optional] 
**SdfDescriptions** | Pointer to [**[]FlowInformation**](FlowInformation.md) | Contains the detected service data flow descriptions if they are deducible. | [optional] 

## Methods

### NewAppDetectionInfo

`func NewAppDetectionInfo(appId string, ) *AppDetectionInfo`

NewAppDetectionInfo instantiates a new AppDetectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDetectionInfoWithDefaults

`func NewAppDetectionInfoWithDefaults() *AppDetectionInfo`

NewAppDetectionInfoWithDefaults instantiates a new AppDetectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppDetectionInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppDetectionInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppDetectionInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInstanceId

`func (o *AppDetectionInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppDetectionInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppDetectionInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppDetectionInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetSdfDescriptions

`func (o *AppDetectionInfo) GetSdfDescriptions() []FlowInformation`

GetSdfDescriptions returns the SdfDescriptions field if non-nil, zero value otherwise.

### GetSdfDescriptionsOk

`func (o *AppDetectionInfo) GetSdfDescriptionsOk() (*[]FlowInformation, bool)`

GetSdfDescriptionsOk returns a tuple with the SdfDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdfDescriptions

`func (o *AppDetectionInfo) SetSdfDescriptions(v []FlowInformation)`

SetSdfDescriptions sets SdfDescriptions field to given value.

### HasSdfDescriptions

`func (o *AppDetectionInfo) HasSdfDescriptions() bool`

HasSdfDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


