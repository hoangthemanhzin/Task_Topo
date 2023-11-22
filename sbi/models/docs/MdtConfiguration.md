# MdtConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [**JobType**](JobType.md) |  | 
**ReportType** | Pointer to [**ReportTypeMdt**](ReportTypeMdt.md) |  | [optional] 
**AreaScope** | Pointer to [**AreaScope**](AreaScope.md) |  | [optional] 
**MeasurementLteList** | Pointer to [**[]MeasurementLteForMdt**](MeasurementLteForMdt.md) |  | [optional] 
**MeasurementNrList** | Pointer to [**[]MeasurementNrForMdt**](MeasurementNrForMdt.md) |  | [optional] 
**SensorMeasurementList** | Pointer to [**[]SensorMeasurement**](SensorMeasurement.md) |  | [optional] 
**ReportingTriggerList** | Pointer to [**[]ReportingTrigger**](ReportingTrigger.md) |  | [optional] 
**ReportInterval** | Pointer to [**ReportIntervalMdt**](ReportIntervalMdt.md) |  | [optional] 
**ReportIntervalNr** | Pointer to [**ReportIntervalNrMdt**](ReportIntervalNrMdt.md) |  | [optional] 
**ReportAmount** | Pointer to [**ReportAmountMdt**](ReportAmountMdt.md) |  | [optional] 
**EventThresholdRsrp** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrpNr** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrq** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrqNr** | Pointer to **int32** |  | [optional] 
**EventList** | Pointer to [**[]EventForMdt**](EventForMdt.md) |  | [optional] 
**LoggingInterval** | Pointer to [**LoggingIntervalMdt**](LoggingIntervalMdt.md) |  | [optional] 
**LoggingIntervalNr** | Pointer to [**LoggingIntervalNrMdt**](LoggingIntervalNrMdt.md) |  | [optional] 
**LoggingDuration** | Pointer to [**LoggingDurationMdt**](LoggingDurationMdt.md) |  | [optional] 
**LoggingDurationNr** | Pointer to [**LoggingDurationNrMdt**](LoggingDurationNrMdt.md) |  | [optional] 
**PositioningMethod** | Pointer to [**PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] 
**AddPositioningMethodList** | Pointer to [**[]PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] 
**CollectionPeriodRmmLte** | Pointer to [**CollectionPeriodRmmLteMdt**](CollectionPeriodRmmLteMdt.md) |  | [optional] 
**CollectionPeriodRmmNr** | Pointer to [**CollectionPeriodRmmNrMdt**](CollectionPeriodRmmNrMdt.md) |  | [optional] 
**MeasurementPeriodLte** | Pointer to [**MeasurementPeriodLteMdt**](MeasurementPeriodLteMdt.md) |  | [optional] 
**MdtAllowedPlmnIdList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**MbsfnAreaList** | Pointer to [**[]MbsfnArea**](MbsfnArea.md) |  | [optional] 
**InterFreqTargetList** | Pointer to [**[]InterFreqTargetInfo**](InterFreqTargetInfo.md) |  | [optional] 

## Methods

### NewMdtConfiguration

`func NewMdtConfiguration(jobType JobType, ) *MdtConfiguration`

NewMdtConfiguration instantiates a new MdtConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdtConfigurationWithDefaults

`func NewMdtConfigurationWithDefaults() *MdtConfiguration`

NewMdtConfigurationWithDefaults instantiates a new MdtConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *MdtConfiguration) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *MdtConfiguration) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *MdtConfiguration) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetReportType

`func (o *MdtConfiguration) GetReportType() ReportTypeMdt`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *MdtConfiguration) GetReportTypeOk() (*ReportTypeMdt, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *MdtConfiguration) SetReportType(v ReportTypeMdt)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *MdtConfiguration) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetAreaScope

`func (o *MdtConfiguration) GetAreaScope() AreaScope`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *MdtConfiguration) GetAreaScopeOk() (*AreaScope, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *MdtConfiguration) SetAreaScope(v AreaScope)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *MdtConfiguration) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.

### GetMeasurementLteList

`func (o *MdtConfiguration) GetMeasurementLteList() []MeasurementLteForMdt`

GetMeasurementLteList returns the MeasurementLteList field if non-nil, zero value otherwise.

### GetMeasurementLteListOk

`func (o *MdtConfiguration) GetMeasurementLteListOk() (*[]MeasurementLteForMdt, bool)`

GetMeasurementLteListOk returns a tuple with the MeasurementLteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementLteList

