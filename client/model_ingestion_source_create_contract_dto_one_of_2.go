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

// checks if the IngestionSourceCreateContractDtoOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionSourceCreateContractDtoOneOf2{}

// IngestionSourceCreateContractDtoOneOf2 struct for IngestionSourceCreateContractDtoOneOf2
type IngestionSourceCreateContractDtoOneOf2 struct {
	Dataset string `json:"dataset"`
	ColumnMapping *IngestionSourceCreateContractDtoOneOf2ColumnMapping `json:"column_mapping,omitempty"`
	Type string `json:"type"`
	SourceName string `json:"source_name"`
	Query *string `json:"query,omitempty"`
	UseDeltaSharing *bool `json:"use_delta_sharing,omitempty"`
	Share *string `json:"share,omitempty"`
	Schema *string `json:"schema,omitempty"`
	Table *string `json:"table,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

type _IngestionSourceCreateContractDtoOneOf2 IngestionSourceCreateContractDtoOneOf2

// NewIngestionSourceCreateContractDtoOneOf2 instantiates a new IngestionSourceCreateContractDtoOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionSourceCreateContractDtoOneOf2(dataset string, type_ string, sourceName string) *IngestionSourceCreateContractDtoOneOf2 {
	this := IngestionSourceCreateContractDtoOneOf2{}
	this.Dataset = dataset
	this.Type = type_
	this.SourceName = sourceName
	return &this
}

// NewIngestionSourceCreateContractDtoOneOf2WithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionSourceCreateContractDtoOneOf2WithDefaults() *IngestionSourceCreateContractDtoOneOf2 {
	this := IngestionSourceCreateContractDtoOneOf2{}
	return &this
}

// GetDataset returns the Dataset field value
func (o *IngestionSourceCreateContractDtoOneOf2) GetDataset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value
func (o *IngestionSourceCreateContractDtoOneOf2) SetDataset(v string) {
	o.Dataset = v
}

// GetColumnMapping returns the ColumnMapping field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetColumnMapping() IngestionSourceCreateContractDtoOneOf2ColumnMapping {
	if o == nil || IsNil(o.ColumnMapping) {
		var ret IngestionSourceCreateContractDtoOneOf2ColumnMapping
		return ret
	}
	return *o.ColumnMapping
}

// GetColumnMappingOk returns a tuple with the ColumnMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOf2ColumnMapping, bool) {
	if o == nil || IsNil(o.ColumnMapping) {
		return nil, false
	}
	return o.ColumnMapping, true
}

// HasColumnMapping returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasColumnMapping() bool {
	if o != nil && !IsNil(o.ColumnMapping) {
		return true
	}

	return false
}

// SetColumnMapping gets a reference to the given IngestionSourceCreateContractDtoOneOf2ColumnMapping and assigns it to the ColumnMapping field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetColumnMapping(v IngestionSourceCreateContractDtoOneOf2ColumnMapping) {
	o.ColumnMapping = &v
}

// GetType returns the Type field value
func (o *IngestionSourceCreateContractDtoOneOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IngestionSourceCreateContractDtoOneOf2) SetType(v string) {
	o.Type = v
}

// GetSourceName returns the SourceName field value
func (o *IngestionSourceCreateContractDtoOneOf2) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value
func (o *IngestionSourceCreateContractDtoOneOf2) SetSourceName(v string) {
	o.SourceName = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetQuery(v string) {
	o.Query = &v
}

// GetUseDeltaSharing returns the UseDeltaSharing field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetUseDeltaSharing() bool {
	if o == nil || IsNil(o.UseDeltaSharing) {
		var ret bool
		return ret
	}
	return *o.UseDeltaSharing
}

// GetUseDeltaSharingOk returns a tuple with the UseDeltaSharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetUseDeltaSharingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDeltaSharing) {
		return nil, false
	}
	return o.UseDeltaSharing, true
}

// HasUseDeltaSharing returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasUseDeltaSharing() bool {
	if o != nil && !IsNil(o.UseDeltaSharing) {
		return true
	}

	return false
}

// SetUseDeltaSharing gets a reference to the given bool and assigns it to the UseDeltaSharing field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetUseDeltaSharing(v bool) {
	o.UseDeltaSharing = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetShare() string {
	if o == nil || IsNil(o.Share) {
		var ret string
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetShareOk() (*string, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasShare() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given string and assigns it to the Share field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetShare(v string) {
	o.Share = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetTable(v string) {
	o.Table = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IngestionSourceCreateContractDtoOneOf2) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o IngestionSourceCreateContractDtoOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestionSourceCreateContractDtoOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataset"] = o.Dataset
	if !IsNil(o.ColumnMapping) {
		toSerialize["column_mapping"] = o.ColumnMapping
	}
	toSerialize["type"] = o.Type
	toSerialize["source_name"] = o.SourceName
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.UseDeltaSharing) {
		toSerialize["use_delta_sharing"] = o.UseDeltaSharing
	}
	if !IsNil(o.Share) {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *IngestionSourceCreateContractDtoOneOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataset",
		"type",
		"source_name",
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

	varIngestionSourceCreateContractDtoOneOf2 := _IngestionSourceCreateContractDtoOneOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestionSourceCreateContractDtoOneOf2)

	if err != nil {
		return err
	}

	*o = IngestionSourceCreateContractDtoOneOf2(varIngestionSourceCreateContractDtoOneOf2)

	return err
}

type NullableIngestionSourceCreateContractDtoOneOf2 struct {
	value *IngestionSourceCreateContractDtoOneOf2
	isSet bool
}

func (v NullableIngestionSourceCreateContractDtoOneOf2) Get() *IngestionSourceCreateContractDtoOneOf2 {
	return v.value
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2) Set(val *IngestionSourceCreateContractDtoOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionSourceCreateContractDtoOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionSourceCreateContractDtoOneOf2(val *IngestionSourceCreateContractDtoOneOf2) *NullableIngestionSourceCreateContractDtoOneOf2 {
	return &NullableIngestionSourceCreateContractDtoOneOf2{value: val, isSet: true}
}

func (v NullableIngestionSourceCreateContractDtoOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


