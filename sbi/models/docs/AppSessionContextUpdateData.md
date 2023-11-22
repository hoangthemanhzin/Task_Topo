# AppSessionContextUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**AfRoutReq** | Pointer to [**AfRoutingRequirementRm**](AfRoutingRequirementRm.md) |  | [optional] 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 
**BdtRefId** | Pointer to **string** | string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154. | [optional] 
**EvSubsc** | Pointer to [**EventsSubscReqDataRm**](EventsSubscReqDataRm.md) |  | [optional] 
**McpttId** | Pointer to **string** | indication of MCPTT service request | [optional] 
**McVideoId** | Pointer to **string** | indication of modification of MCVideo service | [optional] 
**MedComponents** | Pointer to [**map[string]MediaComponentRm**](MediaComponentRm.md) |  | [optional] 
**MpsId** | Pointer to **string** | indication of MPS service request | [optional] 
**McsId** | Pointer to **string** | indication of MCS service request | [optional] 
**PreemptControlInfo** | Pointer to [**PreemptionControlInformationRm**](PreemptionControlInformationRm.md) |  | [optional] 
**ResPrio** | Pointer to [**ReservPriority**](ReservPriority.md) |  | [optional] 
**ServInfStatus** | Pointer to [**ServiceInfoStatus**](ServiceInfoStatus.md) |  | [optional] 
**SipForkInd** | Pointer to [**SipForkingIndication**](SipForkingIndication.md) |  | [optional] 
**SponId** | Pointer to **string** | Contains an identity of a sponsor. | [optional] 
**SponStatus** | Pointer to [**SponsoringStatus**](SponsoringStatus.md) |  | [optional] 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 

## Methods

### NewAppSessionContextUpdateData

`func NewAppSessionContextUpdateData() *AppSessionContextUpdateData`

NewAppSessionContextUpdateData instantiates a new AppSessionContextUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSessionContextUpdateDataWithDefaults

`func NewAppSessionContextUpdateDataWithDefaults() *AppSessionContextUpdateData`

NewAppSessionContextUpdateDataWithDefaults instantiates a new AppSessionContextUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfAppId

`func (o *AppSessionContextUpdateData) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *AppSessionContextUpdateData) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *AppSessionContextUpdateData) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *AppSessionContextUpdateData) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetAfRoutReq

`func (o *AppSessionContextUpdateData) GetAfRoutReq() AfRoutingRequirementRm`

GetAfRoutReq returns the AfRoutReq field if non-nil, zero value otherwise.

### GetAfRoutReqOk

`func (o *AppSessionContextUpdateData) GetAfRoutReqOk() (*AfRoutingRequirementRm, bool)`

GetAfRoutReqOk returns a tuple with the AfRoutReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfRoutReq

`func (o *AppSessionContextUpdateData) SetAfRoutReq(v AfRoutingRequirementRm)`

SetAfRoutReq sets AfRoutReq field to given value.

### HasAfRoutReq

`func (o *AppSessionContextUpdateData) HasAfRoutReq() bool`

HasAfRoutReq returns a boolean if a field has been set.

### SetAfRoutReqNil

`func (o *AppSessionContextUpdateData) SetAfRoutReqNil(b bool)`

 SetAfRoutReqNil sets the value for AfRoutReq to be an explicit nil

### UnsetAfRoutReq
`func (o *AppSessionContextUpdateData) UnsetAfRoutReq()`

UnsetAfRoutReq ensures that no value is present for AfRoutReq, not even an explicit nil
### GetAspId

`func (o *AppSessionContextUpdateData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *AppSessionContextUpdateData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *AppSessionContextUpdateData) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *AppSessionContextUpdateData) HasAspId() bool`

HasAspId returns a boolean if a field has been set.

### GetBdtRefId

`func (o *AppSessionContextUpdateData) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *AppSessionContextUpdateData) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *AppSessionContextUpdateData) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.

### HasBdtRefId

`func (o *AppSessionContextUpdateData) HasBdtRefId() bool`

HasBdtRefId returns a boolean if a field has been set.

### GetEvSubsc

`func (o *AppSessionContextUpdateData) GetEvSubsc() EventsSubscReqDataRm`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *AppSessionContextUpdateData) GetEvSubscOk() (*EventsSubscReqDataRm, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *AppSessionContextUpdateData) SetEvSubsc(v EventsSubscReqDataRm)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *AppSessionContextUpdateData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### SetEvSubscNil

