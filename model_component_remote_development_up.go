/*
API Bunnyshell Environments

Interact with Bunnyshell Platform

API version: 1.1.0
Contact: api+support@bunnyshell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// ComponentRemoteDevelopmentUp A service component represents either an application or a group of applications as a single unit
type ComponentRemoteDevelopmentUp struct {
	Container NullableString `json:"container,omitempty"`
}

// NewComponentRemoteDevelopmentUp instantiates a new ComponentRemoteDevelopmentUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentRemoteDevelopmentUp() *ComponentRemoteDevelopmentUp {
	this := ComponentRemoteDevelopmentUp{}
	return &this
}

// NewComponentRemoteDevelopmentUpWithDefaults instantiates a new ComponentRemoteDevelopmentUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentRemoteDevelopmentUpWithDefaults() *ComponentRemoteDevelopmentUp {
	this := ComponentRemoteDevelopmentUp{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentRemoteDevelopmentUp) GetContainer() string {
	if o == nil || o.Container.Get() == nil {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentRemoteDevelopmentUp) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *ComponentRemoteDevelopmentUp) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *ComponentRemoteDevelopmentUp) SetContainer(v string) {
	o.Container.Set(&v)
}

// SetContainerNil sets the value for Container to be an explicit nil
func (o *ComponentRemoteDevelopmentUp) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *ComponentRemoteDevelopmentUp) UnsetContainer() {
	o.Container.Unset()
}

func (o ComponentRemoteDevelopmentUp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Container.IsSet() {
		toSerialize["container"] = o.Container.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableComponentRemoteDevelopmentUp struct {
	value *ComponentRemoteDevelopmentUp
	isSet bool
}

func (v NullableComponentRemoteDevelopmentUp) Get() *ComponentRemoteDevelopmentUp {
	return v.value
}

func (v *NullableComponentRemoteDevelopmentUp) Set(val *ComponentRemoteDevelopmentUp) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentRemoteDevelopmentUp) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentRemoteDevelopmentUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentRemoteDevelopmentUp(val *ComponentRemoteDevelopmentUp) *NullableComponentRemoteDevelopmentUp {
	return &NullableComponentRemoteDevelopmentUp{value: val, isSet: true}
}

func (v NullableComponentRemoteDevelopmentUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentRemoteDevelopmentUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
