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

// checks if the KubernetesIntegrationCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesIntegrationCollection{}

// KubernetesIntegrationCollection A Kubernetes integration stores connection information for a Kubernetes cluster.
type KubernetesIntegrationCollection struct {
	// Kubernetes integration identifier.
	Id *string `json:"id,omitempty"`
	// Kubernetes integration cluster name.
	ClusterName *string `json:"clusterName,omitempty"`
	// Kubernetes integration cloud name.
	CloudName *string `json:"cloudName,omitempty"`
	// Kubernetes integration cluster provider.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Kubernetes integration status.
	Status *string `json:"status,omitempty"`
	// Organization identifier.
	Organization NullableString `json:"organization,omitempty"`
}

// NewKubernetesIntegrationCollection instantiates a new KubernetesIntegrationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIntegrationCollection() *KubernetesIntegrationCollection {
	this := KubernetesIntegrationCollection{}
	return &this
}

// NewKubernetesIntegrationCollectionWithDefaults instantiates a new KubernetesIntegrationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIntegrationCollectionWithDefaults() *KubernetesIntegrationCollection {
	this := KubernetesIntegrationCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesIntegrationCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIntegrationCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KubernetesIntegrationCollection) SetId(v string) {
	o.Id = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *KubernetesIntegrationCollection) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIntegrationCollection) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *KubernetesIntegrationCollection) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCloudName returns the CloudName field value if set, zero value otherwise.
func (o *KubernetesIntegrationCollection) GetCloudName() string {
	if o == nil || IsNil(o.CloudName) {
		var ret string
		return ret
	}
	return *o.CloudName
}

// GetCloudNameOk returns a tuple with the CloudName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIntegrationCollection) GetCloudNameOk() (*string, bool) {
	if o == nil || IsNil(o.CloudName) {
		return nil, false
	}
	return o.CloudName, true
}

// HasCloudName returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasCloudName() bool {
	if o != nil && !IsNil(o.CloudName) {
		return true
	}

	return false
}

// SetCloudName gets a reference to the given string and assigns it to the CloudName field.
func (o *KubernetesIntegrationCollection) SetCloudName(v string) {
	o.CloudName = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KubernetesIntegrationCollection) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIntegrationCollection) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KubernetesIntegrationCollection) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubernetesIntegrationCollection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIntegrationCollection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KubernetesIntegrationCollection) SetStatus(v string) {
	o.Status = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesIntegrationCollection) GetOrganization() string {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesIntegrationCollection) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesIntegrationCollection) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *KubernetesIntegrationCollection) SetOrganization(v string) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *KubernetesIntegrationCollection) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *KubernetesIntegrationCollection) UnsetOrganization() {
	o.Organization.Unset()
}

func (o KubernetesIntegrationCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesIntegrationCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !IsNil(o.CloudName) {
		toSerialize["cloudName"] = o.CloudName
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesIntegrationCollection struct {
	value *KubernetesIntegrationCollection
	isSet bool
}

func (v NullableKubernetesIntegrationCollection) Get() *KubernetesIntegrationCollection {
	return v.value
}

func (v *NullableKubernetesIntegrationCollection) Set(val *KubernetesIntegrationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIntegrationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIntegrationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIntegrationCollection(val *KubernetesIntegrationCollection) *NullableKubernetesIntegrationCollection {
	return &NullableKubernetesIntegrationCollection{value: val, isSet: true}
}

func (v NullableKubernetesIntegrationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIntegrationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
