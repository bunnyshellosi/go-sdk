# ClusterKubeConfigRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthorityData** | **string** |  | 
**Server** | **string** |  | 

## Methods

### NewClusterKubeConfigRead

`func NewClusterKubeConfigRead(certificateAuthorityData string, server string, ) *ClusterKubeConfigRead`

NewClusterKubeConfigRead instantiates a new ClusterKubeConfigRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterKubeConfigReadWithDefaults

`func NewClusterKubeConfigReadWithDefaults() *ClusterKubeConfigRead`

NewClusterKubeConfigReadWithDefaults instantiates a new ClusterKubeConfigRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthorityData

`func (o *ClusterKubeConfigRead) GetCertificateAuthorityData() string`

GetCertificateAuthorityData returns the CertificateAuthorityData field if non-nil, zero value otherwise.

### GetCertificateAuthorityDataOk

`func (o *ClusterKubeConfigRead) GetCertificateAuthorityDataOk() (*string, bool)`

GetCertificateAuthorityDataOk returns a tuple with the CertificateAuthorityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityData

`func (o *ClusterKubeConfigRead) SetCertificateAuthorityData(v string)`

SetCertificateAuthorityData sets CertificateAuthorityData field to given value.


### GetServer

`func (o *ClusterKubeConfigRead) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ClusterKubeConfigRead) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ClusterKubeConfigRead) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


