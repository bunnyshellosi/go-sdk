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

// checks if the ProjectEditSettingsAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEditSettingsAction{}

// ProjectEditSettingsAction A project holds multiple environments and shared secrets and settings.
type ProjectEditSettingsAction struct {
	Name   NullableString `json:"name"`
	Labels NullableEdit   `json:"labels,omitempty"`
}

// NewProjectEditSettingsAction instantiates a new ProjectEditSettingsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEditSettingsAction(name NullableString) *ProjectEditSettingsAction {
	this := ProjectEditSettingsAction{}
	this.Name = name
	return &this
}

// NewProjectEditSettingsActionWithDefaults instantiates a new ProjectEditSettingsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEditSettingsActionWithDefaults() *ProjectEditSettingsAction {
	this := ProjectEditSettingsAction{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectEditSettingsAction) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditSettingsAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ProjectEditSettingsAction) SetName(v string) {
	o.Name.Set(&v)
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEditSettingsAction) GetLabels() Edit {
	if o == nil || IsNil(o.Labels.Get()) {
		var ret Edit
		return ret
	}
	return *o.Labels.Get()
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditSettingsAction) GetLabelsOk() (*Edit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels.Get(), o.Labels.IsSet()
}

// HasLabels returns a boolean if a field has been set.
func (o *ProjectEditSettingsAction) HasLabels() bool {
	if o != nil && o.Labels.IsSet() {
		return true
	}

	return false
}

// SetLabels gets a reference to the given NullableEdit and assigns it to the Labels field.
func (o *ProjectEditSettingsAction) SetLabels(v Edit) {
	o.Labels.Set(&v)
}

// SetLabelsNil sets the value for Labels to be an explicit nil
func (o *ProjectEditSettingsAction) SetLabelsNil() {
	o.Labels.Set(nil)
}

// UnsetLabels ensures that no value is present for Labels, not even an explicit nil
func (o *ProjectEditSettingsAction) UnsetLabels() {
	o.Labels.Unset()
}

func (o ProjectEditSettingsAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEditSettingsAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if o.Labels.IsSet() {
		toSerialize["labels"] = o.Labels.Get()
	}
	return toSerialize, nil
}

type NullableProjectEditSettingsAction struct {
	value *ProjectEditSettingsAction
	isSet bool
}

func (v NullableProjectEditSettingsAction) Get() *ProjectEditSettingsAction {
	return v.value
}

func (v *NullableProjectEditSettingsAction) Set(val *ProjectEditSettingsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEditSettingsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEditSettingsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEditSettingsAction(val *ProjectEditSettingsAction) *NullableProjectEditSettingsAction {
	return &NullableProjectEditSettingsAction{value: val, isSet: true}
}

func (v NullableProjectEditSettingsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEditSettingsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}