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

// checks if the ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner{}

// ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner struct for ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner
type ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner struct {
	// Type of event criterion for filtering metrics. Options include `value`, `metadata`, `user`, and `user_custom`.
	Type string `json:"type"`
	// Optional column specifying which data attribute to filter on.
	Column *string `json:"column,omitempty"`
	// sql_filter, start_withs, ends_with, and after_exposure are only applicable in Warehouse Native
	Condition string `json:"condition"`
	// Optional array of values for the criterion to match against.
	Values []string `json:"values,omitempty"`
	// If true, overrides null values in criterion evaluation.
	NullVacuousOverride *bool `json:"nullVacuousOverride,omitempty"`
}

type _ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner

// NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner(type_ string, condition string) *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	this := ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner{}
	this.Type = type_
	this.Condition = condition
	return &this
}

// NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInnerWithDefaults() *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	this := ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetType(v string) {
	o.Type = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetColumn() string {
	if o == nil || IsNil(o.Column) {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetColumnOk() (*string, bool) {
	if o == nil || IsNil(o.Column) {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasColumn() bool {
	if o != nil && !IsNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetColumn(v string) {
	o.Column = &v
}

// GetCondition returns the Condition field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetCondition(v string) {
	o.Condition = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetValues(v []string) {
	o.Values = v
}

// GetNullVacuousOverride returns the NullVacuousOverride field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetNullVacuousOverride() bool {
	if o == nil || IsNil(o.NullVacuousOverride) {
		var ret bool
		return ret
	}
	return *o.NullVacuousOverride
}

// GetNullVacuousOverrideOk returns a tuple with the NullVacuousOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) GetNullVacuousOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.NullVacuousOverride) {
		return nil, false
	}
	return o.NullVacuousOverride, true
}

// HasNullVacuousOverride returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) HasNullVacuousOverride() bool {
	if o != nil && !IsNil(o.NullVacuousOverride) {
		return true
	}

	return false
}

// SetNullVacuousOverride gets a reference to the given bool and assigns it to the NullVacuousOverride field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) SetNullVacuousOverride(v bool) {
	o.NullVacuousOverride = &v
}

func (o ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	toSerialize["condition"] = o.Condition
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.NullVacuousOverride) {
		toSerialize["nullVacuousOverride"] = o.NullVacuousOverride
	}
	return toSerialize, nil
}

func (o *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"condition",
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

	varExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner := _ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)

	if err != nil {
		return err
	}

	*o = ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner(varExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)

	return err
}

type NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner struct {
	value *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner
	isSet bool
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) Get() *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	return v.value
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) Set(val *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner(val *ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) *NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	return &NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner{value: val, isSet: true}
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


