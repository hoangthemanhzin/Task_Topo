# LocationQoS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HAccuracy** | Pointer to **float32** |  | [optional] 
**VAccuracy** | Pointer to **float32** |  | [optional] 
**VerticalRequested** | Pointer to **bool** |  | [optional] 
**ResponseTime** | Pointer to [**ResponseTime**](ResponseTime.md) |  | [optional] 
**LcsQosClass** | Pointer to [**LcsQosClass**](LcsQosClass.md) |  | [optional] 

## Methods

### NewLocationQoS

`func NewLocationQoS() *LocationQoS`

NewLocationQoS instantiates a new LocationQoS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationQoSWithDefaults

`func NewLocationQoSWithDefaults() *LocationQoS`

NewLocationQoSWithDefaults instantiates a new LocationQoS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHAccuracy

`func (o *LocationQoS) GetHAccuracy() float32`

GetHAccuracy returns the HAccuracy field if non-nil, zero value otherwise.

### GetHAccuracyOk

`func (o *LocationQoS) GetHAccuracyOk() (*float32, bool)`

GetHAccuracyOk returns a tuple with the HAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHAccuracy

`func (o *LocationQoS) SetHAccuracy(v float32)`

SetHAccuracy sets HAccuracy field to given value.

### HasHAccuracy

`func (o *LocationQoS) HasHAccuracy() bool`

HasHAccuracy returns a boolean if a field has been set.

### GetVAccuracy

`func (o *LocationQoS) GetVAccuracy() float32`

GetVAccuracy returns the VAccuracy field if non-nil, zero value otherwise.

### GetVAccuracyOk

`func (o *LocationQoS) GetVAccuracyOk() (*float32, bool)`

GetVAccuracyOk returns a tuple with the VAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAccuracy

`func (o *LocationQoS) SetVAccuracy(v float32)`

SetVAccuracy sets VAccuracy field to given value.

### HasVAccuracy

`func (o *LocationQoS) HasVAccuracy() bool`

HasVAccuracy returns a boolean if a field has been set.

### GetVerticalRequested

`func (o *LocationQoS) GetVerticalRequested() bool`

GetVerticalRequested returns the VerticalRequested field if non-nil, zero value otherwise.

### GetVerticalRequestedOk

`func (o *LocationQoS) GetVerticalRequestedOk() (*bool, bool)`

GetVerticalRequestedOk returns a tuple with the VerticalRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalRequested

`func (o *LocationQoS) SetVerticalRequested(v bool)`

SetVerticalRequested sets VerticalRequested field to given value.

### HasVerticalRequested

`func (o *LocationQoS) HasVerticalRequested() bool`

HasVerticalRequested returns a boolean if a field has been set.

### GetResponseTime

`func (o *LocationQoS) GetResponseTime() ResponseTime`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *LocationQoS) GetResponseTimeOk() (*ResponseTime, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *LocationQoS) SetResponseTime(v ResponseTime)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *LocationQoS) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetLcsQosClass

`func (o *LocationQoS) GetLcsQosClass() LcsQosClass`

GetLcsQosClass returns the LcsQosClass field if non-nil, zero value otherwise.

### GetLcsQosClassOk

`func (o *LocationQoS) GetLcsQosClassOk() (*LcsQosClass, bool)`

GetLcsQosClassOk returns a tuple with the LcsQosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQosClass

`func (o *LocationQoS) SetLcsQosClass(v LcsQosClass)`

SetLcsQosClass sets LcsQosClass field to given value.

### HasLcsQosClass

`func (o *LocationQoS) HasLcsQosClass() bool`

HasLcsQosClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


