# RuleUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this rule. | [optional] 
**PassPercentage** | Pointer to **float32** | Of the users that meet the conditions of this rule, what percent should return true. | [optional] 
**Conditions** | Pointer to [**[]DynamicConfigDtoRulesInnerConditionsInner**](DynamicConfigDtoRulesInnerConditionsInner.md) | An array of Condition objects. | [optional] 
**Environments** | Pointer to **[]string** | The environments this rule is enabled for. | [optional] 
**Id** | Pointer to **string** | The Statsig ID of this rule. | [optional] 
**BaseID** | Pointer to **string** | The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout | [optional] 
**ReturnValue** | Pointer to **map[string]interface{}** | The return value of the rule. | [optional] 

## Methods

### NewRuleUpdateDto

`func NewRuleUpdateDto() *RuleUpdateDto`

NewRuleUpdateDto instantiates a new RuleUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleUpdateDtoWithDefaults

`func NewRuleUpdateDtoWithDefaults() *RuleUpdateDto`

NewRuleUpdateDtoWithDefaults instantiates a new RuleUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPercentage

`func (o *RuleUpdateDto) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *RuleUpdateDto) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *RuleUpdateDto) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.

### HasPassPercentage

`func (o *RuleUpdateDto) HasPassPercentage() bool`

HasPassPercentage returns a boolean if a field has been set.

### GetConditions

`func (o *RuleUpdateDto) GetConditions() []DynamicConfigDtoRulesInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleUpdateDto) GetConditionsOk() (*[]DynamicConfigDtoRulesInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleUpdateDto) SetConditions(v []DynamicConfigDtoRulesInnerConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RuleUpdateDto) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetEnvironments

`func (o *RuleUpdateDto) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *RuleUpdateDto) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *RuleUpdateDto) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *RuleUpdateDto) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *RuleUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBaseID

`func (o *RuleUpdateDto) GetBaseID() string`

GetBaseID returns the BaseID field if non-nil, zero value otherwise.

### GetBaseIDOk

`func (o *RuleUpdateDto) GetBaseIDOk() (*string, bool)`

GetBaseIDOk returns a tuple with the BaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseID

`func (o *RuleUpdateDto) SetBaseID(v string)`

SetBaseID sets BaseID field to given value.

### HasBaseID

`func (o *RuleUpdateDto) HasBaseID() bool`

HasBaseID returns a boolean if a field has been set.

### GetReturnValue

`func (o *RuleUpdateDto) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *RuleUpdateDto) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *RuleUpdateDto) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *RuleUpdateDto) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


