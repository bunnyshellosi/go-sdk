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

// checks if the EnvironmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentItem{}

// EnvironmentItem An environment holds a collection of buildable and deployable components.
type EnvironmentItem struct {
	// Environment identifier.
	Id *string `json:"id,omitempty"`
	// Environment labels.
	Labels *map[string]string `json:"labels,omitempty"`
	// Environment type: primary or ephemeral.
	Type *string `json:"type,omitempty"`
	// Environment name.
	Name *string `json:"name,omitempty"`
	// Environment k8s namespace.
	Namespace *string `json:"namespace,omitempty"`
	// Service component identifier
	TotalComponents *int32 `json:"totalComponents,omitempty"`
	// Environment operation status.
	OperationStatus *string `json:"operationStatus,omitempty"`
	// Environment cluster status.
	ClusterStatus *string                   `json:"clusterStatus,omitempty"`
	BuildSettings NullableBuildSettingsItem `json:"buildSettings,omitempty"`
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
}

// NewEnvironmentItem instantiates a new EnvironmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentItem() *EnvironmentItem {
	this := EnvironmentItem{}
	return &this
}

// NewEnvironmentItemWithDefaults instantiates a new EnvironmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentItemWithDefaults() *EnvironmentItem {
	this := EnvironmentItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentItem) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *EnvironmentItem) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EnvironmentItem) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *EnvironmentItem) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentItem) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentItem) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *EnvironmentItem) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EnvironmentItem) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *EnvironmentItem) SetNamespace(v string) {
	o.Namespace = &v
}

// GetTotalComponents returns the TotalComponents field value if set, zero value otherwise.
func (o *EnvironmentItem) GetTotalComponents() int32 {
	if o == nil || IsNil(o.TotalComponents) {
		var ret int32
		return ret
	}
	return *o.TotalComponents
}

// GetTotalComponentsOk returns a tuple with the TotalComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetTotalComponentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalComponents) {
		return nil, false
	}
	return o.TotalComponents, true
}

// HasTotalComponents returns a boolean if a field has been set.
func (o *EnvironmentItem) HasTotalComponents() bool {
	if o != nil && !IsNil(o.TotalComponents) {
		return true
	}

	return false
}

// SetTotalComponents gets a reference to the given int32 and assigns it to the TotalComponents field.
func (o *EnvironmentItem) SetTotalComponents(v int32) {
	o.TotalComponents = &v
}

// GetOperationStatus returns the OperationStatus field value if set, zero value otherwise.
func (o *EnvironmentItem) GetOperationStatus() string {
	if o == nil || IsNil(o.OperationStatus) {
		var ret string
		return ret
	}
	return *o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetOperationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OperationStatus) {
		return nil, false
	}
	return o.OperationStatus, true
}

// HasOperationStatus returns a boolean if a field has been set.
func (o *EnvironmentItem) HasOperationStatus() bool {
	if o != nil && !IsNil(o.OperationStatus) {
		return true
	}

	return false
}

// SetOperationStatus gets a reference to the given string and assigns it to the OperationStatus field.
func (o *EnvironmentItem) SetOperationStatus(v string) {
	o.OperationStatus = &v
}

// GetClusterStatus returns the ClusterStatus field value if set, zero value otherwise.
func (o *EnvironmentItem) GetClusterStatus() string {
	if o == nil || IsNil(o.ClusterStatus) {
		var ret string
		return ret
	}
	return *o.ClusterStatus
}

// GetClusterStatusOk returns a tuple with the ClusterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetClusterStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterStatus) {
		return nil, false
	}
	return o.ClusterStatus, true
}

// HasClusterStatus returns a boolean if a field has been set.
func (o *EnvironmentItem) HasClusterStatus() bool {
	if o != nil && !IsNil(o.ClusterStatus) {
		return true
	}

	return false
}

// SetClusterStatus gets a reference to the given string and assigns it to the ClusterStatus field.
func (o *EnvironmentItem) SetClusterStatus(v string) {
	o.ClusterStatus = &v
}

// GetBuildSettings returns the BuildSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentItem) GetBuildSettings() BuildSettingsItem {
	if o == nil || IsNil(o.BuildSettings.Get()) {
		var ret BuildSettingsItem
		return ret
	}
	return *o.BuildSettings.Get()
}

