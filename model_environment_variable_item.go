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

// checks if the EnvironmentVariableItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariableItem{}

// EnvironmentVariableItem An environment variable used during bunnyshell workflows.
type EnvironmentVariableItem struct {
	// Environment variable identifier.
	Id *string `json:"id,omitempty"`
	// Environment variable name.
	Name *string `json:"name,omitempty"`
	// Environment variable value.
	Value NullableString `json:"value,omitempty"`
	// Environment variable marked as secret.
	Secret *bool `json:"secret,omitempty"`
	// Environment identifier.
	Environment *string `json:"environment,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewEnvironmentVariableItem instantiates a new EnvironmentVariableItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableItem() *EnvironmentVariableItem {
	this := EnvironmentVariableItem{}
	return &this
}

// NewEnvironmentVariableItemWithDefaults instantiates a new EnvironmentVariableItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableItemWithDefaults() *EnvironmentVariableItem {
	this := EnvironmentVariableItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentVariableItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentVariableItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentVariableItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentVariableItem) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariableItem) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariableItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EnvironmentVariableItem) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *EnvironmentVariableItem) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EnvironmentVariableItem) UnsetValue() {
	o.Value.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *EnvironmentVariableItem) GetSecret() bool {
	if o == nil || IsNil(o.Secret) {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableItem) GetSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *EnvironmentVariableItem) SetSecret(v bool) {
	o.Secret = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EnvironmentVariableItem) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableItem) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *EnvironmentVariableItem) SetEnvironment(v string) {
	o.Environment = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *EnvironmentVariableItem) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableItem) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *EnvironmentVariableItem) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *EnvironmentVariableItem) SetOrganization(v string) {
	o.Organization = &v
}

func (o EnvironmentVariableItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariableItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableEnvironmentVariableItem struct {
	value *EnvironmentVariableItem
	isSet bool
}

func (v NullableEnvironmentVariableItem) Get() *EnvironmentVariableItem {
	return v.value
}

func (v *NullableEnvironmentVariableItem) Set(val *EnvironmentVariableItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableItem(val *EnvironmentVariableItem) *NullableEnvironmentVariableItem {
	return &NullableEnvironmentVariableItem{value: val, isSet: true}
}

func (v NullableEnvironmentVariableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
