# CagData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CagInfos** | [**map[string]CagInfo**](CagInfo.md) | A map (list of key-value pairs where PlmnId serves as key) of CagInfo | 
**ProvisioningTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCagData

`func NewCagData(cagInfos map[string]CagInfo, ) *CagData`

NewCagData instantiates a new CagData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCagDataWithDefaults

`func NewCagDataWithDefaults() *CagData`

NewCagDataWithDefaults instantiates a new CagData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCagInfos

`func (o *CagData) GetCagInfos() map[string]CagInfo`

GetCagInfos returns the CagInfos field if non-nil, zero value otherwise.

### GetCagInfosOk

`func (o *CagData) GetCagInfosOk() (*map[string]CagInfo, bool)`

GetCagInfosOk returns a tuple with the CagInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagInfos

`func (o *CagData) SetCagInfos(v map[string]CagInfo)`

SetCagInfos sets CagInfos field to given value.


### GetProvisioningTime

`func (o *CagData) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *CagData) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *CagData) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *CagData) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


