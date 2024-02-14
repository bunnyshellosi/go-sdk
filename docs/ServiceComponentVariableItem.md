# ServiceComponentVariableItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Component variable identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Component variable name. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | Component variable value. | [optional] [readonly] 
**Secret** | Pointer to **bool** | Component variable marked as secret. | [optional] [readonly] 
**ServiceComponent** | Pointer to **string** | Service component identifier | [optional] [readonly] 
**Project** | Pointer to **string** | Project identifier. | [optional] [readonly] 
**Environment** | Pointer to **string** | Environment identifier. | [optional] [readonly] 
**Organization** | Pointer to **string** | Organization identifier. | [optional] [readonly] 

## Methods

### NewServiceComponentVariableItem

`func NewServiceComponentVariableItem() *ServiceComponentVariableItem`

NewServiceComponentVariableItem instantiates a new ServiceComponentVariableItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceComponentVariableItemWithDefaults

`func NewServiceComponentVariableItemWithDefaults() *ServiceComponentVariableItem`

NewServiceComponentVariableItemWithDefaults instantiates a new ServiceComponentVariableItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceComponentVariableItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceComponentVariableItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceComponentVariableItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceComponentVariableItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceComponentVariableItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceComponentVariableItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceComponentVariableItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceComponentVariableItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ServiceComponentVariableItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceComponentVariableItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceComponentVariableItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceComponentVariableItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ServiceComponentVariableItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ServiceComponentVariableItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *ServiceComponentVariableItem) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceComponentVariableItem) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceComponentVariableItem) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ServiceComponentVariableItem) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceComponent

`func (o *ServiceComponentVariableItem) GetServiceComponent() string`

GetServiceComponent returns the ServiceComponent field if non-nil, zero value otherwise.

### GetServiceComponentOk

`func (o *ServiceComponentVariableItem) GetServiceComponentOk() (*string, bool)`

GetServiceComponentOk returns a tuple with the ServiceComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceComponent

`func (o *ServiceComponentVariableItem) SetServiceComponent(v string)`

SetServiceComponent sets ServiceComponent field to given value.

### HasServiceComponent

`func (o *ServiceComponentVariableItem) HasServiceComponent() bool`

HasServiceComponent returns a boolean if a field has been set.

### GetProject

`func (o *ServiceComponentVariableItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceComponentVariableItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceComponentVariableItem) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceComponentVariableItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironment

`func (o *ServiceComponentVariableItem) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceComponentVariableItem) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceComponentVariableItem) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ServiceComponentVariableItem) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetOrganization

`func (o *ServiceComponentVariableItem) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServiceComponentVariableItem) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServiceComponentVariableItem) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServiceComponentVariableItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


