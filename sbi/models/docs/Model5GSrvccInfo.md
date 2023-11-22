# Model5GSrvccInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ue5GSrvccCapability** | **bool** |  | 
**StnSr** | Pointer to **string** |  | [optional] 
**CMsisdn** | Pointer to **string** |  | [optional] 

## Methods

### NewModel5GSrvccInfo

`func NewModel5GSrvccInfo(ue5GSrvccCapability bool, ) *Model5GSrvccInfo`

NewModel5GSrvccInfo instantiates a new Model5GSrvccInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GSrvccInfoWithDefaults

`func NewModel5GSrvccInfoWithDefaults() *Model5GSrvccInfo`

NewModel5GSrvccInfoWithDefaults instantiates a new Model5GSrvccInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUe5GSrvccCapability

`func (o *Model5GSrvccInfo) GetUe5GSrvccCapability() bool`

GetUe5GSrvccCapability returns the Ue5GSrvccCapability field if non-nil, zero value otherwise.

### GetUe5GSrvccCapabilityOk

`func (o *Model5GSrvccInfo) GetUe5GSrvccCapabilityOk() (*bool, bool)`

GetUe5GSrvccCapabilityOk returns a tuple with the Ue5GSrvccCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUe5GSrvccCapability

`func (o *Model5GSrvccInfo) SetUe5GSrvccCapability(v bool)`

SetUe5GSrvccCapability sets Ue5GSrvccCapability field to given value.


### GetStnSr

`func (o *Model5GSrvccInfo) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *Model5GSrvccInfo) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *Model5GSrvccInfo) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *Model5GSrvccInfo) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### GetCMsisdn

`func (o *Model5GSrvccInfo) GetCMsisdn() string`

GetCMsisdn returns the CMsisdn field if non-nil, zero value otherwise.

### GetCMsisdnOk

`func (o *Model5GSrvccInfo) GetCMsisdnOk() (*string, bool)`

GetCMsisdnOk returns a tuple with the CMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCMsisdn

`func (o *Model5GSrvccInfo) SetCMsisdn(v string)`

SetCMsisdn sets CMsisdn field to given value.

### HasCMsisdn

`func (o *Model5GSrvccInfo) HasCMsisdn() bool`

HasCMsisdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


