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

// checks if the SimpleResourceConfigItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleResourceConfigItem{}

// SimpleResourceConfigItem struct for SimpleResourceConfigItem
type SimpleResourceConfigItem struct {
	Simple *bool `json:"simple,omitempty"`
	// A list of container configs.
	Containers *map[string]ContainerConfigItem `json:"containers,omitempty"`
}

// NewSimpleResourceConfigItem instantiates a new SimpleResourceConfigItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleResourceConfigItem() *SimpleResourceConfigItem {
	this := SimpleResourceConfigItem{}
	return &this
}

// NewSimpleResourceConfigItemWithDefaults instantiates a new SimpleResourceConfigItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleResourceConfigItemWithDefaults() *SimpleResourceConfigItem {
	this := SimpleResourceConfigItem{}
	return &this
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *SimpleResourceConfigItem) GetSimple() bool {
	if o == nil || IsNil(o.Simple) {
		var ret bool
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleResourceConfigItem) GetSimpleOk() (*bool, bool) {
	if o == nil || IsNil(o.Simple) {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *SimpleResourceConfigItem) HasSimple() bool {
	if o != nil && !IsNil(o.Simple) {
		return true
	}

	return false
}

// SetSimple gets a reference to the given bool and assigns it to the Simple field.
func (o *SimpleResourceConfigItem) SetSimple(v bool) {
	o.Simple = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *SimpleResourceConfigItem) GetContainers() map[string]ContainerConfigItem {
	if o == nil || IsNil(o.Containers) {
		var ret map[string]ContainerConfigItem
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleResourceConfigItem) GetContainersOk() (*map[string]ContainerConfigItem, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *SimpleResourceConfigItem) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given map[string]ContainerConfigItem and assigns it to the Containers field.
func (o *SimpleResourceConfigItem) SetContainers(v map[string]ContainerConfigItem) {
	o.Containers = &v
}

func (o SimpleResourceConfigItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleResourceConfigItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Simple) {
		toSerialize["simple"] = o.Simple
	}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableSimpleResourceConfigItem struct {
	value *SimpleResourceConfigItem
	isSet bool
}

func (v NullableSimpleResourceConfigItem) Get() *SimpleResourceConfigItem {
	return v.value
}

func (v *NullableSimpleResourceConfigItem) Set(val *SimpleResourceConfigItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleResourceConfigItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleResourceConfigItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleResourceConfigItem(val *SimpleResourceConfigItem) *NullableSimpleResourceConfigItem {
	return &NullableSimpleResourceConfigItem{value: val, isSet: true}
}

func (v NullableSimpleResourceConfigItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleResourceConfigItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
