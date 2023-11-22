# N1N2MessageTransferReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N1MessageContainer** | Pointer to [**N1MessageContainer**](N1MessageContainer.md) |  | [optional] 
**N2InfoContainer** | Pointer to [**N2InfoContainer**](N2InfoContainer.md) |  | [optional] 
**MtData** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**SkipInd** | Pointer to **bool** |  | [optional] [default to false]
**LastMsgIndication** | Pointer to **bool** |  | [optional] 
**PduSessionId** | Pointer to **int32** |  | [optional] 
**LcsCorrelationId** | Pointer to **string** |  | [optional] 
**Ppi** | Pointer to **int32** |  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**Var5qi** | Pointer to **int32** |  | [optional] 
**N1n2FailureTxfNotifURI** | Pointer to **string** |  | [optional] 
**SmfReallocationInd** | Pointer to **bool** |  | [optional] [default to false]
**AreaOfValidity** | Pointer to [**AreaOfValidity**](AreaOfValidity.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**MaAcceptedInd** | Pointer to **bool** |  | [optional] [default to false]
**ExtBufSupport** | Pointer to **bool** |  | [optional] [default to false]
**TargetAccess** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewN1N2MessageTransferReqData

`func NewN1N2MessageTransferReqData() *N1N2MessageTransferReqData`

NewN1N2MessageTransferReqData instantiates a new N1N2MessageTransferReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1N2MessageTransferReqDataWithDefaults

`func NewN1N2MessageTransferReqDataWithDefaults() *N1N2MessageTransferReqData`

NewN1N2MessageTransferReqDataWithDefaults instantiates a new N1N2MessageTransferReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN1MessageContainer

`func (o *N1N2MessageTransferReqData) GetN1MessageContainer() N1MessageContainer`

GetN1MessageContainer returns the N1MessageContainer field if non-nil, zero value otherwise.

### GetN1MessageContainerOk

`func (o *N1N2MessageTransferReqData) GetN1MessageContainerOk() (*N1MessageContainer, bool)`

GetN1MessageContainerOk returns a tuple with the N1MessageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageContainer

`func (o *N1N2MessageTransferReqData) SetN1MessageContainer(v N1MessageContainer)`

SetN1MessageContainer sets N1MessageContainer field to given value.

### HasN1MessageContainer

`func (o *N1N2MessageTransferReqData) HasN1MessageContainer() bool`

HasN1MessageContainer returns a boolean if a field has been set.

### GetN2InfoContainer

`func (o *N1N2MessageTransferReqData) GetN2InfoContainer() N2InfoContainer`

GetN2InfoContainer returns the N2InfoContainer field if non-nil, zero value otherwise.

### GetN2InfoContainerOk

`func (o *N1N2MessageTransferReqData) GetN2InfoContainerOk() (*N2InfoContainer, bool)`

GetN2InfoContainerOk returns a tuple with the N2InfoContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InfoContainer

`func (o *N1N2MessageTransferReqData) SetN2InfoContainer(v N2InfoContainer)`

SetN2InfoContainer sets N2InfoContainer field to given value.

### HasN2InfoContainer

`func (o *N1N2MessageTransferReqData) HasN2InfoContainer() bool`

HasN2InfoContainer returns a boolean if a field has been set.

### GetMtData

`func (o *N1N2MessageTransferReqData) GetMtData() RefToBinaryData`

GetMtData returns the MtData field if non-nil, zero value otherwise.

### GetMtDataOk

`func (o *N1N2MessageTransferReqData) GetMtDataOk() (*RefToBinaryData, bool)`

GetMtDataOk returns a tuple with the MtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtData

`func (o *N1N2MessageTransferReqData) SetMtData(v RefToBinaryData)`

SetMtData sets MtData field to given value.

### HasMtData

`func (o *N1N2MessageTransferReqData) HasMtData() bool`

HasMtData returns a boolean if a field has been set.

### GetSkipInd

`func (o *N1N2MessageTransferReqData) GetSkipInd() bool`

GetSkipInd returns the SkipInd field if non-nil, zero value otherwise.

### GetSkipIndOk

`func (o *N1N2MessageTransferReqData) GetSkipIndOk() (*bool, bool)`

GetSkipIndOk returns a tuple with the SkipInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipInd

`func (o *N1N2MessageTransferReqData) SetSkipInd(v bool)`

SetSkipInd sets SkipInd field to given value.

### HasSkipInd

`func (o *N1N2MessageTransferReqData) HasSkipInd() bool`

HasSkipInd returns a boolean if a field has been set.

### GetLastMsgIndication

`func (o *N1N2MessageTransferReqData) GetLastMsgIndication() bool`

GetLastMsgIndication returns the LastMsgIndication field if non-nil, zero value otherwise.

### GetLastMsgIndicationOk

`func (o *N1N2MessageTransferReqData) GetLastMsgIndicationOk() (*bool, bool)`

GetLastMsgIndicationOk returns a tuple with the LastMsgIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMsgIndication

`func (o *N1N2MessageTransferReqData) SetLastMsgIndication(v bool)`

SetLastMsgIndication sets LastMsgIndication field to given value.

### HasLastMsgIndication

`func (o *N1N2MessageTransferReqData) HasLastMsgIndication() bool`

HasLastMsgIndication returns a boolean if a field has been set.

### GetPduSessionId

`func (o *N1N2MessageTransferReqData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *N1N2MessageTransferReqData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *N1N2MessageTransferReqData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *N1N2MessageTransferReqData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetLcsCorrelationId

`func (o *N1N2MessageTransferReqData) GetLcsCorrelationId() string`

GetLcsCorrelationId returns the LcsCorrelationId field if non-nil, zero value otherwise.

### GetLcsCorrelationIdOk

`func (o *N1N2MessageTransferReqData) GetLcsCorrelationIdOk() (*string, bool)`

GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsCorrelationId

`func (o *N1N2MessageTransferReqData) SetLcsCorrelationId(v string)`

SetLcsCorrelationId sets LcsCorrelationId field to given value.

### HasLcsCorrelationId

`func (o *N1N2MessageTransferReqData) HasLcsCorrelationId() bool`

HasLcsCorrelationId returns a boolean if a field has been set.

### GetPpi

`func (o *N1N2MessageTransferReqData) GetPpi() int32`

GetPpi returns the Ppi field if non-nil, zero value otherwise.

### GetPpiOk

`func (o *N1N2MessageTransferReqData) GetPpiOk() (*int32, bool)`

GetPpiOk returns a tuple with the Ppi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpi

`func (o *N1N2MessageTransferReqData) SetPpi(v int32)`

SetPpi sets Ppi field to given value.

### HasPpi

`func (o *N1N2MessageTransferReqData) HasPpi() bool`

HasPpi returns a boolean if a field has been set.

### GetArp

`func (o *N1N2MessageTransferReqData) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *N1N2MessageTransferReqData) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *N1N2MessageTransferReqData) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *N1N2MessageTransferReqData) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetVar5qi

