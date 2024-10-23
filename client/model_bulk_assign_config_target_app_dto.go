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

// checks if the BulkAssignConfigTargetAppDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkAssignConfigTargetAppDto{}

// BulkAssignConfigTargetAppDto struct for BulkAssignConfigTargetAppDto
type BulkAssignConfigTargetAppDto struct {
	// target app ids
	TargetApps []string `json:"targetApps"`
	// Gate IDs to assign to target app(s)
	Gates []string `json:"gates,omitempty"`
	// Dynamic Config IDs to assign to target app(s)
	DynamicConfigs []string `json:"dynamicConfigs,omitempty"`
	// Experiment IDs to assign to target app(s)
	Experiments []string `json:"experiments,omitempty"`
}

type _BulkAssignConfigTargetAppDto BulkAssignConfigTargetAppDto

// NewBulkAssignConfigTargetAppDto instantiates a new BulkAssignConfigTargetAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAssignConfigTargetAppDto(targetApps []string) *BulkAssignConfigTargetAppDto {
	this := BulkAssignConfigTargetAppDto{}
	this.TargetApps = targetApps
	return &this
}

// NewBulkAssignConfigTargetAppDtoWithDefaults instantiates a new BulkAssignConfigTargetAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAssignConfigTargetAppDtoWithDefaults() *BulkAssignConfigTargetAppDto {
	this := BulkAssignConfigTargetAppDto{}
	return &this
}

// GetTargetApps returns the TargetApps field value
func (o *BulkAssignConfigTargetAppDto) GetTargetApps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetApps
}

// GetTargetAppsOk returns a tuple with the TargetApps field value
// and a boolean to check if the value has been set.
func (o *BulkAssignConfigTargetAppDto) GetTargetAppsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetApps, true
}

// SetTargetApps sets field value
func (o *BulkAssignConfigTargetAppDto) SetTargetApps(v []string) {
	o.TargetApps = v
}

// GetGates returns the Gates field value if set, zero value otherwise.
func (o *BulkAssignConfigTargetAppDto) GetGates() []string {
	if o == nil || IsNil(o.Gates) {
		var ret []string
		return ret
	}
	return o.Gates
}

// GetGatesOk returns a tuple with the Gates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAssignConfigTargetAppDto) GetGatesOk() ([]string, bool) {
	if o == nil || IsNil(o.Gates) {
		return nil, false
	}
	return o.Gates, true
}

// HasGates returns a boolean if a field has been set.
func (o *BulkAssignConfigTargetAppDto) HasGates() bool {
	if o != nil && !IsNil(o.Gates) {
		return true
	}

	return false
}

// SetGates gets a reference to the given []string and assigns it to the Gates field.
func (o *BulkAssignConfigTargetAppDto) SetGates(v []string) {
	o.Gates = v
}

// GetDynamicConfigs returns the DynamicConfigs field value if set, zero value otherwise.
func (o *BulkAssignConfigTargetAppDto) GetDynamicConfigs() []string {
	if o == nil || IsNil(o.DynamicConfigs) {
		var ret []string
		return ret
	}
	return o.DynamicConfigs
}

// GetDynamicConfigsOk returns a tuple with the DynamicConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAssignConfigTargetAppDto) GetDynamicConfigsOk() ([]string, bool) {
	if o == nil || IsNil(o.DynamicConfigs) {
		return nil, false
	}
	return o.DynamicConfigs, true
}

// HasDynamicConfigs returns a boolean if a field has been set.
func (o *BulkAssignConfigTargetAppDto) HasDynamicConfigs() bool {
	if o != nil && !IsNil(o.DynamicConfigs) {
		return true
	}

	return false
}

// SetDynamicConfigs gets a reference to the given []string and assigns it to the DynamicConfigs field.
func (o *BulkAssignConfigTargetAppDto) SetDynamicConfigs(v []string) {
	o.DynamicConfigs = v
}

// GetExperiments returns the Experiments field value if set, zero value otherwise.
func (o *BulkAssignConfigTargetAppDto) GetExperiments() []string {
	if o == nil || IsNil(o.Experiments) {
		var ret []string
		return ret
	}
	return o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAssignConfigTargetAppDto) GetExperimentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Experiments) {
		return nil, false
	}
	return o.Experiments, true
}

// HasExperiments returns a boolean if a field has been set.
func (o *BulkAssignConfigTargetAppDto) HasExperiments() bool {
	if o != nil && !IsNil(o.Experiments) {
		return true
	}

	return false
}

// SetExperiments gets a reference to the given []string and assigns it to the Experiments field.
func (o *BulkAssignConfigTargetAppDto) SetExperiments(v []string) {
	o.Experiments = v
}

func (o BulkAssignConfigTargetAppDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkAssignConfigTargetAppDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetApps"] = o.TargetApps
	if !IsNil(o.Gates) {
		toSerialize["gates"] = o.Gates
	}
	if !IsNil(o.DynamicConfigs) {
		toSerialize["dynamicConfigs"] = o.DynamicConfigs
	}
	if !IsNil(o.Experiments) {
		toSerialize["experiments"] = o.Experiments
	}
	return toSerialize, nil
}

func (o *BulkAssignConfigTargetAppDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetApps",
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

	varBulkAssignConfigTargetAppDto := _BulkAssignConfigTargetAppDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkAssignConfigTargetAppDto)

	if err != nil {
		return err
	}

	*o = BulkAssignConfigTargetAppDto(varBulkAssignConfigTargetAppDto)

	return err
}

type NullableBulkAssignConfigTargetAppDto struct {
	value *BulkAssignConfigTargetAppDto
	isSet bool
}

func (v NullableBulkAssignConfigTargetAppDto) Get() *BulkAssignConfigTargetAppDto {
	return v.value
}

func (v *NullableBulkAssignConfigTargetAppDto) Set(val *BulkAssignConfigTargetAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAssignConfigTargetAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAssignConfigTargetAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAssignConfigTargetAppDto(val *BulkAssignConfigTargetAppDto) *NullableBulkAssignConfigTargetAppDto {
	return &NullableBulkAssignConfigTargetAppDto{value: val, isSet: true}
}

func (v NullableBulkAssignConfigTargetAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAssignConfigTargetAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

