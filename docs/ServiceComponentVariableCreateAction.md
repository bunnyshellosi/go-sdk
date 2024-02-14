# ServiceComponentVariableCreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsSecret** | Pointer to **NullableBool** |  | [optional] 
**ServiceComponent** | **string** |  | 

## Methods

### NewServiceComponentVariableCreateAction

`func NewServiceComponentVariableCreateAction(name string, value string, serviceComponent string, ) *ServiceComponentVariableCreateAction`

NewServiceComponentVariableCreateAction instantiates a new ServiceComponentVariableCreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceComponentVariableCreateActionWithDefaults

`func NewServiceComponentVariableCreateActionWithDefaults() *ServiceComponentVariableCreateAction`

NewServiceComponentVariableCreateActionWithDefaults instantiates a new ServiceComponentVariableCreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceComponentVariableCreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceComponentVariableCreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceComponentVariableCreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ServiceComponentVariableCreateAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceComponentVariableCreateAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceComponentVariableCreateAction) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecret

`func (o *ServiceComponentVariableCreateAction) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *ServiceComponentVariableCreateAction) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *ServiceComponentVariableCreateAction) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *ServiceComponentVariableCreateAction) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### SetIsSecretNil

`func (o *ServiceComponentVariableCreateAction) SetIsSecretNil(b bool)`

 SetIsSecretNil sets the value for IsSecret to be an explicit nil

### UnsetIsSecret
`func (o *ServiceComponentVariableCreateAction) UnsetIsSecret()`

UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
### GetServiceComponent

`func (o *ServiceComponentVariableCreateAction) GetServiceComponent() string`

GetServiceComponent returns the ServiceComponent field if non-nil, zero value otherwise.

### GetServiceComponentOk

`func (o *ServiceComponentVariableCreateAction) GetServiceComponentOk() (*string, bool)`

GetServiceComponentOk returns a tuple with the ServiceComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceComponent

`func (o *ServiceComponentVariableCreateAction) SetServiceComponent(v string)`

SetServiceComponent sets ServiceComponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


