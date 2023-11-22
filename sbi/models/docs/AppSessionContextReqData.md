# AppSessionContextReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**AfChargId** | Pointer to **string** |  | [optional] 
**AfReqData** | Pointer to [**AfRequestedData**](AfRequestedData.md) |  | [optional] 
**AfRoutReq** | Pointer to [**AfRoutingRequirement**](AfRoutingRequirement.md) |  | [optional] 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 
**BdtRefId** | Pointer to **string** | string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154. | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**EvSubsc** | Pointer to [**EventsSubscReqData**](EventsSubscReqData.md) |  | [optional] 
**McpttId** | Pointer to **string** | indication of MCPTT service request | [optional] 
**McVideoId** | Pointer to **string** | indication of MCVideo service request | [optional] 
**MedComponents** | Pointer to [**map[string]MediaComponent**](MediaComponent.md) |  | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**MpsId** | Pointer to **string** | indication of MPS service request | [optional] 
**McsId** | Pointer to **string** | indication of MCS service request | [optional] 
**PreemptControlInfo** | Pointer to [**PreemptionControlInformation**](PreemptionControlInformation.md) |  | [optional] 
**ResPrio** | Pointer to [**ReservPriority**](ReservPriority.md) |  | [optional] 
**ServInfStatus** | Pointer to [**ServiceInfoStatus**](ServiceInfoStatus.md) |  | [optional] 
**NotifUri** | **string** |  | 
**ServUrn** | Pointer to **string** | Contains values of the service URN and may include subservices. | [optional] 
**SliceInfo** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SponId** | Pointer to **string** | Contains an identity of a sponsor. | [optional] 
**SponStatus** | Pointer to [**SponsoringStatus**](SponsoringStatus.md) |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**SuppFeat** | **string** |  | 
**UeIpv4** | Pointer to **string** |  | [optional] 
**UeIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**UeMac** | Pointer to **string** |  | [optional] 
**TsnBridgeManCont** | Pointer to [**BridgeManagementContainer**](BridgeManagementContainer.md) |  | [optional] 
**TsnPortManContDstt** | Pointer to [**PortManagementContainer**](PortManagementContainer.md) |  | [optional] 
**TsnPortManContNwtts** | Pointer to [**[]PortManagementContainer**](PortManagementContainer.md) |  | [optional] 

## Methods

### NewAppSessionContextReqData

`func NewAppSessionContextReqData(notifUri string, suppFeat string, ) *AppSessionContextReqData`

NewAppSessionContextReqData instantiates a new AppSessionContextReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSessionContextReqDataWithDefaults

`func NewAppSessionContextReqDataWithDefaults() *AppSessionContextReqData`

NewAppSessionContextReqDataWithDefaults instantiates a new AppSessionContextReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfAppId

`func (o *AppSessionContextReqData) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *AppSessionContextReqData) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *AppSessionContextReqData) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *AppSessionContextReqData) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetAfChargId

`func (o *AppSessionContextReqData) GetAfChargId() string`

GetAfChargId returns the AfChargId field if non-nil, zero value otherwise.

### GetAfChargIdOk

`func (o *AppSessionContextReqData) GetAfChargIdOk() (*string, bool)`

GetAfChargIdOk returns a tuple with the AfChargId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfChargId

`func (o *AppSessionContextReqData) SetAfChargId(v string)`

SetAfChargId sets AfChargId field to given value.

### HasAfChargId

`func (o *AppSessionContextReqData) HasAfChargId() bool`

HasAfChargId returns a boolean if a field has been set.

### GetAfReqData

`func (o *AppSessionContextReqData) GetAfReqData() AfRequestedData`

GetAfReqData returns the AfReqData field if non-nil, zero value otherwise.

### GetAfReqDataOk

`func (o *AppSessionContextReqData) GetAfReqDataOk() (*AfRequestedData, bool)`

GetAfReqDataOk returns a tuple with the AfReqData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfReqData

`func (o *AppSessionContextReqData) SetAfReqData(v AfRequestedData)`

SetAfReqData sets AfReqData field to given value.

### HasAfReqData

`func (o *AppSessionContextReqData) HasAfReqData() bool`

HasAfReqData returns a boolean if a field has been set.

### GetAfRoutReq

`func (o *AppSessionContextReqData) GetAfRoutReq() AfRoutingRequirement`

GetAfRoutReq returns the AfRoutReq field if non-nil, zero value otherwise.

### GetAfRoutReqOk

`func (o *AppSessionContextReqData) GetAfRoutReqOk() (*AfRoutingRequirement, bool)`

GetAfRoutReqOk returns a tuple with the AfRoutReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfRoutReq

`func (o *AppSessionContextReqData) SetAfRoutReq(v AfRoutingRequirement)`

SetAfRoutReq sets AfRoutReq field to given value.

### HasAfRoutReq

`func (o *AppSessionContextReqData) HasAfRoutReq() bool`

HasAfRoutReq returns a boolean if a field has been set.

### GetAspId

`func (o *AppSessionContextReqData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *AppSessionContextReqData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *AppSessionContextReqData) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *AppSessionContextReqData) HasAspId() bool`

HasAspId returns a boolean if a field has been set.

### GetBdtRefId

`func (o *AppSessionContextReqData) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *AppSessionContextReqData) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *AppSessionContextReqData) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.

