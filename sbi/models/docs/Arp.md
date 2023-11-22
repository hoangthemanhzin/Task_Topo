# Arp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityLevel** | **int32** | nullable true shall not be used for this attribute | 
**PreemptCap** | [**PreemptionCapability**](PreemptionCapability.md) |  | 
**PreemptVuln** | [**PreemptionVulnerability**](PreemptionVulnerability.md) |  | 

## Methods

### NewArp

`func NewArp(priorityLevel int32, preemptCap PreemptionCapability, preemptVuln PreemptionVulnerability, ) *Arp`

NewArp instantiates a new Arp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArpWithDefaults

`func NewArpWithDefaults() *Arp`

NewArpWithDefaults instantiates a new Arp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityLevel

`func (o *Arp) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *Arp) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *Arp) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.


### SetPriorityLevelNil

`func (o *Arp) SetPriorityLevelNil(b bool)`

 SetPriorityLevelNil sets the value for PriorityLevel to be an explicit nil

### UnsetPriorityLevel
`func (o *Arp) UnsetPriorityLevel()`

UnsetPriorityLevel ensures that no value is present for PriorityLevel, not even an explicit nil
### GetPreemptCap

`func (o *Arp) GetPreemptCap() PreemptionCapability`

GetPreemptCap returns the PreemptCap field if non-nil, zero value otherwise.

### GetPreemptCapOk

`func (o *Arp) GetPreemptCapOk() (*PreemptionCapability, bool)`

GetPreemptCapOk returns a tuple with the PreemptCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptCap

`func (o *Arp) SetPreemptCap(v PreemptionCapability)`

SetPreemptCap sets PreemptCap field to given value.


### GetPreemptVuln

`func (o *Arp) GetPreemptVuln() PreemptionVulnerability`

GetPreemptVuln returns the PreemptVuln field if non-nil, zero value otherwise.

### GetPreemptVulnOk

`func (o *Arp) GetPreemptVulnOk() (*PreemptionVulnerability, bool)`

GetPreemptVulnOk returns a tuple with the PreemptVuln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptVuln

`func (o *Arp) SetPreemptVuln(v PreemptionVulnerability)`

SetPreemptVuln sets PreemptVuln field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


