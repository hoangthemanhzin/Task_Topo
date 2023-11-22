# AmfStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuamiList** | [**[]Guami**](Guami.md) |  | 
**StatusChange** | [**StatusChange**](StatusChange.md) |  | 
**TargetAmfRemoval** | Pointer to **string** |  | [optional] 
**TargetAmfFailure** | Pointer to **string** |  | [optional] 

## Methods

### NewAmfStatusInfo

`func NewAmfStatusInfo(guamiList []Guami, statusChange StatusChange, ) *AmfStatusInfo`

NewAmfStatusInfo instantiates a new AmfStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfStatusInfoWithDefaults

`func NewAmfStatusInfoWithDefaults() *AmfStatusInfo`

NewAmfStatusInfoWithDefaults instantiates a new AmfStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuamiList

`func (o *AmfStatusInfo) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *AmfStatusInfo) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *AmfStatusInfo) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.


### GetStatusChange

`func (o *AmfStatusInfo) GetStatusChange() StatusChange`

GetStatusChange returns the StatusChange field if non-nil, zero value otherwise.

### GetStatusChangeOk

`func (o *AmfStatusInfo) GetStatusChangeOk() (*StatusChange, bool)`

GetStatusChangeOk returns a tuple with the StatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChange

`func (o *AmfStatusInfo) SetStatusChange(v StatusChange)`

SetStatusChange sets StatusChange field to given value.


### GetTargetAmfRemoval

`func (o *AmfStatusInfo) GetTargetAmfRemoval() string`

GetTargetAmfRemoval returns the TargetAmfRemoval field if non-nil, zero value otherwise.

### GetTargetAmfRemovalOk

`func (o *AmfStatusInfo) GetTargetAmfRemovalOk() (*string, bool)`

GetTargetAmfRemovalOk returns a tuple with the TargetAmfRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmfRemoval

`func (o *AmfStatusInfo) SetTargetAmfRemoval(v string)`

SetTargetAmfRemoval sets TargetAmfRemoval field to given value.

### HasTargetAmfRemoval

`func (o *AmfStatusInfo) HasTargetAmfRemoval() bool`

HasTargetAmfRemoval returns a boolean if a field has been set.

### GetTargetAmfFailure

`func (o *AmfStatusInfo) GetTargetAmfFailure() string`

GetTargetAmfFailure returns the TargetAmfFailure field if non-nil, zero value otherwise.

### GetTargetAmfFailureOk

`func (o *AmfStatusInfo) GetTargetAmfFailureOk() (*string, bool)`

GetTargetAmfFailureOk returns a tuple with the TargetAmfFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmfFailure

`func (o *AmfStatusInfo) SetTargetAmfFailure(v string)`

SetTargetAmfFailure sets TargetAmfFailure field to given value.

### HasTargetAmfFailure

`func (o *AmfStatusInfo) HasTargetAmfFailure() bool`

HasTargetAmfFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


