# ReportingArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaType** | [**ReportingAreaType**](ReportingAreaType.md) |  | 
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 

## Methods

### NewReportingArea

`func NewReportingArea(areaType ReportingAreaType, ) *ReportingArea`

NewReportingArea instantiates a new ReportingArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingAreaWithDefaults

`func NewReportingAreaWithDefaults() *ReportingArea`

NewReportingAreaWithDefaults instantiates a new ReportingArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaType

`func (o *ReportingArea) GetAreaType() ReportingAreaType`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *ReportingArea) GetAreaTypeOk() (*ReportingAreaType, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *ReportingArea) SetAreaType(v ReportingAreaType)`

SetAreaType sets AreaType field to given value.


### GetTai

`func (o *ReportingArea) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *ReportingArea) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *ReportingArea) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *ReportingArea) HasTai() bool`

HasTai returns a boolean if a field has been set.

### GetEcgi

`func (o *ReportingArea) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *ReportingArea) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *ReportingArea) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *ReportingArea) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *ReportingArea) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *ReportingArea) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *ReportingArea) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *ReportingArea) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


