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

// checks if the EnvironmentVariableCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentVariableCollection{}

// EnvironmentVariableCollection An environment variable used during Bunnyshell workflows.
type EnvironmentVariableCollection struct {
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

// NewEnvironmentVariableCollection instantiates a new EnvironmentVariableCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableCollection() *EnvironmentVariableCollection {
	this := EnvironmentVariableCollection{}
	return &this
}

// NewEnvironmentVariableCollectionWithDefaults instantiates a new EnvironmentVariableCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableCollectionWithDefaults() *EnvironmentVariableCollection {
	this := EnvironmentVariableCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentVariableCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentVariableCollection) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariableCollection) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariableCollection) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EnvironmentVariableCollection) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *EnvironmentVariableCollection) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EnvironmentVariableCollection) UnsetValue() {
	o.Value.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetSecret() bool {
	if o == nil || IsNil(o.Secret) {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *EnvironmentVariableCollection) SetSecret(v bool) {
	o.Secret = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *EnvironmentVariableCollection) SetEnvironment(v string) {
	o.Environment = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *EnvironmentVariableCollection) SetOrganization(v string) {
	o.Organization = &v
}

func (o EnvironmentVariableCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentVariableCollection) ToMap() (map[string]interface{}, error) {
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

type NullableEnvironmentVariableCollection struct {
	value *EnvironmentVariableCollection
	isSet bool
}

func (v NullableEnvironmentVariableCollection) Get() *EnvironmentVariableCollection {
	return v.value
}

func (v *NullableEnvironmentVariableCollection) Set(val *EnvironmentVariableCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableCollection(val *EnvironmentVariableCollection) *NullableEnvironmentVariableCollection {
	return &NullableEnvironmentVariableCollection{value: val, isSet: true}
}

func (v NullableEnvironmentVariableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
