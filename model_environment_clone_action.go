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

// checks if the EnvironmentCloneAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentCloneAction{}

// EnvironmentCloneAction An environment holds a collection of buildable and deployable components.
type EnvironmentCloneAction struct {
	Name string `json:"name"`
}

// NewEnvironmentCloneAction instantiates a new EnvironmentCloneAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentCloneAction(name string) *EnvironmentCloneAction {
	this := EnvironmentCloneAction{}
	this.Name = name
	return &this
}

// NewEnvironmentCloneActionWithDefaults instantiates a new EnvironmentCloneAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentCloneActionWithDefaults() *EnvironmentCloneAction {
	this := EnvironmentCloneAction{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentCloneAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCloneAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentCloneAction) SetName(v string) {
	o.Name = v
}

func (o EnvironmentCloneAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentCloneAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableEnvironmentCloneAction struct {
	value *EnvironmentCloneAction
	isSet bool
}

func (v NullableEnvironmentCloneAction) Get() *EnvironmentCloneAction {
	return v.value
}

func (v *NullableEnvironmentCloneAction) Set(val *EnvironmentCloneAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentCloneAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentCloneAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentCloneAction(val *EnvironmentCloneAction) *NullableEnvironmentCloneAction {
	return &NullableEnvironmentCloneAction{value: val, isSet: true}
}

func (v NullableEnvironmentCloneAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentCloneAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
