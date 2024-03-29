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

// checks if the EnvironmentVariableCreateAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariableCreateAction{}

// EnvironmentVariableCreateAction An environment variable used during Bunnyshell workflows.
type EnvironmentVariableCreateAction struct {
	Name        string       `json:"name"`
	Value       string       `json:"value"`
	IsSecret    NullableBool `json:"isSecret,omitempty"`
	Environment string       `json:"environment"`
}

// NewEnvironmentVariableCreateAction instantiates a new EnvironmentVariableCreateAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableCreateAction(name string, value string, environment string) *EnvironmentVariableCreateAction {
	this := EnvironmentVariableCreateAction{}
	this.Name = name
	this.Value = value
	this.Environment = environment
	return &this
}

// NewEnvironmentVariableCreateActionWithDefaults instantiates a new EnvironmentVariableCreateAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableCreateActionWithDefaults() *EnvironmentVariableCreateAction {
	this := EnvironmentVariableCreateAction{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentVariableCreateAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreateAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentVariableCreateAction) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariableCreateAction) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreateAction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariableCreateAction) SetValue(v string) {
	o.Value = v
}

// GetIsSecret returns the IsSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariableCreateAction) GetIsSecret() bool {
	if o == nil || IsNil(o.IsSecret.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSecret.Get()
}

// GetIsSecretOk returns a tuple with the IsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariableCreateAction) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSecret.Get(), o.IsSecret.IsSet()
}

// HasIsSecret returns a boolean if a field has been set.
func (o *EnvironmentVariableCreateAction) HasIsSecret() bool {
	if o != nil && o.IsSecret.IsSet() {
		return true
	}

	return false
}

// SetIsSecret gets a reference to the given NullableBool and assigns it to the IsSecret field.
func (o *EnvironmentVariableCreateAction) SetIsSecret(v bool) {
	o.IsSecret.Set(&v)
}

// SetIsSecretNil sets the value for IsSecret to be an explicit nil
func (o *EnvironmentVariableCreateAction) SetIsSecretNil() {
	o.IsSecret.Set(nil)
}

// UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
func (o *EnvironmentVariableCreateAction) UnsetIsSecret() {
	o.IsSecret.Unset()
}

// GetEnvironment returns the Environment field value
func (o *EnvironmentVariableCreateAction) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreateAction) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *EnvironmentVariableCreateAction) SetEnvironment(v string) {
	o.Environment = v
}

func (o EnvironmentVariableCreateAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariableCreateAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if o.IsSecret.IsSet() {
		toSerialize["isSecret"] = o.IsSecret.Get()
	}
	toSerialize["environment"] = o.Environment
	return toSerialize, nil
}

type NullableEnvironmentVariableCreateAction struct {
	value *EnvironmentVariableCreateAction
	isSet bool
}

func (v NullableEnvironmentVariableCreateAction) Get() *EnvironmentVariableCreateAction {
	return v.value
}

func (v *NullableEnvironmentVariableCreateAction) Set(val *EnvironmentVariableCreateAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableCreateAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableCreateAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableCreateAction(val *EnvironmentVariableCreateAction) *NullableEnvironmentVariableCreateAction {
	return &NullableEnvironmentVariableCreateAction{value: val, isSet: true}
}

func (v NullableEnvironmentVariableCreateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableCreateAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
