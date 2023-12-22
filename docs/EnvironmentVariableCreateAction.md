# EnvironmentVariableCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsSecret** | Pointer to **NullableBool** |  | [optional] 
**Environment** | **string** |  | 

## Methods

### NewEnvironmentVariableCreateAction

`func NewEnvironmentVariableCreateAction(name string, value string, environment string, ) *EnvironmentVariableCreateAction`

NewEnvironmentVariableCreateAction instantiates a new EnvironmentVariableCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableCreateActionWithDefaults

`func NewEnvironmentVariableCreateActionWithDefaults() *EnvironmentVariableCreateAction`

NewEnvironmentVariableCreateActionWithDefaults instantiates a new EnvironmentVariableCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentVariableCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentVariableCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentVariableCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EnvironmentVariableCreateAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableCreateAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableCreateAction) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *EnvironmentVariableCreateAction) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *EnvironmentVariableCreateAction) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *EnvironmentVariableCreateAction) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *EnvironmentVariableCreateAction) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### SetIsSecretNil

`func (o *EnvironmentVariableCreateAction) SetIsSecretNil(b bool)`

 SetIsSecretNil sets the value for IsSecret to be an explicit nil

### UnsetIsSecret
`func (o *EnvironmentVariableCreateAction) UnsetIsSecret()`

UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentVariableCreateAction) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentVariableCreateAction) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentVariableCreateAction) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


