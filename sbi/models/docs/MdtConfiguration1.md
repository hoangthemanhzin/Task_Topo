# MdtConfiguration1

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
**MdtAllowedPlmnIdList** | Pointer to [**[]PlmnId1**](PlmnId1.md) |  | [optional] 
**MbsfnAreaList** | Pointer to [**[]MbsfnArea**](MbsfnArea.md) |  | [optional] 
**InterFreqTargetList** | Pointer to [**[]InterFreqTargetInfo1**](InterFreqTargetInfo1.md) |  | [optional] 

## Methods

### NewMdtConfiguration1

`func NewMdtConfiguration1(jobType JobType, ) *MdtConfiguration1`

NewMdtConfiguration1 instantiates a new MdtConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdtConfiguration1WithDefaults

`func NewMdtConfiguration1WithDefaults() *MdtConfiguration1`

NewMdtConfiguration1WithDefaults instantiates a new MdtConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *MdtConfiguration1) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *MdtConfiguration1) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *MdtConfiguration1) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetReportType

`func (o *MdtConfiguration1) GetReportType() ReportTypeMdt`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *MdtConfiguration1) GetReportTypeOk() (*ReportTypeMdt, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *MdtConfiguration1) SetReportType(v ReportTypeMdt)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *MdtConfiguration1) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetAreaScope

`func (o *MdtConfiguration1) GetAreaScope() AreaScope`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *MdtConfiguration1) GetAreaScopeOk() (*AreaScope, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *MdtConfiguration1) SetAreaScope(v AreaScope)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *MdtConfiguration1) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.

### GetMeasurementLteList

`func (o *MdtConfiguration1) GetMeasurementLteList() []MeasurementLteForMdt`

GetMeasurementLteList returns the MeasurementLteList field if non-nil, zero value otherwise.

### GetMeasurementLteListOk

`func (o *MdtConfiguration1) GetMeasurementLteListOk() (*[]MeasurementLteForMdt, bool)`

GetMeasurementLteListOk returns a tuple with the MeasurementLteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementLteList

`func (o *MdtConfiguration1) SetMeasurementLteList(v []MeasurementLteForMdt)`

SetMeasurementLteList sets MeasurementLteList field to given value.

### HasMeasurementLteList

`func (o *MdtConfiguration1) HasMeasurementLteList() bool`

HasMeasurementLteList returns a boolean if a field has been set.

### GetMeasurementNrList

`func (o *MdtConfiguration1) GetMeasurementNrList() []MeasurementNrForMdt`

GetMeasurementNrList returns the MeasurementNrList field if non-nil, zero value otherwise.

### GetMeasurementNrListOk

`func (o *MdtConfiguration1) GetMeasurementNrListOk() (*[]MeasurementNrForMdt, bool)`

GetMeasurementNrListOk returns a tuple with the MeasurementNrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementNrList

`func (o *MdtConfiguration1) SetMeasurementNrList(v []MeasurementNrForMdt)`

SetMeasurementNrList sets MeasurementNrList field to given value.

### HasMeasurementNrList

`func (o *MdtConfiguration1) HasMeasurementNrList() bool`

HasMeasurementNrList returns a boolean if a field has been set.

### GetSensorMeasurementList

`func (o *MdtConfiguration1) GetSensorMeasurementList() []SensorMeasurement`

GetSensorMeasurementList returns the SensorMeasurementList field if non-nil, zero value otherwise.

### GetSensorMeasurementListOk

`func (o *MdtConfiguration1) GetSensorMeasurementListOk() (*[]SensorMeasurement, bool)`

GetSensorMeasurementListOk returns a tuple with the SensorMeasurementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorMeasurementList

`func (o *MdtConfiguration1) SetSensorMeasurementList(v []SensorMeasurement)`

SetSensorMeasurementList sets SensorMeasurementList field to given value.

### HasSensorMeasurementList

`func (o *MdtConfiguration1) HasSensorMeasurementList() bool`

HasSensorMeasurementList returns a boolean if a field has been set.

### GetReportingTriggerList

`func (o *MdtConfiguration1) GetReportingTriggerList() []ReportingTrigger`

