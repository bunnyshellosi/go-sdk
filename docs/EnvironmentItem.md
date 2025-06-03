# EnvironmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Environment labels. | [optional] [readonly] 
**Type** | Pointer to **string** | Environment type: primary or ephemeral. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment name. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Environment k8s namespace. | [optional] [readonly] 
**UrlHandle** | Pointer to **string** | Environment url handle. | [optional] [readonly] 
**TotalComponents** | Pointer to **int32** | Service component identifier | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Environment operation status. | [optional] [readonly] 
**ClusterStatus** | Pointer to **string** | Environment cluster status. | [optional] [readonly] 
**BuildSettings** | Pointer to [**NullableBuildSettingsItem**](BuildSettingsItem.md) |  | [optional] 
**Project** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**KubernetesIntegration** | Pointer to **NullableString** | Kubernetes integration identifier. | [optional] [readonly] 
**EphemeralKubernetesIntegration** | Pointer to **NullableString** | Kubernetes integration identifier for the ephemeral auto deploy cluster. | [optional] [readonly] 
**HasTerminationProtection** | Pointer to **bool** | Environment termination protection. | [optional] [readonly] 
**HasAutoUpdate** | Pointer to **bool** | Environment auto update. | [optional] [readonly] 
**HasRemoteDevelopmentAllowed** | Pointer to **bool** | Environment remote development allowed. | [optional] [readonly] 
**HasEphemeralAutoDeploy** | Pointer to **bool** | Environment ephemeral auto deploy status. | [optional] [readonly] 
**HasEphemeralCreateOnPr** | Pointer to **bool** | Environment ephemeral create-on-PR status. | [optional] [readonly] 
**HasEphemeralDestroyOnPrClose** | Pointer to **bool** | Environment ephemeral destroy-on-PR close status. | [optional] [readonly] 
**HasEphemeralBranchWhitelist** | Pointer to **bool** | Environment ephemeral branch whitelist status. | [optional] [readonly] 
**EphemeralBranchWhitelistRegex** | Pointer to **NullableString** | Environment ephemeral branch whitelist regex. | [optional] [readonly] 

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

### GetLabels

