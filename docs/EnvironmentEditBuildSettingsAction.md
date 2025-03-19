# EnvironmentEditBuildSettingsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseManagedRegistry** | Pointer to **NullableBool** |  | [optional] 
**RegistryIntegration** | **NullableString** |  | 
**UseManagedCluster** | Pointer to **NullableBool** |  | [optional] 
**KubernetesIntegration** | **NullableString** |  | 
**BuildEngine** | Pointer to **NullableString** |  | [optional] 
**Cpu** | **NullableString** | K8s supports decimal values with step 0.001 That&#39;s why we are using decimal. | 
**Memory** | **NullableInt32** | expressed in Mi | 
**TimeoutSeconds** | **NullableInt32** |  | 
**NamespaceCustomLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnvironmentEditBuildSettingsAction

`func NewEnvironmentEditBuildSettingsAction(registryIntegration NullableString, kubernetesIntegration NullableString, cpu NullableString, memory NullableInt32, timeoutSeconds NullableInt32, ) *EnvironmentEditBuildSettingsAction`

NewEnvironmentEditBuildSettingsAction instantiates a new EnvironmentEditBuildSettingsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditBuildSettingsActionWithDefaults

`func NewEnvironmentEditBuildSettingsActionWithDefaults() *EnvironmentEditBuildSettingsAction`

NewEnvironmentEditBuildSettingsActionWithDefaults instantiates a new EnvironmentEditBuildSettingsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseManagedRegistry

`func (o *EnvironmentEditBuildSettingsAction) GetUseManagedRegistry() bool`

GetUseManagedRegistry returns the UseManagedRegistry field if non-nil, zero value otherwise.

### GetUseManagedRegistryOk

`func (o *EnvironmentEditBuildSettingsAction) GetUseManagedRegistryOk() (*bool, bool)`

GetUseManagedRegistryOk returns a tuple with the UseManagedRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedRegistry

`func (o *EnvironmentEditBuildSettingsAction) SetUseManagedRegistry(v bool)`

SetUseManagedRegistry sets UseManagedRegistry field to given value.

### HasUseManagedRegistry

`func (o *EnvironmentEditBuildSettingsAction) HasUseManagedRegistry() bool`

HasUseManagedRegistry returns a boolean if a field has been set.

### SetUseManagedRegistryNil

`func (o *EnvironmentEditBuildSettingsAction) SetUseManagedRegistryNil(b bool)`

 SetUseManagedRegistryNil sets the value for UseManagedRegistry to be an explicit nil

### UnsetUseManagedRegistry
`func (o *EnvironmentEditBuildSettingsAction) UnsetUseManagedRegistry()`

UnsetUseManagedRegistry ensures that no value is present for UseManagedRegistry, not even an explicit nil
### GetRegistryIntegration

`func (o *EnvironmentEditBuildSettingsAction) GetRegistryIntegration() string`

GetRegistryIntegration returns the RegistryIntegration field if non-nil, zero value otherwise.

### GetRegistryIntegrationOk

`func (o *EnvironmentEditBuildSettingsAction) GetRegistryIntegrationOk() (*string, bool)`

GetRegistryIntegrationOk returns a tuple with the RegistryIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIntegration

`func (o *EnvironmentEditBuildSettingsAction) SetRegistryIntegration(v string)`

SetRegistryIntegration sets RegistryIntegration field to given value.


### SetRegistryIntegrationNil

`func (o *EnvironmentEditBuildSettingsAction) SetRegistryIntegrationNil(b bool)`

 SetRegistryIntegrationNil sets the value for RegistryIntegration to be an explicit nil

### UnsetRegistryIntegration
`func (o *EnvironmentEditBuildSettingsAction) UnsetRegistryIntegration()`

UnsetRegistryIntegration ensures that no value is present for RegistryIntegration, not even an explicit nil
### GetUseManagedCluster

`func (o *EnvironmentEditBuildSettingsAction) GetUseManagedCluster() bool`

GetUseManagedCluster returns the UseManagedCluster field if non-nil, zero value otherwise.

### GetUseManagedClusterOk

`func (o *EnvironmentEditBuildSettingsAction) GetUseManagedClusterOk() (*bool, bool)`

GetUseManagedClusterOk returns a tuple with the UseManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedCluster

`func (o *EnvironmentEditBuildSettingsAction) SetUseManagedCluster(v bool)`

