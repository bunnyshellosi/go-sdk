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

// ResourceListItem struct for ResourceListItem
type ResourceListItem struct {
	// The CPU resources for the container.
	Cpu NullableString `json:"cpu,omitempty"`
	// The Memory resources for the container.
	Memory NullableString `json:"memory,omitempty"`
}

// NewResourceListItem instantiates a new ResourceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceListItem() *ResourceListItem {
	this := ResourceListItem{}
	return &this
}

// NewResourceListItemWithDefaults instantiates a new ResourceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceListItemWithDefaults() *ResourceListItem {
	this := ResourceListItem{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceListItem) GetCpu() string {
	if o == nil || o.Cpu.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceListItem) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *ResourceListItem) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableString and assigns it to the Cpu field.
func (o *ResourceListItem) SetCpu(v string) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *ResourceListItem) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *ResourceListItem) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceListItem) GetMemory() string {
	if o == nil || o.Memory.Get() == nil {
		var ret string
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceListItem) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *ResourceListItem) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableString and assigns it to the Memory field.
func (o *ResourceListItem) SetMemory(v string) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *ResourceListItem) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *ResourceListItem) UnsetMemory() {
	o.Memory.Unset()
}

func (o ResourceListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceListItem struct {
	value *ResourceListItem
	isSet bool
}

func (v NullableResourceListItem) Get() *ResourceListItem {
	return v.value
}

func (v *NullableResourceListItem) Set(val *ResourceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceListItem(val *ResourceListItem) *NullableResourceListItem {
	return &NullableResourceListItem{value: val, isSet: true}
}

func (v NullableResourceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}