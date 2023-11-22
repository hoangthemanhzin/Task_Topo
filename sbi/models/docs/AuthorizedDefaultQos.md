# AuthorizedDefaultQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | Pointer to **int32** |  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**PriorityLevel** | Pointer to **int32** |  | [optional] 
**AverWindow** | Pointer to **int32** |  | [optional] [default to 2000]
**MaxDataBurstVol** | Pointer to **int32** |  | [optional] 
**MaxbrUl** | Pointer to **string** |  | [optional] 
**MaxbrDl** | Pointer to **string** |  | [optional] 
**GbrUl** | Pointer to **string** |  | [optional] 
**GbrDl** | Pointer to **string** |  | [optional] 
**ExtMaxDataBurstVol** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthorizedDefaultQos

`func NewAuthorizedDefaultQos() *AuthorizedDefaultQos`

NewAuthorizedDefaultQos instantiates a new AuthorizedDefaultQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizedDefaultQosWithDefaults

`func NewAuthorizedDefaultQosWithDefaults() *AuthorizedDefaultQos`

NewAuthorizedDefaultQosWithDefaults instantiates a new AuthorizedDefaultQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *AuthorizedDefaultQos) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *AuthorizedDefaultQos) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *AuthorizedDefaultQos) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *AuthorizedDefaultQos) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetArp

`func (o *AuthorizedDefaultQos) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *AuthorizedDefaultQos) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *AuthorizedDefaultQos) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *AuthorizedDefaultQos) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetPriorityLevel

`func (o *AuthorizedDefaultQos) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *AuthorizedDefaultQos) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *AuthorizedDefaultQos) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *AuthorizedDefaultQos) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### SetPriorityLevelNil

`func (o *AuthorizedDefaultQos) SetPriorityLevelNil(b bool)`

 SetPriorityLevelNil sets the value for PriorityLevel to be an explicit nil

### UnsetPriorityLevel
`func (o *AuthorizedDefaultQos) UnsetPriorityLevel()`

UnsetPriorityLevel ensures that no value is present for PriorityLevel, not even an explicit nil
### GetAverWindow

