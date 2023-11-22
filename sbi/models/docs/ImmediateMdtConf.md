# ImmediateMdtConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [**JobType**](JobType.md) |  | 
**MeasurementLteList** | Pointer to [**[]MeasurementLteForMdt**](MeasurementLteForMdt.md) |  | [optional] 
**MeasurementNrList** | Pointer to [**[]MeasurementNrForMdt**](MeasurementNrForMdt.md) |  | [optional] 
**ReportingTriggerList** | Pointer to [**[]ReportingTrigger**](ReportingTrigger.md) |  | [optional] 
**ReportInterval** | Pointer to [**ReportIntervalMdt**](ReportIntervalMdt.md) |  | [optional] 
**ReportIntervalNr** | Pointer to [**ReportIntervalNrMdt**](ReportIntervalNrMdt.md) |  | [optional] 
**ReportAmount** | Pointer to [**ReportAmountMdt**](ReportAmountMdt.md) |  | [optional] 
**EventThresholdRsrp** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrq** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrpNr** | Pointer to **int32** |  | [optional] 
**EventThresholdRsrqNr** | Pointer to **int32** |  | [optional] 
**CollectionPeriodRmmLte** | Pointer to [**CollectionPeriodRmmLteMdt**](CollectionPeriodRmmLteMdt.md) |  | [optional] 
**CollectionPeriodRmmNr** | Pointer to [**CollectionPeriodRmmNrMdt**](CollectionPeriodRmmNrMdt.md) |  | [optional] 
**MeasurementPeriodLte** | Pointer to [**MeasurementPeriodLteMdt**](MeasurementPeriodLteMdt.md) |  | [optional] 
**AreaScope** | Pointer to [**AreaScope**](AreaScope.md) |  | [optional] 
**PositioningMethod** | Pointer to [**PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] 
**AddPositioningMethodList** | Pointer to [**[]PositioningMethodMdt**](PositioningMethodMdt.md) |  | [optional] 
**MdtAllowedPlmnIdList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SensorMeasurementList** | Pointer to [**[]SensorMeasurement**](SensorMeasurement.md) |  | [optional] 

## Methods

### NewImmediateMdtConf

`func NewImmediateMdtConf(jobType JobType, ) *ImmediateMdtConf`

NewImmediateMdtConf instantiates a new ImmediateMdtConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImmediateMdtConfWithDefaults

`func NewImmediateMdtConfWithDefaults() *ImmediateMdtConf`

NewImmediateMdtConfWithDefaults instantiates a new ImmediateMdtConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *ImmediateMdtConf) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ImmediateMdtConf) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ImmediateMdtConf) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetMeasurementLteList

`func (o *ImmediateMdtConf) GetMeasurementLteList() []MeasurementLteForMdt`

GetMeasurementLteList returns the MeasurementLteList field if non-nil, zero value otherwise.

### GetMeasurementLteListOk

`func (o *ImmediateMdtConf) GetMeasurementLteListOk() (*[]MeasurementLteForMdt, bool)`

GetMeasurementLteListOk returns a tuple with the MeasurementLteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementLteList

`func (o *ImmediateMdtConf) SetMeasurementLteList(v []MeasurementLteForMdt)`

SetMeasurementLteList sets MeasurementLteList field to given value.

### HasMeasurementLteList

`func (o *ImmediateMdtConf) HasMeasurementLteList() bool`

HasMeasurementLteList returns a boolean if a field has been set.

### GetMeasurementNrList

`func (o *ImmediateMdtConf) GetMeasurementNrList() []MeasurementNrForMdt`

GetMeasurementNrList returns the MeasurementNrList field if non-nil, zero value otherwise.

### GetMeasurementNrListOk

`func (o *ImmediateMdtConf) GetMeasurementNrListOk() (*[]MeasurementNrForMdt, bool)`

GetMeasurementNrListOk returns a tuple with the MeasurementNrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementNrList

`func (o *ImmediateMdtConf) SetMeasurementNrList(v []MeasurementNrForMdt)`

SetMeasurementNrList sets MeasurementNrList field to given value.

### HasMeasurementNrList

`func (o *ImmediateMdtConf) HasMeasurementNrList() bool`

