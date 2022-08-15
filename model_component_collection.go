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

// ComponentCollection A service component represents either an application or a group of applications as a single unit
type ComponentCollection struct {
	// Service component identifier
	Id *string `json:"id,omitempty"`
	// Service component name
	Name *string `json:"name,omitempty"`
	// Service component cluster status
	ClusterStatus *string `json:"clusterStatus,omitempty"`
	// Service component operation status
	OperationStatus *string `json:"operationStatus,omitempty"`
	// Environment identifier.
	Environment *string `json:"environment,omitempty"`
}

// NewComponentCollection instantiates a new ComponentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentCollection() *ComponentCollection {
	this := ComponentCollection{}
	return &this
}

// NewComponentCollectionWithDefaults instantiates a new ComponentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentCollectionWithDefaults() *ComponentCollection {
	this := ComponentCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentCollection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCollection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentCollection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCollection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentCollection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentCollection) SetName(v string) {
	o.Name = &v
}

// GetClusterStatus returns the ClusterStatus field value if set, zero value otherwise.
func (o *ComponentCollection) GetClusterStatus() string {
	if o == nil || o.ClusterStatus == nil {
		var ret string
		return ret
	}
	return *o.ClusterStatus
}

// GetClusterStatusOk returns a tuple with the ClusterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCollection) GetClusterStatusOk() (*string, bool) {
	if o == nil || o.ClusterStatus == nil {
		return nil, false
	}
	return o.ClusterStatus, true
}

// HasClusterStatus returns a boolean if a field has been set.
func (o *ComponentCollection) HasClusterStatus() bool {
	if o != nil && o.ClusterStatus != nil {
		return true
	}

	return false
}

// SetClusterStatus gets a reference to the given string and assigns it to the ClusterStatus field.
func (o *ComponentCollection) SetClusterStatus(v string) {
	o.ClusterStatus = &v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *ComponentCollection) GetOperationStatus() string {
	if o == nil || o.OperationStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCollection) GetOperationStatusOk() (*string, bool) {
	if o == nil || o.OperationStatus == nil {
		return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *ComponentCollection) HasOperationStatus() bool {
	if o != nil && o.OperationStatus != nil {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given string and assigns it to the OperationStatus field.
func (o *ComponentCollection) SetOperationStatus(v string) {
	o.OperationStatus = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ComponentCollection) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCollection) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ComponentCollection) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ComponentCollection) SetEnvironment(v string) {
	o.Environment = &v
}

func (o ComponentCollection) MarshalJSON() ([]byte, error) {
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
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableComponentCollection struct {
	value *ComponentCollection
	isSet bool
}

func (v NullableComponentCollection) Get() *ComponentCollection {
	return v.value
}

func (v *NullableComponentCollection) Set(val *ComponentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentCollection(val *ComponentCollection) *NullableComponentCollection {
	return &NullableComponentCollection{value: val, isSet: true}
}

func (v NullableComponentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