### HasBdtRefId

`func (o *AppSessionContextReqData) HasBdtRefId() bool`

HasBdtRefId returns a boolean if a field has been set.

### GetDnn

`func (o *AppSessionContextReqData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AppSessionContextReqData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AppSessionContextReqData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AppSessionContextReqData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetEvSubsc

`func (o *AppSessionContextReqData) GetEvSubsc() EventsSubscReqData`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *AppSessionContextReqData) GetEvSubscOk() (*EventsSubscReqData, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *AppSessionContextReqData) SetEvSubsc(v EventsSubscReqData)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *AppSessionContextReqData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### GetMcpttId

`func (o *AppSessionContextReqData) GetMcpttId() string`

GetMcpttId returns the McpttId field if non-nil, zero value otherwise.

### GetMcpttIdOk

`func (o *AppSessionContextReqData) GetMcpttIdOk() (*string, bool)`

GetMcpttIdOk returns a tuple with the McpttId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpttId

`func (o *AppSessionContextReqData) SetMcpttId(v string)`

SetMcpttId sets McpttId field to given value.

### HasMcpttId

`func (o *AppSessionContextReqData) HasMcpttId() bool`

HasMcpttId returns a boolean if a field has been set.

### GetMcVideoId

`func (o *AppSessionContextReqData) GetMcVideoId() string`

GetMcVideoId returns the McVideoId field if non-nil, zero value otherwise.

### GetMcVideoIdOk

`func (o *AppSessionContextReqData) GetMcVideoIdOk() (*string, bool)`

GetMcVideoIdOk returns a tuple with the McVideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcVideoId

`func (o *AppSessionContextReqData) SetMcVideoId(v string)`

SetMcVideoId sets McVideoId field to given value.

### HasMcVideoId

`func (o *AppSessionContextReqData) HasMcVideoId() bool`

HasMcVideoId returns a boolean if a field has been set.

### GetMedComponents

`func (o *AppSessionContextReqData) GetMedComponents() map[string]MediaComponent`

GetMedComponents returns the MedComponents field if non-nil, zero value otherwise.

### GetMedComponentsOk

`func (o *AppSessionContextReqData) GetMedComponentsOk() (*map[string]MediaComponent, bool)`

GetMedComponentsOk returns a tuple with the MedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedComponents

`func (o *AppSessionContextReqData) SetMedComponents(v map[string]MediaComponent)`

SetMedComponents sets MedComponents field to given value.

### HasMedComponents

`func (o *AppSessionContextReqData) HasMedComponents() bool`

HasMedComponents returns a boolean if a field has been set.

### GetIpDomain

`func (o *AppSessionContextReqData) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *AppSessionContextReqData) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *AppSessionContextReqData) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *AppSessionContextReqData) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetMpsId

`func (o *AppSessionContextReqData) GetMpsId() string`

GetMpsId returns the MpsId field if non-nil, zero value otherwise.

### GetMpsIdOk

`func (o *AppSessionContextReqData) GetMpsIdOk() (*string, bool)`

GetMpsIdOk returns a tuple with the MpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsId

`func (o *AppSessionContextReqData) SetMpsId(v string)`

SetMpsId sets MpsId field to given value.

### HasMpsId

`func (o *AppSessionContextReqData) HasMpsId() bool`

HasMpsId returns a boolean if a field has been set.

