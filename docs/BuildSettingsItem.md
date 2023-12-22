# BuildSettingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseManagedRegistry** | Pointer to **bool** | Use the managed container registry. | [optional] 
**RegistryIntegration** | Pointer to **NullableString** | The Container Registry used for the built images. | [optional] 
**UseManagedCluster** | Pointer to **bool** | Use the managed builds Cluster. | [optional] 
**KubernetesIntegration** | Pointer to **NullableString** | The Kubernetes integration cluster used for the image builds. | [optional] 
**Cpu** | Pointer to **NullableString** | The CPU allocated for the build runner. | [optional] 
**Memory** | Pointer to **NullableInt32** | The memory allocated for the build runner. | [optional] 
**LastStatus** | Pointer to **NullableString** | The latest status of the build settings. | [optional] 
**LastError** | Pointer to **NullableString** | The latest status of the build settings. | [optional] 

## Methods

### NewBuildSettingsItem

`func NewBuildSettingsItem() *BuildSettingsItem`

NewBuildSettingsItem instantiates a new BuildSettingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSettingsItemWithDefaults

`func NewBuildSettingsItemWithDefaults() *BuildSettingsItem`

NewBuildSettingsItemWithDefaults instantiates a new BuildSettingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseManagedRegistry

`func (o *BuildSettingsItem) GetUseManagedRegistry() bool`

GetUseManagedRegistry returns the UseManagedRegistry field if non-nil, zero value otherwise.

### GetUseManagedRegistryOk

`func (o *BuildSettingsItem) GetUseManagedRegistryOk() (*bool, bool)`

GetUseManagedRegistryOk returns a tuple with the UseManagedRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedRegistry

`func (o *BuildSettingsItem) SetUseManagedRegistry(v bool)`

SetUseManagedRegistry sets UseManagedRegistry field to given value.

### HasUseManagedRegistry

`func (o *BuildSettingsItem) HasUseManagedRegistry() bool`

HasUseManagedRegistry returns a boolean if a field has been set.

### GetRegistryIntegration

`func (o *BuildSettingsItem) GetRegistryIntegration() string`

GetRegistryIntegration returns the RegistryIntegration field if non-nil, zero value otherwise.

### GetRegistryIntegrationOk

`func (o *BuildSettingsItem) GetRegistryIntegrationOk() (*string, bool)`

GetRegistryIntegrationOk returns a tuple with the RegistryIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegration

`func (o *BuildSettingsItem) SetRegistryIntegration(v string)`

SetRegistryIntegration sets RegistryIntegration field to given value.

### HasRegistryIntegration

`func (o *BuildSettingsItem) HasRegistryIntegration() bool`

HasRegistryIntegration returns a boolean if a field has been set.

### SetRegistryIntegrationNil

`func (o *BuildSettingsItem) SetRegistryIntegrationNil(b bool)`

 SetRegistryIntegrationNil sets the value for RegistryIntegration to be an explicit nil

### UnsetRegistryIntegration
`func (o *BuildSettingsItem) UnsetRegistryIntegration()`

UnsetRegistryIntegration ensures that no value is present for RegistryIntegration, not even an explicit nil
### GetUseManagedCluster

`func (o *BuildSettingsItem) GetUseManagedCluster() bool`

GetUseManagedCluster returns the UseManagedCluster field if non-nil, zero value otherwise.

### GetUseManagedClusterOk

`func (o *BuildSettingsItem) GetUseManagedClusterOk() (*bool, bool)`

GetUseManagedClusterOk returns a tuple with the UseManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedCluster

`func (o *BuildSettingsItem) SetUseManagedCluster(v bool)`

SetUseManagedCluster sets UseManagedCluster field to given value.

### HasUseManagedCluster

`func (o *BuildSettingsItem) HasUseManagedCluster() bool`

HasUseManagedCluster returns a boolean if a field has been set.

### GetKubernetesIntegration

`func (o *BuildSettingsItem) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *BuildSettingsItem) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *BuildSettingsItem) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *BuildSettingsItem) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *BuildSettingsItem) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *BuildSettingsItem) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetCpu

`func (o *BuildSettingsItem) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BuildSettingsItem) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BuildSettingsItem) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *BuildSettingsItem) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *BuildSettingsItem) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *BuildSettingsItem) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *BuildSettingsItem) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *BuildSettingsItem) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *BuildSettingsItem) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *BuildSettingsItem) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *BuildSettingsItem) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *BuildSettingsItem) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetLastStatus

`func (o *BuildSettingsItem) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *BuildSettingsItem) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *BuildSettingsItem) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *BuildSettingsItem) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *BuildSettingsItem) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *BuildSettingsItem) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastError

`func (o *BuildSettingsItem) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *BuildSettingsItem) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *BuildSettingsItem) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *BuildSettingsItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *BuildSettingsItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *BuildSettingsItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


