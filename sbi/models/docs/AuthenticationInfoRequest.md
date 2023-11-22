# AuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ServingNetworkName** | **string** |  | 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**AusfInstanceId** | **string** |  | 
**CellCagInfo** | Pointer to **[]string** |  | [optional] 
**N5gcInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthenticationInfoRequest

`func NewAuthenticationInfoRequest(servingNetworkName string, ausfInstanceId string, ) *AuthenticationInfoRequest`

NewAuthenticationInfoRequest instantiates a new AuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoRequestWithDefaults

`func NewAuthenticationInfoRequestWithDefaults() *AuthenticationInfoRequest`

NewAuthenticationInfoRequestWithDefaults instantiates a new AuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *AuthenticationInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServingNetworkName

`func (o *AuthenticationInfoRequest) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *AuthenticationInfoRequest) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *AuthenticationInfoRequest) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.


### GetResynchronizationInfo

`func (o *AuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *AuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *AuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *AuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetAusfInstanceId

`func (o *AuthenticationInfoRequest) GetAusfInstanceId() string`

GetAusfInstanceId returns the AusfInstanceId field if non-nil, zero value otherwise.

### GetAusfInstanceIdOk

`func (o *AuthenticationInfoRequest) GetAusfInstanceIdOk() (*string, bool)`

GetAusfInstanceIdOk returns a tuple with the AusfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInstanceId

`func (o *AuthenticationInfoRequest) SetAusfInstanceId(v string)`

SetAusfInstanceId sets AusfInstanceId field to given value.


### GetCellCagInfo

`func (o *AuthenticationInfoRequest) GetCellCagInfo() []string`

GetCellCagInfo returns the CellCagInfo field if non-nil, zero value otherwise.

### GetCellCagInfoOk

`func (o *AuthenticationInfoRequest) GetCellCagInfoOk() (*[]string, bool)`

GetCellCagInfoOk returns a tuple with the CellCagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellCagInfo

`func (o *AuthenticationInfoRequest) SetCellCagInfo(v []string)`

SetCellCagInfo sets CellCagInfo field to given value.

### HasCellCagInfo

`func (o *AuthenticationInfoRequest) HasCellCagInfo() bool`

HasCellCagInfo returns a boolean if a field has been set.

### GetN5gcInd

`func (o *AuthenticationInfoRequest) GetN5gcInd() bool`

GetN5gcInd returns the N5gcInd field if non-nil, zero value otherwise.

### GetN5gcIndOk

`func (o *AuthenticationInfoRequest) GetN5gcIndOk() (*bool, bool)`

GetN5gcIndOk returns a tuple with the N5gcInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN5gcInd

`func (o *AuthenticationInfoRequest) SetN5gcInd(v bool)`

SetN5gcInd sets N5gcInd field to given value.

### HasN5gcInd

`func (o *AuthenticationInfoRequest) HasN5gcInd() bool`

HasN5gcInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


