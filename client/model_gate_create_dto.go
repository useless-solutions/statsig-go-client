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

// checks if the GateCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GateCreateDto{}

// GateCreateDto Create a new gate
type GateCreateDto struct {
	IsEnabled *bool `json:"isEnabled,omitempty"`
	Description *string `json:"description,omitempty"`
	Rules []ExternalGateDtoRulesInner `json:"rules,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Type *string `json:"type,omitempty"`
	IdType *string `json:"idType,omitempty"`
	TargetApps *DynamicConfigDtoTargetApps `json:"targetApps,omitempty"`
	CreatorID *nil `json:"creatorID,omitempty"`
	CreatorEmail *nil `json:"creatorEmail,omitempty"`
	Team *nil `json:"team,omitempty"`
	MeasureMetricLifts *bool `json:"measureMetricLifts,omitempty"`
	MonitoringMetrics []ExternalGateDtoMonitoringMetricsInner `json:"monitoringMetrics,omitempty"`
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9_\\\\- ]*$"`
	Id *string `json:"id,omitempty" validate:"regexp=^[a-zA-Z0-9_-]*$"`
}

// NewGateCreateDto instantiates a new GateCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateCreateDto() *GateCreateDto {
	this := GateCreateDto{}
	return &this
}

// NewGateCreateDtoWithDefaults instantiates a new GateCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGateCreateDtoWithDefaults() *GateCreateDto {
	this := GateCreateDto{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *GateCreateDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *GateCreateDto) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *GateCreateDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GateCreateDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GateCreateDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GateCreateDto) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GateCreateDto) GetRules() []ExternalGateDtoRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []ExternalGateDtoRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetRulesOk() ([]ExternalGateDtoRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GateCreateDto) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ExternalGateDtoRulesInner and assigns it to the Rules field.
func (o *GateCreateDto) SetRules(v []ExternalGateDtoRulesInner) {
	o.Rules = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GateCreateDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GateCreateDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GateCreateDto) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GateCreateDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GateCreateDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GateCreateDto) SetType(v string) {
	o.Type = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *GateCreateDto) GetIdType() string {
	if o == nil || IsNil(o.IdType) {
		var ret string
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdType) {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *GateCreateDto) HasIdType() bool {
	if o != nil && !IsNil(o.IdType) {
		return true
	}

	return false
}

// SetIdType gets a reference to the given string and assigns it to the IdType field.
func (o *GateCreateDto) SetIdType(v string) {
	o.IdType = &v
}

// GetTargetApps returns the TargetApps field value if set, zero value otherwise.
func (o *GateCreateDto) GetTargetApps() DynamicConfigDtoTargetApps {
	if o == nil || IsNil(o.TargetApps) {
		var ret DynamicConfigDtoTargetApps
		return ret
	}
	return *o.TargetApps
}

// GetTargetAppsOk returns a tuple with the TargetApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool) {
	if o == nil || IsNil(o.TargetApps) {
		return nil, false
	}
	return o.TargetApps, true
}

// HasTargetApps returns a boolean if a field has been set.
func (o *GateCreateDto) HasTargetApps() bool {
	if o != nil && !IsNil(o.TargetApps) {
		return true
	}

	return false
}

// SetTargetApps gets a reference to the given DynamicConfigDtoTargetApps and assigns it to the TargetApps field.
func (o *GateCreateDto) SetTargetApps(v DynamicConfigDtoTargetApps) {
	o.TargetApps = &v
}

// GetCreatorID returns the CreatorID field value if set, zero value otherwise.
func (o *GateCreateDto) GetCreatorID() nil {
	if o == nil || IsNil(o.CreatorID) {
		var ret nil
		return ret
	}
	return *o.CreatorID
}

// GetCreatorIDOk returns a tuple with the CreatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetCreatorIDOk() (*nil, bool) {
	if o == nil || IsNil(o.CreatorID) {
		return nil, false
	}
	return o.CreatorID, true
}

