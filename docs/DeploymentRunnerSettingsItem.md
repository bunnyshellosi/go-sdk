# DeploymentRunnerSettingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseManagedCluster** | Pointer to **bool** | Use the managed runner Cluster. | [optional] 
**KubernetesIntegration** | Pointer to **NullableString** | The Kubernetes integration cluster used for the runners. | [optional] 
**Cpu** | Pointer to **NullableString** | The CPU allocated for the runner. | [optional] 
**Memory** | Pointer to **NullableInt32** | The memory allocated for the runner. | [optional] 
**Timeout** | Pointer to **NullableInt32** | The max execution time for the runner. | [optional] 
**LastStatus** | Pointer to **NullableString** | The latest status of the deploymentRunner settings. | [optional] 
**LastError** | Pointer to **NullableString** | The latest status of the deploymentRunner settings. | [optional] 

## Methods

### NewDeploymentRunnerSettingsItem

`func NewDeploymentRunnerSettingsItem() *DeploymentRunnerSettingsItem`

NewDeploymentRunnerSettingsItem instantiates a new DeploymentRunnerSettingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRunnerSettingsItemWithDefaults

`func NewDeploymentRunnerSettingsItemWithDefaults() *DeploymentRunnerSettingsItem`

NewDeploymentRunnerSettingsItemWithDefaults instantiates a new DeploymentRunnerSettingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseManagedCluster

`func (o *DeploymentRunnerSettingsItem) GetUseManagedCluster() bool`

GetUseManagedCluster returns the UseManagedCluster field if non-nil, zero value otherwise.

### GetUseManagedClusterOk

`func (o *DeploymentRunnerSettingsItem) GetUseManagedClusterOk() (*bool, bool)`

GetUseManagedClusterOk returns a tuple with the UseManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedCluster

`func (o *DeploymentRunnerSettingsItem) SetUseManagedCluster(v bool)`

SetUseManagedCluster sets UseManagedCluster field to given value.

### HasUseManagedCluster

`func (o *DeploymentRunnerSettingsItem) HasUseManagedCluster() bool`

HasUseManagedCluster returns a boolean if a field has been set.

### GetKubernetesIntegration

`func (o *DeploymentRunnerSettingsItem) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *DeploymentRunnerSettingsItem) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *DeploymentRunnerSettingsItem) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *DeploymentRunnerSettingsItem) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *DeploymentRunnerSettingsItem) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *DeploymentRunnerSettingsItem) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetCpu

`func (o *DeploymentRunnerSettingsItem) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DeploymentRunnerSettingsItem) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DeploymentRunnerSettingsItem) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DeploymentRunnerSettingsItem) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *DeploymentRunnerSettingsItem) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *DeploymentRunnerSettingsItem) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *DeploymentRunnerSettingsItem) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DeploymentRunnerSettingsItem) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DeploymentRunnerSettingsItem) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DeploymentRunnerSettingsItem) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *DeploymentRunnerSettingsItem) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *DeploymentRunnerSettingsItem) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetTimeout

`func (o *DeploymentRunnerSettingsItem) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DeploymentRunnerSettingsItem) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DeploymentRunnerSettingsItem) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DeploymentRunnerSettingsItem) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DeploymentRunnerSettingsItem) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DeploymentRunnerSettingsItem) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetLastStatus

`func (o *DeploymentRunnerSettingsItem) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *DeploymentRunnerSettingsItem) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *DeploymentRunnerSettingsItem) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *DeploymentRunnerSettingsItem) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *DeploymentRunnerSettingsItem) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *DeploymentRunnerSettingsItem) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastError

`func (o *DeploymentRunnerSettingsItem) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *DeploymentRunnerSettingsItem) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *DeploymentRunnerSettingsItem) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *DeploymentRunnerSettingsItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *DeploymentRunnerSettingsItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *DeploymentRunnerSettingsItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


