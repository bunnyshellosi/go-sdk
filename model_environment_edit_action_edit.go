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

// EnvironmentEditActionEdit - Environment edit by type
type EnvironmentEditActionEdit struct {
	Ephemeral *Ephemeral
	Primary   *Primary
}

// EphemeralAsEnvironmentEditActionEdit is a convenience function that returns Ephemeral wrapped in EnvironmentEditActionEdit
func EphemeralAsEnvironmentEditActionEdit(v *Ephemeral) EnvironmentEditActionEdit {
	return EnvironmentEditActionEdit{
		Ephemeral: v,
	}
}

// PrimaryAsEnvironmentEditActionEdit is a convenience function that returns Primary wrapped in EnvironmentEditActionEdit
func PrimaryAsEnvironmentEditActionEdit(v *Primary) EnvironmentEditActionEdit {
	return EnvironmentEditActionEdit{
		Primary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EnvironmentEditActionEdit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ephemeral
	err = newStrictDecoder(data).Decode(&dst.Ephemeral)
	if err == nil {
		jsonEphemeral, _ := json.Marshal(dst.Ephemeral)
		if string(jsonEphemeral) == "{}" { // empty struct
			dst.Ephemeral = nil
		} else {
			match++
		}
	} else {
		dst.Ephemeral = nil
	}

	// try to unmarshal data into Primary
	err = newStrictDecoder(data).Decode(&dst.Primary)
	if err == nil {
		jsonPrimary, _ := json.Marshal(dst.Primary)
		if string(jsonPrimary) == "{}" { // empty struct
			dst.Primary = nil
		} else {
			match++
		}
	} else {
		dst.Primary = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ephemeral = nil
		dst.Primary = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(EnvironmentEditActionEdit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EnvironmentEditActionEdit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EnvironmentEditActionEdit) MarshalJSON() ([]byte, error) {
	if src.Ephemeral != nil {
		return json.Marshal(&src.Ephemeral)
	}

	if src.Primary != nil {
		return json.Marshal(&src.Primary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EnvironmentEditActionEdit) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ephemeral != nil {
		return obj.Ephemeral
	}

	if obj.Primary != nil {
		return obj.Primary
	}

	// all schemas are nil
	return nil
}

type NullableEnvironmentEditActionEdit struct {
	value *EnvironmentEditActionEdit
	isSet bool
}

func (v NullableEnvironmentEditActionEdit) Get() *EnvironmentEditActionEdit {
	return v.value
}

func (v *NullableEnvironmentEditActionEdit) Set(val *EnvironmentEditActionEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentEditActionEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentEditActionEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentEditActionEdit(val *EnvironmentEditActionEdit) *NullableEnvironmentEditActionEdit {
	return &NullableEnvironmentEditActionEdit{value: val, isSet: true}
}

func (v NullableEnvironmentEditActionEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentEditActionEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}