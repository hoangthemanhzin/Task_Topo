# GbrQosFlowInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFbrDl** | **string** |  | 
**MaxFbrUl** | **string** |  | 
**GuaFbrDl** | **string** |  | 
**GuaFbrUl** | **string** |  | 
**NotifControl** | Pointer to [**NotificationControl**](NotificationControl.md) |  | [optional] 
**MaxPacketLossRateDl** | Pointer to **int32** |  | [optional] 
**MaxPacketLossRateUl** | Pointer to **int32** |  | [optional] 
**AlternativeQosProfileList** | Pointer to [**[]AlternativeQosProfile**](AlternativeQosProfile.md) |  | [optional] 

## Methods

### NewGbrQosFlowInformation

`func NewGbrQosFlowInformation(maxFbrDl string, maxFbrUl string, guaFbrDl string, guaFbrUl string, ) *GbrQosFlowInformation`

NewGbrQosFlowInformation instantiates a new GbrQosFlowInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGbrQosFlowInformationWithDefaults

`func NewGbrQosFlowInformationWithDefaults() *GbrQosFlowInformation`

NewGbrQosFlowInformationWithDefaults instantiates a new GbrQosFlowInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFbrDl

`func (o *GbrQosFlowInformation) GetMaxFbrDl() string`

GetMaxFbrDl returns the MaxFbrDl field if non-nil, zero value otherwise.

### GetMaxFbrDlOk

`func (o *GbrQosFlowInformation) GetMaxFbrDlOk() (*string, bool)`

GetMaxFbrDlOk returns a tuple with the MaxFbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFbrDl

`func (o *GbrQosFlowInformation) SetMaxFbrDl(v string)`

SetMaxFbrDl sets MaxFbrDl field to given value.


### GetMaxFbrUl

`func (o *GbrQosFlowInformation) GetMaxFbrUl() string`

GetMaxFbrUl returns the MaxFbrUl field if non-nil, zero value otherwise.

### GetMaxFbrUlOk

`func (o *GbrQosFlowInformation) GetMaxFbrUlOk() (*string, bool)`

GetMaxFbrUlOk returns a tuple with the MaxFbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFbrUl

`func (o *GbrQosFlowInformation) SetMaxFbrUl(v string)`

SetMaxFbrUl sets MaxFbrUl field to given value.


### GetGuaFbrDl

`func (o *GbrQosFlowInformation) GetGuaFbrDl() string`

GetGuaFbrDl returns the GuaFbrDl field if non-nil, zero value otherwise.

### GetGuaFbrDlOk

`func (o *GbrQosFlowInformation) GetGuaFbrDlOk() (*string, bool)`

GetGuaFbrDlOk returns a tuple with the GuaFbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaFbrDl

`func (o *GbrQosFlowInformation) SetGuaFbrDl(v string)`

SetGuaFbrDl sets GuaFbrDl field to given value.


### GetGuaFbrUl

`func (o *GbrQosFlowInformation) GetGuaFbrUl() string`

GetGuaFbrUl returns the GuaFbrUl field if non-nil, zero value otherwise.

### GetGuaFbrUlOk

`func (o *GbrQosFlowInformation) GetGuaFbrUlOk() (*string, bool)`

GetGuaFbrUlOk returns a tuple with the GuaFbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaFbrUl

`func (o *GbrQosFlowInformation) SetGuaFbrUl(v string)`

SetGuaFbrUl sets GuaFbrUl field to given value.


### GetNotifControl

`func (o *GbrQosFlowInformation) GetNotifControl() NotificationControl`

GetNotifControl returns the NotifControl field if non-nil, zero value otherwise.

### GetNotifControlOk

`func (o *GbrQosFlowInformation) GetNotifControlOk() (*NotificationControl, bool)`

GetNotifControlOk returns a tuple with the NotifControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifControl

`func (o *GbrQosFlowInformation) SetNotifControl(v NotificationControl)`

SetNotifControl sets NotifControl field to given value.

### HasNotifControl

`func (o *GbrQosFlowInformation) HasNotifControl() bool`

HasNotifControl returns a boolean if a field has been set.

### GetMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.

### GetMaxPacketLossRateUl

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateUl() int32`

GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateUlOk

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateUlOk() (*int32, bool)`

GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateUl

`func (o *GbrQosFlowInformation) SetMaxPacketLossRateUl(v int32)`

SetMaxPacketLossRateUl sets MaxPacketLossRateUl field to given value.

### HasMaxPacketLossRateUl

`func (o *GbrQosFlowInformation) HasMaxPacketLossRateUl() bool`

HasMaxPacketLossRateUl returns a boolean if a field has been set.

### GetAlternativeQosProfileList

`func (o *GbrQosFlowInformation) GetAlternativeQosProfileList() []AlternativeQosProfile`

GetAlternativeQosProfileList returns the AlternativeQosProfileList field if non-nil, zero value otherwise.

### GetAlternativeQosProfileListOk

`func (o *GbrQosFlowInformation) GetAlternativeQosProfileListOk() (*[]AlternativeQosProfile, bool)`

GetAlternativeQosProfileListOk returns a tuple with the AlternativeQosProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeQosProfileList

`func (o *GbrQosFlowInformation) SetAlternativeQosProfileList(v []AlternativeQosProfile)`

SetAlternativeQosProfileList sets AlternativeQosProfileList field to given value.

### HasAlternativeQosProfileList

`func (o *GbrQosFlowInformation) HasAlternativeQosProfileList() bool`

HasAlternativeQosProfileList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


