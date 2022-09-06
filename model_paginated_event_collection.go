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

// PaginatedEventCollection Paginated result
type PaginatedEventCollection struct {
	Links        *PaginatedLinks          `json:"_links,omitempty"`
	TotalItems   *int32                   `json:"totalItems,omitempty"`
	Page         *int32                   `json:"page,omitempty"`
	ItemsPerPage *int32                   `json:"itemsPerPage,omitempty"`
	Embedded     *EmbeddedEventCollection `json:"_embedded,omitempty"`
}

// NewPaginatedEventCollection instantiates a new PaginatedEventCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEventCollection() *PaginatedEventCollection {
	this := PaginatedEventCollection{}
	return &this
}

// NewPaginatedEventCollectionWithDefaults instantiates a new PaginatedEventCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEventCollectionWithDefaults() *PaginatedEventCollection {
	this := PaginatedEventCollection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedEventCollection) GetLinks() PaginatedLinks {
	if o == nil || o.Links == nil {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEventCollection) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedEventCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *PaginatedEventCollection) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedEventCollection) GetTotalItems() int32 {
	if o == nil || o.TotalItems == nil {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEventCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedEventCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedEventCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedEventCollection) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEventCollection) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedEventCollection) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedEventCollection) SetPage(v int32) {
	o.Page = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *PaginatedEventCollection) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEventCollection) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *PaginatedEventCollection) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *PaginatedEventCollection) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PaginatedEventCollection) GetEmbedded() EmbeddedEventCollection {
	if o == nil || o.Embedded == nil {
		var ret EmbeddedEventCollection
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedEventCollection) GetEmbeddedOk() (*EmbeddedEventCollection, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PaginatedEventCollection) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedEventCollection and assigns it to the Embedded field.
func (o *PaginatedEventCollection) SetEmbedded(v EmbeddedEventCollection) {
	o.Embedded = &v
}

func (o PaginatedEventCollection) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedEventCollection struct {
	value *PaginatedEventCollection
	isSet bool
}

func (v NullablePaginatedEventCollection) Get() *PaginatedEventCollection {
	return v.value
}

func (v *NullablePaginatedEventCollection) Set(val *PaginatedEventCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEventCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEventCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEventCollection(val *PaginatedEventCollection) *NullablePaginatedEventCollection {
	return &NullablePaginatedEventCollection{value: val, isSet: true}
}

func (v NullablePaginatedEventCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEventCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
