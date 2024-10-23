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

// checks if the IngestionCreateDatabricksConnectionContractDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestionCreateDatabricksConnectionContractDto{}

// IngestionCreateDatabricksConnectionContractDto struct for IngestionCreateDatabricksConnectionContractDto
type IngestionCreateDatabricksConnectionContractDto struct {
	Token string `json:"token"`
	Host string `json:"host"`
	Path string `json:"path"`
	DeltaSharingCredentials *string `json:"deltaSharingCredentials,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

type _IngestionCreateDatabricksConnectionContractDto IngestionCreateDatabricksConnectionContractDto

// NewIngestionCreateDatabricksConnectionContractDto instantiates a new IngestionCreateDatabricksConnectionContractDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestionCreateDatabricksConnectionContractDto(token string, host string, path string) *IngestionCreateDatabricksConnectionContractDto {
	this := IngestionCreateDatabricksConnectionContractDto{}
	this.Token = token
	this.Host = host
	this.Path = path
	return &this
}

// NewIngestionCreateDatabricksConnectionContractDtoWithDefaults instantiates a new IngestionCreateDatabricksConnectionContractDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestionCreateDatabricksConnectionContractDtoWithDefaults() *IngestionCreateDatabricksConnectionContractDto {
	this := IngestionCreateDatabricksConnectionContractDto{}
	return &this
}

// GetToken returns the Token field value
func (o *IngestionCreateDatabricksConnectionContractDto) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *IngestionCreateDatabricksConnectionContractDto) SetToken(v string) {
	o.Token = v
}

// GetHost returns the Host field value
func (o *IngestionCreateDatabricksConnectionContractDto) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *IngestionCreateDatabricksConnectionContractDto) SetHost(v string) {
	o.Host = v
}

// GetPath returns the Path field value
func (o *IngestionCreateDatabricksConnectionContractDto) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *IngestionCreateDatabricksConnectionContractDto) SetPath(v string) {
	o.Path = v
}

// GetDeltaSharingCredentials returns the DeltaSharingCredentials field value if set, zero value otherwise.
func (o *IngestionCreateDatabricksConnectionContractDto) GetDeltaSharingCredentials() string {
	if o == nil || IsNil(o.DeltaSharingCredentials) {
		var ret string
		return ret
	}
	return *o.DeltaSharingCredentials
}

// GetDeltaSharingCredentialsOk returns a tuple with the DeltaSharingCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) GetDeltaSharingCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.DeltaSharingCredentials) {
		return nil, false
	}
	return o.DeltaSharingCredentials, true
}

// HasDeltaSharingCredentials returns a boolean if a field has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) HasDeltaSharingCredentials() bool {
	if o != nil && !IsNil(o.DeltaSharingCredentials) {
		return true
	}

	return false
}

// SetDeltaSharingCredentials gets a reference to the given string and assigns it to the DeltaSharingCredentials field.
func (o *IngestionCreateDatabricksConnectionContractDto) SetDeltaSharingCredentials(v string) {
	o.DeltaSharingCredentials = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *IngestionCreateDatabricksConnectionContractDto) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *IngestionCreateDatabricksConnectionContractDto) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *IngestionCreateDatabricksConnectionContractDto) SetVerified(v bool) {
	o.Verified = &v
}

func (o IngestionCreateDatabricksConnectionContractDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestionCreateDatabricksConnectionContractDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["host"] = o.Host
	toSerialize["path"] = o.Path
	if !IsNil(o.DeltaSharingCredentials) {
		toSerialize["deltaSharingCredentials"] = o.DeltaSharingCredentials
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

func (o *IngestionCreateDatabricksConnectionContractDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"host",
		"path",
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

	varIngestionCreateDatabricksConnectionContractDto := _IngestionCreateDatabricksConnectionContractDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestionCreateDatabricksConnectionContractDto)

	if err != nil {
		return err
	}

	*o = IngestionCreateDatabricksConnectionContractDto(varIngestionCreateDatabricksConnectionContractDto)

	return err
}

type NullableIngestionCreateDatabricksConnectionContractDto struct {
	value *IngestionCreateDatabricksConnectionContractDto
	isSet bool
}

func (v NullableIngestionCreateDatabricksConnectionContractDto) Get() *IngestionCreateDatabricksConnectionContractDto {
	return v.value
}

func (v *NullableIngestionCreateDatabricksConnectionContractDto) Set(val *IngestionCreateDatabricksConnectionContractDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionCreateDatabricksConnectionContractDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionCreateDatabricksConnectionContractDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionCreateDatabricksConnectionContractDto(val *IngestionCreateDatabricksConnectionContractDto) *NullableIngestionCreateDatabricksConnectionContractDto {
	return &NullableIngestionCreateDatabricksConnectionContractDto{value: val, isSet: true}
}

func (v NullableIngestionCreateDatabricksConnectionContractDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionCreateDatabricksConnectionContractDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

