# AfExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | Pointer to **string** |  | [optional] 
**AllowedGeographicArea** | Pointer to [**[]GeographicArea**](GeographicArea.md) |  | [optional] 
**PrivacyCheckRelatedAction** | Pointer to [**PrivacyCheckRelatedAction**](PrivacyCheckRelatedAction.md) |  | [optional] 
**ValidTimePeriod** | Pointer to [**ValidTimePeriod**](ValidTimePeriod.md) |  | [optional] 

## Methods

### NewAfExternal

`func NewAfExternal() *AfExternal`

NewAfExternal instantiates a new AfExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfExternalWithDefaults

`func NewAfExternalWithDefaults() *AfExternal`

NewAfExternalWithDefaults instantiates a new AfExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *AfExternal) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *AfExternal) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *AfExternal) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *AfExternal) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetAllowedGeographicArea

`func (o *AfExternal) GetAllowedGeographicArea() []GeographicArea`

GetAllowedGeographicArea returns the AllowedGeographicArea field if non-nil, zero value otherwise.

### GetAllowedGeographicAreaOk

`func (o *AfExternal) GetAllowedGeographicAreaOk() (*[]GeographicArea, bool)`

GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGeographicArea

`func (o *AfExternal) SetAllowedGeographicArea(v []GeographicArea)`

SetAllowedGeographicArea sets AllowedGeographicArea field to given value.

### HasAllowedGeographicArea

`func (o *AfExternal) HasAllowedGeographicArea() bool`

HasAllowedGeographicArea returns a boolean if a field has been set.

### GetPrivacyCheckRelatedAction

`func (o *AfExternal) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction`

GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field if non-nil, zero value otherwise.

### GetPrivacyCheckRelatedActionOk

`func (o *AfExternal) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool)`

GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyCheckRelatedAction

`func (o *AfExternal) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction)`

SetPrivacyCheckRelatedAction sets PrivacyCheckRelatedAction field to given value.

### HasPrivacyCheckRelatedAction

`func (o *AfExternal) HasPrivacyCheckRelatedAction() bool`

HasPrivacyCheckRelatedAction returns a boolean if a field has been set.

### GetValidTimePeriod

`func (o *AfExternal) GetValidTimePeriod() ValidTimePeriod`

GetValidTimePeriod returns the ValidTimePeriod field if non-nil, zero value otherwise.

### GetValidTimePeriodOk

`func (o *AfExternal) GetValidTimePeriodOk() (*ValidTimePeriod, bool)`

GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimePeriod

`func (o *AfExternal) SetValidTimePeriod(v ValidTimePeriod)`

SetValidTimePeriod sets ValidTimePeriod field to given value.

### HasValidTimePeriod

`func (o *AfExternal) HasValidTimePeriod() bool`

HasValidTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


