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

// checks if the AssignmentSourcePartialUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentSourcePartialUpdateDto{}

// AssignmentSourcePartialUpdateDto struct for AssignmentSourcePartialUpdateDto
type AssignmentSourcePartialUpdateDto struct {
	// Unique identifier for the assignment source.
	Name *string `json:"name,omitempty"`
	// Detailed context and purpose of the assignment source.
	Description *string `json:"description,omitempty"`
	// Marks the assignment source as verified for internal trustworthiness.
	IsVerified *bool `json:"isVerified,omitempty"`
	// Tags for categorization and search.
	Tags []string `json:"tags,omitempty"`
	// SQL query defining the data source for assignments.
	Sql interface{} `json:"sql,omitempty"`
	// Column name representing the timestamp of assignments.
	TimestampColumn *string `json:"timestampColumn,omitempty"`
	// Column name for the experiment ID associated with the assignments.
	ExperimentIDColumn *string `json:"experimentIDColumn,omitempty"`
	// Column name for the group ID linked to the assignments.
	GroupIDColumn *string `json:"groupIDColumn,omitempty"`
	// Mappings of Statsig units to their respective columns.
	IdTypeMapping []AssignmentSourceCreationDtoIdTypeMappingInner `json:"idTypeMapping,omitempty"`
	// Specifies if the source can only be edited via the Console API.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}

// NewAssignmentSourcePartialUpdateDto instantiates a new AssignmentSourcePartialUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentSourcePartialUpdateDto() *AssignmentSourcePartialUpdateDto {
	this := AssignmentSourcePartialUpdateDto{}
	return &this
}

// NewAssignmentSourcePartialUpdateDtoWithDefaults instantiates a new AssignmentSourcePartialUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentSourcePartialUpdateDtoWithDefaults() *AssignmentSourcePartialUpdateDto {
	this := AssignmentSourcePartialUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssignmentSourcePartialUpdateDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssignmentSourcePartialUpdateDto) SetDescription(v string) {
	o.Description = &v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetIsVerified() bool {
	if o == nil || IsNil(o.IsVerified) {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVerified) {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasIsVerified() bool {
	if o != nil && !IsNil(o.IsVerified) {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *AssignmentSourcePartialUpdateDto) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AssignmentSourcePartialUpdateDto) SetTags(v []string) {
	o.Tags = v
}

// GetSql returns the Sql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssignmentSourcePartialUpdateDto) GetSql() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssignmentSourcePartialUpdateDto) GetSqlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Sql) {
		return nil, false
	}
	return &o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasSql() bool {
	if o != nil && !IsNil(o.Sql) {
		return true
	}

	return false
}

// SetSql gets a reference to the given interface{} and assigns it to the Sql field.
func (o *AssignmentSourcePartialUpdateDto) SetSql(v interface{}) {
	o.Sql = v
}

// GetTimestampColumn returns the TimestampColumn field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetTimestampColumn() string {
	if o == nil || IsNil(o.TimestampColumn) {
		var ret string
		return ret
	}
	return *o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetTimestampColumnOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampColumn) {
		return nil, false
	}
	return o.TimestampColumn, true
}

// HasTimestampColumn returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasTimestampColumn() bool {
	if o != nil && !IsNil(o.TimestampColumn) {
		return true
	}

	return false
}

// SetTimestampColumn gets a reference to the given string and assigns it to the TimestampColumn field.
func (o *AssignmentSourcePartialUpdateDto) SetTimestampColumn(v string) {
	o.TimestampColumn = &v
}

// GetExperimentIDColumn returns the ExperimentIDColumn field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetExperimentIDColumn() string {
	if o == nil || IsNil(o.ExperimentIDColumn) {
		var ret string
		return ret
	}
	return *o.ExperimentIDColumn
}

// GetExperimentIDColumnOk returns a tuple with the ExperimentIDColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetExperimentIDColumnOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentIDColumn) {
		return nil, false
	}
	return o.ExperimentIDColumn, true
}

