# IptvConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** |  | [optional] 
**InterGroupId** | Pointer to **string** |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AfAppId** | **string** |  | 
**MultiAccCtrls** | [**map[string]MulticastAccessControl**](MulticastAccessControl.md) |  | 
**SuppFeat** | Pointer to **string** |  | [optional] 
**ResUri** | Pointer to **string** |  | [optional] 

## Methods

### NewIptvConfigData

`func NewIptvConfigData(afAppId string, multiAccCtrls map[string]MulticastAccessControl, ) *IptvConfigData`

NewIptvConfigData instantiates a new IptvConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvConfigDataWithDefaults

`func NewIptvConfigDataWithDefaults() *IptvConfigData`

NewIptvConfigDataWithDefaults instantiates a new IptvConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *IptvConfigData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *IptvConfigData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *IptvConfigData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *IptvConfigData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetInterGroupId

`func (o *IptvConfigData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *IptvConfigData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *IptvConfigData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *IptvConfigData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetDnn

`func (o *IptvConfigData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *IptvConfigData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *IptvConfigData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *IptvConfigData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *IptvConfigData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *IptvConfigData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *IptvConfigData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *IptvConfigData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetAfAppId

`func (o *IptvConfigData) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *IptvConfigData) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *IptvConfigData) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.


### GetMultiAccCtrls

`func (o *IptvConfigData) GetMultiAccCtrls() map[string]MulticastAccessControl`

GetMultiAccCtrls returns the MultiAccCtrls field if non-nil, zero value otherwise.

### GetMultiAccCtrlsOk

`func (o *IptvConfigData) GetMultiAccCtrlsOk() (*map[string]MulticastAccessControl, bool)`

GetMultiAccCtrlsOk returns a tuple with the MultiAccCtrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAccCtrls

`func (o *IptvConfigData) SetMultiAccCtrls(v map[string]MulticastAccessControl)`

SetMultiAccCtrls sets MultiAccCtrls field to given value.


### GetSuppFeat

`func (o *IptvConfigData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *IptvConfigData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *IptvConfigData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *IptvConfigData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResUri

`func (o *IptvConfigData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *IptvConfigData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *IptvConfigData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *IptvConfigData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


