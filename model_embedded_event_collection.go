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

// EmbeddedEventCollection struct for EmbeddedEventCollection
type EmbeddedEventCollection struct {
	Item []EventCollection `json:"item,omitempty"`
}

// NewEmbeddedEventCollection instantiates a new EmbeddedEventCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedEventCollection() *EmbeddedEventCollection {
	this := EmbeddedEventCollection{}
	return &this
}

// NewEmbeddedEventCollectionWithDefaults instantiates a new EmbeddedEventCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedEventCollectionWithDefaults() *EmbeddedEventCollection {
	this := EmbeddedEventCollection{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *EmbeddedEventCollection) GetItem() []EventCollection {
	if o == nil || o.Item == nil {
		var ret []EventCollection
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEventCollection) GetItemOk() ([]EventCollection, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *EmbeddedEventCollection) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given []EventCollection and assigns it to the Item field.
func (o *EmbeddedEventCollection) SetItem(v []EventCollection) {
	o.Item = v
}

func (o EmbeddedEventCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedEventCollection struct {
	value *EmbeddedEventCollection
	isSet bool
}

func (v NullableEmbeddedEventCollection) Get() *EmbeddedEventCollection {
	return v.value
}

func (v *NullableEmbeddedEventCollection) Set(val *EmbeddedEventCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedEventCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedEventCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedEventCollection(val *EmbeddedEventCollection) *NullableEmbeddedEventCollection {
	return &NullableEmbeddedEventCollection{value: val, isSet: true}
}

func (v NullableEmbeddedEventCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedEventCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
