# ExperimentPulseResultsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ds** | **string** |  | 
**PrimaryMetrics** | [**[]GatePulseResultsDtoMonitoringMetricsInner**](GatePulseResultsDtoMonitoringMetricsInner.md) |  | 
**SecondaryMetrics** | [**[]GatePulseResultsDtoMonitoringMetricsInner**](GatePulseResultsDtoMonitoringMetricsInner.md) |  | 

## Methods

### NewExperimentPulseResultsDto

`func NewExperimentPulseResultsDto(ds string, primaryMetrics []GatePulseResultsDtoMonitoringMetricsInner, secondaryMetrics []GatePulseResultsDtoMonitoringMetricsInner, ) *ExperimentPulseResultsDto`

NewExperimentPulseResultsDto instantiates a new ExperimentPulseResultsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentPulseResultsDtoWithDefaults

`func NewExperimentPulseResultsDtoWithDefaults() *ExperimentPulseResultsDto`

NewExperimentPulseResultsDtoWithDefaults instantiates a new ExperimentPulseResultsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDs

`func (o *ExperimentPulseResultsDto) GetDs() string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *ExperimentPulseResultsDto) GetDsOk() (*string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *ExperimentPulseResultsDto) SetDs(v string)`

SetDs sets Ds field to given value.


### GetPrimaryMetrics

`func (o *ExperimentPulseResultsDto) GetPrimaryMetrics() []GatePulseResultsDtoMonitoringMetricsInner`

GetPrimaryMetrics returns the PrimaryMetrics field if non-nil, zero value otherwise.

### GetPrimaryMetricsOk

`func (o *ExperimentPulseResultsDto) GetPrimaryMetricsOk() (*[]GatePulseResultsDtoMonitoringMetricsInner, bool)`

GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetrics

`func (o *ExperimentPulseResultsDto) SetPrimaryMetrics(v []GatePulseResultsDtoMonitoringMetricsInner)`

SetPrimaryMetrics sets PrimaryMetrics field to given value.


### GetSecondaryMetrics

`func (o *ExperimentPulseResultsDto) GetSecondaryMetrics() []GatePulseResultsDtoMonitoringMetricsInner`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *ExperimentPulseResultsDto) GetSecondaryMetricsOk() (*[]GatePulseResultsDtoMonitoringMetricsInner, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *ExperimentPulseResultsDto) SetSecondaryMetrics(v []GatePulseResultsDtoMonitoringMetricsInner)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