`func (o *MdtConfiguration) SetMeasurementLteList(v []MeasurementLteForMdt)`

SetMeasurementLteList sets MeasurementLteList field to given value.

### HasMeasurementLteList

`func (o *MdtConfiguration) HasMeasurementLteList() bool`

HasMeasurementLteList returns a boolean if a field has been set.

### GetMeasurementNrList

`func (o *MdtConfiguration) GetMeasurementNrList() []MeasurementNrForMdt`

GetMeasurementNrList returns the MeasurementNrList field if non-nil, zero value otherwise.

### GetMeasurementNrListOk

`func (o *MdtConfiguration) GetMeasurementNrListOk() (*[]MeasurementNrForMdt, bool)`

GetMeasurementNrListOk returns a tuple with the MeasurementNrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementNrList

`func (o *MdtConfiguration) SetMeasurementNrList(v []MeasurementNrForMdt)`

SetMeasurementNrList sets MeasurementNrList field to given value.

### HasMeasurementNrList

`func (o *MdtConfiguration) HasMeasurementNrList() bool`

HasMeasurementNrList returns a boolean if a field has been set.

### GetSensorMeasurementList

`func (o *MdtConfiguration) GetSensorMeasurementList() []SensorMeasurement`

GetSensorMeasurementList returns the SensorMeasurementList field if non-nil, zero value otherwise.

### GetSensorMeasurementListOk

`func (o *MdtConfiguration) GetSensorMeasurementListOk() (*[]SensorMeasurement, bool)`

GetSensorMeasurementListOk returns a tuple with the SensorMeasurementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorMeasurementList

`func (o *MdtConfiguration) SetSensorMeasurementList(v []SensorMeasurement)`

SetSensorMeasurementList sets SensorMeasurementList field to given value.

### HasSensorMeasurementList

`func (o *MdtConfiguration) HasSensorMeasurementList() bool`

HasSensorMeasurementList returns a boolean if a field has been set.

### GetReportingTriggerList

`func (o *MdtConfiguration) GetReportingTriggerList() []ReportingTrigger`

GetReportingTriggerList returns the ReportingTriggerList field if non-nil, zero value otherwise.

### GetReportingTriggerListOk

`func (o *MdtConfiguration) GetReportingTriggerListOk() (*[]ReportingTrigger, bool)`

GetReportingTriggerListOk returns a tuple with the ReportingTriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTriggerList

`func (o *MdtConfiguration) SetReportingTriggerList(v []ReportingTrigger)`

SetReportingTriggerList sets ReportingTriggerList field to given value.

### HasReportingTriggerList

`func (o *MdtConfiguration) HasReportingTriggerList() bool`

HasReportingTriggerList returns a boolean if a field has been set.

### GetReportInterval

`func (o *MdtConfiguration) GetReportInterval() ReportIntervalMdt`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *MdtConfiguration) GetReportIntervalOk() (*ReportIntervalMdt, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *MdtConfiguration) SetReportInterval(v ReportIntervalMdt)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *MdtConfiguration) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetReportIntervalNr

`func (o *MdtConfiguration) GetReportIntervalNr() ReportIntervalNrMdt`

GetReportIntervalNr returns the ReportIntervalNr field if non-nil, zero value otherwise.

### GetReportIntervalNrOk

`func (o *MdtConfiguration) GetReportIntervalNrOk() (*ReportIntervalNrMdt, bool)`

GetReportIntervalNrOk returns a tuple with the ReportIntervalNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalNr

`func (o *MdtConfiguration) SetReportIntervalNr(v ReportIntervalNrMdt)`

SetReportIntervalNr sets ReportIntervalNr field to given value.

### HasReportIntervalNr

`func (o *MdtConfiguration) HasReportIntervalNr() bool`

HasReportIntervalNr returns a boolean if a field has been set.

### GetReportAmount

`func (o *MdtConfiguration) GetReportAmount() ReportAmountMdt`

GetReportAmount returns the ReportAmount field if non-nil, zero value otherwise.

### GetReportAmountOk

`func (o *MdtConfiguration) GetReportAmountOk() (*ReportAmountMdt, bool)`

GetReportAmountOk returns a tuple with the ReportAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAmount

`func (o *MdtConfiguration) SetReportAmount(v ReportAmountMdt)`

SetReportAmount sets ReportAmount field to given value.

### HasReportAmount

`func (o *MdtConfiguration) HasReportAmount() bool`

HasReportAmount returns a boolean if a field has been set.

### GetEventThresholdRsrp

`func (o *MdtConfiguration) GetEventThresholdRsrp() int32`

