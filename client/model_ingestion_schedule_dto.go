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

// checks if the IngestionScheduleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionScheduleDto{}

// IngestionScheduleDto struct for IngestionScheduleDto
type IngestionScheduleDto struct {
	Dataset string `json:"dataset"`
	ScheduledHourPst float32 `json:"scheduled_hour_pst"`
}

type _IngestionScheduleDto IngestionScheduleDto

// NewIngestionScheduleDto instantiates a new IngestionScheduleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionScheduleDto(dataset string, scheduledHourPst float32) *IngestionScheduleDto {
	this := IngestionScheduleDto{}
	this.Dataset = dataset
	this.ScheduledHourPst = scheduledHourPst
	return &this
}

// NewIngestionScheduleDtoWithDefaults instantiates a new IngestionScheduleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionScheduleDtoWithDefaults() *IngestionScheduleDto {
	this := IngestionScheduleDto{}
	return &this
}

// GetDataset returns the Dataset field value
func (o *IngestionScheduleDto) GetDataset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *IngestionScheduleDto) GetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value
func (o *IngestionScheduleDto) SetDataset(v string) {
	o.Dataset = v
}

// GetScheduledHourPst returns the ScheduledHourPst field value
func (o *IngestionScheduleDto) GetScheduledHourPst() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ScheduledHourPst
}

// GetScheduledHourPstOk returns a tuple with the ScheduledHourPst field value
// and a boolean to check if the value has been set.
func (o *IngestionScheduleDto) GetScheduledHourPstOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledHourPst, true
}

// SetScheduledHourPst sets field value
func (o *IngestionScheduleDto) SetScheduledHourPst(v float32) {
	o.ScheduledHourPst = v
}

func (o IngestionScheduleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestionScheduleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataset"] = o.Dataset
	toSerialize["scheduled_hour_pst"] = o.ScheduledHourPst
	return toSerialize, nil
}

func (o *IngestionScheduleDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataset",
		"scheduled_hour_pst",
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

	varIngestionScheduleDto := _IngestionScheduleDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestionScheduleDto)

	if err != nil {
		return err
	}

	*o = IngestionScheduleDto(varIngestionScheduleDto)

	return err
}

type NullableIngestionScheduleDto struct {
	value *IngestionScheduleDto
	isSet bool
}

func (v NullableIngestionScheduleDto) Get() *IngestionScheduleDto {
	return v.value
}

func (v *NullableIngestionScheduleDto) Set(val *IngestionScheduleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionScheduleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionScheduleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionScheduleDto(val *IngestionScheduleDto) *NullableIngestionScheduleDto {
	return &NullableIngestionScheduleDto{value: val, isSet: true}
}

func (v NullableIngestionScheduleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionScheduleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


