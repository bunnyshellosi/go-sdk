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

// checks if the RegistryIntegrationCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryIntegrationCollection{}

// RegistryIntegrationCollection A Registry integration stores connection information for a Container registry.
type RegistryIntegrationCollection struct {
	// Registry integration identifier.
	Id *string `json:"id,omitempty"`
	// Registry integration name.
	Name *string `json:"name,omitempty"`
	// Registry integration provider.
	ProviderName *string `json:"providerName,omitempty"`
	// Registry integration status.
	Status *string `json:"status,omitempty"`
	// Organization identifier.
	Organization *string `json:"organization,omitempty"`
}

// NewRegistryIntegrationCollection instantiates a new RegistryIntegrationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryIntegrationCollection() *RegistryIntegrationCollection {
	this := RegistryIntegrationCollection{}
	return &this
}

// NewRegistryIntegrationCollectionWithDefaults instantiates a new RegistryIntegrationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryIntegrationCollectionWithDefaults() *RegistryIntegrationCollection {
	this := RegistryIntegrationCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistryIntegrationCollection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryIntegrationCollection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistryIntegrationCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistryIntegrationCollection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistryIntegrationCollection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryIntegrationCollection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistryIntegrationCollection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistryIntegrationCollection) SetName(v string) {
	o.Name = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *RegistryIntegrationCollection) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryIntegrationCollection) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *RegistryIntegrationCollection) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *RegistryIntegrationCollection) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegistryIntegrationCollection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryIntegrationCollection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegistryIntegrationCollection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RegistryIntegrationCollection) SetStatus(v string) {
	o.Status = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RegistryIntegrationCollection) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryIntegrationCollection) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RegistryIntegrationCollection) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *RegistryIntegrationCollection) SetOrganization(v string) {
	o.Organization = &v
}

func (o RegistryIntegrationCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryIntegrationCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableRegistryIntegrationCollection struct {
	value *RegistryIntegrationCollection
	isSet bool
}

func (v NullableRegistryIntegrationCollection) Get() *RegistryIntegrationCollection {
	return v.value
}

func (v *NullableRegistryIntegrationCollection) Set(val *RegistryIntegrationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryIntegrationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryIntegrationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryIntegrationCollection(val *RegistryIntegrationCollection) *NullableRegistryIntegrationCollection {
	return &NullableRegistryIntegrationCollection{value: val, isSet: true}
}

func (v NullableRegistryIntegrationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryIntegrationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
