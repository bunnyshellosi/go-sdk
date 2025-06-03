# EnvironmentPartialDeployAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedDependencies** | Pointer to **string** |  | [optional] [default to "none"]
**QueueIfSomethingInProgress** | Pointer to **bool** |  | [optional] 
**IsPartial** | Pointer to **bool** |  | [optional] 
**Components** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnvironmentPartialDeployAction

`func NewEnvironmentPartialDeployAction() *EnvironmentPartialDeployAction`

NewEnvironmentPartialDeployAction instantiates a new EnvironmentPartialDeployAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentPartialDeployActionWithDefaults

`func NewEnvironmentPartialDeployActionWithDefaults() *EnvironmentPartialDeployAction`

NewEnvironmentPartialDeployActionWithDefaults instantiates a new EnvironmentPartialDeployAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedDependencies

`func (o *EnvironmentPartialDeployAction) GetIncludedDependencies() string`

GetIncludedDependencies returns the IncludedDependencies field if non-nil, zero value otherwise.

### GetIncludedDependenciesOk

`func (o *EnvironmentPartialDeployAction) GetIncludedDependenciesOk() (*string, bool)`

GetIncludedDependenciesOk returns a tuple with the IncludedDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedDependencies

`func (o *EnvironmentPartialDeployAction) SetIncludedDependencies(v string)`

SetIncludedDependencies sets IncludedDependencies field to given value.

### HasIncludedDependencies

`func (o *EnvironmentPartialDeployAction) HasIncludedDependencies() bool`

HasIncludedDependencies returns a boolean if a field has been set.

### GetQueueIfSomethingInProgress

`func (o *EnvironmentPartialDeployAction) GetQueueIfSomethingInProgress() bool`

GetQueueIfSomethingInProgress returns the QueueIfSomethingInProgress field if non-nil, zero value otherwise.

### GetQueueIfSomethingInProgressOk

`func (o *EnvironmentPartialDeployAction) GetQueueIfSomethingInProgressOk() (*bool, bool)`

GetQueueIfSomethingInProgressOk returns a tuple with the QueueIfSomethingInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueIfSomethingInProgress

`func (o *EnvironmentPartialDeployAction) SetQueueIfSomethingInProgress(v bool)`

SetQueueIfSomethingInProgress sets QueueIfSomethingInProgress field to given value.

### HasQueueIfSomethingInProgress

`func (o *EnvironmentPartialDeployAction) HasQueueIfSomethingInProgress() bool`

HasQueueIfSomethingInProgress returns a boolean if a field has been set.

### GetIsPartial

`func (o *EnvironmentPartialDeployAction) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *EnvironmentPartialDeployAction) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *EnvironmentPartialDeployAction) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *EnvironmentPartialDeployAction) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetComponents

`func (o *EnvironmentPartialDeployAction) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *EnvironmentPartialDeployAction) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *EnvironmentPartialDeployAction) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *EnvironmentPartialDeployAction) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


