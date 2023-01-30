# EnvironmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | Environment type: primary or ephemeral. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment name. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Environment k8s namespace. | [optional] [readonly] 
**TotalComponents** | Pointer to **int32** | Service component identifier | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Environment operation status. | [optional] [readonly] 
**ClusterStatus** | Pointer to **string** | Environment cluster status. | [optional] [readonly] 
**Project** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**KubernetesIntegration** | Pointer to **NullableString** | Kubernetes integration identifier. | [optional] [readonly] 

## Methods

### NewEnvironmentItem

`func NewEnvironmentItem() *EnvironmentItem`

NewEnvironmentItem instantiates a new EnvironmentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentItemWithDefaults

`func NewEnvironmentItemWithDefaults() *EnvironmentItem`

NewEnvironmentItemWithDefaults instantiates a new EnvironmentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EnvironmentItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EnvironmentItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EnvironmentItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EnvironmentItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTotalComponents

`func (o *EnvironmentItem) GetTotalComponents() int32`

GetTotalComponents returns the TotalComponents field if non-nil, zero value otherwise.

### GetTotalComponentsOk

`func (o *EnvironmentItem) GetTotalComponentsOk() (*int32, bool)`

GetTotalComponentsOk returns a tuple with the TotalComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalComponents

`func (o *EnvironmentItem) SetTotalComponents(v int32)`

SetTotalComponents sets TotalComponents field to given value.

### HasTotalComponents

`func (o *EnvironmentItem) HasTotalComponents() bool`

HasTotalComponents returns a boolean if a field has been set.

### GetOperationStatus

`func (o *EnvironmentItem) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *EnvironmentItem) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *EnvironmentItem) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *EnvironmentItem) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetClusterStatus

`func (o *EnvironmentItem) GetClusterStatus() string`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *EnvironmentItem) GetClusterStatusOk() (*string, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *EnvironmentItem) SetClusterStatus(v string)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *EnvironmentItem) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetProject

`func (o *EnvironmentItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentItem) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetKubernetesIntegration

`func (o *EnvironmentItem) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentItem) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentItem) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *EnvironmentItem) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *EnvironmentItem) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentItem) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


