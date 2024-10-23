# DynamicConfigRuleAddDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this rule. | 
**PassPercentage** | **float32** | Of the users that meet the conditions of this rule, what percent should return true. | 
**Conditions** | [**[]DynamicConfigDtoRulesInnerConditionsInner**](DynamicConfigDtoRulesInnerConditionsInner.md) | An array of Condition objects. | 
**Environments** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | The Statsig ID of this rule. | [optional] 
**BaseID** | Pointer to **string** | The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout | [optional] 
**ReturnValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ReturnValueJson5** | Pointer to **string** |  | [optional] 

## Methods

### NewDynamicConfigRuleAddDto

`func NewDynamicConfigRuleAddDto(name string, passPercentage float32, conditions []DynamicConfigDtoRulesInnerConditionsInner, ) *DynamicConfigRuleAddDto`

NewDynamicConfigRuleAddDto instantiates a new DynamicConfigRuleAddDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicConfigRuleAddDtoWithDefaults

`func NewDynamicConfigRuleAddDtoWithDefaults() *DynamicConfigRuleAddDto`

NewDynamicConfigRuleAddDtoWithDefaults instantiates a new DynamicConfigRuleAddDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DynamicConfigRuleAddDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicConfigRuleAddDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicConfigRuleAddDto) SetName(v string)`

SetName sets Name field to given value.


### GetPassPercentage

`func (o *DynamicConfigRuleAddDto) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *DynamicConfigRuleAddDto) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *DynamicConfigRuleAddDto) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.


### GetConditions

`func (o *DynamicConfigRuleAddDto) GetConditions() []DynamicConfigDtoRulesInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DynamicConfigRuleAddDto) GetConditionsOk() (*[]DynamicConfigDtoRulesInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DynamicConfigRuleAddDto) SetConditions(v []DynamicConfigDtoRulesInnerConditionsInner)`

SetConditions sets Conditions field to given value.


### GetEnvironments

`func (o *DynamicConfigRuleAddDto) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *DynamicConfigRuleAddDto) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *DynamicConfigRuleAddDto) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *DynamicConfigRuleAddDto) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *DynamicConfigRuleAddDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicConfigRuleAddDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicConfigRuleAddDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DynamicConfigRuleAddDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBaseID

`func (o *DynamicConfigRuleAddDto) GetBaseID() string`

GetBaseID returns the BaseID field if non-nil, zero value otherwise.

### GetBaseIDOk

`func (o *DynamicConfigRuleAddDto) GetBaseIDOk() (*string, bool)`

GetBaseIDOk returns a tuple with the BaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseID

`func (o *DynamicConfigRuleAddDto) SetBaseID(v string)`

SetBaseID sets BaseID field to given value.

### HasBaseID

`func (o *DynamicConfigRuleAddDto) HasBaseID() bool`

HasBaseID returns a boolean if a field has been set.

### GetReturnValue

`func (o *DynamicConfigRuleAddDto) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *DynamicConfigRuleAddDto) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *DynamicConfigRuleAddDto) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *DynamicConfigRuleAddDto) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.

### GetReturnValueJson5

`func (o *DynamicConfigRuleAddDto) GetReturnValueJson5() string`

GetReturnValueJson5 returns the ReturnValueJson5 field if non-nil, zero value otherwise.

### GetReturnValueJson5Ok

`func (o *DynamicConfigRuleAddDto) GetReturnValueJson5Ok() (*string, bool)`

GetReturnValueJson5Ok returns a tuple with the ReturnValueJson5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValueJson5

`func (o *DynamicConfigRuleAddDto) SetReturnValueJson5(v string)`

SetReturnValueJson5 sets ReturnValueJson5 field to given value.

### HasReturnValueJson5

`func (o *DynamicConfigRuleAddDto) HasReturnValueJson5() bool`

HasReturnValueJson5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


