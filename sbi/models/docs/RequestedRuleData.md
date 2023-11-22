# RequestedRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefPccRuleIds** | **[]string** | An array of PCC rule id references to the PCC rules associated with the control data. | 
**ReqData** | [**[]RequestedRuleDataType**](RequestedRuleDataType.md) | Array of requested rule data type elements indicating what type of rule data is requested for the corresponding referenced PCC rules. | 

## Methods

### NewRequestedRuleData

`func NewRequestedRuleData(refPccRuleIds []string, reqData []RequestedRuleDataType, ) *RequestedRuleData`

NewRequestedRuleData instantiates a new RequestedRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedRuleDataWithDefaults

`func NewRequestedRuleDataWithDefaults() *RequestedRuleData`

NewRequestedRuleDataWithDefaults instantiates a new RequestedRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefPccRuleIds

`func (o *RequestedRuleData) GetRefPccRuleIds() []string`

GetRefPccRuleIds returns the RefPccRuleIds field if non-nil, zero value otherwise.

### GetRefPccRuleIdsOk

`func (o *RequestedRuleData) GetRefPccRuleIdsOk() (*[]string, bool)`

GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPccRuleIds

`func (o *RequestedRuleData) SetRefPccRuleIds(v []string)`

SetRefPccRuleIds sets RefPccRuleIds field to given value.


### GetReqData

`func (o *RequestedRuleData) GetReqData() []RequestedRuleDataType`

GetReqData returns the ReqData field if non-nil, zero value otherwise.

### GetReqDataOk

`func (o *RequestedRuleData) GetReqDataOk() (*[]RequestedRuleDataType, bool)`

GetReqDataOk returns a tuple with the ReqData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqData

`func (o *RequestedRuleData) SetReqData(v []RequestedRuleDataType)`

SetReqData sets ReqData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


