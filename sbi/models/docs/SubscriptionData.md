# SubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfStatusUri** | **string** |  | 
**GuamiList** | Pointer to [**[]Guami**](Guami.md) |  | [optional] 

## Methods

### NewSubscriptionData

`func NewSubscriptionData(amfStatusUri string, ) *SubscriptionData`

NewSubscriptionData instantiates a new SubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataWithDefaults

`func NewSubscriptionDataWithDefaults() *SubscriptionData`

NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfStatusUri

`func (o *SubscriptionData) GetAmfStatusUri() string`

GetAmfStatusUri returns the AmfStatusUri field if non-nil, zero value otherwise.

### GetAmfStatusUriOk

`func (o *SubscriptionData) GetAmfStatusUriOk() (*string, bool)`

GetAmfStatusUriOk returns a tuple with the AmfStatusUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfStatusUri

`func (o *SubscriptionData) SetAmfStatusUri(v string)`

SetAmfStatusUri sets AmfStatusUri field to given value.


### GetGuamiList

`func (o *SubscriptionData) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *SubscriptionData) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *SubscriptionData) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.

### HasGuamiList

`func (o *SubscriptionData) HasGuamiList() bool`

HasGuamiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


