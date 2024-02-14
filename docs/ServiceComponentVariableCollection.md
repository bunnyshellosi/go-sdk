# ServiceComponentVariableCollection

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

### NewServiceComponentVariableCollection

`func NewServiceComponentVariableCollection() *ServiceComponentVariableCollection`

NewServiceComponentVariableCollection instantiates a new ServiceComponentVariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceComponentVariableCollectionWithDefaults

`func NewServiceComponentVariableCollectionWithDefaults() *ServiceComponentVariableCollection`

NewServiceComponentVariableCollectionWithDefaults instantiates a new ServiceComponentVariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceComponentVariableCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceComponentVariableCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceComponentVariableCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceComponentVariableCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceComponentVariableCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceComponentVariableCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceComponentVariableCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceComponentVariableCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ServiceComponentVariableCollection) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceComponentVariableCollection) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceComponentVariableCollection) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceComponentVariableCollection) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ServiceComponentVariableCollection) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ServiceComponentVariableCollection) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *ServiceComponentVariableCollection) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceComponentVariableCollection) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceComponentVariableCollection) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ServiceComponentVariableCollection) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceComponent

`func (o *ServiceComponentVariableCollection) GetServiceComponent() string`

GetServiceComponent returns the ServiceComponent field if non-nil, zero value otherwise.

### GetServiceComponentOk

`func (o *ServiceComponentVariableCollection) GetServiceComponentOk() (*string, bool)`

GetServiceComponentOk returns a tuple with the ServiceComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceComponent

`func (o *ServiceComponentVariableCollection) SetServiceComponent(v string)`

SetServiceComponent sets ServiceComponent field to given value.

### HasServiceComponent

`func (o *ServiceComponentVariableCollection) HasServiceComponent() bool`

HasServiceComponent returns a boolean if a field has been set.

### GetProject

`func (o *ServiceComponentVariableCollection) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceComponentVariableCollection) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceComponentVariableCollection) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceComponentVariableCollection) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironment

`func (o *ServiceComponentVariableCollection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceComponentVariableCollection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceComponentVariableCollection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ServiceComponentVariableCollection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetOrganization

`func (o *ServiceComponentVariableCollection) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServiceComponentVariableCollection) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServiceComponentVariableCollection) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServiceComponentVariableCollection) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


