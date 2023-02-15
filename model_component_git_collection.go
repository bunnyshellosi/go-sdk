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

// ComponentGitCollection Git info
type ComponentGitCollection struct {
	// Service component identifier
	Id *string `json:"id,omitempty"`
	// Service component name
	Name *string `json:"name,omitempty"`
	// Git repository
	Repository *string `json:"repository,omitempty"`
	// Git ref name
	RefName *string `json:"refName,omitempty"`
	// Git ref sha
	RefSha NullableString `json:"refSha,omitempty"`
	// Git deployed sha
	DeployedSha NullableString `json:"deployedSha,omitempty"`
	// Environment identifier.
	Environment *string `json:"environment,omitempty"`
}

// NewComponentGitCollection instantiates a new ComponentGitCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentGitCollection() *ComponentGitCollection {
	this := ComponentGitCollection{}
	return &this
}

// NewComponentGitCollectionWithDefaults instantiates a new ComponentGitCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentGitCollectionWithDefaults() *ComponentGitCollection {
	this := ComponentGitCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentGitCollection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGitCollection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentGitCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentGitCollection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGitCollection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentGitCollection) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ComponentGitCollection) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGitCollection) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ComponentGitCollection) SetRepository(v string) {
	o.Repository = &v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *ComponentGitCollection) GetRefName() string {
	if o == nil || o.RefName == nil {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGitCollection) GetRefNameOk() (*string, bool) {
	if o == nil || o.RefName == nil {
		return nil, false
	}
	return o.RefName, true
}

// HasRefName returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasRefName() bool {
	if o != nil && o.RefName != nil {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *ComponentGitCollection) SetRefName(v string) {
	o.RefName = &v
}

// GetRefSha returns the RefSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentGitCollection) GetRefSha() string {
	if o == nil || o.RefSha.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefSha.Get()
}

// GetRefShaOk returns a tuple with the RefSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentGitCollection) GetRefShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefSha.Get(), o.RefSha.IsSet()
}

// HasRefSha returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasRefSha() bool {
	if o != nil && o.RefSha.IsSet() {
		return true
	}

	return false
}

// SetRefSha gets a reference to the given NullableString and assigns it to the RefSha field.
func (o *ComponentGitCollection) SetRefSha(v string) {
	o.RefSha.Set(&v)
}

// SetRefShaNil sets the value for RefSha to be an explicit nil
func (o *ComponentGitCollection) SetRefShaNil() {
	o.RefSha.Set(nil)
}

// UnsetRefSha ensures that no value is present for RefSha, not even an explicit nil
func (o *ComponentGitCollection) UnsetRefSha() {
	o.RefSha.Unset()
}

// GetDeployedSha returns the DeployedSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentGitCollection) GetDeployedSha() string {
	if o == nil || o.DeployedSha.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeployedSha.Get()
}

// GetDeployedShaOk returns a tuple with the DeployedSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentGitCollection) GetDeployedShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeployedSha.Get(), o.DeployedSha.IsSet()
}

// HasDeployedSha returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasDeployedSha() bool {
	if o != nil && o.DeployedSha.IsSet() {
		return true
	}

	return false
}

// SetDeployedSha gets a reference to the given NullableString and assigns it to the DeployedSha field.
func (o *ComponentGitCollection) SetDeployedSha(v string) {
	o.DeployedSha.Set(&v)
}

// SetDeployedShaNil sets the value for DeployedSha to be an explicit nil
func (o *ComponentGitCollection) SetDeployedShaNil() {
	o.DeployedSha.Set(nil)
}

// UnsetDeployedSha ensures that no value is present for DeployedSha, not even an explicit nil
func (o *ComponentGitCollection) UnsetDeployedSha() {
	o.DeployedSha.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ComponentGitCollection) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGitCollection) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ComponentGitCollection) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ComponentGitCollection) SetEnvironment(v string) {
	o.Environment = &v
}

func (o ComponentGitCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.RefName != nil {
		toSerialize["refName"] = o.RefName
	}
	if o.RefSha.IsSet() {
		toSerialize["refSha"] = o.RefSha.Get()
	}
	if o.DeployedSha.IsSet() {
		toSerialize["deployedSha"] = o.DeployedSha.Get()
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableComponentGitCollection struct {
	value *ComponentGitCollection
	isSet bool
}

func (v NullableComponentGitCollection) Get() *ComponentGitCollection {
	return v.value
}

func (v *NullableComponentGitCollection) Set(val *ComponentGitCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentGitCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentGitCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentGitCollection(val *ComponentGitCollection) *NullableComponentGitCollection {
	return &NullableComponentGitCollection{value: val, isSet: true}
}

func (v NullableComponentGitCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentGitCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}