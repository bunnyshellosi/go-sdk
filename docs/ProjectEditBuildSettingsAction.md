# ProjectEditBuildSettingsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseManagedRegistry** | Pointer to **NullableBool** |  | [optional] 
**RegistryIntegration** | **NullableString** |  | 
**UseManagedCluster** | Pointer to **NullableBool** |  | [optional] 
**KubernetesIntegration** | **NullableString** |  | 
**Cpu** | **NullableString** | K8s supports decimal values with step 0.001 That&#39;s why we are using decimal. | 
**Memory** | **NullableInt32** | expressed in Mi | 
**TimeoutSeconds** | **NullableInt32** |  | 

## Methods

### NewProjectEditBuildSettingsAction

`func NewProjectEditBuildSettingsAction(registryIntegration NullableString, kubernetesIntegration NullableString, cpu NullableString, memory NullableInt32, timeoutSeconds NullableInt32, ) *ProjectEditBuildSettingsAction`

NewProjectEditBuildSettingsAction instantiates a new ProjectEditBuildSettingsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEditBuildSettingsActionWithDefaults

`func NewProjectEditBuildSettingsActionWithDefaults() *ProjectEditBuildSettingsAction`

NewProjectEditBuildSettingsActionWithDefaults instantiates a new ProjectEditBuildSettingsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseManagedRegistry

`func (o *ProjectEditBuildSettingsAction) GetUseManagedRegistry() bool`

GetUseManagedRegistry returns the UseManagedRegistry field if non-nil, zero value otherwise.

### GetUseManagedRegistryOk

`func (o *ProjectEditBuildSettingsAction) GetUseManagedRegistryOk() (*bool, bool)`

GetUseManagedRegistryOk returns a tuple with the UseManagedRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedRegistry

`func (o *ProjectEditBuildSettingsAction) SetUseManagedRegistry(v bool)`

SetUseManagedRegistry sets UseManagedRegistry field to given value.

### HasUseManagedRegistry

`func (o *ProjectEditBuildSettingsAction) HasUseManagedRegistry() bool`

HasUseManagedRegistry returns a boolean if a field has been set.

### SetUseManagedRegistryNil

`func (o *ProjectEditBuildSettingsAction) SetUseManagedRegistryNil(b bool)`

 SetUseManagedRegistryNil sets the value for UseManagedRegistry to be an explicit nil

### UnsetUseManagedRegistry
`func (o *ProjectEditBuildSettingsAction) UnsetUseManagedRegistry()`

UnsetUseManagedRegistry ensures that no value is present for UseManagedRegistry, not even an explicit nil
### GetRegistryIntegration

`func (o *ProjectEditBuildSettingsAction) GetRegistryIntegration() string`

GetRegistryIntegration returns the RegistryIntegration field if non-nil, zero value otherwise.

### GetRegistryIntegrationOk

`func (o *ProjectEditBuildSettingsAction) GetRegistryIntegrationOk() (*string, bool)`

GetRegistryIntegrationOk returns a tuple with the RegistryIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegration

`func (o *ProjectEditBuildSettingsAction) SetRegistryIntegration(v string)`

SetRegistryIntegration sets RegistryIntegration field to given value.


### SetRegistryIntegrationNil

`func (o *ProjectEditBuildSettingsAction) SetRegistryIntegrationNil(b bool)`

 SetRegistryIntegrationNil sets the value for RegistryIntegration to be an explicit nil

### UnsetRegistryIntegration
`func (o *ProjectEditBuildSettingsAction) UnsetRegistryIntegration()`

UnsetRegistryIntegration ensures that no value is present for RegistryIntegration, not even an explicit nil
### GetUseManagedCluster

`func (o *ProjectEditBuildSettingsAction) GetUseManagedCluster() bool`

GetUseManagedCluster returns the UseManagedCluster field if non-nil, zero value otherwise.

### GetUseManagedClusterOk

`func (o *ProjectEditBuildSettingsAction) GetUseManagedClusterOk() (*bool, bool)`

GetUseManagedClusterOk returns a tuple with the UseManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedCluster

`func (o *ProjectEditBuildSettingsAction) SetUseManagedCluster(v bool)`

SetUseManagedCluster sets UseManagedCluster field to given value.

### HasUseManagedCluster

`func (o *ProjectEditBuildSettingsAction) HasUseManagedCluster() bool`

HasUseManagedCluster returns a boolean if a field has been set.

### SetUseManagedClusterNil

`func (o *ProjectEditBuildSettingsAction) SetUseManagedClusterNil(b bool)`

 SetUseManagedClusterNil sets the value for UseManagedCluster to be an explicit nil

### UnsetUseManagedCluster
`func (o *ProjectEditBuildSettingsAction) UnsetUseManagedCluster()`

UnsetUseManagedCluster ensures that no value is present for UseManagedCluster, not even an explicit nil
### GetKubernetesIntegration

`func (o *ProjectEditBuildSettingsAction) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *ProjectEditBuildSettingsAction) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *ProjectEditBuildSettingsAction) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.


### SetKubernetesIntegrationNil

`func (o *ProjectEditBuildSettingsAction) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *ProjectEditBuildSettingsAction) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetCpu

`func (o *ProjectEditBuildSettingsAction) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectEditBuildSettingsAction) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectEditBuildSettingsAction) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### SetCpuNil

`func (o *ProjectEditBuildSettingsAction) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectEditBuildSettingsAction) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectEditBuildSettingsAction) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectEditBuildSettingsAction) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectEditBuildSettingsAction) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### SetMemoryNil

`func (o *ProjectEditBuildSettingsAction) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectEditBuildSettingsAction) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetTimeoutSeconds

`func (o *ProjectEditBuildSettingsAction) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ProjectEditBuildSettingsAction) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ProjectEditBuildSettingsAction) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### SetTimeoutSecondsNil

`func (o *ProjectEditBuildSettingsAction) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *ProjectEditBuildSettingsAction) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


