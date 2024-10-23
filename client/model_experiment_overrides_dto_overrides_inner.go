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

// checks if the ExperimentOverridesDtoOverridesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperimentOverridesDtoOverridesInner{}

// ExperimentOverridesDtoOverridesInner struct for ExperimentOverridesDtoOverridesInner
type ExperimentOverridesDtoOverridesInner struct {
	Type string `json:"type"`
	// The id of the segment or gate
	Id string `json:"id"`
	// The experiment group which user will be forced into
	GroupID string `json:"groupID"`
}

type _ExperimentOverridesDtoOverridesInner ExperimentOverridesDtoOverridesInner

// NewExperimentOverridesDtoOverridesInner instantiates a new ExperimentOverridesDtoOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentOverridesDtoOverridesInner(type_ string, id string, groupID string) *ExperimentOverridesDtoOverridesInner {
	this := ExperimentOverridesDtoOverridesInner{}
	this.Type = type_
	this.Id = id
	this.GroupID = groupID
	return &this
}

// NewExperimentOverridesDtoOverridesInnerWithDefaults instantiates a new ExperimentOverridesDtoOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentOverridesDtoOverridesInnerWithDefaults() *ExperimentOverridesDtoOverridesInner {
	this := ExperimentOverridesDtoOverridesInner{}
	return &this
}

// GetType returns the Type field value
func (o *ExperimentOverridesDtoOverridesInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExperimentOverridesDtoOverridesInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExperimentOverridesDtoOverridesInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExperimentOverridesDtoOverridesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentOverridesDtoOverridesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExperimentOverridesDtoOverridesInner) SetId(v string) {
	o.Id = v
}

// GetGroupID returns the GroupID field value
func (o *ExperimentOverridesDtoOverridesInner) GetGroupID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupID
}

// GetGroupIDOk returns a tuple with the GroupID field value
// and a boolean to check if the value has been set.
func (o *ExperimentOverridesDtoOverridesInner) GetGroupIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupID, true
}

// SetGroupID sets field value
func (o *ExperimentOverridesDtoOverridesInner) SetGroupID(v string) {
	o.GroupID = v
}

func (o ExperimentOverridesDtoOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentOverridesDtoOverridesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["groupID"] = o.GroupID
	return toSerialize, nil
}

func (o *ExperimentOverridesDtoOverridesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"groupID",
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

	varExperimentOverridesDtoOverridesInner := _ExperimentOverridesDtoOverridesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExperimentOverridesDtoOverridesInner)

	if err != nil {
		return err
	}

	*o = ExperimentOverridesDtoOverridesInner(varExperimentOverridesDtoOverridesInner)

	return err
}

type NullableExperimentOverridesDtoOverridesInner struct {
	value *ExperimentOverridesDtoOverridesInner
	isSet bool
}

func (v NullableExperimentOverridesDtoOverridesInner) Get() *ExperimentOverridesDtoOverridesInner {
	return v.value
}

func (v *NullableExperimentOverridesDtoOverridesInner) Set(val *ExperimentOverridesDtoOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentOverridesDtoOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentOverridesDtoOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentOverridesDtoOverridesInner(val *ExperimentOverridesDtoOverridesInner) *NullableExperimentOverridesDtoOverridesInner {
	return &NullableExperimentOverridesDtoOverridesInner{value: val, isSet: true}
}

func (v NullableExperimentOverridesDtoOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentOverridesDtoOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

