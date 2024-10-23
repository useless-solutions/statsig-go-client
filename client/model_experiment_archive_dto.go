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

// checks if the ExperimentArchiveDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperimentArchiveDto{}

// ExperimentArchiveDto Schema for archiving an experiment
type ExperimentArchiveDto struct {
	// The reason for archiving the experiment
	ArchiveReason *string `json:"archiveReason,omitempty"`
}

// NewExperimentArchiveDto instantiates a new ExperimentArchiveDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentArchiveDto() *ExperimentArchiveDto {
	this := ExperimentArchiveDto{}
	return &this
}

// NewExperimentArchiveDtoWithDefaults instantiates a new ExperimentArchiveDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentArchiveDtoWithDefaults() *ExperimentArchiveDto {
	this := ExperimentArchiveDto{}
	return &this
}

// GetArchiveReason returns the ArchiveReason field value if set, zero value otherwise.
func (o *ExperimentArchiveDto) GetArchiveReason() string {
	if o == nil || IsNil(o.ArchiveReason) {
		var ret string
		return ret
	}
	return *o.ArchiveReason
}

// GetArchiveReasonOk returns a tuple with the ArchiveReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentArchiveDto) GetArchiveReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveReason) {
		return nil, false
	}
	return o.ArchiveReason, true
}

// HasArchiveReason returns a boolean if a field has been set.
func (o *ExperimentArchiveDto) HasArchiveReason() bool {
	if o != nil && !IsNil(o.ArchiveReason) {
		return true
	}

	return false
}

// SetArchiveReason gets a reference to the given string and assigns it to the ArchiveReason field.
func (o *ExperimentArchiveDto) SetArchiveReason(v string) {
	o.ArchiveReason = &v
}

func (o ExperimentArchiveDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentArchiveDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArchiveReason) {
		toSerialize["archiveReason"] = o.ArchiveReason
	}
	return toSerialize, nil
}

type NullableExperimentArchiveDto struct {
	value *ExperimentArchiveDto
	isSet bool
}

func (v NullableExperimentArchiveDto) Get() *ExperimentArchiveDto {
	return v.value
}

func (v *NullableExperimentArchiveDto) Set(val *ExperimentArchiveDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentArchiveDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentArchiveDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentArchiveDto(val *ExperimentArchiveDto) *NullableExperimentArchiveDto {
	return &NullableExperimentArchiveDto{value: val, isSet: true}
}

func (v NullableExperimentArchiveDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentArchiveDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


