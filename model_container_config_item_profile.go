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

// ContainerConfigItemProfile - A profile config or a profile name for local usage with rdev.yaml
type ContainerConfigItemProfile struct {
	ProfileItem *ProfileItem
	String      *string
}

// ProfileItemAsContainerConfigItemProfile is a convenience function that returns ProfileItem wrapped in ContainerConfigItemProfile
func ProfileItemAsContainerConfigItemProfile(v *ProfileItem) ContainerConfigItemProfile {
	return ContainerConfigItemProfile{
		ProfileItem: v,
	}
}

// stringAsContainerConfigItemProfile is a convenience function that returns string wrapped in ContainerConfigItemProfile
func StringAsContainerConfigItemProfile(v *string) ContainerConfigItemProfile {
	return ContainerConfigItemProfile{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContainerConfigItemProfile) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProfileItem
	err = newStrictDecoder(data).Decode(&dst.ProfileItem)
	if err == nil {
		jsonProfileItem, _ := json.Marshal(dst.ProfileItem)
		if string(jsonProfileItem) == "{}" { // empty struct
			dst.ProfileItem = nil
		} else {
			match++
		}
	} else {
		dst.ProfileItem = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ProfileItem = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContainerConfigItemProfile)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContainerConfigItemProfile)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContainerConfigItemProfile) MarshalJSON() ([]byte, error) {
	if src.ProfileItem != nil {
		return json.Marshal(&src.ProfileItem)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContainerConfigItemProfile) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ProfileItem != nil {
		return obj.ProfileItem
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContainerConfigItemProfile struct {
	value *ContainerConfigItemProfile
	isSet bool
}

func (v NullableContainerConfigItemProfile) Get() *ContainerConfigItemProfile {
	return v.value
}

func (v *NullableContainerConfigItemProfile) Set(val *ContainerConfigItemProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerConfigItemProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerConfigItemProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerConfigItemProfile(val *ContainerConfigItemProfile) *NullableContainerConfigItemProfile {
	return &NullableContainerConfigItemProfile{value: val, isSet: true}
}

func (v NullableContainerConfigItemProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerConfigItemProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
