# KeyUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Environments** | Pointer to **[]string** |  | [optional] 
**TargetAppID** | Pointer to [**nil**](nil.md) |  | [optional] 
**SecondaryTargetAppIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKeyUpdateContractDto

`func NewKeyUpdateContractDto() *KeyUpdateContractDto`

NewKeyUpdateContractDto instantiates a new KeyUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUpdateContractDtoWithDefaults

`func NewKeyUpdateContractDtoWithDefaults() *KeyUpdateContractDto`

NewKeyUpdateContractDtoWithDefaults instantiates a new KeyUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KeyUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyUpdateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopes

`func (o *KeyUpdateContractDto) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *KeyUpdateContractDto) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *KeyUpdateContractDto) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *KeyUpdateContractDto) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEnvironments

`func (o *KeyUpdateContractDto) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *KeyUpdateContractDto) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *KeyUpdateContractDto) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *KeyUpdateContractDto) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetTargetAppID

`func (o *KeyUpdateContractDto) GetTargetAppID() nil`

GetTargetAppID returns the TargetAppID field if non-nil, zero value otherwise.

### GetTargetAppIDOk

`func (o *KeyUpdateContractDto) GetTargetAppIDOk() (*nil, bool)`

GetTargetAppIDOk returns a tuple with the TargetAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppID

`func (o *KeyUpdateContractDto) SetTargetAppID(v nil)`

SetTargetAppID sets TargetAppID field to given value.

### HasTargetAppID

`func (o *KeyUpdateContractDto) HasTargetAppID() bool`

HasTargetAppID returns a boolean if a field has been set.

### GetSecondaryTargetAppIDs

`func (o *KeyUpdateContractDto) GetSecondaryTargetAppIDs() []string`

GetSecondaryTargetAppIDs returns the SecondaryTargetAppIDs field if non-nil, zero value otherwise.

### GetSecondaryTargetAppIDsOk

`func (o *KeyUpdateContractDto) GetSecondaryTargetAppIDsOk() (*[]string, bool)`

GetSecondaryTargetAppIDsOk returns a tuple with the SecondaryTargetAppIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTargetAppIDs

`func (o *KeyUpdateContractDto) SetSecondaryTargetAppIDs(v []string)`

SetSecondaryTargetAppIDs sets SecondaryTargetAppIDs field to given value.

### HasSecondaryTargetAppIDs

`func (o *KeyUpdateContractDto) HasSecondaryTargetAppIDs() bool`

HasSecondaryTargetAppIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


