# Primary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "primary"]
**AutoDeployEphemeral** | Pointer to **NullableBool** |  | [optional] 
**TerminationProtection** | Pointer to **NullableBool** |  | [optional] 
**CreateEphemeralOnPrCreate** | Pointer to **NullableBool** |  | [optional] 
**DestroyEphemeralOnPrClose** | Pointer to **NullableBool** |  | [optional] 
**EphemeralKubernetesIntegration** | **NullableString** |  | 
**PrimaryOptions** | Pointer to [**NullableEditPrimaryOptions**](EditPrimaryOptions.md) |  | [optional] 

## Methods

### NewPrimary

`func NewPrimary(ephemeralKubernetesIntegration NullableString, ) *Primary`

NewPrimary instantiates a new Primary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryWithDefaults

`func NewPrimaryWithDefaults() *Primary`

NewPrimaryWithDefaults instantiates a new Primary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Primary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Primary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Primary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Primary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoDeployEphemeral

`func (o *Primary) GetAutoDeployEphemeral() bool`

GetAutoDeployEphemeral returns the AutoDeployEphemeral field if non-nil, zero value otherwise.

### GetAutoDeployEphemeralOk

`func (o *Primary) GetAutoDeployEphemeralOk() (*bool, bool)`

GetAutoDeployEphemeralOk returns a tuple with the AutoDeployEphemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEphemeral

`func (o *Primary) SetAutoDeployEphemeral(v bool)`

SetAutoDeployEphemeral sets AutoDeployEphemeral field to given value.

### HasAutoDeployEphemeral

`func (o *Primary) HasAutoDeployEphemeral() bool`

HasAutoDeployEphemeral returns a boolean if a field has been set.

### SetAutoDeployEphemeralNil

`func (o *Primary) SetAutoDeployEphemeralNil(b bool)`

 SetAutoDeployEphemeralNil sets the value for AutoDeployEphemeral to be an explicit nil

### UnsetAutoDeployEphemeral
`func (o *Primary) UnsetAutoDeployEphemeral()`

UnsetAutoDeployEphemeral ensures that no value is present for AutoDeployEphemeral, not even an explicit nil
### GetTerminationProtection

`func (o *Primary) GetTerminationProtection() bool`

GetTerminationProtection returns the TerminationProtection field if non-nil, zero value otherwise.

### GetTerminationProtectionOk

`func (o *Primary) GetTerminationProtectionOk() (*bool, bool)`

GetTerminationProtectionOk returns a tuple with the TerminationProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtection

`func (o *Primary) SetTerminationProtection(v bool)`

SetTerminationProtection sets TerminationProtection field to given value.

### HasTerminationProtection

`func (o *Primary) HasTerminationProtection() bool`

HasTerminationProtection returns a boolean if a field has been set.

### SetTerminationProtectionNil

`func (o *Primary) SetTerminationProtectionNil(b bool)`

 SetTerminationProtectionNil sets the value for TerminationProtection to be an explicit nil

### UnsetTerminationProtection
`func (o *Primary) UnsetTerminationProtection()`

UnsetTerminationProtection ensures that no value is present for TerminationProtection, not even an explicit nil
### GetCreateEphemeralOnPrCreate

`func (o *Primary) GetCreateEphemeralOnPrCreate() bool`

GetCreateEphemeralOnPrCreate returns the CreateEphemeralOnPrCreate field if non-nil, zero value otherwise.

### GetCreateEphemeralOnPrCreateOk

`func (o *Primary) GetCreateEphemeralOnPrCreateOk() (*bool, bool)`

GetCreateEphemeralOnPrCreateOk returns a tuple with the CreateEphemeralOnPrCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEphemeralOnPrCreate

`func (o *Primary) SetCreateEphemeralOnPrCreate(v bool)`

SetCreateEphemeralOnPrCreate sets CreateEphemeralOnPrCreate field to given value.

### HasCreateEphemeralOnPrCreate

`func (o *Primary) HasCreateEphemeralOnPrCreate() bool`

HasCreateEphemeralOnPrCreate returns a boolean if a field has been set.

### SetCreateEphemeralOnPrCreateNil