### GetMcsId

`func (o *AppSessionContextReqData) GetMcsId() string`

GetMcsId returns the McsId field if non-nil, zero value otherwise.

### GetMcsIdOk

`func (o *AppSessionContextReqData) GetMcsIdOk() (*string, bool)`

GetMcsIdOk returns a tuple with the McsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsId

`func (o *AppSessionContextReqData) SetMcsId(v string)`

SetMcsId sets McsId field to given value.

### HasMcsId

`func (o *AppSessionContextReqData) HasMcsId() bool`

HasMcsId returns a boolean if a field has been set.

### GetPreemptControlInfo

`func (o *AppSessionContextReqData) GetPreemptControlInfo() PreemptionControlInformation`

GetPreemptControlInfo returns the PreemptControlInfo field if non-nil, zero value otherwise.

### GetPreemptControlInfoOk

`func (o *AppSessionContextReqData) GetPreemptControlInfoOk() (*PreemptionControlInformation, bool)`

GetPreemptControlInfoOk returns a tuple with the PreemptControlInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptControlInfo

`func (o *AppSessionContextReqData) SetPreemptControlInfo(v PreemptionControlInformation)`

SetPreemptControlInfo sets PreemptControlInfo field to given value.

### HasPreemptControlInfo

`func (o *AppSessionContextReqData) HasPreemptControlInfo() bool`

HasPreemptControlInfo returns a boolean if a field has been set.

### GetResPrio

`func (o *AppSessionContextReqData) GetResPrio() ReservPriority`

GetResPrio returns the ResPrio field if non-nil, zero value otherwise.

### GetResPrioOk

`func (o *AppSessionContextReqData) GetResPrioOk() (*ReservPriority, bool)`

GetResPrioOk returns a tuple with the ResPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPrio

`func (o *AppSessionContextReqData) SetResPrio(v ReservPriority)`

SetResPrio sets ResPrio field to given value.

### HasResPrio

`func (o *AppSessionContextReqData) HasResPrio() bool`

HasResPrio returns a boolean if a field has been set.

### GetServInfStatus

`func (o *AppSessionContextReqData) GetServInfStatus() ServiceInfoStatus`

GetServInfStatus returns the ServInfStatus field if non-nil, zero value otherwise.

### GetServInfStatusOk

`func (o *AppSessionContextReqData) GetServInfStatusOk() (*ServiceInfoStatus, bool)`

GetServInfStatusOk returns a tuple with the ServInfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServInfStatus

`func (o *AppSessionContextReqData) SetServInfStatus(v ServiceInfoStatus)`

SetServInfStatus sets ServInfStatus field to given value.

### HasServInfStatus

`func (o *AppSessionContextReqData) HasServInfStatus() bool`

HasServInfStatus returns a boolean if a field has been set.

### GetNotifUri

