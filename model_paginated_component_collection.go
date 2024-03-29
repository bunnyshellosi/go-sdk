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

// checks if the PaginatedComponentCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedComponentCollection{}

// PaginatedComponentCollection Paginated result
type PaginatedComponentCollection struct {
	Links        *PaginatedLinks              `json:"_links,omitempty"`
	TotalItems   *int32                       `json:"totalItems,omitempty"`
	Page         *int32                       `json:"page,omitempty"`
	ItemsPerPage *int32                       `json:"itemsPerPage,omitempty"`
	Embedded     *EmbeddedComponentCollection `json:"_embedded,omitempty"`
}

// NewPaginatedComponentCollection instantiates a new PaginatedComponentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedComponentCollection() *PaginatedComponentCollection {
	this := PaginatedComponentCollection{}
	return &this
}

// NewPaginatedComponentCollectionWithDefaults instantiates a new PaginatedComponentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedComponentCollectionWithDefaults() *PaginatedComponentCollection {
	this := PaginatedComponentCollection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedComponentCollection) GetLinks() PaginatedLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentCollection) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedComponentCollection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *PaginatedComponentCollection) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedComponentCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedComponentCollection) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedComponentCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedComponentCollection) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentCollection) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedComponentCollection) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedComponentCollection) SetPage(v int32) {
	o.Page = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *PaginatedComponentCollection) GetItemsPerPage() int32 {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentCollection) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *PaginatedComponentCollection) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *PaginatedComponentCollection) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PaginatedComponentCollection) GetEmbedded() EmbeddedComponentCollection {
	if o == nil || IsNil(o.Embedded) {
		var ret EmbeddedComponentCollection
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedComponentCollection) GetEmbeddedOk() (*EmbeddedComponentCollection, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PaginatedComponentCollection) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedComponentCollection and assigns it to the Embedded field.
func (o *PaginatedComponentCollection) SetEmbedded(v EmbeddedComponentCollection) {
	o.Embedded = &v
}

func (o PaginatedComponentCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedComponentCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.ItemsPerPage) {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	return toSerialize, nil
}

type NullablePaginatedComponentCollection struct {
	value *PaginatedComponentCollection
	isSet bool
}

func (v NullablePaginatedComponentCollection) Get() *PaginatedComponentCollection {
	return v.value
}

func (v *NullablePaginatedComponentCollection) Set(val *PaginatedComponentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedComponentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedComponentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedComponentCollection(val *PaginatedComponentCollection) *NullablePaginatedComponentCollection {
	return &NullablePaginatedComponentCollection{value: val, isSet: true}
}

func (v NullablePaginatedComponentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedComponentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
