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

// checks if the TeamDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamDto{}

// TeamDto struct for TeamDto
type TeamDto struct {
	// The name of the team.
	Name string `json:"name"`
	// Description of the team.
	Description *string `json:"description,omitempty"`
	// The ID of the team.
	Id string `json:"id"`
	// Default gate metrics for the team.
	DefaultGateMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultGateMetrics"`
	// Default primary metrics for experiments in the team.
	DefaultExperimentPrimaryMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultExperimentPrimaryMetrics"`
	// Default secondary metrics for experiments in the team.
	DefaultExperimentSecondaryMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultExperimentSecondaryMetrics"`
	// Default holdout metrics for the team.
	DefaultHoldoutMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultHoldoutMetrics"`
	// Who can change team configurations: \"anyone\" or \"team_only\".
	ChangeTeamConfigs string `json:"changeTeamConfigs"`
	// Who can review and approve changes: \"anyone\", \"team_only\", or \"admin_only\".
	ReviewApproval string `json:"reviewApproval"`
	// Default target applications for the team.
	DefaultTargetApplications []string `json:"defaultTargetApplications"`
	// Default holdout ID for the team, if applicable.
	DefaultHoldoutID *nil `json:"defaultHoldoutID,omitempty"`
	// Whether reviews are required for changes, if applicable.
	RequireReviews *nil `json:"requireReviews,omitempty"`
	// Whether gate templates are required for the team, if applicable.
	RequireGateTemplates *nil `json:"requireGateTemplates,omitempty"`
	// Whether experiment templates are required for the team, if applicable.
	RequireExperimentTemplates *nil `json:"requireExperimentTemplates,omitempty"`
	Members []TeamDtoMembersInner `json:"members"`
	Admins []TeamDtoMembersInner `json:"admins"`
}

type _TeamDto TeamDto

// NewTeamDto instantiates a new TeamDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamDto(name string, id string, defaultGateMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentPrimaryMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentSecondaryMetrics []TeamDtoDefaultGateMetricsInner, defaultHoldoutMetrics []TeamDtoDefaultGateMetricsInner, changeTeamConfigs string, reviewApproval string, defaultTargetApplications []string, members []TeamDtoMembersInner, admins []TeamDtoMembersInner) *TeamDto {
	this := TeamDto{}
	this.Name = name
	this.Id = id
	this.DefaultGateMetrics = defaultGateMetrics
	this.DefaultExperimentPrimaryMetrics = defaultExperimentPrimaryMetrics
	this.DefaultExperimentSecondaryMetrics = defaultExperimentSecondaryMetrics
	this.DefaultHoldoutMetrics = defaultHoldoutMetrics
	this.ChangeTeamConfigs = changeTeamConfigs
	this.ReviewApproval = reviewApproval
	this.DefaultTargetApplications = defaultTargetApplications
	this.Members = members
	this.Admins = admins
	return &this
}

// NewTeamDtoWithDefaults instantiates a new TeamDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDtoWithDefaults() *TeamDto {
	this := TeamDto{}
	return &this
}

// GetName returns the Name field value
func (o *TeamDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamDto) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *TeamDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamDto) SetId(v string) {
	o.Id = v
}

// GetDefaultGateMetrics returns the DefaultGateMetrics field value
func (o *TeamDto) GetDefaultGateMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}

	return o.DefaultGateMetrics
}

// GetDefaultGateMetricsOk returns a tuple with the DefaultGateMetrics field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultGateMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultGateMetrics, true
}

// SetDefaultGateMetrics sets field value
func (o *TeamDto) SetDefaultGateMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultGateMetrics = v
}

// GetDefaultExperimentPrimaryMetrics returns the DefaultExperimentPrimaryMetrics field value
func (o *TeamDto) GetDefaultExperimentPrimaryMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}

	return o.DefaultExperimentPrimaryMetrics
}

// GetDefaultExperimentPrimaryMetricsOk returns a tuple with the DefaultExperimentPrimaryMetrics field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultExperimentPrimaryMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultExperimentPrimaryMetrics, true
}

// SetDefaultExperimentPrimaryMetrics sets field value
func (o *TeamDto) SetDefaultExperimentPrimaryMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultExperimentPrimaryMetrics = v
}

// GetDefaultExperimentSecondaryMetrics returns the DefaultExperimentSecondaryMetrics field value
func (o *TeamDto) GetDefaultExperimentSecondaryMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}

	return o.DefaultExperimentSecondaryMetrics
}

