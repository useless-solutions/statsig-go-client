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

// checks if the RuleUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleUpdateDto{}

// RuleUpdateDto struct for RuleUpdateDto
type RuleUpdateDto struct {
	// The name of this rule.
	Name *string `json:"name,omitempty"`
	// Of the users that meet the conditions of this rule, what percent should return true.
	PassPercentage *float32 `json:"passPercentage,omitempty"`
	// An array of Condition objects.
	Conditions []DynamicConfigDtoRulesInnerConditionsInner `json:"conditions,omitempty"`
	// The environments this rule is enabled for.
	Environments []string `json:"environments,omitempty"`
	// The Statsig ID of this rule.
	Id *string `json:"id,omitempty"`
	// The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout
	BaseID *string `json:"baseID,omitempty"`
	// The return value of the rule.
	ReturnValue map[string]interface{} `json:"returnValue,omitempty"`
}

// NewRuleUpdateDto instantiates a new RuleUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleUpdateDto() *RuleUpdateDto {
	this := RuleUpdateDto{}
	return &this
}

// NewRuleUpdateDtoWithDefaults instantiates a new RuleUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleUpdateDtoWithDefaults() *RuleUpdateDto {
	this := RuleUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleUpdateDto) SetName(v string) {
	o.Name = &v
}

// GetPassPercentage returns the PassPercentage field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetPassPercentage() float32 {
	if o == nil || IsNil(o.PassPercentage) {
		var ret float32
		return ret
	}
	return *o.PassPercentage
}

// GetPassPercentageOk returns a tuple with the PassPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetPassPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.PassPercentage) {
		return nil, false
	}
	return o.PassPercentage, true
}

// HasPassPercentage returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasPassPercentage() bool {
	if o != nil && !IsNil(o.PassPercentage) {
		return true
	}

	return false
}

// SetPassPercentage gets a reference to the given float32 and assigns it to the PassPercentage field.
func (o *RuleUpdateDto) SetPassPercentage(v float32) {
	o.PassPercentage = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetConditions() []DynamicConfigDtoRulesInnerConditionsInner {
	if o == nil || IsNil(o.Conditions) {
		var ret []DynamicConfigDtoRulesInnerConditionsInner
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetConditionsOk() ([]DynamicConfigDtoRulesInnerConditionsInner, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []DynamicConfigDtoRulesInnerConditionsInner and assigns it to the Conditions field.
func (o *RuleUpdateDto) SetConditions(v []DynamicConfigDtoRulesInnerConditionsInner) {
	o.Conditions = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetEnvironments() []string {
	if o == nil || IsNil(o.Environments) {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *RuleUpdateDto) SetEnvironments(v []string) {
	o.Environments = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleUpdateDto) SetId(v string) {
	o.Id = &v
}

// GetBaseID returns the BaseID field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetBaseID() string {
	if o == nil || IsNil(o.BaseID) {
		var ret string
		return ret
	}
	return *o.BaseID
}

// GetBaseIDOk returns a tuple with the BaseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetBaseIDOk() (*string, bool) {
	if o == nil || IsNil(o.BaseID) {
		return nil, false
	}
	return o.BaseID, true
}

// HasBaseID returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasBaseID() bool {
	if o != nil && !IsNil(o.BaseID) {
		return true
	}

	return false
}

// SetBaseID gets a reference to the given string and assigns it to the BaseID field.
func (o *RuleUpdateDto) SetBaseID(v string) {
	o.BaseID = &v
}

// GetReturnValue returns the ReturnValue field value if set, zero value otherwise.
func (o *RuleUpdateDto) GetReturnValue() map[string]interface{} {
	if o == nil || IsNil(o.ReturnValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReturnValue
}

// GetReturnValueOk returns a tuple with the ReturnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdateDto) GetReturnValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReturnValue) {
		return map[string]interface{}{}, false
	}
	return o.ReturnValue, true
}

// HasReturnValue returns a boolean if a field has been set.
func (o *RuleUpdateDto) HasReturnValue() bool {
	if o != nil && !IsNil(o.ReturnValue) {
		return true
	}

	return false
}

// SetReturnValue gets a reference to the given map[string]interface{} and assigns it to the ReturnValue field.
func (o *RuleUpdateDto) SetReturnValue(v map[string]interface{}) {
	o.ReturnValue = v
}

func (o RuleUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PassPercentage) {
		toSerialize["passPercentage"] = o.PassPercentage
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
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

type NullableRuleUpdateDto struct {
	value *RuleUpdateDto
	isSet bool
}

func (v NullableRuleUpdateDto) Get() *RuleUpdateDto {
	return v.value
}

func (v *NullableRuleUpdateDto) Set(val *RuleUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleUpdateDto(val *RuleUpdateDto) *NullableRuleUpdateDto {
	return &NullableRuleUpdateDto{value: val, isSet: true}
}

func (v NullableRuleUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


