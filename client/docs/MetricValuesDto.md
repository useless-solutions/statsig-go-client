# MetricValuesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | Metric value for the given unit_type | 
**UnitType** | **string** | Unit of the metric: stableID, userID, and other custom ids | 
**RowCount** | Pointer to **float32** | Row count for imported metric, optional | [optional] 
**Numerator** | Pointer to **float32** | Numerator of a ratio metric, optional | [optional] 
**Denominator** | Pointer to **float32** | Denominator of a ratio metric, optional | [optional] 

## Methods

### NewMetricValuesDto

`func NewMetricValuesDto(value float32, unitType string, ) *MetricValuesDto`

NewMetricValuesDto instantiates a new MetricValuesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricValuesDtoWithDefaults

`func NewMetricValuesDtoWithDefaults() *MetricValuesDto`

NewMetricValuesDtoWithDefaults instantiates a new MetricValuesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MetricValuesDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricValuesDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricValuesDto) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnitType

`func (o *MetricValuesDto) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *MetricValuesDto) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *MetricValuesDto) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.


### GetRowCount

`func (o *MetricValuesDto) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *MetricValuesDto) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *MetricValuesDto) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *MetricValuesDto) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetNumerator

`func (o *MetricValuesDto) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *MetricValuesDto) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *MetricValuesDto) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.

### HasNumerator

`func (o *MetricValuesDto) HasNumerator() bool`

HasNumerator returns a boolean if a field has been set.

### GetDenominator

`func (o *MetricValuesDto) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *MetricValuesDto) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *MetricValuesDto) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *MetricValuesDto) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