// HasExperimentIDColumn returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasExperimentIDColumn() bool {
	if o != nil && !IsNil(o.ExperimentIDColumn) {
		return true
	}

	return false
}

// SetExperimentIDColumn gets a reference to the given string and assigns it to the ExperimentIDColumn field.
func (o *AssignmentSourcePartialUpdateDto) SetExperimentIDColumn(v string) {
	o.ExperimentIDColumn = &v
}

// GetGroupIDColumn returns the GroupIDColumn field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetGroupIDColumn() string {
	if o == nil || IsNil(o.GroupIDColumn) {
		var ret string
		return ret
	}
	return *o.GroupIDColumn
}

// GetGroupIDColumnOk returns a tuple with the GroupIDColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetGroupIDColumnOk() (*string, bool) {
	if o == nil || IsNil(o.GroupIDColumn) {
		return nil, false
	}
	return o.GroupIDColumn, true
}

// HasGroupIDColumn returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasGroupIDColumn() bool {
	if o != nil && !IsNil(o.GroupIDColumn) {
		return true
	}

	return false
}

// SetGroupIDColumn gets a reference to the given string and assigns it to the GroupIDColumn field.
func (o *AssignmentSourcePartialUpdateDto) SetGroupIDColumn(v string) {
	o.GroupIDColumn = &v
}

// GetIdTypeMapping returns the IdTypeMapping field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetIdTypeMapping() []AssignmentSourceCreationDtoIdTypeMappingInner {
	if o == nil || IsNil(o.IdTypeMapping) {
		var ret []AssignmentSourceCreationDtoIdTypeMappingInner
		return ret
	}
	return o.IdTypeMapping
}

// GetIdTypeMappingOk returns a tuple with the IdTypeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetIdTypeMappingOk() ([]AssignmentSourceCreationDtoIdTypeMappingInner, bool) {
	if o == nil || IsNil(o.IdTypeMapping) {
		return nil, false
	}
	return o.IdTypeMapping, true
}

// HasIdTypeMapping returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasIdTypeMapping() bool {
	if o != nil && !IsNil(o.IdTypeMapping) {
		return true
	}

	return false
}

// SetIdTypeMapping gets a reference to the given []AssignmentSourceCreationDtoIdTypeMappingInner and assigns it to the IdTypeMapping field.
func (o *AssignmentSourcePartialUpdateDto) SetIdTypeMapping(v []AssignmentSourceCreationDtoIdTypeMappingInner) {
	o.IdTypeMapping = v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *AssignmentSourcePartialUpdateDto) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourcePartialUpdateDto) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *AssignmentSourcePartialUpdateDto) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *AssignmentSourcePartialUpdateDto) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

func (o AssignmentSourcePartialUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentSourcePartialUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsVerified) {
		toSerialize["isVerified"] = o.IsVerified
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if !IsNil(o.TimestampColumn) {
		toSerialize["timestampColumn"] = o.TimestampColumn
	}
	if !IsNil(o.ExperimentIDColumn) {
		toSerialize["experimentIDColumn"] = o.ExperimentIDColumn
	}
	if !IsNil(o.GroupIDColumn) {
		toSerialize["groupIDColumn"] = o.GroupIDColumn
	}
	if !IsNil(o.IdTypeMapping) {
		toSerialize["idTypeMapping"] = o.IdTypeMapping
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return toSerialize, nil
}

type NullableAssignmentSourcePartialUpdateDto struct {
	value *AssignmentSourcePartialUpdateDto
	isSet bool
}

func (v NullableAssignmentSourcePartialUpdateDto) Get() *AssignmentSourcePartialUpdateDto {
	return v.value
}

func (v *NullableAssignmentSourcePartialUpdateDto) Set(val *AssignmentSourcePartialUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentSourcePartialUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentSourcePartialUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentSourcePartialUpdateDto(val *AssignmentSourcePartialUpdateDto) *NullableAssignmentSourcePartialUpdateDto {
	return &NullableAssignmentSourcePartialUpdateDto{value: val, isSet: true}
}

func (v NullableAssignmentSourcePartialUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentSourcePartialUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


