# EnvironmentEditDeploymentRunnerSettingsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseManagedCluster** | Pointer to **NullableBool** |  | [optional] 
**KubernetesIntegration** | **NullableString** |  | 
**Cpu** | **NullableString** | K8s supports decimal values with step 0.001 That&#39;s why we are using decimal. | 
**Memory** | **NullableInt32** | expressed in Mi | 
**TimeoutSeconds** | **NullableInt32** |  | 
**NamespaceCustomLabels** | Pointer to **[]string** |  | [optional] 
**RealtimeLogs** | Pointer to **NullableString** |  | [optional] 
**SecurityNetworkPolicies** | Pointer to **NullableString** |  | [optional] 
**DnsRecords** | Pointer to **NullableString** |  | [optional] 
**PublicUrls** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnvironmentEditDeploymentRunnerSettingsAction

`func NewEnvironmentEditDeploymentRunnerSettingsAction(kubernetesIntegration NullableString, cpu NullableString, memory NullableInt32, timeoutSeconds NullableInt32, ) *EnvironmentEditDeploymentRunnerSettingsAction`

NewEnvironmentEditDeploymentRunnerSettingsAction instantiates a new EnvironmentEditDeploymentRunnerSettingsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditDeploymentRunnerSettingsActionWithDefaults

`func NewEnvironmentEditDeploymentRunnerSettingsActionWithDefaults() *EnvironmentEditDeploymentRunnerSettingsAction`

NewEnvironmentEditDeploymentRunnerSettingsActionWithDefaults instantiates a new EnvironmentEditDeploymentRunnerSettingsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseManagedCluster

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetUseManagedCluster() bool`

GetUseManagedCluster returns the UseManagedCluster field if non-nil, zero value otherwise.

### GetUseManagedClusterOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetUseManagedClusterOk() (*bool, bool)`

GetUseManagedClusterOk returns a tuple with the UseManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseManagedCluster

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetUseManagedCluster(v bool)`

SetUseManagedCluster sets UseManagedCluster field to given value.

### HasUseManagedCluster

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasUseManagedCluster() bool`

HasUseManagedCluster returns a boolean if a field has been set.

### SetUseManagedClusterNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetUseManagedClusterNil(b bool)`

 SetUseManagedClusterNil sets the value for UseManagedCluster to be an explicit nil

### UnsetUseManagedCluster
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetUseManagedCluster()`

UnsetUseManagedCluster ensures that no value is present for UseManagedCluster, not even an explicit nil
### GetKubernetesIntegration

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.


### SetKubernetesIntegrationNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetCpu

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### SetCpuNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### SetMemoryNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetTimeoutSeconds

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### SetTimeoutSecondsNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetNamespaceCustomLabels

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetNamespaceCustomLabels() []string`

GetNamespaceCustomLabels returns the NamespaceCustomLabels field if non-nil, zero value otherwise.

### GetNamespaceCustomLabelsOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetNamespaceCustomLabelsOk() (*[]string, bool)`

GetNamespaceCustomLabelsOk returns a tuple with the NamespaceCustomLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceCustomLabels

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetNamespaceCustomLabels(v []string)`

SetNamespaceCustomLabels sets NamespaceCustomLabels field to given value.

### HasNamespaceCustomLabels

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasNamespaceCustomLabels() bool`

HasNamespaceCustomLabels returns a boolean if a field has been set.

### SetNamespaceCustomLabelsNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetNamespaceCustomLabelsNil(b bool)`

 SetNamespaceCustomLabelsNil sets the value for NamespaceCustomLabels to be an explicit nil

### UnsetNamespaceCustomLabels
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetNamespaceCustomLabels()`

UnsetNamespaceCustomLabels ensures that no value is present for NamespaceCustomLabels, not even an explicit nil
### GetRealtimeLogs

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetRealtimeLogs() string`

GetRealtimeLogs returns the RealtimeLogs field if non-nil, zero value otherwise.

### GetRealtimeLogsOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetRealtimeLogsOk() (*string, bool)`

GetRealtimeLogsOk returns a tuple with the RealtimeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeLogs

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetRealtimeLogs(v string)`

SetRealtimeLogs sets RealtimeLogs field to given value.

### HasRealtimeLogs

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasRealtimeLogs() bool`

HasRealtimeLogs returns a boolean if a field has been set.

### SetRealtimeLogsNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetRealtimeLogsNil(b bool)`

 SetRealtimeLogsNil sets the value for RealtimeLogs to be an explicit nil

### UnsetRealtimeLogs
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetRealtimeLogs()`

UnsetRealtimeLogs ensures that no value is present for RealtimeLogs, not even an explicit nil
### GetSecurityNetworkPolicies

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetSecurityNetworkPolicies() string`

GetSecurityNetworkPolicies returns the SecurityNetworkPolicies field if non-nil, zero value otherwise.

### GetSecurityNetworkPoliciesOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetSecurityNetworkPoliciesOk() (*string, bool)`

GetSecurityNetworkPoliciesOk returns a tuple with the SecurityNetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityNetworkPolicies

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetSecurityNetworkPolicies(v string)`

SetSecurityNetworkPolicies sets SecurityNetworkPolicies field to given value.

### HasSecurityNetworkPolicies

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasSecurityNetworkPolicies() bool`

HasSecurityNetworkPolicies returns a boolean if a field has been set.

### SetSecurityNetworkPoliciesNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetSecurityNetworkPoliciesNil(b bool)`

 SetSecurityNetworkPoliciesNil sets the value for SecurityNetworkPolicies to be an explicit nil

### UnsetSecurityNetworkPolicies
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetSecurityNetworkPolicies()`

UnsetSecurityNetworkPolicies ensures that no value is present for SecurityNetworkPolicies, not even an explicit nil
### GetDnsRecords

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetDnsRecords() string`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetDnsRecordsOk() (*string, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetDnsRecords(v string)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### SetDnsRecordsNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetDnsRecordsNil(b bool)`

 SetDnsRecordsNil sets the value for DnsRecords to be an explicit nil

### UnsetDnsRecords
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetDnsRecords()`

UnsetDnsRecords ensures that no value is present for DnsRecords, not even an explicit nil
### GetPublicUrls

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetPublicUrls() string`

GetPublicUrls returns the PublicUrls field if non-nil, zero value otherwise.

### GetPublicUrlsOk

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) GetPublicUrlsOk() (*string, bool)`

GetPublicUrlsOk returns a tuple with the PublicUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrls

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetPublicUrls(v string)`

SetPublicUrls sets PublicUrls field to given value.

### HasPublicUrls

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) HasPublicUrls() bool`

HasPublicUrls returns a boolean if a field has been set.

### SetPublicUrlsNil

`func (o *EnvironmentEditDeploymentRunnerSettingsAction) SetPublicUrlsNil(b bool)`

 SetPublicUrlsNil sets the value for PublicUrls to be an explicit nil

### UnsetPublicUrls
`func (o *EnvironmentEditDeploymentRunnerSettingsAction) UnsetPublicUrls()`

UnsetPublicUrls ensures that no value is present for PublicUrls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