// GetDefaultExperimentSecondaryMetricsOk returns a tuple with the DefaultExperimentSecondaryMetrics field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultExperimentSecondaryMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultExperimentSecondaryMetrics, true
}

// SetDefaultExperimentSecondaryMetrics sets field value
func (o *TeamDto) SetDefaultExperimentSecondaryMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultExperimentSecondaryMetrics = v
}

// GetDefaultHoldoutMetrics returns the DefaultHoldoutMetrics field value
func (o *TeamDto) GetDefaultHoldoutMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}

	return o.DefaultHoldoutMetrics
}

// GetDefaultHoldoutMetricsOk returns a tuple with the DefaultHoldoutMetrics field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultHoldoutMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultHoldoutMetrics, true
}

// SetDefaultHoldoutMetrics sets field value
func (o *TeamDto) SetDefaultHoldoutMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultHoldoutMetrics = v
}

// GetChangeTeamConfigs returns the ChangeTeamConfigs field value
func (o *TeamDto) GetChangeTeamConfigs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeTeamConfigs
}

// GetChangeTeamConfigsOk returns a tuple with the ChangeTeamConfigs field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetChangeTeamConfigsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeTeamConfigs, true
}

// SetChangeTeamConfigs sets field value
func (o *TeamDto) SetChangeTeamConfigs(v string) {
	o.ChangeTeamConfigs = v
}

// GetReviewApproval returns the ReviewApproval field value
func (o *TeamDto) GetReviewApproval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewApproval
}

// GetReviewApprovalOk returns a tuple with the ReviewApproval field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetReviewApprovalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewApproval, true
}

// SetReviewApproval sets field value
func (o *TeamDto) SetReviewApproval(v string) {
	o.ReviewApproval = v
}

// GetDefaultTargetApplications returns the DefaultTargetApplications field value
func (o *TeamDto) GetDefaultTargetApplications() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultTargetApplications
}

// GetDefaultTargetApplicationsOk returns a tuple with the DefaultTargetApplications field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultTargetApplicationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultTargetApplications, true
}

// SetDefaultTargetApplications sets field value
func (o *TeamDto) SetDefaultTargetApplications(v []string) {
	o.DefaultTargetApplications = v
}

// GetDefaultHoldoutID returns the DefaultHoldoutID field value if set, zero value otherwise.
func (o *TeamDto) GetDefaultHoldoutID() nil {
	if o == nil || IsNil(o.DefaultHoldoutID) {
		var ret nil
		return ret
	}
	return *o.DefaultHoldoutID
}

// GetDefaultHoldoutIDOk returns a tuple with the DefaultHoldoutID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamDto) GetDefaultHoldoutIDOk() (*nil, bool) {
	if o == nil || IsNil(o.DefaultHoldoutID) {
		return nil, false
	}
	return o.DefaultHoldoutID, true
}

// HasDefaultHoldoutID returns a boolean if a field has been set.
func (o *TeamDto) HasDefaultHoldoutID() bool {
	if o != nil && !IsNil(o.DefaultHoldoutID) {
		return true
	}

	return false
}

// SetDefaultHoldoutID gets a reference to the given nil and assigns it to the DefaultHoldoutID field.
func (o *TeamDto) SetDefaultHoldoutID(v nil) {
	o.DefaultHoldoutID = &v
}

// GetRequireReviews returns the RequireReviews field value if set, zero value otherwise.
func (o *TeamDto) GetRequireReviews() nil {
	if o == nil || IsNil(o.RequireReviews) {
		var ret nil
		return ret
	}
	return *o.RequireReviews
}

// GetRequireReviewsOk returns a tuple with the RequireReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamDto) GetRequireReviewsOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireReviews) {
		return nil, false
	}
	return o.RequireReviews, true
}

// HasRequireReviews returns a boolean if a field has been set.
func (o *TeamDto) HasRequireReviews() bool {
	if o != nil && !IsNil(o.RequireReviews) {
		return true
	}

	return false
}

// SetRequireReviews gets a reference to the given nil and assigns it to the RequireReviews field.
func (o *TeamDto) SetRequireReviews(v nil) {
	o.RequireReviews = &v
}

// GetRequireGateTemplates returns the RequireGateTemplates field value if set, zero value otherwise.
func (o *TeamDto) GetRequireGateTemplates() nil {
	if o == nil || IsNil(o.RequireGateTemplates) {
		var ret nil
		return ret
	}
	return *o.RequireGateTemplates
}

// GetRequireGateTemplatesOk returns a tuple with the RequireGateTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamDto) GetRequireGateTemplatesOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireGateTemplates) {
		return nil, false
	}
	return o.RequireGateTemplates, true
}

