# SmContextUpdatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**HoState** | Pointer to [**HoState**](HoState.md) |  | [optional] 
**ReleaseEbiList** | Pointer to **[]int32** |  | [optional] 
**AllocatedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**ModifiedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**N1SmMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**EpsBearerSetup** | Pointer to **[]string** |  | [optional] 
**DataForwarding** | Pointer to **bool** |  | [optional] 
**N3DlForwardingTnlList** | Pointer to [**[]IndirectDataForwardingTunnelInfo**](IndirectDataForwardingTunnelInfo.md) |  | [optional] 
**N3UlForwardingTnlList** | Pointer to [**[]IndirectDataForwardingTunnelInfo**](IndirectDataForwardingTunnelInfo.md) |  | [optional] 
**Cause** | Pointer to [**Cause**](Cause.md) |  | [optional] 
**MaAcceptedInd** | Pointer to **bool** |  | [optional] [default to false]
**SupportedFeatures** | Pointer to **string** |  | [optional] 
**ForwardingFTeid** | Pointer to **string** |  | [optional] 
**ForwardingBearerContexts** | Pointer to **[]string** |  | [optional] 
**SelectedSmfId** | Pointer to **string** |  | [optional] 
**SelectedOldSmfId** | Pointer to **string** |  | [optional] 
**AnchorSmfFeatures** | Pointer to [**AnchorSmfFeatures**](AnchorSmfFeatures.md) |  | [optional] 

## Methods

### NewSmContextUpdatedData

`func NewSmContextUpdatedData() *SmContextUpdatedData`

NewSmContextUpdatedData instantiates a new SmContextUpdatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextUpdatedDataWithDefaults

`func NewSmContextUpdatedDataWithDefaults() *SmContextUpdatedData`

NewSmContextUpdatedDataWithDefaults instantiates a new SmContextUpdatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpCnxState

