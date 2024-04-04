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

// TemplateItemVariablesSchemaInner - struct for TemplateItemVariablesSchemaInner
type TemplateItemVariablesSchemaInner struct {
	BooleanTypeItem *BooleanTypeItem
	EnumTypeItem    *EnumTypeItem
	FloatTypeItem   *FloatTypeItem
	IntegerTypeItem *IntegerTypeItem
	StringTypeItem  *StringTypeItem
}

// BooleanTypeItemAsTemplateItemVariablesSchemaInner is a convenience function that returns BooleanTypeItem wrapped in TemplateItemVariablesSchemaInner
func BooleanTypeItemAsTemplateItemVariablesSchemaInner(v *BooleanTypeItem) TemplateItemVariablesSchemaInner {
	return TemplateItemVariablesSchemaInner{
		BooleanTypeItem: v,
	}
}

// EnumTypeItemAsTemplateItemVariablesSchemaInner is a convenience function that returns EnumTypeItem wrapped in TemplateItemVariablesSchemaInner
func EnumTypeItemAsTemplateItemVariablesSchemaInner(v *EnumTypeItem) TemplateItemVariablesSchemaInner {
	return TemplateItemVariablesSchemaInner{
		EnumTypeItem: v,
	}
}

// FloatTypeItemAsTemplateItemVariablesSchemaInner is a convenience function that returns FloatTypeItem wrapped in TemplateItemVariablesSchemaInner
func FloatTypeItemAsTemplateItemVariablesSchemaInner(v *FloatTypeItem) TemplateItemVariablesSchemaInner {
	return TemplateItemVariablesSchemaInner{
		FloatTypeItem: v,
	}
}

// IntegerTypeItemAsTemplateItemVariablesSchemaInner is a convenience function that returns IntegerTypeItem wrapped in TemplateItemVariablesSchemaInner
func IntegerTypeItemAsTemplateItemVariablesSchemaInner(v *IntegerTypeItem) TemplateItemVariablesSchemaInner {
	return TemplateItemVariablesSchemaInner{
		IntegerTypeItem: v,
	}
}

