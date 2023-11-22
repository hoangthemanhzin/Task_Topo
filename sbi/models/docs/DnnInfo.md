# DnnInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | [**AccessAndMobilitySubscriptionDataSubscribedDnnListInner**](AccessAndMobilitySubscriptionDataSubscribedDnnListInner.md) |  | 
**DefaultDnnIndicator** | Pointer to **bool** |  | [optional] 
**LboRoamingAllowed** | Pointer to **bool** |  | [optional] 
**IwkEpsInd** | Pointer to **bool** |  | [optional] 
**DnnBarred** | Pointer to **bool** |  | [optional] 
**InvokeNefInd** | Pointer to **bool** |  | [optional] 
**SmfList** | Pointer to **[]string** |  | [optional] 
**SameSmfInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewDnnInfo

`func NewDnnInfo(dnn AccessAndMobilitySubscriptionDataSubscribedDnnListInner, ) *DnnInfo`

NewDnnInfo instantiates a new DnnInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnInfoWithDefaults

`func NewDnnInfoWithDefaults() *DnnInfo`

NewDnnInfoWithDefaults instantiates a new DnnInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnInfo) GetDnn() AccessAndMobilitySubscriptionDataSubscribedDnnListInner`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnInfo) GetDnnOk() (*AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnInfo) SetDnn(v AccessAndMobilitySubscriptionDataSubscribedDnnListInner)`

SetDnn sets Dnn field to given value.


### GetDefaultDnnIndicator

`func (o *DnnInfo) GetDefaultDnnIndicator() bool`

GetDefaultDnnIndicator returns the DefaultDnnIndicator field if non-nil, zero value otherwise.

### GetDefaultDnnIndicatorOk

`func (o *DnnInfo) GetDefaultDnnIndicatorOk() (*bool, bool)`

GetDefaultDnnIndicatorOk returns a tuple with the DefaultDnnIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDnnIndicator

`func (o *DnnInfo) SetDefaultDnnIndicator(v bool)`

SetDefaultDnnIndicator sets DefaultDnnIndicator field to given value.

### HasDefaultDnnIndicator

`func (o *DnnInfo) HasDefaultDnnIndicator() bool`

HasDefaultDnnIndicator returns a boolean if a field has been set.

### GetLboRoamingAllowed

`func (o *DnnInfo) GetLboRoamingAllowed() bool`

GetLboRoamingAllowed returns the LboRoamingAllowed field if non-nil, zero value otherwise.

### GetLboRoamingAllowedOk

`func (o *DnnInfo) GetLboRoamingAllowedOk() (*bool, bool)`

GetLboRoamingAllowedOk returns a tuple with the LboRoamingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLboRoamingAllowed

`func (o *DnnInfo) SetLboRoamingAllowed(v bool)`

SetLboRoamingAllowed sets LboRoamingAllowed field to given value.

### HasLboRoamingAllowed

`func (o *DnnInfo) HasLboRoamingAllowed() bool`

HasLboRoamingAllowed returns a boolean if a field has been set.

### GetIwkEpsInd

`func (o *DnnInfo) GetIwkEpsInd() bool`

GetIwkEpsInd returns the IwkEpsInd field if non-nil, zero value otherwise.

### GetIwkEpsIndOk

`func (o *DnnInfo) GetIwkEpsIndOk() (*bool, bool)`

GetIwkEpsIndOk returns a tuple with the IwkEpsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwkEpsInd

`func (o *DnnInfo) SetIwkEpsInd(v bool)`

SetIwkEpsInd sets IwkEpsInd field to given value.

### HasIwkEpsInd

`func (o *DnnInfo) HasIwkEpsInd() bool`

HasIwkEpsInd returns a boolean if a field has been set.

### GetDnnBarred

`func (o *DnnInfo) GetDnnBarred() bool`

GetDnnBarred returns the DnnBarred field if non-nil, zero value otherwise.

### GetDnnBarredOk

`func (o *DnnInfo) GetDnnBarredOk() (*bool, bool)`

GetDnnBarredOk returns a tuple with the DnnBarred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnBarred

`func (o *DnnInfo) SetDnnBarred(v bool)`

SetDnnBarred sets DnnBarred field to given value.

### HasDnnBarred

`func (o *DnnInfo) HasDnnBarred() bool`

HasDnnBarred returns a boolean if a field has been set.

### GetInvokeNefInd

`func (o *DnnInfo) GetInvokeNefInd() bool`

GetInvokeNefInd returns the InvokeNefInd field if non-nil, zero value otherwise.

### GetInvokeNefIndOk

`func (o *DnnInfo) GetInvokeNefIndOk() (*bool, bool)`

GetInvokeNefIndOk returns a tuple with the InvokeNefInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeNefInd

`func (o *DnnInfo) SetInvokeNefInd(v bool)`

SetInvokeNefInd sets InvokeNefInd field to given value.

### HasInvokeNefInd

`func (o *DnnInfo) HasInvokeNefInd() bool`

HasInvokeNefInd returns a boolean if a field has been set.

### GetSmfList

`func (o *DnnInfo) GetSmfList() []string`

GetSmfList returns the SmfList field if non-nil, zero value otherwise.

### GetSmfListOk

`func (o *DnnInfo) GetSmfListOk() (*[]string, bool)`

GetSmfListOk returns a tuple with the SmfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfList

`func (o *DnnInfo) SetSmfList(v []string)`

SetSmfList sets SmfList field to given value.

### HasSmfList

`func (o *DnnInfo) HasSmfList() bool`

HasSmfList returns a boolean if a field has been set.

### GetSameSmfInd

`func (o *DnnInfo) GetSameSmfInd() bool`

GetSameSmfInd returns the SameSmfInd field if non-nil, zero value otherwise.

### GetSameSmfIndOk

`func (o *DnnInfo) GetSameSmfIndOk() (*bool, bool)`

GetSameSmfIndOk returns a tuple with the SameSmfInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameSmfInd

`func (o *DnnInfo) SetSameSmfInd(v bool)`

SetSameSmfInd sets SameSmfInd field to given value.

### HasSameSmfInd

`func (o *DnnInfo) HasSameSmfInd() bool`

HasSameSmfInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


