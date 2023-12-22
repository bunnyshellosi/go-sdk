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

// checks if the ProjectItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectItem{}

// ProjectItem A project holds multiple environments and shared secrets and settings.
type ProjectItem struct {
	// Project identifier.
	Id *string `json:"id,omitempty"`
	// Environment labels.
	Labels *map[string]string `json:"labels,omitempty"`
	// Project name.
	Name *string `json:"name,omitempty"`
	// Environment identifier.
	TotalEnvironments *int32                    `json:"totalEnvironments,omitempty"`
	BuildSettings     NullableBuildSettingsItem `json:"buildSettings,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewProjectItem instantiates a new ProjectItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectItem() *ProjectItem {
	this := ProjectItem{}
	return &this
}

// NewProjectItemWithDefaults instantiates a new ProjectItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectItemWithDefaults() *ProjectItem {
	this := ProjectItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectItem) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ProjectItem) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ProjectItem) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ProjectItem) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectItem) SetName(v string) {
	o.Name = &v
}

// GetTotalEnvironments returns the TotalEnvironments field value if set, zero value otherwise.
func (o *ProjectItem) GetTotalEnvironments() int32 {
	if o == nil || IsNil(o.TotalEnvironments) {
		var ret int32
		return ret
	}
	return *o.TotalEnvironments
}

// GetTotalEnvironmentsOk returns a tuple with the TotalEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetTotalEnvironmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalEnvironments) {
		return nil, false
	}
	return o.TotalEnvironments, true
}

// HasTotalEnvironments returns a boolean if a field has been set.
func (o *ProjectItem) HasTotalEnvironments() bool {
	if o != nil && !IsNil(o.TotalEnvironments) {
		return true
	}

	return false
}

// SetTotalEnvironments gets a reference to the given int32 and assigns it to the TotalEnvironments field.
func (o *ProjectItem) SetTotalEnvironments(v int32) {
	o.TotalEnvironments = &v
}

// GetBuildSettings returns the BuildSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectItem) GetBuildSettings() BuildSettingsItem {
	if o == nil || IsNil(o.BuildSettings.Get()) {
		var ret BuildSettingsItem
		return ret
	}
	return *o.BuildSettings.Get()
}

// GetBuildSettingsOk returns a tuple with the BuildSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectItem) GetBuildSettingsOk() (*BuildSettingsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildSettings.Get(), o.BuildSettings.IsSet()
}

// HasBuildSettings returns a boolean if a field has been set.
func (o *ProjectItem) HasBuildSettings() bool {
	if o != nil && o.BuildSettings.IsSet() {
		return true
	}

	return false
}

// SetBuildSettings gets a reference to the given NullableBuildSettingsItem and assigns it to the BuildSettings field.
func (o *ProjectItem) SetBuildSettings(v BuildSettingsItem) {
	o.BuildSettings.Set(&v)
}

// SetBuildSettingsNil sets the value for BuildSettings to be an explicit nil
func (o *ProjectItem) SetBuildSettingsNil() {
	o.BuildSettings.Set(nil)
}

// UnsetBuildSettings ensures that no value is present for BuildSettings, not even an explicit nil
func (o *ProjectItem) UnsetBuildSettings() {
	o.BuildSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectItem) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectItem) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectItem) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *ProjectItem) SetOrganization(v string) {
	o.Organization = &v
}

func (o ProjectItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TotalEnvironments) {
		toSerialize["totalEnvironments"] = o.TotalEnvironments
	}
	if o.BuildSettings.IsSet() {
		toSerialize["buildSettings"] = o.BuildSettings.Get()
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableProjectItem struct {
	value *ProjectItem
	isSet bool
}

func (v NullableProjectItem) Get() *ProjectItem {
	return v.value
}

func (v *NullableProjectItem) Set(val *ProjectItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectItem(val *ProjectItem) *NullableProjectItem {
	return &NullableProjectItem{value: val, isSet: true}
}

func (v NullableProjectItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
