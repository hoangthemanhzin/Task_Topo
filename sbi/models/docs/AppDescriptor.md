# AppDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewAppDescriptor

`func NewAppDescriptor() *AppDescriptor`

NewAppDescriptor instantiates a new AppDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDescriptorWithDefaults

`func NewAppDescriptorWithDefaults() *AppDescriptor`

NewAppDescriptorWithDefaults instantiates a new AppDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsId

`func (o *AppDescriptor) GetOsId() string`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *AppDescriptor) GetOsIdOk() (*string, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *AppDescriptor) SetOsId(v string)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *AppDescriptor) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetAppId

`func (o *AppDescriptor) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppDescriptor) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppDescriptor) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppDescriptor) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


