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

// GitInfo struct for GitInfo
type GitInfo struct {
	Repository NullableString `json:"repository,omitempty"`
	Branch     NullableString `json:"branch,omitempty"`
}

// NewGitInfo instantiates a new GitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitInfo() *GitInfo {
	this := GitInfo{}
	return &this
}

// NewGitInfoWithDefaults instantiates a new GitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitInfoWithDefaults() *GitInfo {
	this := GitInfo{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitInfo) GetRepository() string {
	if o == nil || o.Repository.Get() == nil {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitInfo) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *GitInfo) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *GitInfo) SetRepository(v string) {
	o.Repository.Set(&v)
}

// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *GitInfo) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *GitInfo) UnsetRepository() {
	o.Repository.Unset()
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitInfo) GetBranch() string {
	if o == nil || o.Branch.Get() == nil {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitInfo) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *GitInfo) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *GitInfo) SetBranch(v string) {
	o.Branch.Set(&v)
}

// SetBranchNil sets the value for Branch to be an explicit nil
func (o *GitInfo) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *GitInfo) UnsetBranch() {
	o.Branch.Unset()
}

func (o GitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGitInfo struct {
	value *GitInfo
	isSet bool
}

func (v NullableGitInfo) Get() *GitInfo {
	return v.value
}

func (v *NullableGitInfo) Set(val *GitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitInfo(val *GitInfo) *NullableGitInfo {
	return &NullableGitInfo{value: val, isSet: true}
}

func (v NullableGitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
