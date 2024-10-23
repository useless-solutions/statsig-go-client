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

// checks if the EntityPropertySourcePartialUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityPropertySourcePartialUpdateDto{}

// EntityPropertySourcePartialUpdateDto struct for EntityPropertySourcePartialUpdateDto
type EntityPropertySourcePartialUpdateDto struct {
	// Unique identifier for the entity property source.
	Name *string `json:"name,omitempty"`
	// Detailed context and purpose of the entity property source.
	Description *string `json:"description,omitempty"`
	// Tags for categorization and search.
	Tags []string `json:"tags,omitempty"`
	// SQL query defining the data source
	Sql interface{} `json:"sql,omitempty"`
	// Optional column name for timestamp.
	TimestampColumn *string `json:"timestampColumn,omitempty"`
	// Indicates if the timestamp is treated as a day.
	TimestampAsDay *bool `json:"timestampAsDay,omitempty"`
	// Mappings of Statsig units to their columns.
	IdTypeMapping []EntityPropertySourceDtoIdTypeMappingInner `json:"idTypeMapping,omitempty"`
	// Specifies if the source can only be edited via the Console API.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}

// NewEntityPropertySourcePartialUpdateDto instantiates a new EntityPropertySourcePartialUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityPropertySourcePartialUpdateDto() *EntityPropertySourcePartialUpdateDto {
	this := EntityPropertySourcePartialUpdateDto{}
	return &this
}

// NewEntityPropertySourcePartialUpdateDtoWithDefaults instantiates a new EntityPropertySourcePartialUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityPropertySourcePartialUpdateDtoWithDefaults() *EntityPropertySourcePartialUpdateDto {
	this := EntityPropertySourcePartialUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntityPropertySourcePartialUpdateDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntityPropertySourcePartialUpdateDto) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EntityPropertySourcePartialUpdateDto) SetTags(v []string) {
	o.Tags = v
}

// GetSql returns the Sql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityPropertySourcePartialUpdateDto) GetSql() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityPropertySourcePartialUpdateDto) GetSqlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Sql) {
		return nil, false
	}
	return &o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasSql() bool {
	if o != nil && !IsNil(o.Sql) {
		return true
	}

	return false
}

// SetSql gets a reference to the given interface{} and assigns it to the Sql field.
func (o *EntityPropertySourcePartialUpdateDto) SetSql(v interface{}) {
	o.Sql = v
}

// GetTimestampColumn returns the TimestampColumn field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetTimestampColumn() string {
	if o == nil || IsNil(o.TimestampColumn) {
		var ret string
		return ret
	}
	return *o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetTimestampColumnOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampColumn) {
		return nil, false
	}
	return o.TimestampColumn, true
}

// HasTimestampColumn returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasTimestampColumn() bool {
	if o != nil && !IsNil(o.TimestampColumn) {
		return true
	}

	return false
}

// SetTimestampColumn gets a reference to the given string and assigns it to the TimestampColumn field.
func (o *EntityPropertySourcePartialUpdateDto) SetTimestampColumn(v string) {
	o.TimestampColumn = &v
}

// GetTimestampAsDay returns the TimestampAsDay field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetTimestampAsDay() bool {
	if o == nil || IsNil(o.TimestampAsDay) {
		var ret bool
		return ret
	}
	return *o.TimestampAsDay
}

// GetTimestampAsDayOk returns a tuple with the TimestampAsDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetTimestampAsDayOk() (*bool, bool) {
	if o == nil || IsNil(o.TimestampAsDay) {
		return nil, false
	}
	return o.TimestampAsDay, true
}

// HasTimestampAsDay returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasTimestampAsDay() bool {
	if o != nil && !IsNil(o.TimestampAsDay) {
		return true
	}

	return false
}

// SetTimestampAsDay gets a reference to the given bool and assigns it to the TimestampAsDay field.
func (o *EntityPropertySourcePartialUpdateDto) SetTimestampAsDay(v bool) {
	o.TimestampAsDay = &v
}

// GetIdTypeMapping returns the IdTypeMapping field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetIdTypeMapping() []EntityPropertySourceDtoIdTypeMappingInner {
	if o == nil || IsNil(o.IdTypeMapping) {
		var ret []EntityPropertySourceDtoIdTypeMappingInner
		return ret
	}
	return o.IdTypeMapping
}

// GetIdTypeMappingOk returns a tuple with the IdTypeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetIdTypeMappingOk() ([]EntityPropertySourceDtoIdTypeMappingInner, bool) {
	if o == nil || IsNil(o.IdTypeMapping) {
		return nil, false
	}
	return o.IdTypeMapping, true
}

// HasIdTypeMapping returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasIdTypeMapping() bool {
	if o != nil && !IsNil(o.IdTypeMapping) {
		return true
	}

	return false
}

// SetIdTypeMapping gets a reference to the given []EntityPropertySourceDtoIdTypeMappingInner and assigns it to the IdTypeMapping field.
func (o *EntityPropertySourcePartialUpdateDto) SetIdTypeMapping(v []EntityPropertySourceDtoIdTypeMappingInner) {
	o.IdTypeMapping = v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *EntityPropertySourcePartialUpdateDto) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySourcePartialUpdateDto) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *EntityPropertySourcePartialUpdateDto) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *EntityPropertySourcePartialUpdateDto) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

func (o EntityPropertySourcePartialUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityPropertySourcePartialUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !IsNil(o.TimestampAsDay) {
		toSerialize["timestampAsDay"] = o.TimestampAsDay
	}
	if !IsNil(o.IdTypeMapping) {
		toSerialize["idTypeMapping"] = o.IdTypeMapping
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return toSerialize, nil
}

type NullableEntityPropertySourcePartialUpdateDto struct {
	value *EntityPropertySourcePartialUpdateDto
	isSet bool
}

func (v NullableEntityPropertySourcePartialUpdateDto) Get() *EntityPropertySourcePartialUpdateDto {
	return v.value
}

func (v *NullableEntityPropertySourcePartialUpdateDto) Set(val *EntityPropertySourcePartialUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityPropertySourcePartialUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityPropertySourcePartialUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityPropertySourcePartialUpdateDto(val *EntityPropertySourcePartialUpdateDto) *NullableEntityPropertySourcePartialUpdateDto {
	return &NullableEntityPropertySourcePartialUpdateDto{value: val, isSet: true}
}

func (v NullableEntityPropertySourcePartialUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityPropertySourcePartialUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


