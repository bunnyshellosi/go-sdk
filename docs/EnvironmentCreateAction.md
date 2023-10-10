# EnvironmentCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Project** | **string** |  | 
**Genesis** | Pointer to [**EnvironmentCreateActionGenesis**](EnvironmentCreateActionGenesis.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RemoteDevelopmentAllowed** | Pointer to **bool** |  | [optional] 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**CreateEphemeralOnPrCreate** | Pointer to **bool** |  | [optional] 
**DestroyEphemeralOnPrClose** | Pointer to **bool** |  | [optional] 
**KubernetesIntegration** | Pointer to **NullableString** |  | [optional] 
**EphemeralKubernetesIntegration** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEnvironmentCreateAction

`func NewEnvironmentCreateAction(name string, project string, ) *EnvironmentCreateAction`

NewEnvironmentCreateAction instantiates a new EnvironmentCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateActionWithDefaults

`func NewEnvironmentCreateActionWithDefaults() *EnvironmentCreateAction`

NewEnvironmentCreateActionWithDefaults instantiates a new EnvironmentCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *EnvironmentCreateAction) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentCreateAction) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentCreateAction) SetProject(v string)`

SetProject sets Project field to given value.


### GetGenesis

`func (o *EnvironmentCreateAction) GetGenesis() EnvironmentCreateActionGenesis`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *EnvironmentCreateAction) GetGenesisOk() (*EnvironmentCreateActionGenesis, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *EnvironmentCreateAction) SetGenesis(v EnvironmentCreateActionGenesis)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *EnvironmentCreateAction) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentCreateAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentCreateAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentCreateAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentCreateAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRemoteDevelopmentAllowed

`func (o *EnvironmentCreateAction) GetRemoteDevelopmentAllowed() bool`

GetRemoteDevelopmentAllowed returns the RemoteDevelopmentAllowed field if non-nil, zero value otherwise.

### GetRemoteDevelopmentAllowedOk

`func (o *EnvironmentCreateAction) GetRemoteDevelopmentAllowedOk() (*bool, bool)`

GetRemoteDevelopmentAllowedOk returns a tuple with the RemoteDevelopmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDevelopmentAllowed

`func (o *EnvironmentCreateAction) SetRemoteDevelopmentAllowed(v bool)`

SetRemoteDevelopmentAllowed sets RemoteDevelopmentAllowed field to given value.

### HasRemoteDevelopmentAllowed

`func (o *EnvironmentCreateAction) HasRemoteDevelopmentAllowed() bool`

HasRemoteDevelopmentAllowed returns a boolean if a field has been set.

### GetAutoUpdate

`func (o *EnvironmentCreateAction) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *EnvironmentCreateAction) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *EnvironmentCreateAction) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *EnvironmentCreateAction) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCreateEphemeralOnPrCreate

`func (o *EnvironmentCreateAction) GetCreateEphemeralOnPrCreate() bool`

GetCreateEphemeralOnPrCreate returns the CreateEphemeralOnPrCreate field if non-nil, zero value otherwise.

### GetCreateEphemeralOnPrCreateOk

`func (o *EnvironmentCreateAction) GetCreateEphemeralOnPrCreateOk() (*bool, bool)`

GetCreateEphemeralOnPrCreateOk returns a tuple with the CreateEphemeralOnPrCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEphemeralOnPrCreate

`func (o *EnvironmentCreateAction) SetCreateEphemeralOnPrCreate(v bool)`

SetCreateEphemeralOnPrCreate sets CreateEphemeralOnPrCreate field to given value.

### HasCreateEphemeralOnPrCreate

`func (o *EnvironmentCreateAction) HasCreateEphemeralOnPrCreate() bool`

HasCreateEphemeralOnPrCreate returns a boolean if a field has been set.

### GetDestroyEphemeralOnPrClose

`func (o *EnvironmentCreateAction) GetDestroyEphemeralOnPrClose() bool`

GetDestroyEphemeralOnPrClose returns the DestroyEphemeralOnPrClose field if non-nil, zero value otherwise.

### GetDestroyEphemeralOnPrCloseOk

`func (o *EnvironmentCreateAction) GetDestroyEphemeralOnPrCloseOk() (*bool, bool)`

GetDestroyEphemeralOnPrCloseOk returns a tuple with the DestroyEphemeralOnPrClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyEphemeralOnPrClose

`func (o *EnvironmentCreateAction) SetDestroyEphemeralOnPrClose(v bool)`

SetDestroyEphemeralOnPrClose sets DestroyEphemeralOnPrClose field to given value.

### HasDestroyEphemeralOnPrClose

`func (o *EnvironmentCreateAction) HasDestroyEphemeralOnPrClose() bool`

HasDestroyEphemeralOnPrClose returns a boolean if a field has been set.

### GetKubernetesIntegration

`func (o *EnvironmentCreateAction) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentCreateAction) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentCreateAction) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *EnvironmentCreateAction) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *EnvironmentCreateAction) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentCreateAction) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetEphemeralKubernetesIntegration

`func (o *EnvironmentCreateAction) GetEphemeralKubernetesIntegration() string`

GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field if non-nil, zero value otherwise.

### GetEphemeralKubernetesIntegrationOk

`func (o *EnvironmentCreateAction) GetEphemeralKubernetesIntegrationOk() (*string, bool)`

GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKubernetesIntegration

`func (o *EnvironmentCreateAction) SetEphemeralKubernetesIntegration(v string)`

SetEphemeralKubernetesIntegration sets EphemeralKubernetesIntegration field to given value.

### HasEphemeralKubernetesIntegration

`func (o *EnvironmentCreateAction) HasEphemeralKubernetesIntegration() bool`

HasEphemeralKubernetesIntegration returns a boolean if a field has been set.

### SetEphemeralKubernetesIntegrationNil

`func (o *EnvironmentCreateAction) SetEphemeralKubernetesIntegrationNil(b bool)`

 SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil

### UnsetEphemeralKubernetesIntegration
`func (o *EnvironmentCreateAction) UnsetEphemeralKubernetesIntegration()`

UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
### GetLabels

`func (o *EnvironmentCreateAction) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnvironmentCreateAction) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnvironmentCreateAction) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EnvironmentCreateAction) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


