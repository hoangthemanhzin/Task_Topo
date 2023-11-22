# LadnInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ladn** | **string** |  | 
**Presence** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 

## Methods

### NewLadnInfo

`func NewLadnInfo(ladn string, ) *LadnInfo`

NewLadnInfo instantiates a new LadnInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLadnInfoWithDefaults

`func NewLadnInfoWithDefaults() *LadnInfo`

NewLadnInfoWithDefaults instantiates a new LadnInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLadn

`func (o *LadnInfo) GetLadn() string`

GetLadn returns the Ladn field if non-nil, zero value otherwise.

### GetLadnOk

`func (o *LadnInfo) GetLadnOk() (*string, bool)`

GetLadnOk returns a tuple with the Ladn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLadn

`func (o *LadnInfo) SetLadn(v string)`

SetLadn sets Ladn field to given value.


### GetPresence

`func (o *LadnInfo) GetPresence() PresenceState`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *LadnInfo) GetPresenceOk() (*PresenceState, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *LadnInfo) SetPresence(v PresenceState)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *LadnInfo) HasPresence() bool`

HasPresence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