`func (o *AuthorizedDefaultQos) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *AuthorizedDefaultQos) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *AuthorizedDefaultQos) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *AuthorizedDefaultQos) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### SetAverWindowNil

`func (o *AuthorizedDefaultQos) SetAverWindowNil(b bool)`

 SetAverWindowNil sets the value for AverWindow to be an explicit nil

### UnsetAverWindow
`func (o *AuthorizedDefaultQos) UnsetAverWindow()`

UnsetAverWindow ensures that no value is present for AverWindow, not even an explicit nil
### GetMaxDataBurstVol

`func (o *AuthorizedDefaultQos) GetMaxDataBurstVol() int32`

GetMaxDataBurstVol returns the MaxDataBurstVol field if non-nil, zero value otherwise.

### GetMaxDataBurstVolOk

`func (o *AuthorizedDefaultQos) GetMaxDataBurstVolOk() (*int32, bool)`

GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataBurstVol

`func (o *AuthorizedDefaultQos) SetMaxDataBurstVol(v int32)`

SetMaxDataBurstVol sets MaxDataBurstVol field to given value.

### HasMaxDataBurstVol

`func (o *AuthorizedDefaultQos) HasMaxDataBurstVol() bool`

HasMaxDataBurstVol returns a boolean if a field has been set.

### SetMaxDataBurstVolNil

`func (o *AuthorizedDefaultQos) SetMaxDataBurstVolNil(b bool)`

 SetMaxDataBurstVolNil sets the value for MaxDataBurstVol to be an explicit nil

### UnsetMaxDataBurstVol
`func (o *AuthorizedDefaultQos) UnsetMaxDataBurstVol()`

UnsetMaxDataBurstVol ensures that no value is present for MaxDataBurstVol, not even an explicit nil
### GetMaxbrUl

`func (o *AuthorizedDefaultQos) GetMaxbrUl() string`

GetMaxbrUl returns the MaxbrUl field if non-nil, zero value otherwise.

### GetMaxbrUlOk

`func (o *AuthorizedDefaultQos) GetMaxbrUlOk() (*string, bool)`

GetMaxbrUlOk returns a tuple with the MaxbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrUl

`func (o *AuthorizedDefaultQos) SetMaxbrUl(v string)`

SetMaxbrUl sets MaxbrUl field to given value.

### HasMaxbrUl

`func (o *AuthorizedDefaultQos) HasMaxbrUl() bool`

HasMaxbrUl returns a boolean if a field has been set.

### SetMaxbrUlNil

`func (o *AuthorizedDefaultQos) SetMaxbrUlNil(b bool)`

 SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil

### UnsetMaxbrUl
`func (o *AuthorizedDefaultQos) UnsetMaxbrUl()`

UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
### GetMaxbrDl

`func (o *AuthorizedDefaultQos) GetMaxbrDl() string`

GetMaxbrDl returns the MaxbrDl field if non-nil, zero value otherwise.

### GetMaxbrDlOk

`func (o *AuthorizedDefaultQos) GetMaxbrDlOk() (*string, bool)`

GetMaxbrDlOk returns a tuple with the MaxbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrDl

`func (o *AuthorizedDefaultQos) SetMaxbrDl(v string)`

SetMaxbrDl sets MaxbrDl field to given value.

### HasMaxbrDl

`func (o *AuthorizedDefaultQos) HasMaxbrDl() bool`

HasMaxbrDl returns a boolean if a field has been set.

### SetMaxbrDlNil

`func (o *AuthorizedDefaultQos) SetMaxbrDlNil(b bool)`

 SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil

### UnsetMaxbrDl
`func (o *AuthorizedDefaultQos) UnsetMaxbrDl()`

UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
### GetGbrUl

`func (o *AuthorizedDefaultQos) GetGbrUl() string`

GetGbrUl returns the GbrUl field if non-nil, zero value otherwise.

### GetGbrUlOk

`func (o *AuthorizedDefaultQos) GetGbrUlOk() (*string, bool)`

GetGbrUlOk returns a tuple with the GbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrUl

`func (o *AuthorizedDefaultQos) SetGbrUl(v string)`

SetGbrUl sets GbrUl field to given value.

### HasGbrUl

`func (o *AuthorizedDefaultQos) HasGbrUl() bool`

HasGbrUl returns a boolean if a field has been set.

### SetGbrUlNil

`func (o *AuthorizedDefaultQos) SetGbrUlNil(b bool)`

 SetGbrUlNil sets the value for GbrUl to be an explicit nil

### UnsetGbrUl
`func (o *AuthorizedDefaultQos) UnsetGbrUl()`

UnsetGbrUl ensures that no value is present for GbrUl, not even an explicit nil
### GetGbrDl

`func (o *AuthorizedDefaultQos) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *AuthorizedDefaultQos) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *AuthorizedDefaultQos) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *AuthorizedDefaultQos) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### SetGbrDlNil

`func (o *AuthorizedDefaultQos) SetGbrDlNil(b bool)`

 SetGbrDlNil sets the value for GbrDl to be an explicit nil

### UnsetGbrDl
`func (o *AuthorizedDefaultQos) UnsetGbrDl()`

UnsetGbrDl ensures that no value is present for GbrDl, not even an explicit nil
### GetExtMaxDataBurstVol

`func (o *AuthorizedDefaultQos) GetExtMaxDataBurstVol() int32`

GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field if non-nil, zero value otherwise.

### GetExtMaxDataBurstVolOk

`func (o *AuthorizedDefaultQos) GetExtMaxDataBurstVolOk() (*int32, bool)`

GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMaxDataBurstVol

`func (o *AuthorizedDefaultQos) SetExtMaxDataBurstVol(v int32)`

SetExtMaxDataBurstVol sets ExtMaxDataBurstVol field to given value.

### HasExtMaxDataBurstVol

`func (o *AuthorizedDefaultQos) HasExtMaxDataBurstVol() bool`

HasExtMaxDataBurstVol returns a boolean if a field has been set.

### SetExtMaxDataBurstVolNil

`func (o *AuthorizedDefaultQos) SetExtMaxDataBurstVolNil(b bool)`

 SetExtMaxDataBurstVolNil sets the value for ExtMaxDataBurstVol to be an explicit nil

### UnsetExtMaxDataBurstVol
`func (o *AuthorizedDefaultQos) UnsetExtMaxDataBurstVol()`

UnsetExtMaxDataBurstVol ensures that no value is present for ExtMaxDataBurstVol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


