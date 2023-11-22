# TraceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceRef** | **string** |  | 
**TraceDepth** | [**TraceDepth**](TraceDepth.md) |  | 
**NeTypeList** | **string** |  | 
**EventList** | **string** |  | 
**CollectionEntityIpv4Addr** | Pointer to **string** |  | [optional] 
**CollectionEntityIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**InterfaceList** | Pointer to **string** |  | [optional] 

## Methods

### NewTraceData

`func NewTraceData(traceRef string, traceDepth TraceDepth, neTypeList string, eventList string, ) *TraceData`

NewTraceData instantiates a new TraceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceDataWithDefaults

`func NewTraceDataWithDefaults() *TraceData`

NewTraceDataWithDefaults instantiates a new TraceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceRef

`func (o *TraceData) GetTraceRef() string`

GetTraceRef returns the TraceRef field if non-nil, zero value otherwise.

### GetTraceRefOk

`func (o *TraceData) GetTraceRefOk() (*string, bool)`

GetTraceRefOk returns a tuple with the TraceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRef

`func (o *TraceData) SetTraceRef(v string)`

SetTraceRef sets TraceRef field to given value.


### GetTraceDepth

`func (o *TraceData) GetTraceDepth() TraceDepth`

GetTraceDepth returns the TraceDepth field if non-nil, zero value otherwise.

### GetTraceDepthOk

`func (o *TraceData) GetTraceDepthOk() (*TraceDepth, bool)`

GetTraceDepthOk returns a tuple with the TraceDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceDepth

`func (o *TraceData) SetTraceDepth(v TraceDepth)`

SetTraceDepth sets TraceDepth field to given value.


### GetNeTypeList

`func (o *TraceData) GetNeTypeList() string`

GetNeTypeList returns the NeTypeList field if non-nil, zero value otherwise.

### GetNeTypeListOk

`func (o *TraceData) GetNeTypeListOk() (*string, bool)`

GetNeTypeListOk returns a tuple with the NeTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeTypeList

`func (o *TraceData) SetNeTypeList(v string)`

SetNeTypeList sets NeTypeList field to given value.


### GetEventList

`func (o *TraceData) GetEventList() string`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *TraceData) GetEventListOk() (*string, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *TraceData) SetEventList(v string)`

SetEventList sets EventList field to given value.


### GetCollectionEntityIpv4Addr

`func (o *TraceData) GetCollectionEntityIpv4Addr() string`

GetCollectionEntityIpv4Addr returns the CollectionEntityIpv4Addr field if non-nil, zero value otherwise.

### GetCollectionEntityIpv4AddrOk

`func (o *TraceData) GetCollectionEntityIpv4AddrOk() (*string, bool)`

GetCollectionEntityIpv4AddrOk returns a tuple with the CollectionEntityIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionEntityIpv4Addr

`func (o *TraceData) SetCollectionEntityIpv4Addr(v string)`

SetCollectionEntityIpv4Addr sets CollectionEntityIpv4Addr field to given value.

### HasCollectionEntityIpv4Addr

`func (o *TraceData) HasCollectionEntityIpv4Addr() bool`

HasCollectionEntityIpv4Addr returns a boolean if a field has been set.

### GetCollectionEntityIpv6Addr

`func (o *TraceData) GetCollectionEntityIpv6Addr() Ipv6Addr`

GetCollectionEntityIpv6Addr returns the CollectionEntityIpv6Addr field if non-nil, zero value otherwise.

### GetCollectionEntityIpv6AddrOk

`func (o *TraceData) GetCollectionEntityIpv6AddrOk() (*Ipv6Addr, bool)`

GetCollectionEntityIpv6AddrOk returns a tuple with the CollectionEntityIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionEntityIpv6Addr

`func (o *TraceData) SetCollectionEntityIpv6Addr(v Ipv6Addr)`

SetCollectionEntityIpv6Addr sets CollectionEntityIpv6Addr field to given value.

### HasCollectionEntityIpv6Addr

`func (o *TraceData) HasCollectionEntityIpv6Addr() bool`

HasCollectionEntityIpv6Addr returns a boolean if a field has been set.

### GetInterfaceList

`func (o *TraceData) GetInterfaceList() string`

GetInterfaceList returns the InterfaceList field if non-nil, zero value otherwise.

### GetInterfaceListOk

`func (o *TraceData) GetInterfaceListOk() (*string, bool)`

GetInterfaceListOk returns a tuple with the InterfaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceList

`func (o *TraceData) SetInterfaceList(v string)`

SetInterfaceList sets InterfaceList field to given value.

### HasInterfaceList

`func (o *TraceData) HasInterfaceList() bool`

HasInterfaceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


