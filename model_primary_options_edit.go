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

// checks if the PrimaryOptionsEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimaryOptionsEdit{}

// PrimaryOptionsEdit struct for PrimaryOptionsEdit
type PrimaryOptionsEdit struct {
	WhitelistEnabled NullableBool   `json:"whitelistEnabled,omitempty"`
	BranchWhitelist  NullableString `json:"branchWhitelist,omitempty"`
}

// NewPrimaryOptionsEdit instantiates a new PrimaryOptionsEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimaryOptionsEdit() *PrimaryOptionsEdit {
	this := PrimaryOptionsEdit{}
	return &this
}

// NewPrimaryOptionsEditWithDefaults instantiates a new PrimaryOptionsEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimaryOptionsEditWithDefaults() *PrimaryOptionsEdit {
	this := PrimaryOptionsEdit{}
	return &this
}

// GetWhitelistEnabled returns the WhitelistEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrimaryOptionsEdit) GetWhitelistEnabled() bool {
	if o == nil || IsNil(o.WhitelistEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.WhitelistEnabled.Get()
}

// GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrimaryOptionsEdit) GetWhitelistEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhitelistEnabled.Get(), o.WhitelistEnabled.IsSet()
}

// HasWhitelistEnabled returns a boolean if a field has been set.
func (o *PrimaryOptionsEdit) HasWhitelistEnabled() bool {
	if o != nil && o.WhitelistEnabled.IsSet() {
		return true
	}

	return false
}

// SetWhitelistEnabled gets a reference to the given NullableBool and assigns it to the WhitelistEnabled field.
func (o *PrimaryOptionsEdit) SetWhitelistEnabled(v bool) {
	o.WhitelistEnabled.Set(&v)
}

// SetWhitelistEnabledNil sets the value for WhitelistEnabled to be an explicit nil
func (o *PrimaryOptionsEdit) SetWhitelistEnabledNil() {
	o.WhitelistEnabled.Set(nil)
}

// UnsetWhitelistEnabled ensures that no value is present for WhitelistEnabled, not even an explicit nil
func (o *PrimaryOptionsEdit) UnsetWhitelistEnabled() {
	o.WhitelistEnabled.Unset()
}

// GetBranchWhitelist returns the BranchWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrimaryOptionsEdit) GetBranchWhitelist() string {
	if o == nil || IsNil(o.BranchWhitelist.Get()) {
		var ret string
		return ret
	}
	return *o.BranchWhitelist.Get()
}

// GetBranchWhitelistOk returns a tuple with the BranchWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrimaryOptionsEdit) GetBranchWhitelistOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchWhitelist.Get(), o.BranchWhitelist.IsSet()
}

// HasBranchWhitelist returns a boolean if a field has been set.
func (o *PrimaryOptionsEdit) HasBranchWhitelist() bool {
	if o != nil && o.BranchWhitelist.IsSet() {
		return true
	}

	return false
}

// SetBranchWhitelist gets a reference to the given NullableString and assigns it to the BranchWhitelist field.
func (o *PrimaryOptionsEdit) SetBranchWhitelist(v string) {
	o.BranchWhitelist.Set(&v)
}

// SetBranchWhitelistNil sets the value for BranchWhitelist to be an explicit nil
func (o *PrimaryOptionsEdit) SetBranchWhitelistNil() {
	o.BranchWhitelist.Set(nil)
}

// UnsetBranchWhitelist ensures that no value is present for BranchWhitelist, not even an explicit nil
func (o *PrimaryOptionsEdit) UnsetBranchWhitelist() {
	o.BranchWhitelist.Unset()
}

func (o PrimaryOptionsEdit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrimaryOptionsEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WhitelistEnabled.IsSet() {
		toSerialize["whitelistEnabled"] = o.WhitelistEnabled.Get()
	}
	if o.BranchWhitelist.IsSet() {
		toSerialize["branchWhitelist"] = o.BranchWhitelist.Get()
	}
	return toSerialize, nil
}

type NullablePrimaryOptionsEdit struct {
	value *PrimaryOptionsEdit
	isSet bool
}

func (v NullablePrimaryOptionsEdit) Get() *PrimaryOptionsEdit {
	return v.value
}

func (v *NullablePrimaryOptionsEdit) Set(val *PrimaryOptionsEdit) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimaryOptionsEdit) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimaryOptionsEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimaryOptionsEdit(val *PrimaryOptionsEdit) *NullablePrimaryOptionsEdit {
	return &NullablePrimaryOptionsEdit{value: val, isSet: true}
}

func (v NullablePrimaryOptionsEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrimaryOptionsEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
