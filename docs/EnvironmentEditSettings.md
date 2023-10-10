# EnvironmentEditSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**RemoteDevelopmentAllowed** | Pointer to **NullableBool** |  | [optional] 
**AutoUpdate** | Pointer to **NullableBool** |  | [optional] 
**KubernetesIntegration** | Pointer to **NullableString** |  | [optional] 
**Edit** | Pointer to [**EnvironmentEditSettingsEdit**](EnvironmentEditSettingsEdit.md) |  | [optional] 
**Labels** | Pointer to [**NullableEdit**](Edit.md) |  | [optional] 

## Methods

### NewEnvironmentEditSettings

`func NewEnvironmentEditSettings() *EnvironmentEditSettings`

NewEnvironmentEditSettings instantiates a new EnvironmentEditSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditSettingsWithDefaults

`func NewEnvironmentEditSettingsWithDefaults() *EnvironmentEditSettings`

NewEnvironmentEditSettingsWithDefaults instantiates a new EnvironmentEditSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentEditSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEditSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEditSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEditSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentEditSettings) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentEditSettings) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteDevelopmentAllowed

`func (o *EnvironmentEditSettings) GetRemoteDevelopmentAllowed() bool`

GetRemoteDevelopmentAllowed returns the RemoteDevelopmentAllowed field if non-nil, zero value otherwise.

### GetRemoteDevelopmentAllowedOk

`func (o *EnvironmentEditSettings) GetRemoteDevelopmentAllowedOk() (*bool, bool)`

GetRemoteDevelopmentAllowedOk returns a tuple with the RemoteDevelopmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDevelopmentAllowed

`func (o *EnvironmentEditSettings) SetRemoteDevelopmentAllowed(v bool)`

SetRemoteDevelopmentAllowed sets RemoteDevelopmentAllowed field to given value.

### HasRemoteDevelopmentAllowed

`func (o *EnvironmentEditSettings) HasRemoteDevelopmentAllowed() bool`

HasRemoteDevelopmentAllowed returns a boolean if a field has been set.

### SetRemoteDevelopmentAllowedNil

`func (o *EnvironmentEditSettings) SetRemoteDevelopmentAllowedNil(b bool)`

 SetRemoteDevelopmentAllowedNil sets the value for RemoteDevelopmentAllowed to be an explicit nil

### UnsetRemoteDevelopmentAllowed
`func (o *EnvironmentEditSettings) UnsetRemoteDevelopmentAllowed()`

UnsetRemoteDevelopmentAllowed ensures that no value is present for RemoteDevelopmentAllowed, not even an explicit nil
### GetAutoUpdate

`func (o *EnvironmentEditSettings) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *EnvironmentEditSettings) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *EnvironmentEditSettings) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *EnvironmentEditSettings) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### SetAutoUpdateNil

`func (o *EnvironmentEditSettings) SetAutoUpdateNil(b bool)`

 SetAutoUpdateNil sets the value for AutoUpdate to be an explicit nil

### UnsetAutoUpdate
`func (o *EnvironmentEditSettings) UnsetAutoUpdate()`

UnsetAutoUpdate ensures that no value is present for AutoUpdate, not even an explicit nil
### GetKubernetesIntegration

`func (o *EnvironmentEditSettings) GetKubernetesIntegration() string`

GetKubernetesIntegration returns the KubernetesIntegration field if non-nil, zero value otherwise.

### GetKubernetesIntegrationOk

`func (o *EnvironmentEditSettings) GetKubernetesIntegrationOk() (*string, bool)`

GetKubernetesIntegrationOk returns a tuple with the KubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIntegration

`func (o *EnvironmentEditSettings) SetKubernetesIntegration(v string)`

SetKubernetesIntegration sets KubernetesIntegration field to given value.

### HasKubernetesIntegration

`func (o *EnvironmentEditSettings) HasKubernetesIntegration() bool`

HasKubernetesIntegration returns a boolean if a field has been set.

### SetKubernetesIntegrationNil

`func (o *EnvironmentEditSettings) SetKubernetesIntegrationNil(b bool)`

 SetKubernetesIntegrationNil sets the value for KubernetesIntegration to be an explicit nil

### UnsetKubernetesIntegration
`func (o *EnvironmentEditSettings) UnsetKubernetesIntegration()`

UnsetKubernetesIntegration ensures that no value is present for KubernetesIntegration, not even an explicit nil
### GetEdit

`func (o *EnvironmentEditSettings) GetEdit() EnvironmentEditSettingsEdit`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *EnvironmentEditSettings) GetEditOk() (*EnvironmentEditSettingsEdit, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *EnvironmentEditSettings) SetEdit(v EnvironmentEditSettingsEdit)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *EnvironmentEditSettings) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetLabels

`func (o *EnvironmentEditSettings) GetLabels() Edit`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnvironmentEditSettings) GetLabelsOk() (*Edit, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnvironmentEditSettings) SetLabels(v Edit)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EnvironmentEditSettings) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *EnvironmentEditSettings) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *EnvironmentEditSettings) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


