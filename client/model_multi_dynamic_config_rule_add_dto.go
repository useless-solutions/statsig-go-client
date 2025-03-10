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

// checks if the MultiDynamicConfigRuleAddDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiDynamicConfigRuleAddDto{}

// MultiDynamicConfigRuleAddDto struct for MultiDynamicConfigRuleAddDto
type MultiDynamicConfigRuleAddDto struct {
	// A list of new dynamic config rules
	Rules []DynamicConfigDtoRulesInner `json:"rules"`
}

type _MultiDynamicConfigRuleAddDto MultiDynamicConfigRuleAddDto

// NewMultiDynamicConfigRuleAddDto instantiates a new MultiDynamicConfigRuleAddDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiDynamicConfigRuleAddDto(rules []DynamicConfigDtoRulesInner) *MultiDynamicConfigRuleAddDto {
	this := MultiDynamicConfigRuleAddDto{}
	this.Rules = rules
	return &this
}

// NewMultiDynamicConfigRuleAddDtoWithDefaults instantiates a new MultiDynamicConfigRuleAddDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiDynamicConfigRuleAddDtoWithDefaults() *MultiDynamicConfigRuleAddDto {
	this := MultiDynamicConfigRuleAddDto{}
	return &this
}

// GetRules returns the Rules field value
func (o *MultiDynamicConfigRuleAddDto) GetRules() []DynamicConfigDtoRulesInner {
	if o == nil {
		var ret []DynamicConfigDtoRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *MultiDynamicConfigRuleAddDto) GetRulesOk() ([]DynamicConfigDtoRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *MultiDynamicConfigRuleAddDto) SetRules(v []DynamicConfigDtoRulesInner) {
	o.Rules = v
}

func (o MultiDynamicConfigRuleAddDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiDynamicConfigRuleAddDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rules"] = o.Rules
	return toSerialize, nil
}

func (o *MultiDynamicConfigRuleAddDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rules",
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

	varMultiDynamicConfigRuleAddDto := _MultiDynamicConfigRuleAddDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultiDynamicConfigRuleAddDto)

	if err != nil {
		return err
	}

	*o = MultiDynamicConfigRuleAddDto(varMultiDynamicConfigRuleAddDto)

	return err
}

type NullableMultiDynamicConfigRuleAddDto struct {
	value *MultiDynamicConfigRuleAddDto
	isSet bool
}

func (v NullableMultiDynamicConfigRuleAddDto) Get() *MultiDynamicConfigRuleAddDto {
	return v.value
}

func (v *NullableMultiDynamicConfigRuleAddDto) Set(val *MultiDynamicConfigRuleAddDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiDynamicConfigRuleAddDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiDynamicConfigRuleAddDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiDynamicConfigRuleAddDto(val *MultiDynamicConfigRuleAddDto) *NullableMultiDynamicConfigRuleAddDto {
	return &NullableMultiDynamicConfigRuleAddDto{value: val, isSet: true}
}

func (v NullableMultiDynamicConfigRuleAddDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiDynamicConfigRuleAddDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


