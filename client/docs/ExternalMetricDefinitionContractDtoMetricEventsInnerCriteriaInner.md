# ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of event criterion for filtering metrics. Options include &#x60;value&#x60;, &#x60;metadata&#x60;, &#x60;user&#x60;, and &#x60;user_custom&#x60;. | 
**Column** | Pointer to **string** | Optional column specifying which data attribute to filter on. | [optional] 
**Condition** | **string** | sql_filter, start_withs, ends_with, and after_exposure are only applicable in Warehouse Native | 
**Values** | Pointer to **[]string** | Optional array of values for the criterion to match against. | [optional] 
**NullVacuousOverride** | Pointer to **bool** | If true, overrides null values in criterion evaluation. | [optional] 

## Methods

### NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner

`func NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner(type_ string, condition string, ) *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInnerWithDefaults

`func NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInnerWithDefaults() *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetType(v string)`

SetType sets Type field to given value.


### GetColumn

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetColumn(v string)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetCondition

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetValues

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetNullVacuousOverride

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetNullVacuousOverride() bool`

GetNullVacuousOverride returns the NullVacuousOverride field if non-nil, zero value otherwise.

### GetNullVacuousOverrideOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetNullVacuousOverrideOk() (*bool, bool)`

GetNullVacuousOverrideOk returns a tuple with the NullVacuousOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullVacuousOverride

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetNullVacuousOverride(v bool)`

SetNullVacuousOverride sets NullVacuousOverride field to given value.

### HasNullVacuousOverride

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasNullVacuousOverride() bool`

HasNullVacuousOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


