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

// checks if the EnvironItemEditAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironItemEditAction{}

// EnvironItemEditAction An environment variable used during Bunnyshell workflows.
type EnvironItemEditAction struct {
	Value    NullableString `json:"value,omitempty"`
	IsSecret NullableBool   `json:"isSecret,omitempty"`
}

// NewEnvironItemEditAction instantiates a new EnvironItemEditAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironItemEditAction() *EnvironItemEditAction {
	this := EnvironItemEditAction{}
	return &this
}

// NewEnvironItemEditActionWithDefaults instantiates a new EnvironItemEditAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironItemEditActionWithDefaults() *EnvironItemEditAction {
	this := EnvironItemEditAction{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironItemEditAction) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironItemEditAction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironItemEditAction) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EnvironItemEditAction) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *EnvironItemEditAction) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EnvironItemEditAction) UnsetValue() {
	o.Value.Unset()
}

// GetIsSecret returns the IsSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironItemEditAction) GetIsSecret() bool {
	if o == nil || IsNil(o.IsSecret.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSecret.Get()
}

// GetIsSecretOk returns a tuple with the IsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironItemEditAction) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSecret.Get(), o.IsSecret.IsSet()
}

// HasIsSecret returns a boolean if a field has been set.
func (o *EnvironItemEditAction) HasIsSecret() bool {
	if o != nil && o.IsSecret.IsSet() {
		return true
	}

	return false
}

// SetIsSecret gets a reference to the given NullableBool and assigns it to the IsSecret field.
func (o *EnvironItemEditAction) SetIsSecret(v bool) {
	o.IsSecret.Set(&v)
}

// SetIsSecretNil sets the value for IsSecret to be an explicit nil
func (o *EnvironItemEditAction) SetIsSecretNil() {
	o.IsSecret.Set(nil)
}

// UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
func (o *EnvironItemEditAction) UnsetIsSecret() {
	o.IsSecret.Unset()
}

func (o EnvironItemEditAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironItemEditAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsSecret.IsSet() {
		toSerialize["isSecret"] = o.IsSecret.Get()
	}
	return toSerialize, nil
}

type NullableEnvironItemEditAction struct {
	value *EnvironItemEditAction
	isSet bool
}

func (v NullableEnvironItemEditAction) Get() *EnvironItemEditAction {
	return v.value
}

func (v *NullableEnvironItemEditAction) Set(val *EnvironItemEditAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironItemEditAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironItemEditAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironItemEditAction(val *EnvironItemEditAction) *NullableEnvironItemEditAction {
	return &NullableEnvironItemEditAction{value: val, isSet: true}
}

func (v NullableEnvironItemEditAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironItemEditAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
