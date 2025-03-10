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

// checks if the IngestionSourceCreateContractDtoOneOf2ColumnMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionSourceCreateContractDtoOneOf2ColumnMapping{}

// IngestionSourceCreateContractDtoOneOf2ColumnMapping struct for IngestionSourceCreateContractDtoOneOf2ColumnMapping
type IngestionSourceCreateContractDtoOneOf2ColumnMapping struct {
	// Unique identifier for the experiment.
	Experiment string `json:"experiment"`
	// Unique identifier for the experiment groups.
	GroupId string `json:"group_id"`
	// The unique user identifier this exposure is for. This might not necessarily be a single column for userID - it could be spread across multiple columns for deviceID etc.
	UnitId *string `json:"unit_id,omitempty"`
	// Unix timestamp in seconds of the event (ex. 1729692720)
	Timestamp string `json:"timestamp"`
	Ids map[string]string `json:"ids"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	MetadataObject *string `json:"metadata_object,omitempty"`
	EventValue *string `json:"event_value,omitempty"`
}

type _IngestionSourceCreateContractDtoOneOf2ColumnMapping IngestionSourceCreateContractDtoOneOf2ColumnMapping

// NewIngestionSourceCreateContractDtoOneOf2ColumnMapping instantiates a new IngestionSourceCreateContractDtoOneOf2ColumnMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionSourceCreateContractDtoOneOf2ColumnMapping(experiment string, groupId string, timestamp string, ids map[string]string) *IngestionSourceCreateContractDtoOneOf2ColumnMapping {
	this := IngestionSourceCreateContractDtoOneOf2ColumnMapping{}
	this.Experiment = experiment
	this.GroupId = groupId
	this.Timestamp = timestamp
	this.Ids = ids
	var metadataObject string = "null"
	this.MetadataObject = &metadataObject
	var eventValue string = ""
	this.EventValue = &eventValue
	return &this
}

// NewIngestionSourceCreateContractDtoOneOf2ColumnMappingWithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf2ColumnMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionSourceCreateContractDtoOneOf2ColumnMappingWithDefaults() *IngestionSourceCreateContractDtoOneOf2ColumnMapping {
	this := IngestionSourceCreateContractDtoOneOf2ColumnMapping{}
	var metadataObject string = "null"
	this.MetadataObject = &metadataObject
	var eventValue string = ""
	this.EventValue = &eventValue
	return &this
}

// GetExperiment returns the Experiment field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetExperiment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Experiment
}

// GetExperimentOk returns a tuple with the Experiment field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetExperimentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Experiment, true
}

// SetExperiment sets field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetExperiment(v string) {
	o.Experiment = v
}

// GetGroupId returns the GroupId field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetGroupId(v string) {
	o.GroupId = v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetUnitId() string {
	if o == nil || IsNil(o.UnitId) {
		var ret string
		return ret
	}
	return *o.UnitId
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetUnitIdOk() (*string, bool) {
	if o == nil || IsNil(o.UnitId) {
		return nil, false
	}
	return o.UnitId, true
}

// HasUnitId returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasUnitId() bool {
	if o != nil && !IsNil(o.UnitId) {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given string and assigns it to the UnitId field.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetUnitId(v string) {
	o.UnitId = &v
}

// GetTimestamp returns the Timestamp field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetIds returns the Ids field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetIdsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetIds(v map[string]string) {
	o.Ids = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMetadataObject returns the MetadataObject field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataObject() string {
	if o == nil || IsNil(o.MetadataObject) {
		var ret string
		return ret
	}
	return *o.MetadataObject
}

// GetMetadataObjectOk returns a tuple with the MetadataObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataObjectOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataObject) {
		return nil, false
	}
	return o.MetadataObject, true
}

// HasMetadataObject returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasMetadataObject() bool {
	if o != nil && !IsNil(o.MetadataObject) {
		return true
	}

	return false
}

// SetMetadataObject gets a reference to the given string and assigns it to the MetadataObject field.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetMetadataObject(v string) {
	o.MetadataObject = &v
}

// GetEventValue returns the EventValue field value if set, zero value otherwise.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetEventValue() string {
	if o == nil || IsNil(o.EventValue) {
		var ret string
		return ret
	}
	return *o.EventValue
}

// GetEventValueOk returns a tuple with the EventValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetEventValueOk() (*string, bool) {
	if o == nil || IsNil(o.EventValue) {
		return nil, false
	}
	return o.EventValue, true
}

// HasEventValue returns a boolean if a field has been set.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasEventValue() bool {
	if o != nil && !IsNil(o.EventValue) {
		return true
	}

	return false
}

// SetEventValue gets a reference to the given string and assigns it to the EventValue field.
func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetEventValue(v string) {
	o.EventValue = &v
}

func (o IngestionSourceCreateContractDtoOneOf2ColumnMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestionSourceCreateContractDtoOneOf2ColumnMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["experiment"] = o.Experiment
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.UnitId) {
		toSerialize["unit_id"] = o.UnitId
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["ids"] = o.Ids
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetadataObject) {
		toSerialize["metadata_object"] = o.MetadataObject
	}
	if !IsNil(o.EventValue) {
		toSerialize["event_value"] = o.EventValue
	}
	return toSerialize, nil
}

func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"experiment",
		"group_id",
		"timestamp",
		"ids",
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

	varIngestionSourceCreateContractDtoOneOf2ColumnMapping := _IngestionSourceCreateContractDtoOneOf2ColumnMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestionSourceCreateContractDtoOneOf2ColumnMapping)

	if err != nil {
		return err
	}

	*o = IngestionSourceCreateContractDtoOneOf2ColumnMapping(varIngestionSourceCreateContractDtoOneOf2ColumnMapping)

	return err
}

type NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping struct {
	value *IngestionSourceCreateContractDtoOneOf2ColumnMapping
	isSet bool
}

func (v NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) Get() *IngestionSourceCreateContractDtoOneOf2ColumnMapping {
	return v.value
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) Set(val *IngestionSourceCreateContractDtoOneOf2ColumnMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionSourceCreateContractDtoOneOf2ColumnMapping(val *IngestionSourceCreateContractDtoOneOf2ColumnMapping) *NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping {
	return &NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping{value: val, isSet: true}
}

func (v NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionSourceCreateContractDtoOneOf2ColumnMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


