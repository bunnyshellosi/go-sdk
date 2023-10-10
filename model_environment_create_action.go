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

// checks if the EnvironmentCreateAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentCreateAction{}

// EnvironmentCreateAction An environment holds a collection of buildable and deployable components.
type EnvironmentCreateAction struct {
	Name                           string                          `json:"name"`
	Project                        string                          `json:"project"`
	Genesis                        *EnvironmentCreateActionGenesis `json:"genesis,omitempty"`
	Type                           *string                         `json:"type,omitempty"`
	RemoteDevelopmentAllowed       *bool                           `json:"remoteDevelopmentAllowed,omitempty"`
	AutoUpdate                     *bool                           `json:"autoUpdate,omitempty"`
	CreateEphemeralOnPrCreate      *bool                           `json:"createEphemeralOnPrCreate,omitempty"`
	DestroyEphemeralOnPrClose      *bool                           `json:"destroyEphemeralOnPrClose,omitempty"`
	KubernetesIntegration          NullableString                  `json:"kubernetesIntegration,omitempty"`
	EphemeralKubernetesIntegration NullableString                  `json:"ephemeralKubernetesIntegration,omitempty"`
	Labels                         *map[string]string              `json:"labels,omitempty"`
}

// NewEnvironmentCreateAction instantiates a new EnvironmentCreateAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentCreateAction(name string, project string) *EnvironmentCreateAction {
	this := EnvironmentCreateAction{}
	this.Name = name
	this.Project = project
	return &this
}

// NewEnvironmentCreateActionWithDefaults instantiates a new EnvironmentCreateAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentCreateActionWithDefaults() *EnvironmentCreateAction {
	this := EnvironmentCreateAction{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentCreateAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentCreateAction) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value
func (o *EnvironmentCreateAction) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *EnvironmentCreateAction) SetProject(v string) {
	o.Project = v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetGenesis() EnvironmentCreateActionGenesis {
	if o == nil || IsNil(o.Genesis) {
		var ret EnvironmentCreateActionGenesis
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetGenesisOk() (*EnvironmentCreateActionGenesis, bool) {
	if o == nil || IsNil(o.Genesis) {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasGenesis() bool {
	if o != nil && !IsNil(o.Genesis) {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given EnvironmentCreateActionGenesis and assigns it to the Genesis field.
func (o *EnvironmentCreateAction) SetGenesis(v EnvironmentCreateActionGenesis) {
	o.Genesis = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentCreateAction) SetType(v string) {
	o.Type = &v
}

// GetRemoteDevelopmentAllowed returns the RemoteDevelopmentAllowed field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetRemoteDevelopmentAllowed() bool {
	if o == nil || IsNil(o.RemoteDevelopmentAllowed) {
		var ret bool
		return ret
	}
	return *o.RemoteDevelopmentAllowed
}

// GetRemoteDevelopmentAllowedOk returns a tuple with the RemoteDevelopmentAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetRemoteDevelopmentAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteDevelopmentAllowed) {
		return nil, false
	}
	return o.RemoteDevelopmentAllowed, true
}

// HasRemoteDevelopmentAllowed returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasRemoteDevelopmentAllowed() bool {
	if o != nil && !IsNil(o.RemoteDevelopmentAllowed) {
		return true
	}

	return false
}

// SetRemoteDevelopmentAllowed gets a reference to the given bool and assigns it to the RemoteDevelopmentAllowed field.
func (o *EnvironmentCreateAction) SetRemoteDevelopmentAllowed(v bool) {
	o.RemoteDevelopmentAllowed = &v
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetAutoUpdate() bool {
	if o == nil || IsNil(o.AutoUpdate) {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpdate) {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasAutoUpdate() bool {
	if o != nil && !IsNil(o.AutoUpdate) {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *EnvironmentCreateAction) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetCreateEphemeralOnPrCreate returns the CreateEphemeralOnPrCreate field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetCreateEphemeralOnPrCreate() bool {
	if o == nil || IsNil(o.CreateEphemeralOnPrCreate) {
		var ret bool
		return ret
	}
	return *o.CreateEphemeralOnPrCreate
}

// GetCreateEphemeralOnPrCreateOk returns a tuple with the CreateEphemeralOnPrCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetCreateEphemeralOnPrCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateEphemeralOnPrCreate) {
		return nil, false
	}
	return o.CreateEphemeralOnPrCreate, true
}

// HasCreateEphemeralOnPrCreate returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasCreateEphemeralOnPrCreate() bool {
	if o != nil && !IsNil(o.CreateEphemeralOnPrCreate) {
		return true
	}

	return false
}

// SetCreateEphemeralOnPrCreate gets a reference to the given bool and assigns it to the CreateEphemeralOnPrCreate field.
func (o *EnvironmentCreateAction) SetCreateEphemeralOnPrCreate(v bool) {
	o.CreateEphemeralOnPrCreate = &v
}

// GetDestroyEphemeralOnPrClose returns the DestroyEphemeralOnPrClose field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetDestroyEphemeralOnPrClose() bool {
	if o == nil || IsNil(o.DestroyEphemeralOnPrClose) {
		var ret bool
		return ret
	}
	return *o.DestroyEphemeralOnPrClose
}

// GetDestroyEphemeralOnPrCloseOk returns a tuple with the DestroyEphemeralOnPrClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetDestroyEphemeralOnPrCloseOk() (*bool, bool) {
	if o == nil || IsNil(o.DestroyEphemeralOnPrClose) {
		return nil, false
	}
	return o.DestroyEphemeralOnPrClose, true
}

// HasDestroyEphemeralOnPrClose returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasDestroyEphemeralOnPrClose() bool {
	if o != nil && !IsNil(o.DestroyEphemeralOnPrClose) {
		return true
	}

	return false
}

// SetDestroyEphemeralOnPrClose gets a reference to the given bool and assigns it to the DestroyEphemeralOnPrClose field.
func (o *EnvironmentCreateAction) SetDestroyEphemeralOnPrClose(v bool) {
	o.DestroyEphemeralOnPrClose = &v
}

// GetKubernetesIntegration returns the KubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCreateAction) GetKubernetesIntegration() string {
	if o == nil || IsNil(o.KubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.KubernetesIntegration.Get()
}

// GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentCreateAction) GetKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesIntegration.Get(), o.KubernetesIntegration.IsSet()
}

// HasKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasKubernetesIntegration() bool {
	if o != nil && o.KubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetKubernetesIntegration gets a reference to the given NullableString and assigns it to the KubernetesIntegration field.
func (o *EnvironmentCreateAction) SetKubernetesIntegration(v string) {
	o.KubernetesIntegration.Set(&v)
}

// SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil
func (o *EnvironmentCreateAction) SetKubernetesIntegrationNil() {
	o.KubernetesIntegration.Set(nil)
}

// UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
func (o *EnvironmentCreateAction) UnsetKubernetesIntegration() {
	o.KubernetesIntegration.Unset()
}

// GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCreateAction) GetEphemeralKubernetesIntegration() string {
	if o == nil || IsNil(o.EphemeralKubernetesIntegration.Get()) {
		var ret string
		return ret
	}
	return *o.EphemeralKubernetesIntegration.Get()
}

// GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentCreateAction) GetEphemeralKubernetesIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EphemeralKubernetesIntegration.Get(), o.EphemeralKubernetesIntegration.IsSet()
}

// HasEphemeralKubernetesIntegration returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasEphemeralKubernetesIntegration() bool {
	if o != nil && o.EphemeralKubernetesIntegration.IsSet() {
		return true
	}

	return false
}

// SetEphemeralKubernetesIntegration gets a reference to the given NullableString and assigns it to the EphemeralKubernetesIntegration field.
func (o *EnvironmentCreateAction) SetEphemeralKubernetesIntegration(v string) {
	o.EphemeralKubernetesIntegration.Set(&v)
}

// SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil
func (o *EnvironmentCreateAction) SetEphemeralKubernetesIntegrationNil() {
	o.EphemeralKubernetesIntegration.Set(nil)
}

// UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
func (o *EnvironmentCreateAction) UnsetEphemeralKubernetesIntegration() {
	o.EphemeralKubernetesIntegration.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *EnvironmentCreateAction) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCreateAction) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EnvironmentCreateAction) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *EnvironmentCreateAction) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o EnvironmentCreateAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentCreateAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["project"] = o.Project
	if !IsNil(o.Genesis) {
		toSerialize["genesis"] = o.Genesis
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RemoteDevelopmentAllowed) {
		toSerialize["remoteDevelopmentAllowed"] = o.RemoteDevelopmentAllowed
	}
	if !IsNil(o.AutoUpdate) {
		toSerialize["autoUpdate"] = o.AutoUpdate
	}
	if !IsNil(o.CreateEphemeralOnPrCreate) {
		toSerialize["createEphemeralOnPrCreate"] = o.CreateEphemeralOnPrCreate
	}
	if !IsNil(o.DestroyEphemeralOnPrClose) {
		toSerialize["destroyEphemeralOnPrClose"] = o.DestroyEphemeralOnPrClose
	}
	if o.KubernetesIntegration.IsSet() {
		toSerialize["kubernetesIntegration"] = o.KubernetesIntegration.Get()
	}
	if o.EphemeralKubernetesIntegration.IsSet() {
		toSerialize["ephemeralKubernetesIntegration"] = o.EphemeralKubernetesIntegration.Get()
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableEnvironmentCreateAction struct {
	value *EnvironmentCreateAction
	isSet bool
}

func (v NullableEnvironmentCreateAction) Get() *EnvironmentCreateAction {
	return v.value
}

func (v *NullableEnvironmentCreateAction) Set(val *EnvironmentCreateAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentCreateAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentCreateAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentCreateAction(val *EnvironmentCreateAction) *NullableEnvironmentCreateAction {
	return &NullableEnvironmentCreateAction{value: val, isSet: true}
}

func (v NullableEnvironmentCreateAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentCreateAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
