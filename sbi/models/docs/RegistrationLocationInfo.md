# RegistrationLocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** |  | 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**VgmlcAddress** | Pointer to [**VgmlcAddress**](VgmlcAddress.md) |  | [optional] 
**AccessTypeList** | [**[]AccessType**](AccessType.md) |  | 

## Methods

### NewRegistrationLocationInfo

`func NewRegistrationLocationInfo(amfInstanceId string, accessTypeList []AccessType, ) *RegistrationLocationInfo`

NewRegistrationLocationInfo instantiates a new RegistrationLocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationLocationInfoWithDefaults

`func NewRegistrationLocationInfoWithDefaults() *RegistrationLocationInfo`

NewRegistrationLocationInfoWithDefaults instantiates a new RegistrationLocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *RegistrationLocationInfo) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *RegistrationLocationInfo) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *RegistrationLocationInfo) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetPlmnId

`func (o *RegistrationLocationInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *RegistrationLocationInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *RegistrationLocationInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *RegistrationLocationInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetVgmlcAddress

`func (o *RegistrationLocationInfo) GetVgmlcAddress() VgmlcAddress`

GetVgmlcAddress returns the VgmlcAddress field if non-nil, zero value otherwise.

### GetVgmlcAddressOk

`func (o *RegistrationLocationInfo) GetVgmlcAddressOk() (*VgmlcAddress, bool)`

GetVgmlcAddressOk returns a tuple with the VgmlcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddress

`func (o *RegistrationLocationInfo) SetVgmlcAddress(v VgmlcAddress)`

SetVgmlcAddress sets VgmlcAddress field to given value.

### HasVgmlcAddress

`func (o *RegistrationLocationInfo) HasVgmlcAddress() bool`

HasVgmlcAddress returns a boolean if a field has been set.

### GetAccessTypeList

`func (o *RegistrationLocationInfo) GetAccessTypeList() []AccessType`

GetAccessTypeList returns the AccessTypeList field if non-nil, zero value otherwise.

### GetAccessTypeListOk

`func (o *RegistrationLocationInfo) GetAccessTypeListOk() (*[]AccessType, bool)`

GetAccessTypeListOk returns a tuple with the AccessTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTypeList

`func (o *RegistrationLocationInfo) SetAccessTypeList(v []AccessType)`

SetAccessTypeList sets AccessTypeList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


