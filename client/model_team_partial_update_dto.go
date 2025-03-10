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

// checks if the TeamPartialUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamPartialUpdateDto{}

// TeamPartialUpdateDto struct for TeamPartialUpdateDto
type TeamPartialUpdateDto struct {
	// The name of the team.
	Name *string `json:"name,omitempty"`
	// Description of the team.
	Description *string `json:"description,omitempty"`
	// The ID of the team.
	Id *string `json:"id,omitempty"`
	// Array of member email addresses in the team.
	Members []string `json:"members,omitempty"`
	// Array of admin email addresses in the team.
	Admins []string `json:"admins,omitempty"`
	// Default gate metrics for the team.
	DefaultGateMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultGateMetrics,omitempty"`
	// Default primary metrics for experiments in the team.
	DefaultExperimentPrimaryMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultExperimentPrimaryMetrics,omitempty"`
	// Default secondary metrics for experiments in the team.
	DefaultExperimentSecondaryMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultExperimentSecondaryMetrics,omitempty"`
	// Default holdout metrics for the team.
	DefaultHoldoutMetrics []TeamDtoDefaultGateMetricsInner `json:"defaultHoldoutMetrics,omitempty"`
	// Who can change team configurations: \"anyone\" or \"team_only\".
	ChangeTeamConfigs *string `json:"changeTeamConfigs,omitempty"`
	// Who can review and approve changes: \"anyone\", \"team_only\", or \"admin_only\".
	ReviewApproval *string `json:"reviewApproval,omitempty"`
	// Default target applications for the team.
	DefaultTargetApplications []string `json:"defaultTargetApplications,omitempty"`
	// Default holdout ID for the team, if applicable.
	DefaultHoldoutID *nil `json:"defaultHoldoutID,omitempty"`
	// Whether reviews are required for changes, if applicable.
	RequireReviews *nil `json:"requireReviews,omitempty"`
	// Whether gate templates are required for the team, if applicable.
	RequireGateTemplates *nil `json:"requireGateTemplates,omitempty"`
	// Whether experiment templates are required for the team, if applicable.
	RequireExperimentTemplates *nil `json:"requireExperimentTemplates,omitempty"`
}

// NewTeamPartialUpdateDto instantiates a new TeamPartialUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamPartialUpdateDto() *TeamPartialUpdateDto {
	this := TeamPartialUpdateDto{}
	return &this
}

// NewTeamPartialUpdateDtoWithDefaults instantiates a new TeamPartialUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamPartialUpdateDtoWithDefaults() *TeamPartialUpdateDto {
	this := TeamPartialUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamPartialUpdateDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamPartialUpdateDto) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TeamPartialUpdateDto) SetId(v string) {
	o.Id = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetMembers() []string {
	if o == nil || IsNil(o.Members) {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *TeamPartialUpdateDto) SetMembers(v []string) {
	o.Members = v
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetAdmins() []string {
	if o == nil || IsNil(o.Admins) {
		var ret []string
		return ret
	}
	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetAdminsOk() ([]string, bool) {
	if o == nil || IsNil(o.Admins) {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasAdmins() bool {
	if o != nil && !IsNil(o.Admins) {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given []string and assigns it to the Admins field.
func (o *TeamPartialUpdateDto) SetAdmins(v []string) {
	o.Admins = v
}

// GetDefaultGateMetrics returns the DefaultGateMetrics field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultGateMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil || IsNil(o.DefaultGateMetrics) {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}
	return o.DefaultGateMetrics
}

// GetDefaultGateMetricsOk returns a tuple with the DefaultGateMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultGateMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil || IsNil(o.DefaultGateMetrics) {
		return nil, false
	}
	return o.DefaultGateMetrics, true
}

// HasDefaultGateMetrics returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultGateMetrics() bool {
	if o != nil && !IsNil(o.DefaultGateMetrics) {
		return true
	}

	return false
}

// SetDefaultGateMetrics gets a reference to the given []TeamDtoDefaultGateMetricsInner and assigns it to the DefaultGateMetrics field.
func (o *TeamPartialUpdateDto) SetDefaultGateMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultGateMetrics = v
}

// GetDefaultExperimentPrimaryMetrics returns the DefaultExperimentPrimaryMetrics field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultExperimentPrimaryMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil || IsNil(o.DefaultExperimentPrimaryMetrics) {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}
	return o.DefaultExperimentPrimaryMetrics
}

// GetDefaultExperimentPrimaryMetricsOk returns a tuple with the DefaultExperimentPrimaryMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultExperimentPrimaryMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil || IsNil(o.DefaultExperimentPrimaryMetrics) {
		return nil, false
	}
	return o.DefaultExperimentPrimaryMetrics, true
}

