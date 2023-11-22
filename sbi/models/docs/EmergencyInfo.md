# EmergencyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgwFqdn** | Pointer to **string** |  | [optional] 
**PgwIpAddress** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**SmfInstanceId** | Pointer to **string** |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEmergencyInfo

`func NewEmergencyInfo() *EmergencyInfo`

NewEmergencyInfo instantiates a new EmergencyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmergencyInfoWithDefaults

`func NewEmergencyInfoWithDefaults() *EmergencyInfo`

NewEmergencyInfoWithDefaults instantiates a new EmergencyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgwFqdn

`func (o *EmergencyInfo) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *EmergencyInfo) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *EmergencyInfo) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *EmergencyInfo) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetPgwIpAddress

`func (o *EmergencyInfo) GetPgwIpAddress() IpAddress`

GetPgwIpAddress returns the PgwIpAddress field if non-nil, zero value otherwise.

### GetPgwIpAddressOk

`func (o *EmergencyInfo) GetPgwIpAddressOk() (*IpAddress, bool)`

GetPgwIpAddressOk returns a tuple with the PgwIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddress

`func (o *EmergencyInfo) SetPgwIpAddress(v IpAddress)`

SetPgwIpAddress sets PgwIpAddress field to given value.

### HasPgwIpAddress

`func (o *EmergencyInfo) HasPgwIpAddress() bool`

HasPgwIpAddress returns a boolean if a field has been set.

### GetSmfInstanceId

`func (o *EmergencyInfo) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *EmergencyInfo) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *EmergencyInfo) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.

### HasSmfInstanceId

`func (o *EmergencyInfo) HasSmfInstanceId() bool`

HasSmfInstanceId returns a boolean if a field has been set.

### GetEpdgInd

`func (o *EmergencyInfo) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *EmergencyInfo) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *EmergencyInfo) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *EmergencyInfo) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


