# EnvironmentEditSettingsEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "ephemeral"]
**AutoDeployEphemeral** | Pointer to **NullableBool** |  | [optional] 
**TerminationProtection** | Pointer to **NullableBool** |  | [optional] 
**CreateEphemeralOnPrCreate** | Pointer to **NullableBool** |  | [optional] 
**DestroyEphemeralOnPrClose** | Pointer to **NullableBool** |  | [optional] 
**EphemeralKubernetesIntegration** | **NullableString** |  | 
**PrimaryOptions** | Pointer to [**NullablePrimaryOptionsEdit**](PrimaryOptionsEdit.md) |  | [optional] 

## Methods

### NewEnvironmentEditSettingsEdit

`func NewEnvironmentEditSettingsEdit(ephemeralKubernetesIntegration NullableString, ) *EnvironmentEditSettingsEdit`

NewEnvironmentEditSettingsEdit instantiates a new EnvironmentEditSettingsEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEditSettingsEditWithDefaults

`func NewEnvironmentEditSettingsEditWithDefaults() *EnvironmentEditSettingsEdit`

NewEnvironmentEditSettingsEditWithDefaults instantiates a new EnvironmentEditSettingsEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnvironmentEditSettingsEdit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentEditSettingsEdit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentEditSettingsEdit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentEditSettingsEdit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoDeployEphemeral

`func (o *EnvironmentEditSettingsEdit) GetAutoDeployEphemeral() bool`

GetAutoDeployEphemeral returns the AutoDeployEphemeral field if non-nil, zero value otherwise.

### GetAutoDeployEphemeralOk

`func (o *EnvironmentEditSettingsEdit) GetAutoDeployEphemeralOk() (*bool, bool)`

GetAutoDeployEphemeralOk returns a tuple with the AutoDeployEphemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEphemeral

`func (o *EnvironmentEditSettingsEdit) SetAutoDeployEphemeral(v bool)`

SetAutoDeployEphemeral sets AutoDeployEphemeral field to given value.

### HasAutoDeployEphemeral

`func (o *EnvironmentEditSettingsEdit) HasAutoDeployEphemeral() bool`

HasAutoDeployEphemeral returns a boolean if a field has been set.

### SetAutoDeployEphemeralNil

`func (o *EnvironmentEditSettingsEdit) SetAutoDeployEphemeralNil(b bool)`

 SetAutoDeployEphemeralNil sets the value for AutoDeployEphemeral to be an explicit nil

### UnsetAutoDeployEphemeral
`func (o *EnvironmentEditSettingsEdit) UnsetAutoDeployEphemeral()`

UnsetAutoDeployEphemeral ensures that no value is present for AutoDeployEphemeral, not even an explicit nil
### GetTerminationProtection

`func (o *EnvironmentEditSettingsEdit) GetTerminationProtection() bool`

GetTerminationProtection returns the TerminationProtection field if non-nil, zero value otherwise.

### GetTerminationProtectionOk

`func (o *EnvironmentEditSettingsEdit) GetTerminationProtectionOk() (*bool, bool)`

GetTerminationProtectionOk returns a tuple with the TerminationProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtection

`func (o *EnvironmentEditSettingsEdit) SetTerminationProtection(v bool)`

SetTerminationProtection sets TerminationProtection field to given value.

### HasTerminationProtection

`func (o *EnvironmentEditSettingsEdit) HasTerminationProtection() bool`

HasTerminationProtection returns a boolean if a field has been set.

### SetTerminationProtectionNil

`func (o *EnvironmentEditSettingsEdit) SetTerminationProtectionNil(b bool)`

 SetTerminationProtectionNil sets the value for TerminationProtection to be an explicit nil

### UnsetTerminationProtection
`func (o *EnvironmentEditSettingsEdit) UnsetTerminationProtection()`

UnsetTerminationProtection ensures that no value is present for TerminationProtection, not even an explicit nil
### GetCreateEphemeralOnPrCreate

`func (o *EnvironmentEditSettingsEdit) GetCreateEphemeralOnPrCreate() bool`

GetCreateEphemeralOnPrCreate returns the CreateEphemeralOnPrCreate field if non-nil, zero value otherwise.

### GetCreateEphemeralOnPrCreateOk

`func (o *EnvironmentEditSettingsEdit) GetCreateEphemeralOnPrCreateOk() (*bool, bool)`

GetCreateEphemeralOnPrCreateOk returns a tuple with the CreateEphemeralOnPrCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEphemeralOnPrCreate

`func (o *EnvironmentEditSettingsEdit) SetCreateEphemeralOnPrCreate(v bool)`

SetCreateEphemeralOnPrCreate sets CreateEphemeralOnPrCreate field to given value.

