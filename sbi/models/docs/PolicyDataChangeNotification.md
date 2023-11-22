# PolicyDataChangeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmPolicyData** | Pointer to [**AmPolicyData**](AmPolicyData.md) |  | [optional] 
**UePolicySet** | Pointer to [**UePolicySet**](UePolicySet.md) |  | [optional] 
**PlmnUePolicySet** | Pointer to [**UePolicySet**](UePolicySet.md) |  | [optional] 
**SmPolicyData** | Pointer to [**SmPolicyData**](SmPolicyData.md) |  | [optional] 
**UsageMonData** | Pointer to [**UsageMonData**](UsageMonData.md) |  | [optional] 
**SponsorConnectivityData** | Pointer to [**SponsorConnectivityData**](SponsorConnectivityData.md) |  | [optional] 
**BdtData** | Pointer to [**BdtData**](BdtData.md) |  | [optional] 
**OpSpecData** | Pointer to [**OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md) |  | [optional] 
**OpSpecDataMap** | Pointer to [**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md) |  | [optional] 
**UeId** | Pointer to **string** |  | [optional] 
**SponsorId** | Pointer to **string** |  | [optional] 
**BdtRefId** | Pointer to **string** | string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154. | [optional] 
**UsageMonId** | Pointer to **string** |  | [optional] 
**PlmnId** | Pointer to [**PlmnId1**](PlmnId1.md) |  | [optional] 
**DelResources** | Pointer to **[]string** |  | [optional] 
**NotifId** | Pointer to **string** |  | [optional] 
**ReportedFragments** | Pointer to [**[]NotificationItem**](NotificationItem.md) |  | [optional] 

## Methods

### NewPolicyDataChangeNotification

`func NewPolicyDataChangeNotification() *PolicyDataChangeNotification`

NewPolicyDataChangeNotification instantiates a new PolicyDataChangeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDataChangeNotificationWithDefaults

`func NewPolicyDataChangeNotificationWithDefaults() *PolicyDataChangeNotification`

NewPolicyDataChangeNotificationWithDefaults instantiates a new PolicyDataChangeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmPolicyData

`func (o *PolicyDataChangeNotification) GetAmPolicyData() AmPolicyData`

GetAmPolicyData returns the AmPolicyData field if non-nil, zero value otherwise.

### GetAmPolicyDataOk

`func (o *PolicyDataChangeNotification) GetAmPolicyDataOk() (*AmPolicyData, bool)`

GetAmPolicyDataOk returns a tuple with the AmPolicyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmPolicyData

`func (o *PolicyDataChangeNotification) SetAmPolicyData(v AmPolicyData)`

SetAmPolicyData sets AmPolicyData field to given value.

### HasAmPolicyData

`func (o *PolicyDataChangeNotification) HasAmPolicyData() bool`

HasAmPolicyData returns a boolean if a field has been set.

### GetUePolicySet

`func (o *PolicyDataChangeNotification) GetUePolicySet() UePolicySet`

GetUePolicySet returns the UePolicySet field if non-nil, zero value otherwise.

### GetUePolicySetOk

`func (o *PolicyDataChangeNotification) GetUePolicySetOk() (*UePolicySet, bool)`

GetUePolicySetOk returns a tuple with the UePolicySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicySet

`func (o *PolicyDataChangeNotification) SetUePolicySet(v UePolicySet)`

SetUePolicySet sets UePolicySet field to given value.

### HasUePolicySet

`func (o *PolicyDataChangeNotification) HasUePolicySet() bool`

HasUePolicySet returns a boolean if a field has been set.

### GetPlmnUePolicySet

`func (o *PolicyDataChangeNotification) GetPlmnUePolicySet() UePolicySet`

GetPlmnUePolicySet returns the PlmnUePolicySet field if non-nil, zero value otherwise.

### GetPlmnUePolicySetOk

`func (o *PolicyDataChangeNotification) GetPlmnUePolicySetOk() (*UePolicySet, bool)`

GetPlmnUePolicySetOk returns a tuple with the PlmnUePolicySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnUePolicySet

`func (o *PolicyDataChangeNotification) SetPlmnUePolicySet(v UePolicySet)`

SetPlmnUePolicySet sets PlmnUePolicySet field to given value.

