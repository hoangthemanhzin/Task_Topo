# N1MessageNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N1NotifySubscriptionId** | Pointer to **string** |  | [optional] 
**N1MessageContainer** | [**N1MessageContainer**](N1MessageContainer.md) |  | 
**LcsCorrelationId** | Pointer to **string** |  | [optional] 
**RegistrationCtxtContainer** | Pointer to [**RegistrationContextContainer**](RegistrationContextContainer.md) |  | [optional] 
**NewLmfIdentification** | Pointer to **string** |  | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**CIoT5GSOptimisation** | Pointer to **bool** |  | [optional] [default to false]
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 

## Methods

### NewN1MessageNotification

`func NewN1MessageNotification(n1MessageContainer N1MessageContainer, ) *N1MessageNotification`

NewN1MessageNotification instantiates a new N1MessageNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN1MessageNotificationWithDefaults

`func NewN1MessageNotificationWithDefaults() *N1MessageNotification`

NewN1MessageNotificationWithDefaults instantiates a new N1MessageNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN1NotifySubscriptionId

`func (o *N1MessageNotification) GetN1NotifySubscriptionId() string`

GetN1NotifySubscriptionId returns the N1NotifySubscriptionId field if non-nil, zero value otherwise.

### GetN1NotifySubscriptionIdOk

`func (o *N1MessageNotification) GetN1NotifySubscriptionIdOk() (*string, bool)`

GetN1NotifySubscriptionIdOk returns a tuple with the N1NotifySubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1NotifySubscriptionId

`func (o *N1MessageNotification) SetN1NotifySubscriptionId(v string)`

SetN1NotifySubscriptionId sets N1NotifySubscriptionId field to given value.

### HasN1NotifySubscriptionId

`func (o *N1MessageNotification) HasN1NotifySubscriptionId() bool`

HasN1NotifySubscriptionId returns a boolean if a field has been set.

### GetN1MessageContainer

`func (o *N1MessageNotification) GetN1MessageContainer() N1MessageContainer`

GetN1MessageContainer returns the N1MessageContainer field if non-nil, zero value otherwise.

### GetN1MessageContainerOk

`func (o *N1MessageNotification) GetN1MessageContainerOk() (*N1MessageContainer, bool)`

GetN1MessageContainerOk returns a tuple with the N1MessageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageContainer

`func (o *N1MessageNotification) SetN1MessageContainer(v N1MessageContainer)`

SetN1MessageContainer sets N1MessageContainer field to given value.


### GetLcsCorrelationId

`func (o *N1MessageNotification) GetLcsCorrelationId() string`

GetLcsCorrelationId returns the LcsCorrelationId field if non-nil, zero value otherwise.

### GetLcsCorrelationIdOk

`func (o *N1MessageNotification) GetLcsCorrelationIdOk() (*string, bool)`

GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsCorrelationId

`func (o *N1MessageNotification) SetLcsCorrelationId(v string)`

SetLcsCorrelationId sets LcsCorrelationId field to given value.

### HasLcsCorrelationId

`func (o *N1MessageNotification) HasLcsCorrelationId() bool`

HasLcsCorrelationId returns a boolean if a field has been set.

### GetRegistrationCtxtContainer

`func (o *N1MessageNotification) GetRegistrationCtxtContainer() RegistrationContextContainer`

GetRegistrationCtxtContainer returns the RegistrationCtxtContainer field if non-nil, zero value otherwise.

### GetRegistrationCtxtContainerOk

`func (o *N1MessageNotification) GetRegistrationCtxtContainerOk() (*RegistrationContextContainer, bool)`

GetRegistrationCtxtContainerOk returns a tuple with the RegistrationCtxtContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationCtxtContainer

`func (o *N1MessageNotification) SetRegistrationCtxtContainer(v RegistrationContextContainer)`

SetRegistrationCtxtContainer sets RegistrationCtxtContainer field to given value.

### HasRegistrationCtxtContainer

`func (o *N1MessageNotification) HasRegistrationCtxtContainer() bool`

HasRegistrationCtxtContainer returns a boolean if a field has been set.

### GetNewLmfIdentification

`func (o *N1MessageNotification) GetNewLmfIdentification() string`

GetNewLmfIdentification returns the NewLmfIdentification field if non-nil, zero value otherwise.

### GetNewLmfIdentificationOk

`func (o *N1MessageNotification) GetNewLmfIdentificationOk() (*string, bool)`

GetNewLmfIdentificationOk returns a tuple with the NewLmfIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLmfIdentification

`func (o *N1MessageNotification) SetNewLmfIdentification(v string)`

SetNewLmfIdentification sets NewLmfIdentification field to given value.

### HasNewLmfIdentification

`func (o *N1MessageNotification) HasNewLmfIdentification() bool`

HasNewLmfIdentification returns a boolean if a field has been set.

### GetGuami

`func (o *N1MessageNotification) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *N1MessageNotification) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *N1MessageNotification) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *N1MessageNotification) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetCIoT5GSOptimisation

`func (o *N1MessageNotification) GetCIoT5GSOptimisation() bool`

GetCIoT5GSOptimisation returns the CIoT5GSOptimisation field if non-nil, zero value otherwise.

### GetCIoT5GSOptimisationOk

`func (o *N1MessageNotification) GetCIoT5GSOptimisationOk() (*bool, bool)`

GetCIoT5GSOptimisationOk returns a tuple with the CIoT5GSOptimisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIoT5GSOptimisation

`func (o *N1MessageNotification) SetCIoT5GSOptimisation(v bool)`

SetCIoT5GSOptimisation sets CIoT5GSOptimisation field to given value.

### HasCIoT5GSOptimisation

`func (o *N1MessageNotification) HasCIoT5GSOptimisation() bool`

HasCIoT5GSOptimisation returns a boolean if a field has been set.

### GetEcgi

`func (o *N1MessageNotification) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *N1MessageNotification) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *N1MessageNotification) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *N1MessageNotification) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *N1MessageNotification) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *N1MessageNotification) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *N1MessageNotification) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *N1MessageNotification) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


