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

// checks if the ExternalMetricDefinitionContractDtoMetricEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalMetricDefinitionContractDtoMetricEventsInner{}

// ExternalMetricDefinitionContractDtoMetricEventsInner struct for ExternalMetricDefinitionContractDtoMetricEventsInner
type ExternalMetricDefinitionContractDtoMetricEventsInner struct {
	// The name of the metric event.
	Name string `json:"name"`
	// The type of metric event. Allowed values include: count, count_distinct, value, and metadata.
	Type *nil `json:"type,omitempty"`
	// The key for associated metadata, if applicable.
	MetadataKey *string `json:"metadataKey,omitempty"`
	// Filtering criteria for the metric event, including conditions and values to refine the event data.
	Criteria []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner `json:"criteria,omitempty"`
}

type _ExternalMetricDefinitionContractDtoMetricEventsInner ExternalMetricDefinitionContractDtoMetricEventsInner

// NewExternalMetricDefinitionContractDtoMetricEventsInner instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMetricDefinitionContractDtoMetricEventsInner(name string) *ExternalMetricDefinitionContractDtoMetricEventsInner {
	this := ExternalMetricDefinitionContractDtoMetricEventsInner{}
	this.Name = name
	return &this
}

// NewExternalMetricDefinitionContractDtoMetricEventsInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMetricDefinitionContractDtoMetricEventsInnerWithDefaults() *ExternalMetricDefinitionContractDtoMetricEventsInner {
	this := ExternalMetricDefinitionContractDtoMetricEventsInner{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetType() nil {
	if o == nil || IsNil(o.Type) {
		var ret nil
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetTypeOk() (*nil, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given nil and assigns it to the Type field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetType(v nil) {
	o.Type = &v
}

// GetMetadataKey returns the MetadataKey field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetMetadataKey() string {
	if o == nil || IsNil(o.MetadataKey) {
		var ret string
		return ret
	}
	return *o.MetadataKey
}

// GetMetadataKeyOk returns a tuple with the MetadataKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetMetadataKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataKey) {
		return nil, false
	}
	return o.MetadataKey, true
}

// HasMetadataKey returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasMetadataKey() bool {
	if o != nil && !IsNil(o.MetadataKey) {
		return true
	}

	return false
}

// SetMetadataKey gets a reference to the given string and assigns it to the MetadataKey field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetMetadataKey(v string) {
	o.MetadataKey = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	if o == nil || IsNil(o.Criteria) {
		var ret []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetCriteriaOk() ([]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner and assigns it to the Criteria field.
func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) {
	o.Criteria = v
}

func (o ExternalMetricDefinitionContractDtoMetricEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalMetricDefinitionContractDtoMetricEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MetadataKey) {
		toSerialize["metadataKey"] = o.MetadataKey
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	return toSerialize, nil
}

func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) UnmarshalJSON(data []byte) (err error) {
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

	varExternalMetricDefinitionContractDtoMetricEventsInner := _ExternalMetricDefinitionContractDtoMetricEventsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalMetricDefinitionContractDtoMetricEventsInner)

	if err != nil {
		return err
	}

	*o = ExternalMetricDefinitionContractDtoMetricEventsInner(varExternalMetricDefinitionContractDtoMetricEventsInner)

	return err
}

type NullableExternalMetricDefinitionContractDtoMetricEventsInner struct {
	value *ExternalMetricDefinitionContractDtoMetricEventsInner
	isSet bool
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInner) Get() *ExternalMetricDefinitionContractDtoMetricEventsInner {
	return v.value
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInner) Set(val *ExternalMetricDefinitionContractDtoMetricEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMetricDefinitionContractDtoMetricEventsInner(val *ExternalMetricDefinitionContractDtoMetricEventsInner) *NullableExternalMetricDefinitionContractDtoMetricEventsInner {
	return &NullableExternalMetricDefinitionContractDtoMetricEventsInner{value: val, isSet: true}
}

func (v NullableExternalMetricDefinitionContractDtoMetricEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMetricDefinitionContractDtoMetricEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