// HasDefaultExperimentPrimaryMetrics returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultExperimentPrimaryMetrics() bool {
	if o != nil && !IsNil(o.DefaultExperimentPrimaryMetrics) {
		return true
	}

	return false
}

// SetDefaultExperimentPrimaryMetrics gets a reference to the given []TeamDtoDefaultGateMetricsInner and assigns it to the DefaultExperimentPrimaryMetrics field.
func (o *TeamPartialUpdateDto) SetDefaultExperimentPrimaryMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultExperimentPrimaryMetrics = v
}

// GetDefaultExperimentSecondaryMetrics returns the DefaultExperimentSecondaryMetrics field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultExperimentSecondaryMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil || IsNil(o.DefaultExperimentSecondaryMetrics) {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}
	return o.DefaultExperimentSecondaryMetrics
}

// GetDefaultExperimentSecondaryMetricsOk returns a tuple with the DefaultExperimentSecondaryMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultExperimentSecondaryMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil || IsNil(o.DefaultExperimentSecondaryMetrics) {
		return nil, false
	}
	return o.DefaultExperimentSecondaryMetrics, true
}

// HasDefaultExperimentSecondaryMetrics returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultExperimentSecondaryMetrics() bool {
	if o != nil && !IsNil(o.DefaultExperimentSecondaryMetrics) {
		return true
	}

	return false
}

// SetDefaultExperimentSecondaryMetrics gets a reference to the given []TeamDtoDefaultGateMetricsInner and assigns it to the DefaultExperimentSecondaryMetrics field.
func (o *TeamPartialUpdateDto) SetDefaultExperimentSecondaryMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultExperimentSecondaryMetrics = v
}

// GetDefaultHoldoutMetrics returns the DefaultHoldoutMetrics field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultHoldoutMetrics() []TeamDtoDefaultGateMetricsInner {
	if o == nil || IsNil(o.DefaultHoldoutMetrics) {
		var ret []TeamDtoDefaultGateMetricsInner
		return ret
	}
	return o.DefaultHoldoutMetrics
}

// GetDefaultHoldoutMetricsOk returns a tuple with the DefaultHoldoutMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultHoldoutMetricsOk() ([]TeamDtoDefaultGateMetricsInner, bool) {
	if o == nil || IsNil(o.DefaultHoldoutMetrics) {
		return nil, false
	}
	return o.DefaultHoldoutMetrics, true
}

// HasDefaultHoldoutMetrics returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultHoldoutMetrics() bool {
	if o != nil && !IsNil(o.DefaultHoldoutMetrics) {
		return true
	}

	return false
}

// SetDefaultHoldoutMetrics gets a reference to the given []TeamDtoDefaultGateMetricsInner and assigns it to the DefaultHoldoutMetrics field.
func (o *TeamPartialUpdateDto) SetDefaultHoldoutMetrics(v []TeamDtoDefaultGateMetricsInner) {
	o.DefaultHoldoutMetrics = v
}

// GetChangeTeamConfigs returns the ChangeTeamConfigs field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetChangeTeamConfigs() string {
	if o == nil || IsNil(o.ChangeTeamConfigs) {
		var ret string
		return ret
	}
	return *o.ChangeTeamConfigs
}

