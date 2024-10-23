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

// checks if the IngestionRunDataContractDtoGranularHistoryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionRunDataContractDtoGranularHistoryInner{}

// IngestionRunDataContractDtoGranularHistoryInner struct for IngestionRunDataContractDtoGranularHistoryInner
type IngestionRunDataContractDtoGranularHistoryInner struct {
	Source string `json:"source"`
	LatestSourceStatus string `json:"latestSourceStatus"`
	StatusByDate []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner `json:"statusByDate"`
}

type _IngestionRunDataContractDtoGranularHistoryInner IngestionRunDataContractDtoGranularHistoryInner

// NewIngestionRunDataContractDtoGranularHistoryInner instantiates a new IngestionRunDataContractDtoGranularHistoryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionRunDataContractDtoGranularHistoryInner(source string, latestSourceStatus string, statusByDate []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner) *IngestionRunDataContractDtoGranularHistoryInner {
	this := IngestionRunDataContractDtoGranularHistoryInner{}
	this.Source = source
	this.LatestSourceStatus = latestSourceStatus
	this.StatusByDate = statusByDate
	return &this
}

// NewIngestionRunDataContractDtoGranularHistoryInnerWithDefaults instantiates a new IngestionRunDataContractDtoGranularHistoryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionRunDataContractDtoGranularHistoryInnerWithDefaults() *IngestionRunDataContractDtoGranularHistoryInner {
	this := IngestionRunDataContractDtoGranularHistoryInner{}
	return &this
}

// GetSource returns the Source field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) SetSource(v string) {
	o.Source = v
}

// GetLatestSourceStatus returns the LatestSourceStatus field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetLatestSourceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestSourceStatus
}

// GetLatestSourceStatusOk returns a tuple with the LatestSourceStatus field value
// and a boolean to check if the value has been set.
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetLatestSourceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestSourceStatus, true
}

// SetLatestSourceStatus sets field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) SetLatestSourceStatus(v string) {
	o.LatestSourceStatus = v
}

// GetStatusByDate returns the StatusByDate field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetStatusByDate() []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner {
	if o == nil {
		var ret []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner
		return ret
	}

	return o.StatusByDate
}

// GetStatusByDateOk returns a tuple with the StatusByDate field value
// and a boolean to check if the value has been set.
func (o *IngestionRunDataContractDtoGranularHistoryInner) GetStatusByDateOk() ([]IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusByDate, true
}

// SetStatusByDate sets field value
func (o *IngestionRunDataContractDtoGranularHistoryInner) SetStatusByDate(v []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner) {
	o.StatusByDate = v
}

func (o IngestionRunDataContractDtoGranularHistoryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestionRunDataContractDtoGranularHistoryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["latestSourceStatus"] = o.LatestSourceStatus
	toSerialize["statusByDate"] = o.StatusByDate
	return toSerialize, nil
}

func (o *IngestionRunDataContractDtoGranularHistoryInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"latestSourceStatus",
		"statusByDate",
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

	varIngestionRunDataContractDtoGranularHistoryInner := _IngestionRunDataContractDtoGranularHistoryInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestionRunDataContractDtoGranularHistoryInner)

	if err != nil {
		return err
	}

	*o = IngestionRunDataContractDtoGranularHistoryInner(varIngestionRunDataContractDtoGranularHistoryInner)

	return err
}

type NullableIngestionRunDataContractDtoGranularHistoryInner struct {
	value *IngestionRunDataContractDtoGranularHistoryInner
	isSet bool
}

func (v NullableIngestionRunDataContractDtoGranularHistoryInner) Get() *IngestionRunDataContractDtoGranularHistoryInner {
	return v.value
}

func (v *NullableIngestionRunDataContractDtoGranularHistoryInner) Set(val *IngestionRunDataContractDtoGranularHistoryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionRunDataContractDtoGranularHistoryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionRunDataContractDtoGranularHistoryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionRunDataContractDtoGranularHistoryInner(val *IngestionRunDataContractDtoGranularHistoryInner) *NullableIngestionRunDataContractDtoGranularHistoryInner {
	return &NullableIngestionRunDataContractDtoGranularHistoryInner{value: val, isSet: true}
}

func (v NullableIngestionRunDataContractDtoGranularHistoryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionRunDataContractDtoGranularHistoryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


