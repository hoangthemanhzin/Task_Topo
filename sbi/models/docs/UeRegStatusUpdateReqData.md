# UeRegStatusUpdateReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferStatus** | [**UeContextTransferStatus**](UeContextTransferStatus.md) |  | 
**ToReleaseSessionList** | Pointer to **[]int32** |  | [optional] 
**PcfReselectedInd** | Pointer to **bool** |  | [optional] 
**SmfChangeInfoList** | Pointer to [**[]SmfChangeInfo**](SmfChangeInfo.md) |  | [optional] 

## Methods

### NewUeRegStatusUpdateReqData

`func NewUeRegStatusUpdateReqData(transferStatus UeContextTransferStatus, ) *UeRegStatusUpdateReqData`

NewUeRegStatusUpdateReqData instantiates a new UeRegStatusUpdateReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeRegStatusUpdateReqDataWithDefaults

`func NewUeRegStatusUpdateReqDataWithDefaults() *UeRegStatusUpdateReqData`

NewUeRegStatusUpdateReqDataWithDefaults instantiates a new UeRegStatusUpdateReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferStatus

`func (o *UeRegStatusUpdateReqData) GetTransferStatus() UeContextTransferStatus`

GetTransferStatus returns the TransferStatus field if non-nil, zero value otherwise.

### GetTransferStatusOk

`func (o *UeRegStatusUpdateReqData) GetTransferStatusOk() (*UeContextTransferStatus, bool)`

GetTransferStatusOk returns a tuple with the TransferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferStatus

`func (o *UeRegStatusUpdateReqData) SetTransferStatus(v UeContextTransferStatus)`

SetTransferStatus sets TransferStatus field to given value.


### GetToReleaseSessionList

`func (o *UeRegStatusUpdateReqData) GetToReleaseSessionList() []int32`

GetToReleaseSessionList returns the ToReleaseSessionList field if non-nil, zero value otherwise.

### GetToReleaseSessionListOk

`func (o *UeRegStatusUpdateReqData) GetToReleaseSessionListOk() (*[]int32, bool)`

GetToReleaseSessionListOk returns a tuple with the ToReleaseSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToReleaseSessionList

`func (o *UeRegStatusUpdateReqData) SetToReleaseSessionList(v []int32)`

SetToReleaseSessionList sets ToReleaseSessionList field to given value.

### HasToReleaseSessionList

`func (o *UeRegStatusUpdateReqData) HasToReleaseSessionList() bool`

HasToReleaseSessionList returns a boolean if a field has been set.

### GetPcfReselectedInd

`func (o *UeRegStatusUpdateReqData) GetPcfReselectedInd() bool`

GetPcfReselectedInd returns the PcfReselectedInd field if non-nil, zero value otherwise.

### GetPcfReselectedIndOk

`func (o *UeRegStatusUpdateReqData) GetPcfReselectedIndOk() (*bool, bool)`

GetPcfReselectedIndOk returns a tuple with the PcfReselectedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfReselectedInd

`func (o *UeRegStatusUpdateReqData) SetPcfReselectedInd(v bool)`

SetPcfReselectedInd sets PcfReselectedInd field to given value.

### HasPcfReselectedInd

`func (o *UeRegStatusUpdateReqData) HasPcfReselectedInd() bool`

HasPcfReselectedInd returns a boolean if a field has been set.

### GetSmfChangeInfoList

`func (o *UeRegStatusUpdateReqData) GetSmfChangeInfoList() []SmfChangeInfo`

GetSmfChangeInfoList returns the SmfChangeInfoList field if non-nil, zero value otherwise.

### GetSmfChangeInfoListOk

`func (o *UeRegStatusUpdateReqData) GetSmfChangeInfoListOk() (*[]SmfChangeInfo, bool)`

GetSmfChangeInfoListOk returns a tuple with the SmfChangeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfChangeInfoList

`func (o *UeRegStatusUpdateReqData) SetSmfChangeInfoList(v []SmfChangeInfo)`

SetSmfChangeInfoList sets SmfChangeInfoList field to given value.

### HasSmfChangeInfoList

`func (o *UeRegStatusUpdateReqData) HasSmfChangeInfoList() bool`

HasSmfChangeInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


