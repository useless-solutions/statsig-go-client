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

// LayerContractDtoParametersInnerDefaultValue - The default value for this parameter, which must match the specified type.
type LayerContractDtoParametersInnerDefaultValue struct {
	ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner *[]LayerContractDtoParametersInnerDefaultValueOneOfInner
	Bool *bool
	Float32 *float32
	MapmapOfStringAny *map[string]interface{}
	String *string
}

// []LayerContractDtoParametersInnerDefaultValueOneOfInnerAsLayerContractDtoParametersInnerDefaultValue is a convenience function that returns []LayerContractDtoParametersInnerDefaultValueOneOfInner wrapped in LayerContractDtoParametersInnerDefaultValue
func ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInnerAsLayerContractDtoParametersInnerDefaultValue(v *[]LayerContractDtoParametersInnerDefaultValueOneOfInner) LayerContractDtoParametersInnerDefaultValue {
	return LayerContractDtoParametersInnerDefaultValue{
		ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner: v,
	}
}

// boolAsLayerContractDtoParametersInnerDefaultValue is a convenience function that returns bool wrapped in LayerContractDtoParametersInnerDefaultValue
func BoolAsLayerContractDtoParametersInnerDefaultValue(v *bool) LayerContractDtoParametersInnerDefaultValue {
	return LayerContractDtoParametersInnerDefaultValue{
		Bool: v,
	}
}

// float32AsLayerContractDtoParametersInnerDefaultValue is a convenience function that returns float32 wrapped in LayerContractDtoParametersInnerDefaultValue
func Float32AsLayerContractDtoParametersInnerDefaultValue(v *float32) LayerContractDtoParametersInnerDefaultValue {
	return LayerContractDtoParametersInnerDefaultValue{
		Float32: v,
	}
}

// map[string]interface{}AsLayerContractDtoParametersInnerDefaultValue is a convenience function that returns map[string]interface{} wrapped in LayerContractDtoParametersInnerDefaultValue
func MapmapOfStringAnyAsLayerContractDtoParametersInnerDefaultValue(v *map[string]interface{}) LayerContractDtoParametersInnerDefaultValue {
	return LayerContractDtoParametersInnerDefaultValue{
		MapmapOfStringAny: v,
	}
}

// stringAsLayerContractDtoParametersInnerDefaultValue is a convenience function that returns string wrapped in LayerContractDtoParametersInnerDefaultValue
func StringAsLayerContractDtoParametersInnerDefaultValue(v *string) LayerContractDtoParametersInnerDefaultValue {
	return LayerContractDtoParametersInnerDefaultValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LayerContractDtoParametersInnerDefaultValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner)
	if err == nil {
		jsonArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner, _ := json.Marshal(dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner)
		if string(jsonArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner) == "{}" { // empty struct
			dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner); err != nil {
				dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
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
		dst.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.MapmapOfStringAny = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LayerContractDtoParametersInnerDefaultValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LayerContractDtoParametersInnerDefaultValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LayerContractDtoParametersInnerDefaultValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner != nil {
		return json.Marshal(&src.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LayerContractDtoParametersInnerDefaultValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner != nil {
		return obj.ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableLayerContractDtoParametersInnerDefaultValue struct {
	value *LayerContractDtoParametersInnerDefaultValue
	isSet bool
}

func (v NullableLayerContractDtoParametersInnerDefaultValue) Get() *LayerContractDtoParametersInnerDefaultValue {
	return v.value
}

func (v *NullableLayerContractDtoParametersInnerDefaultValue) Set(val *LayerContractDtoParametersInnerDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLayerContractDtoParametersInnerDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLayerContractDtoParametersInnerDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayerContractDtoParametersInnerDefaultValue(val *LayerContractDtoParametersInnerDefaultValue) *NullableLayerContractDtoParametersInnerDefaultValue {
	return &NullableLayerContractDtoParametersInnerDefaultValue{value: val, isSet: true}
}

func (v NullableLayerContractDtoParametersInnerDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayerContractDtoParametersInnerDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


