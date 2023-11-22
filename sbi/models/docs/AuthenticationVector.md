# AuthenticationVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvType** | [**AvType**](AvType.md) |  | 
**Rand** | **string** |  | 
**Xres** | **string** |  | 
**Autn** | **string** |  | 
**CkPrime** | **string** |  | 
**IkPrime** | **string** |  | 
**XresStar** | **string** |  | 
**Kausf** | **string** |  | 

## Methods

### NewAuthenticationVector

`func NewAuthenticationVector(avType AvType, rand string, xres string, autn string, ckPrime string, ikPrime string, xresStar string, kausf string, ) *AuthenticationVector`

NewAuthenticationVector instantiates a new AuthenticationVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationVectorWithDefaults

`func NewAuthenticationVectorWithDefaults() *AuthenticationVector`

NewAuthenticationVectorWithDefaults instantiates a new AuthenticationVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvType

`func (o *AuthenticationVector) GetAvType() AvType`

GetAvType returns the AvType field if non-nil, zero value otherwise.

### GetAvTypeOk

`func (o *AuthenticationVector) GetAvTypeOk() (*AvType, bool)`

GetAvTypeOk returns a tuple with the AvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvType

`func (o *AuthenticationVector) SetAvType(v AvType)`

SetAvType sets AvType field to given value.


### GetRand

`func (o *AuthenticationVector) GetRand() string`

GetRand returns the Rand field if non-nil, zero value otherwise.

### GetRandOk

`func (o *AuthenticationVector) GetRandOk() (*string, bool)`

GetRandOk returns a tuple with the Rand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRand

`func (o *AuthenticationVector) SetRand(v string)`

SetRand sets Rand field to given value.


### GetXres

`func (o *AuthenticationVector) GetXres() string`

GetXres returns the Xres field if non-nil, zero value otherwise.

### GetXresOk

`func (o *AuthenticationVector) GetXresOk() (*string, bool)`

GetXresOk returns a tuple with the Xres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXres

`func (o *AuthenticationVector) SetXres(v string)`

SetXres sets Xres field to given value.


### GetAutn

`func (o *AuthenticationVector) GetAutn() string`

GetAutn returns the Autn field if non-nil, zero value otherwise.

### GetAutnOk

`func (o *AuthenticationVector) GetAutnOk() (*string, bool)`

GetAutnOk returns a tuple with the Autn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutn

`func (o *AuthenticationVector) SetAutn(v string)`

SetAutn sets Autn field to given value.


### GetCkPrime

`func (o *AuthenticationVector) GetCkPrime() string`

GetCkPrime returns the CkPrime field if non-nil, zero value otherwise.

### GetCkPrimeOk

`func (o *AuthenticationVector) GetCkPrimeOk() (*string, bool)`

GetCkPrimeOk returns a tuple with the CkPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCkPrime

`func (o *AuthenticationVector) SetCkPrime(v string)`

SetCkPrime sets CkPrime field to given value.


### GetIkPrime

`func (o *AuthenticationVector) GetIkPrime() string`

GetIkPrime returns the IkPrime field if non-nil, zero value otherwise.

### GetIkPrimeOk

`func (o *AuthenticationVector) GetIkPrimeOk() (*string, bool)`

GetIkPrimeOk returns a tuple with the IkPrime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkPrime

`func (o *AuthenticationVector) SetIkPrime(v string)`

SetIkPrime sets IkPrime field to given value.


### GetXresStar

`func (o *AuthenticationVector) GetXresStar() string`

GetXresStar returns the XresStar field if non-nil, zero value otherwise.

### GetXresStarOk

`func (o *AuthenticationVector) GetXresStarOk() (*string, bool)`

GetXresStarOk returns a tuple with the XresStar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXresStar

`func (o *AuthenticationVector) SetXresStar(v string)`

SetXresStar sets XresStar field to given value.


### GetKausf

`func (o *AuthenticationVector) GetKausf() string`

GetKausf returns the Kausf field if non-nil, zero value otherwise.

### GetKausfOk

`func (o *AuthenticationVector) GetKausfOk() (*string, bool)`

GetKausfOk returns a tuple with the Kausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKausf

`func (o *AuthenticationVector) SetKausf(v string)`

SetKausf sets Kausf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


