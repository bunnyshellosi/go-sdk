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

// checks if the ServiceComponentVariableEditAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceComponentVariableEditAction{}

// ServiceComponentVariableEditAction A component variable used during Bunnyshell workflows.
type ServiceComponentVariableEditAction struct {
	Value    NullableString `json:"value"`
	IsSecret NullableBool   `json:"isSecret,omitempty"`
}

// NewServiceComponentVariableEditAction instantiates a new ServiceComponentVariableEditAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceComponentVariableEditAction(value NullableString) *ServiceComponentVariableEditAction {
	this := ServiceComponentVariableEditAction{}
	this.Value = value
	return &this
}

// NewServiceComponentVariableEditActionWithDefaults instantiates a new ServiceComponentVariableEditAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceComponentVariableEditActionWithDefaults() *ServiceComponentVariableEditAction {
	this := ServiceComponentVariableEditAction{}
	return &this
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceComponentVariableEditAction) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceComponentVariableEditAction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *ServiceComponentVariableEditAction) SetValue(v string) {
	o.Value.Set(&v)
}

// GetIsSecret returns the IsSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceComponentVariableEditAction) GetIsSecret() bool {
	if o == nil || IsNil(o.IsSecret.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSecret.Get()
}

// GetIsSecretOk returns a tuple with the IsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceComponentVariableEditAction) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSecret.Get(), o.IsSecret.IsSet()
}

// HasIsSecret returns a boolean if a field has been set.
func (o *ServiceComponentVariableEditAction) HasIsSecret() bool {
	if o != nil && o.IsSecret.IsSet() {
		return true
	}

	return false
}

// SetIsSecret gets a reference to the given NullableBool and assigns it to the IsSecret field.
func (o *ServiceComponentVariableEditAction) SetIsSecret(v bool) {
	o.IsSecret.Set(&v)
}

// SetIsSecretNil sets the value for IsSecret to be an explicit nil
func (o *ServiceComponentVariableEditAction) SetIsSecretNil() {
	o.IsSecret.Set(nil)
}

// UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
func (o *ServiceComponentVariableEditAction) UnsetIsSecret() {
	o.IsSecret.Unset()
}

func (o ServiceComponentVariableEditAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceComponentVariableEditAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value.Get()
	if o.IsSecret.IsSet() {
		toSerialize["isSecret"] = o.IsSecret.Get()
	}
	return toSerialize, nil
}

type NullableServiceComponentVariableEditAction struct {
	value *ServiceComponentVariableEditAction
	isSet bool
}

func (v NullableServiceComponentVariableEditAction) Get() *ServiceComponentVariableEditAction {
	return v.value
}

func (v *NullableServiceComponentVariableEditAction) Set(val *ServiceComponentVariableEditAction) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceComponentVariableEditAction) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceComponentVariableEditAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceComponentVariableEditAction(val *ServiceComponentVariableEditAction) *NullableServiceComponentVariableEditAction {
	return &NullableServiceComponentVariableEditAction{value: val, isSet: true}
}

func (v NullableServiceComponentVariableEditAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceComponentVariableEditAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
