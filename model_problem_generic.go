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

// ProblemGeneric Generic Error
type ProblemGeneric struct {
	Title  *string `json:"title,omitempty"`
	Detail *string `json:"detail,omitempty"`
}

// NewProblemGeneric instantiates a new ProblemGeneric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemGeneric() *ProblemGeneric {
	this := ProblemGeneric{}
	return &this
}

// NewProblemGenericWithDefaults instantiates a new ProblemGeneric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemGenericWithDefaults() *ProblemGeneric {
	this := ProblemGeneric{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProblemGeneric) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemGeneric) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemGeneric) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProblemGeneric) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemGeneric) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemGeneric) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemGeneric) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemGeneric) SetDetail(v string) {
	o.Detail = &v
}

func (o ProblemGeneric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableProblemGeneric struct {
	value *ProblemGeneric
	isSet bool
}

func (v NullableProblemGeneric) Get() *ProblemGeneric {
	return v.value
}

func (v *NullableProblemGeneric) Set(val *ProblemGeneric) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemGeneric) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemGeneric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemGeneric(val *ProblemGeneric) *NullableProblemGeneric {
	return &NullableProblemGeneric{value: val, isSet: true}
}

func (v NullableProblemGeneric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemGeneric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
