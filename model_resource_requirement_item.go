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

// ResourceRequirementItem struct for ResourceRequirementItem
type ResourceRequirementItem struct {
	Requests NullableResourceRequirementItemRequests `json:"requests,omitempty"`
	Limits   NullableResourceRequirementItemLimits   `json:"limits,omitempty"`
}

// NewResourceRequirementItem instantiates a new ResourceRequirementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRequirementItem() *ResourceRequirementItem {
	this := ResourceRequirementItem{}
	return &this
}

// NewResourceRequirementItemWithDefaults instantiates a new ResourceRequirementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRequirementItemWithDefaults() *ResourceRequirementItem {
	this := ResourceRequirementItem{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceRequirementItem) GetRequests() ResourceRequirementItemRequests {
	if o == nil || o.Requests.Get() == nil {
		var ret ResourceRequirementItemRequests
		return ret
	}
	return *o.Requests.Get()
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceRequirementItem) GetRequestsOk() (*ResourceRequirementItemRequests, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests.Get(), o.Requests.IsSet()
}

// HasRequests returns a boolean if a field has been set.
func (o *ResourceRequirementItem) HasRequests() bool {
	if o != nil && o.Requests.IsSet() {
		return true
	}

	return false
}

// SetRequests gets a reference to the given NullableResourceRequirementItemRequests and assigns it to the Requests field.
func (o *ResourceRequirementItem) SetRequests(v ResourceRequirementItemRequests) {
	o.Requests.Set(&v)
}

// SetRequestsNil sets the value for Requests to be an explicit nil
func (o *ResourceRequirementItem) SetRequestsNil() {
	o.Requests.Set(nil)
}

// UnsetRequests ensures that no value is present for Requests, not even an explicit nil
func (o *ResourceRequirementItem) UnsetRequests() {
	o.Requests.Unset()
}

// GetLimits returns the Limits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceRequirementItem) GetLimits() ResourceRequirementItemLimits {
	if o == nil || o.Limits.Get() == nil {
		var ret ResourceRequirementItemLimits
		return ret
	}
	return *o.Limits.Get()
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceRequirementItem) GetLimitsOk() (*ResourceRequirementItemLimits, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limits.Get(), o.Limits.IsSet()
}

// HasLimits returns a boolean if a field has been set.
func (o *ResourceRequirementItem) HasLimits() bool {
	if o != nil && o.Limits.IsSet() {
		return true
	}

	return false
}

// SetLimits gets a reference to the given NullableResourceRequirementItemLimits and assigns it to the Limits field.
func (o *ResourceRequirementItem) SetLimits(v ResourceRequirementItemLimits) {
	o.Limits.Set(&v)
}

// SetLimitsNil sets the value for Limits to be an explicit nil
func (o *ResourceRequirementItem) SetLimitsNil() {
	o.Limits.Set(nil)
}

// UnsetLimits ensures that no value is present for Limits, not even an explicit nil
func (o *ResourceRequirementItem) UnsetLimits() {
	o.Limits.Unset()
}

func (o ResourceRequirementItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requests.IsSet() {
		toSerialize["requests"] = o.Requests.Get()
	}
	if o.Limits.IsSet() {
		toSerialize["limits"] = o.Limits.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceRequirementItem struct {
	value *ResourceRequirementItem
	isSet bool
}

func (v NullableResourceRequirementItem) Get() *ResourceRequirementItem {
	return v.value
}

func (v *NullableResourceRequirementItem) Set(val *ResourceRequirementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRequirementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRequirementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRequirementItem(val *ResourceRequirementItem) *NullableResourceRequirementItem {
	return &NullableResourceRequirementItem{value: val, isSet: true}
}

func (v NullableResourceRequirementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRequirementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