`func (o *N1N2MessageTransferReqData) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *N1N2MessageTransferReqData) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *N1N2MessageTransferReqData) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *N1N2MessageTransferReqData) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetN1n2FailureTxfNotifURI

`func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURI() string`

GetN1n2FailureTxfNotifURI returns the N1n2FailureTxfNotifURI field if non-nil, zero value otherwise.

### GetN1n2FailureTxfNotifURIOk

`func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURIOk() (*string, bool)`

GetN1n2FailureTxfNotifURIOk returns a tuple with the N1n2FailureTxfNotifURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1n2FailureTxfNotifURI

`func (o *N1N2MessageTransferReqData) SetN1n2FailureTxfNotifURI(v string)`

SetN1n2FailureTxfNotifURI sets N1n2FailureTxfNotifURI field to given value.

### HasN1n2FailureTxfNotifURI

`func (o *N1N2MessageTransferReqData) HasN1n2FailureTxfNotifURI() bool`

HasN1n2FailureTxfNotifURI returns a boolean if a field has been set.

### GetSmfReallocationInd

`func (o *N1N2MessageTransferReqData) GetSmfReallocationInd() bool`

GetSmfReallocationInd returns the SmfReallocationInd field if non-nil, zero value otherwise.

### GetSmfReallocationIndOk

`func (o *N1N2MessageTransferReqData) GetSmfReallocationIndOk() (*bool, bool)`

GetSmfReallocationIndOk returns a tuple with the SmfReallocationInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfReallocationInd

`func (o *N1N2MessageTransferReqData) SetSmfReallocationInd(v bool)`

SetSmfReallocationInd sets SmfReallocationInd field to given value.

### HasSmfReallocationInd

`func (o *N1N2MessageTransferReqData) HasSmfReallocationInd() bool`

HasSmfReallocationInd returns a boolean if a field has been set.

### GetAreaOfValidity

`func (o *N1N2MessageTransferReqData) GetAreaOfValidity() AreaOfValidity`

GetAreaOfValidity returns the AreaOfValidity field if non-nil, zero value otherwise.

### GetAreaOfValidityOk

`func (o *N1N2MessageTransferReqData) GetAreaOfValidityOk() (*AreaOfValidity, bool)`

GetAreaOfValidityOk returns a tuple with the AreaOfValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaOfValidity

`func (o *N1N2MessageTransferReqData) SetAreaOfValidity(v AreaOfValidity)`

SetAreaOfValidity sets AreaOfValidity field to given value.

### HasAreaOfValidity

`func (o *N1N2MessageTransferReqData) HasAreaOfValidity() bool`

HasAreaOfValidity returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *N1N2MessageTransferReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *N1N2MessageTransferReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *N1N2MessageTransferReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *N1N2MessageTransferReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *N1N2MessageTransferReqData) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *N1N2MessageTransferReqData) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *N1N2MessageTransferReqData) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *N1N2MessageTransferReqData) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.

### GetMaAcceptedInd

`func (o *N1N2MessageTransferReqData) GetMaAcceptedInd() bool`

GetMaAcceptedInd returns the MaAcceptedInd field if non-nil, zero value otherwise.

### GetMaAcceptedIndOk

`func (o *N1N2MessageTransferReqData) GetMaAcceptedIndOk() (*bool, bool)`

GetMaAcceptedIndOk returns a tuple with the MaAcceptedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaAcceptedInd

`func (o *N1N2MessageTransferReqData) SetMaAcceptedInd(v bool)`

SetMaAcceptedInd sets MaAcceptedInd field to given value.

### HasMaAcceptedInd

`func (o *N1N2MessageTransferReqData) HasMaAcceptedInd() bool`

HasMaAcceptedInd returns a boolean if a field has been set.

### GetExtBufSupport

`func (o *N1N2MessageTransferReqData) GetExtBufSupport() bool`

GetExtBufSupport returns the ExtBufSupport field if non-nil, zero value otherwise.

### GetExtBufSupportOk

`func (o *N1N2MessageTransferReqData) GetExtBufSupportOk() (*bool, bool)`

GetExtBufSupportOk returns a tuple with the ExtBufSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtBufSupport

`func (o *N1N2MessageTransferReqData) SetExtBufSupport(v bool)`

SetExtBufSupport sets ExtBufSupport field to given value.

### HasExtBufSupport

`func (o *N1N2MessageTransferReqData) HasExtBufSupport() bool`

HasExtBufSupport returns a boolean if a field has been set.

### GetTargetAccess

`func (o *N1N2MessageTransferReqData) GetTargetAccess() AccessType`

GetTargetAccess returns the TargetAccess field if non-nil, zero value otherwise.

### GetTargetAccessOk

`func (o *N1N2MessageTransferReqData) GetTargetAccessOk() (*AccessType, bool)`

GetTargetAccessOk returns a tuple with the TargetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccess

`func (o *N1N2MessageTransferReqData) SetTargetAccess(v AccessType)`

SetTargetAccess sets TargetAccess field to given value.

### HasTargetAccess

`func (o *N1N2MessageTransferReqData) HasTargetAccess() bool`

HasTargetAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


