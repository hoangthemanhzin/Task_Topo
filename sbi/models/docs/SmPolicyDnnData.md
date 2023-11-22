# SmPolicyDnnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** |  | 
**AllowedServices** | Pointer to **[]string** |  | [optional] 
**SubscCats** | Pointer to **[]string** |  | [optional] 
**GbrUl** | Pointer to **string** |  | [optional] 
**GbrDl** | Pointer to **string** |  | [optional] 
**AdcSupport** | Pointer to **bool** |  | [optional] 
**SubscSpendingLimits** | Pointer to **bool** |  | [optional] 
**Ipv4Index** | Pointer to **int32** |  | [optional] 
**Ipv6Index** | Pointer to **int32** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**ChfInfo** | Pointer to [**ChargingInformation**](ChargingInformation.md) |  | [optional] 
**RefUmDataLimitIds** | Pointer to [**map[string]LimitIdToMonitoringKey**](LimitIdToMonitoringKey.md) |  | [optional] 
**MpsPriority** | Pointer to **bool** |  | [optional] 
**McsPriority** | Pointer to **bool** |  | [optional] 
**ImsSignallingPrio** | Pointer to **bool** |  | [optional] 
**MpsPriorityLevel** | Pointer to **int32** |  | [optional] 
**McsPriorityLevel** | Pointer to **int32** |  | [optional] 
**PraInfos** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) |  | [optional] 
**BdtRefIds** | Pointer to **map[string]string** |  | [optional] 
**LocRoutNotAllowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmPolicyDnnData

`func NewSmPolicyDnnData(dnn string, ) *SmPolicyDnnData`

NewSmPolicyDnnData instantiates a new SmPolicyDnnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDnnDataWithDefaults

`func NewSmPolicyDnnDataWithDefaults() *SmPolicyDnnData`

NewSmPolicyDnnDataWithDefaults instantiates a new SmPolicyDnnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *SmPolicyDnnData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmPolicyDnnData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmPolicyDnnData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetAllowedServices

`func (o *SmPolicyDnnData) GetAllowedServices() []string`

GetAllowedServices returns the AllowedServices field if non-nil, zero value otherwise.

### GetAllowedServicesOk

`func (o *SmPolicyDnnData) GetAllowedServicesOk() (*[]string, bool)`

GetAllowedServicesOk returns a tuple with the AllowedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServices

`func (o *SmPolicyDnnData) SetAllowedServices(v []string)`

SetAllowedServices sets AllowedServices field to given value.

### HasAllowedServices

`func (o *SmPolicyDnnData) HasAllowedServices() bool`

HasAllowedServices returns a boolean if a field has been set.

### GetSubscCats

`func (o *SmPolicyDnnData) GetSubscCats() []string`

GetSubscCats returns the SubscCats field if non-nil, zero value otherwise.

### GetSubscCatsOk

`func (o *SmPolicyDnnData) GetSubscCatsOk() (*[]string, bool)`

GetSubscCatsOk returns a tuple with the SubscCats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscCats

`func (o *SmPolicyDnnData) SetSubscCats(v []string)`

SetSubscCats sets SubscCats field to given value.

### HasSubscCats

`func (o *SmPolicyDnnData) HasSubscCats() bool`

HasSubscCats returns a boolean if a field has been set.

### GetGbrUl

`func (o *SmPolicyDnnData) GetGbrUl() string`

GetGbrUl returns the GbrUl field if non-nil, zero value otherwise.

### GetGbrUlOk

`func (o *SmPolicyDnnData) GetGbrUlOk() (*string, bool)`

GetGbrUlOk returns a tuple with the GbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrUl

`func (o *SmPolicyDnnData) SetGbrUl(v string)`

SetGbrUl sets GbrUl field to given value.

### HasGbrUl

`func (o *SmPolicyDnnData) HasGbrUl() bool`

HasGbrUl returns a boolean if a field has been set.

### GetGbrDl