### HasPlmnUePolicySet

`func (o *PolicyDataChangeNotification) HasPlmnUePolicySet() bool`

HasPlmnUePolicySet returns a boolean if a field has been set.

### GetSmPolicyData

`func (o *PolicyDataChangeNotification) GetSmPolicyData() SmPolicyData`

GetSmPolicyData returns the SmPolicyData field if non-nil, zero value otherwise.

### GetSmPolicyDataOk

`func (o *PolicyDataChangeNotification) GetSmPolicyDataOk() (*SmPolicyData, bool)`

GetSmPolicyDataOk returns a tuple with the SmPolicyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyData

`func (o *PolicyDataChangeNotification) SetSmPolicyData(v SmPolicyData)`

SetSmPolicyData sets SmPolicyData field to given value.

### HasSmPolicyData

`func (o *PolicyDataChangeNotification) HasSmPolicyData() bool`

HasSmPolicyData returns a boolean if a field has been set.

### GetUsageMonData

`func (o *PolicyDataChangeNotification) GetUsageMonData() UsageMonData`

GetUsageMonData returns the UsageMonData field if non-nil, zero value otherwise.

### GetUsageMonDataOk

`func (o *PolicyDataChangeNotification) GetUsageMonDataOk() (*UsageMonData, bool)`

GetUsageMonDataOk returns a tuple with the UsageMonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMonData

`func (o *PolicyDataChangeNotification) SetUsageMonData(v UsageMonData)`

SetUsageMonData sets UsageMonData field to given value.

### HasUsageMonData

`func (o *PolicyDataChangeNotification) HasUsageMonData() bool`

HasUsageMonData returns a boolean if a field has been set.

### GetSponsorConnectivityData

`func (o *PolicyDataChangeNotification) GetSponsorConnectivityData() SponsorConnectivityData`

GetSponsorConnectivityData returns the SponsorConnectivityData field if non-nil, zero value otherwise.

### GetSponsorConnectivityDataOk

`func (o *PolicyDataChangeNotification) GetSponsorConnectivityDataOk() (*SponsorConnectivityData, bool)`

GetSponsorConnectivityDataOk returns a tuple with the SponsorConnectivityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorConnectivityData

`func (o *PolicyDataChangeNotification) SetSponsorConnectivityData(v SponsorConnectivityData)`

SetSponsorConnectivityData sets SponsorConnectivityData field to given value.

### HasSponsorConnectivityData

`func (o *PolicyDataChangeNotification) HasSponsorConnectivityData() bool`

HasSponsorConnectivityData returns a boolean if a field has been set.

### GetBdtData

`func (o *PolicyDataChangeNotification) GetBdtData() BdtData`

GetBdtData returns the BdtData field if non-nil, zero value otherwise.

### GetBdtDataOk

`func (o *PolicyDataChangeNotification) GetBdtDataOk() (*BdtData, bool)`

GetBdtDataOk returns a tuple with the BdtData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtData

`func (o *PolicyDataChangeNotification) SetBdtData(v BdtData)`

SetBdtData sets BdtData field to given value.

### HasBdtData

`func (o *PolicyDataChangeNotification) HasBdtData() bool`

HasBdtData returns a boolean if a field has been set.

### GetOpSpecData

`func (o *PolicyDataChangeNotification) GetOpSpecData() OperatorSpecificDataContainer`

GetOpSpecData returns the OpSpecData field if non-nil, zero value otherwise.

### GetOpSpecDataOk

`func (o *PolicyDataChangeNotification) GetOpSpecDataOk() (*OperatorSpecificDataContainer, bool)`

GetOpSpecDataOk returns a tuple with the OpSpecData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSpecData

`func (o *PolicyDataChangeNotification) SetOpSpecData(v OperatorSpecificDataContainer)`

SetOpSpecData sets OpSpecData field to given value.

### HasOpSpecData

`func (o *PolicyDataChangeNotification) HasOpSpecData() bool`

HasOpSpecData returns a boolean if a field has been set.

### GetOpSpecDataMap

`func (o *PolicyDataChangeNotification) GetOpSpecDataMap() map[string]OperatorSpecificDataContainer`

GetOpSpecDataMap returns the OpSpecDataMap field if non-nil, zero value otherwise.

### GetOpSpecDataMapOk