`func (o *Primary) SetCreateEphemeralOnPrCreateNil(b bool)`

 SetCreateEphemeralOnPrCreateNil sets the value for CreateEphemeralOnPrCreate to be an explicit nil

### UnsetCreateEphemeralOnPrCreate
`func (o *Primary) UnsetCreateEphemeralOnPrCreate()`

UnsetCreateEphemeralOnPrCreate ensures that no value is present for CreateEphemeralOnPrCreate, not even an explicit nil
### GetDestroyEphemeralOnPrClose

`func (o *Primary) GetDestroyEphemeralOnPrClose() bool`

GetDestroyEphemeralOnPrClose returns the DestroyEphemeralOnPrClose field if non-nil, zero value otherwise.

### GetDestroyEphemeralOnPrCloseOk

`func (o *Primary) GetDestroyEphemeralOnPrCloseOk() (*bool, bool)`

GetDestroyEphemeralOnPrCloseOk returns a tuple with the DestroyEphemeralOnPrClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyEphemeralOnPrClose

`func (o *Primary) SetDestroyEphemeralOnPrClose(v bool)`

SetDestroyEphemeralOnPrClose sets DestroyEphemeralOnPrClose field to given value.

### HasDestroyEphemeralOnPrClose

`func (o *Primary) HasDestroyEphemeralOnPrClose() bool`

HasDestroyEphemeralOnPrClose returns a boolean if a field has been set.

### SetDestroyEphemeralOnPrCloseNil

`func (o *Primary) SetDestroyEphemeralOnPrCloseNil(b bool)`

 SetDestroyEphemeralOnPrCloseNil sets the value for DestroyEphemeralOnPrClose to be an explicit nil

### UnsetDestroyEphemeralOnPrClose
`func (o *Primary) UnsetDestroyEphemeralOnPrClose()`

UnsetDestroyEphemeralOnPrClose ensures that no value is present for DestroyEphemeralOnPrClose, not even an explicit nil
### GetEphemeralKubernetesIntegration

`func (o *Primary) GetEphemeralKubernetesIntegration() string`

GetEphemeralKubernetesIntegration returns the EphemeralKubernetesIntegration field if non-nil, zero value otherwise.

### GetEphemeralKubernetesIntegrationOk

`func (o *Primary) GetEphemeralKubernetesIntegrationOk() (*string, bool)`

GetEphemeralKubernetesIntegrationOk returns a tuple with the EphemeralKubernetesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKubernetesIntegration

`func (o *Primary) SetEphemeralKubernetesIntegration(v string)`

SetEphemeralKubernetesIntegration sets EphemeralKubernetesIntegration field to given value.


### SetEphemeralKubernetesIntegrationNil

`func (o *Primary) SetEphemeralKubernetesIntegrationNil(b bool)`

 SetEphemeralKubernetesIntegrationNil sets the value for EphemeralKubernetesIntegration to be an explicit nil

### UnsetEphemeralKubernetesIntegration
`func (o *Primary) UnsetEphemeralKubernetesIntegration()`

UnsetEphemeralKubernetesIntegration ensures that no value is present for EphemeralKubernetesIntegration, not even an explicit nil
### GetPrimaryOptions

`func (o *Primary) GetPrimaryOptions() EditPrimaryOptions`

GetPrimaryOptions returns the PrimaryOptions field if non-nil, zero value otherwise.

### GetPrimaryOptionsOk

`func (o *Primary) GetPrimaryOptionsOk() (*EditPrimaryOptions, bool)`

GetPrimaryOptionsOk returns a tuple with the PrimaryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryOptions

`func (o *Primary) SetPrimaryOptions(v EditPrimaryOptions)`

SetPrimaryOptions sets PrimaryOptions field to given value.

### HasPrimaryOptions

`func (o *Primary) HasPrimaryOptions() bool`

HasPrimaryOptions returns a boolean if a field has been set.

### SetPrimaryOptionsNil

`func (o *Primary) SetPrimaryOptionsNil(b bool)`

 SetPrimaryOptionsNil sets the value for PrimaryOptions to be an explicit nil

### UnsetPrimaryOptions
`func (o *Primary) UnsetPrimaryOptions()`

UnsetPrimaryOptions ensures that no value is present for PrimaryOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


