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

// checks if the PipelineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineItem{}

// PipelineItem A pipeline.
type PipelineItem struct {
	// Pipeline identifier.
	Id *string `json:"id,omitempty"`
	// Pipeline description.
	Description *string `json:"description,omitempty"`
	// Pipeline status.
	Status *string `json:"status,omitempty"`
	// Stage identifier.
	Stages []StageItem `json:"stages,omitempty"`
	// Environment identifier.
	Environment NullableString `json:"environment,omitempty"`
	// Event identifier.
	Event *string `json:"event,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewPipelineItem instantiates a new PipelineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineItem() *PipelineItem {
	this := PipelineItem{}
	return &this
}

// NewPipelineItemWithDefaults instantiates a new PipelineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineItemWithDefaults() *PipelineItem {
	this := PipelineItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PipelineItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PipelineItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PipelineItem) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PipelineItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PipelineItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PipelineItem) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PipelineItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PipelineItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PipelineItem) SetStatus(v string) {
	o.Status = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *PipelineItem) GetStages() []StageItem {
	if o == nil || IsNil(o.Stages) {
		var ret []StageItem
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetStagesOk() ([]StageItem, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *PipelineItem) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []StageItem and assigns it to the Stages field.
func (o *PipelineItem) SetStages(v []StageItem) {
	o.Stages = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PipelineItem) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineItem) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PipelineItem) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *PipelineItem) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *PipelineItem) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *PipelineItem) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *PipelineItem) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *PipelineItem) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *PipelineItem) SetEvent(v string) {
	o.Event = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PipelineItem) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineItem) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PipelineItem) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *PipelineItem) SetOrganization(v string) {
	o.Organization = &v
}

func (o PipelineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullablePipelineItem struct {
	value *PipelineItem
	isSet bool
}

func (v NullablePipelineItem) Get() *PipelineItem {
	return v.value
}

func (v *NullablePipelineItem) Set(val *PipelineItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineItem(val *PipelineItem) *NullablePipelineItem {
	return &NullablePipelineItem{value: val, isSet: true}
}

func (v NullablePipelineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