`func (o *AppSessionContextReqData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *AppSessionContextReqData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *AppSessionContextReqData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetServUrn

`func (o *AppSessionContextReqData) GetServUrn() string`

GetServUrn returns the ServUrn field if non-nil, zero value otherwise.

### GetServUrnOk

`func (o *AppSessionContextReqData) GetServUrnOk() (*string, bool)`

GetServUrnOk returns a tuple with the ServUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServUrn

`func (o *AppSessionContextReqData) SetServUrn(v string)`

SetServUrn sets ServUrn field to given value.

### HasServUrn

`func (o *AppSessionContextReqData) HasServUrn() bool`

HasServUrn returns a boolean if a field has been set.

### GetSliceInfo

`func (o *AppSessionContextReqData) GetSliceInfo() Snssai`

GetSliceInfo returns the SliceInfo field if non-nil, zero value otherwise.

### GetSliceInfoOk

`func (o *AppSessionContextReqData) GetSliceInfoOk() (*Snssai, bool)`

GetSliceInfoOk returns a tuple with the SliceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceInfo

`func (o *AppSessionContextReqData) SetSliceInfo(v Snssai)`

SetSliceInfo sets SliceInfo field to given value.

### HasSliceInfo

`func (o *AppSessionContextReqData) HasSliceInfo() bool`

HasSliceInfo returns a boolean if a field has been set.

### GetSponId

`func (o *AppSessionContextReqData) GetSponId() string`

GetSponId returns the SponId field if non-nil, zero value otherwise.

### GetSponIdOk

`func (o *AppSessionContextReqData) GetSponIdOk() (*string, bool)`

GetSponIdOk returns a tuple with the SponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponId

`func (o *AppSessionContextReqData) SetSponId(v string)`

SetSponId sets SponId field to given value.

### HasSponId

`func (o *AppSessionContextReqData) HasSponId() bool`

HasSponId returns a boolean if a field has been set.

### GetSponStatus

`func (o *AppSessionContextReqData) GetSponStatus() SponsoringStatus`

GetSponStatus returns the SponStatus field if non-nil, zero value otherwise.

### GetSponStatusOk

`func (o *AppSessionContextReqData) GetSponStatusOk() (*SponsoringStatus, bool)`

GetSponStatusOk returns a tuple with the SponStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponStatus

`func (o *AppSessionContextReqData) SetSponStatus(v SponsoringStatus)`

SetSponStatus sets SponStatus field to given value.

### HasSponStatus

`func (o *AppSessionContextReqData) HasSponStatus() bool`

HasSponStatus returns a boolean if a field has been set.

### GetSupi

`func (o *AppSessionContextReqData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *AppSessionContextReqData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *AppSessionContextReqData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *AppSessionContextReqData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *AppSessionContextReqData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AppSessionContextReqData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AppSessionContextReqData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AppSessionContextReqData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AppSessionContextReqData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AppSessionContextReqData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AppSessionContextReqData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.


### GetUeIpv4

`func (o *AppSessionContextReqData) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *AppSessionContextReqData) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *AppSessionContextReqData) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *AppSessionContextReqData) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *AppSessionContextReqData) GetUeIpv6() Ipv6Addr`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *AppSessionContextReqData) GetUeIpv6Ok() (*Ipv6Addr, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *AppSessionContextReqData) SetUeIpv6(v Ipv6Addr)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *AppSessionContextReqData) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetUeMac

`func (o *AppSessionContextReqData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *AppSessionContextReqData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *AppSessionContextReqData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *AppSessionContextReqData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetTsnBridgeManCont

`func (o *AppSessionContextReqData) GetTsnBridgeManCont() BridgeManagementContainer`

GetTsnBridgeManCont returns the TsnBridgeManCont field if non-nil, zero value otherwise.

### GetTsnBridgeManContOk

`func (o *AppSessionContextReqData) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool)`

GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnBridgeManCont

`func (o *AppSessionContextReqData) SetTsnBridgeManCont(v BridgeManagementContainer)`

SetTsnBridgeManCont sets TsnBridgeManCont field to given value.

### HasTsnBridgeManCont

`func (o *AppSessionContextReqData) HasTsnBridgeManCont() bool`

HasTsnBridgeManCont returns a boolean if a field has been set.

### GetTsnPortManContDstt

`func (o *AppSessionContextReqData) GetTsnPortManContDstt() PortManagementContainer`

GetTsnPortManContDstt returns the TsnPortManContDstt field if non-nil, zero value otherwise.

### GetTsnPortManContDsttOk

`func (o *AppSessionContextReqData) GetTsnPortManContDsttOk() (*PortManagementContainer, bool)`

GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContDstt

`func (o *AppSessionContextReqData) SetTsnPortManContDstt(v PortManagementContainer)`

SetTsnPortManContDstt sets TsnPortManContDstt field to given value.

### HasTsnPortManContDstt

`func (o *AppSessionContextReqData) HasTsnPortManContDstt() bool`

HasTsnPortManContDstt returns a boolean if a field has been set.

### GetTsnPortManContNwtts

`func (o *AppSessionContextReqData) GetTsnPortManContNwtts() []PortManagementContainer`

GetTsnPortManContNwtts returns the TsnPortManContNwtts field if non-nil, zero value otherwise.

### GetTsnPortManContNwttsOk

`func (o *AppSessionContextReqData) GetTsnPortManContNwttsOk() (*[]PortManagementContainer, bool)`

GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsnPortManContNwtts

`func (o *AppSessionContextReqData) SetTsnPortManContNwtts(v []PortManagementContainer)`

SetTsnPortManContNwtts sets TsnPortManContNwtts field to given value.

### HasTsnPortManContNwtts

`func (o *AppSessionContextReqData) HasTsnPortManContNwtts() bool`

HasTsnPortManContNwtts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