`func (o *EnvironmentItem) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnvironmentItem) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnvironmentItem) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EnvironmentItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

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

### GetUrlHandle

`func (o *EnvironmentItem) GetUrlHandle() string`

GetUrlHandle returns the UrlHandle field if non-nil, zero value otherwise.

### GetUrlHandleOk

`func (o *EnvironmentItem) GetUrlHandleOk() (*string, bool)`

GetUrlHandleOk returns a tuple with the UrlHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlHandle

`func (o *EnvironmentItem) SetUrlHandle(v string)`

SetUrlHandle sets UrlHandle field to given value.

### HasUrlHandle

`func (o *EnvironmentItem) HasUrlHandle() bool`

HasUrlHandle returns a boolean if a field has been set.

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

### GetBuildSettings

`func (o *EnvironmentItem) GetBuildSettings() BuildSettingsItem`

GetBuildSettings returns the BuildSettings field if non-nil, zero value otherwise.

### GetBuildSettingsOk

`func (o *EnvironmentItem) GetBuildSettingsOk() (*BuildSettingsItem, bool)`

GetBuildSettingsOk returns a tuple with the BuildSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSettings

`func (o *EnvironmentItem) SetBuildSettings(v BuildSettingsItem)`

SetBuildSettings sets BuildSettings field to given value.

### HasBuildSettings

`func (o *EnvironmentItem) HasBuildSettings() bool`

HasBuildSettings returns a boolean if a field has been set.

### SetBuildSettingsNil

`func (o *EnvironmentItem) SetBuildSettingsNil(b bool)`

 SetBuildSettingsNil sets the value for BuildSettings to be an explicit nil

### UnsetBuildSettings
`func (o *EnvironmentItem) UnsetBuildSettings()`

UnsetBuildSettings ensures that no value is present for BuildSettings, not even an explicit nil
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
### GetEphemeralKubernetesIntegration

`func (o *EnvironmentItem) GetEphemeralKubernetesIntegration() string`

GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field if non-nil, zero value otherwise.

### GetEphemeralKubernetesIntegrationOk

`func (o *EnvironmentItem) GetEphemeralKubernetesIntegrationOk() (*string, bool)`

GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKubernetesIntegration

`func (o *EnvironmentItem) SetEphemeralKubernetesIntegration(v string)`

SetEphemeralKubernetesIntegration sets EphemeralKubernetesIntegration field to given value.

### HasEphemeralKubernetesIntegration

`func (o *EnvironmentItem) HasEphemeralKubernetesIntegration() bool`

HasEphemeralKubernetesIntegration returns a boolean if a field has been set.

### SetEphemeralKubernetesIntegrationNil

`func (o *EnvironmentItem) SetEphemeralKubernetesIntegrationNil(b bool)`

 SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil

### UnsetEphemeralKubernetesIntegration
`func (o *EnvironmentItem) UnsetEphemeralKubernetesIntegration()`

UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
### GetHasTerminationProtection

`func (o *EnvironmentItem) GetHasTerminationProtection() bool`

GetHasTerminationProtection returns the HasTerminationProtection field if non-nil, zero value otherwise.

### GetHasTerminationProtectionOk

`func (o *EnvironmentItem) GetHasTerminationProtectionOk() (*bool, bool)`

GetHasTerminationProtectionOk returns a tuple with the HasTerminationProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTerminationProtection

`func (o *EnvironmentItem) SetHasTerminationProtection(v bool)`

SetHasTerminationProtection sets HasTerminationProtection field to given value.

### HasHasTerminationProtection

`func (o *EnvironmentItem) HasHasTerminationProtection() bool`

HasHasTerminationProtection returns a boolean if a field has been set.

### GetHasAutoUpdate

`func (o *EnvironmentItem) GetHasAutoUpdate() bool`

GetHasAutoUpdate returns the HasAutoUpdate field if non-nil, zero value otherwise.

### GetHasAutoUpdateOk

`func (o *EnvironmentItem) GetHasAutoUpdateOk() (*bool, bool)`

GetHasAutoUpdateOk returns a tuple with the HasAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoUpdate

`func (o *EnvironmentItem) SetHasAutoUpdate(v bool)`

SetHasAutoUpdate sets HasAutoUpdate field to given value.

### HasHasAutoUpdate

`func (o *EnvironmentItem) HasHasAutoUpdate() bool`

HasHasAutoUpdate returns a boolean if a field has been set.

### GetHasRemoteDevelopmentAllowed

`func (o *EnvironmentItem) GetHasRemoteDevelopmentAllowed() bool`

GetHasRemoteDevelopmentAllowed returns the HasRemoteDevelopmentAllowed field if non-nil, zero value otherwise.

### GetHasRemoteDevelopmentAllowedOk

`func (o *EnvironmentItem) GetHasRemoteDevelopmentAllowedOk() (*bool, bool)`

GetHasRemoteDevelopmentAllowedOk returns a tuple with the HasRemoteDevelopmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRemoteDevelopmentAllowed

`func (o *EnvironmentItem) SetHasRemoteDevelopmentAllowed(v bool)`

SetHasRemoteDevelopmentAllowed sets HasRemoteDevelopmentAllowed field to given value.

### HasHasRemoteDevelopmentAllowed

`func (o *EnvironmentItem) HasHasRemoteDevelopmentAllowed() bool`

HasHasRemoteDevelopmentAllowed returns a boolean if a field has been set.

### GetHasEphemeralAutoDeploy

`func (o *EnvironmentItem) GetHasEphemeralAutoDeploy() bool`

GetHasEphemeralAutoDeploy returns the HasEphemeralAutoDeploy field if non-nil, zero value otherwise.

### GetHasEphemeralAutoDeployOk

`func (o *EnvironmentItem) GetHasEphemeralAutoDeployOk() (*bool, bool)`

GetHasEphemeralAutoDeployOk returns a tuple with the HasEphemeralAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralAutoDeploy

`func (o *EnvironmentItem) SetHasEphemeralAutoDeploy(v bool)`

SetHasEphemeralAutoDeploy sets HasEphemeralAutoDeploy field to given value.

### HasHasEphemeralAutoDeploy

`func (o *EnvironmentItem) HasHasEphemeralAutoDeploy() bool`

HasHasEphemeralAutoDeploy returns a boolean if a field has been set.

### GetHasEphemeralCreateOnPr

`func (o *EnvironmentItem) GetHasEphemeralCreateOnPr() bool`

GetHasEphemeralCreateOnPr returns the HasEphemeralCreateOnPr field if non-nil, zero value otherwise.

### GetHasEphemeralCreateOnPrOk

`func (o *EnvironmentItem) GetHasEphemeralCreateOnPrOk() (*bool, bool)`

GetHasEphemeralCreateOnPrOk returns a tuple with the HasEphemeralCreateOnPr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralCreateOnPr

`func (o *EnvironmentItem) SetHasEphemeralCreateOnPr(v bool)`

SetHasEphemeralCreateOnPr sets HasEphemeralCreateOnPr field to given value.

### HasHasEphemeralCreateOnPr

`func (o *EnvironmentItem) HasHasEphemeralCreateOnPr() bool`

HasHasEphemeralCreateOnPr returns a boolean if a field has been set.

### GetHasEphemeralDestroyOnPrClose

`func (o *EnvironmentItem) GetHasEphemeralDestroyOnPrClose() bool`

GetHasEphemeralDestroyOnPrClose returns the HasEphemeralDestroyOnPrClose field if non-nil, zero value otherwise.

### GetHasEphemeralDestroyOnPrCloseOk

`func (o *EnvironmentItem) GetHasEphemeralDestroyOnPrCloseOk() (*bool, bool)`

GetHasEphemeralDestroyOnPrCloseOk returns a tuple with the HasEphemeralDestroyOnPrClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralDestroyOnPrClose

`func (o *EnvironmentItem) SetHasEphemeralDestroyOnPrClose(v bool)`

SetHasEphemeralDestroyOnPrClose sets HasEphemeralDestroyOnPrClose field to given value.

### HasHasEphemeralDestroyOnPrClose

`func (o *EnvironmentItem) HasHasEphemeralDestroyOnPrClose() bool`

HasHasEphemeralDestroyOnPrClose returns a boolean if a field has been set.

### GetHasEphemeralBranchWhitelist

`func (o *EnvironmentItem) GetHasEphemeralBranchWhitelist() bool`

GetHasEphemeralBranchWhitelist returns the HasEphemeralBranchWhitelist field if non-nil, zero value otherwise.

### GetHasEphemeralBranchWhitelistOk

`func (o *EnvironmentItem) GetHasEphemeralBranchWhitelistOk() (*bool, bool)`

GetHasEphemeralBranchWhitelistOk returns a tuple with the HasEphemeralBranchWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralBranchWhitelist

`func (o *EnvironmentItem) SetHasEphemeralBranchWhitelist(v bool)`

SetHasEphemeralBranchWhitelist sets HasEphemeralBranchWhitelist field to given value.

### HasHasEphemeralBranchWhitelist

`func (o *EnvironmentItem) HasHasEphemeralBranchWhitelist() bool`

HasHasEphemeralBranchWhitelist returns a boolean if a field has been set.

### GetEphemeralBranchWhitelistRegex

`func (o *EnvironmentItem) GetEphemeralBranchWhitelistRegex() string`

GetEphemeralBranchWhitelistRegex returns the EphemeralBranchWhitelistRegex field if non-nil, zero value otherwise.

### GetEphemeralBranchWhitelistRegexOk

`func (o *EnvironmentItem) GetEphemeralBranchWhitelistRegexOk() (*string, bool)`

GetEphemeralBranchWhitelistRegexOk returns a tuple with the EphemeralBranchWhitelistRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralBranchWhitelistRegex

`func (o *EnvironmentItem) SetEphemeralBranchWhitelistRegex(v string)`

SetEphemeralBranchWhitelistRegex sets EphemeralBranchWhitelistRegex field to given value.

### HasEphemeralBranchWhitelistRegex

`func (o *EnvironmentItem) HasEphemeralBranchWhitelistRegex() bool`

HasEphemeralBranchWhitelistRegex returns a boolean if a field has been set.

### SetEphemeralBranchWhitelistRegexNil

`func (o *EnvironmentItem) SetEphemeralBranchWhitelistRegexNil(b bool)`

 SetEphemeralBranchWhitelistRegexNil sets the value for EphemeralBranchWhitelistRegex to be an explicit nil

### UnsetEphemeralBranchWhitelistRegex
`func (o *EnvironmentItem) UnsetEphemeralBranchWhitelistRegex()`

UnsetEphemeralBranchWhitelistRegex ensures that no value is present for EphemeralBranchWhitelistRegex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


