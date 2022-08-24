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

// ContextWrapperKubeConfigRead struct for ContextWrapperKubeConfigRead
type ContextWrapperKubeConfigRead struct {
	Name    string                 `json:"name"`
	Context *ContextKubeConfigRead `json:"context,omitempty"`
}

// NewContextWrapperKubeConfigRead instantiates a new ContextWrapperKubeConfigRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextWrapperKubeConfigRead(name string) *ContextWrapperKubeConfigRead {
	this := ContextWrapperKubeConfigRead{}
	this.Name = name
	return &this
}

// NewContextWrapperKubeConfigReadWithDefaults instantiates a new ContextWrapperKubeConfigRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWrapperKubeConfigReadWithDefaults() *ContextWrapperKubeConfigRead {
	this := ContextWrapperKubeConfigRead{}
	return &this
}

// GetName returns the Name field value
func (o *ContextWrapperKubeConfigRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContextWrapperKubeConfigRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContextWrapperKubeConfigRead) SetName(v string) {
	o.Name = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ContextWrapperKubeConfigRead) GetContext() ContextKubeConfigRead {
	if o == nil || o.Context == nil {
		var ret ContextKubeConfigRead
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextWrapperKubeConfigRead) GetContextOk() (*ContextKubeConfigRead, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ContextWrapperKubeConfigRead) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given ContextKubeConfigRead and assigns it to the Context field.
func (o *ContextWrapperKubeConfigRead) SetContext(v ContextKubeConfigRead) {
	o.Context = &v
}

func (o ContextWrapperKubeConfigRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableContextWrapperKubeConfigRead struct {
	value *ContextWrapperKubeConfigRead
	isSet bool
}

func (v NullableContextWrapperKubeConfigRead) Get() *ContextWrapperKubeConfigRead {
	return v.value
}

func (v *NullableContextWrapperKubeConfigRead) Set(val *ContextWrapperKubeConfigRead) {
	v.value = val
	v.isSet = true
}

func (v NullableContextWrapperKubeConfigRead) IsSet() bool {
	return v.isSet
}

func (v *NullableContextWrapperKubeConfigRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextWrapperKubeConfigRead(val *ContextWrapperKubeConfigRead) *NullableContextWrapperKubeConfigRead {
	return &NullableContextWrapperKubeConfigRead{value: val, isSet: true}
}

func (v NullableContextWrapperKubeConfigRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextWrapperKubeConfigRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