// HasCreatorID returns a boolean if a field has been set.
func (o *GateCreateDto) HasCreatorID() bool {
	if o != nil && !IsNil(o.CreatorID) {
		return true
	}

	return false
}

// SetCreatorID gets a reference to the given nil and assigns it to the CreatorID field.
func (o *GateCreateDto) SetCreatorID(v nil) {
	o.CreatorID = &v
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise.
func (o *GateCreateDto) GetCreatorEmail() nil {
	if o == nil || IsNil(o.CreatorEmail) {
		var ret nil
		return ret
	}
	return *o.CreatorEmail
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetCreatorEmailOk() (*nil, bool) {
	if o == nil || IsNil(o.CreatorEmail) {
		return nil, false
	}
	return o.CreatorEmail, true
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *GateCreateDto) HasCreatorEmail() bool {
	if o != nil && !IsNil(o.CreatorEmail) {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given nil and assigns it to the CreatorEmail field.
func (o *GateCreateDto) SetCreatorEmail(v nil) {
	o.CreatorEmail = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *GateCreateDto) GetTeam() nil {
	if o == nil || IsNil(o.Team) {
		var ret nil
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetTeamOk() (*nil, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *GateCreateDto) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given nil and assigns it to the Team field.
func (o *GateCreateDto) SetTeam(v nil) {
	o.Team = &v
}

// GetMeasureMetricLifts returns the MeasureMetricLifts field value if set, zero value otherwise.
func (o *GateCreateDto) GetMeasureMetricLifts() bool {
	if o == nil || IsNil(o.MeasureMetricLifts) {
		var ret bool
		return ret
	}
	return *o.MeasureMetricLifts
}

// GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetMeasureMetricLiftsOk() (*bool, bool) {
	if o == nil || IsNil(o.MeasureMetricLifts) {
		return nil, false
	}
	return o.MeasureMetricLifts, true
}

// HasMeasureMetricLifts returns a boolean if a field has been set.
func (o *GateCreateDto) HasMeasureMetricLifts() bool {
	if o != nil && !IsNil(o.MeasureMetricLifts) {
		return true
	}

	return false
}

// SetMeasureMetricLifts gets a reference to the given bool and assigns it to the MeasureMetricLifts field.
func (o *GateCreateDto) SetMeasureMetricLifts(v bool) {
	o.MeasureMetricLifts = &v
}

// GetMonitoringMetrics returns the MonitoringMetrics field value if set, zero value otherwise.
func (o *GateCreateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner {
	if o == nil || IsNil(o.MonitoringMetrics) {
		var ret []ExternalGateDtoMonitoringMetricsInner
		return ret
	}
	return o.MonitoringMetrics
}

// GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetMonitoringMetricsOk() ([]ExternalGateDtoMonitoringMetricsInner, bool) {
	if o == nil || IsNil(o.MonitoringMetrics) {
		return nil, false
	}
	return o.MonitoringMetrics, true
}

// HasMonitoringMetrics returns a boolean if a field has been set.
func (o *GateCreateDto) HasMonitoringMetrics() bool {
	if o != nil && !IsNil(o.MonitoringMetrics) {
		return true
	}

	return false
}

// SetMonitoringMetrics gets a reference to the given []ExternalGateDtoMonitoringMetricsInner and assigns it to the MonitoringMetrics field.
func (o *GateCreateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner) {
	o.MonitoringMetrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GateCreateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GateCreateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GateCreateDto) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GateCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GateCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GateCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GateCreateDto) SetId(v string) {
	o.Id = &v
}

func (o GateCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GateCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGateCreateDto struct {
	value *GateCreateDto
	isSet bool
}

func (v NullableGateCreateDto) Get() *GateCreateDto {
	return v.value
}

func (v *NullableGateCreateDto) Set(val *GateCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGateCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGateCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateCreateDto(val *GateCreateDto) *NullableGateCreateDto {
	return &NullableGateCreateDto{value: val, isSet: true}
}

func (v NullableGateCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

