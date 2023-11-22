# CagAckData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningTime** | **time.Time** |  | 
**UeUpdateStatus** | [**UeUpdateStatus**](UeUpdateStatus.md) |  | 

## Methods

### NewCagAckData

`func NewCagAckData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus, ) *CagAckData`

NewCagAckData instantiates a new CagAckData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCagAckDataWithDefaults

`func NewCagAckDataWithDefaults() *CagAckData`

NewCagAckDataWithDefaults instantiates a new CagAckData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningTime

`func (o *CagAckData) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *CagAckData) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *CagAckData) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetUeUpdateStatus

`func (o *CagAckData) GetUeUpdateStatus() UeUpdateStatus`

GetUeUpdateStatus returns the UeUpdateStatus field if non-nil, zero value otherwise.

### GetUeUpdateStatusOk

`func (o *CagAckData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool)`

GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeUpdateStatus

`func (o *CagAckData) SetUeUpdateStatus(v UeUpdateStatus)`

SetUeUpdateStatus sets UeUpdateStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


