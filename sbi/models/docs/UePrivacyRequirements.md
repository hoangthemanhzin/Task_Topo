# UePrivacyRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LcsServiceAuthInfo** | Pointer to [**LcsServiceAuth**](LcsServiceAuth.md) |  | [optional] 
**CodeWordCheck** | Pointer to **bool** |  | [optional] 

## Methods

### NewUePrivacyRequirements

`func NewUePrivacyRequirements() *UePrivacyRequirements`

NewUePrivacyRequirements instantiates a new UePrivacyRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUePrivacyRequirementsWithDefaults

`func NewUePrivacyRequirementsWithDefaults() *UePrivacyRequirements`

NewUePrivacyRequirementsWithDefaults instantiates a new UePrivacyRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLcsServiceAuthInfo

`func (o *UePrivacyRequirements) GetLcsServiceAuthInfo() LcsServiceAuth`

GetLcsServiceAuthInfo returns the LcsServiceAuthInfo field if non-nil, zero value otherwise.

### GetLcsServiceAuthInfoOk

`func (o *UePrivacyRequirements) GetLcsServiceAuthInfoOk() (*LcsServiceAuth, bool)`

GetLcsServiceAuthInfoOk returns a tuple with the LcsServiceAuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsServiceAuthInfo

`func (o *UePrivacyRequirements) SetLcsServiceAuthInfo(v LcsServiceAuth)`

SetLcsServiceAuthInfo sets LcsServiceAuthInfo field to given value.

### HasLcsServiceAuthInfo

`func (o *UePrivacyRequirements) HasLcsServiceAuthInfo() bool`

HasLcsServiceAuthInfo returns a boolean if a field has been set.

### GetCodeWordCheck

`func (o *UePrivacyRequirements) GetCodeWordCheck() bool`

GetCodeWordCheck returns the CodeWordCheck field if non-nil, zero value otherwise.

### GetCodeWordCheckOk

`func (o *UePrivacyRequirements) GetCodeWordCheckOk() (*bool, bool)`

GetCodeWordCheckOk returns a tuple with the CodeWordCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeWordCheck

`func (o *UePrivacyRequirements) SetCodeWordCheck(v bool)`

SetCodeWordCheck sets CodeWordCheck field to given value.

### HasCodeWordCheck

`func (o *UePrivacyRequirements) HasCodeWordCheck() bool`

HasCodeWordCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


