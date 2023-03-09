# EnvironmentEditAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**RemoteDevelopmentAllowed** | Pointer to **NullableBool** |  | [optional] 
**AutoUpdate** | Pointer to **NullableBool** |  | [optional] 
**KubernetesIntegration** | Pointer to **NullableString** |  | [optional] 
**Edit** | Pointer to [**EnvironmentEditActionEdit**](EnvironmentEditActionEdit.md) |  | [optional] 
**Genesis** | Pointer to [**EnvironmentEditActionGenesis**](EnvironmentEditActionGenesis.md) |  | [optional] 

## Methods

### NewEnvironmentEditAction

`func NewEnvironmentEditAction() *EnvironmentEditAction`

NewEnvironmentEditAction instantiates a new EnvironmentEditAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditActionWithDefaults

`func NewEnvironmentEditActionWithDefaults() *EnvironmentEditAction`

NewEnvironmentEditActionWithDefaults instantiates a new EnvironmentEditAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentEditAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEditAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEditAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEditAction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentEditAction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentEditAction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteDevelopmentAllowed

`func (o *EnvironmentEditAction) GetRemoteDevelopmentAllowed() bool`

GetRemoteDevelopmentAllowed returns the RemoteDevelopmentAllowed field if non-nil, zero value otherwise.

### GetRemoteDevelopmentAllowedOk

`func (o *EnvironmentEditAction) GetRemoteDevelopmentAllowedOk() (*bool, bool)`

GetRemoteDevelopmentAllowedOk returns a tuple with the RemoteDevelopmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDevelopmentAllowed

`func (o *EnvironmentEditAction) SetRemoteDevelopmentAllowed(v bool)`

SetRemoteDevelopmentAllowed sets RemoteDevelopmentAllowed field to given value.

### HasRemoteDevelopmentAllowed

`func (o *EnvironmentEditAction) HasRemoteDevelopmentAllowed() bool`

HasRemoteDevelopmentAllowed returns a boolean if a field has been set.

### SetRemoteDevelopmentAllowedNil

`func (o *EnvironmentEditAction) SetRemoteDevelopmentAllowedNil(b bool)`

 SetRemoteDevelopmentAllowedNil sets the value for RemoteDevelopmentAllowed to be an explicit nil

### UnsetRemoteDevelopmentAllowed
`func (o *EnvironmentEditAction) UnsetRemoteDevelopmentAllowed()`

UnsetRemoteDevelopmentAllowed ensures that no value is present for RemoteDevelopmentAllowed, not even an explicit nil
### GetAutoUpdate

`func (o *EnvironmentEditAction) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *EnvironmentEditAction) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *EnvironmentEditAction) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *EnvironmentEditAction) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### SetAutoUpdateNil

`func (o *EnvironmentEditAction) SetAutoUpdateNil(b bool)`

 SetAutoUpdateNil sets the value for AutoUpdate to be an explicit nil

### UnsetAutoUpdate
`func (o *EnvironmentEditAction) UnsetAutoUpdate()`

UnsetAutoUpdate ensures that no value is present for AutoUpdate, not even an explicit nil
### GetKubernetesIntegration

`func (o *EnvironmentEditAction) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentEditAction) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentEditAction) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *EnvironmentEditAction) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *EnvironmentEditAction) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentEditAction) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetEdit

`func (o *EnvironmentEditAction) GetEdit() EnvironmentEditActionEdit`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *EnvironmentEditAction) GetEditOk() (*EnvironmentEditActionEdit, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *EnvironmentEditAction) SetEdit(v EnvironmentEditActionEdit)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *EnvironmentEditAction) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetGenesis

`func (o *EnvironmentEditAction) GetGenesis() EnvironmentEditActionGenesis`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *EnvironmentEditAction) GetGenesisOk() (*EnvironmentEditActionGenesis, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *EnvironmentEditAction) SetGenesis(v EnvironmentEditActionGenesis)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *EnvironmentEditAction) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