GetReportingTriggerList returns the ReportingTriggerList field if non-nil, zero value otherwise.

### GetReportingTriggerListOk

`func (o *MdtConfiguration1) GetReportingTriggerListOk() (*[]ReportingTrigger, bool)`

GetReportingTriggerListOk returns a tuple with the ReportingTriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTriggerList

`func (o *MdtConfiguration1) SetReportingTriggerList(v []ReportingTrigger)`

SetReportingTriggerList sets ReportingTriggerList field to given value.

### HasReportingTriggerList

`func (o *MdtConfiguration1) HasReportingTriggerList() bool`

HasReportingTriggerList returns a boolean if a field has been set.

### GetReportInterval

`func (o *MdtConfiguration1) GetReportInterval() ReportIntervalMdt`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *MdtConfiguration1) GetReportIntervalOk() (*ReportIntervalMdt, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *MdtConfiguration1) SetReportInterval(v ReportIntervalMdt)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *MdtConfiguration1) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetReportIntervalNr

`func (o *MdtConfiguration1) GetReportIntervalNr() ReportIntervalNrMdt`

GetReportIntervalNr returns the ReportIntervalNr field if non-nil, zero value otherwise.

### GetReportIntervalNrOk

`func (o *MdtConfiguration1) GetReportIntervalNrOk() (*ReportIntervalNrMdt, bool)`

GetReportIntervalNrOk returns a tuple with the ReportIntervalNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalNr

`func (o *MdtConfiguration1) SetReportIntervalNr(v ReportIntervalNrMdt)`

SetReportIntervalNr sets ReportIntervalNr field to given value.

### HasReportIntervalNr

`func (o *MdtConfiguration1) HasReportIntervalNr() bool`

HasReportIntervalNr returns a boolean if a field has been set.

### GetReportAmount

`func (o *MdtConfiguration1) GetReportAmount() ReportAmountMdt`

GetReportAmount returns the ReportAmount field if non-nil, zero value otherwise.

### GetReportAmountOk

`func (o *MdtConfiguration1) GetReportAmountOk() (*ReportAmountMdt, bool)`

GetReportAmountOk returns a tuple with the ReportAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAmount

`func (o *MdtConfiguration1) SetReportAmount(v ReportAmountMdt)`

SetReportAmount sets ReportAmount field to given value.

### HasReportAmount

`func (o *MdtConfiguration1) HasReportAmount() bool`

HasReportAmount returns a boolean if a field has been set.

### GetEventThresholdRsrp

`func (o *MdtConfiguration1) GetEventThresholdRsrp() int32`

GetEventThresholdRsrp returns the EventThresholdRsrp field if non-nil, zero value otherwise.

### GetEventThresholdRsrpOk

`func (o *MdtConfiguration1) GetEventThresholdRsrpOk() (*int32, bool)`

GetEventThresholdRsrpOk returns a tuple with the EventThresholdRsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrp

`func (o *MdtConfiguration1) SetEventThresholdRsrp(v int32)`

SetEventThresholdRsrp sets EventThresholdRsrp field to given value.

### HasEventThresholdRsrp

`func (o *MdtConfiguration1) HasEventThresholdRsrp() bool`

HasEventThresholdRsrp returns a boolean if a field has been set.

### GetEventThresholdRsrpNr

`func (o *MdtConfiguration1) GetEventThresholdRsrpNr() int32`

GetEventThresholdRsrpNr returns the EventThresholdRsrpNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrpNrOk

`func (o *MdtConfiguration1) GetEventThresholdRsrpNrOk() (*int32, bool)`

GetEventThresholdRsrpNrOk returns a tuple with the EventThresholdRsrpNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrpNr

`func (o *MdtConfiguration1) SetEventThresholdRsrpNr(v int32)`

SetEventThresholdRsrpNr sets EventThresholdRsrpNr field to given value.

### HasEventThresholdRsrpNr

`func (o *MdtConfiguration1) HasEventThresholdRsrpNr() bool`

HasEventThresholdRsrpNr returns a boolean if a field has been set.

