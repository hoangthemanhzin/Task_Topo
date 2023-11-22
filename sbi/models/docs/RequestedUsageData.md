# RequestedUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefUmIds** | Pointer to **[]string** | An array of usage monitoring data id references to the usage monitoring data instances for which the PCF is requesting a usage report. This attribute shall only be provided when allUmIds is not set to true. | [optional] 
**AllUmIds** | Pointer to **bool** | This boolean indicates whether requested usage data applies to all usage monitoring data instances. When it&#39;s not included, it means requested usage data shall only apply to the usage monitoring data instances referenced by the refUmIds attribute. | [optional] 

## Methods

### NewRequestedUsageData

`func NewRequestedUsageData() *RequestedUsageData`

NewRequestedUsageData instantiates a new RequestedUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedUsageDataWithDefaults

`func NewRequestedUsageDataWithDefaults() *RequestedUsageData`

NewRequestedUsageDataWithDefaults instantiates a new RequestedUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefUmIds

`func (o *RequestedUsageData) GetRefUmIds() []string`

GetRefUmIds returns the RefUmIds field if non-nil, zero value otherwise.

### GetRefUmIdsOk

`func (o *RequestedUsageData) GetRefUmIdsOk() (*[]string, bool)`

GetRefUmIdsOk returns a tuple with the RefUmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUmIds

`func (o *RequestedUsageData) SetRefUmIds(v []string)`

SetRefUmIds sets RefUmIds field to given value.

### HasRefUmIds

`func (o *RequestedUsageData) HasRefUmIds() bool`

HasRefUmIds returns a boolean if a field has been set.

### GetAllUmIds

`func (o *RequestedUsageData) GetAllUmIds() bool`

GetAllUmIds returns the AllUmIds field if non-nil, zero value otherwise.

### GetAllUmIdsOk

`func (o *RequestedUsageData) GetAllUmIdsOk() (*bool, bool)`

GetAllUmIdsOk returns a tuple with the AllUmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUmIds

`func (o *RequestedUsageData) SetAllUmIds(v bool)`

SetAllUmIds sets AllUmIds field to given value.

### HasAllUmIds

`func (o *RequestedUsageData) HasAllUmIds() bool`

HasAllUmIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


