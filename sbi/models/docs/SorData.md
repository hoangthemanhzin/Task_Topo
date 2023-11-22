# SorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningTime** | **time.Time** |  | 
**UeUpdateStatus** | [**UeUpdateStatus**](UeUpdateStatus.md) |  | 
**SorXmacIue** | Pointer to **string** |  | [optional] 
**SorMacIue** | Pointer to **string** |  | [optional] 

## Methods

### NewSorData

`func NewSorData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus, ) *SorData`

NewSorData instantiates a new SorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorDataWithDefaults

`func NewSorDataWithDefaults() *SorData`

NewSorDataWithDefaults instantiates a new SorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningTime

`func (o *SorData) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *SorData) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *SorData) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetUeUpdateStatus

`func (o *SorData) GetUeUpdateStatus() UeUpdateStatus`

GetUeUpdateStatus returns the UeUpdateStatus field if non-nil, zero value otherwise.

### GetUeUpdateStatusOk

`func (o *SorData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool)`

GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeUpdateStatus

`func (o *SorData) SetUeUpdateStatus(v UeUpdateStatus)`

SetUeUpdateStatus sets UeUpdateStatus field to given value.


### GetSorXmacIue

`func (o *SorData) GetSorXmacIue() string`

GetSorXmacIue returns the SorXmacIue field if non-nil, zero value otherwise.

### GetSorXmacIueOk

`func (o *SorData) GetSorXmacIueOk() (*string, bool)`

GetSorXmacIueOk returns a tuple with the SorXmacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorXmacIue

`func (o *SorData) SetSorXmacIue(v string)`

SetSorXmacIue sets SorXmacIue field to given value.

### HasSorXmacIue

`func (o *SorData) HasSorXmacIue() bool`

HasSorXmacIue returns a boolean if a field has been set.

### GetSorMacIue

`func (o *SorData) GetSorMacIue() string`

GetSorMacIue returns the SorMacIue field if non-nil, zero value otherwise.

### GetSorMacIueOk

`func (o *SorData) GetSorMacIueOk() (*string, bool)`

GetSorMacIueOk returns a tuple with the SorMacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorMacIue

`func (o *SorData) SetSorMacIue(v string)`

SetSorMacIue sets SorMacIue field to given value.

### HasSorMacIue

`func (o *SorData) HasSorMacIue() bool`

HasSorMacIue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


