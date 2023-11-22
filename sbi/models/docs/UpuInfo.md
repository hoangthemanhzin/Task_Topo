# UpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpuDataList** | [**[]UpuData1**](UpuData1.md) |  | 
**UpuRegInd** | **bool** |  | 
**UpuAckInd** | **bool** |  | 
**UpuMacIausf** | Pointer to **string** |  | [optional] 
**CounterUpu** | Pointer to **string** |  | [optional] 
**ProvisioningTime** | **time.Time** |  | 

## Methods

### NewUpuInfo

`func NewUpuInfo(upuDataList []UpuData1, upuRegInd bool, upuAckInd bool, provisioningTime time.Time, ) *UpuInfo`

NewUpuInfo instantiates a new UpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuInfoWithDefaults

`func NewUpuInfoWithDefaults() *UpuInfo`

NewUpuInfoWithDefaults instantiates a new UpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpuDataList

`func (o *UpuInfo) GetUpuDataList() []UpuData1`

GetUpuDataList returns the UpuDataList field if non-nil, zero value otherwise.

### GetUpuDataListOk

`func (o *UpuInfo) GetUpuDataListOk() (*[]UpuData1, bool)`

GetUpuDataListOk returns a tuple with the UpuDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuDataList

`func (o *UpuInfo) SetUpuDataList(v []UpuData1)`

SetUpuDataList sets UpuDataList field to given value.


### GetUpuRegInd

`func (o *UpuInfo) GetUpuRegInd() bool`

GetUpuRegInd returns the UpuRegInd field if non-nil, zero value otherwise.

### GetUpuRegIndOk

`func (o *UpuInfo) GetUpuRegIndOk() (*bool, bool)`

GetUpuRegIndOk returns a tuple with the UpuRegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuRegInd

`func (o *UpuInfo) SetUpuRegInd(v bool)`

SetUpuRegInd sets UpuRegInd field to given value.


### GetUpuAckInd

`func (o *UpuInfo) GetUpuAckInd() bool`

GetUpuAckInd returns the UpuAckInd field if non-nil, zero value otherwise.

### GetUpuAckIndOk

`func (o *UpuInfo) GetUpuAckIndOk() (*bool, bool)`

GetUpuAckIndOk returns a tuple with the UpuAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuAckInd

`func (o *UpuInfo) SetUpuAckInd(v bool)`

SetUpuAckInd sets UpuAckInd field to given value.


### GetUpuMacIausf

`func (o *UpuInfo) GetUpuMacIausf() string`

GetUpuMacIausf returns the UpuMacIausf field if non-nil, zero value otherwise.

### GetUpuMacIausfOk

`func (o *UpuInfo) GetUpuMacIausfOk() (*string, bool)`

GetUpuMacIausfOk returns a tuple with the UpuMacIausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuMacIausf

`func (o *UpuInfo) SetUpuMacIausf(v string)`

SetUpuMacIausf sets UpuMacIausf field to given value.

### HasUpuMacIausf

`func (o *UpuInfo) HasUpuMacIausf() bool`

HasUpuMacIausf returns a boolean if a field has been set.

### GetCounterUpu

`func (o *UpuInfo) GetCounterUpu() string`

GetCounterUpu returns the CounterUpu field if non-nil, zero value otherwise.

### GetCounterUpuOk

`func (o *UpuInfo) GetCounterUpuOk() (*string, bool)`

GetCounterUpuOk returns a tuple with the CounterUpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterUpu

`func (o *UpuInfo) SetCounterUpu(v string)`

SetCounterUpu sets CounterUpu field to given value.

### HasCounterUpu

`func (o *UpuInfo) HasCounterUpu() bool`

HasCounterUpu returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *UpuInfo) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *UpuInfo) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *UpuInfo) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


