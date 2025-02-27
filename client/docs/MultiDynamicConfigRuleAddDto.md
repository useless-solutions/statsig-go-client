# MultiDynamicConfigRuleAddDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**[]DynamicConfigDtoRulesInner**](DynamicConfigDtoRulesInner.md) | A list of new dynamic config rules | [default to []]

## Methods

### NewMultiDynamicConfigRuleAddDto

`func NewMultiDynamicConfigRuleAddDto(rules []DynamicConfigDtoRulesInner, ) *MultiDynamicConfigRuleAddDto`

NewMultiDynamicConfigRuleAddDto instantiates a new MultiDynamicConfigRuleAddDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiDynamicConfigRuleAddDtoWithDefaults

`func NewMultiDynamicConfigRuleAddDtoWithDefaults() *MultiDynamicConfigRuleAddDto`

NewMultiDynamicConfigRuleAddDtoWithDefaults instantiates a new MultiDynamicConfigRuleAddDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *MultiDynamicConfigRuleAddDto) GetRules() []DynamicConfigDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *MultiDynamicConfigRuleAddDto) GetRulesOk() (*[]DynamicConfigDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *MultiDynamicConfigRuleAddDto) SetRules(v []DynamicConfigDtoRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


