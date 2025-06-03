# EnvironmentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | Environment type: primary or ephemeral. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment name. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Environment k8s namespace. | [optional] [readonly] 
**UrlHandle** | Pointer to **string** | Environment url handle. | [optional] [readonly] 
**OperationStatus** | Pointer to **string** | Environment operation status. | [optional] [readonly] 
**ClusterStatus** | Pointer to **string** | Environment cluster status. | [optional] [readonly] 
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

### NewEnvironmentCollection

`func NewEnvironmentCollection() *EnvironmentCollection`

NewEnvironmentCollection instantiates a new EnvironmentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCollectionWithDefaults

`func NewEnvironmentCollectionWithDefaults() *EnvironmentCollection`

NewEnvironmentCollectionWithDefaults instantiates a new EnvironmentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentCollection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentCollection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EnvironmentCollection) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EnvironmentCollection) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EnvironmentCollection) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EnvironmentCollection) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUrlHandle

`func (o *EnvironmentCollection) GetUrlHandle() string`

GetUrlHandle returns the UrlHandle field if non-nil, zero value otherwise.

### GetUrlHandleOk

`func (o *EnvironmentCollection) GetUrlHandleOk() (*string, bool)`

GetUrlHandleOk returns a tuple with the UrlHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlHandle

`func (o *EnvironmentCollection) SetUrlHandle(v string)`

SetUrlHandle sets UrlHandle field to given value.

### HasUrlHandle

`func (o *EnvironmentCollection) HasUrlHandle() bool`

HasUrlHandle returns a boolean if a field has been set.

### GetOperationStatus

`func (o *EnvironmentCollection) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *EnvironmentCollection) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *EnvironmentCollection) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *EnvironmentCollection) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetClusterStatus

`func (o *EnvironmentCollection) GetClusterStatus() string`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *EnvironmentCollection) GetClusterStatusOk() (*string, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *EnvironmentCollection) SetClusterStatus(v string)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *EnvironmentCollection) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetProject

`func (o *EnvironmentCollection) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentCollection) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentCollection) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentCollection) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetKubernetesIntegration

`func (o *EnvironmentCollection) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentCollection) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentCollection) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *EnvironmentCollection) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *EnvironmentCollection) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentCollection) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetEphemeralKubernetesIntegration

`func (o *EnvironmentCollection) GetEphemeralKubernetesIntegration() string`

GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field if non-nil, zero value otherwise.

### GetEphemeralKubernetesIntegrationOk

`func (o *EnvironmentCollection) GetEphemeralKubernetesIntegrationOk() (*string, bool)`

GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKubernetesIntegration

`func (o *EnvironmentCollection) SetEphemeralKubernetesIntegration(v string)`

SetEphemeralKubernetesIntegration sets EphemeralKubernetesIntegration field to given value.

### HasEphemeralKubernetesIntegration

`func (o *EnvironmentCollection) HasEphemeralKubernetesIntegration() bool`

HasEphemeralKubernetesIntegration returns a boolean if a field has been set.

### SetEphemeralKubernetesIntegrationNil

`func (o *EnvironmentCollection) SetEphemeralKubernetesIntegrationNil(b bool)`

 SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil

### UnsetEphemeralKubernetesIntegration
`func (o *EnvironmentCollection) UnsetEphemeralKubernetesIntegration()`

UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
### GetHasTerminationProtection

`func (o *EnvironmentCollection) GetHasTerminationProtection() bool`

GetHasTerminationProtection returns the HasTerminationProtection field if non-nil, zero value otherwise.

### GetHasTerminationProtectionOk

`func (o *EnvironmentCollection) GetHasTerminationProtectionOk() (*bool, bool)`

GetHasTerminationProtectionOk returns a tuple with the HasTerminationProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTerminationProtection

`func (o *EnvironmentCollection) SetHasTerminationProtection(v bool)`

SetHasTerminationProtection sets HasTerminationProtection field to given value.

### HasHasTerminationProtection

`func (o *EnvironmentCollection) HasHasTerminationProtection() bool`

HasHasTerminationProtection returns a boolean if a field has been set.

### GetHasAutoUpdate

`func (o *EnvironmentCollection) GetHasAutoUpdate() bool`

GetHasAutoUpdate returns the HasAutoUpdate field if non-nil, zero value otherwise.

### GetHasAutoUpdateOk

`func (o *EnvironmentCollection) GetHasAutoUpdateOk() (*bool, bool)`

GetHasAutoUpdateOk returns a tuple with the HasAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoUpdate

`func (o *EnvironmentCollection) SetHasAutoUpdate(v bool)`

SetHasAutoUpdate sets HasAutoUpdate field to given value.

### HasHasAutoUpdate

`func (o *EnvironmentCollection) HasHasAutoUpdate() bool`

HasHasAutoUpdate returns a boolean if a field has been set.

### GetHasRemoteDevelopmentAllowed

`func (o *EnvironmentCollection) GetHasRemoteDevelopmentAllowed() bool`

GetHasRemoteDevelopmentAllowed returns the HasRemoteDevelopmentAllowed field if non-nil, zero value otherwise.

### GetHasRemoteDevelopmentAllowedOk

`func (o *EnvironmentCollection) GetHasRemoteDevelopmentAllowedOk() (*bool, bool)`

GetHasRemoteDevelopmentAllowedOk returns a tuple with the HasRemoteDevelopmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRemoteDevelopmentAllowed

`func (o *EnvironmentCollection) SetHasRemoteDevelopmentAllowed(v bool)`

SetHasRemoteDevelopmentAllowed sets HasRemoteDevelopmentAllowed field to given value.

### HasHasRemoteDevelopmentAllowed

`func (o *EnvironmentCollection) HasHasRemoteDevelopmentAllowed() bool`

HasHasRemoteDevelopmentAllowed returns a boolean if a field has been set.

### GetHasEphemeralAutoDeploy

`func (o *EnvironmentCollection) GetHasEphemeralAutoDeploy() bool`

GetHasEphemeralAutoDeploy returns the HasEphemeralAutoDeploy field if non-nil, zero value otherwise.

### GetHasEphemeralAutoDeployOk

`func (o *EnvironmentCollection) GetHasEphemeralAutoDeployOk() (*bool, bool)`

GetHasEphemeralAutoDeployOk returns a tuple with the HasEphemeralAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralAutoDeploy

`func (o *EnvironmentCollection) SetHasEphemeralAutoDeploy(v bool)`

SetHasEphemeralAutoDeploy sets HasEphemeralAutoDeploy field to given value.

### HasHasEphemeralAutoDeploy

`func (o *EnvironmentCollection) HasHasEphemeralAutoDeploy() bool`

HasHasEphemeralAutoDeploy returns a boolean if a field has been set.

### GetHasEphemeralCreateOnPr

`func (o *EnvironmentCollection) GetHasEphemeralCreateOnPr() bool`

GetHasEphemeralCreateOnPr returns the HasEphemeralCreateOnPr field if non-nil, zero value otherwise.

### GetHasEphemeralCreateOnPrOk

`func (o *EnvironmentCollection) GetHasEphemeralCreateOnPrOk() (*bool, bool)`

GetHasEphemeralCreateOnPrOk returns a tuple with the HasEphemeralCreateOnPr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralCreateOnPr

`func (o *EnvironmentCollection) SetHasEphemeralCreateOnPr(v bool)`

SetHasEphemeralCreateOnPr sets HasEphemeralCreateOnPr field to given value.

### HasHasEphemeralCreateOnPr

`func (o *EnvironmentCollection) HasHasEphemeralCreateOnPr() bool`

HasHasEphemeralCreateOnPr returns a boolean if a field has been set.

### GetHasEphemeralDestroyOnPrClose

`func (o *EnvironmentCollection) GetHasEphemeralDestroyOnPrClose() bool`

GetHasEphemeralDestroyOnPrClose returns the HasEphemeralDestroyOnPrClose field if non-nil, zero value otherwise.

### GetHasEphemeralDestroyOnPrCloseOk

`func (o *EnvironmentCollection) GetHasEphemeralDestroyOnPrCloseOk() (*bool, bool)`

GetHasEphemeralDestroyOnPrCloseOk returns a tuple with the HasEphemeralDestroyOnPrClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralDestroyOnPrClose

`func (o *EnvironmentCollection) SetHasEphemeralDestroyOnPrClose(v bool)`

SetHasEphemeralDestroyOnPrClose sets HasEphemeralDestroyOnPrClose field to given value.

### HasHasEphemeralDestroyOnPrClose

`func (o *EnvironmentCollection) HasHasEphemeralDestroyOnPrClose() bool`

HasHasEphemeralDestroyOnPrClose returns a boolean if a field has been set.

### GetHasEphemeralBranchWhitelist

`func (o *EnvironmentCollection) GetHasEphemeralBranchWhitelist() bool`

GetHasEphemeralBranchWhitelist returns the HasEphemeralBranchWhitelist field if non-nil, zero value otherwise.

### GetHasEphemeralBranchWhitelistOk

`func (o *EnvironmentCollection) GetHasEphemeralBranchWhitelistOk() (*bool, bool)`

GetHasEphemeralBranchWhitelistOk returns a tuple with the HasEphemeralBranchWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEphemeralBranchWhitelist

`func (o *EnvironmentCollection) SetHasEphemeralBranchWhitelist(v bool)`

SetHasEphemeralBranchWhitelist sets HasEphemeralBranchWhitelist field to given value.

### HasHasEphemeralBranchWhitelist

`func (o *EnvironmentCollection) HasHasEphemeralBranchWhitelist() bool`

HasHasEphemeralBranchWhitelist returns a boolean if a field has been set.

### GetEphemeralBranchWhitelistRegex

`func (o *EnvironmentCollection) GetEphemeralBranchWhitelistRegex() string`

GetEphemeralBranchWhitelistRegex returns the EphemeralBranchWhitelistRegex field if non-nil, zero value otherwise.

### GetEphemeralBranchWhitelistRegexOk

`func (o *EnvironmentCollection) GetEphemeralBranchWhitelistRegexOk() (*string, bool)`

GetEphemeralBranchWhitelistRegexOk returns a tuple with the EphemeralBranchWhitelistRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralBranchWhitelistRegex

`func (o *EnvironmentCollection) SetEphemeralBranchWhitelistRegex(v string)`

SetEphemeralBranchWhitelistRegex sets EphemeralBranchWhitelistRegex field to given value.

### HasEphemeralBranchWhitelistRegex

`func (o *EnvironmentCollection) HasEphemeralBranchWhitelistRegex() bool`

HasEphemeralBranchWhitelistRegex returns a boolean if a field has been set.

### SetEphemeralBranchWhitelistRegexNil

`func (o *EnvironmentCollection) SetEphemeralBranchWhitelistRegexNil(b bool)`

 SetEphemeralBranchWhitelistRegexNil sets the value for EphemeralBranchWhitelistRegex to be an explicit nil

### UnsetEphemeralBranchWhitelistRegex
`func (o *EnvironmentCollection) UnsetEphemeralBranchWhitelistRegex()`

UnsetEphemeralBranchWhitelistRegex ensures that no value is present for EphemeralBranchWhitelistRegex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


