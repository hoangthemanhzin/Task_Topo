# BdtReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspId** | **string** | Contains an identity of an application service provider. | 
**DesTimeInt** | [**TimeWindow**](TimeWindow.md) |  | 
**Dnn** | Pointer to **string** |  | [optional] 
**InterGroupId** | Pointer to **string** |  | [optional] 
**NotifUri** | Pointer to **string** |  | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**NumOfUes** | **int32** | Indicates a number of UEs. | 
**VolPerUe** | [**UsageThreshold**](UsageThreshold.md) |  | 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 
**TrafficDes** | Pointer to **string** | Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w. | [optional] 
**WarnNotifReq** | Pointer to **bool** | Indicates whether the BDT warning notification is enabled or disabled. | [optional] [default to false]

## Methods

### NewBdtReqData

`func NewBdtReqData(aspId string, desTimeInt TimeWindow, numOfUes int32, volPerUe UsageThreshold, ) *BdtReqData`

NewBdtReqData instantiates a new BdtReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtReqDataWithDefaults

`func NewBdtReqDataWithDefaults() *BdtReqData`

NewBdtReqDataWithDefaults instantiates a new BdtReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspId

`func (o *BdtReqData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *BdtReqData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *BdtReqData) SetAspId(v string)`

SetAspId sets AspId field to given value.


### GetDesTimeInt

`func (o *BdtReqData) GetDesTimeInt() TimeWindow`

GetDesTimeInt returns the DesTimeInt field if non-nil, zero value otherwise.

### GetDesTimeIntOk

`func (o *BdtReqData) GetDesTimeIntOk() (*TimeWindow, bool)`

GetDesTimeIntOk returns a tuple with the DesTimeInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesTimeInt

`func (o *BdtReqData) SetDesTimeInt(v TimeWindow)`

SetDesTimeInt sets DesTimeInt field to given value.


### GetDnn

`func (o *BdtReqData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *BdtReqData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *BdtReqData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *BdtReqData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetInterGroupId

`func (o *BdtReqData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *BdtReqData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *BdtReqData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *BdtReqData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetNotifUri

`func (o *BdtReqData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *BdtReqData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *BdtReqData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *BdtReqData) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *BdtReqData) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *BdtReqData) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *BdtReqData) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *BdtReqData) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetNumOfUes

`func (o *BdtReqData) GetNumOfUes() int32`

GetNumOfUes returns the NumOfUes field if non-nil, zero value otherwise.

### GetNumOfUesOk

`func (o *BdtReqData) GetNumOfUesOk() (*int32, bool)`

GetNumOfUesOk returns a tuple with the NumOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUes

`func (o *BdtReqData) SetNumOfUes(v int32)`

SetNumOfUes sets NumOfUes field to given value.


### GetVolPerUe

`func (o *BdtReqData) GetVolPerUe() UsageThreshold`

GetVolPerUe returns the VolPerUe field if non-nil, zero value otherwise.

### GetVolPerUeOk

`func (o *BdtReqData) GetVolPerUeOk() (*UsageThreshold, bool)`

GetVolPerUeOk returns a tuple with the VolPerUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolPerUe

`func (o *BdtReqData) SetVolPerUe(v UsageThreshold)`

SetVolPerUe sets VolPerUe field to given value.


### GetSnssai

`func (o *BdtReqData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *BdtReqData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *BdtReqData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *BdtReqData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSuppFeat

`func (o *BdtReqData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *BdtReqData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *BdtReqData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *BdtReqData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetTrafficDes

`func (o *BdtReqData) GetTrafficDes() string`

GetTrafficDes returns the TrafficDes field if non-nil, zero value otherwise.

### GetTrafficDesOk

`func (o *BdtReqData) GetTrafficDesOk() (*string, bool)`

GetTrafficDesOk returns a tuple with the TrafficDes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDes

`func (o *BdtReqData) SetTrafficDes(v string)`

SetTrafficDes sets TrafficDes field to given value.

### HasTrafficDes

`func (o *BdtReqData) HasTrafficDes() bool`

HasTrafficDes returns a boolean if a field has been set.

### GetWarnNotifReq

`func (o *BdtReqData) GetWarnNotifReq() bool`

GetWarnNotifReq returns the WarnNotifReq field if non-nil, zero value otherwise.

### GetWarnNotifReqOk

`func (o *BdtReqData) GetWarnNotifReqOk() (*bool, bool)`

GetWarnNotifReqOk returns a tuple with the WarnNotifReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnNotifReq

`func (o *BdtReqData) SetWarnNotifReq(v bool)`

SetWarnNotifReq sets WarnNotifReq field to given value.

### HasWarnNotifReq

`func (o *BdtReqData) HasWarnNotifReq() bool`

HasWarnNotifReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


