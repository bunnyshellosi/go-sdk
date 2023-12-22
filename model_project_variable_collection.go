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

// checks if the ProjectVariableCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectVariableCollection{}

// ProjectVariableCollection A project variable used during Bunnyshell workflows.
type ProjectVariableCollection struct {
	// Project variable identifier.
	Id *string `json:"id,omitempty"`
	// Project variable name.
	Name *string `json:"name,omitempty"`
	// Project variable value.
	Value NullableString `json:"value,omitempty"`
	// Project variable marked as secret.
	Secret *bool `json:"secret,omitempty"`
	// Project identifier.
	Project *string `json:"project,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewProjectVariableCollection instantiates a new ProjectVariableCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectVariableCollection() *ProjectVariableCollection {
	this := ProjectVariableCollection{}
	return &this
}

// NewProjectVariableCollectionWithDefaults instantiates a new ProjectVariableCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectVariableCollectionWithDefaults() *ProjectVariableCollection {
	this := ProjectVariableCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectVariableCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectVariableCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectVariableCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectVariableCollection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectVariableCollection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectVariableCollection) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectVariableCollection) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectVariableCollection) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectVariableCollection) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectVariableCollection) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectVariableCollection) UnsetValue() {
	o.Value.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProjectVariableCollection) GetSecret() bool {
	if o == nil || IsNil(o.Secret) {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectVariableCollection) GetSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *ProjectVariableCollection) SetSecret(v bool) {
	o.Secret = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectVariableCollection) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectVariableCollection) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ProjectVariableCollection) SetProject(v string) {
	o.Project = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectVariableCollection) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectVariableCollection) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectVariableCollection) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *ProjectVariableCollection) SetOrganization(v string) {
	o.Organization = &v
}

func (o ProjectVariableCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectVariableCollection) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableProjectVariableCollection struct {
	value *ProjectVariableCollection
	isSet bool
}

func (v NullableProjectVariableCollection) Get() *ProjectVariableCollection {
	return v.value
}

func (v *NullableProjectVariableCollection) Set(val *ProjectVariableCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectVariableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectVariableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectVariableCollection(val *ProjectVariableCollection) *NullableProjectVariableCollection {
	return &NullableProjectVariableCollection{value: val, isSet: true}
}

func (v NullableProjectVariableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectVariableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
