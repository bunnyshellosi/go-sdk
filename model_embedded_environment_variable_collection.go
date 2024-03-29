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

// checks if the EmbeddedEnvironmentVariableCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedEnvironmentVariableCollection{}

// EmbeddedEnvironmentVariableCollection struct for EmbeddedEnvironmentVariableCollection
type EmbeddedEnvironmentVariableCollection struct {
	Item []EnvironmentVariableCollection `json:"item,omitempty"`
}

// NewEmbeddedEnvironmentVariableCollection instantiates a new EmbeddedEnvironmentVariableCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedEnvironmentVariableCollection() *EmbeddedEnvironmentVariableCollection {
	this := EmbeddedEnvironmentVariableCollection{}
	return &this
}

// NewEmbeddedEnvironmentVariableCollectionWithDefaults instantiates a new EmbeddedEnvironmentVariableCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedEnvironmentVariableCollectionWithDefaults() *EmbeddedEnvironmentVariableCollection {
	this := EmbeddedEnvironmentVariableCollection{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *EmbeddedEnvironmentVariableCollection) GetItem() []EnvironmentVariableCollection {
	if o == nil || IsNil(o.Item) {
		var ret []EnvironmentVariableCollection
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEnvironmentVariableCollection) GetItemOk() ([]EnvironmentVariableCollection, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *EmbeddedEnvironmentVariableCollection) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []EnvironmentVariableCollection and assigns it to the Item field.
func (o *EmbeddedEnvironmentVariableCollection) SetItem(v []EnvironmentVariableCollection) {
	o.Item = v
}

func (o EmbeddedEnvironmentVariableCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedEnvironmentVariableCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableEmbeddedEnvironmentVariableCollection struct {
	value *EmbeddedEnvironmentVariableCollection
	isSet bool
}

func (v NullableEmbeddedEnvironmentVariableCollection) Get() *EmbeddedEnvironmentVariableCollection {
	return v.value
}

func (v *NullableEmbeddedEnvironmentVariableCollection) Set(val *EmbeddedEnvironmentVariableCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedEnvironmentVariableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedEnvironmentVariableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedEnvironmentVariableCollection(val *EmbeddedEnvironmentVariableCollection) *NullableEmbeddedEnvironmentVariableCollection {
	return &NullableEmbeddedEnvironmentVariableCollection{value: val, isSet: true}
}

func (v NullableEmbeddedEnvironmentVariableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedEnvironmentVariableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
