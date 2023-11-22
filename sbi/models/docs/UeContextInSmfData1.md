# UeContextInSmfData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessions** | Pointer to [**map[string]PduSession1**](PduSession1.md) | A map (list of key-value pairs where PduSessionId serves as key) of PduSessions | [optional] 
**PgwInfo** | Pointer to [**[]PgwInfo1**](PgwInfo1.md) |  | [optional] 
**EmergencyInfo** | Pointer to [**EmergencyInfo1**](EmergencyInfo1.md) |  | [optional] 

## Methods

### NewUeContextInSmfData1

`func NewUeContextInSmfData1() *UeContextInSmfData1`

NewUeContextInSmfData1 instantiates a new UeContextInSmfData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextInSmfData1WithDefaults

`func NewUeContextInSmfData1WithDefaults() *UeContextInSmfData1`

NewUeContextInSmfData1WithDefaults instantiates a new UeContextInSmfData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessions

`func (o *UeContextInSmfData1) GetPduSessions() map[string]PduSession1`

GetPduSessions returns the PduSessions field if non-nil, zero value otherwise.

### GetPduSessionsOk

`func (o *UeContextInSmfData1) GetPduSessionsOk() (*map[string]PduSession1, bool)`

GetPduSessionsOk returns a tuple with the PduSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessions

`func (o *UeContextInSmfData1) SetPduSessions(v map[string]PduSession1)`

SetPduSessions sets PduSessions field to given value.

### HasPduSessions

`func (o *UeContextInSmfData1) HasPduSessions() bool`

HasPduSessions returns a boolean if a field has been set.

### GetPgwInfo

`func (o *UeContextInSmfData1) GetPgwInfo() []PgwInfo1`

GetPgwInfo returns the PgwInfo field if non-nil, zero value otherwise.

### GetPgwInfoOk

`func (o *UeContextInSmfData1) GetPgwInfoOk() (*[]PgwInfo1, bool)`

GetPgwInfoOk returns a tuple with the PgwInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwInfo

`func (o *UeContextInSmfData1) SetPgwInfo(v []PgwInfo1)`

SetPgwInfo sets PgwInfo field to given value.

### HasPgwInfo

`func (o *UeContextInSmfData1) HasPgwInfo() bool`

HasPgwInfo returns a boolean if a field has been set.

### GetEmergencyInfo

`func (o *UeContextInSmfData1) GetEmergencyInfo() EmergencyInfo1`

GetEmergencyInfo returns the EmergencyInfo field if non-nil, zero value otherwise.

### GetEmergencyInfoOk

`func (o *UeContextInSmfData1) GetEmergencyInfoOk() (*EmergencyInfo1, bool)`

GetEmergencyInfoOk returns a tuple with the EmergencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyInfo

`func (o *UeContextInSmfData1) SetEmergencyInfo(v EmergencyInfo1)`

SetEmergencyInfo sets EmergencyInfo field to given value.

### HasEmergencyInfo

`func (o *UeContextInSmfData1) HasEmergencyInfo() bool`

HasEmergencyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


