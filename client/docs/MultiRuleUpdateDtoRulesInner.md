# MultiRuleUpdateDtoRulesInner

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

### NewMultiRuleUpdateDtoRulesInner

`func NewMultiRuleUpdateDtoRulesInner() *MultiRuleUpdateDtoRulesInner`

NewMultiRuleUpdateDtoRulesInner instantiates a new MultiRuleUpdateDtoRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiRuleUpdateDtoRulesInnerWithDefaults

`func NewMultiRuleUpdateDtoRulesInnerWithDefaults() *MultiRuleUpdateDtoRulesInner`

NewMultiRuleUpdateDtoRulesInnerWithDefaults instantiates a new MultiRuleUpdateDtoRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MultiRuleUpdateDtoRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultiRuleUpdateDtoRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultiRuleUpdateDtoRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MultiRuleUpdateDtoRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPercentage

`func (o *MultiRuleUpdateDtoRulesInner) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *MultiRuleUpdateDtoRulesInner) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *MultiRuleUpdateDtoRulesInner) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.

### HasPassPercentage

`func (o *MultiRuleUpdateDtoRulesInner) HasPassPercentage() bool`

HasPassPercentage returns a boolean if a field has been set.

### GetConditions

`func (o *MultiRuleUpdateDtoRulesInner) GetConditions() []DynamicConfigDtoRulesInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MultiRuleUpdateDtoRulesInner) GetConditionsOk() (*[]DynamicConfigDtoRulesInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MultiRuleUpdateDtoRulesInner) SetConditions(v []DynamicConfigDtoRulesInnerConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MultiRuleUpdateDtoRulesInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetEnvironments

`func (o *MultiRuleUpdateDtoRulesInner) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *MultiRuleUpdateDtoRulesInner) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *MultiRuleUpdateDtoRulesInner) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *MultiRuleUpdateDtoRulesInner) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *MultiRuleUpdateDtoRulesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiRuleUpdateDtoRulesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiRuleUpdateDtoRulesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MultiRuleUpdateDtoRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBaseID

`func (o *MultiRuleUpdateDtoRulesInner) GetBaseID() string`

GetBaseID returns the BaseID field if non-nil, zero value otherwise.

### GetBaseIDOk

`func (o *MultiRuleUpdateDtoRulesInner) GetBaseIDOk() (*string, bool)`

GetBaseIDOk returns a tuple with the BaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseID

`func (o *MultiRuleUpdateDtoRulesInner) SetBaseID(v string)`

SetBaseID sets BaseID field to given value.

### HasBaseID

`func (o *MultiRuleUpdateDtoRulesInner) HasBaseID() bool`

HasBaseID returns a boolean if a field has been set.

### GetReturnValue

`func (o *MultiRuleUpdateDtoRulesInner) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *MultiRuleUpdateDtoRulesInner) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *MultiRuleUpdateDtoRulesInner) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *MultiRuleUpdateDtoRulesInner) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


