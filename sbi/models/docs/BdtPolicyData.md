# BdtPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BdtRefId** | **string** | string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154. | 
**TransfPolicies** | [**[]TransferPolicy**](TransferPolicy.md) | Contains transfer policies. | 
**SelTransPolicyId** | Pointer to **int32** | Contains an identity of the selected transfer policy. | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewBdtPolicyData

`func NewBdtPolicyData(bdtRefId string, transfPolicies []TransferPolicy, ) *BdtPolicyData`

NewBdtPolicyData instantiates a new BdtPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtPolicyDataWithDefaults

`func NewBdtPolicyDataWithDefaults() *BdtPolicyData`

NewBdtPolicyDataWithDefaults instantiates a new BdtPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBdtRefId

`func (o *BdtPolicyData) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *BdtPolicyData) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *BdtPolicyData) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.


### GetTransfPolicies

`func (o *BdtPolicyData) GetTransfPolicies() []TransferPolicy`

GetTransfPolicies returns the TransfPolicies field if non-nil, zero value otherwise.

### GetTransfPoliciesOk

`func (o *BdtPolicyData) GetTransfPoliciesOk() (*[]TransferPolicy, bool)`

GetTransfPoliciesOk returns a tuple with the TransfPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfPolicies

`func (o *BdtPolicyData) SetTransfPolicies(v []TransferPolicy)`

SetTransfPolicies sets TransfPolicies field to given value.


### GetSelTransPolicyId

`func (o *BdtPolicyData) GetSelTransPolicyId() int32`

GetSelTransPolicyId returns the SelTransPolicyId field if non-nil, zero value otherwise.

### GetSelTransPolicyIdOk

`func (o *BdtPolicyData) GetSelTransPolicyIdOk() (*int32, bool)`

GetSelTransPolicyIdOk returns a tuple with the SelTransPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelTransPolicyId

`func (o *BdtPolicyData) SetSelTransPolicyId(v int32)`

SetSelTransPolicyId sets SelTransPolicyId field to given value.

### HasSelTransPolicyId

`func (o *BdtPolicyData) HasSelTransPolicyId() bool`

HasSelTransPolicyId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *BdtPolicyData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *BdtPolicyData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *BdtPolicyData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *BdtPolicyData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


