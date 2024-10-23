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

// checks if the ExternalGateDtoActiveReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGateDtoActiveReview{}

// ExternalGateDtoActiveReview struct for ExternalGateDtoActiveReview
type ExternalGateDtoActiveReview struct {
	ReviewID string `json:"reviewID"`
	ReviewStatus string `json:"reviewStatus"`
	Description string `json:"description"`
}

type _ExternalGateDtoActiveReview ExternalGateDtoActiveReview

// NewExternalGateDtoActiveReview instantiates a new ExternalGateDtoActiveReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGateDtoActiveReview(reviewID string, reviewStatus string, description string) *ExternalGateDtoActiveReview {
	this := ExternalGateDtoActiveReview{}
	this.ReviewID = reviewID
	this.ReviewStatus = reviewStatus
	this.Description = description
	return &this
}

// NewExternalGateDtoActiveReviewWithDefaults instantiates a new ExternalGateDtoActiveReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGateDtoActiveReviewWithDefaults() *ExternalGateDtoActiveReview {
	this := ExternalGateDtoActiveReview{}
	return &this
}

// GetReviewID returns the ReviewID field value
func (o *ExternalGateDtoActiveReview) GetReviewID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewID
}

// GetReviewIDOk returns a tuple with the ReviewID field value
// and a boolean to check if the value has been set.
func (o *ExternalGateDtoActiveReview) GetReviewIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewID, true
}

// SetReviewID sets field value
func (o *ExternalGateDtoActiveReview) SetReviewID(v string) {
	o.ReviewID = v
}

// GetReviewStatus returns the ReviewStatus field value
func (o *ExternalGateDtoActiveReview) GetReviewStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewStatus
}

// GetReviewStatusOk returns a tuple with the ReviewStatus field value
// and a boolean to check if the value has been set.
func (o *ExternalGateDtoActiveReview) GetReviewStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewStatus, true
}

// SetReviewStatus sets field value
func (o *ExternalGateDtoActiveReview) SetReviewStatus(v string) {
	o.ReviewStatus = v
}

// GetDescription returns the Description field value
func (o *ExternalGateDtoActiveReview) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExternalGateDtoActiveReview) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExternalGateDtoActiveReview) SetDescription(v string) {
	o.Description = v
}

func (o ExternalGateDtoActiveReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGateDtoActiveReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviewID"] = o.ReviewID
	toSerialize["reviewStatus"] = o.ReviewStatus
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *ExternalGateDtoActiveReview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviewID",
		"reviewStatus",
		"description",
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

	varExternalGateDtoActiveReview := _ExternalGateDtoActiveReview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalGateDtoActiveReview)

	if err != nil {
		return err
	}

	*o = ExternalGateDtoActiveReview(varExternalGateDtoActiveReview)

	return err
}

type NullableExternalGateDtoActiveReview struct {
	value *ExternalGateDtoActiveReview
	isSet bool
}

func (v NullableExternalGateDtoActiveReview) Get() *ExternalGateDtoActiveReview {
	return v.value
}

func (v *NullableExternalGateDtoActiveReview) Set(val *ExternalGateDtoActiveReview) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGateDtoActiveReview) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGateDtoActiveReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGateDtoActiveReview(val *ExternalGateDtoActiveReview) *NullableExternalGateDtoActiveReview {
	return &NullableExternalGateDtoActiveReview{value: val, isSet: true}
}

func (v NullableExternalGateDtoActiveReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGateDtoActiveReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


