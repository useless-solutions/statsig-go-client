# ExternalMetricDefinitionContractDtoLineage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** | List of event names that contribute to the metric’s calculation. | 
**Metrics** | **[]string** | List of metric names that are part of this metric’s lineage. | 

## Methods

### NewExternalMetricDefinitionContractDtoLineage

`func NewExternalMetricDefinitionContractDtoLineage(events []string, metrics []string, ) *ExternalMetricDefinitionContractDtoLineage`

NewExternalMetricDefinitionContractDtoLineage instantiates a new ExternalMetricDefinitionContractDtoLineage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMetricDefinitionContractDtoLineageWithDefaults

`func NewExternalMetricDefinitionContractDtoLineageWithDefaults() *ExternalMetricDefinitionContractDtoLineage`

NewExternalMetricDefinitionContractDtoLineageWithDefaults instantiates a new ExternalMetricDefinitionContractDtoLineage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ExternalMetricDefinitionContractDtoLineage) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ExternalMetricDefinitionContractDtoLineage) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ExternalMetricDefinitionContractDtoLineage) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetMetrics

`func (o *ExternalMetricDefinitionContractDtoLineage) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ExternalMetricDefinitionContractDtoLineage) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ExternalMetricDefinitionContractDtoLineage) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


