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

// checks if the ProjectEditBuildSettingsAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEditBuildSettingsAction{}

// ProjectEditBuildSettingsAction A project holds multiple environments and shared secrets and settings.
type ProjectEditBuildSettingsAction struct {
	UseManagedRegistry    NullableBool   `json:"useManagedRegistry,omitempty"`
	RegistryIntegration   NullableString `json:"registryIntegration"`
	UseManagedCluster     NullableBool   `json:"useManagedCluster,omitempty"`
	KubernetesIntegration NullableString `json:"kubernetesIntegration"`
	BuildEngine           NullableString `json:"buildEngine,omitempty"`
	// K8s supports decimal values with step 0.001 That's why we are using decimal.
	Cpu NullableString `json:"cpu"`
	// expressed in Mi
	Memory                NullableInt32 `json:"memory"`
	TimeoutSeconds        NullableInt32 `json:"timeoutSeconds"`
	NamespaceCustomLabels []string      `json:"namespaceCustomLabels,omitempty"`
}

// NewProjectEditBuildSettingsAction instantiates a new ProjectEditBuildSettingsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEditBuildSettingsAction(registryIntegration NullableString, kubernetesIntegration NullableString, cpu NullableString, memory NullableInt32, timeoutSeconds NullableInt32) *ProjectEditBuildSettingsAction {
	this := ProjectEditBuildSettingsAction{}
	this.RegistryIntegration = registryIntegration
	this.KubernetesIntegration = kubernetesIntegration
	this.Cpu = cpu
	this.Memory = memory
	this.TimeoutSeconds = timeoutSeconds
	return &this
}

// NewProjectEditBuildSettingsActionWithDefaults instantiates a new ProjectEditBuildSettingsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEditBuildSettingsActionWithDefaults() *ProjectEditBuildSettingsAction {
	this := ProjectEditBuildSettingsAction{}
	return &this
}

// GetUseManagedRegistry returns the UseManagedRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEditBuildSettingsAction) GetUseManagedRegistry() bool {
	if o == nil || IsNil(o.UseManagedRegistry.Get()) {
		var ret bool
		return ret
	}
	return *o.UseManagedRegistry.Get()
}

// GetUseManagedRegistryOk returns a tuple with the UseManagedRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetUseManagedRegistryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseManagedRegistry.Get(), o.UseManagedRegistry.IsSet()
}

// HasUseManagedRegistry returns a boolean if a field has been set.
func (o *ProjectEditBuildSettingsAction) HasUseManagedRegistry() bool {
	if o != nil && o.UseManagedRegistry.IsSet() {
		return true
	}

	return false
}

// SetUseManagedRegistry gets a reference to the given NullableBool and assigns it to the UseManagedRegistry field.
func (o *ProjectEditBuildSettingsAction) SetUseManagedRegistry(v bool) {
	o.UseManagedRegistry.Set(&v)
}

// SetUseManagedRegistryNil sets the value for UseManagedRegistry to be an explicit nil
func (o *ProjectEditBuildSettingsAction) SetUseManagedRegistryNil() {
	o.UseManagedRegistry.Set(nil)
}

// UnsetUseManagedRegistry ensures that no value is present for UseManagedRegistry, not even an explicit nil
func (o *ProjectEditBuildSettingsAction) UnsetUseManagedRegistry() {
	o.UseManagedRegistry.Unset()
}

// GetRegistryIntegration returns the RegistryIntegration field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectEditBuildSettingsAction) GetRegistryIntegration() string {
	if o == nil || o.RegistryIntegration.Get() == nil {
		var ret string
		return ret
	}

	return *o.RegistryIntegration.Get()
}

// GetRegistryIntegrationOk returns a tuple with the RegistryIntegration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetRegistryIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistryIntegration.Get(), o.RegistryIntegration.IsSet()
}

// SetRegistryIntegration sets field value
func (o *ProjectEditBuildSettingsAction) SetRegistryIntegration(v string) {
	o.RegistryIntegration.Set(&v)
}

// GetUseManagedCluster returns the UseManagedCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEditBuildSettingsAction) GetUseManagedCluster() bool {
	if o == nil || IsNil(o.UseManagedCluster.Get()) {
		var ret bool
		return ret
	}
	return *o.UseManagedCluster.Get()
}

// GetUseManagedClusterOk returns a tuple with the UseManagedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetUseManagedClusterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseManagedCluster.Get(), o.UseManagedCluster.IsSet()
}

// HasUseManagedCluster returns a boolean if a field has been set.
func (o *ProjectEditBuildSettingsAction) HasUseManagedCluster() bool {
	if o != nil && o.UseManagedCluster.IsSet() {
		return true
	}

	return false
}

// SetUseManagedCluster gets a reference to the given NullableBool and assigns it to the UseManagedCluster field.
func (o *ProjectEditBuildSettingsAction) SetUseManagedCluster(v bool) {
	o.UseManagedCluster.Set(&v)
}

// SetUseManagedClusterNil sets the value for UseManagedCluster to be an explicit nil
func (o *ProjectEditBuildSettingsAction) SetUseManagedClusterNil() {
	o.UseManagedCluster.Set(nil)
}

// UnsetUseManagedCluster ensures that no value is present for UseManagedCluster, not even an explicit nil
func (o *ProjectEditBuildSettingsAction) UnsetUseManagedCluster() {
	o.UseManagedCluster.Unset()
}

