/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateTargetAppDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTargetAppDto{}

// UpdateTargetAppDto struct for UpdateTargetAppDto
type UpdateTargetAppDto struct {
	// name of the target app
	Name *string `json:"name,omitempty"`
	// a description of the target app
	Description *string `json:"description,omitempty"`
}

// NewUpdateTargetAppDto instantiates a new UpdateTargetAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTargetAppDto() *UpdateTargetAppDto {
	this := UpdateTargetAppDto{}
	return &this
}

// NewUpdateTargetAppDtoWithDefaults instantiates a new UpdateTargetAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTargetAppDtoWithDefaults() *UpdateTargetAppDto {
	this := UpdateTargetAppDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTargetAppDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetAppDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTargetAppDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTargetAppDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateTargetAppDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateTargetAppDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateTargetAppDto) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateTargetAppDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTargetAppDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableUpdateTargetAppDto struct {
	value *UpdateTargetAppDto
	isSet bool
}

func (v NullableUpdateTargetAppDto) Get() *UpdateTargetAppDto {
	return v.value
}

func (v *NullableUpdateTargetAppDto) Set(val *UpdateTargetAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTargetAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTargetAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTargetAppDto(val *UpdateTargetAppDto) *NullableUpdateTargetAppDto {
	return &NullableUpdateTargetAppDto{value: val, isSet: true}
}

func (v NullableUpdateTargetAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTargetAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

