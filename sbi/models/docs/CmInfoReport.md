# CmInfoReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldCmInfoList** | Pointer to [**[]CmInfo**](CmInfo.md) |  | [optional] 
**NewCmInfoList** | [**[]CmInfo**](CmInfo.md) |  | 

## Methods

### NewCmInfoReport

`func NewCmInfoReport(newCmInfoList []CmInfo, ) *CmInfoReport`

NewCmInfoReport instantiates a new CmInfoReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmInfoReportWithDefaults

`func NewCmInfoReportWithDefaults() *CmInfoReport`

NewCmInfoReportWithDefaults instantiates a new CmInfoReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldCmInfoList

`func (o *CmInfoReport) GetOldCmInfoList() []CmInfo`

GetOldCmInfoList returns the OldCmInfoList field if non-nil, zero value otherwise.

### GetOldCmInfoListOk

`func (o *CmInfoReport) GetOldCmInfoListOk() (*[]CmInfo, bool)`

GetOldCmInfoListOk returns a tuple with the OldCmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCmInfoList

`func (o *CmInfoReport) SetOldCmInfoList(v []CmInfo)`

SetOldCmInfoList sets OldCmInfoList field to given value.

### HasOldCmInfoList

`func (o *CmInfoReport) HasOldCmInfoList() bool`

HasOldCmInfoList returns a boolean if a field has been set.

### GetNewCmInfoList

`func (o *CmInfoReport) GetNewCmInfoList() []CmInfo`

GetNewCmInfoList returns the NewCmInfoList field if non-nil, zero value otherwise.

### GetNewCmInfoListOk

`func (o *CmInfoReport) GetNewCmInfoListOk() (*[]CmInfo, bool)`

GetNewCmInfoListOk returns a tuple with the NewCmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCmInfoList

`func (o *CmInfoReport) SetNewCmInfoList(v []CmInfo)`

SetNewCmInfoList sets NewCmInfoList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