// GetBuildSettingsOk returns a tuple with the BuildSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentItem) GetBuildSettingsOk() (*BuildSettingsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildSettings.Get(), o.BuildSettings.IsSet()
}

// HasBuildSettings returns a boolean if a field has been set.
func (o *EnvironmentItem) HasBuildSettings() bool {
	if o != nil && o.BuildSettings.IsSet() {
		return true
	}

	return false
}

// SetBuildSettings gets a reference to the given NullableBuildSettingsItem and assigns it to the BuildSettings field.
func (o *EnvironmentItem) SetBuildSettings(v BuildSettingsItem) {
	o.BuildSettings.Set(&v)
}

// SetBuildSettingsNil sets the value for BuildSettings to be an explicit nil
func (o *EnvironmentItem) SetBuildSettingsNil() {
	o.BuildSettings.Set(nil)
}

// UnsetBuildSettings ensures that no value is present for BuildSettings, not even an explicit nil
func (o *EnvironmentItem) UnsetBuildSettings() {
	o.BuildSettings.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *EnvironmentItem) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentItem) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *EnvironmentItem) SetProject(v string) {
	o.Project = &v
}

// GetKubernetesIntegration returns the KubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentItem) GetKubernetesIntegration() string {
	if o == nil || IsNil(o.KubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesIntegration.Get()
}

// GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentItem) GetKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesIntegration.Get(), o.KubernetesIntegration.IsSet()
}

// HasKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentItem) HasKubernetesIntegration() bool {
	if o != nil && o.KubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetKubernetesIntegration gets a reference to the given NullableString and assigns it to the KubernetesIntegration field.
func (o *EnvironmentItem) SetKubernetesIntegration(v string) {
	o.KubernetesIntegration.Set(&v)
}

// SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil
func (o *EnvironmentItem) SetKubernetesIntegrationNil() {
	o.KubernetesIntegration.Set(nil)
}

// UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
func (o *EnvironmentItem) UnsetKubernetesIntegration() {
	o.KubernetesIntegration.Unset()
}

// GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentItem) GetEphemeralKubernetesIntegration() string {
	if o == nil || IsNil(o.EphemeralKubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.EphemeralKubernetesIntegration.Get()
}

// GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentItem) GetEphemeralKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EphemeralKubernetesIntegration.Get(), o.EphemeralKubernetesIntegration.IsSet()
}

// HasEphemeralKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentItem) HasEphemeralKubernetesIntegration() bool {
	if o != nil && o.EphemeralKubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetEphemeralKubernetesIntegration gets a reference to the given NullableString and assigns it to the EphemeralKubernetesIntegration field.
func (o *EnvironmentItem) SetEphemeralKubernetesIntegration(v string) {
	o.EphemeralKubernetesIntegration.Set(&v)
}

// SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil
func (o *EnvironmentItem) SetEphemeralKubernetesIntegrationNil() {
	o.EphemeralKubernetesIntegration.Set(nil)
}

// UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
func (o *EnvironmentItem) UnsetEphemeralKubernetesIntegration() {
	o.EphemeralKubernetesIntegration.Unset()
}

// GetHasTerminationProtection returns the HasTerminationProtection field value if set, zero value otherwise.
func (o *EnvironmentItem) GetHasTerminationProtection() bool {
	if o == nil || IsNil(o.HasTerminationProtection) {
		var ret bool
		return ret
	}
	return *o.HasTerminationProtection
}

// GetHasTerminationProtectionOk returns a tuple with the HasTerminationProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetHasTerminationProtectionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTerminationProtection) {
		return nil, false
	}
	return o.HasTerminationProtection, true
}

// HasHasTerminationProtection returns a boolean if a field has been set.
func (o *EnvironmentItem) HasHasTerminationProtection() bool {
	if o != nil && !IsNil(o.HasTerminationProtection) {
		return true
	}

	return false
}

// SetHasTerminationProtection gets a reference to the given bool and assigns it to the HasTerminationProtection field.
func (o *EnvironmentItem) SetHasTerminationProtection(v bool) {
	o.HasTerminationProtection = &v
}

// GetHasEphemeralAutoDeploy returns the HasEphemeralAutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentItem) GetHasEphemeralAutoDeploy() bool {
	if o == nil || IsNil(o.HasEphemeralAutoDeploy) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralAutoDeploy
}