SetUseManagedCluster sets UseManagedCluster field to given value.

### HasUseManagedCluster

`func (o *EnvironmentEditBuildSettingsAction) HasUseManagedCluster() bool`

HasUseManagedCluster returns a boolean if a field has been set.

### SetUseManagedClusterNil

`func (o *EnvironmentEditBuildSettingsAction) SetUseManagedClusterNil(b bool)`

 SetUseManagedClusterNil sets the value for UseManagedCluster to be an explicit nil

### UnsetUseManagedCluster
`func (o *EnvironmentEditBuildSettingsAction) UnsetUseManagedCluster()`

UnsetUseManagedCluster ensures that no value is present for UseManagedCluster, not even an explicit nil
### GetKubernetesIntegration

`func (o *EnvironmentEditBuildSettingsAction) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentEditBuildSettingsAction) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentEditBuildSettingsAction) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.


### SetKubernetesIntegrationNil

`func (o *EnvironmentEditBuildSettingsAction) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentEditBuildSettingsAction) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetBuildEngine

`func (o *EnvironmentEditBuildSettingsAction) GetBuildEngine() string`

GetBuildEngine returns the BuildEngine field if non-nil, zero value otherwise.

### GetBuildEngineOk

`func (o *EnvironmentEditBuildSettingsAction) GetBuildEngineOk() (*string, bool)`

GetBuildEngineOk returns a tuple with the BuildEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEngine

`func (o *EnvironmentEditBuildSettingsAction) SetBuildEngine(v string)`

SetBuildEngine sets BuildEngine field to given value.

### HasBuildEngine

`func (o *EnvironmentEditBuildSettingsAction) HasBuildEngine() bool`

HasBuildEngine returns a boolean if a field has been set.

### SetBuildEngineNil

`func (o *EnvironmentEditBuildSettingsAction) SetBuildEngineNil(b bool)`

 SetBuildEngineNil sets the value for BuildEngine to be an explicit nil

### UnsetBuildEngine
`func (o *EnvironmentEditBuildSettingsAction) UnsetBuildEngine()`

UnsetBuildEngine ensures that no value is present for BuildEngine, not even an explicit nil
### GetCpu

`func (o *EnvironmentEditBuildSettingsAction) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentEditBuildSettingsAction) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentEditBuildSettingsAction) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### SetCpuNil

`func (o *EnvironmentEditBuildSettingsAction) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentEditBuildSettingsAction) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentEditBuildSettingsAction) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentEditBuildSettingsAction) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentEditBuildSettingsAction) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### SetMemoryNil

`func (o *EnvironmentEditBuildSettingsAction) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentEditBuildSettingsAction) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetTimeoutSeconds

`func (o *EnvironmentEditBuildSettingsAction) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *EnvironmentEditBuildSettingsAction) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *EnvironmentEditBuildSettingsAction) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### SetTimeoutSecondsNil

`func (o *EnvironmentEditBuildSettingsAction) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *EnvironmentEditBuildSettingsAction) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetNamespaceCustomLabels

`func (o *EnvironmentEditBuildSettingsAction) GetNamespaceCustomLabels() []string`

GetNamespaceCustomLabels returns the NamespaceCustomLabels field if non-nil, zero value otherwise.

### GetNamespaceCustomLabelsOk

`func (o *EnvironmentEditBuildSettingsAction) GetNamespaceCustomLabelsOk() (*[]string, bool)`

GetNamespaceCustomLabelsOk returns a tuple with the NamespaceCustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceCustomLabels

`func (o *EnvironmentEditBuildSettingsAction) SetNamespaceCustomLabels(v []string)`

SetNamespaceCustomLabels sets NamespaceCustomLabels field to given value.

### HasNamespaceCustomLabels

`func (o *EnvironmentEditBuildSettingsAction) HasNamespaceCustomLabels() bool`

HasNamespaceCustomLabels returns a boolean if a field has been set.

### SetNamespaceCustomLabelsNil

`func (o *EnvironmentEditBuildSettingsAction) SetNamespaceCustomLabelsNil(b bool)`

 SetNamespaceCustomLabelsNil sets the value for NamespaceCustomLabels to be an explicit nil

### UnsetNamespaceCustomLabels
`func (o *EnvironmentEditBuildSettingsAction) UnsetNamespaceCustomLabels()`

UnsetNamespaceCustomLabels ensures that no value is present for NamespaceCustomLabels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


