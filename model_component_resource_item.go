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

// checks if the ComponentResourceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentResourceItem{}

// ComponentResourceItem A service component represents either an application or a group of applications as a single unit
type ComponentResourceItem struct {
	// Kubernetes resource kind.
	Kind *string `json:"kind,omitempty"`
	// Kubernetes resource name.
	Name *string `json:"name,omitempty"`
	// Kubernetes resource namespace.
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewComponentResourceItem instantiates a new ComponentResourceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentResourceItem() *ComponentResourceItem {
	this := ComponentResourceItem{}
	return &this
}

// NewComponentResourceItemWithDefaults instantiates a new ComponentResourceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentResourceItemWithDefaults() *ComponentResourceItem {
	this := ComponentResourceItem{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ComponentResourceItem) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceItem) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ComponentResourceItem) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ComponentResourceItem) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentResourceItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentResourceItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentResourceItem) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceItem) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceItem) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *ComponentResourceItem) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *ComponentResourceItem) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *ComponentResourceItem) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *ComponentResourceItem) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o ComponentResourceItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentResourceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	return toSerialize, nil
}

type NullableComponentResourceItem struct {
	value *ComponentResourceItem
	isSet bool
}

func (v NullableComponentResourceItem) Get() *ComponentResourceItem {
	return v.value
}

func (v *NullableComponentResourceItem) Set(val *ComponentResourceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentResourceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentResourceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentResourceItem(val *ComponentResourceItem) *NullableComponentResourceItem {
	return &NullableComponentResourceItem{value: val, isSet: true}
}

func (v NullableComponentResourceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentResourceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
