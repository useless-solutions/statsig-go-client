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

// checks if the ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response{}

// ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response struct for ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response
type ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response instantiates a new ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response() *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response {
	this := ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response{}
	return &this
}

// NewConsoleV1ExperimentsControllerGenRemoveAssignmentSource200ResponseWithDefaults instantiates a new ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleV1ExperimentsControllerGenRemoveAssignmentSource200ResponseWithDefaults() *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response {
	this := ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) SetMessage(v string) {
	o.Message = &v
}

func (o ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response struct {
	value *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response
	isSet bool
}

func (v NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) Get() *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response {
	return v.value
}

func (v *NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) Set(val *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response(val *ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) *NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response {
	return &NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response{value: val, isSet: true}
}

func (v NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


