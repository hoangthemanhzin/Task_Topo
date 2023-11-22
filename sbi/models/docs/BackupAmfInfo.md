# BackupAmfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupAmf** | **string** |  | 
**GuamiList** | Pointer to [**[]Guami**](Guami.md) |  | [optional] 

## Methods

### NewBackupAmfInfo

`func NewBackupAmfInfo(backupAmf string, ) *BackupAmfInfo`

NewBackupAmfInfo instantiates a new BackupAmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAmfInfoWithDefaults

`func NewBackupAmfInfoWithDefaults() *BackupAmfInfo`

NewBackupAmfInfoWithDefaults instantiates a new BackupAmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupAmf

`func (o *BackupAmfInfo) GetBackupAmf() string`

GetBackupAmf returns the BackupAmf field if non-nil, zero value otherwise.

### GetBackupAmfOk

`func (o *BackupAmfInfo) GetBackupAmfOk() (*string, bool)`

GetBackupAmfOk returns a tuple with the BackupAmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmf

`func (o *BackupAmfInfo) SetBackupAmf(v string)`

SetBackupAmf sets BackupAmf field to given value.


### GetGuamiList

`func (o *BackupAmfInfo) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *BackupAmfInfo) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *BackupAmfInfo) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.

### HasGuamiList

`func (o *BackupAmfInfo) HasGuamiList() bool`

HasGuamiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