// GetHasEphemeralAutoDeployOk returns a tuple with the HasEphemeralAutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetHasEphemeralAutoDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralAutoDeploy) {
		return nil, false
	}
	return o.HasEphemeralAutoDeploy, true
}

// HasHasEphemeralAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentItem) HasHasEphemeralAutoDeploy() bool {
	if o != nil && !IsNil(o.HasEphemeralAutoDeploy) {
		return true
	}

	return false
}

// SetHasEphemeralAutoDeploy gets a reference to the given bool and assigns it to the HasEphemeralAutoDeploy field.
func (o *EnvironmentItem) SetHasEphemeralAutoDeploy(v bool) {
	o.HasEphemeralAutoDeploy = &v
}

// GetHasEphemeralCreateOnPr returns the HasEphemeralCreateOnPr field value if set, zero value otherwise.
func (o *EnvironmentItem) GetHasEphemeralCreateOnPr() bool {
	if o == nil || IsNil(o.HasEphemeralCreateOnPr) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralCreateOnPr
}

// GetHasEphemeralCreateOnPrOk returns a tuple with the HasEphemeralCreateOnPr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetHasEphemeralCreateOnPrOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralCreateOnPr) {
		return nil, false
	}
	return o.HasEphemeralCreateOnPr, true
}

// HasHasEphemeralCreateOnPr returns a boolean if a field has been set.
func (o *EnvironmentItem) HasHasEphemeralCreateOnPr() bool {
	if o != nil && !IsNil(o.HasEphemeralCreateOnPr) {
		return true
	}

	return false
}

// SetHasEphemeralCreateOnPr gets a reference to the given bool and assigns it to the HasEphemeralCreateOnPr field.
func (o *EnvironmentItem) SetHasEphemeralCreateOnPr(v bool) {
	o.HasEphemeralCreateOnPr = &v
}

// GetHasEphemeralDestroyOnPrClose returns the HasEphemeralDestroyOnPrClose field value if set, zero value otherwise.
func (o *EnvironmentItem) GetHasEphemeralDestroyOnPrClose() bool {
	if o == nil || IsNil(o.HasEphemeralDestroyOnPrClose) {
		var ret bool
		return ret
	}
	return *o.HasEphemeralDestroyOnPrClose
}

// GetHasEphemeralDestroyOnPrCloseOk returns a tuple with the HasEphemeralDestroyOnPrClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentItem) GetHasEphemeralDestroyOnPrCloseOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEphemeralDestroyOnPrClose) {
		return nil, false
	}
	return o.HasEphemeralDestroyOnPrClose, true
}

// HasHasEphemeralDestroyOnPrClose returns a boolean if a field has been set.
func (o *EnvironmentItem) HasHasEphemeralDestroyOnPrClose() bool {
	if o != nil && !IsNil(o.HasEphemeralDestroyOnPrClose) {
		return true
	}

	return false
}

// SetHasEphemeralDestroyOnPrClose gets a reference to the given bool and assigns it to the HasEphemeralDestroyOnPrClose field.
func (o *EnvironmentItem) SetHasEphemeralDestroyOnPrClose(v bool) {
	o.HasEphemeralDestroyOnPrClose = &v
}

func (o EnvironmentItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
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
	if !IsNil(o.TotalComponents) {
		toSerialize["totalComponents"] = o.TotalComponents
	}
	if !IsNil(o.OperationStatus) {
		toSerialize["operationStatus"] = o.OperationStatus
	}
	if !IsNil(o.ClusterStatus) {
		toSerialize["clusterStatus"] = o.ClusterStatus
	}
	if o.BuildSettings.IsSet() {
		toSerialize["buildSettings"] = o.BuildSettings.Get()
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
	return toSerialize, nil
}

type NullableEnvironmentItem struct {
	value *EnvironmentItem
	isSet bool
}

func (v NullableEnvironmentItem) Get() *EnvironmentItem {
	return v.value
}

func (v *NullableEnvironmentItem) Set(val *EnvironmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentItem(val *EnvironmentItem) *NullableEnvironmentItem {
	return &NullableEnvironmentItem{value: val, isSet: true}
}

func (v NullableEnvironmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
