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

// checks if the TargetAppDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetAppDto{}

// TargetAppDto struct for TargetAppDto
type TargetAppDto struct {
	// id of target app
	Id *string `json:"id,omitempty"`
	// name of the target app
	Name string `json:"name"`
}

type _TargetAppDto TargetAppDto

// NewTargetAppDto instantiates a new TargetAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetAppDto(name string) *TargetAppDto {
	this := TargetAppDto{}
	this.Name = name
	return &this
}

// NewTargetAppDtoWithDefaults instantiates a new TargetAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetAppDtoWithDefaults() *TargetAppDto {
	this := TargetAppDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TargetAppDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetAppDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TargetAppDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TargetAppDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *TargetAppDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TargetAppDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TargetAppDto) SetName(v string) {
	o.Name = v
}

func (o TargetAppDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetAppDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *TargetAppDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTargetAppDto := _TargetAppDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetAppDto)

	if err != nil {
		return err
	}

	*o = TargetAppDto(varTargetAppDto)

	return err
}

type NullableTargetAppDto struct {
	value *TargetAppDto
	isSet bool
}

func (v NullableTargetAppDto) Get() *TargetAppDto {
	return v.value
}

func (v *NullableTargetAppDto) Set(val *TargetAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetAppDto(val *TargetAppDto) *NullableTargetAppDto {
	return &NullableTargetAppDto{value: val, isSet: true}
}

func (v NullableTargetAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


