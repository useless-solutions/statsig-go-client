# DynamicConfigFullUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Is the dynamic config enabled | [default to true]
**Description** | **string** | A brief summary of what the dynamic config is being used for | 
**Rules** | [**[]DynamicConfigDtoRulesInner**](DynamicConfigDtoRulesInner.md) | An array of Rule objects | 
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

## Methods

### NewDynamicConfigFullUpdateDto

`func NewDynamicConfigFullUpdateDto(isEnabled bool, description string, rules []DynamicConfigDtoRulesInner, ) *DynamicConfigFullUpdateDto`

NewDynamicConfigFullUpdateDto instantiates a new DynamicConfigFullUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicConfigFullUpdateDtoWithDefaults

`func NewDynamicConfigFullUpdateDtoWithDefaults() *DynamicConfigFullUpdateDto`

NewDynamicConfigFullUpdateDtoWithDefaults instantiates a new DynamicConfigFullUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DynamicConfigFullUpdateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DynamicConfigFullUpdateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DynamicConfigFullUpdateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDescription

`func (o *DynamicConfigFullUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicConfigFullUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicConfigFullUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *DynamicConfigFullUpdateDto) GetRules() []DynamicConfigDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DynamicConfigFullUpdateDto) GetRulesOk() (*[]DynamicConfigDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DynamicConfigFullUpdateDto) SetRules(v []DynamicConfigDtoRulesInner)`

SetRules sets Rules field to given value.


### GetDefaultValue

`func (o *DynamicConfigFullUpdateDto) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DynamicConfigFullUpdateDto) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DynamicConfigFullUpdateDto) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DynamicConfigFullUpdateDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDefaultValueJson5

`func (o *DynamicConfigFullUpdateDto) GetDefaultValueJson5() string`

GetDefaultValueJson5 returns the DefaultValueJson5 field if non-nil, zero value otherwise.

### GetDefaultValueJson5Ok

`func (o *DynamicConfigFullUpdateDto) GetDefaultValueJson5Ok() (*string, bool)`

GetDefaultValueJson5Ok returns a tuple with the DefaultValueJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueJson5

`func (o *DynamicConfigFullUpdateDto) SetDefaultValueJson5(v string)`

SetDefaultValueJson5 sets DefaultValueJson5 field to given value.

### HasDefaultValueJson5

`func (o *DynamicConfigFullUpdateDto) HasDefaultValueJson5() bool`

HasDefaultValueJson5 returns a boolean if a field has been set.

### GetIdType

`func (o *DynamicConfigFullUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *DynamicConfigFullUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *DynamicConfigFullUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *DynamicConfigFullUpdateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTags

`func (o *DynamicConfigFullUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicConfigFullUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicConfigFullUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicConfigFullUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatorID

`func (o *DynamicConfigFullUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *DynamicConfigFullUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *DynamicConfigFullUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *DynamicConfigFullUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *DynamicConfigFullUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *DynamicConfigFullUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *DynamicConfigFullUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *DynamicConfigFullUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetSchema

`func (o *DynamicConfigFullUpdateDto) GetSchema() nil`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DynamicConfigFullUpdateDto) GetSchemaOk() (*nil, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DynamicConfigFullUpdateDto) SetSchema(v nil)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DynamicConfigFullUpdateDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaJson5

`func (o *DynamicConfigFullUpdateDto) GetSchemaJson5() nil`

GetSchemaJson5 returns the SchemaJson5 field if non-nil, zero value otherwise.

### GetSchemaJson5Ok

`func (o *DynamicConfigFullUpdateDto) GetSchemaJson5Ok() (*nil, bool)`

GetSchemaJson5Ok returns a tuple with the SchemaJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaJson5

`func (o *DynamicConfigFullUpdateDto) SetSchemaJson5(v nil)`

SetSchemaJson5 sets SchemaJson5 field to given value.

### HasSchemaJson5

`func (o *DynamicConfigFullUpdateDto) HasSchemaJson5() bool`

HasSchemaJson5 returns a boolean if a field has been set.

### GetTargetApps

`func (o *DynamicConfigFullUpdateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *DynamicConfigFullUpdateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *DynamicConfigFullUpdateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *DynamicConfigFullUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTeam

`func (o *DynamicConfigFullUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DynamicConfigFullUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DynamicConfigFullUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DynamicConfigFullUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


