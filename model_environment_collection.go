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

// checks if the EnvironmentCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentCollection{}

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
	// Kubernetes integration identifier.
	KubernetesIntegration NullableString `json:"kubernetesIntegration,omitempty"`
	// Kubernetes integration identifier for the ephemeral auto deploy cluster.
	EphemeralKubernetesIntegration NullableString `json:"ephemeralKubernetesIntegration,omitempty"`
	// Environment termination protection.
	HasTerminationProtection *bool `json:"hasTerminationProtection,omitempty"`
	// Environment ephemeral auto deploy status.
	HasEphemeralAutoDeploy *bool `json:"hasEphemeralAutoDeploy,omitempty"`
	// Environment ephemeral create-on-PR status.
	HasEphemeralCreateOnPr *bool `json:"hasEphemeralCreateOnPr,omitempty"`
	// Environment ephemeral destroy-on-PR close status.
	HasEphemeralDestroyOnPrClose *bool `json:"hasEphemeralDestroyOnPrClose,omitempty"`
	// Environment ephemeral branch whitelist status.
	HasEphemeralBranchWhitelist *bool `json:"hasEphemeralBranchWhitelist,omitempty"`
	// Environment ephemeral branch whitelist regex.
	EphemeralBranchWhitelistRegex NullableString `json:"ephemeralBranchWhitelistRegex,omitempty"`
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
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
	if o == nil || IsNil(o.OperationStatus) {
		var ret string
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetOperationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperationStatus) {
		return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasOperationStatus() bool {
	if o != nil && !IsNil(o.OperationStatus) {
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
	if o == nil || IsNil(o.ClusterStatus) {
		var ret string
		return ret
	}
	return *o.ClusterStatus
}

// GetClusterStatusOk returns a tuple with the ClusterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetClusterStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterStatus) {
		return nil, false
	}
	return o.ClusterStatus, true
}

// HasClusterStatus returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasClusterStatus() bool {
	if o != nil && !IsNil(o.ClusterStatus) {
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
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *EnvironmentCollection) SetProject(v string) {
	o.Project = &v
}

// GetKubernetesIntegration returns the KubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCollection) GetKubernetesIntegration() string {
	if o == nil || IsNil(o.KubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesIntegration.Get()
}

// GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentCollection) GetKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesIntegration.Get(), o.KubernetesIntegration.IsSet()
}

// HasKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasKubernetesIntegration() bool {
	if o != nil && o.KubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetKubernetesIntegration gets a reference to the given NullableString and assigns it to the KubernetesIntegration field.
func (o *EnvironmentCollection) SetKubernetesIntegration(v string) {
	o.KubernetesIntegration.Set(&v)
}

// SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil
func (o *EnvironmentCollection) SetKubernetesIntegrationNil() {
	o.KubernetesIntegration.Set(nil)
}

// UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
func (o *EnvironmentCollection) UnsetKubernetesIntegration() {
	o.KubernetesIntegration.Unset()
}

// GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCollection) GetEphemeralKubernetesIntegration() string {
	if o == nil || IsNil(o.EphemeralKubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.EphemeralKubernetesIntegration.Get()
}

// GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentCollection) GetEphemeralKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EphemeralKubernetesIntegration.Get(), o.EphemeralKubernetesIntegration.IsSet()
}

// HasEphemeralKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasEphemeralKubernetesIntegration() bool {
	if o != nil && o.EphemeralKubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetEphemeralKubernetesIntegration gets a reference to the given NullableString and assigns it to the EphemeralKubernetesIntegration field.
func (o *EnvironmentCollection) SetEphemeralKubernetesIntegration(v string) {
	o.EphemeralKubernetesIntegration.Set(&v)
}

// SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil
func (o *EnvironmentCollection) SetEphemeralKubernetesIntegrationNil() {
	o.EphemeralKubernetesIntegration.Set(nil)
}

// UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
func (o *EnvironmentCollection) UnsetEphemeralKubernetesIntegration() {
	o.EphemeralKubernetesIntegration.Unset()
}

// GetHasTerminationProtection returns the HasTerminationProtection field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetHasTerminationProtection() bool {
	if o == nil || IsNil(o.HasTerminationProtection) {
		var ret bool
		return ret
	}
	return *o.HasTerminationProtection
}