// GetKubernetesIntegration returns the KubernetesIntegration field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectEditBuildSettingsAction) GetKubernetesIntegration() string {
	if o == nil || o.KubernetesIntegration.Get() == nil {
		var ret string
		return ret
	}

	return *o.KubernetesIntegration.Get()
}

// GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesIntegration.Get(), o.KubernetesIntegration.IsSet()
}

// SetKubernetesIntegration sets field value
func (o *ProjectEditBuildSettingsAction) SetKubernetesIntegration(v string) {
	o.KubernetesIntegration.Set(&v)
}

// GetBuildEngine returns the BuildEngine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEditBuildSettingsAction) GetBuildEngine() string {
	if o == nil || IsNil(o.BuildEngine.Get()) {
		var ret string
		return ret
	}
	return *o.BuildEngine.Get()
}

// GetBuildEngineOk returns a tuple with the BuildEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetBuildEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildEngine.Get(), o.BuildEngine.IsSet()
}

// HasBuildEngine returns a boolean if a field has been set.
func (o *ProjectEditBuildSettingsAction) HasBuildEngine() bool {
	if o != nil && o.BuildEngine.IsSet() {
		return true
	}

	return false
}

// SetBuildEngine gets a reference to the given NullableString and assigns it to the BuildEngine field.
func (o *ProjectEditBuildSettingsAction) SetBuildEngine(v string) {
	o.BuildEngine.Set(&v)
}

// SetBuildEngineNil sets the value for BuildEngine to be an explicit nil
func (o *ProjectEditBuildSettingsAction) SetBuildEngineNil() {
	o.BuildEngine.Set(nil)
}

// UnsetBuildEngine ensures that no value is present for BuildEngine, not even an explicit nil
func (o *ProjectEditBuildSettingsAction) UnsetBuildEngine() {
	o.BuildEngine.Unset()
}

// GetCpu returns the Cpu field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectEditBuildSettingsAction) GetCpu() string {
	if o == nil || o.Cpu.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// SetCpu sets field value
func (o *ProjectEditBuildSettingsAction) SetCpu(v string) {
	o.Cpu.Set(&v)
}

// GetMemory returns the Memory field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectEditBuildSettingsAction) GetMemory() int32 {
	if o == nil || o.Memory.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// SetMemory sets field value
func (o *ProjectEditBuildSettingsAction) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// GetTimeoutSeconds returns the TimeoutSeconds field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ProjectEditBuildSettingsAction) GetTimeoutSeconds() int32 {
	if o == nil || o.TimeoutSeconds.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TimeoutSeconds.Get()
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeoutSeconds.Get(), o.TimeoutSeconds.IsSet()
}

// SetTimeoutSeconds sets field value
func (o *ProjectEditBuildSettingsAction) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds.Set(&v)
}

// GetNamespaceCustomLabels returns the NamespaceCustomLabels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEditBuildSettingsAction) GetNamespaceCustomLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NamespaceCustomLabels
}

// GetNamespaceCustomLabelsOk returns a tuple with the NamespaceCustomLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEditBuildSettingsAction) GetNamespaceCustomLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.NamespaceCustomLabels) {
		return nil, false
	}
	return o.NamespaceCustomLabels, true
}

// HasNamespaceCustomLabels returns a boolean if a field has been set.
func (o *ProjectEditBuildSettingsAction) HasNamespaceCustomLabels() bool {
	if o != nil && IsNil(o.NamespaceCustomLabels) {
		return true
	}

	return false
}

// SetNamespaceCustomLabels gets a reference to the given []string and assigns it to the NamespaceCustomLabels field.
func (o *ProjectEditBuildSettingsAction) SetNamespaceCustomLabels(v []string) {
	o.NamespaceCustomLabels = v
}

func (o ProjectEditBuildSettingsAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEditBuildSettingsAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UseManagedRegistry.IsSet() {
		toSerialize["useManagedRegistry"] = o.UseManagedRegistry.Get()
	}
	toSerialize["registryIntegration"] = o.RegistryIntegration.Get()
	if o.UseManagedCluster.IsSet() {
		toSerialize["useManagedCluster"] = o.UseManagedCluster.Get()
	}
	toSerialize["kubernetesIntegration"] = o.KubernetesIntegration.Get()
	if o.BuildEngine.IsSet() {
		toSerialize["buildEngine"] = o.BuildEngine.Get()
	}
	toSerialize["cpu"] = o.Cpu.Get()
	toSerialize["memory"] = o.Memory.Get()
	toSerialize["timeoutSeconds"] = o.TimeoutSeconds.Get()
	if o.NamespaceCustomLabels != nil {
		toSerialize["namespaceCustomLabels"] = o.NamespaceCustomLabels
	}
	return toSerialize, nil
}

type NullableProjectEditBuildSettingsAction struct {
	value *ProjectEditBuildSettingsAction
	isSet bool
}

func (v NullableProjectEditBuildSettingsAction) Get() *ProjectEditBuildSettingsAction {
	return v.value
}

func (v *NullableProjectEditBuildSettingsAction) Set(val *ProjectEditBuildSettingsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEditBuildSettingsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEditBuildSettingsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEditBuildSettingsAction(val *ProjectEditBuildSettingsAction) *NullableProjectEditBuildSettingsAction {
	return &NullableProjectEditBuildSettingsAction{value: val, isSet: true}
}

func (v NullableProjectEditBuildSettingsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEditBuildSettingsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
