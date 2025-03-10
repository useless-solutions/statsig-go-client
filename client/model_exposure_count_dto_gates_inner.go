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

// checks if the ExposureCountDtoGatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExposureCountDtoGatesInner{}

// ExposureCountDtoGatesInner struct for ExposureCountDtoGatesInner
type ExposureCountDtoGatesInner struct {
	Id string `json:"id"`
	PastDay float32 `json:"pastDay"`
	Past7Days float32 `json:"past7Days"`
}

type _ExposureCountDtoGatesInner ExposureCountDtoGatesInner

// NewExposureCountDtoGatesInner instantiates a new ExposureCountDtoGatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExposureCountDtoGatesInner(id string, pastDay float32, past7Days float32) *ExposureCountDtoGatesInner {
	this := ExposureCountDtoGatesInner{}
	this.Id = id
	this.PastDay = pastDay
	this.Past7Days = past7Days
	return &this
}

// NewExposureCountDtoGatesInnerWithDefaults instantiates a new ExposureCountDtoGatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExposureCountDtoGatesInnerWithDefaults() *ExposureCountDtoGatesInner {
	this := ExposureCountDtoGatesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ExposureCountDtoGatesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExposureCountDtoGatesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExposureCountDtoGatesInner) SetId(v string) {
	o.Id = v
}

// GetPastDay returns the PastDay field value
func (o *ExposureCountDtoGatesInner) GetPastDay() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PastDay
}

// GetPastDayOk returns a tuple with the PastDay field value
// and a boolean to check if the value has been set.
func (o *ExposureCountDtoGatesInner) GetPastDayOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PastDay, true
}

// SetPastDay sets field value
func (o *ExposureCountDtoGatesInner) SetPastDay(v float32) {
	o.PastDay = v
}

// GetPast7Days returns the Past7Days field value
func (o *ExposureCountDtoGatesInner) GetPast7Days() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Past7Days
}

// GetPast7DaysOk returns a tuple with the Past7Days field value
// and a boolean to check if the value has been set.
func (o *ExposureCountDtoGatesInner) GetPast7DaysOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Past7Days, true
}

// SetPast7Days sets field value
func (o *ExposureCountDtoGatesInner) SetPast7Days(v float32) {
	o.Past7Days = v
}

func (o ExposureCountDtoGatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExposureCountDtoGatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["pastDay"] = o.PastDay
	toSerialize["past7Days"] = o.Past7Days
	return toSerialize, nil
}

func (o *ExposureCountDtoGatesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"pastDay",
		"past7Days",
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

	varExposureCountDtoGatesInner := _ExposureCountDtoGatesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExposureCountDtoGatesInner)

	if err != nil {
		return err
	}

	*o = ExposureCountDtoGatesInner(varExposureCountDtoGatesInner)

	return err
}

type NullableExposureCountDtoGatesInner struct {
	value *ExposureCountDtoGatesInner
	isSet bool
}

func (v NullableExposureCountDtoGatesInner) Get() *ExposureCountDtoGatesInner {
	return v.value
}

func (v *NullableExposureCountDtoGatesInner) Set(val *ExposureCountDtoGatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExposureCountDtoGatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExposureCountDtoGatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExposureCountDtoGatesInner(val *ExposureCountDtoGatesInner) *NullableExposureCountDtoGatesInner {
	return &NullableExposureCountDtoGatesInner{value: val, isSet: true}
}

func (v NullableExposureCountDtoGatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExposureCountDtoGatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


