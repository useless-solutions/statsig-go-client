# DynamicConfigPartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Is the dynamic config enabled | [optional] [default to true]
**Description** | Pointer to **string** | A brief summary of what the dynamic config is being used for | [optional] 
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

## Methods

### NewDynamicConfigPartialUpdateDto

`func NewDynamicConfigPartialUpdateDto() *DynamicConfigPartialUpdateDto`

NewDynamicConfigPartialUpdateDto instantiates a new DynamicConfigPartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicConfigPartialUpdateDtoWithDefaults

`func NewDynamicConfigPartialUpdateDtoWithDefaults() *DynamicConfigPartialUpdateDto`

NewDynamicConfigPartialUpdateDtoWithDefaults instantiates a new DynamicConfigPartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DynamicConfigPartialUpdateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DynamicConfigPartialUpdateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DynamicConfigPartialUpdateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DynamicConfigPartialUpdateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicConfigPartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicConfigPartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicConfigPartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicConfigPartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *DynamicConfigPartialUpdateDto) GetRules() []DynamicConfigDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DynamicConfigPartialUpdateDto) GetRulesOk() (*[]DynamicConfigDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DynamicConfigPartialUpdateDto) SetRules(v []DynamicConfigDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DynamicConfigPartialUpdateDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetDefaultValue

`func (o *DynamicConfigPartialUpdateDto) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DynamicConfigPartialUpdateDto) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DynamicConfigPartialUpdateDto) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DynamicConfigPartialUpdateDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDefaultValueJson5

`func (o *DynamicConfigPartialUpdateDto) GetDefaultValueJson5() string`

GetDefaultValueJson5 returns the DefaultValueJson5 field if non-nil, zero value otherwise.

### GetDefaultValueJson5Ok

`func (o *DynamicConfigPartialUpdateDto) GetDefaultValueJson5Ok() (*string, bool)`

GetDefaultValueJson5Ok returns a tuple with the DefaultValueJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueJson5

`func (o *DynamicConfigPartialUpdateDto) SetDefaultValueJson5(v string)`

SetDefaultValueJson5 sets DefaultValueJson5 field to given value.

### HasDefaultValueJson5

`func (o *DynamicConfigPartialUpdateDto) HasDefaultValueJson5() bool`

HasDefaultValueJson5 returns a boolean if a field has been set.

### GetIdType

`func (o *DynamicConfigPartialUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *DynamicConfigPartialUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *DynamicConfigPartialUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *DynamicConfigPartialUpdateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTags

`func (o *DynamicConfigPartialUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicConfigPartialUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicConfigPartialUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicConfigPartialUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatorID

`func (o *DynamicConfigPartialUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *DynamicConfigPartialUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *DynamicConfigPartialUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *DynamicConfigPartialUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *DynamicConfigPartialUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *DynamicConfigPartialUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *DynamicConfigPartialUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *DynamicConfigPartialUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetSchema

`func (o *DynamicConfigPartialUpdateDto) GetSchema() nil`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DynamicConfigPartialUpdateDto) GetSchemaOk() (*nil, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DynamicConfigPartialUpdateDto) SetSchema(v nil)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DynamicConfigPartialUpdateDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaJson5

`func (o *DynamicConfigPartialUpdateDto) GetSchemaJson5() nil`

GetSchemaJson5 returns the SchemaJson5 field if non-nil, zero value otherwise.

### GetSchemaJson5Ok

`func (o *DynamicConfigPartialUpdateDto) GetSchemaJson5Ok() (*nil, bool)`

GetSchemaJson5Ok returns a tuple with the SchemaJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaJson5

`func (o *DynamicConfigPartialUpdateDto) SetSchemaJson5(v nil)`

SetSchemaJson5 sets SchemaJson5 field to given value.

### HasSchemaJson5

`func (o *DynamicConfigPartialUpdateDto) HasSchemaJson5() bool`

HasSchemaJson5 returns a boolean if a field has been set.

### GetTargetApps

`func (o *DynamicConfigPartialUpdateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *DynamicConfigPartialUpdateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *DynamicConfigPartialUpdateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *DynamicConfigPartialUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTeam

`func (o *DynamicConfigPartialUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DynamicConfigPartialUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DynamicConfigPartialUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DynamicConfigPartialUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


