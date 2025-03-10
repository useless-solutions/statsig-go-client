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

// checks if the MetricSourceContractDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricSourceContractDto{}

// MetricSourceContractDto struct for MetricSourceContractDto
type MetricSourceContractDto struct {
	// The name of the metric source, serving as its primary identifier.
	Name string `json:"name"`
	// A detailed description of the metric source, providing context and usage information.
	Description string `json:"description"`
	// Optional tags for categorizing the metric source and improving searchability.
	Tags []string `json:"tags,omitempty"`
	// The SQL query or statement used to extract data from the metric source.
	Sql string `json:"sql"`
	// The name of the column containing timestamp data for the metric source.
	TimestampColumn string `json:"timestampColumn"`
	// Indicates whether the timestamp should be treated as a day-level granularity.
	TimestampAsDay *bool `json:"timestampAsDay,omitempty"`
	// Array defining the mapping between Statsig unit IDs and their respective source columns.
	IdTypeMapping []MetricSourceContractDtoIdTypeMappingInner `json:"idTypeMapping"`
	// The type of source, indicating whether it is a database table or a custom query.
	SourceType *string `json:"sourceType,omitempty"`
	// The name of the database table if the source type is \"table\".
	TableName *string `json:"tableName,omitempty"`
	// Optional array defining mappings for custom fields using specific formulas.
	CustomFieldMapping []MetricSourceContractDtoCustomFieldMappingInner `json:"customFieldMapping,omitempty"`
	// Specifies if the source can only be edited via the Console API.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}

type _MetricSourceContractDto MetricSourceContractDto

// NewMetricSourceContractDto instantiates a new MetricSourceContractDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricSourceContractDto(name string, description string, sql string, timestampColumn string, idTypeMapping []MetricSourceContractDtoIdTypeMappingInner) *MetricSourceContractDto {
	this := MetricSourceContractDto{}
	this.Name = name
	this.Description = description
	this.Sql = sql
	this.TimestampColumn = timestampColumn
	this.IdTypeMapping = idTypeMapping
	return &this
}

// NewMetricSourceContractDtoWithDefaults instantiates a new MetricSourceContractDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricSourceContractDtoWithDefaults() *MetricSourceContractDto {
	this := MetricSourceContractDto{}
	return &this
}

// GetName returns the Name field value
func (o *MetricSourceContractDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricSourceContractDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *MetricSourceContractDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MetricSourceContractDto) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricSourceContractDto) SetTags(v []string) {
	o.Tags = v
}

// GetSql returns the Sql field value
func (o *MetricSourceContractDto) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *MetricSourceContractDto) SetSql(v string) {
	o.Sql = v
}

// GetTimestampColumn returns the TimestampColumn field value
func (o *MetricSourceContractDto) GetTimestampColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimestampColumn
}

// GetTimestampColumnOk returns a tuple with the TimestampColumn field value
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetTimestampColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampColumn, true
}

// SetTimestampColumn sets field value
func (o *MetricSourceContractDto) SetTimestampColumn(v string) {
	o.TimestampColumn = v
}

// GetTimestampAsDay returns the TimestampAsDay field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetTimestampAsDay() bool {
	if o == nil || IsNil(o.TimestampAsDay) {
		var ret bool
		return ret
	}
	return *o.TimestampAsDay
}

// GetTimestampAsDayOk returns a tuple with the TimestampAsDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetTimestampAsDayOk() (*bool, bool) {
	if o == nil || IsNil(o.TimestampAsDay) {
		return nil, false
	}
	return o.TimestampAsDay, true
}

// HasTimestampAsDay returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasTimestampAsDay() bool {
	if o != nil && !IsNil(o.TimestampAsDay) {
		return true
	}

	return false
}

// SetTimestampAsDay gets a reference to the given bool and assigns it to the TimestampAsDay field.
func (o *MetricSourceContractDto) SetTimestampAsDay(v bool) {
	o.TimestampAsDay = &v
}

// GetIdTypeMapping returns the IdTypeMapping field value
func (o *MetricSourceContractDto) GetIdTypeMapping() []MetricSourceContractDtoIdTypeMappingInner {
	if o == nil {
		var ret []MetricSourceContractDtoIdTypeMappingInner
		return ret
	}

	return o.IdTypeMapping
}

// GetIdTypeMappingOk returns a tuple with the IdTypeMapping field value
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetIdTypeMappingOk() ([]MetricSourceContractDtoIdTypeMappingInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdTypeMapping, true
}

// SetIdTypeMapping sets field value
func (o *MetricSourceContractDto) SetIdTypeMapping(v []MetricSourceContractDtoIdTypeMappingInner) {
	o.IdTypeMapping = v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *MetricSourceContractDto) SetSourceType(v string) {
	o.SourceType = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *MetricSourceContractDto) SetTableName(v string) {
	o.TableName = &v
}

// GetCustomFieldMapping returns the CustomFieldMapping field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetCustomFieldMapping() []MetricSourceContractDtoCustomFieldMappingInner {
	if o == nil || IsNil(o.CustomFieldMapping) {
		var ret []MetricSourceContractDtoCustomFieldMappingInner
		return ret
	}
	return o.CustomFieldMapping
}

// GetCustomFieldMappingOk returns a tuple with the CustomFieldMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetCustomFieldMappingOk() ([]MetricSourceContractDtoCustomFieldMappingInner, bool) {
	if o == nil || IsNil(o.CustomFieldMapping) {
		return nil, false
	}
	return o.CustomFieldMapping, true
}

// HasCustomFieldMapping returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasCustomFieldMapping() bool {
	if o != nil && !IsNil(o.CustomFieldMapping) {
		return true
	}

	return false
}

// SetCustomFieldMapping gets a reference to the given []MetricSourceContractDtoCustomFieldMappingInner and assigns it to the CustomFieldMapping field.
func (o *MetricSourceContractDto) SetCustomFieldMapping(v []MetricSourceContractDtoCustomFieldMappingInner) {
	o.CustomFieldMapping = v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *MetricSourceContractDto) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSourceContractDto) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *MetricSourceContractDto) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *MetricSourceContractDto) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

func (o MetricSourceContractDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricSourceContractDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["sql"] = o.Sql
	toSerialize["timestampColumn"] = o.TimestampColumn
	if !IsNil(o.TimestampAsDay) {
		toSerialize["timestampAsDay"] = o.TimestampAsDay
	}
	toSerialize["idTypeMapping"] = o.IdTypeMapping
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.TableName) {
		toSerialize["tableName"] = o.TableName
	}
	if !IsNil(o.CustomFieldMapping) {
		toSerialize["customFieldMapping"] = o.CustomFieldMapping
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return toSerialize, nil
}

func (o *MetricSourceContractDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"sql",
		"timestampColumn",
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

	varMetricSourceContractDto := _MetricSourceContractDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetricSourceContractDto)

	if err != nil {
		return err
	}

	*o = MetricSourceContractDto(varMetricSourceContractDto)

	return err
}

type NullableMetricSourceContractDto struct {
	value *MetricSourceContractDto
	isSet bool
}

func (v NullableMetricSourceContractDto) Get() *MetricSourceContractDto {
	return v.value
}

func (v *NullableMetricSourceContractDto) Set(val *MetricSourceContractDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricSourceContractDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricSourceContractDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricSourceContractDto(val *MetricSourceContractDto) *NullableMetricSourceContractDto {
	return &NullableMetricSourceContractDto{value: val, isSet: true}
}

func (v NullableMetricSourceContractDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricSourceContractDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


