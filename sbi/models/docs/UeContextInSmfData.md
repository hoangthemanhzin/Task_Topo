# UeContextInSmfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessions** | Pointer to [**map[string]PduSession**](PduSession.md) | A map (list of key-value pairs where PduSessionId serves as key) of PduSessions | [optional] 
**PgwInfo** | Pointer to [**[]PgwInfo**](PgwInfo.md) |  | [optional] 
**EmergencyInfo** | Pointer to [**EmergencyInfo**](EmergencyInfo.md) |  | [optional] 

## Methods

### NewUeContextInSmfData

`func NewUeContextInSmfData() *UeContextInSmfData`

NewUeContextInSmfData instantiates a new UeContextInSmfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextInSmfDataWithDefaults

`func NewUeContextInSmfDataWithDefaults() *UeContextInSmfData`

NewUeContextInSmfDataWithDefaults instantiates a new UeContextInSmfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessions

`func (o *UeContextInSmfData) GetPduSessions() map[string]PduSession`

GetPduSessions returns the PduSessions field if non-nil, zero value otherwise.

### GetPduSessionsOk

`func (o *UeContextInSmfData) GetPduSessionsOk() (*map[string]PduSession, bool)`

GetPduSessionsOk returns a tuple with the PduSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessions

`func (o *UeContextInSmfData) SetPduSessions(v map[string]PduSession)`

SetPduSessions sets PduSessions field to given value.

### HasPduSessions

`func (o *UeContextInSmfData) HasPduSessions() bool`

HasPduSessions returns a boolean if a field has been set.

### GetPgwInfo

`func (o *UeContextInSmfData) GetPgwInfo() []PgwInfo`

GetPgwInfo returns the PgwInfo field if non-nil, zero value otherwise.

### GetPgwInfoOk

`func (o *UeContextInSmfData) GetPgwInfoOk() (*[]PgwInfo, bool)`

GetPgwInfoOk returns a tuple with the PgwInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwInfo

`func (o *UeContextInSmfData) SetPgwInfo(v []PgwInfo)`

SetPgwInfo sets PgwInfo field to given value.

### HasPgwInfo

`func (o *UeContextInSmfData) HasPgwInfo() bool`

HasPgwInfo returns a boolean if a field has been set.

### GetEmergencyInfo

`func (o *UeContextInSmfData) GetEmergencyInfo() EmergencyInfo`

GetEmergencyInfo returns the EmergencyInfo field if non-nil, zero value otherwise.

### GetEmergencyInfoOk

`func (o *UeContextInSmfData) GetEmergencyInfoOk() (*EmergencyInfo, bool)`

GetEmergencyInfoOk returns a tuple with the EmergencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyInfo

`func (o *UeContextInSmfData) SetEmergencyInfo(v EmergencyInfo)`

SetEmergencyInfo sets EmergencyInfo field to given value.

### HasEmergencyInfo

`func (o *UeContextInSmfData) HasEmergencyInfo() bool`

HasEmergencyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