`func (o *SmContextUpdatedData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *SmContextUpdatedData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *SmContextUpdatedData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *SmContextUpdatedData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetHoState

`func (o *SmContextUpdatedData) GetHoState() HoState`

GetHoState returns the HoState field if non-nil, zero value otherwise.

### GetHoStateOk

`func (o *SmContextUpdatedData) GetHoStateOk() (*HoState, bool)`

GetHoStateOk returns a tuple with the HoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoState

`func (o *SmContextUpdatedData) SetHoState(v HoState)`

SetHoState sets HoState field to given value.

### HasHoState

`func (o *SmContextUpdatedData) HasHoState() bool`

HasHoState returns a boolean if a field has been set.

### GetReleaseEbiList

`func (o *SmContextUpdatedData) GetReleaseEbiList() []int32`

GetReleaseEbiList returns the ReleaseEbiList field if non-nil, zero value otherwise.

### GetReleaseEbiListOk

`func (o *SmContextUpdatedData) GetReleaseEbiListOk() (*[]int32, bool)`

GetReleaseEbiListOk returns a tuple with the ReleaseEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseEbiList

`func (o *SmContextUpdatedData) SetReleaseEbiList(v []int32)`

SetReleaseEbiList sets ReleaseEbiList field to given value.

### HasReleaseEbiList

`func (o *SmContextUpdatedData) HasReleaseEbiList() bool`

HasReleaseEbiList returns a boolean if a field has been set.

### GetAllocatedEbiList

`func (o *SmContextUpdatedData) GetAllocatedEbiList() []EbiArpMapping`

GetAllocatedEbiList returns the AllocatedEbiList field if non-nil, zero value otherwise.

### GetAllocatedEbiListOk

`func (o *SmContextUpdatedData) GetAllocatedEbiListOk() (*[]EbiArpMapping, bool)`

GetAllocatedEbiListOk returns a tuple with the AllocatedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedEbiList

`func (o *SmContextUpdatedData) SetAllocatedEbiList(v []EbiArpMapping)`

SetAllocatedEbiList sets AllocatedEbiList field to given value.

### HasAllocatedEbiList

`func (o *SmContextUpdatedData) HasAllocatedEbiList() bool`

HasAllocatedEbiList returns a boolean if a field has been set.

### GetModifiedEbiList

`func (o *SmContextUpdatedData) GetModifiedEbiList() []EbiArpMapping`

GetModifiedEbiList returns the ModifiedEbiList field if non-nil, zero value otherwise.

### GetModifiedEbiListOk

`func (o *SmContextUpdatedData) GetModifiedEbiListOk() (*[]EbiArpMapping, bool)`

GetModifiedEbiListOk returns a tuple with the ModifiedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedEbiList

`func (o *SmContextUpdatedData) SetModifiedEbiList(v []EbiArpMapping)`

SetModifiedEbiList sets ModifiedEbiList field to given value.

### HasModifiedEbiList

`func (o *SmContextUpdatedData) HasModifiedEbiList() bool`

HasModifiedEbiList returns a boolean if a field has been set.

### GetN1SmMsg

`func (o *SmContextUpdatedData) GetN1SmMsg() RefToBinaryData`

GetN1SmMsg returns the N1SmMsg field if non-nil, zero value otherwise.

### GetN1SmMsgOk

`func (o *SmContextUpdatedData) GetN1SmMsgOk() (*RefToBinaryData, bool)`

GetN1SmMsgOk returns a tuple with the N1SmMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmMsg

`func (o *SmContextUpdatedData) SetN1SmMsg(v RefToBinaryData)`

SetN1SmMsg sets N1SmMsg field to given value.

### HasN1SmMsg

`func (o *SmContextUpdatedData) HasN1SmMsg() bool`

HasN1SmMsg returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextUpdatedData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextUpdatedData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextUpdatedData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextUpdatedData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextUpdatedData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextUpdatedData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextUpdatedData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextUpdatedData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetEpsBearerSetup

`func (o *SmContextUpdatedData) GetEpsBearerSetup() []string`

GetEpsBearerSetup returns the EpsBearerSetup field if non-nil, zero value otherwise.

### GetEpsBearerSetupOk

`func (o *SmContextUpdatedData) GetEpsBearerSetupOk() (*[]string, bool)`

GetEpsBearerSetupOk returns a tuple with the EpsBearerSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerSetup

`func (o *SmContextUpdatedData) SetEpsBearerSetup(v []string)`

SetEpsBearerSetup sets EpsBearerSetup field to given value.

### HasEpsBearerSetup

`func (o *SmContextUpdatedData) HasEpsBearerSetup() bool`

HasEpsBearerSetup returns a boolean if a field has been set.

### GetDataForwarding

`func (o *SmContextUpdatedData) GetDataForwarding() bool`

GetDataForwarding returns the DataForwarding field if non-nil, zero value otherwise.

### GetDataForwardingOk

`func (o *SmContextUpdatedData) GetDataForwardingOk() (*bool, bool)`

GetDataForwardingOk returns a tuple with the DataForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwarding

`func (o *SmContextUpdatedData) SetDataForwarding(v bool)`

SetDataForwarding sets DataForwarding field to given value.

### HasDataForwarding

`func (o *SmContextUpdatedData) HasDataForwarding() bool`

HasDataForwarding returns a boolean if a field has been set.

### GetN3DlForwardingTnlList

`func (o *SmContextUpdatedData) GetN3DlForwardingTnlList() []IndirectDataForwardingTunnelInfo`

GetN3DlForwardingTnlList returns the N3DlForwardingTnlList field if non-nil, zero value otherwise.

### GetN3DlForwardingTnlListOk

`func (o *SmContextUpdatedData) GetN3DlForwardingTnlListOk() (*[]IndirectDataForwardingTunnelInfo, bool)`

GetN3DlForwardingTnlListOk returns a tuple with the N3DlForwardingTnlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3DlForwardingTnlList

`func (o *SmContextUpdatedData) SetN3DlForwardingTnlList(v []IndirectDataForwardingTunnelInfo)`

SetN3DlForwardingTnlList sets N3DlForwardingTnlList field to given value.

### HasN3DlForwardingTnlList

`func (o *SmContextUpdatedData) HasN3DlForwardingTnlList() bool`

HasN3DlForwardingTnlList returns a boolean if a field has been set.

### GetN3UlForwardingTnlList

`func (o *SmContextUpdatedData) GetN3UlForwardingTnlList() []IndirectDataForwardingTunnelInfo`

GetN3UlForwardingTnlList returns the N3UlForwardingTnlList field if non-nil, zero value otherwise.

### GetN3UlForwardingTnlListOk

`func (o *SmContextUpdatedData) GetN3UlForwardingTnlListOk() (*[]IndirectDataForwardingTunnelInfo, bool)`

GetN3UlForwardingTnlListOk returns a tuple with the N3UlForwardingTnlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3UlForwardingTnlList

`func (o *SmContextUpdatedData) SetN3UlForwardingTnlList(v []IndirectDataForwardingTunnelInfo)`

SetN3UlForwardingTnlList sets N3UlForwardingTnlList field to given value.

### HasN3UlForwardingTnlList

`func (o *SmContextUpdatedData) HasN3UlForwardingTnlList() bool`

HasN3UlForwardingTnlList returns a boolean if a field has been set.

### GetCause

`func (o *SmContextUpdatedData) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SmContextUpdatedData) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SmContextUpdatedData) SetCause(v Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *SmContextUpdatedData) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetMaAcceptedInd

