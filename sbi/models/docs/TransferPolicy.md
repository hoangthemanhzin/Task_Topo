# TransferPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBitRateDl** | Pointer to **string** |  | [optional] 
**MaxBitRateUl** | Pointer to **string** |  | [optional] 
**RatingGroup** | **int32** | Indicates a rating group for the recommended time window. | 
**RecTimeInt** | [**TimeWindow**](TimeWindow.md) |  | 
**TransPolicyId** | **int32** | Contains an identity of a transfer policy. | 

## Methods

### NewTransferPolicy

`func NewTransferPolicy(ratingGroup int32, recTimeInt TimeWindow, transPolicyId int32, ) *TransferPolicy`

NewTransferPolicy instantiates a new TransferPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferPolicyWithDefaults

`func NewTransferPolicyWithDefaults() *TransferPolicy`

NewTransferPolicyWithDefaults instantiates a new TransferPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBitRateDl

`func (o *TransferPolicy) GetMaxBitRateDl() string`

GetMaxBitRateDl returns the MaxBitRateDl field if non-nil, zero value otherwise.

### GetMaxBitRateDlOk

`func (o *TransferPolicy) GetMaxBitRateDlOk() (*string, bool)`

GetMaxBitRateDlOk returns a tuple with the MaxBitRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRateDl

`func (o *TransferPolicy) SetMaxBitRateDl(v string)`

SetMaxBitRateDl sets MaxBitRateDl field to given value.

### HasMaxBitRateDl

`func (o *TransferPolicy) HasMaxBitRateDl() bool`

HasMaxBitRateDl returns a boolean if a field has been set.

### GetMaxBitRateUl

`func (o *TransferPolicy) GetMaxBitRateUl() string`

GetMaxBitRateUl returns the MaxBitRateUl field if non-nil, zero value otherwise.

### GetMaxBitRateUlOk

`func (o *TransferPolicy) GetMaxBitRateUlOk() (*string, bool)`

GetMaxBitRateUlOk returns a tuple with the MaxBitRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRateUl

`func (o *TransferPolicy) SetMaxBitRateUl(v string)`

SetMaxBitRateUl sets MaxBitRateUl field to given value.

### HasMaxBitRateUl

`func (o *TransferPolicy) HasMaxBitRateUl() bool`

HasMaxBitRateUl returns a boolean if a field has been set.

### GetRatingGroup

`func (o *TransferPolicy) GetRatingGroup() int32`

GetRatingGroup returns the RatingGroup field if non-nil, zero value otherwise.

### GetRatingGroupOk

`func (o *TransferPolicy) GetRatingGroupOk() (*int32, bool)`

GetRatingGroupOk returns a tuple with the RatingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingGroup

`func (o *TransferPolicy) SetRatingGroup(v int32)`

SetRatingGroup sets RatingGroup field to given value.


### GetRecTimeInt

`func (o *TransferPolicy) GetRecTimeInt() TimeWindow`

GetRecTimeInt returns the RecTimeInt field if non-nil, zero value otherwise.

### GetRecTimeIntOk

`func (o *TransferPolicy) GetRecTimeIntOk() (*TimeWindow, bool)`

GetRecTimeIntOk returns a tuple with the RecTimeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecTimeInt

`func (o *TransferPolicy) SetRecTimeInt(v TimeWindow)`

SetRecTimeInt sets RecTimeInt field to given value.


### GetTransPolicyId

`func (o *TransferPolicy) GetTransPolicyId() int32`

GetTransPolicyId returns the TransPolicyId field if non-nil, zero value otherwise.

### GetTransPolicyIdOk

`func (o *TransferPolicy) GetTransPolicyIdOk() (*int32, bool)`

GetTransPolicyIdOk returns a tuple with the TransPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransPolicyId

`func (o *TransferPolicy) SetTransPolicyId(v int32)`

SetTransPolicyId sets TransPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


