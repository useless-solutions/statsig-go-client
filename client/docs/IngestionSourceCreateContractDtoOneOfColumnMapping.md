# IngestionSourceCreateContractDtoOneOfColumnMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitId** | **string** | The unique user identifier this metric is for. This might not necessarily be a user_id - it could be a custom_id of some kind. Make sure this is in the same format as your logged unit_ids. | 
**IdType** | **string** | The id_type the unit_id represents. Must be valid id_type. Default Statsig types are user_id/stable_id, but you may have generated custom id_types. Make sure this matches (case sensitive) a customID in your project, or you won’t get experiment results. | 
**Dateid** | **string** | Date of the daily metric, ISO formatted (ex. 2024-10-23). We’ll load custom metrics to whatever date you use here. | 
**MetricName** | **string** | String format. Not null. Length &lt; 128 characters. | 
**MetricValue** | Pointer to **string** | Numeric value for the metric. This OR both of numerator and denominator need to be provided. | [optional] [default to "null"]
**Numerator** | Pointer to **string** | Required for ratio metrics. If present along with a denominator in any record, the metric will be treated as ratio and only calculated for users with non-null denominators | [optional] [default to "null"]
**Denominator** | Pointer to **string** | Required for ratio metrics. If present along with a numerator in any record, the metric will be treated as ratio and only calculated for users with non-null numerators. | [optional] [default to "null"]

## Methods

### NewIngestionSourceCreateContractDtoOneOfColumnMapping

`func NewIngestionSourceCreateContractDtoOneOfColumnMapping(unitId string, idType string, dateid string, metricName string, ) *IngestionSourceCreateContractDtoOneOfColumnMapping`

NewIngestionSourceCreateContractDtoOneOfColumnMapping instantiates a new IngestionSourceCreateContractDtoOneOfColumnMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoOneOfColumnMappingWithDefaults

`func NewIngestionSourceCreateContractDtoOneOfColumnMappingWithDefaults() *IngestionSourceCreateContractDtoOneOfColumnMapping`

NewIngestionSourceCreateContractDtoOneOfColumnMappingWithDefaults instantiates a new IngestionSourceCreateContractDtoOneOfColumnMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitId

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.


### GetIdType

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetDateid

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetDateid() string`

GetDateid returns the Dateid field if non-nil, zero value otherwise.

### GetDateidOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetDateidOk() (*string, bool)`

GetDateidOk returns a tuple with the Dateid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateid

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetDateid(v string)`

SetDateid sets Dateid field to given value.


### GetMetricName

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricValue

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetMetricValue() string`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetMetricValueOk() (*string, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetMetricValue(v string)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.

### GetNumerator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetNumerator() string`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetNumeratorOk() (*string, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetNumerator(v string)`

SetNumerator sets Numerator field to given value.

### HasNumerator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) HasNumerator() bool`

HasNumerator returns a boolean if a field has been set.

### GetDenominator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetDenominator() string`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) GetDenominatorOk() (*string, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) SetDenominator(v string)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *IngestionSourceCreateContractDtoOneOfColumnMapping) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


