# MetricsUpdateContractDtoWarehouseNative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | Allowed: count┃sum┃mean┃daily_participation┃ratio┃funnel┃count_distinct┃percentile | [optional] 
**MetricSourceName** | Pointer to **string** | For Count, Sum, Mean, User Count aggregation types: the name of metric source | [optional] 
**Criteria** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner**](ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner.md) | MetricEventCrtieria | [optional] 
**WaitForCohortWindow** | Pointer to **bool** |  | [optional] 
**DenominatorCriteria** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner**](ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner.md) | MetricEventCriteria | [optional] 
**DenominatorAggregation** | Pointer to **string** | Allowed: count┃sum┃mean┃daily_participation┃ratio┃funnel┃count_distinct┃percentile | [optional] 
**DenominatorCustomRollupEnd** | Pointer to **float32** | Custom end date for rollup in days since exposure. | [optional] 
**DenominatorCustomRollupStart** | Pointer to **float32** | Custom start date for rollup in days since exposure. | [optional] 
**DenominatorMetricSourceName** | Pointer to **string** | Name of the metric source for the denominator. | [optional] 
**DenominatorRollupTimeWindow** | Pointer to **string** | Time window for the denominator metric. Specify \&quot;custom\&quot; for a custom window. | [optional] 
**DenominatorValueColumn** | Pointer to **string** | Column name for the denominator’s value. | [optional] 
**FunnelCalculationWindow** | Pointer to **float32** | Duration for counting funnel events in days. | [optional] 
**FunnelCountDistinct** | Pointer to **string** | Allowed: users┃sessions for distinct count method in funnel events. | [optional] 
**FunnelEvents** | Pointer to [**[]ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner**](ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner.md) | List of funnel events with associated criteria and identifiers. | [optional] 
**FunnelStartCriteria** | Pointer to **string** | Allowed: start_event┃exposure to determine funnel start criteria. | [optional] 
**MetricDimensionColumns** | Pointer to **[]string** | Specify metadata columns for breaking down metric analysis. | [optional] 
**MetricBakeDays** | Pointer to **float32** | Number of days for metric baking; specify duration for analysis. | [optional] 
**NumeratorAggregation** | Pointer to **string** | Aggregation type for numerator; Allowed: count┃sum┃mean┃daily_participation┃ratio┃funnel┃count_distinct┃percentile. | [optional] 
**ValueColumn** | Pointer to **string** | Column name representing the metric’s value. | [optional] 
**WinsorizationHigh** | Pointer to **float32** | High threshold for winsorization; must be between 0 and 1. | [optional] 
**WinsorizationLow** | Pointer to **float32** | Low threshold for winsorization; must be between 0 and 1. | [optional] 
**CupedAttributionWindow** | Pointer to [**nil**](nil.md) | Attribution window for CUPED adjustments in days. | [optional] 
**OnlyIncludeUsersWithConversionEvent** | Pointer to **bool** | Flag to include only users with a conversion event in the metric. | [optional] 
**Percentile** | Pointer to **float32** | Percentile value for statistical calculations. | [optional] 
**ValueThreshold** | Pointer to **float32** | Threshold value for filtering metrics. | [optional] 
**Cap** | Pointer to **float32** | Maximum cap for metric values. | [optional] 
**RollupTimeWindow** | Pointer to **string** | General time window for rollup; can specify custom settings. | [optional] 
**CustomRollUpStart** | Pointer to **float32** | Custom start date for rollup in days since exposure. | [optional] 
**CustomRollUpEnd** | Pointer to **float32** | Custom end date for rollup in days since exposure. | [optional] 
**AllowNullRatioDenominator** | Pointer to **bool** | Include units which do not have a denominator. Only applicable to ratios. | [optional] 

## Methods

### NewMetricsUpdateContractDtoWarehouseNative

`func NewMetricsUpdateContractDtoWarehouseNative() *MetricsUpdateContractDtoWarehouseNative`

NewMetricsUpdateContractDtoWarehouseNative instantiates a new MetricsUpdateContractDtoWarehouseNative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsUpdateContractDtoWarehouseNativeWithDefaults

`func NewMetricsUpdateContractDtoWarehouseNativeWithDefaults() *MetricsUpdateContractDtoWarehouseNative`

NewMetricsUpdateContractDtoWarehouseNativeWithDefaults instantiates a new MetricsUpdateContractDtoWarehouseNative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricSourceName() string`

GetMetricSourceName returns the MetricSourceName field if non-nil, zero value otherwise.

### GetMetricSourceNameOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricSourceNameOk() (*string, bool)`