### GetEventThresholdRsrq

`func (o *MdtConfiguration1) GetEventThresholdRsrq() int32`

GetEventThresholdRsrq returns the EventThresholdRsrq field if non-nil, zero value otherwise.

### GetEventThresholdRsrqOk

`func (o *MdtConfiguration1) GetEventThresholdRsrqOk() (*int32, bool)`

GetEventThresholdRsrqOk returns a tuple with the EventThresholdRsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrq

`func (o *MdtConfiguration1) SetEventThresholdRsrq(v int32)`

SetEventThresholdRsrq sets EventThresholdRsrq field to given value.

### HasEventThresholdRsrq

`func (o *MdtConfiguration1) HasEventThresholdRsrq() bool`

HasEventThresholdRsrq returns a boolean if a field has been set.

### GetEventThresholdRsrqNr

`func (o *MdtConfiguration1) GetEventThresholdRsrqNr() int32`

GetEventThresholdRsrqNr returns the EventThresholdRsrqNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrqNrOk

`func (o *MdtConfiguration1) GetEventThresholdRsrqNrOk() (*int32, bool)`

GetEventThresholdRsrqNrOk returns a tuple with the EventThresholdRsrqNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrqNr

`func (o *MdtConfiguration1) SetEventThresholdRsrqNr(v int32)`

SetEventThresholdRsrqNr sets EventThresholdRsrqNr field to given value.

### HasEventThresholdRsrqNr

`func (o *MdtConfiguration1) HasEventThresholdRsrqNr() bool`

HasEventThresholdRsrqNr returns a boolean if a field has been set.

### GetEventList

