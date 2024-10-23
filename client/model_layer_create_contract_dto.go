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

// checks if the LayerCreateContractDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LayerCreateContractDto{}

// LayerCreateContractDto struct for LayerCreateContractDto
type LayerCreateContractDto struct {
	// The unique name of the layer, formatted as an ID (e.g., \"A Layer\" becomes \"a_layer\").
	Name string `json:"name" validate:"regexp=^[A-Za-z0-9_ -]*$"`
	// A helpful description of what this layer does, providing context for its purpose.
	Description *string `json:"description,omitempty"`
	// The type of ID used for this layer, essential for validation in services.
	IdType string `json:"idType"`
	TargetApps *LayerCreateContractDtoTargetApps `json:"targetApps,omitempty"`
	// Optional identifier for the team responsible for this layer.
	Team *string `json:"team,omitempty"`
}

type _LayerCreateContractDto LayerCreateContractDto

// NewLayerCreateContractDto instantiates a new LayerCreateContractDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayerCreateContractDto(name string, idType string) *LayerCreateContractDto {
	this := LayerCreateContractDto{}
	this.Name = name
	this.IdType = idType
	return &this
}

// NewLayerCreateContractDtoWithDefaults instantiates a new LayerCreateContractDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayerCreateContractDtoWithDefaults() *LayerCreateContractDto {
	this := LayerCreateContractDto{}
	return &this
}

// GetName returns the Name field value
func (o *LayerCreateContractDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LayerCreateContractDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LayerCreateContractDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LayerCreateContractDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerCreateContractDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LayerCreateContractDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LayerCreateContractDto) SetDescription(v string) {
	o.Description = &v
}

// GetIdType returns the IdType field value
func (o *LayerCreateContractDto) GetIdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *LayerCreateContractDto) GetIdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *LayerCreateContractDto) SetIdType(v string) {
	o.IdType = v
}

// GetTargetApps returns the TargetApps field value if set, zero value otherwise.
func (o *LayerCreateContractDto) GetTargetApps() LayerCreateContractDtoTargetApps {
	if o == nil || IsNil(o.TargetApps) {
		var ret LayerCreateContractDtoTargetApps
		return ret
	}
	return *o.TargetApps
}

// GetTargetAppsOk returns a tuple with the TargetApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerCreateContractDto) GetTargetAppsOk() (*LayerCreateContractDtoTargetApps, bool) {
	if o == nil || IsNil(o.TargetApps) {
		return nil, false
	}
	return o.TargetApps, true
}

// HasTargetApps returns a boolean if a field has been set.
func (o *LayerCreateContractDto) HasTargetApps() bool {
	if o != nil && !IsNil(o.TargetApps) {
		return true
	}

	return false
}

// SetTargetApps gets a reference to the given LayerCreateContractDtoTargetApps and assigns it to the TargetApps field.
func (o *LayerCreateContractDto) SetTargetApps(v LayerCreateContractDtoTargetApps) {
	o.TargetApps = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *LayerCreateContractDto) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerCreateContractDto) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *LayerCreateContractDto) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *LayerCreateContractDto) SetTeam(v string) {
	o.Team = &v
}

func (o LayerCreateContractDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LayerCreateContractDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["idType"] = o.IdType
	if !IsNil(o.TargetApps) {
		toSerialize["targetApps"] = o.TargetApps
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	return toSerialize, nil
}

func (o *LayerCreateContractDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"idType",
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

	varLayerCreateContractDto := _LayerCreateContractDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLayerCreateContractDto)

	if err != nil {
		return err
	}

	*o = LayerCreateContractDto(varLayerCreateContractDto)

	return err
}

type NullableLayerCreateContractDto struct {
	value *LayerCreateContractDto
	isSet bool
}

func (v NullableLayerCreateContractDto) Get() *LayerCreateContractDto {
	return v.value
}

func (v *NullableLayerCreateContractDto) Set(val *LayerCreateContractDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLayerCreateContractDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLayerCreateContractDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayerCreateContractDto(val *LayerCreateContractDto) *NullableLayerCreateContractDto {
	return &NullableLayerCreateContractDto{value: val, isSet: true}
}

func (v NullableLayerCreateContractDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayerCreateContractDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

