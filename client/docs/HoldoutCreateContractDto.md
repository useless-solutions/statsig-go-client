# HoldoutCreateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the holdout | 
**Description** | Pointer to **string** | description of the holdout | [optional] 
**IdType** | Pointer to **string** | type of id | [optional] 
**TeamID** | Pointer to [**nil**](nil.md) | id of the team | [optional] 

## Methods

### NewHoldoutCreateContractDto

`func NewHoldoutCreateContractDto(name string, ) *HoldoutCreateContractDto`

NewHoldoutCreateContractDto instantiates a new HoldoutCreateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutCreateContractDtoWithDefaults

`func NewHoldoutCreateContractDtoWithDefaults() *HoldoutCreateContractDto`

NewHoldoutCreateContractDtoWithDefaults instantiates a new HoldoutCreateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HoldoutCreateContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HoldoutCreateContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HoldoutCreateContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HoldoutCreateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutCreateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutCreateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldoutCreateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdType

`func (o *HoldoutCreateContractDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *HoldoutCreateContractDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *HoldoutCreateContractDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *HoldoutCreateContractDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTeamID

`func (o *HoldoutCreateContractDto) GetTeamID() nil`

GetTeamID returns the TeamID field if non-nil, zero value otherwise.

### GetTeamIDOk

`func (o *HoldoutCreateContractDto) GetTeamIDOk() (*nil, bool)`

GetTeamIDOk returns a tuple with the TeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamID

`func (o *HoldoutCreateContractDto) SetTeamID(v nil)`

SetTeamID sets TeamID field to given value.

### HasTeamID

`func (o *HoldoutCreateContractDto) HasTeamID() bool`

HasTeamID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


