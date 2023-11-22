# SorInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteeringContainer** | Pointer to [**SteeringContainer**](SteeringContainer.md) |  | [optional] 
**AckInd** | **bool** |  | 
**SorMacIausf** | Pointer to **string** |  | [optional] 
**Countersor** | Pointer to **string** |  | [optional] 
**ProvisioningTime** | **time.Time** |  | 

## Methods

### NewSorInfo1

`func NewSorInfo1(ackInd bool, provisioningTime time.Time, ) *SorInfo1`

NewSorInfo1 instantiates a new SorInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorInfo1WithDefaults

`func NewSorInfo1WithDefaults() *SorInfo1`

NewSorInfo1WithDefaults instantiates a new SorInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteeringContainer

`func (o *SorInfo1) GetSteeringContainer() SteeringContainer`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *SorInfo1) GetSteeringContainerOk() (*SteeringContainer, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *SorInfo1) SetSteeringContainer(v SteeringContainer)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *SorInfo1) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetAckInd

`func (o *SorInfo1) GetAckInd() bool`

GetAckInd returns the AckInd field if non-nil, zero value otherwise.

### GetAckIndOk

`func (o *SorInfo1) GetAckIndOk() (*bool, bool)`

GetAckIndOk returns a tuple with the AckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckInd

`func (o *SorInfo1) SetAckInd(v bool)`

SetAckInd sets AckInd field to given value.


### GetSorMacIausf

`func (o *SorInfo1) GetSorMacIausf() string`

GetSorMacIausf returns the SorMacIausf field if non-nil, zero value otherwise.

### GetSorMacIausfOk

`func (o *SorInfo1) GetSorMacIausfOk() (*string, bool)`

GetSorMacIausfOk returns a tuple with the SorMacIausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorMacIausf

`func (o *SorInfo1) SetSorMacIausf(v string)`

SetSorMacIausf sets SorMacIausf field to given value.

### HasSorMacIausf

`func (o *SorInfo1) HasSorMacIausf() bool`

HasSorMacIausf returns a boolean if a field has been set.

### GetCountersor

`func (o *SorInfo1) GetCountersor() string`

GetCountersor returns the Countersor field if non-nil, zero value otherwise.

### GetCountersorOk

`func (o *SorInfo1) GetCountersorOk() (*string, bool)`

GetCountersorOk returns a tuple with the Countersor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountersor

`func (o *SorInfo1) SetCountersor(v string)`

SetCountersor sets Countersor field to given value.

### HasCountersor

`func (o *SorInfo1) HasCountersor() bool`

HasCountersor returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *SorInfo1) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *SorInfo1) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *SorInfo1) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