// StringTypeItemAsTemplateItemVariablesSchemaInner is a convenience function that returns StringTypeItem wrapped in TemplateItemVariablesSchemaInner
func StringTypeItemAsTemplateItemVariablesSchemaInner(v *StringTypeItem) TemplateItemVariablesSchemaInner {
	return TemplateItemVariablesSchemaInner{
		StringTypeItem: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TemplateItemVariablesSchemaInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BooleanType-item'
	if jsonDict["type"] == "BooleanType-item" {
		// try to unmarshal JSON data into BooleanTypeItem
		err = json.Unmarshal(data, &dst.BooleanTypeItem)
		if err == nil {
			return nil // data stored in dst.BooleanTypeItem, return on the first match
		} else {
			dst.BooleanTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as BooleanTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EnumType-item'
	if jsonDict["type"] == "EnumType-item" {
		// try to unmarshal JSON data into EnumTypeItem
		err = json.Unmarshal(data, &dst.EnumTypeItem)
		if err == nil {
			return nil // data stored in dst.EnumTypeItem, return on the first match
		} else {
			dst.EnumTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as EnumTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FloatType-item'
	if jsonDict["type"] == "FloatType-item" {
		// try to unmarshal JSON data into FloatTypeItem
		err = json.Unmarshal(data, &dst.FloatTypeItem)
		if err == nil {
			return nil // data stored in dst.FloatTypeItem, return on the first match
		} else {
			dst.FloatTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as FloatTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IntegerType-item'
	if jsonDict["type"] == "IntegerType-item" {
		// try to unmarshal JSON data into IntegerTypeItem
		err = json.Unmarshal(data, &dst.IntegerTypeItem)
		if err == nil {
			return nil // data stored in dst.IntegerTypeItem, return on the first match
		} else {
			dst.IntegerTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as IntegerTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'StringType-item'
	if jsonDict["type"] == "StringType-item" {
		// try to unmarshal JSON data into StringTypeItem
		err = json.Unmarshal(data, &dst.StringTypeItem)
		if err == nil {
			return nil // data stored in dst.StringTypeItem, return on the first match
		} else {
			dst.StringTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as StringTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'bool'
	if jsonDict["type"] == "bool" {
		// try to unmarshal JSON data into BooleanTypeItem
		err = json.Unmarshal(data, &dst.BooleanTypeItem)
		if err == nil {
			return nil // data stored in dst.BooleanTypeItem, return on the first match
		} else {
			dst.BooleanTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as BooleanTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'enum'
	if jsonDict["type"] == "enum" {
		// try to unmarshal JSON data into EnumTypeItem
		err = json.Unmarshal(data, &dst.EnumTypeItem)
		if err == nil {
			return nil // data stored in dst.EnumTypeItem, return on the first match
		} else {
			dst.EnumTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as EnumTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'float'
	if jsonDict["type"] == "float" {
		// try to unmarshal JSON data into FloatTypeItem
		err = json.Unmarshal(data, &dst.FloatTypeItem)
		if err == nil {
			return nil // data stored in dst.FloatTypeItem, return on the first match
		} else {
			dst.FloatTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as FloatTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'int'
	if jsonDict["type"] == "int" {
		// try to unmarshal JSON data into IntegerTypeItem
		err = json.Unmarshal(data, &dst.IntegerTypeItem)
		if err == nil {
			return nil // data stored in dst.IntegerTypeItem, return on the first match
		} else {
			dst.IntegerTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as IntegerTypeItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'string'
	if jsonDict["type"] == "string" {
		// try to unmarshal JSON data into StringTypeItem
		err = json.Unmarshal(data, &dst.StringTypeItem)
		if err == nil {
			return nil // data stored in dst.StringTypeItem, return on the first match
		} else {
			dst.StringTypeItem = nil
			return fmt.Errorf("failed to unmarshal TemplateItemVariablesSchemaInner as StringTypeItem: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TemplateItemVariablesSchemaInner) MarshalJSON() ([]byte, error) {
	if src.BooleanTypeItem != nil {
		return json.Marshal(&src.BooleanTypeItem)
	}

	if src.EnumTypeItem != nil {
		return json.Marshal(&src.EnumTypeItem)
	}

	if src.FloatTypeItem != nil {
		return json.Marshal(&src.FloatTypeItem)
	}

	if src.IntegerTypeItem != nil {
		return json.Marshal(&src.IntegerTypeItem)
	}

	if src.StringTypeItem != nil {
		return json.Marshal(&src.StringTypeItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TemplateItemVariablesSchemaInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BooleanTypeItem != nil {
		return obj.BooleanTypeItem
	}

	if obj.EnumTypeItem != nil {
		return obj.EnumTypeItem
	}

	if obj.FloatTypeItem != nil {
		return obj.FloatTypeItem
	}

	if obj.IntegerTypeItem != nil {
		return obj.IntegerTypeItem
	}

	if obj.StringTypeItem != nil {
		return obj.StringTypeItem
	}

	// all schemas are nil
	return nil
}

type NullableTemplateItemVariablesSchemaInner struct {
	value *TemplateItemVariablesSchemaInner
	isSet bool
}

func (v NullableTemplateItemVariablesSchemaInner) Get() *TemplateItemVariablesSchemaInner {
	return v.value
}

func (v *NullableTemplateItemVariablesSchemaInner) Set(val *TemplateItemVariablesSchemaInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateItemVariablesSchemaInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateItemVariablesSchemaInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateItemVariablesSchemaInner(val *TemplateItemVariablesSchemaInner) *NullableTemplateItemVariablesSchemaInner {
	return &NullableTemplateItemVariablesSchemaInner{value: val, isSet: true}
}

func (v NullableTemplateItemVariablesSchemaInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateItemVariablesSchemaInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
