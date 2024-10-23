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

// checks if the RuleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleDto{}

// RuleDto struct for RuleDto
type RuleDto struct {
	// The name of this rule.
	Name string `json:"name"`
	// Of the users that meet the conditions of this rule, what percent should return true.
	PassPercentage float32 `json:"passPercentage"`
	// An array of Condition objects.
	Conditions []DynamicConfigDtoRulesInnerConditionsInner `json:"conditions"`
	// The environments this rule is enabled for.
	Environments []string `json:"environments,omitempty"`
	// The Statsig ID of this rule.
	Id *string `json:"id,omitempty"`
	// The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout
	BaseID *string `json:"baseID,omitempty"`
	// The return value of the rule.
	ReturnValue map[string]interface{} `json:"returnValue,omitempty"`
}

type _RuleDto RuleDto

// NewRuleDto instantiates a new RuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDto(name string, passPercentage float32, conditions []DynamicConfigDtoRulesInnerConditionsInner) *RuleDto {
	this := RuleDto{}
	this.Name = name
	this.PassPercentage = passPercentage
	this.Conditions = conditions
	return &this
}

// NewRuleDtoWithDefaults instantiates a new RuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDtoWithDefaults() *RuleDto {
	this := RuleDto{}
	return &this
}

// GetName returns the Name field value
func (o *RuleDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleDto) SetName(v string) {
	o.Name = v
}

// GetPassPercentage returns the PassPercentage field value
func (o *RuleDto) GetPassPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PassPercentage
}

// GetPassPercentageOk returns a tuple with the PassPercentage field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetPassPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassPercentage, true
}

// SetPassPercentage sets field value
func (o *RuleDto) SetPassPercentage(v float32) {
	o.PassPercentage = v
}

// GetConditions returns the Conditions field value
func (o *RuleDto) GetConditions() []DynamicConfigDtoRulesInnerConditionsInner {
	if o == nil {
		var ret []DynamicConfigDtoRulesInnerConditionsInner
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *RuleDto) GetConditionsOk() ([]DynamicConfigDtoRulesInnerConditionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *RuleDto) SetConditions(v []DynamicConfigDtoRulesInnerConditionsInner) {
	o.Conditions = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *RuleDto) GetEnvironments() []string {
	if o == nil || IsNil(o.Environments) {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *RuleDto) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *RuleDto) SetEnvironments(v []string) {
	o.Environments = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleDto) SetId(v string) {
	o.Id = &v
}

// GetBaseID returns the BaseID field value if set, zero value otherwise.
func (o *RuleDto) GetBaseID() string {
	if o == nil || IsNil(o.BaseID) {
		var ret string
		return ret
	}
	return *o.BaseID
}

// GetBaseIDOk returns a tuple with the BaseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetBaseIDOk() (*string, bool) {
	if o == nil || IsNil(o.BaseID) {
		return nil, false
	}
	return o.BaseID, true
}

// HasBaseID returns a boolean if a field has been set.
func (o *RuleDto) HasBaseID() bool {
	if o != nil && !IsNil(o.BaseID) {
		return true
	}

	return false
}

// SetBaseID gets a reference to the given string and assigns it to the BaseID field.
func (o *RuleDto) SetBaseID(v string) {
	o.BaseID = &v
}

// GetReturnValue returns the ReturnValue field value if set, zero value otherwise.
func (o *RuleDto) GetReturnValue() map[string]interface{} {
	if o == nil || IsNil(o.ReturnValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReturnValue
}

// GetReturnValueOk returns a tuple with the ReturnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDto) GetReturnValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReturnValue) {
		return map[string]interface{}{}, false
	}
	return o.ReturnValue, true
}

// HasReturnValue returns a boolean if a field has been set.
func (o *RuleDto) HasReturnValue() bool {
	if o != nil && !IsNil(o.ReturnValue) {
		return true
	}

	return false
}

// SetReturnValue gets a reference to the given map[string]interface{} and assigns it to the ReturnValue field.
func (o *RuleDto) SetReturnValue(v map[string]interface{}) {
	o.ReturnValue = v
}

func (o RuleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["passPercentage"] = o.PassPercentage
	toSerialize["conditions"] = o.Conditions
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.BaseID) {
		toSerialize["baseID"] = o.BaseID
	}
	if !IsNil(o.ReturnValue) {
		toSerialize["returnValue"] = o.ReturnValue
	}
	return toSerialize, nil
}

func (o *RuleDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"passPercentage",
		"conditions",
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

	varRuleDto := _RuleDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleDto)

	if err != nil {
		return err
	}

	*o = RuleDto(varRuleDto)

	return err
}

type NullableRuleDto struct {
	value *RuleDto
	isSet bool
}

func (v NullableRuleDto) Get() *RuleDto {
	return v.value
}

func (v *NullableRuleDto) Set(val *RuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDto(val *RuleDto) *NullableRuleDto {
	return &NullableRuleDto{value: val, isSet: true}
}

func (v NullableRuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

