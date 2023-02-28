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

// PaginatedTemplatesRepositoryCollection Paginated result
type PaginatedTemplatesRepositoryCollection struct {
	Links        *PaginatedLinks                        `json:"_links,omitempty"`
	TotalItems   *int32                                 `json:"totalItems,omitempty"`
	Page         *int32                                 `json:"page,omitempty"`
	ItemsPerPage *int32                                 `json:"itemsPerPage,omitempty"`
	Embedded     *EmbeddedTemplatesRepositoryCollection `json:"_embedded,omitempty"`
}

// NewPaginatedTemplatesRepositoryCollection instantiates a new PaginatedTemplatesRepositoryCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTemplatesRepositoryCollection() *PaginatedTemplatesRepositoryCollection {
	this := PaginatedTemplatesRepositoryCollection{}
	return &this
}

// NewPaginatedTemplatesRepositoryCollectionWithDefaults instantiates a new PaginatedTemplatesRepositoryCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTemplatesRepositoryCollectionWithDefaults() *PaginatedTemplatesRepositoryCollection {
	this := PaginatedTemplatesRepositoryCollection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedTemplatesRepositoryCollection) GetLinks() PaginatedLinks {
	if o == nil || o.Links == nil {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTemplatesRepositoryCollection) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedTemplatesRepositoryCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *PaginatedTemplatesRepositoryCollection) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedTemplatesRepositoryCollection) GetTotalItems() int32 {
	if o == nil || o.TotalItems == nil {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTemplatesRepositoryCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedTemplatesRepositoryCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedTemplatesRepositoryCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedTemplatesRepositoryCollection) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTemplatesRepositoryCollection) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedTemplatesRepositoryCollection) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedTemplatesRepositoryCollection) SetPage(v int32) {
	o.Page = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *PaginatedTemplatesRepositoryCollection) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTemplatesRepositoryCollection) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *PaginatedTemplatesRepositoryCollection) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *PaginatedTemplatesRepositoryCollection) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PaginatedTemplatesRepositoryCollection) GetEmbedded() EmbeddedTemplatesRepositoryCollection {
	if o == nil || o.Embedded == nil {
		var ret EmbeddedTemplatesRepositoryCollection
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTemplatesRepositoryCollection) GetEmbeddedOk() (*EmbeddedTemplatesRepositoryCollection, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PaginatedTemplatesRepositoryCollection) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedTemplatesRepositoryCollection and assigns it to the Embedded field.
func (o *PaginatedTemplatesRepositoryCollection) SetEmbedded(v EmbeddedTemplatesRepositoryCollection) {
	o.Embedded = &v
}

func (o PaginatedTemplatesRepositoryCollection) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedTemplatesRepositoryCollection struct {
	value *PaginatedTemplatesRepositoryCollection
	isSet bool
}

func (v NullablePaginatedTemplatesRepositoryCollection) Get() *PaginatedTemplatesRepositoryCollection {
	return v.value
}

func (v *NullablePaginatedTemplatesRepositoryCollection) Set(val *PaginatedTemplatesRepositoryCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTemplatesRepositoryCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTemplatesRepositoryCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTemplatesRepositoryCollection(val *PaginatedTemplatesRepositoryCollection) *NullablePaginatedTemplatesRepositoryCollection {
	return &NullablePaginatedTemplatesRepositoryCollection{value: val, isSet: true}
}

func (v NullablePaginatedTemplatesRepositoryCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTemplatesRepositoryCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
