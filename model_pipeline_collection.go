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

// checks if the PipelineCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineCollection{}

// PipelineCollection A pipeline.
type PipelineCollection struct {
	// Pipeline identifier.
	Id *string `json:"id,omitempty"`
	// Pipeline description.
	Description *string `json:"description,omitempty"`
	// Pipeline status.
	Status *string `json:"status,omitempty"`
	// Environment identifier.
	Environment NullableString `json:"environment,omitempty"`
	// Event identifier.
	Event *string `json:"event,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewPipelineCollection instantiates a new PipelineCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCollection() *PipelineCollection {
	this := PipelineCollection{}
	return &this
}

// NewPipelineCollectionWithDefaults instantiates a new PipelineCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCollectionWithDefaults() *PipelineCollection {
	this := PipelineCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PipelineCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PipelineCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PipelineCollection) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PipelineCollection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PipelineCollection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PipelineCollection) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PipelineCollection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCollection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PipelineCollection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PipelineCollection) SetStatus(v string) {
	o.Status = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PipelineCollection) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCollection) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PipelineCollection) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *PipelineCollection) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *PipelineCollection) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *PipelineCollection) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *PipelineCollection) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCollection) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *PipelineCollection) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *PipelineCollection) SetEvent(v string) {
	o.Event = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PipelineCollection) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCollection) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PipelineCollection) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *PipelineCollection) SetOrganization(v string) {
	o.Organization = &v
}

func (o PipelineCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineCollection) ToMap() (map[string]interface{}, error) {
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

type NullablePipelineCollection struct {
	value *PipelineCollection
	isSet bool
}

func (v NullablePipelineCollection) Get() *PipelineCollection {
	return v.value
}

func (v *NullablePipelineCollection) Set(val *PipelineCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCollection(val *PipelineCollection) *NullablePipelineCollection {
	return &NullablePipelineCollection{value: val, isSet: true}
}

func (v NullablePipelineCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
