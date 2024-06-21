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

// TemplateValidateActionSource - struct for TemplateValidateActionSource
type TemplateValidateActionSource struct {
	ValidateSourceGit    *ValidateSourceGit
	ValidateSourceString *ValidateSourceString
}

// ValidateSourceGitAsTemplateValidateActionSource is a convenience function that returns ValidateSourceGit wrapped in TemplateValidateActionSource
func ValidateSourceGitAsTemplateValidateActionSource(v *ValidateSourceGit) TemplateValidateActionSource {
	return TemplateValidateActionSource{
		ValidateSourceGit: v,
	}
}

// ValidateSourceStringAsTemplateValidateActionSource is a convenience function that returns ValidateSourceString wrapped in TemplateValidateActionSource
func ValidateSourceStringAsTemplateValidateActionSource(v *ValidateSourceString) TemplateValidateActionSource {
	return TemplateValidateActionSource{
		ValidateSourceString: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TemplateValidateActionSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ValidateSourceGit
	err = newStrictDecoder(data).Decode(&dst.ValidateSourceGit)
	if err == nil {
		jsonValidateSourceGit, _ := json.Marshal(dst.ValidateSourceGit)
		if string(jsonValidateSourceGit) == "{}" { // empty struct
			dst.ValidateSourceGit = nil
		} else {
			match++
		}
	} else {
		dst.ValidateSourceGit = nil
	}

	// try to unmarshal data into ValidateSourceString
	err = newStrictDecoder(data).Decode(&dst.ValidateSourceString)
	if err == nil {
		jsonValidateSourceString, _ := json.Marshal(dst.ValidateSourceString)
		if string(jsonValidateSourceString) == "{}" { // empty struct
			dst.ValidateSourceString = nil
		} else {
			match++
		}
	} else {
		dst.ValidateSourceString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ValidateSourceGit = nil
		dst.ValidateSourceString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TemplateValidateActionSource)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TemplateValidateActionSource)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TemplateValidateActionSource) MarshalJSON() ([]byte, error) {
	if src.ValidateSourceGit != nil {
		return json.Marshal(&src.ValidateSourceGit)
	}

	if src.ValidateSourceString != nil {
		return json.Marshal(&src.ValidateSourceString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TemplateValidateActionSource) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ValidateSourceGit != nil {
		return obj.ValidateSourceGit
	}

	if obj.ValidateSourceString != nil {
		return obj.ValidateSourceString
	}

	// all schemas are nil
	return nil
}

type NullableTemplateValidateActionSource struct {
	value *TemplateValidateActionSource
	isSet bool
}

func (v NullableTemplateValidateActionSource) Get() *TemplateValidateActionSource {
	return v.value
}

func (v *NullableTemplateValidateActionSource) Set(val *TemplateValidateActionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateValidateActionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateValidateActionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateValidateActionSource(val *TemplateValidateActionSource) *NullableTemplateValidateActionSource {
	return &NullableTemplateValidateActionSource{value: val, isSet: true}
}

func (v NullableTemplateValidateActionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateValidateActionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