HasMeasurementNrList returns a boolean if a field has been set.

### GetReportingTriggerList

`func (o *ImmediateMdtConf) GetReportingTriggerList() []ReportingTrigger`

GetReportingTriggerList returns the ReportingTriggerList field if non-nil, zero value otherwise.

### GetReportingTriggerListOk

`func (o *ImmediateMdtConf) GetReportingTriggerListOk() (*[]ReportingTrigger, bool)`

GetReportingTriggerListOk returns a tuple with the ReportingTriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTriggerList

`func (o *ImmediateMdtConf) SetReportingTriggerList(v []ReportingTrigger)`

SetReportingTriggerList sets ReportingTriggerList field to given value.

### HasReportingTriggerList

`func (o *ImmediateMdtConf) HasReportingTriggerList() bool`

HasReportingTriggerList returns a boolean if a field has been set.

### GetReportInterval

`func (o *ImmediateMdtConf) GetReportInterval() ReportIntervalMdt`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *ImmediateMdtConf) GetReportIntervalOk() (*ReportIntervalMdt, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *ImmediateMdtConf) SetReportInterval(v ReportIntervalMdt)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *ImmediateMdtConf) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetReportIntervalNr

`func (o *ImmediateMdtConf) GetReportIntervalNr() ReportIntervalNrMdt`

GetReportIntervalNr returns the ReportIntervalNr field if non-nil, zero value otherwise.

### GetReportIntervalNrOk

`func (o *ImmediateMdtConf) GetReportIntervalNrOk() (*ReportIntervalNrMdt, bool)`

GetReportIntervalNrOk returns a tuple with the ReportIntervalNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalNr

`func (o *ImmediateMdtConf) SetReportIntervalNr(v ReportIntervalNrMdt)`

SetReportIntervalNr sets ReportIntervalNr field to given value.

### HasReportIntervalNr

`func (o *ImmediateMdtConf) HasReportIntervalNr() bool`

HasReportIntervalNr returns a boolean if a field has been set.

### GetReportAmount

`func (o *ImmediateMdtConf) GetReportAmount() ReportAmountMdt`

GetReportAmount returns the ReportAmount field if non-nil, zero value otherwise.

### GetReportAmountOk

`func (o *ImmediateMdtConf) GetReportAmountOk() (*ReportAmountMdt, bool)`

GetReportAmountOk returns a tuple with the ReportAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAmount

`func (o *ImmediateMdtConf) SetReportAmount(v ReportAmountMdt)`

SetReportAmount sets ReportAmount field to given value.

### HasReportAmount

`func (o *ImmediateMdtConf) HasReportAmount() bool`

HasReportAmount returns a boolean if a field has been set.

### GetEventThresholdRsrp

`func (o *ImmediateMdtConf) GetEventThresholdRsrp() int32`

GetEventThresholdRsrp returns the EventThresholdRsrp field if non-nil, zero value otherwise.

### GetEventThresholdRsrpOk

`func (o *ImmediateMdtConf) GetEventThresholdRsrpOk() (*int32, bool)`

GetEventThresholdRsrpOk returns a tuple with the EventThresholdRsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrp

`func (o *ImmediateMdtConf) SetEventThresholdRsrp(v int32)`

SetEventThresholdRsrp sets EventThresholdRsrp field to given value.

### HasEventThresholdRsrp

`func (o *ImmediateMdtConf) HasEventThresholdRsrp() bool`

HasEventThresholdRsrp returns a boolean if a field has been set.

### GetEventThresholdRsrq

`func (o *ImmediateMdtConf) GetEventThresholdRsrq() int32`

GetEventThresholdRsrq returns the EventThresholdRsrq field if non-nil, zero value otherwise.

### GetEventThresholdRsrqOk

`func (o *ImmediateMdtConf) GetEventThresholdRsrqOk() (*int32, bool)`

GetEventThresholdRsrqOk returns a tuple with the EventThresholdRsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrq

`func (o *ImmediateMdtConf) SetEventThresholdRsrq(v int32)`

SetEventThresholdRsrq sets EventThresholdRsrq field to given value.

### HasEventThresholdRsrq

`func (o *ImmediateMdtConf) HasEventThresholdRsrq() bool`