// GetChangeTeamConfigsOk returns a tuple with the ChangeTeamConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetChangeTeamConfigsOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeTeamConfigs) {
		return nil, false
	}
	return o.ChangeTeamConfigs, true
}

// HasChangeTeamConfigs returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasChangeTeamConfigs() bool {
	if o != nil && !IsNil(o.ChangeTeamConfigs) {
		return true
	}

	return false
}

// SetChangeTeamConfigs gets a reference to the given string and assigns it to the ChangeTeamConfigs field.
func (o *TeamPartialUpdateDto) SetChangeTeamConfigs(v string) {
	o.ChangeTeamConfigs = &v
}

// GetReviewApproval returns the ReviewApproval field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetReviewApproval() string {
	if o == nil || IsNil(o.ReviewApproval) {
		var ret string
		return ret
	}
	return *o.ReviewApproval
}

// GetReviewApprovalOk returns a tuple with the ReviewApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetReviewApprovalOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewApproval) {
		return nil, false
	}
	return o.ReviewApproval, true
}

// HasReviewApproval returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasReviewApproval() bool {
	if o != nil && !IsNil(o.ReviewApproval) {
		return true
	}

	return false
}

// SetReviewApproval gets a reference to the given string and assigns it to the ReviewApproval field.
func (o *TeamPartialUpdateDto) SetReviewApproval(v string) {
	o.ReviewApproval = &v
}

// GetDefaultTargetApplications returns the DefaultTargetApplications field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultTargetApplications() []string {
	if o == nil || IsNil(o.DefaultTargetApplications) {
		var ret []string
		return ret
	}
	return o.DefaultTargetApplications
}

// GetDefaultTargetApplicationsOk returns a tuple with the DefaultTargetApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultTargetApplicationsOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultTargetApplications) {
		return nil, false
	}
	return o.DefaultTargetApplications, true
}

// HasDefaultTargetApplications returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultTargetApplications() bool {
	if o != nil && !IsNil(o.DefaultTargetApplications) {
		return true
	}

	return false
}

// SetDefaultTargetApplications gets a reference to the given []string and assigns it to the DefaultTargetApplications field.
func (o *TeamPartialUpdateDto) SetDefaultTargetApplications(v []string) {
	o.DefaultTargetApplications = v
}

// GetDefaultHoldoutID returns the DefaultHoldoutID field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetDefaultHoldoutID() nil {
	if o == nil || IsNil(o.DefaultHoldoutID) {
		var ret nil
		return ret
	}
	return *o.DefaultHoldoutID
}

// GetDefaultHoldoutIDOk returns a tuple with the DefaultHoldoutID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetDefaultHoldoutIDOk() (*nil, bool) {
	if o == nil || IsNil(o.DefaultHoldoutID) {
		return nil, false
	}
	return o.DefaultHoldoutID, true
}

// HasDefaultHoldoutID returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasDefaultHoldoutID() bool {
	if o != nil && !IsNil(o.DefaultHoldoutID) {
		return true
	}

	return false
}

// SetDefaultHoldoutID gets a reference to the given nil and assigns it to the DefaultHoldoutID field.
func (o *TeamPartialUpdateDto) SetDefaultHoldoutID(v nil) {
	o.DefaultHoldoutID = &v
}

// GetRequireReviews returns the RequireReviews field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetRequireReviews() nil {
	if o == nil || IsNil(o.RequireReviews) {
		var ret nil
		return ret
	}
	return *o.RequireReviews
}

// GetRequireReviewsOk returns a tuple with the RequireReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetRequireReviewsOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireReviews) {
		return nil, false
	}
	return o.RequireReviews, true
}

// HasRequireReviews returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasRequireReviews() bool {
	if o != nil && !IsNil(o.RequireReviews) {
		return true
	}

	return false
}

// SetRequireReviews gets a reference to the given nil and assigns it to the RequireReviews field.
func (o *TeamPartialUpdateDto) SetRequireReviews(v nil) {
	o.RequireReviews = &v
}

