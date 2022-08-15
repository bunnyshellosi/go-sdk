/*
API Bunnyshell Environments

Interact with Bunnyshell Platform

API version: 1.0.1
Contact: api+support@bunnyshell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// EmbeddedComponentCollection struct for EmbeddedComponentCollection
type EmbeddedComponentCollection struct {
	Item []ComponentCollection `json:"item,omitempty"`
}

// NewEmbeddedComponentCollection instantiates a new EmbeddedComponentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedComponentCollection() *EmbeddedComponentCollection {
	this := EmbeddedComponentCollection{}
	return &this
}

// NewEmbeddedComponentCollectionWithDefaults instantiates a new EmbeddedComponentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedComponentCollectionWithDefaults() *EmbeddedComponentCollection {
	this := EmbeddedComponentCollection{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *EmbeddedComponentCollection) GetItem() []ComponentCollection {
	if o == nil || o.Item == nil {
		var ret []ComponentCollection
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedComponentCollection) GetItemOk() ([]ComponentCollection, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *EmbeddedComponentCollection) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given []ComponentCollection and assigns it to the Item field.
func (o *EmbeddedComponentCollection) SetItem(v []ComponentCollection) {
	o.Item = v
}

func (o EmbeddedComponentCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedComponentCollection struct {
	value *EmbeddedComponentCollection
	isSet bool
}

func (v NullableEmbeddedComponentCollection) Get() *EmbeddedComponentCollection {
	return v.value
}

func (v *NullableEmbeddedComponentCollection) Set(val *EmbeddedComponentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedComponentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedComponentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedComponentCollection(val *EmbeddedComponentCollection) *NullableEmbeddedComponentCollection {
	return &NullableEmbeddedComponentCollection{value: val, isSet: true}
}

func (v NullableEmbeddedComponentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedComponentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
