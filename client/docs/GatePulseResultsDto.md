# GatePulseResultsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ds** | **string** |  | 
**MonitoringMetrics** | [**[]GatePulseResultsDtoMonitoringMetricsInner**](GatePulseResultsDtoMonitoringMetricsInner.md) |  | 

## Methods

### NewGatePulseResultsDto

`func NewGatePulseResultsDto(ds string, monitoringMetrics []GatePulseResultsDtoMonitoringMetricsInner, ) *GatePulseResultsDto`

NewGatePulseResultsDto instantiates a new GatePulseResultsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatePulseResultsDtoWithDefaults

`func NewGatePulseResultsDtoWithDefaults() *GatePulseResultsDto`

NewGatePulseResultsDtoWithDefaults instantiates a new GatePulseResultsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDs

`func (o *GatePulseResultsDto) GetDs() string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *GatePulseResultsDto) GetDsOk() (*string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *GatePulseResultsDto) SetDs(v string)`

SetDs sets Ds field to given value.


### GetMonitoringMetrics

`func (o *GatePulseResultsDto) GetMonitoringMetrics() []GatePulseResultsDtoMonitoringMetricsInner`

GetMonitoringMetrics returns the MonitoringMetrics field if non-nil, zero value otherwise.

### GetMonitoringMetricsOk

`func (o *GatePulseResultsDto) GetMonitoringMetricsOk() (*[]GatePulseResultsDtoMonitoringMetricsInner, bool)`

GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMetrics

`func (o *GatePulseResultsDto) SetMonitoringMetrics(v []GatePulseResultsDtoMonitoringMetricsInner)`

SetMonitoringMetrics sets MonitoringMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


