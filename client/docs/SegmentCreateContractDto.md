# SegmentCreateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the segment | 
**Id** | Pointer to **string** | optional id of the segment (defaults to name) | [optional] 
**Description** | Pointer to **string** | description of the segment | [optional] 
**Type** | **string** | type of the segment | 
**IdType** | Pointer to **string** | type of id | [optional] [default to "userID"]
**Tags** | Pointer to **[]string** | optional tags for categorization | [optional] 
**CreatorID** | Pointer to [**nil**](nil.md) | the Statsig ID of the creator of this experiment | [optional] 
**CreatorEmail** | Pointer to [**nil**](nil.md) | the email of the creator of this experiment | [optional] 
**Team** | Pointer to [**nil**](nil.md) | optional identifier for the responsible team (enterprise only) | [optional] 
**Rules** | Pointer to [**[]SegmentCreateContractDtoRulesInner**](SegmentCreateContractDtoRulesInner.md) | Rule Object | [optional] 

## Methods

### NewSegmentCreateContractDto

`func NewSegmentCreateContractDto(name string, type_ string, ) *SegmentCreateContractDto`

NewSegmentCreateContractDto instantiates a new SegmentCreateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentCreateContractDtoWithDefaults

`func NewSegmentCreateContractDtoWithDefaults() *SegmentCreateContractDto`

NewSegmentCreateContractDtoWithDefaults instantiates a new SegmentCreateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentCreateContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentCreateContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentCreateContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *SegmentCreateContractDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentCreateContractDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentCreateContractDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SegmentCreateContractDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *SegmentCreateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentCreateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentCreateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentCreateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SegmentCreateContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SegmentCreateContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SegmentCreateContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetIdType

`func (o *SegmentCreateContractDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *SegmentCreateContractDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *SegmentCreateContractDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *SegmentCreateContractDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTags

`func (o *SegmentCreateContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SegmentCreateContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SegmentCreateContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SegmentCreateContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatorID

`func (o *SegmentCreateContractDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *SegmentCreateContractDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *SegmentCreateContractDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *SegmentCreateContractDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *SegmentCreateContractDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentCreateContractDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentCreateContractDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *SegmentCreateContractDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTeam

`func (o *SegmentCreateContractDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *SegmentCreateContractDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *SegmentCreateContractDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *SegmentCreateContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetRules

`func (o *SegmentCreateContractDto) GetRules() []SegmentCreateContractDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SegmentCreateContractDto) GetRulesOk() (*[]SegmentCreateContractDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SegmentCreateContractDto) SetRules(v []SegmentCreateContractDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SegmentCreateContractDto) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


