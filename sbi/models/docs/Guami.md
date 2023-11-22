# Guami

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**AmfId** | **string** |  | 

## Methods

### NewGuami

`func NewGuami(plmnId PlmnIdNid, amfId string, ) *Guami`

NewGuami instantiates a new Guami object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuamiWithDefaults

`func NewGuamiWithDefaults() *Guami`

NewGuamiWithDefaults instantiates a new Guami object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *Guami) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Guami) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Guami) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.


### GetAmfId

`func (o *Guami) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *Guami) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *Guami) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


