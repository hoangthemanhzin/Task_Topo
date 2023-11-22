# NssaaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Status** | [**AuthStatus**](AuthStatus.md) |  | 

## Methods

### NewNssaaStatus

`func NewNssaaStatus(snssai Snssai, status AuthStatus, ) *NssaaStatus`

NewNssaaStatus instantiates a new NssaaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssaaStatusWithDefaults

`func NewNssaaStatusWithDefaults() *NssaaStatus`

NewNssaaStatusWithDefaults instantiates a new NssaaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *NssaaStatus) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NssaaStatus) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NssaaStatus) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetStatus

`func (o *NssaaStatus) GetStatus() AuthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NssaaStatus) GetStatusOk() (*AuthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NssaaStatus) SetStatus(v AuthStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