HasEventThresholdRsrq returns a boolean if a field has been set.

### GetEventThresholdRsrpNr

`func (o *ImmediateMdtConf) GetEventThresholdRsrpNr() int32`

GetEventThresholdRsrpNr returns the EventThresholdRsrpNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrpNrOk

`func (o *ImmediateMdtConf) GetEventThresholdRsrpNrOk() (*int32, bool)`

GetEventThresholdRsrpNrOk returns a tuple with the EventThresholdRsrpNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrpNr

`func (o *ImmediateMdtConf) SetEventThresholdRsrpNr(v int32)`

SetEventThresholdRsrpNr sets EventThresholdRsrpNr field to given value.

### HasEventThresholdRsrpNr

`func (o *ImmediateMdtConf) HasEventThresholdRsrpNr() bool`

HasEventThresholdRsrpNr returns a boolean if a field has been set.

### GetEventThresholdRsrqNr

`func (o *ImmediateMdtConf) GetEventThresholdRsrqNr() int32`

GetEventThresholdRsrqNr returns the EventThresholdRsrqNr field if non-nil, zero value otherwise.

### GetEventThresholdRsrqNrOk

`func (o *ImmediateMdtConf) GetEventThresholdRsrqNrOk() (*int32, bool)`

GetEventThresholdRsrqNrOk returns a tuple with the EventThresholdRsrqNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThresholdRsrqNr

`func (o *ImmediateMdtConf) SetEventThresholdRsrqNr(v int32)`

SetEventThresholdRsrqNr sets EventThresholdRsrqNr field to given value.

### HasEventThresholdRsrqNr

`func (o *ImmediateMdtConf) HasEventThresholdRsrqNr() bool`

HasEventThresholdRsrqNr returns a boolean if a field has been set.

### GetCollectionPeriodRmmLte

`func (o *ImmediateMdtConf) GetCollectionPeriodRmmLte() CollectionPeriodRmmLteMdt`

GetCollectionPeriodRmmLte returns the CollectionPeriodRmmLte field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmLteOk

`func (o *ImmediateMdtConf) GetCollectionPeriodRmmLteOk() (*CollectionPeriodRmmLteMdt, bool)`

GetCollectionPeriodRmmLteOk returns a tuple with the CollectionPeriodRmmLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmLte

`func (o *ImmediateMdtConf) SetCollectionPeriodRmmLte(v CollectionPeriodRmmLteMdt)`

SetCollectionPeriodRmmLte sets CollectionPeriodRmmLte field to given value.

### HasCollectionPeriodRmmLte

`func (o *ImmediateMdtConf) HasCollectionPeriodRmmLte() bool`

HasCollectionPeriodRmmLte returns a boolean if a field has been set.

### GetCollectionPeriodRmmNr

`func (o *ImmediateMdtConf) GetCollectionPeriodRmmNr() CollectionPeriodRmmNrMdt`

GetCollectionPeriodRmmNr returns the CollectionPeriodRmmNr field if non-nil, zero value otherwise.

### GetCollectionPeriodRmmNrOk

`func (o *ImmediateMdtConf) GetCollectionPeriodRmmNrOk() (*CollectionPeriodRmmNrMdt, bool)`

GetCollectionPeriodRmmNrOk returns a tuple with the CollectionPeriodRmmNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPeriodRmmNr

`func (o *ImmediateMdtConf) SetCollectionPeriodRmmNr(v CollectionPeriodRmmNrMdt)`

SetCollectionPeriodRmmNr sets CollectionPeriodRmmNr field to given value.

### HasCollectionPeriodRmmNr

`func (o *ImmediateMdtConf) HasCollectionPeriodRmmNr() bool`

HasCollectionPeriodRmmNr returns a boolean if a field has been set.

### GetMeasurementPeriodLte

`func (o *ImmediateMdtConf) GetMeasurementPeriodLte() MeasurementPeriodLteMdt`

GetMeasurementPeriodLte returns the MeasurementPeriodLte field if non-nil, zero value otherwise.

### GetMeasurementPeriodLteOk

`func (o *ImmediateMdtConf) GetMeasurementPeriodLteOk() (*MeasurementPeriodLteMdt, bool)`

