# GatePulseResultsDtoMonitoringMetricsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricID** | **string** |  | 
**AbsoluteChange** | Pointer to **float32** |  | [optional] 
**AbsoluteConfidenceIntervalDelta** | Pointer to **float32** |  | [optional] 
**PercentChange** | Pointer to **float32** |  | [optional] 
**SequentialTestingConfidenceIntervalDelta** | Pointer to **float32** |  | [optional] 
**PercentConfidenceIntervalDelta** | Pointer to **float32** |  | [optional] 
**PercentSequentialTestingConfidenceIntervalDelta** | Pointer to **float32** |  | [optional] 
**TestMean** | Pointer to **float32** |  | [optional] 
**ControlMean** | Pointer to **float32** |  | [optional] 
**TestStd** | Pointer to **float32** |  | [optional] 
**ControlStd** | Pointer to **float32** |  | [optional] 
**TestUnits** | Pointer to **float32** |  | [optional] 
**ControlUnits** | Pointer to **float32** |  | [optional] 
**PValue** | Pointer to **float32** |  | [optional] 

## Methods

### NewGatePulseResultsDtoMonitoringMetricsInner

`func NewGatePulseResultsDtoMonitoringMetricsInner(metricID string, ) *GatePulseResultsDtoMonitoringMetricsInner`

NewGatePulseResultsDtoMonitoringMetricsInner instantiates a new GatePulseResultsDtoMonitoringMetricsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatePulseResultsDtoMonitoringMetricsInnerWithDefaults

`func NewGatePulseResultsDtoMonitoringMetricsInnerWithDefaults() *GatePulseResultsDtoMonitoringMetricsInner`

NewGatePulseResultsDtoMonitoringMetricsInnerWithDefaults instantiates a new GatePulseResultsDtoMonitoringMetricsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricID

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetMetricID() string`

GetMetricID returns the MetricID field if non-nil, zero value otherwise.

### GetMetricIDOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetMetricIDOk() (*string, bool)`

GetMetricIDOk returns a tuple with the MetricID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricID

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetMetricID(v string)`

SetMetricID sets MetricID field to given value.


### GetAbsoluteChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetAbsoluteChange() float32`

GetAbsoluteChange returns the AbsoluteChange field if non-nil, zero value otherwise.

### GetAbsoluteChangeOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetAbsoluteChangeOk() (*float32, bool)`

GetAbsoluteChangeOk returns a tuple with the AbsoluteChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetAbsoluteChange(v float32)`

SetAbsoluteChange sets AbsoluteChange field to given value.

### HasAbsoluteChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasAbsoluteChange() bool`

HasAbsoluteChange returns a boolean if a field has been set.

### GetAbsoluteConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetAbsoluteConfidenceIntervalDelta() float32`

GetAbsoluteConfidenceIntervalDelta returns the AbsoluteConfidenceIntervalDelta field if non-nil, zero value otherwise.

### GetAbsoluteConfidenceIntervalDeltaOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetAbsoluteConfidenceIntervalDeltaOk() (*float32, bool)`

GetAbsoluteConfidenceIntervalDeltaOk returns a tuple with the AbsoluteConfidenceIntervalDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetAbsoluteConfidenceIntervalDelta(v float32)`

SetAbsoluteConfidenceIntervalDelta sets AbsoluteConfidenceIntervalDelta field to given value.

### HasAbsoluteConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasAbsoluteConfidenceIntervalDelta() bool`

HasAbsoluteConfidenceIntervalDelta returns a boolean if a field has been set.

### GetPercentChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentChange() float32`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentChangeOk() (*float32, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetPercentChange(v float32)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetSequentialTestingConfidenceIntervalDelta() float32`

GetSequentialTestingConfidenceIntervalDelta returns the SequentialTestingConfidenceIntervalDelta field if non-nil, zero value otherwise.

### GetSequentialTestingConfidenceIntervalDeltaOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetSequentialTestingConfidenceIntervalDeltaOk() (*float32, bool)`

GetSequentialTestingConfidenceIntervalDeltaOk returns a tuple with the SequentialTestingConfidenceIntervalDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetSequentialTestingConfidenceIntervalDelta(v float32)`

SetSequentialTestingConfidenceIntervalDelta sets SequentialTestingConfidenceIntervalDelta field to given value.

### HasSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasSequentialTestingConfidenceIntervalDelta() bool`

HasSequentialTestingConfidenceIntervalDelta returns a boolean if a field has been set.

### GetPercentConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentConfidenceIntervalDelta() float32`

GetPercentConfidenceIntervalDelta returns the PercentConfidenceIntervalDelta field if non-nil, zero value otherwise.

### GetPercentConfidenceIntervalDeltaOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentConfidenceIntervalDeltaOk() (*float32, bool)`

GetPercentConfidenceIntervalDeltaOk returns a tuple with the PercentConfidenceIntervalDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetPercentConfidenceIntervalDelta(v float32)`

SetPercentConfidenceIntervalDelta sets PercentConfidenceIntervalDelta field to given value.

### HasPercentConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasPercentConfidenceIntervalDelta() bool`

HasPercentConfidenceIntervalDelta returns a boolean if a field has been set.

### GetPercentSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentSequentialTestingConfidenceIntervalDelta() float32`

GetPercentSequentialTestingConfidenceIntervalDelta returns the PercentSequentialTestingConfidenceIntervalDelta field if non-nil, zero value otherwise.

### GetPercentSequentialTestingConfidenceIntervalDeltaOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPercentSequentialTestingConfidenceIntervalDeltaOk() (*float32, bool)`

GetPercentSequentialTestingConfidenceIntervalDeltaOk returns a tuple with the PercentSequentialTestingConfidenceIntervalDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetPercentSequentialTestingConfidenceIntervalDelta(v float32)`

SetPercentSequentialTestingConfidenceIntervalDelta sets PercentSequentialTestingConfidenceIntervalDelta field to given value.

### HasPercentSequentialTestingConfidenceIntervalDelta

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasPercentSequentialTestingConfidenceIntervalDelta() bool`

HasPercentSequentialTestingConfidenceIntervalDelta returns a boolean if a field has been set.

### GetTestMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestMean() float32`

GetTestMean returns the TestMean field if non-nil, zero value otherwise.

### GetTestMeanOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestMeanOk() (*float32, bool)`

GetTestMeanOk returns a tuple with the TestMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetTestMean(v float32)`

SetTestMean sets TestMean field to given value.

### HasTestMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasTestMean() bool`

HasTestMean returns a boolean if a field has been set.

### GetControlMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlMean() float32`

GetControlMean returns the ControlMean field if non-nil, zero value otherwise.

### GetControlMeanOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlMeanOk() (*float32, bool)`

GetControlMeanOk returns a tuple with the ControlMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetControlMean(v float32)`

SetControlMean sets ControlMean field to given value.

### HasControlMean

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasControlMean() bool`

HasControlMean returns a boolean if a field has been set.

### GetTestStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestStd() float32`

GetTestStd returns the TestStd field if non-nil, zero value otherwise.

### GetTestStdOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestStdOk() (*float32, bool)`

GetTestStdOk returns a tuple with the TestStd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetTestStd(v float32)`

SetTestStd sets TestStd field to given value.

### HasTestStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasTestStd() bool`

HasTestStd returns a boolean if a field has been set.

### GetControlStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlStd() float32`

GetControlStd returns the ControlStd field if non-nil, zero value otherwise.

### GetControlStdOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlStdOk() (*float32, bool)`

GetControlStdOk returns a tuple with the ControlStd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetControlStd(v float32)`

SetControlStd sets ControlStd field to given value.

### HasControlStd

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasControlStd() bool`

HasControlStd returns a boolean if a field has been set.

### GetTestUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestUnits() float32`

GetTestUnits returns the TestUnits field if non-nil, zero value otherwise.

### GetTestUnitsOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetTestUnitsOk() (*float32, bool)`

GetTestUnitsOk returns a tuple with the TestUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetTestUnits(v float32)`

SetTestUnits sets TestUnits field to given value.

### HasTestUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasTestUnits() bool`

HasTestUnits returns a boolean if a field has been set.

### GetControlUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlUnits() float32`

GetControlUnits returns the ControlUnits field if non-nil, zero value otherwise.

### GetControlUnitsOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetControlUnitsOk() (*float32, bool)`

GetControlUnitsOk returns a tuple with the ControlUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetControlUnits(v float32)`

SetControlUnits sets ControlUnits field to given value.

### HasControlUnits

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasControlUnits() bool`

HasControlUnits returns a boolean if a field has been set.

### GetPValue

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *GatePulseResultsDtoMonitoringMetricsInner) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *GatePulseResultsDtoMonitoringMetricsInner) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *GatePulseResultsDtoMonitoringMetricsInner) HasPValue() bool`

HasPValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