`func (o *SmContextUpdatedData) GetMaAcceptedInd() bool`

GetMaAcceptedInd returns the MaAcceptedInd field if non-nil, zero value otherwise.

### GetMaAcceptedIndOk

`func (o *SmContextUpdatedData) GetMaAcceptedIndOk() (*bool, bool)`

GetMaAcceptedIndOk returns a tuple with the MaAcceptedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaAcceptedInd

`func (o *SmContextUpdatedData) SetMaAcceptedInd(v bool)`

SetMaAcceptedInd sets MaAcceptedInd field to given value.

### HasMaAcceptedInd

`func (o *SmContextUpdatedData) HasMaAcceptedInd() bool`

HasMaAcceptedInd returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextUpdatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextUpdatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextUpdatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextUpdatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetForwardingFTeid

`func (o *SmContextUpdatedData) GetForwardingFTeid() string`

GetForwardingFTeid returns the ForwardingFTeid field if non-nil, zero value otherwise.

### GetForwardingFTeidOk

`func (o *SmContextUpdatedData) GetForwardingFTeidOk() (*string, bool)`

GetForwardingFTeidOk returns a tuple with the ForwardingFTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingFTeid

`func (o *SmContextUpdatedData) SetForwardingFTeid(v string)`

SetForwardingFTeid sets ForwardingFTeid field to given value.

### HasForwardingFTeid

`func (o *SmContextUpdatedData) HasForwardingFTeid() bool`

HasForwardingFTeid returns a boolean if a field has been set.

### GetForwardingBearerContexts

`func (o *SmContextUpdatedData) GetForwardingBearerContexts() []string`

GetForwardingBearerContexts returns the ForwardingBearerContexts field if non-nil, zero value otherwise.

### GetForwardingBearerContextsOk

`func (o *SmContextUpdatedData) GetForwardingBearerContextsOk() (*[]string, bool)`

GetForwardingBearerContextsOk returns a tuple with the ForwardingBearerContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingBearerContexts

`func (o *SmContextUpdatedData) SetForwardingBearerContexts(v []string)`

SetForwardingBearerContexts sets ForwardingBearerContexts field to given value.

### HasForwardingBearerContexts

`func (o *SmContextUpdatedData) HasForwardingBearerContexts() bool`

HasForwardingBearerContexts returns a boolean if a field has been set.

### GetSelectedSmfId

`func (o *SmContextUpdatedData) GetSelectedSmfId() string`

GetSelectedSmfId returns the SelectedSmfId field if non-nil, zero value otherwise.

### GetSelectedSmfIdOk

`func (o *SmContextUpdatedData) GetSelectedSmfIdOk() (*string, bool)`

GetSelectedSmfIdOk returns a tuple with the SelectedSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSmfId

`func (o *SmContextUpdatedData) SetSelectedSmfId(v string)`

SetSelectedSmfId sets SelectedSmfId field to given value.

### HasSelectedSmfId

`func (o *SmContextUpdatedData) HasSelectedSmfId() bool`

HasSelectedSmfId returns a boolean if a field has been set.

### GetSelectedOldSmfId

`func (o *SmContextUpdatedData) GetSelectedOldSmfId() string`

GetSelectedOldSmfId returns the SelectedOldSmfId field if non-nil, zero value otherwise.

### GetSelectedOldSmfIdOk

`func (o *SmContextUpdatedData) GetSelectedOldSmfIdOk() (*string, bool)`

GetSelectedOldSmfIdOk returns a tuple with the SelectedOldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOldSmfId

`func (o *SmContextUpdatedData) SetSelectedOldSmfId(v string)`

SetSelectedOldSmfId sets SelectedOldSmfId field to given value.

### HasSelectedOldSmfId

`func (o *SmContextUpdatedData) HasSelectedOldSmfId() bool`

HasSelectedOldSmfId returns a boolean if a field has been set.

### GetAnchorSmfFeatures

`func (o *SmContextUpdatedData) GetAnchorSmfFeatures() AnchorSmfFeatures`

GetAnchorSmfFeatures returns the AnchorSmfFeatures field if non-nil, zero value otherwise.

### GetAnchorSmfFeaturesOk

`func (o *SmContextUpdatedData) GetAnchorSmfFeaturesOk() (*AnchorSmfFeatures, bool)`

GetAnchorSmfFeaturesOk returns a tuple with the AnchorSmfFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorSmfFeatures

`func (o *SmContextUpdatedData) SetAnchorSmfFeatures(v AnchorSmfFeatures)`

SetAnchorSmfFeatures sets AnchorSmfFeatures field to given value.

### HasAnchorSmfFeatures

`func (o *SmContextUpdatedData) HasAnchorSmfFeatures() bool`

HasAnchorSmfFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