GetMetricSourceNameOk returns a tuple with the MetricSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) SetMetricSourceName(v string)`

SetMetricSourceName sets MetricSourceName field to given value.

### HasMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) HasMetricSourceName() bool`

HasMetricSourceName returns a boolean if a field has been set.

### GetCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCriteriaOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) SetCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetWaitForCohortWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWaitForCohortWindow() bool`

GetWaitForCohortWindow returns the WaitForCohortWindow field if non-nil, zero value otherwise.

### GetWaitForCohortWindowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWaitForCohortWindowOk() (*bool, bool)`

GetWaitForCohortWindowOk returns a tuple with the WaitForCohortWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCohortWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetWaitForCohortWindow(v bool)`

SetWaitForCohortWindow sets WaitForCohortWindow field to given value.

### HasWaitForCohortWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasWaitForCohortWindow() bool`

HasWaitForCohortWindow returns a boolean if a field has been set.

### GetDenominatorCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

GetDenominatorCriteria returns the DenominatorCriteria field if non-nil, zero value otherwise.

### GetDenominatorCriteriaOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCriteriaOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool)`

GetDenominatorCriteriaOk returns a tuple with the DenominatorCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)`

SetDenominatorCriteria sets DenominatorCriteria field to given value.

### HasDenominatorCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorCriteria() bool`

HasDenominatorCriteria returns a boolean if a field has been set.

### GetDenominatorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorAggregation() string`

GetDenominatorAggregation returns the DenominatorAggregation field if non-nil, zero value otherwise.

### GetDenominatorAggregationOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorAggregationOk() (*string, bool)`

GetDenominatorAggregationOk returns a tuple with the DenominatorAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorAggregation(v string)`

SetDenominatorAggregation sets DenominatorAggregation field to given value.

### HasDenominatorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorAggregation() bool`

HasDenominatorAggregation returns a boolean if a field has been set.

### GetDenominatorCustomRollupEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCustomRollupEnd() float32`

GetDenominatorCustomRollupEnd returns the DenominatorCustomRollupEnd field if non-nil, zero value otherwise.

### GetDenominatorCustomRollupEndOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCustomRollupEndOk() (*float32, bool)`

GetDenominatorCustomRollupEndOk returns a tuple with the DenominatorCustomRollupEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorCustomRollupEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorCustomRollupEnd(v float32)`

SetDenominatorCustomRollupEnd sets DenominatorCustomRollupEnd field to given value.

### HasDenominatorCustomRollupEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorCustomRollupEnd() bool`

HasDenominatorCustomRollupEnd returns a boolean if a field has been set.

### GetDenominatorCustomRollupStart

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCustomRollupStart() float32`

GetDenominatorCustomRollupStart returns the DenominatorCustomRollupStart field if non-nil, zero value otherwise.

### GetDenominatorCustomRollupStartOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorCustomRollupStartOk() (*float32, bool)`

GetDenominatorCustomRollupStartOk returns a tuple with the DenominatorCustomRollupStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorCustomRollupStart

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorCustomRollupStart(v float32)`

SetDenominatorCustomRollupStart sets DenominatorCustomRollupStart field to given value.

### HasDenominatorCustomRollupStart

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorCustomRollupStart() bool`

HasDenominatorCustomRollupStart returns a boolean if a field has been set.

### GetDenominatorMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorMetricSourceName() string`

GetDenominatorMetricSourceName returns the DenominatorMetricSourceName field if non-nil, zero value otherwise.

### GetDenominatorMetricSourceNameOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorMetricSourceNameOk() (*string, bool)`

GetDenominatorMetricSourceNameOk returns a tuple with the DenominatorMetricSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorMetricSourceName(v string)`

SetDenominatorMetricSourceName sets DenominatorMetricSourceName field to given value.

### HasDenominatorMetricSourceName

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorMetricSourceName() bool`

HasDenominatorMetricSourceName returns a boolean if a field has been set.

### GetDenominatorRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorRollupTimeWindow() string`

GetDenominatorRollupTimeWindow returns the DenominatorRollupTimeWindow field if non-nil, zero value otherwise.

### GetDenominatorRollupTimeWindowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorRollupTimeWindowOk() (*string, bool)`

