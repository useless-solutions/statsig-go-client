# MetricValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** |  | 
**UnitType** | **string** |  | 
**Numerator** | Pointer to **float32** |  | [optional] 
**Denominator** | Pointer to **float32** |  | [optional] 
**InputRows** | Pointer to **float32** |  | [optional] 
**MetricName** | **string** |  | 
**MetricType** | **string** |  | 

## Methods

### NewMetricValueDto

`func NewMetricValueDto(value float32, unitType string, metricName string, metricType string, ) *MetricValueDto`

NewMetricValueDto instantiates a new MetricValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricValueDtoWithDefaults

`func NewMetricValueDtoWithDefaults() *MetricValueDto`

NewMetricValueDtoWithDefaults instantiates a new MetricValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *MetricValueDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricValueDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricValueDto) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnitType

`func (o *MetricValueDto) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *MetricValueDto) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *MetricValueDto) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.


### GetNumerator

`func (o *MetricValueDto) GetNumerator() float32`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *MetricValueDto) GetNumeratorOk() (*float32, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *MetricValueDto) SetNumerator(v float32)`

SetNumerator sets Numerator field to given value.

### HasNumerator

`func (o *MetricValueDto) HasNumerator() bool`

HasNumerator returns a boolean if a field has been set.

### GetDenominator

`func (o *MetricValueDto) GetDenominator() float32`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *MetricValueDto) GetDenominatorOk() (*float32, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *MetricValueDto) SetDenominator(v float32)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *MetricValueDto) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.

### GetInputRows

`func (o *MetricValueDto) GetInputRows() float32`

GetInputRows returns the InputRows field if non-nil, zero value otherwise.

### GetInputRowsOk

`func (o *MetricValueDto) GetInputRowsOk() (*float32, bool)`

GetInputRowsOk returns a tuple with the InputRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRows

`func (o *MetricValueDto) SetInputRows(v float32)`

SetInputRows sets InputRows field to given value.

### HasInputRows

`func (o *MetricValueDto) HasInputRows() bool`

HasInputRows returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricValueDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricValueDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricValueDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricType

`func (o *MetricValueDto) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricValueDto) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricValueDto) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


