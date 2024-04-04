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
)

// checks if the EnumTypeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumTypeItem{}

// EnumTypeItem struct for EnumTypeItem
type EnumTypeItem struct {
	DefaultValue *EnumTypeItemDefaultValue `json:"defaultValue,omitempty"`
	// The available values for the enum.
	Values []EnumTypeItemValuesInner `json:"values,omitempty"`
	// A variable used within the template.
	Name *string `json:"name,omitempty"`
	// The variable description
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewEnumTypeItem instantiates a new EnumTypeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumTypeItem() *EnumTypeItem {
	this := EnumTypeItem{}
	return &this
}

// NewEnumTypeItemWithDefaults instantiates a new EnumTypeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumTypeItemWithDefaults() *EnumTypeItem {
	this := EnumTypeItem{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *EnumTypeItem) GetDefaultValue() EnumTypeItemDefaultValue {
	if o == nil || IsNil(o.DefaultValue) {
		var ret EnumTypeItemDefaultValue
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTypeItem) GetDefaultValueOk() (*EnumTypeItemDefaultValue, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *EnumTypeItem) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given EnumTypeItemDefaultValue and assigns it to the DefaultValue field.
func (o *EnumTypeItem) SetDefaultValue(v EnumTypeItemDefaultValue) {
	o.DefaultValue = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EnumTypeItem) GetValues() []EnumTypeItemValuesInner {
	if o == nil || IsNil(o.Values) {
		var ret []EnumTypeItemValuesInner
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTypeItem) GetValuesOk() ([]EnumTypeItemValuesInner, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EnumTypeItem) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EnumTypeItemValuesInner and assigns it to the Values field.
func (o *EnumTypeItem) SetValues(v []EnumTypeItemValuesInner) {
	o.Values = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnumTypeItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTypeItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnumTypeItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnumTypeItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnumTypeItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTypeItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnumTypeItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnumTypeItem) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnumTypeItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTypeItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnumTypeItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnumTypeItem) SetType(v string) {
	o.Type = &v
}

func (o EnumTypeItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumTypeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEnumTypeItem struct {
	value *EnumTypeItem
	isSet bool
}

func (v NullableEnumTypeItem) Get() *EnumTypeItem {
	return v.value
}

func (v *NullableEnumTypeItem) Set(val *EnumTypeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTypeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTypeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTypeItem(val *EnumTypeItem) *NullableEnumTypeItem {
	return &NullableEnumTypeItem{value: val, isSet: true}
}

func (v NullableEnumTypeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTypeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
