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

// ClusterWrapperKubeConfigRead struct for ClusterWrapperKubeConfigRead
type ClusterWrapperKubeConfigRead struct {
	Name    string                 `json:"name"`
	Cluster *ClusterKubeConfigRead `json:"cluster,omitempty"`
}

// NewClusterWrapperKubeConfigRead instantiates a new ClusterWrapperKubeConfigRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterWrapperKubeConfigRead(name string) *ClusterWrapperKubeConfigRead {
	this := ClusterWrapperKubeConfigRead{}
	this.Name = name
	return &this
}

// NewClusterWrapperKubeConfigReadWithDefaults instantiates a new ClusterWrapperKubeConfigRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWrapperKubeConfigReadWithDefaults() *ClusterWrapperKubeConfigRead {
	this := ClusterWrapperKubeConfigRead{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterWrapperKubeConfigRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterWrapperKubeConfigRead) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterWrapperKubeConfigRead) SetName(v string) {
	o.Name = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClusterWrapperKubeConfigRead) GetCluster() ClusterKubeConfigRead {
	if o == nil || o.Cluster == nil {
		var ret ClusterKubeConfigRead
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterWrapperKubeConfigRead) GetClusterOk() (*ClusterKubeConfigRead, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClusterWrapperKubeConfigRead) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterKubeConfigRead and assigns it to the Cluster field.
func (o *ClusterWrapperKubeConfigRead) SetCluster(v ClusterKubeConfigRead) {
	o.Cluster = &v
}

func (o ClusterWrapperKubeConfigRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

type NullableClusterWrapperKubeConfigRead struct {
	value *ClusterWrapperKubeConfigRead
	isSet bool
}

func (v NullableClusterWrapperKubeConfigRead) Get() *ClusterWrapperKubeConfigRead {
	return v.value
}

func (v *NullableClusterWrapperKubeConfigRead) Set(val *ClusterWrapperKubeConfigRead) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterWrapperKubeConfigRead) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterWrapperKubeConfigRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterWrapperKubeConfigRead(val *ClusterWrapperKubeConfigRead) *NullableClusterWrapperKubeConfigRead {
	return &NullableClusterWrapperKubeConfigRead{value: val, isSet: true}
}

func (v NullableClusterWrapperKubeConfigRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterWrapperKubeConfigRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