// GetHasTerminationProtectionOk returns a tuple with the HasTerminationProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetHasTerminationProtectionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTerminationProtection) {
		return nil, false
	}
	return o.HasTerminationProtection, true
}

// HasHasTerminationProtection returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasHasTerminationProtection() bool {
	if o != nil && !IsNil(o.HasTerminationProtection) {
		return true
	}

	return false
}

// SetHasTerminationProtection gets a reference to the given bool and assigns it to the HasTerminationProtection field.
func (o *EnvironmentCollection) SetHasTerminationProtection(v bool) {
	o.HasTerminationProtection = &v
}

// GetHasEphemeralAutoDeploy returns the HasEphemeralAutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetHasEphemeralAutoDeploy() bool {
	if o == nil || IsNil(o.HasEphemeralAutoDeploy) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralAutoDeploy
}

// GetHasEphemeralAutoDeployOk returns a tuple with the HasEphemeralAutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetHasEphemeralAutoDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralAutoDeploy) {
		return nil, false
	}
	return o.HasEphemeralAutoDeploy, true
}

// HasHasEphemeralAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasHasEphemeralAutoDeploy() bool {
	if o != nil && !IsNil(o.HasEphemeralAutoDeploy) {
		return true
	}

	return false
}

// SetHasEphemeralAutoDeploy gets a reference to the given bool and assigns it to the HasEphemeralAutoDeploy field.
func (o *EnvironmentCollection) SetHasEphemeralAutoDeploy(v bool) {
	o.HasEphemeralAutoDeploy = &v
}

// GetHasEphemeralCreateOnPr returns the HasEphemeralCreateOnPr field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetHasEphemeralCreateOnPr() bool {
	if o == nil || IsNil(o.HasEphemeralCreateOnPr) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralCreateOnPr
}

// GetHasEphemeralCreateOnPrOk returns a tuple with the HasEphemeralCreateOnPr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetHasEphemeralCreateOnPrOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralCreateOnPr) {
		return nil, false
	}
	return o.HasEphemeralCreateOnPr, true
}

// HasHasEphemeralCreateOnPr returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasHasEphemeralCreateOnPr() bool {
	if o != nil && !IsNil(o.HasEphemeralCreateOnPr) {
		return true
	}

	return false
}

// SetHasEphemeralCreateOnPr gets a reference to the given bool and assigns it to the HasEphemeralCreateOnPr field.
func (o *EnvironmentCollection) SetHasEphemeralCreateOnPr(v bool) {
	o.HasEphemeralCreateOnPr = &v
}

// GetHasEphemeralDestroyOnPrClose returns the HasEphemeralDestroyOnPrClose field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetHasEphemeralDestroyOnPrClose() bool {
	if o == nil || IsNil(o.HasEphemeralDestroyOnPrClose) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralDestroyOnPrClose
}

// GetHasEphemeralDestroyOnPrCloseOk returns a tuple with the HasEphemeralDestroyOnPrClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetHasEphemeralDestroyOnPrCloseOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralDestroyOnPrClose) {
		return nil, false
	}
	return o.HasEphemeralDestroyOnPrClose, true
}

// HasHasEphemeralDestroyOnPrClose returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasHasEphemeralDestroyOnPrClose() bool {
	if o != nil && !IsNil(o.HasEphemeralDestroyOnPrClose) {
		return true
	}

	return false
}

// SetHasEphemeralDestroyOnPrClose gets a reference to the given bool and assigns it to the HasEphemeralDestroyOnPrClose field.
func (o *EnvironmentCollection) SetHasEphemeralDestroyOnPrClose(v bool) {
	o.HasEphemeralDestroyOnPrClose = &v
}

// GetHasEphemeralBranchWhitelist returns the HasEphemeralBranchWhitelist field value if set, zero value otherwise.
func (o *EnvironmentCollection) GetHasEphemeralBranchWhitelist() bool {
	if o == nil || IsNil(o.HasEphemeralBranchWhitelist) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralBranchWhitelist
}