GetMeasurementPeriodLteOk returns a tuple with the MeasurementPeriodLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriodLte

`func (o *ImmediateMdtConf) SetMeasurementPeriodLte(v MeasurementPeriodLteMdt)`

SetMeasurementPeriodLte sets MeasurementPeriodLte field to given value.

### HasMeasurementPeriodLte

`func (o *ImmediateMdtConf) HasMeasurementPeriodLte() bool`

HasMeasurementPeriodLte returns a boolean if a field has been set.

### GetAreaScope

`func (o *ImmediateMdtConf) GetAreaScope() AreaScope`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *ImmediateMdtConf) GetAreaScopeOk() (*AreaScope, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *ImmediateMdtConf) SetAreaScope(v AreaScope)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *ImmediateMdtConf) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.

### GetPositioningMethod

`func (o *ImmediateMdtConf) GetPositioningMethod() PositioningMethodMdt`

GetPositioningMethod returns the PositioningMethod field if non-nil, zero value otherwise.

### GetPositioningMethodOk

`func (o *ImmediateMdtConf) GetPositioningMethodOk() (*PositioningMethodMdt, bool)`

GetPositioningMethodOk returns a tuple with the PositioningMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningMethod

`func (o *ImmediateMdtConf) SetPositioningMethod(v PositioningMethodMdt)`

SetPositioningMethod sets PositioningMethod field to given value.

### HasPositioningMethod

`func (o *ImmediateMdtConf) HasPositioningMethod() bool`

HasPositioningMethod returns a boolean if a field has been set.

### GetAddPositioningMethodList

`func (o *ImmediateMdtConf) GetAddPositioningMethodList() []PositioningMethodMdt`

GetAddPositioningMethodList returns the AddPositioningMethodList field if non-nil, zero value otherwise.

### GetAddPositioningMethodListOk

`func (o *ImmediateMdtConf) GetAddPositioningMethodListOk() (*[]PositioningMethodMdt, bool)`

GetAddPositioningMethodListOk returns a tuple with the AddPositioningMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPositioningMethodList

`func (o *ImmediateMdtConf) SetAddPositioningMethodList(v []PositioningMethodMdt)`

SetAddPositioningMethodList sets AddPositioningMethodList field to given value.

### HasAddPositioningMethodList

`func (o *ImmediateMdtConf) HasAddPositioningMethodList() bool`

HasAddPositioningMethodList returns a boolean if a field has been set.

### GetMdtAllowedPlmnIdList

`func (o *ImmediateMdtConf) GetMdtAllowedPlmnIdList() []PlmnId`

GetMdtAllowedPlmnIdList returns the MdtAllowedPlmnIdList field if non-nil, zero value otherwise.

### GetMdtAllowedPlmnIdListOk

`func (o *ImmediateMdtConf) GetMdtAllowedPlmnIdListOk() (*[]PlmnId, bool)`

GetMdtAllowedPlmnIdListOk returns a tuple with the MdtAllowedPlmnIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtAllowedPlmnIdList

`func (o *ImmediateMdtConf) SetMdtAllowedPlmnIdList(v []PlmnId)`

SetMdtAllowedPlmnIdList sets MdtAllowedPlmnIdList field to given value.

### HasMdtAllowedPlmnIdList

`func (o *ImmediateMdtConf) HasMdtAllowedPlmnIdList() bool`

HasMdtAllowedPlmnIdList returns a boolean if a field has been set.

### GetSensorMeasurementList

`func (o *ImmediateMdtConf) GetSensorMeasurementList() []SensorMeasurement`

GetSensorMeasurementList returns the SensorMeasurementList field if non-nil, zero value otherwise.

### GetSensorMeasurementListOk

`func (o *ImmediateMdtConf) GetSensorMeasurementListOk() (*[]SensorMeasurement, bool)`

GetSensorMeasurementListOk returns a tuple with the SensorMeasurementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorMeasurementList

`func (o *ImmediateMdtConf) SetSensorMeasurementList(v []SensorMeasurement)`

SetSensorMeasurementList sets SensorMeasurementList field to given value.

### HasSensorMeasurementList

`func (o *ImmediateMdtConf) HasSensorMeasurementList() bool`

HasSensorMeasurementList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