GetDenominatorRollupTimeWindowOk returns a tuple with the DenominatorRollupTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorRollupTimeWindow(v string)`

SetDenominatorRollupTimeWindow sets DenominatorRollupTimeWindow field to given value.

### HasDenominatorRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorRollupTimeWindow() bool`

HasDenominatorRollupTimeWindow returns a boolean if a field has been set.

### GetDenominatorValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorValueColumn() string`

GetDenominatorValueColumn returns the DenominatorValueColumn field if non-nil, zero value otherwise.

### GetDenominatorValueColumnOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetDenominatorValueColumnOk() (*string, bool)`

GetDenominatorValueColumnOk returns a tuple with the DenominatorValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) SetDenominatorValueColumn(v string)`

SetDenominatorValueColumn sets DenominatorValueColumn field to given value.

### HasDenominatorValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) HasDenominatorValueColumn() bool`

HasDenominatorValueColumn returns a boolean if a field has been set.

### GetFunnelCalculationWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelCalculationWindow() float32`

GetFunnelCalculationWindow returns the FunnelCalculationWindow field if non-nil, zero value otherwise.

### GetFunnelCalculationWindowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelCalculationWindowOk() (*float32, bool)`

GetFunnelCalculationWindowOk returns a tuple with the FunnelCalculationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCalculationWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetFunnelCalculationWindow(v float32)`

SetFunnelCalculationWindow sets FunnelCalculationWindow field to given value.

### HasFunnelCalculationWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasFunnelCalculationWindow() bool`

HasFunnelCalculationWindow returns a boolean if a field has been set.

### GetFunnelCountDistinct

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelCountDistinct() string`

GetFunnelCountDistinct returns the FunnelCountDistinct field if non-nil, zero value otherwise.

### GetFunnelCountDistinctOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelCountDistinctOk() (*string, bool)`

GetFunnelCountDistinctOk returns a tuple with the FunnelCountDistinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCountDistinct

`func (o *MetricsUpdateContractDtoWarehouseNative) SetFunnelCountDistinct(v string)`

SetFunnelCountDistinct sets FunnelCountDistinct field to given value.

### HasFunnelCountDistinct

`func (o *MetricsUpdateContractDtoWarehouseNative) HasFunnelCountDistinct() bool`

HasFunnelCountDistinct returns a boolean if a field has been set.

### GetFunnelEvents

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelEvents() []ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner`

GetFunnelEvents returns the FunnelEvents field if non-nil, zero value otherwise.

### GetFunnelEventsOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelEventsOk() (*[]ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner, bool)`

GetFunnelEventsOk returns a tuple with the FunnelEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelEvents

`func (o *MetricsUpdateContractDtoWarehouseNative) SetFunnelEvents(v []ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner)`

SetFunnelEvents sets FunnelEvents field to given value.

### HasFunnelEvents

`func (o *MetricsUpdateContractDtoWarehouseNative) HasFunnelEvents() bool`

HasFunnelEvents returns a boolean if a field has been set.

### GetFunnelStartCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelStartCriteria() string`

GetFunnelStartCriteria returns the FunnelStartCriteria field if non-nil, zero value otherwise.

### GetFunnelStartCriteriaOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetFunnelStartCriteriaOk() (*string, bool)`

GetFunnelStartCriteriaOk returns a tuple with the FunnelStartCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStartCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) SetFunnelStartCriteria(v string)`

SetFunnelStartCriteria sets FunnelStartCriteria field to given value.

### HasFunnelStartCriteria

`func (o *MetricsUpdateContractDtoWarehouseNative) HasFunnelStartCriteria() bool`

HasFunnelStartCriteria returns a boolean if a field has been set.

### GetMetricDimensionColumns

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricDimensionColumns() []string`

GetMetricDimensionColumns returns the MetricDimensionColumns field if non-nil, zero value otherwise.

### GetMetricDimensionColumnsOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricDimensionColumnsOk() (*[]string, bool)`

GetMetricDimensionColumnsOk returns a tuple with the MetricDimensionColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDimensionColumns

`func (o *MetricsUpdateContractDtoWarehouseNative) SetMetricDimensionColumns(v []string)`

SetMetricDimensionColumns sets MetricDimensionColumns field to given value.

### HasMetricDimensionColumns

`func (o *MetricsUpdateContractDtoWarehouseNative) HasMetricDimensionColumns() bool`

HasMetricDimensionColumns returns a boolean if a field has been set.

### GetMetricBakeDays

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricBakeDays() float32`

GetMetricBakeDays returns the MetricBakeDays field if non-nil, zero value otherwise.

### GetMetricBakeDaysOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetMetricBakeDaysOk() (*float32, bool)`