### HasCreateEphemeralOnPrCreate

`func (o *EnvironmentEditSettingsEdit) HasCreateEphemeralOnPrCreate() bool`

HasCreateEphemeralOnPrCreate returns a boolean if a field has been set.

### SetCreateEphemeralOnPrCreateNil

`func (o *EnvironmentEditSettingsEdit) SetCreateEphemeralOnPrCreateNil(b bool)`

 SetCreateEphemeralOnPrCreateNil sets the value for CreateEphemeralOnPrCreate to be an explicit nil

### UnsetCreateEphemeralOnPrCreate
`func (o *EnvironmentEditSettingsEdit) UnsetCreateEphemeralOnPrCreate()`

UnsetCreateEphemeralOnPrCreate ensures that no value is present for CreateEphemeralOnPrCreate, not even an explicit nil
### GetDestroyEphemeralOnPrClose

`func (o *EnvironmentEditSettingsEdit) GetDestroyEphemeralOnPrClose() bool`

GetDestroyEphemeralOnPrClose returns the DestroyEphemeralOnPrClose field if non-nil, zero value otherwise.

### GetDestroyEphemeralOnPrCloseOk

`func (o *EnvironmentEditSettingsEdit) GetDestroyEphemeralOnPrCloseOk() (*bool, bool)`

GetDestroyEphemeralOnPrCloseOk returns a tuple with the DestroyEphemeralOnPrClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyEphemeralOnPrClose

`func (o *EnvironmentEditSettingsEdit) SetDestroyEphemeralOnPrClose(v bool)`

SetDestroyEphemeralOnPrClose sets DestroyEphemeralOnPrClose field to given value.

### HasDestroyEphemeralOnPrClose

`func (o *EnvironmentEditSettingsEdit) HasDestroyEphemeralOnPrClose() bool`

HasDestroyEphemeralOnPrClose returns a boolean if a field has been set.

### SetDestroyEphemeralOnPrCloseNil

`func (o *EnvironmentEditSettingsEdit) SetDestroyEphemeralOnPrCloseNil(b bool)`

 SetDestroyEphemeralOnPrCloseNil sets the value for DestroyEphemeralOnPrClose to be an explicit nil

### UnsetDestroyEphemeralOnPrClose
`func (o *EnvironmentEditSettingsEdit) UnsetDestroyEphemeralOnPrClose()`

UnsetDestroyEphemeralOnPrClose ensures that no value is present for DestroyEphemeralOnPrClose, not even an explicit nil
### GetEphemeralKubernetesIntegration

`func (o *EnvironmentEditSettingsEdit) GetEphemeralKubernetesIntegration() string`

GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field if non-nil, zero value otherwise.

### GetEphemeralKubernetesIntegrationOk

`func (o *EnvironmentEditSettingsEdit) GetEphemeralKubernetesIntegrationOk() (*string, bool)`

GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKubernetesIntegration

`func (o *EnvironmentEditSettingsEdit) SetEphemeralKubernetesIntegration(v string)`

SetEphemeralKubernetesIntegration sets EphemeralKubernetesIntegration field to given value.


### SetEphemeralKubernetesIntegrationNil

`func (o *EnvironmentEditSettingsEdit) SetEphemeralKubernetesIntegrationNil(b bool)`

 SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil

### UnsetEphemeralKubernetesIntegration
`func (o *EnvironmentEditSettingsEdit) UnsetEphemeralKubernetesIntegration()`

UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
### GetPrimaryOptions

`func (o *EnvironmentEditSettingsEdit) GetPrimaryOptions() PrimaryOptionsEdit`

GetPrimaryOptions returns the PrimaryOptions field if non-nil, zero value otherwise.

### GetPrimaryOptionsOk

`func (o *EnvironmentEditSettingsEdit) GetPrimaryOptionsOk() (*PrimaryOptionsEdit, bool)`

GetPrimaryOptionsOk returns a tuple with the PrimaryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryOptions

`func (o *EnvironmentEditSettingsEdit) SetPrimaryOptions(v PrimaryOptionsEdit)`

SetPrimaryOptions sets PrimaryOptions field to given value.

### HasPrimaryOptions

`func (o *EnvironmentEditSettingsEdit) HasPrimaryOptions() bool`

HasPrimaryOptions returns a boolean if a field has been set.

### SetPrimaryOptionsNil

`func (o *EnvironmentEditSettingsEdit) SetPrimaryOptionsNil(b bool)`

 SetPrimaryOptionsNil sets the value for PrimaryOptions to be an explicit nil

### UnsetPrimaryOptions
`func (o *EnvironmentEditSettingsEdit) UnsetPrimaryOptions()`

UnsetPrimaryOptions ensures that no value is present for PrimaryOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


