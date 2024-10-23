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

// checks if the ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner{}

// ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner struct for ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner
type ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner struct {
	// Optional array of criteria to filter the funnel events, defined by various types and conditions.
	Criteria []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner `json:"criteria,omitempty"`
	// Optional name of the metric source associated with the funnel event.
	MetricSourceName *string `json:"metricSourceName,omitempty"`
	// Optional step name for the funnel event, can be null if not specified.
	Name *nil `json:"name,omitempty"`
	// Name of column which being used as session identifier. Funnel event with the same metric source
	SessionIdentifierField *nil `json:"sessionIdentifierField,omitempty"`
}

// NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner instantiates a new ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner() *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner {
	this := ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner{}
	return &this
}

// NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInnerWithDefaults() *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner {
	this := ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner {
	if o == nil || IsNil(o.Criteria) {
		var ret []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetCriteriaOk() ([]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner and assigns it to the Criteria field.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner) {
	o.Criteria = v
}

// GetMetricSourceName returns the MetricSourceName field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetMetricSourceName() string {
	if o == nil || IsNil(o.MetricSourceName) {
		var ret string
		return ret
	}
	return *o.MetricSourceName
}

// GetMetricSourceNameOk returns a tuple with the MetricSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetMetricSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricSourceName) {
		return nil, false
	}
	return o.MetricSourceName, true
}

// HasMetricSourceName returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasMetricSourceName() bool {
	if o != nil && !IsNil(o.MetricSourceName) {
		return true
	}

	return false
}

// SetMetricSourceName gets a reference to the given string and assigns it to the MetricSourceName field.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetMetricSourceName(v string) {
	o.MetricSourceName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetName() nil {
	if o == nil || IsNil(o.Name) {
		var ret nil
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetNameOk() (*nil, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given nil and assigns it to the Name field.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetName(v nil) {
	o.Name = &v
}

// GetSessionIdentifierField returns the SessionIdentifierField field value if set, zero value otherwise.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetSessionIdentifierField() nil {
	if o == nil || IsNil(o.SessionIdentifierField) {
		var ret nil
		return ret
	}
	return *o.SessionIdentifierField
}

// GetSessionIdentifierFieldOk returns a tuple with the SessionIdentifierField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetSessionIdentifierFieldOk() (*nil, bool) {
	if o == nil || IsNil(o.SessionIdentifierField) {
		return nil, false
	}
	return o.SessionIdentifierField, true
}

// HasSessionIdentifierField returns a boolean if a field has been set.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasSessionIdentifierField() bool {
	if o != nil && !IsNil(o.SessionIdentifierField) {
		return true
	}

	return false
}

// SetSessionIdentifierField gets a reference to the given nil and assigns it to the SessionIdentifierField field.
func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetSessionIdentifierField(v nil) {
	o.SessionIdentifierField = &v
}

func (o ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.MetricSourceName) {
		toSerialize["metricSourceName"] = o.MetricSourceName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SessionIdentifierField) {
		toSerialize["sessionIdentifierField"] = o.SessionIdentifierField
	}
	return toSerialize, nil
}

type NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner struct {
	value *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner
	isSet bool
}

func (v NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) Get() *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner {
	return v.value
}

func (v *NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) Set(val *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner(val *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) *NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner {
	return &NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner{value: val, isSet: true}
}

func (v NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


