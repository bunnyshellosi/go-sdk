# KubernetesIntegrationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Kubernetes integration identifier. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Kubernetes integration cluster name. | [optional] [readonly] 
**CloudName** | Pointer to **string** | Kubernetes integration cloud name. | [optional] [readonly] 
**CloudProvider** | Pointer to **string** | Kubernetes integration cluster provider. | [optional] [readonly] 
**Status** | Pointer to **string** | Kubernetes integration status. | [optional] [readonly] 
**Organization** | Pointer to **NullableString** | Organization identifier. | [optional] [readonly] 

## Methods

### NewKubernetesIntegrationItem

`func NewKubernetesIntegrationItem() *KubernetesIntegrationItem`

NewKubernetesIntegrationItem instantiates a new KubernetesIntegrationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesIntegrationItemWithDefaults

`func NewKubernetesIntegrationItemWithDefaults() *KubernetesIntegrationItem`

NewKubernetesIntegrationItemWithDefaults instantiates a new KubernetesIntegrationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesIntegrationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesIntegrationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesIntegrationItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesIntegrationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterName

`func (o *KubernetesIntegrationItem) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *KubernetesIntegrationItem) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *KubernetesIntegrationItem) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *KubernetesIntegrationItem) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCloudName

`func (o *KubernetesIntegrationItem) GetCloudName() string`

GetCloudName returns the CloudName field if non-nil, zero value otherwise.

### GetCloudNameOk

`func (o *KubernetesIntegrationItem) GetCloudNameOk() (*string, bool)`

GetCloudNameOk returns a tuple with the CloudName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudName

`func (o *KubernetesIntegrationItem) SetCloudName(v string)`

SetCloudName sets CloudName field to given value.

### HasCloudName

`func (o *KubernetesIntegrationItem) HasCloudName() bool`

HasCloudName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *KubernetesIntegrationItem) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KubernetesIntegrationItem) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KubernetesIntegrationItem) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KubernetesIntegrationItem) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesIntegrationItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesIntegrationItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesIntegrationItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesIntegrationItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesIntegrationItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesIntegrationItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesIntegrationItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesIntegrationItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *KubernetesIntegrationItem) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *KubernetesIntegrationItem) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


