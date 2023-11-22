# AssignedEbiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** |  | 
**AssignedEbiList** | [**[]EbiArpMapping**](EbiArpMapping.md) |  | 
**FailedArpList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 
**ReleasedEbiList** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewAssignedEbiData

`func NewAssignedEbiData(pduSessionId int32, assignedEbiList []EbiArpMapping, ) *AssignedEbiData`

NewAssignedEbiData instantiates a new AssignedEbiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedEbiDataWithDefaults

`func NewAssignedEbiDataWithDefaults() *AssignedEbiData`

NewAssignedEbiDataWithDefaults instantiates a new AssignedEbiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *AssignedEbiData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *AssignedEbiData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *AssignedEbiData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetAssignedEbiList

`func (o *AssignedEbiData) GetAssignedEbiList() []EbiArpMapping`

GetAssignedEbiList returns the AssignedEbiList field if non-nil, zero value otherwise.

### GetAssignedEbiListOk

`func (o *AssignedEbiData) GetAssignedEbiListOk() (*[]EbiArpMapping, bool)`

GetAssignedEbiListOk returns a tuple with the AssignedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEbiList

`func (o *AssignedEbiData) SetAssignedEbiList(v []EbiArpMapping)`

SetAssignedEbiList sets AssignedEbiList field to given value.


### GetFailedArpList

`func (o *AssignedEbiData) GetFailedArpList() []Arp`

GetFailedArpList returns the FailedArpList field if non-nil, zero value otherwise.

### GetFailedArpListOk

`func (o *AssignedEbiData) GetFailedArpListOk() (*[]Arp, bool)`

GetFailedArpListOk returns a tuple with the FailedArpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedArpList

`func (o *AssignedEbiData) SetFailedArpList(v []Arp)`

SetFailedArpList sets FailedArpList field to given value.

### HasFailedArpList

`func (o *AssignedEbiData) HasFailedArpList() bool`

HasFailedArpList returns a boolean if a field has been set.

### GetReleasedEbiList

`func (o *AssignedEbiData) GetReleasedEbiList() []int32`

GetReleasedEbiList returns the ReleasedEbiList field if non-nil, zero value otherwise.

### GetReleasedEbiListOk

`func (o *AssignedEbiData) GetReleasedEbiListOk() (*[]int32, bool)`

GetReleasedEbiListOk returns a tuple with the ReleasedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedEbiList

`func (o *AssignedEbiData) SetReleasedEbiList(v []int32)`

SetReleasedEbiList sets ReleasedEbiList field to given value.

### HasReleasedEbiList

`func (o *AssignedEbiData) HasReleasedEbiList() bool`

HasReleasedEbiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


