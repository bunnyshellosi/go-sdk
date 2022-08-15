/*
API Bunnyshell Environments

Interact with Bunnyshell Platform

API version: 1.0.1
Contact: api+support@bunnyshell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// ComponentItem A service component represents either an application or a group of applications as a single unit
type ComponentItem struct {
	// Service component identifier
	Id *string `json:"id,omitempty"`
	// Service component name
	Name *string `json:"name,omitempty"`
	// Service component cluster status
	ClusterStatus *string `json:"clusterStatus,omitempty"`
	// Service component operation status
	OperationStatus *string `json:"operationStatus,omitempty"`
	// Service component URLs
	PublicURLs []string `json:"publicURLs,omitempty"`
	// Environment identifier.
	Environment *string `json:"environment,omitempty"`
}

// NewComponentItem instantiates a new ComponentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentItem() *ComponentItem {
	this := ComponentItem{}
	return &this
}

// NewComponentItemWithDefaults instantiates a new ComponentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentItemWithDefaults() *ComponentItem {
	this := ComponentItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentItem) SetName(v string) {
	o.Name = &v
}

// GetClusterStatus returns the ClusterStatus field value if set, zero value otherwise.
func (o *ComponentItem) GetClusterStatus() string {
	if o == nil || o.ClusterStatus == nil {
		var ret string
		return ret
	}
	return *o.ClusterStatus
}

// GetClusterStatusOk returns a tuple with the ClusterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetClusterStatusOk() (*string, bool) {
	if o == nil || o.ClusterStatus == nil {
		return nil, false
	}
	return o.ClusterStatus, true
}

// HasClusterStatus returns a boolean if a field has been set.
func (o *ComponentItem) HasClusterStatus() bool {
	if o != nil && o.ClusterStatus != nil {
		return true
	}

	return false
}

// SetClusterStatus gets a reference to the given string and assigns it to the ClusterStatus field.
func (o *ComponentItem) SetClusterStatus(v string) {
	o.ClusterStatus = &v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *ComponentItem) GetOperationStatus() string {
	if o == nil || o.OperationStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetOperationStatusOk() (*string, bool) {
	if o == nil || o.OperationStatus == nil {
		return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *ComponentItem) HasOperationStatus() bool {
	if o != nil && o.OperationStatus != nil {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given string and assigns it to the OperationStatus field.
func (o *ComponentItem) SetOperationStatus(v string) {
	o.OperationStatus = &v
}

// GetPublicURLs returns the PublicURLs field value if set, zero value otherwise.
func (o *ComponentItem) GetPublicURLs() []string {
	if o == nil || o.PublicURLs == nil {
		var ret []string
		return ret
	}
	return o.PublicURLs
}

// GetPublicURLsOk returns a tuple with the PublicURLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetPublicURLsOk() ([]string, bool) {
	if o == nil || o.PublicURLs == nil {
		return nil, false
	}
	return o.PublicURLs, true
}

// HasPublicURLs returns a boolean if a field has been set.
func (o *ComponentItem) HasPublicURLs() bool {
	if o != nil && o.PublicURLs != nil {
		return true
	}

	return false
}

// SetPublicURLs gets a reference to the given []string and assigns it to the PublicURLs field.
func (o *ComponentItem) SetPublicURLs(v []string) {
	o.PublicURLs = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ComponentItem) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentItem) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ComponentItem) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ComponentItem) SetEnvironment(v string) {
	o.Environment = &v
}

func (o ComponentItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClusterStatus != nil {
		toSerialize["clusterStatus"] = o.ClusterStatus
	}
	if o.OperationStatus != nil {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	if o.PublicURLs != nil {
		toSerialize["publicURLs"] = o.PublicURLs
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableComponentItem struct {
	value *ComponentItem
	isSet bool
}

func (v NullableComponentItem) Get() *ComponentItem {
	return v.value
}

func (v *NullableComponentItem) Set(val *ComponentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentItem(val *ComponentItem) *NullableComponentItem {
	return &NullableComponentItem{value: val, isSet: true}
}

func (v NullableComponentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
