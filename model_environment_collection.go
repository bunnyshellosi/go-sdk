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

// EnvironmentCollection An environment holds a collection of buildable and deployable components.
type EnvironmentCollection struct {
	// Environment identifier.
	Id *string `json:"id,omitempty"`
	// Environment type: primary or ephemeral.
	Type *string `json:"type,omitempty"`
	// Environment name.
	Name *string `json:"name,omitempty"`
	// Environment k8s namespace.
	Namespace *string `json:"namespace,omitempty"`
	// Environment operation status.
	OperationStatus *string `json:"operationStatus,omitempty"`
	// Environment cluster status.
	ClusterStatus *string `json:"clusterStatus,omitempty"`
	// Project identifier.
	Project *string `json:"project,omitempty"`
}

// NewEnvironmentCollection instantiates a new EnvironmentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentCollection() *EnvironmentCollection {
	this := EnvironmentCollection{}
	return &this
}

// NewEnvironmentCollectionWithDefaults instantiates a new EnvironmentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentCollectionWithDefaults() *EnvironmentCollection {
	this := EnvironmentCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentCollection) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentCollection) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentCollection) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *EnvironmentCollection) SetNamespace(v string) {
	o.Namespace = &v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetOperationStatus() string {
	if o == nil || o.OperationStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetOperationStatusOk() (*string, bool) {
	if o == nil || o.OperationStatus == nil {
		return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasOperationStatus() bool {
	if o != nil && o.OperationStatus != nil {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given string and assigns it to the OperationStatus field.
func (o *EnvironmentCollection) SetOperationStatus(v string) {
	o.OperationStatus = &v
}

// GetClusterStatus returns the ClusterStatus field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetClusterStatus() string {
	if o == nil || o.ClusterStatus == nil {
		var ret string
		return ret
	}
	return *o.ClusterStatus
}

// GetClusterStatusOk returns a tuple with the ClusterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetClusterStatusOk() (*string, bool) {
	if o == nil || o.ClusterStatus == nil {
		return nil, false
	}
	return o.ClusterStatus, true
}

// HasClusterStatus returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasClusterStatus() bool {
	if o != nil && o.ClusterStatus != nil {
		return true
	}

	return false
}

// SetClusterStatus gets a reference to the given string and assigns it to the ClusterStatus field.
func (o *EnvironmentCollection) SetClusterStatus(v string) {
	o.ClusterStatus = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *EnvironmentCollection) SetProject(v string) {
	o.Project = &v
}

func (o EnvironmentCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.OperationStatus != nil {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	if o.ClusterStatus != nil {
		toSerialize["clusterStatus"] = o.ClusterStatus
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentCollection struct {
	value *EnvironmentCollection
	isSet bool
}

func (v NullableEnvironmentCollection) Get() *EnvironmentCollection {
	return v.value
}

func (v *NullableEnvironmentCollection) Set(val *EnvironmentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentCollection(val *EnvironmentCollection) *NullableEnvironmentCollection {
	return &NullableEnvironmentCollection{value: val, isSet: true}
}

func (v NullableEnvironmentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
