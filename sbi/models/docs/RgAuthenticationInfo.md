# RgAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suci** | **string** |  | 
**AuthenticatedInd** | **bool** |  | [default to false]
**SupportedFeatures** | Pointer to **string** |  | [optional] 

## Methods

### NewRgAuthenticationInfo

`func NewRgAuthenticationInfo(suci string, authenticatedInd bool, ) *RgAuthenticationInfo`

NewRgAuthenticationInfo instantiates a new RgAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRgAuthenticationInfoWithDefaults

`func NewRgAuthenticationInfoWithDefaults() *RgAuthenticationInfo`

NewRgAuthenticationInfoWithDefaults instantiates a new RgAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuci

`func (o *RgAuthenticationInfo) GetSuci() string`

GetSuci returns the Suci field if non-nil, zero value otherwise.

### GetSuciOk

`func (o *RgAuthenticationInfo) GetSuciOk() (*string, bool)`

GetSuciOk returns a tuple with the Suci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuci

`func (o *RgAuthenticationInfo) SetSuci(v string)`

SetSuci sets Suci field to given value.


### GetAuthenticatedInd

`func (o *RgAuthenticationInfo) GetAuthenticatedInd() bool`

GetAuthenticatedInd returns the AuthenticatedInd field if non-nil, zero value otherwise.

### GetAuthenticatedIndOk

`func (o *RgAuthenticationInfo) GetAuthenticatedIndOk() (*bool, bool)`

GetAuthenticatedIndOk returns a tuple with the AuthenticatedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedInd

`func (o *RgAuthenticationInfo) SetAuthenticatedInd(v bool)`

SetAuthenticatedInd sets AuthenticatedInd field to given value.


### GetSupportedFeatures

`func (o *RgAuthenticationInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RgAuthenticationInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RgAuthenticationInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RgAuthenticationInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