`func (o *AppSessionContextUpdateData) SetEvSubscNil(b bool)`

 SetEvSubscNil sets the value for EvSubsc to be an explicit nil

### UnsetEvSubsc
`func (o *AppSessionContextUpdateData) UnsetEvSubsc()`

UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil
### GetMcpttId

`func (o *AppSessionContextUpdateData) GetMcpttId() string`

GetMcpttId returns the McpttId field if non-nil, zero value otherwise.

### GetMcpttIdOk

`func (o *AppSessionContextUpdateData) GetMcpttIdOk() (*string, bool)`

GetMcpttIdOk returns a tuple with the McpttId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpttId

`func (o *AppSessionContextUpdateData) SetMcpttId(v string)`

SetMcpttId sets McpttId field to given value.

### HasMcpttId

`func (o *AppSessionContextUpdateData) HasMcpttId() bool`

HasMcpttId returns a boolean if a field has been set.

### GetMcVideoId

`func (o *AppSessionContextUpdateData) GetMcVideoId() string`

GetMcVideoId returns the McVideoId field if non-nil, zero value otherwise.

### GetMcVideoIdOk

`func (o *AppSessionContextUpdateData) GetMcVideoIdOk() (*string, bool)`

GetMcVideoIdOk returns a tuple with the McVideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcVideoId

`func (o *AppSessionContextUpdateData) SetMcVideoId(v string)`

SetMcVideoId sets McVideoId field to given value.

### HasMcVideoId

`func (o *AppSessionContextUpdateData) HasMcVideoId() bool`

HasMcVideoId returns a boolean if a field has been set.

### GetMedComponents

`func (o *AppSessionContextUpdateData) GetMedComponents() map[string]MediaComponentRm`

GetMedComponents returns the MedComponents field if non-nil, zero value otherwise.

### GetMedComponentsOk

`func (o *AppSessionContextUpdateData) GetMedComponentsOk() (*map[string]MediaComponentRm, bool)`

GetMedComponentsOk returns a tuple with the MedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedComponents

`func (o *AppSessionContextUpdateData) SetMedComponents(v map[string]MediaComponentRm)`

SetMedComponents sets MedComponents field to given value.

### HasMedComponents

`func (o *AppSessionContextUpdateData) HasMedComponents() bool`

HasMedComponents returns a boolean if a field has been set.

### GetMpsId

`func (o *AppSessionContextUpdateData) GetMpsId() string`

GetMpsId returns the MpsId field if non-nil, zero value otherwise.

### GetMpsIdOk

`func (o *AppSessionContextUpdateData) GetMpsIdOk() (*string, bool)`

GetMpsIdOk returns a tuple with the MpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsId

`func (o *AppSessionContextUpdateData) SetMpsId(v string)`

SetMpsId sets MpsId field to given value.

### HasMpsId

`func (o *AppSessionContextUpdateData) HasMpsId() bool`

HasMpsId returns a boolean if a field has been set.

### GetMcsId

`func (o *AppSessionContextUpdateData) GetMcsId() string`

GetMcsId returns the McsId field if non-nil, zero value otherwise.

### GetMcsIdOk

`func (o *AppSessionContextUpdateData) GetMcsIdOk() (*string, bool)`

GetMcsIdOk returns a tuple with the McsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsId

`func (o *AppSessionContextUpdateData) SetMcsId(v string)`

SetMcsId sets McsId field to given value.

### HasMcsId

`func (o *AppSessionContextUpdateData) HasMcsId() bool`

HasMcsId returns a boolean if a field has been set.

### GetPreemptControlInfo

`func (o *AppSessionContextUpdateData) GetPreemptControlInfo() PreemptionControlInformationRm`

GetPreemptControlInfo returns the PreemptControlInfo field if non-nil, zero value otherwise.

### GetPreemptControlInfoOk

`func (o *AppSessionContextUpdateData) GetPreemptControlInfoOk() (*PreemptionControlInformationRm, bool)`

GetPreemptControlInfoOk returns a tuple with the PreemptControlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptControlInfo

