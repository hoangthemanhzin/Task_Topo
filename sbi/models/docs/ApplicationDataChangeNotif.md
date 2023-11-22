# ApplicationDataChangeNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IptvConfigData** | Pointer to [**IptvConfigData**](IptvConfigData.md) |  | [optional] 
**PfdData** | Pointer to [**PfdChangeNotification**](PfdChangeNotification.md) |  | [optional] 
**BdtPolicyData** | Pointer to [**BdtPolicyData**](BdtPolicyData.md) |  | [optional] 
**ResUri** | **string** |  | 
**SerParamData** | Pointer to [**ServiceParameterData**](ServiceParameterData.md) |  | [optional] 

## Methods

### NewApplicationDataChangeNotif

`func NewApplicationDataChangeNotif(resUri string, ) *ApplicationDataChangeNotif`

NewApplicationDataChangeNotif instantiates a new ApplicationDataChangeNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataChangeNotifWithDefaults

`func NewApplicationDataChangeNotifWithDefaults() *ApplicationDataChangeNotif`

NewApplicationDataChangeNotifWithDefaults instantiates a new ApplicationDataChangeNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIptvConfigData

`func (o *ApplicationDataChangeNotif) GetIptvConfigData() IptvConfigData`

GetIptvConfigData returns the IptvConfigData field if non-nil, zero value otherwise.

### GetIptvConfigDataOk

`func (o *ApplicationDataChangeNotif) GetIptvConfigDataOk() (*IptvConfigData, bool)`

GetIptvConfigDataOk returns a tuple with the IptvConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvConfigData

`func (o *ApplicationDataChangeNotif) SetIptvConfigData(v IptvConfigData)`

SetIptvConfigData sets IptvConfigData field to given value.

### HasIptvConfigData

`func (o *ApplicationDataChangeNotif) HasIptvConfigData() bool`

HasIptvConfigData returns a boolean if a field has been set.

### GetPfdData

`func (o *ApplicationDataChangeNotif) GetPfdData() PfdChangeNotification`

GetPfdData returns the PfdData field if non-nil, zero value otherwise.

### GetPfdDataOk

`func (o *ApplicationDataChangeNotif) GetPfdDataOk() (*PfdChangeNotification, bool)`

GetPfdDataOk returns a tuple with the PfdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdData

`func (o *ApplicationDataChangeNotif) SetPfdData(v PfdChangeNotification)`

SetPfdData sets PfdData field to given value.

### HasPfdData

`func (o *ApplicationDataChangeNotif) HasPfdData() bool`

HasPfdData returns a boolean if a field has been set.

### GetBdtPolicyData

`func (o *ApplicationDataChangeNotif) GetBdtPolicyData() BdtPolicyData`

GetBdtPolicyData returns the BdtPolicyData field if non-nil, zero value otherwise.

### GetBdtPolicyDataOk

`func (o *ApplicationDataChangeNotif) GetBdtPolicyDataOk() (*BdtPolicyData, bool)`

GetBdtPolicyDataOk returns a tuple with the BdtPolicyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtPolicyData

`func (o *ApplicationDataChangeNotif) SetBdtPolicyData(v BdtPolicyData)`

SetBdtPolicyData sets BdtPolicyData field to given value.

### HasBdtPolicyData

`func (o *ApplicationDataChangeNotif) HasBdtPolicyData() bool`

HasBdtPolicyData returns a boolean if a field has been set.

### GetResUri

`func (o *ApplicationDataChangeNotif) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *ApplicationDataChangeNotif) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *ApplicationDataChangeNotif) SetResUri(v string)`

SetResUri sets ResUri field to given value.


### GetSerParamData

`func (o *ApplicationDataChangeNotif) GetSerParamData() ServiceParameterData`

GetSerParamData returns the SerParamData field if non-nil, zero value otherwise.

### GetSerParamDataOk

`func (o *ApplicationDataChangeNotif) GetSerParamDataOk() (*ServiceParameterData, bool)`

GetSerParamDataOk returns a tuple with the SerParamData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerParamData

`func (o *ApplicationDataChangeNotif) SetSerParamData(v ServiceParameterData)`

SetSerParamData sets SerParamData field to given value.

### HasSerParamData

`func (o *ApplicationDataChangeNotif) HasSerParamData() bool`

HasSerParamData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


