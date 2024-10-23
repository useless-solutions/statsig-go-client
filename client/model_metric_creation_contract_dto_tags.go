/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// MetricCreationContractDtoTags - Tags associated with the metric for categorization and searchability.
type MetricCreationContractDtoTags struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsMetricCreationContractDtoTags is a convenience function that returns []string wrapped in MetricCreationContractDtoTags
func ArrayOfStringAsMetricCreationContractDtoTags(v *[]string) MetricCreationContractDtoTags {
	return MetricCreationContractDtoTags{
		ArrayOfString: v,
	}
}

// stringAsMetricCreationContractDtoTags is a convenience function that returns string wrapped in MetricCreationContractDtoTags
func StringAsMetricCreationContractDtoTags(v *string) MetricCreationContractDtoTags {
	return MetricCreationContractDtoTags{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MetricCreationContractDtoTags) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MetricCreationContractDtoTags)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MetricCreationContractDtoTags)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MetricCreationContractDtoTags) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MetricCreationContractDtoTags) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableMetricCreationContractDtoTags struct {
	value *MetricCreationContractDtoTags
	isSet bool
}

func (v NullableMetricCreationContractDtoTags) Get() *MetricCreationContractDtoTags {
	return v.value
}

func (v *NullableMetricCreationContractDtoTags) Set(val *MetricCreationContractDtoTags) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricCreationContractDtoTags) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricCreationContractDtoTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricCreationContractDtoTags(val *MetricCreationContractDtoTags) *NullableMetricCreationContractDtoTags {
	return &NullableMetricCreationContractDtoTags{value: val, isSet: true}
}

func (v NullableMetricCreationContractDtoTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricCreationContractDtoTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

