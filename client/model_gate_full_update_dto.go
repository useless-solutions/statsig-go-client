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

// checks if the GateFullUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GateFullUpdateDto{}

// GateFullUpdateDto struct for GateFullUpdateDto
type GateFullUpdateDto struct {
	IsEnabled bool `json:"isEnabled"`
	Description string `json:"description"`
	Rules []ExternalGateDtoRulesInner `json:"rules"`
	Tags []string `json:"tags,omitempty"`
	Type *string `json:"type,omitempty"`
	IdType *string `json:"idType,omitempty"`
	TargetApps *DynamicConfigDtoTargetApps `json:"targetApps,omitempty"`
	CreatorID *nil `json:"creatorID,omitempty"`
	CreatorEmail *nil `json:"creatorEmail,omitempty"`
	Team *nil `json:"team,omitempty"`
	MeasureMetricLifts *bool `json:"measureMetricLifts,omitempty"`
	MonitoringMetrics []ExternalGateDtoMonitoringMetricsInner `json:"monitoringMetrics,omitempty"`
}

type _GateFullUpdateDto GateFullUpdateDto

// NewGateFullUpdateDto instantiates a new GateFullUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateFullUpdateDto(isEnabled bool, description string, rules []ExternalGateDtoRulesInner) *GateFullUpdateDto {
	this := GateFullUpdateDto{}
	this.IsEnabled = isEnabled
	this.Description = description
	this.Rules = rules
	return &this
}

// NewGateFullUpdateDtoWithDefaults instantiates a new GateFullUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGateFullUpdateDtoWithDefaults() *GateFullUpdateDto {
	this := GateFullUpdateDto{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GateFullUpdateDto) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GateFullUpdateDto) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetDescription returns the Description field value
func (o *GateFullUpdateDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GateFullUpdateDto) SetDescription(v string) {
	o.Description = v
}

// GetRules returns the Rules field value
func (o *GateFullUpdateDto) GetRules() []ExternalGateDtoRulesInner {
	if o == nil {
		var ret []ExternalGateDtoRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetRulesOk() ([]ExternalGateDtoRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *GateFullUpdateDto) SetRules(v []ExternalGateDtoRulesInner) {
	o.Rules = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GateFullUpdateDto) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GateFullUpdateDto) SetType(v string) {
	o.Type = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetIdType() string {
	if o == nil || IsNil(o.IdType) {
		var ret string
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdType) {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasIdType() bool {
	if o != nil && !IsNil(o.IdType) {
		return true
	}

	return false
}

// SetIdType gets a reference to the given string and assigns it to the IdType field.
func (o *GateFullUpdateDto) SetIdType(v string) {
	o.IdType = &v
}

// GetTargetApps returns the TargetApps field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetTargetApps() DynamicConfigDtoTargetApps {
	if o == nil || IsNil(o.TargetApps) {
		var ret DynamicConfigDtoTargetApps
		return ret
	}
	return *o.TargetApps
}

// GetTargetAppsOk returns a tuple with the TargetApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool) {
	if o == nil || IsNil(o.TargetApps) {
		return nil, false
	}
	return o.TargetApps, true
}

// HasTargetApps returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasTargetApps() bool {
	if o != nil && !IsNil(o.TargetApps) {
		return true
	}

	return false
}

// SetTargetApps gets a reference to the given DynamicConfigDtoTargetApps and assigns it to the TargetApps field.
func (o *GateFullUpdateDto) SetTargetApps(v DynamicConfigDtoTargetApps) {
	o.TargetApps = &v
}

// GetCreatorID returns the CreatorID field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetCreatorID() nil {
	if o == nil || IsNil(o.CreatorID) {
		var ret nil
		return ret
	}
	return *o.CreatorID
}

// GetCreatorIDOk returns a tuple with the CreatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetCreatorIDOk() (*nil, bool) {
	if o == nil || IsNil(o.CreatorID) {
		return nil, false
	}
	return o.CreatorID, true
}

// HasCreatorID returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasCreatorID() bool {
	if o != nil && !IsNil(o.CreatorID) {
		return true
	}

	return false
}

