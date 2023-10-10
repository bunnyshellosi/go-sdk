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

// checks if the PaginatedProjectCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedProjectCollection{}

// PaginatedProjectCollection Paginated result
type PaginatedProjectCollection struct {
	Links        *PaginatedLinks            `json:"_links,omitempty"`
	TotalItems   *int32                     `json:"totalItems,omitempty"`
	Page         *int32                     `json:"page,omitempty"`
	ItemsPerPage *int32                     `json:"itemsPerPage,omitempty"`
	Embedded     *EmbeddedProjectCollection `json:"_embedded,omitempty"`
}

// NewPaginatedProjectCollection instantiates a new PaginatedProjectCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProjectCollection() *PaginatedProjectCollection {
	this := PaginatedProjectCollection{}
	return &this
}

// NewPaginatedProjectCollectionWithDefaults instantiates a new PaginatedProjectCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProjectCollectionWithDefaults() *PaginatedProjectCollection {
	this := PaginatedProjectCollection{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedProjectCollection) GetLinks() PaginatedLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProjectCollection) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedProjectCollection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *PaginatedProjectCollection) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedProjectCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProjectCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedProjectCollection) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedProjectCollection) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedProjectCollection) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProjectCollection) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedProjectCollection) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedProjectCollection) SetPage(v int32) {
	o.Page = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *PaginatedProjectCollection) GetItemsPerPage() int32 {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProjectCollection) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *PaginatedProjectCollection) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *PaginatedProjectCollection) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PaginatedProjectCollection) GetEmbedded() EmbeddedProjectCollection {
	if o == nil || IsNil(o.Embedded) {
		var ret EmbeddedProjectCollection
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProjectCollection) GetEmbeddedOk() (*EmbeddedProjectCollection, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PaginatedProjectCollection) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedProjectCollection and assigns it to the Embedded field.
func (o *PaginatedProjectCollection) SetEmbedded(v EmbeddedProjectCollection) {
	o.Embedded = &v
}

func (o PaginatedProjectCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedProjectCollection) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedProjectCollection struct {
	value *PaginatedProjectCollection
	isSet bool
}

func (v NullablePaginatedProjectCollection) Get() *PaginatedProjectCollection {
	return v.value
}

func (v *NullablePaginatedProjectCollection) Set(val *PaginatedProjectCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProjectCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProjectCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProjectCollection(val *PaginatedProjectCollection) *NullablePaginatedProjectCollection {
	return &NullablePaginatedProjectCollection{value: val, isSet: true}
}

func (v NullablePaginatedProjectCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProjectCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
