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

// checks if the UpdateOverridesContractDtoOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOverridesContractDtoOneOf1{}

// UpdateOverridesContractDtoOneOf1 Contract for updating ID overrides
type UpdateOverridesContractDtoOneOf1 struct {
	// List of user IDs
	PassingUserIDs []string `json:"passingUserIDs"`
	// List of user IDs
	FailingUserIDs []string `json:"failingUserIDs"`
	// Optional list of custom IDs
	PassingCustomIDs []string `json:"passingCustomIDs,omitempty"`
	// Optional list of custom IDs
	FailingCustomIDs []string `json:"failingCustomIDs,omitempty"`
}

type _UpdateOverridesContractDtoOneOf1 UpdateOverridesContractDtoOneOf1

// NewUpdateOverridesContractDtoOneOf1 instantiates a new UpdateOverridesContractDtoOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOverridesContractDtoOneOf1(passingUserIDs []string, failingUserIDs []string) *UpdateOverridesContractDtoOneOf1 {
	this := UpdateOverridesContractDtoOneOf1{}
	this.PassingUserIDs = passingUserIDs
	this.FailingUserIDs = failingUserIDs
	return &this
}

// NewUpdateOverridesContractDtoOneOf1WithDefaults instantiates a new UpdateOverridesContractDtoOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOverridesContractDtoOneOf1WithDefaults() *UpdateOverridesContractDtoOneOf1 {
	this := UpdateOverridesContractDtoOneOf1{}
	return &this
}

// GetPassingUserIDs returns the PassingUserIDs field value
func (o *UpdateOverridesContractDtoOneOf1) GetPassingUserIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PassingUserIDs
}

// GetPassingUserIDsOk returns a tuple with the PassingUserIDs field value
// and a boolean to check if the value has been set.
func (o *UpdateOverridesContractDtoOneOf1) GetPassingUserIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassingUserIDs, true
}

// SetPassingUserIDs sets field value
func (o *UpdateOverridesContractDtoOneOf1) SetPassingUserIDs(v []string) {
	o.PassingUserIDs = v
}

// GetFailingUserIDs returns the FailingUserIDs field value
func (o *UpdateOverridesContractDtoOneOf1) GetFailingUserIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FailingUserIDs
}

// GetFailingUserIDsOk returns a tuple with the FailingUserIDs field value
// and a boolean to check if the value has been set.
func (o *UpdateOverridesContractDtoOneOf1) GetFailingUserIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailingUserIDs, true
}

// SetFailingUserIDs sets field value
func (o *UpdateOverridesContractDtoOneOf1) SetFailingUserIDs(v []string) {
	o.FailingUserIDs = v
}

// GetPassingCustomIDs returns the PassingCustomIDs field value if set, zero value otherwise.
func (o *UpdateOverridesContractDtoOneOf1) GetPassingCustomIDs() []string {
	if o == nil || IsNil(o.PassingCustomIDs) {
		var ret []string
		return ret
	}
	return o.PassingCustomIDs
}

// GetPassingCustomIDsOk returns a tuple with the PassingCustomIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOverridesContractDtoOneOf1) GetPassingCustomIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.PassingCustomIDs) {
		return nil, false
	}
	return o.PassingCustomIDs, true
}

// HasPassingCustomIDs returns a boolean if a field has been set.
func (o *UpdateOverridesContractDtoOneOf1) HasPassingCustomIDs() bool {
	if o != nil && !IsNil(o.PassingCustomIDs) {
		return true
	}

	return false
}

// SetPassingCustomIDs gets a reference to the given []string and assigns it to the PassingCustomIDs field.
func (o *UpdateOverridesContractDtoOneOf1) SetPassingCustomIDs(v []string) {
	o.PassingCustomIDs = v
}

// GetFailingCustomIDs returns the FailingCustomIDs field value if set, zero value otherwise.
func (o *UpdateOverridesContractDtoOneOf1) GetFailingCustomIDs() []string {
	if o == nil || IsNil(o.FailingCustomIDs) {
		var ret []string
		return ret
	}
	return o.FailingCustomIDs
}

// GetFailingCustomIDsOk returns a tuple with the FailingCustomIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOverridesContractDtoOneOf1) GetFailingCustomIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailingCustomIDs) {
		return nil, false
	}
	return o.FailingCustomIDs, true
}

// HasFailingCustomIDs returns a boolean if a field has been set.
func (o *UpdateOverridesContractDtoOneOf1) HasFailingCustomIDs() bool {
	if o != nil && !IsNil(o.FailingCustomIDs) {
		return true
	}

	return false
}

// SetFailingCustomIDs gets a reference to the given []string and assigns it to the FailingCustomIDs field.
func (o *UpdateOverridesContractDtoOneOf1) SetFailingCustomIDs(v []string) {
	o.FailingCustomIDs = v
}

func (o UpdateOverridesContractDtoOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOverridesContractDtoOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["passingUserIDs"] = o.PassingUserIDs
	toSerialize["failingUserIDs"] = o.FailingUserIDs
	if !IsNil(o.PassingCustomIDs) {
		toSerialize["passingCustomIDs"] = o.PassingCustomIDs
	}
	if !IsNil(o.FailingCustomIDs) {
		toSerialize["failingCustomIDs"] = o.FailingCustomIDs
	}
	return toSerialize, nil
}

func (o *UpdateOverridesContractDtoOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"passingUserIDs",
		"failingUserIDs",
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

	varUpdateOverridesContractDtoOneOf1 := _UpdateOverridesContractDtoOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOverridesContractDtoOneOf1)

	if err != nil {
		return err
	}

	*o = UpdateOverridesContractDtoOneOf1(varUpdateOverridesContractDtoOneOf1)

	return err
}

type NullableUpdateOverridesContractDtoOneOf1 struct {
	value *UpdateOverridesContractDtoOneOf1
	isSet bool
}

func (v NullableUpdateOverridesContractDtoOneOf1) Get() *UpdateOverridesContractDtoOneOf1 {
	return v.value
}

func (v *NullableUpdateOverridesContractDtoOneOf1) Set(val *UpdateOverridesContractDtoOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOverridesContractDtoOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOverridesContractDtoOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOverridesContractDtoOneOf1(val *UpdateOverridesContractDtoOneOf1) *NullableUpdateOverridesContractDtoOneOf1 {
	return &NullableUpdateOverridesContractDtoOneOf1{value: val, isSet: true}
}

func (v NullableUpdateOverridesContractDtoOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOverridesContractDtoOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


