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

// ClusterKubeConfigRead struct for ClusterKubeConfigRead
type ClusterKubeConfigRead struct {
	CertificateAuthorityData string `json:"certificate-authority-data"`
	Server                   string `json:"server"`
}

// NewClusterKubeConfigRead instantiates a new ClusterKubeConfigRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterKubeConfigRead(certificateAuthorityData string, server string) *ClusterKubeConfigRead {
	this := ClusterKubeConfigRead{}
	this.CertificateAuthorityData = certificateAuthorityData
	this.Server = server
	return &this
}

// NewClusterKubeConfigReadWithDefaults instantiates a new ClusterKubeConfigRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterKubeConfigReadWithDefaults() *ClusterKubeConfigRead {
	this := ClusterKubeConfigRead{}
	return &this
}

// GetCertificateAuthorityData returns the CertificateAuthorityData field value
func (o *ClusterKubeConfigRead) GetCertificateAuthorityData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateAuthorityData
}

// GetCertificateAuthorityDataOk returns a tuple with the CertificateAuthorityData field value
// and a boolean to check if the value has been set.
func (o *ClusterKubeConfigRead) GetCertificateAuthorityDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateAuthorityData, true
}

// SetCertificateAuthorityData sets field value
func (o *ClusterKubeConfigRead) SetCertificateAuthorityData(v string) {
	o.CertificateAuthorityData = v
}

// GetServer returns the Server field value
func (o *ClusterKubeConfigRead) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ClusterKubeConfigRead) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ClusterKubeConfigRead) SetServer(v string) {
	o.Server = v
}

func (o ClusterKubeConfigRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificate-authority-data"] = o.CertificateAuthorityData
	}
	if true {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableClusterKubeConfigRead struct {
	value *ClusterKubeConfigRead
	isSet bool
}

func (v NullableClusterKubeConfigRead) Get() *ClusterKubeConfigRead {
	return v.value
}

func (v *NullableClusterKubeConfigRead) Set(val *ClusterKubeConfigRead) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterKubeConfigRead) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterKubeConfigRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterKubeConfigRead(val *ClusterKubeConfigRead) *NullableClusterKubeConfigRead {
	return &NullableClusterKubeConfigRead{value: val, isSet: true}
}

func (v NullableClusterKubeConfigRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterKubeConfigRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}