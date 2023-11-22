# EnableUeReachabilityRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachability** | [**UeReachability**](UeReachability.md) |  | 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewEnableUeReachabilityRspData

`func NewEnableUeReachabilityRspData(reachability UeReachability, ) *EnableUeReachabilityRspData`

NewEnableUeReachabilityRspData instantiates a new EnableUeReachabilityRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableUeReachabilityRspDataWithDefaults

`func NewEnableUeReachabilityRspDataWithDefaults() *EnableUeReachabilityRspData`

NewEnableUeReachabilityRspDataWithDefaults instantiates a new EnableUeReachabilityRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachability

`func (o *EnableUeReachabilityRspData) GetReachability() UeReachability`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *EnableUeReachabilityRspData) GetReachabilityOk() (*UeReachability, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *EnableUeReachabilityRspData) SetReachability(v UeReachability)`

SetReachability sets Reachability field to given value.


### GetSupportedFeatures

`func (o *EnableUeReachabilityRspData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EnableUeReachabilityRspData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EnableUeReachabilityRspData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EnableUeReachabilityRspData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


