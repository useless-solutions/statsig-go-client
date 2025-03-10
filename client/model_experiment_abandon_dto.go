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

// checks if the ExperimentAbandonDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperimentAbandonDto{}

// ExperimentAbandonDto struct for ExperimentAbandonDto
type ExperimentAbandonDto struct {
	// The reason for making the decision to update the experiment status
	DecisionReason string `json:"decisionReason"`
}

type _ExperimentAbandonDto ExperimentAbandonDto

// NewExperimentAbandonDto instantiates a new ExperimentAbandonDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentAbandonDto(decisionReason string) *ExperimentAbandonDto {
	this := ExperimentAbandonDto{}
	this.DecisionReason = decisionReason
	return &this
}

// NewExperimentAbandonDtoWithDefaults instantiates a new ExperimentAbandonDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentAbandonDtoWithDefaults() *ExperimentAbandonDto {
	this := ExperimentAbandonDto{}
	return &this
}

// GetDecisionReason returns the DecisionReason field value
func (o *ExperimentAbandonDto) GetDecisionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DecisionReason
}

// GetDecisionReasonOk returns a tuple with the DecisionReason field value
// and a boolean to check if the value has been set.
func (o *ExperimentAbandonDto) GetDecisionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecisionReason, true
}

// SetDecisionReason sets field value
func (o *ExperimentAbandonDto) SetDecisionReason(v string) {
	o.DecisionReason = v
}

func (o ExperimentAbandonDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentAbandonDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["decisionReason"] = o.DecisionReason
	return toSerialize, nil
}

func (o *ExperimentAbandonDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"decisionReason",
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

	varExperimentAbandonDto := _ExperimentAbandonDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExperimentAbandonDto)

	if err != nil {
		return err
	}

	*o = ExperimentAbandonDto(varExperimentAbandonDto)

	return err
}

type NullableExperimentAbandonDto struct {
	value *ExperimentAbandonDto
	isSet bool
}

func (v NullableExperimentAbandonDto) Get() *ExperimentAbandonDto {
	return v.value
}

func (v *NullableExperimentAbandonDto) Set(val *ExperimentAbandonDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentAbandonDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentAbandonDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentAbandonDto(val *ExperimentAbandonDto) *NullableExperimentAbandonDto {
	return &NullableExperimentAbandonDto{value: val, isSet: true}
}

func (v NullableExperimentAbandonDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentAbandonDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


