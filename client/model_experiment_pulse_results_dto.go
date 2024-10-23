/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExperimentPulseResultsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperimentPulseResultsDto{}

// ExperimentPulseResultsDto struct for ExperimentPulseResultsDto
type ExperimentPulseResultsDto struct {
	Ds string `json:"ds"`
	PrimaryMetrics []GatePulseResultsDtoMonitoringMetricsInner `json:"primaryMetrics"`
	SecondaryMetrics []GatePulseResultsDtoMonitoringMetricsInner `json:"secondaryMetrics"`
}

type _ExperimentPulseResultsDto ExperimentPulseResultsDto

// NewExperimentPulseResultsDto instantiates a new ExperimentPulseResultsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentPulseResultsDto(ds string, primaryMetrics []GatePulseResultsDtoMonitoringMetricsInner, secondaryMetrics []GatePulseResultsDtoMonitoringMetricsInner) *ExperimentPulseResultsDto {
	this := ExperimentPulseResultsDto{}
	this.Ds = ds
	this.PrimaryMetrics = primaryMetrics
	this.SecondaryMetrics = secondaryMetrics
	return &this
}

// NewExperimentPulseResultsDtoWithDefaults instantiates a new ExperimentPulseResultsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentPulseResultsDtoWithDefaults() *ExperimentPulseResultsDto {
	this := ExperimentPulseResultsDto{}
	return &this
}

// GetDs returns the Ds field value
func (o *ExperimentPulseResultsDto) GetDs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ds
}

// GetDsOk returns a tuple with the Ds field value
// and a boolean to check if the value has been set.
func (o *ExperimentPulseResultsDto) GetDsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ds, true
}

// SetDs sets field value
func (o *ExperimentPulseResultsDto) SetDs(v string) {
	o.Ds = v
}

// GetPrimaryMetrics returns the PrimaryMetrics field value
func (o *ExperimentPulseResultsDto) GetPrimaryMetrics() []GatePulseResultsDtoMonitoringMetricsInner {
	if o == nil {
		var ret []GatePulseResultsDtoMonitoringMetricsInner
		return ret
	}

	return o.PrimaryMetrics
}

// GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field value
// and a boolean to check if the value has been set.
func (o *ExperimentPulseResultsDto) GetPrimaryMetricsOk() ([]GatePulseResultsDtoMonitoringMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMetrics, true
}

// SetPrimaryMetrics sets field value
func (o *ExperimentPulseResultsDto) SetPrimaryMetrics(v []GatePulseResultsDtoMonitoringMetricsInner) {
	o.PrimaryMetrics = v
}

// GetSecondaryMetrics returns the SecondaryMetrics field value
func (o *ExperimentPulseResultsDto) GetSecondaryMetrics() []GatePulseResultsDtoMonitoringMetricsInner {
	if o == nil {
		var ret []GatePulseResultsDtoMonitoringMetricsInner
		return ret
	}

	return o.SecondaryMetrics
}

// GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field value
// and a boolean to check if the value has been set.
func (o *ExperimentPulseResultsDto) GetSecondaryMetricsOk() ([]GatePulseResultsDtoMonitoringMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryMetrics, true
}

// SetSecondaryMetrics sets field value
func (o *ExperimentPulseResultsDto) SetSecondaryMetrics(v []GatePulseResultsDtoMonitoringMetricsInner) {
	o.SecondaryMetrics = v
}

func (o ExperimentPulseResultsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentPulseResultsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ds"] = o.Ds
	toSerialize["primaryMetrics"] = o.PrimaryMetrics
	toSerialize["secondaryMetrics"] = o.SecondaryMetrics
	return toSerialize, nil
}

func (o *ExperimentPulseResultsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ds",
		"primaryMetrics",
		"secondaryMetrics",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExperimentPulseResultsDto := _ExperimentPulseResultsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExperimentPulseResultsDto)

	if err != nil {
		return err
	}

	*o = ExperimentPulseResultsDto(varExperimentPulseResultsDto)

	return err
}

type NullableExperimentPulseResultsDto struct {
	value *ExperimentPulseResultsDto
	isSet bool
}

func (v NullableExperimentPulseResultsDto) Get() *ExperimentPulseResultsDto {
	return v.value
}

func (v *NullableExperimentPulseResultsDto) Set(val *ExperimentPulseResultsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentPulseResultsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentPulseResultsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentPulseResultsDto(val *ExperimentPulseResultsDto) *NullableExperimentPulseResultsDto {
	return &NullableExperimentPulseResultsDto{value: val, isSet: true}
}

func (v NullableExperimentPulseResultsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentPulseResultsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


