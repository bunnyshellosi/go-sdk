/*
API Bunnyshell Environments

Interact with Bunnyshell Platform

API version: 1.0.1
Contact: api+support@bunnyshell.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// PaginatedLinks struct for PaginatedLinks
type PaginatedLinks struct {
	Self *PaginatedLinkSelf  `json:"self,omitempty"`
	Item []PaginatedLinkSelf `json:"item,omitempty"`
}

// NewPaginatedLinks instantiates a new PaginatedLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLinks() *PaginatedLinks {
	this := PaginatedLinks{}
	return &this
}

// NewPaginatedLinksWithDefaults instantiates a new PaginatedLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLinksWithDefaults() *PaginatedLinks {
	this := PaginatedLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PaginatedLinks) GetSelf() PaginatedLinkSelf {
	if o == nil || o.Self == nil {
		var ret PaginatedLinkSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetSelfOk() (*PaginatedLinkSelf, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PaginatedLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given PaginatedLinkSelf and assigns it to the Self field.
func (o *PaginatedLinks) SetSelf(v PaginatedLinkSelf) {
	o.Self = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *PaginatedLinks) GetItem() []PaginatedLinkSelf {
	if o == nil || o.Item == nil {
		var ret []PaginatedLinkSelf
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetItemOk() ([]PaginatedLinkSelf, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *PaginatedLinks) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given []PaginatedLinkSelf and assigns it to the Item field.
func (o *PaginatedLinks) SetItem(v []PaginatedLinkSelf) {
	o.Item = v
}

func (o PaginatedLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLinks struct {
	value *PaginatedLinks
	isSet bool
}

func (v NullablePaginatedLinks) Get() *PaginatedLinks {
	return v.value
}

func (v *NullablePaginatedLinks) Set(val *PaginatedLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLinks(val *PaginatedLinks) *NullablePaginatedLinks {
	return &NullablePaginatedLinks{value: val, isSet: true}
}

func (v NullablePaginatedLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
