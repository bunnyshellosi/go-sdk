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

// checks if the BooleanTypeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanTypeItem{}

// BooleanTypeItem struct for BooleanTypeItem
type BooleanTypeItem struct {
	DefaultValue NullableBooleanValueItem `json:"defaultValue,omitempty"`
	// A variable used within the template.
	Name *string `json:"name,omitempty"`
	// The variable description
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewBooleanTypeItem instantiates a new BooleanTypeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanTypeItem() *BooleanTypeItem {
	this := BooleanTypeItem{}
	return &this
}

// NewBooleanTypeItemWithDefaults instantiates a new BooleanTypeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanTypeItemWithDefaults() *BooleanTypeItem {
	this := BooleanTypeItem{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BooleanTypeItem) GetDefaultValue() BooleanValueItem {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret BooleanValueItem
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BooleanTypeItem) GetDefaultValueOk() (*BooleanValueItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BooleanTypeItem) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableBooleanValueItem and assigns it to the DefaultValue field.
func (o *BooleanTypeItem) SetDefaultValue(v BooleanValueItem) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *BooleanTypeItem) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *BooleanTypeItem) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BooleanTypeItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTypeItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BooleanTypeItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BooleanTypeItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BooleanTypeItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTypeItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BooleanTypeItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BooleanTypeItem) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BooleanTypeItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanTypeItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BooleanTypeItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BooleanTypeItem) SetType(v string) {
	o.Type = &v
}

func (o BooleanTypeItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanTypeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValue.IsSet() {
		toSerialize["defaultValue"] = o.DefaultValue.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBooleanTypeItem struct {
	value *BooleanTypeItem
	isSet bool
}

func (v NullableBooleanTypeItem) Get() *BooleanTypeItem {
	return v.value
}

func (v *NullableBooleanTypeItem) Set(val *BooleanTypeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanTypeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanTypeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanTypeItem(val *BooleanTypeItem) *NullableBooleanTypeItem {
	return &NullableBooleanTypeItem{value: val, isSet: true}
}

func (v NullableBooleanTypeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanTypeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}