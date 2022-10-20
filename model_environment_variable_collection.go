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

// EnvironmentVariableCollection An environment variable used during bunnyshell workflows.
type EnvironmentVariableCollection struct {
	// Environment variable identifier.
	Id *string `json:"id,omitempty"`
	// Environment variable name.
	Name *string `json:"name,omitempty"`
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
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasId() bool {
	if o != nil && o.Id != nil {
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentVariableCollection) SetName(v string) {
	o.Name = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *EnvironmentVariableCollection) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasSecret() bool {
	if o != nil && o.Secret != nil {
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
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
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
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCollection) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *EnvironmentVariableCollection) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *EnvironmentVariableCollection) SetOrganization(v string) {
	o.Organization = &v
}

func (o EnvironmentVariableCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
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