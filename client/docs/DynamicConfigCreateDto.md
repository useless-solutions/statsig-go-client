# DynamicConfigCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Is the dynamic config enabled | [optional] [default to true]
**Description** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]DynamicConfigDtoRulesInner**](DynamicConfigDtoRulesInner.md) | An array of Rule objects | [optional] 
**DefaultValue** | Pointer to **map[string]interface{}** | The fallback JSON object when no rules are triggered | [optional] 
**DefaultValueJson5** | Pointer to **string** | Can include comments. If provided with defaultValue, must parse to the same JSON | [optional] 
**IdType** | Pointer to **string** | The type of ID which the dynamic config is based on. | [optional] 
**Tags** | Pointer to **[]string** | The list of tag names attached to the dynamic config | [optional] 
**CreatorID** | Pointer to [**nil**](nil.md) |  | [optional] 
**CreatorEmail** | Pointer to [**nil**](nil.md) |  | [optional] 
**Schema** | Pointer to [**nil**](nil.md) | A schema using JSON Schema Draft 2020-12 to enforce return values of this dynamic config&#39;s rules. | [optional] 
**SchemaJson5** | Pointer to [**nil**](nil.md) | &#x60;schema&#x60; except with Json5 comments. Optional and should parse to same json as &#x60;schema&#x60;. | [optional] 
**TargetApps** | Pointer to [**DynamicConfigDtoTargetApps**](DynamicConfigDtoTargetApps.md) |  | [optional] 
**Team** | Pointer to [**nil**](nil.md) |  | [optional] 
**Name** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewDynamicConfigCreateDto

`func NewDynamicConfigCreateDto(name string, ) *DynamicConfigCreateDto`

NewDynamicConfigCreateDto instantiates a new DynamicConfigCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicConfigCreateDtoWithDefaults

`func NewDynamicConfigCreateDtoWithDefaults() *DynamicConfigCreateDto`

NewDynamicConfigCreateDtoWithDefaults instantiates a new DynamicConfigCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DynamicConfigCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DynamicConfigCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DynamicConfigCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DynamicConfigCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicConfigCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicConfigCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicConfigCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicConfigCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *DynamicConfigCreateDto) GetRules() []DynamicConfigDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DynamicConfigCreateDto) GetRulesOk() (*[]DynamicConfigDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DynamicConfigCreateDto) SetRules(v []DynamicConfigDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DynamicConfigCreateDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetDefaultValue

`func (o *DynamicConfigCreateDto) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DynamicConfigCreateDto) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DynamicConfigCreateDto) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DynamicConfigCreateDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDefaultValueJson5

`func (o *DynamicConfigCreateDto) GetDefaultValueJson5() string`

GetDefaultValueJson5 returns the DefaultValueJson5 field if non-nil, zero value otherwise.

### GetDefaultValueJson5Ok

`func (o *DynamicConfigCreateDto) GetDefaultValueJson5Ok() (*string, bool)`

GetDefaultValueJson5Ok returns a tuple with the DefaultValueJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueJson5

`func (o *DynamicConfigCreateDto) SetDefaultValueJson5(v string)`

SetDefaultValueJson5 sets DefaultValueJson5 field to given value.

### HasDefaultValueJson5

`func (o *DynamicConfigCreateDto) HasDefaultValueJson5() bool`

HasDefaultValueJson5 returns a boolean if a field has been set.

### GetIdType

`func (o *DynamicConfigCreateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *DynamicConfigCreateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *DynamicConfigCreateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *DynamicConfigCreateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTags

`func (o *DynamicConfigCreateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicConfigCreateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicConfigCreateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicConfigCreateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatorID

`func (o *DynamicConfigCreateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *DynamicConfigCreateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *DynamicConfigCreateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *DynamicConfigCreateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *DynamicConfigCreateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *DynamicConfigCreateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *DynamicConfigCreateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *DynamicConfigCreateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetSchema

`func (o *DynamicConfigCreateDto) GetSchema() nil`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DynamicConfigCreateDto) GetSchemaOk() (*nil, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DynamicConfigCreateDto) SetSchema(v nil)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DynamicConfigCreateDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaJson5

`func (o *DynamicConfigCreateDto) GetSchemaJson5() nil`

GetSchemaJson5 returns the SchemaJson5 field if non-nil, zero value otherwise.

### GetSchemaJson5Ok

`func (o *DynamicConfigCreateDto) GetSchemaJson5Ok() (*nil, bool)`

GetSchemaJson5Ok returns a tuple with the SchemaJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaJson5

`func (o *DynamicConfigCreateDto) SetSchemaJson5(v nil)`

SetSchemaJson5 sets SchemaJson5 field to given value.

### HasSchemaJson5

`func (o *DynamicConfigCreateDto) HasSchemaJson5() bool`

HasSchemaJson5 returns a boolean if a field has been set.

### GetTargetApps

`func (o *DynamicConfigCreateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *DynamicConfigCreateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *DynamicConfigCreateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *DynamicConfigCreateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTeam

`func (o *DynamicConfigCreateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DynamicConfigCreateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DynamicConfigCreateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DynamicConfigCreateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetName

`func (o *DynamicConfigCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicConfigCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicConfigCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *DynamicConfigCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicConfigCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicConfigCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DynamicConfigCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