`func (o *AppSessionContextUpdateData) SetPreemptControlInfo(v PreemptionControlInformationRm)`

SetPreemptControlInfo sets PreemptControlInfo field to given value.

### HasPreemptControlInfo

`func (o *AppSessionContextUpdateData) HasPreemptControlInfo() bool`

HasPreemptControlInfo returns a boolean if a field has been set.

### GetResPrio

`func (o *AppSessionContextUpdateData) GetResPrio() ReservPriority`

GetResPrio returns the ResPrio field if non-nil, zero value otherwise.

### GetResPrioOk

`func (o *AppSessionContextUpdateData) GetResPrioOk() (*ReservPriority, bool)`

GetResPrioOk returns a tuple with the ResPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPrio

`func (o *AppSessionContextUpdateData) SetResPrio(v ReservPriority)`

SetResPrio sets ResPrio field to given value.

### HasResPrio

`func (o *AppSessionContextUpdateData) HasResPrio() bool`

HasResPrio returns a boolean if a field has been set.

### GetServInfStatus

`func (o *AppSessionContextUpdateData) GetServInfStatus() ServiceInfoStatus`

GetServInfStatus returns the ServInfStatus field if non-nil, zero value otherwise.

### GetServInfStatusOk

`func (o *AppSessionContextUpdateData) GetServInfStatusOk() (*ServiceInfoStatus, bool)`

GetServInfStatusOk returns a tuple with the ServInfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServInfStatus

`func (o *AppSessionContextUpdateData) SetServInfStatus(v ServiceInfoStatus)`

SetServInfStatus sets ServInfStatus field to given value.

### HasServInfStatus

`func (o *AppSessionContextUpdateData) HasServInfStatus() bool`

HasServInfStatus returns a boolean if a field has been set.

### GetSipForkInd

`func (o *AppSessionContextUpdateData) GetSipForkInd() SipForkingIndication`

GetSipForkInd returns the SipForkInd field if non-nil, zero value otherwise.

### GetSipForkIndOk

`func (o *AppSessionContextUpdateData) GetSipForkIndOk() (*SipForkingIndication, bool)`

GetSipForkIndOk returns a tuple with the SipForkInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipForkInd

`func (o *AppSessionContextUpdateData) SetSipForkInd(v SipForkingIndication)`

SetSipForkInd sets SipForkInd field to given value.

### HasSipForkInd

`func (o *AppSessionContextUpdateData) HasSipForkInd() bool`

HasSipForkInd returns a boolean if a field has been set.

### GetSponId

`func (o *AppSessionContextUpdateData) GetSponId() string`

GetSponId returns the SponId field if non-nil, zero value otherwise.

### GetSponIdOk

`func (o *AppSessionContextUpdateData) GetSponIdOk() (*string, bool)`

GetSponIdOk returns a tuple with the SponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponId

`func (o *AppSessionContextUpdateData) SetSponId(v string)`

SetSponId sets SponId field to given value.

### HasSponId

`func (o *AppSessionContextUpdateData) HasSponId() bool`

HasSponId returns a boolean if a field has been set.

### GetSponStatus

`func (o *AppSessionContextUpdateData) GetSponStatus() SponsoringStatus`

GetSponStatus returns the SponStatus field if non-nil, zero value otherwise.

### GetSponStatusOk

`func (o *AppSessionContextUpdateData) GetSponStatusOk() (*SponsoringStatus, bool)`

GetSponStatusOk returns a tuple with the SponStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponStatus

`func (o *AppSessionContextUpdateData) SetSponStatus(v SponsoringStatus)`

SetSponStatus sets SponStatus field to given value.

### HasSponStatus

`func (o *AppSessionContextUpdateData) HasSponStatus() bool`

HasSponStatus returns a boolean if a field has been set.

### GetTsnBridgeManCont

`func (o *AppSessionContextUpdateData) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *AppSessionContextUpdateData) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *AppSessionContextUpdateData) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *AppSessionContextUpdateData) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *AppSessionContextUpdateData) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *AppSessionContextUpdateData) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *AppSessionContextUpdateData) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *AppSessionContextUpdateData) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *AppSessionContextUpdateData) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *AppSessionContextUpdateData) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *AppSessionContextUpdateData) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *AppSessionContextUpdateData) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


