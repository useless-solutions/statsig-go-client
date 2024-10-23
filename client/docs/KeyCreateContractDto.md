# KeyCreateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Type** | **string** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Environments** | Pointer to **[]string** |  | [optional] 
**TargetAppID** | Pointer to **string** |  | [optional] 
**SecondaryTargetAppIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKeyCreateContractDto

`func NewKeyCreateContractDto(description string, type_ string, ) *KeyCreateContractDto`

NewKeyCreateContractDto instantiates a new KeyCreateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyCreateContractDtoWithDefaults

`func NewKeyCreateContractDtoWithDefaults() *KeyCreateContractDto`

NewKeyCreateContractDtoWithDefaults instantiates a new KeyCreateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KeyCreateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyCreateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyCreateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *KeyCreateContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyCreateContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyCreateContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetScopes

`func (o *KeyCreateContractDto) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *KeyCreateContractDto) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *KeyCreateContractDto) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *KeyCreateContractDto) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEnvironments

`func (o *KeyCreateContractDto) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *KeyCreateContractDto) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *KeyCreateContractDto) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *KeyCreateContractDto) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetTargetAppID

`func (o *KeyCreateContractDto) GetTargetAppID() string`

GetTargetAppID returns the TargetAppID field if non-nil, zero value otherwise.

### GetTargetAppIDOk

`func (o *KeyCreateContractDto) GetTargetAppIDOk() (*string, bool)`

GetTargetAppIDOk returns a tuple with the TargetAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppID

`func (o *KeyCreateContractDto) SetTargetAppID(v string)`

SetTargetAppID sets TargetAppID field to given value.

### HasTargetAppID

`func (o *KeyCreateContractDto) HasTargetAppID() bool`

HasTargetAppID returns a boolean if a field has been set.

### GetSecondaryTargetAppIDs

`func (o *KeyCreateContractDto) GetSecondaryTargetAppIDs() []string`

GetSecondaryTargetAppIDs returns the SecondaryTargetAppIDs field if non-nil, zero value otherwise.

### GetSecondaryTargetAppIDsOk

`func (o *KeyCreateContractDto) GetSecondaryTargetAppIDsOk() (*[]string, bool)`

GetSecondaryTargetAppIDsOk returns a tuple with the SecondaryTargetAppIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTargetAppIDs

`func (o *KeyCreateContractDto) SetSecondaryTargetAppIDs(v []string)`

SetSecondaryTargetAppIDs sets SecondaryTargetAppIDs field to given value.

### HasSecondaryTargetAppIDs

`func (o *KeyCreateContractDto) HasSecondaryTargetAppIDs() bool`

HasSecondaryTargetAppIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


