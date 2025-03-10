/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MultipleEventCountDtoInnerEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipleEventCountDtoInnerEventsInner{}

// MultipleEventCountDtoInnerEventsInner struct for MultipleEventCountDtoInnerEventsInner
type MultipleEventCountDtoInnerEventsInner struct {
	Event string `json:"event"`
	Count float32 `json:"count"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type _MultipleEventCountDtoInnerEventsInner MultipleEventCountDtoInnerEventsInner

// NewMultipleEventCountDtoInnerEventsInner instantiates a new MultipleEventCountDtoInnerEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleEventCountDtoInnerEventsInner(event string, count float32, lastUpdatedTime time.Time) *MultipleEventCountDtoInnerEventsInner {
	this := MultipleEventCountDtoInnerEventsInner{}
	this.Event = event
	this.Count = count
	this.LastUpdatedTime = lastUpdatedTime
	return &this
}

// NewMultipleEventCountDtoInnerEventsInnerWithDefaults instantiates a new MultipleEventCountDtoInnerEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleEventCountDtoInnerEventsInnerWithDefaults() *MultipleEventCountDtoInnerEventsInner {
	this := MultipleEventCountDtoInnerEventsInner{}
	return &this
}

// GetEvent returns the Event field value
func (o *MultipleEventCountDtoInnerEventsInner) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *MultipleEventCountDtoInnerEventsInner) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *MultipleEventCountDtoInnerEventsInner) SetEvent(v string) {
	o.Event = v
}

// GetCount returns the Count field value
func (o *MultipleEventCountDtoInnerEventsInner) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MultipleEventCountDtoInnerEventsInner) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *MultipleEventCountDtoInnerEventsInner) SetCount(v float32) {
	o.Count = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *MultipleEventCountDtoInnerEventsInner) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *MultipleEventCountDtoInnerEventsInner) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *MultipleEventCountDtoInnerEventsInner) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

func (o MultipleEventCountDtoInnerEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipleEventCountDtoInnerEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["count"] = o.Count
	toSerialize["last_updated_time"] = o.LastUpdatedTime
	return toSerialize, nil
}

func (o *MultipleEventCountDtoInnerEventsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"count",
		"last_updated_time",
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

	varMultipleEventCountDtoInnerEventsInner := _MultipleEventCountDtoInnerEventsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultipleEventCountDtoInnerEventsInner)

	if err != nil {
		return err
	}

	*o = MultipleEventCountDtoInnerEventsInner(varMultipleEventCountDtoInnerEventsInner)

	return err
}

type NullableMultipleEventCountDtoInnerEventsInner struct {
	value *MultipleEventCountDtoInnerEventsInner
	isSet bool
}

func (v NullableMultipleEventCountDtoInnerEventsInner) Get() *MultipleEventCountDtoInnerEventsInner {
	return v.value
}

func (v *NullableMultipleEventCountDtoInnerEventsInner) Set(val *MultipleEventCountDtoInnerEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleEventCountDtoInnerEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleEventCountDtoInnerEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleEventCountDtoInnerEventsInner(val *MultipleEventCountDtoInnerEventsInner) *NullableMultipleEventCountDtoInnerEventsInner {
	return &NullableMultipleEventCountDtoInnerEventsInner{value: val, isSet: true}
}

func (v NullableMultipleEventCountDtoInnerEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleEventCountDtoInnerEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