// HasRequireGateTemplates returns a boolean if a field has been set.
func (o *TeamDto) HasRequireGateTemplates() bool {
	if o != nil && !IsNil(o.RequireGateTemplates) {
		return true
	}

	return false
}

// SetRequireGateTemplates gets a reference to the given nil and assigns it to the RequireGateTemplates field.
func (o *TeamDto) SetRequireGateTemplates(v nil) {
	o.RequireGateTemplates = &v
}

// GetRequireExperimentTemplates returns the RequireExperimentTemplates field value if set, zero value otherwise.
func (o *TeamDto) GetRequireExperimentTemplates() nil {
	if o == nil || IsNil(o.RequireExperimentTemplates) {
		var ret nil
		return ret
	}
	return *o.RequireExperimentTemplates
}

// GetRequireExperimentTemplatesOk returns a tuple with the RequireExperimentTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamDto) GetRequireExperimentTemplatesOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireExperimentTemplates) {
		return nil, false
	}
	return o.RequireExperimentTemplates, true
}

// HasRequireExperimentTemplates returns a boolean if a field has been set.
func (o *TeamDto) HasRequireExperimentTemplates() bool {
	if o != nil && !IsNil(o.RequireExperimentTemplates) {
		return true
	}

	return false
}

// SetRequireExperimentTemplates gets a reference to the given nil and assigns it to the RequireExperimentTemplates field.
func (o *TeamDto) SetRequireExperimentTemplates(v nil) {
	o.RequireExperimentTemplates = &v
}

// GetMembers returns the Members field value
func (o *TeamDto) GetMembers() []TeamDtoMembersInner {
	if o == nil {
		var ret []TeamDtoMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetMembersOk() ([]TeamDtoMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *TeamDto) SetMembers(v []TeamDtoMembersInner) {
	o.Members = v
}

// GetAdmins returns the Admins field value
func (o *TeamDto) GetAdmins() []TeamDtoMembersInner {
	if o == nil {
		var ret []TeamDtoMembersInner
		return ret
	}

	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value
// and a boolean to check if the value has been set.
func (o *TeamDto) GetAdminsOk() ([]TeamDtoMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Admins, true
}

// SetAdmins sets field value
func (o *TeamDto) SetAdmins(v []TeamDtoMembersInner) {
	o.Admins = v
}

func (o TeamDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["defaultGateMetrics"] = o.DefaultGateMetrics
	toSerialize["defaultExperimentPrimaryMetrics"] = o.DefaultExperimentPrimaryMetrics
	toSerialize["defaultExperimentSecondaryMetrics"] = o.DefaultExperimentSecondaryMetrics
	toSerialize["defaultHoldoutMetrics"] = o.DefaultHoldoutMetrics
	toSerialize["changeTeamConfigs"] = o.ChangeTeamConfigs
	toSerialize["reviewApproval"] = o.ReviewApproval
	toSerialize["defaultTargetApplications"] = o.DefaultTargetApplications
	if !IsNil(o.DefaultHoldoutID) {
		toSerialize["defaultHoldoutID"] = o.DefaultHoldoutID
	}
	if !IsNil(o.RequireReviews) {
		toSerialize["requireReviews"] = o.RequireReviews
	}
	if !IsNil(o.RequireGateTemplates) {
		toSerialize["requireGateTemplates"] = o.RequireGateTemplates
	}
	if !IsNil(o.RequireExperimentTemplates) {
		toSerialize["requireExperimentTemplates"] = o.RequireExperimentTemplates
	}
	toSerialize["members"] = o.Members
	toSerialize["admins"] = o.Admins
	return toSerialize, nil
}

func (o *TeamDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"defaultGateMetrics",
		"defaultExperimentPrimaryMetrics",
		"defaultExperimentSecondaryMetrics",
		"defaultHoldoutMetrics",
		"changeTeamConfigs",
		"reviewApproval",
		"defaultTargetApplications",
		"members",
		"admins",
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

	varTeamDto := _TeamDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamDto)

	if err != nil {
		return err
	}

	*o = TeamDto(varTeamDto)

	return err
}

type NullableTeamDto struct {
	value *TeamDto
	isSet bool
}

func (v NullableTeamDto) Get() *TeamDto {
	return v.value
}

func (v *NullableTeamDto) Set(val *TeamDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamDto(val *TeamDto) *NullableTeamDto {
	return &NullableTeamDto{value: val, isSet: true}
}

func (v NullableTeamDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

