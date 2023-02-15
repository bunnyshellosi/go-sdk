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

// PaginatedComponentGitCollection Paginated result
type PaginatedComponentGitCollection struct {
	Links        *PaginatedLinks                 `json:"_links,omitempty"`
	TotalItems   *int32                          `json:"totalItems,omitempty"`
	Page         *int32                          `json:"page,omitempty"`
	ItemsPerPage *int32                          `json:"itemsPerPage,omitempty"`
	Embedded     *EmbeddedComponentGitCollection `json:"_embedded,omitempty"`
}

// NewPaginatedComponentGitCollection instantiates a new PaginatedComponentGitCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedComponentGitCollection() *PaginatedComponentGitCollection {
	this := PaginatedComponentGitCollection{}
	return &this
}

// NewPaginatedComponentGitCollectionWithDefaults instantiates a new PaginatedComponentGitCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedComponentGitCollectionWithDefaults() *PaginatedComponentGitCollection {
	this := PaginatedComponentGitCollection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedComponentGitCollection) GetLinks() PaginatedLinks {
	if o == nil || o.Links == nil {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentGitCollection) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedComponentGitCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *PaginatedComponentGitCollection) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedComponentGitCollection) GetTotalItems() int32 {
	if o == nil || o.TotalItems == nil {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentGitCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedComponentGitCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedComponentGitCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedComponentGitCollection) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentGitCollection) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedComponentGitCollection) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedComponentGitCollection) SetPage(v int32) {
	o.Page = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *PaginatedComponentGitCollection) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentGitCollection) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *PaginatedComponentGitCollection) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *PaginatedComponentGitCollection) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PaginatedComponentGitCollection) GetEmbedded() EmbeddedComponentGitCollection {
	if o == nil || o.Embedded == nil {
		var ret EmbeddedComponentGitCollection
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentGitCollection) GetEmbeddedOk() (*EmbeddedComponentGitCollection, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PaginatedComponentGitCollection) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedComponentGitCollection and assigns it to the Embedded field.
func (o *PaginatedComponentGitCollection) SetEmbedded(v EmbeddedComponentGitCollection) {
	o.Embedded = &v
}

func (o PaginatedComponentGitCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.TotalItems != nil {
		toSerialize["totalItems"] = o.TotalItems
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.ItemsPerPage != nil {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedComponentGitCollection struct {
	value *PaginatedComponentGitCollection
	isSet bool
}

func (v NullablePaginatedComponentGitCollection) Get() *PaginatedComponentGitCollection {
	return v.value
}

func (v *NullablePaginatedComponentGitCollection) Set(val *PaginatedComponentGitCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedComponentGitCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedComponentGitCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedComponentGitCollection(val *PaginatedComponentGitCollection) *NullablePaginatedComponentGitCollection {
	return &NullablePaginatedComponentGitCollection{value: val, isSet: true}
}

func (v NullablePaginatedComponentGitCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedComponentGitCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}