`func (o *SmPolicyDnnData) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *SmPolicyDnnData) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *SmPolicyDnnData) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *SmPolicyDnnData) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### GetAdcSupport

`func (o *SmPolicyDnnData) GetAdcSupport() bool`

GetAdcSupport returns the AdcSupport field if non-nil, zero value otherwise.

### GetAdcSupportOk

`func (o *SmPolicyDnnData) GetAdcSupportOk() (*bool, bool)`

GetAdcSupportOk returns a tuple with the AdcSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdcSupport

`func (o *SmPolicyDnnData) SetAdcSupport(v bool)`

SetAdcSupport sets AdcSupport field to given value.

### HasAdcSupport

`func (o *SmPolicyDnnData) HasAdcSupport() bool`

HasAdcSupport returns a boolean if a field has been set.

### GetSubscSpendingLimits

`func (o *SmPolicyDnnData) GetSubscSpendingLimits() bool`

GetSubscSpendingLimits returns the SubscSpendingLimits field if non-nil, zero value otherwise.

### GetSubscSpendingLimitsOk

`func (o *SmPolicyDnnData) GetSubscSpendingLimitsOk() (*bool, bool)`

GetSubscSpendingLimitsOk returns a tuple with the SubscSpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscSpendingLimits

`func (o *SmPolicyDnnData) SetSubscSpendingLimits(v bool)`

SetSubscSpendingLimits sets SubscSpendingLimits field to given value.

### HasSubscSpendingLimits

`func (o *SmPolicyDnnData) HasSubscSpendingLimits() bool`

HasSubscSpendingLimits returns a boolean if a field has been set.

### GetIpv4Index

`func (o *SmPolicyDnnData) GetIpv4Index() int32`

GetIpv4Index returns the Ipv4Index field if non-nil, zero value otherwise.

### GetIpv4IndexOk

`func (o *SmPolicyDnnData) GetIpv4IndexOk() (*int32, bool)`

GetIpv4IndexOk returns a tuple with the Ipv4Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Index

`func (o *SmPolicyDnnData) SetIpv4Index(v int32)`

SetIpv4Index sets Ipv4Index field to given value.

### HasIpv4Index

`func (o *SmPolicyDnnData) HasIpv4Index() bool`

HasIpv4Index returns a boolean if a field has been set.

### GetIpv6Index

`func (o *SmPolicyDnnData) GetIpv6Index() int32`

GetIpv6Index returns the Ipv6Index field if non-nil, zero value otherwise.

### GetIpv6IndexOk

`func (o *SmPolicyDnnData) GetIpv6IndexOk() (*int32, bool)`

GetIpv6IndexOk returns a tuple with the Ipv6Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Index

`func (o *SmPolicyDnnData) SetIpv6Index(v int32)`

SetIpv6Index sets Ipv6Index field to given value.

### HasIpv6Index

`func (o *SmPolicyDnnData) HasIpv6Index() bool`

HasIpv6Index returns a boolean if a field has been set.

### GetOffline

`func (o *SmPolicyDnnData) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *SmPolicyDnnData) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *SmPolicyDnnData) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *SmPolicyDnnData) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOnline

`func (o *SmPolicyDnnData) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SmPolicyDnnData) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SmPolicyDnnData) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SmPolicyDnnData) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetChfInfo

`func (o *SmPolicyDnnData) GetChfInfo() ChargingInformation`

GetChfInfo returns the ChfInfo field if non-nil, zero value otherwise.

### GetChfInfoOk

`func (o *SmPolicyDnnData) GetChfInfoOk() (*ChargingInformation, bool)`

GetChfInfoOk returns a tuple with the ChfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfo

`func (o *SmPolicyDnnData) SetChfInfo(v ChargingInformation)`

SetChfInfo sets ChfInfo field to given value.

### HasChfInfo

`func (o *SmPolicyDnnData) HasChfInfo() bool`

HasChfInfo returns a boolean if a field has been set.

### GetRefUmDataLimitIds

`func (o *SmPolicyDnnData) GetRefUmDataLimitIds() map[string]LimitIdToMonitoringKey`

GetRefUmDataLimitIds returns the RefUmDataLimitIds field if non-nil, zero value otherwise.

### GetRefUmDataLimitIdsOk

`func (o *SmPolicyDnnData) GetRefUmDataLimitIdsOk() (*map[string]LimitIdToMonitoringKey, bool)`

GetRefUmDataLimitIdsOk returns a tuple with the RefUmDataLimitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmDataLimitIds

`func (o *SmPolicyDnnData) SetRefUmDataLimitIds(v map[string]LimitIdToMonitoringKey)`

SetRefUmDataLimitIds sets RefUmDataLimitIds field to given value.

### HasRefUmDataLimitIds

`func (o *SmPolicyDnnData) HasRefUmDataLimitIds() bool`

HasRefUmDataLimitIds returns a boolean if a field has been set.

### GetMpsPriority

`func (o *SmPolicyDnnData) GetMpsPriority() bool`

GetMpsPriority returns the MpsPriority field if non-nil, zero value otherwise.

### GetMpsPriorityOk

`func (o *SmPolicyDnnData) GetMpsPriorityOk() (*bool, bool)`

GetMpsPriorityOk returns a tuple with the MpsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsPriority

`func (o *SmPolicyDnnData) SetMpsPriority(v bool)`

SetMpsPriority sets MpsPriority field to given value.

### HasMpsPriority

`func (o *SmPolicyDnnData) HasMpsPriority() bool`

HasMpsPriority returns a boolean if a field has been set.

### GetMcsPriority

`func (o *SmPolicyDnnData) GetMcsPriority() bool`

GetMcsPriority returns the McsPriority field if non-nil, zero value otherwise.

### GetMcsPriorityOk

`func (o *SmPolicyDnnData) GetMcsPriorityOk() (*bool, bool)`

GetMcsPriorityOk returns a tuple with the McsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsPriority

`func (o *SmPolicyDnnData) SetMcsPriority(v bool)`

SetMcsPriority sets McsPriority field to given value.

### HasMcsPriority

`func (o *SmPolicyDnnData) HasMcsPriority() bool`

HasMcsPriority returns a boolean if a field has been set.

### GetImsSignallingPrio

`func (o *SmPolicyDnnData) GetImsSignallingPrio() bool`

GetImsSignallingPrio returns the ImsSignallingPrio field if non-nil, zero value otherwise.

### GetImsSignallingPrioOk

`func (o *SmPolicyDnnData) GetImsSignallingPrioOk() (*bool, bool)`

GetImsSignallingPrioOk returns a tuple with the ImsSignallingPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsSignallingPrio

`func (o *SmPolicyDnnData) SetImsSignallingPrio(v bool)`

SetImsSignallingPrio sets ImsSignallingPrio field to given value.

### HasImsSignallingPrio

`func (o *SmPolicyDnnData) HasImsSignallingPrio() bool`

HasImsSignallingPrio returns a boolean if a field has been set.

### GetMpsPriorityLevel

`func (o *SmPolicyDnnData) GetMpsPriorityLevel() int32`

GetMpsPriorityLevel returns the MpsPriorityLevel field if non-nil, zero value otherwise.

### GetMpsPriorityLevelOk

`func (o *SmPolicyDnnData) GetMpsPriorityLevelOk() (*int32, bool)`

GetMpsPriorityLevelOk returns a tuple with the MpsPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpsPriorityLevel

`func (o *SmPolicyDnnData) SetMpsPriorityLevel(v int32)`

SetMpsPriorityLevel sets MpsPriorityLevel field to given value.

### HasMpsPriorityLevel

`func (o *SmPolicyDnnData) HasMpsPriorityLevel() bool`

HasMpsPriorityLevel returns a boolean if a field has been set.

### GetMcsPriorityLevel

`func (o *SmPolicyDnnData) GetMcsPriorityLevel() int32`

GetMcsPriorityLevel returns the McsPriorityLevel field if non-nil, zero value otherwise.

### GetMcsPriorityLevelOk

`func (o *SmPolicyDnnData) GetMcsPriorityLevelOk() (*int32, bool)`

GetMcsPriorityLevelOk returns a tuple with the McsPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsPriorityLevel

`func (o *SmPolicyDnnData) SetMcsPriorityLevel(v int32)`

SetMcsPriorityLevel sets McsPriorityLevel field to given value.

### HasMcsPriorityLevel

`func (o *SmPolicyDnnData) HasMcsPriorityLevel() bool`

HasMcsPriorityLevel returns a boolean if a field has been set.

### GetPraInfos

`func (o *SmPolicyDnnData) GetPraInfos() map[string]PresenceInfo`

GetPraInfos returns the PraInfos field if non-nil, zero value otherwise.

### GetPraInfosOk

`func (o *SmPolicyDnnData) GetPraInfosOk() (*map[string]PresenceInfo, bool)`

GetPraInfosOk returns a tuple with the PraInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInfos

`func (o *SmPolicyDnnData) SetPraInfos(v map[string]PresenceInfo)`

SetPraInfos sets PraInfos field to given value.

### HasPraInfos

`func (o *SmPolicyDnnData) HasPraInfos() bool`

HasPraInfos returns a boolean if a field has been set.

### GetBdtRefIds

`func (o *SmPolicyDnnData) GetBdtRefIds() map[string]string`

GetBdtRefIds returns the BdtRefIds field if non-nil, zero value otherwise.

### GetBdtRefIdsOk

`func (o *SmPolicyDnnData) GetBdtRefIdsOk() (*map[string]string, bool)`

GetBdtRefIdsOk returns a tuple with the BdtRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefIds

`func (o *SmPolicyDnnData) SetBdtRefIds(v map[string]string)`

SetBdtRefIds sets BdtRefIds field to given value.

### HasBdtRefIds

`func (o *SmPolicyDnnData) HasBdtRefIds() bool`

HasBdtRefIds returns a boolean if a field has been set.

### SetBdtRefIdsNil

`func (o *SmPolicyDnnData) SetBdtRefIdsNil(b bool)`

 SetBdtRefIdsNil sets the value for BdtRefIds to be an explicit nil

### UnsetBdtRefIds
`func (o *SmPolicyDnnData) UnsetBdtRefIds()`

UnsetBdtRefIds ensures that no value is present for BdtRefIds, not even an explicit nil
### GetLocRoutNotAllowed

`func (o *SmPolicyDnnData) GetLocRoutNotAllowed() bool`

GetLocRoutNotAllowed returns the LocRoutNotAllowed field if non-nil, zero value otherwise.

### GetLocRoutNotAllowedOk

`func (o *SmPolicyDnnData) GetLocRoutNotAllowedOk() (*bool, bool)`

GetLocRoutNotAllowedOk returns a tuple with the LocRoutNotAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocRoutNotAllowed

`func (o *SmPolicyDnnData) SetLocRoutNotAllowed(v bool)`

SetLocRoutNotAllowed sets LocRoutNotAllowed field to given value.

### HasLocRoutNotAllowed

`func (o *SmPolicyDnnData) HasLocRoutNotAllowed() bool`

HasLocRoutNotAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


