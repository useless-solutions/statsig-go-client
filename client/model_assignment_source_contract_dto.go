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

// checks if the AssignmentSourceContractDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentSourceContractDto{}

// AssignmentSourceContractDto struct for AssignmentSourceContractDto
type AssignmentSourceContractDto struct {
	// Unique identifier for the assignment source.
	Name string `json:"name"`
	// Detailed context and purpose of the assignment source.
	Description string `json:"description"`
	// Marks the assignment source as verified for internal trustworthiness.
	IsVerified *bool `json:"isVerified,omitempty"`
	// Tags for categorization and search.
	Tags []string `json:"tags"`
	// SQL query defining the data source for assignments.
	Sql string `json:"sql"`
	// Column name representing the timestamp of assignments.
	TimestampColumn string `json:"timestampColumn"`
	// Column name for the experiment ID associated with the assignments.
	ExperimentIDColumn string `json:"experimentIDColumn"`
	// Column name for the group ID linked to the assignments.
	GroupIDColumn string `json:"groupIDColumn"`
	// Mappings of Statsig units to their respective columns.
	IdTypeMapping []AssignmentSourceCreationDtoIdTypeMappingInner `json:"idTypeMapping"`
	// Specifies if the source can only be edited via the Console API.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}

type _AssignmentSourceContractDto AssignmentSourceContractDto

// NewAssignmentSourceContractDto instantiates a new AssignmentSourceContractDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentSourceContractDto(name string, description string, tags []string, sql string, timestampColumn string, experimentIDColumn string, groupIDColumn string, idTypeMapping []AssignmentSourceCreationDtoIdTypeMappingInner) *AssignmentSourceContractDto {
	this := AssignmentSourceContractDto{}
	this.Name = name
	this.Description = description
	this.Tags = tags
	this.Sql = sql
	this.TimestampColumn = timestampColumn
	this.ExperimentIDColumn = experimentIDColumn
	this.GroupIDColumn = groupIDColumn
	this.IdTypeMapping = idTypeMapping
	return &this
}

// NewAssignmentSourceContractDtoWithDefaults instantiates a new AssignmentSourceContractDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentSourceContractDtoWithDefaults() *AssignmentSourceContractDto {
	this := AssignmentSourceContractDto{}
	return &this
}

// GetName returns the Name field value
func (o *AssignmentSourceContractDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssignmentSourceContractDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AssignmentSourceContractDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AssignmentSourceContractDto) SetDescription(v string) {
	o.Description = v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *AssignmentSourceContractDto) GetIsVerified() bool {
	if o == nil || IsNil(o.IsVerified) {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVerified) {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *AssignmentSourceContractDto) HasIsVerified() bool {
	if o != nil && !IsNil(o.IsVerified) {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *AssignmentSourceContractDto) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetTags returns the Tags field value
func (o *AssignmentSourceContractDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *AssignmentSourceContractDto) SetTags(v []string) {
	o.Tags = v
}

// GetSql returns the Sql field value
func (o *AssignmentSourceContractDto) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *AssignmentSourceContractDto) SetSql(v string) {
	o.Sql = v
}

// GetTimestampColumn returns the TimestampColumn field value
func (o *AssignmentSourceContractDto) GetTimestampColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetTimestampColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampColumn, true
}

// SetTimestampColumn sets field value
func (o *AssignmentSourceContractDto) SetTimestampColumn(v string) {
	o.TimestampColumn = v
}

// GetExperimentIDColumn returns the ExperimentIDColumn field value
func (o *AssignmentSourceContractDto) GetExperimentIDColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExperimentIDColumn
}

// GetExperimentIDColumnOk returns a tuple with the ExperimentIDColumn field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetExperimentIDColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentIDColumn, true
}

// SetExperimentIDColumn sets field value
func (o *AssignmentSourceContractDto) SetExperimentIDColumn(v string) {
	o.ExperimentIDColumn = v
}

// GetGroupIDColumn returns the GroupIDColumn field value
func (o *AssignmentSourceContractDto) GetGroupIDColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupIDColumn
}

// GetGroupIDColumnOk returns a tuple with the GroupIDColumn field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetGroupIDColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupIDColumn, true
}

// SetGroupIDColumn sets field value
func (o *AssignmentSourceContractDto) SetGroupIDColumn(v string) {
	o.GroupIDColumn = v
}

// GetIdTypeMapping returns the IdTypeMapping field value
func (o *AssignmentSourceContractDto) GetIdTypeMapping() []AssignmentSourceCreationDtoIdTypeMappingInner {
	if o == nil {
		var ret []AssignmentSourceCreationDtoIdTypeMappingInner
		return ret
	}

	return o.IdTypeMapping
}

// GetIdTypeMappingOk returns a tuple with the IdTypeMapping field value
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetIdTypeMappingOk() ([]AssignmentSourceCreationDtoIdTypeMappingInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdTypeMapping, true
}

// SetIdTypeMapping sets field value
func (o *AssignmentSourceContractDto) SetIdTypeMapping(v []AssignmentSourceCreationDtoIdTypeMappingInner) {
	o.IdTypeMapping = v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *AssignmentSourceContractDto) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentSourceContractDto) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *AssignmentSourceContractDto) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *AssignmentSourceContractDto) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

func (o AssignmentSourceContractDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentSourceContractDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.IsVerified) {
		toSerialize["isVerified"] = o.IsVerified
	}
	toSerialize["tags"] = o.Tags
	toSerialize["sql"] = o.Sql
	toSerialize["timestampColumn"] = o.TimestampColumn
	toSerialize["experimentIDColumn"] = o.ExperimentIDColumn
	toSerialize["groupIDColumn"] = o.GroupIDColumn
	toSerialize["idTypeMapping"] = o.IdTypeMapping
	if !IsNil(o.IsReadOnly) {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return toSerialize, nil
}

func (o *AssignmentSourceContractDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"tags",
		"sql",
		"timestampColumn",
		"experimentIDColumn",
		"groupIDColumn",
		"idTypeMapping",
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

	varAssignmentSourceContractDto := _AssignmentSourceContractDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssignmentSourceContractDto)

	if err != nil {
		return err
	}

	*o = AssignmentSourceContractDto(varAssignmentSourceContractDto)

	return err
}

type NullableAssignmentSourceContractDto struct {
	value *AssignmentSourceContractDto
	isSet bool
}

func (v NullableAssignmentSourceContractDto) Get() *AssignmentSourceContractDto {
	return v.value
}

func (v *NullableAssignmentSourceContractDto) Set(val *AssignmentSourceContractDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentSourceContractDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentSourceContractDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentSourceContractDto(val *AssignmentSourceContractDto) *NullableAssignmentSourceContractDto {
	return &NullableAssignmentSourceContractDto{value: val, isSet: true}
}

func (v NullableAssignmentSourceContractDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentSourceContractDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


