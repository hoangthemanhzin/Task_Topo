# UeContextTransferReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**TransferReason**](TransferReason.md) |  | 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**RegRequest** | Pointer to [**N1MessageContainer**](N1MessageContainer.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewUeContextTransferReqData

`func NewUeContextTransferReqData(reason TransferReason, accessType AccessType, ) *UeContextTransferReqData`

NewUeContextTransferReqData instantiates a new UeContextTransferReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextTransferReqDataWithDefaults

`func NewUeContextTransferReqDataWithDefaults() *UeContextTransferReqData`

NewUeContextTransferReqDataWithDefaults instantiates a new UeContextTransferReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *UeContextTransferReqData) GetReason() TransferReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UeContextTransferReqData) GetReasonOk() (*TransferReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UeContextTransferReqData) SetReason(v TransferReason)`

SetReason sets Reason field to given value.


### GetAccessType

`func (o *UeContextTransferReqData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UeContextTransferReqData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UeContextTransferReqData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetPlmnId

`func (o *UeContextTransferReqData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *UeContextTransferReqData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *UeContextTransferReqData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *UeContextTransferReqData) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetRegRequest

`func (o *UeContextTransferReqData) GetRegRequest() N1MessageContainer`

GetRegRequest returns the RegRequest field if non-nil, zero value otherwise.

### GetRegRequestOk

`func (o *UeContextTransferReqData) GetRegRequestOk() (*N1MessageContainer, bool)`

GetRegRequestOk returns a tuple with the RegRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegRequest

`func (o *UeContextTransferReqData) SetRegRequest(v N1MessageContainer)`

SetRegRequest sets RegRequest field to given value.

### HasRegRequest

`func (o *UeContextTransferReqData) HasRegRequest() bool`

HasRegRequest returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *UeContextTransferReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *UeContextTransferReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *UeContextTransferReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *UeContextTransferReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


