# PcEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**PcEvent**](PcEvent.md) |  | 
**AccType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AddAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**RelAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**AnGwAddr** | Pointer to [**AnGwAddress**](AnGwAddress.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**Supi** | Pointer to **string** |  | [optional] 
**Gpsi** | Pointer to **string** |  | [optional] 
**TimeStamp** | **time.Time** |  | 
**PduSessionInfo** | Pointer to [**PduSessionInformation**](PduSessionInformation.md) |  | [optional] 
**RepServices** | Pointer to [**ServiceIdentification**](ServiceIdentification.md) |  | [optional] 

## Methods

### NewPcEventNotification

`func NewPcEventNotification(event PcEvent, timeStamp time.Time, ) *PcEventNotification`

NewPcEventNotification instantiates a new PcEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcEventNotificationWithDefaults

`func NewPcEventNotificationWithDefaults() *PcEventNotification`

NewPcEventNotificationWithDefaults instantiates a new PcEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PcEventNotification) GetEvent() PcEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PcEventNotification) GetEventOk() (*PcEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PcEventNotification) SetEvent(v PcEvent)`

SetEvent sets Event field to given value.


### GetAccType

`func (o *PcEventNotification) GetAccType() AccessType`

GetAccType returns the AccType field if non-nil, zero value otherwise.

### GetAccTypeOk

`func (o *PcEventNotification) GetAccTypeOk() (*AccessType, bool)`

GetAccTypeOk returns a tuple with the AccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccType

`func (o *PcEventNotification) SetAccType(v AccessType)`

SetAccType sets AccType field to given value.

### HasAccType

`func (o *PcEventNotification) HasAccType() bool`

HasAccType returns a boolean if a field has been set.

### GetAddAccessInfo

`func (o *PcEventNotification) GetAddAccessInfo() AdditionalAccessInfo`

GetAddAccessInfo returns the AddAccessInfo field if non-nil, zero value otherwise.

### GetAddAccessInfoOk

`func (o *PcEventNotification) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetAddAccessInfoOk returns a tuple with the AddAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessInfo

`func (o *PcEventNotification) SetAddAccessInfo(v AdditionalAccessInfo)`

SetAddAccessInfo sets AddAccessInfo field to given value.

### HasAddAccessInfo

`func (o *PcEventNotification) HasAddAccessInfo() bool`

HasAddAccessInfo returns a boolean if a field has been set.

### GetRelAccessInfo

`func (o *PcEventNotification) GetRelAccessInfo() AdditionalAccessInfo`

GetRelAccessInfo returns the RelAccessInfo field if non-nil, zero value otherwise.

### GetRelAccessInfoOk

`func (o *PcEventNotification) GetRelAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetRelAccessInfoOk returns a tuple with the RelAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelAccessInfo

`func (o *PcEventNotification) SetRelAccessInfo(v AdditionalAccessInfo)`

SetRelAccessInfo sets RelAccessInfo field to given value.

### HasRelAccessInfo

`func (o *PcEventNotification) HasRelAccessInfo() bool`

HasRelAccessInfo returns a boolean if a field has been set.

### GetAnGwAddr

`func (o *PcEventNotification) GetAnGwAddr() AnGwAddress`

GetAnGwAddr returns the AnGwAddr field if non-nil, zero value otherwise.

### GetAnGwAddrOk

`func (o *PcEventNotification) GetAnGwAddrOk() (*AnGwAddress, bool)`

GetAnGwAddrOk returns a tuple with the AnGwAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnGwAddr

`func (o *PcEventNotification) SetAnGwAddr(v AnGwAddress)`

SetAnGwAddr sets AnGwAddr field to given value.

### HasAnGwAddr

`func (o *PcEventNotification) HasAnGwAddr() bool`

HasAnGwAddr returns a boolean if a field has been set.

### GetRatType

`func (o *PcEventNotification) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *PcEventNotification) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *PcEventNotification) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *PcEventNotification) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetPlmnId

`func (o *PcEventNotification) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PcEventNotification) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PcEventNotification) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PcEventNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetSupi

`func (o *PcEventNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PcEventNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PcEventNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PcEventNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *PcEventNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PcEventNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PcEventNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PcEventNotification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetTimeStamp

`func (o *PcEventNotification) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PcEventNotification) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PcEventNotification) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetPduSessionInfo

`func (o *PcEventNotification) GetPduSessionInfo() PduSessionInformation`

GetPduSessionInfo returns the PduSessionInfo field if non-nil, zero value otherwise.

### GetPduSessionInfoOk

`func (o *PcEventNotification) GetPduSessionInfoOk() (*PduSessionInformation, bool)`

GetPduSessionInfoOk returns a tuple with the PduSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionInfo

`func (o *PcEventNotification) SetPduSessionInfo(v PduSessionInformation)`

SetPduSessionInfo sets PduSessionInfo field to given value.

### HasPduSessionInfo

`func (o *PcEventNotification) HasPduSessionInfo() bool`

HasPduSessionInfo returns a boolean if a field has been set.

### GetRepServices

`func (o *PcEventNotification) GetRepServices() ServiceIdentification`

GetRepServices returns the RepServices field if non-nil, zero value otherwise.

### GetRepServicesOk

`func (o *PcEventNotification) GetRepServicesOk() (*ServiceIdentification, bool)`

GetRepServicesOk returns a tuple with the RepServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepServices

`func (o *PcEventNotification) SetRepServices(v ServiceIdentification)`

SetRepServices sets RepServices field to given value.

### HasRepServices

`func (o *PcEventNotification) HasRepServices() bool`

HasRepServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