// GetRequireGateTemplates returns the RequireGateTemplates field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetRequireGateTemplates() nil {
	if o == nil || IsNil(o.RequireGateTemplates) {
		var ret nil
		return ret
	}
	return *o.RequireGateTemplates
}

// GetRequireGateTemplatesOk returns a tuple with the RequireGateTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetRequireGateTemplatesOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireGateTemplates) {
		return nil, false
	}
	return o.RequireGateTemplates, true
}

// HasRequireGateTemplates returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasRequireGateTemplates() bool {
	if o != nil && !IsNil(o.RequireGateTemplates) {
		return true
	}

	return false
}

// SetRequireGateTemplates gets a reference to the given nil and assigns it to the RequireGateTemplates field.
func (o *TeamPartialUpdateDto) SetRequireGateTemplates(v nil) {
	o.RequireGateTemplates = &v
}

// GetRequireExperimentTemplates returns the RequireExperimentTemplates field value if set, zero value otherwise.
func (o *TeamPartialUpdateDto) GetRequireExperimentTemplates() nil {
	if o == nil || IsNil(o.RequireExperimentTemplates) {
		var ret nil
		return ret
	}
	return *o.RequireExperimentTemplates
}

// GetRequireExperimentTemplatesOk returns a tuple with the RequireExperimentTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPartialUpdateDto) GetRequireExperimentTemplatesOk() (*nil, bool) {
	if o == nil || IsNil(o.RequireExperimentTemplates) {
		return nil, false
	}
	return o.RequireExperimentTemplates, true
}

// HasRequireExperimentTemplates returns a boolean if a field has been set.
func (o *TeamPartialUpdateDto) HasRequireExperimentTemplates() bool {
	if o != nil && !IsNil(o.RequireExperimentTemplates) {
		return true
	}

	return false
}

// SetRequireExperimentTemplates gets a reference to the given nil and assigns it to the RequireExperimentTemplates field.
func (o *TeamPartialUpdateDto) SetRequireExperimentTemplates(v nil) {
	o.RequireExperimentTemplates = &v
}

func (o TeamPartialUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamPartialUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Admins) {
		toSerialize["admins"] = o.Admins
	}
	if !IsNil(o.DefaultGateMetrics) {
		toSerialize["defaultGateMetrics"] = o.DefaultGateMetrics
	}
	if !IsNil(o.DefaultExperimentPrimaryMetrics) {
		toSerialize["defaultExperimentPrimaryMetrics"] = o.DefaultExperimentPrimaryMetrics
	}
	if !IsNil(o.DefaultExperimentSecondaryMetrics) {
		toSerialize["defaultExperimentSecondaryMetrics"] = o.DefaultExperimentSecondaryMetrics
	}
	if !IsNil(o.DefaultHoldoutMetrics) {
		toSerialize["defaultHoldoutMetrics"] = o.DefaultHoldoutMetrics
	}
	if !IsNil(o.ChangeTeamConfigs) {
		toSerialize["changeTeamConfigs"] = o.ChangeTeamConfigs
	}
	if !IsNil(o.ReviewApproval) {
		toSerialize["reviewApproval"] = o.ReviewApproval
	}
	if !IsNil(o.DefaultTargetApplications) {
		toSerialize["defaultTargetApplications"] = o.DefaultTargetApplications
	}
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
	return toSerialize, nil
}

type NullableTeamPartialUpdateDto struct {
	value *TeamPartialUpdateDto
	isSet bool
}

func (v NullableTeamPartialUpdateDto) Get() *TeamPartialUpdateDto {
	return v.value
}

func (v *NullableTeamPartialUpdateDto) Set(val *TeamPartialUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamPartialUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamPartialUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamPartialUpdateDto(val *TeamPartialUpdateDto) *NullableTeamPartialUpdateDto {
	return &NullableTeamPartialUpdateDto{value: val, isSet: true}
}

func (v NullableTeamPartialUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamPartialUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