`func (o *PolicyDataChangeNotification) GetOpSpecDataMapOk() (*map[string]OperatorSpecificDataContainer, bool)`

GetOpSpecDataMapOk returns a tuple with the OpSpecDataMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSpecDataMap

`func (o *PolicyDataChangeNotification) SetOpSpecDataMap(v map[string]OperatorSpecificDataContainer)`

SetOpSpecDataMap sets OpSpecDataMap field to given value.

### HasOpSpecDataMap

`func (o *PolicyDataChangeNotification) HasOpSpecDataMap() bool`

HasOpSpecDataMap returns a boolean if a field has been set.

### GetUeId

`func (o *PolicyDataChangeNotification) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *PolicyDataChangeNotification) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *PolicyDataChangeNotification) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *PolicyDataChangeNotification) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetSponsorId

`func (o *PolicyDataChangeNotification) GetSponsorId() string`

GetSponsorId returns the SponsorId field if non-nil, zero value otherwise.

### GetSponsorIdOk

`func (o *PolicyDataChangeNotification) GetSponsorIdOk() (*string, bool)`

GetSponsorIdOk returns a tuple with the SponsorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorId

`func (o *PolicyDataChangeNotification) SetSponsorId(v string)`

SetSponsorId sets SponsorId field to given value.

### HasSponsorId

`func (o *PolicyDataChangeNotification) HasSponsorId() bool`

HasSponsorId returns a boolean if a field has been set.

### GetBdtRefId

`func (o *PolicyDataChangeNotification) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *PolicyDataChangeNotification) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *PolicyDataChangeNotification) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.

### HasBdtRefId

`func (o *PolicyDataChangeNotification) HasBdtRefId() bool`

HasBdtRefId returns a boolean if a field has been set.

### GetUsageMonId

`func (o *PolicyDataChangeNotification) GetUsageMonId() string`

GetUsageMonId returns the UsageMonId field if non-nil, zero value otherwise.

### GetUsageMonIdOk

`func (o *PolicyDataChangeNotification) GetUsageMonIdOk() (*string, bool)`

GetUsageMonIdOk returns a tuple with the UsageMonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMonId

`func (o *PolicyDataChangeNotification) SetUsageMonId(v string)`

SetUsageMonId sets UsageMonId field to given value.

### HasUsageMonId

`func (o *PolicyDataChangeNotification) HasUsageMonId() bool`

HasUsageMonId returns a boolean if a field has been set.

### GetPlmnId

`func (o *PolicyDataChangeNotification) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PolicyDataChangeNotification) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PolicyDataChangeNotification) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PolicyDataChangeNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetDelResources

`func (o *PolicyDataChangeNotification) GetDelResources() []string`

GetDelResources returns the DelResources field if non-nil, zero value otherwise.

### GetDelResourcesOk

`func (o *PolicyDataChangeNotification) GetDelResourcesOk() (*[]string, bool)`

GetDelResourcesOk returns a tuple with the DelResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelResources

`func (o *PolicyDataChangeNotification) SetDelResources(v []string)`

SetDelResources sets DelResources field to given value.

### HasDelResources

`func (o *PolicyDataChangeNotification) HasDelResources() bool`

HasDelResources returns a boolean if a field has been set.

### GetNotifId

`func (o *PolicyDataChangeNotification) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *PolicyDataChangeNotification) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *PolicyDataChangeNotification) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.

### HasNotifId

`func (o *PolicyDataChangeNotification) HasNotifId() bool`

HasNotifId returns a boolean if a field has been set.

### GetReportedFragments

`func (o *PolicyDataChangeNotification) GetReportedFragments() []NotificationItem`

GetReportedFragments returns the ReportedFragments field if non-nil, zero value otherwise.

### GetReportedFragmentsOk

`func (o *PolicyDataChangeNotification) GetReportedFragmentsOk() (*[]NotificationItem, bool)`

GetReportedFragmentsOk returns a tuple with the ReportedFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedFragments

`func (o *PolicyDataChangeNotification) SetReportedFragments(v []NotificationItem)`

SetReportedFragments sets ReportedFragments field to given value.

### HasReportedFragments

`func (o *PolicyDataChangeNotification) HasReportedFragments() bool`

HasReportedFragments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


