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

// checks if the IntegerValueItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegerValueItem{}

// IntegerValueItem struct for IntegerValueItem
type IntegerValueItem struct {
	Value *int32  `json:"value,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// NewIntegerValueItem instantiates a new IntegerValueItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerValueItem() *IntegerValueItem {
	this := IntegerValueItem{}
	return &this
}

// NewIntegerValueItemWithDefaults instantiates a new IntegerValueItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerValueItemWithDefaults() *IntegerValueItem {
	this := IntegerValueItem{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IntegerValueItem) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerValueItem) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IntegerValueItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *IntegerValueItem) SetValue(v int32) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegerValueItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerValueItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegerValueItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IntegerValueItem) SetType(v string) {
	o.Type = &v
}

func (o IntegerValueItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegerValueItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIntegerValueItem struct {
	value *IntegerValueItem
	isSet bool
}

func (v NullableIntegerValueItem) Get() *IntegerValueItem {
	return v.value
}

func (v *NullableIntegerValueItem) Set(val *IntegerValueItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerValueItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerValueItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerValueItem(val *IntegerValueItem) *NullableIntegerValueItem {
	return &NullableIntegerValueItem{value: val, isSet: true}
}

func (v NullableIntegerValueItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerValueItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}