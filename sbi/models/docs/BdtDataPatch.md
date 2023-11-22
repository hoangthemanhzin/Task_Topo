# BdtDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransPolicy** | Pointer to [**TransferPolicy**](TransferPolicy.md) |  | [optional] 
**BdtpStatus** | Pointer to [**BdtPolicyStatus**](BdtPolicyStatus.md) |  | [optional] 

## Methods

### NewBdtDataPatch

`func NewBdtDataPatch() *BdtDataPatch`

NewBdtDataPatch instantiates a new BdtDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtDataPatchWithDefaults

`func NewBdtDataPatchWithDefaults() *BdtDataPatch`

NewBdtDataPatchWithDefaults instantiates a new BdtDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransPolicy

`func (o *BdtDataPatch) GetTransPolicy() TransferPolicy`

GetTransPolicy returns the TransPolicy field if non-nil, zero value otherwise.

### GetTransPolicyOk

`func (o *BdtDataPatch) GetTransPolicyOk() (*TransferPolicy, bool)`

GetTransPolicyOk returns a tuple with the TransPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransPolicy

`func (o *BdtDataPatch) SetTransPolicy(v TransferPolicy)`

SetTransPolicy sets TransPolicy field to given value.

### HasTransPolicy

`func (o *BdtDataPatch) HasTransPolicy() bool`

HasTransPolicy returns a boolean if a field has been set.

### GetBdtpStatus

`func (o *BdtDataPatch) GetBdtpStatus() BdtPolicyStatus`

GetBdtpStatus returns the BdtpStatus field if non-nil, zero value otherwise.

### GetBdtpStatusOk

`func (o *BdtDataPatch) GetBdtpStatusOk() (*BdtPolicyStatus, bool)`

GetBdtpStatusOk returns a tuple with the BdtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtpStatus

`func (o *BdtDataPatch) SetBdtpStatus(v BdtPolicyStatus)`

SetBdtpStatus sets BdtpStatus field to given value.

### HasBdtpStatus

`func (o *BdtDataPatch) HasBdtpStatus() bool`

HasBdtpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


