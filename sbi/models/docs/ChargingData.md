# ChargingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChgId** | **string** | Univocally identifies the charging control policy data within a PDU session. | 
**MeteringMethod** | Pointer to [**MeteringMethod**](MeteringMethod.md) |  | [optional] 
**Offline** | Pointer to **bool** | Indicates the offline charging is applicable to the PCC rule when it is included and set to true. | [optional] 
**Online** | Pointer to **bool** | Indicates the online charging is applicable to the PCC rule when it is included and set to true. | [optional] 
**SdfHandl** | Pointer to **bool** | Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request. | [optional] 
**RatingGroup** | Pointer to **int32** |  | [optional] 
**ReportingLevel** | Pointer to [**ReportingLevel**](ReportingLevel.md) |  | [optional] 
**ServiceId** | Pointer to **int32** |  | [optional] 
**SponsorId** | Pointer to **string** | Indicates the sponsor identity. | [optional] 
**AppSvcProvId** | Pointer to **string** | Indicates the application service provider identity. | [optional] 
**AfChargingIdentifier** | Pointer to **int32** |  | [optional] 
**AfChargId** | Pointer to **string** |  | [optional] 

## Methods

### NewChargingData

`func NewChargingData(chgId string, ) *ChargingData`

NewChargingData instantiates a new ChargingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingDataWithDefaults

`func NewChargingDataWithDefaults() *ChargingData`

NewChargingDataWithDefaults instantiates a new ChargingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChgId

`func (o *ChargingData) GetChgId() string`

GetChgId returns the ChgId field if non-nil, zero value otherwise.

### GetChgIdOk

`func (o *ChargingData) GetChgIdOk() (*string, bool)`

GetChgIdOk returns a tuple with the ChgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChgId

`func (o *ChargingData) SetChgId(v string)`

SetChgId sets ChgId field to given value.


### GetMeteringMethod

`func (o *ChargingData) GetMeteringMethod() MeteringMethod`

GetMeteringMethod returns the MeteringMethod field if non-nil, zero value otherwise.

### GetMeteringMethodOk

`func (o *ChargingData) GetMeteringMethodOk() (*MeteringMethod, bool)`

GetMeteringMethodOk returns a tuple with the MeteringMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringMethod

`func (o *ChargingData) SetMeteringMethod(v MeteringMethod)`

SetMeteringMethod sets MeteringMethod field to given value.

### HasMeteringMethod

`func (o *ChargingData) HasMeteringMethod() bool`

HasMeteringMethod returns a boolean if a field has been set.

### GetOffline

`func (o *ChargingData) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *ChargingData) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *ChargingData) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *ChargingData) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetOnline

`func (o *ChargingData) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ChargingData) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ChargingData) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ChargingData) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetSdfHandl

`func (o *ChargingData) GetSdfHandl() bool`

GetSdfHandl returns the SdfHandl field if non-nil, zero value otherwise.

### GetSdfHandlOk

`func (o *ChargingData) GetSdfHandlOk() (*bool, bool)`

GetSdfHandlOk returns a tuple with the SdfHandl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdfHandl

`func (o *ChargingData) SetSdfHandl(v bool)`

SetSdfHandl sets SdfHandl field to given value.

### HasSdfHandl

`func (o *ChargingData) HasSdfHandl() bool`

HasSdfHandl returns a boolean if a field has been set.

### GetRatingGroup

`func (o *ChargingData) GetRatingGroup() int32`

GetRatingGroup returns the RatingGroup field if non-nil, zero value otherwise.

### GetRatingGroupOk

`func (o *ChargingData) GetRatingGroupOk() (*int32, bool)`

GetRatingGroupOk returns a tuple with the RatingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingGroup

`func (o *ChargingData) SetRatingGroup(v int32)`

SetRatingGroup sets RatingGroup field to given value.

### HasRatingGroup

`func (o *ChargingData) HasRatingGroup() bool`

HasRatingGroup returns a boolean if a field has been set.

### GetReportingLevel

`func (o *ChargingData) GetReportingLevel() ReportingLevel`

GetReportingLevel returns the ReportingLevel field if non-nil, zero value otherwise.

### GetReportingLevelOk

`func (o *ChargingData) GetReportingLevelOk() (*ReportingLevel, bool)`

GetReportingLevelOk returns a tuple with the ReportingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLevel

`func (o *ChargingData) SetReportingLevel(v ReportingLevel)`

SetReportingLevel sets ReportingLevel field to given value.

### HasReportingLevel

`func (o *ChargingData) HasReportingLevel() bool`

HasReportingLevel returns a boolean if a field has been set.

### GetServiceId

`func (o *ChargingData) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ChargingData) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ChargingData) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ChargingData) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSponsorId

`func (o *ChargingData) GetSponsorId() string`

GetSponsorId returns the SponsorId field if non-nil, zero value otherwise.

### GetSponsorIdOk

`func (o *ChargingData) GetSponsorIdOk() (*string, bool)`

GetSponsorIdOk returns a tuple with the SponsorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorId

`func (o *ChargingData) SetSponsorId(v string)`

SetSponsorId sets SponsorId field to given value.

### HasSponsorId

`func (o *ChargingData) HasSponsorId() bool`

HasSponsorId returns a boolean if a field has been set.

### GetAppSvcProvId

`func (o *ChargingData) GetAppSvcProvId() string`

GetAppSvcProvId returns the AppSvcProvId field if non-nil, zero value otherwise.

### GetAppSvcProvIdOk

`func (o *ChargingData) GetAppSvcProvIdOk() (*string, bool)`

GetAppSvcProvIdOk returns a tuple with the AppSvcProvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSvcProvId

`func (o *ChargingData) SetAppSvcProvId(v string)`

SetAppSvcProvId sets AppSvcProvId field to given value.

### HasAppSvcProvId

`func (o *ChargingData) HasAppSvcProvId() bool`

HasAppSvcProvId returns a boolean if a field has been set.

### GetAfChargingIdentifier

`func (o *ChargingData) GetAfChargingIdentifier() int32`

GetAfChargingIdentifier returns the AfChargingIdentifier field if non-nil, zero value otherwise.

### GetAfChargingIdentifierOk

`func (o *ChargingData) GetAfChargingIdentifierOk() (*int32, bool)`

GetAfChargingIdentifierOk returns a tuple with the AfChargingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfChargingIdentifier

`func (o *ChargingData) SetAfChargingIdentifier(v int32)`

SetAfChargingIdentifier sets AfChargingIdentifier field to given value.

### HasAfChargingIdentifier

`func (o *ChargingData) HasAfChargingIdentifier() bool`

HasAfChargingIdentifier returns a boolean if a field has been set.

### GetAfChargId

`func (o *ChargingData) GetAfChargId() string`

GetAfChargId returns the AfChargId field if non-nil, zero value otherwise.

### GetAfChargIdOk

`func (o *ChargingData) GetAfChargIdOk() (*string, bool)`

GetAfChargIdOk returns a tuple with the AfChargId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfChargId

`func (o *ChargingData) SetAfChargId(v string)`

SetAfChargId sets AfChargId field to given value.

### HasAfChargId

`func (o *ChargingData) HasAfChargId() bool`

HasAfChargId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


