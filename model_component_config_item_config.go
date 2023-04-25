/*
API Bunnyshell Environments

Interact with Bunnyshell Platform

API version: 1.1.0
Contact: osi+support@bunnyshell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// ComponentConfigItemConfig - A list of container configs.
type ComponentConfigItemConfig struct {
	ArrayOfExtendedResourceConfigItem *[]ExtendedResourceConfigItem
	ArrayOfSimpleResourceConfigItem   *[]SimpleResourceConfigItem
}

// []ExtendedResourceConfigItemAsComponentConfigItemConfig is a convenience function that returns []ExtendedResourceConfigItem wrapped in ComponentConfigItemConfig
func ArrayOfExtendedResourceConfigItemAsComponentConfigItemConfig(v *[]ExtendedResourceConfigItem) ComponentConfigItemConfig {
	return ComponentConfigItemConfig{
		ArrayOfExtendedResourceConfigItem: v,
	}
}

// []SimpleResourceConfigItemAsComponentConfigItemConfig is a convenience function that returns []SimpleResourceConfigItem wrapped in ComponentConfigItemConfig
func ArrayOfSimpleResourceConfigItemAsComponentConfigItemConfig(v *[]SimpleResourceConfigItem) ComponentConfigItemConfig {
	return ComponentConfigItemConfig{
		ArrayOfSimpleResourceConfigItem: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComponentConfigItemConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfExtendedResourceConfigItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfExtendedResourceConfigItem)
	if err == nil {
		jsonArrayOfExtendedResourceConfigItem, _ := json.Marshal(dst.ArrayOfExtendedResourceConfigItem)
		if string(jsonArrayOfExtendedResourceConfigItem) == "{}" { // empty struct
			dst.ArrayOfExtendedResourceConfigItem = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfExtendedResourceConfigItem = nil
	}

	// try to unmarshal data into ArrayOfSimpleResourceConfigItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSimpleResourceConfigItem)
	if err == nil {
		jsonArrayOfSimpleResourceConfigItem, _ := json.Marshal(dst.ArrayOfSimpleResourceConfigItem)
		if string(jsonArrayOfSimpleResourceConfigItem) == "{}" { // empty struct
			dst.ArrayOfSimpleResourceConfigItem = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfSimpleResourceConfigItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfExtendedResourceConfigItem = nil
		dst.ArrayOfSimpleResourceConfigItem = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ComponentConfigItemConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ComponentConfigItemConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComponentConfigItemConfig) MarshalJSON() ([]byte, error) {
	if src.ArrayOfExtendedResourceConfigItem != nil {
		return json.Marshal(&src.ArrayOfExtendedResourceConfigItem)
	}

	if src.ArrayOfSimpleResourceConfigItem != nil {
		return json.Marshal(&src.ArrayOfSimpleResourceConfigItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComponentConfigItemConfig) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfExtendedResourceConfigItem != nil {
		return obj.ArrayOfExtendedResourceConfigItem
	}

	if obj.ArrayOfSimpleResourceConfigItem != nil {
		return obj.ArrayOfSimpleResourceConfigItem
	}

	// all schemas are nil
	return nil
}

type NullableComponentConfigItemConfig struct {
	value *ComponentConfigItemConfig
	isSet bool
}

func (v NullableComponentConfigItemConfig) Get() *ComponentConfigItemConfig {
	return v.value
}

func (v *NullableComponentConfigItemConfig) Set(val *ComponentConfigItemConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentConfigItemConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentConfigItemConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentConfigItemConfig(val *ComponentConfigItemConfig) *NullableComponentConfigItemConfig {
	return &NullableComponentConfigItemConfig{value: val, isSet: true}
}

func (v NullableComponentConfigItemConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentConfigItemConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
