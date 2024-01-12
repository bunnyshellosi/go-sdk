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

// checks if the WorkflowItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowItem{}

// WorkflowItem A workflow.
type WorkflowItem struct {
	// Workflow identifier.
	Id *string `json:"id,omitempty"`
	// Workflow description.
	Description *string `json:"description,omitempty"`
	// Workflow status.
	Status *string `json:"status,omitempty"`
	// Environment identifier.
	Environment NullableString `json:"environment,omitempty"`
	// Event identifier.
	Event *string `json:"event,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
	// Workflow duration in seconds.
	Duration NullableInt32 `json:"duration,omitempty"`
	// Number of jobs.
	JobsCount *int32 `json:"jobsCount,omitempty"`
	// Number of completed jobs.
	CompletedJobsCount *int32 `json:"completedJobsCount,omitempty"`
	// Workflow web URL.
	WebUrl *string `json:"webUrl,omitempty"`
}

// NewWorkflowItem instantiates a new WorkflowItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowItem() *WorkflowItem {
	this := WorkflowItem{}
	return &this
}

// NewWorkflowItemWithDefaults instantiates a new WorkflowItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowItemWithDefaults() *WorkflowItem {
	this := WorkflowItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowItem) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowItem) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowItem) SetStatus(v string) {
	o.Status = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowItem) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowItem) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *WorkflowItem) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *WorkflowItem) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *WorkflowItem) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *WorkflowItem) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *WorkflowItem) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *WorkflowItem) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *WorkflowItem) SetEvent(v string) {
	o.Event = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowItem) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowItem) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *WorkflowItem) SetOrganization(v string) {
	o.Organization = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowItem) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowItem) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkflowItem) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *WorkflowItem) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil
func (o *WorkflowItem) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *WorkflowItem) UnsetDuration() {
	o.Duration.Unset()
}

// GetJobsCount returns the JobsCount field value if set, zero value otherwise.
func (o *WorkflowItem) GetJobsCount() int32 {
	if o == nil || IsNil(o.JobsCount) {
		var ret int32
		return ret
	}
	return *o.JobsCount
}

// GetJobsCountOk returns a tuple with the JobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetJobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.JobsCount) {
		return nil, false
	}
	return o.JobsCount, true
}

// HasJobsCount returns a boolean if a field has been set.
func (o *WorkflowItem) HasJobsCount() bool {
	if o != nil && !IsNil(o.JobsCount) {
		return true
	}

	return false
}

// SetJobsCount gets a reference to the given int32 and assigns it to the JobsCount field.
func (o *WorkflowItem) SetJobsCount(v int32) {
	o.JobsCount = &v
}

// GetCompletedJobsCount returns the CompletedJobsCount field value if set, zero value otherwise.
func (o *WorkflowItem) GetCompletedJobsCount() int32 {
	if o == nil || IsNil(o.CompletedJobsCount) {
		var ret int32
		return ret
	}
	return *o.CompletedJobsCount
}

// GetCompletedJobsCountOk returns a tuple with the CompletedJobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetCompletedJobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CompletedJobsCount) {
		return nil, false
	}
	return o.CompletedJobsCount, true
}

// HasCompletedJobsCount returns a boolean if a field has been set.
func (o *WorkflowItem) HasCompletedJobsCount() bool {
	if o != nil && !IsNil(o.CompletedJobsCount) {
		return true
	}

	return false
}

// SetCompletedJobsCount gets a reference to the given int32 and assigns it to the CompletedJobsCount field.
func (o *WorkflowItem) SetCompletedJobsCount(v int32) {
	o.CompletedJobsCount = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *WorkflowItem) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowItem) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *WorkflowItem) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *WorkflowItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o WorkflowItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowItem) ToMap() (map[string]interface{}, error) {
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
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if !IsNil(o.JobsCount) {
		toSerialize["jobsCount"] = o.JobsCount
	}
	if !IsNil(o.CompletedJobsCount) {
		toSerialize["completedJobsCount"] = o.CompletedJobsCount
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableWorkflowItem struct {
	value *WorkflowItem
	isSet bool
}

func (v NullableWorkflowItem) Get() *WorkflowItem {
	return v.value
}

func (v *NullableWorkflowItem) Set(val *WorkflowItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowItem(val *WorkflowItem) *NullableWorkflowItem {
	return &NullableWorkflowItem{value: val, isSet: true}
}

func (v NullableWorkflowItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