GetEventThresholdRsrp returns the EventThresholdRsrp field if non-nil, zero value otherwise.

### GetEventThresholdRsrpOk

`func (o *MdtConfiguration) GetEventThresholdRsrpOk() (*int32, bool)`

GetEventThresholdRsrpOk returns a tuple with the EventThresholdRsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrp

`func (o *MdtConfiguration) SetEventThresholdRsrp(v int32)`

SetEventThresholdRsrp sets EventThresholdRsrp field to given value.

### HasEventThresholdRsrp

`func (o *MdtConfiguration) HasEventThresholdRsrp() bool`

HasEventThresholdRsrp returns a boolean if a field has been set.

### GetEventThresholdRsrpNr

`func (o *MdtConfiguration) GetEventThresholdRsrpNr() int32`

GetEventThresholdRsrpNr returns the EventThresholdRsrpNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrpNrOk

`func (o *MdtConfiguration) GetEventThresholdRsrpNrOk() (*int32, bool)`

GetEventThresholdRsrpNrOk returns a tuple with the EventThresholdRsrpNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrpNr

`func (o *MdtConfiguration) SetEventThresholdRsrpNr(v int32)`

SetEventThresholdRsrpNr sets EventThresholdRsrpNr field to given value.

### HasEventThresholdRsrpNr

`func (o *MdtConfiguration) HasEventThresholdRsrpNr() bool`

HasEventThresholdRsrpNr returns a boolean if a field has been set.

### GetEventThresholdRsrq

`func (o *MdtConfiguration) GetEventThresholdRsrq() int32`

GetEventThresholdRsrq returns the EventThresholdRsrq field if non-nil, zero value otherwise.

### GetEventThresholdRsrqOk

`func (o *MdtConfiguration) GetEventThresholdRsrqOk() (*int32, bool)`

GetEventThresholdRsrqOk returns a tuple with the EventThresholdRsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrq

`func (o *MdtConfiguration) SetEventThresholdRsrq(v int32)`

SetEventThresholdRsrq sets EventThresholdRsrq field to given value.

### HasEventThresholdRsrq

`func (o *MdtConfiguration) HasEventThresholdRsrq() bool`

HasEventThresholdRsrq returns a boolean if a field has been set.

### GetEventThresholdRsrqNr

`func (o *MdtConfiguration) GetEventThresholdRsrqNr() int32`

GetEventThresholdRsrqNr returns the EventThresholdRsrqNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrqNrOk

`func (o *MdtConfiguration) GetEventThresholdRsrqNrOk() (*int32, bool)`

GetEventThresholdRsrqNrOk returns a tuple with the EventThresholdRsrqNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrqNr

`func (o *MdtConfiguration) SetEventThresholdRsrqNr(v int32)`

SetEventThresholdRsrqNr sets EventThresholdRsrqNr field to given value.

### HasEventThresholdRsrqNr

`func (o *MdtConfiguration) HasEventThresholdRsrqNr() bool`

HasEventThresholdRsrqNr returns a boolean if a field has been set.

### GetEventList