`func (o *MdtConfiguration1) GetEventList() []EventForMdt`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *MdtConfiguration1) GetEventListOk() (*[]EventForMdt, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *MdtConfiguration1) SetEventList(v []EventForMdt)`

SetEventList sets EventList field to given value.

### HasEventList

`func (o *MdtConfiguration1) HasEventList() bool`

HasEventList returns a boolean if a field has been set.

### GetLoggingInterval

`func (o *MdtConfiguration1) GetLoggingInterval() LoggingIntervalMdt`

GetLoggingInterval returns the LoggingInterval field if non-nil, zero value otherwise.

### GetLoggingIntervalOk

`func (o *MdtConfiguration1) GetLoggingIntervalOk() (*LoggingIntervalMdt, bool)`

GetLoggingIntervalOk returns a tuple with the LoggingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingInterval

`func (o *MdtConfiguration1) SetLoggingInterval(v LoggingIntervalMdt)`

SetLoggingInterval sets LoggingInterval field to given value.

### HasLoggingInterval

`func (o *MdtConfiguration1) HasLoggingInterval() bool`

HasLoggingInterval returns a boolean if a field has been set.

### GetLoggingIntervalNr

`func (o *MdtConfiguration1) GetLoggingIntervalNr() LoggingIntervalNrMdt`

GetLoggingIntervalNr returns the LoggingIntervalNr field if non-nil, zero value otherwise.

### GetLoggingIntervalNrOk

`func (o *MdtConfiguration1) GetLoggingIntervalNrOk() (*LoggingIntervalNrMdt, bool)`

GetLoggingIntervalNrOk returns a tuple with the LoggingIntervalNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingIntervalNr

`func (o *MdtConfiguration1) SetLoggingIntervalNr(v LoggingIntervalNrMdt)`

SetLoggingIntervalNr sets LoggingIntervalNr field to given value.

### HasLoggingIntervalNr

`func (o *MdtConfiguration1) HasLoggingIntervalNr() bool`

HasLoggingIntervalNr returns a boolean if a field has been set.

### GetLoggingDuration

`func (o *MdtConfiguration1) GetLoggingDuration() LoggingDurationMdt`

GetLoggingDuration returns the LoggingDuration field if non-nil, zero value otherwise.

### GetLoggingDurationOk

`func (o *MdtConfiguration1) GetLoggingDurationOk() (*LoggingDurationMdt, bool)`

GetLoggingDurationOk returns a tuple with the LoggingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingDuration

`func (o *MdtConfiguration1) SetLoggingDuration(v LoggingDurationMdt)`

SetLoggingDuration sets LoggingDuration field to given value.

### HasLoggingDuration

`func (o *MdtConfiguration1) HasLoggingDuration() bool`

HasLoggingDuration returns a boolean if a field has been set.

### GetLoggingDurationNr

`func (o *MdtConfiguration1) GetLoggingDurationNr() LoggingDurationNrMdt`

GetLoggingDurationNr returns the LoggingDurationNr field if non-nil, zero value otherwise.

### GetLoggingDurationNrOk

`func (o *MdtConfiguration1) GetLoggingDurationNrOk() (*LoggingDurationNrMdt, bool)`

GetLoggingDurationNrOk returns a tuple with the LoggingDurationNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingDurationNr

`func (o *MdtConfiguration1) SetLoggingDurationNr(v LoggingDurationNrMdt)`

SetLoggingDurationNr sets LoggingDurationNr field to given value.

### HasLoggingDurationNr

`func (o *MdtConfiguration1) HasLoggingDurationNr() bool`

HasLoggingDurationNr returns a boolean if a field has been set.

### GetPositioningMethod

`func (o *MdtConfiguration1) GetPositioningMethod() PositioningMethodMdt`

GetPositioningMethod returns the PositioningMethod field if non-nil, zero value otherwise.

### GetPositioningMethodOk

`func (o *MdtConfiguration1) GetPositioningMethodOk() (*PositioningMethodMdt, bool)`

GetPositioningMethodOk returns a tuple with the PositioningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningMethod

`func (o *MdtConfiguration1) SetPositioningMethod(v PositioningMethodMdt)`

SetPositioningMethod sets PositioningMethod field to given value.

### HasPositioningMethod

`func (o *MdtConfiguration1) HasPositioningMethod() bool`

HasPositioningMethod returns a boolean if a field has been set.

### GetAddPositioningMethodList

`func (o *MdtConfiguration1) GetAddPositioningMethodList() []PositioningMethodMdt`

GetAddPositioningMethodList returns the AddPositioningMethodList field if non-nil, zero value otherwise.

### GetAddPositioningMethodListOk

`func (o *MdtConfiguration1) GetAddPositioningMethodListOk() (*[]PositioningMethodMdt, bool)`

GetAddPositioningMethodListOk returns a tuple with the AddPositioningMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPositioningMethodList

`func (o *MdtConfiguration1) SetAddPositioningMethodList(v []PositioningMethodMdt)`

SetAddPositioningMethodList sets AddPositioningMethodList field to given value.

### HasAddPositioningMethodList

`func (o *MdtConfiguration1) HasAddPositioningMethodList() bool`

HasAddPositioningMethodList returns a boolean if a field has been set.

### GetCollectionPeriodRmmLte

`func (o *MdtConfiguration1) GetCollectionPeriodRmmLte() CollectionPeriodRmmLteMdt`

GetCollectionPeriodRmmLte returns the CollectionPeriodRmmLte field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmLteOk

`func (o *MdtConfiguration1) GetCollectionPeriodRmmLteOk() (*CollectionPeriodRmmLteMdt, bool)`

GetCollectionPeriodRmmLteOk returns a tuple with the CollectionPeriodRmmLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmLte

`func (o *MdtConfiguration1) SetCollectionPeriodRmmLte(v CollectionPeriodRmmLteMdt)`

SetCollectionPeriodRmmLte sets CollectionPeriodRmmLte field to given value.

### HasCollectionPeriodRmmLte

`func (o *MdtConfiguration1) HasCollectionPeriodRmmLte() bool`

HasCollectionPeriodRmmLte returns a boolean if a field has been set.

### GetCollectionPeriodRmmNr

`func (o *MdtConfiguration1) GetCollectionPeriodRmmNr() CollectionPeriodRmmNrMdt`

GetCollectionPeriodRmmNr returns the CollectionPeriodRmmNr field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmNrOk

`func (o *MdtConfiguration1) GetCollectionPeriodRmmNrOk() (*CollectionPeriodRmmNrMdt, bool)`

GetCollectionPeriodRmmNrOk returns a tuple with the CollectionPeriodRmmNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmNr

`func (o *MdtConfiguration1) SetCollectionPeriodRmmNr(v CollectionPeriodRmmNrMdt)`

SetCollectionPeriodRmmNr sets CollectionPeriodRmmNr field to given value.

### HasCollectionPeriodRmmNr

`func (o *MdtConfiguration1) HasCollectionPeriodRmmNr() bool`

HasCollectionPeriodRmmNr returns a boolean if a field has been set.

### GetMeasurementPeriodLte

`func (o *MdtConfiguration1) GetMeasurementPeriodLte() MeasurementPeriodLteMdt`

GetMeasurementPeriodLte returns the MeasurementPeriodLte field if non-nil, zero value otherwise.

### GetMeasurementPeriodLteOk

`func (o *MdtConfiguration1) GetMeasurementPeriodLteOk() (*MeasurementPeriodLteMdt, bool)`

GetMeasurementPeriodLteOk returns a tuple with the MeasurementPeriodLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriodLte

`func (o *MdtConfiguration1) SetMeasurementPeriodLte(v MeasurementPeriodLteMdt)`

SetMeasurementPeriodLte sets MeasurementPeriodLte field to given value.

### HasMeasurementPeriodLte

`func (o *MdtConfiguration1) HasMeasurementPeriodLte() bool`

HasMeasurementPeriodLte returns a boolean if a field has been set.

### GetMdtAllowedPlmnIdList

`func (o *MdtConfiguration1) GetMdtAllowedPlmnIdList() []PlmnId1`

GetMdtAllowedPlmnIdList returns the MdtAllowedPlmnIdList field if non-nil, zero value otherwise.

### GetMdtAllowedPlmnIdListOk

`func (o *MdtConfiguration1) GetMdtAllowedPlmnIdListOk() (*[]PlmnId1, bool)`

GetMdtAllowedPlmnIdListOk returns a tuple with the MdtAllowedPlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtAllowedPlmnIdList

`func (o *MdtConfiguration1) SetMdtAllowedPlmnIdList(v []PlmnId1)`

SetMdtAllowedPlmnIdList sets MdtAllowedPlmnIdList field to given value.

### HasMdtAllowedPlmnIdList

`func (o *MdtConfiguration1) HasMdtAllowedPlmnIdList() bool`

HasMdtAllowedPlmnIdList returns a boolean if a field has been set.

### GetMbsfnAreaList

`func (o *MdtConfiguration1) GetMbsfnAreaList() []MbsfnArea`

GetMbsfnAreaList returns the MbsfnAreaList field if non-nil, zero value otherwise.

### GetMbsfnAreaListOk

`func (o *MdtConfiguration1) GetMbsfnAreaListOk() (*[]MbsfnArea, bool)`

GetMbsfnAreaListOk returns a tuple with the MbsfnAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsfnAreaList

`func (o *MdtConfiguration1) SetMbsfnAreaList(v []MbsfnArea)`

SetMbsfnAreaList sets MbsfnAreaList field to given value.

### HasMbsfnAreaList

`func (o *MdtConfiguration1) HasMbsfnAreaList() bool`

HasMbsfnAreaList returns a boolean if a field has been set.

### GetInterFreqTargetList

`func (o *MdtConfiguration1) GetInterFreqTargetList() []InterFreqTargetInfo1`

GetInterFreqTargetList returns the InterFreqTargetList field if non-nil, zero value otherwise.

### GetInterFreqTargetListOk

`func (o *MdtConfiguration1) GetInterFreqTargetListOk() (*[]InterFreqTargetInfo1, bool)`

GetInterFreqTargetListOk returns a tuple with the InterFreqTargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterFreqTargetList

`func (o *MdtConfiguration1) SetInterFreqTargetList(v []InterFreqTargetInfo1)`

SetInterFreqTargetList sets InterFreqTargetList field to given value.

### HasInterFreqTargetList

`func (o *MdtConfiguration1) HasInterFreqTargetList() bool`

HasInterFreqTargetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


