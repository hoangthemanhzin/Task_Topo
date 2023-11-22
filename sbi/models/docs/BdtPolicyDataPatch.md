# BdtPolicyDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelTransPolicyId** | **int32** | Contains an identity (i.e. transPolicyId value) of the selected transfer policy. If the BdtNotification_5G feature is supported value 0 indicates that no transfer policy is selected. | 

## Methods

### NewBdtPolicyDataPatch

`func NewBdtPolicyDataPatch(selTransPolicyId int32, ) *BdtPolicyDataPatch`

NewBdtPolicyDataPatch instantiates a new BdtPolicyDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtPolicyDataPatchWithDefaults

`func NewBdtPolicyDataPatchWithDefaults() *BdtPolicyDataPatch`

NewBdtPolicyDataPatchWithDefaults instantiates a new BdtPolicyDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelTransPolicyId

`func (o *BdtPolicyDataPatch) GetSelTransPolicyId() int32`

GetSelTransPolicyId returns the SelTransPolicyId field if non-nil, zero value otherwise.

### GetSelTransPolicyIdOk

`func (o *BdtPolicyDataPatch) GetSelTransPolicyIdOk() (*int32, bool)`

GetSelTransPolicyIdOk returns a tuple with the SelTransPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelTransPolicyId

`func (o *BdtPolicyDataPatch) SetSelTransPolicyId(v int32)`

SetSelTransPolicyId sets SelTransPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