// SetCreatorID gets a reference to the given nil and assigns it to the CreatorID field.
func (o *GateFullUpdateDto) SetCreatorID(v nil) {
	o.CreatorID = &v
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetCreatorEmail() nil {
	if o == nil || IsNil(o.CreatorEmail) {
		var ret nil
		return ret
	}
	return *o.CreatorEmail
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetCreatorEmailOk() (*nil, bool) {
	if o == nil || IsNil(o.CreatorEmail) {
		return nil, false
	}
	return o.CreatorEmail, true
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasCreatorEmail() bool {
	if o != nil && !IsNil(o.CreatorEmail) {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given nil and assigns it to the CreatorEmail field.
func (o *GateFullUpdateDto) SetCreatorEmail(v nil) {
	o.CreatorEmail = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetTeam() nil {
	if o == nil || IsNil(o.Team) {
		var ret nil
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetTeamOk() (*nil, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given nil and assigns it to the Team field.
func (o *GateFullUpdateDto) SetTeam(v nil) {
	o.Team = &v
}

// GetMeasureMetricLifts returns the MeasureMetricLifts field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetMeasureMetricLifts() bool {
	if o == nil || IsNil(o.MeasureMetricLifts) {
		var ret bool
		return ret
	}
	return *o.MeasureMetricLifts
}

// GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetMeasureMetricLiftsOk() (*bool, bool) {
	if o == nil || IsNil(o.MeasureMetricLifts) {
		return nil, false
	}
	return o.MeasureMetricLifts, true
}

// HasMeasureMetricLifts returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasMeasureMetricLifts() bool {
	if o != nil && !IsNil(o.MeasureMetricLifts) {
		return true
	}

	return false
}

// SetMeasureMetricLifts gets a reference to the given bool and assigns it to the MeasureMetricLifts field.
func (o *GateFullUpdateDto) SetMeasureMetricLifts(v bool) {
	o.MeasureMetricLifts = &v
}

// GetMonitoringMetrics returns the MonitoringMetrics field value if set, zero value otherwise.
func (o *GateFullUpdateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner {
	if o == nil || IsNil(o.MonitoringMetrics) {
		var ret []ExternalGateDtoMonitoringMetricsInner
		return ret
	}
	return o.MonitoringMetrics
}

// GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateFullUpdateDto) GetMonitoringMetricsOk() ([]ExternalGateDtoMonitoringMetricsInner, bool) {
	if o == nil || IsNil(o.MonitoringMetrics) {
		return nil, false
	}
	return o.MonitoringMetrics, true
}

// HasMonitoringMetrics returns a boolean if a field has been set.
func (o *GateFullUpdateDto) HasMonitoringMetrics() bool {
	if o != nil && !IsNil(o.MonitoringMetrics) {
		return true
	}

	return false
}

// SetMonitoringMetrics gets a reference to the given []ExternalGateDtoMonitoringMetricsInner and assigns it to the MonitoringMetrics field.
func (o *GateFullUpdateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner) {
	o.MonitoringMetrics = v
}

func (o GateFullUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GateFullUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["description"] = o.Description
	toSerialize["rules"] = o.Rules
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IdType) {
		toSerialize["idType"] = o.IdType
	}
	if !IsNil(o.TargetApps) {
		toSerialize["targetApps"] = o.TargetApps
	}
	if !IsNil(o.CreatorID) {
		toSerialize["creatorID"] = o.CreatorID
	}
	if !IsNil(o.CreatorEmail) {
		toSerialize["creatorEmail"] = o.CreatorEmail
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.MeasureMetricLifts) {
		toSerialize["measureMetricLifts"] = o.MeasureMetricLifts
	}
	if !IsNil(o.MonitoringMetrics) {
		toSerialize["monitoringMetrics"] = o.MonitoringMetrics
	}
	return toSerialize, nil
}

func (o *GateFullUpdateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isEnabled",
		"description",
		"rules",
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

	varGateFullUpdateDto := _GateFullUpdateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGateFullUpdateDto)

	if err != nil {
		return err
	}

	*o = GateFullUpdateDto(varGateFullUpdateDto)

	return err
}

type NullableGateFullUpdateDto struct {
	value *GateFullUpdateDto
	isSet bool
}

func (v NullableGateFullUpdateDto) Get() *GateFullUpdateDto {
	return v.value
}

func (v *NullableGateFullUpdateDto) Set(val *GateFullUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGateFullUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGateFullUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateFullUpdateDto(val *GateFullUpdateDto) *NullableGateFullUpdateDto {
	return &NullableGateFullUpdateDto{value: val, isSet: true}
}

func (v NullableGateFullUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateFullUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

