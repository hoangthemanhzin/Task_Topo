# NrppaInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfId** | **string** |  | 
**NrppaPdu** | [**N2InfoContent**](N2InfoContent.md) |  | 
**ServiceInstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewNrppaInformation

`func NewNrppaInformation(nfId string, nrppaPdu N2InfoContent, ) *NrppaInformation`

NewNrppaInformation instantiates a new NrppaInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrppaInformationWithDefaults

`func NewNrppaInformationWithDefaults() *NrppaInformation`

NewNrppaInformationWithDefaults instantiates a new NrppaInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfId

`func (o *NrppaInformation) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *NrppaInformation) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *NrppaInformation) SetNfId(v string)`

SetNfId sets NfId field to given value.


### GetNrppaPdu

`func (o *NrppaInformation) GetNrppaPdu() N2InfoContent`

GetNrppaPdu returns the NrppaPdu field if non-nil, zero value otherwise.

### GetNrppaPduOk

`func (o *NrppaInformation) GetNrppaPduOk() (*N2InfoContent, bool)`

GetNrppaPduOk returns a tuple with the NrppaPdu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrppaPdu

`func (o *NrppaInformation) SetNrppaPdu(v N2InfoContent)`

SetNrppaPdu sets NrppaPdu field to given value.


### GetServiceInstanceId

`func (o *NrppaInformation) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *NrppaInformation) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *NrppaInformation) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *NrppaInformation) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


