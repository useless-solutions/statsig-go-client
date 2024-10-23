# KeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**nil**](nil.md) |  | 
**Type** | **string** |  | 
**Description** | **string** |  | 
**Scopes** | **[]string** |  | 
**Environments** | Pointer to **[]string** |  | [optional] 
**PrimaryTargetApp** | Pointer to [**nil**](nil.md) |  | [optional] 
**SecondaryTargetApps** | Pointer to **[]string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewKeyDto

`func NewKeyDto(key nil, type_ string, description string, scopes []string, status string, ) *KeyDto`

NewKeyDto instantiates a new KeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDtoWithDefaults

`func NewKeyDtoWithDefaults() *KeyDto`

NewKeyDtoWithDefaults instantiates a new KeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyDto) GetKey() nil`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyDto) GetKeyOk() (*nil, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyDto) SetKey(v nil)`

SetKey sets Key field to given value.


### GetType

`func (o *KeyDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyDto) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *KeyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScopes

`func (o *KeyDto) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *KeyDto) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *KeyDto) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetEnvironments

`func (o *KeyDto) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *KeyDto) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *KeyDto) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *KeyDto) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPrimaryTargetApp

`func (o *KeyDto) GetPrimaryTargetApp() nil`

GetPrimaryTargetApp returns the PrimaryTargetApp field if non-nil, zero value otherwise.

### GetPrimaryTargetAppOk

`func (o *KeyDto) GetPrimaryTargetAppOk() (*nil, bool)`

GetPrimaryTargetAppOk returns a tuple with the PrimaryTargetApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTargetApp

`func (o *KeyDto) SetPrimaryTargetApp(v nil)`

SetPrimaryTargetApp sets PrimaryTargetApp field to given value.

### HasPrimaryTargetApp

`func (o *KeyDto) HasPrimaryTargetApp() bool`

HasPrimaryTargetApp returns a boolean if a field has been set.

### GetSecondaryTargetApps

`func (o *KeyDto) GetSecondaryTargetApps() []string`

GetSecondaryTargetApps returns the SecondaryTargetApps field if non-nil, zero value otherwise.

### GetSecondaryTargetAppsOk

`func (o *KeyDto) GetSecondaryTargetAppsOk() (*[]string, bool)`

GetSecondaryTargetAppsOk returns a tuple with the SecondaryTargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTargetApps

`func (o *KeyDto) SetSecondaryTargetApps(v []string)`

SetSecondaryTargetApps sets SecondaryTargetApps field to given value.

### HasSecondaryTargetApps

`func (o *KeyDto) HasSecondaryTargetApps() bool`

HasSecondaryTargetApps returns a boolean if a field has been set.

### GetStatus

`func (o *KeyDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyDto) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


