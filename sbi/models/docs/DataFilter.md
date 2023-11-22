# DataFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataInd** | [**DataInd**](DataInd.md) |  | 
**Dnns** | Pointer to **[]string** |  | [optional] 
**Snssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**InternalGroupIds** | Pointer to **[]string** |  | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**UeIpv4s** | Pointer to **[]string** |  | [optional] 
**UeIpv6s** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**UeMacs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDataFilter

`func NewDataFilter(dataInd DataInd, ) *DataFilter`

NewDataFilter instantiates a new DataFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFilterWithDefaults

`func NewDataFilterWithDefaults() *DataFilter`

NewDataFilterWithDefaults instantiates a new DataFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataInd

`func (o *DataFilter) GetDataInd() DataInd`

GetDataInd returns the DataInd field if non-nil, zero value otherwise.

### GetDataIndOk

`func (o *DataFilter) GetDataIndOk() (*DataInd, bool)`

GetDataIndOk returns a tuple with the DataInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInd

`func (o *DataFilter) SetDataInd(v DataInd)`

SetDataInd sets DataInd field to given value.


### GetDnns

`func (o *DataFilter) GetDnns() []string`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *DataFilter) GetDnnsOk() (*[]string, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *DataFilter) SetDnns(v []string)`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *DataFilter) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetSnssais

`func (o *DataFilter) GetSnssais() []Snssai`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *DataFilter) GetSnssaisOk() (*[]Snssai, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *DataFilter) SetSnssais(v []Snssai)`

SetSnssais sets Snssais field to given value.

### HasSnssais

`func (o *DataFilter) HasSnssais() bool`

HasSnssais returns a boolean if a field has been set.

### GetInternalGroupIds

`func (o *DataFilter) GetInternalGroupIds() []string`

GetInternalGroupIds returns the InternalGroupIds field if non-nil, zero value otherwise.

### GetInternalGroupIdsOk

`func (o *DataFilter) GetInternalGroupIdsOk() (*[]string, bool)`

GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIds

`func (o *DataFilter) SetInternalGroupIds(v []string)`

SetInternalGroupIds sets InternalGroupIds field to given value.

### HasInternalGroupIds

`func (o *DataFilter) HasInternalGroupIds() bool`

HasInternalGroupIds returns a boolean if a field has been set.

### GetSupis

`func (o *DataFilter) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *DataFilter) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *DataFilter) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *DataFilter) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetAppIds

`func (o *DataFilter) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *DataFilter) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *DataFilter) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *DataFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetUeIpv4s

`func (o *DataFilter) GetUeIpv4s() []string`

GetUeIpv4s returns the UeIpv4s field if non-nil, zero value otherwise.

### GetUeIpv4sOk

`func (o *DataFilter) GetUeIpv4sOk() (*[]string, bool)`

GetUeIpv4sOk returns a tuple with the UeIpv4s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4s

`func (o *DataFilter) SetUeIpv4s(v []string)`

SetUeIpv4s sets UeIpv4s field to given value.

### HasUeIpv4s

`func (o *DataFilter) HasUeIpv4s() bool`

HasUeIpv4s returns a boolean if a field has been set.

### GetUeIpv6s

`func (o *DataFilter) GetUeIpv6s() []Ipv6Addr`

GetUeIpv6s returns the UeIpv6s field if non-nil, zero value otherwise.

### GetUeIpv6sOk

`func (o *DataFilter) GetUeIpv6sOk() (*[]Ipv6Addr, bool)`

GetUeIpv6sOk returns a tuple with the UeIpv6s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6s

`func (o *DataFilter) SetUeIpv6s(v []Ipv6Addr)`

SetUeIpv6s sets UeIpv6s field to given value.

### HasUeIpv6s

`func (o *DataFilter) HasUeIpv6s() bool`

HasUeIpv6s returns a boolean if a field has been set.

### GetUeMacs

`func (o *DataFilter) GetUeMacs() []string`

GetUeMacs returns the UeMacs field if non-nil, zero value otherwise.

### GetUeMacsOk

`func (o *DataFilter) GetUeMacsOk() (*[]string, bool)`

GetUeMacsOk returns a tuple with the UeMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMacs

`func (o *DataFilter) SetUeMacs(v []string)`

SetUeMacs sets UeMacs field to given value.

### HasUeMacs

`func (o *DataFilter) HasUeMacs() bool`

HasUeMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


