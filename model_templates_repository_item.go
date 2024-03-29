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

// checks if the TemplatesRepositoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatesRepositoryItem{}

// TemplatesRepositoryItem A templates repository holds multiple templates.
type TemplatesRepositoryItem struct {
	// Templates repository identifier.
	Id *string `json:"id,omitempty"`
	// Templates repository name.
	Name *string `json:"name,omitempty"`
	// Templates repository git repository URL.
	GitRepositoryUrl *string `json:"gitRepositoryUrl,omitempty"`
	// Templates repository git repository reference.
	GitRef *string `json:"gitRef,omitempty"`
	// Templates repository last synced git repository SHA.
	LastSyncSha NullableString `json:"lastSyncSha,omitempty"`
	// Organization identifier.
	Organization NullableString `json:"organization,omitempty"`
}

// NewTemplatesRepositoryItem instantiates a new TemplatesRepositoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatesRepositoryItem() *TemplatesRepositoryItem {
	this := TemplatesRepositoryItem{}
	return &this
}

// NewTemplatesRepositoryItemWithDefaults instantiates a new TemplatesRepositoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesRepositoryItemWithDefaults() *TemplatesRepositoryItem {
	this := TemplatesRepositoryItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplatesRepositoryItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplatesRepositoryItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplatesRepositoryItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplatesRepositoryItem) SetName(v string) {
	o.Name = &v
}

// GetGitRepositoryUrl returns the GitRepositoryUrl field value if set, zero value otherwise.
func (o *TemplatesRepositoryItem) GetGitRepositoryUrl() string {
	if o == nil || IsNil(o.GitRepositoryUrl) {
		var ret string
		return ret
	}
	return *o.GitRepositoryUrl
}

// GetGitRepositoryUrlOk returns a tuple with the GitRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryItem) GetGitRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GitRepositoryUrl) {
		return nil, false
	}
	return o.GitRepositoryUrl, true
}

// HasGitRepositoryUrl returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasGitRepositoryUrl() bool {
	if o != nil && !IsNil(o.GitRepositoryUrl) {
		return true
	}

	return false
}

// SetGitRepositoryUrl gets a reference to the given string and assigns it to the GitRepositoryUrl field.
func (o *TemplatesRepositoryItem) SetGitRepositoryUrl(v string) {
	o.GitRepositoryUrl = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *TemplatesRepositoryItem) GetGitRef() string {
	if o == nil || IsNil(o.GitRef) {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryItem) GetGitRefOk() (*string, bool) {
	if o == nil || IsNil(o.GitRef) {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasGitRef() bool {
	if o != nil && !IsNil(o.GitRef) {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *TemplatesRepositoryItem) SetGitRef(v string) {
	o.GitRef = &v
}

// GetLastSyncSha returns the LastSyncSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatesRepositoryItem) GetLastSyncSha() string {
	if o == nil || IsNil(o.LastSyncSha.Get()) {
		var ret string
		return ret
	}
	return *o.LastSyncSha.Get()
}

// GetLastSyncShaOk returns a tuple with the LastSyncSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatesRepositoryItem) GetLastSyncShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncSha.Get(), o.LastSyncSha.IsSet()
}

// HasLastSyncSha returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasLastSyncSha() bool {
	if o != nil && o.LastSyncSha.IsSet() {
		return true
	}

	return false
}

// SetLastSyncSha gets a reference to the given NullableString and assigns it to the LastSyncSha field.
func (o *TemplatesRepositoryItem) SetLastSyncSha(v string) {
	o.LastSyncSha.Set(&v)
}

// SetLastSyncShaNil sets the value for LastSyncSha to be an explicit nil
func (o *TemplatesRepositoryItem) SetLastSyncShaNil() {
	o.LastSyncSha.Set(nil)
}

// UnsetLastSyncSha ensures that no value is present for LastSyncSha, not even an explicit nil
func (o *TemplatesRepositoryItem) UnsetLastSyncSha() {
	o.LastSyncSha.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatesRepositoryItem) GetOrganization() string {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatesRepositoryItem) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *TemplatesRepositoryItem) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *TemplatesRepositoryItem) SetOrganization(v string) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *TemplatesRepositoryItem) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *TemplatesRepositoryItem) UnsetOrganization() {
	o.Organization.Unset()
}

func (o TemplatesRepositoryItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatesRepositoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.GitRepositoryUrl) {
		toSerialize["gitRepositoryUrl"] = o.GitRepositoryUrl
	}
	if !IsNil(o.GitRef) {
		toSerialize["gitRef"] = o.GitRef
	}
	if o.LastSyncSha.IsSet() {
		toSerialize["lastSyncSha"] = o.LastSyncSha.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	return toSerialize, nil
}

type NullableTemplatesRepositoryItem struct {
	value *TemplatesRepositoryItem
	isSet bool
}

func (v NullableTemplatesRepositoryItem) Get() *TemplatesRepositoryItem {
	return v.value
}

func (v *NullableTemplatesRepositoryItem) Set(val *TemplatesRepositoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatesRepositoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatesRepositoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatesRepositoryItem(val *TemplatesRepositoryItem) *NullableTemplatesRepositoryItem {
	return &NullableTemplatesRepositoryItem{value: val, isSet: true}
}

func (v NullableTemplatesRepositoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatesRepositoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
