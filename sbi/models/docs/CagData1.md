# CagData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CagInfos** | [**map[string]CagInfo1**](CagInfo1.md) | A map (list of key-value pairs where PlmnId serves as key) of CagInfo | 
**ProvisioningTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCagData1

`func NewCagData1(cagInfos map[string]CagInfo1, ) *CagData1`

NewCagData1 instantiates a new CagData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCagData1WithDefaults

`func NewCagData1WithDefaults() *CagData1`

NewCagData1WithDefaults instantiates a new CagData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCagInfos

`func (o *CagData1) GetCagInfos() map[string]CagInfo1`

GetCagInfos returns the CagInfos field if non-nil, zero value otherwise.

### GetCagInfosOk

`func (o *CagData1) GetCagInfosOk() (*map[string]CagInfo1, bool)`

GetCagInfosOk returns a tuple with the CagInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagInfos

`func (o *CagData1) SetCagInfos(v map[string]CagInfo1)`

SetCagInfos sets CagInfos field to given value.


### GetProvisioningTime

`func (o *CagData1) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *CagData1) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *CagData1) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.

### HasProvisioningTime

`func (o *CagData1) HasProvisioningTime() bool`

HasProvisioningTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


