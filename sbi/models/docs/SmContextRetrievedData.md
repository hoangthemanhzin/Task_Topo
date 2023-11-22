# SmContextRetrievedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeEpsPdnConnection** | **string** |  | 
**SmContext** | Pointer to [**SmContext**](SmContext.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**DlDataWaitingInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSmContextRetrievedData

`func NewSmContextRetrievedData(ueEpsPdnConnection string, ) *SmContextRetrievedData`

NewSmContextRetrievedData instantiates a new SmContextRetrievedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextRetrievedDataWithDefaults

`func NewSmContextRetrievedDataWithDefaults() *SmContextRetrievedData`

NewSmContextRetrievedDataWithDefaults instantiates a new SmContextRetrievedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeEpsPdnConnection

`func (o *SmContextRetrievedData) GetUeEpsPdnConnection() string`

GetUeEpsPdnConnection returns the UeEpsPdnConnection field if non-nil, zero value otherwise.

### GetUeEpsPdnConnectionOk

`func (o *SmContextRetrievedData) GetUeEpsPdnConnectionOk() (*string, bool)`

GetUeEpsPdnConnectionOk returns a tuple with the UeEpsPdnConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeEpsPdnConnection

`func (o *SmContextRetrievedData) SetUeEpsPdnConnection(v string)`

SetUeEpsPdnConnection sets UeEpsPdnConnection field to given value.


### GetSmContext

`func (o *SmContextRetrievedData) GetSmContext() SmContext`

GetSmContext returns the SmContext field if non-nil, zero value otherwise.

### GetSmContextOk

`func (o *SmContextRetrievedData) GetSmContextOk() (*SmContext, bool)`

GetSmContextOk returns a tuple with the SmContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContext

`func (o *SmContextRetrievedData) SetSmContext(v SmContext)`

SetSmContext sets SmContext field to given value.

### HasSmContext

`func (o *SmContextRetrievedData) HasSmContext() bool`

HasSmContext returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *SmContextRetrievedData) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextRetrievedData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextRetrievedData) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextRetrievedData) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextRetrievedData) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextRetrievedData) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextRetrievedData) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextRetrievedData) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetDlDataWaitingInd

`func (o *SmContextRetrievedData) GetDlDataWaitingInd() bool`

GetDlDataWaitingInd returns the DlDataWaitingInd field if non-nil, zero value otherwise.

### GetDlDataWaitingIndOk

`func (o *SmContextRetrievedData) GetDlDataWaitingIndOk() (*bool, bool)`

GetDlDataWaitingIndOk returns a tuple with the DlDataWaitingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDataWaitingInd

`func (o *SmContextRetrievedData) SetDlDataWaitingInd(v bool)`

SetDlDataWaitingInd sets DlDataWaitingInd field to given value.

### HasDlDataWaitingInd

`func (o *SmContextRetrievedData) HasDlDataWaitingInd() bool`

HasDlDataWaitingInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