`func (o *MdtConfiguration) GetEventList() []EventForMdt`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *MdtConfiguration) GetEventListOk() (*[]EventForMdt, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *MdtConfiguration) SetEventList(v []EventForMdt)`

SetEventList sets EventList field to given value.

### HasEventList

`func (o *MdtConfiguration) HasEventList() bool`

HasEventList returns a boolean if a field has been set.

### GetLoggingInterval

`func (o *MdtConfiguration) GetLoggingInterval() LoggingIntervalMdt`

GetLoggingInterval returns the LoggingInterval field if non-nil, zero value otherwise.

### GetLoggingIntervalOk

`func (o *MdtConfiguration) GetLoggingIntervalOk() (*LoggingIntervalMdt, bool)`

GetLoggingIntervalOk returns a tuple with the LoggingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingInterval

`func (o *MdtConfiguration) SetLoggingInterval(v LoggingIntervalMdt)`

SetLoggingInterval sets LoggingInterval field to given value.

### HasLoggingInterval

`func (o *MdtConfiguration) HasLoggingInterval() bool`

HasLoggingInterval returns a boolean if a field has been set.

### GetLoggingIntervalNr

`func (o *MdtConfiguration) GetLoggingIntervalNr() LoggingIntervalNrMdt`

GetLoggingIntervalNr returns the LoggingIntervalNr field if non-nil, zero value otherwise.

### GetLoggingIntervalNrOk

`func (o *MdtConfiguration) GetLoggingIntervalNrOk() (*LoggingIntervalNrMdt, bool)`

GetLoggingIntervalNrOk returns a tuple with the LoggingIntervalNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingIntervalNr

`func (o *MdtConfiguration) SetLoggingIntervalNr(v LoggingIntervalNrMdt)`

SetLoggingIntervalNr sets LoggingIntervalNr field to given value.

### HasLoggingIntervalNr

`func (o *MdtConfiguration) HasLoggingIntervalNr() bool`

HasLoggingIntervalNr returns a boolean if a field has been set.

### GetLoggingDuration

`func (o *MdtConfiguration) GetLoggingDuration() LoggingDurationMdt`

GetLoggingDuration returns the LoggingDuration field if non-nil, zero value otherwise.

### GetLoggingDurationOk

`func (o *MdtConfiguration) GetLoggingDurationOk() (*LoggingDurationMdt, bool)`

GetLoggingDurationOk returns a tuple with the LoggingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingDuration

`func (o *MdtConfiguration) SetLoggingDuration(v LoggingDurationMdt)`

SetLoggingDuration sets LoggingDuration field to given value.

### HasLoggingDuration

`func (o *MdtConfiguration) HasLoggingDuration() bool`

HasLoggingDuration returns a boolean if a field has been set.

### GetLoggingDurationNr

`func (o *MdtConfiguration) GetLoggingDurationNr() LoggingDurationNrMdt`

GetLoggingDurationNr returns the LoggingDurationNr field if non-nil, zero value otherwise.

### GetLoggingDurationNrOk

`func (o *MdtConfiguration) GetLoggingDurationNrOk() (*LoggingDurationNrMdt, bool)`

GetLoggingDurationNrOk returns a tuple with the LoggingDurationNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingDurationNr

`func (o *MdtConfiguration) SetLoggingDurationNr(v LoggingDurationNrMdt)`

SetLoggingDurationNr sets LoggingDurationNr field to given value.

### HasLoggingDurationNr

`func (o *MdtConfiguration) HasLoggingDurationNr() bool`

HasLoggingDurationNr returns a boolean if a field has been set.

### GetPositioningMethod

`func (o *MdtConfiguration) GetPositioningMethod() PositioningMethodMdt`

GetPositioningMethod returns the PositioningMethod field if non-nil, zero value otherwise.

### GetPositioningMethodOk

`func (o *MdtConfiguration) GetPositioningMethodOk() (*PositioningMethodMdt, bool)`

GetPositioningMethodOk returns a tuple with the PositioningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningMethod

`func (o *MdtConfiguration) SetPositioningMethod(v PositioningMethodMdt)`

SetPositioningMethod sets PositioningMethod field to given value.

### HasPositioningMethod

`func (o *MdtConfiguration) HasPositioningMethod() bool`

HasPositioningMethod returns a boolean if a field has been set.

### GetAddPositioningMethodList

`func (o *MdtConfiguration) GetAddPositioningMethodList() []PositioningMethodMdt`

GetAddPositioningMethodList returns the AddPositioningMethodList field if non-nil, zero value otherwise.

### GetAddPositioningMethodListOk

`func (o *MdtConfiguration) GetAddPositioningMethodListOk() (*[]PositioningMethodMdt, bool)`

GetAddPositioningMethodListOk returns a tuple with the AddPositioningMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPositioningMethodList

`func (o *MdtConfiguration) SetAddPositioningMethodList(v []PositioningMethodMdt)`

SetAddPositioningMethodList sets AddPositioningMethodList field to given value.

### HasAddPositioningMethodList

`func (o *MdtConfiguration) HasAddPositioningMethodList() bool`

HasAddPositioningMethodList returns a boolean if a field has been set.

### GetCollectionPeriodRmmLte

`func (o *MdtConfiguration) GetCollectionPeriodRmmLte() CollectionPeriodRmmLteMdt`

GetCollectionPeriodRmmLte returns the CollectionPeriodRmmLte field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmLteOk

`func (o *MdtConfiguration) GetCollectionPeriodRmmLteOk() (*CollectionPeriodRmmLteMdt, bool)`

GetCollectionPeriodRmmLteOk returns a tuple with the CollectionPeriodRmmLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmLte

`func (o *MdtConfiguration) SetCollectionPeriodRmmLte(v CollectionPeriodRmmLteMdt)`

SetCollectionPeriodRmmLte sets CollectionPeriodRmmLte field to given value.

### HasCollectionPeriodRmmLte

`func (o *MdtConfiguration) HasCollectionPeriodRmmLte() bool`

HasCollectionPeriodRmmLte returns a boolean if a field has been set.

### GetCollectionPeriodRmmNr

`func (o *MdtConfiguration) GetCollectionPeriodRmmNr() CollectionPeriodRmmNrMdt`

GetCollectionPeriodRmmNr returns the CollectionPeriodRmmNr field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmNrOk

`func (o *MdtConfiguration) GetCollectionPeriodRmmNrOk() (*CollectionPeriodRmmNrMdt, bool)`

GetCollectionPeriodRmmNrOk returns a tuple with the CollectionPeriodRmmNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmNr

`func (o *MdtConfiguration) SetCollectionPeriodRmmNr(v CollectionPeriodRmmNrMdt)`

SetCollectionPeriodRmmNr sets CollectionPeriodRmmNr field to given value.

### HasCollectionPeriodRmmNr

`func (o *MdtConfiguration) HasCollectionPeriodRmmNr() bool`

HasCollectionPeriodRmmNr returns a boolean if a field has been set.

### GetMeasurementPeriodLte

`func (o *MdtConfiguration) GetMeasurementPeriodLte() MeasurementPeriodLteMdt`

GetMeasurementPeriodLte returns the MeasurementPeriodLte field if non-nil, zero value otherwise.

### GetMeasurementPeriodLteOk

`func (o *MdtConfiguration) GetMeasurementPeriodLteOk() (*MeasurementPeriodLteMdt, bool)`

GetMeasurementPeriodLteOk returns a tuple with the MeasurementPeriodLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriodLte

`func (o *MdtConfiguration) SetMeasurementPeriodLte(v MeasurementPeriodLteMdt)`

SetMeasurementPeriodLte sets MeasurementPeriodLte field to given value.

### HasMeasurementPeriodLte

`func (o *MdtConfiguration) HasMeasurementPeriodLte() bool`

HasMeasurementPeriodLte returns a boolean if a field has been set.

### GetMdtAllowedPlmnIdList

`func (o *MdtConfiguration) GetMdtAllowedPlmnIdList() []PlmnId`

GetMdtAllowedPlmnIdList returns the MdtAllowedPlmnIdList field if non-nil, zero value otherwise.

### GetMdtAllowedPlmnIdListOk

`func (o *MdtConfiguration) GetMdtAllowedPlmnIdListOk() (*[]PlmnId, bool)`

GetMdtAllowedPlmnIdListOk returns a tuple with the MdtAllowedPlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtAllowedPlmnIdList

`func (o *MdtConfiguration) SetMdtAllowedPlmnIdList(v []PlmnId)`

SetMdtAllowedPlmnIdList sets MdtAllowedPlmnIdList field to given value.

### HasMdtAllowedPlmnIdList

`func (o *MdtConfiguration) HasMdtAllowedPlmnIdList() bool`

HasMdtAllowedPlmnIdList returns a boolean if a field has been set.

### GetMbsfnAreaList

`func (o *MdtConfiguration) GetMbsfnAreaList() []MbsfnArea`

GetMbsfnAreaList returns the MbsfnAreaList field if non-nil, zero value otherwise.

### GetMbsfnAreaListOk

`func (o *MdtConfiguration) GetMbsfnAreaListOk() (*[]MbsfnArea, bool)`

GetMbsfnAreaListOk returns a tuple with the MbsfnAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsfnAreaList

`func (o *MdtConfiguration) SetMbsfnAreaList(v []MbsfnArea)`

SetMbsfnAreaList sets MbsfnAreaList field to given value.

### HasMbsfnAreaList

`func (o *MdtConfiguration) HasMbsfnAreaList() bool`

HasMbsfnAreaList returns a boolean if a field has been set.

### GetInterFreqTargetList

`func (o *MdtConfiguration) GetInterFreqTargetList() []InterFreqTargetInfo`

GetInterFreqTargetList returns the InterFreqTargetList field if non-nil, zero value otherwise.

### GetInterFreqTargetListOk

`func (o *MdtConfiguration) GetInterFreqTargetListOk() (*[]InterFreqTargetInfo, bool)`

GetInterFreqTargetListOk returns a tuple with the InterFreqTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterFreqTargetList

`func (o *MdtConfiguration) SetInterFreqTargetList(v []InterFreqTargetInfo)`

SetInterFreqTargetList sets InterFreqTargetList field to given value.

### HasInterFreqTargetList

`func (o *MdtConfiguration) HasInterFreqTargetList() bool`

HasInterFreqTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


