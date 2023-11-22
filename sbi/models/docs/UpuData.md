# UpuData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningTime** | **time.Time** |  | 
**UeUpdateStatus** | [**UeUpdateStatus**](UeUpdateStatus.md) |  | 
**UpuXmacIue** | Pointer to **string** |  | [optional] 
**UpuMacIue** | Pointer to **string** |  | [optional] 

## Methods

### NewUpuData

`func NewUpuData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus, ) *UpuData`

NewUpuData instantiates a new UpuData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuDataWithDefaults

`func NewUpuDataWithDefaults() *UpuData`

NewUpuDataWithDefaults instantiates a new UpuData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningTime

`func (o *UpuData) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *UpuData) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *UpuData) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetUeUpdateStatus

`func (o *UpuData) GetUeUpdateStatus() UeUpdateStatus`

GetUeUpdateStatus returns the UeUpdateStatus field if non-nil, zero value otherwise.

### GetUeUpdateStatusOk

`func (o *UpuData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool)`

GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeUpdateStatus

`func (o *UpuData) SetUeUpdateStatus(v UeUpdateStatus)`

SetUeUpdateStatus sets UeUpdateStatus field to given value.


### GetUpuXmacIue

`func (o *UpuData) GetUpuXmacIue() string`

GetUpuXmacIue returns the UpuXmacIue field if non-nil, zero value otherwise.

### GetUpuXmacIueOk

`func (o *UpuData) GetUpuXmacIueOk() (*string, bool)`

GetUpuXmacIueOk returns a tuple with the UpuXmacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuXmacIue

`func (o *UpuData) SetUpuXmacIue(v string)`

SetUpuXmacIue sets UpuXmacIue field to given value.

### HasUpuXmacIue

`func (o *UpuData) HasUpuXmacIue() bool`

HasUpuXmacIue returns a boolean if a field has been set.

### GetUpuMacIue

`func (o *UpuData) GetUpuMacIue() string`

GetUpuMacIue returns the UpuMacIue field if non-nil, zero value otherwise.

### GetUpuMacIueOk

`func (o *UpuData) GetUpuMacIueOk() (*string, bool)`

GetUpuMacIueOk returns a tuple with the UpuMacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuMacIue

`func (o *UpuData) SetUpuMacIue(v string)`

SetUpuMacIue sets UpuMacIue field to given value.

### HasUpuMacIue

`func (o *UpuData) HasUpuMacIue() bool`

HasUpuMacIue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


