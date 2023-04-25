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

// ProfileItem struct for ProfileItem
type ProfileItem struct {
	// The command to be run when starting the container.
	Command []string `json:"command,omitempty"`
	// The port mapping for the container.
	PortMapping []string `json:"portMapping,omitempty"`
	// The environ for the container.
	Environ *map[string]string `json:"environ,omitempty"`
	// The sync paths for the container.
	SyncPaths    []SyncPathItem                  `json:"syncPaths,omitempty"`
	Requirements NullableProfileItemRequirements `json:"requirements,omitempty"`
}

// NewProfileItem instantiates a new ProfileItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileItem() *ProfileItem {
	this := ProfileItem{}
	return &this
}

// NewProfileItemWithDefaults instantiates a new ProfileItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileItemWithDefaults() *ProfileItem {
	this := ProfileItem{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *ProfileItem) GetCommand() []string {
	if o == nil || o.Command == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileItem) GetCommandOk() ([]string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *ProfileItem) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *ProfileItem) SetCommand(v []string) {
	o.Command = v
}

// GetPortMapping returns the PortMapping field value if set, zero value otherwise.
func (o *ProfileItem) GetPortMapping() []string {
	if o == nil || o.PortMapping == nil {
		var ret []string
		return ret
	}
	return o.PortMapping
}

// GetPortMappingOk returns a tuple with the PortMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileItem) GetPortMappingOk() ([]string, bool) {
	if o == nil || o.PortMapping == nil {
		return nil, false
	}
	return o.PortMapping, true
}

// HasPortMapping returns a boolean if a field has been set.
func (o *ProfileItem) HasPortMapping() bool {
	if o != nil && o.PortMapping != nil {
		return true
	}

	return false
}

// SetPortMapping gets a reference to the given []string and assigns it to the PortMapping field.
func (o *ProfileItem) SetPortMapping(v []string) {
	o.PortMapping = v
}

// GetEnviron returns the Environ field value if set, zero value otherwise.
func (o *ProfileItem) GetEnviron() map[string]string {
	if o == nil || o.Environ == nil {
		var ret map[string]string
		return ret
	}
	return *o.Environ
}

// GetEnvironOk returns a tuple with the Environ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileItem) GetEnvironOk() (*map[string]string, bool) {
	if o == nil || o.Environ == nil {
		return nil, false
	}
	return o.Environ, true
}

// HasEnviron returns a boolean if a field has been set.
func (o *ProfileItem) HasEnviron() bool {
	if o != nil && o.Environ != nil {
		return true
	}

	return false
}

// SetEnviron gets a reference to the given map[string]string and assigns it to the Environ field.
func (o *ProfileItem) SetEnviron(v map[string]string) {
	o.Environ = &v
}

// GetSyncPaths returns the SyncPaths field value if set, zero value otherwise.
func (o *ProfileItem) GetSyncPaths() []SyncPathItem {
	if o == nil || o.SyncPaths == nil {
		var ret []SyncPathItem
		return ret
	}
	return o.SyncPaths
}

// GetSyncPathsOk returns a tuple with the SyncPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileItem) GetSyncPathsOk() ([]SyncPathItem, bool) {
	if o == nil || o.SyncPaths == nil {
		return nil, false
	}
	return o.SyncPaths, true
}

// HasSyncPaths returns a boolean if a field has been set.
func (o *ProfileItem) HasSyncPaths() bool {
	if o != nil && o.SyncPaths != nil {
		return true
	}

	return false
}

// SetSyncPaths gets a reference to the given []SyncPathItem and assigns it to the SyncPaths field.
func (o *ProfileItem) SetSyncPaths(v []SyncPathItem) {
	o.SyncPaths = v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileItem) GetRequirements() ProfileItemRequirements {
	if o == nil || o.Requirements.Get() == nil {
		var ret ProfileItemRequirements
		return ret
	}
	return *o.Requirements.Get()
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileItem) GetRequirementsOk() (*ProfileItemRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requirements.Get(), o.Requirements.IsSet()
}

// HasRequirements returns a boolean if a field has been set.
func (o *ProfileItem) HasRequirements() bool {
	if o != nil && o.Requirements.IsSet() {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given NullableProfileItemRequirements and assigns it to the Requirements field.
func (o *ProfileItem) SetRequirements(v ProfileItemRequirements) {
	o.Requirements.Set(&v)
}

// SetRequirementsNil sets the value for Requirements to be an explicit nil
func (o *ProfileItem) SetRequirementsNil() {
	o.Requirements.Set(nil)
}

// UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
func (o *ProfileItem) UnsetRequirements() {
	o.Requirements.Unset()
}

func (o ProfileItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.PortMapping != nil {
		toSerialize["portMapping"] = o.PortMapping
	}
	if o.Environ != nil {
		toSerialize["environ"] = o.Environ
	}
	if o.SyncPaths != nil {
		toSerialize["syncPaths"] = o.SyncPaths
	}
	if o.Requirements.IsSet() {
		toSerialize["requirements"] = o.Requirements.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProfileItem struct {
	value *ProfileItem
	isSet bool
}

func (v NullableProfileItem) Get() *ProfileItem {
	return v.value
}

func (v *NullableProfileItem) Set(val *ProfileItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileItem(val *ProfileItem) *NullableProfileItem {
	return &NullableProfileItem{value: val, isSet: true}
}

func (v NullableProfileItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
