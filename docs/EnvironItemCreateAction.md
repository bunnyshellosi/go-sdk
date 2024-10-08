# EnvironItemCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsSecret** | Pointer to **bool** |  | [optional] 
**Environment** | **string** |  | 

## Methods

### NewEnvironItemCreateAction

`func NewEnvironItemCreateAction(groupName string, name string, value string, environment string, ) *EnvironItemCreateAction`

NewEnvironItemCreateAction instantiates a new EnvironItemCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironItemCreateActionWithDefaults

`func NewEnvironItemCreateActionWithDefaults() *EnvironItemCreateAction`

NewEnvironItemCreateActionWithDefaults instantiates a new EnvironItemCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *EnvironItemCreateAction) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EnvironItemCreateAction) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EnvironItemCreateAction) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetName

`func (o *EnvironItemCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironItemCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironItemCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EnvironItemCreateAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironItemCreateAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironItemCreateAction) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *EnvironItemCreateAction) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *EnvironItemCreateAction) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *EnvironItemCreateAction) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *EnvironItemCreateAction) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvironItemCreateAction) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironItemCreateAction) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironItemCreateAction) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


