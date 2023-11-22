# UpSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpIntegr** | [**UpIntegrity**](UpIntegrity.md) |  | 
**UpConfid** | [**UpConfidentiality**](UpConfidentiality.md) |  | 

## Methods

### NewUpSecurity

`func NewUpSecurity(upIntegr UpIntegrity, upConfid UpConfidentiality, ) *UpSecurity`

NewUpSecurity instantiates a new UpSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpSecurityWithDefaults

`func NewUpSecurityWithDefaults() *UpSecurity`

NewUpSecurityWithDefaults instantiates a new UpSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpIntegr

`func (o *UpSecurity) GetUpIntegr() UpIntegrity`

GetUpIntegr returns the UpIntegr field if non-nil, zero value otherwise.

### GetUpIntegrOk

`func (o *UpSecurity) GetUpIntegrOk() (*UpIntegrity, bool)`

GetUpIntegrOk returns a tuple with the UpIntegr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpIntegr

`func (o *UpSecurity) SetUpIntegr(v UpIntegrity)`

SetUpIntegr sets UpIntegr field to given value.


### GetUpConfid

`func (o *UpSecurity) GetUpConfid() UpConfidentiality`

GetUpConfid returns the UpConfid field if non-nil, zero value otherwise.

### GetUpConfidOk

`func (o *UpSecurity) GetUpConfidOk() (*UpConfidentiality, bool)`

GetUpConfidOk returns a tuple with the UpConfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpConfid

`func (o *UpSecurity) SetUpConfid(v UpConfidentiality)`

SetUpConfid sets UpConfid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


