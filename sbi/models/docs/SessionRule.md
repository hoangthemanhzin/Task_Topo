# SessionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSessAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**AuthDefQos** | Pointer to [**AuthorizedDefaultQos**](AuthorizedDefaultQos.md) |  | [optional] 
**SessRuleId** | **string** | Univocally identifies the session rule within a PDU session. | 
**RefUmData** | Pointer to **string** | A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12. | [optional] 
**RefUmN3gData** | Pointer to **string** | A reference to UsageMonitoringData policy decision type to apply for Non-3GPP access. It is the umId described in subclause 5.6.2.12. | [optional] 
**RefCondData** | Pointer to **string** | A reference to the condition data. It is the condId described in subclause 5.6.2.9. | [optional] 

## Methods

### NewSessionRule

`func NewSessionRule(sessRuleId string, ) *SessionRule`

NewSessionRule instantiates a new SessionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRuleWithDefaults

`func NewSessionRuleWithDefaults() *SessionRule`

NewSessionRuleWithDefaults instantiates a new SessionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSessAmbr

`func (o *SessionRule) GetAuthSessAmbr() Ambr`

GetAuthSessAmbr returns the AuthSessAmbr field if non-nil, zero value otherwise.

### GetAuthSessAmbrOk

`func (o *SessionRule) GetAuthSessAmbrOk() (*Ambr, bool)`

GetAuthSessAmbrOk returns a tuple with the AuthSessAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessAmbr

`func (o *SessionRule) SetAuthSessAmbr(v Ambr)`

SetAuthSessAmbr sets AuthSessAmbr field to given value.

### HasAuthSessAmbr

`func (o *SessionRule) HasAuthSessAmbr() bool`

HasAuthSessAmbr returns a boolean if a field has been set.

### GetAuthDefQos

`func (o *SessionRule) GetAuthDefQos() AuthorizedDefaultQos`

GetAuthDefQos returns the AuthDefQos field if non-nil, zero value otherwise.

### GetAuthDefQosOk

`func (o *SessionRule) GetAuthDefQosOk() (*AuthorizedDefaultQos, bool)`

GetAuthDefQosOk returns a tuple with the AuthDefQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDefQos

`func (o *SessionRule) SetAuthDefQos(v AuthorizedDefaultQos)`

SetAuthDefQos sets AuthDefQos field to given value.

### HasAuthDefQos

`func (o *SessionRule) HasAuthDefQos() bool`

HasAuthDefQos returns a boolean if a field has been set.

### GetSessRuleId

`func (o *SessionRule) GetSessRuleId() string`

GetSessRuleId returns the SessRuleId field if non-nil, zero value otherwise.

### GetSessRuleIdOk

`func (o *SessionRule) GetSessRuleIdOk() (*string, bool)`

GetSessRuleIdOk returns a tuple with the SessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessRuleId

`func (o *SessionRule) SetSessRuleId(v string)`

SetSessRuleId sets SessRuleId field to given value.


### GetRefUmData

`func (o *SessionRule) GetRefUmData() string`

GetRefUmData returns the RefUmData field if non-nil, zero value otherwise.

### GetRefUmDataOk

`func (o *SessionRule) GetRefUmDataOk() (*string, bool)`

GetRefUmDataOk returns a tuple with the RefUmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmData

`func (o *SessionRule) SetRefUmData(v string)`

SetRefUmData sets RefUmData field to given value.

### HasRefUmData

`func (o *SessionRule) HasRefUmData() bool`

HasRefUmData returns a boolean if a field has been set.

### SetRefUmDataNil

`func (o *SessionRule) SetRefUmDataNil(b bool)`

 SetRefUmDataNil sets the value for RefUmData to be an explicit nil

### UnsetRefUmData
`func (o *SessionRule) UnsetRefUmData()`

UnsetRefUmData ensures that no value is present for RefUmData, not even an explicit nil
### GetRefUmN3gData

`func (o *SessionRule) GetRefUmN3gData() string`

GetRefUmN3gData returns the RefUmN3gData field if non-nil, zero value otherwise.

### GetRefUmN3gDataOk

`func (o *SessionRule) GetRefUmN3gDataOk() (*string, bool)`

GetRefUmN3gDataOk returns a tuple with the RefUmN3gData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmN3gData

`func (o *SessionRule) SetRefUmN3gData(v string)`

SetRefUmN3gData sets RefUmN3gData field to given value.

### HasRefUmN3gData

`func (o *SessionRule) HasRefUmN3gData() bool`

HasRefUmN3gData returns a boolean if a field has been set.

### SetRefUmN3gDataNil

`func (o *SessionRule) SetRefUmN3gDataNil(b bool)`

 SetRefUmN3gDataNil sets the value for RefUmN3gData to be an explicit nil

### UnsetRefUmN3gData
`func (o *SessionRule) UnsetRefUmN3gData()`

UnsetRefUmN3gData ensures that no value is present for RefUmN3gData, not even an explicit nil
### GetRefCondData

`func (o *SessionRule) GetRefCondData() string`

GetRefCondData returns the RefCondData field if non-nil, zero value otherwise.

### GetRefCondDataOk

`func (o *SessionRule) GetRefCondDataOk() (*string, bool)`

GetRefCondDataOk returns a tuple with the RefCondData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCondData

`func (o *SessionRule) SetRefCondData(v string)`

SetRefCondData sets RefCondData field to given value.

### HasRefCondData

`func (o *SessionRule) HasRefCondData() bool`

HasRefCondData returns a boolean if a field has been set.

### SetRefCondDataNil

`func (o *SessionRule) SetRefCondDataNil(b bool)`

 SetRefCondDataNil sets the value for RefCondData to be an explicit nil

### UnsetRefCondData
`func (o *SessionRule) UnsetRefCondData()`

UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


