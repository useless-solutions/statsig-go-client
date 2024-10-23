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

// checks if the ConsoleV1MetricsControllerGenCreateMetricSource201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleV1MetricsControllerGenCreateMetricSource201Response{}

// ConsoleV1MetricsControllerGenCreateMetricSource201Response struct for ConsoleV1MetricsControllerGenCreateMetricSource201Response
type ConsoleV1MetricsControllerGenCreateMetricSource201Response struct {
	// A simple string explaining the result of the operation.
	Message string `json:"message"`
	Data MetricSourceContractDto `json:"data"`
}

type _ConsoleV1MetricsControllerGenCreateMetricSource201Response ConsoleV1MetricsControllerGenCreateMetricSource201Response

// NewConsoleV1MetricsControllerGenCreateMetricSource201Response instantiates a new ConsoleV1MetricsControllerGenCreateMetricSource201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleV1MetricsControllerGenCreateMetricSource201Response(message string, data MetricSourceContractDto) *ConsoleV1MetricsControllerGenCreateMetricSource201Response {
	this := ConsoleV1MetricsControllerGenCreateMetricSource201Response{}
	this.Message = message
	this.Data = data
	return &this
}

// NewConsoleV1MetricsControllerGenCreateMetricSource201ResponseWithDefaults instantiates a new ConsoleV1MetricsControllerGenCreateMetricSource201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleV1MetricsControllerGenCreateMetricSource201ResponseWithDefaults() *ConsoleV1MetricsControllerGenCreateMetricSource201Response {
	this := ConsoleV1MetricsControllerGenCreateMetricSource201Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) GetData() MetricSourceContractDto {
	if o == nil {
		var ret MetricSourceContractDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) GetDataOk() (*MetricSourceContractDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) SetData(v MetricSourceContractDto) {
	o.Data = v
}

func (o ConsoleV1MetricsControllerGenCreateMetricSource201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleV1MetricsControllerGenCreateMetricSource201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ConsoleV1MetricsControllerGenCreateMetricSource201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"data",
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

	varConsoleV1MetricsControllerGenCreateMetricSource201Response := _ConsoleV1MetricsControllerGenCreateMetricSource201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConsoleV1MetricsControllerGenCreateMetricSource201Response)

	if err != nil {
		return err
	}

	*o = ConsoleV1MetricsControllerGenCreateMetricSource201Response(varConsoleV1MetricsControllerGenCreateMetricSource201Response)

	return err
}

type NullableConsoleV1MetricsControllerGenCreateMetricSource201Response struct {
	value *ConsoleV1MetricsControllerGenCreateMetricSource201Response
	isSet bool
}

func (v NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) Get() *ConsoleV1MetricsControllerGenCreateMetricSource201Response {
	return v.value
}

func (v *NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) Set(val *ConsoleV1MetricsControllerGenCreateMetricSource201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleV1MetricsControllerGenCreateMetricSource201Response(val *ConsoleV1MetricsControllerGenCreateMetricSource201Response) *NullableConsoleV1MetricsControllerGenCreateMetricSource201Response {
	return &NullableConsoleV1MetricsControllerGenCreateMetricSource201Response{value: val, isSet: true}
}

func (v NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleV1MetricsControllerGenCreateMetricSource201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


