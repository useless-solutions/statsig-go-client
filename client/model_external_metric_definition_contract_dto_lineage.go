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

// checks if the ExternalMetricDefinitionContractDtoLineage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalMetricDefinitionContractDtoLineage{}

// ExternalMetricDefinitionContractDtoLineage struct for ExternalMetricDefinitionContractDtoLineage
type ExternalMetricDefinitionContractDtoLineage struct {
	// List of event names that contribute to the metric’s calculation.
	Events []string `json:"events"`
	// List of metric names that are part of this metric’s lineage.
	Metrics []string `json:"metrics"`
}

type _ExternalMetricDefinitionContractDtoLineage ExternalMetricDefinitionContractDtoLineage

// NewExternalMetricDefinitionContractDtoLineage instantiates a new ExternalMetricDefinitionContractDtoLineage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMetricDefinitionContractDtoLineage(events []string, metrics []string) *ExternalMetricDefinitionContractDtoLineage {
	this := ExternalMetricDefinitionContractDtoLineage{}
	this.Events = events
	this.Metrics = metrics
	return &this
}

// NewExternalMetricDefinitionContractDtoLineageWithDefaults instantiates a new ExternalMetricDefinitionContractDtoLineage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMetricDefinitionContractDtoLineageWithDefaults() *ExternalMetricDefinitionContractDtoLineage {
	this := ExternalMetricDefinitionContractDtoLineage{}
	return &this
}

// GetEvents returns the Events field value
func (o *ExternalMetricDefinitionContractDtoLineage) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoLineage) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ExternalMetricDefinitionContractDtoLineage) SetEvents(v []string) {
	o.Events = v
}

// GetMetrics returns the Metrics field value
func (o *ExternalMetricDefinitionContractDtoLineage) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoLineage) GetMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *ExternalMetricDefinitionContractDtoLineage) SetMetrics(v []string) {
	o.Metrics = v
}

func (o ExternalMetricDefinitionContractDtoLineage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalMetricDefinitionContractDtoLineage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

func (o *ExternalMetricDefinitionContractDtoLineage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
		"metrics",
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

	varExternalMetricDefinitionContractDtoLineage := _ExternalMetricDefinitionContractDtoLineage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalMetricDefinitionContractDtoLineage)

	if err != nil {
		return err
	}

	*o = ExternalMetricDefinitionContractDtoLineage(varExternalMetricDefinitionContractDtoLineage)

	return err
}

type NullableExternalMetricDefinitionContractDtoLineage struct {
	value *ExternalMetricDefinitionContractDtoLineage
	isSet bool
}

func (v NullableExternalMetricDefinitionContractDtoLineage) Get() *ExternalMetricDefinitionContractDtoLineage {
	return v.value
}

func (v *NullableExternalMetricDefinitionContractDtoLineage) Set(val *ExternalMetricDefinitionContractDtoLineage) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMetricDefinitionContractDtoLineage) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMetricDefinitionContractDtoLineage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMetricDefinitionContractDtoLineage(val *ExternalMetricDefinitionContractDtoLineage) *NullableExternalMetricDefinitionContractDtoLineage {
	return &NullableExternalMetricDefinitionContractDtoLineage{value: val, isSet: true}
}

func (v NullableExternalMetricDefinitionContractDtoLineage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMetricDefinitionContractDtoLineage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