// GetHasEphemeralBranchWhitelistOk returns a tuple with the HasEphemeralBranchWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCollection) GetHasEphemeralBranchWhitelistOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralBranchWhitelist) {
		return nil, false
	}
	return o.HasEphemeralBranchWhitelist, true
}

// HasHasEphemeralBranchWhitelist returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasHasEphemeralBranchWhitelist() bool {
	if o != nil && !IsNil(o.HasEphemeralBranchWhitelist) {
		return true
	}

	return false
}

// SetHasEphemeralBranchWhitelist gets a reference to the given bool and assigns it to the HasEphemeralBranchWhitelist field.
func (o *EnvironmentCollection) SetHasEphemeralBranchWhitelist(v bool) {
	o.HasEphemeralBranchWhitelist = &v
}

// GetEphemeralBranchWhitelistRegex returns the EphemeralBranchWhitelistRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCollection) GetEphemeralBranchWhitelistRegex() string {
	if o == nil || IsNil(o.EphemeralBranchWhitelistRegex.Get()) {
		var ret string
		return ret
	}
	return *o.EphemeralBranchWhitelistRegex.Get()
}

// GetEphemeralBranchWhitelistRegexOk returns a tuple with the EphemeralBranchWhitelistRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentCollection) GetEphemeralBranchWhitelistRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EphemeralBranchWhitelistRegex.Get(), o.EphemeralBranchWhitelistRegex.IsSet()
}

// HasEphemeralBranchWhitelistRegex returns a boolean if a field has been set.
func (o *EnvironmentCollection) HasEphemeralBranchWhitelistRegex() bool {
	if o != nil && o.EphemeralBranchWhitelistRegex.IsSet() {
		return true
	}

	return false
}

// SetEphemeralBranchWhitelistRegex gets a reference to the given NullableString and assigns it to the EphemeralBranchWhitelistRegex field.
func (o *EnvironmentCollection) SetEphemeralBranchWhitelistRegex(v string) {
	o.EphemeralBranchWhitelistRegex.Set(&v)
}

// SetEphemeralBranchWhitelistRegexNil sets the value for EphemeralBranchWhitelistRegex to be an explicit nil
func (o *EnvironmentCollection) SetEphemeralBranchWhitelistRegexNil() {
	o.EphemeralBranchWhitelistRegex.Set(nil)
}

// UnsetEphemeralBranchWhitelistRegex ensures that no value is present for EphemeralBranchWhitelistRegex, not even an explicit nil
func (o *EnvironmentCollection) UnsetEphemeralBranchWhitelistRegex() {
	o.EphemeralBranchWhitelistRegex.Unset()
}

func (o EnvironmentCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.OperationStatus) {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	if !IsNil(o.ClusterStatus) {
		toSerialize["clusterStatus"] = o.ClusterStatus
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if o.KubernetesIntegration.IsSet() {
		toSerialize["kubernetesIntegration"] = o.KubernetesIntegration.Get()
	}
	if o.EphemeralKubernetesIntegration.IsSet() {
		toSerialize["ephemeralKubernetesIntegration"] = o.EphemeralKubernetesIntegration.Get()
	}
	if !IsNil(o.HasTerminationProtection) {
		toSerialize["hasTerminationProtection"] = o.HasTerminationProtection
	}
	if !IsNil(o.HasEphemeralAutoDeploy) {
		toSerialize["hasEphemeralAutoDeploy"] = o.HasEphemeralAutoDeploy
	}
	if !IsNil(o.HasEphemeralCreateOnPr) {
		toSerialize["hasEphemeralCreateOnPr"] = o.HasEphemeralCreateOnPr
	}
	if !IsNil(o.HasEphemeralDestroyOnPrClose) {
		toSerialize["hasEphemeralDestroyOnPrClose"] = o.HasEphemeralDestroyOnPrClose
	}
	if !IsNil(o.HasEphemeralBranchWhitelist) {
		toSerialize["hasEphemeralBranchWhitelist"] = o.HasEphemeralBranchWhitelist
	}
	if o.EphemeralBranchWhitelistRegex.IsSet() {
		toSerialize["ephemeralBranchWhitelistRegex"] = o.EphemeralBranchWhitelistRegex.Get()
	}
	return toSerialize, nil
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