GetMetricBakeDaysOk returns a tuple with the MetricBakeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricBakeDays

`func (o *MetricsUpdateContractDtoWarehouseNative) SetMetricBakeDays(v float32)`

SetMetricBakeDays sets MetricBakeDays field to given value.

### HasMetricBakeDays

`func (o *MetricsUpdateContractDtoWarehouseNative) HasMetricBakeDays() bool`

HasMetricBakeDays returns a boolean if a field has been set.

### GetNumeratorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) GetNumeratorAggregation() string`

GetNumeratorAggregation returns the NumeratorAggregation field if non-nil, zero value otherwise.

### GetNumeratorAggregationOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetNumeratorAggregationOk() (*string, bool)`

GetNumeratorAggregationOk returns a tuple with the NumeratorAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) SetNumeratorAggregation(v string)`

SetNumeratorAggregation sets NumeratorAggregation field to given value.

### HasNumeratorAggregation

`func (o *MetricsUpdateContractDtoWarehouseNative) HasNumeratorAggregation() bool`

HasNumeratorAggregation returns a boolean if a field has been set.

### GetValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) GetValueColumn() string`

GetValueColumn returns the ValueColumn field if non-nil, zero value otherwise.

### GetValueColumnOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetValueColumnOk() (*string, bool)`

GetValueColumnOk returns a tuple with the ValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) SetValueColumn(v string)`

SetValueColumn sets ValueColumn field to given value.

### HasValueColumn

`func (o *MetricsUpdateContractDtoWarehouseNative) HasValueColumn() bool`

HasValueColumn returns a boolean if a field has been set.

### GetWinsorizationHigh

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWinsorizationHigh() float32`

GetWinsorizationHigh returns the WinsorizationHigh field if non-nil, zero value otherwise.

### GetWinsorizationHighOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWinsorizationHighOk() (*float32, bool)`

GetWinsorizationHighOk returns a tuple with the WinsorizationHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorizationHigh

`func (o *MetricsUpdateContractDtoWarehouseNative) SetWinsorizationHigh(v float32)`

SetWinsorizationHigh sets WinsorizationHigh field to given value.

### HasWinsorizationHigh

`func (o *MetricsUpdateContractDtoWarehouseNative) HasWinsorizationHigh() bool`

HasWinsorizationHigh returns a boolean if a field has been set.

### GetWinsorizationLow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWinsorizationLow() float32`

GetWinsorizationLow returns the WinsorizationLow field if non-nil, zero value otherwise.

### GetWinsorizationLowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetWinsorizationLowOk() (*float32, bool)`

GetWinsorizationLowOk returns a tuple with the WinsorizationLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorizationLow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetWinsorizationLow(v float32)`

SetWinsorizationLow sets WinsorizationLow field to given value.

### HasWinsorizationLow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasWinsorizationLow() bool`

HasWinsorizationLow returns a boolean if a field has been set.

### GetCupedAttributionWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCupedAttributionWindow() nil`

GetCupedAttributionWindow returns the CupedAttributionWindow field if non-nil, zero value otherwise.

### GetCupedAttributionWindowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCupedAttributionWindowOk() (*nil, bool)`

GetCupedAttributionWindowOk returns a tuple with the CupedAttributionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCupedAttributionWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetCupedAttributionWindow(v nil)`

SetCupedAttributionWindow sets CupedAttributionWindow field to given value.

### HasCupedAttributionWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasCupedAttributionWindow() bool`

HasCupedAttributionWindow returns a boolean if a field has been set.

### GetOnlyIncludeUsersWithConversionEvent

`func (o *MetricsUpdateContractDtoWarehouseNative) GetOnlyIncludeUsersWithConversionEvent() bool`

GetOnlyIncludeUsersWithConversionEvent returns the OnlyIncludeUsersWithConversionEvent field if non-nil, zero value otherwise.

### GetOnlyIncludeUsersWithConversionEventOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetOnlyIncludeUsersWithConversionEventOk() (*bool, bool)`

GetOnlyIncludeUsersWithConversionEventOk returns a tuple with the OnlyIncludeUsersWithConversionEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIncludeUsersWithConversionEvent

`func (o *MetricsUpdateContractDtoWarehouseNative) SetOnlyIncludeUsersWithConversionEvent(v bool)`

SetOnlyIncludeUsersWithConversionEvent sets OnlyIncludeUsersWithConversionEvent field to given value.

