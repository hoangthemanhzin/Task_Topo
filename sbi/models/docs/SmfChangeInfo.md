# SmfChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionIdList** | **[]int32** |  | 
**SmfChangeInd** | [**SmfChangeIndication**](SmfChangeIndication.md) |  | 

## Methods

### NewSmfChangeInfo

`func NewSmfChangeInfo(pduSessionIdList []int32, smfChangeInd SmfChangeIndication, ) *SmfChangeInfo`

NewSmfChangeInfo instantiates a new SmfChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfChangeInfoWithDefaults

`func NewSmfChangeInfoWithDefaults() *SmfChangeInfo`

NewSmfChangeInfoWithDefaults instantiates a new SmfChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionIdList

`func (o *SmfChangeInfo) GetPduSessionIdList() []int32`

GetPduSessionIdList returns the PduSessionIdList field if non-nil, zero value otherwise.

### GetPduSessionIdListOk

`func (o *SmfChangeInfo) GetPduSessionIdListOk() (*[]int32, bool)`

GetPduSessionIdListOk returns a tuple with the PduSessionIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionIdList

`func (o *SmfChangeInfo) SetPduSessionIdList(v []int32)`

SetPduSessionIdList sets PduSessionIdList field to given value.


### GetSmfChangeInd

`func (o *SmfChangeInfo) GetSmfChangeInd() SmfChangeIndication`

GetSmfChangeInd returns the SmfChangeInd field if non-nil, zero value otherwise.

### GetSmfChangeIndOk

`func (o *SmfChangeInfo) GetSmfChangeIndOk() (*SmfChangeIndication, bool)`

GetSmfChangeIndOk returns a tuple with the SmfChangeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfChangeInd

`func (o *SmfChangeInfo) SetSmfChangeInd(v SmfChangeIndication)`

SetSmfChangeInd sets SmfChangeInd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


