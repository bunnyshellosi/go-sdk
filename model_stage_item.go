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

// checks if the StageItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageItem{}

// StageItem struct for StageItem
type StageItem struct {
	// Stage identifier.
	Id *string `json:"id,omitempty"`
	// Stage name.
	Name *string `json:"name,omitempty"`
	// Stage status.
	Status *string `json:"status,omitempty"`
	// Stage duration.
	Duration NullableInt32 `json:"duration,omitempty"`
	// Stage total jobs.
	JobsCount *int32 `json:"jobsCount,omitempty"`
	// Stage completed jobs.
	CompletedJobsCount *int32 `json:"completedJobsCount,omitempty"`
	// Event identifier.
	Pipeline *string `json:"pipeline,omitempty"`
}

// NewStageItem instantiates a new StageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageItem() *StageItem {
	this := StageItem{}
	return &this
}

// NewStageItemWithDefaults instantiates a new StageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageItemWithDefaults() *StageItem {
	this := StageItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StageItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StageItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StageItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StageItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StageItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StageItem) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StageItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StageItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StageItem) SetStatus(v string) {
	o.Status = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StageItem) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StageItem) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *StageItem) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *StageItem) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil
func (o *StageItem) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *StageItem) UnsetDuration() {
	o.Duration.Unset()
}

// GetJobsCount returns the JobsCount field value if set, zero value otherwise.
func (o *StageItem) GetJobsCount() int32 {
	if o == nil || IsNil(o.JobsCount) {
		var ret int32
		return ret
	}
	return *o.JobsCount
}

// GetJobsCountOk returns a tuple with the JobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetJobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.JobsCount) {
		return nil, false
	}
	return o.JobsCount, true
}

// HasJobsCount returns a boolean if a field has been set.
func (o *StageItem) HasJobsCount() bool {
	if o != nil && !IsNil(o.JobsCount) {
		return true
	}

	return false
}

// SetJobsCount gets a reference to the given int32 and assigns it to the JobsCount field.
func (o *StageItem) SetJobsCount(v int32) {
	o.JobsCount = &v
}

// GetCompletedJobsCount returns the CompletedJobsCount field value if set, zero value otherwise.
func (o *StageItem) GetCompletedJobsCount() int32 {
	if o == nil || IsNil(o.CompletedJobsCount) {
		var ret int32
		return ret
	}
	return *o.CompletedJobsCount
}

// GetCompletedJobsCountOk returns a tuple with the CompletedJobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetCompletedJobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CompletedJobsCount) {
		return nil, false
	}
	return o.CompletedJobsCount, true
}

// HasCompletedJobsCount returns a boolean if a field has been set.
func (o *StageItem) HasCompletedJobsCount() bool {
	if o != nil && !IsNil(o.CompletedJobsCount) {
		return true
	}

	return false
}

// SetCompletedJobsCount gets a reference to the given int32 and assigns it to the CompletedJobsCount field.
func (o *StageItem) SetCompletedJobsCount(v int32) {
	o.CompletedJobsCount = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *StageItem) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageItem) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *StageItem) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *StageItem) SetPipeline(v string) {
	o.Pipeline = &v
}

func (o StageItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	return toSerialize, nil
}

type NullableStageItem struct {
	value *StageItem
	isSet bool
}

func (v NullableStageItem) Get() *StageItem {
	return v.value
}

func (v *NullableStageItem) Set(val *StageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageItem(val *StageItem) *NullableStageItem {
	return &NullableStageItem{value: val, isSet: true}
}

func (v NullableStageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
