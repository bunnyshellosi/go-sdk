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

// TemplatesRepositoryCollection A templates repository holds multiple templates.
type TemplatesRepositoryCollection struct {
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

// NewTemplatesRepositoryCollection instantiates a new TemplatesRepositoryCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatesRepositoryCollection() *TemplatesRepositoryCollection {
	this := TemplatesRepositoryCollection{}
	return &this
}

// NewTemplatesRepositoryCollectionWithDefaults instantiates a new TemplatesRepositoryCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesRepositoryCollectionWithDefaults() *TemplatesRepositoryCollection {
	this := TemplatesRepositoryCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplatesRepositoryCollection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryCollection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplatesRepositoryCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplatesRepositoryCollection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryCollection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplatesRepositoryCollection) SetName(v string) {
	o.Name = &v
}

// GetGitRepositoryUrl returns the GitRepositoryUrl field value if set, zero value otherwise.
func (o *TemplatesRepositoryCollection) GetGitRepositoryUrl() string {
	if o == nil || o.GitRepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.GitRepositoryUrl
}

// GetGitRepositoryUrlOk returns a tuple with the GitRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryCollection) GetGitRepositoryUrlOk() (*string, bool) {
	if o == nil || o.GitRepositoryUrl == nil {
		return nil, false
	}
	return o.GitRepositoryUrl, true
}

// HasGitRepositoryUrl returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasGitRepositoryUrl() bool {
	if o != nil && o.GitRepositoryUrl != nil {
		return true
	}

	return false
}

// SetGitRepositoryUrl gets a reference to the given string and assigns it to the GitRepositoryUrl field.
func (o *TemplatesRepositoryCollection) SetGitRepositoryUrl(v string) {
	o.GitRepositoryUrl = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *TemplatesRepositoryCollection) GetGitRef() string {
	if o == nil || o.GitRef == nil {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesRepositoryCollection) GetGitRefOk() (*string, bool) {
	if o == nil || o.GitRef == nil {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasGitRef() bool {
	if o != nil && o.GitRef != nil {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *TemplatesRepositoryCollection) SetGitRef(v string) {
	o.GitRef = &v
}

// GetLastSyncSha returns the LastSyncSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatesRepositoryCollection) GetLastSyncSha() string {
	if o == nil || o.LastSyncSha.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSyncSha.Get()
}

// GetLastSyncShaOk returns a tuple with the LastSyncSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatesRepositoryCollection) GetLastSyncShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSyncSha.Get(), o.LastSyncSha.IsSet()
}

// HasLastSyncSha returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasLastSyncSha() bool {
	if o != nil && o.LastSyncSha.IsSet() {
		return true
	}

	return false
}

// SetLastSyncSha gets a reference to the given NullableString and assigns it to the LastSyncSha field.
func (o *TemplatesRepositoryCollection) SetLastSyncSha(v string) {
	o.LastSyncSha.Set(&v)
}

// SetLastSyncShaNil sets the value for LastSyncSha to be an explicit nil
func (o *TemplatesRepositoryCollection) SetLastSyncShaNil() {
	o.LastSyncSha.Set(nil)
}

// UnsetLastSyncSha ensures that no value is present for LastSyncSha, not even an explicit nil
func (o *TemplatesRepositoryCollection) UnsetLastSyncSha() {
	o.LastSyncSha.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatesRepositoryCollection) GetOrganization() string {
	if o == nil || o.Organization.Get() == nil {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatesRepositoryCollection) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *TemplatesRepositoryCollection) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *TemplatesRepositoryCollection) SetOrganization(v string) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *TemplatesRepositoryCollection) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *TemplatesRepositoryCollection) UnsetOrganization() {
	o.Organization.Unset()
}

func (o TemplatesRepositoryCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GitRepositoryUrl != nil {
		toSerialize["gitRepositoryUrl"] = o.GitRepositoryUrl
	}
	if o.GitRef != nil {
		toSerialize["gitRef"] = o.GitRef
	}
	if o.LastSyncSha.IsSet() {
		toSerialize["lastSyncSha"] = o.LastSyncSha.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTemplatesRepositoryCollection struct {
	value *TemplatesRepositoryCollection
	isSet bool
}

func (v NullableTemplatesRepositoryCollection) Get() *TemplatesRepositoryCollection {
	return v.value
}

func (v *NullableTemplatesRepositoryCollection) Set(val *TemplatesRepositoryCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatesRepositoryCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatesRepositoryCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatesRepositoryCollection(val *TemplatesRepositoryCollection) *NullableTemplatesRepositoryCollection {
	return &NullableTemplatesRepositoryCollection{value: val, isSet: true}
}

func (v NullableTemplatesRepositoryCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatesRepositoryCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}