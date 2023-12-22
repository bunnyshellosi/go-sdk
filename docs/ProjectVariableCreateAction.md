# ProjectVariableCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsSecret** | Pointer to **NullableBool** |  | [optional] 
**Project** | **string** |  | 

## Methods

### NewProjectVariableCreateAction

`func NewProjectVariableCreateAction(name string, value string, project string, ) *ProjectVariableCreateAction`

NewProjectVariableCreateAction instantiates a new ProjectVariableCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectVariableCreateActionWithDefaults

`func NewProjectVariableCreateActionWithDefaults() *ProjectVariableCreateAction`

NewProjectVariableCreateActionWithDefaults instantiates a new ProjectVariableCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectVariableCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectVariableCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectVariableCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ProjectVariableCreateAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectVariableCreateAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectVariableCreateAction) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *ProjectVariableCreateAction) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *ProjectVariableCreateAction) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *ProjectVariableCreateAction) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *ProjectVariableCreateAction) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### SetIsSecretNil

`func (o *ProjectVariableCreateAction) SetIsSecretNil(b bool)`

 SetIsSecretNil sets the value for IsSecret to be an explicit nil

### UnsetIsSecret
`func (o *ProjectVariableCreateAction) UnsetIsSecret()`

UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
### GetProject

`func (o *ProjectVariableCreateAction) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectVariableCreateAction) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectVariableCreateAction) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


