# SegmentCreateContractDtoRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this rule. | 
**PassPercentage** | **float32** | Of the users that meet the conditions of this rule, what percent should return true. | 
**Conditions** | [**[]SegmentCreateContractDtoRulesInnerConditionsInner**](SegmentCreateContractDtoRulesInnerConditionsInner.md) |  | 
**Environments** | Pointer to **[]string** | The environments this rule is enabled for. | [optional] 
**Id** | Pointer to **string** | The Statsig ID of this rule. | [optional] 
**BaseID** | Pointer to **string** | The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout | [optional] 
**ReturnValue** | Pointer to **map[string]interface{}** | The return value of the rule. | [optional] 

## Methods

### NewSegmentCreateContractDtoRulesInner

`func NewSegmentCreateContractDtoRulesInner(name string, passPercentage float32, conditions []SegmentCreateContractDtoRulesInnerConditionsInner, ) *SegmentCreateContractDtoRulesInner`

NewSegmentCreateContractDtoRulesInner instantiates a new SegmentCreateContractDtoRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentCreateContractDtoRulesInnerWithDefaults

`func NewSegmentCreateContractDtoRulesInnerWithDefaults() *SegmentCreateContractDtoRulesInner`

NewSegmentCreateContractDtoRulesInnerWithDefaults instantiates a new SegmentCreateContractDtoRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentCreateContractDtoRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentCreateContractDtoRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentCreateContractDtoRulesInner) SetName(v string)`

SetName sets Name field to given value.


### GetPassPercentage

`func (o *SegmentCreateContractDtoRulesInner) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *SegmentCreateContractDtoRulesInner) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *SegmentCreateContractDtoRulesInner) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.


### GetConditions

`func (o *SegmentCreateContractDtoRulesInner) GetConditions() []SegmentCreateContractDtoRulesInnerConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SegmentCreateContractDtoRulesInner) GetConditionsOk() (*[]SegmentCreateContractDtoRulesInnerConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SegmentCreateContractDtoRulesInner) SetConditions(v []SegmentCreateContractDtoRulesInnerConditionsInner)`

SetConditions sets Conditions field to given value.


### GetEnvironments

`func (o *SegmentCreateContractDtoRulesInner) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *SegmentCreateContractDtoRulesInner) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *SegmentCreateContractDtoRulesInner) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *SegmentCreateContractDtoRulesInner) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *SegmentCreateContractDtoRulesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentCreateContractDtoRulesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentCreateContractDtoRulesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SegmentCreateContractDtoRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBaseID

`func (o *SegmentCreateContractDtoRulesInner) GetBaseID() string`

GetBaseID returns the BaseID field if non-nil, zero value otherwise.

### GetBaseIDOk

`func (o *SegmentCreateContractDtoRulesInner) GetBaseIDOk() (*string, bool)`

GetBaseIDOk returns a tuple with the BaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseID

`func (o *SegmentCreateContractDtoRulesInner) SetBaseID(v string)`

SetBaseID sets BaseID field to given value.

### HasBaseID

`func (o *SegmentCreateContractDtoRulesInner) HasBaseID() bool`

HasBaseID returns a boolean if a field has been set.

### GetReturnValue

`func (o *SegmentCreateContractDtoRulesInner) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *SegmentCreateContractDtoRulesInner) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *SegmentCreateContractDtoRulesInner) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *SegmentCreateContractDtoRulesInner) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