### HasOnlyIncludeUsersWithConversionEvent

`func (o *MetricsUpdateContractDtoWarehouseNative) HasOnlyIncludeUsersWithConversionEvent() bool`

HasOnlyIncludeUsersWithConversionEvent returns a boolean if a field has been set.

### GetPercentile

`func (o *MetricsUpdateContractDtoWarehouseNative) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *MetricsUpdateContractDtoWarehouseNative) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *MetricsUpdateContractDtoWarehouseNative) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetValueThreshold

`func (o *MetricsUpdateContractDtoWarehouseNative) GetValueThreshold() float32`

GetValueThreshold returns the ValueThreshold field if non-nil, zero value otherwise.

### GetValueThresholdOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetValueThresholdOk() (*float32, bool)`

GetValueThresholdOk returns a tuple with the ValueThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueThreshold

`func (o *MetricsUpdateContractDtoWarehouseNative) SetValueThreshold(v float32)`

SetValueThreshold sets ValueThreshold field to given value.

### HasValueThreshold

`func (o *MetricsUpdateContractDtoWarehouseNative) HasValueThreshold() bool`

HasValueThreshold returns a boolean if a field has been set.

### GetCap

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCap() float32`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCapOk() (*float32, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *MetricsUpdateContractDtoWarehouseNative) SetCap(v float32)`

SetCap sets Cap field to given value.

### HasCap

`func (o *MetricsUpdateContractDtoWarehouseNative) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) GetRollupTimeWindow() string`

GetRollupTimeWindow returns the RollupTimeWindow field if non-nil, zero value otherwise.

### GetRollupTimeWindowOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetRollupTimeWindowOk() (*string, bool)`

GetRollupTimeWindowOk returns a tuple with the RollupTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) SetRollupTimeWindow(v string)`

SetRollupTimeWindow sets RollupTimeWindow field to given value.

### HasRollupTimeWindow

`func (o *MetricsUpdateContractDtoWarehouseNative) HasRollupTimeWindow() bool`

HasRollupTimeWindow returns a boolean if a field has been set.

### GetCustomRollUpStart

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCustomRollUpStart() float32`

GetCustomRollUpStart returns the CustomRollUpStart field if non-nil, zero value otherwise.

### GetCustomRollUpStartOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCustomRollUpStartOk() (*float32, bool)`

GetCustomRollUpStartOk returns a tuple with the CustomRollUpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpStart

`func (o *MetricsUpdateContractDtoWarehouseNative) SetCustomRollUpStart(v float32)`

SetCustomRollUpStart sets CustomRollUpStart field to given value.

### HasCustomRollUpStart

`func (o *MetricsUpdateContractDtoWarehouseNative) HasCustomRollUpStart() bool`

HasCustomRollUpStart returns a boolean if a field has been set.

### GetCustomRollUpEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCustomRollUpEnd() float32`

GetCustomRollUpEnd returns the CustomRollUpEnd field if non-nil, zero value otherwise.

### GetCustomRollUpEndOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetCustomRollUpEndOk() (*float32, bool)`

GetCustomRollUpEndOk returns a tuple with the CustomRollUpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) SetCustomRollUpEnd(v float32)`

SetCustomRollUpEnd sets CustomRollUpEnd field to given value.

### HasCustomRollUpEnd

`func (o *MetricsUpdateContractDtoWarehouseNative) HasCustomRollUpEnd() bool`

HasCustomRollUpEnd returns a boolean if a field has been set.

### GetAllowNullRatioDenominator

`func (o *MetricsUpdateContractDtoWarehouseNative) GetAllowNullRatioDenominator() bool`

GetAllowNullRatioDenominator returns the AllowNullRatioDenominator field if non-nil, zero value otherwise.

### GetAllowNullRatioDenominatorOk

`func (o *MetricsUpdateContractDtoWarehouseNative) GetAllowNullRatioDenominatorOk() (*bool, bool)`

GetAllowNullRatioDenominatorOk returns a tuple with the AllowNullRatioDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNullRatioDenominator

`func (o *MetricsUpdateContractDtoWarehouseNative) SetAllowNullRatioDenominator(v bool)`

SetAllowNullRatioDenominator sets AllowNullRatioDenominator field to given value.

### HasAllowNullRatioDenominator

`func (o *MetricsUpdateContractDtoWarehouseNative) HasAllowNullRatioDenominator() bool`

HasAllowNullRatioDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


